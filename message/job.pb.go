// Code generated by protoc-gen-go.
// source: job.proto
// DO NOT EDIT!

package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Name struct {
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
}

func (m *Name) Reset()                    { *m = Name{} }
func (m *Name) String() string            { return proto.CompactTextString(m) }
func (*Name) ProtoMessage()               {}
func (*Name) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Name) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Job struct {
	Name               string   `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Region             string   `protobuf:"bytes,2,opt,name=Region" json:"Region,omitempty"`
	Schedule           string   `protobuf:"bytes,3,opt,name=schedule" json:"schedule,omitempty"`
	Shell              bool     `protobuf:"varint,4,opt,name=Shell" json:"Shell,omitempty"`
	Command            string   `protobuf:"bytes,5,opt,name=Command" json:"Command,omitempty"`
	Expression         string   `protobuf:"bytes,6,opt,name=Expression" json:"Expression,omitempty"`
	BeingDependentJobs []string `protobuf:"bytes,7,rep,name=BeingDependentJobs" json:"BeingDependentJobs,omitempty"`
	ParentJob          string   `protobuf:"bytes,8,opt,name=ParentJob" json:"ParentJob,omitempty"`
	Parallel           bool     `protobuf:"varint,9,opt,name=Parallel" json:"Parallel,omitempty"`
	Concurrent         bool     `protobuf:"varint,10,opt,name=Concurrent" json:"Concurrent,omitempty"`
	Disable            bool     `protobuf:"varint,11,opt,name=Disable" json:"Disable,omitempty"`
	SchedulerNodeName  string   `protobuf:"bytes,12,opt,name=SchedulerNodeName" json:"SchedulerNodeName,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Job) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Job) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *Job) GetSchedule() string {
	if m != nil {
		return m.Schedule
	}
	return ""
}

func (m *Job) GetShell() bool {
	if m != nil {
		return m.Shell
	}
	return false
}

func (m *Job) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Job) GetExpression() string {
	if m != nil {
		return m.Expression
	}
	return ""
}

func (m *Job) GetBeingDependentJobs() []string {
	if m != nil {
		return m.BeingDependentJobs
	}
	return nil
}

func (m *Job) GetParentJob() string {
	if m != nil {
		return m.ParentJob
	}
	return ""
}

func (m *Job) GetParallel() bool {
	if m != nil {
		return m.Parallel
	}
	return false
}

func (m *Job) GetConcurrent() bool {
	if m != nil {
		return m.Concurrent
	}
	return false
}

func (m *Job) GetDisable() bool {
	if m != nil {
		return m.Disable
	}
	return false
}

func (m *Job) GetSchedulerNodeName() string {
	if m != nil {
		return m.SchedulerNodeName
	}
	return ""
}

type JobStatus struct {
	Name            string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	SuccessCount    int64  `protobuf:"varint,2,opt,name=SuccessCount" json:"SuccessCount,omitempty"`
	ErrorCount      int64  `protobuf:"varint,3,opt,name=ErrorCount" json:"ErrorCount,omitempty"`
	LastHandleAgent string `protobuf:"bytes,4,opt,name=LastHandleAgent" json:"LastHandleAgent,omitempty"`
	LastSuccess     string `protobuf:"bytes,5,opt,name=LastSuccess" json:"LastSuccess,omitempty"`
	LastError       string `protobuf:"bytes,6,opt,name=LastError" json:"LastError,omitempty"`
}

func (m *JobStatus) Reset()                    { *m = JobStatus{} }
func (m *JobStatus) String() string            { return proto.CompactTextString(m) }
func (*JobStatus) ProtoMessage()               {}
func (*JobStatus) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *JobStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JobStatus) GetSuccessCount() int64 {
	if m != nil {
		return m.SuccessCount
	}
	return 0
}

func (m *JobStatus) GetErrorCount() int64 {
	if m != nil {
		return m.ErrorCount
	}
	return 0
}

func (m *JobStatus) GetLastHandleAgent() string {
	if m != nil {
		return m.LastHandleAgent
	}
	return ""
}

func (m *JobStatus) GetLastSuccess() string {
	if m != nil {
		return m.LastSuccess
	}
	return ""
}

func (m *JobStatus) GetLastError() string {
	if m != nil {
		return m.LastError
	}
	return ""
}

type Execution struct {
	Name       string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Cmd        string `protobuf:"bytes,2,opt,name=Cmd" json:"Cmd,omitempty"`
	Output     []byte `protobuf:"bytes,3,opt,name=Output,proto3" json:"Output,omitempty"`
	Succeed    bool   `protobuf:"varint,4,opt,name=Succeed" json:"Succeed,omitempty"`
	StartTime  string `protobuf:"bytes,5,opt,name=StartTime" json:"StartTime,omitempty"`
	FinishTime string `protobuf:"bytes,6,opt,name=FinishTime" json:"FinishTime,omitempty"`
	NodeName   string `protobuf:"bytes,7,opt,name=NodeName" json:"NodeName,omitempty"`
	JobName    string `protobuf:"bytes,8,opt,name=JobName" json:"JobName,omitempty"`
	Retries    int64  `protobuf:"varint,9,opt,name=Retries" json:"Retries,omitempty"`
}

func (m *Execution) Reset()                    { *m = Execution{} }
func (m *Execution) String() string            { return proto.CompactTextString(m) }
func (*Execution) ProtoMessage()               {}
func (*Execution) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *Execution) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Execution) GetCmd() string {
	if m != nil {
		return m.Cmd
	}
	return ""
}

func (m *Execution) GetOutput() []byte {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *Execution) GetSucceed() bool {
	if m != nil {
		return m.Succeed
	}
	return false
}

func (m *Execution) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *Execution) GetFinishTime() string {
	if m != nil {
		return m.FinishTime
	}
	return ""
}

func (m *Execution) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Execution) GetJobName() string {
	if m != nil {
		return m.JobName
	}
	return ""
}

func (m *Execution) GetRetries() int64 {
	if m != nil {
		return m.Retries
	}
	return 0
}

type Result struct {
	Status  int32  `protobuf:"varint,1,opt,name=Status" json:"Status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *Result) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Result) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Name)(nil), "message.Name")
	proto.RegisterType((*Job)(nil), "message.Job")
	proto.RegisterType((*JobStatus)(nil), "message.JobStatus")
	proto.RegisterType((*Execution)(nil), "message.Execution")
	proto.RegisterType((*Result)(nil), "message.Result")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Job service

type JobClient interface {
	GetJob(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Job, error)
	GetExecution(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Execution, error)
	ExecDone(ctx context.Context, in *Execution, opts ...grpc.CallOption) (*Result, error)
	SetJob(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Result, error)
}

type jobClient struct {
	cc *grpc.ClientConn
}

func NewJobClient(cc *grpc.ClientConn) JobClient {
	return &jobClient{cc}
}

func (c *jobClient) GetJob(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := grpc.Invoke(ctx, "/message.job/GetJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) GetExecution(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Execution, error) {
	out := new(Execution)
	err := grpc.Invoke(ctx, "/message.job/GetExecution", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) ExecDone(ctx context.Context, in *Execution, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/message.job/ExecDone", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) SetJob(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/message.job/SetJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Job service

type JobServer interface {
	GetJob(context.Context, *Name) (*Job, error)
	GetExecution(context.Context, *Name) (*Execution, error)
	ExecDone(context.Context, *Execution) (*Result, error)
	SetJob(context.Context, *Job) (*Result, error)
}

func RegisterJobServer(s *grpc.Server, srv JobServer) {
	s.RegisterService(&_Job_serviceDesc, srv)
}

func _Job_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.job/GetJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).GetJob(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_GetExecution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).GetExecution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.job/GetExecution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).GetExecution(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_ExecDone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Execution)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).ExecDone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.job/ExecDone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).ExecDone(ctx, req.(*Execution))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_SetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).SetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.job/SetJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).SetJob(ctx, req.(*Job))
	}
	return interceptor(ctx, in, info, handler)
}

var _Job_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.job",
	HandlerType: (*JobServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJob",
			Handler:    _Job_GetJob_Handler,
		},
		{
			MethodName: "GetExecution",
			Handler:    _Job_GetExecution_Handler,
		},
		{
			MethodName: "ExecDone",
			Handler:    _Job_ExecDone_Handler,
		},
		{
			MethodName: "SetJob",
			Handler:    _Job_SetJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job.proto",
}

func init() { proto.RegisterFile("job.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x54, 0xdd, 0x8e, 0xd3, 0x3c,
	0x10, 0xfd, 0xba, 0xe9, 0xef, 0x6c, 0x3f, 0x2d, 0x58, 0x08, 0x59, 0x15, 0x42, 0x55, 0x6e, 0x58,
	0x24, 0x54, 0x09, 0xf6, 0x8e, 0x3b, 0x68, 0x17, 0x10, 0x82, 0x05, 0xa5, 0xbc, 0x40, 0xda, 0x58,
	0x6d, 0x50, 0x1a, 0x57, 0xb6, 0x23, 0xed, 0xeb, 0xf1, 0x00, 0x3c, 0x0b, 0xe2, 0x0d, 0x98, 0x19,
	0x3b, 0x49, 0x77, 0x37, 0x57, 0xcd, 0x39, 0x67, 0xec, 0x39, 0xf3, 0xe3, 0xc2, 0xe4, 0xa7, 0xde,
	0x2c, 0x8e, 0x46, 0x3b, 0x2d, 0x46, 0x07, 0x65, 0x6d, 0xba, 0x53, 0xf1, 0x0c, 0xfa, 0x37, 0xe9,
	0x41, 0x09, 0xe1, 0x7f, 0x65, 0x6f, 0xde, 0xbb, 0x9c, 0x24, 0xfc, 0x1d, 0xff, 0x39, 0x83, 0xe8,
	0xb3, 0xde, 0x74, 0x69, 0xe2, 0x29, 0x0c, 0x13, 0xb5, 0xcb, 0x75, 0x29, 0xcf, 0x98, 0x0d, 0x48,
	0xcc, 0x60, 0x6c, 0xb7, 0x7b, 0x95, 0x55, 0x85, 0x92, 0x11, 0x2b, 0x0d, 0x16, 0x4f, 0x60, 0xb0,
	0xde, 0xab, 0xa2, 0x90, 0x7d, 0x14, 0xc6, 0x89, 0x07, 0x42, 0xc2, 0x68, 0xa9, 0x0f, 0x87, 0xb4,
	0xcc, 0xe4, 0x80, 0x0f, 0xd4, 0x50, 0x3c, 0x07, 0xb8, 0xbe, 0x3d, 0x1a, 0x74, 0x4a, 0x79, 0x86,
	0x2c, 0x9e, 0x30, 0x62, 0x01, 0xe2, 0xbd, 0xca, 0xcb, 0xdd, 0x4a, 0x1d, 0x55, 0x99, 0xa9, 0xd2,
	0xa1, 0x59, 0x2b, 0x47, 0xf3, 0x08, 0xe3, 0x3a, 0x14, 0xf1, 0x0c, 0x26, 0xdf, 0x53, 0xe3, 0x91,
	0x1c, 0xf3, 0x75, 0x2d, 0x41, 0xce, 0x11, 0xa4, 0x45, 0xa1, 0x0a, 0x39, 0x61, 0x83, 0x0d, 0x26,
	0x27, 0x4b, 0x5d, 0x6e, 0x2b, 0x43, 0xc1, 0x12, 0x58, 0x3d, 0x61, 0xa8, 0x86, 0x55, 0x6e, 0xd3,
	0x0d, 0x16, 0x7d, 0xce, 0x62, 0x0d, 0xc5, 0x2b, 0x78, 0xbc, 0x0e, 0xf5, 0x9b, 0x1b, 0x9d, 0x29,
	0x6e, 0xe4, 0x94, 0x73, 0x3f, 0x14, 0xe2, 0xdf, 0x3d, 0x98, 0xa0, 0x97, 0xb5, 0x4b, 0x5d, 0x65,
	0x3b, 0xfb, 0x1e, 0xc3, 0x74, 0x5d, 0x6d, 0xb7, 0xd8, 0x81, 0xa5, 0xae, 0xd0, 0x0b, 0x75, 0x3f,
	0x4a, 0xee, 0x70, 0xdc, 0x37, 0x63, 0xb4, 0xf1, 0x11, 0x11, 0x47, 0x9c, 0x30, 0xe2, 0x12, 0x2e,
	0xbe, 0xa4, 0xd6, 0x7d, 0xc2, 0x1e, 0x17, 0xea, 0xdd, 0x8e, 0x4a, 0xea, 0x73, 0x8a, 0xfb, 0xb4,
	0x98, 0xc3, 0x39, 0x51, 0xe1, 0xf6, 0x30, 0x9f, 0x53, 0x8a, 0x7a, 0x4a, 0x90, 0x6f, 0x0f, 0x23,
	0x6a, 0x89, 0xf8, 0x2f, 0xd6, 0x73, 0x7d, 0xab, 0xb6, 0x95, 0xa3, 0x79, 0x75, 0xd5, 0xf3, 0x08,
	0xa2, 0xe5, 0x21, 0x0b, 0x4b, 0x44, 0x9f, 0xb4, 0x59, 0xdf, 0x2a, 0x77, 0xac, 0xbc, 0xf3, 0x69,
	0x12, 0x10, 0xf5, 0x98, 0x93, 0xaa, 0x2c, 0xec, 0x4f, 0x0d, 0xc9, 0x03, 0x76, 0xcc, 0xb8, 0x1f,
	0x39, 0x5e, 0xee, 0x3d, 0xb6, 0x04, 0x75, 0xe3, 0x43, 0x5e, 0xe6, 0x76, 0xcf, 0x72, 0xd8, 0xa2,
	0x96, 0xa1, 0xb9, 0x37, 0x83, 0x19, 0xf9, 0x8d, 0xad, 0x31, 0xe5, 0xc4, 0x71, 0xb0, 0xe4, 0xf7,
	0xa5, 0x86, 0xa4, 0x24, 0xca, 0x99, 0x5c, 0x59, 0x5e, 0x96, 0x28, 0xa9, 0x61, 0xfc, 0x96, 0x5e,
	0x86, 0xad, 0x0a, 0x47, 0x95, 0xf8, 0x49, 0x72, 0xc5, 0x83, 0x24, 0x20, 0x3a, 0xfb, 0xd5, 0x3f,
	0xbf, 0x50, 0x77, 0x0d, 0xdf, 0xfc, 0xea, 0x41, 0x84, 0x8f, 0x54, 0xbc, 0x80, 0xe1, 0x47, 0xc5,
	0x5b, 0xf9, 0xff, 0x22, 0xbc, 0xd4, 0x05, 0xa5, 0x9d, 0x4d, 0x1b, 0x88, 0x62, 0xfc, 0x9f, 0xb8,
	0x82, 0x29, 0x06, 0xb6, 0x2d, 0xbe, 0x17, 0x2e, 0x1a, 0xd8, 0x84, 0xe0, 0xa1, 0xd7, 0x30, 0x26,
	0xb8, 0xd2, 0x25, 0xbe, 0xfb, 0x87, 0x11, 0xb3, 0x8b, 0x86, 0xf3, 0x85, 0xe0, 0x91, 0x97, 0x58,
	0x8a, 0x37, 0x74, 0xc7, 0x41, 0x47, 0xe8, 0x66, 0xc8, 0xff, 0x30, 0x57, 0xff, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x6e, 0x75, 0x6b, 0xf0, 0x6e, 0x04, 0x00, 0x00,
}
