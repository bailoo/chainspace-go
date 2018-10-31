// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types.proto

package sbac

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Opcode int32

const (
	Opcode_UNKNOWN         Opcode = 0
	Opcode_ADD_TRANSACTION Opcode = 1
	Opcode_QUERY_OBJECT    Opcode = 2
	Opcode_CREATE_OBJECT   Opcode = 3
	Opcode_STATES          Opcode = 4
	Opcode_SBAC            Opcode = 5
	Opcode_CREATE_OBJECTS  Opcode = 6
)

var Opcode_name = map[int32]string{
	0: "UNKNOWN",
	1: "ADD_TRANSACTION",
	2: "QUERY_OBJECT",
	3: "CREATE_OBJECT",
	4: "STATES",
	5: "SBAC",
	6: "CREATE_OBJECTS",
}
var Opcode_value = map[string]int32{
	"UNKNOWN":         0,
	"ADD_TRANSACTION": 1,
	"QUERY_OBJECT":    2,
	"CREATE_OBJECT":   3,
	"STATES":          4,
	"SBAC":            5,
	"CREATE_OBJECTS":  6,
}

func (x Opcode) String() string {
	return proto.EnumName(Opcode_name, int32(x))
}
func (Opcode) EnumDescriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

type ObjectStatus int32

const (
	ObjectStatus_ACTIVE   ObjectStatus = 0
	ObjectStatus_INACTIVE ObjectStatus = 1
	ObjectStatus_LOCKED   ObjectStatus = 2
)

var ObjectStatus_name = map[int32]string{
	0: "ACTIVE",
	1: "INACTIVE",
	2: "LOCKED",
}
var ObjectStatus_value = map[string]int32{
	"ACTIVE":   0,
	"INACTIVE": 1,
	"LOCKED":   2,
}

func (x ObjectStatus) String() string {
	return proto.EnumName(ObjectStatus_name, int32(x))
}
func (ObjectStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptorTypes, []int{1} }

type SBACOp int32

const (
	SBACOp_Phase1 SBACOp = 0
	SBACOp_Phase2 SBACOp = 1
	SBACOp_Commit SBACOp = 2
)

var SBACOp_name = map[int32]string{
	0: "Phase1",
	1: "Phase2",
	2: "Commit",
}
var SBACOp_value = map[string]int32{
	"Phase1": 0,
	"Phase2": 1,
	"Commit": 2,
}

func (x SBACOp) String() string {
	return proto.EnumName(SBACOp_name, int32(x))
}
func (SBACOp) EnumDescriptor() ([]byte, []int) { return fileDescriptorTypes, []int{2} }

type SBACDecision int32

const (
	SBACDecision_ACCEPT SBACDecision = 0
	SBACDecision_REJECT SBACDecision = 1
)

var SBACDecision_name = map[int32]string{
	0: "ACCEPT",
	1: "REJECT",
}
var SBACDecision_value = map[string]int32{
	"ACCEPT": 0,
	"REJECT": 1,
}

func (x SBACDecision) String() string {
	return proto.EnumName(SBACDecision_name, int32(x))
}
func (SBACDecision) EnumDescriptor() ([]byte, []int) { return fileDescriptorTypes, []int{3} }

type ConsensusOp int32

const (
	ConsensusOp_Consensus1      ConsensusOp = 0
	ConsensusOp_Consensus2      ConsensusOp = 1
	ConsensusOp_ConsensusCommit ConsensusOp = 2
)

var ConsensusOp_name = map[int32]string{
	0: "Consensus1",
	1: "Consensus2",
	2: "ConsensusCommit",
}
var ConsensusOp_value = map[string]int32{
	"Consensus1":      0,
	"Consensus2":      1,
	"ConsensusCommit": 2,
}

func (x ConsensusOp) String() string {
	return proto.EnumName(ConsensusOp_name, int32(x))
}
func (ConsensusOp) EnumDescriptor() ([]byte, []int) { return fileDescriptorTypes, []int{4} }

type Transaction struct {
	Traces []*Trace `protobuf:"bytes,1,rep,name=traces" json:"traces,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

func (m *Transaction) GetTraces() []*Trace {
	if m != nil {
		return m.Traces
	}
	return nil
}

type Strings struct {
	Strs []string `protobuf:"bytes,1,rep,name=strs" json:"strs,omitempty"`
}

func (m *Strings) Reset()                    { *m = Strings{} }
func (m *Strings) String() string            { return proto.CompactTextString(m) }
func (*Strings) ProtoMessage()               {}
func (*Strings) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{1} }

func (m *Strings) GetStrs() []string {
	if m != nil {
		return m.Strs
	}
	return nil
}

type Trace struct {
	ContractID               string     `protobuf:"bytes,1,opt,name=contractID,proto3" json:"contractID,omitempty"`
	Procedure                string     `protobuf:"bytes,2,opt,name=procedure,proto3" json:"procedure,omitempty"`
	InputObjectVersionIDs    [][]byte   `protobuf:"bytes,3,rep,name=inputObjectVersionIDs" json:"inputObjectVersionIDs,omitempty"`
	InputReferenceVersionIDs [][]byte   `protobuf:"bytes,4,rep,name=inputReferenceVersionIDs" json:"inputReferenceVersionIDs,omitempty"`
	OutputObjects            [][]byte   `protobuf:"bytes,5,rep,name=outputObjects" json:"outputObjects,omitempty"`
	Parameters               [][]byte   `protobuf:"bytes,6,rep,name=parameters" json:"parameters,omitempty"`
	Returns                  [][]byte   `protobuf:"bytes,7,rep,name=returns" json:"returns,omitempty"`
	Labels                   []*Strings `protobuf:"bytes,8,rep,name=labels" json:"labels,omitempty"`
	Dependencies             []*Trace   `protobuf:"bytes,9,rep,name=dependencies" json:"dependencies,omitempty"`
	InputObjects             [][]byte   `protobuf:"bytes,10,rep,name=inputObjects" json:"inputObjects,omitempty"`
	InputReferences          [][]byte   `protobuf:"bytes,11,rep,name=inputReferences" json:"inputReferences,omitempty"`
}

func (m *Trace) Reset()                    { *m = Trace{} }
func (m *Trace) String() string            { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()               {}
func (*Trace) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{2} }

func (m *Trace) GetContractID() string {
	if m != nil {
		return m.ContractID
	}
	return ""
}

func (m *Trace) GetProcedure() string {
	if m != nil {
		return m.Procedure
	}
	return ""
}

func (m *Trace) GetInputObjectVersionIDs() [][]byte {
	if m != nil {
		return m.InputObjectVersionIDs
	}
	return nil
}

func (m *Trace) GetInputReferenceVersionIDs() [][]byte {
	if m != nil {
		return m.InputReferenceVersionIDs
	}
	return nil
}

func (m *Trace) GetOutputObjects() [][]byte {
	if m != nil {
		return m.OutputObjects
	}
	return nil
}

func (m *Trace) GetParameters() [][]byte {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *Trace) GetReturns() [][]byte {
	if m != nil {
		return m.Returns
	}
	return nil
}

func (m *Trace) GetLabels() []*Strings {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Trace) GetDependencies() []*Trace {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *Trace) GetInputObjects() [][]byte {
	if m != nil {
		return m.InputObjects
	}
	return nil
}

func (m *Trace) GetInputReferences() [][]byte {
	if m != nil {
		return m.InputReferences
	}
	return nil
}

type Object struct {
	VersionID []byte       `protobuf:"bytes,1,opt,name=versionID,proto3" json:"versionID,omitempty"`
	Value     []byte       `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Status    ObjectStatus `protobuf:"varint,3,opt,name=status,proto3,enum=sbac.ObjectStatus" json:"status,omitempty"`
	Labels    []string     `protobuf:"bytes,4,rep,name=labels" json:"labels,omitempty"`
}

func (m *Object) Reset()                    { *m = Object{} }
func (m *Object) String() string            { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()               {}
func (*Object) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{3} }

func (m *Object) GetVersionID() []byte {
	if m != nil {
		return m.VersionID
	}
	return nil
}

func (m *Object) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Object) GetStatus() ObjectStatus {
	if m != nil {
		return m.Status
	}
	return ObjectStatus_ACTIVE
}

func (m *Object) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type AddTransactionRequest struct {
	Tx        *Transaction      `protobuf:"bytes,1,opt,name=tx" json:"tx,omitempty"`
	Evidences map[uint64][]byte `protobuf:"bytes,2,rep,name=evidences" json:"evidences,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *AddTransactionRequest) Reset()                    { *m = AddTransactionRequest{} }
func (m *AddTransactionRequest) String() string            { return proto.CompactTextString(m) }
func (*AddTransactionRequest) ProtoMessage()               {}
func (*AddTransactionRequest) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{4} }

func (m *AddTransactionRequest) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *AddTransactionRequest) GetEvidences() map[uint64][]byte {
	if m != nil {
		return m.Evidences
	}
	return nil
}

type ObjectTraceIDPair struct {
	TraceID []byte    `protobuf:"bytes,1,opt,name=traceID,proto3" json:"traceID,omitempty"`
	Objects []*Object `protobuf:"bytes,2,rep,name=objects" json:"objects,omitempty"`
}

func (m *ObjectTraceIDPair) Reset()                    { *m = ObjectTraceIDPair{} }
func (m *ObjectTraceIDPair) String() string            { return proto.CompactTextString(m) }
func (*ObjectTraceIDPair) ProtoMessage()               {}
func (*ObjectTraceIDPair) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{5} }

func (m *ObjectTraceIDPair) GetTraceID() []byte {
	if m != nil {
		return m.TraceID
	}
	return nil
}

func (m *ObjectTraceIDPair) GetObjects() []*Object {
	if m != nil {
		return m.Objects
	}
	return nil
}

type ObjectList struct {
	List []*Object `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
}

func (m *ObjectList) Reset()                    { *m = ObjectList{} }
func (m *ObjectList) String() string            { return proto.CompactTextString(m) }
func (*ObjectList) ProtoMessage()               {}
func (*ObjectList) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{6} }

func (m *ObjectList) GetList() []*Object {
	if m != nil {
		return m.List
	}
	return nil
}

type AddTransactionResponse struct {
	Objects []*Object `protobuf:"bytes,1,rep,name=objects" json:"objects,omitempty"`
}

func (m *AddTransactionResponse) Reset()                    { *m = AddTransactionResponse{} }
func (m *AddTransactionResponse) String() string            { return proto.CompactTextString(m) }
func (*AddTransactionResponse) ProtoMessage()               {}
func (*AddTransactionResponse) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{7} }

func (m *AddTransactionResponse) GetObjects() []*Object {
	if m != nil {
		return m.Objects
	}
	return nil
}

type StateReport struct {
	HashID          uint32          `protobuf:"varint,1,opt,name=hashID,proto3" json:"hashID,omitempty"`
	CommitDecisions map[uint64]bool `protobuf:"bytes,2,rep,name=commitDecisions" json:"commitDecisions,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Phase1Decisions map[uint64]bool `protobuf:"bytes,3,rep,name=phase1Decisions" json:"phase1Decisions,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Phase2Decisions map[uint64]bool `protobuf:"bytes,4,rep,name=phase2Decisions" json:"phase2Decisions,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	State           string          `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	PendingEvents   int32           `protobuf:"varint,6,opt,name=pendingEvents,proto3" json:"pendingEvents,omitempty"`
}

func (m *StateReport) Reset()                    { *m = StateReport{} }
func (m *StateReport) String() string            { return proto.CompactTextString(m) }
func (*StateReport) ProtoMessage()               {}
func (*StateReport) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{8} }

func (m *StateReport) GetHashID() uint32 {
	if m != nil {
		return m.HashID
	}
	return 0
}

func (m *StateReport) GetCommitDecisions() map[uint64]bool {
	if m != nil {
		return m.CommitDecisions
	}
	return nil
}

func (m *StateReport) GetPhase1Decisions() map[uint64]bool {
	if m != nil {
		return m.Phase1Decisions
	}
	return nil
}

func (m *StateReport) GetPhase2Decisions() map[uint64]bool {
	if m != nil {
		return m.Phase2Decisions
	}
	return nil
}

func (m *StateReport) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *StateReport) GetPendingEvents() int32 {
	if m != nil {
		return m.PendingEvents
	}
	return 0
}

type StatesReportResponse struct {
	States        []*StateReport `protobuf:"bytes,1,rep,name=states" json:"states,omitempty"`
	EventsInQueue int32          `protobuf:"varint,2,opt,name=eventsInQueue,proto3" json:"eventsInQueue,omitempty"`
}

func (m *StatesReportResponse) Reset()                    { *m = StatesReportResponse{} }
func (m *StatesReportResponse) String() string            { return proto.CompactTextString(m) }
func (*StatesReportResponse) ProtoMessage()               {}
func (*StatesReportResponse) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{9} }

func (m *StatesReportResponse) GetStates() []*StateReport {
	if m != nil {
		return m.States
	}
	return nil
}

func (m *StatesReportResponse) GetEventsInQueue() int32 {
	if m != nil {
		return m.EventsInQueue
	}
	return 0
}

type QueryObjectRequest struct {
	VersionID []byte `protobuf:"bytes,1,opt,name=versionID,proto3" json:"versionID,omitempty"`
}

func (m *QueryObjectRequest) Reset()                    { *m = QueryObjectRequest{} }
func (m *QueryObjectRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryObjectRequest) ProtoMessage()               {}
func (*QueryObjectRequest) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{10} }

func (m *QueryObjectRequest) GetVersionID() []byte {
	if m != nil {
		return m.VersionID
	}
	return nil
}

type QueryObjectResponse struct {
	Object *Object `protobuf:"bytes,1,opt,name=object" json:"object,omitempty"`
	Error  string  `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *QueryObjectResponse) Reset()                    { *m = QueryObjectResponse{} }
func (m *QueryObjectResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryObjectResponse) ProtoMessage()               {}
func (*QueryObjectResponse) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{11} }

func (m *QueryObjectResponse) GetObject() *Object {
	if m != nil {
		return m.Object
	}
	return nil
}

func (m *QueryObjectResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type CreateObjectRequest struct {
	Object []byte `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
}

func (m *CreateObjectRequest) Reset()                    { *m = CreateObjectRequest{} }
func (m *CreateObjectRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateObjectRequest) ProtoMessage()               {}
func (*CreateObjectRequest) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{12} }

func (m *CreateObjectRequest) GetObject() []byte {
	if m != nil {
		return m.Object
	}
	return nil
}

type CreateObjectResponse struct {
	ID    []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *CreateObjectResponse) Reset()                    { *m = CreateObjectResponse{} }
func (m *CreateObjectResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateObjectResponse) ProtoMessage()               {}
func (*CreateObjectResponse) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{13} }

func (m *CreateObjectResponse) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *CreateObjectResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type CreateObjectsRequest struct {
	Objects [][]byte `protobuf:"bytes,1,rep,name=objects" json:"objects,omitempty"`
}

func (m *CreateObjectsRequest) Reset()                    { *m = CreateObjectsRequest{} }
func (m *CreateObjectsRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateObjectsRequest) ProtoMessage()               {}
func (*CreateObjectsRequest) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{14} }

func (m *CreateObjectsRequest) GetObjects() [][]byte {
	if m != nil {
		return m.Objects
	}
	return nil
}

type CreateObjectsResponse struct {
	IDs   [][]byte `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
	Error string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *CreateObjectsResponse) Reset()                    { *m = CreateObjectsResponse{} }
func (m *CreateObjectsResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateObjectsResponse) ProtoMessage()               {}
func (*CreateObjectsResponse) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{15} }

func (m *CreateObjectsResponse) GetIDs() [][]byte {
	if m != nil {
		return m.IDs
	}
	return nil
}

func (m *CreateObjectsResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type SBACMessage struct {
	Op        SBACOp            `protobuf:"varint,1,opt,name=op,proto3,enum=sbac.SBACOp" json:"op,omitempty"`
	Decision  SBACDecision      `protobuf:"varint,2,opt,name=decision,proto3,enum=sbac.SBACDecision" json:"decision,omitempty"`
	TxID      []byte            `protobuf:"bytes,3,opt,name=txId,proto3" json:"txId,omitempty"`
	Tx        *Transaction      `protobuf:"bytes,4,opt,name=tx" json:"tx,omitempty"`
	Evidences map[uint64][]byte `protobuf:"bytes,5,rep,name=evidences" json:"evidences,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Signature []byte            `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	PeerID    uint64            `protobuf:"varint,7,opt,name=peerId,proto3" json:"peerId,omitempty"`
}

func (m *SBACMessage) Reset()                    { *m = SBACMessage{} }
func (m *SBACMessage) String() string            { return proto.CompactTextString(m) }
func (*SBACMessage) ProtoMessage()               {}
func (*SBACMessage) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{16} }

func (m *SBACMessage) GetOp() SBACOp {
	if m != nil {
		return m.Op
	}
	return SBACOp_Phase1
}

func (m *SBACMessage) GetDecision() SBACDecision {
	if m != nil {
		return m.Decision
	}
	return SBACDecision_ACCEPT
}

func (m *SBACMessage) GetTxID() []byte {
	if m != nil {
		return m.TxID
	}
	return nil
}

func (m *SBACMessage) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *SBACMessage) GetEvidences() map[uint64][]byte {
	if m != nil {
		return m.Evidences
	}
	return nil
}

func (m *SBACMessage) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *SBACMessage) GetPeerID() uint64 {
	if m != nil {
		return m.PeerID
	}
	return 0
}

type SBACMessageAck struct {
	LastID uint64 `protobuf:"varint,1,opt,name=lastId,proto3" json:"lastId,omitempty"`
}

func (m *SBACMessageAck) Reset()                    { *m = SBACMessageAck{} }
func (m *SBACMessageAck) String() string            { return proto.CompactTextString(m) }
func (*SBACMessageAck) ProtoMessage()               {}
func (*SBACMessageAck) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{17} }

func (m *SBACMessageAck) GetLastID() uint64 {
	if m != nil {
		return m.LastID
	}
	return 0
}

type ConsensusTransaction struct {
	TxID      []byte            `protobuf:"bytes,1,opt,name=txId,proto3" json:"txId,omitempty"`
	Tx        *Transaction      `protobuf:"bytes,2,opt,name=tx" json:"tx,omitempty"`
	Evidences map[uint64][]byte `protobuf:"bytes,3,rep,name=evidences" json:"evidences,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Op        ConsensusOp       `protobuf:"varint,4,opt,name=op,proto3,enum=sbac.ConsensusOp" json:"op,omitempty"`
	Initiator uint64            `protobuf:"varint,5,opt,name=initiator,proto3" json:"initiator,omitempty"`
}

func (m *ConsensusTransaction) Reset()                    { *m = ConsensusTransaction{} }
func (m *ConsensusTransaction) String() string            { return proto.CompactTextString(m) }
func (*ConsensusTransaction) ProtoMessage()               {}
func (*ConsensusTransaction) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{18} }

func (m *ConsensusTransaction) GetTxID() []byte {
	if m != nil {
		return m.TxID
	}
	return nil
}

func (m *ConsensusTransaction) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *ConsensusTransaction) GetEvidences() map[uint64][]byte {
	if m != nil {
		return m.Evidences
	}
	return nil
}

func (m *ConsensusTransaction) GetOp() ConsensusOp {
	if m != nil {
		return m.Op
	}
	return ConsensusOp_Consensus1
}

func (m *ConsensusTransaction) GetInitiator() uint64 {
	if m != nil {
		return m.Initiator
	}
	return 0
}

func init() {
	proto.RegisterType((*Transaction)(nil), "sbac.Transaction")
	proto.RegisterType((*Strings)(nil), "sbac.strings")
	proto.RegisterType((*Trace)(nil), "sbac.Trace")
	proto.RegisterType((*Object)(nil), "sbac.Object")
	proto.RegisterType((*AddTransactionRequest)(nil), "sbac.AddTransactionRequest")
	proto.RegisterType((*ObjectTraceIDPair)(nil), "sbac.ObjectTraceIDPair")
	proto.RegisterType((*ObjectList)(nil), "sbac.ObjectList")
	proto.RegisterType((*AddTransactionResponse)(nil), "sbac.AddTransactionResponse")
	proto.RegisterType((*StateReport)(nil), "sbac.StateReport")
	proto.RegisterType((*StatesReportResponse)(nil), "sbac.StatesReportResponse")
	proto.RegisterType((*QueryObjectRequest)(nil), "sbac.QueryObjectRequest")
	proto.RegisterType((*QueryObjectResponse)(nil), "sbac.QueryObjectResponse")
	proto.RegisterType((*CreateObjectRequest)(nil), "sbac.CreateObjectRequest")
	proto.RegisterType((*CreateObjectResponse)(nil), "sbac.CreateObjectResponse")
	proto.RegisterType((*CreateObjectsRequest)(nil), "sbac.CreateObjectsRequest")
	proto.RegisterType((*CreateObjectsResponse)(nil), "sbac.CreateObjectsResponse")
	proto.RegisterType((*SBACMessage)(nil), "sbac.SBACMessage")
	proto.RegisterType((*SBACMessageAck)(nil), "sbac.SBACMessageAck")
	proto.RegisterType((*ConsensusTransaction)(nil), "sbac.ConsensusTransaction")
	proto.RegisterEnum("sbac.Opcode", Opcode_name, Opcode_value)
	proto.RegisterEnum("sbac.ObjectStatus", ObjectStatus_name, ObjectStatus_value)
	proto.RegisterEnum("sbac.SBACOp", SBACOp_name, SBACOp_value)
	proto.RegisterEnum("sbac.SBACDecision", SBACDecision_name, SBACDecision_value)
	proto.RegisterEnum("sbac.ConsensusOp", ConsensusOp_name, ConsensusOp_value)
}

func init() { proto.RegisterFile("sbac/types.proto", fileDescriptorTypes) }

var fileDescriptorTypes = []byte{
	// 1293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x57, 0xdd, 0x72, 0xdb, 0x44,
	0x14, 0xae, 0x65, 0x59, 0x8e, 0x8f, 0x9d, 0xc4, 0xd9, 0xfc, 0x8c, 0xda, 0x09, 0x38, 0x15, 0x6d,
	0x27, 0x35, 0xd4, 0x01, 0xd3, 0x0b, 0x86, 0x61, 0x18, 0xfc, 0x37, 0x54, 0xb4, 0xc4, 0xc9, 0xc6,
	0x2d, 0x70, 0xd5, 0x91, 0xa5, 0xad, 0x23, 0x9a, 0x48, 0x42, 0xbb, 0xce, 0x24, 0x2f, 0xc0, 0x2d,
	0x8f, 0xc0, 0x05, 0x6f, 0xc2, 0x73, 0x30, 0xbe, 0xf0, 0x93, 0x30, 0xfb, 0x23, 0x59, 0x4a, 0x9d,
	0x30, 0xb4, 0x77, 0x3e, 0xe7, 0x7c, 0xe7, 0x3b, 0xbb, 0xdf, 0x9e, 0x3d, 0x2b, 0x43, 0x9d, 0x8e,
	0x1d, 0xf7, 0x80, 0x5d, 0x45, 0x84, 0xb6, 0xa2, 0x38, 0x64, 0x21, 0xd2, 0xb9, 0xe7, 0xde, 0x93,
	0x89, 0xcf, 0x4e, 0xa7, 0xe3, 0x96, 0x1b, 0x9e, 0x1f, 0x4c, 0xc2, 0x49, 0x78, 0x20, 0x82, 0xe3,
	0xe9, 0x1b, 0x61, 0x09, 0x43, 0xfc, 0x92, 0x49, 0x56, 0x1b, 0xaa, 0xa3, 0xd8, 0x09, 0xa8, 0xe3,
	0x32, 0x3f, 0x0c, 0xd0, 0x27, 0x60, 0xb0, 0xd8, 0x71, 0x09, 0x35, 0x0b, 0x7b, 0xc5, 0xfd, 0x6a,
	0xbb, 0xda, 0xe2, 0xa4, 0xad, 0x11, 0xf7, 0x61, 0x15, 0xb2, 0x3e, 0x82, 0x32, 0x65, 0xb1, 0x1f,
	0x4c, 0x28, 0x42, 0xa0, 0x53, 0x16, 0x4b, 0x74, 0x05, 0x8b, 0xdf, 0xd6, 0xef, 0x3a, 0x94, 0x44,
	0x02, 0x6a, 0x01, 0xb8, 0x61, 0xc0, 0xb3, 0x98, 0xdd, 0x37, 0x0b, 0x7b, 0x85, 0xfd, 0x4a, 0x77,
	0x6d, 0x3e, 0x6b, 0x40, 0x2f, 0xf5, 0xe2, 0x0c, 0x02, 0xed, 0x42, 0x25, 0x8a, 0x43, 0x97, 0x78,
	0xd3, 0x98, 0x98, 0x1a, 0x87, 0xe3, 0x85, 0x03, 0x0d, 0x61, 0xdb, 0x0f, 0xa2, 0x29, 0x1b, 0x8e,
	0x7f, 0x25, 0x2e, 0x7b, 0x45, 0x62, 0xea, 0x87, 0x81, 0xdd, 0xa7, 0x66, 0x71, 0xaf, 0xb8, 0x5f,
	0xeb, 0xde, 0x9d, 0xcf, 0x1a, 0xdb, 0xf6, 0x32, 0x00, 0x5e, 0x9e, 0x87, 0x7e, 0x06, 0x53, 0x04,
	0x30, 0x79, 0x43, 0x62, 0x12, 0xb8, 0x24, 0xc3, 0xa9, 0x0b, 0xce, 0xdd, 0xf9, 0xac, 0x61, 0xda,
	0x37, 0x60, 0xf0, 0x8d, 0xd9, 0xe8, 0x01, 0xac, 0x86, 0x53, 0x96, 0xd6, 0xa4, 0x66, 0x89, 0xd3,
	0xe1, 0xbc, 0x13, 0x7d, 0x0c, 0x10, 0x39, 0xb1, 0x73, 0x4e, 0x18, 0x89, 0xa9, 0x69, 0x08, 0x48,
	0xc6, 0x83, 0x4c, 0x28, 0xc7, 0x84, 0x4d, 0xe3, 0x80, 0x9a, 0x65, 0x11, 0x4c, 0x4c, 0xf4, 0x10,
	0x8c, 0x33, 0x67, 0x4c, 0xce, 0xa8, 0xb9, 0x22, 0x8e, 0x69, 0x55, 0x1e, 0x93, 0x3a, 0x15, 0xac,
	0x82, 0xe8, 0x00, 0x6a, 0x1e, 0x89, 0x48, 0xe0, 0x91, 0xc0, 0xf5, 0x09, 0x35, 0x2b, 0xef, 0x9e,
	0x69, 0x0e, 0x80, 0x2c, 0xa8, 0x65, 0xa4, 0xa2, 0x26, 0x88, 0xb2, 0x39, 0x1f, 0xda, 0x87, 0xf5,
	0xfc, 0xbe, 0xa9, 0x59, 0x15, 0xb0, 0xeb, 0x6e, 0xeb, 0x8f, 0x02, 0x18, 0x32, 0x0b, 0x7d, 0x0a,
	0x95, 0x8b, 0x44, 0x1e, 0xd1, 0x08, 0xb5, 0xee, 0xea, 0x7c, 0xd6, 0xa8, 0xa4, 0x9a, 0xe1, 0x45,
	0x1c, 0x6d, 0x41, 0xe9, 0xc2, 0x39, 0x9b, 0xca, 0x16, 0xa8, 0x61, 0x69, 0xa0, 0x26, 0x18, 0x94,
	0x39, 0x6c, 0xca, 0xcf, 0xbb, 0xb0, 0xbf, 0xd6, 0x46, 0x72, 0x1b, 0xb2, 0xc0, 0x89, 0x88, 0x60,
	0x85, 0x40, 0x3b, 0xa9, 0x3e, 0xba, 0x68, 0x4c, 0x65, 0x59, 0x7f, 0x17, 0x60, 0xbb, 0xe3, 0x79,
	0x99, 0x8e, 0xc7, 0xe4, 0xb7, 0x29, 0xa1, 0x0c, 0xdd, 0x07, 0x8d, 0x5d, 0x8a, 0x95, 0x55, 0xdb,
	0x1b, 0xa9, 0x40, 0x29, 0x4a, 0x63, 0x97, 0xe8, 0x19, 0x54, 0xc8, 0x85, 0xef, 0xc9, 0x2d, 0x6b,
	0x42, 0xca, 0xa6, 0x44, 0x2e, 0xa5, 0x6c, 0x0d, 0x12, 0xf0, 0x20, 0x60, 0xf1, 0x15, 0x5e, 0x24,
	0xdf, 0xfb, 0x06, 0xd6, 0xf2, 0x41, 0x54, 0x87, 0xe2, 0x5b, 0x72, 0x25, 0xea, 0xeb, 0x98, 0xff,
	0x5c, 0x2e, 0xc2, 0xd7, 0xda, 0x57, 0x05, 0x6b, 0x0c, 0x1b, 0x72, 0xd3, 0xe2, 0x04, 0xed, 0xfe,
	0x91, 0xe3, 0xc7, 0xe8, 0x21, 0x94, 0x99, 0x34, 0x95, 0xbc, 0xd5, 0xf9, 0xac, 0x51, 0x56, 0x08,
	0x9c, 0xc4, 0xd0, 0x23, 0x28, 0x87, 0xea, 0x6c, 0xe5, 0x0e, 0x6a, 0x59, 0x15, 0x71, 0x12, 0xb4,
	0x5a, 0x00, 0xd2, 0xf5, 0xc2, 0xa7, 0x0c, 0xed, 0x81, 0x7e, 0xe6, 0x53, 0xa6, 0x66, 0x42, 0x3e,
	0x45, 0x44, 0xac, 0xef, 0x60, 0xe7, 0xba, 0x08, 0x34, 0x0a, 0x03, 0x4a, 0xb2, 0x15, 0x0b, 0xb7,
	0x55, 0xfc, 0x4b, 0x87, 0x2a, 0x3f, 0x45, 0x82, 0x49, 0x14, 0xc6, 0x8c, 0x1f, 0xe1, 0xa9, 0x43,
	0x4f, 0xd5, 0x7e, 0x56, 0xb1, 0xb2, 0xd0, 0x11, 0xac, 0xbb, 0xe1, 0xf9, 0xb9, 0xcf, 0xfa, 0xc4,
	0xf5, 0x79, 0xc3, 0x24, 0x3b, 0x79, 0x24, 0x79, 0x33, 0x1c, 0xad, 0x5e, 0x1e, 0x28, 0xcf, 0xe1,
	0x7a, 0x3a, 0x67, 0x8c, 0x4e, 0x1d, 0x4a, 0xbe, 0x58, 0x30, 0x16, 0x6f, 0x62, 0x3c, 0xca, 0x03,
	0x15, 0xe3, 0xb5, 0xf4, 0x94, 0xb1, 0xbd, 0x60, 0xd4, 0x6f, 0x65, 0x6c, 0x2f, 0x65, 0x5c, 0x78,
	0x79, 0x37, 0xf0, 0xd6, 0x26, 0x66, 0x49, 0x4c, 0x45, 0x69, 0xf0, 0x31, 0xc3, 0x2f, 0xaf, 0x1f,
	0x4c, 0x06, 0x17, 0x24, 0x60, 0x7c, 0x86, 0x14, 0xf6, 0x4b, 0x38, 0xef, 0xbc, 0xd7, 0x85, 0xad,
	0x65, 0x42, 0xfc, 0x57, 0xcf, 0xad, 0x64, 0x7a, 0x8e, 0x73, 0x2c, 0xdb, 0xfa, 0x7b, 0x71, 0xb4,
	0xdf, 0x9f, 0xc3, 0x9a, 0xc0, 0x96, 0x10, 0x8f, 0x4a, 0xf5, 0xd2, 0x2e, 0x7b, 0x2c, 0x87, 0x43,
	0xfa, 0x6e, 0x6d, 0xbc, 0x23, 0x34, 0x56, 0x00, 0x2e, 0x1a, 0x11, 0xc2, 0xd8, 0xc1, 0xf1, 0x94,
	0xa8, 0x22, 0x25, 0x9c, 0x77, 0x5a, 0x1d, 0x40, 0xc7, 0x53, 0x12, 0x5f, 0xa9, 0x36, 0x55, 0x53,
	0xe2, 0xff, 0x8c, 0x31, 0xeb, 0x18, 0x36, 0x73, 0x14, 0x6a, 0xa9, 0x0f, 0xc0, 0x90, 0x3d, 0xaf,
	0xa6, 0x4d, 0xfe, 0x3e, 0xa8, 0x18, 0x97, 0x80, 0xc4, 0x71, 0x18, 0xab, 0x67, 0x50, 0x1a, 0xd6,
	0x13, 0xd8, 0xec, 0xc5, 0xc4, 0x61, 0x24, 0xbf, 0xac, 0x9d, 0x1c, 0x65, 0x2d, 0x21, 0xb1, 0xfa,
	0xb0, 0x95, 0x87, 0xab, 0x25, 0xec, 0x80, 0xe6, 0x7b, 0x6a, 0xfd, 0xc6, 0x7c, 0xd6, 0xd0, 0xec,
	0x3e, 0xd6, 0x7c, 0xef, 0x86, 0xa2, 0x9f, 0xe7, 0x59, 0x68, 0x52, 0xd5, 0xcc, 0xdf, 0xec, 0xda,
	0xe2, 0x2e, 0x3f, 0x83, 0xed, 0x6b, 0x19, 0xaa, 0xf0, 0x5d, 0x28, 0xfa, 0x9e, 0x82, 0x77, 0xcb,
	0xf3, 0x59, 0xa3, 0xc8, 0xdf, 0x51, 0xee, 0xbb, 0xa1, 0xf6, 0x3f, 0x1a, 0x54, 0x4f, 0xba, 0x9d,
	0xde, 0x8f, 0x84, 0x52, 0x67, 0x42, 0xd0, 0x2e, 0x68, 0x61, 0x24, 0x56, 0xbe, 0x96, 0x08, 0xc7,
	0xc3, 0xc3, 0x08, 0x6b, 0x61, 0x84, 0x5a, 0xb0, 0xe2, 0xa9, 0xde, 0x12, 0x34, 0xe9, 0x23, 0xc1,
	0x31, 0x49, 0xd7, 0xe1, 0x14, 0x83, 0x76, 0x41, 0x67, 0x97, 0xb6, 0x27, 0x1e, 0x94, 0x5a, 0x77,
	0x65, 0x3e, 0x6b, 0xe8, 0xa3, 0x4b, 0xbb, 0x8f, 0x85, 0x57, 0x3d, 0x09, 0xfa, 0x6d, 0x4f, 0xc2,
	0xb7, 0xd9, 0x27, 0xa1, 0x24, 0x3a, 0x6f, 0x6f, 0x51, 0x51, 0x2d, 0xfa, 0xe6, 0x87, 0x80, 0x7f,
	0xf0, 0x50, 0x7f, 0x12, 0x38, 0x8c, 0x7f, 0xf0, 0x18, 0xe2, 0xec, 0x16, 0x0e, 0x64, 0x81, 0x11,
	0x11, 0x12, 0xdb, 0x9e, 0x59, 0xe6, 0x77, 0xa3, 0x0b, 0xf3, 0x59, 0xc3, 0x38, 0xe2, 0x9e, 0x3e,
	0x56, 0x91, 0x0f, 0x7c, 0x4a, 0x9e, 0xc2, 0x5a, 0x66, 0xa1, 0x1d, 0xf7, 0x2d, 0xaf, 0x79, 0xe6,
	0x50, 0x66, 0xcb, 0xf6, 0x50, 0x35, 0x5f, 0x70, 0x4f, 0x1f, 0xab, 0x88, 0xf5, 0xa7, 0xc6, 0x27,
	0x4a, 0x40, 0x49, 0x40, 0xa7, 0x34, 0xfb, 0xf5, 0x98, 0xe8, 0x59, 0xb8, 0x45, 0x4f, 0xed, 0x36,
	0x3d, 0xbf, 0xcf, 0xea, 0x29, 0x87, 0xf0, 0x63, 0x89, 0x5c, 0x56, 0xef, 0x16, 0x61, 0xef, 0x8b,
	0x3e, 0xd1, 0x45, 0x0f, 0x6c, 0x5c, 0x63, 0x50, 0xcd, 0xb2, 0x0b, 0x15, 0x3f, 0xf0, 0x99, 0xef,
	0xb0, 0x30, 0x16, 0x63, 0x55, 0xc7, 0x0b, 0xc7, 0x87, 0xe9, 0xda, 0xbc, 0x02, 0x63, 0x18, 0xb9,
	0xa1, 0x47, 0x50, 0x15, 0xca, 0x2f, 0x0f, 0x9f, 0x1f, 0x0e, 0x7f, 0x3a, 0xac, 0xdf, 0x41, 0x9b,
	0xb0, 0xde, 0xe9, 0xf7, 0x5f, 0x8f, 0x70, 0xe7, 0xf0, 0xa4, 0xd3, 0x1b, 0xd9, 0xc3, 0xc3, 0x7a,
	0x01, 0xd5, 0xa1, 0x76, 0xfc, 0x72, 0x80, 0x7f, 0x79, 0x3d, 0xec, 0xfe, 0x30, 0xe8, 0x8d, 0xea,
	0x1a, 0xda, 0x80, 0xd5, 0x1e, 0x1e, 0x74, 0x46, 0x83, 0xc4, 0x55, 0x44, 0x00, 0xc6, 0xc9, 0xa8,
	0x33, 0x1a, 0x9c, 0xd4, 0x75, 0xb4, 0x02, 0x3a, 0x3f, 0xb4, 0x7a, 0x09, 0x21, 0x58, 0xcb, 0x01,
	0x4f, 0xea, 0x46, 0xf3, 0x29, 0xd4, 0xb2, 0x9f, 0x44, 0x3c, 0x93, 0x97, 0x7a, 0x35, 0xa8, 0xdf,
	0x41, 0x35, 0x58, 0xb1, 0x0f, 0x95, 0x55, 0xe0, 0x91, 0x17, 0xc3, 0xde, 0xf3, 0x41, 0xbf, 0xae,
	0x35, 0x3f, 0x03, 0x43, 0xde, 0x23, 0xee, 0x95, 0x93, 0xbe, 0x7e, 0x27, 0xfd, 0xdd, 0x96, 0x68,
	0xf9, 0x8a, 0xd4, 0xb5, 0xe6, 0x23, 0xa8, 0x65, 0x6f, 0x94, 0xac, 0xd1, 0x1b, 0x1c, 0x8d, 0x64,
	0x0e, 0x1e, 0x88, 0x55, 0x17, 0x9a, 0x5d, 0xa8, 0x66, 0x54, 0x47, 0x6b, 0x00, 0xa9, 0xc9, 0xe9,
	0xb3, 0x36, 0x2f, 0xb1, 0x09, 0xeb, 0xa9, 0x9d, 0xd4, 0x1a, 0x1b, 0xe2, 0x7f, 0xca, 0x97, 0xff,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x44, 0x03, 0x34, 0xc4, 0xf0, 0x0c, 0x00, 0x00,
}
