package sp

// import (
// 	reflect "reflect"
// 	sync "sync"

// 	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
// 	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
// )

// const (
// 	// Verify that this generated code is sufficiently up-to-date.
// 	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
// 	// Verify that runtime/protoimpl is sufficiently up-to-date.
// 	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
// )

// type FieldDescriptorProto_Type int32

// const (
// 	// 0 is reserved for errors.
// 	// Order is weird for historical reasons.
// 	FieldDescriptorProto_TYPE_DOUBLE FieldDescriptorProto_Type = 1
// 	FieldDescriptorProto_TYPE_FLOAT  FieldDescriptorProto_Type = 2
// 	// Not ZigZag encoded.  Negative numbers take 10 bytes.  Use TYPE_SINT64 if
// 	// negative values are likely.
// 	FieldDescriptorProto_TYPE_INT64  FieldDescriptorProto_Type = 3
// 	FieldDescriptorProto_TYPE_UINT64 FieldDescriptorProto_Type = 4
// 	// Not ZigZag encoded.  Negative numbers take 10 bytes.  Use TYPE_SINT32 if
// 	// negative values are likely.
// 	FieldDescriptorProto_TYPE_INT32   FieldDescriptorProto_Type = 5
// 	FieldDescriptorProto_TYPE_FIXED64 FieldDescriptorProto_Type = 6
// 	FieldDescriptorProto_TYPE_FIXED32 FieldDescriptorProto_Type = 7
// 	FieldDescriptorProto_TYPE_BOOL    FieldDescriptorProto_Type = 8
// 	FieldDescriptorProto_TYPE_STRING  FieldDescriptorProto_Type = 9
// 	// Tag-delimited aggregate.
// 	// Group type is deprecated and not supported in proto3. However, Proto3
// 	// implementations should still be able to parse the group wire format and
// 	// treat group fields as unknown fields.
// 	FieldDescriptorProto_TYPE_GROUP   FieldDescriptorProto_Type = 10
// 	FieldDescriptorProto_TYPE_MESSAGE FieldDescriptorProto_Type = 11 // Length-delimited aggregate.
// 	// New in version 2.
// 	FieldDescriptorProto_TYPE_BYTES    FieldDescriptorProto_Type = 12
// 	FieldDescriptorProto_TYPE_UINT32   FieldDescriptorProto_Type = 13
// 	FieldDescriptorProto_TYPE_ENUM     FieldDescriptorProto_Type = 14
// 	FieldDescriptorProto_TYPE_SFIXED32 FieldDescriptorProto_Type = 15
// 	FieldDescriptorProto_TYPE_SFIXED64 FieldDescriptorProto_Type = 16
// 	FieldDescriptorProto_TYPE_SINT32   FieldDescriptorProto_Type = 17 // Uses ZigZag encoding.
// 	FieldDescriptorProto_TYPE_SINT64   FieldDescriptorProto_Type = 18 // Uses ZigZag encoding.
// )

// // Enum value maps for FieldDescriptorProto_Type.
// var (
// 	FieldDescriptorProto_Type_name = map[int32]string{
// 		1:  "TYPE_DOUBLE",
// 		2:  "TYPE_FLOAT",
// 		3:  "TYPE_INT64",
// 		4:  "TYPE_UINT64",
// 		5:  "TYPE_INT32",
// 		6:  "TYPE_FIXED64",
// 		7:  "TYPE_FIXED32",
// 		8:  "TYPE_BOOL",
// 		9:  "TYPE_STRING",
// 		10: "TYPE_GROUP",
// 		11: "TYPE_MESSAGE",
// 		12: "TYPE_BYTES",
// 		13: "TYPE_UINT32",
// 		14: "TYPE_ENUM",
// 		15: "TYPE_SFIXED32",
// 		16: "TYPE_SFIXED64",
// 		17: "TYPE_SINT32",
// 		18: "TYPE_SINT64",
// 	}
// 	FieldDescriptorProto_Type_value = map[string]int32{
// 		"TYPE_DOUBLE":   1,
// 		"TYPE_FLOAT":    2,
// 		"TYPE_INT64":    3,
// 		"TYPE_UINT64":   4,
// 		"TYPE_INT32":    5,
// 		"TYPE_FIXED64":  6,
// 		"TYPE_FIXED32":  7,
// 		"TYPE_BOOL":     8,
// 		"TYPE_STRING":   9,
// 		"TYPE_GROUP":    10,
// 		"TYPE_MESSAGE":  11,
// 		"TYPE_BYTES":    12,
// 		"TYPE_UINT32":   13,
// 		"TYPE_ENUM":     14,
// 		"TYPE_SFIXED32": 15,
// 		"TYPE_SFIXED64": 16,
// 		"TYPE_SINT32":   17,
// 		"TYPE_SINT64":   18,
// 	}
// )

// func (x FieldDescriptorProto_Type) Enum() *FieldDescriptorProto_Type {
// 	p := new(FieldDescriptorProto_Type)
// 	*p = x
// 	return p
// }

// func (x FieldDescriptorProto_Type) String() string {
// 	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
// }

// func (FieldDescriptorProto_Type) Descriptor() protoreflect.EnumDescriptor {
// 	return file_descriptor_proto_enumTypes[0].Descriptor()
// }

// func (FieldDescriptorProto_Type) Type() protoreflect.EnumType {
// 	return &file_descriptor_proto_enumTypes[0]
// }

// func (x FieldDescriptorProto_Type) Number() protoreflect.EnumNumber {
// 	return protoreflect.EnumNumber(x)
// }

// // Deprecated: Do not use.
// func (x *FieldDescriptorProto_Type) UnmarshalJSON(b []byte) error {
// 	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
// 	if err != nil {
// 		return err
// 	}
// 	*x = FieldDescriptorProto_Type(num)
// 	return nil
// }

// // Deprecated: Use FieldDescriptorProto_Type.Descriptor instead.
// func (FieldDescriptorProto_Type) EnumDescriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{4, 0}
// }

// type FieldDescriptorProto_Label int32

// const (
// 	// 0 is reserved for errors
// 	FieldDescriptorProto_LABEL_OPTIONAL FieldDescriptorProto_Label = 1
// 	FieldDescriptorProto_LABEL_REQUIRED FieldDescriptorProto_Label = 2
// 	FieldDescriptorProto_LABEL_REPEATED FieldDescriptorProto_Label = 3
// )

// // Enum value maps for FieldDescriptorProto_Label.
// var (
// 	FieldDescriptorProto_Label_name = map[int32]string{
// 		1: "LABEL_OPTIONAL",
// 		2: "LABEL_REQUIRED",
// 		3: "LABEL_REPEATED",
// 	}
// 	FieldDescriptorProto_Label_value = map[string]int32{
// 		"LABEL_OPTIONAL": 1,
// 		"LABEL_REQUIRED": 2,
// 		"LABEL_REPEATED": 3,
// 	}
// )

// func (x FieldDescriptorProto_Label) Enum() *FieldDescriptorProto_Label {
// 	p := new(FieldDescriptorProto_Label)
// 	*p = x
// 	return p
// }

// func (x FieldDescriptorProto_Label) String() string {
// 	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
// }

// func (FieldDescriptorProto_Label) Descriptor() protoreflect.EnumDescriptor {
// 	return file_descriptor_proto_enumTypes[1].Descriptor()
// }

// func (FieldDescriptorProto_Label) Type() protoreflect.EnumType {
// 	return &file_descriptor_proto_enumTypes[1]
// }

// func (x FieldDescriptorProto_Label) Number() protoreflect.EnumNumber {
// 	return protoreflect.EnumNumber(x)
// }

// // Deprecated: Do not use.
// func (x *FieldDescriptorProto_Label) UnmarshalJSON(b []byte) error {
// 	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
// 	if err != nil {
// 		return err
// 	}
// 	*x = FieldDescriptorProto_Label(num)
// 	return nil
// }

// // Deprecated: Use FieldDescriptorProto_Label.Descriptor instead.
// func (FieldDescriptorProto_Label) EnumDescriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{4, 1}
// }

// // Generated classes can be optimized for speed or code size.
// type FileOptions_OptimizeMode int32

// const (
// 	FileOptions_SPEED FileOptions_OptimizeMode = 1 // Generate complete code for parsing, serialization,
// 	// etc.
// 	FileOptions_CODE_SIZE    FileOptions_OptimizeMode = 2 // Use ReflectionOps to implement these methods.
// 	FileOptions_LITE_RUNTIME FileOptions_OptimizeMode = 3 // Generate code using MessageLite and the lite runtime.
// )

// // Enum value maps for FileOptions_OptimizeMode.
// var (
// 	FileOptions_OptimizeMode_name = map[int32]string{
// 		1: "SPEED",
// 		2: "CODE_SIZE",
// 		3: "LITE_RUNTIME",
// 	}
// 	FileOptions_OptimizeMode_value = map[string]int32{
// 		"SPEED":        1,
// 		"CODE_SIZE":    2,
// 		"LITE_RUNTIME": 3,
// 	}
// )

// func (x FileOptions_OptimizeMode) Enum() *FileOptions_OptimizeMode {
// 	p := new(FileOptions_OptimizeMode)
// 	*p = x
// 	return p
// }

// func (x FileOptions_OptimizeMode) String() string {
// 	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
// }

// func (FileOptions_OptimizeMode) Descriptor() protoreflect.EnumDescriptor {
// 	return file_descriptor_proto_enumTypes[2].Descriptor()
// }

// func (FileOptions_OptimizeMode) Type() protoreflect.EnumType {
// 	return &file_descriptor_proto_enumTypes[2]
// }

// func (x FileOptions_OptimizeMode) Number() protoreflect.EnumNumber {
// 	return protoreflect.EnumNumber(x)
// }

// // Deprecated: Do not use.
// func (x *FileOptions_OptimizeMode) UnmarshalJSON(b []byte) error {
// 	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
// 	if err != nil {
// 		return err
// 	}
// 	*x = FileOptions_OptimizeMode(num)
// 	return nil
// }

// // Deprecated: Use FileOptions_OptimizeMode.Descriptor instead.
// func (FileOptions_OptimizeMode) EnumDescriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{10, 0}
// }

// type FieldOptions_CType int32

// const (
// 	// Default mode.
// 	FieldOptions_STRING       FieldOptions_CType = 0
// 	FieldOptions_CORD         FieldOptions_CType = 1
// 	FieldOptions_STRING_PIECE FieldOptions_CType = 2
// )

// // Enum value maps for FieldOptions_CType.
// var (
// 	FieldOptions_CType_name = map[int32]string{
// 		0: "STRING",
// 		1: "CORD",
// 		2: "STRING_PIECE",
// 	}
// 	FieldOptions_CType_value = map[string]int32{
// 		"STRING":       0,
// 		"CORD":         1,
// 		"STRING_PIECE": 2,
// 	}
// )

// func (x FieldOptions_CType) Enum() *FieldOptions_CType {
// 	p := new(FieldOptions_CType)
// 	*p = x
// 	return p
// }

// func (x FieldOptions_CType) String() string {
// 	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
// }

// func (FieldOptions_CType) Descriptor() protoreflect.EnumDescriptor {
// 	return file_descriptor_proto_enumTypes[3].Descriptor()
// }

// func (FieldOptions_CType) Type() protoreflect.EnumType {
// 	return &file_descriptor_proto_enumTypes[3]
// }

// func (x FieldOptions_CType) Number() protoreflect.EnumNumber {
// 	return protoreflect.EnumNumber(x)
// }

// // Deprecated: Do not use.
// func (x *FieldOptions_CType) UnmarshalJSON(b []byte) error {
// 	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
// 	if err != nil {
// 		return err
// 	}
// 	*x = FieldOptions_CType(num)
// 	return nil
// }

// // Deprecated: Use FieldOptions_CType.Descriptor instead.
// func (FieldOptions_CType) EnumDescriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{12, 0}
// }

// type FieldOptions_JSType int32

// const (
// 	// Use the default type.
// 	FieldOptions_JS_NORMAL FieldOptions_JSType = 0
// 	// Use JavaScript strings.
// 	FieldOptions_JS_STRING FieldOptions_JSType = 1
// 	// Use JavaScript numbers.
// 	FieldOptions_JS_NUMBER FieldOptions_JSType = 2
// )

// // Enum value maps for FieldOptions_JSType.
// var (
// 	FieldOptions_JSType_name = map[int32]string{
// 		0: "JS_NORMAL",
// 		1: "JS_STRING",
// 		2: "JS_NUMBER",
// 	}
// 	FieldOptions_JSType_value = map[string]int32{
// 		"JS_NORMAL": 0,
// 		"JS_STRING": 1,
// 		"JS_NUMBER": 2,
// 	}
// )

// func (x FieldOptions_JSType) Enum() *FieldOptions_JSType {
// 	p := new(FieldOptions_JSType)
// 	*p = x
// 	return p
// }

// func (x FieldOptions_JSType) String() string {
// 	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
// }

// func (FieldOptions_JSType) Descriptor() protoreflect.EnumDescriptor {
// 	return file_descriptor_proto_enumTypes[4].Descriptor()
// }

// func (FieldOptions_JSType) Type() protoreflect.EnumType {
// 	return &file_descriptor_proto_enumTypes[4]
// }

// func (x FieldOptions_JSType) Number() protoreflect.EnumNumber {
// 	return protoreflect.EnumNumber(x)
// }

// // Deprecated: Do not use.
// func (x *FieldOptions_JSType) UnmarshalJSON(b []byte) error {
// 	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
// 	if err != nil {
// 		return err
// 	}
// 	*x = FieldOptions_JSType(num)
// 	return nil
// }

// // Deprecated: Use FieldOptions_JSType.Descriptor instead.
// func (FieldOptions_JSType) EnumDescriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{12, 1}
// }

// // Is this method side-effect-free (or safe in HTTP parlance), or idempotent,
// // or neither? HTTP based RPC implementation may choose GET verb for safe
// // methods, and PUT verb for idempotent methods instead of the default POST.
// type MethodOptions_IdempotencyLevel int32

// const (
// 	MethodOptions_IDEMPOTENCY_UNKNOWN MethodOptions_IdempotencyLevel = 0
// 	MethodOptions_NO_SIDE_EFFECTS     MethodOptions_IdempotencyLevel = 1 // implies idempotent
// 	MethodOptions_IDEMPOTENT          MethodOptions_IdempotencyLevel = 2 // idempotent, but may have side effects
// )

// // Enum value maps for MethodOptions_IdempotencyLevel.
// var (
// 	MethodOptions_IdempotencyLevel_name = map[int32]string{
// 		0: "IDEMPOTENCY_UNKNOWN",
// 		1: "NO_SIDE_EFFECTS",
// 		2: "IDEMPOTENT",
// 	}
// 	MethodOptions_IdempotencyLevel_value = map[string]int32{
// 		"IDEMPOTENCY_UNKNOWN": 0,
// 		"NO_SIDE_EFFECTS":     1,
// 		"IDEMPOTENT":          2,
// 	}
// )

// func (x MethodOptions_IdempotencyLevel) Enum() *MethodOptions_IdempotencyLevel {
// 	p := new(MethodOptions_IdempotencyLevel)
// 	*p = x
// 	return p
// }

// func (x MethodOptions_IdempotencyLevel) String() string {
// 	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
// }

// func (MethodOptions_IdempotencyLevel) Descriptor() protoreflect.EnumDescriptor {
// 	return file_descriptor_proto_enumTypes[5].Descriptor()
// }

// func (MethodOptions_IdempotencyLevel) Type() protoreflect.EnumType {
// 	return &file_descriptor_proto_enumTypes[5]
// }

// func (x MethodOptions_IdempotencyLevel) Number() protoreflect.EnumNumber {
// 	return protoreflect.EnumNumber(x)
// }

// // Deprecated: Do not use.
// func (x *MethodOptions_IdempotencyLevel) UnmarshalJSON(b []byte) error {
// 	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
// 	if err != nil {
// 		return err
// 	}
// 	*x = MethodOptions_IdempotencyLevel(num)
// 	return nil
// }

// // Deprecated: Use MethodOptions_IdempotencyLevel.Descriptor instead.
// func (MethodOptions_IdempotencyLevel) EnumDescriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{17, 0}
// }

// // The protocol compiler can output a FileDescriptorSet containing the .proto
// // files it parses.
// type FileDescriptorSet struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	File []*FileDescriptorProto `protobuf:"bytes,1,rep,name=file" json:"file,omitempty"`
// }

// func (x *FileDescriptorSet) Reset() {
// 	*x = FileDescriptorSet{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[0]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *FileDescriptorSet) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*FileDescriptorSet) ProtoMessage() {}

// func (x *FileDescriptorSet) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[0]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use FileDescriptorSet.ProtoReflect.Descriptor instead.
// func (*FileDescriptorSet) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{0}
// }

// func (x *FileDescriptorSet) GetFile() []*FileDescriptorProto {
// 	if x != nil {
// 		return x.File
// 	}
// 	return nil
// }

// // Describes a complete .proto file.
// type FileDescriptorProto struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	Name    *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`       // file name, relative to root of source tree
// 	Package *string `protobuf:"bytes,2,opt,name=package" json:"package,omitempty"` // e.g. "foo", "foo.bar", etc.
// 	// Names of files imported by this file.
// 	Dependency []string `protobuf:"bytes,3,rep,name=dependency" json:"dependency,omitempty"`
// 	// Indexes of the public imported files in the dependency list above.
// 	PublicDependency []int32 `protobuf:"varint,10,rep,name=public_dependency,json=publicDependency" json:"public_dependency,omitempty"`
// 	// Indexes of the weak imported files in the dependency list.
// 	// For Google-internal migration only. Do not use.
// 	WeakDependency []int32 `protobuf:"varint,11,rep,name=weak_dependency,json=weakDependency" json:"weak_dependency,omitempty"`
// 	// All top-level definitions in this file.
// 	MessageType []*DescriptorProto        `protobuf:"bytes,4,rep,name=message_type,json=messageType" json:"message_type,omitempty"`
// 	EnumType    []*EnumDescriptorProto    `protobuf:"bytes,5,rep,name=enum_type,json=enumType" json:"enum_type,omitempty"`
// 	Service     []*ServiceDescriptorProto `protobuf:"bytes,6,rep,name=service" json:"service,omitempty"`
// 	Extension   []*FieldDescriptorProto   `protobuf:"bytes,7,rep,name=extension" json:"extension,omitempty"`
// 	Options     *FileOptions              `protobuf:"bytes,8,opt,name=options" json:"options,omitempty"`
// 	// This field contains optional information about the original source code.
// 	// You may safely remove this entire field without harming runtime
// 	// functionality of the descriptors -- the information is needed only by
// 	// development tools.
// 	SourceCodeInfo *SourceCodeInfo `protobuf:"bytes,9,opt,name=source_code_info,json=sourceCodeInfo" json:"source_code_info,omitempty"`
// 	// The syntax of the proto file.
// 	// The supported values are "proto2" and "proto3".
// 	Syntax *string `protobuf:"bytes,12,opt,name=syntax" json:"syntax,omitempty"`
// }

// func (x *FileDescriptorProto) Reset() {
// 	*x = FileDescriptorProto{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[1]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *FileDescriptorProto) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*FileDescriptorProto) ProtoMessage() {}

// func (x *FileDescriptorProto) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[1]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use FileDescriptorProto.ProtoReflect.Descriptor instead.
// func (*FileDescriptorProto) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{1}
// }

// func (x *FileDescriptorProto) GetName() string {
// 	if x != nil && x.Name != nil {
// 		return *x.Name
// 	}
// 	return ""
// }

// func (x *FileDescriptorProto) GetPackage() string {
// 	if x != nil && x.Package != nil {
// 		return *x.Package
// 	}
// 	return ""
// }

// func (x *FileDescriptorProto) GetDependency() []string {
// 	if x != nil {
// 		return x.Dependency
// 	}
// 	return nil
// }

// func (x *FileDescriptorProto) GetPublicDependency() []int32 {
// 	if x != nil {
// 		return x.PublicDependency
// 	}
// 	return nil
// }

// func (x *FileDescriptorProto) GetWeakDependency() []int32 {
// 	if x != nil {
// 		return x.WeakDependency
// 	}
// 	return nil
// }

// func (x *FileDescriptorProto) GetMessageType() []*DescriptorProto {
// 	if x != nil {
// 		return x.MessageType
// 	}
// 	return nil
// }

// func (x *FileDescriptorProto) GetEnumType() []*EnumDescriptorProto {
// 	if x != nil {
// 		return x.EnumType
// 	}
// 	return nil
// }

// func (x *FileDescriptorProto) GetService() []*ServiceDescriptorProto {
// 	if x != nil {
// 		return x.Service
// 	}
// 	return nil
// }

// func (x *FileDescriptorProto) GetExtension() []*FieldDescriptorProto {
// 	if x != nil {
// 		return x.Extension
// 	}
// 	return nil
// }

// func (x *FileDescriptorProto) GetOptions() *FileOptions {
// 	if x != nil {
// 		return x.Options
// 	}
// 	return nil
// }

// func (x *FileDescriptorProto) GetSourceCodeInfo() *SourceCodeInfo {
// 	if x != nil {
// 		return x.SourceCodeInfo
// 	}
// 	return nil
// }

// func (x *FileDescriptorProto) GetSyntax() string {
// 	if x != nil && x.Syntax != nil {
// 		return *x.Syntax
// 	}
// 	return ""
// }

// // Describes a message type.
// type DescriptorProto struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	Name           *string                           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
// 	Field          []*FieldDescriptorProto           `protobuf:"bytes,2,rep,name=field" json:"field,omitempty"`
// 	Extension      []*FieldDescriptorProto           `protobuf:"bytes,6,rep,name=extension" json:"extension,omitempty"`
// 	NestedType     []*DescriptorProto                `protobuf:"bytes,3,rep,name=nested_type,json=nestedType" json:"nested_type,omitempty"`
// 	EnumType       []*EnumDescriptorProto            `protobuf:"bytes,4,rep,name=enum_type,json=enumType" json:"enum_type,omitempty"`
// 	ExtensionRange []*DescriptorProto_ExtensionRange `protobuf:"bytes,5,rep,name=extension_range,json=extensionRange" json:"extension_range,omitempty"`
// 	OneofDecl      []*OneofDescriptorProto           `protobuf:"bytes,8,rep,name=oneof_decl,json=oneofDecl" json:"oneof_decl,omitempty"`
// 	Options        *MessageOptions                   `protobuf:"bytes,7,opt,name=options" json:"options,omitempty"`
// 	ReservedRange  []*DescriptorProto_ReservedRange  `protobuf:"bytes,9,rep,name=reserved_range,json=reservedRange" json:"reserved_range,omitempty"`
// 	// Reserved field names, which may not be used by fields in the same message.
// 	// A given name may only be reserved once.
// 	ReservedName []string `protobuf:"bytes,10,rep,name=reserved_name,json=reservedName" json:"reserved_name,omitempty"`
// }

// func (x *DescriptorProto) Reset() {
// 	*x = DescriptorProto{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[2]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *DescriptorProto) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*DescriptorProto) ProtoMessage() {}

// func (x *DescriptorProto) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[2]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use DescriptorProto.ProtoReflect.Descriptor instead.
// func (*DescriptorProto) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{2}
// }

// func (x *DescriptorProto) GetName() string {
// 	if x != nil && x.Name != nil {
// 		return *x.Name
// 	}
// 	return ""
// }

// func (x *DescriptorProto) GetField() []*FieldDescriptorProto {
// 	if x != nil {
// 		return x.Field
// 	}
// 	return nil
// }

// func (x *DescriptorProto) GetExtension() []*FieldDescriptorProto {
// 	if x != nil {
// 		return x.Extension
// 	}
// 	return nil
// }

// func (x *DescriptorProto) GetNestedType() []*DescriptorProto {
// 	if x != nil {
// 		return x.NestedType
// 	}
// 	return nil
// }

// func (x *DescriptorProto) GetEnumType() []*EnumDescriptorProto {
// 	if x != nil {
// 		return x.EnumType
// 	}
// 	return nil
// }

// func (x *DescriptorProto) GetExtensionRange() []*DescriptorProto_ExtensionRange {
// 	if x != nil {
// 		return x.ExtensionRange
// 	}
// 	return nil
// }

// func (x *DescriptorProto) GetOneofDecl() []*OneofDescriptorProto {
// 	if x != nil {
// 		return x.OneofDecl
// 	}
// 	return nil
// }

// func (x *DescriptorProto) GetOptions() *MessageOptions {
// 	if x != nil {
// 		return x.Options
// 	}
// 	return nil
// }

// func (x *DescriptorProto) GetReservedRange() []*DescriptorProto_ReservedRange {
// 	if x != nil {
// 		return x.ReservedRange
// 	}
// 	return nil
// }

// func (x *DescriptorProto) GetReservedName() []string {
// 	if x != nil {
// 		return x.ReservedName
// 	}
// 	return nil
// }

// type ExtensionRangeOptions struct {
// 	state           protoimpl.MessageState
// 	sizeCache       protoimpl.SizeCache
// 	unknownFields   protoimpl.UnknownFields
// 	extensionFields protoimpl.ExtensionFields

// 	// The parser stores options it doesn't recognize here. See above.
// 	UninterpretedOption []*UninterpretedOption `protobuf:"bytes,999,rep,name=uninterpreted_option,json=uninterpretedOption" json:"uninterpreted_option,omitempty"`
// }

// func (x *ExtensionRangeOptions) Reset() {
// 	*x = ExtensionRangeOptions{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[3]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *ExtensionRangeOptions) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*ExtensionRangeOptions) ProtoMessage() {}

// func (x *ExtensionRangeOptions) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[3]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use ExtensionRangeOptions.ProtoReflect.Descriptor instead.
// func (*ExtensionRangeOptions) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{3}
// }

// func (x *ExtensionRangeOptions) GetUninterpretedOption() []*UninterpretedOption {
// 	if x != nil {
// 		return x.UninterpretedOption
// 	}
// 	return nil
// }

// // Describes a field within a message.
// type FieldDescriptorProto struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	Name   *string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
// 	Number *int32                      `protobuf:"varint,3,opt,name=number" json:"number,omitempty"`
// 	Label  *FieldDescriptorProto_Label `protobuf:"varint,4,opt,name=label,enum=google.protobuf.FieldDescriptorProto_Label" json:"label,omitempty"`
// 	// If type_name is set, this need not be set.  If both this and type_name
// 	// are set, this must be one of TYPE_ENUM, TYPE_MESSAGE or TYPE_GROUP.
// 	Type *FieldDescriptorProto_Type `protobuf:"varint,5,opt,name=type,enum=google.protobuf.FieldDescriptorProto_Type" json:"type,omitempty"`
// 	// For message and enum types, this is the name of the type.  If the name
// 	// starts with a '.', it is fully-qualified.  Otherwise, C++-like scoping
// 	// rules are used to find the type (i.e. first the nested types within this
// 	// message are searched, then within the parent, on up to the root
// 	// namespace).
// 	TypeName *string `protobuf:"bytes,6,opt,name=type_name,json=typeName" json:"type_name,omitempty"`
// 	// For extensions, this is the name of the type being extended.  It is
// 	// resolved in the same manner as type_name.
// 	Extendee *string `protobuf:"bytes,2,opt,name=extendee" json:"extendee,omitempty"`
// 	// For numeric types, contains the original text representation of the value.
// 	// For booleans, "true" or "false".
// 	// For strings, contains the default text contents (not escaped in any way).
// 	// For bytes, contains the C escaped value.  All bytes >= 128 are escaped.
// 	// TODO(kenton):  Base-64 encode?
// 	DefaultValue *string `protobuf:"bytes,7,opt,name=default_value,json=defaultValue" json:"default_value,omitempty"`
// 	// If set, gives the index of a oneof in the containing type's oneof_decl
// 	// list.  This field is a member of that oneof.
// 	OneofIndex *int32 `protobuf:"varint,9,opt,name=oneof_index,json=oneofIndex" json:"oneof_index,omitempty"`
// 	// JSON name of this field. The value is set by protocol compiler. If the
// 	// user has set a "json_name" option on this field, that option's value
// 	// will be used. Otherwise, it's deduced from the field's name by converting
// 	// it to camelCase.
// 	JsonName *string       `protobuf:"bytes,10,opt,name=json_name,json=jsonName" json:"json_name,omitempty"`
// 	Options  *FieldOptions `protobuf:"bytes,8,opt,name=options" json:"options,omitempty"`
// }

// func (x *FieldDescriptorProto) Reset() {
// 	*x = FieldDescriptorProto{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[4]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *FieldDescriptorProto) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*FieldDescriptorProto) ProtoMessage() {}

// func (x *FieldDescriptorProto) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[4]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use FieldDescriptorProto.ProtoReflect.Descriptor instead.
// func (*FieldDescriptorProto) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{4}
// }

// func (x *FieldDescriptorProto) GetName() string {
// 	if x != nil && x.Name != nil {
// 		return *x.Name
// 	}
// 	return ""
// }

// func (x *FieldDescriptorProto) GetNumber() int32 {
// 	if x != nil && x.Number != nil {
// 		return *x.Number
// 	}
// 	return 0
// }

// func (x *FieldDescriptorProto) GetLabel() FieldDescriptorProto_Label {
// 	if x != nil && x.Label != nil {
// 		return *x.Label
// 	}
// 	return FieldDescriptorProto_LABEL_OPTIONAL
// }

// func (x *FieldDescriptorProto) GetType() FieldDescriptorProto_Type {
// 	if x != nil && x.Type != nil {
// 		return *x.Type
// 	}
// 	return FieldDescriptorProto_TYPE_DOUBLE
// }

// func (x *FieldDescriptorProto) GetTypeName() string {
// 	if x != nil && x.TypeName != nil {
// 		return *x.TypeName
// 	}
// 	return ""
// }

// func (x *FieldDescriptorProto) GetExtendee() string {
// 	if x != nil && x.Extendee != nil {
// 		return *x.Extendee
// 	}
// 	return ""
// }

// func (x *FieldDescriptorProto) GetDefaultValue() string {
// 	if x != nil && x.DefaultValue != nil {
// 		return *x.DefaultValue
// 	}
// 	return ""
// }

// func (x *FieldDescriptorProto) GetOneofIndex() int32 {
// 	if x != nil && x.OneofIndex != nil {
// 		return *x.OneofIndex
// 	}
// 	return 0
// }

// func (x *FieldDescriptorProto) GetJsonName() string {
// 	if x != nil && x.JsonName != nil {
// 		return *x.JsonName
// 	}
// 	return ""
// }

// func (x *FieldDescriptorProto) GetOptions() *FieldOptions {
// 	if x != nil {
// 		return x.Options
// 	}
// 	return nil
// }

// // Describes a oneof.
// type OneofDescriptorProto struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	Name    *string       `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
// 	Options *OneofOptions `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
// }

// func (x *OneofDescriptorProto) Reset() {
// 	*x = OneofDescriptorProto{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[5]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *OneofDescriptorProto) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*OneofDescriptorProto) ProtoMessage() {}

// func (x *OneofDescriptorProto) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[5]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use OneofDescriptorProto.ProtoReflect.Descriptor instead.
// func (*OneofDescriptorProto) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{5}
// }

// func (x *OneofDescriptorProto) GetName() string {
// 	if x != nil && x.Name != nil {
// 		return *x.Name
// 	}
// 	return ""
// }

// func (x *OneofDescriptorProto) GetOptions() *OneofOptions {
// 	if x != nil {
// 		return x.Options
// 	}
// 	return nil
// }

// // Describes an enum type.
// type EnumDescriptorProto struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	Name    *string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
// 	Value   []*EnumValueDescriptorProto `protobuf:"bytes,2,rep,name=value" json:"value,omitempty"`
// 	Options *EnumOptions                `protobuf:"bytes,3,opt,name=options" json:"options,omitempty"`
// 	// Range of reserved numeric values. Reserved numeric values may not be used
// 	// by enum values in the same enum declaration. Reserved ranges may not
// 	// overlap.
// 	ReservedRange []*EnumDescriptorProto_EnumReservedRange `protobuf:"bytes,4,rep,name=reserved_range,json=reservedRange" json:"reserved_range,omitempty"`
// 	// Reserved enum value names, which may not be reused. A given name may only
// 	// be reserved once.
// 	ReservedName []string `protobuf:"bytes,5,rep,name=reserved_name,json=reservedName" json:"reserved_name,omitempty"`
// }

// func (x *EnumDescriptorProto) Reset() {
// 	*x = EnumDescriptorProto{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[6]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *EnumDescriptorProto) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*EnumDescriptorProto) ProtoMessage() {}

// func (x *EnumDescriptorProto) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[6]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use EnumDescriptorProto.ProtoReflect.Descriptor instead.
// func (*EnumDescriptorProto) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{6}
// }

// func (x *EnumDescriptorProto) GetName() string {
// 	if x != nil && x.Name != nil {
// 		return *x.Name
// 	}
// 	return ""
// }

// func (x *EnumDescriptorProto) GetValue() []*EnumValueDescriptorProto {
// 	if x != nil {
// 		return x.Value
// 	}
// 	return nil
// }

// func (x *EnumDescriptorProto) GetOptions() *EnumOptions {
// 	if x != nil {
// 		return x.Options
// 	}
// 	return nil
// }

// func (x *EnumDescriptorProto) GetReservedRange() []*EnumDescriptorProto_EnumReservedRange {
// 	if x != nil {
// 		return x.ReservedRange
// 	}
// 	return nil
// }

// func (x *EnumDescriptorProto) GetReservedName() []string {
// 	if x != nil {
// 		return x.ReservedName
// 	}
// 	return nil
// }

// // Describes a value within an enum.
// type EnumValueDescriptorProto struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	Name    *string           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
// 	Number  *int32            `protobuf:"varint,2,opt,name=number" json:"number,omitempty"`
// 	Options *EnumValueOptions `protobuf:"bytes,3,opt,name=options" json:"options,omitempty"`
// }

// func (x *EnumValueDescriptorProto) Reset() {
// 	*x = EnumValueDescriptorProto{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[7]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *EnumValueDescriptorProto) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*EnumValueDescriptorProto) ProtoMessage() {}

// func (x *EnumValueDescriptorProto) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[7]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use EnumValueDescriptorProto.ProtoReflect.Descriptor instead.
// func (*EnumValueDescriptorProto) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{7}
// }

// func (x *EnumValueDescriptorProto) GetName() string {
// 	if x != nil && x.Name != nil {
// 		return *x.Name
// 	}
// 	return ""
// }

// func (x *EnumValueDescriptorProto) GetNumber() int32 {
// 	if x != nil && x.Number != nil {
// 		return *x.Number
// 	}
// 	return 0
// }

// func (x *EnumValueDescriptorProto) GetOptions() *EnumValueOptions {
// 	if x != nil {
// 		return x.Options
// 	}
// 	return nil
// }

// // Describes a service.
// type ServiceDescriptorProto struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	Name    *string                  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
// 	Method  []*MethodDescriptorProto `protobuf:"bytes,2,rep,name=method" json:"method,omitempty"`
// 	Options *ServiceOptions          `protobuf:"bytes,3,opt,name=options" json:"options,omitempty"`
// }

// func (x *ServiceDescriptorProto) Reset() {
// 	*x = ServiceDescriptorProto{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[8]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *ServiceDescriptorProto) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*ServiceDescriptorProto) ProtoMessage() {}

// func (x *ServiceDescriptorProto) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[8]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use ServiceDescriptorProto.ProtoReflect.Descriptor instead.
// func (*ServiceDescriptorProto) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{8}
// }

// func (x *ServiceDescriptorProto) GetName() string {
// 	if x != nil && x.Name != nil {
// 		return *x.Name
// 	}
// 	return ""
// }

// func (x *ServiceDescriptorProto) GetMethod() []*MethodDescriptorProto {
// 	if x != nil {
// 		return x.Method
// 	}
// 	return nil
// }

// func (x *ServiceDescriptorProto) GetOptions() *ServiceOptions {
// 	if x != nil {
// 		return x.Options
// 	}
// 	return nil
// }

// // Describes a method of a service.
// type MethodDescriptorProto struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
// 	// Input and output type names.  These are resolved in the same way as
// 	// FieldDescriptorProto.type_name, but must refer to a message type.
// 	InputType  *string        `protobuf:"bytes,2,opt,name=input_type,json=inputType" json:"input_type,omitempty"`
// 	OutputType *string        `protobuf:"bytes,3,opt,name=output_type,json=outputType" json:"output_type,omitempty"`
// 	Options    *MethodOptions `protobuf:"bytes,4,opt,name=options" json:"options,omitempty"`
// 	// Identifies if client streams multiple client messages
// 	ClientStreaming *bool `protobuf:"varint,5,opt,name=client_streaming,json=clientStreaming,def=0" json:"client_streaming,omitempty"`
// 	// Identifies if server streams multiple server messages
// 	ServerStreaming *bool `protobuf:"varint,6,opt,name=server_streaming,json=serverStreaming,def=0" json:"server_streaming,omitempty"`
// }

// // Default values for MethodDescriptorProto fields.
// const (
// 	Default_MethodDescriptorProto_ClientStreaming = bool(false)
// 	Default_MethodDescriptorProto_ServerStreaming = bool(false)
// )

// func (x *MethodDescriptorProto) Reset() {
// 	*x = MethodDescriptorProto{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[9]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *MethodDescriptorProto) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*MethodDescriptorProto) ProtoMessage() {}

// func (x *MethodDescriptorProto) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[9]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use MethodDescriptorProto.ProtoReflect.Descriptor instead.
// func (*MethodDescriptorProto) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{9}
// }

// func (x *MethodDescriptorProto) GetName() string {
// 	if x != nil && x.Name != nil {
// 		return *x.Name
// 	}
// 	return ""
// }

// func (x *MethodDescriptorProto) GetInputType() string {
// 	if x != nil && x.InputType != nil {
// 		return *x.InputType
// 	}
// 	return ""
// }

// func (x *MethodDescriptorProto) GetOutputType() string {
// 	if x != nil && x.OutputType != nil {
// 		return *x.OutputType
// 	}
// 	return ""
// }

// func (x *MethodDescriptorProto) GetOptions() *MethodOptions {
// 	if x != nil {
// 		return x.Options
// 	}
// 	return nil
// }

// func (x *MethodDescriptorProto) GetClientStreaming() bool {
// 	if x != nil && x.ClientStreaming != nil {
// 		return *x.ClientStreaming
// 	}
// 	return Default_MethodDescriptorProto_ClientStreaming
// }

// func (x *MethodDescriptorProto) GetServerStreaming() bool {
// 	if x != nil && x.ServerStreaming != nil {
// 		return *x.ServerStreaming
// 	}
// 	return Default_MethodDescriptorProto_ServerStreaming
// }

// type FileOptions struct {
// 	state           protoimpl.MessageState
// 	sizeCache       protoimpl.SizeCache
// 	unknownFields   protoimpl.UnknownFields
// 	extensionFields protoimpl.ExtensionFields

// 	// Sets the Java package where classes generated from this .proto will be
// 	// placed.  By default, the proto package is used, but this is often
// 	// inappropriate because proto packages do not normally start with backwards
// 	// domain names.
// 	JavaPackage *string `protobuf:"bytes,1,opt,name=java_package,json=javaPackage" json:"java_package,omitempty"`
// 	// If set, all the classes from the .proto file are wrapped in a single
// 	// outer class with the given name.  This applies to both Proto1
// 	// (equivalent to the old "--one_java_file" option) and Proto2 (where
// 	// a .proto always translates to a single class, but you may want to
// 	// explicitly choose the class name).
// 	JavaOuterClassname *string `protobuf:"bytes,8,opt,name=java_outer_classname,json=javaOuterClassname" json:"java_outer_classname,omitempty"`
// 	// If set true, then the Java code generator will generate a separate .java
// 	// file for each top-level message, enum, and service defined in the .proto
// 	// file.  Thus, these types will *not* be nested inside the outer class
// 	// named by java_outer_classname.  However, the outer class will still be
// 	// generated to contain the file's getDescriptor() method as well as any
// 	// top-level extensions defined in the file.
// 	JavaMultipleFiles *bool `protobuf:"varint,10,opt,name=java_multiple_files,json=javaMultipleFiles,def=0" json:"java_multiple_files,omitempty"`
// 	// This option does nothing.
// 	//
// 	// Deprecated: Do not use.
// 	JavaGenerateEqualsAndHash *bool `protobuf:"varint,20,opt,name=java_generate_equals_and_hash,json=javaGenerateEqualsAndHash" json:"java_generate_equals_and_hash,omitempty"`
// 	// If set true, then the Java2 code generator will generate code that
// 	// throws an exception whenever an attempt is made to assign a non-UTF-8
// 	// byte sequence to a string field.
// 	// Message reflection will do the same.
// 	// However, an extension field still accepts non-UTF-8 byte sequences.
// 	// This option has no effect on when used with the lite runtime.
// 	JavaStringCheckUtf8 *bool                     `protobuf:"varint,27,opt,name=java_string_check_utf8,json=javaStringCheckUtf8,def=0" json:"java_string_check_utf8,omitempty"`
// 	OptimizeFor         *FileOptions_OptimizeMode `protobuf:"varint,9,opt,name=optimize_for,json=optimizeFor,enum=google.protobuf.FileOptions_OptimizeMode,def=1" json:"optimize_for,omitempty"`
// 	// Sets the Go package where structs generated from this .proto will be
// 	// placed. If omitted, the Go package will be derived from the following:
// 	//   - The basename of the package import path, if provided.
// 	//   - Otherwise, the package statement in the .proto file, if present.
// 	//   - Otherwise, the basename of the .proto file, without extension.
// 	GoPackage *string `protobuf:"bytes,11,opt,name=go_package,json=goPackage" json:"go_package,omitempty"`
// 	// Should generic services be generated in each language?  "Generic" services
// 	// are not specific to any particular RPC system.  They are generated by the
// 	// main code generators in each language (without additional plugins).
// 	// Generic services were the only kind of service generation supported by
// 	// early versions of google.protobuf.
// 	//
// 	// Generic services are now considered deprecated in favor of using plugins
// 	// that generate code specific to your particular RPC system.  Therefore,
// 	// these default to false.  Old code which depends on generic services should
// 	// explicitly set them to true.
// 	CcGenericServices   *bool `protobuf:"varint,16,opt,name=cc_generic_services,json=ccGenericServices,def=0" json:"cc_generic_services,omitempty"`
// 	JavaGenericServices *bool `protobuf:"varint,17,opt,name=java_generic_services,json=javaGenericServices,def=0" json:"java_generic_services,omitempty"`
// 	PyGenericServices   *bool `protobuf:"varint,18,opt,name=py_generic_services,json=pyGenericServices,def=0" json:"py_generic_services,omitempty"`
// 	PhpGenericServices  *bool `protobuf:"varint,42,opt,name=php_generic_services,json=phpGenericServices,def=0" json:"php_generic_services,omitempty"`
// 	// Is this file deprecated?
// 	// Depending on the target platform, this can emit Deprecated annotations
// 	// for everything in the file, or it will be completely ignored; in the very
// 	// least, this is a formalization for deprecating files.
// 	Deprecated *bool `protobuf:"varint,23,opt,name=deprecated,def=0" json:"deprecated,omitempty"`
// 	// Enables the use of arenas for the proto messages in this file. This applies
// 	// only to generated classes for C++.
// 	CcEnableArenas *bool `protobuf:"varint,31,opt,name=cc_enable_arenas,json=ccEnableArenas,def=0" json:"cc_enable_arenas,omitempty"`
// 	// Sets the objective c class prefix which is prepended to all objective c
// 	// generated classes from this .proto. There is no default.
// 	ObjcClassPrefix *string `protobuf:"bytes,36,opt,name=objc_class_prefix,json=objcClassPrefix" json:"objc_class_prefix,omitempty"`
// 	// Namespace for generated classes; defaults to the package.
// 	CsharpNamespace *string `protobuf:"bytes,37,opt,name=csharp_namespace,json=csharpNamespace" json:"csharp_namespace,omitempty"`
// 	// By default Swift generators will take the proto package and CamelCase it
// 	// replacing '.' with underscore and use that to prefix the types/symbols
// 	// defined. When this options is provided, they will use this value instead
// 	// to prefix the types/symbols defined.
// 	SwiftPrefix *string `protobuf:"bytes,39,opt,name=swift_prefix,json=swiftPrefix" json:"swift_prefix,omitempty"`
// 	// Sets the php class prefix which is prepended to all php generated classes
// 	// from this .proto. Default is empty.
// 	PhpClassPrefix *string `protobuf:"bytes,40,opt,name=php_class_prefix,json=phpClassPrefix" json:"php_class_prefix,omitempty"`
// 	// Use this option to change the namespace of php generated classes. Default
// 	// is empty. When this option is empty, the package name will be used for
// 	// determining the namespace.
// 	PhpNamespace *string `protobuf:"bytes,41,opt,name=php_namespace,json=phpNamespace" json:"php_namespace,omitempty"`
// 	// The parser stores options it doesn't recognize here.
// 	// See the documentation for the "Options" section above.
// 	UninterpretedOption []*UninterpretedOption `protobuf:"bytes,999,rep,name=uninterpreted_option,json=uninterpretedOption" json:"uninterpreted_option,omitempty"`
// }

// // Default values for FileOptions fields.
// const (
// 	Default_FileOptions_JavaMultipleFiles   = bool(false)
// 	Default_FileOptions_JavaStringCheckUtf8 = bool(false)
// 	Default_FileOptions_OptimizeFor         = FileOptions_SPEED
// 	Default_FileOptions_CcGenericServices   = bool(false)
// 	Default_FileOptions_JavaGenericServices = bool(false)
// 	Default_FileOptions_PyGenericServices   = bool(false)
// 	Default_FileOptions_PhpGenericServices  = bool(false)
// 	Default_FileOptions_Deprecated          = bool(false)
// 	Default_FileOptions_CcEnableArenas      = bool(false)
// )

// func (x *FileOptions) Reset() {
// 	*x = FileOptions{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[10]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *FileOptions) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*FileOptions) ProtoMessage() {}

// func (x *FileOptions) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[10]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use FileOptions.ProtoReflect.Descriptor instead.
// func (*FileOptions) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{10}
// }

// func (x *FileOptions) GetJavaPackage() string {
// 	if x != nil && x.JavaPackage != nil {
// 		return *x.JavaPackage
// 	}
// 	return ""
// }

// func (x *FileOptions) GetJavaOuterClassname() string {
// 	if x != nil && x.JavaOuterClassname != nil {
// 		return *x.JavaOuterClassname
// 	}
// 	return ""
// }

// func (x *FileOptions) GetJavaMultipleFiles() bool {
// 	if x != nil && x.JavaMultipleFiles != nil {
// 		return *x.JavaMultipleFiles
// 	}
// 	return Default_FileOptions_JavaMultipleFiles
// }

// // Deprecated: Do not use.
// func (x *FileOptions) GetJavaGenerateEqualsAndHash() bool {
// 	if x != nil && x.JavaGenerateEqualsAndHash != nil {
// 		return *x.JavaGenerateEqualsAndHash
// 	}
// 	return false
// }

// func (x *FileOptions) GetJavaStringCheckUtf8() bool {
// 	if x != nil && x.JavaStringCheckUtf8 != nil {
// 		return *x.JavaStringCheckUtf8
// 	}
// 	return Default_FileOptions_JavaStringCheckUtf8
// }

// func (x *FileOptions) GetOptimizeFor() FileOptions_OptimizeMode {
// 	if x != nil && x.OptimizeFor != nil {
// 		return *x.OptimizeFor
// 	}
// 	return Default_FileOptions_OptimizeFor
// }

// func (x *FileOptions) GetGoPackage() string {
// 	if x != nil && x.GoPackage != nil {
// 		return *x.GoPackage
// 	}
// 	return ""
// }

// func (x *FileOptions) GetCcGenericServices() bool {
// 	if x != nil && x.CcGenericServices != nil {
// 		return *x.CcGenericServices
// 	}
// 	return Default_FileOptions_CcGenericServices
// }

// func (x *FileOptions) GetJavaGenericServices() bool {
// 	if x != nil && x.JavaGenericServices != nil {
// 		return *x.JavaGenericServices
// 	}
// 	return Default_FileOptions_JavaGenericServices
// }

// func (x *FileOptions) GetPyGenericServices() bool {
// 	if x != nil && x.PyGenericServices != nil {
// 		return *x.PyGenericServices
// 	}
// 	return Default_FileOptions_PyGenericServices
// }

// func (x *FileOptions) GetPhpGenericServices() bool {
// 	if x != nil && x.PhpGenericServices != nil {
// 		return *x.PhpGenericServices
// 	}
// 	return Default_FileOptions_PhpGenericServices
// }

// func (x *FileOptions) GetDeprecated() bool {
// 	if x != nil && x.Deprecated != nil {
// 		return *x.Deprecated
// 	}
// 	return Default_FileOptions_Deprecated
// }

// func (x *FileOptions) GetCcEnableArenas() bool {
// 	if x != nil && x.CcEnableArenas != nil {
// 		return *x.CcEnableArenas
// 	}
// 	return Default_FileOptions_CcEnableArenas
// }

// func (x *FileOptions) GetObjcClassPrefix() string {
// 	if x != nil && x.ObjcClassPrefix != nil {
// 		return *x.ObjcClassPrefix
// 	}
// 	return ""
// }

// func (x *FileOptions) GetCsharpNamespace() string {
// 	if x != nil && x.CsharpNamespace != nil {
// 		return *x.CsharpNamespace
// 	}
// 	return ""
// }

// func (x *FileOptions) GetSwiftPrefix() string {
// 	if x != nil && x.SwiftPrefix != nil {
// 		return *x.SwiftPrefix
// 	}
// 	return ""
// }

// func (x *FileOptions) GetPhpClassPrefix() string {
// 	if x != nil && x.PhpClassPrefix != nil {
// 		return *x.PhpClassPrefix
// 	}
// 	return ""
// }

// func (x *FileOptions) GetPhpNamespace() string {
// 	if x != nil && x.PhpNamespace != nil {
// 		return *x.PhpNamespace
// 	}
// 	return ""
// }

// func (x *FileOptions) GetUninterpretedOption() []*UninterpretedOption {
// 	if x != nil {
// 		return x.UninterpretedOption
// 	}
// 	return nil
// }

// type MessageOptions struct {
// 	state           protoimpl.MessageState
// 	sizeCache       protoimpl.SizeCache
// 	unknownFields   protoimpl.UnknownFields
// 	extensionFields protoimpl.ExtensionFields

// 	// Set true to use the old proto1 MessageSet wire format for extensions.
// 	// This is provided for backwards-compatibility with the MessageSet wire
// 	// format.  You should not use this for any other reason:  It's less
// 	// efficient, has fewer features, and is more complicated.
// 	//
// 	// The message must be defined exactly as follows:
// 	//   message Foo {
// 	//     option message_set_wire_format = true;
// 	//     extensions 4 to max;
// 	//   }
// 	// Note that the message cannot have any defined fields; MessageSets only
// 	// have extensions.
// 	//
// 	// All extensions of your type must be singular messages; e.g. they cannot
// 	// be int32s, enums, or repeated messages.
// 	//
// 	// Because this is an option, the above two restrictions are not enforced by
// 	// the protocol compiler.
// 	MessageSetWireFormat *bool `protobuf:"varint,1,opt,name=message_set_wire_format,json=messageSetWireFormat,def=0" json:"message_set_wire_format,omitempty"`
// 	// Disables the generation of the standard "descriptor()" accessor, which can
// 	// conflict with a field of the same name.  This is meant to make migration
// 	// from proto1 easier; new code should avoid fields named "descriptor".
// 	NoStandardDescriptorAccessor *bool `protobuf:"varint,2,opt,name=no_standard_descriptor_accessor,json=noStandardDescriptorAccessor,def=0" json:"no_standard_descriptor_accessor,omitempty"`
// 	// Is this message deprecated?
// 	// Depending on the target platform, this can emit Deprecated annotations
// 	// for the message, or it will be completely ignored; in the very least,
// 	// this is a formalization for deprecating messages.
// 	Deprecated *bool `protobuf:"varint,3,opt,name=deprecated,def=0" json:"deprecated,omitempty"`
// 	// Whether the message is an automatically generated map entry type for the
// 	// maps field.
// 	//
// 	// For maps fields:
// 	//     map<KeyType, ValueType> map_field = 1;
// 	// The parsed descriptor looks like:
// 	//     message MapFieldEntry {
// 	//         option map_entry = true;
// 	//         optional KeyType key = 1;
// 	//         optional ValueType value = 2;
// 	//     }
// 	//     repeated MapFieldEntry map_field = 1;
// 	//
// 	// Implementations may choose not to generate the map_entry=true message, but
// 	// use a native map in the target language to hold the keys and values.
// 	// The reflection APIs in such implementions still need to work as
// 	// if the field is a repeated message field.
// 	//
// 	// NOTE: Do not set the option in .proto files. Always use the maps syntax
// 	// instead. The option should only be implicitly set by the proto compiler
// 	// parser.
// 	MapEntry *bool `protobuf:"varint,7,opt,name=map_entry,json=mapEntry" json:"map_entry,omitempty"`
// 	// The parser stores options it doesn't recognize here. See above.
// 	UninterpretedOption []*UninterpretedOption `protobuf:"bytes,999,rep,name=uninterpreted_option,json=uninterpretedOption" json:"uninterpreted_option,omitempty"`
// }

// // Default values for MessageOptions fields.
// const (
// 	Default_MessageOptions_MessageSetWireFormat         = bool(false)
// 	Default_MessageOptions_NoStandardDescriptorAccessor = bool(false)
// 	Default_MessageOptions_Deprecated                   = bool(false)
// )

// func (x *MessageOptions) Reset() {
// 	*x = MessageOptions{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[11]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *MessageOptions) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*MessageOptions) ProtoMessage() {}

// func (x *MessageOptions) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[11]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use MessageOptions.ProtoReflect.Descriptor instead.
// func (*MessageOptions) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{11}
// }

// func (x *MessageOptions) GetMessageSetWireFormat() bool {
// 	if x != nil && x.MessageSetWireFormat != nil {
// 		return *x.MessageSetWireFormat
// 	}
// 	return Default_MessageOptions_MessageSetWireFormat
// }

// func (x *MessageOptions) GetNoStandardDescriptorAccessor() bool {
// 	if x != nil && x.NoStandardDescriptorAccessor != nil {
// 		return *x.NoStandardDescriptorAccessor
// 	}
// 	return Default_MessageOptions_NoStandardDescriptorAccessor
// }

// func (x *MessageOptions) GetDeprecated() bool {
// 	if x != nil && x.Deprecated != nil {
// 		return *x.Deprecated
// 	}
// 	return Default_MessageOptions_Deprecated
// }

// func (x *MessageOptions) GetMapEntry() bool {
// 	if x != nil && x.MapEntry != nil {
// 		return *x.MapEntry
// 	}
// 	return false
// }

// func (x *MessageOptions) GetUninterpretedOption() []*UninterpretedOption {
// 	if x != nil {
// 		return x.UninterpretedOption
// 	}
// 	return nil
// }

// type FieldOptions struct {
// 	state           protoimpl.MessageState
// 	sizeCache       protoimpl.SizeCache
// 	unknownFields   protoimpl.UnknownFields
// 	extensionFields protoimpl.ExtensionFields

// 	// The ctype option instructs the C++ code generator to use a different
// 	// representation of the field than it normally would.  See the specific
// 	// options below.  This option is not yet implemented in the open source
// 	// release -- sorry, we'll try to include it in a future version!
// 	Ctype *FieldOptions_CType `protobuf:"varint,1,opt,name=ctype,enum=google.protobuf.FieldOptions_CType,def=0" json:"ctype,omitempty"`
// 	// The packed option can be enabled for repeated primitive fields to enable
// 	// a more efficient representation on the wire. Rather than repeatedly
// 	// writing the tag and type for each element, the entire array is encoded as
// 	// a single length-delimited blob. In proto3, only explicit setting it to
// 	// false will avoid using packed encoding.
// 	Packed *bool `protobuf:"varint,2,opt,name=packed" json:"packed,omitempty"`
// 	// The jstype option determines the JavaScript type used for values of the
// 	// field.  The option is permitted only for 64 bit integral and fixed types
// 	// (int64, uint64, sint64, fixed64, sfixed64).  A field with jstype JS_STRING
// 	// is represented as JavaScript string, which avoids loss of precision that
// 	// can happen when a large value is converted to a floating point JavaScript.
// 	// Specifying JS_NUMBER for the jstype causes the generated JavaScript code to
// 	// use the JavaScript "number" type.  The behavior of the default option
// 	// JS_NORMAL is implementation dependent.
// 	//
// 	// This option is an enum to permit additional types to be added, e.g.
// 	// goog.math.Integer.
// 	Jstype *FieldOptions_JSType `protobuf:"varint,6,opt,name=jstype,enum=google.protobuf.FieldOptions_JSType,def=0" json:"jstype,omitempty"`
// 	// Should this field be parsed lazily?  Lazy applies only to message-type
// 	// fields.  It means that when the outer message is initially parsed, the
// 	// inner message's contents will not be parsed but instead stored in encoded
// 	// form.  The inner message will actually be parsed when it is first accessed.
// 	//
// 	// This is only a hint.  Implementations are free to choose whether to use
// 	// eager or lazy parsing regardless of the value of this option.  However,
// 	// setting this option true suggests that the protocol author believes that
// 	// using lazy parsing on this field is worth the additional bookkeeping
// 	// overhead typically needed to implement it.
// 	//
// 	// This option does not affect the public interface of any generated code;
// 	// all method signatures remain the same.  Furthermore, thread-safety of the
// 	// interface is not affected by this option; const methods remain safe to
// 	// call from multiple threads concurrently, while non-const methods continue
// 	// to require exclusive access.
// 	//
// 	//
// 	// Note that implementations may choose not to check required fields within
// 	// a lazy sub-message.  That is, calling IsInitialized() on the outer message
// 	// may return true even if the inner message has missing required fields.
// 	// This is necessary because otherwise the inner message would have to be
// 	// parsed in order to perform the check, defeating the purpose of lazy
// 	// parsing.  An implementation which chooses not to check required fields
// 	// must be consistent about it.  That is, for any particular sub-message, the
// 	// implementation must either *always* check its required fields, or *never*
// 	// check its required fields, regardless of whether or not the message has
// 	// been parsed.
// 	Lazy *bool `protobuf:"varint,5,opt,name=lazy,def=0" json:"lazy,omitempty"`
// 	// Is this field deprecated?
// 	// Depending on the target platform, this can emit Deprecated annotations
// 	// for accessors, or it will be completely ignored; in the very least, this
// 	// is a formalization for deprecating fields.
// 	Deprecated *bool `protobuf:"varint,3,opt,name=deprecated,def=0" json:"deprecated,omitempty"`
// 	// For Google-internal migration only. Do not use.
// 	Weak *bool `protobuf:"varint,10,opt,name=weak,def=0" json:"weak,omitempty"`
// 	// The parser stores options it doesn't recognize here. See above.
// 	UninterpretedOption []*UninterpretedOption `protobuf:"bytes,999,rep,name=uninterpreted_option,json=uninterpretedOption" json:"uninterpreted_option,omitempty"`
// }

// // Default values for FieldOptions fields.
// const (
// 	Default_FieldOptions_Ctype      = FieldOptions_STRING
// 	Default_FieldOptions_Jstype     = FieldOptions_JS_NORMAL
// 	Default_FieldOptions_Lazy       = bool(false)
// 	Default_FieldOptions_Deprecated = bool(false)
// 	Default_FieldOptions_Weak       = bool(false)
// )

// func (x *FieldOptions) Reset() {
// 	*x = FieldOptions{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[12]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *FieldOptions) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*FieldOptions) ProtoMessage() {}

// func (x *FieldOptions) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[12]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use FieldOptions.ProtoReflect.Descriptor instead.
// func (*FieldOptions) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{12}
// }

// func (x *FieldOptions) GetCtype() FieldOptions_CType {
// 	if x != nil && x.Ctype != nil {
// 		return *x.Ctype
// 	}
// 	return Default_FieldOptions_Ctype
// }

// func (x *FieldOptions) GetPacked() bool {
// 	if x != nil && x.Packed != nil {
// 		return *x.Packed
// 	}
// 	return false
// }

// func (x *FieldOptions) GetJstype() FieldOptions_JSType {
// 	if x != nil && x.Jstype != nil {
// 		return *x.Jstype
// 	}
// 	return Default_FieldOptions_Jstype
// }

// func (x *FieldOptions) GetLazy() bool {
// 	if x != nil && x.Lazy != nil {
// 		return *x.Lazy
// 	}
// 	return Default_FieldOptions_Lazy
// }

// func (x *FieldOptions) GetDeprecated() bool {
// 	if x != nil && x.Deprecated != nil {
// 		return *x.Deprecated
// 	}
// 	return Default_FieldOptions_Deprecated
// }

// func (x *FieldOptions) GetWeak() bool {
// 	if x != nil && x.Weak != nil {
// 		return *x.Weak
// 	}
// 	return Default_FieldOptions_Weak
// }

// func (x *FieldOptions) GetUninterpretedOption() []*UninterpretedOption {
// 	if x != nil {
// 		return x.UninterpretedOption
// 	}
// 	return nil
// }

// type OneofOptions struct {
// 	state           protoimpl.MessageState
// 	sizeCache       protoimpl.SizeCache
// 	unknownFields   protoimpl.UnknownFields
// 	extensionFields protoimpl.ExtensionFields

// 	// The parser stores options it doesn't recognize here. See above.
// 	UninterpretedOption []*UninterpretedOption `protobuf:"bytes,999,rep,name=uninterpreted_option,json=uninterpretedOption" json:"uninterpreted_option,omitempty"`
// }

// func (x *OneofOptions) Reset() {
// 	*x = OneofOptions{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[13]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *OneofOptions) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*OneofOptions) ProtoMessage() {}

// func (x *OneofOptions) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[13]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use OneofOptions.ProtoReflect.Descriptor instead.
// func (*OneofOptions) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{13}
// }

// func (x *OneofOptions) GetUninterpretedOption() []*UninterpretedOption {
// 	if x != nil {
// 		return x.UninterpretedOption
// 	}
// 	return nil
// }

// type EnumOptions struct {
// 	state           protoimpl.MessageState
// 	sizeCache       protoimpl.SizeCache
// 	unknownFields   protoimpl.UnknownFields
// 	extensionFields protoimpl.ExtensionFields

// 	// Set this option to true to allow mapping different tag names to the same
// 	// value.
// 	AllowAlias *bool `protobuf:"varint,2,opt,name=allow_alias,json=allowAlias" json:"allow_alias,omitempty"`
// 	// Is this enum deprecated?
// 	// Depending on the target platform, this can emit Deprecated annotations
// 	// for the enum, or it will be completely ignored; in the very least, this
// 	// is a formalization for deprecating enums.
// 	Deprecated *bool `protobuf:"varint,3,opt,name=deprecated,def=0" json:"deprecated,omitempty"`
// 	// The parser stores options it doesn't recognize here. See above.
// 	UninterpretedOption []*UninterpretedOption `protobuf:"bytes,999,rep,name=uninterpreted_option,json=uninterpretedOption" json:"uninterpreted_option,omitempty"`
// }

// // Default values for EnumOptions fields.
// const (
// 	Default_EnumOptions_Deprecated = bool(false)
// )

// func (x *EnumOptions) Reset() {
// 	*x = EnumOptions{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[14]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *EnumOptions) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*EnumOptions) ProtoMessage() {}

// func (x *EnumOptions) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[14]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use EnumOptions.ProtoReflect.Descriptor instead.
// func (*EnumOptions) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{14}
// }

// func (x *EnumOptions) GetAllowAlias() bool {
// 	if x != nil && x.AllowAlias != nil {
// 		return *x.AllowAlias
// 	}
// 	return false
// }

// func (x *EnumOptions) GetDeprecated() bool {
// 	if x != nil && x.Deprecated != nil {
// 		return *x.Deprecated
// 	}
// 	return Default_EnumOptions_Deprecated
// }

// func (x *EnumOptions) GetUninterpretedOption() []*UninterpretedOption {
// 	if x != nil {
// 		return x.UninterpretedOption
// 	}
// 	return nil
// }

// type EnumValueOptions struct {
// 	state           protoimpl.MessageState
// 	sizeCache       protoimpl.SizeCache
// 	unknownFields   protoimpl.UnknownFields
// 	extensionFields protoimpl.ExtensionFields

// 	// Is this enum value deprecated?
// 	// Depending on the target platform, this can emit Deprecated annotations
// 	// for the enum value, or it will be completely ignored; in the very least,
// 	// this is a formalization for deprecating enum values.
// 	Deprecated *bool `protobuf:"varint,1,opt,name=deprecated,def=0" json:"deprecated,omitempty"`
// 	// The parser stores options it doesn't recognize here. See above.
// 	UninterpretedOption []*UninterpretedOption `protobuf:"bytes,999,rep,name=uninterpreted_option,json=uninterpretedOption" json:"uninterpreted_option,omitempty"`
// }

// // Default values for EnumValueOptions fields.
// const (
// 	Default_EnumValueOptions_Deprecated = bool(false)
// )

// func (x *EnumValueOptions) Reset() {
// 	*x = EnumValueOptions{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[15]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *EnumValueOptions) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*EnumValueOptions) ProtoMessage() {}

// func (x *EnumValueOptions) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[15]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use EnumValueOptions.ProtoReflect.Descriptor instead.
// func (*EnumValueOptions) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{15}
// }

// func (x *EnumValueOptions) GetDeprecated() bool {
// 	if x != nil && x.Deprecated != nil {
// 		return *x.Deprecated
// 	}
// 	return Default_EnumValueOptions_Deprecated
// }

// func (x *EnumValueOptions) GetUninterpretedOption() []*UninterpretedOption {
// 	if x != nil {
// 		return x.UninterpretedOption
// 	}
// 	return nil
// }

// type ServiceOptions struct {
// 	state           protoimpl.MessageState
// 	sizeCache       protoimpl.SizeCache
// 	unknownFields   protoimpl.UnknownFields
// 	extensionFields protoimpl.ExtensionFields

// 	// Is this service deprecated?
// 	// Depending on the target platform, this can emit Deprecated annotations
// 	// for the service, or it will be completely ignored; in the very least,
// 	// this is a formalization for deprecating services.
// 	Deprecated *bool `protobuf:"varint,33,opt,name=deprecated,def=0" json:"deprecated,omitempty"`
// 	// The parser stores options it doesn't recognize here. See above.
// 	UninterpretedOption []*UninterpretedOption `protobuf:"bytes,999,rep,name=uninterpreted_option,json=uninterpretedOption" json:"uninterpreted_option,omitempty"`
// }

// // Default values for ServiceOptions fields.
// const (
// 	Default_ServiceOptions_Deprecated = bool(false)
// )

// func (x *ServiceOptions) Reset() {
// 	*x = ServiceOptions{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[16]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *ServiceOptions) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*ServiceOptions) ProtoMessage() {}

// func (x *ServiceOptions) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[16]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use ServiceOptions.ProtoReflect.Descriptor instead.
// func (*ServiceOptions) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{16}
// }

// func (x *ServiceOptions) GetDeprecated() bool {
// 	if x != nil && x.Deprecated != nil {
// 		return *x.Deprecated
// 	}
// 	return Default_ServiceOptions_Deprecated
// }

// func (x *ServiceOptions) GetUninterpretedOption() []*UninterpretedOption {
// 	if x != nil {
// 		return x.UninterpretedOption
// 	}
// 	return nil
// }

// type MethodOptions struct {
// 	state           protoimpl.MessageState
// 	sizeCache       protoimpl.SizeCache
// 	unknownFields   protoimpl.UnknownFields
// 	extensionFields protoimpl.ExtensionFields

// 	// Is this method deprecated?
// 	// Depending on the target platform, this can emit Deprecated annotations
// 	// for the method, or it will be completely ignored; in the very least,
// 	// this is a formalization for deprecating methods.
// 	Deprecated       *bool                           `protobuf:"varint,33,opt,name=deprecated,def=0" json:"deprecated,omitempty"`
// 	IdempotencyLevel *MethodOptions_IdempotencyLevel `protobuf:"varint,34,opt,name=idempotency_level,json=idempotencyLevel,enum=google.protobuf.MethodOptions_IdempotencyLevel,def=0" json:"idempotency_level,omitempty"`
// 	// The parser stores options it doesn't recognize here. See above.
// 	UninterpretedOption []*UninterpretedOption `protobuf:"bytes,999,rep,name=uninterpreted_option,json=uninterpretedOption" json:"uninterpreted_option,omitempty"`
// }

// // Default values for MethodOptions fields.
// const (
// 	Default_MethodOptions_Deprecated       = bool(false)
// 	Default_MethodOptions_IdempotencyLevel = MethodOptions_IDEMPOTENCY_UNKNOWN
// )

// func (x *MethodOptions) Reset() {
// 	*x = MethodOptions{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[17]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *MethodOptions) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*MethodOptions) ProtoMessage() {}

// func (x *MethodOptions) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[17]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use MethodOptions.ProtoReflect.Descriptor instead.
// func (*MethodOptions) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{17}
// }

// func (x *MethodOptions) GetDeprecated() bool {
// 	if x != nil && x.Deprecated != nil {
// 		return *x.Deprecated
// 	}
// 	return Default_MethodOptions_Deprecated
// }

// func (x *MethodOptions) GetIdempotencyLevel() MethodOptions_IdempotencyLevel {
// 	if x != nil && x.IdempotencyLevel != nil {
// 		return *x.IdempotencyLevel
// 	}
// 	return Default_MethodOptions_IdempotencyLevel
// }

// func (x *MethodOptions) GetUninterpretedOption() []*UninterpretedOption {
// 	if x != nil {
// 		return x.UninterpretedOption
// 	}
// 	return nil
// }

// // A message representing a option the parser does not recognize. This only
// // appears in options protos created by the compiler::Parser class.
// // DescriptorPool resolves these when building Descriptor objects. Therefore,
// // options protos in descriptor objects (e.g. returned by Descriptor::options(),
// // or produced by Descriptor::CopyTo()) will never have UninterpretedOptions
// // in them.
// type UninterpretedOption struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	Name []*UninterpretedOption_NamePart `protobuf:"bytes,2,rep,name=name" json:"name,omitempty"`
// 	// The value of the uninterpreted option, in whatever type the tokenizer
// 	// identified it as during parsing. Exactly one of these should be set.
// 	IdentifierValue  *string  `protobuf:"bytes,3,opt,name=identifier_value,json=identifierValue" json:"identifier_value,omitempty"`
// 	PositiveIntValue *uint64  `protobuf:"varint,4,opt,name=positive_int_value,json=positiveIntValue" json:"positive_int_value,omitempty"`
// 	NegativeIntValue *int64   `protobuf:"varint,5,opt,name=negative_int_value,json=negativeIntValue" json:"negative_int_value,omitempty"`
// 	DoubleValue      *float64 `protobuf:"fixed64,6,opt,name=double_value,json=doubleValue" json:"double_value,omitempty"`
// 	StringValue      []byte   `protobuf:"bytes,7,opt,name=string_value,json=stringValue" json:"string_value,omitempty"`
// 	AggregateValue   *string  `protobuf:"bytes,8,opt,name=aggregate_value,json=aggregateValue" json:"aggregate_value,omitempty"`
// }

// func (x *UninterpretedOption) Reset() {
// 	*x = UninterpretedOption{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[18]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *UninterpretedOption) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*UninterpretedOption) ProtoMessage() {}

// func (x *UninterpretedOption) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[18]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use UninterpretedOption.ProtoReflect.Descriptor instead.
// func (*UninterpretedOption) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{18}
// }

// func (x *UninterpretedOption) GetName() []*UninterpretedOption_NamePart {
// 	if x != nil {
// 		return x.Name
// 	}
// 	return nil
// }

// func (x *UninterpretedOption) GetIdentifierValue() string {
// 	if x != nil && x.IdentifierValue != nil {
// 		return *x.IdentifierValue
// 	}
// 	return ""
// }

// func (x *UninterpretedOption) GetPositiveIntValue() uint64 {
// 	if x != nil && x.PositiveIntValue != nil {
// 		return *x.PositiveIntValue
// 	}
// 	return 0
// }

// func (x *UninterpretedOption) GetNegativeIntValue() int64 {
// 	if x != nil && x.NegativeIntValue != nil {
// 		return *x.NegativeIntValue
// 	}
// 	return 0
// }

// func (x *UninterpretedOption) GetDoubleValue() float64 {
// 	if x != nil && x.DoubleValue != nil {
// 		return *x.DoubleValue
// 	}
// 	return 0
// }

// func (x *UninterpretedOption) GetStringValue() []byte {
// 	if x != nil {
// 		return x.StringValue
// 	}
// 	return nil
// }

// func (x *UninterpretedOption) GetAggregateValue() string {
// 	if x != nil && x.AggregateValue != nil {
// 		return *x.AggregateValue
// 	}
// 	return ""
// }

// // Encapsulates information about the original source file from which a
// // FileDescriptorProto was generated.
// type SourceCodeInfo struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	// A Location identifies a piece of source code in a .proto file which
// 	// corresponds to a particular definition.  This information is intended
// 	// to be useful to IDEs, code indexers, documentation generators, and similar
// 	// tools.
// 	//
// 	// For example, say we have a file like:
// 	//   message Foo {
// 	//     optional string foo = 1;
// 	//   }
// 	// Let's look at just the field definition:
// 	//   optional string foo = 1;
// 	//   ^       ^^     ^^  ^  ^^^
// 	//   a       bc     de  f  ghi
// 	// We have the following locations:
// 	//   span   path               represents
// 	//   [a,i)  [ 4, 0, 2, 0 ]     The whole field definition.
// 	//   [a,b)  [ 4, 0, 2, 0, 4 ]  The label (optional).
// 	//   [c,d)  [ 4, 0, 2, 0, 5 ]  The type (string).
// 	//   [e,f)  [ 4, 0, 2, 0, 1 ]  The name (foo).
// 	//   [g,h)  [ 4, 0, 2, 0, 3 ]  The number (1).
// 	//
// 	// Notes:
// 	// - A location may refer to a repeated field itself (i.e. not to any
// 	//   particular index within it).  This is used whenever a set of elements are
// 	//   logically enclosed in a single code segment.  For example, an entire
// 	//   extend block (possibly containing multiple extension definitions) will
// 	//   have an outer location whose path refers to the "extensions" repeated
// 	//   field without an index.
// 	// - Multiple locations may have the same path.  This happens when a single
// 	//   logical declaration is spread out across multiple places.  The most
// 	//   obvious example is the "extend" block again -- there may be multiple
// 	//   extend blocks in the same scope, each of which will have the same path.
// 	// - A location's span is not always a subset of its parent's span.  For
// 	//   example, the "extendee" of an extension declaration appears at the
// 	//   beginning of the "extend" block and is shared by all extensions within
// 	//   the block.
// 	// - Just because a location's span is a subset of some other location's span
// 	//   does not mean that it is a descendent.  For example, a "group" defines
// 	//   both a type and a field in a single declaration.  Thus, the locations
// 	//   corresponding to the type and field and their components will overlap.
// 	// - Code which tries to interpret locations should probably be designed to
// 	//   ignore those that it doesn't understand, as more types of locations could
// 	//   be recorded in the future.
// 	Location []*SourceCodeInfo_Location `protobuf:"bytes,1,rep,name=location" json:"location,omitempty"`
// }

// func (x *SourceCodeInfo) Reset() {
// 	*x = SourceCodeInfo{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[19]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *SourceCodeInfo) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*SourceCodeInfo) ProtoMessage() {}

// func (x *SourceCodeInfo) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[19]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use SourceCodeInfo.ProtoReflect.Descriptor instead.
// func (*SourceCodeInfo) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{19}
// }

// func (x *SourceCodeInfo) GetLocation() []*SourceCodeInfo_Location {
// 	if x != nil {
// 		return x.Location
// 	}
// 	return nil
// }

// // Describes the relationship between generated code and its original source
// // file. A GeneratedCodeInfo message is associated with only one generated
// // source file, but may contain references to different source .proto files.
// type GeneratedCodeInfo struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	// An Annotation connects some span of text in generated code to an element
// 	// of its generating .proto file.
// 	Annotation []*GeneratedCodeInfo_Annotation `protobuf:"bytes,1,rep,name=annotation" json:"annotation,omitempty"`
// }

// func (x *GeneratedCodeInfo) Reset() {
// 	*x = GeneratedCodeInfo{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[20]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *GeneratedCodeInfo) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*GeneratedCodeInfo) ProtoMessage() {}

// func (x *GeneratedCodeInfo) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[20]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use GeneratedCodeInfo.ProtoReflect.Descriptor instead.
// func (*GeneratedCodeInfo) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{20}
// }

// func (x *GeneratedCodeInfo) GetAnnotation() []*GeneratedCodeInfo_Annotation {
// 	if x != nil {
// 		return x.Annotation
// 	}
// 	return nil
// }

// type DescriptorProto_ExtensionRange struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	Start   *int32                 `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
// 	End     *int32                 `protobuf:"varint,2,opt,name=end" json:"end,omitempty"`
// 	Options *ExtensionRangeOptions `protobuf:"bytes,3,opt,name=options" json:"options,omitempty"`
// }

// func (x *DescriptorProto_ExtensionRange) Reset() {
// 	*x = DescriptorProto_ExtensionRange{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[21]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *DescriptorProto_ExtensionRange) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*DescriptorProto_ExtensionRange) ProtoMessage() {}

// func (x *DescriptorProto_ExtensionRange) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[21]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use DescriptorProto_ExtensionRange.ProtoReflect.Descriptor instead.
// func (*DescriptorProto_ExtensionRange) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{2, 0}
// }

// func (x *DescriptorProto_ExtensionRange) GetStart() int32 {
// 	if x != nil && x.Start != nil {
// 		return *x.Start
// 	}
// 	return 0
// }

// func (x *DescriptorProto_ExtensionRange) GetEnd() int32 {
// 	if x != nil && x.End != nil {
// 		return *x.End
// 	}
// 	return 0
// }

// func (x *DescriptorProto_ExtensionRange) GetOptions() *ExtensionRangeOptions {
// 	if x != nil {
// 		return x.Options
// 	}
// 	return nil
// }

// // Range of reserved tag numbers. Reserved tag numbers may not be used by
// // fields or extension ranges in the same message. Reserved ranges may
// // not overlap.
// type DescriptorProto_ReservedRange struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	Start *int32 `protobuf:"varint,1,opt,name=start" json:"start,omitempty"` // Inclusive.
// 	End   *int32 `protobuf:"varint,2,opt,name=end" json:"end,omitempty"`     // Exclusive.
// }

// func (x *DescriptorProto_ReservedRange) Reset() {
// 	*x = DescriptorProto_ReservedRange{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[22]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *DescriptorProto_ReservedRange) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*DescriptorProto_ReservedRange) ProtoMessage() {}

// func (x *DescriptorProto_ReservedRange) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[22]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use DescriptorProto_ReservedRange.ProtoReflect.Descriptor instead.
// func (*DescriptorProto_ReservedRange) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{2, 1}
// }

// func (x *DescriptorProto_ReservedRange) GetStart() int32 {
// 	if x != nil && x.Start != nil {
// 		return *x.Start
// 	}
// 	return 0
// }

// func (x *DescriptorProto_ReservedRange) GetEnd() int32 {
// 	if x != nil && x.End != nil {
// 		return *x.End
// 	}
// 	return 0
// }

// // Range of reserved numeric values. Reserved values may not be used by
// // entries in the same enum. Reserved ranges may not overlap.
// //
// // Note that this is distinct from DescriptorProto.ReservedRange in that it
// // is inclusive such that it can appropriately represent the entire int32
// // domain.
// type EnumDescriptorProto_EnumReservedRange struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	Start *int32 `protobuf:"varint,1,opt,name=start" json:"start,omitempty"` // Inclusive.
// 	End   *int32 `protobuf:"varint,2,opt,name=end" json:"end,omitempty"`     // Inclusive.
// }

// func (x *EnumDescriptorProto_EnumReservedRange) Reset() {
// 	*x = EnumDescriptorProto_EnumReservedRange{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[23]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *EnumDescriptorProto_EnumReservedRange) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*EnumDescriptorProto_EnumReservedRange) ProtoMessage() {}

// func (x *EnumDescriptorProto_EnumReservedRange) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[23]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use EnumDescriptorProto_EnumReservedRange.ProtoReflect.Descriptor instead.
// func (*EnumDescriptorProto_EnumReservedRange) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{6, 0}
// }

// func (x *EnumDescriptorProto_EnumReservedRange) GetStart() int32 {
// 	if x != nil && x.Start != nil {
// 		return *x.Start
// 	}
// 	return 0
// }

// func (x *EnumDescriptorProto_EnumReservedRange) GetEnd() int32 {
// 	if x != nil && x.End != nil {
// 		return *x.End
// 	}
// 	return 0
// }

// // The name of the uninterpreted option.  Each string represents a segment in
// // a dot-separated name.  is_extension is true iff a segment represents an
// // extension (denoted with parentheses in options specs in .proto files).
// // E.g.,{ ["foo", false], ["bar.baz", true], ["qux", false] } represents
// // "foo.(bar.baz).qux".
// type UninterpretedOption_NamePart struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	NamePart    *string `protobuf:"bytes,1,req,name=name_part,json=namePart" json:"name_part,omitempty"`
// 	IsExtension *bool   `protobuf:"varint,2,req,name=is_extension,json=isExtension" json:"is_extension,omitempty"`
// }

// func (x *UninterpretedOption_NamePart) Reset() {
// 	*x = UninterpretedOption_NamePart{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[24]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *UninterpretedOption_NamePart) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*UninterpretedOption_NamePart) ProtoMessage() {}

// func (x *UninterpretedOption_NamePart) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[24]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use UninterpretedOption_NamePart.ProtoReflect.Descriptor instead.
// func (*UninterpretedOption_NamePart) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{18, 0}
// }

// func (x *UninterpretedOption_NamePart) GetNamePart() string {
// 	if x != nil && x.NamePart != nil {
// 		return *x.NamePart
// 	}
// 	return ""
// }

// func (x *UninterpretedOption_NamePart) GetIsExtension() bool {
// 	if x != nil && x.IsExtension != nil {
// 		return *x.IsExtension
// 	}
// 	return false
// }

// type SourceCodeInfo_Location struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	// Identifies which part of the FileDescriptorProto was defined at this
// 	// location.
// 	//
// 	// Each element is a field number or an index.  They form a path from
// 	// the root FileDescriptorProto to the place where the definition.  For
// 	// example, this path:
// 	//   [ 4, 3, 2, 7, 1 ]
// 	// refers to:
// 	//   file.message_type(3)  // 4, 3
// 	//       .field(7)         // 2, 7
// 	//       .name()           // 1
// 	// This is because FileDescriptorProto.message_type has field number 4:
// 	//   repeated DescriptorProto message_type = 4;
// 	// and DescriptorProto.field has field number 2:
// 	//   repeated FieldDescriptorProto field = 2;
// 	// and FieldDescriptorProto.name has field number 1:
// 	//   optional string name = 1;
// 	//
// 	// Thus, the above path gives the location of a field name.  If we removed
// 	// the last element:
// 	//   [ 4, 3, 2, 7 ]
// 	// this path refers to the whole field declaration (from the beginning
// 	// of the label to the terminating semicolon).
// 	Path []int32 `protobuf:"varint,1,rep,packed,name=path" json:"path,omitempty"`
// 	// Always has exactly three or four elements: start line, start column,
// 	// end line (optional, otherwise assumed same as start line), end column.
// 	// These are packed into a single field for efficiency.  Note that line
// 	// and column numbers are zero-based -- typically you will want to add
// 	// 1 to each before displaying to a user.
// 	Span []int32 `protobuf:"varint,2,rep,packed,name=span" json:"span,omitempty"`
// 	// If this SourceCodeInfo represents a complete declaration, these are any
// 	// comments appearing before and after the declaration which appear to be
// 	// attached to the declaration.
// 	//
// 	// A series of line comments appearing on consecutive lines, with no other
// 	// tokens appearing on those lines, will be treated as a single comment.
// 	//
// 	// leading_detached_comments will keep paragraphs of comments that appear
// 	// before (but not connected to) the current element. Each paragraph,
// 	// separated by empty lines, will be one comment element in the repeated
// 	// field.
// 	//
// 	// Only the comment content is provided; comment markers (e.g. //) are
// 	// stripped out.  For block comments, leading whitespace and an asterisk
// 	// will be stripped from the beginning of each line other than the first.
// 	// Newlines are included in the output.
// 	//
// 	// Examples:
// 	//
// 	//   optional int32 foo = 1;  // Comment attached to foo.
// 	//   // Comment attached to bar.
// 	//   optional int32 bar = 2;
// 	//
// 	//   optional string baz = 3;
// 	//   // Comment attached to baz.
// 	//   // Another line attached to baz.
// 	//
// 	//   // Comment attached to qux.
// 	//   //
// 	//   // Another line attached to qux.
// 	//   optional double qux = 4;
// 	//
// 	//   // Detached comment for corge. This is not leading or trailing comments
// 	//   // to qux or corge because there are blank lines separating it from
// 	//   // both.
// 	//
// 	//   // Detached comment for corge paragraph 2.
// 	//
// 	//   optional string corge = 5;
// 	//   /* Block comment attached
// 	//    * to corge.  Leading asterisks
// 	//    * will be removed. */
// 	//   /* Block comment attached to
// 	//    * grault. */
// 	//   optional int32 grault = 6;
// 	//
// 	//   // ignored detached comments.
// 	LeadingComments         *string  `protobuf:"bytes,3,opt,name=leading_comments,json=leadingComments" json:"leading_comments,omitempty"`
// 	TrailingComments        *string  `protobuf:"bytes,4,opt,name=trailing_comments,json=trailingComments" json:"trailing_comments,omitempty"`
// 	LeadingDetachedComments []string `protobuf:"bytes,6,rep,name=leading_detached_comments,json=leadingDetachedComments" json:"leading_detached_comments,omitempty"`
// }

// func (x *SourceCodeInfo_Location) Reset() {
// 	*x = SourceCodeInfo_Location{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[25]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *SourceCodeInfo_Location) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*SourceCodeInfo_Location) ProtoMessage() {}

// func (x *SourceCodeInfo_Location) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[25]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use SourceCodeInfo_Location.ProtoReflect.Descriptor instead.
// func (*SourceCodeInfo_Location) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{19, 0}
// }

// func (x *SourceCodeInfo_Location) GetPath() []int32 {
// 	if x != nil {
// 		return x.Path
// 	}
// 	return nil
// }

// func (x *SourceCodeInfo_Location) GetSpan() []int32 {
// 	if x != nil {
// 		return x.Span
// 	}
// 	return nil
// }

// func (x *SourceCodeInfo_Location) GetLeadingComments() string {
// 	if x != nil && x.LeadingComments != nil {
// 		return *x.LeadingComments
// 	}
// 	return ""
// }

// func (x *SourceCodeInfo_Location) GetTrailingComments() string {
// 	if x != nil && x.TrailingComments != nil {
// 		return *x.TrailingComments
// 	}
// 	return ""
// }

// func (x *SourceCodeInfo_Location) GetLeadingDetachedComments() []string {
// 	if x != nil {
// 		return x.LeadingDetachedComments
// 	}
// 	return nil
// }

// type GeneratedCodeInfo_Annotation struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	// Identifies the element in the original source .proto file. This field
// 	// is formatted the same as SourceCodeInfo.Location.path.
// 	Path []int32 `protobuf:"varint,1,rep,packed,name=path" json:"path,omitempty"`
// 	// Identifies the filesystem path to the original source .proto.
// 	SourceFile *string `protobuf:"bytes,2,opt,name=source_file,json=sourceFile" json:"source_file,omitempty"`
// 	// Identifies the starting offset in bytes in the generated code
// 	// that relates to the identified object.
// 	Begin *int32 `protobuf:"varint,3,opt,name=begin" json:"begin,omitempty"`
// 	// Identifies the ending offset in bytes in the generated code that
// 	// relates to the identified offset. The end offset should be one past
// 	// the last relevant byte (so the length of the text = end - begin).
// 	End *int32 `protobuf:"varint,4,opt,name=end" json:"end,omitempty"`
// }

// func (x *GeneratedCodeInfo_Annotation) Reset() {
// 	*x = GeneratedCodeInfo_Annotation{}
// 	if protoimpl.UnsafeEnabled {
// 		mi := &file_descriptor_proto_msgTypes[26]
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		ms.StoreMessageInfo(mi)
// 	}
// }

// func (x *GeneratedCodeInfo_Annotation) String() string {
// 	return protoimpl.X.MessageStringOf(x)
// }

// func (*GeneratedCodeInfo_Annotation) ProtoMessage() {}

// func (x *GeneratedCodeInfo_Annotation) ProtoReflect() protoreflect.Message {
// 	mi := &file_descriptor_proto_msgTypes[26]
// 	if protoimpl.UnsafeEnabled && x != nil {
// 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
// 		if ms.LoadMessageInfo() == nil {
// 			ms.StoreMessageInfo(mi)
// 		}
// 		return ms
// 	}
// 	return mi.MessageOf(x)
// }

// // Deprecated: Use GeneratedCodeInfo_Annotation.ProtoReflect.Descriptor instead.
// func (*GeneratedCodeInfo_Annotation) Descriptor() ([]byte, []int) {
// 	return file_descriptor_proto_rawDescGZIP(), []int{20, 0}
// }

// func (x *GeneratedCodeInfo_Annotation) GetPath() []int32 {
// 	if x != nil {
// 		return x.Path
// 	}
// 	return nil
// }

// func (x *GeneratedCodeInfo_Annotation) GetSourceFile() string {
// 	if x != nil && x.SourceFile != nil {
// 		return *x.SourceFile
// 	}
// 	return ""
// }

// func (x *GeneratedCodeInfo_Annotation) GetBegin() int32 {
// 	if x != nil && x.Begin != nil {
// 		return *x.Begin
// 	}
// 	return 0
// }

// func (x *GeneratedCodeInfo_Annotation) GetEnd() int32 {
// 	if x != nil && x.End != nil {
// 		return *x.End
// 	}
// 	return 0
// }

// var File_descriptor_proto protoreflect.FileDescriptor

// var file_descriptor_proto_rawDesc = []byte{
// 	0x0a, 0x10, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
// 	0x74, 0x6f, 0x12, 0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
// 	0x62, 0x75, 0x66, 0x22, 0x4d, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72,
// 	0x69, 0x70, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x12, 0x38, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65,
// 	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
// 	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73,
// 	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x04, 0x66, 0x69,
// 	0x6c, 0x65, 0x22, 0xe4, 0x04, 0x0a, 0x13, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72,
// 	0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
// 	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
// 	0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
// 	0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x65,
// 	0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65,
// 	0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x75, 0x62, 0x6c,
// 	0x69, 0x63, 0x5f, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0a, 0x20,
// 	0x03, 0x28, 0x05, 0x52, 0x10, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x44, 0x65, 0x70, 0x65, 0x6e,
// 	0x64, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x77, 0x65, 0x61, 0x6b, 0x5f, 0x64, 0x65,
// 	0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0e,
// 	0x77, 0x65, 0x61, 0x6b, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x43,
// 	0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
// 	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
// 	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
// 	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54,
// 	0x79, 0x70, 0x65, 0x12, 0x41, 0x0a, 0x09, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65,
// 	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
// 	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x44, 0x65, 0x73,
// 	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x08, 0x65, 0x6e,
// 	0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x41, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
// 	0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
// 	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
// 	0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
// 	0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x09, 0x65, 0x78, 0x74,
// 	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67,
// 	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
// 	0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72,
// 	0x6f, 0x74, 0x6f, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x36,
// 	0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
// 	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
// 	0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f,
// 	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x49, 0x0a, 0x10, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
// 	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
// 	0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
// 	0x75, 0x66, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66,
// 	0x6f, 0x52, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66,
// 	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x18, 0x0c, 0x20, 0x01, 0x28,
// 	0x09, 0x52, 0x06, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x22, 0xb9, 0x06, 0x0a, 0x0f, 0x44, 0x65,
// 	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x0a,
// 	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
// 	0x65, 0x12, 0x3b, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
// 	0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
// 	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
// 	0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x43,
// 	0x0a, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x03, 0x28,
// 	0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
// 	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
// 	0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
// 	0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x0b, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x79,
// 	0x70, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
// 	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
// 	0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x0a, 0x6e, 0x65, 0x73, 0x74,
// 	0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x41, 0x0a, 0x09, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x74,
// 	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
// 	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
// 	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52,
// 	0x08, 0x65, 0x6e, 0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x58, 0x0a, 0x0f, 0x65, 0x78, 0x74,
// 	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x05, 0x20, 0x03,
// 	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
// 	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50,
// 	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61,
// 	0x6e, 0x67, 0x65, 0x52, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61,
// 	0x6e, 0x67, 0x65, 0x12, 0x44, 0x0a, 0x0a, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x64, 0x65, 0x63,
// 	0x6c, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
// 	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x44,
// 	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x09,
// 	0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x44, 0x65, 0x63, 0x6c, 0x12, 0x39, 0x0a, 0x07, 0x6f, 0x70, 0x74,
// 	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x6f,
// 	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73,
// 	0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74,
// 	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x55, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64,
// 	0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67,
// 	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
// 	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
// 	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0d, 0x72, 0x65,
// 	0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72,
// 	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x03,
// 	0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65,
// 	0x1a, 0x7a, 0x0a, 0x0e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x6e,
// 	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
// 	0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18,
// 	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x40, 0x0a, 0x07, 0x6f, 0x70,
// 	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f,
// 	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x78,
// 	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
// 	0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x37, 0x0a, 0x0d,
// 	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a,
// 	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74,
// 	0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
// 	0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x7c, 0x0a, 0x15, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
// 	0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x58,
// 	0x0a, 0x14, 0x75, 0x6e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x64, 0x5f,
// 	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0xe7, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
// 	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
// 	0x55, 0x6e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x74,
// 	0x69, 0x6f, 0x6e, 0x52, 0x13, 0x75, 0x6e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74,
// 	0x65, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x09, 0x08, 0xe8, 0x07, 0x10, 0x80, 0x80,
// 	0x80, 0x80, 0x02, 0x22, 0x98, 0x06, 0x0a, 0x14, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x73,
// 	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04,
// 	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
// 	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
// 	0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65,
// 	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
// 	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44,
// 	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c,
// 	0x61, 0x62, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x3e, 0x0a, 0x04, 0x74,
// 	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
// 	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
// 	0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
// 	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74,
// 	0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
// 	0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x74, 0x65,
// 	0x6e, 0x64, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x74, 0x65,
// 	0x6e, 0x64, 0x65, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f,
// 	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x66,
// 	0x61, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x6e, 0x65,
// 	0x6f, 0x66, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
// 	0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x73,
// 	0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a,
// 	0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
// 	0x6e, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
// 	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
// 	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
// 	0x22, 0xb6, 0x02, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50,
// 	0x45, 0x5f, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59,
// 	0x50, 0x45, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59,
// 	0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59,
// 	0x50, 0x45, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x54,
// 	0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x54,
// 	0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x58, 0x45, 0x44, 0x36, 0x34, 0x10, 0x06, 0x12, 0x10, 0x0a,
// 	0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x58, 0x45, 0x44, 0x33, 0x32, 0x10, 0x07, 0x12,
// 	0x0d, 0x0a, 0x09, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x08, 0x12, 0x0f,
// 	0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x09, 0x12,
// 	0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x0a, 0x12,
// 	0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10,
// 	0x0b, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x59, 0x54, 0x45, 0x53, 0x10,
// 	0x0c, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x33, 0x32,
// 	0x10, 0x0d, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x10,
// 	0x0e, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x46, 0x49, 0x58, 0x45, 0x44,
// 	0x33, 0x32, 0x10, 0x0f, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x46, 0x49,
// 	0x58, 0x45, 0x44, 0x36, 0x34, 0x10, 0x10, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f,
// 	0x53, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0x11, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45,
// 	0x5f, 0x53, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x12, 0x22, 0x43, 0x0a, 0x05, 0x4c, 0x61, 0x62,
// 	0x65, 0x6c, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x5f, 0x4f, 0x50, 0x54, 0x49,
// 	0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x5f,
// 	0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x41,
// 	0x42, 0x45, 0x4c, 0x5f, 0x52, 0x45, 0x50, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x03, 0x22, 0x63,
// 	0x0a, 0x14, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
// 	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
// 	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x6f, 0x70,
// 	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f,
// 	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x6e,
// 	0x65, 0x6f, 0x66, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69,
// 	0x6f, 0x6e, 0x73, 0x22, 0xe3, 0x02, 0x0a, 0x13, 0x45, 0x6e, 0x75, 0x6d, 0x44, 0x65, 0x73, 0x63,
// 	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e,
// 	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
// 	0x3f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29,
// 	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
// 	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
// 	0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
// 	0x12, 0x36, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
// 	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
// 	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
// 	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5d, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x65,
// 	0x72, 0x76, 0x65, 0x64, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
// 	0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
// 	0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
// 	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x65, 0x72,
// 	0x76, 0x65, 0x64, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
// 	0x65, 0x64, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72,
// 	0x76, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c,
// 	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x3b, 0x0a, 0x11,
// 	0x45, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x52, 0x61, 0x6e, 0x67,
// 	0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
// 	0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02,
// 	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x18, 0x45, 0x6e,
// 	0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
// 	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
// 	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75,
// 	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62,
// 	0x65, 0x72, 0x12, 0x3b, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20,
// 	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
// 	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f,
// 	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
// 	0xa7, 0x01, 0x0a, 0x16, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72,
// 	0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
// 	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e,
// 	0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26,
// 	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
// 	0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
// 	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x39,
// 	0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
// 	0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
// 	0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
// 	0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x89, 0x02, 0x0a, 0x15, 0x4d, 0x65,
// 	0x74, 0x68, 0x6f, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72,
// 	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
// 	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74,
// 	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x70,
// 	0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
// 	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x75, 0x74,
// 	0x70, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
// 	0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
// 	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f,
// 	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
// 	0x73, 0x12, 0x30, 0x0a, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x65,
// 	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c,
// 	0x73, 0x65, 0x52, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
// 	0x69, 0x6e, 0x67, 0x12, 0x30, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x74,
// 	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66,
// 	0x61, 0x6c, 0x73, 0x65, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65,
// 	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x22, 0xb9, 0x08, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70,
// 	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6a, 0x61, 0x76, 0x61, 0x5f, 0x70, 0x61,
// 	0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6a, 0x61, 0x76,
// 	0x61, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x6a, 0x61, 0x76, 0x61,
// 	0x5f, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x6e, 0x61, 0x6d, 0x65,
// 	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6a, 0x61, 0x76, 0x61, 0x4f, 0x75, 0x74, 0x65,
// 	0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x13, 0x6a, 0x61,
// 	0x76, 0x61, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65,
// 	0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x11,
// 	0x6a, 0x61, 0x76, 0x61, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65,
// 	0x73, 0x12, 0x44, 0x0a, 0x1d, 0x6a, 0x61, 0x76, 0x61, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
// 	0x74, 0x65, 0x5f, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x73, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x68, 0x61,
// 	0x73, 0x68, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x42, 0x02, 0x18, 0x01, 0x52, 0x19, 0x6a, 0x61,
// 	0x76, 0x61, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x73,
// 	0x41, 0x6e, 0x64, 0x48, 0x61, 0x73, 0x68, 0x12, 0x3a, 0x0a, 0x16, 0x6a, 0x61, 0x76, 0x61, 0x5f,
// 	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x75, 0x74, 0x66,
// 	0x38, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x13,
// 	0x6a, 0x61, 0x76, 0x61, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55,
// 	0x74, 0x66, 0x38, 0x12, 0x53, 0x0a, 0x0c, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x5f,
// 	0x66, 0x6f, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
// 	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65,
// 	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x65,
// 	0x4d, 0x6f, 0x64, 0x65, 0x3a, 0x05, 0x53, 0x50, 0x45, 0x45, 0x44, 0x52, 0x0b, 0x6f, 0x70, 0x74,
// 	0x69, 0x6d, 0x69, 0x7a, 0x65, 0x46, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x6f, 0x5f, 0x70,
// 	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x6f,
// 	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x35, 0x0a, 0x13, 0x63, 0x63, 0x5f, 0x67, 0x65,
// 	0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x10,
// 	0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x11, 0x63, 0x63, 0x47,
// 	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x39,
// 	0x0a, 0x15, 0x6a, 0x61, 0x76, 0x61, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x73,
// 	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66,
// 	0x61, 0x6c, 0x73, 0x65, 0x52, 0x13, 0x6a, 0x61, 0x76, 0x61, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69,
// 	0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x13, 0x70, 0x79, 0x5f,
// 	0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
// 	0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x11, 0x70,
// 	0x79, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
// 	0x12, 0x37, 0x0a, 0x14, 0x70, 0x68, 0x70, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f,
// 	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05,
// 	0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x12, 0x70, 0x68, 0x70, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69,
// 	0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0a, 0x64, 0x65, 0x70,
// 	0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66,
// 	0x61, 0x6c, 0x73, 0x65, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64,
// 	0x12, 0x2f, 0x0a, 0x10, 0x63, 0x63, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x72,
// 	0x65, 0x6e, 0x61, 0x73, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73,
// 	0x65, 0x52, 0x0e, 0x63, 0x63, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x72, 0x65, 0x6e, 0x61,
// 	0x73, 0x12, 0x2a, 0x0a, 0x11, 0x6f, 0x62, 0x6a, 0x63, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f,
// 	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x24, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x62,
// 	0x6a, 0x63, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x29, 0x0a,
// 	0x10, 0x63, 0x73, 0x68, 0x61, 0x72, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
// 	0x65, 0x18, 0x25, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x73, 0x68, 0x61, 0x72, 0x70, 0x4e,
// 	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x77, 0x69, 0x66,
// 	0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x27, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
// 	0x73, 0x77, 0x69, 0x66, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x28, 0x0a, 0x10, 0x70,
// 	0x68, 0x70, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18,
// 	0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x68, 0x70, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x50,
// 	0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x68, 0x70, 0x5f, 0x6e, 0x61, 0x6d,
// 	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x29, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x68,
// 	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x14, 0x75, 0x6e,
// 	0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69,
// 	0x6f, 0x6e, 0x18, 0xe7, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
// 	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x6e, 0x69, 0x6e,
// 	0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
// 	0x13, 0x75, 0x6e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x64, 0x4f, 0x70,
// 	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x0a, 0x0c, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x65,
// 	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x50, 0x45, 0x45, 0x44, 0x10, 0x01, 0x12,
// 	0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x10, 0x02, 0x12, 0x10,
// 	0x0a, 0x0c, 0x4c, 0x49, 0x54, 0x45, 0x5f, 0x52, 0x55, 0x4e, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x03,
// 	0x2a, 0x09, 0x08, 0xe8, 0x07, 0x10, 0x80, 0x80, 0x80, 0x80, 0x02, 0x4a, 0x04, 0x08, 0x26, 0x10,
// 	0x27, 0x22, 0xd1, 0x02, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74,
// 	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3c, 0x0a, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
// 	0x73, 0x65, 0x74, 0x5f, 0x77, 0x69, 0x72, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18,
// 	0x01, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x14, 0x6d, 0x65,
// 	0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x57, 0x69, 0x72, 0x65, 0x46, 0x6f, 0x72, 0x6d,
// 	0x61, 0x74, 0x12, 0x4c, 0x0a, 0x1f, 0x6e, 0x6f, 0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72,
// 	0x64, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x63, 0x63,
// 	0x65, 0x73, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c,
// 	0x73, 0x65, 0x52, 0x1c, 0x6e, 0x6f, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x44, 0x65,
// 	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72,
// 	0x12, 0x25, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03,
// 	0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x0a, 0x64, 0x65, 0x70,
// 	0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x70, 0x5f, 0x65,
// 	0x6e, 0x74, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6d, 0x61, 0x70, 0x45,
// 	0x6e, 0x74, 0x72, 0x79, 0x12, 0x58, 0x0a, 0x14, 0x75, 0x6e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70,
// 	0x72, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0xe7, 0x07, 0x20,
// 	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
// 	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x6e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65,
// 	0x74, 0x65, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x75, 0x6e, 0x69, 0x6e, 0x74,
// 	0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x09,
// 	0x08, 0xe8, 0x07, 0x10, 0x80, 0x80, 0x80, 0x80, 0x02, 0x4a, 0x04, 0x08, 0x08, 0x10, 0x09, 0x4a,
// 	0x04, 0x08, 0x09, 0x10, 0x0a, 0x22, 0xe2, 0x03, 0x0a, 0x0c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
// 	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x41, 0x0a, 0x05, 0x63, 0x74, 0x79, 0x70, 0x65, 0x18,
// 	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
// 	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74,
// 	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x06, 0x53, 0x54, 0x52, 0x49,
// 	0x4e, 0x47, 0x52, 0x05, 0x63, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x63,
// 	0x6b, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x61, 0x63, 0x6b, 0x65,
// 	0x64, 0x12, 0x47, 0x0a, 0x06, 0x6a, 0x73, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
// 	0x0e, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
// 	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
// 	0x2e, 0x4a, 0x53, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x09, 0x4a, 0x53, 0x5f, 0x4e, 0x4f, 0x52, 0x4d,
// 	0x41, 0x4c, 0x52, 0x06, 0x6a, 0x73, 0x74, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x6c, 0x61,
// 	0x7a, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52,
// 	0x04, 0x6c, 0x61, 0x7a, 0x79, 0x12, 0x25, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61,
// 	0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65,
// 	0x52, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x04,
// 	0x77, 0x65, 0x61, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73,
// 	0x65, 0x52, 0x04, 0x77, 0x65, 0x61, 0x6b, 0x12, 0x58, 0x0a, 0x14, 0x75, 0x6e, 0x69, 0x6e, 0x74,
// 	0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
// 	0xe7, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
// 	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x6e, 0x69, 0x6e, 0x74, 0x65, 0x72,
// 	0x70, 0x72, 0x65, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x75, 0x6e,
// 	0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
// 	0x6e, 0x22, 0x2f, 0x0a, 0x05, 0x43, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54,
// 	0x52, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x4f, 0x52, 0x44, 0x10, 0x01,
// 	0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x49, 0x45, 0x43, 0x45,
// 	0x10, 0x02, 0x22, 0x35, 0x0a, 0x06, 0x4a, 0x53, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09,
// 	0x4a, 0x53, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4a,
// 	0x53, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4a, 0x53,
// 	0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x02, 0x2a, 0x09, 0x08, 0xe8, 0x07, 0x10, 0x80,
// 	0x80, 0x80, 0x80, 0x02, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x22, 0x73, 0x0a, 0x0c, 0x4f, 0x6e,
// 	0x65, 0x6f, 0x66, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x58, 0x0a, 0x14, 0x75, 0x6e,
// 	0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69,
// 	0x6f, 0x6e, 0x18, 0xe7, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
// 	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x6e, 0x69, 0x6e,
// 	0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
// 	0x13, 0x75, 0x6e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x64, 0x4f, 0x70,
// 	0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x09, 0x08, 0xe8, 0x07, 0x10, 0x80, 0x80, 0x80, 0x80, 0x02, 0x22,
// 	0xc0, 0x01, 0x0a, 0x0b, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
// 	0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x02,
// 	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x6c, 0x69, 0x61, 0x73,
// 	0x12, 0x25, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03,
// 	0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x0a, 0x64, 0x65, 0x70,
// 	0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x58, 0x0a, 0x14, 0x75, 0x6e, 0x69, 0x6e, 0x74,
// 	0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
// 	0xe7, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
// 	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x6e, 0x69, 0x6e, 0x74, 0x65, 0x72,
// 	0x70, 0x72, 0x65, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x75, 0x6e,
// 	0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
// 	0x6e, 0x2a, 0x09, 0x08, 0xe8, 0x07, 0x10, 0x80, 0x80, 0x80, 0x80, 0x02, 0x4a, 0x04, 0x08, 0x05,
// 	0x10, 0x06, 0x22, 0x9e, 0x01, 0x0a, 0x10, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65,
// 	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65,
// 	0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c,
// 	0x73, 0x65, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x58,
// 	0x0a, 0x14, 0x75, 0x6e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x64, 0x5f,
// 	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0xe7, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
// 	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
// 	0x55, 0x6e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x74,
// 	0x69, 0x6f, 0x6e, 0x52, 0x13, 0x75, 0x6e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74,
// 	0x65, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x09, 0x08, 0xe8, 0x07, 0x10, 0x80, 0x80,
// 	0x80, 0x80, 0x02, 0x22, 0x9c, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f,
// 	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63,
// 	0x61, 0x74, 0x65, 0x64, 0x18, 0x21, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73,
// 	0x65, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x58, 0x0a,
// 	0x14, 0x75, 0x6e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x6f,
// 	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0xe7, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67,
// 	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
// 	0x6e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x74, 0x69,
// 	0x6f, 0x6e, 0x52, 0x13, 0x75, 0x6e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65,
// 	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x09, 0x08, 0xe8, 0x07, 0x10, 0x80, 0x80, 0x80,
// 	0x80, 0x02, 0x22, 0xe0, 0x02, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74,
// 	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74,
// 	0x65, 0x64, 0x18, 0x21, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52,
// 	0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x71, 0x0a, 0x11, 0x69,
// 	0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
// 	0x18, 0x22, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
// 	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f,
// 	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e,
// 	0x63, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x3a, 0x13, 0x49, 0x44, 0x45, 0x4d, 0x50, 0x4f, 0x54,
// 	0x45, 0x4e, 0x43, 0x59, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x52, 0x10, 0x69, 0x64,
// 	0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x58,
// 	0x0a, 0x14, 0x75, 0x6e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x64, 0x5f,
// 	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0xe7, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
// 	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
// 	0x55, 0x6e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x74,
// 	0x69, 0x6f, 0x6e, 0x52, 0x13, 0x75, 0x6e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74,
// 	0x65, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x50, 0x0a, 0x10, 0x49, 0x64, 0x65, 0x6d,
// 	0x70, 0x6f, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x13,
// 	0x49, 0x44, 0x45, 0x4d, 0x50, 0x4f, 0x54, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
// 	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x4f, 0x5f, 0x53, 0x49, 0x44, 0x45,
// 	0x5f, 0x45, 0x46, 0x46, 0x45, 0x43, 0x54, 0x53, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x44,
// 	0x45, 0x4d, 0x50, 0x4f, 0x54, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x2a, 0x09, 0x08, 0xe8, 0x07, 0x10,
// 	0x80, 0x80, 0x80, 0x80, 0x02, 0x22, 0x9a, 0x03, 0x0a, 0x13, 0x55, 0x6e, 0x69, 0x6e, 0x74, 0x65,
// 	0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a,
// 	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f,
// 	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x6e,
// 	0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x65, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
// 	0x6e, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x72, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
// 	0x12, 0x29, 0x0a, 0x10, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x76,
// 	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x64, 0x65, 0x6e,
// 	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x70,
// 	0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75,
// 	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76,
// 	0x65, 0x49, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x6e, 0x65, 0x67,
// 	0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
// 	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x49,
// 	0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c,
// 	0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x64,
// 	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74,
// 	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c,
// 	0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x27, 0x0a,
// 	0x0f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
// 	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
// 	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x4a, 0x0a, 0x08, 0x4e, 0x61, 0x6d, 0x65, 0x50, 0x61,
// 	0x72, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x18,
// 	0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x72, 0x74, 0x12,
// 	0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18,
// 	0x02, 0x20, 0x02, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
// 	0x6f, 0x6e, 0x22, 0xa7, 0x02, 0x0a, 0x0e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x64,
// 	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x44, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
// 	0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
// 	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
// 	0x43, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
// 	0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0xce, 0x01, 0x0a, 0x08,
// 	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
// 	0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x42, 0x02, 0x10, 0x01, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
// 	0x12, 0x16, 0x0a, 0x04, 0x73, 0x70, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x42, 0x02,
// 	0x10, 0x01, 0x52, 0x04, 0x73, 0x70, 0x61, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x6c, 0x65, 0x61, 0x64,
// 	0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01,
// 	0x28, 0x09, 0x52, 0x0f, 0x6c, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
// 	0x6e, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x5f,
// 	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
// 	0x74, 0x72, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
// 	0x12, 0x3a, 0x0a, 0x19, 0x6c, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x74, 0x61,
// 	0x63, 0x68, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20,
// 	0x03, 0x28, 0x09, 0x52, 0x17, 0x6c, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61,
// 	0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xd1, 0x01, 0x0a,
// 	0x11, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x49, 0x6e,
// 	0x66, 0x6f, 0x12, 0x4d, 0x0a, 0x0a, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
// 	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
// 	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
// 	0x65, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
// 	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
// 	0x6e, 0x1a, 0x6d, 0x0a, 0x0a, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
// 	0x16, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x42, 0x02, 0x10,
// 	0x01, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63,
// 	0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f,
// 	0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x65, 0x67, 0x69,
// 	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x10,
// 	0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x65, 0x6e, 0x64,
// 	0x42, 0x57, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
// 	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x42, 0x10, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
// 	0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x48, 0x01, 0x5a, 0x06, 0x2e, 0x2f, 0x74,
// 	0x65, 0x73, 0x74, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x47, 0x50, 0x42, 0xaa, 0x02, 0x1a, 0x47,
// 	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52,
// 	0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
// }

// var (
// 	file_descriptor_proto_rawDescOnce sync.Once
// 	file_descriptor_proto_rawDescData = file_descriptor_proto_rawDesc
// )

// func file_descriptor_proto_rawDescGZIP() []byte {
// 	file_descriptor_proto_rawDescOnce.Do(func() {
// 		file_descriptor_proto_rawDescData = protoimpl.X.CompressGZIP(file_descriptor_proto_rawDescData)
// 	})
// 	return file_descriptor_proto_rawDescData
// }

// var file_descriptor_proto_enumTypes = make([]protoimpl.EnumInfo, 6)
// var file_descriptor_proto_msgTypes = make([]protoimpl.MessageInfo, 27)
// var file_descriptor_proto_goTypes = []interface{}{
// 	(FieldDescriptorProto_Type)(0),                // 0: google.protobuf.FieldDescriptorProto.Type
// 	(FieldDescriptorProto_Label)(0),               // 1: google.protobuf.FieldDescriptorProto.Label
// 	(FileOptions_OptimizeMode)(0),                 // 2: google.protobuf.FileOptions.OptimizeMode
// 	(FieldOptions_CType)(0),                       // 3: google.protobuf.FieldOptions.CType
// 	(FieldOptions_JSType)(0),                      // 4: google.protobuf.FieldOptions.JSType
// 	(MethodOptions_IdempotencyLevel)(0),           // 5: google.protobuf.MethodOptions.IdempotencyLevel
// 	(*FileDescriptorSet)(nil),                     // 6: google.protobuf.FileDescriptorSet
// 	(*FileDescriptorProto)(nil),                   // 7: google.protobuf.FileDescriptorProto
// 	(*DescriptorProto)(nil),                       // 8: google.protobuf.DescriptorProto
// 	(*ExtensionRangeOptions)(nil),                 // 9: google.protobuf.ExtensionRangeOptions
// 	(*FieldDescriptorProto)(nil),                  // 10: google.protobuf.FieldDescriptorProto
// 	(*OneofDescriptorProto)(nil),                  // 11: google.protobuf.OneofDescriptorProto
// 	(*EnumDescriptorProto)(nil),                   // 12: google.protobuf.EnumDescriptorProto
// 	(*EnumValueDescriptorProto)(nil),              // 13: google.protobuf.EnumValueDescriptorProto
// 	(*ServiceDescriptorProto)(nil),                // 14: google.protobuf.ServiceDescriptorProto
// 	(*MethodDescriptorProto)(nil),                 // 15: google.protobuf.MethodDescriptorProto
// 	(*FileOptions)(nil),                           // 16: google.protobuf.FileOptions
// 	(*MessageOptions)(nil),                        // 17: google.protobuf.MessageOptions
// 	(*FieldOptions)(nil),                          // 18: google.protobuf.FieldOptions
// 	(*OneofOptions)(nil),                          // 19: google.protobuf.OneofOptions
// 	(*EnumOptions)(nil),                           // 20: google.protobuf.EnumOptions
// 	(*EnumValueOptions)(nil),                      // 21: google.protobuf.EnumValueOptions
// 	(*ServiceOptions)(nil),                        // 22: google.protobuf.ServiceOptions
// 	(*MethodOptions)(nil),                         // 23: google.protobuf.MethodOptions
// 	(*UninterpretedOption)(nil),                   // 24: google.protobuf.UninterpretedOption
// 	(*SourceCodeInfo)(nil),                        // 25: google.protobuf.SourceCodeInfo
// 	(*GeneratedCodeInfo)(nil),                     // 26: google.protobuf.GeneratedCodeInfo
// 	(*DescriptorProto_ExtensionRange)(nil),        // 27: google.protobuf.DescriptorProto.ExtensionRange
// 	(*DescriptorProto_ReservedRange)(nil),         // 28: google.protobuf.DescriptorProto.ReservedRange
// 	(*EnumDescriptorProto_EnumReservedRange)(nil), // 29: google.protobuf.EnumDescriptorProto.EnumReservedRange
// 	(*UninterpretedOption_NamePart)(nil),          // 30: google.protobuf.UninterpretedOption.NamePart
// 	(*SourceCodeInfo_Location)(nil),               // 31: google.protobuf.SourceCodeInfo.Location
// 	(*GeneratedCodeInfo_Annotation)(nil),          // 32: google.protobuf.GeneratedCodeInfo.Annotation
// }
// var file_descriptor_proto_depIdxs = []int32{
// 	7,  // 0: google.protobuf.FileDescriptorSet.file:type_name -> google.protobuf.FileDescriptorProto
// 	8,  // 1: google.protobuf.FileDescriptorProto.message_type:type_name -> google.protobuf.DescriptorProto
// 	12, // 2: google.protobuf.FileDescriptorProto.enum_type:type_name -> google.protobuf.EnumDescriptorProto
// 	14, // 3: google.protobuf.FileDescriptorProto.service:type_name -> google.protobuf.ServiceDescriptorProto
// 	10, // 4: google.protobuf.FileDescriptorProto.extension:type_name -> google.protobuf.FieldDescriptorProto
// 	16, // 5: google.protobuf.FileDescriptorProto.options:type_name -> google.protobuf.FileOptions
// 	25, // 6: google.protobuf.FileDescriptorProto.source_code_info:type_name -> google.protobuf.SourceCodeInfo
// 	10, // 7: google.protobuf.DescriptorProto.field:type_name -> google.protobuf.FieldDescriptorProto
// 	10, // 8: google.protobuf.DescriptorProto.extension:type_name -> google.protobuf.FieldDescriptorProto
// 	8,  // 9: google.protobuf.DescriptorProto.nested_type:type_name -> google.protobuf.DescriptorProto
// 	12, // 10: google.protobuf.DescriptorProto.enum_type:type_name -> google.protobuf.EnumDescriptorProto
// 	27, // 11: google.protobuf.DescriptorProto.extension_range:type_name -> google.protobuf.DescriptorProto.ExtensionRange
// 	11, // 12: google.protobuf.DescriptorProto.oneof_decl:type_name -> google.protobuf.OneofDescriptorProto
// 	17, // 13: google.protobuf.DescriptorProto.options:type_name -> google.protobuf.MessageOptions
// 	28, // 14: google.protobuf.DescriptorProto.reserved_range:type_name -> google.protobuf.DescriptorProto.ReservedRange
// 	24, // 15: google.protobuf.ExtensionRangeOptions.uninterpreted_option:type_name -> google.protobuf.UninterpretedOption
// 	1,  // 16: google.protobuf.FieldDescriptorProto.label:type_name -> google.protobuf.FieldDescriptorProto.Label
// 	0,  // 17: google.protobuf.FieldDescriptorProto.type:type_name -> google.protobuf.FieldDescriptorProto.Type
// 	18, // 18: google.protobuf.FieldDescriptorProto.options:type_name -> google.protobuf.FieldOptions
// 	19, // 19: google.protobuf.OneofDescriptorProto.options:type_name -> google.protobuf.OneofOptions
// 	13, // 20: google.protobuf.EnumDescriptorProto.value:type_name -> google.protobuf.EnumValueDescriptorProto
// 	20, // 21: google.protobuf.EnumDescriptorProto.options:type_name -> google.protobuf.EnumOptions
// 	29, // 22: google.protobuf.EnumDescriptorProto.reserved_range:type_name -> google.protobuf.EnumDescriptorProto.EnumReservedRange
// 	21, // 23: google.protobuf.EnumValueDescriptorProto.options:type_name -> google.protobuf.EnumValueOptions
// 	15, // 24: google.protobuf.ServiceDescriptorProto.method:type_name -> google.protobuf.MethodDescriptorProto
// 	22, // 25: google.protobuf.ServiceDescriptorProto.options:type_name -> google.protobuf.ServiceOptions
// 	23, // 26: google.protobuf.MethodDescriptorProto.options:type_name -> google.protobuf.MethodOptions
// 	2,  // 27: google.protobuf.FileOptions.optimize_for:type_name -> google.protobuf.FileOptions.OptimizeMode
// 	24, // 28: google.protobuf.FileOptions.uninterpreted_option:type_name -> google.protobuf.UninterpretedOption
// 	24, // 29: google.protobuf.MessageOptions.uninterpreted_option:type_name -> google.protobuf.UninterpretedOption
// 	3,  // 30: google.protobuf.FieldOptions.ctype:type_name -> google.protobuf.FieldOptions.CType
// 	4,  // 31: google.protobuf.FieldOptions.jstype:type_name -> google.protobuf.FieldOptions.JSType
// 	24, // 32: google.protobuf.FieldOptions.uninterpreted_option:type_name -> google.protobuf.UninterpretedOption
// 	24, // 33: google.protobuf.OneofOptions.uninterpreted_option:type_name -> google.protobuf.UninterpretedOption
// 	24, // 34: google.protobuf.EnumOptions.uninterpreted_option:type_name -> google.protobuf.UninterpretedOption
// 	24, // 35: google.protobuf.EnumValueOptions.uninterpreted_option:type_name -> google.protobuf.UninterpretedOption
// 	24, // 36: google.protobuf.ServiceOptions.uninterpreted_option:type_name -> google.protobuf.UninterpretedOption
// 	5,  // 37: google.protobuf.MethodOptions.idempotency_level:type_name -> google.protobuf.MethodOptions.IdempotencyLevel
// 	24, // 38: google.protobuf.MethodOptions.uninterpreted_option:type_name -> google.protobuf.UninterpretedOption
// 	30, // 39: google.protobuf.UninterpretedOption.name:type_name -> google.protobuf.UninterpretedOption.NamePart
// 	31, // 40: google.protobuf.SourceCodeInfo.location:type_name -> google.protobuf.SourceCodeInfo.Location
// 	32, // 41: google.protobuf.GeneratedCodeInfo.annotation:type_name -> google.protobuf.GeneratedCodeInfo.Annotation
// 	9,  // 42: google.protobuf.DescriptorProto.ExtensionRange.options:type_name -> google.protobuf.ExtensionRangeOptions
// 	43, // [43:43] is the sub-list for method output_type
// 	43, // [43:43] is the sub-list for method input_type
// 	43, // [43:43] is the sub-list for extension type_name
// 	43, // [43:43] is the sub-list for extension extendee
// 	0,  // [0:43] is the sub-list for field type_name
// }

// func init() { file_descriptor_proto_init() }
// func file_descriptor_proto_init() {
// 	if File_descriptor_proto != nil {
// 		return
// 	}
// 	if !protoimpl.UnsafeEnabled {
// 		file_descriptor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*FileDescriptorSet); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*FileDescriptorProto); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*DescriptorProto); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*ExtensionRangeOptions); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			case 3:
// 				return &v.extensionFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*FieldDescriptorProto); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*OneofDescriptorProto); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*EnumDescriptorProto); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*EnumValueDescriptorProto); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*ServiceDescriptorProto); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*MethodDescriptorProto); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*FileOptions); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			case 3:
// 				return &v.extensionFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*MessageOptions); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			case 3:
// 				return &v.extensionFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*FieldOptions); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			case 3:
// 				return &v.extensionFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*OneofOptions); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			case 3:
// 				return &v.extensionFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*EnumOptions); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			case 3:
// 				return &v.extensionFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*EnumValueOptions); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			case 3:
// 				return &v.extensionFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*ServiceOptions); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			case 3:
// 				return &v.extensionFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*MethodOptions); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			case 3:
// 				return &v.extensionFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*UninterpretedOption); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[19].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*SourceCodeInfo); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[20].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*GeneratedCodeInfo); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[21].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*DescriptorProto_ExtensionRange); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[22].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*DescriptorProto_ReservedRange); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[23].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*EnumDescriptorProto_EnumReservedRange); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[24].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*UninterpretedOption_NamePart); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[25].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*SourceCodeInfo_Location); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			default:
// 				return nil
// 			}
// 		}
// 		file_descriptor_proto_msgTypes[26].Exporter = func(v interface{}, i int) interface{} {
// 			switch v := v.(*GeneratedCodeInfo_Annotation); i {
// 			case 0:
// 				return &v.state
// 			case 1:
// 				return &v.sizeCache
// 			case 2:
// 				return &v.unknownFields
// 			default:
// 				return nil
// 			}
// 		}
// 	}
// 	type x struct{}
// 	out := protoimpl.TypeBuilder{
// 		File: protoimpl.DescBuilder{
// 			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
// 			RawDescriptor: file_descriptor_proto_rawDesc,
// 			NumEnums:      6,
// 			NumMessages:   27,
// 			NumExtensions: 0,
// 			NumServices:   0,
// 		},
// 		GoTypes:           file_descriptor_proto_goTypes,
// 		DependencyIndexes: file_descriptor_proto_depIdxs,
// 		EnumInfos:         file_descriptor_proto_enumTypes,
// 		MessageInfos:      file_descriptor_proto_msgTypes,
// 	}.Build()
// 	File_descriptor_proto = out.File
// 	file_descriptor_proto_rawDesc = nil
// 	file_descriptor_proto_goTypes = nil
// 	file_descriptor_proto_depIdxs = nil
// }
