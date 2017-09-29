package api

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	"version.uuzu.com/zhuhuipeng/djob/log"
	pb "version.uuzu.com/zhuhuipeng/djob/message"
)

type Backend interface {
	JobModify(*pb.Job) (*pb.Job, error)
	JobInfo(name, region string) (*pb.Job, error)
	JobDelete(name, region string) (*pb.Job, error)
	JobList(region string) ([]*pb.Job, error)
	JobRun(name, region string) (*pb.Execution, error)
	JobStatus(name, region string) (*pb.JobStatus, error)
}

type KayPair struct {
	Key  string
	Cert string
}

var jsonContentType = []string{"application/json; charset=utf-8"}

//jsonString implement output a protobuf obj as a json string and json content type
type pbjson struct {
	data interface{}
}

func (j pbjson) Render(w http.ResponseWriter) error {
	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{EmitDefaults: true}
	if err := marshaler.Marshal(&buf, j.data.(proto.Message)); err != nil {
		return err
	}
	w.Write(buf.Bytes())
	return nil
}

func (j pbjson) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = jsonContentType
	}
}

type jsonpbBinding struct{}

func (jsonpbBinding) Name() string {
	return "jsonpb"
}

func (jsonpbBinding) Bind(req *http.Request, obj interface{}) error {
	jsondec := json.NewDecoder(req.Body)
	unmarshaler := &jsonpb.Unmarshaler{AllowUnknownFields: false}
	if err := unmarshaler.UnmarshalNext(jsondec, obj.(proto.Message)); err != nil {
		return err
	}

	return nil
}

type APIServer struct {
	bindIP    string
	bindPort  int
	tokenList map[string]string
	backend   Backend
	loger     *logrus.Entry
	keyPair   *KayPair
	tls       bool
	router    *gin.Engine
	server    *http.Server
}

func NewAPIServer(ip string, port int,
	tokens map[string]string, tls bool, pair *KayPair, backend Backend) (*APIServer, error) {
	if len(tokens) == 0 {
		return nil, errors.New("Have no tokens")
	}
	// reverse tokens
	n := make(map[string]string)
	for k, v := range tokens {
		if _, exist := n[v]; exist {
			return nil, errors.New("Have repetition token")
		}
		n[v] = k
	}
	return &APIServer{
		bindIP:    ip,
		bindPort:  port,
		loger:     log.Loger,
		tokenList: n,
		tls:       tls,
		keyPair:   pair,
		backend:   backend,
	}, nil
}

func (a *APIServer) prepareGin() *gin.Engine {
	r := gin.New()
	r.Use(a.logMiddleware())
	r.Use(a.tokenAuthMiddleware())
	r.Use(gin.Recovery())
	if a.tls {
		r.Use(a.tlsHeaderMiddleware())
	}
	web := r.Group("/web")
	web.Use(gzip.Gzip(gzip.DefaultCompression))

	jobAPI := web.Group("/job")
	jobAPI.POST("/", a.modJob)
	jobAPI.GET("/:region", a.getJobList)
	jobAPI.GET("/:region/:name", a.getJob)
	jobAPI.DELETE("/:region/:name", a.deleteJob)
	jobAPI.GET("/:region/:name/run", a.runJob)
	jobAPI.GET("/:region/:name/status", a.getJobStatus)

	//executionAPI := web.Group("/execution")
	//executionAPI.GET("/:region", a.getAllExecutions)
	//executionAPI.GET("/:region/:name", a.getJobAllExecutions)
	//executionAPI.GET("/:region/:name/:group", a.getExecution)
	return r
}

func (a *APIServer) getJobStatus(c *gin.Context) {
	name := c.Params.ByName("name")
	region := c.Params.ByName("region")
	js, err := a.backend.JobStatus(name, region)
	if err != nil {
		a.respondWithError(http.StatusBadRequest, &pb.RespStatus{Status: http.StatusBadRequest, Message: err.Error()}, c)
		return
	}
	resp := pb.RespStatus{
		Status:  0,
		Message: "succeed",
		Data:    js,
	}
	c.Render(http.StatusOK, pbjson{data: &resp})
}

func (a *APIServer) runJob(c *gin.Context) {
	name := c.Params.ByName("name")
	region := c.Params.ByName("region")
	ex, err := a.backend.JobRun(name, region)
	if err != nil {
		a.respondWithError(http.StatusBadRequest, &pb.RespJob{Status: http.StatusBadRequest, Message: err.Error()}, c)
		return
	}
	resp := pb.RespExec{
		Status:  0,
		Message: "succeed",
		Data:    []*pb.Execution{ex},
	}
	c.Render(http.StatusOK, pbjson{data: &resp})
}

func (a *APIServer) deleteJob(c *gin.Context) {
	name := c.Params.ByName("name")
	region := c.Params.ByName("region")
	job, err := a.backend.JobDelete(name, region)
	if err != nil {
		a.respondWithError(http.StatusBadRequest, &pb.RespJob{Status: http.StatusBadRequest, Message: err.Error()}, c)
		return
	}
	resp := pb.RespJob{
		Status:  0,
		Message: "succeed",
		Data:    []*pb.Job{job},
	}
	c.Render(http.StatusOK, pbjson{data: &resp})
}

// TODO: add a data filter
func (a *APIServer) getJobList(c *gin.Context) {
	region := c.Params.ByName("region")
	jobs, err := a.backend.JobList(region)
	if err != nil {
		a.respondWithError(http.StatusNotFound, &pb.RespJob{Status: http.StatusNotFound, Message: err.Error()}, c)
		return
	}
	resp := pb.RespJob{
		Status:  0,
		Message: "succeed",
		Data:    jobs,
	}
	c.Render(http.StatusOK, pbjson{data: &resp})
}
func (a *APIServer) getJob(c *gin.Context) {
	name := c.Params.ByName("name")
	region := c.Params.ByName("region")
	job, err := a.backend.JobInfo(name, region)
	if err != nil {
		a.respondWithError(http.StatusBadRequest, &pb.RespJob{Status: http.StatusBadRequest, Message: err.Error()}, c)
		return
	}
	resp := pb.RespJob{
		Status:  0,
		Message: "succeed",
		Data:    []*pb.Job{job},
	}
	c.Render(http.StatusOK, pbjson{data: &resp})
}

func (a *APIServer) modJob(c *gin.Context) {
	var (
		job  pb.Job
		resp pb.RespJob
		err  error
	)
	err = c.MustBindWith(&job, jsonpbBinding{})
	if err != nil {
		a.respondWithError(http.StatusBadRequest, &pb.RespJob{Status: http.StatusBadRequest, Message: err.Error()}, c)
		return
	}
	rj, err := a.backend.JobModify(&job)
	if err != nil {
		a.respondWithError(http.StatusBadRequest, &pb.RespJob{Status: http.StatusBadRequest, Message: err.Error()}, c)
		return
	}
	resp.Status = 0
	resp.Message = "succeed"
	resp.Data = append(resp.Data, rj)

	c.Render(http.StatusOK, pbjson{data: &resp})
}

func (a *APIServer) tlsHeaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		c.Writer.Header().Set("Strict-Transport-Security",
			"max-age=63072000; includeSubDomains")
	}
}

func (a *APIServer) logMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		entry := a.loger.WithFields(logrus.Fields{
			"client_ip":  c.ClientIP(),
			"method":     c.Request.Method,
			"path":       path,
			"status":     c.Writer.Status(),
			"latency":    latency,
			"user-agent": c.Request.UserAgent(),
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info()
		}
	}
}

func (a *APIServer) respondWithError(code int, pb interface{}, c *gin.Context) {
	c.Render(code, pbjson{data: pb.(proto.Message)})
	c.AbortWithStatus(code)
}

func (a *APIServer) tokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("X-Auth-Token")

		if token == "" {
			a.respondWithError(http.StatusUnauthorized, &pb.RespJob{Status: http.StatusUnauthorized, Message: "API token required"}, c)
			return
		}
		if _, exist := a.tokenList[token]; exist {
			c.Next()
		} else {
			a.respondWithError(http.StatusUnauthorized, &pb.RespJob{Status: http.StatusUnauthorized, Message: "API token Error"}, c)
			return
		}
	}
}

func (a *APIServer) Run() {
	r := a.prepareGin()
	a.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", a.bindIP, a.bindPort),
		Handler: r,
	}
	if a.tls {
		a.server.TLSConfig = &tls.Config{
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP256, tls.CurveP384, tls.CurveP521},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			},
		}
		a.server.TLSNextProto = make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0)
		go a.server.ListenAndServeTLS(a.keyPair.Cert, a.keyPair.Key)
		return
	}
	go a.server.ListenAndServe()
	return
}

func (a *APIServer) Stop(wait time.Duration) error {
	a.loger.Infof("API-server: shutdown in %d second", wait)
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	if err := a.server.Shutdown(ctx); err != nil {
		return err
	}
	a.loger.Info("API-server: bye-bye")
	return nil
}
