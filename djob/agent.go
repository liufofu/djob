package djob

import (
	"version.uuzu.com/zhuhuipeng/djob/rpc"
	pb "version.uuzu.com/zhuhuipeng/djob/message"
	"time"

	"errors"
	"fmt"
	"sync"

	"github.com/Sirupsen/logrus"
	"github.com/docker/libkv/store"
	"github.com/hashicorp/memberlist"
	"github.com/hashicorp/serf/serf"
	"version.uuzu.com/zhuhuipeng/djob/scheduler"
	"version.uuzu.com/zhuhuipeng/djob/web/api"
)

var (
	ErrLockTimeout = errors.New("locking timeout")
	LockTimeout    = 2 * time.Second
)

const gracefulTime = 5 * time.Second

type Agent struct {
	jobLockers  map[string]store.Locker
	config      *Config
	serf        *serf.Serf
	eventCh     chan serf.Event
	ready       bool
	rpcServer   *rpc.RpcServer
	rpcClient   *rpc.RpcClient
	store       *KVStore
	memberCache map[string]map[string]string
	mutex       sync.Mutex
	scheduler   *scheduler.Scheduler
	apiServer   *api.APIServer
	version     string
	runJobCh    chan *pb.Job
}

func (a *Agent) setupSerf() *serf.Serf {
	encryptKey, err := a.config.EncryptKey()
	if err != nil {
		Log.Fatal(err)
		return nil
	}

	serfConfig := serf.DefaultConfig()

	//noinspection GoBinaryAndUnaryExpressionTypesCompatibility
	serfConfig.MemberlistConfig = memberlist.DefaultWANConfig()

	serfConfig.MemberlistConfig.BindAddr = a.config.SerfBindIP
	serfConfig.MemberlistConfig.BindPort = a.config.SerfBindPort
	serfConfig.MemberlistConfig.AdvertiseAddr = a.config.SerfAdvertiseIP
	serfConfig.MemberlistConfig.AdvertisePort = a.config.SerfAdvertisePort
	serfConfig.MemberlistConfig.SecretKey = encryptKey
	serfConfig.NodeName = a.config.Nodename
	serfConfig.Tags = a.config.Tags
	serfConfig.SnapshotPath = a.config.SerfSnapshotPath
	serfConfig.CoalescePeriod = 3 * time.Second
	serfConfig.QuiescentPeriod = time.Second
	serfConfig.UserCoalescePeriod = 3 * time.Second
	serfConfig.UserQuiescentPeriod = time.Second
	serfConfig.EnableNameConflictResolution = true

	//	a.eventCh = make(chan serf.Event, 64)
	serfConfig.EventCh = a.eventCh

	Log.Info("agent: Djob agent starting")

	s, err := serf.Create(serfConfig)
	if err != nil {
		Log.Fatal(err)
		return nil
	}
	//a.memberCache = make(map[string]map[string]string)

	return s
}

// serfJion let serf intence jion a serf clust
func (a *Agent) serfJion(addrs []string, replay bool) (n int, err error) {
	Log.Infof("agent: joining: %v replay: %v", addrs, replay)
	ignoreOld := !replay
	n, err = a.serf.Join(addrs, ignoreOld)
	if n > 0 {
		Log.Infof("agent: joined: %d nodes", n)
	}
	if err != nil {
		Log.Warnf("agent: error joining: %v", err)
	}
	return
}

func (a *Agent) lockJob(jobName string) (store.Locker, error) {

	//reNewCh := make(chan struct{})

	lockkey := fmt.Sprintf("%s/job_locks/%s", a.store.keyspace, jobName)

	//l, err := a.store.Client.NewLock(lockkey, &store.LockOptions{RenewLock:reNewCh})
	l, err := a.store.Client.NewLock(lockkey, &store.LockOptions{})
	if err != nil {
		Log.WithField("jobName", jobName).WithError(err).Fatal("agent: New lock failed")
	}

	errCh := make(chan error)
	freeCh := make(chan struct{})
	timeoutCh := time.After(LockTimeout)
	stoplockingCh := make(chan struct{})

	go func() {
		_, err = l.Lock(stoplockingCh)
		if err != nil {
			errCh <- err
			return
		}
		freeCh <- struct{}{}
	}()

	select {
	case <-freeCh:
		return l, nil
	case err := <-errCh:
		return nil, err
	case <-timeoutCh:
		stoplockingCh <- struct{}{}
		return nil, ErrLockTimeout
	}
}

func (a *Agent) mainLoop() {
	serfShutdownCh := a.serf.ShutdownCh()
	Log.Info("agent: Listen for event")
	for {

		select {
		// handle serf event
		case e := <-a.eventCh:
			Log.WithFields(logrus.Fields{
				"event": e.String(),
			}).Debug("agent: Received event")

			if memberevent, ok := e.(serf.MemberEvent); ok {
				var memberNames []string
				for _, member := range memberevent.Members {
					memberNames = append(memberNames, member.Name)
				}
				Log.WithFields(logrus.Fields{
					"node":    a.config.Nodename,
					"members": memberNames,
					"event":   e.EventType(),
				}).Debug("agent: Member event got")

				/* go a.handleMemberCache(memberevent.Type, memberevent.Members) */
			}

			// handle custom query event
			if e.EventType() == serf.EventQuery {
				query := e.(*serf.Query)

				switch qname := query.Name; qname {
				case QueryNewJob:
					if a.config.Server {
						Log.WithFields(logrus.Fields{
							"query":   query.Name,
							"payload": string(query.Payload),
							"at":      query.LTime,
						}).Debug("agent: Server receive a add new job event")

						go a.receiveNewJobQuery(query)
					}
				case QueryRunJob:
					Log.WithFields(logrus.Fields{
						"query":   query.Name,
						"payload": string(query.Payload),
						"at":      query.LTime,
					}).Debug("agent: Running job")
				case QueryRPCConfig:
					if a.config.Server {
						Log.WithFields(logrus.Fields{
							"query":   query.Name,
							"payload": string(query.Payload),
							"at":      query.LTime,
						}).Debug("agent: Server receive a rpc config query")
					}
				default:
					Log.Warn("agent: get a unknow message")
					Log.WithFields(logrus.Fields{
						"query":   query.Name,
						"payload": string(query.Payload),
						"at":      query.LTime,
					}).Debug("agent: get a unknow message")
				}
			}

		case <-serfShutdownCh:
			Log.Warn("agent: Serf shutdown detected, quitting")
			return
		}
	}
}

// Run func start a Agent process
func (a *Agent) Run() {
	var err error
	if err != nil {
		Log.Fatalln(err)
	}
	if a.serf = a.setupSerf(); a.serf == nil {
		Log.Fatalln("Start serf failed!")
	}
	a.serfJion(a.config.SerfJoin, true)
	// TODO: add prometheus support
	// start prometheus client
	var tls rpc.TlsOpt
	var keyPair api.KayPair
	if a.config.RPCTls {
		tls = rpc.TlsOpt{
			CertFile: a.config.CertFile,
			KeyFile:  a.config.KeyFile,
			CaFile:   a.config.CAFile,
		}
		keyPair = api.KayPair{
			Key:  a.config.KeyFile,
			Cert: a.config.CertFile,
		}
	}

	if a.config.Server {
		a.scheduler = scheduler.New(a.runJobCh)
		a.scheduler.Start()
		a.store, err = NewStore(a.config.JobStore, a.config.JobStoreServers, a.config.JobStoreKeyspace)
		if err != nil {
			Log.WithFields(logrus.Fields{
				"backend":  a.config.JobStore,
				"servers":  a.config.JobStoreServers,
				"keyspace": a.config.JobStoreKeyspace,
			}).Debug("Agent: Connect Backend Failed")
			Log.Fatalf("Agent: Connent Backend Failed: %s", err)
		}
		// run rpc server
		a.rpcServer = rpc.NewRPCServer(a.config.RPCBindIP, a.config.RPCBindPort, a, &tls)
		go func() {
			if err := a.rpcServer.Run(); err != nil {
				Log.Fatalf("Agent: Start RPC Srever Failed: %s", err)
			}
			Log.Info("Agent: RPC Server started")
		}()

		a.loadAllJob(a.config.Region)

		a.apiServer, err = api.NewAPIServer(a.config.APIBindIP, a.config.APIBindPort, Log, make(map[string]string), a.config.RPCTls, &keyPair)
		if err != nil {
			Log.Fatalf("Agent: New API Server Failed: %s", err)
		}
		go func() {
			if err := a.apiServer.Run(); err != nil {
				Log.Fatalf("Agent: Start API Server Failed: %s", err)
			}
			Log.Info("Agent: API Server started")
		}()

	}

	a.mainLoop()
}

func (a *Agent) loadAllJob(region string) {

}

func (a *Agent) Reload(args []string) {
	newConf, err := newConfig(args, a.version)
	if err != nil {
		Log.Warn(err)
		return
	}
	a.config = newConf

	a.serf.SetTags(a.config.Tags)
}

func (a *Agent) Stop(graceful bool) int {

	if !graceful {
		return 0
	}

	gracefulCh := make(chan struct{})

	Log.Info("agent: Gracefully shutting down agent...")
	go func() {
		var wg sync.WaitGroup
		if a.config.Server {
			// shutdown scheduler and unlock all jobs
			go func() {
				wg.Add(1)
				a.scheduler.Stop()
				for _, locker := range a.jobLockers {
					locker.Unlock()
				}
				wg.Done()
			}()
			// graceful shutdown rpc server
			go func() {
				wg.Add(1)
				a.rpcServer.Shutdown(gracefulTime)
				wg.Done()
			}()
			// gracefull shutdown api server
			go func() {
				wg.Add(1)
				if err := a.apiServer.Stop(gracefulTime); err != nil {
					Log.Errorf("Agent:Graceful shutdown Api server failed: %s", err)
				}
				wg.Done()
			}()
		}
		go func() {
			wg.Add(1)
			if err := a.serf.Leave(); err != nil {
				Log.Errorf("Agent: Graceful shutdown down serf failed: %s", err)
			}
			a.serf.Shutdown()
			wg.Done()
		}()
		wg.Wait()
		gracefulCh <- struct{}{}
	}()

	select {
	case <-time.After(gracefulTime):
		return 1
	case <-gracefulCh:
		return 0
	}
}

func New(args []string, version string) *Agent {
	config, err := newConfig(args, version)
	if err != nil {
		Log.Fatalln(err)
	}

	return &Agent{
		jobLockers:  make(map[string]store.Locker),
		eventCh:     make(chan serf.Event, 64),
		memberCache: make(map[string]map[string]string),
		runJobCh:    make(chan *pb.Job),
		config:      config,
		version:     version,
	}

}

