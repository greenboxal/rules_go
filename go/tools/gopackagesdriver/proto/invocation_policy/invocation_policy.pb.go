// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/main/protobuf/invocation_policy.proto

package blaze_invocation_policy

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The --invocation_policy flag takes a base64-encoded binary-serialized or text
// formatted InvocationPolicy message.
type InvocationPolicy struct {
	// Order matters.
	// After expanding policies on expansion flags or flags with implicit
	// requirements, only the final policy on a specific flag will be enforced
	// onto the user's command line.
	FlagPolicies         []*FlagPolicy `protobuf:"bytes,1,rep,name=flag_policies,json=flagPolicies" json:"flag_policies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *InvocationPolicy) Reset()         { *m = InvocationPolicy{} }
func (m *InvocationPolicy) String() string { return proto.CompactTextString(m) }
func (*InvocationPolicy) ProtoMessage()    {}
func (*InvocationPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_e130ca0062eb2b66, []int{0}
}

func (m *InvocationPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvocationPolicy.Unmarshal(m, b)
}
func (m *InvocationPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvocationPolicy.Marshal(b, m, deterministic)
}
func (m *InvocationPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvocationPolicy.Merge(m, src)
}
func (m *InvocationPolicy) XXX_Size() int {
	return xxx_messageInfo_InvocationPolicy.Size(m)
}
func (m *InvocationPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_InvocationPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_InvocationPolicy proto.InternalMessageInfo

func (m *InvocationPolicy) GetFlagPolicies() []*FlagPolicy {
	if m != nil {
		return m.FlagPolicies
	}
	return nil
}

// A policy for controlling the value of a flag.
type FlagPolicy struct {
	// The name of the flag to enforce this policy on.
	//
	// Note that this should be the full name of the flag, not the abbreviated
	// name of the flag. If the user specifies the abbreviated name of a flag,
	// that flag will be matched using its full name.
	//
	// The "no" prefix will not be parsed, so for boolean flags, use
	// the flag's full name and explicitly set it to true or false.
	FlagName *string `protobuf:"bytes,1,opt,name=flag_name,json=flagName" json:"flag_name,omitempty"`
	// If set, this flag policy is applied only if one of the given commands or a
	// command that inherits from one of the given commands is being run. For
	// instance, if "build" is one of the commands here, then this policy will
	// apply to any command that inherits from build, such as info, coverage, or
	// test. If empty, this flag policy is applied for all commands. This allows
	// the policy setter to add all policies to the proto without having to
	// determine which Bazel command the user is actually running. Additionally,
	// Bazel allows multiple flags to be defined by the same name, and the
	// specific flag definition is determined by the command.
	Commands []string `protobuf:"bytes,2,rep,name=commands" json:"commands,omitempty"`
	// Types that are valid to be assigned to Operation:
	//	*FlagPolicy_SetValue
	//	*FlagPolicy_UseDefault
	//	*FlagPolicy_DisallowValues
	//	*FlagPolicy_AllowValues
	Operation            isFlagPolicy_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *FlagPolicy) Reset()         { *m = FlagPolicy{} }
func (m *FlagPolicy) String() string { return proto.CompactTextString(m) }
func (*FlagPolicy) ProtoMessage()    {}
func (*FlagPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_e130ca0062eb2b66, []int{1}
}

func (m *FlagPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlagPolicy.Unmarshal(m, b)
}
func (m *FlagPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlagPolicy.Marshal(b, m, deterministic)
}
func (m *FlagPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlagPolicy.Merge(m, src)
}
func (m *FlagPolicy) XXX_Size() int {
	return xxx_messageInfo_FlagPolicy.Size(m)
}
func (m *FlagPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_FlagPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_FlagPolicy proto.InternalMessageInfo

func (m *FlagPolicy) GetFlagName() string {
	if m != nil && m.FlagName != nil {
		return *m.FlagName
	}
	return ""
}

func (m *FlagPolicy) GetCommands() []string {
	if m != nil {
		return m.Commands
	}
	return nil
}

type isFlagPolicy_Operation interface {
	isFlagPolicy_Operation()
}

type FlagPolicy_SetValue struct {
	SetValue *SetValue `protobuf:"bytes,3,opt,name=set_value,json=setValue,oneof"`
}

type FlagPolicy_UseDefault struct {
	UseDefault *UseDefault `protobuf:"bytes,4,opt,name=use_default,json=useDefault,oneof"`
}

type FlagPolicy_DisallowValues struct {
	DisallowValues *DisallowValues `protobuf:"bytes,5,opt,name=disallow_values,json=disallowValues,oneof"`
}

type FlagPolicy_AllowValues struct {
	AllowValues *AllowValues `protobuf:"bytes,6,opt,name=allow_values,json=allowValues,oneof"`
}

func (*FlagPolicy_SetValue) isFlagPolicy_Operation() {}

func (*FlagPolicy_UseDefault) isFlagPolicy_Operation() {}

func (*FlagPolicy_DisallowValues) isFlagPolicy_Operation() {}

func (*FlagPolicy_AllowValues) isFlagPolicy_Operation() {}

func (m *FlagPolicy) GetOperation() isFlagPolicy_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *FlagPolicy) GetSetValue() *SetValue {
	if x, ok := m.GetOperation().(*FlagPolicy_SetValue); ok {
		return x.SetValue
	}
	return nil
}

func (m *FlagPolicy) GetUseDefault() *UseDefault {
	if x, ok := m.GetOperation().(*FlagPolicy_UseDefault); ok {
		return x.UseDefault
	}
	return nil
}

func (m *FlagPolicy) GetDisallowValues() *DisallowValues {
	if x, ok := m.GetOperation().(*FlagPolicy_DisallowValues); ok {
		return x.DisallowValues
	}
	return nil
}

func (m *FlagPolicy) GetAllowValues() *AllowValues {
	if x, ok := m.GetOperation().(*FlagPolicy_AllowValues); ok {
		return x.AllowValues
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FlagPolicy) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FlagPolicy_SetValue)(nil),
		(*FlagPolicy_UseDefault)(nil),
		(*FlagPolicy_DisallowValues)(nil),
		(*FlagPolicy_AllowValues)(nil),
	}
}

type SetValue struct {
	// Use this value for the specified flag, overriding any default or user-set
	// value (unless append is set to true for repeatable flags).
	//
	// This field is repeated for repeatable flags. It is an error to set
	// multiple values for a flag that is not actually a repeatable flag.
	// This requires at least 1 value, if even the empty string.
	//
	// If the flag allows multiple values, all of its values are replaced with the
	// value or values from the policy (i.e., no diffing or merging is performed),
	// unless the append field (see below) is set to true.
	//
	// Note that some flags are tricky. For example, some flags look like boolean
	// flags, but are actually Void expansion flags that expand into other flags.
	// The Bazel flag parser will accept "--void_flag=false", but because
	// the flag is Void, the "=false" is ignored. It can get even trickier, like
	// "--novoid_flag" which is also an expansion flag with the type Void whose
	// name is explicitly "novoid_flag" and which expands into other flags that
	// are the opposite of "--void_flag". For expansion flags, it's best to
	// explicitly override the flags they expand into.
	//
	// Other flags may be differently tricky: A flag could have a converter that
	// converts some string to a list of values, but that flag may not itself have
	// allowMultiple set to true.
	//
	// An example is "--test_tag_filters": this flag sets its converter to
	// CommaSeparatedOptionListConverter, but does not set allowMultiple to true.
	// So "--test_tag_filters=foo,bar" results in ["foo", "bar"], however
	// "--test_tag_filters=foo --test_tag_filters=bar" results in just ["bar"]
	// since the 2nd value overrides the 1st.
	//
	// Similarly, "--test_tag_filters=foo,bar --test_tag_filters=baz,qux" results
	// in ["baz", "qux"]. For flags like these, the policy should specify
	// "foo,bar" instead of separately specifying "foo" and "bar" so that the
	// converter is appropriately invoked.
	//
	// Note that the opposite is not necessarily
	// true: for a flag that specifies allowMultiple=true, "--flag=foo,bar"
	// may fail to parse or result in an unexpected value.
	FlagValue []string `protobuf:"bytes,1,rep,name=flag_value,json=flagValue" json:"flag_value,omitempty"`
	// Whether to allow this policy to be overridden by user-specified values.
	// When set, if the user specified a value for this flag, use the value
	// from the user, otherwise use the value specified in this policy.
	Overridable *bool `protobuf:"varint,2,opt,name=overridable" json:"overridable,omitempty"`
	// If true, and if the flag named in the policy is a repeatable flag, then
	// the values listed in flag_value do not replace all the user-set or default
	// values of the flag, but instead append to them. If the flag is not
	// repeatable, then this has no effect.
	Append               *bool    `protobuf:"varint,3,opt,name=append" json:"append,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetValue) Reset()         { *m = SetValue{} }
func (m *SetValue) String() string { return proto.CompactTextString(m) }
func (*SetValue) ProtoMessage()    {}
func (*SetValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_e130ca0062eb2b66, []int{2}
}

func (m *SetValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetValue.Unmarshal(m, b)
}
func (m *SetValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetValue.Marshal(b, m, deterministic)
}
func (m *SetValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetValue.Merge(m, src)
}
func (m *SetValue) XXX_Size() int {
	return xxx_messageInfo_SetValue.Size(m)
}
func (m *SetValue) XXX_DiscardUnknown() {
	xxx_messageInfo_SetValue.DiscardUnknown(m)
}

var xxx_messageInfo_SetValue proto.InternalMessageInfo

func (m *SetValue) GetFlagValue() []string {
	if m != nil {
		return m.FlagValue
	}
	return nil
}

func (m *SetValue) GetOverridable() bool {
	if m != nil && m.Overridable != nil {
		return *m.Overridable
	}
	return false
}

func (m *SetValue) GetAppend() bool {
	if m != nil && m.Append != nil {
		return *m.Append
	}
	return false
}

type UseDefault struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UseDefault) Reset()         { *m = UseDefault{} }
func (m *UseDefault) String() string { return proto.CompactTextString(m) }
func (*UseDefault) ProtoMessage()    {}
func (*UseDefault) Descriptor() ([]byte, []int) {
	return fileDescriptor_e130ca0062eb2b66, []int{3}
}

func (m *UseDefault) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UseDefault.Unmarshal(m, b)
}
func (m *UseDefault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UseDefault.Marshal(b, m, deterministic)
}
func (m *UseDefault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UseDefault.Merge(m, src)
}
func (m *UseDefault) XXX_Size() int {
	return xxx_messageInfo_UseDefault.Size(m)
}
func (m *UseDefault) XXX_DiscardUnknown() {
	xxx_messageInfo_UseDefault.DiscardUnknown(m)
}

var xxx_messageInfo_UseDefault proto.InternalMessageInfo

type DisallowValues struct {
	// It is an error for the user to use any of these values (that is, the Bazel
	// command will fail), unless new_value or use_default is set.
	//
	// For repeatable flags, if any one of the values in the flag matches a value
	// in the list of disallowed values, an error is thrown.
	//
	// Care must be taken for flags with complicated converters. For example,
	// it's possible for a repeated flag to be of type List<List<T>>, so that
	// "--foo=a,b --foo=c,d" results in foo=[["a","b"], ["c", "d"]]. In this case,
	// it is not possible to disallow just "b", nor will ["b", "a"] match, nor
	// will ["b", "c"] (but ["a", "b"] will still match).
	DisallowedValues []string `protobuf:"bytes,1,rep,name=disallowed_values,json=disallowedValues" json:"disallowed_values,omitempty"`
	// Types that are valid to be assigned to ReplacementValue:
	//	*DisallowValues_NewValue
	//	*DisallowValues_UseDefault
	ReplacementValue     isDisallowValues_ReplacementValue `protobuf_oneof:"replacement_value"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *DisallowValues) Reset()         { *m = DisallowValues{} }
func (m *DisallowValues) String() string { return proto.CompactTextString(m) }
func (*DisallowValues) ProtoMessage()    {}
func (*DisallowValues) Descriptor() ([]byte, []int) {
	return fileDescriptor_e130ca0062eb2b66, []int{4}
}

func (m *DisallowValues) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisallowValues.Unmarshal(m, b)
}
func (m *DisallowValues) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisallowValues.Marshal(b, m, deterministic)
}
func (m *DisallowValues) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisallowValues.Merge(m, src)
}
func (m *DisallowValues) XXX_Size() int {
	return xxx_messageInfo_DisallowValues.Size(m)
}
func (m *DisallowValues) XXX_DiscardUnknown() {
	xxx_messageInfo_DisallowValues.DiscardUnknown(m)
}

var xxx_messageInfo_DisallowValues proto.InternalMessageInfo

func (m *DisallowValues) GetDisallowedValues() []string {
	if m != nil {
		return m.DisallowedValues
	}
	return nil
}

type isDisallowValues_ReplacementValue interface {
	isDisallowValues_ReplacementValue()
}

type DisallowValues_NewValue struct {
	NewValue string `protobuf:"bytes,3,opt,name=new_value,json=newValue,oneof"`
}

type DisallowValues_UseDefault struct {
	UseDefault *UseDefault `protobuf:"bytes,4,opt,name=use_default,json=useDefault,oneof"`
}

func (*DisallowValues_NewValue) isDisallowValues_ReplacementValue() {}

func (*DisallowValues_UseDefault) isDisallowValues_ReplacementValue() {}

func (m *DisallowValues) GetReplacementValue() isDisallowValues_ReplacementValue {
	if m != nil {
		return m.ReplacementValue
	}
	return nil
}

func (m *DisallowValues) GetNewValue() string {
	if x, ok := m.GetReplacementValue().(*DisallowValues_NewValue); ok {
		return x.NewValue
	}
	return ""
}

func (m *DisallowValues) GetUseDefault() *UseDefault {
	if x, ok := m.GetReplacementValue().(*DisallowValues_UseDefault); ok {
		return x.UseDefault
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DisallowValues) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DisallowValues_NewValue)(nil),
		(*DisallowValues_UseDefault)(nil),
	}
}

type AllowValues struct {
	// It is an error for the user to use any value not in this list, unless
	// new_value or use_default is set.
	AllowedValues []string `protobuf:"bytes,1,rep,name=allowed_values,json=allowedValues" json:"allowed_values,omitempty"`
	// Types that are valid to be assigned to ReplacementValue:
	//	*AllowValues_NewValue
	//	*AllowValues_UseDefault
	ReplacementValue     isAllowValues_ReplacementValue `protobuf_oneof:"replacement_value"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *AllowValues) Reset()         { *m = AllowValues{} }
func (m *AllowValues) String() string { return proto.CompactTextString(m) }
func (*AllowValues) ProtoMessage()    {}
func (*AllowValues) Descriptor() ([]byte, []int) {
	return fileDescriptor_e130ca0062eb2b66, []int{5}
}

func (m *AllowValues) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllowValues.Unmarshal(m, b)
}
func (m *AllowValues) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllowValues.Marshal(b, m, deterministic)
}
func (m *AllowValues) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowValues.Merge(m, src)
}
func (m *AllowValues) XXX_Size() int {
	return xxx_messageInfo_AllowValues.Size(m)
}
func (m *AllowValues) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowValues.DiscardUnknown(m)
}

var xxx_messageInfo_AllowValues proto.InternalMessageInfo

func (m *AllowValues) GetAllowedValues() []string {
	if m != nil {
		return m.AllowedValues
	}
	return nil
}

type isAllowValues_ReplacementValue interface {
	isAllowValues_ReplacementValue()
}

type AllowValues_NewValue struct {
	NewValue string `protobuf:"bytes,3,opt,name=new_value,json=newValue,oneof"`
}

type AllowValues_UseDefault struct {
	UseDefault *UseDefault `protobuf:"bytes,4,opt,name=use_default,json=useDefault,oneof"`
}

func (*AllowValues_NewValue) isAllowValues_ReplacementValue() {}

func (*AllowValues_UseDefault) isAllowValues_ReplacementValue() {}

func (m *AllowValues) GetReplacementValue() isAllowValues_ReplacementValue {
	if m != nil {
		return m.ReplacementValue
	}
	return nil
}

func (m *AllowValues) GetNewValue() string {
	if x, ok := m.GetReplacementValue().(*AllowValues_NewValue); ok {
		return x.NewValue
	}
	return ""
}

func (m *AllowValues) GetUseDefault() *UseDefault {
	if x, ok := m.GetReplacementValue().(*AllowValues_UseDefault); ok {
		return x.UseDefault
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AllowValues) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AllowValues_NewValue)(nil),
		(*AllowValues_UseDefault)(nil),
	}
}

func init() {
	proto.RegisterType((*InvocationPolicy)(nil), "blaze.invocation_policy.InvocationPolicy")
	proto.RegisterType((*FlagPolicy)(nil), "blaze.invocation_policy.FlagPolicy")
	proto.RegisterType((*SetValue)(nil), "blaze.invocation_policy.SetValue")
	proto.RegisterType((*UseDefault)(nil), "blaze.invocation_policy.UseDefault")
	proto.RegisterType((*DisallowValues)(nil), "blaze.invocation_policy.DisallowValues")
	proto.RegisterType((*AllowValues)(nil), "blaze.invocation_policy.AllowValues")
}

func init() {
	proto.RegisterFile("src/main/protobuf/invocation_policy.proto", fileDescriptor_e130ca0062eb2b66)
}

var fileDescriptor_e130ca0062eb2b66 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x6d, 0xda, 0x75, 0x49, 0x6e, 0xba, 0xb5, 0x3b, 0x82, 0x06, 0x65, 0x21, 0x46, 0xc5, 0xca,
	0xe2, 0x14, 0xf6, 0x0b, 0xb4, 0x2c, 0x4b, 0xd7, 0x07, 0x91, 0x88, 0x3e, 0x09, 0x65, 0x92, 0xb9,
	0x2d, 0x03, 0x93, 0x99, 0x90, 0x49, 0x5a, 0xf4, 0xc3, 0x7c, 0xf5, 0xc3, 0x7c, 0x91, 0x4e, 0x92,
	0x6d, 0x83, 0xe6, 0xcd, 0x87, 0x7d, 0x9b, 0x7b, 0xce, 0xe9, 0xe1, 0xdc, 0x73, 0x4b, 0xe0, 0x8d,
	0x29, 0xd2, 0x79, 0xc6, 0x84, 0x9a, 0xe7, 0x85, 0x2e, 0x75, 0x52, 0xad, 0xe7, 0x42, 0x6d, 0x75,
	0xca, 0x4a, 0xa1, 0xd5, 0x2a, 0xd7, 0x52, 0xa4, 0xdf, 0xa9, 0xa5, 0xc8, 0x93, 0x44, 0xb2, 0x1f,
	0x48, 0xff, 0xa2, 0xa3, 0x6f, 0x30, 0xbd, 0xbd, 0x03, 0x3f, 0x59, 0x8c, 0x2c, 0xe1, 0x6c, 0x2d,
	0xd9, 0xa6, 0x96, 0x08, 0x34, 0x81, 0x13, 0x8e, 0x66, 0xfe, 0xd5, 0x0b, 0xda, 0x63, 0x42, 0x6f,
	0x24, 0xdb, 0xd4, 0xbf, 0x8d, 0xc7, 0xeb, 0xf6, 0x2d, 0xd0, 0x44, 0xbf, 0x87, 0x00, 0x07, 0x92,
	0x3c, 0x03, 0xcf, 0x1a, 0x2b, 0x96, 0x61, 0xe0, 0x84, 0xce, 0xcc, 0x8b, 0xdd, 0x3d, 0xf0, 0x91,
	0x65, 0x48, 0x9e, 0x82, 0x9b, 0xea, 0x2c, 0x63, 0x8a, 0x9b, 0x60, 0x18, 0x8e, 0xf6, 0x5c, 0x3b,
	0x93, 0x77, 0xe0, 0x19, 0x2c, 0x57, 0x5b, 0x26, 0x2b, 0x0c, 0x46, 0xa1, 0x33, 0xf3, 0xaf, 0x9e,
	0xf7, 0xa6, 0xf9, 0x8c, 0xe5, 0xd7, 0xbd, 0x70, 0x39, 0x88, 0x5d, 0xd3, 0xbc, 0xc9, 0x0d, 0xf8,
	0x95, 0xc1, 0x15, 0xc7, 0x35, 0xab, 0x64, 0x19, 0x9c, 0x58, 0x8f, 0xfe, 0x8d, 0xbe, 0x18, 0xbc,
	0xae, 0xa5, 0xcb, 0x41, 0x0c, 0xd5, 0xdd, 0x44, 0x62, 0x78, 0xc8, 0x85, 0x61, 0x52, 0xea, 0x5d,
	0x1d, 0xc7, 0x04, 0x0f, 0xac, 0xd7, 0xeb, 0x5e, 0xaf, 0xeb, 0x46, 0x6f, 0x83, 0x98, 0xe5, 0x20,
	0x9e, 0xf0, 0x0e, 0x42, 0x6e, 0x61, 0xdc, 0x31, 0x3c, 0xb5, 0x86, 0x2f, 0x7b, 0x0d, 0xdf, 0x77,
	0xdc, 0xfc, 0x23, 0xab, 0x85, 0x0f, 0x9e, 0xce, 0xb1, 0xb0, 0xf2, 0x28, 0x05, 0xb7, 0xed, 0x82,
	0x5c, 0x00, 0xd8, 0xea, 0xeb, 0x0a, 0x1d, 0xdb, 0xaf, 0x3d, 0x46, 0x4d, 0x87, 0xe0, 0xeb, 0x2d,
	0x16, 0x85, 0xe0, 0x2c, 0x91, 0x18, 0x0c, 0x43, 0x67, 0xe6, 0xc6, 0xc7, 0x10, 0x79, 0x0c, 0xa7,
	0x2c, 0xcf, 0x51, 0x71, 0xdb, 0xbf, 0x1b, 0x37, 0x53, 0x34, 0x06, 0x38, 0x94, 0x15, 0xfd, 0x72,
	0x60, 0xd2, 0xdd, 0x97, 0x5c, 0xc2, 0x79, 0xbb, 0x2f, 0xf2, 0x76, 0xc5, 0x3a, 0xc0, 0xf4, 0x40,
	0x34, 0xe2, 0x0b, 0xf0, 0x14, 0xee, 0x8e, 0x0e, 0xed, 0xed, 0xaf, 0xa8, 0x70, 0xf7, 0x5f, 0xaf,
	0xb8, 0x78, 0x04, 0xe7, 0x05, 0xe6, 0x92, 0xa5, 0x98, 0xa1, 0x6a, 0xfe, 0x57, 0x1f, 0x4e, 0xdc,
	0xe1, 0x74, 0x14, 0xfd, 0x74, 0xc0, 0x3f, 0x2a, 0x98, 0xbc, 0x82, 0xc9, 0x3f, 0xb3, 0x9f, 0xdd,
	0xb3, 0xe0, 0x8b, 0xb7, 0x70, 0x99, 0xea, 0x8c, 0x6e, 0xb4, 0xde, 0x48, 0xa4, 0x1c, 0xb7, 0xa5,
	0xd6, 0xd2, 0xd0, 0xa4, 0x12, 0x92, 0x53, 0x29, 0x12, 0x5a, 0x54, 0xaa, 0x14, 0x19, 0xd6, 0x5f,
	0x84, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x89, 0xdf, 0xe8, 0x3d, 0x04, 0x00, 0x00,
}
