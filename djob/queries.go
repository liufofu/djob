package djob

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/gogo/protobuf/proto"
	"github.com/hashicorp/serf/serf"
	"math/rand"
	pb "version.uuzu.com/zhuhuipeng/djob/message"
)

const (
	QueryNewJob    = "job:new"
	QueryRunJob    = "job:run"
	QueryRPCConfig = "rpc:config"
	QueryJobCount  = "job:count"
	QueryJobDelete = "job:delete"
)

func (a *Agent) sendJobDeleteQuery(name, region, nodeName string) (*pb.Result, error) {
	params := &serf.QueryParam{
		FilterNodes: []string{nodeName},
		FilterTags: map[string]string{
			"server": "true",
			"region": region,
		},
		RequestAck: true,
	}

	qp := &pb.JobQueryParams{
		Name:           name,
		Region:         region,
		SourceNodeName: a.config.Nodename,
	}
	qpPb, _ := proto.Marshal(qp)

	Log.WithFields(logrus.Fields{
		"query_name": QueryJobDelete,
		"job_name":   name,
		"job_region": region,
		"payload":    qp.String(),
	}).Debug("Agent: Sending query")

	qr, err := a.serf.Query(QueryJobDelete, qpPb, params)
	if err != nil {
		Log.WithField("query", QueryJobDelete).WithError(err).Fatal("Agent: Sending query error")
	}
	defer qr.Close()

	ackCh := qr.AckCh()
	respCh := qr.ResponseCh()
	var payloadPb []byte
	for {
		if len(payloadPb) != 0 {
			break
		}
		select {
		case ack, ok := <-ackCh:
			if ok {
				Log.WithFields(logrus.Fields{
					"query": QueryJobDelete,
					"from":  ack,
				}).Debug("Agent: Received ack")
			}

		case resp, ok := <-respCh:
			if ok {
				payloadPb = resp.Payload
				Log.WithFields(logrus.Fields{
					"query":   QueryJobDelete,
					"payload": string(resp.Payload),
				}).Debug("Agent: Received response")
			}

		}
	}

	var result pb.Result
	if err := proto.Unmarshal(payloadPb, &result); err != nil {
		Log.WithError(err).Error("Agent: Decode respond failed")
		return nil, err
	}

	return &result, nil
}

func (a *Agent) receiveJobDeleteQuery(query *serf.Query) {
	var params *pb.JobQueryParams
	if err := proto.Unmarshal(query.Payload, params); err != nil {
		Log.WithFields(logrus.Fields{
			"query":   query.Name,
			"payload": string(query.Payload),
		}).WithError(err).Error("Agent: Decode payload failed")
		rb, _ := genrateResultPb(1, "Decode failed")
		query.Respond(rb)
		return
	}
	if params.Region != a.config.Region {
		Log.WithFields(logrus.Fields{
			"name":   params.Name,
			"region": params.Region,
		}).Debug("Agent: receive a job from other region")
		rb, _ := genrateResultPb(2, "region error")
		query.Respond(rb)
		return
	}
	if !a.haveIt(params.Name) {
		Log.WithFields(logrus.Fields{
			"query": query.Name,
			"name":  params.Name,
		}).Error("Agent: have no job locker")
		rb, _ := genrateResultPb(10, "have no job locker")
		query.Respond(rb)
		return
	}
	_, err := a.store.DeleteJob(params.Name, params.Region)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"query": query.Name,
		}).WithError(err).Error("Agent: Delete job failed")
		rb, _ := genrateResultPb(7, "Delete job failed")
		query.Respond(rb)
		return
	}
	if err := a.deleteLocker(params.Name); err != nil {
		Log.WithError(err).Error("Agent: Unlock job failed")
		rb, _ := genrateResultPb(11, "Unlock job failed")
		query.Respond(rb)
		return
	}
	rb, _ := genrateResultPb(0, "Succeed")
	if err := query.Respond(rb); err != nil {
		Log.WithError(err).Fatal("Agent: serf query Respond error")
	}
	return
}

func (a *Agent) sendJobCountQuery(region string) ([]*pb.JobCountResp, error) {

	params, err := a.createSerfQueryParam(fmt.Sprintf("server=='true'&&region=='%s'", region))
	if err != nil {
		return nil, err
	}
	params.FilterTags = map[string]string{
		"server": "true",
		"region": region,
	}
	qr, err := a.serf.Query(QueryJobCount, nil, params)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"query": QueryJobCount,
			"error": err,
		}).Fatal("Agent: Error sending serf query")
	}
	defer qr.Close()

	ackCh := qr.AckCh()
	respCh := qr.ResponseCh()
	var payloads [][]byte
	for {
		if len(payloads) == len(params.FilterNodes) {
			break
		}
		select {

		case ack, ok := <-ackCh:
			if ok {
				Log.WithFields(logrus.Fields{
					"query": QueryJobCount,
					"from":  ack,
				}).Debug("Agent: Received ack")
			}
		case resp, ok := <-respCh:
			if ok {
				Log.WithFields(logrus.Fields{
					"query": QueryJobCount,
					"resp":  string(resp.Payload),
				}).Debug("Agent: Received resp")
				payloads = append(payloads, resp.Payload)
			}
		}
	}

	var r []*pb.JobCountResp
	for _, v := range payloads {
		var p pb.JobCountResp
		proto.Unmarshal(v, &p)
		r = append(r, &p)
	}
	return r, nil
}

func (a *Agent) receiveJobCountQuery(query *serf.Query) {
	resp := pb.JobCountResp{
		Name:  a.config.Nodename,
		Count: int64(a.scheduler.JobCount()),
	}
	respPb, _ := proto.Marshal(&resp)
	Log.WithFields(logrus.Fields{
		"name":   resp.Name,
		"count":  resp.Count,
		"respPb": string(respPb),
	}).Debug("Agent: respone of job:count")
	if err := query.Respond(respPb); err != nil {
		Log.WithError(err).Fatal("Agent: serf quert Responed failed")
	}
	return
}

// sendNreJobQuery func used to notice a server there is a now job need to be add
func (a *Agent) sendNewJobQuery(name, region, nodeName string) (*pb.Result, error) {

	params := &serf.QueryParam{
		FilterNodes: []string{nodeName},
		FilterTags: map[string]string{
			"server": "true",
			"region": region,
		},
		RequestAck: true,
	}

	qp := &pb.JobQueryParams{
		Name:           name,
		Region:         region,
		SourceNodeName: a.config.Nodename,
	}
	qpPb, _ := proto.Marshal(qp)

	Log.WithFields(logrus.Fields{
		"query_name": QueryNewJob,
		"job_name":   name,
		"job_region": region,
		"payload":    qp.String(),
	}).Debug("Agent: Sending query")

	qr, err := a.serf.Query(QueryNewJob, qpPb, params)
	if err != nil {
		Log.WithField("query", QueryNewJob).WithError(err).Fatal("Agent: Sending query error")
	}
	defer qr.Close()

	ackCh := qr.AckCh()
	respCh := qr.ResponseCh()
	var payloadPb []byte
	for {
		if len(payloadPb) != 0 {
			break
		}
		select {
		case ack, ok := <-ackCh:
			if ok {
				Log.WithFields(logrus.Fields{
					"query": QueryNewJob,
					"from":  ack,
				}).Debug("Agent: Received ack")
			}

		case resp, ok := <-respCh:
			if ok {
				Log.WithFields(logrus.Fields{
					"query":   QueryNewJob,
					"payload": string(resp.Payload),
				}).Debug("Agent: Received response")
				payloadPb = resp.Payload
			}

		}
	}

	var result pb.Result
	if err := proto.Unmarshal(payloadPb, &result); err != nil {
		Log.WithError(err).Error("Agent: Decode respond failed")
		return nil, err
	}

	return &result, nil
}

func (a *Agent) receiveNewJobQuery(query *serf.Query) {
	var params pb.JobQueryParams
	if err := proto.Unmarshal(query.Payload, &params); err != nil {
		Log.WithFields(logrus.Fields{
			"query":   query.Name,
			"payload": string(query.Payload),
		}).WithError(err).Error("Agent: Decode payload failed")
		rb, _ := genrateResultPb(1, "Decode payload failed")
		query.Respond(rb)
		return
	}

	if !a.haveIt(params.Name) {
		locker, err := a.lock(params.Name, params.Region, pb.Job{})
		if err != nil {
			Log.WithFields(logrus.Fields{
				"JobName": params.Name,
				"Region":  params.Region,
			}).WithError(err).Error("Agent: try lock a job failed")
			rb, _ := genrateResultPb(9, "try lock a job failed")
			query.Respond(rb)
			return
		}
		a.addLocker(params.Name, locker)
	}

	if params.Region != a.config.Region {
		Log.WithFields(logrus.Fields{
			"name":   params.Name,
			"region": params.Region,
		}).Debug("Agent: receive a job from other region")
		rb, _ := genrateResultPb(2, "region error")
		query.Respond(rb)
		return
	}
	ip, port, err := a.sendGetRPCConfigQuery(params.SourceNodeName, params.Region)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"SourceNodeName": params.SourceNodeName,
		}).WithError(err).Error("Agent: Get RPC config Failed")
		rb, _ := genrateResultPb(3, fmt.Sprintf("Get RPC config failed: %s", params.SourceNodeName))
		query.Respond(rb)
		return
	}

	rpcClient := a.newRPCClient(ip, port)
	defer rpcClient.Shutdown()

	job, err := rpcClient.GetJob(params.Name, params.Region)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"Name":    query.Name,
			"RPCAddr": fmt.Sprintf("%s:%d", ip, port),
		}).WithError(err).Error("Agent: RPC Client get job failed")
		rb, _ := genrateResultPb(4, "RPC Client get job failed")
		query.Respond(rb)
		return
	}

	job.SchedulerNodeName = a.config.Nodename

	if err := a.store.SetJob(job); err != nil {
		Log.WithFields(logrus.Fields{
			"JobName": job.Name,
			"Region":  job.Region,
		}).WithError(err).Error("Agent: Save job to store failed")
		rb, _ := genrateResultPb(5, "Save job to store failed")
		query.Respond(rb)
		return
	}

	a.scheduler.DeleteJob(job.Name)

	// add job
	err = a.scheduler.AddJob(job)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"job_name":  job.Name,
			"scheduler": job.Schedule,
		}).WithError(err).Error("Add Job to scheduler failed")
		rb, _ := genrateResultPb(6, "Add Job to scheduler failed")
		query.Respond(rb)
		return
	}
	Log.Infof("Agent: send job %s to newJobCh", job.Name)
	Log.WithFields(logrus.Fields{
		"query":   query.Name,
		"payload": string(query.Payload),
	}).Debug("Agent: send job to newJobCh")
	rb, _ := genrateResultPb(0, "succeed")
	if err := query.Respond(rb); err != nil {
		Log.WithError(err).Fatal("Agent: serf query Respond error")
	}
	return
}

func (a *Agent) sendRunJobQuery(ex *pb.Execution) {

	job, err := a.store.GetJob(ex.JobName, ex.Region)
	if err != nil {
		Log.WithError(err).Fatal("Agent GetJob failed")
	}

	exPb, err := proto.Marshal(ex)
	if err != nil {
		Log.WithError(err).Fatal("Agent: Encode failed")
	}

	var params *serf.QueryParam
	if ex.Retries < 1 {
		params, err = a.createSerfQueryParam(job.Expression)
	} else {
		params = &serf.QueryParam{
			FilterNodes: []string{ex.RunNodeName},
			RequestAck:  true,
		}
	}

	if err != nil {
		Log.WithFields(logrus.Fields{
			"Query": QueryRunJob,
			"stage": "prepare",
		}).WithError(err).Error("Agent: Create serf query param failed")
		return
	}

	Log.WithFields(logrus.Fields{
		"query_name": QueryNewJob,
		"job_name":   job.Name,
		"job_region": job.Region,
		"payload":    string(exPb),
	}).Debug("Agent: Sending query")

	qr, err := a.serf.Query(QueryRunJob, exPb, params)
	if err != nil {
		Log.WithField("query", QueryRunJob).WithError(err).Fatal("Agent: Sending query error")
	}
	defer qr.Close()

	ackCh := qr.AckCh()
	respCh := qr.ResponseCh()
	for !qr.Finished() {
		select {
		case ack, ok := <-ackCh:
			if ok {
				Log.WithFields(logrus.Fields{
					"query": QueryRunJob,
					"from":  ack,
				}).Debug("Agent: Received ack")
			}

		case resp, ok := <-respCh:
			if ok {
				Log.WithFields(logrus.Fields{
					"query":   QueryRunJob,
					"payload": string(resp.Payload),
				}).Debug("Agent: Received response")
			}

		}
	}
	Log.WithFields(logrus.Fields{
		"query": QueryRunJob,
	}).Debug("Agent: Done receiving acks and responses")

}

func (a *Agent) receiveRunJobQuery(query *serf.Query) {
	var ex pb.Execution
	if err := proto.Unmarshal(query.Payload, &ex); err != nil {
		Log.WithError(err).Fatal("Agent: Error decode query payload")
	}

	Log.WithFields(logrus.Fields{
		"job":    ex.JobName,
		"region": ex.Region,
		"group":  ex.Group,
	}).Info("Agent: Starting job")

	ip, port, err := a.sendGetRPCConfigQuery(ex.SchedulerNodeName, ex.Region)
	if err != nil {
		Log.WithError(err).Fatal("Agent: get rpc config failed")
	}

	rpcc := a.newRPCClient(ip, port)
	defer rpcc.Shutdown()

	job, err := rpcc.GetJob(ex.JobName, ex.Region)
	if err != nil {
		Log.WithError(err).Fatal("Agent: rpc call GetJob Failed")
	}
	logrus.WithFields(logrus.Fields{
		"job":    job.Name,
		"region": job.Region,
		"cmd":    job.Command,
	}).Debug("Agent: Got job by rpc call GetJob")

	go func() {
		if err := a.execJob(job, &ex); err != nil {
			Log.WithError(err).Error("Proc: Exec job Failed")
		}
		Log.WithFields(logrus.Fields{
			"job":    ex.JobName,
			"region": ex.Region,
			"group":  ex.Group,
		}).Info("Agent: Job done")
	}()
}

func (a *Agent) sendGetRPCConfigQuery(nodeName, region string) (string, int, error) {
	params, err := a.createSerfQueryParam(fmt.Sprintf("server=='true'&&region=='%s'", region))
	if err != nil {
		return "", 0, err
	}

	if nodeName != "" {
		params.FilterNodes = []string{nodeName}
	}

	qr, err := a.serf.Query(QueryRPCConfig, nil, params)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"query": QueryRPCConfig,
			"error": err,
		}).Fatal("Agent: Error sending serf query")
	}

	defer qr.Close()

	ackCh := qr.AckCh()
	respCh := qr.ResponseCh()
	var payloads [][]byte
	for {
		if len(payloads) == len(params.FilterNodes) {
			break
		}
		select {
		case ack, ok := <-ackCh:
			if ok {
				Log.WithFields(logrus.Fields{
					"query": QueryRPCConfig,
					"from":  ack,
				}).Debug("Agent: Received ack")
			}
		case resp, ok := <-respCh:
			if ok {
				payloads = append(payloads, resp.Payload)
				Log.WithFields(logrus.Fields{
					"query":   QueryRPCConfig,
					"from":    resp.From,
					"payload": string(resp.Payload),
				}).Debug("Agent: Received response")
			}
		}
	}
	Log.WithFields(logrus.Fields{
		"query": QueryRPCConfig,
	}).Debug("Agent: Received ack and response Done")

	var payload pb.GetRPCConfigResp
	if err := proto.Unmarshal(payloads[rand.Intn(len(payloads))], &payload); err != nil {
		Log.WithError(err).Error("Agent: Payload decode failed")
		return "", 0, nil
	}
	return payload.Ip, int(payload.Port), nil
}

func (a *Agent) receiveGetRPCConfigQuery(query *serf.Query) {
	resp := &pb.GetRPCConfigResp{
		Ip:   a.config.RPCBindIP,
		Port: int32(a.config.RPCBindPort),
	}
	respPb, _ := proto.Marshal(resp)
	if err := query.Respond(respPb); err != nil {
		Log.WithError(err).Fatal("Agent: serf query Respond error")
	}
	return
}
