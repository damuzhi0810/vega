// Code generated by protoc-gen-go. DO NOT EDIT.
// source: oracles/v1/oracle_spec.proto

package v1

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

// Status describe the status of the oracle spec
type OracleSpec_Status int32

const (
	// The default value.
	OracleSpec_STATUS_UNSPECIFIED OracleSpec_Status = 0
	// STATUS_ACTIVE describes an active oracle spec.
	OracleSpec_STATUS_ACTIVE OracleSpec_Status = 1
	// STATUS_DEACTIVATED describes an oracle spec that is not listening to data
	// anymore.
	OracleSpec_STATUS_DEACTIVATED OracleSpec_Status = 2
)

var OracleSpec_Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "STATUS_ACTIVE",
	2: "STATUS_DEACTIVATED",
}

var OracleSpec_Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"STATUS_ACTIVE":      1,
	"STATUS_DEACTIVATED": 2,
}

func (x OracleSpec_Status) String() string {
	return proto.EnumName(OracleSpec_Status_name, int32(x))
}

func (OracleSpec_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_42a9664f2aebf360, []int{1, 0}
}

// Type describes the type of properties that are supported by the oracle
// engine.
type PropertyKey_Type int32

const (
	// The default value.
	PropertyKey_TYPE_UNSPECIFIED PropertyKey_Type = 0
	// Any type.
	PropertyKey_TYPE_EMPTY PropertyKey_Type = 1
	// Integer type.
	PropertyKey_TYPE_INTEGER PropertyKey_Type = 2
	// String type.
	PropertyKey_TYPE_STRING PropertyKey_Type = 3
	// Boolean type.
	PropertyKey_TYPE_BOOLEAN PropertyKey_Type = 4
	// Any floating point decimal type.
	PropertyKey_TYPE_DECIMAL PropertyKey_Type = 5
	// Timestamp date type.
	PropertyKey_TYPE_TIMESTAMP PropertyKey_Type = 6
)

var PropertyKey_Type_name = map[int32]string{
	0: "TYPE_UNSPECIFIED",
	1: "TYPE_EMPTY",
	2: "TYPE_INTEGER",
	3: "TYPE_STRING",
	4: "TYPE_BOOLEAN",
	5: "TYPE_DECIMAL",
	6: "TYPE_TIMESTAMP",
}

var PropertyKey_Type_value = map[string]int32{
	"TYPE_UNSPECIFIED": 0,
	"TYPE_EMPTY":       1,
	"TYPE_INTEGER":     2,
	"TYPE_STRING":      3,
	"TYPE_BOOLEAN":     4,
	"TYPE_DECIMAL":     5,
	"TYPE_TIMESTAMP":   6,
}

func (x PropertyKey_Type) String() string {
	return proto.EnumName(PropertyKey_Type_name, int32(x))
}

func (PropertyKey_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_42a9664f2aebf360, []int{3, 0}
}

// Comparator describes the type of comparison.
type Condition_Operator int32

const (
	// The default value
	Condition_OPERATOR_UNSPECIFIED Condition_Operator = 0
	// Verify if the property values are strictly equal or not.
	Condition_OPERATOR_EQUALS Condition_Operator = 1
	// Verify if the oracle data value is greater than the Condition value.
	Condition_OPERATOR_GREATER_THAN Condition_Operator = 2
	// Verify if the oracle data value is greater than or equal to the Condition
	// value.
	Condition_OPERATOR_GREATER_THAN_OR_EQUAL Condition_Operator = 3
	// Verify if the oracle data value is less than the Condition value.
	Condition_OPERATOR_LESS_THAN Condition_Operator = 4
	// Verify if the oracle data value is less or equal to than the Condition
	// value.
	Condition_OPERATOR_LESS_THAN_OR_EQUAL Condition_Operator = 5
)

var Condition_Operator_name = map[int32]string{
	0: "OPERATOR_UNSPECIFIED",
	1: "OPERATOR_EQUALS",
	2: "OPERATOR_GREATER_THAN",
	3: "OPERATOR_GREATER_THAN_OR_EQUAL",
	4: "OPERATOR_LESS_THAN",
	5: "OPERATOR_LESS_THAN_OR_EQUAL",
}

var Condition_Operator_value = map[string]int32{
	"OPERATOR_UNSPECIFIED":           0,
	"OPERATOR_EQUALS":                1,
	"OPERATOR_GREATER_THAN":          2,
	"OPERATOR_GREATER_THAN_OR_EQUAL": 3,
	"OPERATOR_LESS_THAN":             4,
	"OPERATOR_LESS_THAN_OR_EQUAL":    5,
}

func (x Condition_Operator) String() string {
	return proto.EnumName(Condition_Operator_name, int32(x))
}

func (Condition_Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_42a9664f2aebf360, []int{4, 0}
}

// An oracle spec describe the oracle data that a product (or a risk model)
// wants to get from the oracle engine.
type OracleSpecConfiguration struct {
	// pubKeys is the list of authorized public keys that signed the data for this
	// oracle. All the public keys in the oracle data should be contained in these
	// public keys.
	PubKeys []string `protobuf:"bytes,1,rep,name=pub_keys,json=pubKeys,proto3" json:"pub_keys,omitempty"`
	// filters describes which oracle data are considered of interest or not for
	// the product (or the risk model).
	Filters              []*Filter `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *OracleSpecConfiguration) Reset()         { *m = OracleSpecConfiguration{} }
func (m *OracleSpecConfiguration) String() string { return proto.CompactTextString(m) }
func (*OracleSpecConfiguration) ProtoMessage()    {}
func (*OracleSpecConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_42a9664f2aebf360, []int{0}
}

func (m *OracleSpecConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OracleSpecConfiguration.Unmarshal(m, b)
}
func (m *OracleSpecConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OracleSpecConfiguration.Marshal(b, m, deterministic)
}
func (m *OracleSpecConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleSpecConfiguration.Merge(m, src)
}
func (m *OracleSpecConfiguration) XXX_Size() int {
	return xxx_messageInfo_OracleSpecConfiguration.Size(m)
}
func (m *OracleSpecConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleSpecConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_OracleSpecConfiguration proto.InternalMessageInfo

func (m *OracleSpecConfiguration) GetPubKeys() []string {
	if m != nil {
		return m.PubKeys
	}
	return nil
}

func (m *OracleSpecConfiguration) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

// An oracle spec describe the oracle data that a product (or a risk model)
// wants to get from the oracle engine.
// This message contains additional information used by the API.
type OracleSpec struct {
	// id is a hash generated from the OracleSpec data.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Creation Date time
	CreatedAt int64 `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last Updated timestamp
	UpdatedAt int64 `protobuf:"varint,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// pubKeys is the list of authorized public keys that signed the data for this
	// oracle. All the public keys in the oracle data should be contained in these
	// public keys.
	PubKeys []string `protobuf:"bytes,4,rep,name=pub_keys,json=pubKeys,proto3" json:"pub_keys,omitempty"`
	// filters describes which oracle data are considered of interest or not for
	// the product (or the risk model).
	Filters []*Filter `protobuf:"bytes,5,rep,name=filters,proto3" json:"filters,omitempty"`
	// status describes the status of the oracle spec
	Status               OracleSpec_Status `protobuf:"varint,6,opt,name=status,proto3,enum=oracles.v1.OracleSpec_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *OracleSpec) Reset()         { *m = OracleSpec{} }
func (m *OracleSpec) String() string { return proto.CompactTextString(m) }
func (*OracleSpec) ProtoMessage()    {}
func (*OracleSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_42a9664f2aebf360, []int{1}
}

func (m *OracleSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OracleSpec.Unmarshal(m, b)
}
func (m *OracleSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OracleSpec.Marshal(b, m, deterministic)
}
func (m *OracleSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleSpec.Merge(m, src)
}
func (m *OracleSpec) XXX_Size() int {
	return xxx_messageInfo_OracleSpec.Size(m)
}
func (m *OracleSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleSpec.DiscardUnknown(m)
}

var xxx_messageInfo_OracleSpec proto.InternalMessageInfo

func (m *OracleSpec) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OracleSpec) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *OracleSpec) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *OracleSpec) GetPubKeys() []string {
	if m != nil {
		return m.PubKeys
	}
	return nil
}

func (m *OracleSpec) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *OracleSpec) GetStatus() OracleSpec_Status {
	if m != nil {
		return m.Status
	}
	return OracleSpec_STATUS_UNSPECIFIED
}

// Filter describes the conditions under which an oracle data is considered of
// interest or not.
type Filter struct {
	// key is the oracle data property key targeted by the filter.
	Key *PropertyKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// conditions are the conditions that should be matched by the data to be
	//considered of interest.
	Conditions           []*Condition `protobuf:"bytes,2,rep,name=conditions,proto3" json:"conditions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_42a9664f2aebf360, []int{2}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetKey() *PropertyKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Filter) GetConditions() []*Condition {
	if m != nil {
		return m.Conditions
	}
	return nil
}

// PropertyKey describes the property key contained in an oracle data.
type PropertyKey struct {
	// name is the name of the property.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// type is the type of the property.
	Type                 PropertyKey_Type `protobuf:"varint,2,opt,name=type,proto3,enum=oracles.v1.PropertyKey_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PropertyKey) Reset()         { *m = PropertyKey{} }
func (m *PropertyKey) String() string { return proto.CompactTextString(m) }
func (*PropertyKey) ProtoMessage()    {}
func (*PropertyKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_42a9664f2aebf360, []int{3}
}

func (m *PropertyKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PropertyKey.Unmarshal(m, b)
}
func (m *PropertyKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PropertyKey.Marshal(b, m, deterministic)
}
func (m *PropertyKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PropertyKey.Merge(m, src)
}
func (m *PropertyKey) XXX_Size() int {
	return xxx_messageInfo_PropertyKey.Size(m)
}
func (m *PropertyKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PropertyKey.DiscardUnknown(m)
}

var xxx_messageInfo_PropertyKey proto.InternalMessageInfo

func (m *PropertyKey) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PropertyKey) GetType() PropertyKey_Type {
	if m != nil {
		return m.Type
	}
	return PropertyKey_TYPE_UNSPECIFIED
}

// Condition describes the condition that must be validated by the
type Condition struct {
	// comparator is the type of comparison to make on the value.
	Operator Condition_Operator `protobuf:"varint,1,opt,name=operator,proto3,enum=oracles.v1.Condition_Operator" json:"operator,omitempty"`
	// value is used by the comparator.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Condition) Reset()         { *m = Condition{} }
func (m *Condition) String() string { return proto.CompactTextString(m) }
func (*Condition) ProtoMessage()    {}
func (*Condition) Descriptor() ([]byte, []int) {
	return fileDescriptor_42a9664f2aebf360, []int{4}
}

func (m *Condition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Condition.Unmarshal(m, b)
}
func (m *Condition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Condition.Marshal(b, m, deterministic)
}
func (m *Condition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Condition.Merge(m, src)
}
func (m *Condition) XXX_Size() int {
	return xxx_messageInfo_Condition.Size(m)
}
func (m *Condition) XXX_DiscardUnknown() {
	xxx_messageInfo_Condition.DiscardUnknown(m)
}

var xxx_messageInfo_Condition proto.InternalMessageInfo

func (m *Condition) GetOperator() Condition_Operator {
	if m != nil {
		return m.Operator
	}
	return Condition_OPERATOR_UNSPECIFIED
}

func (m *Condition) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterEnum("oracles.v1.OracleSpec_Status", OracleSpec_Status_name, OracleSpec_Status_value)
	proto.RegisterEnum("oracles.v1.PropertyKey_Type", PropertyKey_Type_name, PropertyKey_Type_value)
	proto.RegisterEnum("oracles.v1.Condition_Operator", Condition_Operator_name, Condition_Operator_value)
	proto.RegisterType((*OracleSpecConfiguration)(nil), "oracles.v1.OracleSpecConfiguration")
	proto.RegisterType((*OracleSpec)(nil), "oracles.v1.OracleSpec")
	proto.RegisterType((*Filter)(nil), "oracles.v1.Filter")
	proto.RegisterType((*PropertyKey)(nil), "oracles.v1.PropertyKey")
	proto.RegisterType((*Condition)(nil), "oracles.v1.Condition")
}

func init() { proto.RegisterFile("oracles/v1/oracle_spec.proto", fileDescriptor_42a9664f2aebf360) }

var fileDescriptor_42a9664f2aebf360 = []byte{
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xcf, 0x4e, 0xdb, 0x4a,
	0x14, 0xc6, 0xaf, 0x9d, 0x3f, 0x90, 0x93, 0x7b, 0xc3, 0xdc, 0x53, 0x28, 0x41, 0x05, 0x1a, 0x79,
	0x95, 0x56, 0xc8, 0x29, 0xa9, 0xd8, 0x74, 0x67, 0x92, 0x81, 0x5a, 0xe4, 0x5f, 0xc7, 0x43, 0x25,
	0xba, 0x89, 0x1c, 0x67, 0x40, 0x2e, 0x69, 0x6c, 0xd9, 0x4e, 0x24, 0x3f, 0x41, 0xd7, 0x7d, 0x8a,
	0xae, 0xfb, 0x40, 0x7d, 0x97, 0xca, 0x63, 0xc7, 0x49, 0x29, 0xa8, 0x3b, 0x9f, 0xef, 0xfb, 0x9d,
	0xc9, 0x37, 0xe7, 0x44, 0x03, 0x87, 0x5e, 0x60, 0x3b, 0x33, 0x11, 0xb6, 0x96, 0xa7, 0xad, 0xf4,
	0x73, 0x1c, 0xfa, 0xc2, 0xd1, 0xfd, 0xc0, 0x8b, 0x3c, 0x84, 0xcc, 0xd5, 0x97, 0xa7, 0xda, 0x04,
	0xf6, 0x87, 0xb2, 0xb2, 0x7c, 0xe1, 0x74, 0xbc, 0xf9, 0xad, 0x7b, 0xb7, 0x08, 0xec, 0xc8, 0xf5,
	0xe6, 0x78, 0x00, 0xdb, 0xfe, 0x62, 0x32, 0xbe, 0x17, 0x71, 0x58, 0x57, 0x1a, 0x85, 0x66, 0x85,
	0x6d, 0xf9, 0x8b, 0xc9, 0x95, 0x88, 0x43, 0x3c, 0x81, 0xad, 0x5b, 0x77, 0x16, 0x89, 0x20, 0xac,
	0xab, 0x8d, 0x42, 0xb3, 0xda, 0x46, 0x7d, 0x7d, 0xa6, 0x7e, 0x21, 0x2d, 0xb6, 0x42, 0xb4, 0xef,
	0x2a, 0xc0, 0xfa, 0x47, 0xb0, 0x06, 0xaa, 0x3b, 0xad, 0x2b, 0x0d, 0xa5, 0x59, 0x61, 0xaa, 0x3b,
	0xc5, 0x23, 0x00, 0x27, 0x10, 0x76, 0x24, 0xa6, 0x63, 0x3b, 0xaa, 0xab, 0x0d, 0xa5, 0x59, 0x60,
	0x95, 0x4c, 0x31, 0xa2, 0xc4, 0x5e, 0xf8, 0xd3, 0x95, 0x5d, 0x48, 0xed, 0x4c, 0x31, 0xa2, 0xdf,
	0x52, 0x16, 0x9f, 0x4c, 0x59, 0xfa, 0x6b, 0x4a, 0x3c, 0x83, 0x72, 0x18, 0xd9, 0xd1, 0x22, 0xac,
	0x97, 0x1b, 0x4a, 0xb3, 0xd6, 0x3e, 0xda, 0x84, 0xd7, 0xf1, 0x75, 0x4b, 0x42, 0x2c, 0x83, 0xb5,
	0x2b, 0x28, 0xa7, 0x0a, 0x3e, 0x07, 0xb4, 0xb8, 0xc1, 0xaf, 0xad, 0xf1, 0xf5, 0xc0, 0x1a, 0xd1,
	0x8e, 0x79, 0x61, 0xd2, 0x2e, 0xf9, 0x07, 0xff, 0x87, 0xff, 0x32, 0xdd, 0xe8, 0x70, 0xf3, 0x23,
	0x25, 0xca, 0x06, 0xda, 0xa5, 0x52, 0x34, 0x38, 0xed, 0x12, 0x55, 0xfb, 0x0c, 0xe5, 0x34, 0x16,
	0xbe, 0x82, 0xc2, 0xbd, 0x88, 0xe5, 0x94, 0xaa, 0xed, 0xfd, 0xcd, 0x28, 0xa3, 0xc0, 0xf3, 0x45,
	0x10, 0xc5, 0x57, 0x22, 0x66, 0x09, 0x83, 0x67, 0x00, 0x8e, 0x37, 0x9f, 0xba, 0xc9, 0xd2, 0x56,
	0xfb, 0xd8, 0xdb, 0xec, 0xe8, 0xac, 0x5c, 0xb6, 0x01, 0x6a, 0x3f, 0x15, 0xa8, 0x6e, 0x9c, 0x85,
	0x08, 0xc5, 0xb9, 0xfd, 0x45, 0x64, 0x8b, 0x91, 0xdf, 0xf8, 0x06, 0x8a, 0x51, 0xec, 0x0b, 0xb9,
	0x94, 0x5a, 0xfb, 0xf0, 0x89, 0x18, 0x3a, 0x8f, 0x7d, 0xc1, 0x24, 0xa9, 0x7d, 0x55, 0xa0, 0x98,
	0x94, 0xb8, 0x0b, 0x84, 0xdf, 0x8c, 0xe8, 0x83, 0x59, 0xd4, 0x00, 0xa4, 0x4a, 0xfb, 0x23, 0x7e,
	0x43, 0x14, 0x24, 0xf0, 0xaf, 0xac, 0xcd, 0x01, 0xa7, 0x97, 0x94, 0x11, 0x15, 0x77, 0xa0, 0x2a,
	0x15, 0x8b, 0x33, 0x73, 0x70, 0x49, 0x0a, 0x39, 0x72, 0x3e, 0x1c, 0xf6, 0xa8, 0x31, 0x20, 0xc5,
	0x5c, 0xe9, 0xd2, 0x8e, 0xd9, 0x37, 0x7a, 0xa4, 0x84, 0x08, 0x35, 0xa9, 0x70, 0xb3, 0x4f, 0x2d,
	0x6e, 0xf4, 0x47, 0xa4, 0xac, 0x7d, 0x53, 0xa1, 0x92, 0xdf, 0x1c, 0xdf, 0xc1, 0x76, 0x92, 0xd7,
	0x8e, 0xbc, 0x40, 0xde, 0xb0, 0xd6, 0x3e, 0x7e, 0x74, 0x44, 0xfa, 0x30, 0xa3, 0x58, 0xce, 0xe3,
	0x2e, 0x94, 0x96, 0xf6, 0x6c, 0x91, 0x8e, 0xa1, 0xc2, 0xd2, 0x42, 0xfb, 0xa1, 0xc0, 0xf6, 0x0a,
	0xc6, 0x3a, 0xec, 0x0e, 0x47, 0x94, 0x19, 0x7c, 0xc8, 0x1e, 0xdc, 0xf8, 0x19, 0xec, 0xe4, 0x0e,
	0xfd, 0x70, 0x6d, 0xf4, 0x2c, 0xa2, 0xe0, 0x01, 0xec, 0xe5, 0xe2, 0x25, 0xa3, 0x06, 0xa7, 0x6c,
	0xcc, 0xdf, 0x1b, 0x03, 0xa2, 0xa2, 0x06, 0xc7, 0x8f, 0x5a, 0xe3, 0x55, 0x3f, 0x29, 0x24, 0x7f,
	0x9f, 0x9c, 0xe9, 0x51, 0xcb, 0x4a, 0x7b, 0x8b, 0xf8, 0x12, 0x5e, 0xfc, 0xa9, 0xaf, 0x1b, 0x4b,
	0xe7, 0x27, 0x9f, 0x5e, 0x3b, 0xde, 0x54, 0xe8, 0x4b, 0x71, 0x67, 0xcb, 0xb7, 0xc0, 0xf1, 0x66,
	0xba, 0xeb, 0xb5, 0x92, 0xba, 0x25, 0x85, 0xd6, 0xfa, 0xe5, 0x98, 0x94, 0xa5, 0xf2, 0xf6, 0x57,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x47, 0xd3, 0x30, 0x4e, 0x04, 0x00, 0x00,
}