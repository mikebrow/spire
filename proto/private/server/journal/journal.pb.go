// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v4.24.4
// source: private/server/journal/journal.proto

package journal

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	// Status is unknown.
	Status_UNKNOWN Status = 0
	// This holds a new authority that was prepared for future uses.
	Status_PREPARED Status = 2
	// This holds the active authority that is currently being used for
	// signing operations.
	Status_ACTIVE Status = 3
	// This holds an old authority that is no longer used.
	Status_OLD Status = 4
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "UNKNOWN",
		2: "PREPARED",
		3: "ACTIVE",
		4: "OLD",
	}
	Status_value = map[string]int32{
		"UNKNOWN":  0,
		"PREPARED": 2,
		"ACTIVE":   3,
		"OLD":      4,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_private_server_journal_journal_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_private_server_journal_journal_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_private_server_journal_journal_proto_rawDescGZIP(), []int{0}
}

type X509CAEntry struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Which X509 CA slot this entry occupied.
	SlotId string `protobuf:"bytes,1,opt,name=slot_id,json=slotId,proto3" json:"slot_id,omitempty"`
	// When the CA was issued (unix epoch in seconds)
	IssuedAt int64 `protobuf:"varint,2,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	// DER encoded CA certificate
	Certificate []byte `protobuf:"bytes,3,opt,name=certificate,proto3" json:"certificate,omitempty"`
	// DER encoded upstream CA chain. See the X509CA struct for details.
	UpstreamChain [][]byte `protobuf:"bytes,4,rep,name=upstream_chain,json=upstreamChain,proto3" json:"upstream_chain,omitempty"`
	// The entry status
	Status Status `protobuf:"varint,5,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
	// The X.509 Subject Key Identifier (SKID)
	AuthorityId string `protobuf:"bytes,6,opt,name=authority_id,json=authorityId,proto3" json:"authority_id,omitempty"`
	// When the CA expires (unix epoch in seconds)
	NotAfter int64 `protobuf:"varint,7,opt,name=not_after,json=notAfter,proto3" json:"not_after,omitempty"`
	// The X.509 Authority Subject Key Identifier (SKID)
	UpstreamAuthorityId string `protobuf:"bytes,8,opt,name=upstream_authority_id,json=upstreamAuthorityId,proto3" json:"upstream_authority_id,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *X509CAEntry) Reset() {
	*x = X509CAEntry{}
	mi := &file_private_server_journal_journal_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *X509CAEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*X509CAEntry) ProtoMessage() {}

func (x *X509CAEntry) ProtoReflect() protoreflect.Message {
	mi := &file_private_server_journal_journal_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use X509CAEntry.ProtoReflect.Descriptor instead.
func (*X509CAEntry) Descriptor() ([]byte, []int) {
	return file_private_server_journal_journal_proto_rawDescGZIP(), []int{0}
}

func (x *X509CAEntry) GetSlotId() string {
	if x != nil {
		return x.SlotId
	}
	return ""
}

func (x *X509CAEntry) GetIssuedAt() int64 {
	if x != nil {
		return x.IssuedAt
	}
	return 0
}

func (x *X509CAEntry) GetCertificate() []byte {
	if x != nil {
		return x.Certificate
	}
	return nil
}

func (x *X509CAEntry) GetUpstreamChain() [][]byte {
	if x != nil {
		return x.UpstreamChain
	}
	return nil
}

func (x *X509CAEntry) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UNKNOWN
}

func (x *X509CAEntry) GetAuthorityId() string {
	if x != nil {
		return x.AuthorityId
	}
	return ""
}

func (x *X509CAEntry) GetNotAfter() int64 {
	if x != nil {
		return x.NotAfter
	}
	return 0
}

func (x *X509CAEntry) GetUpstreamAuthorityId() string {
	if x != nil {
		return x.UpstreamAuthorityId
	}
	return ""
}

type JWTKeyEntry struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Which JWT Key slot this entry occupied.
	SlotId string `protobuf:"bytes,1,opt,name=slot_id,json=slotId,proto3" json:"slot_id,omitempty"`
	// When the key was issued (unix epoch in seconds)
	IssuedAt int64 `protobuf:"varint,2,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	// When the key expires unix epoch in seconds)
	NotAfter int64 `protobuf:"varint,3,opt,name=not_after,json=notAfter,proto3" json:"not_after,omitempty"`
	// JWT key id (i.e. "kid" claim)
	Kid string `protobuf:"bytes,4,opt,name=kid,proto3" json:"kid,omitempty"`
	// PKIX encoded public key
	PublicKey []byte `protobuf:"bytes,5,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// The entry status
	Status Status `protobuf:"varint,6,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
	// The JWT key ID
	AuthorityId   string `protobuf:"bytes,7,opt,name=authority_id,json=authorityId,proto3" json:"authority_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JWTKeyEntry) Reset() {
	*x = JWTKeyEntry{}
	mi := &file_private_server_journal_journal_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JWTKeyEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTKeyEntry) ProtoMessage() {}

func (x *JWTKeyEntry) ProtoReflect() protoreflect.Message {
	mi := &file_private_server_journal_journal_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWTKeyEntry.ProtoReflect.Descriptor instead.
func (*JWTKeyEntry) Descriptor() ([]byte, []int) {
	return file_private_server_journal_journal_proto_rawDescGZIP(), []int{1}
}

func (x *JWTKeyEntry) GetSlotId() string {
	if x != nil {
		return x.SlotId
	}
	return ""
}

func (x *JWTKeyEntry) GetIssuedAt() int64 {
	if x != nil {
		return x.IssuedAt
	}
	return 0
}

func (x *JWTKeyEntry) GetNotAfter() int64 {
	if x != nil {
		return x.NotAfter
	}
	return 0
}

func (x *JWTKeyEntry) GetKid() string {
	if x != nil {
		return x.Kid
	}
	return ""
}

func (x *JWTKeyEntry) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *JWTKeyEntry) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UNKNOWN
}

func (x *JWTKeyEntry) GetAuthorityId() string {
	if x != nil {
		return x.AuthorityId
	}
	return ""
}

type Entries struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	X509CAs       []*X509CAEntry         `protobuf:"bytes,1,rep,name=x509CAs,proto3" json:"x509CAs,omitempty"`
	JwtKeys       []*JWTKeyEntry         `protobuf:"bytes,2,rep,name=jwtKeys,proto3" json:"jwtKeys,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Entries) Reset() {
	*x = Entries{}
	mi := &file_private_server_journal_journal_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Entries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entries) ProtoMessage() {}

func (x *Entries) ProtoReflect() protoreflect.Message {
	mi := &file_private_server_journal_journal_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entries.ProtoReflect.Descriptor instead.
func (*Entries) Descriptor() ([]byte, []int) {
	return file_private_server_journal_journal_proto_rawDescGZIP(), []int{2}
}

func (x *Entries) GetX509CAs() []*X509CAEntry {
	if x != nil {
		return x.X509CAs
	}
	return nil
}

func (x *Entries) GetJwtKeys() []*JWTKeyEntry {
	if x != nil {
		return x.JwtKeys
	}
	return nil
}

var File_private_server_journal_journal_proto protoreflect.FileDescriptor

var file_private_server_journal_journal_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x02, 0x0a, 0x0b, 0x58, 0x35, 0x30, 0x39, 0x43,
	0x41, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0d, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x74,
	0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x6f,
	0x74, 0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x15, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0xd5, 0x01, 0x0a, 0x0b, 0x4a,
	0x57, 0x54, 0x4b, 0x65, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6c,
	0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6c, 0x6f,
	0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1f,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x22, 0x59, 0x0a, 0x07, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x26, 0x0a,
	0x07, 0x78, 0x35, 0x30, 0x39, 0x43, 0x41, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x58, 0x35, 0x30, 0x39, 0x43, 0x41, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x78, 0x35,
	0x30, 0x39, 0x43, 0x41, 0x73, 0x12, 0x26, 0x0a, 0x07, 0x6a, 0x77, 0x74, 0x4b, 0x65, 0x79, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4a, 0x57, 0x54, 0x4b, 0x65, 0x79, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6a, 0x77, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x2a, 0x38, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x45, 0x50, 0x41, 0x52, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x03, 0x12, 0x07,
	0x0a, 0x03, 0x4f, 0x4c, 0x44, 0x10, 0x04, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x2f, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_private_server_journal_journal_proto_rawDescOnce sync.Once
	file_private_server_journal_journal_proto_rawDescData = file_private_server_journal_journal_proto_rawDesc
)

func file_private_server_journal_journal_proto_rawDescGZIP() []byte {
	file_private_server_journal_journal_proto_rawDescOnce.Do(func() {
		file_private_server_journal_journal_proto_rawDescData = protoimpl.X.CompressGZIP(file_private_server_journal_journal_proto_rawDescData)
	})
	return file_private_server_journal_journal_proto_rawDescData
}

var file_private_server_journal_journal_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_private_server_journal_journal_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_private_server_journal_journal_proto_goTypes = []any{
	(Status)(0),         // 0: Status
	(*X509CAEntry)(nil), // 1: X509CAEntry
	(*JWTKeyEntry)(nil), // 2: JWTKeyEntry
	(*Entries)(nil),     // 3: Entries
}
var file_private_server_journal_journal_proto_depIdxs = []int32{
	0, // 0: X509CAEntry.status:type_name -> Status
	0, // 1: JWTKeyEntry.status:type_name -> Status
	1, // 2: Entries.x509CAs:type_name -> X509CAEntry
	2, // 3: Entries.jwtKeys:type_name -> JWTKeyEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_private_server_journal_journal_proto_init() }
func file_private_server_journal_journal_proto_init() {
	if File_private_server_journal_journal_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_private_server_journal_journal_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_private_server_journal_journal_proto_goTypes,
		DependencyIndexes: file_private_server_journal_journal_proto_depIdxs,
		EnumInfos:         file_private_server_journal_journal_proto_enumTypes,
		MessageInfos:      file_private_server_journal_journal_proto_msgTypes,
	}.Build()
	File_private_server_journal_journal_proto = out.File
	file_private_server_journal_journal_proto_rawDesc = nil
	file_private_server_journal_journal_proto_goTypes = nil
	file_private_server_journal_journal_proto_depIdxs = nil
}
