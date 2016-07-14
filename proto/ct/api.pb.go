// Code generated by protoc-gen-go.
// source: proto/ct/api.proto
// DO NOT EDIT!

/*
Package ct is a generated protocol buffer package.

It is generated from these files:
	proto/ct/api.proto
	proto/ct/ct.proto

It has these top-level messages:
	AddChainRequest
	AddChainResponse
	GetSTHResponse
	Base64LeafEntry
	GetEntriesResponse
	GetConsistencyProofResponse
	GetAuditProofResponse
	GetAcceptedRootsResponse
	GetEntryAndProofResponse
	DigitallySigned
	X509ChainEntry
	PreCert
	CertInfo
	PrecertChainEntry
	XJSONEntry
	LogEntry
	LogID
	SctExtension
	SignedCertificateTimestamp
	SignedCertificateTimestampList
	SignedEntry
	TimestampedEntry
	MerkleTreeLeaf
	MerkleAuditProof
	ShortMerkleAuditProof
	LoggedEntryPB
	SthExtension
	SignedTreeHead
	SSLClientCTData
	ClusterNodeState
	ClusterControl
	ClusterConfig
	SequenceMapping
*/
package ct

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// AddChainRequest represents the JSON request body sent to the add-chain CT
// method.
type AddChainRequest struct {
	Chain []string `protobuf:"bytes,1,rep,name=chain" json:"chain,omitempty"`
}

func (m *AddChainRequest) Reset()                    { *m = AddChainRequest{} }
func (m *AddChainRequest) String() string            { return proto.CompactTextString(m) }
func (*AddChainRequest) ProtoMessage()               {}
func (*AddChainRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// AddChainResponse represents the JSON response to the add-chain CT method.
// An SCT represents a Log's promise to integrate a [pre-]certificate into the
// log within a defined period of time.
type AddChainResponse struct {
	SctVersion Version `protobuf:"varint,1,opt,name=sct_version,json=sctVersion,enum=ct.Version" json:"sct_version,omitempty"`
	Id         string  `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Timestamp  uint64  `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Extensions string  `protobuf:"bytes,4,opt,name=extensions" json:"extensions,omitempty"`
	Signature  string  `protobuf:"bytes,5,opt,name=signature" json:"signature,omitempty"`
}

func (m *AddChainResponse) Reset()                    { *m = AddChainResponse{} }
func (m *AddChainResponse) String() string            { return proto.CompactTextString(m) }
func (*AddChainResponse) ProtoMessage()               {}
func (*AddChainResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// GetSTHResponse respresents the JSON response to the get-sth CT method
type GetSTHResponse struct {
	TreeSize          uint64 `protobuf:"varint,1,opt,name=tree_size,json=treeSize" json:"tree_size,omitempty"`
	Timestamp         uint64 `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Sha256RootHash    string `protobuf:"bytes,3,opt,name=sha256_root_hash,json=sha256RootHash" json:"sha256_root_hash,omitempty"`
	TreeHeadSignature string `protobuf:"bytes,4,opt,name=tree_head_signature,json=treeHeadSignature" json:"tree_head_signature,omitempty"`
}

func (m *GetSTHResponse) Reset()                    { *m = GetSTHResponse{} }
func (m *GetSTHResponse) String() string            { return proto.CompactTextString(m) }
func (*GetSTHResponse) ProtoMessage()               {}
func (*GetSTHResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Base64LeafEntry respresents a Base64 encoded leaf entry
type Base64LeafEntry struct {
	LeafInput string `protobuf:"bytes,1,opt,name=leaf_input,json=leafInput" json:"leaf_input,omitempty"`
	ExtraData string `protobuf:"bytes,2,opt,name=extra_data,json=extraData" json:"extra_data,omitempty"`
}

func (m *Base64LeafEntry) Reset()                    { *m = Base64LeafEntry{} }
func (m *Base64LeafEntry) String() string            { return proto.CompactTextString(m) }
func (*Base64LeafEntry) ProtoMessage()               {}
func (*Base64LeafEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// GetEntriesReponse respresents the JSON response to the CT get-entries method
type GetEntriesResponse struct {
	Entries []*Base64LeafEntry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *GetEntriesResponse) Reset()                    { *m = GetEntriesResponse{} }
func (m *GetEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*GetEntriesResponse) ProtoMessage()               {}
func (*GetEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetEntriesResponse) GetEntries() []*Base64LeafEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// GetConsistencyProofResponse represents the JSON response to the CT get-consistency-proof method
type GetConsistencyProofResponse struct {
	Consistency []string `protobuf:"bytes,1,rep,name=consistency" json:"consistency,omitempty"`
}

func (m *GetConsistencyProofResponse) Reset()                    { *m = GetConsistencyProofResponse{} }
func (m *GetConsistencyProofResponse) String() string            { return proto.CompactTextString(m) }
func (*GetConsistencyProofResponse) ProtoMessage()               {}
func (*GetConsistencyProofResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// GetAuditProofResponse represents the JSON response to the CT get-audit-proof method
type GetAuditProofResponse struct {
	Hash     []string `protobuf:"bytes,1,rep,name=hash" json:"hash,omitempty"`
	TreeSize uint64   `protobuf:"varint,2,opt,name=tree_size,json=treeSize" json:"tree_size,omitempty"`
}

func (m *GetAuditProofResponse) Reset()                    { *m = GetAuditProofResponse{} }
func (m *GetAuditProofResponse) String() string            { return proto.CompactTextString(m) }
func (*GetAuditProofResponse) ProtoMessage()               {}
func (*GetAuditProofResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// GetAcceptedRootsResponse represents the JSON response to the CT get-roots method.
type GetAcceptedRootsResponse struct {
	Certificates []string `protobuf:"bytes,1,rep,name=certificates" json:"certificates,omitempty"`
}

func (m *GetAcceptedRootsResponse) Reset()                    { *m = GetAcceptedRootsResponse{} }
func (m *GetAcceptedRootsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetAcceptedRootsResponse) ProtoMessage()               {}
func (*GetAcceptedRootsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// GetEntryAndProodReponse represents the JSON response to the CT get-entry-and-proof method
type GetEntryAndProofResponse struct {
	LeafInput string   `protobuf:"bytes,1,opt,name=leaf_input,json=leafInput" json:"leaf_input,omitempty"`
	ExtraData string   `protobuf:"bytes,2,opt,name=extra_data,json=extraData" json:"extra_data,omitempty"`
	AuditPath []string `protobuf:"bytes,3,rep,name=audit_path,json=auditPath" json:"audit_path,omitempty"`
}

func (m *GetEntryAndProofResponse) Reset()                    { *m = GetEntryAndProofResponse{} }
func (m *GetEntryAndProofResponse) String() string            { return proto.CompactTextString(m) }
func (*GetEntryAndProofResponse) ProtoMessage()               {}
func (*GetEntryAndProofResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto.RegisterType((*AddChainRequest)(nil), "ct.AddChainRequest")
	proto.RegisterType((*AddChainResponse)(nil), "ct.AddChainResponse")
	proto.RegisterType((*GetSTHResponse)(nil), "ct.GetSTHResponse")
	proto.RegisterType((*Base64LeafEntry)(nil), "ct.Base64LeafEntry")
	proto.RegisterType((*GetEntriesResponse)(nil), "ct.GetEntriesResponse")
	proto.RegisterType((*GetConsistencyProofResponse)(nil), "ct.GetConsistencyProofResponse")
	proto.RegisterType((*GetAuditProofResponse)(nil), "ct.GetAuditProofResponse")
	proto.RegisterType((*GetAcceptedRootsResponse)(nil), "ct.GetAcceptedRootsResponse")
	proto.RegisterType((*GetEntryAndProofResponse)(nil), "ct.GetEntryAndProofResponse")
}

func init() { proto.RegisterFile("proto/ct/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x93, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0xd5, 0xac, 0x03, 0x72, 0x45, 0xdd, 0xe6, 0x81, 0x14, 0x31, 0x40, 0x93, 0x5f, 0xe8,
	0x03, 0x74, 0x52, 0x81, 0x3d, 0x82, 0x4a, 0x41, 0x14, 0x09, 0x89, 0xc9, 0x45, 0xbc, 0x46, 0xc6,
	0xb9, 0x12, 0x4b, 0x2c, 0x09, 0xf1, 0x15, 0x18, 0xdf, 0x87, 0x37, 0x3e, 0x24, 0x67, 0xa7, 0x24,
	0x4b, 0x5f, 0x79, 0xf3, 0xfd, 0xce, 0x77, 0xf7, 0xff, 0x9f, 0x13, 0x10, 0x55, 0x5d, 0x52, 0x79,
	0x66, 0xe8, 0x4c, 0x57, 0x76, 0x1a, 0x02, 0x11, 0x19, 0xba, 0x77, 0xd4, 0x72, 0x43, 0x0d, 0x96,
	0x8f, 0xe0, 0x60, 0x9e, 0x65, 0x8b, 0x5c, 0xdb, 0x42, 0xe1, 0xb7, 0x0d, 0x3a, 0x12, 0x77, 0x60,
	0xdf, 0xf8, 0x38, 0x19, 0x9c, 0xee, 0x4d, 0x62, 0xd5, 0x04, 0xf2, 0xcf, 0x00, 0x0e, 0xbb, 0x9b,
	0xae, 0x2a, 0x0b, 0x87, 0xe2, 0x31, 0x8c, 0x9c, 0xa1, 0xf4, 0x3b, 0xd6, 0xce, 0x96, 0xbe, 0x60,
	0x30, 0x19, 0xcf, 0x46, 0x53, 0xee, 0xfe, 0xa9, 0x41, 0x0a, 0x38, 0xbf, 0x3d, 0x8b, 0x31, 0x44,
	0x36, 0x4b, 0x22, 0xbe, 0x14, 0x2b, 0x3e, 0x89, 0xfb, 0x10, 0x93, 0xbd, 0xe4, 0x91, 0xfa, 0xb2,
	0x4a, 0xf6, 0x18, 0x0f, 0x55, 0x07, 0xc4, 0x43, 0x00, 0xfc, 0x49, 0x58, 0xf8, 0x52, 0x97, 0x0c,
	0x43, 0xd5, 0x35, 0xe2, 0xab, 0x9d, 0xfd, 0x52, 0x68, 0xda, 0xd4, 0x98, 0xec, 0x87, 0x74, 0x07,
	0xe4, 0xef, 0x01, 0x8c, 0xdf, 0x22, 0xad, 0x3e, 0x2e, 0x5b, 0xb1, 0x27, 0x3c, 0xae, 0x46, 0x4c,
	0x9d, 0xfd, 0x85, 0x41, 0xea, 0x50, 0xdd, 0xf2, 0x60, 0xc5, 0x71, 0x5f, 0x4b, 0xb4, 0xab, 0x65,
	0x02, 0x87, 0x2e, 0xd7, 0xb3, 0xe7, 0xe7, 0x69, 0x5d, 0x96, 0x94, 0xe6, 0xda, 0xe5, 0x41, 0x70,
	0xac, 0xc6, 0x0d, 0x57, 0x8c, 0x97, 0x4c, 0xc5, 0x14, 0x8e, 0xc3, 0x90, 0x1c, 0x75, 0x96, 0x76,
	0xfa, 0x1a, 0xf9, 0x47, 0x3e, 0xb5, 0xe4, 0xcc, 0xaa, 0xd5, 0xf9, 0x01, 0x0e, 0x5e, 0x69, 0x87,
	0xe7, 0xcf, 0xde, 0xa3, 0x5e, 0xbf, 0x29, 0xa8, 0xbe, 0x12, 0x0f, 0x00, 0xbe, 0x72, 0x90, 0xda,
	0xa2, 0xda, 0x50, 0x10, 0xca, 0xce, 0x3c, 0x79, 0xe7, 0x81, 0x4f, 0xf3, 0x16, 0x6a, 0x9d, 0x66,
	0x9a, 0xf4, 0x76, 0x9b, 0x71, 0x20, 0xaf, 0x19, 0xc8, 0x05, 0x08, 0xf6, 0xed, 0x3b, 0x59, 0x74,
	0xad, 0xf7, 0x27, 0x70, 0x13, 0x1b, 0x14, 0x5e, 0x75, 0x34, 0x3b, 0xf6, 0x8f, 0xb4, 0x33, 0x59,
	0xfd, 0xbb, 0x23, 0x5f, 0xc2, 0x09, 0x37, 0x59, 0x70, 0xa9, 0x75, 0xbc, 0x70, 0x73, 0x75, 0xc1,
	0xbe, 0xd7, 0x6d, 0xb7, 0x53, 0x18, 0x99, 0x2e, 0xb7, 0xfd, 0x4e, 0xae, 0x23, 0xb9, 0x84, 0xbb,
	0xdc, 0x60, 0xbe, 0xc9, 0x2c, 0xf5, 0x4b, 0x05, 0x0c, 0xc3, 0xf6, 0x9a, 0x9a, 0x70, 0xee, 0x3f,
	0x4c, 0xd4, 0x7f, 0x18, 0xf9, 0x02, 0x12, 0xdf, 0xc9, 0x18, 0xac, 0x08, 0x33, 0xbf, 0xe7, 0xce,
	0x95, 0x84, 0xdb, 0x06, 0x6b, 0xb2, 0x6b, 0x6b, 0x34, 0x6d, 0xad, 0xc5, 0xaa, 0xc7, 0xe4, 0x8f,
	0x50, 0x1f, 0xfc, 0xcd, 0x8b, 0xac, 0x2f, 0xe6, 0xbf, 0x36, 0xed, 0xd3, 0xda, 0x1b, 0x4c, 0x2b,
	0x4d, 0xfe, 0x73, 0xf0, 0xb3, 0xe3, 0x40, 0x2e, 0x18, 0x7c, 0xbe, 0x11, 0x7e, 0xb0, 0xa7, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x81, 0x3b, 0xb2, 0x5e, 0x8d, 0x03, 0x00, 0x00,
}