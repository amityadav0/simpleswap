// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: simpleswap/simpleswap/pool_params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type PoolParams struct {
	// swapFee is ranged from 0 to 10000.
	SwapFee github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=swapFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"swapFee"`
	// Amplifier parameters for stable pool.
	Amp *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=amp,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amp,omitempty"`
}

func (m *PoolParams) Reset()         { *m = PoolParams{} }
func (m *PoolParams) String() string { return proto.CompactTextString(m) }
func (*PoolParams) ProtoMessage()    {}
func (*PoolParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_a81ab124ed3b94ba, []int{0}
}
func (m *PoolParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolParams.Merge(m, src)
}
func (m *PoolParams) XXX_Size() int {
	return m.Size()
}
func (m *PoolParams) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolParams.DiscardUnknown(m)
}

var xxx_messageInfo_PoolParams proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PoolParams)(nil), "simpleswap.simpleswap.PoolParams")
}

func init() {
	proto.RegisterFile("simpleswap/simpleswap/pool_params.proto", fileDescriptor_a81ab124ed3b94ba)
}

var fileDescriptor_a81ab124ed3b94ba = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0xce, 0xcc, 0x2d,
	0xc8, 0x49, 0x2d, 0x2e, 0x4f, 0x2c, 0xd0, 0x47, 0x62, 0x16, 0xe4, 0xe7, 0xe7, 0xc4, 0x17, 0x24,
	0x16, 0x25, 0xe6, 0x16, 0xeb, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x89, 0x22, 0x64, 0xf5, 0x10,
	0x4c, 0x29, 0xb9, 0xe4, 0xfc, 0xe2, 0xdc, 0xfc, 0x62, 0xfd, 0xa4, 0xc4, 0xe2, 0x54, 0xfd, 0x32,
	0xc3, 0xa4, 0xd4, 0x92, 0x44, 0x43, 0xfd, 0xe4, 0xfc, 0xcc, 0x3c, 0x88, 0x36, 0x29, 0x91, 0xf4,
	0xfc, 0xf4, 0x7c, 0x30, 0x53, 0x1f, 0xc4, 0x82, 0x88, 0x2a, 0x4d, 0x61, 0xe4, 0xe2, 0x0a, 0xc8,
	0xcf, 0xcf, 0x09, 0x00, 0xdb, 0x20, 0xe4, 0xc1, 0xc5, 0x0e, 0x32, 0xcc, 0x2d, 0x35, 0x55, 0x82,
	0x51, 0x81, 0x51, 0x83, 0xd3, 0x49, 0xef, 0xc4, 0x3d, 0x79, 0x86, 0x5b, 0xf7, 0xe4, 0xd5, 0xd2,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xa1, 0x16, 0x41, 0x28, 0xdd, 0xe2,
	0x94, 0x6c, 0xfd, 0x92, 0xca, 0x82, 0xd4, 0x62, 0x3d, 0x97, 0xd4, 0xe4, 0x20, 0x98, 0x76, 0x21,
	0x1b, 0x2e, 0xe6, 0xc4, 0xdc, 0x02, 0x09, 0x26, 0xb0, 0x29, 0x5a, 0x44, 0x9a, 0xe0, 0x99, 0x57,
	0x12, 0x04, 0xd2, 0xe6, 0x64, 0x7e, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e,
	0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51,
	0xb2, 0x48, 0x61, 0x53, 0x81, 0x1c, 0x50, 0x60, 0x13, 0x92, 0xd8, 0xc0, 0xde, 0x32, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x89, 0x83, 0xb1, 0x09, 0x4e, 0x01, 0x00, 0x00,
}

func (m *PoolParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amp != nil {
		{
			size := m.Amp.Size()
			i -= size
			if _, err := m.Amp.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintPoolParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.SwapFee.Size()
		i -= size
		if _, err := m.SwapFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPoolParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintPoolParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovPoolParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PoolParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.SwapFee.Size()
	n += 1 + l + sovPoolParams(uint64(l))
	if m.Amp != nil {
		l = m.Amp.Size()
		n += 1 + l + sovPoolParams(uint64(l))
	}
	return n
}

func sovPoolParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPoolParams(x uint64) (n int) {
	return sovPoolParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PoolParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPoolParams
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
			return fmt.Errorf("proto: PoolParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolParams
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
				return ErrInvalidLengthPoolParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoolParams
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
				return ErrInvalidLengthPoolParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoolParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.Amp = &v
			if err := m.Amp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPoolParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPoolParams
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
func skipPoolParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPoolParams
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
					return 0, ErrIntOverflowPoolParams
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
					return 0, ErrIntOverflowPoolParams
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
				return 0, ErrInvalidLengthPoolParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPoolParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPoolParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPoolParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPoolParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPoolParams = fmt.Errorf("proto: unexpected end of group")
)
