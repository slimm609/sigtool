// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sign/hdr.proto

package sign

import (
	bytes "bytes"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

//
// Every encrypted file starts with a header describing the
// Block Size, Salt, Recipient keys etc. Header represents a
// decoded version of this information. It is encoded in
// protobuf format before writing to disk.
type Header struct {
	ChunkSize uint32        `protobuf:"varint,1,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"`
	Salt      []byte        `protobuf:"bytes,2,opt,name=salt,proto3" json:"salt,omitempty"`
	Keys      []*WrappedKey `protobuf:"bytes,3,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (m *Header) Reset()      { *m = Header{} }
func (*Header) ProtoMessage() {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_85aff542c746609f, []int{0}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Header.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return m.Size()
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetChunkSize() uint32 {
	if m != nil {
		return m.ChunkSize
	}
	return 0
}

func (m *Header) GetSalt() []byte {
	if m != nil {
		return m.Salt
	}
	return nil
}

func (m *Header) GetKeys() []*WrappedKey {
	if m != nil {
		return m.Keys
	}
	return nil
}

//
// A file encryption key is wrapped by a recipient specific public
// key. WrappedKey describes such a wrapped key.
type WrappedKey struct {
	PkHash []byte `protobuf:"bytes,1,opt,name=pk_hash,json=pkHash,proto3" json:"pk_hash,omitempty"`
	Pk     []byte `protobuf:"bytes,2,opt,name=pk,proto3" json:"pk,omitempty"`
	Nonce  []byte `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Key    []byte `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *WrappedKey) Reset()      { *m = WrappedKey{} }
func (*WrappedKey) ProtoMessage() {}
func (*WrappedKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_85aff542c746609f, []int{1}
}
func (m *WrappedKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WrappedKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WrappedKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WrappedKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WrappedKey.Merge(m, src)
}
func (m *WrappedKey) XXX_Size() int {
	return m.Size()
}
func (m *WrappedKey) XXX_DiscardUnknown() {
	xxx_messageInfo_WrappedKey.DiscardUnknown(m)
}

var xxx_messageInfo_WrappedKey proto.InternalMessageInfo

func (m *WrappedKey) GetPkHash() []byte {
	if m != nil {
		return m.PkHash
	}
	return nil
}

func (m *WrappedKey) GetPk() []byte {
	if m != nil {
		return m.Pk
	}
	return nil
}

func (m *WrappedKey) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *WrappedKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func init() {
	proto.RegisterType((*Header)(nil), "sign.header")
	proto.RegisterType((*WrappedKey)(nil), "sign.wrapped_key")
}

func init() { proto.RegisterFile("sign/hdr.proto", fileDescriptor_85aff542c746609f) }

var fileDescriptor_85aff542c746609f = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4a, 0xc4, 0x40,
	0x14, 0x86, 0xe7, 0x25, 0x31, 0xe2, 0xdb, 0x75, 0xd1, 0x41, 0x30, 0x8d, 0x8f, 0xb0, 0x20, 0xa4,
	0x8a, 0xa0, 0x9e, 0xc0, 0xca, 0x3a, 0xf6, 0x86, 0xec, 0x66, 0x70, 0xc2, 0x48, 0x32, 0x64, 0x56,
	0x24, 0x5b, 0x79, 0x04, 0x8f, 0xe1, 0x51, 0x2c, 0x53, 0x6e, 0x69, 0x26, 0x8d, 0xe5, 0x1e, 0x41,
	0x32, 0x5a, 0xd8, 0xfd, 0xff, 0xf7, 0xe0, 0x7d, 0xf0, 0xe3, 0xc2, 0x54, 0x4f, 0xf5, 0x95, 0x2c,
	0xdb, 0x54, 0xb7, 0xcd, 0xa6, 0xe1, 0xc1, 0xd4, 0x97, 0x2b, 0x0c, 0xa5, 0x28, 0x4a, 0xd1, 0xf2,
	0x0b, 0xc4, 0xb5, 0x7c, 0xa9, 0x55, 0x6e, 0xaa, 0xad, 0x88, 0x20, 0x86, 0xe4, 0x38, 0x3b, 0x72,
	0xe4, 0xa1, 0xda, 0x0a, 0xce, 0x31, 0x30, 0xc5, 0xf3, 0x26, 0xf2, 0x62, 0x48, 0xe6, 0x99, 0xcb,
	0xfc, 0x12, 0x03, 0x25, 0x3a, 0x13, 0xf9, 0xb1, 0x9f, 0xcc, 0xae, 0x4f, 0xd3, 0xe9, 0x63, 0xfa,
	0xda, 0x16, 0x5a, 0x8b, 0x32, 0x57, 0xa2, 0xcb, 0xdc, 0x79, 0xf9, 0x88, 0xb3, 0x7f, 0x90, 0x9f,
	0xe3, 0xa1, 0x56, 0xb9, 0x2c, 0x8c, 0x74, 0x96, 0x79, 0x16, 0x6a, 0x75, 0x5f, 0x18, 0xc9, 0x17,
	0xe8, 0x69, 0xf5, 0x27, 0xf0, 0xb4, 0xe2, 0x67, 0x78, 0x50, 0x37, 0xf5, 0x5a, 0x44, 0xbe, 0x43,
	0xbf, 0x85, 0x9f, 0xa0, 0xaf, 0x44, 0x17, 0x05, 0x8e, 0x4d, 0xf1, 0xee, 0xb6, 0x1f, 0x88, 0xed,
	0x06, 0x62, 0xfb, 0x81, 0xe0, 0xcd, 0x12, 0x7c, 0x58, 0x82, 0x4f, 0x4b, 0xd0, 0x5b, 0x82, 0x2f,
	0x4b, 0xf0, 0x6d, 0x89, 0xed, 0x2d, 0xc1, 0xfb, 0x48, 0xac, 0x1f, 0x89, 0xed, 0x46, 0x62, 0xab,
	0xd0, 0xcd, 0x70, 0xf3, 0x13, 0x00, 0x00, 0xff, 0xff, 0xde, 0xf2, 0x28, 0xc0, 0x18, 0x01, 0x00,
	0x00,
}

func (this *Header) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Header)
	if !ok {
		that2, ok := that.(Header)
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
	if this.ChunkSize != that1.ChunkSize {
		return false
	}
	if !bytes.Equal(this.Salt, that1.Salt) {
		return false
	}
	if len(this.Keys) != len(that1.Keys) {
		return false
	}
	for i := range this.Keys {
		if !this.Keys[i].Equal(that1.Keys[i]) {
			return false
		}
	}
	return true
}
func (this *WrappedKey) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WrappedKey)
	if !ok {
		that2, ok := that.(WrappedKey)
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
	if !bytes.Equal(this.PkHash, that1.PkHash) {
		return false
	}
	if !bytes.Equal(this.Pk, that1.Pk) {
		return false
	}
	if !bytes.Equal(this.Nonce, that1.Nonce) {
		return false
	}
	if !bytes.Equal(this.Key, that1.Key) {
		return false
	}
	return true
}
func (this *Header) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&sign.Header{")
	s = append(s, "ChunkSize: "+fmt.Sprintf("%#v", this.ChunkSize)+",\n")
	s = append(s, "Salt: "+fmt.Sprintf("%#v", this.Salt)+",\n")
	if this.Keys != nil {
		s = append(s, "Keys: "+fmt.Sprintf("%#v", this.Keys)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *WrappedKey) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&sign.WrappedKey{")
	s = append(s, "PkHash: "+fmt.Sprintf("%#v", this.PkHash)+",\n")
	s = append(s, "Pk: "+fmt.Sprintf("%#v", this.Pk)+",\n")
	s = append(s, "Nonce: "+fmt.Sprintf("%#v", this.Nonce)+",\n")
	s = append(s, "Key: "+fmt.Sprintf("%#v", this.Key)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringHdr(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Header) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Header) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Header) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Keys) > 0 {
		for iNdEx := len(m.Keys) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Keys[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHdr(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Salt) > 0 {
		i -= len(m.Salt)
		copy(dAtA[i:], m.Salt)
		i = encodeVarintHdr(dAtA, i, uint64(len(m.Salt)))
		i--
		dAtA[i] = 0x12
	}
	if m.ChunkSize != 0 {
		i = encodeVarintHdr(dAtA, i, uint64(m.ChunkSize))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *WrappedKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WrappedKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WrappedKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintHdr(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Nonce) > 0 {
		i -= len(m.Nonce)
		copy(dAtA[i:], m.Nonce)
		i = encodeVarintHdr(dAtA, i, uint64(len(m.Nonce)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Pk) > 0 {
		i -= len(m.Pk)
		copy(dAtA[i:], m.Pk)
		i = encodeVarintHdr(dAtA, i, uint64(len(m.Pk)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PkHash) > 0 {
		i -= len(m.PkHash)
		copy(dAtA[i:], m.PkHash)
		i = encodeVarintHdr(dAtA, i, uint64(len(m.PkHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHdr(dAtA []byte, offset int, v uint64) int {
	offset -= sovHdr(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Header) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChunkSize != 0 {
		n += 1 + sovHdr(uint64(m.ChunkSize))
	}
	l = len(m.Salt)
	if l > 0 {
		n += 1 + l + sovHdr(uint64(l))
	}
	if len(m.Keys) > 0 {
		for _, e := range m.Keys {
			l = e.Size()
			n += 1 + l + sovHdr(uint64(l))
		}
	}
	return n
}

func (m *WrappedKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PkHash)
	if l > 0 {
		n += 1 + l + sovHdr(uint64(l))
	}
	l = len(m.Pk)
	if l > 0 {
		n += 1 + l + sovHdr(uint64(l))
	}
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + sovHdr(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovHdr(uint64(l))
	}
	return n
}

func sovHdr(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHdr(x uint64) (n int) {
	return sovHdr(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Header) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForKeys := "[]*WrappedKey{"
	for _, f := range this.Keys {
		repeatedStringForKeys += strings.Replace(fmt.Sprintf("%v", f), "WrappedKey", "WrappedKey", 1) + ","
	}
	repeatedStringForKeys += "}"
	s := strings.Join([]string{`&Header{`,
		`ChunkSize:` + fmt.Sprintf("%v", this.ChunkSize) + `,`,
		`Salt:` + fmt.Sprintf("%v", this.Salt) + `,`,
		`Keys:` + repeatedStringForKeys + `,`,
		`}`,
	}, "")
	return s
}
func (this *WrappedKey) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&WrappedKey{`,
		`PkHash:` + fmt.Sprintf("%v", this.PkHash) + `,`,
		`Pk:` + fmt.Sprintf("%v", this.Pk) + `,`,
		`Nonce:` + fmt.Sprintf("%v", this.Nonce) + `,`,
		`Key:` + fmt.Sprintf("%v", this.Key) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringHdr(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Header) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHdr
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
			return fmt.Errorf("proto: header: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: header: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChunkSize", wireType)
			}
			m.ChunkSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHdr
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChunkSize |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Salt", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHdr
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
				return ErrInvalidLengthHdr
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHdr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Salt = append(m.Salt[:0], dAtA[iNdEx:postIndex]...)
			if m.Salt == nil {
				m.Salt = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHdr
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
				return ErrInvalidLengthHdr
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHdr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Keys = append(m.Keys, &WrappedKey{})
			if err := m.Keys[len(m.Keys)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHdr(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHdr
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHdr
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
func (m *WrappedKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHdr
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
			return fmt.Errorf("proto: wrapped_key: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: wrapped_key: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PkHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHdr
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
				return ErrInvalidLengthHdr
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHdr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PkHash = append(m.PkHash[:0], dAtA[iNdEx:postIndex]...)
			if m.PkHash == nil {
				m.PkHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pk", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHdr
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
				return ErrInvalidLengthHdr
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHdr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pk = append(m.Pk[:0], dAtA[iNdEx:postIndex]...)
			if m.Pk == nil {
				m.Pk = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHdr
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
				return ErrInvalidLengthHdr
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHdr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = append(m.Nonce[:0], dAtA[iNdEx:postIndex]...)
			if m.Nonce == nil {
				m.Nonce = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHdr
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
				return ErrInvalidLengthHdr
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHdr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHdr(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHdr
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHdr
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
func skipHdr(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHdr
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
					return 0, ErrIntOverflowHdr
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
					return 0, ErrIntOverflowHdr
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
				return 0, ErrInvalidLengthHdr
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHdr
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHdr
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHdr        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHdr          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHdr = fmt.Errorf("proto: unexpected end of group")
)
