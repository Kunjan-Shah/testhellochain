// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: estimator/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the estimator module's genesis state.
type GenesisState struct {
	Params          Params        `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	ApiHitsList     []ApiHits     `protobuf:"bytes,2,rep,name=apiHitsList,proto3" json:"apiHitsList"`
	ApiHitsCount    uint64        `protobuf:"varint,3,opt,name=apiHitsCount,proto3" json:"apiHitsCount,omitempty"`
	ApiDataList     []ApiData     `protobuf:"bytes,4,rep,name=apiDataList,proto3" json:"apiDataList"`
	ApiDataCount    uint64        `protobuf:"varint,5,opt,name=apiDataCount,proto3" json:"apiDataCount,omitempty"`
	ApiCountMapList []ApiCountMap `protobuf:"bytes,6,rep,name=apiCountMapList,proto3" json:"apiCountMapList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c50fcec85101d329, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetApiHitsList() []ApiHits {
	if m != nil {
		return m.ApiHitsList
	}
	return nil
}

func (m *GenesisState) GetApiHitsCount() uint64 {
	if m != nil {
		return m.ApiHitsCount
	}
	return 0
}

func (m *GenesisState) GetApiDataList() []ApiData {
	if m != nil {
		return m.ApiDataList
	}
	return nil
}

func (m *GenesisState) GetApiDataCount() uint64 {
	if m != nil {
		return m.ApiDataCount
	}
	return 0
}

func (m *GenesisState) GetApiCountMapList() []ApiCountMap {
	if m != nil {
		return m.ApiCountMapList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "testhellochain.estimator.GenesisState")
}

func init() { proto.RegisterFile("estimator/genesis.proto", fileDescriptor_c50fcec85101d329) }

var fileDescriptor_c50fcec85101d329 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0xc7, 0x9b, 0xdf, 0xf6, 0xdb, 0x21, 0x1b, 0x08, 0x45, 0xb4, 0x0c, 0x8c, 0x75, 0x20, 0xec,
	0xd4, 0xc2, 0xbc, 0x79, 0x10, 0x9c, 0x82, 0x0a, 0x0a, 0x32, 0xf1, 0xe2, 0x65, 0x3c, 0xce, 0xb0,
	0x05, 0xb6, 0x25, 0x2c, 0x8f, 0xa0, 0x67, 0xdf, 0x80, 0x2f, 0x6b, 0xc7, 0x1d, 0x3d, 0x89, 0xb4,
	0x6f, 0x44, 0x9a, 0x64, 0xed, 0x2a, 0x54, 0xbc, 0x25, 0xcf, 0xf7, 0xcf, 0xe7, 0x81, 0x87, 0xee,
	0x72, 0x8d, 0x62, 0x06, 0x28, 0x17, 0xf1, 0x98, 0xcf, 0xb9, 0x16, 0x3a, 0x52, 0x0b, 0x89, 0xd2,
	0x0f, 0x90, 0x6b, 0x9c, 0xf0, 0xe9, 0x54, 0x8e, 0x26, 0x20, 0xe6, 0x51, 0xee, 0x6b, 0x6f, 0x8f,
	0xe5, 0x58, 0x1a, 0x53, 0x9c, 0xbd, 0xac, 0xbf, 0xbd, 0x53, 0x14, 0x29, 0x58, 0xc0, 0xcc, 0xf5,
	0xb4, 0x83, 0x62, 0x0e, 0x4a, 0x0c, 0x27, 0x02, 0x2b, 0x94, 0x27, 0x40, 0x70, 0xca, 0x5e, 0x59,
	0x19, 0xc9, 0xe7, 0x39, 0x0e, 0x67, 0xa0, 0xac, 0xdc, 0x79, 0xab, 0xd1, 0xd6, 0x85, 0x5d, 0xf6,
	0x0e, 0x01, 0xb9, 0x7f, 0x42, 0x1b, 0x96, 0x19, 0x90, 0x90, 0x74, 0x9b, 0xbd, 0x30, 0xaa, 0x5a,
	0x3e, 0xba, 0x35, 0xbe, 0x7e, 0x7d, 0xf9, 0xb9, 0xef, 0x0d, 0x5c, 0xca, 0xbf, 0xa2, 0x4d, 0x50,
	0xe2, 0x52, 0xa0, 0xbe, 0x16, 0x1a, 0x83, 0x7f, 0x61, 0xad, 0xdb, 0xec, 0x1d, 0x54, 0x97, 0x9c,
	0x5a, 0xb3, 0x6b, 0xd9, 0xcc, 0xfa, 0x1d, 0xda, 0x72, 0xdf, 0xb3, 0x6c, 0xeb, 0xa0, 0x16, 0x92,
	0x6e, 0x7d, 0x50, 0x9a, 0x39, 0xdc, 0x39, 0x20, 0x18, 0x5c, 0xfd, 0x0f, 0xb8, 0xcc, 0xbc, 0x81,
	0x5b, 0x67, 0x1d, 0x2e, 0xfb, 0x5a, 0xdc, 0xff, 0x1c, 0x97, 0xcf, 0xfc, 0x7b, 0xba, 0x05, 0x4a,
	0x98, 0xf7, 0x0d, 0x28, 0x83, 0x6c, 0x18, 0xe4, 0xe1, 0xaf, 0xc8, 0x75, 0xc0, 0x61, 0x7f, 0x76,
	0xf4, 0x8f, 0x97, 0x09, 0x23, 0xab, 0x84, 0x91, 0xaf, 0x84, 0x91, 0xf7, 0x94, 0x79, 0xab, 0x94,
	0x79, 0x1f, 0x29, 0xf3, 0x1e, 0xc2, 0x72, 0x6d, 0xfc, 0x12, 0x17, 0xf7, 0xc4, 0x57, 0xc5, 0xf5,
	0x63, 0xc3, 0x1c, 0xf2, 0xe8, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xf1, 0xb9, 0x55, 0x7e, 0x02,
	0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ApiCountMapList) > 0 {
		for iNdEx := len(m.ApiCountMapList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ApiCountMapList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.ApiDataCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ApiDataCount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.ApiDataList) > 0 {
		for iNdEx := len(m.ApiDataList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ApiDataList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.ApiHitsCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ApiHitsCount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ApiHitsList) > 0 {
		for iNdEx := len(m.ApiHitsList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ApiHitsList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.ApiHitsList) > 0 {
		for _, e := range m.ApiHitsList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.ApiHitsCount != 0 {
		n += 1 + sovGenesis(uint64(m.ApiHitsCount))
	}
	if len(m.ApiDataList) > 0 {
		for _, e := range m.ApiDataList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.ApiDataCount != 0 {
		n += 1 + sovGenesis(uint64(m.ApiDataCount))
	}
	if len(m.ApiCountMapList) > 0 {
		for _, e := range m.ApiCountMapList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiHitsList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiHitsList = append(m.ApiHitsList, ApiHits{})
			if err := m.ApiHitsList[len(m.ApiHitsList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiHitsCount", wireType)
			}
			m.ApiHitsCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ApiHitsCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiDataList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiDataList = append(m.ApiDataList, ApiData{})
			if err := m.ApiDataList[len(m.ApiDataList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiDataCount", wireType)
			}
			m.ApiDataCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ApiDataCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiCountMapList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiCountMapList = append(m.ApiCountMapList, ApiCountMap{})
			if err := m.ApiCountMapList[len(m.ApiCountMapList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)