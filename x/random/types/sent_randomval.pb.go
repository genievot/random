// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: random/sent_randomval.proto

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

type SentRandomval struct {
	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Userid  string `protobuf:"bytes,2,opt,name=userid,proto3" json:"userid,omitempty"`
	Vrv     string `protobuf:"bytes,3,opt,name=vrv,proto3" json:"vrv,omitempty"`
	Creator string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *SentRandomval) Reset()         { *m = SentRandomval{} }
func (m *SentRandomval) String() string { return proto.CompactTextString(m) }
func (*SentRandomval) ProtoMessage()    {}
func (*SentRandomval) Descriptor() ([]byte, []int) {
	return fileDescriptor_289185b0e5852e56, []int{0}
}
func (m *SentRandomval) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SentRandomval) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SentRandomval.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SentRandomval) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SentRandomval.Merge(m, src)
}
func (m *SentRandomval) XXX_Size() int {
	return m.Size()
}
func (m *SentRandomval) XXX_DiscardUnknown() {
	xxx_messageInfo_SentRandomval.DiscardUnknown(m)
}

var xxx_messageInfo_SentRandomval proto.InternalMessageInfo

func (m *SentRandomval) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SentRandomval) GetUserid() string {
	if m != nil {
		return m.Userid
	}
	return ""
}

func (m *SentRandomval) GetVrv() string {
	if m != nil {
		return m.Vrv
	}
	return ""
}

func (m *SentRandomval) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*SentRandomval)(nil), "genievot.random.random.SentRandomval")
}

func init() { proto.RegisterFile("random/sent_randomval.proto", fileDescriptor_289185b0e5852e56) }

var fileDescriptor_289185b0e5852e56 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x4a, 0xcc, 0x4b,
	0xc9, 0xcf, 0xd5, 0x2f, 0x4e, 0xcd, 0x2b, 0x89, 0x87, 0xb0, 0xcb, 0x12, 0x73, 0xf4, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0xc4, 0xd2, 0x53, 0xf3, 0x32, 0x53, 0xcb, 0xf2, 0x4b, 0xf4, 0x20, 0x32,
	0x50, 0x4a, 0x4a, 0x24, 0x3d, 0x3f, 0x3d, 0x1f, 0xac, 0x44, 0x1f, 0xc4, 0x82, 0xa8, 0x56, 0x4a,
	0xe6, 0xe2, 0x0d, 0x4e, 0xcd, 0x2b, 0x09, 0x82, 0x19, 0x22, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x22,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x12, 0xc4, 0x94, 0x99, 0x22, 0x24, 0xc6, 0xc5, 0x56, 0x5a, 0x9c,
	0x5a, 0x94, 0x99, 0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe5, 0x09, 0x09, 0x70, 0x31,
	0x97, 0x15, 0x95, 0x49, 0x30, 0x83, 0x05, 0x41, 0x4c, 0x21, 0x09, 0x2e, 0xf6, 0xe4, 0xa2, 0xd4,
	0xc4, 0x92, 0xfc, 0x22, 0x09, 0x16, 0xb0, 0x28, 0x8c, 0xeb, 0xe4, 0x7c, 0xe2, 0x91, 0x1c, 0xe3,
	0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c,
	0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x9a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9,
	0xb9, 0xfa, 0x30, 0x77, 0xeb, 0x43, 0x7d, 0x57, 0x01, 0x63, 0x94, 0x54, 0x16, 0xa4, 0x16, 0x27,
	0xb1, 0x81, 0x1d, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x5c, 0x4a, 0x7c, 0xfd, 0x00,
	0x00, 0x00,
}

func (m *SentRandomval) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SentRandomval) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SentRandomval) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSentRandomval(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Vrv) > 0 {
		i -= len(m.Vrv)
		copy(dAtA[i:], m.Vrv)
		i = encodeVarintSentRandomval(dAtA, i, uint64(len(m.Vrv)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Userid) > 0 {
		i -= len(m.Userid)
		copy(dAtA[i:], m.Userid)
		i = encodeVarintSentRandomval(dAtA, i, uint64(len(m.Userid)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintSentRandomval(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSentRandomval(dAtA []byte, offset int, v uint64) int {
	offset -= sovSentRandomval(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SentRandomval) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovSentRandomval(uint64(m.Id))
	}
	l = len(m.Userid)
	if l > 0 {
		n += 1 + l + sovSentRandomval(uint64(l))
	}
	l = len(m.Vrv)
	if l > 0 {
		n += 1 + l + sovSentRandomval(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSentRandomval(uint64(l))
	}
	return n
}

func sovSentRandomval(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSentRandomval(x uint64) (n int) {
	return sovSentRandomval(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SentRandomval) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSentRandomval
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
			return fmt.Errorf("proto: SentRandomval: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SentRandomval: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSentRandomval
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Userid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSentRandomval
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
				return ErrInvalidLengthSentRandomval
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSentRandomval
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Userid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vrv", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSentRandomval
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
				return ErrInvalidLengthSentRandomval
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSentRandomval
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vrv = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSentRandomval
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
				return ErrInvalidLengthSentRandomval
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSentRandomval
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSentRandomval(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSentRandomval
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
func skipSentRandomval(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSentRandomval
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
					return 0, ErrIntOverflowSentRandomval
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
					return 0, ErrIntOverflowSentRandomval
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
				return 0, ErrInvalidLengthSentRandomval
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSentRandomval
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSentRandomval
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSentRandomval        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSentRandomval          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSentRandomval = fmt.Errorf("proto: unexpected end of group")
)
