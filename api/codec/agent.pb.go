// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: agent.proto

package codec

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AgentMessage struct {
	Metadata *AgentMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Results  []*Result      `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	Logs     []string       `protobuf:"bytes,3,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (m *AgentMessage) Reset()      { *m = AgentMessage{} }
func (*AgentMessage) ProtoMessage() {}
func (*AgentMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{0}
}
func (m *AgentMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AgentMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AgentMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AgentMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentMessage.Merge(m, src)
}
func (m *AgentMessage) XXX_Size() int {
	return m.Size()
}
func (m *AgentMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AgentMessage proto.InternalMessageInfo

func (m *AgentMessage) GetMetadata() *AgentMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *AgentMessage) GetResults() []*Result {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *AgentMessage) GetLogs() []string {
	if m != nil {
		return m.Logs
	}
	return nil
}

type AgentMetadata struct {
	AgentID     string `protobuf:"bytes,1,opt,name=agentID,proto3" json:"agentID,omitempty"`
	MachineUUID string `protobuf:"bytes,2,opt,name=machineUUID,proto3" json:"machineUUID,omitempty"`
	SessionID   string `protobuf:"bytes,3,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	Hostname    string `protobuf:"bytes,4,opt,name=hostname,proto3" json:"hostname,omitempty"`
	PrimaryIP   string `protobuf:"bytes,5,opt,name=primaryIP,proto3" json:"primaryIP,omitempty"`
	PrimaryMAC  string `protobuf:"bytes,6,opt,name=primaryMAC,proto3" json:"primaryMAC,omitempty"`
}

func (m *AgentMetadata) Reset()      { *m = AgentMetadata{} }
func (*AgentMetadata) ProtoMessage() {}
func (*AgentMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{1}
}
func (m *AgentMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AgentMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AgentMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AgentMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentMetadata.Merge(m, src)
}
func (m *AgentMetadata) XXX_Size() int {
	return m.Size()
}
func (m *AgentMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_AgentMetadata proto.InternalMessageInfo

func (m *AgentMetadata) GetAgentID() string {
	if m != nil {
		return m.AgentID
	}
	return ""
}

func (m *AgentMetadata) GetMachineUUID() string {
	if m != nil {
		return m.MachineUUID
	}
	return ""
}

func (m *AgentMetadata) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *AgentMetadata) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *AgentMetadata) GetPrimaryIP() string {
	if m != nil {
		return m.PrimaryIP
	}
	return ""
}

func (m *AgentMetadata) GetPrimaryMAC() string {
	if m != nil {
		return m.PrimaryMAC
	}
	return ""
}

type Result struct {
	Id            string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Output        []string         `protobuf:"bytes,2,rep,name=output,proto3" json:"output,omitempty"`
	Error         string           `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	ExecStartTime *types.Timestamp `protobuf:"bytes,4,opt,name=execStartTime,proto3" json:"execStartTime,omitempty"`
	ExecStopTime  *types.Timestamp `protobuf:"bytes,5,opt,name=execStopTime,proto3" json:"execStopTime,omitempty"`
}

func (m *Result) Reset()      { *m = Result{} }
func (*Result) ProtoMessage() {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{2}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Result.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return m.Size()
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Result) GetOutput() []string {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *Result) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Result) GetExecStartTime() *types.Timestamp {
	if m != nil {
		return m.ExecStartTime
	}
	return nil
}

func (m *Result) GetExecStopTime() *types.Timestamp {
	if m != nil {
		return m.ExecStopTime
	}
	return nil
}

func init() {
	proto.RegisterType((*AgentMessage)(nil), "codec.AgentMessage")
	proto.RegisterType((*AgentMetadata)(nil), "codec.AgentMetadata")
	proto.RegisterType((*Result)(nil), "codec.Result")
}

func init() { proto.RegisterFile("agent.proto", fileDescriptor_56ede974c0020f77) }

var fileDescriptor_56ede974c0020f77 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xb1, 0x6e, 0x13, 0x41,
	0x10, 0xbd, 0xb5, 0x63, 0x27, 0x9e, 0x8b, 0x29, 0x56, 0x16, 0x5a, 0x59, 0x68, 0xb0, 0xdc, 0xe0,
	0xea, 0x8c, 0x4c, 0x8b, 0x10, 0x01, 0x37, 0x2e, 0x22, 0xa1, 0x23, 0xf9, 0x80, 0x8d, 0xbd, 0x5c,
	0x4e, 0xf2, 0xdd, 0x9e, 0x76, 0xf7, 0x24, 0x28, 0x90, 0xf8, 0x04, 0x3e, 0x83, 0x4f, 0xa0, 0xa7,
	0x49, 0xe9, 0x32, 0x25, 0xb7, 0x6e, 0x28, 0x53, 0x52, 0x22, 0xef, 0xde, 0x19, 0xbb, 0x4a, 0x37,
	0xef, 0xbd, 0x79, 0x6f, 0x67, 0x76, 0x20, 0xe4, 0x89, 0xc8, 0x4d, 0x54, 0x28, 0x69, 0x24, 0xed,
	0x2c, 0xe5, 0x4a, 0x2c, 0x87, 0xcf, 0x13, 0x29, 0x93, 0xb5, 0x98, 0x3a, 0xf2, 0xa6, 0xfc, 0x34,
	0x35, 0x69, 0x26, 0xb4, 0xe1, 0x59, 0xe1, 0xfb, 0x86, 0x83, 0x44, 0x26, 0xd2, 0x95, 0xd3, 0x5d,
	0xe5, 0xd9, 0xf1, 0x57, 0x38, 0xbf, 0xd8, 0x85, 0x5d, 0x0a, 0xad, 0x79, 0x22, 0xe8, 0x4b, 0x38,
	0xcb, 0x84, 0xe1, 0x2b, 0x6e, 0x38, 0x23, 0x23, 0x32, 0x09, 0x67, 0x83, 0xc8, 0x3d, 0x10, 0xd5,
	0x6d, 0x5e, 0x8b, 0xf7, 0x5d, 0xf4, 0x05, 0x9c, 0x2a, 0xa1, 0xcb, 0xb5, 0xd1, 0xac, 0x35, 0x6a,
	0x4f, 0xc2, 0x59, 0xbf, 0x36, 0xc4, 0x8e, 0x8d, 0x1b, 0x95, 0x52, 0x38, 0x59, 0xcb, 0x44, 0xb3,
	0xf6, 0xa8, 0x3d, 0xe9, 0xc5, 0xae, 0x1e, 0xff, 0x22, 0xd0, 0x3f, 0x0a, 0xa6, 0x0c, 0x4e, 0xdd,
	0x76, 0x8b, 0xb9, 0x7b, 0xbf, 0x17, 0x37, 0x90, 0x8e, 0x20, 0xcc, 0xf8, 0xf2, 0x36, 0xcd, 0xc5,
	0xf5, 0xf5, 0x62, 0xce, 0x5a, 0x4e, 0x3d, 0xa4, 0xe8, 0x33, 0xe8, 0x69, 0xa1, 0x75, 0x2a, 0xf3,
	0xc5, 0x9c, 0xb5, 0x9d, 0xfe, 0x9f, 0xa0, 0x43, 0x38, 0xbb, 0x95, 0xda, 0xe4, 0x3c, 0x13, 0xec,
	0xc4, 0x89, 0x7b, 0xbc, 0x73, 0x16, 0x2a, 0xcd, 0xb8, 0xfa, 0xb2, 0xf8, 0xc0, 0x3a, 0xde, 0xb9,
	0x27, 0x28, 0x02, 0xd4, 0xe0, 0xf2, 0xe2, 0x3d, 0xeb, 0x3a, 0xf9, 0x80, 0x19, 0xdf, 0x11, 0xe8,
	0xfa, 0x6d, 0xe9, 0x13, 0x68, 0xa5, 0xab, 0x7a, 0xf2, 0x56, 0xba, 0xa2, 0x4f, 0xa1, 0x2b, 0x4b,
	0x53, 0x94, 0xc6, 0x7d, 0x4e, 0x2f, 0xae, 0x11, 0x1d, 0x40, 0x47, 0x28, 0x25, 0x55, 0x3d, 0xa6,
	0x07, 0xf4, 0x2d, 0xf4, 0xc5, 0x67, 0xb1, 0xfc, 0x68, 0xb8, 0x32, 0x57, 0x69, 0x3d, 0x67, 0x38,
	0x1b, 0x46, 0xfe, 0xb8, 0x51, 0x73, 0xdc, 0xe8, 0xaa, 0x39, 0x6e, 0x7c, 0x6c, 0xa0, 0x6f, 0xe0,
	0xdc, 0x13, 0xb2, 0x70, 0x01, 0x9d, 0x47, 0x03, 0x8e, 0xfa, 0xdf, 0xbd, 0xde, 0x54, 0x18, 0xdc,
	0x57, 0x18, 0x3c, 0x54, 0x48, 0xfe, 0x56, 0x48, 0xbe, 0x59, 0x24, 0x3f, 0x2c, 0x92, 0x9f, 0x16,
	0xc9, 0x9d, 0x45, 0xb2, 0xb1, 0x48, 0x7e, 0x5b, 0x24, 0x7f, 0x2c, 0x06, 0x0f, 0x16, 0xc9, 0xf7,
	0x2d, 0x06, 0x9b, 0x2d, 0x06, 0xf7, 0x5b, 0x0c, 0x6e, 0xba, 0x2e, 0xff, 0xd5, 0xbf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x31, 0xec, 0x86, 0xa1, 0xa1, 0x02, 0x00, 0x00,
}

func (this *AgentMessage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AgentMessage)
	if !ok {
		that2, ok := that.(AgentMessage)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if len(this.Results) != len(that1.Results) {
		return false
	}
	for i := range this.Results {
		if !this.Results[i].Equal(that1.Results[i]) {
			return false
		}
	}
	if len(this.Logs) != len(that1.Logs) {
		return false
	}
	for i := range this.Logs {
		if this.Logs[i] != that1.Logs[i] {
			return false
		}
	}
	return true
}
func (this *AgentMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AgentMetadata)
	if !ok {
		that2, ok := that.(AgentMetadata)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.AgentID != that1.AgentID {
		return false
	}
	if this.MachineUUID != that1.MachineUUID {
		return false
	}
	if this.SessionID != that1.SessionID {
		return false
	}
	if this.Hostname != that1.Hostname {
		return false
	}
	if this.PrimaryIP != that1.PrimaryIP {
		return false
	}
	if this.PrimaryMAC != that1.PrimaryMAC {
		return false
	}
	return true
}
func (this *Result) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Result)
	if !ok {
		that2, ok := that.(Result)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if len(this.Output) != len(that1.Output) {
		return false
	}
	for i := range this.Output {
		if this.Output[i] != that1.Output[i] {
			return false
		}
	}
	if this.Error != that1.Error {
		return false
	}
	if !this.ExecStartTime.Equal(that1.ExecStartTime) {
		return false
	}
	if !this.ExecStopTime.Equal(that1.ExecStopTime) {
		return false
	}
	return true
}
func (this *AgentMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&codec.AgentMessage{")
	if this.Metadata != nil {
		s = append(s, "Metadata: "+fmt.Sprintf("%#v", this.Metadata)+",\n")
	}
	if this.Results != nil {
		s = append(s, "Results: "+fmt.Sprintf("%#v", this.Results)+",\n")
	}
	s = append(s, "Logs: "+fmt.Sprintf("%#v", this.Logs)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *AgentMetadata) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&codec.AgentMetadata{")
	s = append(s, "AgentID: "+fmt.Sprintf("%#v", this.AgentID)+",\n")
	s = append(s, "MachineUUID: "+fmt.Sprintf("%#v", this.MachineUUID)+",\n")
	s = append(s, "SessionID: "+fmt.Sprintf("%#v", this.SessionID)+",\n")
	s = append(s, "Hostname: "+fmt.Sprintf("%#v", this.Hostname)+",\n")
	s = append(s, "PrimaryIP: "+fmt.Sprintf("%#v", this.PrimaryIP)+",\n")
	s = append(s, "PrimaryMAC: "+fmt.Sprintf("%#v", this.PrimaryMAC)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Result) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&codec.Result{")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "Output: "+fmt.Sprintf("%#v", this.Output)+",\n")
	s = append(s, "Error: "+fmt.Sprintf("%#v", this.Error)+",\n")
	if this.ExecStartTime != nil {
		s = append(s, "ExecStartTime: "+fmt.Sprintf("%#v", this.ExecStartTime)+",\n")
	}
	if this.ExecStopTime != nil {
		s = append(s, "ExecStopTime: "+fmt.Sprintf("%#v", this.ExecStopTime)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringAgent(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *AgentMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AgentMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAgent(dAtA, i, uint64(m.Metadata.Size()))
		n1, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Results) > 0 {
		for _, msg := range m.Results {
			dAtA[i] = 0x12
			i++
			i = encodeVarintAgent(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Logs) > 0 {
		for _, s := range m.Logs {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *AgentMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AgentMetadata) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AgentID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.AgentID)))
		i += copy(dAtA[i:], m.AgentID)
	}
	if len(m.MachineUUID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.MachineUUID)))
		i += copy(dAtA[i:], m.MachineUUID)
	}
	if len(m.SessionID) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.SessionID)))
		i += copy(dAtA[i:], m.SessionID)
	}
	if len(m.Hostname) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.Hostname)))
		i += copy(dAtA[i:], m.Hostname)
	}
	if len(m.PrimaryIP) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.PrimaryIP)))
		i += copy(dAtA[i:], m.PrimaryIP)
	}
	if len(m.PrimaryMAC) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.PrimaryMAC)))
		i += copy(dAtA[i:], m.PrimaryMAC)
	}
	return i, nil
}

func (m *Result) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Result) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Output) > 0 {
		for _, s := range m.Output {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Error) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.Error)))
		i += copy(dAtA[i:], m.Error)
	}
	if m.ExecStartTime != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintAgent(dAtA, i, uint64(m.ExecStartTime.Size()))
		n2, err := m.ExecStartTime.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.ExecStopTime != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintAgent(dAtA, i, uint64(m.ExecStopTime.Size()))
		n3, err := m.ExecStopTime.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func encodeVarintAgent(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedAgentMessage(r randyAgent, easy bool) *AgentMessage {
	this := &AgentMessage{}
	if r.Intn(10) != 0 {
		this.Metadata = NewPopulatedAgentMetadata(r, easy)
	}
	if r.Intn(10) != 0 {
		v1 := r.Intn(5)
		this.Results = make([]*Result, v1)
		for i := 0; i < v1; i++ {
			this.Results[i] = NewPopulatedResult(r, easy)
		}
	}
	v2 := r.Intn(10)
	this.Logs = make([]string, v2)
	for i := 0; i < v2; i++ {
		this.Logs[i] = string(randStringAgent(r))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedAgentMetadata(r randyAgent, easy bool) *AgentMetadata {
	this := &AgentMetadata{}
	this.AgentID = string(randStringAgent(r))
	this.MachineUUID = string(randStringAgent(r))
	this.SessionID = string(randStringAgent(r))
	this.Hostname = string(randStringAgent(r))
	this.PrimaryIP = string(randStringAgent(r))
	this.PrimaryMAC = string(randStringAgent(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedResult(r randyAgent, easy bool) *Result {
	this := &Result{}
	this.Id = string(randStringAgent(r))
	v3 := r.Intn(10)
	this.Output = make([]string, v3)
	for i := 0; i < v3; i++ {
		this.Output[i] = string(randStringAgent(r))
	}
	this.Error = string(randStringAgent(r))
	if r.Intn(10) != 0 {
		this.ExecStartTime = types.NewPopulatedTimestamp(r, easy)
	}
	if r.Intn(10) != 0 {
		this.ExecStopTime = types.NewPopulatedTimestamp(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyAgent interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneAgent(r randyAgent) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringAgent(r randyAgent) string {
	v4 := r.Intn(100)
	tmps := make([]rune, v4)
	for i := 0; i < v4; i++ {
		tmps[i] = randUTF8RuneAgent(r)
	}
	return string(tmps)
}
func randUnrecognizedAgent(r randyAgent, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldAgent(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldAgent(dAtA []byte, r randyAgent, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateAgent(dAtA, uint64(key))
		v5 := r.Int63()
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		dAtA = encodeVarintPopulateAgent(dAtA, uint64(v5))
	case 1:
		dAtA = encodeVarintPopulateAgent(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateAgent(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateAgent(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateAgent(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateAgent(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *AgentMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovAgent(uint64(l))
	}
	if len(m.Results) > 0 {
		for _, e := range m.Results {
			l = e.Size()
			n += 1 + l + sovAgent(uint64(l))
		}
	}
	if len(m.Logs) > 0 {
		for _, s := range m.Logs {
			l = len(s)
			n += 1 + l + sovAgent(uint64(l))
		}
	}
	return n
}

func (m *AgentMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AgentID)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.MachineUUID)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.SessionID)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.Hostname)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.PrimaryIP)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.PrimaryMAC)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	return n
}

func (m *Result) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	if len(m.Output) > 0 {
		for _, s := range m.Output {
			l = len(s)
			n += 1 + l + sovAgent(uint64(l))
		}
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	if m.ExecStartTime != nil {
		l = m.ExecStartTime.Size()
		n += 1 + l + sovAgent(uint64(l))
	}
	if m.ExecStopTime != nil {
		l = m.ExecStopTime.Size()
		n += 1 + l + sovAgent(uint64(l))
	}
	return n
}

func sovAgent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAgent(x uint64) (n int) {
	return sovAgent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *AgentMessage) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AgentMessage{`,
		`Metadata:` + strings.Replace(fmt.Sprintf("%v", this.Metadata), "AgentMetadata", "AgentMetadata", 1) + `,`,
		`Results:` + strings.Replace(fmt.Sprintf("%v", this.Results), "Result", "Result", 1) + `,`,
		`Logs:` + fmt.Sprintf("%v", this.Logs) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AgentMetadata) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AgentMetadata{`,
		`AgentID:` + fmt.Sprintf("%v", this.AgentID) + `,`,
		`MachineUUID:` + fmt.Sprintf("%v", this.MachineUUID) + `,`,
		`SessionID:` + fmt.Sprintf("%v", this.SessionID) + `,`,
		`Hostname:` + fmt.Sprintf("%v", this.Hostname) + `,`,
		`PrimaryIP:` + fmt.Sprintf("%v", this.PrimaryIP) + `,`,
		`PrimaryMAC:` + fmt.Sprintf("%v", this.PrimaryMAC) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Result) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Result{`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`Output:` + fmt.Sprintf("%v", this.Output) + `,`,
		`Error:` + fmt.Sprintf("%v", this.Error) + `,`,
		`ExecStartTime:` + strings.Replace(fmt.Sprintf("%v", this.ExecStartTime), "Timestamp", "types.Timestamp", 1) + `,`,
		`ExecStopTime:` + strings.Replace(fmt.Sprintf("%v", this.ExecStopTime), "Timestamp", "types.Timestamp", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringAgent(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *AgentMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAgent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AgentMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AgentMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &AgentMetadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = append(m.Results, &Result{})
			if err := m.Results[len(m.Results)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Logs = append(m.Logs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAgent
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAgent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AgentMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAgent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AgentMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AgentMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AgentID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AgentID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MachineUUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MachineUUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hostname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hostname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrimaryIP", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrimaryIP = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrimaryMAC", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrimaryMAC = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAgent
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAgent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Result) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAgent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Result: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Result: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Output", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Output = append(m.Output, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecStartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExecStartTime == nil {
				m.ExecStartTime = &types.Timestamp{}
			}
			if err := m.ExecStartTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecStopTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAgent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExecStopTime == nil {
				m.ExecStopTime = &types.Timestamp{}
			}
			if err := m.ExecStopTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAgent
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAgent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAgent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAgent
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAgent
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthAgent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAgent
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAgent(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthAgent
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAgent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAgent   = fmt.Errorf("proto: integer overflow")
)
