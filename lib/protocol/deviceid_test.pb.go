// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: deviceid_test.proto

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	deviceid_test.proto

It has these top-level messages:
	TestOldDeviceID
	TestNewDeviceID
*/
package protocol

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TestOldDeviceID struct {
	Test []byte `protobuf:"bytes,1,opt,name=test,proto3" json:"test,omitempty"`
}

func (m *TestOldDeviceID) Reset()                    { *m = TestOldDeviceID{} }
func (m *TestOldDeviceID) String() string            { return proto.CompactTextString(m) }
func (*TestOldDeviceID) ProtoMessage()               {}
func (*TestOldDeviceID) Descriptor() ([]byte, []int) { return fileDescriptorDeviceidTest, []int{0} }

type TestNewDeviceID struct {
	Test DeviceID `protobuf:"bytes,1,opt,name=test,proto3,customtype=DeviceID" json:"test"`
}

func (m *TestNewDeviceID) Reset()                    { *m = TestNewDeviceID{} }
func (m *TestNewDeviceID) String() string            { return proto.CompactTextString(m) }
func (*TestNewDeviceID) ProtoMessage()               {}
func (*TestNewDeviceID) Descriptor() ([]byte, []int) { return fileDescriptorDeviceidTest, []int{1} }

func init() {
	proto.RegisterType((*TestOldDeviceID)(nil), "protocol.TestOldDeviceID")
	proto.RegisterType((*TestNewDeviceID)(nil), "protocol.TestNewDeviceID")
}
func (m *TestOldDeviceID) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestOldDeviceID) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Test) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDeviceidTest(dAtA, i, uint64(len(m.Test)))
		i += copy(dAtA[i:], m.Test)
	}
	return i, nil
}

func (m *TestNewDeviceID) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestNewDeviceID) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintDeviceidTest(dAtA, i, uint64(m.Test.ProtoSize()))
	n1, err := m.Test.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func encodeFixed64DeviceidTest(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32DeviceidTest(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintDeviceidTest(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TestOldDeviceID) ProtoSize() (n int) {
	var l int
	_ = l
	l = len(m.Test)
	if l > 0 {
		n += 1 + l + sovDeviceidTest(uint64(l))
	}
	return n
}

func (m *TestNewDeviceID) ProtoSize() (n int) {
	var l int
	_ = l
	l = m.Test.ProtoSize()
	n += 1 + l + sovDeviceidTest(uint64(l))
	return n
}

func sovDeviceidTest(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDeviceidTest(x uint64) (n int) {
	return sovDeviceidTest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TestOldDeviceID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceidTest
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TestOldDeviceID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestOldDeviceID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Test", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceidTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDeviceidTest
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Test = append(m.Test[:0], dAtA[iNdEx:postIndex]...)
			if m.Test == nil {
				m.Test = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceidTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceidTest
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
func (m *TestNewDeviceID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceidTest
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TestNewDeviceID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestNewDeviceID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Test", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceidTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDeviceidTest
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Test.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceidTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceidTest
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
func skipDeviceidTest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeviceidTest
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
					return 0, ErrIntOverflowDeviceidTest
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDeviceidTest
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthDeviceidTest
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDeviceidTest
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipDeviceidTest(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthDeviceidTest = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeviceidTest   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("deviceid_test.proto", fileDescriptorDeviceidTest) }

var fileDescriptorDeviceidTest = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x49, 0x2d, 0xcb,
	0x4c, 0x4e, 0xcd, 0x4c, 0x89, 0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39, 0x52, 0xca, 0x45, 0xa9, 0x05, 0xf9, 0xc5, 0xfa, 0x60, 0x7e,
	0x52, 0x69, 0x9a, 0x7e, 0x7a, 0x7e, 0x7a, 0x3e, 0x98, 0x03, 0x66, 0x41, 0x94, 0x2b, 0xa9, 0x72,
	0xf1, 0x87, 0xa4, 0x16, 0x97, 0xf8, 0xe7, 0xa4, 0xb8, 0x80, 0x0d, 0xf3, 0x74, 0x11, 0x12, 0xe2,
	0x62, 0x01, 0x99, 0x27, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x13, 0x04, 0x66, 0x2b, 0x99, 0x43, 0x94,
	0xf9, 0xa5, 0x96, 0xc3, 0x95, 0xa9, 0x20, 0x2b, 0x73, 0x12, 0x38, 0x71, 0x4f, 0x9e, 0xe1, 0xd6,
	0x3d, 0x79, 0x0e, 0x98, 0x3c, 0x44, 0xa3, 0x93, 0xcc, 0x89, 0x87, 0x72, 0x0c, 0x17, 0x1e, 0xca,
	0x31, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x0c, 0x2f, 0x1e,
	0xc9, 0x31, 0x2c, 0x78, 0x2c, 0xc7, 0x98, 0xc4, 0x06, 0x76, 0x84, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x82, 0x8f, 0xec, 0x3e, 0xca, 0x00, 0x00, 0x00,
}
