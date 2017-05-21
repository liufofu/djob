// Code generated by protoc-gen-go.
// source: job.proto
// DO NOT EDIT!

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	job.proto

It has these top-level messages:
	Name
	Job
	Execution
	Result
*/
package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Name struct {
	JobName string `protobuf:"bytes,1,opt,name=jobName" json:"jobName,omitempty"`
}

func (m *Name) Reset()                    { *m = Name{} }
func (m *Name) String() string            { return proto.CompactTextString(m) }
func (*Name) ProtoMessage()               {}
func (*Name) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Name) GetJobName() string {
	if m != nil {
		return m.JobName
	}
	return ""
}

type Job struct {
	JobName string `protobuf:"bytes,1,opt,name=jobName" json:"jobName,omitempty"`
	Cmd     string `protobuf:"bytes,2,opt,name=cmd" json:"cmd,omitempty"`
	Shell   bool   `protobuf:"varint,3,opt,name=shell" json:"shell,omitempty"`
	Running bool   `protobuf:"varint,4,opt,name=running" json:"running,omitempty"`
	Enabled bool   `protobuf:"varint,5,opt,name=enabled" json:"enabled,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Job) GetJobName() string {
	if m != nil {
		return m.JobName
	}
	return ""
}

func (m *Job) GetCmd() string {
	if m != nil {
		return m.Cmd
	}
	return ""
}

func (m *Job) GetShell() bool {
	if m != nil {
		return m.Shell
	}
	return false
}

func (m *Job) GetRunning() bool {
	if m != nil {
		return m.Running
	}
	return false
}

func (m *Job) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type Execution struct {
	Name       string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Cmd        string                     `protobuf:"bytes,2,opt,name=cmd" json:"cmd,omitempty"`
	Output     string                     `protobuf:"bytes,3,opt,name=output" json:"output,omitempty"`
	Succeed    bool                       `protobuf:"varint,4,opt,name=succeed" json:"succeed,omitempty"`
	StartTime  *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=startTime" json:"startTime,omitempty"`
	FinishTime *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=finishTime" json:"finishTime,omitempty"`
}

func (m *Execution) Reset()                    { *m = Execution{} }
func (m *Execution) String() string            { return proto.CompactTextString(m) }
func (*Execution) ProtoMessage()               {}
func (*Execution) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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

func (m *Execution) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *Execution) GetSucceed() bool {
	if m != nil {
		return m.Succeed
	}
	return false
}

func (m *Execution) GetStartTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Execution) GetFinishTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.FinishTime
	}
	return nil
}

type Result struct {
	Err bool `protobuf:"varint,1,opt,name=err" json:"err,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Result) GetErr() bool {
	if m != nil {
		return m.Err
	}
	return false
}

func init() {
	proto.RegisterType((*Name)(nil), "message.Name")
	proto.RegisterType((*Job)(nil), "message.Job")
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
	ExecDone(ctx context.Context, in *Execution, opts ...grpc.CallOption) (*Result, error)
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

func (c *jobClient) ExecDone(ctx context.Context, in *Execution, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/message.job/ExecDone", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Job service

type JobServer interface {
	GetJob(context.Context, *Name) (*Job, error)
	ExecDone(context.Context, *Execution) (*Result, error)
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

var _Job_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.job",
	HandlerType: (*JobServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJob",
			Handler:    _Job_GetJob_Handler,
		},
		{
			MethodName: "ExecDone",
			Handler:    _Job_ExecDone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job.proto",
}

func init() { proto.RegisterFile("job.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x5b, 0xd3, 0xc6, 0x66, 0x2a, 0xa2, 0x8b, 0x48, 0xc8, 0xc5, 0xb2, 0x20, 0x78, 0xda,
	0x42, 0xbd, 0x88, 0x67, 0x45, 0xf0, 0xe0, 0x61, 0xe9, 0xd5, 0x43, 0x93, 0x4e, 0xdb, 0x48, 0x76,
	0xb7, 0x24, 0xbb, 0x20, 0xf8, 0x49, 0xfd, 0x36, 0xee, 0x9f, 0xac, 0xf5, 0xa0, 0x78, 0x9b, 0xf7,
	0x66, 0x26, 0x6f, 0xf6, 0x17, 0xc8, 0xde, 0x54, 0xc9, 0xf6, 0xad, 0xd2, 0x8a, 0xa4, 0x95, 0x12,
	0x42, 0xc9, 0xe2, 0x6a, 0xab, 0xd4, 0xb6, 0xc1, 0xb9, 0x77, 0x4b, 0xb3, 0x99, 0xeb, 0x5a, 0x60,
	0xa7, 0x57, 0x62, 0x1f, 0x06, 0xe9, 0x0c, 0x46, 0x2f, 0x2b, 0x81, 0x24, 0x87, 0x63, 0xbb, 0xed,
	0xca, 0x7c, 0x38, 0x1b, 0xde, 0x64, 0x3c, 0x4a, 0xfa, 0x01, 0xc9, 0xb3, 0x2a, 0xff, 0x1e, 0x20,
	0x67, 0x90, 0x54, 0x62, 0x9d, 0x1f, 0x79, 0xd7, 0x95, 0xe4, 0x02, 0xc6, 0xdd, 0x0e, 0x9b, 0x26,
	0x4f, 0xac, 0x37, 0xe1, 0x41, 0xb8, 0x2f, 0xb4, 0x46, 0xca, 0x5a, 0x6e, 0xf3, 0x91, 0xf7, 0xa3,
	0x74, 0x1d, 0x94, 0xab, 0xb2, 0xc1, 0x75, 0x3e, 0x0e, 0x9d, 0x5e, 0xd2, 0xcf, 0x21, 0x64, 0x8f,
	0xef, 0x58, 0x19, 0x5d, 0x2b, 0x49, 0x08, 0x8c, 0xe4, 0xe1, 0x00, 0x5f, 0xff, 0x92, 0x7e, 0x09,
	0xa9, 0x32, 0x7a, 0x6f, 0xb4, 0x8f, 0xcf, 0x78, 0xaf, 0x5c, 0x4a, 0x67, 0xaa, 0x0a, 0x6d, 0x4a,
	0x9f, 0xdf, 0x4b, 0x72, 0x07, 0x99, 0x65, 0xd2, 0xea, 0xa5, 0x85, 0xe3, 0x2f, 0x98, 0x2e, 0x0a,
	0x16, 0xc8, 0xb1, 0x48, 0x8e, 0x2d, 0x23, 0x39, 0x7e, 0x18, 0x26, 0xf7, 0x00, 0x9b, 0x5a, 0xd6,
	0xdd, 0xce, 0xaf, 0xa6, 0xff, 0xae, 0xfe, 0x98, 0xa6, 0x05, 0xa4, 0x1c, 0x3b, 0xd3, 0x68, 0xf7,
	0x06, 0x6c, 0x5b, 0xff, 0xac, 0x09, 0x77, 0xe5, 0xe2, 0x15, 0x12, 0x8b, 0x97, 0x5c, 0x43, 0xfa,
	0x84, 0xda, 0xe1, 0x3f, 0x61, 0xe1, 0x8f, 0x32, 0x87, 0xbc, 0x98, 0x46, 0x65, 0x5b, 0x74, 0x40,
	0xe6, 0x30, 0x71, 0x90, 0x1e, 0x94, 0x44, 0x72, 0x1e, 0x5b, 0xdf, 0xd8, 0x8a, 0xd3, 0x68, 0x85,
	0x38, 0x3a, 0x28, 0x53, 0x7f, 0xda, 0xed, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x25, 0xa9, 0x5b,
	0xf2, 0x32, 0x02, 0x00, 0x00,
}
