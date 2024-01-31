package gen

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SplitflapState_ModuleState_State int32

const (
	// Keep in sync with splitflap_module_data.h!
	SplitflapState_ModuleState_NORMAL         SplitflapState_ModuleState_State = 0
	SplitflapState_ModuleState_LOOK_FOR_HOME  SplitflapState_ModuleState_State = 1
	SplitflapState_ModuleState_SENSOR_ERROR   SplitflapState_ModuleState_State = 2
	SplitflapState_ModuleState_PANIC          SplitflapState_ModuleState_State = 3
	SplitflapState_ModuleState_STATE_DISABLED SplitflapState_ModuleState_State = 4
)

// Enum value maps for SplitflapState_ModuleState_State.
var (
	SplitflapState_ModuleState_State_name = map[int32]string{
		0: "NORMAL",
		1: "LOOK_FOR_HOME",
		2: "SENSOR_ERROR",
		3: "PANIC",
		4: "STATE_DISABLED",
	}
	SplitflapState_ModuleState_State_value = map[string]int32{
		"NORMAL":         0,
		"LOOK_FOR_HOME":  1,
		"SENSOR_ERROR":   2,
		"PANIC":          3,
		"STATE_DISABLED": 4,
	}
)

func (x SplitflapState_ModuleState_State) Enum() *SplitflapState_ModuleState_State {
	p := new(SplitflapState_ModuleState_State)
	*p = x
	return p
}

func (x SplitflapState_ModuleState_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SplitflapState_ModuleState_State) Descriptor() protoreflect.EnumDescriptor {
	return file_test_proto_enumTypes[0].Descriptor()
}

func (SplitflapState_ModuleState_State) Type() protoreflect.EnumType {
	return &file_test_proto_enumTypes[0]
}

func (x SplitflapState_ModuleState_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SplitflapState_ModuleState_State.Descriptor instead.
func (SplitflapState_ModuleState_State) EnumDescriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0, 0, 0}
}

type SupervisorState_State int32

const (
	SupervisorState_UNKNOWN                  SupervisorState_State = 0
	SupervisorState_STARTING_VERIFY_PSU_OFF  SupervisorState_State = 1
	SupervisorState_STARTING_VERIFY_VOLTAGES SupervisorState_State = 2
	SupervisorState_STARTING_ENABLE_CHANNELS SupervisorState_State = 3
	SupervisorState_NORMAL                   SupervisorState_State = 4
	SupervisorState_FAULT                    SupervisorState_State = 5
)

// Enum value maps for SupervisorState_State.
var (
	SupervisorState_State_name = map[int32]string{
		0: "UNKNOWN",
		1: "STARTING_VERIFY_PSU_OFF",
		2: "STARTING_VERIFY_VOLTAGES",
		3: "STARTING_ENABLE_CHANNELS",
		4: "NORMAL",
		5: "FAULT",
	}
	SupervisorState_State_value = map[string]int32{
		"UNKNOWN":                  0,
		"STARTING_VERIFY_PSU_OFF":  1,
		"STARTING_VERIFY_VOLTAGES": 2,
		"STARTING_ENABLE_CHANNELS": 3,
		"NORMAL":                   4,
		"FAULT":                    5,
	}
)

func (x SupervisorState_State) Enum() *SupervisorState_State {
	p := new(SupervisorState_State)
	*p = x
	return p
}

func (x SupervisorState_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SupervisorState_State) Descriptor() protoreflect.EnumDescriptor {
	return file_test_proto_enumTypes[1].Descriptor()
}

func (SupervisorState_State) Type() protoreflect.EnumType {
	return &file_test_proto_enumTypes[1]
}

func (x SupervisorState_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SupervisorState_State.Descriptor instead.
func (SupervisorState_State) EnumDescriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{3, 0}
}

type SupervisorState_FaultInfo_FaultType int32

const (
	SupervisorState_FaultInfo_UNKNOWN                    SupervisorState_FaultInfo_FaultType = 0
	SupervisorState_FaultInfo_NONE                       SupervisorState_FaultInfo_FaultType = 1
	SupervisorState_FaultInfo_INRUSH_CURRENT_NOT_SETTLED SupervisorState_FaultInfo_FaultType = 2
	SupervisorState_FaultInfo_SPLITFLAP_SHUTDOWN         SupervisorState_FaultInfo_FaultType = 3
	SupervisorState_FaultInfo_OUT_OF_RANGE               SupervisorState_FaultInfo_FaultType = 4
	SupervisorState_FaultInfo_OVER_CURRENT               SupervisorState_FaultInfo_FaultType = 5
	SupervisorState_FaultInfo_UNEXPECTED_POWER           SupervisorState_FaultInfo_FaultType = 6
)

// Enum value maps for SupervisorState_FaultInfo_FaultType.
var (
	SupervisorState_FaultInfo_FaultType_name = map[int32]string{
		0: "UNKNOWN",
		1: "NONE",
		2: "INRUSH_CURRENT_NOT_SETTLED",
		3: "SPLITFLAP_SHUTDOWN",
		4: "OUT_OF_RANGE",
		5: "OVER_CURRENT",
		6: "UNEXPECTED_POWER",
	}
	SupervisorState_FaultInfo_FaultType_value = map[string]int32{
		"UNKNOWN":                    0,
		"NONE":                       1,
		"INRUSH_CURRENT_NOT_SETTLED": 2,
		"SPLITFLAP_SHUTDOWN":         3,
		"OUT_OF_RANGE":               4,
		"OVER_CURRENT":               5,
		"UNEXPECTED_POWER":           6,
	}
)

func (x SupervisorState_FaultInfo_FaultType) Enum() *SupervisorState_FaultInfo_FaultType {
	p := new(SupervisorState_FaultInfo_FaultType)
	*p = x
	return p
}

func (x SupervisorState_FaultInfo_FaultType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SupervisorState_FaultInfo_FaultType) Descriptor() protoreflect.EnumDescriptor {
	return file_test_proto_enumTypes[2].Descriptor()
}

func (SupervisorState_FaultInfo_FaultType) Type() protoreflect.EnumType {
	return &file_test_proto_enumTypes[2]
}

func (x SupervisorState_FaultInfo_FaultType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SupervisorState_FaultInfo_FaultType.Descriptor instead.
func (SupervisorState_FaultInfo_FaultType) EnumDescriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{3, 1, 0}
}

type SplitflapCommand_ModuleCommand_Action int32

const (
	SplitflapCommand_ModuleCommand_NO_OP          SplitflapCommand_ModuleCommand_Action = 0
	SplitflapCommand_ModuleCommand_GO_TO_FLAP     SplitflapCommand_ModuleCommand_Action = 1
	SplitflapCommand_ModuleCommand_RESET_AND_HOME SplitflapCommand_ModuleCommand_Action = 2
)

// Enum value maps for SplitflapCommand_ModuleCommand_Action.
var (
	SplitflapCommand_ModuleCommand_Action_name = map[int32]string{
		0: "NO_OP",
		1: "GO_TO_FLAP",
		2: "RESET_AND_HOME",
	}
	SplitflapCommand_ModuleCommand_Action_value = map[string]int32{
		"NO_OP":          0,
		"GO_TO_FLAP":     1,
		"RESET_AND_HOME": 2,
	}
)

func (x SplitflapCommand_ModuleCommand_Action) Enum() *SplitflapCommand_ModuleCommand_Action {
	p := new(SplitflapCommand_ModuleCommand_Action)
	*p = x
	return p
}

func (x SplitflapCommand_ModuleCommand_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SplitflapCommand_ModuleCommand_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_test_proto_enumTypes[3].Descriptor()
}

func (SplitflapCommand_ModuleCommand_Action) Type() protoreflect.EnumType {
	return &file_test_proto_enumTypes[3]
}

func (x SplitflapCommand_ModuleCommand_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SplitflapCommand_ModuleCommand_Action.Descriptor instead.
func (SplitflapCommand_ModuleCommand_Action) EnumDescriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{5, 0, 0}
}

type SplitflapState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Modules []*SplitflapState_ModuleState `protobuf:"bytes,1,rep,name=modules,proto3" json:"modules,omitempty"`
}

func (x *SplitflapState) Reset() {
	*x = SplitflapState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SplitflapState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SplitflapState) ProtoMessage() {}

func (x *SplitflapState) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SplitflapState.ProtoReflect.Descriptor instead.
func (*SplitflapState) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *SplitflapState) GetModules() []*SplitflapState_ModuleState {
	if x != nil {
		return x.Modules
	}
	return nil
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *Log) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce uint32 `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{2}
}

func (x *Ack) GetNonce() uint32 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

type SupervisorState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UptimeMillis  uint32                               `protobuf:"varint,1,opt,name=uptime_millis,json=uptimeMillis,proto3" json:"uptime_millis,omitempty"`
	State         SupervisorState_State                `protobuf:"varint,2,opt,name=state,proto3,enum=PB.SupervisorState_State" json:"state,omitempty"`
	PowerChannels []*SupervisorState_PowerChannelState `protobuf:"bytes,3,rep,name=power_channels,json=powerChannels,proto3" json:"power_channels,omitempty"`
	FaultInfo     *SupervisorState_FaultInfo           `protobuf:"bytes,4,opt,name=fault_info,json=faultInfo,proto3" json:"fault_info,omitempty"`
}

func (x *SupervisorState) Reset() {
	*x = SupervisorState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupervisorState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupervisorState) ProtoMessage() {}

func (x *SupervisorState) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupervisorState.ProtoReflect.Descriptor instead.
func (*SupervisorState) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{3}
}

func (x *SupervisorState) GetUptimeMillis() uint32 {
	if x != nil {
		return x.UptimeMillis
	}
	return 0
}

func (x *SupervisorState) GetState() SupervisorState_State {
	if x != nil {
		return x.State
	}
	return SupervisorState_UNKNOWN
}

func (x *SupervisorState) GetPowerChannels() []*SupervisorState_PowerChannelState {
	if x != nil {
		return x.PowerChannels
	}
	return nil
}

func (x *SupervisorState) GetFaultInfo() *SupervisorState_FaultInfo {
	if x != nil {
		return x.FaultInfo
	}
	return nil
}

type FromSplitflap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//	*FromSplitflap_SplitflapState
	//	*FromSplitflap_Log
	//	*FromSplitflap_Ack
	//	*FromSplitflap_SupervisorState
	Payload isFromSplitflap_Payload `protobuf_oneof:"payload"`
}

func (x *FromSplitflap) Reset() {
	*x = FromSplitflap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FromSplitflap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FromSplitflap) ProtoMessage() {}

func (x *FromSplitflap) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FromSplitflap.ProtoReflect.Descriptor instead.
func (*FromSplitflap) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{4}
}

func (m *FromSplitflap) GetPayload() isFromSplitflap_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *FromSplitflap) GetSplitflapState() *SplitflapState {
	if x, ok := x.GetPayload().(*FromSplitflap_SplitflapState); ok {
		return x.SplitflapState
	}
	return nil
}

func (x *FromSplitflap) GetLog() *Log {
	if x, ok := x.GetPayload().(*FromSplitflap_Log); ok {
		return x.Log
	}
	return nil
}

func (x *FromSplitflap) GetAck() *Ack {
	if x, ok := x.GetPayload().(*FromSplitflap_Ack); ok {
		return x.Ack
	}
	return nil
}

func (x *FromSplitflap) GetSupervisorState() *SupervisorState {
	if x, ok := x.GetPayload().(*FromSplitflap_SupervisorState); ok {
		return x.SupervisorState
	}
	return nil
}

type isFromSplitflap_Payload interface {
	isFromSplitflap_Payload()
}

type FromSplitflap_SplitflapState struct {
	SplitflapState *SplitflapState `protobuf:"bytes,1,opt,name=splitflap_state,json=splitflapState,proto3,oneof"`
}

type FromSplitflap_Log struct {
	Log *Log `protobuf:"bytes,2,opt,name=log,proto3,oneof"`
}

type FromSplitflap_Ack struct {
	Ack *Ack `protobuf:"bytes,3,opt,name=ack,proto3,oneof"`
}

type FromSplitflap_SupervisorState struct {
	SupervisorState *SupervisorState `protobuf:"bytes,4,opt,name=supervisor_state,json=supervisorState,proto3,oneof"`
}

func (*FromSplitflap_SplitflapState) isFromSplitflap_Payload() {}

func (*FromSplitflap_Log) isFromSplitflap_Payload() {}

func (*FromSplitflap_Ack) isFromSplitflap_Payload() {}

func (*FromSplitflap_SupervisorState) isFromSplitflap_Payload() {}

type SplitflapCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Modules []*SplitflapCommand_ModuleCommand `protobuf:"bytes,2,rep,name=modules,proto3" json:"modules,omitempty"`
}

func (x *SplitflapCommand) Reset() {
	*x = SplitflapCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SplitflapCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SplitflapCommand) ProtoMessage() {}

func (x *SplitflapCommand) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SplitflapCommand.ProtoReflect.Descriptor instead.
func (*SplitflapCommand) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{5}
}

func (x *SplitflapCommand) GetModules() []*SplitflapCommand_ModuleCommand {
	if x != nil {
		return x.Modules
	}
	return nil
}

type SplitflapConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Modules []*SplitflapConfig_ModuleConfig `protobuf:"bytes,1,rep,name=modules,proto3" json:"modules,omitempty"`
}

func (x *SplitflapConfig) Reset() {
	*x = SplitflapConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SplitflapConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SplitflapConfig) ProtoMessage() {}

func (x *SplitflapConfig) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SplitflapConfig.ProtoReflect.Descriptor instead.
func (*SplitflapConfig) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{6}
}

func (x *SplitflapConfig) GetModules() []*SplitflapConfig_ModuleConfig {
	if x != nil {
		return x.Modules
	}
	return nil
}

type RequestState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RequestState) Reset() {
	*x = RequestState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestState) ProtoMessage() {}

func (x *RequestState) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestState.ProtoReflect.Descriptor instead.
func (*RequestState) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{7}
}

type ToSplitflap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce uint32 `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// Types that are assignable to Payload:
	//	*ToSplitflap_SplitflapCommand
	//	*ToSplitflap_SplitflapConfig
	//	*ToSplitflap_RequestState
	Payload isToSplitflap_Payload `protobuf_oneof:"payload"`
}

func (x *ToSplitflap) Reset() {
	*x = ToSplitflap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToSplitflap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToSplitflap) ProtoMessage() {}

func (x *ToSplitflap) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToSplitflap.ProtoReflect.Descriptor instead.
func (*ToSplitflap) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{8}
}

func (x *ToSplitflap) GetNonce() uint32 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (m *ToSplitflap) GetPayload() isToSplitflap_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *ToSplitflap) GetSplitflapCommand() *SplitflapCommand {
	if x, ok := x.GetPayload().(*ToSplitflap_SplitflapCommand); ok {
		return x.SplitflapCommand
	}
	return nil
}

func (x *ToSplitflap) GetSplitflapConfig() *SplitflapConfig {
	if x, ok := x.GetPayload().(*ToSplitflap_SplitflapConfig); ok {
		return x.SplitflapConfig
	}
	return nil
}

func (x *ToSplitflap) GetRequestState() *RequestState {
	if x, ok := x.GetPayload().(*ToSplitflap_RequestState); ok {
		return x.RequestState
	}
	return nil
}

type isToSplitflap_Payload interface {
	isToSplitflap_Payload()
}

type ToSplitflap_SplitflapCommand struct {
	SplitflapCommand *SplitflapCommand `protobuf:"bytes,2,opt,name=splitflap_command,json=splitflapCommand,proto3,oneof"`
}

type ToSplitflap_SplitflapConfig struct {
	SplitflapConfig *SplitflapConfig `protobuf:"bytes,3,opt,name=splitflap_config,json=splitflapConfig,proto3,oneof"`
}

type ToSplitflap_RequestState struct {
	RequestState *RequestState `protobuf:"bytes,4,opt,name=request_state,json=requestState,proto3,oneof"`
}

func (*ToSplitflap_SplitflapCommand) isToSplitflap_Payload() {}

func (*ToSplitflap_SplitflapConfig) isToSplitflap_Payload() {}

func (*ToSplitflap_RequestState) isToSplitflap_Payload() {}

type SplitflapState_ModuleState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State               SplitflapState_ModuleState_State `protobuf:"varint,1,opt,name=state,proto3,enum=PB.SplitflapState_ModuleState_State" json:"state,omitempty"`
	FlapIndex           uint32                           `protobuf:"varint,2,opt,name=flap_index,json=flapIndex,proto3" json:"flap_index,omitempty"`
	Moving              bool                             `protobuf:"varint,3,opt,name=moving,proto3" json:"moving,omitempty"`
	HomeState           bool                             `protobuf:"varint,4,opt,name=home_state,json=homeState,proto3" json:"home_state,omitempty"`
	CountUnexpectedHome uint32                           `protobuf:"varint,5,opt,name=count_unexpected_home,json=countUnexpectedHome,proto3" json:"count_unexpected_home,omitempty"`
	CountMissedHome     uint32                           `protobuf:"varint,6,opt,name=count_missed_home,json=countMissedHome,proto3" json:"count_missed_home,omitempty"`
}

func (x *SplitflapState_ModuleState) Reset() {
	*x = SplitflapState_ModuleState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SplitflapState_ModuleState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SplitflapState_ModuleState) ProtoMessage() {}

func (x *SplitflapState_ModuleState) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SplitflapState_ModuleState.ProtoReflect.Descriptor instead.
func (*SplitflapState_ModuleState) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SplitflapState_ModuleState) GetState() SplitflapState_ModuleState_State {
	if x != nil {
		return x.State
	}
	return SplitflapState_ModuleState_NORMAL
}

func (x *SplitflapState_ModuleState) GetFlapIndex() uint32 {
	if x != nil {
		return x.FlapIndex
	}
	return 0
}

func (x *SplitflapState_ModuleState) GetMoving() bool {
	if x != nil {
		return x.Moving
	}
	return false
}

func (x *SplitflapState_ModuleState) GetHomeState() bool {
	if x != nil {
		return x.HomeState
	}
	return false
}

func (x *SplitflapState_ModuleState) GetCountUnexpectedHome() uint32 {
	if x != nil {
		return x.CountUnexpectedHome
	}
	return 0
}

func (x *SplitflapState_ModuleState) GetCountMissedHome() uint32 {
	if x != nil {
		return x.CountMissedHome
	}
	return 0
}

type SupervisorState_PowerChannelState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoltageVolts float32 `protobuf:"fixed32,1,opt,name=voltage_volts,json=voltageVolts,proto3" json:"voltage_volts,omitempty"`
	CurrentAmps  float32 `protobuf:"fixed32,2,opt,name=current_amps,json=currentAmps,proto3" json:"current_amps,omitempty"`
	On           bool    `protobuf:"varint,3,opt,name=on,proto3" json:"on,omitempty"`
}

func (x *SupervisorState_PowerChannelState) Reset() {
	*x = SupervisorState_PowerChannelState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupervisorState_PowerChannelState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupervisorState_PowerChannelState) ProtoMessage() {}

func (x *SupervisorState_PowerChannelState) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupervisorState_PowerChannelState.ProtoReflect.Descriptor instead.
func (*SupervisorState_PowerChannelState) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{3, 0}
}

func (x *SupervisorState_PowerChannelState) GetVoltageVolts() float32 {
	if x != nil {
		return x.VoltageVolts
	}
	return 0
}

func (x *SupervisorState_PowerChannelState) GetCurrentAmps() float32 {
	if x != nil {
		return x.CurrentAmps
	}
	return 0
}

func (x *SupervisorState_PowerChannelState) GetOn() bool {
	if x != nil {
		return x.On
	}
	return false
}

type SupervisorState_FaultInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     SupervisorState_FaultInfo_FaultType `protobuf:"varint,1,opt,name=type,proto3,enum=PB.SupervisorState_FaultInfo_FaultType" json:"type,omitempty"`
	Msg      string                              `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	TsMillis uint32                              `protobuf:"varint,3,opt,name=ts_millis,json=tsMillis,proto3" json:"ts_millis,omitempty"`
}

func (x *SupervisorState_FaultInfo) Reset() {
	*x = SupervisorState_FaultInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupervisorState_FaultInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupervisorState_FaultInfo) ProtoMessage() {}

func (x *SupervisorState_FaultInfo) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupervisorState_FaultInfo.ProtoReflect.Descriptor instead.
func (*SupervisorState_FaultInfo) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{3, 1}
}

func (x *SupervisorState_FaultInfo) GetType() SupervisorState_FaultInfo_FaultType {
	if x != nil {
		return x.Type
	}
	return SupervisorState_FaultInfo_UNKNOWN
}

func (x *SupervisorState_FaultInfo) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SupervisorState_FaultInfo) GetTsMillis() uint32 {
	if x != nil {
		return x.TsMillis
	}
	return 0
}

type SplitflapCommand_ModuleCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action SplitflapCommand_ModuleCommand_Action `protobuf:"varint,1,opt,name=action,proto3,enum=PB.SplitflapCommand_ModuleCommand_Action" json:"action,omitempty"`
	Param  uint32                                `protobuf:"varint,2,opt,name=param,proto3" json:"param,omitempty"`
}

func (x *SplitflapCommand_ModuleCommand) Reset() {
	*x = SplitflapCommand_ModuleCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SplitflapCommand_ModuleCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SplitflapCommand_ModuleCommand) ProtoMessage() {}

func (x *SplitflapCommand_ModuleCommand) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SplitflapCommand_ModuleCommand.ProtoReflect.Descriptor instead.
func (*SplitflapCommand_ModuleCommand) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{5, 0}
}

func (x *SplitflapCommand_ModuleCommand) GetAction() SplitflapCommand_ModuleCommand_Action {
	if x != nil {
		return x.Action
	}
	return SplitflapCommand_ModuleCommand_NO_OP
}

func (x *SplitflapCommand_ModuleCommand) GetParam() uint32 {
	if x != nil {
		return x.Param
	}
	return 0
}

type SplitflapConfig_ModuleConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetFlapIndex uint32 `protobuf:"varint,1,opt,name=target_flap_index,json=targetFlapIndex,proto3" json:"target_flap_index,omitempty"`
	//*
	// Value that triggers a movement upon change. If unused, only changes to target_flap_index
	// will trigger a movement. This can be used to trigger a full revolution back to the *same*
	// flap index.
	//
	// NOTE: Must be < 256
	MovementNonce uint32 `protobuf:"varint,2,opt,name=movement_nonce,json=movementNonce,proto3" json:"movement_nonce,omitempty"`
	//*
	// Value that triggers a reset (clear error counters, re-home) upon change. If unused,
	// module will only re-home upon recoverable errors, and error counters will continue
	// to increase until overflow.
	//
	// NOTE: Must be < 256
	ResetNonce uint32 `protobuf:"varint,3,opt,name=reset_nonce,json=resetNonce,proto3" json:"reset_nonce,omitempty"`
}

func (x *SplitflapConfig_ModuleConfig) Reset() {
	*x = SplitflapConfig_ModuleConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SplitflapConfig_ModuleConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SplitflapConfig_ModuleConfig) ProtoMessage() {}

func (x *SplitflapConfig_ModuleConfig) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SplitflapConfig_ModuleConfig.ProtoReflect.Descriptor instead.
func (*SplitflapConfig_ModuleConfig) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{6, 0}
}

func (x *SplitflapConfig_ModuleConfig) GetTargetFlapIndex() uint32 {
	if x != nil {
		return x.TargetFlapIndex
	}
	return 0
}

func (x *SplitflapConfig_ModuleConfig) GetMovementNonce() uint32 {
	if x != nil {
		return x.MovementNonce
	}
	return 0
}

func (x *SplitflapConfig_ModuleConfig) GetResetNonce() uint32 {
	if x != nil {
		return x.ResetNonce
	}
	return 0
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x50, 0x42,
	0x1a, 0x0c, 0x6e, 0x61, 0x6e, 0x6f, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2,
	0x03, 0x0a, 0x0e, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x66, 0x6c, 0x61, 0x70, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x40, 0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x50, 0x42, 0x2e, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x66, 0x6c, 0x61,
	0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x42, 0x06, 0x92, 0x3f, 0x03, 0x10, 0xff, 0x01, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x73, 0x1a, 0xed, 0x02, 0x0a, 0x0b, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x24, 0x2e, 0x50, 0x42, 0x2e, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x66, 0x6c, 0x61,
	0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x24, 0x0a, 0x0a, 0x66, 0x6c, 0x61, 0x70, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x42, 0x05, 0x92, 0x3f, 0x02, 0x38, 0x08, 0x52, 0x09, 0x66, 0x6c, 0x61, 0x70,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x76, 0x69, 0x6e, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6d, 0x6f, 0x76, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a,
	0x0a, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x68, 0x6f, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x15,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x5f, 0x68, 0x6f, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x05, 0x92, 0x3f, 0x02,
	0x38, 0x08, 0x52, 0x13, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x11, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x6d, 0x69, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x68, 0x6f, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x05, 0x92, 0x3f, 0x02, 0x38, 0x08, 0x52, 0x0f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4d, 0x69, 0x73, 0x73, 0x65, 0x64, 0x48, 0x6f, 0x6d, 0x65, 0x22, 0x57, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12,
	0x11, 0x0a, 0x0d, 0x4c, 0x4f, 0x4f, 0x4b, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x48, 0x4f, 0x4d, 0x45,
	0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45, 0x4e, 0x53, 0x4f, 0x52, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x41, 0x4e, 0x49, 0x43, 0x10, 0x03, 0x12,
	0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45,
	0x44, 0x10, 0x04, 0x22, 0x1f, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x18, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0x92, 0x3f, 0x03, 0x70, 0xff, 0x01, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x22, 0x1b, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x22, 0x87, 0x06, 0x0a, 0x0f, 0x53, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x75, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x12, 0x2f, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x50, 0x42, 0x2e, 0x53,
	0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x70,
	0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x50, 0x42, 0x2e, 0x53, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69,
	0x73, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x05, 0x92, 0x3f, 0x02, 0x10,
	0x05, 0x52, 0x0d, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73,
	0x12, 0x3c, 0x0a, 0x0a, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x50, 0x42, 0x2e, 0x53, 0x75, 0x70, 0x65, 0x72, 0x76,
	0x69, 0x73, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x09, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x6b,
	0x0a, 0x11, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x76,
	0x6f, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x76, 0x6f, 0x6c, 0x74,
	0x61, 0x67, 0x65, 0x56, 0x6f, 0x6c, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x70, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6e, 0x1a, 0x96, 0x02, 0x0a, 0x09,
	0x46, 0x61, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3b, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x50, 0x42, 0x2e, 0x53, 0x75, 0x70,
	0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x46, 0x61, 0x75,
	0x6c, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x06, 0x92, 0x3f, 0x03, 0x70, 0xff, 0x01, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x73, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x74, 0x73, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x22, 0x94, 0x01,
	0x0a, 0x09, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x49, 0x4e, 0x52, 0x55, 0x53, 0x48, 0x5f, 0x43, 0x55, 0x52,
	0x52, 0x45, 0x4e, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x50, 0x4c, 0x49, 0x54, 0x46, 0x4c, 0x41, 0x50, 0x5f,
	0x53, 0x48, 0x55, 0x54, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x55,
	0x54, 0x5f, 0x4f, 0x46, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c,
	0x4f, 0x56, 0x45, 0x52, 0x5f, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x54, 0x10, 0x05, 0x12, 0x14,
	0x0a, 0x10, 0x55, 0x4e, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x50, 0x4f, 0x57,
	0x45, 0x52, 0x10, 0x06, 0x22, 0x84, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x53,
	0x54, 0x41, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x5f, 0x50,
	0x53, 0x55, 0x5f, 0x4f, 0x46, 0x46, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x54, 0x41, 0x52,
	0x54, 0x49, 0x4e, 0x47, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x5f, 0x56, 0x4f, 0x4c, 0x54,
	0x41, 0x47, 0x45, 0x53, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x54, 0x41, 0x52, 0x54, 0x49,
	0x4e, 0x47, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45,
	0x4c, 0x53, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x04,
	0x12, 0x09, 0x0a, 0x05, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x05, 0x22, 0xd5, 0x01, 0x0a, 0x0d,
	0x46, 0x72, 0x6f, 0x6d, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x66, 0x6c, 0x61, 0x70, 0x12, 0x3d, 0x0a,
	0x0f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x66, 0x6c, 0x61, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x50, 0x42, 0x2e, 0x53, 0x70, 0x6c, 0x69,
	0x74, 0x66, 0x6c, 0x61, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x73, 0x70,
	0x6c, 0x69, 0x74, 0x66, 0x6c, 0x61, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x03,
	0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x42, 0x2e, 0x4c,
	0x6f, 0x67, 0x48, 0x00, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x1b, 0x0a, 0x03, 0x61, 0x63, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x42, 0x2e, 0x41, 0x63, 0x6b, 0x48,
	0x00, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x12, 0x40, 0x0a, 0x10, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76,
	0x69, 0x73, 0x6f, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x50, 0x42, 0x2e, 0x53, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69,
	0x73, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x22, 0x83, 0x02, 0x0a, 0x10, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x66, 0x6c, 0x61,
	0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x44, 0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x50, 0x42, 0x2e, 0x53,
	0x70, 0x6c, 0x69, 0x74, 0x66, 0x6c, 0x61, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x42, 0x06, 0x92,
	0x3f, 0x03, 0x10, 0xff, 0x01, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x1a, 0xa8,
	0x01, 0x0a, 0x0d, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x41, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x29, 0x2e, 0x50, 0x42, 0x2e, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x66, 0x6c, 0x61, 0x70, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x05, 0x92, 0x3f, 0x02, 0x38, 0x08, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x22, 0x37, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x4f,
	0x5f, 0x4f, 0x50, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x47, 0x4f, 0x5f, 0x54, 0x4f, 0x5f, 0x46,
	0x4c, 0x41, 0x50, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x41,
	0x4e, 0x44, 0x5f, 0x48, 0x4f, 0x4d, 0x45, 0x10, 0x02, 0x22, 0xef, 0x01, 0x0a, 0x0f, 0x53, 0x70,
	0x6c, 0x69, 0x74, 0x66, 0x6c, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x42, 0x0a,
	0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x50, 0x42, 0x2e, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x66, 0x6c, 0x61, 0x70, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x42, 0x06, 0x92, 0x3f, 0x03, 0x10, 0xff, 0x01, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x73, 0x1a, 0x97, 0x01, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x31, 0x0a, 0x11, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x66, 0x6c, 0x61,
	0x70, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x05, 0x92,
	0x3f, 0x02, 0x38, 0x08, 0x52, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x46, 0x6c, 0x61, 0x70,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2c, 0x0a, 0x0e, 0x6d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x05, 0x92,
	0x3f, 0x02, 0x38, 0x08, 0x52, 0x0d, 0x6d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x6f,
	0x6e, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x05, 0x92, 0x3f, 0x02, 0x38, 0x08, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x65, 0x74, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0xee, 0x01, 0x0a, 0x0b,
	0x54, 0x6f, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x66, 0x6c, 0x61, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x12, 0x43, 0x0a, 0x11, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x66, 0x6c, 0x61, 0x70, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x50,
	0x42, 0x2e, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x66, 0x6c, 0x61, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x48, 0x00, 0x52, 0x10, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x66, 0x6c, 0x61, 0x70, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x40, 0x0a, 0x10, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x66,
	0x6c, 0x61, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x50, 0x42, 0x2e, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x66, 0x6c, 0x61, 0x70, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x66, 0x6c,
	0x61, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x37, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x50, 0x42, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x08, 0x5a, 0x06,
	0x2e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_test_proto_goTypes = []interface{}{
	(SplitflapState_ModuleState_State)(0),      // 0: PB.SplitflapState.ModuleState.State
	(SupervisorState_State)(0),                 // 1: PB.SupervisorState.State
	(SupervisorState_FaultInfo_FaultType)(0),   // 2: PB.SupervisorState.FaultInfo.FaultType
	(SplitflapCommand_ModuleCommand_Action)(0), // 3: PB.SplitflapCommand.ModuleCommand.Action
	(*SplitflapState)(nil),                     // 4: PB.SplitflapState
	(*Log)(nil),                                // 5: PB.Log
	(*Ack)(nil),                                // 6: PB.Ack
	(*SupervisorState)(nil),                    // 7: PB.SupervisorState
	(*FromSplitflap)(nil),                      // 8: PB.FromSplitflap
	(*SplitflapCommand)(nil),                   // 9: PB.SplitflapCommand
	(*SplitflapConfig)(nil),                    // 10: PB.SplitflapConfig
	(*RequestState)(nil),                       // 11: PB.RequestState
	(*ToSplitflap)(nil),                        // 12: PB.ToSplitflap
	(*SplitflapState_ModuleState)(nil),         // 13: PB.SplitflapState.ModuleState
	(*SupervisorState_PowerChannelState)(nil),  // 14: PB.SupervisorState.PowerChannelState
	(*SupervisorState_FaultInfo)(nil),          // 15: PB.SupervisorState.FaultInfo
	(*SplitflapCommand_ModuleCommand)(nil),     // 16: PB.SplitflapCommand.ModuleCommand
	(*SplitflapConfig_ModuleConfig)(nil),       // 17: PB.SplitflapConfig.ModuleConfig
}
var file_test_proto_depIdxs = []int32{
	13, // 0: PB.SplitflapState.modules:type_name -> PB.SplitflapState.ModuleState
	1,  // 1: PB.SupervisorState.state:type_name -> PB.SupervisorState.State
	14, // 2: PB.SupervisorState.power_channels:type_name -> PB.SupervisorState.PowerChannelState
	15, // 3: PB.SupervisorState.fault_info:type_name -> PB.SupervisorState.FaultInfo
	4,  // 4: PB.FromSplitflap._state:type_name -> PB.SplitflapState
	5,  // 5: PB.FromSplitflap.log:type_name -> PB.Log
	6,  // 6: PB.FromSplitflap.ack:type_name -> PB.Ack
	7,  // 7: PB.FromSplitflap.supervisor_state:type_name -> PB.SupervisorState
	16, // 8: PB.SplitflapCommand.modules:type_name -> PB.SplitflapCommand.ModuleCommand
	17, // 9: PB.SplitflapConfig.modules:type_name -> PB.SplitflapConfig.ModuleConfig
	9,  // 10: PB.ToSplitflap._command:type_name -> PB.SplitflapCommand
	10, // 11: PB.ToSplitflap._config:type_name -> PB.SplitflapConfig
	11, // 12: PB.ToSplitflap.request_state:type_name -> PB.RequestState
	0,  // 13: PB.SplitflapState.ModuleState.state:type_name -> PB.SplitflapState.ModuleState.State
	2,  // 14: PB.SupervisorState.FaultInfo.type:type_name -> PB.SupervisorState.FaultInfo.FaultType
	3,  // 15: PB.SplitflapCommand.ModuleCommand.action:type_name -> PB.SplitflapCommand.ModuleCommand.Action
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	// file_nanopb_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SplitflapState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupervisorState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FromSplitflap); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SplitflapCommand); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			// switch v := v.(*SplitflapConfig); i {
			// case 0:
			// 	return &v.state
			// case 1:
			// 	return &v.sizeCache
			// case 2:
			// 	return &v.unknownFields
			// default:
			return nil
			// }
		}
		file_test_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToSplitflap); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SplitflapState_ModuleState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupervisorState_PowerChannelState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupervisorState_FaultInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SplitflapCommand_ModuleCommand); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SplitflapConfig_ModuleConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_test_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*FromSplitflap_SplitflapState)(nil),
		(*FromSplitflap_Log)(nil),
		(*FromSplitflap_Ack)(nil),
		(*FromSplitflap_SupervisorState)(nil),
	}
	file_test_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*ToSplitflap_SplitflapCommand)(nil),
		(*ToSplitflap_SplitflapConfig)(nil),
		(*ToSplitflap_RequestState)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		EnumInfos:         file_test_proto_enumTypes,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}
