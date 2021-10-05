// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: shmsg.proto

package shmsg

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type G1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	G1Bytes []byte `protobuf:"bytes,1,opt,name=g1bytes,proto3" json:"g1bytes,omitempty"` // Unmarshal with new(G1).Unmarshal(msg.g1bytes)
}

func (x *G1) Reset() {
	*x = G1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shmsg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *G1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*G1) ProtoMessage() {}

func (x *G1) ProtoReflect() protoreflect.Message {
	mi := &file_shmsg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use G1.ProtoReflect.Descriptor instead.
func (*G1) Descriptor() ([]byte, []int) {
	return file_shmsg_proto_rawDescGZIP(), []int{0}
}

func (x *G1) GetG1Bytes() []byte {
	if x != nil {
		return x.G1Bytes
	}
	return nil
}

type G2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	G2Bytes []byte `protobuf:"bytes,1,opt,name=g2bytes,proto3" json:"g2bytes,omitempty"` // Unmarshal with new(G2).Unmarshal(msg.g2bytes)
}

func (x *G2) Reset() {
	*x = G2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shmsg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *G2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*G2) ProtoMessage() {}

func (x *G2) ProtoReflect() protoreflect.Message {
	mi := &file_shmsg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use G2.ProtoReflect.Descriptor instead.
func (*G2) Descriptor() ([]byte, []int) {
	return file_shmsg_proto_rawDescGZIP(), []int{1}
}

func (x *G2) GetG2Bytes() []byte {
	if x != nil {
		return x.G2Bytes
	}
	return nil
}

type GT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gtbytes []byte `protobuf:"bytes,1,opt,name=gtbytes,proto3" json:"gtbytes,omitempty"` // Unmarshal with new(GT).Unmarshal(msg.g2bytes)
}

func (x *GT) Reset() {
	*x = GT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shmsg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GT) ProtoMessage() {}

func (x *GT) ProtoReflect() protoreflect.Message {
	mi := &file_shmsg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GT.ProtoReflect.Descriptor instead.
func (*GT) Descriptor() ([]byte, []int) {
	return file_shmsg_proto_rawDescGZIP(), []int{2}
}

func (x *GT) GetGtbytes() []byte {
	if x != nil {
		return x.Gtbytes
	}
	return nil
}

type BatchConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartBatchIndex   uint64   `protobuf:"varint,1,opt,name=start_batch_index,json=startBatchIndex,proto3" json:"start_batch_index,omitempty"`
	Keypers           [][]byte `protobuf:"bytes,2,rep,name=keypers,proto3" json:"keypers,omitempty"`
	Threshold         uint64   `protobuf:"varint,3,opt,name=threshold,proto3" json:"threshold,omitempty"`
	ConfigIndex       uint64   `protobuf:"varint,5,opt,name=config_index,json=configIndex,proto3" json:"config_index,omitempty"`
	Started           bool     `protobuf:"varint,6,opt,name=started,proto3" json:"started,omitempty"`
	ValidatorsUpdated bool     `protobuf:"varint,7,opt,name=validatorsUpdated,proto3" json:"validatorsUpdated,omitempty"`
}

func (x *BatchConfig) Reset() {
	*x = BatchConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shmsg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchConfig) ProtoMessage() {}

func (x *BatchConfig) ProtoReflect() protoreflect.Message {
	mi := &file_shmsg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchConfig.ProtoReflect.Descriptor instead.
func (*BatchConfig) Descriptor() ([]byte, []int) {
	return file_shmsg_proto_rawDescGZIP(), []int{3}
}

func (x *BatchConfig) GetStartBatchIndex() uint64 {
	if x != nil {
		return x.StartBatchIndex
	}
	return 0
}

func (x *BatchConfig) GetKeypers() [][]byte {
	if x != nil {
		return x.Keypers
	}
	return nil
}

func (x *BatchConfig) GetThreshold() uint64 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

func (x *BatchConfig) GetConfigIndex() uint64 {
	if x != nil {
		return x.ConfigIndex
	}
	return 0
}

func (x *BatchConfig) GetStarted() bool {
	if x != nil {
		return x.Started
	}
	return false
}

func (x *BatchConfig) GetValidatorsUpdated() bool {
	if x != nil {
		return x.ValidatorsUpdated
	}
	return false
}

type BatchConfigStarted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchConfigIndex uint64 `protobuf:"varint,1,opt,name=batch_config_index,json=batchConfigIndex,proto3" json:"batch_config_index,omitempty"`
}

func (x *BatchConfigStarted) Reset() {
	*x = BatchConfigStarted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shmsg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchConfigStarted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchConfigStarted) ProtoMessage() {}

func (x *BatchConfigStarted) ProtoReflect() protoreflect.Message {
	mi := &file_shmsg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchConfigStarted.ProtoReflect.Descriptor instead.
func (*BatchConfigStarted) Descriptor() ([]byte, []int) {
	return file_shmsg_proto_rawDescGZIP(), []int{4}
}

func (x *BatchConfigStarted) GetBatchConfigIndex() uint64 {
	if x != nil {
		return x.BatchConfigIndex
	}
	return 0
}

type CheckIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValidatorPublicKey  []byte `protobuf:"bytes,1,opt,name=validator_public_key,json=validatorPublicKey,proto3" json:"validator_public_key,omitempty"`    // 32 byte ed25519 public key
	EncryptionPublicKey []byte `protobuf:"bytes,2,opt,name=encryption_public_key,json=encryptionPublicKey,proto3" json:"encryption_public_key,omitempty"` // compressed ecies public key
}

func (x *CheckIn) Reset() {
	*x = CheckIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shmsg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIn) ProtoMessage() {}

func (x *CheckIn) ProtoReflect() protoreflect.Message {
	mi := &file_shmsg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIn.ProtoReflect.Descriptor instead.
func (*CheckIn) Descriptor() ([]byte, []int) {
	return file_shmsg_proto_rawDescGZIP(), []int{5}
}

func (x *CheckIn) GetValidatorPublicKey() []byte {
	if x != nil {
		return x.ValidatorPublicKey
	}
	return nil
}

func (x *CheckIn) GetEncryptionPublicKey() []byte {
	if x != nil {
		return x.EncryptionPublicKey
	}
	return nil
}

type PolyEval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eon            uint64   `protobuf:"varint,1,opt,name=eon,proto3" json:"eon,omitempty"`
	Receivers      [][]byte `protobuf:"bytes,2,rep,name=receivers,proto3" json:"receivers,omitempty"`
	EncryptedEvals [][]byte `protobuf:"bytes,3,rep,name=encrypted_evals,json=encryptedEvals,proto3" json:"encrypted_evals,omitempty"`
}

func (x *PolyEval) Reset() {
	*x = PolyEval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shmsg_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolyEval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolyEval) ProtoMessage() {}

func (x *PolyEval) ProtoReflect() protoreflect.Message {
	mi := &file_shmsg_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolyEval.ProtoReflect.Descriptor instead.
func (*PolyEval) Descriptor() ([]byte, []int) {
	return file_shmsg_proto_rawDescGZIP(), []int{6}
}

func (x *PolyEval) GetEon() uint64 {
	if x != nil {
		return x.Eon
	}
	return 0
}

func (x *PolyEval) GetReceivers() [][]byte {
	if x != nil {
		return x.Receivers
	}
	return nil
}

func (x *PolyEval) GetEncryptedEvals() [][]byte {
	if x != nil {
		return x.EncryptedEvals
	}
	return nil
}

type PolyCommitment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eon    uint64   `protobuf:"varint,1,opt,name=eon,proto3" json:"eon,omitempty"`
	Gammas [][]byte `protobuf:"bytes,2,rep,name=gammas,proto3" json:"gammas,omitempty"`
}

func (x *PolyCommitment) Reset() {
	*x = PolyCommitment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shmsg_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolyCommitment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolyCommitment) ProtoMessage() {}

func (x *PolyCommitment) ProtoReflect() protoreflect.Message {
	mi := &file_shmsg_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolyCommitment.ProtoReflect.Descriptor instead.
func (*PolyCommitment) Descriptor() ([]byte, []int) {
	return file_shmsg_proto_rawDescGZIP(), []int{7}
}

func (x *PolyCommitment) GetEon() uint64 {
	if x != nil {
		return x.Eon
	}
	return 0
}

func (x *PolyCommitment) GetGammas() [][]byte {
	if x != nil {
		return x.Gammas
	}
	return nil
}

type Accusation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eon     uint64   `protobuf:"varint,1,opt,name=eon,proto3" json:"eon,omitempty"`
	Accused [][]byte `protobuf:"bytes,2,rep,name=accused,proto3" json:"accused,omitempty"`
}

func (x *Accusation) Reset() {
	*x = Accusation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shmsg_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Accusation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Accusation) ProtoMessage() {}

func (x *Accusation) ProtoReflect() protoreflect.Message {
	mi := &file_shmsg_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Accusation.ProtoReflect.Descriptor instead.
func (*Accusation) Descriptor() ([]byte, []int) {
	return file_shmsg_proto_rawDescGZIP(), []int{8}
}

func (x *Accusation) GetEon() uint64 {
	if x != nil {
		return x.Eon
	}
	return 0
}

func (x *Accusation) GetAccused() [][]byte {
	if x != nil {
		return x.Accused
	}
	return nil
}

type Apology struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eon       uint64   `protobuf:"varint,1,opt,name=eon,proto3" json:"eon,omitempty"`
	Accusers  [][]byte `protobuf:"bytes,2,rep,name=accusers,proto3" json:"accusers,omitempty"`
	PolyEvals [][]byte `protobuf:"bytes,3,rep,name=poly_evals,json=polyEvals,proto3" json:"poly_evals,omitempty"`
}

func (x *Apology) Reset() {
	*x = Apology{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shmsg_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Apology) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Apology) ProtoMessage() {}

func (x *Apology) ProtoReflect() protoreflect.Message {
	mi := &file_shmsg_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Apology.ProtoReflect.Descriptor instead.
func (*Apology) Descriptor() ([]byte, []int) {
	return file_shmsg_proto_rawDescGZIP(), []int{9}
}

func (x *Apology) GetEon() uint64 {
	if x != nil {
		return x.Eon
	}
	return 0
}

func (x *Apology) GetAccusers() [][]byte {
	if x != nil {
		return x.Accusers
	}
	return nil
}

func (x *Apology) GetPolyEvals() [][]byte {
	if x != nil {
		return x.PolyEvals
	}
	return nil
}

type EonStartVote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartBatchIndex uint64 `protobuf:"varint,1,opt,name=start_batch_index,json=startBatchIndex,proto3" json:"start_batch_index,omitempty"`
}

func (x *EonStartVote) Reset() {
	*x = EonStartVote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shmsg_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EonStartVote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EonStartVote) ProtoMessage() {}

func (x *EonStartVote) ProtoReflect() protoreflect.Message {
	mi := &file_shmsg_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EonStartVote.ProtoReflect.Descriptor instead.
func (*EonStartVote) Descriptor() ([]byte, []int) {
	return file_shmsg_proto_rawDescGZIP(), []int{10}
}

func (x *EonStartVote) GetStartBatchIndex() uint64 {
	if x != nil {
		return x.StartBatchIndex
	}
	return 0
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//	*Message_BatchConfig
	//	*Message_BatchConfigStarted
	//	*Message_CheckIn
	//	*Message_PolyEval
	//	*Message_PolyCommitment
	//	*Message_Accusation
	//	*Message_Apology
	//	*Message_EonStartVote
	Payload isMessage_Payload `protobuf_oneof:"payload"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shmsg_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_shmsg_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_shmsg_proto_rawDescGZIP(), []int{11}
}

func (m *Message) GetPayload() isMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Message) GetBatchConfig() *BatchConfig {
	if x, ok := x.GetPayload().(*Message_BatchConfig); ok {
		return x.BatchConfig
	}
	return nil
}

func (x *Message) GetBatchConfigStarted() *BatchConfigStarted {
	if x, ok := x.GetPayload().(*Message_BatchConfigStarted); ok {
		return x.BatchConfigStarted
	}
	return nil
}

func (x *Message) GetCheckIn() *CheckIn {
	if x, ok := x.GetPayload().(*Message_CheckIn); ok {
		return x.CheckIn
	}
	return nil
}

func (x *Message) GetPolyEval() *PolyEval {
	if x, ok := x.GetPayload().(*Message_PolyEval); ok {
		return x.PolyEval
	}
	return nil
}

func (x *Message) GetPolyCommitment() *PolyCommitment {
	if x, ok := x.GetPayload().(*Message_PolyCommitment); ok {
		return x.PolyCommitment
	}
	return nil
}

func (x *Message) GetAccusation() *Accusation {
	if x, ok := x.GetPayload().(*Message_Accusation); ok {
		return x.Accusation
	}
	return nil
}

func (x *Message) GetApology() *Apology {
	if x, ok := x.GetPayload().(*Message_Apology); ok {
		return x.Apology
	}
	return nil
}

func (x *Message) GetEonStartVote() *EonStartVote {
	if x, ok := x.GetPayload().(*Message_EonStartVote); ok {
		return x.EonStartVote
	}
	return nil
}

type isMessage_Payload interface {
	isMessage_Payload()
}

type Message_BatchConfig struct {
	BatchConfig *BatchConfig `protobuf:"bytes,4,opt,name=batch_config,json=batchConfig,proto3,oneof"`
}

type Message_BatchConfigStarted struct {
	BatchConfigStarted *BatchConfigStarted `protobuf:"bytes,6,opt,name=batch_config_started,json=batchConfigStarted,proto3,oneof"`
}

type Message_CheckIn struct {
	CheckIn *CheckIn `protobuf:"bytes,7,opt,name=check_in,json=checkIn,proto3,oneof"`
}

type Message_PolyEval struct {
	// DKG messages
	PolyEval *PolyEval `protobuf:"bytes,9,opt,name=poly_eval,json=polyEval,proto3,oneof"`
}

type Message_PolyCommitment struct {
	PolyCommitment *PolyCommitment `protobuf:"bytes,10,opt,name=poly_commitment,json=polyCommitment,proto3,oneof"`
}

type Message_Accusation struct {
	Accusation *Accusation `protobuf:"bytes,11,opt,name=accusation,proto3,oneof"`
}

type Message_Apology struct {
	Apology *Apology `protobuf:"bytes,12,opt,name=apology,proto3,oneof"`
}

type Message_EonStartVote struct {
	EonStartVote *EonStartVote `protobuf:"bytes,13,opt,name=eon_start_vote,json=eonStartVote,proto3,oneof"`
}

func (*Message_BatchConfig) isMessage_Payload() {}

func (*Message_BatchConfigStarted) isMessage_Payload() {}

func (*Message_CheckIn) isMessage_Payload() {}

func (*Message_PolyEval) isMessage_Payload() {}

func (*Message_PolyCommitment) isMessage_Payload() {}

func (*Message_Accusation) isMessage_Payload() {}

func (*Message_Apology) isMessage_Payload() {}

func (*Message_EonStartVote) isMessage_Payload() {}

type MessageWithNonce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg         *Message `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	ChainId     []byte   `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	RandomNonce uint64   `protobuf:"varint,3,opt,name=random_nonce,json=randomNonce,proto3" json:"random_nonce,omitempty"`
}

func (x *MessageWithNonce) Reset() {
	*x = MessageWithNonce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shmsg_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithNonce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithNonce) ProtoMessage() {}

func (x *MessageWithNonce) ProtoReflect() protoreflect.Message {
	mi := &file_shmsg_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithNonce.ProtoReflect.Descriptor instead.
func (*MessageWithNonce) Descriptor() ([]byte, []int) {
	return file_shmsg_proto_rawDescGZIP(), []int{12}
}

func (x *MessageWithNonce) GetMsg() *Message {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *MessageWithNonce) GetChainId() []byte {
	if x != nil {
		return x.ChainId
	}
	return nil
}

func (x *MessageWithNonce) GetRandomNonce() uint64 {
	if x != nil {
		return x.RandomNonce
	}
	return 0
}

var File_shmsg_proto protoreflect.FileDescriptor

var file_shmsg_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x68, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73,
	0x68, 0x6d, 0x73, 0x67, 0x22, 0x1e, 0x0a, 0x02, 0x47, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x31,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x67, 0x31, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x22, 0x1e, 0x0a, 0x02, 0x47, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x32,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x67, 0x32, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x22, 0x1e, 0x0a, 0x02, 0x47, 0x54, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x74,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x67, 0x74, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x22, 0xdc, 0x01, 0x0a, 0x0b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x70, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x70, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x22, 0x42, 0x0a, 0x12, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x62, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x6f, 0x0a, 0x07, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x49, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x12, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x15, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x13, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x63, 0x0a, 0x08, 0x50, 0x6f, 0x6c, 0x79,
	0x45, 0x76, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x65, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x72, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65,
	0x64, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0e, 0x65,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x45, 0x76, 0x61, 0x6c, 0x73, 0x22, 0x3a, 0x0a,
	0x0e, 0x50, 0x6f, 0x6c, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x65, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d, 0x6d, 0x61, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x6d, 0x61, 0x73, 0x22, 0x38, 0x0a, 0x0a, 0x41, 0x63, 0x63,
	0x75, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x65, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63,
	0x75, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x63, 0x63, 0x75,
	0x73, 0x65, 0x64, 0x22, 0x56, 0x0a, 0x07, 0x41, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x65, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x08, 0x61, 0x63, 0x63, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x6f, 0x6c, 0x79, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c,
	0x52, 0x09, 0x70, 0x6f, 0x6c, 0x79, 0x45, 0x76, 0x61, 0x6c, 0x73, 0x22, 0x3a, 0x0a, 0x0c, 0x45,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0xd9, 0x03, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x68, 0x6d, 0x73,
	0x67, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52,
	0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4d, 0x0a, 0x14,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x68, 0x6d,
	0x73, 0x67, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x12, 0x62, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x2b, 0x0a, 0x08, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x73, 0x68, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x48, 0x00, 0x52,
	0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x12, 0x2e, 0x0a, 0x09, 0x70, 0x6f, 0x6c, 0x79,
	0x5f, 0x65, 0x76, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x68,
	0x6d, 0x73, 0x67, 0x2e, 0x50, 0x6f, 0x6c, 0x79, 0x45, 0x76, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x08,
	0x70, 0x6f, 0x6c, 0x79, 0x45, 0x76, 0x61, 0x6c, 0x12, 0x40, 0x0a, 0x0f, 0x70, 0x6f, 0x6c, 0x79,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x73, 0x68, 0x6d, 0x73, 0x67, 0x2e, 0x50, 0x6f, 0x6c, 0x79, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x70, 0x6f, 0x6c, 0x79,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x0a, 0x61, 0x63,
	0x63, 0x75, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x73, 0x68, 0x6d, 0x73, 0x67, 0x2e, 0x41, 0x63, 0x63, 0x75, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x75, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2a, 0x0a, 0x07, 0x61, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x73, 0x68, 0x6d, 0x73, 0x67, 0x2e, 0x41, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79,
	0x48, 0x00, 0x52, 0x07, 0x61, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x12, 0x3b, 0x0a, 0x0e, 0x65,
	0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x68, 0x6d, 0x73, 0x67, 0x2e, 0x45, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x65, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x22, 0x72, 0x0a, 0x10, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69,
	0x74, 0x68, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x68, 0x6d, 0x73, 0x67, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x5f, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x72, 0x61, 0x6e, 0x64,
	0x6f, 0x6d, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x73, 0x68,
	0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shmsg_proto_rawDescOnce sync.Once
	file_shmsg_proto_rawDescData = file_shmsg_proto_rawDesc
)

func file_shmsg_proto_rawDescGZIP() []byte {
	file_shmsg_proto_rawDescOnce.Do(func() {
		file_shmsg_proto_rawDescData = protoimpl.X.CompressGZIP(file_shmsg_proto_rawDescData)
	})
	return file_shmsg_proto_rawDescData
}

var file_shmsg_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_shmsg_proto_goTypes = []interface{}{
	(*G1)(nil),                 // 0: shmsg.G1
	(*G2)(nil),                 // 1: shmsg.G2
	(*GT)(nil),                 // 2: shmsg.GT
	(*BatchConfig)(nil),        // 3: shmsg.BatchConfig
	(*BatchConfigStarted)(nil), // 4: shmsg.BatchConfigStarted
	(*CheckIn)(nil),            // 5: shmsg.CheckIn
	(*PolyEval)(nil),           // 6: shmsg.PolyEval
	(*PolyCommitment)(nil),     // 7: shmsg.PolyCommitment
	(*Accusation)(nil),         // 8: shmsg.Accusation
	(*Apology)(nil),            // 9: shmsg.Apology
	(*EonStartVote)(nil),       // 10: shmsg.EonStartVote
	(*Message)(nil),            // 11: shmsg.Message
	(*MessageWithNonce)(nil),   // 12: shmsg.MessageWithNonce
}
var file_shmsg_proto_depIdxs = []int32{
	3,  // 0: shmsg.Message.batch_config:type_name -> shmsg.BatchConfig
	4,  // 1: shmsg.Message.batch_config_started:type_name -> shmsg.BatchConfigStarted
	5,  // 2: shmsg.Message.check_in:type_name -> shmsg.CheckIn
	6,  // 3: shmsg.Message.poly_eval:type_name -> shmsg.PolyEval
	7,  // 4: shmsg.Message.poly_commitment:type_name -> shmsg.PolyCommitment
	8,  // 5: shmsg.Message.accusation:type_name -> shmsg.Accusation
	9,  // 6: shmsg.Message.apology:type_name -> shmsg.Apology
	10, // 7: shmsg.Message.eon_start_vote:type_name -> shmsg.EonStartVote
	11, // 8: shmsg.MessageWithNonce.msg:type_name -> shmsg.Message
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_shmsg_proto_init() }
func file_shmsg_proto_init() {
	if File_shmsg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shmsg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*G1); i {
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
		file_shmsg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*G2); i {
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
		file_shmsg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GT); i {
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
		file_shmsg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchConfig); i {
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
		file_shmsg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchConfigStarted); i {
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
		file_shmsg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckIn); i {
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
		file_shmsg_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolyEval); i {
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
		file_shmsg_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolyCommitment); i {
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
		file_shmsg_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Accusation); i {
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
		file_shmsg_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Apology); i {
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
		file_shmsg_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EonStartVote); i {
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
		file_shmsg_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_shmsg_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithNonce); i {
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
	file_shmsg_proto_msgTypes[11].OneofWrappers = []interface{}{
		(*Message_BatchConfig)(nil),
		(*Message_BatchConfigStarted)(nil),
		(*Message_CheckIn)(nil),
		(*Message_PolyEval)(nil),
		(*Message_PolyCommitment)(nil),
		(*Message_Accusation)(nil),
		(*Message_Apology)(nil),
		(*Message_EonStartVote)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shmsg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shmsg_proto_goTypes,
		DependencyIndexes: file_shmsg_proto_depIdxs,
		MessageInfos:      file_shmsg_proto_msgTypes,
	}.Build()
	File_shmsg_proto = out.File
	file_shmsg_proto_rawDesc = nil
	file_shmsg_proto_goTypes = nil
	file_shmsg_proto_depIdxs = nil
}
