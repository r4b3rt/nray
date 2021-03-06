// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schemas/messages.proto

package nraySchema

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Everything a server sends is packed in a nray server message
type NrayServerMessage struct {
	// Types that are valid to be assigned to MessageContent:
	//	*NrayServerMessage_RegisteredNode
	//	*NrayServerMessage_JobBatch
	//	*NrayServerMessage_HeartbeatAck
	//	*NrayServerMessage_WorkDoneAck
	//	*NrayServerMessage_GoodbyeAck
	//	*NrayServerMessage_NodeIsUnregistered
	MessageContent       isNrayServerMessage_MessageContent `protobuf_oneof:"MessageContent"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *NrayServerMessage) Reset()         { *m = NrayServerMessage{} }
func (m *NrayServerMessage) String() string { return proto.CompactTextString(m) }
func (*NrayServerMessage) ProtoMessage()    {}
func (*NrayServerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_1723a75bcb31ddc3, []int{0}
}

func (m *NrayServerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NrayServerMessage.Unmarshal(m, b)
}
func (m *NrayServerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NrayServerMessage.Marshal(b, m, deterministic)
}
func (m *NrayServerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NrayServerMessage.Merge(m, src)
}
func (m *NrayServerMessage) XXX_Size() int {
	return xxx_messageInfo_NrayServerMessage.Size(m)
}
func (m *NrayServerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_NrayServerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_NrayServerMessage proto.InternalMessageInfo

type isNrayServerMessage_MessageContent interface {
	isNrayServerMessage_MessageContent()
}

type NrayServerMessage_RegisteredNode struct {
	RegisteredNode *RegisteredNode `protobuf:"bytes,1,opt,name=registeredNode,proto3,oneof"`
}

type NrayServerMessage_JobBatch struct {
	JobBatch *MoreWorkReply `protobuf:"bytes,2,opt,name=jobBatch,proto3,oneof"`
}

type NrayServerMessage_HeartbeatAck struct {
	HeartbeatAck *HeartbeatAck `protobuf:"bytes,3,opt,name=heartbeatAck,proto3,oneof"`
}

type NrayServerMessage_WorkDoneAck struct {
	WorkDoneAck *WorkDoneAck `protobuf:"bytes,4,opt,name=workDoneAck,proto3,oneof"`
}

type NrayServerMessage_GoodbyeAck struct {
	GoodbyeAck *GoodbyeAck `protobuf:"bytes,5,opt,name=goodbyeAck,proto3,oneof"`
}

type NrayServerMessage_NodeIsUnregistered struct {
	NodeIsUnregistered *Unregistered `protobuf:"bytes,6,opt,name=nodeIsUnregistered,proto3,oneof"`
}

func (*NrayServerMessage_RegisteredNode) isNrayServerMessage_MessageContent() {}

func (*NrayServerMessage_JobBatch) isNrayServerMessage_MessageContent() {}

func (*NrayServerMessage_HeartbeatAck) isNrayServerMessage_MessageContent() {}

func (*NrayServerMessage_WorkDoneAck) isNrayServerMessage_MessageContent() {}

func (*NrayServerMessage_GoodbyeAck) isNrayServerMessage_MessageContent() {}

func (*NrayServerMessage_NodeIsUnregistered) isNrayServerMessage_MessageContent() {}

func (m *NrayServerMessage) GetMessageContent() isNrayServerMessage_MessageContent {
	if m != nil {
		return m.MessageContent
	}
	return nil
}

func (m *NrayServerMessage) GetRegisteredNode() *RegisteredNode {
	if x, ok := m.GetMessageContent().(*NrayServerMessage_RegisteredNode); ok {
		return x.RegisteredNode
	}
	return nil
}

func (m *NrayServerMessage) GetJobBatch() *MoreWorkReply {
	if x, ok := m.GetMessageContent().(*NrayServerMessage_JobBatch); ok {
		return x.JobBatch
	}
	return nil
}

func (m *NrayServerMessage) GetHeartbeatAck() *HeartbeatAck {
	if x, ok := m.GetMessageContent().(*NrayServerMessage_HeartbeatAck); ok {
		return x.HeartbeatAck
	}
	return nil
}

func (m *NrayServerMessage) GetWorkDoneAck() *WorkDoneAck {
	if x, ok := m.GetMessageContent().(*NrayServerMessage_WorkDoneAck); ok {
		return x.WorkDoneAck
	}
	return nil
}

func (m *NrayServerMessage) GetGoodbyeAck() *GoodbyeAck {
	if x, ok := m.GetMessageContent().(*NrayServerMessage_GoodbyeAck); ok {
		return x.GoodbyeAck
	}
	return nil
}

func (m *NrayServerMessage) GetNodeIsUnregistered() *Unregistered {
	if x, ok := m.GetMessageContent().(*NrayServerMessage_NodeIsUnregistered); ok {
		return x.NodeIsUnregistered
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*NrayServerMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*NrayServerMessage_RegisteredNode)(nil),
		(*NrayServerMessage_JobBatch)(nil),
		(*NrayServerMessage_HeartbeatAck)(nil),
		(*NrayServerMessage_WorkDoneAck)(nil),
		(*NrayServerMessage_GoodbyeAck)(nil),
		(*NrayServerMessage_NodeIsUnregistered)(nil),
	}
}

// Everything the nodes sends is packed in a nray node message
type NrayNodeMessage struct {
	// Types that are valid to be assigned to MessageContent:
	//	*NrayNodeMessage_NodeRegister
	//	*NrayNodeMessage_Heartbeat
	//	*NrayNodeMessage_MoreWork
	//	*NrayNodeMessage_WorkDone
	//	*NrayNodeMessage_Goodbye
	MessageContent       isNrayNodeMessage_MessageContent `protobuf_oneof:"MessageContent"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *NrayNodeMessage) Reset()         { *m = NrayNodeMessage{} }
func (m *NrayNodeMessage) String() string { return proto.CompactTextString(m) }
func (*NrayNodeMessage) ProtoMessage()    {}
func (*NrayNodeMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_1723a75bcb31ddc3, []int{1}
}

func (m *NrayNodeMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NrayNodeMessage.Unmarshal(m, b)
}
func (m *NrayNodeMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NrayNodeMessage.Marshal(b, m, deterministic)
}
func (m *NrayNodeMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NrayNodeMessage.Merge(m, src)
}
func (m *NrayNodeMessage) XXX_Size() int {
	return xxx_messageInfo_NrayNodeMessage.Size(m)
}
func (m *NrayNodeMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_NrayNodeMessage.DiscardUnknown(m)
}

var xxx_messageInfo_NrayNodeMessage proto.InternalMessageInfo

type isNrayNodeMessage_MessageContent interface {
	isNrayNodeMessage_MessageContent()
}

type NrayNodeMessage_NodeRegister struct {
	NodeRegister *NodeRegister `protobuf:"bytes,1,opt,name=nodeRegister,proto3,oneof"`
}

type NrayNodeMessage_Heartbeat struct {
	Heartbeat *Heartbeat `protobuf:"bytes,2,opt,name=heartbeat,proto3,oneof"`
}

type NrayNodeMessage_MoreWork struct {
	MoreWork *MoreWorkRequest `protobuf:"bytes,3,opt,name=moreWork,proto3,oneof"`
}

type NrayNodeMessage_WorkDone struct {
	WorkDone *WorkDone `protobuf:"bytes,4,opt,name=workDone,proto3,oneof"`
}

type NrayNodeMessage_Goodbye struct {
	Goodbye *Goodbye `protobuf:"bytes,5,opt,name=goodbye,proto3,oneof"`
}

func (*NrayNodeMessage_NodeRegister) isNrayNodeMessage_MessageContent() {}

func (*NrayNodeMessage_Heartbeat) isNrayNodeMessage_MessageContent() {}

func (*NrayNodeMessage_MoreWork) isNrayNodeMessage_MessageContent() {}

func (*NrayNodeMessage_WorkDone) isNrayNodeMessage_MessageContent() {}

func (*NrayNodeMessage_Goodbye) isNrayNodeMessage_MessageContent() {}

func (m *NrayNodeMessage) GetMessageContent() isNrayNodeMessage_MessageContent {
	if m != nil {
		return m.MessageContent
	}
	return nil
}

func (m *NrayNodeMessage) GetNodeRegister() *NodeRegister {
	if x, ok := m.GetMessageContent().(*NrayNodeMessage_NodeRegister); ok {
		return x.NodeRegister
	}
	return nil
}

func (m *NrayNodeMessage) GetHeartbeat() *Heartbeat {
	if x, ok := m.GetMessageContent().(*NrayNodeMessage_Heartbeat); ok {
		return x.Heartbeat
	}
	return nil
}

func (m *NrayNodeMessage) GetMoreWork() *MoreWorkRequest {
	if x, ok := m.GetMessageContent().(*NrayNodeMessage_MoreWork); ok {
		return x.MoreWork
	}
	return nil
}

func (m *NrayNodeMessage) GetWorkDone() *WorkDone {
	if x, ok := m.GetMessageContent().(*NrayNodeMessage_WorkDone); ok {
		return x.WorkDone
	}
	return nil
}

func (m *NrayNodeMessage) GetGoodbye() *Goodbye {
	if x, ok := m.GetMessageContent().(*NrayNodeMessage_Goodbye); ok {
		return x.Goodbye
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*NrayNodeMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*NrayNodeMessage_NodeRegister)(nil),
		(*NrayNodeMessage_Heartbeat)(nil),
		(*NrayNodeMessage_MoreWork)(nil),
		(*NrayNodeMessage_WorkDone)(nil),
		(*NrayNodeMessage_Goodbye)(nil),
	}
}

// ScanTargets may be dnsNames/ips and ports.
//ip is encoded as byte array.
//dnsName is a FQDN and mustn't contain a protocol specification
type ScanTargets struct {
	Rhosts               []string `protobuf:"bytes,1,rep,name=rhosts,proto3" json:"rhosts,omitempty"`
	Tcpports             []uint32 `protobuf:"varint,2,rep,packed,name=tcpports,proto3" json:"tcpports,omitempty"`
	Udpports             []uint32 `protobuf:"varint,3,rep,packed,name=udpports,proto3" json:"udpports,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanTargets) Reset()         { *m = ScanTargets{} }
func (m *ScanTargets) String() string { return proto.CompactTextString(m) }
func (*ScanTargets) ProtoMessage()    {}
func (*ScanTargets) Descriptor() ([]byte, []int) {
	return fileDescriptor_1723a75bcb31ddc3, []int{2}
}

func (m *ScanTargets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanTargets.Unmarshal(m, b)
}
func (m *ScanTargets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanTargets.Marshal(b, m, deterministic)
}
func (m *ScanTargets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanTargets.Merge(m, src)
}
func (m *ScanTargets) XXX_Size() int {
	return xxx_messageInfo_ScanTargets.Size(m)
}
func (m *ScanTargets) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanTargets.DiscardUnknown(m)
}

var xxx_messageInfo_ScanTargets proto.InternalMessageInfo

func (m *ScanTargets) GetRhosts() []string {
	if m != nil {
		return m.Rhosts
	}
	return nil
}

func (m *ScanTargets) GetTcpports() []uint32 {
	if m != nil {
		return m.Tcpports
	}
	return nil
}

func (m *ScanTargets) GetUdpports() []uint32 {
	if m != nil {
		return m.Udpports
	}
	return nil
}

// This message is sent every time a node registers at
//the server. It contains a unique node ID so there are
//not multiple scanner nodes running on the same machine
type NodeRegister struct {
	MachineID            string   `protobuf:"bytes,1,opt,name=machineID,proto3" json:"machineID,omitempty"`
	PreferredPool        int32    `protobuf:"varint,2,opt,name=preferredPool,proto3" json:"preferredPool,omitempty"`
	PreferredNodeName    string   `protobuf:"bytes,3,opt,name=preferredNodeName,proto3" json:"preferredNodeName,omitempty"`
	Envinfo              *Event   `protobuf:"bytes,4,opt,name=envinfo,proto3" json:"envinfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeRegister) Reset()         { *m = NodeRegister{} }
func (m *NodeRegister) String() string { return proto.CompactTextString(m) }
func (*NodeRegister) ProtoMessage()    {}
func (*NodeRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_1723a75bcb31ddc3, []int{3}
}

func (m *NodeRegister) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRegister.Unmarshal(m, b)
}
func (m *NodeRegister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRegister.Marshal(b, m, deterministic)
}
func (m *NodeRegister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRegister.Merge(m, src)
}
func (m *NodeRegister) XXX_Size() int {
	return xxx_messageInfo_NodeRegister.Size(m)
}
func (m *NodeRegister) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRegister.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRegister proto.InternalMessageInfo

func (m *NodeRegister) GetMachineID() string {
	if m != nil {
		return m.MachineID
	}
	return ""
}

func (m *NodeRegister) GetPreferredPool() int32 {
	if m != nil {
		return m.PreferredPool
	}
	return 0
}

func (m *NodeRegister) GetPreferredNodeName() string {
	if m != nil {
		return m.PreferredNodeName
	}
	return ""
}

func (m *NodeRegister) GetEnvinfo() *Event {
	if m != nil {
		return m.Envinfo
	}
	return nil
}

// This message is sent by the server and indicates that
//the server does not know this node and that the node should
//register again
type Unregistered struct {
	NodeID               string   `protobuf:"bytes,1,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Unregistered) Reset()         { *m = Unregistered{} }
func (m *Unregistered) String() string { return proto.CompactTextString(m) }
func (*Unregistered) ProtoMessage()    {}
func (*Unregistered) Descriptor() ([]byte, []int) {
	return fileDescriptor_1723a75bcb31ddc3, []int{4}
}

func (m *Unregistered) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Unregistered.Unmarshal(m, b)
}
func (m *Unregistered) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Unregistered.Marshal(b, m, deterministic)
}
func (m *Unregistered) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unregistered.Merge(m, src)
}
func (m *Unregistered) XXX_Size() int {
	return xxx_messageInfo_Unregistered.Size(m)
}
func (m *Unregistered) XXX_DiscardUnknown() {
	xxx_messageInfo_Unregistered.DiscardUnknown(m)
}

var xxx_messageInfo_Unregistered proto.InternalMessageInfo

func (m *Unregistered) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

// This message tells the scanner that it is registered at
//the server and assigns a unique scanner ID as well as it carries
//the timestamp of the server, so nodes that have no correct clock
//can still perform somewhat accurate timestamping of events
type RegisteredNode struct {
	NodeID      string               `protobuf:"bytes,1,opt,name=NodeID,proto3" json:"NodeID,omitempty"`
	ServerClock *timestamp.Timestamp `protobuf:"bytes,2,opt,name=ServerClock,proto3" json:"ServerClock,omitempty"`
	//int32 pool = 3;
	Scannerconfig        []byte   `protobuf:"bytes,4,opt,name=scannerconfig,proto3" json:"scannerconfig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisteredNode) Reset()         { *m = RegisteredNode{} }
func (m *RegisteredNode) String() string { return proto.CompactTextString(m) }
func (*RegisteredNode) ProtoMessage()    {}
func (*RegisteredNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_1723a75bcb31ddc3, []int{5}
}

func (m *RegisteredNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisteredNode.Unmarshal(m, b)
}
func (m *RegisteredNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisteredNode.Marshal(b, m, deterministic)
}
func (m *RegisteredNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisteredNode.Merge(m, src)
}
func (m *RegisteredNode) XXX_Size() int {
	return xxx_messageInfo_RegisteredNode.Size(m)
}
func (m *RegisteredNode) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisteredNode.DiscardUnknown(m)
}

var xxx_messageInfo_RegisteredNode proto.InternalMessageInfo

func (m *RegisteredNode) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

func (m *RegisteredNode) GetServerClock() *timestamp.Timestamp {
	if m != nil {
		return m.ServerClock
	}
	return nil
}

func (m *RegisteredNode) GetScannerconfig() []byte {
	if m != nil {
		return m.Scannerconfig
	}
	return nil
}

// A heartbeat message that is sent regularly from any node
//to the server to signal that it is still alive
type Heartbeat struct {
	NodeID               string               `protobuf:"bytes,1,opt,name=NodeID,proto3" json:"NodeID,omitempty"`
	BeatTime             *timestamp.Timestamp `protobuf:"bytes,2,opt,name=BeatTime,proto3" json:"BeatTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Heartbeat) Reset()         { *m = Heartbeat{} }
func (m *Heartbeat) String() string { return proto.CompactTextString(m) }
func (*Heartbeat) ProtoMessage()    {}
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_1723a75bcb31ddc3, []int{6}
}

func (m *Heartbeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Heartbeat.Unmarshal(m, b)
}
func (m *Heartbeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Heartbeat.Marshal(b, m, deterministic)
}
func (m *Heartbeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Heartbeat.Merge(m, src)
}
func (m *Heartbeat) XXX_Size() int {
	return xxx_messageInfo_Heartbeat.Size(m)
}
func (m *Heartbeat) XXX_DiscardUnknown() {
	xxx_messageInfo_Heartbeat.DiscardUnknown(m)
}

var xxx_messageInfo_Heartbeat proto.InternalMessageInfo

func (m *Heartbeat) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

func (m *Heartbeat) GetBeatTime() *timestamp.Timestamp {
	if m != nil {
		return m.BeatTime
	}
	return nil
}

// Acknowledgement of a heartbeat message. A server
//may indicate to stop scanning and if the scanner should
//exit
type HeartbeatAck struct {
	Scanning             bool     `protobuf:"varint,1,opt,name=Scanning,proto3" json:"Scanning,omitempty"`
	Running              bool     `protobuf:"varint,2,opt,name=Running,proto3" json:"Running,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatAck) Reset()         { *m = HeartbeatAck{} }
func (m *HeartbeatAck) String() string { return proto.CompactTextString(m) }
func (*HeartbeatAck) ProtoMessage()    {}
func (*HeartbeatAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_1723a75bcb31ddc3, []int{7}
}

func (m *HeartbeatAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatAck.Unmarshal(m, b)
}
func (m *HeartbeatAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatAck.Marshal(b, m, deterministic)
}
func (m *HeartbeatAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatAck.Merge(m, src)
}
func (m *HeartbeatAck) XXX_Size() int {
	return xxx_messageInfo_HeartbeatAck.Size(m)
}
func (m *HeartbeatAck) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatAck.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatAck proto.InternalMessageInfo

func (m *HeartbeatAck) GetScanning() bool {
	if m != nil {
		return m.Scanning
	}
	return false
}

func (m *HeartbeatAck) GetRunning() bool {
	if m != nil {
		return m.Running
	}
	return false
}

// Sent by a node if it has no work
type MoreWorkRequest struct {
	NodeID               string   `protobuf:"bytes,1,opt,name=NodeID,proto3" json:"NodeID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MoreWorkRequest) Reset()         { *m = MoreWorkRequest{} }
func (m *MoreWorkRequest) String() string { return proto.CompactTextString(m) }
func (*MoreWorkRequest) ProtoMessage()    {}
func (*MoreWorkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1723a75bcb31ddc3, []int{8}
}

func (m *MoreWorkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoreWorkRequest.Unmarshal(m, b)
}
func (m *MoreWorkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoreWorkRequest.Marshal(b, m, deterministic)
}
func (m *MoreWorkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoreWorkRequest.Merge(m, src)
}
func (m *MoreWorkRequest) XXX_Size() int {
	return xxx_messageInfo_MoreWorkRequest.Size(m)
}
func (m *MoreWorkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MoreWorkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MoreWorkRequest proto.InternalMessageInfo

func (m *MoreWorkRequest) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

// Contains more work for a scanner
type MoreWorkReply struct {
	Batchid              uint64       `protobuf:"varint,1,opt,name=batchid,proto3" json:"batchid,omitempty"`
	Targets              *ScanTargets `protobuf:"bytes,3,opt,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MoreWorkReply) Reset()         { *m = MoreWorkReply{} }
func (m *MoreWorkReply) String() string { return proto.CompactTextString(m) }
func (*MoreWorkReply) ProtoMessage()    {}
func (*MoreWorkReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_1723a75bcb31ddc3, []int{9}
}

func (m *MoreWorkReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoreWorkReply.Unmarshal(m, b)
}
func (m *MoreWorkReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoreWorkReply.Marshal(b, m, deterministic)
}
func (m *MoreWorkReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoreWorkReply.Merge(m, src)
}
func (m *MoreWorkReply) XXX_Size() int {
	return xxx_messageInfo_MoreWorkReply.Size(m)
}
func (m *MoreWorkReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MoreWorkReply.DiscardUnknown(m)
}

var xxx_messageInfo_MoreWorkReply proto.InternalMessageInfo

func (m *MoreWorkReply) GetBatchid() uint64 {
	if m != nil {
		return m.Batchid
	}
	return 0
}

func (m *MoreWorkReply) GetTargets() *ScanTargets {
	if m != nil {
		return m.Targets
	}
	return nil
}

// Indicates that a node is done with a work batch
//and contains the results
type WorkDone struct {
	NodeID               string   `protobuf:"bytes,1,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	Batchid              uint64   `protobuf:"varint,2,opt,name=batchid,proto3" json:"batchid,omitempty"`
	Events               []*Event `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkDone) Reset()         { *m = WorkDone{} }
func (m *WorkDone) String() string { return proto.CompactTextString(m) }
func (*WorkDone) ProtoMessage()    {}
func (*WorkDone) Descriptor() ([]byte, []int) {
	return fileDescriptor_1723a75bcb31ddc3, []int{10}
}

func (m *WorkDone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkDone.Unmarshal(m, b)
}
func (m *WorkDone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkDone.Marshal(b, m, deterministic)
}
func (m *WorkDone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkDone.Merge(m, src)
}
func (m *WorkDone) XXX_Size() int {
	return xxx_messageInfo_WorkDone.Size(m)
}
func (m *WorkDone) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkDone.DiscardUnknown(m)
}

var xxx_messageInfo_WorkDone proto.InternalMessageInfo

func (m *WorkDone) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

func (m *WorkDone) GetBatchid() uint64 {
	if m != nil {
		return m.Batchid
	}
	return 0
}

func (m *WorkDone) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

// Acknowledges a WorkDone packet
type WorkDoneAck struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkDoneAck) Reset()         { *m = WorkDoneAck{} }
func (m *WorkDoneAck) String() string { return proto.CompactTextString(m) }
func (*WorkDoneAck) ProtoMessage()    {}
func (*WorkDoneAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_1723a75bcb31ddc3, []int{11}
}

func (m *WorkDoneAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkDoneAck.Unmarshal(m, b)
}
func (m *WorkDoneAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkDoneAck.Marshal(b, m, deterministic)
}
func (m *WorkDoneAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkDoneAck.Merge(m, src)
}
func (m *WorkDoneAck) XXX_Size() int {
	return xxx_messageInfo_WorkDoneAck.Size(m)
}
func (m *WorkDoneAck) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkDoneAck.DiscardUnknown(m)
}

var xxx_messageInfo_WorkDoneAck proto.InternalMessageInfo

// Node is going to exit
type Goodbye struct {
	NodeID               string   `protobuf:"bytes,1,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Goodbye) Reset()         { *m = Goodbye{} }
func (m *Goodbye) String() string { return proto.CompactTextString(m) }
func (*Goodbye) ProtoMessage()    {}
func (*Goodbye) Descriptor() ([]byte, []int) {
	return fileDescriptor_1723a75bcb31ddc3, []int{12}
}

func (m *Goodbye) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Goodbye.Unmarshal(m, b)
}
func (m *Goodbye) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Goodbye.Marshal(b, m, deterministic)
}
func (m *Goodbye) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Goodbye.Merge(m, src)
}
func (m *Goodbye) XXX_Size() int {
	return xxx_messageInfo_Goodbye.Size(m)
}
func (m *Goodbye) XXX_DiscardUnknown() {
	xxx_messageInfo_Goodbye.DiscardUnknown(m)
}

var xxx_messageInfo_Goodbye proto.InternalMessageInfo

func (m *Goodbye) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

// Server ACKs the goodbye and signals if
//the node is allowed to exit
type GoodbyeAck struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoodbyeAck) Reset()         { *m = GoodbyeAck{} }
func (m *GoodbyeAck) String() string { return proto.CompactTextString(m) }
func (*GoodbyeAck) ProtoMessage()    {}
func (*GoodbyeAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_1723a75bcb31ddc3, []int{13}
}

func (m *GoodbyeAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoodbyeAck.Unmarshal(m, b)
}
func (m *GoodbyeAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoodbyeAck.Marshal(b, m, deterministic)
}
func (m *GoodbyeAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodbyeAck.Merge(m, src)
}
func (m *GoodbyeAck) XXX_Size() int {
	return xxx_messageInfo_GoodbyeAck.Size(m)
}
func (m *GoodbyeAck) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodbyeAck.DiscardUnknown(m)
}

var xxx_messageInfo_GoodbyeAck proto.InternalMessageInfo

func (m *GoodbyeAck) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*NrayServerMessage)(nil), "nraySchema.NrayServerMessage")
	proto.RegisterType((*NrayNodeMessage)(nil), "nraySchema.NrayNodeMessage")
	proto.RegisterType((*ScanTargets)(nil), "nraySchema.ScanTargets")
	proto.RegisterType((*NodeRegister)(nil), "nraySchema.NodeRegister")
	proto.RegisterType((*Unregistered)(nil), "nraySchema.Unregistered")
	proto.RegisterType((*RegisteredNode)(nil), "nraySchema.RegisteredNode")
	proto.RegisterType((*Heartbeat)(nil), "nraySchema.Heartbeat")
	proto.RegisterType((*HeartbeatAck)(nil), "nraySchema.HeartbeatAck")
	proto.RegisterType((*MoreWorkRequest)(nil), "nraySchema.MoreWorkRequest")
	proto.RegisterType((*MoreWorkReply)(nil), "nraySchema.MoreWorkReply")
	proto.RegisterType((*WorkDone)(nil), "nraySchema.WorkDone")
	proto.RegisterType((*WorkDoneAck)(nil), "nraySchema.WorkDoneAck")
	proto.RegisterType((*Goodbye)(nil), "nraySchema.Goodbye")
	proto.RegisterType((*GoodbyeAck)(nil), "nraySchema.GoodbyeAck")
}

func init() { proto.RegisterFile("schemas/messages.proto", fileDescriptor_1723a75bcb31ddc3) }

var fileDescriptor_1723a75bcb31ddc3 = []byte{
	// 743 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x6e, 0x13, 0x3d,
	0x10, 0xdd, 0x24, 0x6d, 0x7e, 0x26, 0x3f, 0xfd, 0xea, 0xaf, 0x94, 0x25, 0x54, 0xa2, 0xac, 0x10,
	0x6a, 0x05, 0x4a, 0x44, 0x11, 0x7f, 0x02, 0x55, 0xa2, 0x0d, 0x62, 0x8b, 0xd4, 0x08, 0xb9, 0x45,
	0xbd, 0x00, 0x2e, 0x36, 0x1b, 0x67, 0x13, 0x92, 0xd8, 0xc1, 0xeb, 0x14, 0xf5, 0x1d, 0x78, 0x0f,
	0xae, 0x78, 0x01, 0x9e, 0x0e, 0xd9, 0x6b, 0x6f, 0xbc, 0x6d, 0x22, 0x71, 0x39, 0x7b, 0xce, 0xf1,
	0x38, 0xe7, 0x8c, 0x27, 0xb0, 0x1d, 0x87, 0x43, 0x32, 0x0d, 0xe2, 0xf6, 0x94, 0xc4, 0x71, 0x10,
	0x91, 0xb8, 0x35, 0xe3, 0x4c, 0x30, 0x04, 0x94, 0x07, 0x57, 0x67, 0x0a, 0x6b, 0xde, 0x8b, 0x18,
	0x8b, 0x26, 0xa4, 0xad, 0x90, 0xde, 0x7c, 0xd0, 0x16, 0xa3, 0x29, 0x89, 0x45, 0x30, 0x9d, 0x25,
	0xe4, 0xe6, 0x96, 0x39, 0x84, 0x5c, 0x12, 0x2a, 0xf4, 0x11, 0xde, 0xaf, 0x02, 0x6c, 0x76, 0xe5,
	0x29, 0x84, 0x5f, 0x12, 0x7e, 0x9a, 0x9c, 0x8f, 0x3a, 0xd0, 0xe0, 0x24, 0x1a, 0xc5, 0x82, 0x70,
	0xd2, 0xef, 0xb2, 0x3e, 0x71, 0x73, 0xbb, 0xb9, 0xbd, 0xea, 0x41, 0xb3, 0xb5, 0xe8, 0xd8, 0xc2,
	0x19, 0x86, 0xef, 0xe0, 0x6b, 0x1a, 0xf4, 0x02, 0xca, 0xdf, 0x58, 0xef, 0x28, 0x10, 0xe1, 0xd0,
	0xcd, 0x2b, 0xfd, 0x1d, 0x5b, 0x7f, 0xca, 0x38, 0xb9, 0x60, 0x7c, 0x8c, 0xc9, 0x6c, 0x72, 0xe5,
	0x3b, 0x38, 0x25, 0xa3, 0x43, 0xa8, 0x0d, 0x49, 0xc0, 0x45, 0x8f, 0x04, 0xe2, 0x6d, 0x38, 0x76,
	0x0b, 0x4a, 0xec, 0xda, 0x62, 0xdf, 0xc2, 0x7d, 0x07, 0x67, 0xf8, 0xe8, 0x35, 0x54, 0x7f, 0x30,
	0x3e, 0xee, 0x30, 0x4a, 0xa4, 0x7c, 0x4d, 0xc9, 0x6f, 0xdb, 0xf2, 0x8b, 0x05, 0xec, 0x3b, 0xd8,
	0x66, 0xa3, 0x97, 0x00, 0x11, 0x63, 0xfd, 0xde, 0x95, 0xd2, 0xae, 0x2b, 0xed, 0xb6, 0xad, 0x7d,
	0x9f, 0xa2, 0xbe, 0x83, 0x2d, 0x2e, 0xfa, 0x00, 0x88, 0xb2, 0x3e, 0x39, 0x89, 0x3f, 0xd1, 0x85,
	0x13, 0x6e, 0xf1, 0xe6, 0xe5, 0x6d, 0xdc, 0x77, 0xf0, 0x12, 0xd5, 0xd1, 0x7f, 0xd0, 0xd0, 0x61,
	0x1c, 0x33, 0x2a, 0x08, 0x15, 0xde, 0x9f, 0x3c, 0x6c, 0xc8, 0xa4, 0xa4, 0xb5, 0x26, 0xa7, 0x43,
	0xa8, 0x49, 0xad, 0x49, 0x42, 0xa7, 0x94, 0xe9, 0xd5, 0xb5, 0x70, 0x69, 0x94, 0xcd, 0x47, 0xcf,
	0xa0, 0x92, 0x1a, 0xa7, 0x23, 0xba, 0xb5, 0xd4, 0x65, 0xdf, 0xc1, 0x0b, 0x26, 0x7a, 0x05, 0xe5,
	0xa9, 0x0e, 0x4f, 0x67, 0x73, 0x77, 0x79, 0xb0, 0xdf, 0xe7, 0x24, 0x96, 0xda, 0x94, 0x8e, 0x0e,
	0xa0, 0x6c, 0xcc, 0xd6, 0xb9, 0x6c, 0x2d, 0xcb, 0x45, 0x6a, 0x0c, 0x0f, 0xb5, 0xa1, 0xa4, 0x5d,
	0xd6, 0x71, 0xfc, 0xbf, 0x24, 0x0e, 0xdf, 0xc1, 0x86, 0xb5, 0xc4, 0xbc, 0xaf, 0x50, 0x3d, 0x0b,
	0x03, 0x7a, 0x1e, 0xf0, 0x88, 0x88, 0x18, 0x6d, 0x43, 0x91, 0x0f, 0x59, 0x2c, 0x62, 0x37, 0xb7,
	0x5b, 0xd8, 0xab, 0x60, 0x5d, 0xa1, 0x26, 0x94, 0x45, 0x38, 0x9b, 0x31, 0x2e, 0x62, 0x37, 0xbf,
	0x5b, 0xd8, 0xab, 0xe3, 0xb4, 0x96, 0xd8, 0xbc, 0xaf, 0xb1, 0x42, 0x82, 0x99, 0xda, 0xfb, 0x9d,
	0x83, 0x9a, 0x6d, 0x34, 0xda, 0x81, 0xca, 0x34, 0x08, 0x87, 0x23, 0x4a, 0x4e, 0x3a, 0x2a, 0x95,
	0x0a, 0x5e, 0x7c, 0x40, 0x0f, 0xa0, 0x3e, 0xe3, 0x64, 0x40, 0x38, 0x27, 0xfd, 0x8f, 0x8c, 0x4d,
	0x94, 0xf5, 0xeb, 0x38, 0xfb, 0x11, 0x3d, 0x86, 0xcd, 0xf4, 0x83, 0x3c, 0xbc, 0x1b, 0x4c, 0x89,
	0xb2, 0xbb, 0x82, 0x6f, 0x02, 0xe8, 0x11, 0x94, 0x08, 0xbd, 0x1c, 0xd1, 0x01, 0xd3, 0xbe, 0x6e,
	0xda, 0x26, 0xbd, 0x93, 0x6f, 0x1e, 0x1b, 0x86, 0xf7, 0x10, 0x6a, 0xf6, 0xb4, 0x49, 0x3f, 0xd4,
	0x0c, 0x9a, 0xbb, 0xea, 0xca, 0xfb, 0x99, 0x83, 0x46, 0xf6, 0x99, 0x4b, 0x6a, 0x37, 0x43, 0x4d,
	0x2a, 0xf4, 0x06, 0xaa, 0xc9, 0x0e, 0x39, 0x9e, 0xb0, 0x70, 0xac, 0x87, 0xa9, 0xd9, 0x4a, 0xb6,
	0x52, 0xcb, 0x6c, 0xa5, 0xd6, 0xb9, 0xd9, 0x4a, 0xd8, 0xa6, 0x4b, 0x47, 0xe2, 0x30, 0xa0, 0x94,
	0xf0, 0x90, 0xd1, 0xc1, 0x28, 0x52, 0xbf, 0xa1, 0x86, 0xb3, 0x1f, 0xbd, 0xcf, 0x50, 0x49, 0x27,
	0x72, 0xe5, 0x45, 0x9e, 0x43, 0xf9, 0x88, 0x04, 0x42, 0x36, 0xfa, 0x87, 0x5b, 0xa4, 0x5c, 0xaf,
	0x03, 0x35, 0x7b, 0xa9, 0xc8, 0xbc, 0xe5, 0xc8, 0xd0, 0x11, 0x8d, 0x54, 0x87, 0x32, 0x4e, 0x6b,
	0xe4, 0x42, 0x09, 0xcf, 0x13, 0x28, 0xaf, 0x20, 0x53, 0x7a, 0xfb, 0xb0, 0x71, 0x6d, 0xfc, 0x57,
	0x5d, 0xd4, 0xfb, 0x02, 0xf5, 0xcc, 0x0a, 0x94, 0xa7, 0xf6, 0xe4, 0xfe, 0x1b, 0xf5, 0x15, 0x73,
	0x0d, 0x9b, 0x12, 0x3d, 0x81, 0x92, 0x48, 0x46, 0x57, 0xbf, 0xb7, 0xcc, 0x32, 0xb3, 0x26, 0x1b,
	0x1b, 0x9e, 0x17, 0x41, 0xd9, 0x3c, 0xa6, 0x55, 0xf1, 0xda, 0x0d, 0xf3, 0xd9, 0x86, 0xfb, 0x50,
	0x4c, 0xfe, 0x26, 0xd4, 0xa8, 0x2f, 0x1d, 0x26, 0x4d, 0xf0, 0xea, 0x50, 0xb5, 0xb6, 0xa9, 0x77,
	0x1f, 0x4a, 0xfa, 0x45, 0xae, 0x9c, 0xaa, 0x1d, 0x80, 0xc5, 0x0e, 0x45, 0x0d, 0xc8, 0xb3, 0xb1,
	0x76, 0x38, 0xcf, 0xc6, 0xbd, 0xa2, 0x4a, 0xe9, 0xe9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb5,
	0x46, 0x96, 0x20, 0xf5, 0x06, 0x00, 0x00,
}
