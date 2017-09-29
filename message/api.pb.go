// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	api.proto
	job.proto
	serfQueryParams.proto

It has these top-level messages:
	RespJob
	RespExec
	RespStatus
	RespJobs
	RespExecs
	RespStatuses
	Params
	Job
	JobStatus
	Execution
	JobQueryParams
	GetRPCConfigResp
	JobCountResp
	QueryResult
*/
package message

import "github.com/golang/protobuf/proto"
import "fmt"
import "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RespJob struct {
	Status  int32  `protobuf:"varint,1,opt,name=Status" json:"Status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	Node    string `protobuf:"bytes,3,opt,name=Node" json:"Node,omitempty"`
	Result  []*Job `protobuf:"bytes,4,rep,name=Result" json:"Result,omitempty"`
}

func (m *RespJob) Reset()                    { *m = RespJob{} }
func (m *RespJob) String() string            { return proto.CompactTextString(m) }
func (*RespJob) ProtoMessage()               {}
func (*RespJob) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RespJob) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *RespJob) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RespJob) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *RespJob) GetResult() []*Job {
	if m != nil {
		return m.Result
	}
	return nil
}

type RespExec struct {
	Status  int32        `protobuf:"varint,1,opt,name=Status" json:"Status,omitempty"`
	Message string       `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	Result  []*Execution `protobuf:"bytes,3,rep,name=Result" json:"Result,omitempty"`
}

func (m *RespExec) Reset()                    { *m = RespExec{} }
func (m *RespExec) String() string            { return proto.CompactTextString(m) }
func (*RespExec) ProtoMessage()               {}
func (*RespExec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RespExec) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *RespExec) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RespExec) GetResult() []*Execution {
	if m != nil {
		return m.Result
	}
	return nil
}

type RespStatus struct {
	Status  int32        `protobuf:"varint,1,opt,name=Status" json:"Status,omitempty"`
	Message string       `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	Result  []*JobStatus `protobuf:"bytes,3,rep,name=Result" json:"Result,omitempty"`
}

func (m *RespStatus) Reset()                    { *m = RespStatus{} }
func (m *RespStatus) String() string            { return proto.CompactTextString(m) }
func (*RespStatus) ProtoMessage()               {}
func (*RespStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RespStatus) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *RespStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RespStatus) GetResult() []*JobStatus {
	if m != nil {
		return m.Result
	}
	return nil
}

type RespJobs struct {
	Data []*RespJob `protobuf:"bytes,1,rep,name=Data" json:"Data,omitempty"`
}

func (m *RespJobs) Reset()                    { *m = RespJobs{} }
func (m *RespJobs) String() string            { return proto.CompactTextString(m) }
func (*RespJobs) ProtoMessage()               {}
func (*RespJobs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RespJobs) GetData() []*RespJob {
	if m != nil {
		return m.Data
	}
	return nil
}

type RespExecs struct {
	Data []*RespExec `protobuf:"bytes,1,rep,name=Data" json:"Data,omitempty"`
}

func (m *RespExecs) Reset()                    { *m = RespExecs{} }
func (m *RespExecs) String() string            { return proto.CompactTextString(m) }
func (*RespExecs) ProtoMessage()               {}
func (*RespExecs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RespExecs) GetData() []*RespExec {
	if m != nil {
		return m.Data
	}
	return nil
}

type RespStatuses struct {
	Data []*RespStatus `protobuf:"bytes,1,rep,name=Data" json:"Data,omitempty"`
}

func (m *RespStatuses) Reset()                    { *m = RespStatuses{} }
func (m *RespStatuses) String() string            { return proto.CompactTextString(m) }
func (*RespStatuses) ProtoMessage()               {}
func (*RespStatuses) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RespStatuses) GetData() []*RespStatus {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*RespJob)(nil), "message.RespJob")
	proto.RegisterType((*RespExec)(nil), "message.RespExec")
	proto.RegisterType((*RespStatus)(nil), "message.RespStatus")
	proto.RegisterType((*RespJobs)(nil), "message.RespJobs")
	proto.RegisterType((*RespExecs)(nil), "message.RespExecs")
	proto.RegisterType((*RespStatuses)(nil), "message.RespStatuses")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x91, 0xcd, 0x4a, 0xc5, 0x30,
	0x10, 0x46, 0xb9, 0xb6, 0xb6, 0x66, 0xbc, 0x0b, 0x1d, 0x41, 0x82, 0xab, 0x52, 0x14, 0x8b, 0x8b,
	0x22, 0x75, 0xe1, 0x0b, 0xe8, 0x46, 0xd0, 0x45, 0x7c, 0x82, 0x54, 0x83, 0xb6, 0xa8, 0x29, 0x26,
	0x01, 0x1f, 0xdf, 0xfc, 0x35, 0x92, 0xad, 0xdc, 0x5d, 0x93, 0x6f, 0xce, 0x9c, 0xce, 0x04, 0x08,
	0x5f, 0xa6, 0x7e, 0xf9, 0x96, 0x5a, 0x62, 0xfd, 0x29, 0x94, 0xe2, 0x6f, 0xe2, 0x8c, 0xcc, 0x72,
	0x0c, 0x77, 0xad, 0x81, 0x9a, 0x09, 0xb5, 0x3c, 0xc8, 0x11, 0x4f, 0xa1, 0x7a, 0xd6, 0x5c, 0x1b,
	0x45, 0x37, 0xcd, 0xa6, 0xdb, 0x67, 0xf1, 0x84, 0x14, 0xea, 0xc7, 0x00, 0xd2, 0x3d, 0x1b, 0x10,
	0xb6, 0x1e, 0x11, 0xa1, 0x7c, 0x92, 0xaf, 0x82, 0x16, 0xfe, 0xda, 0x7f, 0xe3, 0x39, 0x54, 0xb6,
	0xa1, 0xf9, 0xd0, 0xb4, 0x6c, 0x8a, 0xee, 0x70, 0xd8, 0xf6, 0xd1, 0xda, 0x5b, 0x07, 0x8b, 0x59,
	0xfb, 0x0e, 0x07, 0x4e, 0x7b, 0xff, 0x23, 0x5e, 0xfe, 0xe1, 0xbd, 0x4a, 0x8e, 0xc2, 0x3b, 0x30,
	0x39, 0x5c, 0x43, 0xa3, 0x27, 0xf9, 0x95, 0x4c, 0x33, 0x80, 0x33, 0xc5, 0x9e, 0xbb, 0x74, 0xd9,
	0x79, 0x02, 0x9d, 0x5c, 0xd7, 0x61, 0x2a, 0x1b, 0x28, 0xbb, 0x87, 0xf2, 0x8e, 0x6b, 0x6e, 0x3d,
	0x8e, 0x3a, 0x4a, 0x54, 0x2c, 0x60, 0x3e, 0x6d, 0x07, 0x20, 0xeb, 0x1e, 0x14, 0x5e, 0x64, 0xc8,
	0x71, 0x86, 0xb8, 0x8a, 0xc8, 0xdc, 0xc2, 0xf6, 0x6f, 0x22, 0xa1, 0xf0, 0x32, 0xc3, 0x4e, 0x32,
	0x2c, 0xfe, 0xa0, 0x2f, 0x18, 0x2b, 0xff, 0xe4, 0x37, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x25,
	0x74, 0x05, 0x4c, 0x13, 0x02, 0x00, 0x00,
}
