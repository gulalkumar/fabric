// Code generated by protoc-gen-go.
// source: common/configtx.proto
// DO NOT EDIT!

package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ConfigEnvelope is designed to contain _all_ configuration for a chain with no dependency
// on previous configuration transactions.
//
// It is generated with the following scheme:
//   1. Retrieve the existing configuration
//   2. Note the highest configuration sequence number, store it and increment it by one
//   3. Modify desired ConfigItems, setting each LastModified to the stored and incremented sequence number
//     a) Note that the ConfigItem has a ChannelHeader header attached to it, who's type is set to CONFIGURATION_ITEM
//   4. Create Config message containing the new configuration, marshal it into ConfigEnvelope.config and encode the required signatures
//     a) Each signature is of type ConfigSignature
//     b) The ConfigSignature signature is over the concatenation of signatureHeader and the Config bytes (which includes a ChannelHeader)
//   5. Submit new Config for ordering in Envelope signed by submitter
//     a) The Envelope Payload has data set to the marshaled ConfigEnvelope
//     b) The Envelope Payload has a header of type Header.Type.CONFIGURATION_TRANSACTION
//
// The configuration manager will verify:
//   1. All configuration items and the envelope refer to the correct chain
//   2. Some configuration item has been added or modified
//   3. No existing configuration item has been ommitted
//   4. All configuration changes have a LastModification of one more than the last configuration's highest LastModification number
//   5. All configuration changes satisfy the corresponding modification policy
type ConfigEnvelope struct {
	Config     []byte             `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	Signatures []*ConfigSignature `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty"`
}

func (m *ConfigEnvelope) Reset()                    { *m = ConfigEnvelope{} }
func (m *ConfigEnvelope) String() string            { return proto.CompactTextString(m) }
func (*ConfigEnvelope) ProtoMessage()               {}
func (*ConfigEnvelope) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ConfigEnvelope) GetSignatures() []*ConfigSignature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

type ConfigGroupSchema struct {
	Groups   map[string]*ConfigGroupSchema  `protobuf:"bytes,1,rep,name=groups" json:"groups,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Values   map[string]*ConfigValueSchema  `protobuf:"bytes,2,rep,name=values" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Policies map[string]*ConfigPolicySchema `protobuf:"bytes,3,rep,name=policies" json:"policies,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ConfigGroupSchema) Reset()                    { *m = ConfigGroupSchema{} }
func (m *ConfigGroupSchema) String() string            { return proto.CompactTextString(m) }
func (*ConfigGroupSchema) ProtoMessage()               {}
func (*ConfigGroupSchema) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ConfigGroupSchema) GetGroups() map[string]*ConfigGroupSchema {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *ConfigGroupSchema) GetValues() map[string]*ConfigValueSchema {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *ConfigGroupSchema) GetPolicies() map[string]*ConfigPolicySchema {
	if m != nil {
		return m.Policies
	}
	return nil
}

type ConfigValueSchema struct {
}

func (m *ConfigValueSchema) Reset()                    { *m = ConfigValueSchema{} }
func (m *ConfigValueSchema) String() string            { return proto.CompactTextString(m) }
func (*ConfigValueSchema) ProtoMessage()               {}
func (*ConfigValueSchema) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type ConfigPolicySchema struct {
}

func (m *ConfigPolicySchema) Reset()                    { *m = ConfigPolicySchema{} }
func (m *ConfigPolicySchema) String() string            { return proto.CompactTextString(m) }
func (*ConfigPolicySchema) ProtoMessage()               {}
func (*ConfigPolicySchema) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

// Config represents the config for a particular channel
type Config struct {
	Header  *ChannelHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Channel *ConfigGroup   `protobuf:"bytes,2,opt,name=channel" json:"channel,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *Config) GetHeader() *ChannelHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Config) GetChannel() *ConfigGroup {
	if m != nil {
		return m.Channel
	}
	return nil
}

// ConfigGroup is the hierarchical data structure for holding config
type ConfigGroup struct {
	Version   uint64                   `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Groups    map[string]*ConfigGroup  `protobuf:"bytes,2,rep,name=groups" json:"groups,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Values    map[string]*ConfigValue  `protobuf:"bytes,3,rep,name=values" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Policies  map[string]*ConfigPolicy `protobuf:"bytes,4,rep,name=policies" json:"policies,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ModPolicy string                   `protobuf:"bytes,5,opt,name=mod_policy,json=modPolicy" json:"mod_policy,omitempty"`
}

func (m *ConfigGroup) Reset()                    { *m = ConfigGroup{} }
func (m *ConfigGroup) String() string            { return proto.CompactTextString(m) }
func (*ConfigGroup) ProtoMessage()               {}
func (*ConfigGroup) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *ConfigGroup) GetGroups() map[string]*ConfigGroup {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *ConfigGroup) GetValues() map[string]*ConfigValue {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *ConfigGroup) GetPolicies() map[string]*ConfigPolicy {
	if m != nil {
		return m.Policies
	}
	return nil
}

// ConfigValue represents an individual piece of config data
type ConfigValue struct {
	Version   uint64 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Value     []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ModPolicy string `protobuf:"bytes,3,opt,name=mod_policy,json=modPolicy" json:"mod_policy,omitempty"`
}

func (m *ConfigValue) Reset()                    { *m = ConfigValue{} }
func (m *ConfigValue) String() string            { return proto.CompactTextString(m) }
func (*ConfigValue) ProtoMessage()               {}
func (*ConfigValue) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

type ConfigPolicy struct {
	Version   uint64  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Policy    *Policy `protobuf:"bytes,2,opt,name=policy" json:"policy,omitempty"`
	ModPolicy string  `protobuf:"bytes,3,opt,name=mod_policy,json=modPolicy" json:"mod_policy,omitempty"`
}

func (m *ConfigPolicy) Reset()                    { *m = ConfigPolicy{} }
func (m *ConfigPolicy) String() string            { return proto.CompactTextString(m) }
func (*ConfigPolicy) ProtoMessage()               {}
func (*ConfigPolicy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *ConfigPolicy) GetPolicy() *Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type ConfigSignature struct {
	SignatureHeader []byte `protobuf:"bytes,1,opt,name=signature_header,json=signatureHeader,proto3" json:"signature_header,omitempty"`
	Signature       []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *ConfigSignature) Reset()                    { *m = ConfigSignature{} }
func (m *ConfigSignature) String() string            { return proto.CompactTextString(m) }
func (*ConfigSignature) ProtoMessage()               {}
func (*ConfigSignature) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func init() {
	proto.RegisterType((*ConfigEnvelope)(nil), "common.ConfigEnvelope")
	proto.RegisterType((*ConfigGroupSchema)(nil), "common.ConfigGroupSchema")
	proto.RegisterType((*ConfigValueSchema)(nil), "common.ConfigValueSchema")
	proto.RegisterType((*ConfigPolicySchema)(nil), "common.ConfigPolicySchema")
	proto.RegisterType((*Config)(nil), "common.Config")
	proto.RegisterType((*ConfigGroup)(nil), "common.ConfigGroup")
	proto.RegisterType((*ConfigValue)(nil), "common.ConfigValue")
	proto.RegisterType((*ConfigPolicy)(nil), "common.ConfigPolicy")
	proto.RegisterType((*ConfigSignature)(nil), "common.ConfigSignature")
}

func init() { proto.RegisterFile("common/configtx.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x95, 0x51, 0x8f, 0x93, 0x40,
	0x10, 0xc7, 0x43, 0xb9, 0xe3, 0xec, 0x50, 0xef, 0xce, 0x6d, 0x4f, 0x91, 0x68, 0xac, 0x24, 0x6a,
	0x4f, 0xd3, 0x62, 0xea, 0x43, 0x8d, 0xc9, 0xbd, 0x78, 0xb9, 0xe8, 0xd3, 0x45, 0x39, 0xa3, 0xc9,
	0xc5, 0xa4, 0xa1, 0x74, 0x0b, 0x44, 0x60, 0x09, 0xd0, 0x46, 0x3e, 0xab, 0xdf, 0xc0, 0x4f, 0x61,
	0xd8, 0x5d, 0x70, 0x69, 0x69, 0x9b, 0x7b, 0x82, 0xdd, 0x99, 0xff, 0x6f, 0x86, 0xd9, 0x3f, 0x00,
	0x67, 0x0e, 0x09, 0x43, 0x12, 0x99, 0x0e, 0x89, 0x16, 0xbe, 0x9b, 0xfd, 0x1e, 0xc5, 0x09, 0xc9,
	0x08, 0x52, 0xd8, 0xb6, 0xde, 0xad, 0xc2, 0xc5, 0x85, 0x05, 0xf5, 0x52, 0x13, 0x93, 0xc0, 0x77,
	0x7c, 0x9c, 0xb2, 0x6d, 0xc3, 0x86, 0xe3, 0x4b, 0x4a, 0xb9, 0x8a, 0x56, 0x38, 0x20, 0x31, 0x46,
	0x0f, 0x41, 0x61, 0x5c, 0x4d, 0xea, 0x4b, 0x83, 0x8e, 0xc5, 0x57, 0x68, 0x02, 0x90, 0xfa, 0x6e,
	0x64, 0x67, 0xcb, 0x04, 0xa7, 0x5a, 0xab, 0x2f, 0x0f, 0xd4, 0xf1, 0xa3, 0x11, 0xaf, 0xc1, 0x18,
	0x37, 0x65, 0xdc, 0x12, 0x52, 0x8d, 0x3f, 0x32, 0x3c, 0x60, 0xf1, 0x4f, 0x09, 0x59, 0xc6, 0x37,
	0x8e, 0x87, 0x43, 0x1b, 0x5d, 0x80, 0xe2, 0x16, 0xcb, 0x54, 0x93, 0x28, 0xea, 0x45, 0x1d, 0x25,
	0xa4, 0x8e, 0xe8, 0x7d, 0x7a, 0x15, 0x65, 0x49, 0x6e, 0x71, 0x51, 0x21, 0x5f, 0xd9, 0xc1, 0xb2,
	0xea, 0x64, 0x87, 0xfc, 0x3b, 0xcd, 0xe3, 0x72, 0x26, 0x42, 0x97, 0x70, 0xaf, 0x1c, 0x84, 0x26,
	0x53, 0xc0, 0xab, 0xed, 0x80, 0x2f, 0x3c, 0x93, 0x21, 0x2a, 0xa1, 0xfe, 0x0d, 0x54, 0xa1, 0x35,
	0x74, 0x0a, 0xf2, 0x2f, 0x9c, 0xd3, 0xa9, 0xb5, 0xad, 0xe2, 0x16, 0x99, 0x70, 0x48, 0xeb, 0x69,
	0xad, 0xbe, 0x34, 0x50, 0xc7, 0x8f, 0xb7, 0x96, 0xb0, 0x58, 0xde, 0x87, 0xd6, 0x7b, 0xa9, 0xa0,
	0x0a, 0x1d, 0xdf, 0x99, 0x4a, 0xb5, 0x9b, 0xd4, 0x1f, 0x70, 0xbf, 0xf6, 0x18, 0x0d, 0xdc, 0xb7,
	0x75, 0xae, 0x5e, 0xe7, 0x52, 0x75, 0xbe, 0x01, 0x36, 0xba, 0xe5, 0xe1, 0x0a, 0x85, 0x8d, 0x1e,
	0xa0, 0x4d, 0x95, 0xb1, 0x00, 0x85, 0xed, 0xa2, 0x21, 0x28, 0x1e, 0xb6, 0xe7, 0x38, 0xa1, 0xf5,
	0xd5, 0xf1, 0x59, 0x55, 0xcb, 0xb3, 0xa3, 0x08, 0x07, 0x9f, 0x69, 0xd0, 0xe2, 0x49, 0x68, 0x08,
	0x47, 0x0e, 0x0b, 0xf0, 0xde, 0xba, 0x0d, 0x93, 0xb4, 0xca, 0x1c, 0xe3, 0xaf, 0x0c, 0xaa, 0x10,
	0x40, 0x1a, 0x1c, 0xad, 0x70, 0x92, 0xfa, 0x24, 0xa2, 0xe5, 0x0e, 0xac, 0x72, 0x89, 0x26, 0x95,
	0x09, 0x99, 0x8b, 0x9e, 0x35, 0x70, 0x1b, 0xed, 0x37, 0xa9, 0xec, 0x27, 0x6f, 0x17, 0x36, 0x19,
	0xef, 0x42, 0x30, 0xde, 0x01, 0x95, 0x3e, 0x6f, 0x92, 0x6e, 0xb1, 0x1c, 0x7a, 0x0a, 0x10, 0x92,
	0xf9, 0x94, 0xae, 0x73, 0xed, 0x90, 0x1e, 0x5e, 0x3b, 0x24, 0x73, 0x36, 0x67, 0xfd, 0x7a, 0x9f,
	0x23, 0xcf, 0xeb, 0x67, 0xdc, 0x38, 0x47, 0xc1, 0x35, 0xd7, 0xfb, 0xbc, 0xb8, 0x9b, 0x47, 0xb5,
	0x22, 0xef, 0xeb, 0x7e, 0x17, 0xbe, 0xae, 0x13, 0x7b, 0x4d, 0x2e, 0x14, 0xfd, 0xf7, 0xb3, 0x3c,
	0x6b, 0x5a, 0x6c, 0xc7, 0x59, 0xf7, 0x44, 0x70, 0x87, 0x23, 0xd6, 0x06, 0x2a, 0xaf, 0x0d, 0xd4,
	0x20, 0xd0, 0x11, 0x0b, 0xef, 0xc0, 0xbf, 0x04, 0x85, 0x43, 0x58, 0xe3, 0xc7, 0x65, 0xe3, 0xbc,
	0x65, 0x1e, 0xdd, 0x57, 0xf0, 0x16, 0x4e, 0xd6, 0xbe, 0xa5, 0xe8, 0x1c, 0x4e, 0xab, 0xaf, 0xe9,
	0x54, 0x78, 0x6d, 0x3a, 0xd6, 0x49, 0xb5, 0xcf, 0x5e, 0x18, 0xf4, 0x04, 0xda, 0xd5, 0x16, 0x7f,
	0xce, 0xff, 0x1b, 0x1f, 0x87, 0xb7, 0x6f, 0x5c, 0x3f, 0xf3, 0x96, 0xb3, 0xa2, 0x35, 0xd3, 0xcb,
	0x63, 0x9c, 0x04, 0x78, 0xee, 0xe2, 0xc4, 0x5c, 0xd8, 0xb3, 0xc4, 0x77, 0x4c, 0xfa, 0x4b, 0x48,
	0xf9, 0x7f, 0x63, 0xa6, 0xd0, 0xe5, 0xbb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x6c, 0xc3,
	0x6a, 0x6e, 0x06, 0x00, 0x00,
}
