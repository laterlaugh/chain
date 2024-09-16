// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tunnel/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState represents the initial state of the blockchain.
type GenesisState struct {
	// Params is all parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// port_id is the port id of the module.
	PortID string `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	// tunnel_count is the number of tunnels.
	TunnelCount uint64 `protobuf:"varint,3,opt,name=tunnel_count,json=tunnelCount,proto3" json:"tunnel_count,omitempty"`
	// tunnels is the list of tunnels.
	Tunnels []Tunnel `protobuf:"bytes,4,rep,name=tunnels,proto3" json:"tunnels"`
	// signal_prices_infos is the signal prices info.
	SignalPricesInfos []SignalPricesInfo `protobuf:"bytes,5,rep,name=signal_prices_infos,json=signalPricesInfos,proto3" json:"signal_prices_infos"`
	// TotalFees is the type for the total fees collected by the tunnel
	TotalFees TotalFees `protobuf:"bytes,6,opt,name=total_fees,json=totalFees,proto3" json:"total_fees"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb87d9373a74da2e, []int{0}
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

func (m *GenesisState) GetPortID() string {
	if m != nil {
		return m.PortID
	}
	return ""
}

func (m *GenesisState) GetTunnelCount() uint64 {
	if m != nil {
		return m.TunnelCount
	}
	return 0
}

func (m *GenesisState) GetTunnels() []Tunnel {
	if m != nil {
		return m.Tunnels
	}
	return nil
}

func (m *GenesisState) GetSignalPricesInfos() []SignalPricesInfo {
	if m != nil {
		return m.SignalPricesInfos
	}
	return nil
}

func (m *GenesisState) GetTotalFees() TotalFees {
	if m != nil {
		return m.TotalFees
	}
	return TotalFees{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "tunnel.v1beta1.GenesisState")
}

func init() { proto.RegisterFile("tunnel/v1beta1/genesis.proto", fileDescriptor_cb87d9373a74da2e) }

var fileDescriptor_cb87d9373a74da2e = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x5b, 0xe0, 0x96, 0x30, 0x90, 0x9b, 0xdc, 0xb9, 0x37, 0x37, 0x15, 0x4d, 0xa9, 0xba,
	0x61, 0xd5, 0x09, 0x68, 0x5c, 0xba, 0x40, 0xa3, 0xe9, 0x8e, 0x80, 0x71, 0xe1, 0xa6, 0x99, 0x96,
	0xa1, 0x4c, 0x52, 0x66, 0x9a, 0xce, 0x40, 0xf4, 0x2d, 0x7c, 0x2c, 0x96, 0x2c, 0x5d, 0x18, 0x62,
	0xca, 0x8b, 0x98, 0xce, 0x14, 0x13, 0x51, 0x77, 0x73, 0xfe, 0xff, 0x9b, 0xff, 0x9c, 0x9c, 0x03,
	0x8e, 0xe4, 0x82, 0x31, 0x92, 0xa0, 0x65, 0x2f, 0x24, 0x12, 0xf7, 0x50, 0x4c, 0x18, 0x11, 0x54,
	0x78, 0x69, 0xc6, 0x25, 0x87, 0xbf, 0xb5, 0xeb, 0x95, 0x6e, 0xfb, 0x5f, 0xcc, 0x63, 0xae, 0x2c,
	0x54, 0xbc, 0x34, 0xd5, 0x3e, 0xdc, 0xcb, 0x48, 0x71, 0x86, 0xe7, 0xe2, 0x07, 0xb3, 0x4c, 0x54,
	0xe6, 0xc9, 0x6b, 0x05, 0xb4, 0x6e, 0x75, 0xc7, 0xb1, 0xc4, 0x92, 0xc0, 0x73, 0x60, 0xe9, 0xdf,
	0xb6, 0xe9, 0x9a, 0xdd, 0x66, 0xff, 0xbf, 0xf7, 0x79, 0x02, 0x6f, 0xa8, 0xdc, 0x41, 0x6d, 0xb5,
	0xe9, 0x18, 0xa3, 0x92, 0x85, 0xa7, 0xa0, 0x9e, 0xf2, 0x4c, 0x06, 0x74, 0x62, 0x57, 0x5c, 0xb3,
	0xdb, 0x18, 0x80, 0x7c, 0xd3, 0xb1, 0x86, 0x3c, 0x93, 0xfe, 0xf5, 0xc8, 0x2a, 0x2c, 0x7f, 0x02,
	0x8f, 0x41, 0x4b, 0x67, 0x05, 0x11, 0x5f, 0x30, 0x69, 0x57, 0x5d, 0xb3, 0x5b, 0x1b, 0x35, 0xb5,
	0x76, 0x55, 0x48, 0xf0, 0x02, 0xd4, 0x75, 0x29, 0xec, 0x9a, 0x5b, 0xfd, 0xae, 0xfd, 0x9d, 0x2a,
	0xcb, 0xf6, 0x3b, 0x18, 0xde, 0x83, 0xbf, 0x82, 0xc6, 0x0c, 0x27, 0x41, 0x9a, 0xd1, 0x88, 0x88,
	0x80, 0xb2, 0x29, 0x17, 0xf6, 0x2f, 0x95, 0xe1, 0xee, 0x67, 0x8c, 0x15, 0x3a, 0x54, 0xa4, 0xcf,
	0xa6, 0xbc, 0x4c, 0xfb, 0x23, 0xf6, 0x74, 0x01, 0x2f, 0x01, 0x90, 0x5c, 0xe2, 0x24, 0x98, 0x12,
	0x22, 0x6c, 0x4b, 0x6d, 0xe4, 0xe0, 0xcb, 0x48, 0x05, 0x71, 0x43, 0xc8, 0x6e, 0x29, 0x0d, 0xf9,
	0x21, 0xf8, 0xab, 0xdc, 0x31, 0xd7, 0xb9, 0x63, 0xbe, 0xe5, 0x8e, 0xf9, 0xbc, 0x75, 0x8c, 0xf5,
	0xd6, 0x31, 0x5e, 0xb6, 0x8e, 0xf1, 0x80, 0x62, 0x2a, 0x67, 0x8b, 0xd0, 0x8b, 0xf8, 0x1c, 0x85,
	0x98, 0x4d, 0xd4, 0x39, 0x22, 0x9e, 0xa0, 0x68, 0x86, 0x29, 0x43, 0xcb, 0x3e, 0x7a, 0x2c, 0x2f,
	0x85, 0xe4, 0x53, 0x4a, 0x44, 0x68, 0x29, 0xe2, 0xec, 0x3d, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x9b,
	0xb9, 0x50, 0x30, 0x02, 0x00, 0x00,
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
	{
		size, err := m.TotalFees.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.SignalPricesInfos) > 0 {
		for iNdEx := len(m.SignalPricesInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SignalPricesInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Tunnels) > 0 {
		for iNdEx := len(m.Tunnels) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tunnels[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.TunnelCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TunnelCount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.PortID) > 0 {
		i -= len(m.PortID)
		copy(dAtA[i:], m.PortID)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortID)))
		i--
		dAtA[i] = 0x12
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
	l = len(m.PortID)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.TunnelCount != 0 {
		n += 1 + sovGenesis(uint64(m.TunnelCount))
	}
	if len(m.Tunnels) > 0 {
		for _, e := range m.Tunnels {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SignalPricesInfos) > 0 {
		for _, e := range m.SignalPricesInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.TotalFees.Size()
	n += 1 + l + sovGenesis(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field PortID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TunnelCount", wireType)
			}
			m.TunnelCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TunnelCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tunnels", wireType)
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
			m.Tunnels = append(m.Tunnels, Tunnel{})
			if err := m.Tunnels[len(m.Tunnels)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignalPricesInfos", wireType)
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
			m.SignalPricesInfos = append(m.SignalPricesInfos, SignalPricesInfo{})
			if err := m.SignalPricesInfos[len(m.SignalPricesInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalFees", wireType)
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
			if err := m.TotalFees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
