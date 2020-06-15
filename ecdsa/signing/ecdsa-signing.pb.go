// Copyright © 2019 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: protob/ecdsa-signing.proto

package signing

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//
// Represents a P2P message sent to each party during Phase 1 of the GG20 ECDSA TSS signing protocol.
type SignRound1Message1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C               []byte   `protobuf:"bytes,1,opt,name=c,proto3" json:"c,omitempty"`
	RangeProofAlice [][]byte `protobuf:"bytes,2,rep,name=range_proof_alice,json=rangeProofAlice,proto3" json:"range_proof_alice,omitempty"`
}

func (x *SignRound1Message1) Reset() {
	*x = SignRound1Message1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_signing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRound1Message1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRound1Message1) ProtoMessage() {}

func (x *SignRound1Message1) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_signing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRound1Message1.ProtoReflect.Descriptor instead.
func (*SignRound1Message1) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_signing_proto_rawDescGZIP(), []int{0}
}

func (x *SignRound1Message1) GetC() []byte {
	if x != nil {
		return x.C
	}
	return nil
}

func (x *SignRound1Message1) GetRangeProofAlice() [][]byte {
	if x != nil {
		return x.RangeProofAlice
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Phase 1 of the GG20 ECDSA TSS signing protocol.
type SignRound1Message2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commitment []byte `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
}

func (x *SignRound1Message2) Reset() {
	*x = SignRound1Message2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_signing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRound1Message2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRound1Message2) ProtoMessage() {}

func (x *SignRound1Message2) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_signing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRound1Message2.ProtoReflect.Descriptor instead.
func (*SignRound1Message2) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_signing_proto_rawDescGZIP(), []int{1}
}

func (x *SignRound1Message2) GetCommitment() []byte {
	if x != nil {
		return x.Commitment
	}
	return nil
}

//
// Represents a P2P message sent to each party during Phase 2 of the GG20 ECDSA TSS signing protocol.
type SignRound2Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C1         []byte   `protobuf:"bytes,1,opt,name=c1,proto3" json:"c1,omitempty"`
	C2         []byte   `protobuf:"bytes,2,opt,name=c2,proto3" json:"c2,omitempty"`
	ProofBob   [][]byte `protobuf:"bytes,3,rep,name=proof_bob,json=proofBob,proto3" json:"proof_bob,omitempty"`
	ProofBobWc [][]byte `protobuf:"bytes,4,rep,name=proof_bob_wc,json=proofBobWc,proto3" json:"proof_bob_wc,omitempty"`
}

func (x *SignRound2Message) Reset() {
	*x = SignRound2Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_signing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRound2Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRound2Message) ProtoMessage() {}

func (x *SignRound2Message) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_signing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRound2Message.ProtoReflect.Descriptor instead.
func (*SignRound2Message) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_signing_proto_rawDescGZIP(), []int{2}
}

func (x *SignRound2Message) GetC1() []byte {
	if x != nil {
		return x.C1
	}
	return nil
}

func (x *SignRound2Message) GetC2() []byte {
	if x != nil {
		return x.C2
	}
	return nil
}

func (x *SignRound2Message) GetProofBob() [][]byte {
	if x != nil {
		return x.ProofBob
	}
	return nil
}

func (x *SignRound2Message) GetProofBobWc() [][]byte {
	if x != nil {
		return x.ProofBobWc
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Phase 3 of the GG20 ECDSA TSS signing protocol.
type SignRound3Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeltaI       []byte `protobuf:"bytes,1,opt,name=delta_i,json=deltaI,proto3" json:"delta_i,omitempty"`
	TIX          []byte `protobuf:"bytes,2,opt,name=t_i_x,json=tIX,proto3" json:"t_i_x,omitempty"`
	TIY          []byte `protobuf:"bytes,3,opt,name=t_i_y,json=tIY,proto3" json:"t_i_y,omitempty"`
	TProofAlphaX []byte `protobuf:"bytes,4,opt,name=t_proof_alpha_x,json=tProofAlphaX,proto3" json:"t_proof_alpha_x,omitempty"`
	TProofAlphaY []byte `protobuf:"bytes,5,opt,name=t_proof_alpha_y,json=tProofAlphaY,proto3" json:"t_proof_alpha_y,omitempty"`
	TProofT      []byte `protobuf:"bytes,6,opt,name=t_proof_t,json=tProofT,proto3" json:"t_proof_t,omitempty"`
	TProofU      []byte `protobuf:"bytes,7,opt,name=t_proof_u,json=tProofU,proto3" json:"t_proof_u,omitempty"`
}

func (x *SignRound3Message) Reset() {
	*x = SignRound3Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_signing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRound3Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRound3Message) ProtoMessage() {}

func (x *SignRound3Message) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_signing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRound3Message.ProtoReflect.Descriptor instead.
func (*SignRound3Message) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_signing_proto_rawDescGZIP(), []int{3}
}

func (x *SignRound3Message) GetDeltaI() []byte {
	if x != nil {
		return x.DeltaI
	}
	return nil
}

func (x *SignRound3Message) GetTIX() []byte {
	if x != nil {
		return x.TIX
	}
	return nil
}

func (x *SignRound3Message) GetTIY() []byte {
	if x != nil {
		return x.TIY
	}
	return nil
}

func (x *SignRound3Message) GetTProofAlphaX() []byte {
	if x != nil {
		return x.TProofAlphaX
	}
	return nil
}

func (x *SignRound3Message) GetTProofAlphaY() []byte {
	if x != nil {
		return x.TProofAlphaY
	}
	return nil
}

func (x *SignRound3Message) GetTProofT() []byte {
	if x != nil {
		return x.TProofT
	}
	return nil
}

func (x *SignRound3Message) GetTProofU() []byte {
	if x != nil {
		return x.TProofU
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Phase 4 of the GG20 ECDSA TSS signing protocol.
type SignRound4Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeCommitment [][]byte `protobuf:"bytes,1,rep,name=de_commitment,json=deCommitment,proto3" json:"de_commitment,omitempty"`
}

func (x *SignRound4Message) Reset() {
	*x = SignRound4Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_signing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRound4Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRound4Message) ProtoMessage() {}

func (x *SignRound4Message) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_signing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRound4Message.ProtoReflect.Descriptor instead.
func (*SignRound4Message) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_signing_proto_rawDescGZIP(), []int{4}
}

func (x *SignRound4Message) GetDeCommitment() [][]byte {
	if x != nil {
		return x.DeCommitment
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Phase 5 of the GG20 ECDSA TSS signing protocol.
type SignRound5Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RIX            []byte   `protobuf:"bytes,1,opt,name=r_i_x,json=rIX,proto3" json:"r_i_x,omitempty"`
	RIY            []byte   `protobuf:"bytes,2,opt,name=r_i_y,json=rIY,proto3" json:"r_i_y,omitempty"`
	ProofPdlWSlack [][]byte `protobuf:"bytes,3,rep,name=proof_pdl_w_slack,json=proofPdlWSlack,proto3" json:"proof_pdl_w_slack,omitempty"`
}

func (x *SignRound5Message) Reset() {
	*x = SignRound5Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_signing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRound5Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRound5Message) ProtoMessage() {}

func (x *SignRound5Message) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_signing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRound5Message.ProtoReflect.Descriptor instead.
func (*SignRound5Message) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_signing_proto_rawDescGZIP(), []int{5}
}

func (x *SignRound5Message) GetRIX() []byte {
	if x != nil {
		return x.RIX
	}
	return nil
}

func (x *SignRound5Message) GetRIY() []byte {
	if x != nil {
		return x.RIY
	}
	return nil
}

func (x *SignRound5Message) GetProofPdlWSlack() [][]byte {
	if x != nil {
		return x.ProofPdlWSlack
	}
	return nil
}

//
// Represents a BROADCAST message sent to all parties during Phase 6 of the GG20 ECDSA TSS signing protocol.
type SignRound6Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Content:
	//	*SignRound6Message_Success
	//	*SignRound6Message_Abort
	Content isSignRound6Message_Content `protobuf_oneof:"content"`
}

func (x *SignRound6Message) Reset() {
	*x = SignRound6Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_signing_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRound6Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRound6Message) ProtoMessage() {}

func (x *SignRound6Message) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_signing_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRound6Message.ProtoReflect.Descriptor instead.
func (*SignRound6Message) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_signing_proto_rawDescGZIP(), []int{6}
}

func (m *SignRound6Message) GetContent() isSignRound6Message_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *SignRound6Message) GetSuccess() *SignRound6Message_SuccessData {
	if x, ok := x.GetContent().(*SignRound6Message_Success); ok {
		return x.Success
	}
	return nil
}

func (x *SignRound6Message) GetAbort() *SignRound6Message_AbortData {
	if x, ok := x.GetContent().(*SignRound6Message_Abort); ok {
		return x.Abort
	}
	return nil
}

type isSignRound6Message_Content interface {
	isSignRound6Message_Content()
}

type SignRound6Message_Success struct {
	Success *SignRound6Message_SuccessData `protobuf:"bytes,1,opt,name=success,proto3,oneof"`
}

type SignRound6Message_Abort struct {
	Abort *SignRound6Message_AbortData `protobuf:"bytes,2,opt,name=abort,proto3,oneof"`
}

func (*SignRound6Message_Success) isSignRound6Message_Content() {}

func (*SignRound6Message_Abort) isSignRound6Message_Content() {}

//
// Represents a BROADCAST message sent to all parties during Phase 7 of the GG20 ECDSA TSS signing protocol.
type SignRound7Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SI []byte `protobuf:"bytes,1,opt,name=s_i,json=sI,proto3" json:"s_i,omitempty"`
}

func (x *SignRound7Message) Reset() {
	*x = SignRound7Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_signing_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRound7Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRound7Message) ProtoMessage() {}

func (x *SignRound7Message) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_signing_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRound7Message.ProtoReflect.Descriptor instead.
func (*SignRound7Message) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_signing_proto_rawDescGZIP(), []int{7}
}

func (x *SignRound7Message) GetSI() []byte {
	if x != nil {
		return x.SI
	}
	return nil
}

type SignRound6Message_SuccessData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SIX           []byte `protobuf:"bytes,1,opt,name=s_i_x,json=sIX,proto3" json:"s_i_x,omitempty"`
	SIY           []byte `protobuf:"bytes,2,opt,name=s_i_y,json=sIY,proto3" json:"s_i_y,omitempty"`
	StProofAlphaX []byte `protobuf:"bytes,3,opt,name=st_proof_alpha_x,json=stProofAlphaX,proto3" json:"st_proof_alpha_x,omitempty"`
	StProofAlphaY []byte `protobuf:"bytes,4,opt,name=st_proof_alpha_y,json=stProofAlphaY,proto3" json:"st_proof_alpha_y,omitempty"`
	StProofBetaX  []byte `protobuf:"bytes,5,opt,name=st_proof_beta_x,json=stProofBetaX,proto3" json:"st_proof_beta_x,omitempty"`
	StProofBetaY  []byte `protobuf:"bytes,6,opt,name=st_proof_beta_y,json=stProofBetaY,proto3" json:"st_proof_beta_y,omitempty"`
	StProofT      []byte `protobuf:"bytes,7,opt,name=st_proof_t,json=stProofT,proto3" json:"st_proof_t,omitempty"`
	StProofU      []byte `protobuf:"bytes,8,opt,name=st_proof_u,json=stProofU,proto3" json:"st_proof_u,omitempty"`
}

func (x *SignRound6Message_SuccessData) Reset() {
	*x = SignRound6Message_SuccessData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_signing_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRound6Message_SuccessData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRound6Message_SuccessData) ProtoMessage() {}

func (x *SignRound6Message_SuccessData) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_signing_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRound6Message_SuccessData.ProtoReflect.Descriptor instead.
func (*SignRound6Message_SuccessData) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_signing_proto_rawDescGZIP(), []int{6, 0}
}

func (x *SignRound6Message_SuccessData) GetSIX() []byte {
	if x != nil {
		return x.SIX
	}
	return nil
}

func (x *SignRound6Message_SuccessData) GetSIY() []byte {
	if x != nil {
		return x.SIY
	}
	return nil
}

func (x *SignRound6Message_SuccessData) GetStProofAlphaX() []byte {
	if x != nil {
		return x.StProofAlphaX
	}
	return nil
}

func (x *SignRound6Message_SuccessData) GetStProofAlphaY() []byte {
	if x != nil {
		return x.StProofAlphaY
	}
	return nil
}

func (x *SignRound6Message_SuccessData) GetStProofBetaX() []byte {
	if x != nil {
		return x.StProofBetaX
	}
	return nil
}

func (x *SignRound6Message_SuccessData) GetStProofBetaY() []byte {
	if x != nil {
		return x.StProofBetaY
	}
	return nil
}

func (x *SignRound6Message_SuccessData) GetStProofT() []byte {
	if x != nil {
		return x.StProofT
	}
	return nil
}

func (x *SignRound6Message_SuccessData) GetStProofU() []byte {
	if x != nil {
		return x.StProofU
	}
	return nil
}

type SignRound6Message_AbortData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KI           []byte   `protobuf:"bytes,1,opt,name=k_i,json=kI,proto3" json:"k_i,omitempty"`
	KIRandomness []byte   `protobuf:"bytes,2,opt,name=k_i_randomness,json=kIRandomness,proto3" json:"k_i_randomness,omitempty"`
	GammaI       []byte   `protobuf:"bytes,3,opt,name=gamma_i,json=gammaI,proto3" json:"gamma_i,omitempty"`
	AlphaIJ      [][]byte `protobuf:"bytes,4,rep,name=alpha_i_j,json=alphaIJ,proto3" json:"alpha_i_j,omitempty"`
	BetaJI       [][]byte `protobuf:"bytes,5,rep,name=beta_j_i,json=betaJI,proto3" json:"beta_j_i,omitempty"`
}

func (x *SignRound6Message_AbortData) Reset() {
	*x = SignRound6Message_AbortData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protob_ecdsa_signing_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRound6Message_AbortData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRound6Message_AbortData) ProtoMessage() {}

func (x *SignRound6Message_AbortData) ProtoReflect() protoreflect.Message {
	mi := &file_protob_ecdsa_signing_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRound6Message_AbortData.ProtoReflect.Descriptor instead.
func (*SignRound6Message_AbortData) Descriptor() ([]byte, []int) {
	return file_protob_ecdsa_signing_proto_rawDescGZIP(), []int{6, 1}
}

func (x *SignRound6Message_AbortData) GetKI() []byte {
	if x != nil {
		return x.KI
	}
	return nil
}

func (x *SignRound6Message_AbortData) GetKIRandomness() []byte {
	if x != nil {
		return x.KIRandomness
	}
	return nil
}

func (x *SignRound6Message_AbortData) GetGammaI() []byte {
	if x != nil {
		return x.GammaI
	}
	return nil
}

func (x *SignRound6Message_AbortData) GetAlphaIJ() [][]byte {
	if x != nil {
		return x.AlphaIJ
	}
	return nil
}

func (x *SignRound6Message_AbortData) GetBetaJI() [][]byte {
	if x != nil {
		return x.BetaJI
	}
	return nil
}

var File_protob_ecdsa_signing_proto protoreflect.FileDescriptor

var file_protob_ecdsa_signing_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x2f, 0x65, 0x63, 0x64, 0x73, 0x61, 0x2d, 0x73,
	0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x12,
	0x53, 0x69, 0x67, 0x6e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x31, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x31, 0x12, 0x0c, 0x0a, 0x01, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x63,
	0x12, 0x2a, 0x0a, 0x11, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f,
	0x61, 0x6c, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0f, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x41, 0x6c, 0x69, 0x63, 0x65, 0x22, 0x34, 0x0a, 0x12,
	0x53, 0x69, 0x67, 0x6e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x31, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x72, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x32,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x02, 0x63, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x32, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x02, 0x63, 0x32, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6f, 0x66,
	0x5f, 0x62, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6f,
	0x66, 0x42, 0x6f, 0x62, 0x12, 0x20, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x62, 0x6f,
	0x62, 0x5f, 0x77, 0x63, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6f,
	0x66, 0x42, 0x6f, 0x62, 0x57, 0x63, 0x22, 0xda, 0x01, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e, 0x52,
	0x6f, 0x75, 0x6e, 0x64, 0x33, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x64, 0x65, 0x6c, 0x74, 0x61, 0x5f, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x64,
	0x65, 0x6c, 0x74, 0x61, 0x49, 0x12, 0x12, 0x0a, 0x05, 0x74, 0x5f, 0x69, 0x5f, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x74, 0x49, 0x58, 0x12, 0x12, 0x0a, 0x05, 0x74, 0x5f, 0x69,
	0x5f, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x74, 0x49, 0x59, 0x12, 0x25, 0x0a,
	0x0f, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5f, 0x78,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x41, 0x6c,
	0x70, 0x68, 0x61, 0x58, 0x12, 0x25, 0x0a, 0x0f, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x5f, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x74,
	0x50, 0x72, 0x6f, 0x6f, 0x66, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x59, 0x12, 0x1a, 0x0a, 0x09, 0x74,
	0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x54, 0x12, 0x1a, 0x0a, 0x09, 0x74, 0x5f, 0x70, 0x72, 0x6f,
	0x6f, 0x66, 0x5f, 0x75, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x50, 0x72, 0x6f,
	0x6f, 0x66, 0x55, 0x22, 0x38, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x6f, 0x75, 0x6e, 0x64,
	0x34, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x0c, 0x64, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x66, 0x0a,
	0x11, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x35, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x05, 0x72, 0x5f, 0x69, 0x5f, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x72, 0x49, 0x58, 0x12, 0x12, 0x0a, 0x05, 0x72, 0x5f, 0x69, 0x5f, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x72, 0x49, 0x59, 0x12, 0x29, 0x0a, 0x11, 0x70, 0x72,
	0x6f, 0x6f, 0x66, 0x5f, 0x70, 0x64, 0x6c, 0x5f, 0x77, 0x5f, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x50, 0x64, 0x6c, 0x57,
	0x53, 0x6c, 0x61, 0x63, 0x6b, 0x22, 0xb8, 0x04, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x6f,
	0x75, 0x6e, 0x64, 0x36, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x36, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x34, 0x0a, 0x05, 0x61, 0x62, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x6f, 0x75,
	0x6e, 0x64, 0x36, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x62, 0x6f, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x05, 0x61, 0x62, 0x6f, 0x72, 0x74, 0x1a, 0x91, 0x02,
	0x0a, 0x0b, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x05, 0x73, 0x5f, 0x69, 0x5f, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x49,
	0x58, 0x12, 0x12, 0x0a, 0x05, 0x73, 0x5f, 0x69, 0x5f, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x73, 0x49, 0x59, 0x12, 0x27, 0x0a, 0x10, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x6f,
	0x66, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5f, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0d, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x58, 0x12, 0x27,
	0x0a, 0x10, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x5f, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6f,
	0x66, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x59, 0x12, 0x25, 0x0a, 0x0f, 0x73, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x6f, 0x66, 0x5f, 0x62, 0x65, 0x74, 0x61, 0x5f, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0c, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x42, 0x65, 0x74, 0x61, 0x58, 0x12, 0x25,
	0x0a, 0x0f, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x62, 0x65, 0x74, 0x61, 0x5f,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66,
	0x42, 0x65, 0x74, 0x61, 0x59, 0x12, 0x1c, 0x0a, 0x0a, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x6f,
	0x66, 0x5f, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x6f, 0x66, 0x54, 0x12, 0x1c, 0x0a, 0x0a, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f,
	0x75, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66,
	0x55, 0x1a, 0x91, 0x01, 0x0a, 0x09, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x0f, 0x0a, 0x03, 0x6b, 0x5f, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x6b, 0x49,
	0x12, 0x24, 0x0a, 0x0e, 0x6b, 0x5f, 0x69, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x6e, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6b, 0x49, 0x52, 0x61, 0x6e, 0x64,
	0x6f, 0x6d, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x6d, 0x61, 0x5f,
	0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x6d, 0x61, 0x49, 0x12,
	0x1a, 0x0a, 0x09, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5f, 0x69, 0x5f, 0x6a, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x07, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x49, 0x4a, 0x12, 0x18, 0x0a, 0x08, 0x62,
	0x65, 0x74, 0x61, 0x5f, 0x6a, 0x5f, 0x69, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x62,
	0x65, 0x74, 0x61, 0x4a, 0x49, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x24, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x37, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0f, 0x0a, 0x03, 0x73, 0x5f, 0x69, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x02, 0x73, 0x49, 0x42, 0x0f, 0x5a, 0x0d, 0x65, 0x63, 0x64, 0x73, 0x61, 0x2f,
	0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protob_ecdsa_signing_proto_rawDescOnce sync.Once
	file_protob_ecdsa_signing_proto_rawDescData = file_protob_ecdsa_signing_proto_rawDesc
)

func file_protob_ecdsa_signing_proto_rawDescGZIP() []byte {
	file_protob_ecdsa_signing_proto_rawDescOnce.Do(func() {
		file_protob_ecdsa_signing_proto_rawDescData = protoimpl.X.CompressGZIP(file_protob_ecdsa_signing_proto_rawDescData)
	})
	return file_protob_ecdsa_signing_proto_rawDescData
}

var file_protob_ecdsa_signing_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_protob_ecdsa_signing_proto_goTypes = []interface{}{
	(*SignRound1Message1)(nil),            // 0: SignRound1Message1
	(*SignRound1Message2)(nil),            // 1: SignRound1Message2
	(*SignRound2Message)(nil),             // 2: SignRound2Message
	(*SignRound3Message)(nil),             // 3: SignRound3Message
	(*SignRound4Message)(nil),             // 4: SignRound4Message
	(*SignRound5Message)(nil),             // 5: SignRound5Message
	(*SignRound6Message)(nil),             // 6: SignRound6Message
	(*SignRound7Message)(nil),             // 7: SignRound7Message
	(*SignRound6Message_SuccessData)(nil), // 8: SignRound6Message.SuccessData
	(*SignRound6Message_AbortData)(nil),   // 9: SignRound6Message.AbortData
}
var file_protob_ecdsa_signing_proto_depIdxs = []int32{
	8, // 0: SignRound6Message.success:type_name -> SignRound6Message.SuccessData
	9, // 1: SignRound6Message.abort:type_name -> SignRound6Message.AbortData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protob_ecdsa_signing_proto_init() }
func file_protob_ecdsa_signing_proto_init() {
	if File_protob_ecdsa_signing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protob_ecdsa_signing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRound1Message1); i {
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
		file_protob_ecdsa_signing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRound1Message2); i {
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
		file_protob_ecdsa_signing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRound2Message); i {
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
		file_protob_ecdsa_signing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRound3Message); i {
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
		file_protob_ecdsa_signing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRound4Message); i {
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
		file_protob_ecdsa_signing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRound5Message); i {
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
		file_protob_ecdsa_signing_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRound6Message); i {
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
		file_protob_ecdsa_signing_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRound7Message); i {
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
		file_protob_ecdsa_signing_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRound6Message_SuccessData); i {
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
		file_protob_ecdsa_signing_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRound6Message_AbortData); i {
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
	file_protob_ecdsa_signing_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*SignRound6Message_Success)(nil),
		(*SignRound6Message_Abort)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protob_ecdsa_signing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protob_ecdsa_signing_proto_goTypes,
		DependencyIndexes: file_protob_ecdsa_signing_proto_depIdxs,
		MessageInfos:      file_protob_ecdsa_signing_proto_msgTypes,
	}.Build()
	File_protob_ecdsa_signing_proto = out.File
	file_protob_ecdsa_signing_proto_rawDesc = nil
	file_protob_ecdsa_signing_proto_goTypes = nil
	file_protob_ecdsa_signing_proto_depIdxs = nil
}
