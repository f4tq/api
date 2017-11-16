// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proxy/v1/config/ingress_rule.proto

package istio_proxy_v1_config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ingress rules are routing rules applied to the ingress proxy pool. The
// ingress proxes serve as the receiving edge proxy for the entire mesh, but
// can also be addressed from inside the mesh.  Each ingress rule defines a
// destination service and port. Rules that do not resolve to a service or a
// port in the mesh should be ignored.
//
// The routing rules for the destination service are applied at the ingress
// proxy. That means the routing rule match conditions are composed and its
// actions are enforced. The traffic splitting for the destination service is
// also effective.
//
// WARNING: This API is experimental and under active development
type IngressRule struct {
	// REQUIRED: Port on which the ingress proxy listens and applies the rule.
	Port int32 `protobuf:"varint,1,opt,name=port" json:"port,omitempty"`
	// Optional TLS secret path to apply server-side TLS context on the port.
	// It is up to the underlying secret store to interpret the path to the secret.
	TlsSecret string `protobuf:"bytes,2,opt,name=tls_secret,json=tlsSecret" json:"tls_secret,omitempty"`
	// RECOMMENDED. Precedence is used to disambiguate the order of application
	// of rules. A higher number takes priority. If not specified, the value is
	// assumed to be 0.  The order of application for rules with the same
	// precedence is unspecified.
	Precedence int32 `protobuf:"varint,3,opt,name=precedence" json:"precedence,omitempty"`
	// Match conditions to be satisfied for the ingress rule to be
	// activated.
	Match *MatchCondition `protobuf:"bytes,4,opt,name=match" json:"match,omitempty"`
	// REQUIRED: Destination uniquely identifies the destination service.
	//
	// *Note:* The ingress rule destination specification represents all version
	// of the service and therefore the IstioService's labels field MUST be empty.
	//
	Destination *IstioService `protobuf:"bytes,5,opt,name=destination" json:"destination,omitempty"`
	// REQUIRED: Destination port identifies a port on the destination service for routing.
	//
	// Types that are valid to be assigned to DestinationServicePort:
	//	*IngressRule_DestinationPort
	//	*IngressRule_DestinationPortName
	DestinationServicePort isIngressRule_DestinationServicePort `protobuf_oneof:"destination_service_port"`
}

func (m *IngressRule) Reset()                    { *m = IngressRule{} }
func (m *IngressRule) String() string            { return proto.CompactTextString(m) }
func (*IngressRule) ProtoMessage()               {}
func (*IngressRule) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type isIngressRule_DestinationServicePort interface {
	isIngressRule_DestinationServicePort()
}

type IngressRule_DestinationPort struct {
	DestinationPort int32 `protobuf:"varint,6,opt,name=destination_port,json=destinationPort,oneof"`
}
type IngressRule_DestinationPortName struct {
	DestinationPortName string `protobuf:"bytes,7,opt,name=destination_port_name,json=destinationPortName,oneof"`
}

func (*IngressRule_DestinationPort) isIngressRule_DestinationServicePort()     {}
func (*IngressRule_DestinationPortName) isIngressRule_DestinationServicePort() {}

func (m *IngressRule) GetDestinationServicePort() isIngressRule_DestinationServicePort {
	if m != nil {
		return m.DestinationServicePort
	}
	return nil
}

func (m *IngressRule) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *IngressRule) GetTlsSecret() string {
	if m != nil {
		return m.TlsSecret
	}
	return ""
}

func (m *IngressRule) GetPrecedence() int32 {
	if m != nil {
		return m.Precedence
	}
	return 0
}

func (m *IngressRule) GetMatch() *MatchCondition {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *IngressRule) GetDestination() *IstioService {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *IngressRule) GetDestinationPort() int32 {
	if x, ok := m.GetDestinationServicePort().(*IngressRule_DestinationPort); ok {
		return x.DestinationPort
	}
	return 0
}

func (m *IngressRule) GetDestinationPortName() string {
	if x, ok := m.GetDestinationServicePort().(*IngressRule_DestinationPortName); ok {
		return x.DestinationPortName
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*IngressRule) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _IngressRule_OneofMarshaler, _IngressRule_OneofUnmarshaler, _IngressRule_OneofSizer, []interface{}{
		(*IngressRule_DestinationPort)(nil),
		(*IngressRule_DestinationPortName)(nil),
	}
}

func _IngressRule_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*IngressRule)
	// destination_service_port
	switch x := m.DestinationServicePort.(type) {
	case *IngressRule_DestinationPort:
		b.EncodeVarint(6<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.DestinationPort))
	case *IngressRule_DestinationPortName:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.DestinationPortName)
	case nil:
	default:
		return fmt.Errorf("IngressRule.DestinationServicePort has unexpected type %T", x)
	}
	return nil
}

func _IngressRule_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*IngressRule)
	switch tag {
	case 6: // destination_service_port.destination_port
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.DestinationServicePort = &IngressRule_DestinationPort{int32(x)}
		return true, err
	case 7: // destination_service_port.destination_port_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.DestinationServicePort = &IngressRule_DestinationPortName{x}
		return true, err
	default:
		return false, nil
	}
}

func _IngressRule_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*IngressRule)
	// destination_service_port
	switch x := m.DestinationServicePort.(type) {
	case *IngressRule_DestinationPort:
		n += proto.SizeVarint(6<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.DestinationPort))
	case *IngressRule_DestinationPortName:
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.DestinationPortName)))
		n += len(x.DestinationPortName)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*IngressRule)(nil), "istio.proxy.v1.config.IngressRule")
}

func init() { proto.RegisterFile("proxy/v1/config/ingress_rule.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xd1, 0x4a, 0xf3, 0x30,
	0x1c, 0xc5, 0xd7, 0x7d, 0xdb, 0x3e, 0xf6, 0xdf, 0x85, 0x12, 0x19, 0x84, 0x81, 0x52, 0x26, 0xc2,
	0x40, 0x48, 0x99, 0x7a, 0xe7, 0x9d, 0x22, 0x6c, 0x17, 0x8a, 0x64, 0x0f, 0x50, 0x6a, 0xfa, 0x77,
	0x06, 0xda, 0xa4, 0x24, 0x69, 0xd1, 0xe7, 0xf5, 0x45, 0x24, 0xe9, 0x4d, 0x28, 0x7a, 0x97, 0x9c,
	0x9c, 0xdf, 0xe1, 0x9c, 0xc0, 0xba, 0x31, 0xfa, 0xf3, 0x2b, 0xeb, 0xb6, 0x99, 0xd0, 0xea, 0x5d,
	0x1e, 0x33, 0xa9, 0x8e, 0x06, 0xad, 0xcd, 0x4d, 0x5b, 0x21, 0x6b, 0x8c, 0x76, 0x9a, 0x2c, 0xa5,
	0x75, 0x52, 0xb3, 0xe0, 0x64, 0xdd, 0x96, 0xf5, 0xce, 0x55, 0x3a, 0x44, 0x8d, 0x6e, 0x1d, 0x46,
	0xe0, 0xfa, 0x7b, 0x0c, 0x8b, 0x7d, 0x9f, 0xc7, 0xdb, 0x0a, 0x09, 0x81, 0x49, 0xa3, 0x8d, 0xa3,
	0x49, 0x9a, 0x6c, 0xa6, 0x3c, 0x9c, 0xc9, 0x39, 0x80, 0xab, 0x6c, 0x6e, 0x51, 0x18, 0x74, 0x74,
	0x9c, 0x26, 0x9b, 0x39, 0x9f, 0xbb, 0xca, 0x1e, 0x82, 0x40, 0x2e, 0x00, 0x1a, 0x83, 0x02, 0x4b,
	0x54, 0x02, 0xe9, 0xbf, 0x00, 0x46, 0x0a, 0xb9, 0x87, 0x69, 0x5d, 0x38, 0xf1, 0x41, 0x27, 0x69,
	0xb2, 0x59, 0xdc, 0x5c, 0xb1, 0x5f, 0xbb, 0xb2, 0x67, 0xef, 0x79, 0xd4, 0xaa, 0x94, 0x4e, 0x6a,
	0xc5, 0x7b, 0x86, 0x3c, 0xc1, 0xa2, 0x44, 0xeb, 0xa4, 0x2a, 0xbc, 0x4a, 0xa7, 0x21, 0xe2, 0xf2,
	0x8f, 0x88, 0xbd, 0x57, 0x0f, 0x68, 0x3a, 0x29, 0x90, 0xc7, 0x1c, 0xb9, 0x86, 0xd3, 0xe8, 0x9a,
	0x87, 0x89, 0x33, 0xdf, 0x74, 0x37, 0xe2, 0x27, 0xd1, 0xcb, 0xab, 0xdf, 0x7b, 0x07, 0xcb, 0xa1,
	0x39, 0x57, 0x45, 0x8d, 0xf4, 0xbf, 0x9f, 0xbe, 0x1b, 0xf1, 0xb3, 0x01, 0xf1, 0x52, 0xd4, 0xf8,
	0xb0, 0x02, 0x1a, 0x53, 0xb6, 0xaf, 0x11, 0xe8, 0xb7, 0x59, 0xf8, 0xec, 0xdb, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x68, 0x87, 0x6f, 0x6c, 0xcb, 0x01, 0x00, 0x00,
}
