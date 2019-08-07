// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/main/protobuf/command_line.proto

package command_line

import (
	fmt "fmt"
	option_filters "github.com/bazelbuild/rules_go/go/tools/gopackagesdriver/proto/option_filters"
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

// Representation of a Bazel command line.
type CommandLine struct {
	// A title for this command line value, to differentiate it from others.
	// In particular, a single invocation may wish to report both the literal and
	// canonical command lines, and this label would be used to differentiate
	// between both versions. This is a string for flexibility.
	CommandLineLabel string `protobuf:"bytes,1,opt,name=command_line_label,json=commandLineLabel,proto3" json:"command_line_label,omitempty"`
	// A Bazel command line is made of distinct parts. For example,
	//    `bazel --nomaster_bazelrc test --nocache_test_results //foo:aTest`
	// has the executable "bazel", a startup flag, a command "test", a command
	// flag, and a test target. There could be many more flags and targets, or
	// none (`bazel info` for example), but the basic structure is there. The
	// command line should be broken down into these logical sections here.
	Sections             []*CommandLineSection `protobuf:"bytes,2,rep,name=sections,proto3" json:"sections,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CommandLine) Reset()         { *m = CommandLine{} }
func (m *CommandLine) String() string { return proto.CompactTextString(m) }
func (*CommandLine) ProtoMessage()    {}
func (*CommandLine) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6c030aeb09bb06d, []int{0}
}

func (m *CommandLine) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandLine.Unmarshal(m, b)
}
func (m *CommandLine) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandLine.Marshal(b, m, deterministic)
}
func (m *CommandLine) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandLine.Merge(m, src)
}
func (m *CommandLine) XXX_Size() int {
	return xxx_messageInfo_CommandLine.Size(m)
}
func (m *CommandLine) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandLine.DiscardUnknown(m)
}

var xxx_messageInfo_CommandLine proto.InternalMessageInfo

func (m *CommandLine) GetCommandLineLabel() string {
	if m != nil {
		return m.CommandLineLabel
	}
	return ""
}

func (m *CommandLine) GetSections() []*CommandLineSection {
	if m != nil {
		return m.Sections
	}
	return nil
}

// A section of the Bazel command line.
type CommandLineSection struct {
	// The name of this section, such as "startup_option" or "command".
	SectionLabel string `protobuf:"bytes,1,opt,name=section_label,json=sectionLabel,proto3" json:"section_label,omitempty"`
	// Types that are valid to be assigned to SectionType:
	//	*CommandLineSection_ChunkList
	//	*CommandLineSection_OptionList
	SectionType          isCommandLineSection_SectionType `protobuf_oneof:"section_type"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *CommandLineSection) Reset()         { *m = CommandLineSection{} }
func (m *CommandLineSection) String() string { return proto.CompactTextString(m) }
func (*CommandLineSection) ProtoMessage()    {}
func (*CommandLineSection) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6c030aeb09bb06d, []int{1}
}

func (m *CommandLineSection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandLineSection.Unmarshal(m, b)
}
func (m *CommandLineSection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandLineSection.Marshal(b, m, deterministic)
}
func (m *CommandLineSection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandLineSection.Merge(m, src)
}
func (m *CommandLineSection) XXX_Size() int {
	return xxx_messageInfo_CommandLineSection.Size(m)
}
func (m *CommandLineSection) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandLineSection.DiscardUnknown(m)
}

var xxx_messageInfo_CommandLineSection proto.InternalMessageInfo

func (m *CommandLineSection) GetSectionLabel() string {
	if m != nil {
		return m.SectionLabel
	}
	return ""
}

type isCommandLineSection_SectionType interface {
	isCommandLineSection_SectionType()
}

type CommandLineSection_ChunkList struct {
	ChunkList *ChunkList `protobuf:"bytes,2,opt,name=chunk_list,json=chunkList,proto3,oneof"`
}

type CommandLineSection_OptionList struct {
	OptionList *OptionList `protobuf:"bytes,3,opt,name=option_list,json=optionList,proto3,oneof"`
}

func (*CommandLineSection_ChunkList) isCommandLineSection_SectionType() {}

func (*CommandLineSection_OptionList) isCommandLineSection_SectionType() {}

func (m *CommandLineSection) GetSectionType() isCommandLineSection_SectionType {
	if m != nil {
		return m.SectionType
	}
	return nil
}

func (m *CommandLineSection) GetChunkList() *ChunkList {
	if x, ok := m.GetSectionType().(*CommandLineSection_ChunkList); ok {
		return x.ChunkList
	}
	return nil
}

func (m *CommandLineSection) GetOptionList() *OptionList {
	if x, ok := m.GetSectionType().(*CommandLineSection_OptionList); ok {
		return x.OptionList
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CommandLineSection) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CommandLineSection_ChunkList)(nil),
		(*CommandLineSection_OptionList)(nil),
	}
}

// Wrapper to allow a list of strings in the "oneof" section_type.
type ChunkList struct {
	Chunk                []string `protobuf:"bytes,1,rep,name=chunk,proto3" json:"chunk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChunkList) Reset()         { *m = ChunkList{} }
func (m *ChunkList) String() string { return proto.CompactTextString(m) }
func (*ChunkList) ProtoMessage()    {}
func (*ChunkList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6c030aeb09bb06d, []int{2}
}

func (m *ChunkList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChunkList.Unmarshal(m, b)
}
func (m *ChunkList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChunkList.Marshal(b, m, deterministic)
}
func (m *ChunkList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChunkList.Merge(m, src)
}
func (m *ChunkList) XXX_Size() int {
	return xxx_messageInfo_ChunkList.Size(m)
}
func (m *ChunkList) XXX_DiscardUnknown() {
	xxx_messageInfo_ChunkList.DiscardUnknown(m)
}

var xxx_messageInfo_ChunkList proto.InternalMessageInfo

func (m *ChunkList) GetChunk() []string {
	if m != nil {
		return m.Chunk
	}
	return nil
}

// Wrapper to allow a list of options in the "oneof" section_type.
type OptionList struct {
	Option               []*Option `protobuf:"bytes,1,rep,name=option,proto3" json:"option,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *OptionList) Reset()         { *m = OptionList{} }
func (m *OptionList) String() string { return proto.CompactTextString(m) }
func (*OptionList) ProtoMessage()    {}
func (*OptionList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6c030aeb09bb06d, []int{3}
}

func (m *OptionList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OptionList.Unmarshal(m, b)
}
func (m *OptionList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OptionList.Marshal(b, m, deterministic)
}
func (m *OptionList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OptionList.Merge(m, src)
}
func (m *OptionList) XXX_Size() int {
	return xxx_messageInfo_OptionList.Size(m)
}
func (m *OptionList) XXX_DiscardUnknown() {
	xxx_messageInfo_OptionList.DiscardUnknown(m)
}

var xxx_messageInfo_OptionList proto.InternalMessageInfo

func (m *OptionList) GetOption() []*Option {
	if m != nil {
		return m.Option
	}
	return nil
}

// A single command line option.
//
// This represents the option itself, but does not take into account the type of
// option or how the parser interpreted it. If this option is part of a command
// line that represents the actual input that Bazel received, it would, for
// example, include expansion flags as they are. However, if this option
// represents the canonical form of the command line, with the values as Bazel
// understands them, then the expansion flag, which has no value, would not
// appear, and the flags it expands to would.
type Option struct {
	// How the option looks with the option and its value combined. Depending on
	// the purpose of this command line report, this could be the canonical
	// form, or the way that the flag was set.
	//
	// Some examples: this might be `--foo=bar` form, or `--foo bar` with a space;
	// for boolean flags, `--nobaz` is accepted on top of `--baz=false` and other
	// negating values, or for a positive value, the unqualified `--baz` form
	// is also accepted. This could also be a short `-b`, if the flag has an
	// abbreviated form.
	CombinedForm string `protobuf:"bytes,1,opt,name=combined_form,json=combinedForm,proto3" json:"combined_form,omitempty"`
	// The canonical name of the option, without the preceding dashes.
	OptionName string `protobuf:"bytes,2,opt,name=option_name,json=optionName,proto3" json:"option_name,omitempty"`
	// The value of the flag, or unset for flags that do not take values.
	// Especially for boolean flags, this should be in canonical form, the
	// combined_form field above gives room for showing the flag as it was set
	// if that is preferred.
	OptionValue string `protobuf:"bytes,3,opt,name=option_value,json=optionValue,proto3" json:"option_value,omitempty"`
	// This flag's tagged effects. See OptionEffectTag's java documentation for
	// details.
	EffectTags []option_filters.OptionEffectTag `protobuf:"varint,4,rep,packed,name=effect_tags,json=effectTags,proto3,enum=options.OptionEffectTag" json:"effect_tags,omitempty"`
	// Metadata about the flag. See OptionMetadataTag's java documentation for
	// details.
	MetadataTags         []option_filters.OptionMetadataTag `protobuf:"varint,5,rep,packed,name=metadata_tags,json=metadataTags,proto3,enum=options.OptionMetadataTag" json:"metadata_tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *Option) Reset()         { *m = Option{} }
func (m *Option) String() string { return proto.CompactTextString(m) }
func (*Option) ProtoMessage()    {}
func (*Option) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6c030aeb09bb06d, []int{4}
}

func (m *Option) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Option.Unmarshal(m, b)
}
func (m *Option) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Option.Marshal(b, m, deterministic)
}
func (m *Option) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Option.Merge(m, src)
}
func (m *Option) XXX_Size() int {
	return xxx_messageInfo_Option.Size(m)
}
func (m *Option) XXX_DiscardUnknown() {
	xxx_messageInfo_Option.DiscardUnknown(m)
}

var xxx_messageInfo_Option proto.InternalMessageInfo

func (m *Option) GetCombinedForm() string {
	if m != nil {
		return m.CombinedForm
	}
	return ""
}

func (m *Option) GetOptionName() string {
	if m != nil {
		return m.OptionName
	}
	return ""
}

func (m *Option) GetOptionValue() string {
	if m != nil {
		return m.OptionValue
	}
	return ""
}

func (m *Option) GetEffectTags() []option_filters.OptionEffectTag {
	if m != nil {
		return m.EffectTags
	}
	return nil
}

func (m *Option) GetMetadataTags() []option_filters.OptionMetadataTag {
	if m != nil {
		return m.MetadataTags
	}
	return nil
}

func init() {
	proto.RegisterType((*CommandLine)(nil), "command_line.CommandLine")
	proto.RegisterType((*CommandLineSection)(nil), "command_line.CommandLineSection")
	proto.RegisterType((*ChunkList)(nil), "command_line.ChunkList")
	proto.RegisterType((*OptionList)(nil), "command_line.OptionList")
	proto.RegisterType((*Option)(nil), "command_line.Option")
}

func init() {
	proto.RegisterFile("src/main/protobuf/command_line.proto", fileDescriptor_a6c030aeb09bb06d)
}

var fileDescriptor_a6c030aeb09bb06d = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xdf, 0x6a, 0xd4, 0x40,
	0x14, 0xc6, 0x4d, 0xd7, 0x2e, 0xe6, 0x64, 0x5b, 0x64, 0x28, 0x18, 0x7a, 0x63, 0x1a, 0x45, 0x02,
	0xd6, 0x59, 0x58, 0x6f, 0xfc, 0x07, 0x42, 0x45, 0xf1, 0x62, 0x55, 0x88, 0xe2, 0x6d, 0x98, 0x4c,
	0x26, 0xeb, 0xe0, 0xfc, 0x59, 0x32, 0x93, 0x42, 0xdf, 0xce, 0x17, 0xf2, 0x1d, 0x24, 0x33, 0x93,
	0xec, 0x6e, 0xeb, 0x55, 0x72, 0xbe, 0x7c, 0xbf, 0xf3, 0x9d, 0x73, 0x08, 0x3c, 0x35, 0x1d, 0x5d,
	0x4a, 0xc2, 0xd5, 0x72, 0xdb, 0x69, 0xab, 0xeb, 0xbe, 0x5d, 0x52, 0x2d, 0x25, 0x51, 0x4d, 0x25,
	0xb8, 0x62, 0xd8, 0xa9, 0x68, 0xb1, 0xaf, 0x9d, 0x3f, 0xbb, 0xcb, 0xe8, 0xad, 0xe5, 0x5a, 0x55,
	0x2d, 0x17, 0x96, 0x75, 0xc6, 0x53, 0xf9, 0x0d, 0x24, 0x1f, 0x3c, 0xb7, 0xe6, 0x8a, 0xa1, 0x4b,
	0x40, 0xfb, 0x6d, 0x2a, 0x41, 0x6a, 0x26, 0xd2, 0x28, 0x8b, 0x8a, 0xb8, 0x7c, 0x48, 0x77, 0xc6,
	0xf5, 0xa0, 0xa3, 0x77, 0xf0, 0xc0, 0x30, 0x3a, 0x74, 0x35, 0xe9, 0x51, 0x36, 0x2b, 0x92, 0x55,
	0x86, 0x0f, 0x26, 0xdb, 0x6b, 0xfd, 0xdd, 0x1b, 0xcb, 0x89, 0xc8, 0xff, 0x44, 0x80, 0xee, 0x1a,
	0xd0, 0x13, 0x38, 0x09, 0x96, 0x83, 0xf4, 0x45, 0x10, 0x7d, 0xf2, 0x2b, 0x00, 0xfa, 0xab, 0x57,
	0xbf, 0x2b, 0xc1, 0x8d, 0x4d, 0x8f, 0xb2, 0xa8, 0x48, 0x56, 0x8f, 0x6e, 0x65, 0x0f, 0xdf, 0xd7,
	0xdc, 0xd8, 0xcf, 0xf7, 0xca, 0x98, 0x8e, 0x05, 0x7a, 0x0b, 0x49, 0x38, 0x84, 0x43, 0x67, 0x0e,
	0x4d, 0x0f, 0xd1, 0x6f, 0xce, 0x10, 0x58, 0xd0, 0x53, 0x75, 0x75, 0x0a, 0xe3, 0x18, 0x95, 0xbd,
	0xd9, 0xb2, 0xfc, 0x02, 0xe2, 0x29, 0x06, 0x9d, 0xc1, 0xb1, 0x8b, 0x49, 0xa3, 0x6c, 0x56, 0xc4,
	0xa5, 0x2f, 0xf2, 0x37, 0x00, 0xbb, 0x76, 0xe8, 0x12, 0xe6, 0xbe, 0x9d, 0x33, 0x25, 0xab, 0xb3,
	0xff, 0x05, 0x97, 0xc1, 0x93, 0xff, 0x8d, 0x60, 0xee, 0xa5, 0xe1, 0x2a, 0x54, 0xcb, 0x9a, 0x2b,
	0xd6, 0x54, 0xad, 0xee, 0xe4, 0x78, 0x95, 0x51, 0xfc, 0xa4, 0x3b, 0x89, 0x1e, 0x4f, 0xbb, 0x29,
	0x22, 0x99, 0x3b, 0x4b, 0x3c, 0xce, 0xff, 0x95, 0x48, 0x86, 0x2e, 0x60, 0x11, 0x0c, 0xd7, 0x44,
	0xf4, 0xcc, 0x6d, 0x1f, 0x97, 0x01, 0xfa, 0x39, 0x48, 0xe8, 0x35, 0x24, 0xac, 0x6d, 0x19, 0xb5,
	0x95, 0x25, 0x1b, 0x93, 0xde, 0xcf, 0x66, 0xc5, 0xe9, 0x2a, 0xc5, 0xde, 0x62, 0xc2, 0x84, 0x1f,
	0x9d, 0xe3, 0x07, 0xd9, 0x94, 0xc0, 0xc6, 0x57, 0x83, 0xde, 0xc3, 0x89, 0x64, 0x96, 0x34, 0xc4,
	0x12, 0x0f, 0x1f, 0x3b, 0xf8, 0xfc, 0x16, 0xfc, 0x25, 0x78, 0x06, 0x7c, 0x21, 0x77, 0x85, 0xb9,
	0x7a, 0x01, 0xcf, 0xa9, 0x96, 0x78, 0xa3, 0xf5, 0x46, 0x30, 0xdc, 0xb0, 0x6b, 0xab, 0xb5, 0x30,
	0xb8, 0xee, 0xb9, 0x68, 0xb0, 0xe0, 0x35, 0xee, 0x7a, 0x65, 0xb9, 0x0c, 0x7f, 0x7c, 0x3d, 0x77,
	0x8f, 0x97, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xae, 0x90, 0x2a, 0x02, 0x20, 0x03, 0x00, 0x00,
}
