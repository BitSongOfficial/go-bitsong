// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fantoken/fantoken.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/x/bank/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type SocialProof struct {
	Social   string    `protobuf:"bytes,1,opt,name=social,proto3" json:"social,omitempty"`
	Handler  string    `protobuf:"bytes,2,opt,name=handler,proto3" json:"handler,omitempty"`
	ProofUri string    `protobuf:"bytes,3,opt,name=proof_uri,json=proofUri,proto3" json:"proof_uri,omitempty"`
	ExpireAt time.Time `protobuf:"bytes,4,opt,name=expire_at,json=expireAt,proto3,stdtime" json:"expire_at" yaml:"expire_at"`
}

func (m *SocialProof) Reset()         { *m = SocialProof{} }
func (m *SocialProof) String() string { return proto.CompactTextString(m) }
func (*SocialProof) ProtoMessage()    {}
func (*SocialProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4844cddd6ff898e, []int{0}
}
func (m *SocialProof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SocialProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SocialProof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SocialProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocialProof.Merge(m, src)
}
func (m *SocialProof) XXX_Size() int {
	return m.Size()
}
func (m *SocialProof) XXX_DiscardUnknown() {
	xxx_messageInfo_SocialProof.DiscardUnknown(m)
}

var xxx_messageInfo_SocialProof proto.InternalMessageInfo

// FanToken defines a standard for the fungible token
type FanToken struct {
	Name         string                                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MaxSupply    github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=max_supply,json=maxSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_supply" yaml:"max_supply"`
	Mintable     bool                                   `protobuf:"varint,3,opt,name=mintable,proto3" json:"mintable,omitempty"`
	Owner        string                                 `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	MetaData     types.Metadata                         `protobuf:"bytes,5,opt,name=meta_data,json=metaData,proto3" json:"meta_data" yaml:"meta_data"`
	SocialProofs []*SocialProof                         `protobuf:"bytes,6,rep,name=social_proofs,json=socialProofs,proto3" json:"social_proofs,omitempty"`
}

func (m *FanToken) Reset()      { *m = FanToken{} }
func (*FanToken) ProtoMessage() {}
func (*FanToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4844cddd6ff898e, []int{1}
}
func (m *FanToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FanToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FanToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FanToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FanToken.Merge(m, src)
}
func (m *FanToken) XXX_Size() int {
	return m.Size()
}
func (m *FanToken) XXX_DiscardUnknown() {
	xxx_messageInfo_FanToken.DiscardUnknown(m)
}

var xxx_messageInfo_FanToken proto.InternalMessageInfo

// Params defines fantoken module's parameters
type Params struct {
	IssuePrice types1.Coin `protobuf:"bytes,2,opt,name=issue_price,json=issuePrice,proto3" json:"issue_price" yaml:"issue_price"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4844cddd6ff898e, []int{2}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SocialProof)(nil), "bitsong.fantoken.SocialProof")
	proto.RegisterType((*FanToken)(nil), "bitsong.fantoken.FanToken")
	proto.RegisterType((*Params)(nil), "bitsong.fantoken.Params")
}

func init() { proto.RegisterFile("fantoken/fantoken.proto", fileDescriptor_b4844cddd6ff898e) }

var fileDescriptor_b4844cddd6ff898e = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0xbd, 0x6e, 0xdb, 0x3c,
	0x14, 0x86, 0xa5, 0xfc, 0xf8, 0x93, 0xe8, 0x7c, 0x40, 0x4a, 0x04, 0xad, 0xea, 0xb6, 0x52, 0xa0,
	0xa1, 0xc8, 0x52, 0x09, 0x49, 0x36, 0x6f, 0x55, 0x8a, 0x02, 0x1d, 0x0a, 0xb8, 0x8a, 0xd3, 0xa1,
	0x8b, 0x40, 0xd9, 0xb4, 0xc2, 0x5a, 0x24, 0x05, 0x91, 0x6e, 0xed, 0x3b, 0xe8, 0x98, 0xb1, 0xa3,
	0xef, 0xa1, 0x37, 0xe1, 0x31, 0x63, 0xd0, 0xc1, 0x6d, 0xed, 0xa5, 0x73, 0xae, 0xa0, 0x10, 0xf5,
	0x63, 0xb7, 0x93, 0xce, 0xc3, 0xc3, 0x43, 0xbe, 0xe7, 0xbc, 0x14, 0x78, 0x34, 0x42, 0x4c, 0xf2,
	0x31, 0x66, 0x7e, 0x1d, 0x78, 0x59, 0xce, 0x25, 0x87, 0x87, 0x31, 0x91, 0x82, 0xb3, 0xc4, 0xab,
	0xd7, 0x3b, 0xf6, 0x80, 0x0b, 0xca, 0x85, 0x1f, 0x23, 0x81, 0xfd, 0x4f, 0xa7, 0x31, 0x96, 0xe8,
	0xd4, 0x1f, 0x70, 0x52, 0x55, 0x6c, 0xe5, 0xd9, 0xb8, 0xc9, 0x17, 0x50, 0xe5, 0x8f, 0x12, 0x9e,
	0x70, 0x15, 0xfa, 0x45, 0x54, 0xad, 0x3a, 0x09, 0xe7, 0x49, 0x8a, 0x7d, 0x45, 0xf1, 0x64, 0xe4,
	0x4b, 0x42, 0xb1, 0x90, 0x88, 0x66, 0xe5, 0x06, 0xf7, 0x9b, 0x0e, 0xda, 0x97, 0x7c, 0x40, 0x50,
	0xda, 0xcb, 0x39, 0x1f, 0xc1, 0x87, 0xa0, 0x25, 0x14, 0x5a, 0xfa, 0xb1, 0x7e, 0x62, 0x86, 0x15,
	0x41, 0x0b, 0xfc, 0x77, 0x8d, 0xd8, 0x30, 0xc5, 0xb9, 0xb5, 0xa3, 0x12, 0x35, 0xc2, 0x27, 0xc0,
	0xcc, 0x8a, 0xd2, 0x68, 0x92, 0x13, 0x6b, 0x57, 0xe5, 0x0c, 0xb5, 0x70, 0x95, 0x13, 0x78, 0x05,
	0x4c, 0x3c, 0xcd, 0x48, 0x8e, 0x23, 0x24, 0xad, 0xbd, 0x63, 0xfd, 0xa4, 0x7d, 0xd6, 0xf1, 0x4a,
	0x4d, 0x5e, 0xad, 0xc9, 0xeb, 0xd7, 0x9a, 0x82, 0xa7, 0x8b, 0xa5, 0xa3, 0xdd, 0x2f, 0x9d, 0xc3,
	0x19, 0xa2, 0x69, 0xd7, 0x6d, 0x4a, 0xdd, 0x9b, 0x1f, 0x8e, 0x1e, 0x1a, 0x25, 0xbf, 0x94, 0xee,
	0xdd, 0x0e, 0x30, 0x5e, 0x23, 0xd6, 0x2f, 0x26, 0x07, 0x21, 0xd8, 0x63, 0x88, 0xe2, 0x4a, 0xb0,
	0x8a, 0x61, 0x0c, 0x00, 0x45, 0xd3, 0x48, 0x4c, 0xb2, 0x2c, 0x9d, 0x29, 0xc5, 0x07, 0xc1, 0x45,
	0x71, 0xf8, 0xf7, 0xa5, 0xf3, 0x3c, 0x21, 0xf2, 0x7a, 0x12, 0x7b, 0x03, 0x4e, 0xfd, 0x6a, 0xa8,
	0xe5, 0xe7, 0x85, 0x18, 0x8e, 0x7d, 0x39, 0xcb, 0xb0, 0xf0, 0xde, 0x30, 0x79, 0xbf, 0x74, 0x1e,
	0x94, 0x32, 0x36, 0x27, 0xb9, 0xa1, 0x49, 0xd1, 0xf4, 0x52, 0xc5, 0xb0, 0x03, 0x0c, 0x4a, 0x98,
	0x44, 0x71, 0x8a, 0x55, 0xdf, 0x46, 0xd8, 0x30, 0x3c, 0x02, 0xfb, 0xfc, 0x33, 0xc3, 0xb9, 0xea,
	0xd9, 0x0c, 0x4b, 0x80, 0x7d, 0x60, 0x52, 0x2c, 0x51, 0x34, 0x44, 0x12, 0x59, 0xfb, 0x6a, 0x1a,
	0xcf, 0xbc, 0xf2, 0x6e, 0x4f, 0x59, 0x59, 0xf9, 0xea, 0xbd, 0xc5, 0x12, 0x15, 0x9b, 0x02, 0xeb,
	0xef, 0x81, 0x34, 0xd5, 0x6e, 0x68, 0x14, 0xf1, 0x2b, 0x24, 0x11, 0x0c, 0xc0, 0xff, 0xa5, 0x49,
	0x91, 0x1a, 0xbb, 0xb0, 0x5a, 0xc7, 0xbb, 0xea, 0xe4, 0x7f, 0xdf, 0x98, 0xb7, 0x65, 0x74, 0x78,
	0x20, 0x36, 0x20, 0xba, 0xc6, 0x97, 0xb9, 0xa3, 0x7d, 0x9d, 0x3b, 0x9a, 0xfb, 0x11, 0xb4, 0x7a,
	0x28, 0x47, 0x54, 0xc0, 0xf7, 0xa0, 0x4d, 0x84, 0x98, 0xe0, 0x28, 0xcb, 0xc9, 0x00, 0xab, 0x21,
	0xb6, 0xcf, 0x1e, 0x6f, 0xf4, 0x0a, 0xdc, 0xe8, 0xbd, 0xe0, 0x84, 0x05, 0x9d, 0x4a, 0x2b, 0x2c,
	0xb5, 0x6e, 0xd5, 0xba, 0x21, 0x50, 0xd4, 0x2b, 0xa0, 0x6b, 0x14, 0xf7, 0xfc, 0x9e, 0x3b, 0x7a,
	0xf0, 0x6e, 0xf1, 0xcb, 0xd6, 0x16, 0x2b, 0x5b, 0xbf, 0x5d, 0xd9, 0xfa, 0xcf, 0x95, 0xad, 0xdf,
	0xac, 0x6d, 0xed, 0x76, 0x6d, 0x6b, 0x77, 0x6b, 0x5b, 0xfb, 0x70, 0xbe, 0xe5, 0x53, 0xd5, 0x0a,
	0x1f, 0x8d, 0x48, 0xa1, 0xba, 0x66, 0x7f, 0xda, 0xfc, 0x58, 0xa5, 0x71, 0x71, 0x4b, 0xbd, 0xaa,
	0xf3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x45, 0xd1, 0x65, 0x78, 0x7a, 0x03, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if !this.IssuePrice.Equal(&that1.IssuePrice) {
		return false
	}
	return true
}
func (m *SocialProof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SocialProof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SocialProof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.ExpireAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpireAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintFantoken(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if len(m.ProofUri) > 0 {
		i -= len(m.ProofUri)
		copy(dAtA[i:], m.ProofUri)
		i = encodeVarintFantoken(dAtA, i, uint64(len(m.ProofUri)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Handler) > 0 {
		i -= len(m.Handler)
		copy(dAtA[i:], m.Handler)
		i = encodeVarintFantoken(dAtA, i, uint64(len(m.Handler)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Social) > 0 {
		i -= len(m.Social)
		copy(dAtA[i:], m.Social)
		i = encodeVarintFantoken(dAtA, i, uint64(len(m.Social)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FanToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FanToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FanToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SocialProofs) > 0 {
		for iNdEx := len(m.SocialProofs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SocialProofs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFantoken(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	{
		size, err := m.MetaData.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintFantoken(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintFantoken(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x22
	}
	if m.Mintable {
		i--
		if m.Mintable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.MaxSupply.Size()
		i -= size
		if _, err := m.MaxSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintFantoken(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintFantoken(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.IssuePrice.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintFantoken(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}

func encodeVarintFantoken(dAtA []byte, offset int, v uint64) int {
	offset -= sovFantoken(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SocialProof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Social)
	if l > 0 {
		n += 1 + l + sovFantoken(uint64(l))
	}
	l = len(m.Handler)
	if l > 0 {
		n += 1 + l + sovFantoken(uint64(l))
	}
	l = len(m.ProofUri)
	if l > 0 {
		n += 1 + l + sovFantoken(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpireAt)
	n += 1 + l + sovFantoken(uint64(l))
	return n
}

func (m *FanToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovFantoken(uint64(l))
	}
	l = m.MaxSupply.Size()
	n += 1 + l + sovFantoken(uint64(l))
	if m.Mintable {
		n += 2
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovFantoken(uint64(l))
	}
	l = m.MetaData.Size()
	n += 1 + l + sovFantoken(uint64(l))
	if len(m.SocialProofs) > 0 {
		for _, e := range m.SocialProofs {
			l = e.Size()
			n += 1 + l + sovFantoken(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.IssuePrice.Size()
	n += 1 + l + sovFantoken(uint64(l))
	return n
}

func sovFantoken(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFantoken(x uint64) (n int) {
	return sovFantoken(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SocialProof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFantoken
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
			return fmt.Errorf("proto: SocialProof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SocialProof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Social", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFantoken
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
				return ErrInvalidLengthFantoken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFantoken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Social = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Handler", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFantoken
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
				return ErrInvalidLengthFantoken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFantoken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Handler = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofUri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFantoken
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
				return ErrInvalidLengthFantoken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFantoken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProofUri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpireAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFantoken
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
				return ErrInvalidLengthFantoken
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFantoken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.ExpireAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFantoken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFantoken
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFantoken
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
func (m *FanToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFantoken
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
			return fmt.Errorf("proto: FanToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FanToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFantoken
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
				return ErrInvalidLengthFantoken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFantoken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSupply", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFantoken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFantoken
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFantoken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mintable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFantoken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Mintable = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFantoken
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
				return ErrInvalidLengthFantoken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFantoken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFantoken
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
				return ErrInvalidLengthFantoken
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFantoken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MetaData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SocialProofs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFantoken
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
				return ErrInvalidLengthFantoken
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFantoken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SocialProofs = append(m.SocialProofs, &SocialProof{})
			if err := m.SocialProofs[len(m.SocialProofs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFantoken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFantoken
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFantoken
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFantoken
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssuePrice", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFantoken
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
				return ErrInvalidLengthFantoken
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFantoken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.IssuePrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFantoken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFantoken
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFantoken
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
func skipFantoken(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFantoken
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
					return 0, ErrIntOverflowFantoken
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFantoken
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
				return 0, ErrInvalidLengthFantoken
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFantoken
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFantoken
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFantoken        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFantoken          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFantoken = fmt.Errorf("proto: unexpected end of group")
)