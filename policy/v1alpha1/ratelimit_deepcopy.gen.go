// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: policy/v1alpha1/ratelimit.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "istio.io/api/type/v1beta1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using RateLimit within kubernetes types, where deepcopy-gen is used.
func (in *RateLimit) DeepCopyInto(out *RateLimit) {
	p := proto.Clone(in).(*RateLimit)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimit. Required by controller-gen.
func (in *RateLimit) DeepCopy() *RateLimit {
	if in == nil {
		return nil
	}
	out := new(RateLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RateLimit. Required by controller-gen.
func (in *RateLimit) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RouteSelector within kubernetes types, where deepcopy-gen is used.
func (in *RouteSelector) DeepCopyInto(out *RouteSelector) {
	p := proto.Clone(in).(*RouteSelector)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteSelector. Required by controller-gen.
func (in *RouteSelector) DeepCopy() *RouteSelector {
	if in == nil {
		return nil
	}
	out := new(RouteSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RouteSelector. Required by controller-gen.
func (in *RouteSelector) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RateLimitDescriptor within kubernetes types, where deepcopy-gen is used.
func (in *RateLimitDescriptor) DeepCopyInto(out *RateLimitDescriptor) {
	p := proto.Clone(in).(*RateLimitDescriptor)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitDescriptor. Required by controller-gen.
func (in *RateLimitDescriptor) DeepCopy() *RateLimitDescriptor {
	if in == nil {
		return nil
	}
	out := new(RateLimitDescriptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitDescriptor. Required by controller-gen.
func (in *RateLimitDescriptor) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RateLimitEntry within kubernetes types, where deepcopy-gen is used.
func (in *RateLimitEntry) DeepCopyInto(out *RateLimitEntry) {
	p := proto.Clone(in).(*RateLimitEntry)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitEntry. Required by controller-gen.
func (in *RateLimitEntry) DeepCopy() *RateLimitEntry {
	if in == nil {
		return nil
	}
	out := new(RateLimitEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitEntry. Required by controller-gen.
func (in *RateLimitEntry) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RateLimitEntry_RequestHeader within kubernetes types, where deepcopy-gen is used.
func (in *RateLimitEntry_RequestHeader) DeepCopyInto(out *RateLimitEntry_RequestHeader) {
	p := proto.Clone(in).(*RateLimitEntry_RequestHeader)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitEntry_RequestHeader. Required by controller-gen.
func (in *RateLimitEntry_RequestHeader) DeepCopy() *RateLimitEntry_RequestHeader {
	if in == nil {
		return nil
	}
	out := new(RateLimitEntry_RequestHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitEntry_RequestHeader. Required by controller-gen.
func (in *RateLimitEntry_RequestHeader) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RateLimitEntry_RequestHeaderMatch within kubernetes types, where deepcopy-gen is used.
func (in *RateLimitEntry_RequestHeaderMatch) DeepCopyInto(out *RateLimitEntry_RequestHeaderMatch) {
	p := proto.Clone(in).(*RateLimitEntry_RequestHeaderMatch)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitEntry_RequestHeaderMatch. Required by controller-gen.
func (in *RateLimitEntry_RequestHeaderMatch) DeepCopy() *RateLimitEntry_RequestHeaderMatch {
	if in == nil {
		return nil
	}
	out := new(RateLimitEntry_RequestHeaderMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitEntry_RequestHeaderMatch. Required by controller-gen.
func (in *RateLimitEntry_RequestHeaderMatch) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RateLimitEntry_RequestHeaderMatch_HeaderMatchCondition within kubernetes types, where deepcopy-gen is used.
func (in *RateLimitEntry_RequestHeaderMatch_HeaderMatchCondition) DeepCopyInto(out *RateLimitEntry_RequestHeaderMatch_HeaderMatchCondition) {
	p := proto.Clone(in).(*RateLimitEntry_RequestHeaderMatch_HeaderMatchCondition)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitEntry_RequestHeaderMatch_HeaderMatchCondition. Required by controller-gen.
func (in *RateLimitEntry_RequestHeaderMatch_HeaderMatchCondition) DeepCopy() *RateLimitEntry_RequestHeaderMatch_HeaderMatchCondition {
	if in == nil {
		return nil
	}
	out := new(RateLimitEntry_RequestHeaderMatch_HeaderMatchCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitEntry_RequestHeaderMatch_HeaderMatchCondition. Required by controller-gen.
func (in *RateLimitEntry_RequestHeaderMatch_HeaderMatchCondition) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RateLimitEntry_ConstantEntry within kubernetes types, where deepcopy-gen is used.
func (in *RateLimitEntry_ConstantEntry) DeepCopyInto(out *RateLimitEntry_ConstantEntry) {
	p := proto.Clone(in).(*RateLimitEntry_ConstantEntry)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitEntry_ConstantEntry. Required by controller-gen.
func (in *RateLimitEntry_ConstantEntry) DeepCopy() *RateLimitEntry_ConstantEntry {
	if in == nil {
		return nil
	}
	out := new(RateLimitEntry_ConstantEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitEntry_ConstantEntry. Required by controller-gen.
func (in *RateLimitEntry_ConstantEntry) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using RateLimitEntry_Expression within kubernetes types, where deepcopy-gen is used.
func (in *RateLimitEntry_Expression) DeepCopyInto(out *RateLimitEntry_Expression) {
	p := proto.Clone(in).(*RateLimitEntry_Expression)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitEntry_Expression. Required by controller-gen.
func (in *RateLimitEntry_Expression) DeepCopy() *RateLimitEntry_Expression {
	if in == nil {
		return nil
	}
	out := new(RateLimitEntry_Expression)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitEntry_Expression. Required by controller-gen.
func (in *RateLimitEntry_Expression) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ProviderRef within kubernetes types, where deepcopy-gen is used.
func (in *ProviderRef) DeepCopyInto(out *ProviderRef) {
	p := proto.Clone(in).(*ProviderRef)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderRef. Required by controller-gen.
func (in *ProviderRef) DeepCopy() *ProviderRef {
	if in == nil {
		return nil
	}
	out := new(ProviderRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ProviderRef. Required by controller-gen.
func (in *ProviderRef) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LocalRateLimit within kubernetes types, where deepcopy-gen is used.
func (in *LocalRateLimit) DeepCopyInto(out *LocalRateLimit) {
	p := proto.Clone(in).(*LocalRateLimit)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalRateLimit. Required by controller-gen.
func (in *LocalRateLimit) DeepCopy() *LocalRateLimit {
	if in == nil {
		return nil
	}
	out := new(LocalRateLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LocalRateLimit. Required by controller-gen.
func (in *LocalRateLimit) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LocalRateLimit_TokenBucket within kubernetes types, where deepcopy-gen is used.
func (in *LocalRateLimit_TokenBucket) DeepCopyInto(out *LocalRateLimit_TokenBucket) {
	p := proto.Clone(in).(*LocalRateLimit_TokenBucket)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalRateLimit_TokenBucket. Required by controller-gen.
func (in *LocalRateLimit_TokenBucket) DeepCopy() *LocalRateLimit_TokenBucket {
	if in == nil {
		return nil
	}
	out := new(LocalRateLimit_TokenBucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LocalRateLimit_TokenBucket. Required by controller-gen.
func (in *LocalRateLimit_TokenBucket) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LocalRateLimit_LocalRateLimitDescriptor within kubernetes types, where deepcopy-gen is used.
func (in *LocalRateLimit_LocalRateLimitDescriptor) DeepCopyInto(out *LocalRateLimit_LocalRateLimitDescriptor) {
	p := proto.Clone(in).(*LocalRateLimit_LocalRateLimitDescriptor)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalRateLimit_LocalRateLimitDescriptor. Required by controller-gen.
func (in *LocalRateLimit_LocalRateLimitDescriptor) DeepCopy() *LocalRateLimit_LocalRateLimitDescriptor {
	if in == nil {
		return nil
	}
	out := new(LocalRateLimit_LocalRateLimitDescriptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LocalRateLimit_LocalRateLimitDescriptor. Required by controller-gen.
func (in *LocalRateLimit_LocalRateLimitDescriptor) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LocalRateLimit_LocalRateLimitDescriptor_DescriptorEntry within kubernetes types, where deepcopy-gen is used.
func (in *LocalRateLimit_LocalRateLimitDescriptor_DescriptorEntry) DeepCopyInto(out *LocalRateLimit_LocalRateLimitDescriptor_DescriptorEntry) {
	p := proto.Clone(in).(*LocalRateLimit_LocalRateLimitDescriptor_DescriptorEntry)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalRateLimit_LocalRateLimitDescriptor_DescriptorEntry. Required by controller-gen.
func (in *LocalRateLimit_LocalRateLimitDescriptor_DescriptorEntry) DeepCopy() *LocalRateLimit_LocalRateLimitDescriptor_DescriptorEntry {
	if in == nil {
		return nil
	}
	out := new(LocalRateLimit_LocalRateLimitDescriptor_DescriptorEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LocalRateLimit_LocalRateLimitDescriptor_DescriptorEntry. Required by controller-gen.
func (in *LocalRateLimit_LocalRateLimitDescriptor_DescriptorEntry) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}