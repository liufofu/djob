// Code generated by protoc-gen-go.
// source: message/api.proto
// DO NOT EDIT!

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	message/api.proto
	message/serfQueryParams.proto
	message/job.proto

It has these top-level messages:
	ApiJobResponse
	ApiExecutionResponse
	ApiJobStatusResponse
	ApiStringResponse
	Pageing
	SearchCondition
	ApiJobQueryString
	ApiJobStatusQueryString
	ApiExecutionQueryString
	ApiSearchQueryString
	JobQueryParams
	GetRPCConfigResp
	JobCountResp
	QueryResult
	Search
	Params
	Result
	Job
	JobStatus
	Execution
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

type ApiJobResponse struct {
	Succeed    bool   `protobuf:"varint,1,opt,name=Succeed" json:"Succeed,omitempty"`
	Message    string `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	MaxPageNum int32  `protobuf:"varint,3,opt,name=MaxPageNum" json:"MaxPageNum,omitempty"`
	Data       []*Job `protobuf:"bytes,4,rep,name=Data" json:"Data,omitempty"`
}

func (m *ApiJobResponse) Reset()                    { *m = ApiJobResponse{} }
func (m *ApiJobResponse) String() string            { return proto.CompactTextString(m) }
func (*ApiJobResponse) ProtoMessage()               {}
func (*ApiJobResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ApiJobResponse) GetSucceed() bool {
	if m != nil {
		return m.Succeed
	}
	return false
}

func (m *ApiJobResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ApiJobResponse) GetMaxPageNum() int32 {
	if m != nil {
		return m.MaxPageNum
	}
	return 0
}

func (m *ApiJobResponse) GetData() []*Job {
	if m != nil {
		return m.Data
	}
	return nil
}

type ApiExecutionResponse struct {
	Succeed    bool         `protobuf:"varint,1,opt,name=Succeed" json:"Succeed,omitempty"`
	Message    string       `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	MaxPageNum int32        `protobuf:"varint,3,opt,name=MaxPageNum" json:"MaxPageNum,omitempty"`
	Data       []*Execution `protobuf:"bytes,4,rep,name=Data" json:"Data,omitempty"`
}

func (m *ApiExecutionResponse) Reset()                    { *m = ApiExecutionResponse{} }
func (m *ApiExecutionResponse) String() string            { return proto.CompactTextString(m) }
func (*ApiExecutionResponse) ProtoMessage()               {}
func (*ApiExecutionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ApiExecutionResponse) GetSucceed() bool {
	if m != nil {
		return m.Succeed
	}
	return false
}

func (m *ApiExecutionResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ApiExecutionResponse) GetMaxPageNum() int32 {
	if m != nil {
		return m.MaxPageNum
	}
	return 0
}

func (m *ApiExecutionResponse) GetData() []*Execution {
	if m != nil {
		return m.Data
	}
	return nil
}

type ApiJobStatusResponse struct {
	Succeed    bool         `protobuf:"varint,1,opt,name=Succeed" json:"Succeed,omitempty"`
	Message    string       `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	MaxPageNum int32        `protobuf:"varint,3,opt,name=MaxPageNum" json:"MaxPageNum,omitempty"`
	Data       []*JobStatus `protobuf:"bytes,4,rep,name=Data" json:"Data,omitempty"`
}

func (m *ApiJobStatusResponse) Reset()                    { *m = ApiJobStatusResponse{} }
func (m *ApiJobStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*ApiJobStatusResponse) ProtoMessage()               {}
func (*ApiJobStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ApiJobStatusResponse) GetSucceed() bool {
	if m != nil {
		return m.Succeed
	}
	return false
}

func (m *ApiJobStatusResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ApiJobStatusResponse) GetMaxPageNum() int32 {
	if m != nil {
		return m.MaxPageNum
	}
	return 0
}

func (m *ApiJobStatusResponse) GetData() []*JobStatus {
	if m != nil {
		return m.Data
	}
	return nil
}

type ApiStringResponse struct {
	Succeed    bool     `protobuf:"varint,1,opt,name=Succeed" json:"Succeed,omitempty"`
	Message    string   `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	MaxPageNum int32    `protobuf:"varint,3,opt,name=MaxPageNum" json:"MaxPageNum,omitempty"`
	Data       []string `protobuf:"bytes,4,rep,name=Data" json:"Data,omitempty"`
}

func (m *ApiStringResponse) Reset()                    { *m = ApiStringResponse{} }
func (m *ApiStringResponse) String() string            { return proto.CompactTextString(m) }
func (*ApiStringResponse) ProtoMessage()               {}
func (*ApiStringResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ApiStringResponse) GetSucceed() bool {
	if m != nil {
		return m.Succeed
	}
	return false
}

func (m *ApiStringResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ApiStringResponse) GetMaxPageNum() int32 {
	if m != nil {
		return m.MaxPageNum
	}
	return 0
}

func (m *ApiStringResponse) GetData() []string {
	if m != nil {
		return m.Data
	}
	return nil
}

type Pageing struct {
	// @inject_tag: form:"pagenum"
	PageNum int32 `protobuf:"varint,1,opt,name=PageNum" json:"PageNum,omitempty" form:"pagenum"`
	// @inject_tag: form:"pagesize"
	PageSize int32 `protobuf:"varint,2,opt,name=PageSize" json:"PageSize,omitempty" form:"pagesize"`
	// @inject_tag: form:"maxpage"
	OutMaxPage bool `protobuf:"varint,3,opt,name=OutMaxPage" json:"OutMaxPage,omitempty" form:"maxpage"`
}

func (m *Pageing) Reset()                    { *m = Pageing{} }
func (m *Pageing) String() string            { return proto.CompactTextString(m) }
func (*Pageing) ProtoMessage()               {}
func (*Pageing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Pageing) GetPageNum() int32 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *Pageing) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *Pageing) GetOutMaxPage() bool {
	if m != nil {
		return m.OutMaxPage
	}
	return false
}

type SearchCondition struct {
	// @inject_tag: form:"conditions"
	Conditions []string `protobuf:"bytes,1,rep,name=Conditions" json:"Conditions,omitempty" form:"conditions"`
	// @inject_tag: form:"links"
	Links []string `protobuf:"bytes,2,rep,name=Links" json:"Links,omitempty" form:"links"`
}

func (m *SearchCondition) Reset()                    { *m = SearchCondition{} }
func (m *SearchCondition) String() string            { return proto.CompactTextString(m) }
func (*SearchCondition) ProtoMessage()               {}
func (*SearchCondition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SearchCondition) GetConditions() []string {
	if m != nil {
		return m.Conditions
	}
	return nil
}

func (m *SearchCondition) GetLinks() []string {
	if m != nil {
		return m.Links
	}
	return nil
}

type ApiJobQueryString struct {
	// @inject_tag: form:"job"
	Job *Job `protobuf:"bytes,1,opt,name=Job" json:"Job,omitempty" form:"job"`
	// @inject_tag: form:"pageing"
	Pageing *Pageing `protobuf:"bytes,2,opt,name=Pageing" json:"Pageing,omitempty" form:"pageing"`
}

func (m *ApiJobQueryString) Reset()                    { *m = ApiJobQueryString{} }
func (m *ApiJobQueryString) String() string            { return proto.CompactTextString(m) }
func (*ApiJobQueryString) ProtoMessage()               {}
func (*ApiJobQueryString) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ApiJobQueryString) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *ApiJobQueryString) GetPageing() *Pageing {
	if m != nil {
		return m.Pageing
	}
	return nil
}

type ApiJobStatusQueryString struct {
	// @inject_tag: form:"status"
	Status *JobStatus `protobuf:"bytes,1,opt,name=status" json:"status,omitempty" form:"status"`
	// @inject_tag: form:"pageing"
	Pageing *Pageing `protobuf:"bytes,2,opt,name=Pageing" json:"Pageing,omitempty" form:"pageing"`
}

func (m *ApiJobStatusQueryString) Reset()                    { *m = ApiJobStatusQueryString{} }
func (m *ApiJobStatusQueryString) String() string            { return proto.CompactTextString(m) }
func (*ApiJobStatusQueryString) ProtoMessage()               {}
func (*ApiJobStatusQueryString) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ApiJobStatusQueryString) GetStatus() *JobStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ApiJobStatusQueryString) GetPageing() *Pageing {
	if m != nil {
		return m.Pageing
	}
	return nil
}

type ApiExecutionQueryString struct {
	// @inject_tag: form:"execution"
	Execution *Execution `protobuf:"bytes,1,opt,name=Execution" json:"Execution,omitempty" form:"execution"`
	// @inject_tag: form:"pageing"
	Pageing *Pageing `protobuf:"bytes,2,opt,name=Pageing" json:"Pageing,omitempty" form:"pageing"`
}

func (m *ApiExecutionQueryString) Reset()                    { *m = ApiExecutionQueryString{} }
func (m *ApiExecutionQueryString) String() string            { return proto.CompactTextString(m) }
func (*ApiExecutionQueryString) ProtoMessage()               {}
func (*ApiExecutionQueryString) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ApiExecutionQueryString) GetExecution() *Execution {
	if m != nil {
		return m.Execution
	}
	return nil
}

func (m *ApiExecutionQueryString) GetPageing() *Pageing {
	if m != nil {
		return m.Pageing
	}
	return nil
}

type ApiSearchQueryString struct {
	// @inject_tag: form:"q"
	SearchCondition *SearchCondition `protobuf:"bytes,1,opt,name=SearchCondition" json:"SearchCondition,omitempty" form:"q"`
	// @inject_tag: form:"pageing"
	Pageing *Pageing `protobuf:"bytes,2,opt,name=Pageing" json:"Pageing,omitempty" form:"pageing"`
}

func (m *ApiSearchQueryString) Reset()                    { *m = ApiSearchQueryString{} }
func (m *ApiSearchQueryString) String() string            { return proto.CompactTextString(m) }
func (*ApiSearchQueryString) ProtoMessage()               {}
func (*ApiSearchQueryString) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ApiSearchQueryString) GetSearchCondition() *SearchCondition {
	if m != nil {
		return m.SearchCondition
	}
	return nil
}

func (m *ApiSearchQueryString) GetPageing() *Pageing {
	if m != nil {
		return m.Pageing
	}
	return nil
}

func init() {
	proto.RegisterType((*ApiJobResponse)(nil), "message.ApiJobResponse")
	proto.RegisterType((*ApiExecutionResponse)(nil), "message.ApiExecutionResponse")
	proto.RegisterType((*ApiJobStatusResponse)(nil), "message.ApiJobStatusResponse")
	proto.RegisterType((*ApiStringResponse)(nil), "message.ApiStringResponse")
	proto.RegisterType((*Pageing)(nil), "message.Pageing")
	proto.RegisterType((*SearchCondition)(nil), "message.SearchCondition")
	proto.RegisterType((*ApiJobQueryString)(nil), "message.ApiJobQueryString")
	proto.RegisterType((*ApiJobStatusQueryString)(nil), "message.ApiJobStatusQueryString")
	proto.RegisterType((*ApiExecutionQueryString)(nil), "message.ApiExecutionQueryString")
	proto.RegisterType((*ApiSearchQueryString)(nil), "message.ApiSearchQueryString")
}

func init() { proto.RegisterFile("message/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x94, 0xcb, 0x4e, 0xc2, 0x40,
	0x14, 0x86, 0x53, 0xb9, 0x1f, 0x8c, 0xca, 0x84, 0xc4, 0x86, 0x85, 0x21, 0x5d, 0x18, 0xc2, 0x02,
	0x0d, 0x3e, 0x81, 0xb7, 0x98, 0x18, 0xf1, 0x32, 0x7d, 0x00, 0x53, 0xca, 0x04, 0xab, 0xa1, 0xad,
	0x9d, 0x36, 0xa2, 0xee, 0x79, 0x00, 0x9f, 0xd8, 0xb9, 0x75, 0x3a, 0x80, 0x1b, 0x16, 0xc4, 0xdd,
	0x9c, 0xf3, 0xcf, 0xcc, 0xf7, 0xc1, 0x19, 0x80, 0xd6, 0x8c, 0x50, 0xea, 0x4d, 0xc9, 0x89, 0x17,
	0x07, 0x83, 0x38, 0x89, 0xd2, 0x08, 0xd5, 0x54, 0xab, 0xa3, 0xb3, 0xd7, 0x68, 0x2c, 0x33, 0x67,
	0x61, 0xc1, 0xde, 0x79, 0x1c, 0xdc, 0x46, 0x63, 0x4c, 0x68, 0x1c, 0x85, 0x94, 0x20, 0x1b, 0x6a,
	0x6e, 0xe6, 0xfb, 0x84, 0x4c, 0x6c, 0xab, 0x6b, 0xf5, 0xea, 0x38, 0x2f, 0x79, 0x32, 0x92, 0x37,
	0xd8, 0x3b, 0x2c, 0x69, 0xe0, 0xbc, 0x44, 0x47, 0x00, 0x23, 0x6f, 0xfe, 0xc8, 0x96, 0xf7, 0xd9,
	0xcc, 0x2e, 0xb1, 0xb0, 0x82, 0x8d, 0x0e, 0xea, 0x42, 0xf9, 0xca, 0x4b, 0x3d, 0xbb, 0xdc, 0x2d,
	0xf5, 0x9a, 0xc3, 0xdd, 0x81, 0x12, 0x19, 0x70, 0xae, 0x48, 0x9c, 0x1f, 0x0b, 0xda, 0x4c, 0xe4,
	0x7a, 0x4e, 0xfc, 0x2c, 0x0d, 0xa2, 0x70, 0xab, 0x3a, 0xc7, 0x4b, 0x3a, 0x48, 0xeb, 0x14, 0xf4,
	0x25, 0x29, 0x66, 0xe9, 0xa6, 0x5e, 0x9a, 0xd1, 0x7f, 0x91, 0x2a, 0xe8, 0x52, 0xea, 0x1b, 0x5a,
	0xcc, 0xc9, 0x4d, 0x93, 0x20, 0x9c, 0x6e, 0x55, 0x08, 0x19, 0x42, 0x0d, 0x05, 0x7f, 0x86, 0x1a,
	0x8f, 0x19, 0x9a, 0x5f, 0x9c, 0x9f, 0xb5, 0xc4, 0xd9, 0xbc, 0x44, 0x1d, 0xa8, 0xf3, 0xa5, 0x1b,
	0x7c, 0x49, 0x66, 0x05, 0xeb, 0x9a, 0x43, 0x1f, 0xb2, 0x54, 0x51, 0x04, 0xb4, 0x8e, 0x8d, 0x8e,
	0x73, 0x03, 0xfb, 0x2e, 0xf1, 0x12, 0xff, 0xe5, 0x32, 0x0a, 0x27, 0x01, 0x9f, 0x05, 0x3f, 0xa2,
	0x0b, 0xca, 0x58, 0xdc, 0xc6, 0xe8, 0xa0, 0x36, 0x54, 0xee, 0x82, 0xf0, 0x8d, 0x32, 0x16, 0x8f,
	0x64, 0xc1, 0x4c, 0x5b, 0x72, 0x74, 0x4f, 0x19, 0x49, 0x3e, 0xe5, 0xd7, 0xc5, 0xae, 0x2a, 0xb1,
	0x8e, 0xf0, 0x5d, 0x7d, 0x86, 0x3c, 0x40, 0x7d, 0xfd, 0xf1, 0x84, 0x78, 0x73, 0x78, 0xa0, 0xf7,
	0xa8, 0x3e, 0xce, 0x37, 0x38, 0xef, 0x70, 0x68, 0xbe, 0x0d, 0x13, 0xd3, 0x87, 0x2a, 0x15, 0x4d,
	0x45, 0xfa, 0x6b, 0x98, 0x6a, 0xc7, 0x46, 0xc8, 0x0f, 0x81, 0xd4, 0xaf, 0xd4, 0x44, 0x9e, 0x42,
	0x43, 0xf7, 0xd7, 0xa8, 0xc5, 0xbb, 0x2e, 0x36, 0x6d, 0x04, 0x5e, 0xc8, 0x1f, 0x82, 0x9c, 0x8c,
	0x89, 0xbd, 0x58, 0x1b, 0x97, 0x82, 0xdb, 0xfa, 0xb2, 0x95, 0x1c, 0xaf, 0xcd, 0x77, 0x03, 0x91,
	0x71, 0x55, 0xfc, 0x6d, 0x9d, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0x34, 0x01, 0x17, 0xf7, 0xe7,
	0x04, 0x00, 0x00,
}
