// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/monitor/monitor.proto

package monitor // import "github.com/tsavola/gate/server/monitor"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import event "github.com/tsavola/gate/server/event"

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

type State struct {
	ProgramsLoaded       uint32   `protobuf:"varint,1,opt,name=programs_loaded,json=programsLoaded,proto3" json:"programs_loaded,omitempty"`
	ProgramLinks         uint32   `protobuf:"varint,2,opt,name=program_links,json=programLinks,proto3" json:"program_links,omitempty"`
	Instances            uint32   `protobuf:"varint,3,opt,name=instances,proto3" json:"instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_611652bc795644a5, []int{0}
}
func (m *State) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_State.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(dst, src)
}
func (m *State) XXX_Size() int {
	return m.Size()
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

type Event struct {
	Type                 event.Event_Type `protobuf:"varint,1,opt,name=type,proto3,enum=event.Event_Type" json:"type,omitempty"`
	Event                []byte           `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	Error                string           `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_611652bc795644a5, []int{1}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Event.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return m.Size()
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func init() {
	proto.RegisterType((*State)(nil), "monitor.State")
	proto.RegisterType((*Event)(nil), "monitor.Event")
}
func (m *State) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *State) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ProgramsLoaded != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMonitor(dAtA, i, uint64(m.ProgramsLoaded))
	}
	if m.ProgramLinks != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMonitor(dAtA, i, uint64(m.ProgramLinks))
	}
	if m.Instances != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMonitor(dAtA, i, uint64(m.Instances))
	}
	return i, nil
}

func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMonitor(dAtA, i, uint64(m.Type))
	}
	if len(m.Event) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMonitor(dAtA, i, uint64(len(m.Event)))
		i += copy(dAtA[i:], m.Event)
	}
	if len(m.Error) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMonitor(dAtA, i, uint64(len(m.Error)))
		i += copy(dAtA[i:], m.Error)
	}
	return i, nil
}

func encodeVarintMonitor(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *State) Size() (n int) {
	var l int
	_ = l
	if m.ProgramsLoaded != 0 {
		n += 1 + sovMonitor(uint64(m.ProgramsLoaded))
	}
	if m.ProgramLinks != 0 {
		n += 1 + sovMonitor(uint64(m.ProgramLinks))
	}
	if m.Instances != 0 {
		n += 1 + sovMonitor(uint64(m.Instances))
	}
	return n
}

func (m *Event) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovMonitor(uint64(m.Type))
	}
	l = len(m.Event)
	if l > 0 {
		n += 1 + l + sovMonitor(uint64(l))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovMonitor(uint64(l))
	}
	return n
}

func sovMonitor(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMonitor(x uint64) (n int) {
	return sovMonitor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *State) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMonitor
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
			return fmt.Errorf("proto: State: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: State: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProgramsLoaded", wireType)
			}
			m.ProgramsLoaded = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProgramsLoaded |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProgramLinks", wireType)
			}
			m.ProgramLinks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProgramLinks |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Instances", wireType)
			}
			m.Instances = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Instances |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMonitor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMonitor
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
func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMonitor
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
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (event.Event_Type(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitor
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
				return ErrInvalidLengthMonitor
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Event = append(m.Event[:0], dAtA[iNdEx:postIndex]...)
			if m.Event == nil {
				m.Event = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMonitor
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMonitor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMonitor
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
func skipMonitor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMonitor
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
					return 0, ErrIntOverflowMonitor
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
					return 0, ErrIntOverflowMonitor
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
				return 0, ErrInvalidLengthMonitor
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMonitor
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
				next, err := skipMonitor(dAtA[start:])
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
	ErrInvalidLengthMonitor = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMonitor   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("server/monitor/monitor.proto", fileDescriptor_monitor_611652bc795644a5)
}

var fileDescriptor_monitor_611652bc795644a5 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0xcf, 0xcd, 0xcf, 0xcb, 0x2c, 0xc9, 0x87, 0xd3, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0xec, 0x50, 0xae, 0x94, 0x04, 0x54, 0x59, 0x6a, 0x59, 0x6a, 0x5e, 0x09, 0x84, 0x84,
	0x28, 0x51, 0x2a, 0xe6, 0x62, 0x0d, 0x2e, 0x49, 0x2c, 0x49, 0x15, 0x52, 0xe7, 0xe2, 0x2f, 0x28,
	0xca, 0x4f, 0x2f, 0x4a, 0xcc, 0x2d, 0x8e, 0xcf, 0xc9, 0x4f, 0x4c, 0x49, 0x4d, 0x91, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0d, 0xe2, 0x83, 0x09, 0xfb, 0x80, 0x45, 0x85, 0x94, 0xb9, 0x78, 0xa1, 0x22,
	0xf1, 0x39, 0x99, 0x79, 0xd9, 0xc5, 0x12, 0x4c, 0x60, 0x65, 0x3c, 0x50, 0x41, 0x1f, 0x90, 0x98,
	0x90, 0x0c, 0x17, 0x67, 0x66, 0x5e, 0x71, 0x49, 0x62, 0x5e, 0x72, 0x6a, 0xb1, 0x04, 0x33, 0x58,
	0x01, 0x42, 0x40, 0x29, 0x8a, 0x8b, 0xd5, 0x15, 0xe4, 0x06, 0x21, 0x55, 0x2e, 0x96, 0x92, 0xca,
	0x82, 0x54, 0xb0, 0x4d, 0x7c, 0x46, 0x82, 0x7a, 0x10, 0x97, 0x81, 0xe5, 0xf4, 0x42, 0x2a, 0x0b,
	0x52, 0x83, 0xc0, 0xd2, 0x42, 0x22, 0x5c, 0xac, 0x60, 0x19, 0xb0, 0x55, 0x3c, 0x41, 0x10, 0x0e,
	0x58, 0xb4, 0xa8, 0x28, 0xbf, 0x08, 0x6c, 0x3e, 0x67, 0x10, 0x84, 0xe3, 0x64, 0x73, 0xe2, 0xa1,
	0x1c, 0xc3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1,
	0xb1, 0x1c, 0x43, 0x94, 0x5a, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e,
	0x49, 0x71, 0x62, 0x59, 0x7e, 0x4e, 0xa2, 0x7e, 0x7a, 0x62, 0x49, 0xaa, 0x3e, 0x6a, 0xf8, 0x25,
	0xb1, 0x81, 0x43, 0xc5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x25, 0xb8, 0xde, 0x58, 0x01,
	0x00, 0x00,
}
