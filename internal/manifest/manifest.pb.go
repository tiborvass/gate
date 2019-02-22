// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internal/manifest/manifest.proto

package manifest

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

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

type ByteRange struct {
	Offset               int64    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Length               int64    `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ByteRange) Reset()         { *m = ByteRange{} }
func (m *ByteRange) String() string { return proto.CompactTextString(m) }
func (*ByteRange) ProtoMessage()    {}
func (*ByteRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_manifest_0e5e7cdd968e7cdd, []int{0}
}
func (m *ByteRange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ByteRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ByteRange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ByteRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ByteRange.Merge(dst, src)
}
func (m *ByteRange) XXX_Size() int {
	return m.Size()
}
func (m *ByteRange) XXX_DiscardUnknown() {
	xxx_messageInfo_ByteRange.DiscardUnknown(m)
}

var xxx_messageInfo_ByteRange proto.InternalMessageInfo

type Archive struct {
	ModuleSize           int64             `protobuf:"varint,1,opt,name=module_size,json=moduleSize,proto3" json:"module_size,omitempty"`
	Sections             []ByteRange       `protobuf:"bytes,2,rep,name=sections" json:"sections"`
	StackSection         ByteRange         `protobuf:"bytes,3,opt,name=stack_section,json=stackSection" json:"stack_section"`
	GlobalTypes          []byte            `protobuf:"bytes,4,opt,name=global_types,json=globalTypes,proto3" json:"global_types,omitempty"`
	EntryIndexes         map[string]uint32 `protobuf:"bytes,5,rep,name=entry_indexes,json=entryIndexes" json:"entry_indexes" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	EntryAddrs           map[uint32]uint32 `protobuf:"bytes,6,rep,name=entry_addrs,json=entryAddrs" json:"entry_addrs" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	CallSitesSize        uint32            `protobuf:"varint,7,opt,name=call_sites_size,json=callSitesSize,proto3" json:"call_sites_size,omitempty"`
	FuncAddrsSize        uint32            `protobuf:"varint,8,opt,name=func_addrs_size,json=funcAddrsSize,proto3" json:"func_addrs_size,omitempty"`
	Exe                  Executable        `protobuf:"bytes,9,opt,name=exe" json:"exe"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Archive) Reset()         { *m = Archive{} }
func (m *Archive) String() string { return proto.CompactTextString(m) }
func (*Archive) ProtoMessage()    {}
func (*Archive) Descriptor() ([]byte, []int) {
	return fileDescriptor_manifest_0e5e7cdd968e7cdd, []int{1}
}
func (m *Archive) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Archive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Archive.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Archive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Archive.Merge(dst, src)
}
func (m *Archive) XXX_Size() int {
	return m.Size()
}
func (m *Archive) XXX_DiscardUnknown() {
	xxx_messageInfo_Archive.DiscardUnknown(m)
}

var xxx_messageInfo_Archive proto.InternalMessageInfo

type Executable struct {
	TextAddr             uint64   `protobuf:"varint,1,opt,name=text_addr,json=textAddr,proto3" json:"text_addr,omitempty"`
	TextSize             uint32   `protobuf:"varint,2,opt,name=text_size,json=textSize,proto3" json:"text_size,omitempty"`
	StackSize            uint32   `protobuf:"varint,3,opt,name=stack_size,json=stackSize,proto3" json:"stack_size,omitempty"`
	StackUsage           uint32   `protobuf:"varint,4,opt,name=stack_usage,json=stackUsage,proto3" json:"stack_usage,omitempty"`
	GlobalsSize          uint32   `protobuf:"varint,5,opt,name=globals_size,json=globalsSize,proto3" json:"globals_size,omitempty"`
	MemoryDataSize       uint32   `protobuf:"varint,6,opt,name=memory_data_size,json=memoryDataSize,proto3" json:"memory_data_size,omitempty"`
	MemorySize           uint32   `protobuf:"varint,7,opt,name=memory_size,json=memorySize,proto3" json:"memory_size,omitempty"`
	MemorySizeLimit      uint32   `protobuf:"varint,8,opt,name=memory_size_limit,json=memorySizeLimit,proto3" json:"memory_size_limit,omitempty"`
	InitRoutine          int32    `protobuf:"varint,9,opt,name=init_routine,json=initRoutine,proto3" json:"init_routine,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Executable) Reset()         { *m = Executable{} }
func (m *Executable) String() string { return proto.CompactTextString(m) }
func (*Executable) ProtoMessage()    {}
func (*Executable) Descriptor() ([]byte, []int) {
	return fileDescriptor_manifest_0e5e7cdd968e7cdd, []int{2}
}
func (m *Executable) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Executable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Executable.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Executable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Executable.Merge(dst, src)
}
func (m *Executable) XXX_Size() int {
	return m.Size()
}
func (m *Executable) XXX_DiscardUnknown() {
	xxx_messageInfo_Executable.DiscardUnknown(m)
}

var xxx_messageInfo_Executable proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ByteRange)(nil), "manifest.ByteRange")
	proto.RegisterType((*Archive)(nil), "manifest.Archive")
	proto.RegisterMapType((map[uint32]uint32)(nil), "manifest.Archive.EntryAddrsEntry")
	proto.RegisterMapType((map[string]uint32)(nil), "manifest.Archive.EntryIndexesEntry")
	proto.RegisterType((*Executable)(nil), "manifest.Executable")
}
func (m *ByteRange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ByteRange) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Offset != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintManifest(dAtA, i, uint64(m.Offset))
	}
	if m.Length != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintManifest(dAtA, i, uint64(m.Length))
	}
	return i, nil
}

func (m *Archive) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Archive) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ModuleSize != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintManifest(dAtA, i, uint64(m.ModuleSize))
	}
	if len(m.Sections) > 0 {
		for _, msg := range m.Sections {
			dAtA[i] = 0x12
			i++
			i = encodeVarintManifest(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintManifest(dAtA, i, uint64(m.StackSection.Size()))
	n1, err := m.StackSection.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.GlobalTypes) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintManifest(dAtA, i, uint64(len(m.GlobalTypes)))
		i += copy(dAtA[i:], m.GlobalTypes)
	}
	if len(m.EntryIndexes) > 0 {
		for k, _ := range m.EntryIndexes {
			dAtA[i] = 0x2a
			i++
			v := m.EntryIndexes[k]
			mapSize := 1 + len(k) + sovManifest(uint64(len(k))) + 1 + sovManifest(uint64(v))
			i = encodeVarintManifest(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintManifest(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x10
			i++
			i = encodeVarintManifest(dAtA, i, uint64(v))
		}
	}
	if len(m.EntryAddrs) > 0 {
		for k, _ := range m.EntryAddrs {
			dAtA[i] = 0x32
			i++
			v := m.EntryAddrs[k]
			mapSize := 1 + sovManifest(uint64(k)) + 1 + sovManifest(uint64(v))
			i = encodeVarintManifest(dAtA, i, uint64(mapSize))
			dAtA[i] = 0x8
			i++
			i = encodeVarintManifest(dAtA, i, uint64(k))
			dAtA[i] = 0x10
			i++
			i = encodeVarintManifest(dAtA, i, uint64(v))
		}
	}
	if m.CallSitesSize != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintManifest(dAtA, i, uint64(m.CallSitesSize))
	}
	if m.FuncAddrsSize != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintManifest(dAtA, i, uint64(m.FuncAddrsSize))
	}
	dAtA[i] = 0x4a
	i++
	i = encodeVarintManifest(dAtA, i, uint64(m.Exe.Size()))
	n2, err := m.Exe.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func (m *Executable) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Executable) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TextAddr != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintManifest(dAtA, i, uint64(m.TextAddr))
	}
	if m.TextSize != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintManifest(dAtA, i, uint64(m.TextSize))
	}
	if m.StackSize != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintManifest(dAtA, i, uint64(m.StackSize))
	}
	if m.StackUsage != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintManifest(dAtA, i, uint64(m.StackUsage))
	}
	if m.GlobalsSize != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintManifest(dAtA, i, uint64(m.GlobalsSize))
	}
	if m.MemoryDataSize != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintManifest(dAtA, i, uint64(m.MemoryDataSize))
	}
	if m.MemorySize != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintManifest(dAtA, i, uint64(m.MemorySize))
	}
	if m.MemorySizeLimit != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintManifest(dAtA, i, uint64(m.MemorySizeLimit))
	}
	if m.InitRoutine != 0 {
		dAtA[i] = 0x48
		i++
		i = encodeVarintManifest(dAtA, i, uint64(m.InitRoutine))
	}
	return i, nil
}

func encodeVarintManifest(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ByteRange) Size() (n int) {
	var l int
	_ = l
	if m.Offset != 0 {
		n += 1 + sovManifest(uint64(m.Offset))
	}
	if m.Length != 0 {
		n += 1 + sovManifest(uint64(m.Length))
	}
	return n
}

func (m *Archive) Size() (n int) {
	var l int
	_ = l
	if m.ModuleSize != 0 {
		n += 1 + sovManifest(uint64(m.ModuleSize))
	}
	if len(m.Sections) > 0 {
		for _, e := range m.Sections {
			l = e.Size()
			n += 1 + l + sovManifest(uint64(l))
		}
	}
	l = m.StackSection.Size()
	n += 1 + l + sovManifest(uint64(l))
	l = len(m.GlobalTypes)
	if l > 0 {
		n += 1 + l + sovManifest(uint64(l))
	}
	if len(m.EntryIndexes) > 0 {
		for k, v := range m.EntryIndexes {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovManifest(uint64(len(k))) + 1 + sovManifest(uint64(v))
			n += mapEntrySize + 1 + sovManifest(uint64(mapEntrySize))
		}
	}
	if len(m.EntryAddrs) > 0 {
		for k, v := range m.EntryAddrs {
			_ = k
			_ = v
			mapEntrySize := 1 + sovManifest(uint64(k)) + 1 + sovManifest(uint64(v))
			n += mapEntrySize + 1 + sovManifest(uint64(mapEntrySize))
		}
	}
	if m.CallSitesSize != 0 {
		n += 1 + sovManifest(uint64(m.CallSitesSize))
	}
	if m.FuncAddrsSize != 0 {
		n += 1 + sovManifest(uint64(m.FuncAddrsSize))
	}
	l = m.Exe.Size()
	n += 1 + l + sovManifest(uint64(l))
	return n
}

func (m *Executable) Size() (n int) {
	var l int
	_ = l
	if m.TextAddr != 0 {
		n += 1 + sovManifest(uint64(m.TextAddr))
	}
	if m.TextSize != 0 {
		n += 1 + sovManifest(uint64(m.TextSize))
	}
	if m.StackSize != 0 {
		n += 1 + sovManifest(uint64(m.StackSize))
	}
	if m.StackUsage != 0 {
		n += 1 + sovManifest(uint64(m.StackUsage))
	}
	if m.GlobalsSize != 0 {
		n += 1 + sovManifest(uint64(m.GlobalsSize))
	}
	if m.MemoryDataSize != 0 {
		n += 1 + sovManifest(uint64(m.MemoryDataSize))
	}
	if m.MemorySize != 0 {
		n += 1 + sovManifest(uint64(m.MemorySize))
	}
	if m.MemorySizeLimit != 0 {
		n += 1 + sovManifest(uint64(m.MemorySizeLimit))
	}
	if m.InitRoutine != 0 {
		n += 1 + sovManifest(uint64(m.InitRoutine))
	}
	return n
}

func sovManifest(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozManifest(x uint64) (n int) {
	return sovManifest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ByteRange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManifest
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
			return fmt.Errorf("proto: ByteRange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ByteRange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Length", wireType)
			}
			m.Length = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Length |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipManifest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManifest
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
func (m *Archive) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManifest
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
			return fmt.Errorf("proto: Archive: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Archive: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleSize", wireType)
			}
			m.ModuleSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ModuleSize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sections", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthManifest
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sections = append(m.Sections, ByteRange{})
			if err := m.Sections[len(m.Sections)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StackSection", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthManifest
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StackSection.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalTypes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
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
				return ErrInvalidLengthManifest
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GlobalTypes = append(m.GlobalTypes[:0], dAtA[iNdEx:postIndex]...)
			if m.GlobalTypes == nil {
				m.GlobalTypes = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntryIndexes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthManifest
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EntryIndexes == nil {
				m.EntryIndexes = make(map[string]uint32)
			}
			var mapkey string
			var mapvalue uint32
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowManifest
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowManifest
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthManifest
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowManifest
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= (uint32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipManifest(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthManifest
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.EntryIndexes[mapkey] = mapvalue
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntryAddrs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthManifest
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EntryAddrs == nil {
				m.EntryAddrs = make(map[uint32]uint32)
			}
			var mapkey uint32
			var mapvalue uint32
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowManifest
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowManifest
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= (uint32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowManifest
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= (uint32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipManifest(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthManifest
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.EntryAddrs[mapkey] = mapvalue
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallSitesSize", wireType)
			}
			m.CallSitesSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CallSitesSize |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FuncAddrsSize", wireType)
			}
			m.FuncAddrsSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FuncAddrsSize |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exe", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthManifest
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Exe.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipManifest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManifest
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
func (m *Executable) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManifest
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
			return fmt.Errorf("proto: Executable: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Executable: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TextAddr", wireType)
			}
			m.TextAddr = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TextAddr |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TextSize", wireType)
			}
			m.TextSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TextSize |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StackSize", wireType)
			}
			m.StackSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StackSize |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StackUsage", wireType)
			}
			m.StackUsage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StackUsage |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalsSize", wireType)
			}
			m.GlobalsSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GlobalsSize |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemoryDataSize", wireType)
			}
			m.MemoryDataSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemoryDataSize |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemorySize", wireType)
			}
			m.MemorySize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemorySize |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemorySizeLimit", wireType)
			}
			m.MemorySizeLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemorySizeLimit |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitRoutine", wireType)
			}
			m.InitRoutine = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManifest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitRoutine |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipManifest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManifest
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
func skipManifest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowManifest
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
					return 0, ErrIntOverflowManifest
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
					return 0, ErrIntOverflowManifest
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
				return 0, ErrInvalidLengthManifest
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowManifest
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
				next, err := skipManifest(dAtA[start:])
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
	ErrInvalidLengthManifest = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowManifest   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("internal/manifest/manifest.proto", fileDescriptor_manifest_0e5e7cdd968e7cdd)
}

var fileDescriptor_manifest_0e5e7cdd968e7cdd = []byte{
	// 564 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xdf, 0x8a, 0xd3, 0x4e,
	0x14, 0xc7, 0x37, 0x4d, 0xdb, 0x6d, 0x4f, 0x36, 0xbf, 0xee, 0xce, 0xaf, 0x48, 0x58, 0xb1, 0xcd,
	0x56, 0x58, 0x8a, 0x48, 0x05, 0x45, 0x10, 0x45, 0x65, 0x8b, 0x0b, 0x0a, 0xe2, 0x45, 0xaa, 0xd7,
	0x61, 0x9a, 0x9e, 0x76, 0x87, 0x4d, 0x93, 0x25, 0x33, 0x59, 0xda, 0x7d, 0x0a, 0xdf, 0xc9, 0x9b,
	0x5e, 0xfa, 0x00, 0x22, 0x5a, 0x5f, 0x44, 0x66, 0x4e, 0x9a, 0x16, 0x45, 0xbd, 0xcb, 0x7c, 0xcf,
	0x67, 0xbe, 0x73, 0xfe, 0x11, 0xf0, 0x45, 0xa2, 0x30, 0x4b, 0x78, 0xfc, 0x60, 0xce, 0x13, 0x31,
	0x45, 0xa9, 0xca, 0x8f, 0xc1, 0x55, 0x96, 0xaa, 0x94, 0x35, 0x36, 0xe7, 0xde, 0x33, 0x68, 0x0e,
	0x97, 0x0a, 0x03, 0x9e, 0xcc, 0x90, 0xdd, 0x82, 0x7a, 0x3a, 0x9d, 0x4a, 0x54, 0x9e, 0xe5, 0x5b,
	0x7d, 0x3b, 0x28, 0x4e, 0x5a, 0x8f, 0x31, 0x99, 0xa9, 0x0b, 0xaf, 0x42, 0x3a, 0x9d, 0x7a, 0x5f,
	0xaa, 0xb0, 0x7f, 0x96, 0x45, 0x17, 0xe2, 0x1a, 0x59, 0x17, 0x9c, 0x79, 0x3a, 0xc9, 0x63, 0x0c,
	0xa5, 0xb8, 0xc1, 0xc2, 0x00, 0x48, 0x1a, 0x89, 0x1b, 0x64, 0x8f, 0xa1, 0x21, 0x31, 0x52, 0x22,
	0x4d, 0xa4, 0x57, 0xf1, 0xed, 0xbe, 0xf3, 0xf0, 0xff, 0x41, 0x99, 0x56, 0x99, 0xc3, 0xb0, 0xba,
	0xfa, 0xda, 0xdd, 0x0b, 0x4a, 0x94, 0xbd, 0x00, 0x57, 0x2a, 0x1e, 0x5d, 0x86, 0x85, 0xe2, 0xd9,
	0xbe, 0xf5, 0xf7, 0xbb, 0x07, 0x86, 0x1f, 0x11, 0xce, 0x4e, 0xe0, 0x60, 0x16, 0xa7, 0x63, 0x1e,
	0x87, 0x6a, 0x79, 0x85, 0xd2, 0xab, 0xfa, 0x56, 0xff, 0x20, 0x70, 0x48, 0x7b, 0xaf, 0x25, 0xf6,
	0x0e, 0x5c, 0x4c, 0x54, 0xb6, 0x0c, 0x45, 0x32, 0xc1, 0x05, 0x4a, 0xaf, 0x66, 0xd2, 0xbb, 0xbb,
	0x7d, 0xa2, 0x28, 0x72, 0x70, 0xae, 0xb1, 0x37, 0x44, 0x99, 0xef, 0xcd, 0x93, 0xb8, 0x13, 0x60,
	0xaf, 0xc1, 0x21, 0x3f, 0x3e, 0x99, 0x64, 0xd2, 0xab, 0x1b, 0xb7, 0x93, 0x3f, 0xb8, 0x9d, 0x69,
	0x66, 0xd7, 0x0b, 0xb0, 0x94, 0xd9, 0x29, 0xb4, 0x22, 0x1e, 0xc7, 0xa1, 0x14, 0x0a, 0x25, 0x35,
	0x76, 0xdf, 0xb7, 0xfa, 0x6e, 0xe0, 0x6a, 0x79, 0xa4, 0x55, 0xd3, 0xdb, 0x53, 0x68, 0x4d, 0xf3,
	0x24, 0xa2, 0x07, 0x89, 0x6b, 0x10, 0xa7, 0x65, 0xe3, 0x65, 0xb8, 0xfb, 0x60, 0xe3, 0x02, 0xbd,
	0xa6, 0x69, 0x61, 0x7b, 0x9b, 0xd1, 0xf9, 0x02, 0xa3, 0x5c, 0xf1, 0x71, 0xbc, 0xe9, 0xa1, 0xc6,
	0x8e, 0x5f, 0xc2, 0xd1, 0x6f, 0x05, 0xb3, 0x43, 0xb0, 0x2f, 0x71, 0x69, 0xe6, 0xdb, 0x0c, 0xf4,
	0x27, 0x6b, 0x43, 0xed, 0x9a, 0xc7, 0x39, 0x9a, 0xe5, 0x70, 0x03, 0x3a, 0x3c, 0xad, 0x3c, 0xb1,
	0x8e, 0x9f, 0x43, 0xeb, 0x97, 0x1a, 0x77, 0xaf, 0xbb, 0xff, 0xb8, 0xde, 0xfb, 0x54, 0x01, 0xd8,
	0x66, 0xc6, 0x6e, 0x43, 0x53, 0xe1, 0x42, 0x99, 0x22, 0x8d, 0x41, 0x35, 0x68, 0x68, 0x41, 0xbb,
	0x97, 0x41, 0x53, 0x3b, 0x39, 0x99, 0xa0, 0x29, 0xfb, 0x0e, 0x40, 0xb1, 0x43, 0x3a, 0x6a, 0x9b,
	0x68, 0x93, 0xb6, 0x44, 0x87, 0xbb, 0xe0, 0x50, 0x38, 0x97, 0x7c, 0x86, 0x66, 0x43, 0xdc, 0x80,
	0x6e, 0x7c, 0xd0, 0xca, 0x76, 0x87, 0x8a, 0xde, 0xd6, 0x0c, 0x51, 0xec, 0x10, 0x75, 0xb6, 0x0f,
	0x87, 0x73, 0x9c, 0xa7, 0xd9, 0x32, 0x9c, 0x70, 0xc5, 0x09, 0xab, 0x1b, 0xec, 0x3f, 0xd2, 0x5f,
	0x71, 0xc5, 0x37, 0xaf, 0x15, 0xe4, 0xce, 0x3c, 0x81, 0x24, 0x03, 0xdc, 0x83, 0xa3, 0x1d, 0x20,
	0x8c, 0xc5, 0x5c, 0xa8, 0x62, 0x9c, 0xad, 0x2d, 0xf6, 0x56, 0xcb, 0x3a, 0x33, 0x91, 0x08, 0x15,
	0x66, 0x69, 0xae, 0x44, 0x42, 0x93, 0xad, 0x05, 0x8e, 0xd6, 0x02, 0x92, 0x86, 0xed, 0xd5, 0xf7,
	0xce, 0xde, 0x6a, 0xdd, 0xb1, 0x3e, 0xaf, 0x3b, 0xd6, 0xb7, 0x75, 0xc7, 0xfa, 0xf8, 0xa3, 0xb3,
	0x37, 0xae, 0x9b, 0x1f, 0xc1, 0xa3, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xdf, 0x1e, 0xf2,
	0x2c, 0x04, 0x00, 0x00,
}