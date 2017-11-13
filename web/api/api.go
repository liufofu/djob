/*
 * Copyright (c) 2017.  Harrison Zhu <wcg6121@gmail.com>
 * This file is part of djob <https://github.com/HZ89/djob>.
 *
 * djob is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * djob is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with djob.  If not, see <http://www.gnu.org/licenses/>.
 */

package api

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	"version.uuzu.com/zhuhuipeng/djob/errors"
	"version.uuzu.com/zhuhuipeng/djob/log"
	pb "version.uuzu.com/zhuhuipeng/djob/message"
)

type ApiController interface {
	ListRegions() ([]string, error)
	AddJob(*pb.Job) (*pb.Job, error)
	ModifyJob(*pb.Job) (*pb.Job, error)
	DeleteJob(*pb.Job) (*pb.Job, error)
	ListJob(name, region string) ([]*pb.Job, error)
	RunJob(name, region string) (*pb.Execution, error)
	GetStatus(name, region string) (*pb.JobStatus, error)
	ListExecutions(name, region string, group int64) ([]*pb.Execution, error)
	Search(interface{}, *pb.Search) ([]interface{}, int, error)
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
	mc        ApiController
	loger     *logrus.Entry
	keyPair   *KayPair
	tls       bool
	router    *gin.Engine
	server    *http.Server
}

func NewAPIServer(ip string, port int,
	tokens map[string]string, tls bool, pair *KayPair, backend ApiController) (*APIServer, error) {
	if len(tokens) == 0 {
		return nil, errors.ErrMissApiToken
	}
	// reverse tokens
	n := make(map[string]string)
	for k, v := range tokens {
		if _, exist := n[v]; exist {
			return nil, errors.ErrRepetionToken
		}
		n[v] = k
	}
	return &APIServer{
		bindIP:    ip,
		bindPort:  port,
		loger:     log.FmdLoger,
		tokenList: n,
		tls:       tls,
		keyPair:   pair,
		mc:        backend,
	}, nil
}

func (a *APIServer) prepareGin() *gin.Engine {

	if a.loger.Logger.Level.String() == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(a.logMiddleware())
	r.Use(a.tokenAuthMiddleware())
	r.Use(gin.Recovery())
	if a.tls {
		r.Use(a.tlsHeaderMiddleware())
	}
	web := r.Group("/api")
	web.Use(gzip.Gzip(gzip.DefaultCompression))
	web.GET("/region", a.listRegions)

	jobAPI := web.Group("/job")
	jobAPI.POST("/", a.addJob)
	jobAPI.PUT("/", a.modifyJob)
	jobAPI.GET("/", a.listAllJobs)
	jobAPI.GET("/:region", a.listJobsBelongToRegion)
	jobAPI.GET("/:region/:name", a.listTheJob)
	jobAPI.DELETE("/:region/:name", a.deleteJob)
	jobAPI.GET("/:region/:name/run", a.runJob)
	jobAPI.GET("/:region/:name/status", a.getJobStatus)
	jobAPI.POST("/search", a.jobSearch)

	executionAPI := web.Group("/execution")
	executionAPI.GET("/", a.listAllExecutions)
	executionAPI.GET("/:region", a.listExecutionsBelongToRgeion)
	executionAPI.GET("/:region/:name", a.listExecutionsBelongToJob)
	executionAPI.GET("/:region/:name/:group", a.listExecutionInGroup)
	executionAPI.POST("/search", a.executionSearch)

	return r
}

func (a *APIServer) listAllExecutions(c *gin.Context) {
	a.listExecutions("", "", 0, c)
}

func (a *APIServer) listExecutionsBelongToRgeion(c *gin.Context) {
	region := c.Params.ByName("region")
	a.listExecutions("", region, 0, c)
}

func (a *APIServer) listExecutionsBelongToJob(c *gin.Context) {
	region := c.Params.ByName("region")
	name := c.Params.ByName("name")
	a.listExecutions(name, region, 0, c)
}

func (a *APIServer) listExecutionInGroup(c *gin.Context) {
	region := c.Params.ByName("region")
	name := c.Params.ByName("name")
	groupStr := c.Params.ByName("group")
	group, err := strconv.ParseInt(groupStr, 10, 64)
	if err != nil {
		a.respondWithError(http.StatusUnprocessableEntity, &pb.ApiStringResponse{Succeed: false, Message: err.Error()}, c)
		return
	}
	a.listExecutions(name, region, group, c)
}

func (a *APIServer) listExecutions(name, region string, group int64, c *gin.Context) {
	outs, err := a.mc.ListExecutions(name, region, group)
	if err != nil {
		a.respondWithError(http.StatusInternalServerError, &pb.ApiStringResponse{Succeed: false, Message: err.Error()}, c)
		return
	}
	resp := &pb.ApiExecutionResponse{Succeed: true, Message: "Succeed", Data: outs}
	c.Render(http.StatusOK, pbjson{data: &resp})
}

func (a *APIServer) jobSearch(c *gin.Context) {
	var search pb.Search
	err := c.MustBindWith(&search, jsonpbBinding{})
	if err != nil {
		a.respondWithError(http.StatusInternalServerError, &pb.ApiStringResponse{Succeed: false, Message: err.Error()}, c)
		return
	}
	a.search(&pb.Job{}, &search, c)
}

func (a *APIServer) executionSearch(c *gin.Context) {
	var search pb.Search
	err := c.MustBindWith(&search, jsonpbBinding{})
	if err != nil {
		a.respondWithError(http.StatusInternalServerError, &pb.ApiStringResponse{Succeed: false, Message: err.Error()}, c)
		return
	}
	a.search(&pb.Execution{}, &search, c)
}

func (a *APIServer) search(obj interface{}, search *pb.Search, c *gin.Context) {
	outs, count, err := a.mc.Search(obj, search)

	if err != nil {
		a.respondWithError(http.StatusInternalServerError, &pb.ApiStringResponse{Succeed: false, Message: err.Error()}, c)
		return
	}

	switch obj.(type) {
	case *pb.Job:
		resp := &pb.ApiJobResponse{}
		for _, out := range outs {
			if t, ok := out.(*pb.Job); ok {
				resp.Data = append(resp.Data, t)
			} else {
				a.respondWithError(http.StatusInternalServerError, &pb.ApiJobResponse{Succeed: false, Message: errors.ErrNotExpectation.Error()}, c)
				log.FmdLoger.WithError(errors.ErrNotExpectation).Fatal("API: Search result error")
				return
			}
		}
		resp.MaxPageNum = int32(count)
		resp.Succeed = true
		c.Render(http.StatusOK, pbjson{data: &resp})
	case *pb.Execution:
		resp := &pb.ApiExecutionResponse{}
		for _, out := range outs {
			if t, ok := out.(*pb.Execution); ok {
				resp.Data = append(resp.Data, t)
			} else {
				a.respondWithError(http.StatusInternalServerError, &pb.ApiJobResponse{Succeed: false, Message: errors.ErrNotExpectation.Error()}, c)
				log.FmdLoger.WithError(errors.ErrNotExpectation).Fatal("API: Search result error")
				return
			}
		}
		resp.MaxPageNum = int32(count)
		resp.Succeed = true
		c.Render(http.StatusOK, pbjson{data: &resp})
	default:
		a.respondWithError(http.StatusInternalServerError, &pb.ApiStringResponse{Succeed: false, Message: errors.ErrType.Error()}, c)
		return
	}
}

func (a *APIServer) listRegions(c *gin.Context) {
	var resp pb.ApiStringResponse
	var err error
	resp.Data, err = a.mc.ListRegions()
	if err != nil {
		resp.Message = err.Error()
		a.respondWithError(http.StatusInternalServerError, &resp, c)
		return
	}
	resp.Succeed = true
	resp.Message = "Succeed"
	c.Render(http.StatusOK, pbjson{data: &resp})
}

func (a *APIServer) addJob(c *gin.Context) {
	var job pb.Job
	err := c.MustBindWith(&job, jsonpbBinding{})
	if err != nil {
		a.respondWithError(http.StatusBadRequest, &pb.ApiJobResponse{Succeed: false, Message: err.Error()}, c)
		return
	}
	a.jobCRUD(&job, pb.Ops_ADD, c)
}

func (a *APIServer) modifyJob(c *gin.Context) {
	var job pb.Job
	err := c.MustBindWith(&job, jsonpbBinding{})
	if err != nil {
		a.respondWithError(http.StatusBadRequest, &pb.ApiJobResponse{Succeed: false, Message: err.Error()}, c)
		return
	}
	a.jobCRUD(&job, pb.Ops_MODIFY, c)
}

func (a *APIServer) listAllJobs(c *gin.Context) {
	a.jobCRUD(&pb.Job{}, pb.Ops_READ, c)
}

func (a *APIServer) listJobsBelongToRegion(c *gin.Context) {
	region := c.Params.ByName("region")
	a.jobCRUD(&pb.Job{Region: region}, pb.Ops_READ, c)
}

func (a *APIServer) listTheJob(c *gin.Context) {
	name := c.Params.ByName("name")
	region := c.Params.ByName("region")
	a.jobCRUD(&pb.Job{Name: name, Region: region}, pb.Ops_READ, c)
}

func (a *APIServer) deleteJob(c *gin.Context) {
	name := c.Params.ByName("name")
	region := c.Params.ByName("region")
	a.jobCRUD(&pb.Job{Name: name, Region: region}, pb.Ops_DELETE, c)
}

func (a *APIServer) jobCRUD(in *pb.Job, ops pb.Ops, c *gin.Context) {

	if ops == pb.Ops_ADD || ops == pb.Ops_MODIFY {
		// this fields are forbidden to change by user
		in.SchedulerNodeName = ""
		in.ParentJobName = ""
	}

	resp := pb.ApiJobResponse{}
	var err error
	var out *pb.Job
	switch ops {
	case pb.Ops_READ:
		resp.Data, err = a.mc.ListJob(in.Name, in.Region)
	case pb.Ops_ADD:
		out, err = a.mc.AddJob(in)
	case pb.Ops_MODIFY:
		out, err = a.mc.ModifyJob(in)
	case pb.Ops_DELETE:
		out, err = a.mc.DeleteJob(in)
	default:
		err = errors.ErrUnknownOps
	}
	if err != nil {
		resp.Succeed = false
		resp.Message = err.Error()
		a.respondWithError(http.StatusInternalServerError, &resp, c)
		return
	}
	if out != nil {
		resp.Data = append(resp.Data, out)
	}
	resp.Succeed = true
	resp.Message = "Succeed"
	c.Render(http.StatusOK, pbjson{data: &resp})
}

func (a *APIServer) getJobStatus(c *gin.Context) {
	var resp pb.ApiJobStatusResponse
	name := c.Params.ByName("name")
	region := c.Params.ByName("region")
	status, err := a.mc.GetStatus(name, region)
	if err != nil {
		resp.Succeed = false
		resp.Message = err.Error()
		a.respondWithError(http.StatusInternalServerError, &resp, c)
		return
	}
	resp.Succeed = true
	resp.Message = "Succeed"
	resp.Data = append(resp.Data, status)
	c.Render(http.StatusOK, pbjson{data: &resp})
}

func (a *APIServer) runJob(c *gin.Context) {
	var resp pb.ApiExecutionResponse
	name := c.Params.ByName("name")
	region := c.Params.ByName("region")
	execution, err := a.mc.RunJob(name, region)
	if err != nil {
		resp.Succeed = false
		resp.Message = err.Error()
		a.respondWithError(http.StatusInternalServerError, &resp, c)
		return
	}
	resp.Data = append(resp.Data, execution)
	resp.Succeed = true
	resp.Message = "Succeed"
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
			a.respondWithError(http.StatusUnauthorized, &pb.ApiJobResponse{Succeed: false, Message: "API token required"}, c)
			return
		}
		if _, exist := a.tokenList[token]; exist {
			c.Next()
		} else {
			a.respondWithError(http.StatusUnauthorized, &pb.ApiJobResponse{Succeed: false, Message: "API token Error"}, c)
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
