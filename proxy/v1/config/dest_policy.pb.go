// Code generated by protoc-gen-go.
// source: proxy/v1/config/dest_policy.proto
// DO NOT EDIT!

/*
Package istio_proxy_v1_config is a generated protocol buffer package.

It is generated from these files:
	proxy/v1/config/dest_policy.proto
	proxy/v1/config/egress_rule.proto
	proxy/v1/config/http_fault.proto
	proxy/v1/config/ingress_rule.proto
	proxy/v1/config/l4_fault.proto
	proxy/v1/config/proxy_mesh.proto
	proxy/v1/config/route_rule.proto

It has these top-level messages:
	DestinationPolicy
	DestinationVersionPolicy
	LoadBalancing
	CircuitBreaker
	EgressRule
	HTTPFaultInjection
	IngressRule
	L4FaultInjection
	ProxyMeshConfig
	RouteRule
	MatchCondition
	DestinationWeight
	L4MatchAttributes
	HTTPRedirect
	HTTPRewrite
	StringMatch
	HTTPTimeout
	HTTPRetry
*/
package istio_proxy_v1_config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"
import google_protobuf1 "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Load balancing algorithms supported by Envoy proxy.
type LoadBalancing_SimpleLBPolicy int32

const (
	// Simple round robin policy.
	LoadBalancing_ROUND_ROBIN LoadBalancing_SimpleLBPolicy = 0
	// The least request load balancer uses an O(1) algorithm which selects
	// two random healthy hosts and picks the host which has fewer active
	// requests.
	LoadBalancing_LEAST_CONN LoadBalancing_SimpleLBPolicy = 1
	// The random load balancer selects a random healthy host. The random
	// load balancer generally performs better than round robin if no health
	// checking policy is configured.
	LoadBalancing_RANDOM LoadBalancing_SimpleLBPolicy = 2
)

var LoadBalancing_SimpleLBPolicy_name = map[int32]string{
	0: "ROUND_ROBIN",
	1: "LEAST_CONN",
	2: "RANDOM",
}
var LoadBalancing_SimpleLBPolicy_value = map[string]int32{
	"ROUND_ROBIN": 0,
	"LEAST_CONN":  1,
	"RANDOM":      2,
}

func (x LoadBalancing_SimpleLBPolicy) String() string {
	return proto.EnumName(LoadBalancing_SimpleLBPolicy_name, int32(x))
}
func (LoadBalancing_SimpleLBPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 0}
}

// DestinationPolicy defines client/caller-side policies that determine how
// to handle traffic bound to a particular destination service. The policy
// specifies configuration for load balancing and circuit breakers.  For
// example, a simple load balancing policy for the reviews service would
// look as follows:
//
//     destination: reviews.default.svc.cluster.local
//     policy:
//     - loadBalancing:
//         name: RANDOM
//
// Policies are applicable per individual service versions. ONLY
// ONE policy can be defined per service version. Policy CANNOT be empty.
type DestinationPolicy struct {
	// REQUIRED. Service name for which the service version is defined. The
	// value MUST BE a fully-qualified domain name,
	// e.g. _my-service.default.svc.cluster.local_.
	Destination string `protobuf:"bytes,1,opt,name=destination" json:"destination,omitempty"`
	// REQUIRED. List of policies, one per service version.
	Policy []*DestinationVersionPolicy `protobuf:"bytes,2,rep,name=policy" json:"policy,omitempty"`
}

func (m *DestinationPolicy) Reset()                    { *m = DestinationPolicy{} }
func (m *DestinationPolicy) String() string            { return proto.CompactTextString(m) }
func (*DestinationPolicy) ProtoMessage()               {}
func (*DestinationPolicy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DestinationPolicy) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *DestinationPolicy) GetPolicy() []*DestinationVersionPolicy {
	if m != nil {
		return m.Policy
	}
	return nil
}

// A destination policy can be restricted to a particular version of a
// service or applied to all versions. The tags field in the
// DestinationVersionPolicy allow restricting the scope of a
// DestinationPolicy. For example, the following load balancing policy
// applies to version v1 of the reviews service running in the prod
// environment:
//
//     destination: reviews.default.svc.cluster.local
//     policy:
//     - tags:
//         env: prod
//         version: v1
//       loadBalancing:
//         name: RANDOM
//
// If tags are omitted, the policy applies for all versions of the
// service. Policy CANNOT BE empty.
// *Note:* Destination policies will be applied only if the corresponding
// tagged instances are explicity routed to. In other words, for every
// destination policy defined, atleast one route rule must refer to the
// service version indicated in the destination policy.
type DestinationVersionPolicy struct {
	// Optional set of tags that identify a particular version of the
	// destination service. If omitted, the policy will apply to all versions
	// of the service. (-- N.B. The map is used instead of
	// pstruct due to lack of serialization support in golang protobuf
	// library (see https://github.com/golang/protobuf/pull/208) --)
	Tags map[string]string `protobuf:"bytes,1,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Load balancing policy.
	LoadBalancing *LoadBalancing `protobuf:"bytes,2,opt,name=load_balancing,json=loadBalancing" json:"load_balancing,omitempty"`
	// Circuit breaker policy.
	CircuitBreaker *CircuitBreaker `protobuf:"bytes,3,opt,name=circuit_breaker,json=circuitBreaker" json:"circuit_breaker,omitempty"`
	// (-- Other custom policy implementations --)
	Custom *google_protobuf.Any `protobuf:"bytes,100,opt,name=custom" json:"custom,omitempty"`
}

func (m *DestinationVersionPolicy) Reset()                    { *m = DestinationVersionPolicy{} }
func (m *DestinationVersionPolicy) String() string            { return proto.CompactTextString(m) }
func (*DestinationVersionPolicy) ProtoMessage()               {}
func (*DestinationVersionPolicy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DestinationVersionPolicy) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *DestinationVersionPolicy) GetLoadBalancing() *LoadBalancing {
	if m != nil {
		return m.LoadBalancing
	}
	return nil
}

func (m *DestinationVersionPolicy) GetCircuitBreaker() *CircuitBreaker {
	if m != nil {
		return m.CircuitBreaker
	}
	return nil
}

func (m *DestinationVersionPolicy) GetCustom() *google_protobuf.Any {
	if m != nil {
		return m.Custom
	}
	return nil
}

// Load balancing policy to use when forwarding traffic. These policies
// directly correlate to [load balancer
// types](https://lyft.github.io/envoy/docs/intro/arch_overview/load_balancing.html)
// supported by Envoy. Example,
//
//     destination: reviews.default.svc.cluster.local
//     policy:
//     - loadBalancing:
//         name: RANDOM
//
type LoadBalancing struct {
	// Types that are valid to be assigned to LbPolicy:
	//	*LoadBalancing_Name
	//	*LoadBalancing_Custom
	LbPolicy isLoadBalancing_LbPolicy `protobuf_oneof:"lb_policy"`
}

func (m *LoadBalancing) Reset()                    { *m = LoadBalancing{} }
func (m *LoadBalancing) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancing) ProtoMessage()               {}
func (*LoadBalancing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isLoadBalancing_LbPolicy interface {
	isLoadBalancing_LbPolicy()
}

type LoadBalancing_Name struct {
	Name LoadBalancing_SimpleLBPolicy `protobuf:"varint,1,opt,name=name,enum=istio.proxy.v1.config.LoadBalancing_SimpleLBPolicy,oneof"`
}
type LoadBalancing_Custom struct {
	Custom *google_protobuf.Any `protobuf:"bytes,2,opt,name=custom,oneof"`
}

func (*LoadBalancing_Name) isLoadBalancing_LbPolicy()   {}
func (*LoadBalancing_Custom) isLoadBalancing_LbPolicy() {}

func (m *LoadBalancing) GetLbPolicy() isLoadBalancing_LbPolicy {
	if m != nil {
		return m.LbPolicy
	}
	return nil
}

func (m *LoadBalancing) GetName() LoadBalancing_SimpleLBPolicy {
	if x, ok := m.GetLbPolicy().(*LoadBalancing_Name); ok {
		return x.Name
	}
	return LoadBalancing_ROUND_ROBIN
}

func (m *LoadBalancing) GetCustom() *google_protobuf.Any {
	if x, ok := m.GetLbPolicy().(*LoadBalancing_Custom); ok {
		return x.Custom
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LoadBalancing) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LoadBalancing_OneofMarshaler, _LoadBalancing_OneofUnmarshaler, _LoadBalancing_OneofSizer, []interface{}{
		(*LoadBalancing_Name)(nil),
		(*LoadBalancing_Custom)(nil),
	}
}

func _LoadBalancing_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LoadBalancing)
	// lb_policy
	switch x := m.LbPolicy.(type) {
	case *LoadBalancing_Name:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Name))
	case *LoadBalancing_Custom:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Custom); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LoadBalancing.LbPolicy has unexpected type %T", x)
	}
	return nil
}

func _LoadBalancing_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LoadBalancing)
	switch tag {
	case 1: // lb_policy.name
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.LbPolicy = &LoadBalancing_Name{LoadBalancing_SimpleLBPolicy(x)}
		return true, err
	case 2: // lb_policy.custom
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf.Any)
		err := b.DecodeMessage(msg)
		m.LbPolicy = &LoadBalancing_Custom{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LoadBalancing_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LoadBalancing)
	// lb_policy
	switch x := m.LbPolicy.(type) {
	case *LoadBalancing_Name:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Name))
	case *LoadBalancing_Custom:
		s := proto.Size(x.Custom)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Circuit breaker configuration for Envoy. The circuit breaker
// implementation is fine-grained in that it tracks the success/failure
// rates of individual hosts in the load balancing pool. Hosts that
// continually return errors for API calls are ejected from the pool for a
// pre-defined period of time.
// See Envoy's
// [circuit breaker](https://lyft.github.io/envoy/docs/intro/arch_overview/circuit_breaking.html)
// and [outlier detection](https://lyft.github.io/envoy/docs/intro/arch_overview/outlier.html)
// for more details.
type CircuitBreaker struct {
	// Types that are valid to be assigned to CbPolicy:
	//	*CircuitBreaker_SimpleCb
	//	*CircuitBreaker_Custom
	CbPolicy isCircuitBreaker_CbPolicy `protobuf_oneof:"cb_policy"`
}

func (m *CircuitBreaker) Reset()                    { *m = CircuitBreaker{} }
func (m *CircuitBreaker) String() string            { return proto.CompactTextString(m) }
func (*CircuitBreaker) ProtoMessage()               {}
func (*CircuitBreaker) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isCircuitBreaker_CbPolicy interface {
	isCircuitBreaker_CbPolicy()
}

type CircuitBreaker_SimpleCb struct {
	SimpleCb *CircuitBreaker_SimpleCircuitBreakerPolicy `protobuf:"bytes,1,opt,name=simple_cb,json=simpleCb,oneof"`
}
type CircuitBreaker_Custom struct {
	Custom *google_protobuf.Any `protobuf:"bytes,2,opt,name=custom,oneof"`
}

func (*CircuitBreaker_SimpleCb) isCircuitBreaker_CbPolicy() {}
func (*CircuitBreaker_Custom) isCircuitBreaker_CbPolicy()   {}

func (m *CircuitBreaker) GetCbPolicy() isCircuitBreaker_CbPolicy {
	if m != nil {
		return m.CbPolicy
	}
	return nil
}

func (m *CircuitBreaker) GetSimpleCb() *CircuitBreaker_SimpleCircuitBreakerPolicy {
	if x, ok := m.GetCbPolicy().(*CircuitBreaker_SimpleCb); ok {
		return x.SimpleCb
	}
	return nil
}

func (m *CircuitBreaker) GetCustom() *google_protobuf.Any {
	if x, ok := m.GetCbPolicy().(*CircuitBreaker_Custom); ok {
		return x.Custom
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CircuitBreaker) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CircuitBreaker_OneofMarshaler, _CircuitBreaker_OneofUnmarshaler, _CircuitBreaker_OneofSizer, []interface{}{
		(*CircuitBreaker_SimpleCb)(nil),
		(*CircuitBreaker_Custom)(nil),
	}
}

func _CircuitBreaker_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CircuitBreaker)
	// cb_policy
	switch x := m.CbPolicy.(type) {
	case *CircuitBreaker_SimpleCb:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SimpleCb); err != nil {
			return err
		}
	case *CircuitBreaker_Custom:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Custom); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CircuitBreaker.CbPolicy has unexpected type %T", x)
	}
	return nil
}

func _CircuitBreaker_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CircuitBreaker)
	switch tag {
	case 1: // cb_policy.simple_cb
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CircuitBreaker_SimpleCircuitBreakerPolicy)
		err := b.DecodeMessage(msg)
		m.CbPolicy = &CircuitBreaker_SimpleCb{msg}
		return true, err
	case 2: // cb_policy.custom
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf.Any)
		err := b.DecodeMessage(msg)
		m.CbPolicy = &CircuitBreaker_Custom{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CircuitBreaker_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CircuitBreaker)
	// cb_policy
	switch x := m.CbPolicy.(type) {
	case *CircuitBreaker_SimpleCb:
		s := proto.Size(x.SimpleCb)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CircuitBreaker_Custom:
		s := proto.Size(x.Custom)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A simple circuit breaker can be set based on a number of criteria such as
// connection and request limits. For example, the following destination
// policy sets a limit of 100 connections to "reviews" service version
// "v1" backends.
//
//     destination: reviews.default.svc.cluster.local
//     policy:
//     - tags:
//         version: v1
//       circuitBreaker:
//         simpleCb:
//           maxConnections: 100
//
// The following destination policy sets a limit of 100 connections and
// 1000 concurrent requests, with no more than 10 req/connection to
// "reviews" service version "v1" backends. In addition, it configures
// hosts to be scanned every 5 mins, such that any host that fails 7
// consecutive times with 5XX error code will be ejected for 15 minutes.
//
//     destination: reviews.default.svc.cluster.local
//     policy:
//     - tags:
//         version: v1
//       circuitBreaker:
//         simpleCb:
//           maxConnections: 100
//           httpMaxRequests: 1000
//           httpMaxRequestsPerConnection: 10
//           httpConsecutiveErrors: 7
//           sleepWindow: 15m
//           httpDetectionInterval: 5m
//
type CircuitBreaker_SimpleCircuitBreakerPolicy struct {
	// Maximum number of connections to a backend.
	MaxConnections int32 `protobuf:"varint,1,opt,name=max_connections,json=maxConnections" json:"max_connections,omitempty"`
	// Maximum number of pending requests to a backend. Default 1024
	HttpMaxPendingRequests int32 `protobuf:"varint,2,opt,name=http_max_pending_requests,json=httpMaxPendingRequests" json:"http_max_pending_requests,omitempty"`
	// Maximum number of requests to a backend. Default 1024
	HttpMaxRequests int32 `protobuf:"varint,3,opt,name=http_max_requests,json=httpMaxRequests" json:"http_max_requests,omitempty"`
	// Minimum time the circuit will be closed. format: 1h/1m/1s/1ms. MUST
	// BE >=1ms. Default is 30s.
	SleepWindow *google_protobuf1.Duration `protobuf:"bytes,4,opt,name=sleep_window,json=sleepWindow" json:"sleep_window,omitempty"`
	// Number of 5XX errors before circuit is opened. Defaults to 5.
	HttpConsecutiveErrors int32 `protobuf:"varint,5,opt,name=http_consecutive_errors,json=httpConsecutiveErrors" json:"http_consecutive_errors,omitempty"`
	// Time interval between ejection sweep analysis. format:
	// 1h/1m/1s/1ms. MUST BE >=1ms. Default is 10s.
	HttpDetectionInterval *google_protobuf1.Duration `protobuf:"bytes,6,opt,name=http_detection_interval,json=httpDetectionInterval" json:"http_detection_interval,omitempty"`
	// Maximum number of requests per connection to a backend. Setting this
	// parameter to 1 disables keep alive.
	HttpMaxRequestsPerConnection int32 `protobuf:"varint,7,opt,name=http_max_requests_per_connection,json=httpMaxRequestsPerConnection" json:"http_max_requests_per_connection,omitempty"`
	// Maximum % of hosts in the load balancing pool for the destination
	// service that can be ejected by the circuit breaker. Defaults to
	// 10%.
	HttpMaxEjectionPercent int32 `protobuf:"varint,8,opt,name=http_max_ejection_percent,json=httpMaxEjectionPercent" json:"http_max_ejection_percent,omitempty"`
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) Reset() {
	*m = CircuitBreaker_SimpleCircuitBreakerPolicy{}
}
func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) String() string { return proto.CompactTextString(m) }
func (*CircuitBreaker_SimpleCircuitBreakerPolicy) ProtoMessage()    {}
func (*CircuitBreaker_SimpleCircuitBreakerPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetMaxConnections() int32 {
	if m != nil {
		return m.MaxConnections
	}
	return 0
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetHttpMaxPendingRequests() int32 {
	if m != nil {
		return m.HttpMaxPendingRequests
	}
	return 0
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetHttpMaxRequests() int32 {
	if m != nil {
		return m.HttpMaxRequests
	}
	return 0
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetSleepWindow() *google_protobuf1.Duration {
	if m != nil {
		return m.SleepWindow
	}
	return nil
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetHttpConsecutiveErrors() int32 {
	if m != nil {
		return m.HttpConsecutiveErrors
	}
	return 0
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetHttpDetectionInterval() *google_protobuf1.Duration {
	if m != nil {
		return m.HttpDetectionInterval
	}
	return nil
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetHttpMaxRequestsPerConnection() int32 {
	if m != nil {
		return m.HttpMaxRequestsPerConnection
	}
	return 0
}

func (m *CircuitBreaker_SimpleCircuitBreakerPolicy) GetHttpMaxEjectionPercent() int32 {
	if m != nil {
		return m.HttpMaxEjectionPercent
	}
	return 0
}

func init() {
	proto.RegisterType((*DestinationPolicy)(nil), "istio.proxy.v1.config.DestinationPolicy")
	proto.RegisterType((*DestinationVersionPolicy)(nil), "istio.proxy.v1.config.DestinationVersionPolicy")
	proto.RegisterType((*LoadBalancing)(nil), "istio.proxy.v1.config.LoadBalancing")
	proto.RegisterType((*CircuitBreaker)(nil), "istio.proxy.v1.config.CircuitBreaker")
	proto.RegisterType((*CircuitBreaker_SimpleCircuitBreakerPolicy)(nil), "istio.proxy.v1.config.CircuitBreaker.SimpleCircuitBreakerPolicy")
	proto.RegisterEnum("istio.proxy.v1.config.LoadBalancing_SimpleLBPolicy", LoadBalancing_SimpleLBPolicy_name, LoadBalancing_SimpleLBPolicy_value)
}

func init() { proto.RegisterFile("proxy/v1/config/dest_policy.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 689 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5d, 0x6f, 0x12, 0x4d,
	0x14, 0x2e, 0x50, 0x78, 0xcb, 0xe1, 0x2d, 0xd0, 0x49, 0xab, 0x5b, 0x62, 0x0c, 0x12, 0x8d, 0x8d,
	0x31, 0x4b, 0x4a, 0x13, 0xb5, 0x46, 0x13, 0xcb, 0x87, 0xb6, 0xb1, 0x05, 0x9c, 0x56, 0xbd, 0xdc,
	0x0c, 0xbb, 0x53, 0x1c, 0xbb, 0xcc, 0xac, 0x33, 0x03, 0x85, 0x1b, 0xff, 0x8f, 0x7f, 0xc8, 0x7f,
	0xe1, 0xbd, 0x97, 0x86, 0xd9, 0xe5, 0xab, 0x95, 0xd8, 0xde, 0x31, 0xe7, 0xf9, 0x38, 0xe7, 0x39,
	0x33, 0x2c, 0x3c, 0x08, 0xa4, 0x18, 0x8e, 0xca, 0x83, 0xdd, 0xb2, 0x2b, 0xf8, 0x39, 0xeb, 0x96,
	0x3d, 0xaa, 0xb4, 0x13, 0x08, 0x9f, 0xb9, 0x23, 0x3b, 0x90, 0x42, 0x0b, 0xb4, 0xc5, 0x94, 0x66,
	0xc2, 0x36, 0x44, 0x7b, 0xb0, 0x6b, 0x87, 0xc4, 0xc2, 0x76, 0x57, 0x88, 0xae, 0x4f, 0xcb, 0x86,
	0xd4, 0xe9, 0x9f, 0x97, 0x09, 0x8f, 0x14, 0x85, 0xfb, 0x57, 0x21, 0xaf, 0x2f, 0x89, 0x66, 0x82,
	0x87, 0x78, 0xe9, 0x3b, 0x6c, 0xd4, 0xa9, 0xd2, 0x8c, 0x9b, 0x62, 0xdb, 0x34, 0x43, 0x45, 0xc8,
	0x78, 0xb3, 0xa2, 0x15, 0x2b, 0xc6, 0x76, 0xd2, 0x78, 0xbe, 0x84, 0xde, 0x41, 0x2a, 0x1c, 0xcc,
	0x8a, 0x17, 0x13, 0x3b, 0x99, 0x4a, 0xd9, 0xfe, 0xeb, 0x64, 0xf6, 0x9c, 0xf7, 0x27, 0x2a, 0xd5,
	0xb4, 0x05, 0x8e, 0xe4, 0xa5, 0x5f, 0x71, 0xb0, 0x96, 0x91, 0xd0, 0x09, 0xac, 0x6a, 0xd2, 0x55,
	0x56, 0xcc, 0xf4, 0xd8, 0xbf, 0x65, 0x0f, 0xfb, 0x8c, 0x74, 0x55, 0x83, 0x6b, 0x39, 0xc2, 0xc6,
	0x06, 0xbd, 0x87, 0xac, 0x2f, 0x88, 0xe7, 0x74, 0x88, 0x4f, 0xb8, 0xcb, 0x78, 0xd7, 0x8a, 0x17,
	0x63, 0x3b, 0x99, 0xca, 0xc3, 0x25, 0xc6, 0xc7, 0x82, 0x78, 0xd5, 0x09, 0x17, 0xaf, 0xfb, 0xf3,
	0x47, 0xd4, 0x84, 0x9c, 0xcb, 0xa4, 0xdb, 0x67, 0xda, 0xe9, 0x48, 0x4a, 0x2e, 0xa8, 0xb4, 0x12,
	0xc6, 0xed, 0xd1, 0x12, 0xb7, 0x5a, 0xc8, 0xae, 0x86, 0x64, 0x9c, 0x75, 0x17, 0xce, 0xe8, 0x29,
	0xa4, 0xdc, 0xbe, 0xd2, 0xa2, 0x67, 0x79, 0xc6, 0x66, 0xd3, 0x0e, 0x6f, 0xce, 0x9e, 0xdc, 0x9c,
	0x7d, 0xc0, 0x47, 0x38, 0xe2, 0x14, 0x9e, 0x43, 0x7a, 0x9a, 0x0e, 0xe5, 0x21, 0x71, 0x41, 0x47,
	0xd1, 0x35, 0x8d, 0x7f, 0xa2, 0x4d, 0x48, 0x0e, 0x88, 0xdf, 0xa7, 0x26, 0x60, 0x1a, 0x87, 0x87,
	0x97, 0xf1, 0x17, 0xb1, 0xd2, 0xcf, 0x18, 0xac, 0x2f, 0xe4, 0x42, 0x47, 0xb0, 0xca, 0x49, 0x8f,
	0x1a, 0x79, 0xb6, 0xb2, 0x77, 0x93, 0x5d, 0xd8, 0xa7, 0xac, 0x17, 0xf8, 0xf4, 0xb8, 0x1a, 0x2e,
	0xfa, 0x70, 0x05, 0x1b, 0x0b, 0x64, 0x4f, 0x33, 0xc4, 0x97, 0x67, 0x38, 0x5c, 0x99, 0xa4, 0x28,
	0xbd, 0x86, 0xec, 0xa2, 0x13, 0xca, 0x41, 0x06, 0xb7, 0x3e, 0x36, 0xeb, 0x0e, 0x6e, 0x55, 0x8f,
	0x9a, 0xf9, 0x15, 0x94, 0x05, 0x38, 0x6e, 0x1c, 0x9c, 0x9e, 0x39, 0xb5, 0x56, 0xb3, 0x99, 0x8f,
	0x21, 0x80, 0x14, 0x3e, 0x68, 0xd6, 0x5b, 0x27, 0xf9, 0x78, 0x35, 0x03, 0x69, 0xbf, 0x13, 0xfd,
	0x41, 0x4a, 0x3f, 0x92, 0x90, 0x5d, 0x5c, 0x31, 0x72, 0x20, 0xad, 0x8c, 0xbd, 0xe3, 0x76, 0x4c,
	0xbc, 0x4c, 0xe5, 0xcd, 0x8d, 0x2e, 0x27, 0xca, 0xb7, 0x58, 0x9c, 0x66, 0x5d, 0x0b, 0x4d, 0x6b,
	0x9d, 0xdb, 0xe6, 0x2d, 0xfc, 0x4e, 0x40, 0x61, 0xb9, 0x35, 0x7a, 0x0c, 0xb9, 0x1e, 0x19, 0x3a,
	0xae, 0xe0, 0x9c, 0xba, 0xe3, 0xe7, 0xac, 0xcc, 0xd4, 0x49, 0x9c, 0xed, 0x91, 0x61, 0x6d, 0x56,
	0x45, 0xfb, 0xb0, 0xfd, 0x45, 0xeb, 0xc0, 0x19, 0xb3, 0x03, 0xca, 0x3d, 0xc6, 0xbb, 0x8e, 0xa4,
	0xdf, 0xfa, 0x54, 0x69, 0x65, 0x46, 0x49, 0xe2, 0x3b, 0x63, 0xc2, 0x09, 0x19, 0xb6, 0x43, 0x18,
	0x47, 0x28, 0x7a, 0x02, 0x1b, 0x53, 0xe9, 0x54, 0x92, 0x30, 0x92, 0x5c, 0x24, 0x99, 0x72, 0x5f,
	0xc1, 0xff, 0xca, 0xa7, 0x34, 0x70, 0x2e, 0x19, 0xf7, 0xc4, 0xa5, 0xb5, 0x6a, 0x42, 0x6e, 0x5f,
	0x0b, 0x59, 0x8f, 0x3e, 0x29, 0x38, 0x63, 0xe8, 0x9f, 0x0d, 0x1b, 0x3d, 0x83, 0xbb, 0xa6, 0x93,
	0x2b, 0xb8, 0xa2, 0x6e, 0x5f, 0xb3, 0x01, 0x75, 0xa8, 0x94, 0x42, 0x2a, 0x2b, 0x69, 0xfa, 0x6d,
	0x8d, 0xe1, 0xda, 0x0c, 0x6d, 0x18, 0x10, 0x7d, 0x88, 0x74, 0x1e, 0xd5, 0x61, 0x5e, 0x87, 0x71,
	0x4d, 0xe5, 0x80, 0xf8, 0x56, 0xea, 0x5f, 0x03, 0x18, 0xcb, 0xfa, 0x44, 0x78, 0x14, 0xe9, 0xd0,
	0x5b, 0x28, 0x5e, 0x0b, 0xed, 0x04, 0x54, 0xce, 0xad, 0xda, 0xfa, 0xcf, 0xcc, 0x74, 0xef, 0xca,
	0x0e, 0xda, 0x54, 0xce, 0x16, 0xbf, 0xb0, 0x77, 0xfa, 0x35, 0x9a, 0x2e, 0xa0, 0xd2, 0xa5, 0x5c,
	0x5b, 0x6b, 0x0b, 0x7b, 0x6f, 0x44, 0x70, 0x3b, 0x44, 0xc7, 0x6f, 0xd5, 0x9d, 0xbc, 0xd5, 0x4e,
	0xca, 0x4c, 0xbe, 0xf7, 0x27, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x49, 0xd2, 0x57, 0xf2, 0x05, 0x00,
	0x00,
}
