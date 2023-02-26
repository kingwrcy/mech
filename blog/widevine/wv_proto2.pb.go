package widevine

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// todo: fill (for this top-level type, it might be impossible/difficult)
type LicenseType int32

const (
	LicenseType_ZERO    LicenseType = 0
	LicenseType_DEFAULT LicenseType = 1 // 1 is STREAMING/temporary license; on recent versions may go up to 3 (latest x86); it might be persist/don't persist type, unconfirmed
	LicenseType_OFFLINE LicenseType = 2
)

// Enum value maps for LicenseType.
var (
	LicenseType_name = map[int32]string{
		0: "ZERO",
		1: "DEFAULT",
		2: "OFFLINE",
	}
	LicenseType_value = map[string]int32{
		"ZERO":    0,
		"DEFAULT": 1,
		"OFFLINE": 2,
	}
)

func (x LicenseType) Enum() *LicenseType {
	p := new(LicenseType)
	*p = x
	return p
}

func (x LicenseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LicenseType) Descriptor() protoreflect.EnumDescriptor {
	return file_wv_proto2_proto_enumTypes[0].Descriptor()
}

func (LicenseType) Type() protoreflect.EnumType {
	return &file_wv_proto2_proto_enumTypes[0]
}

func (x LicenseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *LicenseType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = LicenseType(num)
	return nil
}

// Deprecated: Use LicenseType.Descriptor instead.
func (LicenseType) EnumDescriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{0}
}

// todo: fill (for this top-level type, it might be impossible/difficult)
// this is just a guess because these globals got lost, but really, do we need more?
type ProtocolVersion int32

const (
	ProtocolVersion_CURRENT ProtocolVersion = 21 // don't have symbols for this
)

// Enum value maps for ProtocolVersion.
var (
	ProtocolVersion_name = map[int32]string{
		21: "CURRENT",
	}
	ProtocolVersion_value = map[string]int32{
		"CURRENT": 21,
	}
)

func (x ProtocolVersion) Enum() *ProtocolVersion {
	p := new(ProtocolVersion)
	*p = x
	return p
}

func (x ProtocolVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProtocolVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_wv_proto2_proto_enumTypes[1].Descriptor()
}

func (ProtocolVersion) Type() protoreflect.EnumType {
	return &file_wv_proto2_proto_enumTypes[1]
}

func (x ProtocolVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ProtocolVersion) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ProtocolVersion(num)
	return nil
}

// Deprecated: Use ProtocolVersion.Descriptor instead.
func (ProtocolVersion) EnumDescriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{1}
}

type ClientIdentification_TokenType int32

const (
	ClientIdentification_KEYBOX                         ClientIdentification_TokenType = 0
	ClientIdentification_DEVICE_CERTIFICATE             ClientIdentification_TokenType = 1
	ClientIdentification_REMOTE_ATTESTATION_CERTIFICATE ClientIdentification_TokenType = 2
)

// Enum value maps for ClientIdentification_TokenType.
var (
	ClientIdentification_TokenType_name = map[int32]string{
		0: "KEYBOX",
		1: "DEVICE_CERTIFICATE",
		2: "REMOTE_ATTESTATION_CERTIFICATE",
	}
	ClientIdentification_TokenType_value = map[string]int32{
		"KEYBOX":                         0,
		"DEVICE_CERTIFICATE":             1,
		"REMOTE_ATTESTATION_CERTIFICATE": 2,
	}
)

func (x ClientIdentification_TokenType) Enum() *ClientIdentification_TokenType {
	p := new(ClientIdentification_TokenType)
	*p = x
	return p
}

func (x ClientIdentification_TokenType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientIdentification_TokenType) Descriptor() protoreflect.EnumDescriptor {
	return file_wv_proto2_proto_enumTypes[2].Descriptor()
}

func (ClientIdentification_TokenType) Type() protoreflect.EnumType {
	return &file_wv_proto2_proto_enumTypes[2]
}

func (x ClientIdentification_TokenType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ClientIdentification_TokenType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ClientIdentification_TokenType(num)
	return nil
}

// Deprecated: Use ClientIdentification_TokenType.Descriptor instead.
func (ClientIdentification_TokenType) EnumDescriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{0, 0}
}

type ClientIdentification_ClientCapabilities_HdcpVersion int32

const (
	ClientIdentification_ClientCapabilities_HDCP_NONE ClientIdentification_ClientCapabilities_HdcpVersion = 0
	ClientIdentification_ClientCapabilities_HDCP_V1   ClientIdentification_ClientCapabilities_HdcpVersion = 1
	ClientIdentification_ClientCapabilities_HDCP_V2   ClientIdentification_ClientCapabilities_HdcpVersion = 2
	ClientIdentification_ClientCapabilities_HDCP_V2_1 ClientIdentification_ClientCapabilities_HdcpVersion = 3
	ClientIdentification_ClientCapabilities_HDCP_V2_2 ClientIdentification_ClientCapabilities_HdcpVersion = 4
)

// Enum value maps for ClientIdentification_ClientCapabilities_HdcpVersion.
var (
	ClientIdentification_ClientCapabilities_HdcpVersion_name = map[int32]string{
		0: "HDCP_NONE",
		1: "HDCP_V1",
		2: "HDCP_V2",
		3: "HDCP_V2_1",
		4: "HDCP_V2_2",
	}
	ClientIdentification_ClientCapabilities_HdcpVersion_value = map[string]int32{
		"HDCP_NONE": 0,
		"HDCP_V1":   1,
		"HDCP_V2":   2,
		"HDCP_V2_1": 3,
		"HDCP_V2_2": 4,
	}
)

func (x ClientIdentification_ClientCapabilities_HdcpVersion) Enum() *ClientIdentification_ClientCapabilities_HdcpVersion {
	p := new(ClientIdentification_ClientCapabilities_HdcpVersion)
	*p = x
	return p
}

func (x ClientIdentification_ClientCapabilities_HdcpVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientIdentification_ClientCapabilities_HdcpVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_wv_proto2_proto_enumTypes[3].Descriptor()
}

func (ClientIdentification_ClientCapabilities_HdcpVersion) Type() protoreflect.EnumType {
	return &file_wv_proto2_proto_enumTypes[3]
}

func (x ClientIdentification_ClientCapabilities_HdcpVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ClientIdentification_ClientCapabilities_HdcpVersion) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ClientIdentification_ClientCapabilities_HdcpVersion(num)
	return nil
}

// Deprecated: Use ClientIdentification_ClientCapabilities_HdcpVersion.Descriptor instead.
func (ClientIdentification_ClientCapabilities_HdcpVersion) EnumDescriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{0, 1, 0}
}

type DeviceCertificate_CertificateType int32

const (
	DeviceCertificate_ROOT         DeviceCertificate_CertificateType = 0
	DeviceCertificate_INTERMEDIATE DeviceCertificate_CertificateType = 1
	DeviceCertificate_USER_DEVICE  DeviceCertificate_CertificateType = 2
	DeviceCertificate_SERVICE      DeviceCertificate_CertificateType = 3
)

// Enum value maps for DeviceCertificate_CertificateType.
var (
	DeviceCertificate_CertificateType_name = map[int32]string{
		0: "ROOT",
		1: "INTERMEDIATE",
		2: "USER_DEVICE",
		3: "SERVICE",
	}
	DeviceCertificate_CertificateType_value = map[string]int32{
		"ROOT":         0,
		"INTERMEDIATE": 1,
		"USER_DEVICE":  2,
		"SERVICE":      3,
	}
)

func (x DeviceCertificate_CertificateType) Enum() *DeviceCertificate_CertificateType {
	p := new(DeviceCertificate_CertificateType)
	*p = x
	return p
}

func (x DeviceCertificate_CertificateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceCertificate_CertificateType) Descriptor() protoreflect.EnumDescriptor {
	return file_wv_proto2_proto_enumTypes[4].Descriptor()
}

func (DeviceCertificate_CertificateType) Type() protoreflect.EnumType {
	return &file_wv_proto2_proto_enumTypes[4]
}

func (x DeviceCertificate_CertificateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *DeviceCertificate_CertificateType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = DeviceCertificate_CertificateType(num)
	return nil
}

// Deprecated: Use DeviceCertificate_CertificateType.Descriptor instead.
func (DeviceCertificate_CertificateType) EnumDescriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{1, 0}
}

type DeviceCertificateStatus_CertificateStatus int32

const (
	DeviceCertificateStatus_VALID   DeviceCertificateStatus_CertificateStatus = 0
	DeviceCertificateStatus_REVOKED DeviceCertificateStatus_CertificateStatus = 1
)

// Enum value maps for DeviceCertificateStatus_CertificateStatus.
var (
	DeviceCertificateStatus_CertificateStatus_name = map[int32]string{
		0: "VALID",
		1: "REVOKED",
	}
	DeviceCertificateStatus_CertificateStatus_value = map[string]int32{
		"VALID":   0,
		"REVOKED": 1,
	}
)

func (x DeviceCertificateStatus_CertificateStatus) Enum() *DeviceCertificateStatus_CertificateStatus {
	p := new(DeviceCertificateStatus_CertificateStatus)
	*p = x
	return p
}

func (x DeviceCertificateStatus_CertificateStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceCertificateStatus_CertificateStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_wv_proto2_proto_enumTypes[5].Descriptor()
}

func (DeviceCertificateStatus_CertificateStatus) Type() protoreflect.EnumType {
	return &file_wv_proto2_proto_enumTypes[5]
}

func (x DeviceCertificateStatus_CertificateStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *DeviceCertificateStatus_CertificateStatus) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = DeviceCertificateStatus_CertificateStatus(num)
	return nil
}

// Deprecated: Use DeviceCertificateStatus_CertificateStatus.Descriptor instead.
func (DeviceCertificateStatus_CertificateStatus) EnumDescriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{2, 0}
}

type License_KeyContainer_KeyType int32

const (
	License_KeyContainer_SIGNING          License_KeyContainer_KeyType = 1
	License_KeyContainer_CONTENT          License_KeyContainer_KeyType = 2
	License_KeyContainer_KEY_CONTROL      License_KeyContainer_KeyType = 3
	License_KeyContainer_OPERATOR_SESSION License_KeyContainer_KeyType = 4
)

// Enum value maps for License_KeyContainer_KeyType.
var (
	License_KeyContainer_KeyType_name = map[int32]string{
		1: "SIGNING",
		2: "CONTENT",
		3: "KEY_CONTROL",
		4: "OPERATOR_SESSION",
	}
	License_KeyContainer_KeyType_value = map[string]int32{
		"SIGNING":          1,
		"CONTENT":          2,
		"KEY_CONTROL":      3,
		"OPERATOR_SESSION": 4,
	}
)

func (x License_KeyContainer_KeyType) Enum() *License_KeyContainer_KeyType {
	p := new(License_KeyContainer_KeyType)
	*p = x
	return p
}

func (x License_KeyContainer_KeyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (License_KeyContainer_KeyType) Descriptor() protoreflect.EnumDescriptor {
	return file_wv_proto2_proto_enumTypes[6].Descriptor()
}

func (License_KeyContainer_KeyType) Type() protoreflect.EnumType {
	return &file_wv_proto2_proto_enumTypes[6]
}

func (x License_KeyContainer_KeyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *License_KeyContainer_KeyType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = License_KeyContainer_KeyType(num)
	return nil
}

// Deprecated: Use License_KeyContainer_KeyType.Descriptor instead.
func (License_KeyContainer_KeyType) EnumDescriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{6, 1, 0}
}

type License_KeyContainer_SecurityLevel int32

const (
	License_KeyContainer_SW_SECURE_CRYPTO License_KeyContainer_SecurityLevel = 1
	License_KeyContainer_SW_SECURE_DECODE License_KeyContainer_SecurityLevel = 2
	License_KeyContainer_HW_SECURE_CRYPTO License_KeyContainer_SecurityLevel = 3
	License_KeyContainer_HW_SECURE_DECODE License_KeyContainer_SecurityLevel = 4
	License_KeyContainer_HW_SECURE_ALL    License_KeyContainer_SecurityLevel = 5
)

// Enum value maps for License_KeyContainer_SecurityLevel.
var (
	License_KeyContainer_SecurityLevel_name = map[int32]string{
		1: "SW_SECURE_CRYPTO",
		2: "SW_SECURE_DECODE",
		3: "HW_SECURE_CRYPTO",
		4: "HW_SECURE_DECODE",
		5: "HW_SECURE_ALL",
	}
	License_KeyContainer_SecurityLevel_value = map[string]int32{
		"SW_SECURE_CRYPTO": 1,
		"SW_SECURE_DECODE": 2,
		"HW_SECURE_CRYPTO": 3,
		"HW_SECURE_DECODE": 4,
		"HW_SECURE_ALL":    5,
	}
)

func (x License_KeyContainer_SecurityLevel) Enum() *License_KeyContainer_SecurityLevel {
	p := new(License_KeyContainer_SecurityLevel)
	*p = x
	return p
}

func (x License_KeyContainer_SecurityLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (License_KeyContainer_SecurityLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_wv_proto2_proto_enumTypes[7].Descriptor()
}

func (License_KeyContainer_SecurityLevel) Type() protoreflect.EnumType {
	return &file_wv_proto2_proto_enumTypes[7]
}

func (x License_KeyContainer_SecurityLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *License_KeyContainer_SecurityLevel) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = License_KeyContainer_SecurityLevel(num)
	return nil
}

// Deprecated: Use License_KeyContainer_SecurityLevel.Descriptor instead.
func (License_KeyContainer_SecurityLevel) EnumDescriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{6, 1, 1}
}

type License_KeyContainer_OutputProtection_CGMS int32

const (
	License_KeyContainer_OutputProtection_COPY_FREE  License_KeyContainer_OutputProtection_CGMS = 0
	License_KeyContainer_OutputProtection_COPY_ONCE  License_KeyContainer_OutputProtection_CGMS = 2
	License_KeyContainer_OutputProtection_COPY_NEVER License_KeyContainer_OutputProtection_CGMS = 3
	License_KeyContainer_OutputProtection_CGMS_NONE  License_KeyContainer_OutputProtection_CGMS = 42 // PC default!
)

// Enum value maps for License_KeyContainer_OutputProtection_CGMS.
var (
	License_KeyContainer_OutputProtection_CGMS_name = map[int32]string{
		0:  "COPY_FREE",
		2:  "COPY_ONCE",
		3:  "COPY_NEVER",
		42: "CGMS_NONE",
	}
	License_KeyContainer_OutputProtection_CGMS_value = map[string]int32{
		"COPY_FREE":  0,
		"COPY_ONCE":  2,
		"COPY_NEVER": 3,
		"CGMS_NONE":  42,
	}
)

func (x License_KeyContainer_OutputProtection_CGMS) Enum() *License_KeyContainer_OutputProtection_CGMS {
	p := new(License_KeyContainer_OutputProtection_CGMS)
	*p = x
	return p
}

func (x License_KeyContainer_OutputProtection_CGMS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (License_KeyContainer_OutputProtection_CGMS) Descriptor() protoreflect.EnumDescriptor {
	return file_wv_proto2_proto_enumTypes[8].Descriptor()
}

func (License_KeyContainer_OutputProtection_CGMS) Type() protoreflect.EnumType {
	return &file_wv_proto2_proto_enumTypes[8]
}

func (x License_KeyContainer_OutputProtection_CGMS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *License_KeyContainer_OutputProtection_CGMS) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = License_KeyContainer_OutputProtection_CGMS(num)
	return nil
}

// Deprecated: Use License_KeyContainer_OutputProtection_CGMS.Descriptor instead.
func (License_KeyContainer_OutputProtection_CGMS) EnumDescriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{6, 1, 0, 0}
}

type LicenseError_Error int32

const (
	LicenseError_INVALID_DEVICE_CERTIFICATE LicenseError_Error = 1
	LicenseError_REVOKED_DEVICE_CERTIFICATE LicenseError_Error = 2
	LicenseError_SERVICE_UNAVAILABLE        LicenseError_Error = 3
)

// Enum value maps for LicenseError_Error.
var (
	LicenseError_Error_name = map[int32]string{
		1: "INVALID_DEVICE_CERTIFICATE",
		2: "REVOKED_DEVICE_CERTIFICATE",
		3: "SERVICE_UNAVAILABLE",
	}
	LicenseError_Error_value = map[string]int32{
		"INVALID_DEVICE_CERTIFICATE": 1,
		"REVOKED_DEVICE_CERTIFICATE": 2,
		"SERVICE_UNAVAILABLE":        3,
	}
)

func (x LicenseError_Error) Enum() *LicenseError_Error {
	p := new(LicenseError_Error)
	*p = x
	return p
}

func (x LicenseError_Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LicenseError_Error) Descriptor() protoreflect.EnumDescriptor {
	return file_wv_proto2_proto_enumTypes[9].Descriptor()
}

func (LicenseError_Error) Type() protoreflect.EnumType {
	return &file_wv_proto2_proto_enumTypes[9]
}

func (x LicenseError_Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *LicenseError_Error) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = LicenseError_Error(num)
	return nil
}

// Deprecated: Use LicenseError_Error.Descriptor instead.
func (LicenseError_Error) EnumDescriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{7, 0}
}

type LicenseRequest_RequestType int32

const (
	LicenseRequest_NEW     LicenseRequest_RequestType = 1
	LicenseRequest_RENEWAL LicenseRequest_RequestType = 2
	LicenseRequest_RELEASE LicenseRequest_RequestType = 3
)

// Enum value maps for LicenseRequest_RequestType.
var (
	LicenseRequest_RequestType_name = map[int32]string{
		1: "NEW",
		2: "RENEWAL",
		3: "RELEASE",
	}
	LicenseRequest_RequestType_value = map[string]int32{
		"NEW":     1,
		"RENEWAL": 2,
		"RELEASE": 3,
	}
)

func (x LicenseRequest_RequestType) Enum() *LicenseRequest_RequestType {
	p := new(LicenseRequest_RequestType)
	*p = x
	return p
}

func (x LicenseRequest_RequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LicenseRequest_RequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_wv_proto2_proto_enumTypes[10].Descriptor()
}

func (LicenseRequest_RequestType) Type() protoreflect.EnumType {
	return &file_wv_proto2_proto_enumTypes[10]
}

func (x LicenseRequest_RequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *LicenseRequest_RequestType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = LicenseRequest_RequestType(num)
	return nil
}

// Deprecated: Use LicenseRequest_RequestType.Descriptor instead.
func (LicenseRequest_RequestType) EnumDescriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{8, 0}
}

type LicenseRequestRaw_RequestType int32

const (
	LicenseRequestRaw_NEW     LicenseRequestRaw_RequestType = 1
	LicenseRequestRaw_RENEWAL LicenseRequestRaw_RequestType = 2
	LicenseRequestRaw_RELEASE LicenseRequestRaw_RequestType = 3
)

// Enum value maps for LicenseRequestRaw_RequestType.
var (
	LicenseRequestRaw_RequestType_name = map[int32]string{
		1: "NEW",
		2: "RENEWAL",
		3: "RELEASE",
	}
	LicenseRequestRaw_RequestType_value = map[string]int32{
		"NEW":     1,
		"RENEWAL": 2,
		"RELEASE": 3,
	}
)

func (x LicenseRequestRaw_RequestType) Enum() *LicenseRequestRaw_RequestType {
	p := new(LicenseRequestRaw_RequestType)
	*p = x
	return p
}

func (x LicenseRequestRaw_RequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LicenseRequestRaw_RequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_wv_proto2_proto_enumTypes[11].Descriptor()
}

func (LicenseRequestRaw_RequestType) Type() protoreflect.EnumType {
	return &file_wv_proto2_proto_enumTypes[11]
}

func (x LicenseRequestRaw_RequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *LicenseRequestRaw_RequestType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = LicenseRequestRaw_RequestType(num)
	return nil
}

// Deprecated: Use LicenseRequestRaw_RequestType.Descriptor instead.
func (LicenseRequestRaw_RequestType) EnumDescriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{9, 0}
}

type ProvisionedDeviceInfo_WvSecurityLevel int32

const (
	ProvisionedDeviceInfo_LEVEL_UNSPECIFIED ProvisionedDeviceInfo_WvSecurityLevel = 0
	ProvisionedDeviceInfo_LEVEL_1           ProvisionedDeviceInfo_WvSecurityLevel = 1
	ProvisionedDeviceInfo_LEVEL_2           ProvisionedDeviceInfo_WvSecurityLevel = 2
	ProvisionedDeviceInfo_LEVEL_3           ProvisionedDeviceInfo_WvSecurityLevel = 3
)

// Enum value maps for ProvisionedDeviceInfo_WvSecurityLevel.
var (
	ProvisionedDeviceInfo_WvSecurityLevel_name = map[int32]string{
		0: "LEVEL_UNSPECIFIED",
		1: "LEVEL_1",
		2: "LEVEL_2",
		3: "LEVEL_3",
	}
	ProvisionedDeviceInfo_WvSecurityLevel_value = map[string]int32{
		"LEVEL_UNSPECIFIED": 0,
		"LEVEL_1":           1,
		"LEVEL_2":           2,
		"LEVEL_3":           3,
	}
)

func (x ProvisionedDeviceInfo_WvSecurityLevel) Enum() *ProvisionedDeviceInfo_WvSecurityLevel {
	p := new(ProvisionedDeviceInfo_WvSecurityLevel)
	*p = x
	return p
}

func (x ProvisionedDeviceInfo_WvSecurityLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProvisionedDeviceInfo_WvSecurityLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_wv_proto2_proto_enumTypes[12].Descriptor()
}

func (ProvisionedDeviceInfo_WvSecurityLevel) Type() protoreflect.EnumType {
	return &file_wv_proto2_proto_enumTypes[12]
}

func (x ProvisionedDeviceInfo_WvSecurityLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ProvisionedDeviceInfo_WvSecurityLevel) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ProvisionedDeviceInfo_WvSecurityLevel(num)
	return nil
}

// Deprecated: Use ProvisionedDeviceInfo_WvSecurityLevel.Descriptor instead.
func (ProvisionedDeviceInfo_WvSecurityLevel) EnumDescriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{10, 0}
}

type SignedMessage_MessageType int32

const (
	SignedMessage_LICENSE_REQUEST             SignedMessage_MessageType = 1
	SignedMessage_LICENSE                     SignedMessage_MessageType = 2
	SignedMessage_ERROR_RESPONSE              SignedMessage_MessageType = 3
	SignedMessage_SERVICE_CERTIFICATE_REQUEST SignedMessage_MessageType = 4
	SignedMessage_SERVICE_CERTIFICATE         SignedMessage_MessageType = 5
)

// Enum value maps for SignedMessage_MessageType.
var (
	SignedMessage_MessageType_name = map[int32]string{
		1: "LICENSE_REQUEST",
		2: "LICENSE",
		3: "ERROR_RESPONSE",
		4: "SERVICE_CERTIFICATE_REQUEST",
		5: "SERVICE_CERTIFICATE",
	}
	SignedMessage_MessageType_value = map[string]int32{
		"LICENSE_REQUEST":             1,
		"LICENSE":                     2,
		"ERROR_RESPONSE":              3,
		"SERVICE_CERTIFICATE_REQUEST": 4,
		"SERVICE_CERTIFICATE":         5,
	}
)

func (x SignedMessage_MessageType) Enum() *SignedMessage_MessageType {
	p := new(SignedMessage_MessageType)
	*p = x
	return p
}

func (x SignedMessage_MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SignedMessage_MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_wv_proto2_proto_enumTypes[13].Descriptor()
}

func (SignedMessage_MessageType) Type() protoreflect.EnumType {
	return &file_wv_proto2_proto_enumTypes[13]
}

func (x SignedMessage_MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *SignedMessage_MessageType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = SignedMessage_MessageType(num)
	return nil
}

// Deprecated: Use SignedMessage_MessageType.Descriptor instead.
func (SignedMessage_MessageType) EnumDescriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{20, 0}
}

type WidevineCencHeader_Algorithm int32

const (
	WidevineCencHeader_UNENCRYPTED WidevineCencHeader_Algorithm = 0
	WidevineCencHeader_AESCTR      WidevineCencHeader_Algorithm = 1
)

// Enum value maps for WidevineCencHeader_Algorithm.
var (
	WidevineCencHeader_Algorithm_name = map[int32]string{
		0: "UNENCRYPTED",
		1: "AESCTR",
	}
	WidevineCencHeader_Algorithm_value = map[string]int32{
		"UNENCRYPTED": 0,
		"AESCTR":      1,
	}
)

func (x WidevineCencHeader_Algorithm) Enum() *WidevineCencHeader_Algorithm {
	p := new(WidevineCencHeader_Algorithm)
	*p = x
	return p
}

func (x WidevineCencHeader_Algorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WidevineCencHeader_Algorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_wv_proto2_proto_enumTypes[14].Descriptor()
}

func (WidevineCencHeader_Algorithm) Type() protoreflect.EnumType {
	return &file_wv_proto2_proto_enumTypes[14]
}

func (x WidevineCencHeader_Algorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *WidevineCencHeader_Algorithm) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = WidevineCencHeader_Algorithm(num)
	return nil
}

// Deprecated: Use WidevineCencHeader_Algorithm.Descriptor instead.
func (WidevineCencHeader_Algorithm) EnumDescriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{21, 0}
}

type SignedLicenseRequest_MessageType int32

const (
	SignedLicenseRequest_LICENSE_REQUEST             SignedLicenseRequest_MessageType = 1
	SignedLicenseRequest_LICENSE                     SignedLicenseRequest_MessageType = 2
	SignedLicenseRequest_ERROR_RESPONSE              SignedLicenseRequest_MessageType = 3
	SignedLicenseRequest_SERVICE_CERTIFICATE_REQUEST SignedLicenseRequest_MessageType = 4
	SignedLicenseRequest_SERVICE_CERTIFICATE         SignedLicenseRequest_MessageType = 5
)

// Enum value maps for SignedLicenseRequest_MessageType.
var (
	SignedLicenseRequest_MessageType_name = map[int32]string{
		1: "LICENSE_REQUEST",
		2: "LICENSE",
		3: "ERROR_RESPONSE",
		4: "SERVICE_CERTIFICATE_REQUEST",
		5: "SERVICE_CERTIFICATE",
	}
	SignedLicenseRequest_MessageType_value = map[string]int32{
		"LICENSE_REQUEST":             1,
		"LICENSE":                     2,
		"ERROR_RESPONSE":              3,
		"SERVICE_CERTIFICATE_REQUEST": 4,
		"SERVICE_CERTIFICATE":         5,
	}
)

func (x SignedLicenseRequest_MessageType) Enum() *SignedLicenseRequest_MessageType {
	p := new(SignedLicenseRequest_MessageType)
	*p = x
	return p
}

func (x SignedLicenseRequest_MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SignedLicenseRequest_MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_wv_proto2_proto_enumTypes[15].Descriptor()
}

func (SignedLicenseRequest_MessageType) Type() protoreflect.EnumType {
	return &file_wv_proto2_proto_enumTypes[15]
}

func (x SignedLicenseRequest_MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *SignedLicenseRequest_MessageType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = SignedLicenseRequest_MessageType(num)
	return nil
}

// Deprecated: Use SignedLicenseRequest_MessageType.Descriptor instead.
func (SignedLicenseRequest_MessageType) EnumDescriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{22, 0}
}

type SignedLicenseRequestRaw_MessageType int32

const (
	SignedLicenseRequestRaw_LICENSE_REQUEST             SignedLicenseRequestRaw_MessageType = 1
	SignedLicenseRequestRaw_LICENSE                     SignedLicenseRequestRaw_MessageType = 2
	SignedLicenseRequestRaw_ERROR_RESPONSE              SignedLicenseRequestRaw_MessageType = 3
	SignedLicenseRequestRaw_SERVICE_CERTIFICATE_REQUEST SignedLicenseRequestRaw_MessageType = 4
	SignedLicenseRequestRaw_SERVICE_CERTIFICATE         SignedLicenseRequestRaw_MessageType = 5
)

// Enum value maps for SignedLicenseRequestRaw_MessageType.
var (
	SignedLicenseRequestRaw_MessageType_name = map[int32]string{
		1: "LICENSE_REQUEST",
		2: "LICENSE",
		3: "ERROR_RESPONSE",
		4: "SERVICE_CERTIFICATE_REQUEST",
		5: "SERVICE_CERTIFICATE",
	}
	SignedLicenseRequestRaw_MessageType_value = map[string]int32{
		"LICENSE_REQUEST":             1,
		"LICENSE":                     2,
		"ERROR_RESPONSE":              3,
		"SERVICE_CERTIFICATE_REQUEST": 4,
		"SERVICE_CERTIFICATE":         5,
	}
)

func (x SignedLicenseRequestRaw_MessageType) Enum() *SignedLicenseRequestRaw_MessageType {
	p := new(SignedLicenseRequestRaw_MessageType)
	*p = x
	return p
}

func (x SignedLicenseRequestRaw_MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SignedLicenseRequestRaw_MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_wv_proto2_proto_enumTypes[16].Descriptor()
}

func (SignedLicenseRequestRaw_MessageType) Type() protoreflect.EnumType {
	return &file_wv_proto2_proto_enumTypes[16]
}

func (x SignedLicenseRequestRaw_MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *SignedLicenseRequestRaw_MessageType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = SignedLicenseRequestRaw_MessageType(num)
	return nil
}

// Deprecated: Use SignedLicenseRequestRaw_MessageType.Descriptor instead.
func (SignedLicenseRequestRaw_MessageType) EnumDescriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{23, 0}
}

type SignedLicense_MessageType int32

const (
	SignedLicense_LICENSE_REQUEST             SignedLicense_MessageType = 1
	SignedLicense_LICENSE                     SignedLicense_MessageType = 2
	SignedLicense_ERROR_RESPONSE              SignedLicense_MessageType = 3
	SignedLicense_SERVICE_CERTIFICATE_REQUEST SignedLicense_MessageType = 4
	SignedLicense_SERVICE_CERTIFICATE         SignedLicense_MessageType = 5
)

// Enum value maps for SignedLicense_MessageType.
var (
	SignedLicense_MessageType_name = map[int32]string{
		1: "LICENSE_REQUEST",
		2: "LICENSE",
		3: "ERROR_RESPONSE",
		4: "SERVICE_CERTIFICATE_REQUEST",
		5: "SERVICE_CERTIFICATE",
	}
	SignedLicense_MessageType_value = map[string]int32{
		"LICENSE_REQUEST":             1,
		"LICENSE":                     2,
		"ERROR_RESPONSE":              3,
		"SERVICE_CERTIFICATE_REQUEST": 4,
		"SERVICE_CERTIFICATE":         5,
	}
)

func (x SignedLicense_MessageType) Enum() *SignedLicense_MessageType {
	p := new(SignedLicense_MessageType)
	*p = x
	return p
}

func (x SignedLicense_MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SignedLicense_MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_wv_proto2_proto_enumTypes[17].Descriptor()
}

func (SignedLicense_MessageType) Type() protoreflect.EnumType {
	return &file_wv_proto2_proto_enumTypes[17]
}

func (x SignedLicense_MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *SignedLicense_MessageType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = SignedLicense_MessageType(num)
	return nil
}

// Deprecated: Use SignedLicense_MessageType.Descriptor instead.
func (SignedLicense_MessageType) EnumDescriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{24, 0}
}

type SignedServiceCertificate_MessageType int32

const (
	SignedServiceCertificate_LICENSE_REQUEST             SignedServiceCertificate_MessageType = 1
	SignedServiceCertificate_LICENSE                     SignedServiceCertificate_MessageType = 2
	SignedServiceCertificate_ERROR_RESPONSE              SignedServiceCertificate_MessageType = 3
	SignedServiceCertificate_SERVICE_CERTIFICATE_REQUEST SignedServiceCertificate_MessageType = 4
	SignedServiceCertificate_SERVICE_CERTIFICATE         SignedServiceCertificate_MessageType = 5
)

// Enum value maps for SignedServiceCertificate_MessageType.
var (
	SignedServiceCertificate_MessageType_name = map[int32]string{
		1: "LICENSE_REQUEST",
		2: "LICENSE",
		3: "ERROR_RESPONSE",
		4: "SERVICE_CERTIFICATE_REQUEST",
		5: "SERVICE_CERTIFICATE",
	}
	SignedServiceCertificate_MessageType_value = map[string]int32{
		"LICENSE_REQUEST":             1,
		"LICENSE":                     2,
		"ERROR_RESPONSE":              3,
		"SERVICE_CERTIFICATE_REQUEST": 4,
		"SERVICE_CERTIFICATE":         5,
	}
)

func (x SignedServiceCertificate_MessageType) Enum() *SignedServiceCertificate_MessageType {
	p := new(SignedServiceCertificate_MessageType)
	*p = x
	return p
}

func (x SignedServiceCertificate_MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SignedServiceCertificate_MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_wv_proto2_proto_enumTypes[18].Descriptor()
}

func (SignedServiceCertificate_MessageType) Type() protoreflect.EnumType {
	return &file_wv_proto2_proto_enumTypes[18]
}

func (x SignedServiceCertificate_MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *SignedServiceCertificate_MessageType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = SignedServiceCertificate_MessageType(num)
	return nil
}

// Deprecated: Use SignedServiceCertificate_MessageType.Descriptor instead.
func (SignedServiceCertificate_MessageType) EnumDescriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{25, 0}
}

// from x86 (partial), most of it from the ARM version:
type ClientIdentification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type *ClientIdentification_TokenType `protobuf:"varint,1,req,name=Type,enum=ClientIdentification_TokenType" json:"Type,omitempty"`
	//optional bytes Token = 2; // by default the client treats this as blob, but it's usually a DeviceCertificate, so for usefulness sake, I'm replacing it with this one:
	Token               *SignedDeviceCertificate                 `protobuf:"bytes,2,opt,name=Token" json:"Token,omitempty"` // use this when parsing, "bytes" when building a client id blob
	ClientInfo          []*ClientIdentification_NameValue        `protobuf:"bytes,3,rep,name=ClientInfo" json:"ClientInfo,omitempty"`
	ProviderClientToken []byte                                   `protobuf:"bytes,4,opt,name=ProviderClientToken" json:"ProviderClientToken,omitempty"`
	LicenseCounter      *uint32                                  `protobuf:"varint,5,opt,name=LicenseCounter" json:"LicenseCounter,omitempty"`
	XClientCapabilities *ClientIdentification_ClientCapabilities `protobuf:"bytes,6,opt,name=_ClientCapabilities,json=ClientCapabilities" json:"_ClientCapabilities,omitempty"` // how should we deal with duped names? will have to look at proto docs later
	XFileHashes         *FileHashes                              `protobuf:"bytes,7,opt,name=_FileHashes,json=FileHashes" json:"_FileHashes,omitempty"`                         // vmp blob goes here
}

func (x *ClientIdentification) Reset() {
	*x = ClientIdentification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientIdentification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientIdentification) ProtoMessage() {}

func (x *ClientIdentification) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientIdentification.ProtoReflect.Descriptor instead.
func (*ClientIdentification) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{0}
}

func (x *ClientIdentification) GetType() ClientIdentification_TokenType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ClientIdentification_KEYBOX
}

func (x *ClientIdentification) GetToken() *SignedDeviceCertificate {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *ClientIdentification) GetClientInfo() []*ClientIdentification_NameValue {
	if x != nil {
		return x.ClientInfo
	}
	return nil
}

func (x *ClientIdentification) GetProviderClientToken() []byte {
	if x != nil {
		return x.ProviderClientToken
	}
	return nil
}

func (x *ClientIdentification) GetLicenseCounter() uint32 {
	if x != nil && x.LicenseCounter != nil {
		return *x.LicenseCounter
	}
	return 0
}

func (x *ClientIdentification) GetXClientCapabilities() *ClientIdentification_ClientCapabilities {
	if x != nil {
		return x.XClientCapabilities
	}
	return nil
}

func (x *ClientIdentification) GetXFileHashes() *FileHashes {
	if x != nil {
		return x.XFileHashes
	}
	return nil
}

type DeviceCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type                 *DeviceCertificate_CertificateType `protobuf:"varint,1,req,name=Type,enum=DeviceCertificate_CertificateType" json:"Type,omitempty"` // the compiled code reused this as ProvisionedDeviceInfo.WvSecurityLevel, however that is incorrect (compiler aliased it as they're both identical as a structure)
	SerialNumber         []byte                             `protobuf:"bytes,2,opt,name=SerialNumber" json:"SerialNumber,omitempty"`
	CreationTimeSeconds  *uint32                            `protobuf:"varint,3,opt,name=CreationTimeSeconds" json:"CreationTimeSeconds,omitempty"`
	PublicKey            []byte                             `protobuf:"bytes,4,opt,name=PublicKey" json:"PublicKey,omitempty"`
	SystemId             *uint32                            `protobuf:"varint,5,opt,name=SystemId" json:"SystemId,omitempty"`
	TestDeviceDeprecated *uint32                            `protobuf:"varint,6,opt,name=TestDeviceDeprecated" json:"TestDeviceDeprecated,omitempty"` // is it bool or int?
	ServiceId            []byte                             `protobuf:"bytes,7,opt,name=ServiceId" json:"ServiceId,omitempty"`                        // service URL for service certificates
}

func (x *DeviceCertificate) Reset() {
	*x = DeviceCertificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceCertificate) ProtoMessage() {}

func (x *DeviceCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceCertificate.ProtoReflect.Descriptor instead.
func (*DeviceCertificate) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceCertificate) GetType() DeviceCertificate_CertificateType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return DeviceCertificate_ROOT
}

func (x *DeviceCertificate) GetSerialNumber() []byte {
	if x != nil {
		return x.SerialNumber
	}
	return nil
}

func (x *DeviceCertificate) GetCreationTimeSeconds() uint32 {
	if x != nil && x.CreationTimeSeconds != nil {
		return *x.CreationTimeSeconds
	}
	return 0
}

func (x *DeviceCertificate) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *DeviceCertificate) GetSystemId() uint32 {
	if x != nil && x.SystemId != nil {
		return *x.SystemId
	}
	return 0
}

func (x *DeviceCertificate) GetTestDeviceDeprecated() uint32 {
	if x != nil && x.TestDeviceDeprecated != nil {
		return *x.TestDeviceDeprecated
	}
	return 0
}

func (x *DeviceCertificate) GetServiceId() []byte {
	if x != nil {
		return x.ServiceId
	}
	return nil
}

// missing some references,
type DeviceCertificateStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerialNumber []byte                                     `protobuf:"bytes,1,opt,name=SerialNumber" json:"SerialNumber,omitempty"`
	Status       *DeviceCertificateStatus_CertificateStatus `protobuf:"varint,2,opt,name=Status,enum=DeviceCertificateStatus_CertificateStatus" json:"Status,omitempty"`
	DeviceInfo   *ProvisionedDeviceInfo                     `protobuf:"bytes,4,opt,name=DeviceInfo" json:"DeviceInfo,omitempty"` // where is 3? is it deprecated?
}

func (x *DeviceCertificateStatus) Reset() {
	*x = DeviceCertificateStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceCertificateStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceCertificateStatus) ProtoMessage() {}

func (x *DeviceCertificateStatus) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceCertificateStatus.ProtoReflect.Descriptor instead.
func (*DeviceCertificateStatus) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{2}
}

func (x *DeviceCertificateStatus) GetSerialNumber() []byte {
	if x != nil {
		return x.SerialNumber
	}
	return nil
}

func (x *DeviceCertificateStatus) GetStatus() DeviceCertificateStatus_CertificateStatus {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return DeviceCertificateStatus_VALID
}

func (x *DeviceCertificateStatus) GetDeviceInfo() *ProvisionedDeviceInfo {
	if x != nil {
		return x.DeviceInfo
	}
	return nil
}

type DeviceCertificateStatusList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreationTimeSeconds *uint32                    `protobuf:"varint,1,opt,name=CreationTimeSeconds" json:"CreationTimeSeconds,omitempty"`
	CertificateStatus   []*DeviceCertificateStatus `protobuf:"bytes,2,rep,name=CertificateStatus" json:"CertificateStatus,omitempty"`
}

func (x *DeviceCertificateStatusList) Reset() {
	*x = DeviceCertificateStatusList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceCertificateStatusList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceCertificateStatusList) ProtoMessage() {}

func (x *DeviceCertificateStatusList) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceCertificateStatusList.ProtoReflect.Descriptor instead.
func (*DeviceCertificateStatusList) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{3}
}

func (x *DeviceCertificateStatusList) GetCreationTimeSeconds() uint32 {
	if x != nil && x.CreationTimeSeconds != nil {
		return *x.CreationTimeSeconds
	}
	return 0
}

func (x *DeviceCertificateStatusList) GetCertificateStatus() []*DeviceCertificateStatus {
	if x != nil {
		return x.CertificateStatus
	}
	return nil
}

type EncryptedClientIdentification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceId                      *string `protobuf:"bytes,1,req,name=ServiceId" json:"ServiceId,omitempty"`
	ServiceCertificateSerialNumber []byte  `protobuf:"bytes,2,opt,name=ServiceCertificateSerialNumber" json:"ServiceCertificateSerialNumber,omitempty"`
	EncryptedClientId              []byte  `protobuf:"bytes,3,req,name=EncryptedClientId" json:"EncryptedClientId,omitempty"`
	EncryptedClientIdIv            []byte  `protobuf:"bytes,4,req,name=EncryptedClientIdIv" json:"EncryptedClientIdIv,omitempty"`
	EncryptedPrivacyKey            []byte  `protobuf:"bytes,5,req,name=EncryptedPrivacyKey" json:"EncryptedPrivacyKey,omitempty"`
}

func (x *EncryptedClientIdentification) Reset() {
	*x = EncryptedClientIdentification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptedClientIdentification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptedClientIdentification) ProtoMessage() {}

func (x *EncryptedClientIdentification) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptedClientIdentification.ProtoReflect.Descriptor instead.
func (*EncryptedClientIdentification) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{4}
}

func (x *EncryptedClientIdentification) GetServiceId() string {
	if x != nil && x.ServiceId != nil {
		return *x.ServiceId
	}
	return ""
}

func (x *EncryptedClientIdentification) GetServiceCertificateSerialNumber() []byte {
	if x != nil {
		return x.ServiceCertificateSerialNumber
	}
	return nil
}

func (x *EncryptedClientIdentification) GetEncryptedClientId() []byte {
	if x != nil {
		return x.EncryptedClientId
	}
	return nil
}

func (x *EncryptedClientIdentification) GetEncryptedClientIdIv() []byte {
	if x != nil {
		return x.EncryptedClientIdIv
	}
	return nil
}

func (x *EncryptedClientIdentification) GetEncryptedPrivacyKey() []byte {
	if x != nil {
		return x.EncryptedPrivacyKey
	}
	return nil
}

type LicenseIdentification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId            []byte       `protobuf:"bytes,1,opt,name=RequestId" json:"RequestId,omitempty"`
	SessionId            []byte       `protobuf:"bytes,2,opt,name=SessionId" json:"SessionId,omitempty"`
	PurchaseId           []byte       `protobuf:"bytes,3,opt,name=PurchaseId" json:"PurchaseId,omitempty"`
	Type                 *LicenseType `protobuf:"varint,4,opt,name=Type,enum=LicenseType" json:"Type,omitempty"`
	Version              *uint32      `protobuf:"varint,5,opt,name=Version" json:"Version,omitempty"`
	ProviderSessionToken []byte       `protobuf:"bytes,6,opt,name=ProviderSessionToken" json:"ProviderSessionToken,omitempty"`
}

func (x *LicenseIdentification) Reset() {
	*x = LicenseIdentification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LicenseIdentification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LicenseIdentification) ProtoMessage() {}

func (x *LicenseIdentification) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LicenseIdentification.ProtoReflect.Descriptor instead.
func (*LicenseIdentification) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{5}
}

func (x *LicenseIdentification) GetRequestId() []byte {
	if x != nil {
		return x.RequestId
	}
	return nil
}

func (x *LicenseIdentification) GetSessionId() []byte {
	if x != nil {
		return x.SessionId
	}
	return nil
}

func (x *LicenseIdentification) GetPurchaseId() []byte {
	if x != nil {
		return x.PurchaseId
	}
	return nil
}

func (x *LicenseIdentification) GetType() LicenseType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return LicenseType_ZERO
}

func (x *LicenseIdentification) GetVersion() uint32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

func (x *LicenseIdentification) GetProviderSessionToken() []byte {
	if x != nil {
		return x.ProviderSessionToken
	}
	return nil
}

type License struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                        *LicenseIdentification  `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	XPolicy                   *License_Policy         `protobuf:"bytes,2,opt,name=_Policy,json=Policy" json:"_Policy,omitempty"` // duped names, etc
	Key                       []*License_KeyContainer `protobuf:"bytes,3,rep,name=Key" json:"Key,omitempty"`
	LicenseStartTime          *uint32                 `protobuf:"varint,4,opt,name=LicenseStartTime" json:"LicenseStartTime,omitempty"`
	RemoteAttestationVerified *uint32                 `protobuf:"varint,5,opt,name=RemoteAttestationVerified" json:"RemoteAttestationVerified,omitempty"` // bool?
	ProviderClientToken       []byte                  `protobuf:"bytes,6,opt,name=ProviderClientToken" json:"ProviderClientToken,omitempty"`
	// there might be more, check with newer versions (I see field 7-8 in a lic)
	// this appeared in latest x86:
	ProtectionScheme *uint32 `protobuf:"varint,7,opt,name=ProtectionScheme" json:"ProtectionScheme,omitempty"` // type unconfirmed fully, but it's likely as WidevineCencHeader describesit (fourcc)
}

func (x *License) Reset() {
	*x = License{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *License) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*License) ProtoMessage() {}

func (x *License) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use License.ProtoReflect.Descriptor instead.
func (*License) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{6}
}

func (x *License) GetId() *LicenseIdentification {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *License) GetXPolicy() *License_Policy {
	if x != nil {
		return x.XPolicy
	}
	return nil
}

func (x *License) GetKey() []*License_KeyContainer {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *License) GetLicenseStartTime() uint32 {
	if x != nil && x.LicenseStartTime != nil {
		return *x.LicenseStartTime
	}
	return 0
}

func (x *License) GetRemoteAttestationVerified() uint32 {
	if x != nil && x.RemoteAttestationVerified != nil {
		return *x.RemoteAttestationVerified
	}
	return 0
}

func (x *License) GetProviderClientToken() []byte {
	if x != nil {
		return x.ProviderClientToken
	}
	return nil
}

func (x *License) GetProtectionScheme() uint32 {
	if x != nil && x.ProtectionScheme != nil {
		return *x.ProtectionScheme
	}
	return 0
}

type LicenseError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//LicenseRequest.RequestType ErrorCode; // clang mismatch
	ErrorCode *LicenseError_Error `protobuf:"varint,1,opt,name=ErrorCode,enum=LicenseError_Error" json:"ErrorCode,omitempty"`
}

func (x *LicenseError) Reset() {
	*x = LicenseError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LicenseError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LicenseError) ProtoMessage() {}

func (x *LicenseError) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LicenseError.ProtoReflect.Descriptor instead.
func (*LicenseError) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{7}
}

func (x *LicenseError) GetErrorCode() LicenseError_Error {
	if x != nil && x.ErrorCode != nil {
		return *x.ErrorCode
	}
	return LicenseError_INVALID_DEVICE_CERTIFICATE
}

type LicenseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId                  *ClientIdentification                 `protobuf:"bytes,1,opt,name=ClientId" json:"ClientId,omitempty"`
	ContentId                 *LicenseRequest_ContentIdentification `protobuf:"bytes,2,opt,name=ContentId" json:"ContentId,omitempty"`
	Type                      *LicenseRequest_RequestType           `protobuf:"varint,3,opt,name=Type,enum=LicenseRequest_RequestType" json:"Type,omitempty"`
	RequestTime               *uint32                               `protobuf:"varint,4,opt,name=RequestTime" json:"RequestTime,omitempty"`
	KeyControlNonceDeprecated []byte                                `protobuf:"bytes,5,opt,name=KeyControlNonceDeprecated" json:"KeyControlNonceDeprecated,omitempty"`
	ProtocolVersion           *ProtocolVersion                      `protobuf:"varint,6,opt,name=ProtocolVersion,enum=ProtocolVersion" json:"ProtocolVersion,omitempty"` // lacking symbols for this
	KeyControlNonce           *uint32                               `protobuf:"varint,7,opt,name=KeyControlNonce" json:"KeyControlNonce,omitempty"`
	EncryptedClientId         *EncryptedClientIdentification        `protobuf:"bytes,8,opt,name=EncryptedClientId" json:"EncryptedClientId,omitempty"`
}

func (x *LicenseRequest) Reset() {
	*x = LicenseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LicenseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LicenseRequest) ProtoMessage() {}

func (x *LicenseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LicenseRequest.ProtoReflect.Descriptor instead.
func (*LicenseRequest) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{8}
}

func (x *LicenseRequest) GetClientId() *ClientIdentification {
	if x != nil {
		return x.ClientId
	}
	return nil
}

func (x *LicenseRequest) GetContentId() *LicenseRequest_ContentIdentification {
	if x != nil {
		return x.ContentId
	}
	return nil
}

func (x *LicenseRequest) GetType() LicenseRequest_RequestType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return LicenseRequest_NEW
}

func (x *LicenseRequest) GetRequestTime() uint32 {
	if x != nil && x.RequestTime != nil {
		return *x.RequestTime
	}
	return 0
}

func (x *LicenseRequest) GetKeyControlNonceDeprecated() []byte {
	if x != nil {
		return x.KeyControlNonceDeprecated
	}
	return nil
}

func (x *LicenseRequest) GetProtocolVersion() ProtocolVersion {
	if x != nil && x.ProtocolVersion != nil {
		return *x.ProtocolVersion
	}
	return ProtocolVersion_CURRENT
}

func (x *LicenseRequest) GetKeyControlNonce() uint32 {
	if x != nil && x.KeyControlNonce != nil {
		return *x.KeyControlNonce
	}
	return 0
}

func (x *LicenseRequest) GetEncryptedClientId() *EncryptedClientIdentification {
	if x != nil {
		return x.EncryptedClientId
	}
	return nil
}

// raw pssh hack
type LicenseRequestRaw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId                  *ClientIdentification                    `protobuf:"bytes,1,opt,name=ClientId" json:"ClientId,omitempty"`
	ContentId                 *LicenseRequestRaw_ContentIdentification `protobuf:"bytes,2,opt,name=ContentId" json:"ContentId,omitempty"`
	Type                      *LicenseRequestRaw_RequestType           `protobuf:"varint,3,opt,name=Type,enum=LicenseRequestRaw_RequestType" json:"Type,omitempty"`
	RequestTime               *uint32                                  `protobuf:"varint,4,opt,name=RequestTime" json:"RequestTime,omitempty"`
	KeyControlNonceDeprecated []byte                                   `protobuf:"bytes,5,opt,name=KeyControlNonceDeprecated" json:"KeyControlNonceDeprecated,omitempty"`
	ProtocolVersion           *ProtocolVersion                         `protobuf:"varint,6,opt,name=ProtocolVersion,enum=ProtocolVersion" json:"ProtocolVersion,omitempty"` // lacking symbols for this
	KeyControlNonce           *uint32                                  `protobuf:"varint,7,opt,name=KeyControlNonce" json:"KeyControlNonce,omitempty"`
	EncryptedClientId         *EncryptedClientIdentification           `protobuf:"bytes,8,opt,name=EncryptedClientId" json:"EncryptedClientId,omitempty"`
}

func (x *LicenseRequestRaw) Reset() {
	*x = LicenseRequestRaw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LicenseRequestRaw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LicenseRequestRaw) ProtoMessage() {}

func (x *LicenseRequestRaw) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LicenseRequestRaw.ProtoReflect.Descriptor instead.
func (*LicenseRequestRaw) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{9}
}

func (x *LicenseRequestRaw) GetClientId() *ClientIdentification {
	if x != nil {
		return x.ClientId
	}
	return nil
}

func (x *LicenseRequestRaw) GetContentId() *LicenseRequestRaw_ContentIdentification {
	if x != nil {
		return x.ContentId
	}
	return nil
}

func (x *LicenseRequestRaw) GetType() LicenseRequestRaw_RequestType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return LicenseRequestRaw_NEW
}

func (x *LicenseRequestRaw) GetRequestTime() uint32 {
	if x != nil && x.RequestTime != nil {
		return *x.RequestTime
	}
	return 0
}

func (x *LicenseRequestRaw) GetKeyControlNonceDeprecated() []byte {
	if x != nil {
		return x.KeyControlNonceDeprecated
	}
	return nil
}

func (x *LicenseRequestRaw) GetProtocolVersion() ProtocolVersion {
	if x != nil && x.ProtocolVersion != nil {
		return *x.ProtocolVersion
	}
	return ProtocolVersion_CURRENT
}

func (x *LicenseRequestRaw) GetKeyControlNonce() uint32 {
	if x != nil && x.KeyControlNonce != nil {
		return *x.KeyControlNonce
	}
	return 0
}

func (x *LicenseRequestRaw) GetEncryptedClientId() *EncryptedClientIdentification {
	if x != nil {
		return x.EncryptedClientId
	}
	return nil
}

type ProvisionedDeviceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SystemId      *uint32                                `protobuf:"varint,1,opt,name=SystemId" json:"SystemId,omitempty"`
	Soc           *string                                `protobuf:"bytes,2,opt,name=Soc" json:"Soc,omitempty"`
	Manufacturer  *string                                `protobuf:"bytes,3,opt,name=Manufacturer" json:"Manufacturer,omitempty"`
	Model         *string                                `protobuf:"bytes,4,opt,name=Model" json:"Model,omitempty"`
	DeviceType    *string                                `protobuf:"bytes,5,opt,name=DeviceType" json:"DeviceType,omitempty"`
	ModelYear     *uint32                                `protobuf:"varint,6,opt,name=ModelYear" json:"ModelYear,omitempty"`
	SecurityLevel *ProvisionedDeviceInfo_WvSecurityLevel `protobuf:"varint,7,opt,name=SecurityLevel,enum=ProvisionedDeviceInfo_WvSecurityLevel" json:"SecurityLevel,omitempty"`
	TestDevice    *uint32                                `protobuf:"varint,8,opt,name=TestDevice" json:"TestDevice,omitempty"` // bool?
}

func (x *ProvisionedDeviceInfo) Reset() {
	*x = ProvisionedDeviceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProvisionedDeviceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvisionedDeviceInfo) ProtoMessage() {}

func (x *ProvisionedDeviceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvisionedDeviceInfo.ProtoReflect.Descriptor instead.
func (*ProvisionedDeviceInfo) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{10}
}

func (x *ProvisionedDeviceInfo) GetSystemId() uint32 {
	if x != nil && x.SystemId != nil {
		return *x.SystemId
	}
	return 0
}

func (x *ProvisionedDeviceInfo) GetSoc() string {
	if x != nil && x.Soc != nil {
		return *x.Soc
	}
	return ""
}

func (x *ProvisionedDeviceInfo) GetManufacturer() string {
	if x != nil && x.Manufacturer != nil {
		return *x.Manufacturer
	}
	return ""
}

func (x *ProvisionedDeviceInfo) GetModel() string {
	if x != nil && x.Model != nil {
		return *x.Model
	}
	return ""
}

func (x *ProvisionedDeviceInfo) GetDeviceType() string {
	if x != nil && x.DeviceType != nil {
		return *x.DeviceType
	}
	return ""
}

func (x *ProvisionedDeviceInfo) GetModelYear() uint32 {
	if x != nil && x.ModelYear != nil {
		return *x.ModelYear
	}
	return 0
}

func (x *ProvisionedDeviceInfo) GetSecurityLevel() ProvisionedDeviceInfo_WvSecurityLevel {
	if x != nil && x.SecurityLevel != nil {
		return *x.SecurityLevel
	}
	return ProvisionedDeviceInfo_LEVEL_UNSPECIFIED
}

func (x *ProvisionedDeviceInfo) GetTestDevice() uint32 {
	if x != nil && x.TestDevice != nil {
		return *x.TestDevice
	}
	return 0
}

// todo: fill
type ProvisioningOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProvisioningOptions) Reset() {
	*x = ProvisioningOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProvisioningOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvisioningOptions) ProtoMessage() {}

func (x *ProvisioningOptions) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvisioningOptions.ProtoReflect.Descriptor instead.
func (*ProvisioningOptions) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{11}
}

// todo: fill
type ProvisioningRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProvisioningRequest) Reset() {
	*x = ProvisioningRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProvisioningRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvisioningRequest) ProtoMessage() {}

func (x *ProvisioningRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvisioningRequest.ProtoReflect.Descriptor instead.
func (*ProvisioningRequest) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{12}
}

// todo: fill
type ProvisioningResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProvisioningResponse) Reset() {
	*x = ProvisioningResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProvisioningResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvisioningResponse) ProtoMessage() {}

func (x *ProvisioningResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvisioningResponse.ProtoReflect.Descriptor instead.
func (*ProvisioningResponse) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{13}
}

type RemoteAttestation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Certificate *EncryptedClientIdentification `protobuf:"bytes,1,opt,name=Certificate" json:"Certificate,omitempty"`
	Salt        *string                        `protobuf:"bytes,2,opt,name=Salt" json:"Salt,omitempty"`
	Signature   *string                        `protobuf:"bytes,3,opt,name=Signature" json:"Signature,omitempty"`
}

func (x *RemoteAttestation) Reset() {
	*x = RemoteAttestation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteAttestation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteAttestation) ProtoMessage() {}

func (x *RemoteAttestation) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteAttestation.ProtoReflect.Descriptor instead.
func (*RemoteAttestation) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{14}
}

func (x *RemoteAttestation) GetCertificate() *EncryptedClientIdentification {
	if x != nil {
		return x.Certificate
	}
	return nil
}

func (x *RemoteAttestation) GetSalt() string {
	if x != nil && x.Salt != nil {
		return *x.Salt
	}
	return ""
}

func (x *RemoteAttestation) GetSignature() string {
	if x != nil && x.Signature != nil {
		return *x.Signature
	}
	return ""
}

// todo: fill
type SessionInit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SessionInit) Reset() {
	*x = SessionInit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionInit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionInit) ProtoMessage() {}

func (x *SessionInit) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionInit.ProtoReflect.Descriptor instead.
func (*SessionInit) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{15}
}

// todo: fill
type SessionState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SessionState) Reset() {
	*x = SessionState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionState) ProtoMessage() {}

func (x *SessionState) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionState.ProtoReflect.Descriptor instead.
func (*SessionState) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{16}
}

// todo: fill
type SignedCertificateStatusList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignedCertificateStatusList) Reset() {
	*x = SignedCertificateStatusList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedCertificateStatusList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedCertificateStatusList) ProtoMessage() {}

func (x *SignedCertificateStatusList) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedCertificateStatusList.ProtoReflect.Descriptor instead.
func (*SignedCertificateStatusList) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{17}
}

type SignedDeviceCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//optional bytes DeviceCertificate = 1; // again, they use a buffer where it's supposed to be a message, so we'll replace it with what it really is:
	XDeviceCertificate *DeviceCertificate       `protobuf:"bytes,1,opt,name=_DeviceCertificate,json=DeviceCertificate" json:"_DeviceCertificate,omitempty"` // how should we deal with duped names? will have to look at proto docs later
	Signature          []byte                   `protobuf:"bytes,2,opt,name=Signature" json:"Signature,omitempty"`
	Signer             *SignedDeviceCertificate `protobuf:"bytes,3,opt,name=Signer" json:"Signer,omitempty"`
}

func (x *SignedDeviceCertificate) Reset() {
	*x = SignedDeviceCertificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedDeviceCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedDeviceCertificate) ProtoMessage() {}

func (x *SignedDeviceCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedDeviceCertificate.ProtoReflect.Descriptor instead.
func (*SignedDeviceCertificate) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{18}
}

func (x *SignedDeviceCertificate) GetXDeviceCertificate() *DeviceCertificate {
	if x != nil {
		return x.XDeviceCertificate
	}
	return nil
}

func (x *SignedDeviceCertificate) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *SignedDeviceCertificate) GetSigner() *SignedDeviceCertificate {
	if x != nil {
		return x.Signer
	}
	return nil
}

// todo: fill
type SignedProvisioningMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignedProvisioningMessage) Reset() {
	*x = SignedProvisioningMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[19]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedProvisioningMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedProvisioningMessage) ProtoMessage() {}

func (x *SignedProvisioningMessage) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[19]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedProvisioningMessage.ProtoReflect.Descriptor instead.
func (*SignedProvisioningMessage) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{19}
}

// the root of all messages, from either server or client
type SignedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type *SignedMessage_MessageType `protobuf:"varint,1,opt,name=Type,enum=SignedMessage_MessageType" json:"Type,omitempty"` // has in incorrect overlap with License_KeyContainer_SecurityLevel
	Msg  []byte                     `protobuf:"bytes,2,opt,name=Msg" json:"Msg,omitempty"`                                   // this has to be casted dynamically, to LicenseRequest, License or LicenseError (? unconfirmed), for Request, no other fields but Type need to be present
	// for SERVICE_CERTIFICATE, only Type and Msg are present, and it's just a DeviceCertificate with CertificateType set to SERVICE
	Signature         []byte             `protobuf:"bytes,3,opt,name=Signature" json:"Signature,omitempty"`   // might be different type of signatures (ex. RSA vs AES CMAC(??), unconfirmed for now)
	SessionKey        []byte             `protobuf:"bytes,4,opt,name=SessionKey" json:"SessionKey,omitempty"` // often RSA wrapped for licenses
	RemoteAttestation *RemoteAttestation `protobuf:"bytes,5,opt,name=RemoteAttestation" json:"RemoteAttestation,omitempty"`
}

func (x *SignedMessage) Reset() {
	*x = SignedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[20]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedMessage) ProtoMessage() {}

func (x *SignedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[20]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedMessage.ProtoReflect.Descriptor instead.
func (*SignedMessage) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{20}
}

func (x *SignedMessage) GetType() SignedMessage_MessageType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return SignedMessage_LICENSE_REQUEST
}

func (x *SignedMessage) GetMsg() []byte {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *SignedMessage) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *SignedMessage) GetSessionKey() []byte {
	if x != nil {
		return x.SessionKey
	}
	return nil
}

func (x *SignedMessage) GetRemoteAttestation() *RemoteAttestation {
	if x != nil {
		return x.RemoteAttestation
	}
	return nil
}

// This message is copied from google's docs, not reversed:
type WidevineCencHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Algorithm *WidevineCencHeader_Algorithm `protobuf:"varint,1,opt,name=algorithm,enum=WidevineCencHeader_Algorithm" json:"algorithm,omitempty"`
	KeyId     [][]byte                      `protobuf:"bytes,2,rep,name=key_id,json=keyId" json:"key_id,omitempty"`
	// Content provider name.
	Provider *string `protobuf:"bytes,3,opt,name=provider" json:"provider,omitempty"`
	// A content identifier, specified by content provider.
	ContentId []byte `protobuf:"bytes,4,opt,name=content_id,json=contentId" json:"content_id,omitempty"`
	// Track type. Acceptable values are SD, HD and AUDIO. Used to
	// differentiate content keys used by an asset.
	TrackTypeDeprecated *string `protobuf:"bytes,5,opt,name=track_type_deprecated,json=trackTypeDeprecated" json:"track_type_deprecated,omitempty"`
	// The name of a registered policy to be used for this asset.
	Policy *string `protobuf:"bytes,6,opt,name=policy" json:"policy,omitempty"`
	// Crypto period index, for media using key rotation.
	CryptoPeriodIndex *uint32 `protobuf:"varint,7,opt,name=crypto_period_index,json=cryptoPeriodIndex" json:"crypto_period_index,omitempty"`
	// Optional protected context for group content. The grouped_license is a
	// serialized SignedMessage.
	GroupedLicense []byte `protobuf:"bytes,8,opt,name=grouped_license,json=groupedLicense" json:"grouped_license,omitempty"`
	// Protection scheme identifying the encryption algorithm.
	// Represented as one of the following 4CC values:
	// 'cenc' (AESCTR), 'cbc1' (AESCBC),
	// 'cens' (AESCTR subsample), 'cbcs' (AESCBC subsample).
	ProtectionScheme *uint32 `protobuf:"varint,9,opt,name=protection_scheme,json=protectionScheme" json:"protection_scheme,omitempty"`
	// Optional. For media using key rotation, this represents the duration
	// of each crypto period in seconds.
	CryptoPeriodSeconds *uint32 `protobuf:"varint,10,opt,name=crypto_period_seconds,json=cryptoPeriodSeconds" json:"crypto_period_seconds,omitempty"`
}

func (x *WidevineCencHeader) Reset() {
	*x = WidevineCencHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[21]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WidevineCencHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WidevineCencHeader) ProtoMessage() {}

func (x *WidevineCencHeader) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[21]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WidevineCencHeader.ProtoReflect.Descriptor instead.
func (*WidevineCencHeader) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{21}
}

func (x *WidevineCencHeader) GetAlgorithm() WidevineCencHeader_Algorithm {
	if x != nil && x.Algorithm != nil {
		return *x.Algorithm
	}
	return WidevineCencHeader_UNENCRYPTED
}

func (x *WidevineCencHeader) GetKeyId() [][]byte {
	if x != nil {
		return x.KeyId
	}
	return nil
}

func (x *WidevineCencHeader) GetProvider() string {
	if x != nil && x.Provider != nil {
		return *x.Provider
	}
	return ""
}

func (x *WidevineCencHeader) GetContentId() []byte {
	if x != nil {
		return x.ContentId
	}
	return nil
}

func (x *WidevineCencHeader) GetTrackTypeDeprecated() string {
	if x != nil && x.TrackTypeDeprecated != nil {
		return *x.TrackTypeDeprecated
	}
	return ""
}

func (x *WidevineCencHeader) GetPolicy() string {
	if x != nil && x.Policy != nil {
		return *x.Policy
	}
	return ""
}

func (x *WidevineCencHeader) GetCryptoPeriodIndex() uint32 {
	if x != nil && x.CryptoPeriodIndex != nil {
		return *x.CryptoPeriodIndex
	}
	return 0
}

func (x *WidevineCencHeader) GetGroupedLicense() []byte {
	if x != nil {
		return x.GroupedLicense
	}
	return nil
}

func (x *WidevineCencHeader) GetProtectionScheme() uint32 {
	if x != nil && x.ProtectionScheme != nil {
		return *x.ProtectionScheme
	}
	return 0
}

func (x *WidevineCencHeader) GetCryptoPeriodSeconds() uint32 {
	if x != nil && x.CryptoPeriodSeconds != nil {
		return *x.CryptoPeriodSeconds
	}
	return 0
}

// from here on, it's just for testing, these messages don't exist in the binaries, I'm adding them to avoid detecting type programmatically
type SignedLicenseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type *SignedLicenseRequest_MessageType `protobuf:"varint,1,opt,name=Type,enum=SignedLicenseRequest_MessageType" json:"Type,omitempty"` // has in incorrect overlap with License_KeyContainer_SecurityLevel
	Msg  *LicenseRequest                   `protobuf:"bytes,2,opt,name=Msg" json:"Msg,omitempty"`                                          // this has to be casted dynamically, to LicenseRequest, License or LicenseError (? unconfirmed), for Request, no other fields but Type need to be present
	// for SERVICE_CERTIFICATE, only Type and Msg are present, and it's just a DeviceCertificate with CertificateType set to SERVICE
	Signature         []byte             `protobuf:"bytes,3,opt,name=Signature" json:"Signature,omitempty"`   // might be different type of signatures (ex. RSA vs AES CMAC(??), unconfirmed for now)
	SessionKey        []byte             `protobuf:"bytes,4,opt,name=SessionKey" json:"SessionKey,omitempty"` // often RSA wrapped for licenses
	RemoteAttestation *RemoteAttestation `protobuf:"bytes,5,opt,name=RemoteAttestation" json:"RemoteAttestation,omitempty"`
}

func (x *SignedLicenseRequest) Reset() {
	*x = SignedLicenseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[22]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedLicenseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedLicenseRequest) ProtoMessage() {}

func (x *SignedLicenseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[22]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedLicenseRequest.ProtoReflect.Descriptor instead.
func (*SignedLicenseRequest) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{22}
}

func (x *SignedLicenseRequest) GetType() SignedLicenseRequest_MessageType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return SignedLicenseRequest_LICENSE_REQUEST
}

func (x *SignedLicenseRequest) GetMsg() *LicenseRequest {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *SignedLicenseRequest) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *SignedLicenseRequest) GetSessionKey() []byte {
	if x != nil {
		return x.SessionKey
	}
	return nil
}

func (x *SignedLicenseRequest) GetRemoteAttestation() *RemoteAttestation {
	if x != nil {
		return x.RemoteAttestation
	}
	return nil
}

// hack
type SignedLicenseRequestRaw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type *SignedLicenseRequestRaw_MessageType `protobuf:"varint,1,opt,name=Type,enum=SignedLicenseRequestRaw_MessageType" json:"Type,omitempty"` // has in incorrect overlap with License_KeyContainer_SecurityLevel
	Msg  *LicenseRequestRaw                   `protobuf:"bytes,2,opt,name=Msg" json:"Msg,omitempty"`                                             // this has to be casted dynamically, to LicenseRequest, License or LicenseError (? unconfirmed), for Request, no other fields but Type need to be present
	// for SERVICE_CERTIFICATE, only Type and Msg are present, and it's just a DeviceCertificate with CertificateType set to SERVICE
	Signature         []byte             `protobuf:"bytes,3,opt,name=Signature" json:"Signature,omitempty"`   // might be different type of signatures (ex. RSA vs AES CMAC(??), unconfirmed for now)
	SessionKey        []byte             `protobuf:"bytes,4,opt,name=SessionKey" json:"SessionKey,omitempty"` // often RSA wrapped for licenses
	RemoteAttestation *RemoteAttestation `protobuf:"bytes,5,opt,name=RemoteAttestation" json:"RemoteAttestation,omitempty"`
}

func (x *SignedLicenseRequestRaw) Reset() {
	*x = SignedLicenseRequestRaw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[23]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedLicenseRequestRaw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedLicenseRequestRaw) ProtoMessage() {}

func (x *SignedLicenseRequestRaw) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[23]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedLicenseRequestRaw.ProtoReflect.Descriptor instead.
func (*SignedLicenseRequestRaw) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{23}
}

func (x *SignedLicenseRequestRaw) GetType() SignedLicenseRequestRaw_MessageType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return SignedLicenseRequestRaw_LICENSE_REQUEST
}

func (x *SignedLicenseRequestRaw) GetMsg() *LicenseRequestRaw {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *SignedLicenseRequestRaw) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *SignedLicenseRequestRaw) GetSessionKey() []byte {
	if x != nil {
		return x.SessionKey
	}
	return nil
}

func (x *SignedLicenseRequestRaw) GetRemoteAttestation() *RemoteAttestation {
	if x != nil {
		return x.RemoteAttestation
	}
	return nil
}

type SignedLicense struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type *SignedLicense_MessageType `protobuf:"varint,1,opt,name=Type,enum=SignedLicense_MessageType" json:"Type,omitempty"` // has in incorrect overlap with License_KeyContainer_SecurityLevel
	Msg  *License                   `protobuf:"bytes,2,opt,name=Msg" json:"Msg,omitempty"`                                   // this has to be casted dynamically, to LicenseRequest, License or LicenseError (? unconfirmed), for Request, no other fields but Type need to be present
	// for SERVICE_CERTIFICATE, only Type and Msg are present, and it's just a DeviceCertificate with CertificateType set to SERVICE
	Signature         []byte             `protobuf:"bytes,3,opt,name=Signature" json:"Signature,omitempty"`   // might be different type of signatures (ex. RSA vs AES CMAC(??), unconfirmed for now)
	SessionKey        []byte             `protobuf:"bytes,4,opt,name=SessionKey" json:"SessionKey,omitempty"` // often RSA wrapped for licenses
	RemoteAttestation *RemoteAttestation `protobuf:"bytes,5,opt,name=RemoteAttestation" json:"RemoteAttestation,omitempty"`
}

func (x *SignedLicense) Reset() {
	*x = SignedLicense{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[24]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedLicense) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedLicense) ProtoMessage() {}

func (x *SignedLicense) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[24]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedLicense.ProtoReflect.Descriptor instead.
func (*SignedLicense) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{24}
}

func (x *SignedLicense) GetType() SignedLicense_MessageType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return SignedLicense_LICENSE_REQUEST
}

func (x *SignedLicense) GetMsg() *License {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *SignedLicense) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *SignedLicense) GetSessionKey() []byte {
	if x != nil {
		return x.SessionKey
	}
	return nil
}

func (x *SignedLicense) GetRemoteAttestation() *RemoteAttestation {
	if x != nil {
		return x.RemoteAttestation
	}
	return nil
}

type SignedServiceCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type *SignedServiceCertificate_MessageType `protobuf:"varint,1,opt,name=Type,enum=SignedServiceCertificate_MessageType" json:"Type,omitempty"` // has in incorrect overlap with License_KeyContainer_SecurityLevel
	Msg  *SignedDeviceCertificate              `protobuf:"bytes,2,opt,name=Msg" json:"Msg,omitempty"`                                              // this has to be casted dynamically, to LicenseRequest, License or LicenseError (? unconfirmed), for Request, no other fields but Type need to be present
	// for SERVICE_CERTIFICATE, only Type and Msg are present, and it's just a DeviceCertificate with CertificateType set to SERVICE
	Signature         []byte             `protobuf:"bytes,3,opt,name=Signature" json:"Signature,omitempty"`   // might be different type of signatures (ex. RSA vs AES CMAC(??), unconfirmed for now)
	SessionKey        []byte             `protobuf:"bytes,4,opt,name=SessionKey" json:"SessionKey,omitempty"` // often RSA wrapped for licenses
	RemoteAttestation *RemoteAttestation `protobuf:"bytes,5,opt,name=RemoteAttestation" json:"RemoteAttestation,omitempty"`
}

func (x *SignedServiceCertificate) Reset() {
	*x = SignedServiceCertificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[25]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedServiceCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedServiceCertificate) ProtoMessage() {}

func (x *SignedServiceCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[25]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedServiceCertificate.ProtoReflect.Descriptor instead.
func (*SignedServiceCertificate) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{25}
}

func (x *SignedServiceCertificate) GetType() SignedServiceCertificate_MessageType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return SignedServiceCertificate_LICENSE_REQUEST
}

func (x *SignedServiceCertificate) GetMsg() *SignedDeviceCertificate {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *SignedServiceCertificate) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *SignedServiceCertificate) GetSessionKey() []byte {
	if x != nil {
		return x.SessionKey
	}
	return nil
}

func (x *SignedServiceCertificate) GetRemoteAttestation() *RemoteAttestation {
	if x != nil {
		return x.RemoteAttestation
	}
	return nil
}

//vmp support
type FileHashes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signer     []byte                  `protobuf:"bytes,1,opt,name=signer" json:"signer,omitempty"`
	Signatures []*FileHashes_Signature `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty"`
}

func (x *FileHashes) Reset() {
	*x = FileHashes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[26]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileHashes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileHashes) ProtoMessage() {}

func (x *FileHashes) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[26]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileHashes.ProtoReflect.Descriptor instead.
func (*FileHashes) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{26}
}

func (x *FileHashes) GetSigner() []byte {
	if x != nil {
		return x.Signer
	}
	return nil
}

func (x *FileHashes) GetSignatures() []*FileHashes_Signature {
	if x != nil {
		return x.Signatures
	}
	return nil
}

type ClientIdentification_NameValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  *string `protobuf:"bytes,1,req,name=Name" json:"Name,omitempty"`
	Value *string `protobuf:"bytes,2,req,name=Value" json:"Value,omitempty"`
}

func (x *ClientIdentification_NameValue) Reset() {
	*x = ClientIdentification_NameValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[27]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientIdentification_NameValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientIdentification_NameValue) ProtoMessage() {}

func (x *ClientIdentification_NameValue) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[27]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientIdentification_NameValue.ProtoReflect.Descriptor instead.
func (*ClientIdentification_NameValue) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ClientIdentification_NameValue) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ClientIdentification_NameValue) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

type ClientIdentification_ClientCapabilities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientToken                *uint32                                              `protobuf:"varint,1,opt,name=ClientToken" json:"ClientToken,omitempty"`
	SessionToken               *uint32                                              `protobuf:"varint,2,opt,name=SessionToken" json:"SessionToken,omitempty"`
	VideoResolutionConstraints *uint32                                              `protobuf:"varint,3,opt,name=VideoResolutionConstraints" json:"VideoResolutionConstraints,omitempty"`
	MaxHdcpVersion             *ClientIdentification_ClientCapabilities_HdcpVersion `protobuf:"varint,4,opt,name=MaxHdcpVersion,enum=ClientIdentification_ClientCapabilities_HdcpVersion" json:"MaxHdcpVersion,omitempty"`
	OemCryptoApiVersion        *uint32                                              `protobuf:"varint,5,opt,name=OemCryptoApiVersion" json:"OemCryptoApiVersion,omitempty"`
}

func (x *ClientIdentification_ClientCapabilities) Reset() {
	*x = ClientIdentification_ClientCapabilities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[28]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientIdentification_ClientCapabilities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientIdentification_ClientCapabilities) ProtoMessage() {}

func (x *ClientIdentification_ClientCapabilities) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[28]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientIdentification_ClientCapabilities.ProtoReflect.Descriptor instead.
func (*ClientIdentification_ClientCapabilities) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ClientIdentification_ClientCapabilities) GetClientToken() uint32 {
	if x != nil && x.ClientToken != nil {
		return *x.ClientToken
	}
	return 0
}

func (x *ClientIdentification_ClientCapabilities) GetSessionToken() uint32 {
	if x != nil && x.SessionToken != nil {
		return *x.SessionToken
	}
	return 0
}

func (x *ClientIdentification_ClientCapabilities) GetVideoResolutionConstraints() uint32 {
	if x != nil && x.VideoResolutionConstraints != nil {
		return *x.VideoResolutionConstraints
	}
	return 0
}

func (x *ClientIdentification_ClientCapabilities) GetMaxHdcpVersion() ClientIdentification_ClientCapabilities_HdcpVersion {
	if x != nil && x.MaxHdcpVersion != nil {
		return *x.MaxHdcpVersion
	}
	return ClientIdentification_ClientCapabilities_HDCP_NONE
}

func (x *ClientIdentification_ClientCapabilities) GetOemCryptoApiVersion() uint32 {
	if x != nil && x.OemCryptoApiVersion != nil {
		return *x.OemCryptoApiVersion
	}
	return 0
}

type License_Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CanPlay                        *bool   `protobuf:"varint,1,opt,name=CanPlay" json:"CanPlay,omitempty"` // changed from uint32 to bool
	CanPersist                     *bool   `protobuf:"varint,2,opt,name=CanPersist" json:"CanPersist,omitempty"`
	CanRenew                       *bool   `protobuf:"varint,3,opt,name=CanRenew" json:"CanRenew,omitempty"`
	RentalDurationSeconds          *uint32 `protobuf:"varint,4,opt,name=RentalDurationSeconds" json:"RentalDurationSeconds,omitempty"`
	PlaybackDurationSeconds        *uint32 `protobuf:"varint,5,opt,name=PlaybackDurationSeconds" json:"PlaybackDurationSeconds,omitempty"`
	LicenseDurationSeconds         *uint32 `protobuf:"varint,6,opt,name=LicenseDurationSeconds" json:"LicenseDurationSeconds,omitempty"`
	RenewalRecoveryDurationSeconds *uint32 `protobuf:"varint,7,opt,name=RenewalRecoveryDurationSeconds" json:"RenewalRecoveryDurationSeconds,omitempty"`
	RenewalServerUrl               *string `protobuf:"bytes,8,opt,name=RenewalServerUrl" json:"RenewalServerUrl,omitempty"`
	RenewalDelaySeconds            *uint32 `protobuf:"varint,9,opt,name=RenewalDelaySeconds" json:"RenewalDelaySeconds,omitempty"`
	RenewalRetryIntervalSeconds    *uint32 `protobuf:"varint,10,opt,name=RenewalRetryIntervalSeconds" json:"RenewalRetryIntervalSeconds,omitempty"`
	RenewWithUsage                 *bool   `protobuf:"varint,11,opt,name=RenewWithUsage" json:"RenewWithUsage,omitempty"` // was uint32
}

func (x *License_Policy) Reset() {
	*x = License_Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[29]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *License_Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*License_Policy) ProtoMessage() {}

func (x *License_Policy) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[29]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use License_Policy.ProtoReflect.Descriptor instead.
func (*License_Policy) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{6, 0}
}

func (x *License_Policy) GetCanPlay() bool {
	if x != nil && x.CanPlay != nil {
		return *x.CanPlay
	}
	return false
}

func (x *License_Policy) GetCanPersist() bool {
	if x != nil && x.CanPersist != nil {
		return *x.CanPersist
	}
	return false
}

func (x *License_Policy) GetCanRenew() bool {
	if x != nil && x.CanRenew != nil {
		return *x.CanRenew
	}
	return false
}

func (x *License_Policy) GetRentalDurationSeconds() uint32 {
	if x != nil && x.RentalDurationSeconds != nil {
		return *x.RentalDurationSeconds
	}
	return 0
}

func (x *License_Policy) GetPlaybackDurationSeconds() uint32 {
	if x != nil && x.PlaybackDurationSeconds != nil {
		return *x.PlaybackDurationSeconds
	}
	return 0
}

func (x *License_Policy) GetLicenseDurationSeconds() uint32 {
	if x != nil && x.LicenseDurationSeconds != nil {
		return *x.LicenseDurationSeconds
	}
	return 0
}

func (x *License_Policy) GetRenewalRecoveryDurationSeconds() uint32 {
	if x != nil && x.RenewalRecoveryDurationSeconds != nil {
		return *x.RenewalRecoveryDurationSeconds
	}
	return 0
}

func (x *License_Policy) GetRenewalServerUrl() string {
	if x != nil && x.RenewalServerUrl != nil {
		return *x.RenewalServerUrl
	}
	return ""
}

func (x *License_Policy) GetRenewalDelaySeconds() uint32 {
	if x != nil && x.RenewalDelaySeconds != nil {
		return *x.RenewalDelaySeconds
	}
	return 0
}

func (x *License_Policy) GetRenewalRetryIntervalSeconds() uint32 {
	if x != nil && x.RenewalRetryIntervalSeconds != nil {
		return *x.RenewalRetryIntervalSeconds
	}
	return 0
}

func (x *License_Policy) GetRenewWithUsage() bool {
	if x != nil && x.RenewWithUsage != nil {
		return *x.RenewWithUsage
	}
	return false
}

type License_KeyContainer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                             []byte                                              `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	Iv                             []byte                                              `protobuf:"bytes,2,opt,name=Iv" json:"Iv,omitempty"`
	Key                            []byte                                              `protobuf:"bytes,3,opt,name=Key" json:"Key,omitempty"`
	Type                           *License_KeyContainer_KeyType                       `protobuf:"varint,4,opt,name=Type,enum=License_KeyContainer_KeyType" json:"Type,omitempty"`
	Level                          *License_KeyContainer_SecurityLevel                 `protobuf:"varint,5,opt,name=Level,enum=License_KeyContainer_SecurityLevel" json:"Level,omitempty"`
	RequiredProtection             *License_KeyContainer_OutputProtection              `protobuf:"bytes,6,opt,name=RequiredProtection" json:"RequiredProtection,omitempty"`
	RequestedProtection            *License_KeyContainer_OutputProtection              `protobuf:"bytes,7,opt,name=RequestedProtection" json:"RequestedProtection,omitempty"`
	XKeyControl                    *License_KeyContainer_KeyControl                    `protobuf:"bytes,8,opt,name=_KeyControl,json=KeyControl" json:"_KeyControl,omitempty"`                                                          // duped names, etc
	XOperatorSessionKeyPermissions *License_KeyContainer_OperatorSessionKeyPermissions `protobuf:"bytes,9,opt,name=_OperatorSessionKeyPermissions,json=OperatorSessionKeyPermissions" json:"_OperatorSessionKeyPermissions,omitempty"` // duped names, etc
	VideoResolutionConstraints     []*License_KeyContainer_VideoResolutionConstraint   `protobuf:"bytes,10,rep,name=VideoResolutionConstraints" json:"VideoResolutionConstraints,omitempty"`
}

func (x *License_KeyContainer) Reset() {
	*x = License_KeyContainer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[30]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *License_KeyContainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*License_KeyContainer) ProtoMessage() {}

func (x *License_KeyContainer) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[30]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use License_KeyContainer.ProtoReflect.Descriptor instead.
func (*License_KeyContainer) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{6, 1}
}

func (x *License_KeyContainer) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *License_KeyContainer) GetIv() []byte {
	if x != nil {
		return x.Iv
	}
	return nil
}

func (x *License_KeyContainer) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *License_KeyContainer) GetType() License_KeyContainer_KeyType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return License_KeyContainer_SIGNING
}

func (x *License_KeyContainer) GetLevel() License_KeyContainer_SecurityLevel {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return License_KeyContainer_SW_SECURE_CRYPTO
}

func (x *License_KeyContainer) GetRequiredProtection() *License_KeyContainer_OutputProtection {
	if x != nil {
		return x.RequiredProtection
	}
	return nil
}

func (x *License_KeyContainer) GetRequestedProtection() *License_KeyContainer_OutputProtection {
	if x != nil {
		return x.RequestedProtection
	}
	return nil
}

func (x *License_KeyContainer) GetXKeyControl() *License_KeyContainer_KeyControl {
	if x != nil {
		return x.XKeyControl
	}
	return nil
}

func (x *License_KeyContainer) GetXOperatorSessionKeyPermissions() *License_KeyContainer_OperatorSessionKeyPermissions {
	if x != nil {
		return x.XOperatorSessionKeyPermissions
	}
	return nil
}

func (x *License_KeyContainer) GetVideoResolutionConstraints() []*License_KeyContainer_VideoResolutionConstraint {
	if x != nil {
		return x.VideoResolutionConstraints
	}
	return nil
}

type License_KeyContainer_OutputProtection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hdcp      *ClientIdentification_ClientCapabilities_HdcpVersion `protobuf:"varint,1,opt,name=Hdcp,enum=ClientIdentification_ClientCapabilities_HdcpVersion" json:"Hdcp,omitempty"` // it's most likely a copy of Hdcp version available here, but compiler optimized it away
	CgmsFlags *License_KeyContainer_OutputProtection_CGMS          `protobuf:"varint,2,opt,name=CgmsFlags,enum=License_KeyContainer_OutputProtection_CGMS" json:"CgmsFlags,omitempty"`
}

func (x *License_KeyContainer_OutputProtection) Reset() {
	*x = License_KeyContainer_OutputProtection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[31]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *License_KeyContainer_OutputProtection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*License_KeyContainer_OutputProtection) ProtoMessage() {}

func (x *License_KeyContainer_OutputProtection) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[31]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use License_KeyContainer_OutputProtection.ProtoReflect.Descriptor instead.
func (*License_KeyContainer_OutputProtection) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{6, 1, 0}
}

func (x *License_KeyContainer_OutputProtection) GetHdcp() ClientIdentification_ClientCapabilities_HdcpVersion {
	if x != nil && x.Hdcp != nil {
		return *x.Hdcp
	}
	return ClientIdentification_ClientCapabilities_HDCP_NONE
}

func (x *License_KeyContainer_OutputProtection) GetCgmsFlags() License_KeyContainer_OutputProtection_CGMS {
	if x != nil && x.CgmsFlags != nil {
		return *x.CgmsFlags
	}
	return License_KeyContainer_OutputProtection_COPY_FREE
}

type License_KeyContainer_KeyControl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyControlBlock []byte `protobuf:"bytes,1,req,name=KeyControlBlock" json:"KeyControlBlock,omitempty"` // what is this?
	Iv              []byte `protobuf:"bytes,2,req,name=Iv" json:"Iv,omitempty"`
}

func (x *License_KeyContainer_KeyControl) Reset() {
	*x = License_KeyContainer_KeyControl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[32]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *License_KeyContainer_KeyControl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*License_KeyContainer_KeyControl) ProtoMessage() {}

func (x *License_KeyContainer_KeyControl) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[32]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use License_KeyContainer_KeyControl.ProtoReflect.Descriptor instead.
func (*License_KeyContainer_KeyControl) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{6, 1, 1}
}

func (x *License_KeyContainer_KeyControl) GetKeyControlBlock() []byte {
	if x != nil {
		return x.KeyControlBlock
	}
	return nil
}

func (x *License_KeyContainer_KeyControl) GetIv() []byte {
	if x != nil {
		return x.Iv
	}
	return nil
}

type License_KeyContainer_OperatorSessionKeyPermissions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllowEncrypt         *uint32 `protobuf:"varint,1,opt,name=AllowEncrypt" json:"AllowEncrypt,omitempty"`
	AllowDecrypt         *uint32 `protobuf:"varint,2,opt,name=AllowDecrypt" json:"AllowDecrypt,omitempty"`
	AllowSign            *uint32 `protobuf:"varint,3,opt,name=AllowSign" json:"AllowSign,omitempty"`
	AllowSignatureVerify *uint32 `protobuf:"varint,4,opt,name=AllowSignatureVerify" json:"AllowSignatureVerify,omitempty"`
}

func (x *License_KeyContainer_OperatorSessionKeyPermissions) Reset() {
	*x = License_KeyContainer_OperatorSessionKeyPermissions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[33]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *License_KeyContainer_OperatorSessionKeyPermissions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*License_KeyContainer_OperatorSessionKeyPermissions) ProtoMessage() {}

func (x *License_KeyContainer_OperatorSessionKeyPermissions) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[33]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use License_KeyContainer_OperatorSessionKeyPermissions.ProtoReflect.Descriptor instead.
func (*License_KeyContainer_OperatorSessionKeyPermissions) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{6, 1, 2}
}

func (x *License_KeyContainer_OperatorSessionKeyPermissions) GetAllowEncrypt() uint32 {
	if x != nil && x.AllowEncrypt != nil {
		return *x.AllowEncrypt
	}
	return 0
}

func (x *License_KeyContainer_OperatorSessionKeyPermissions) GetAllowDecrypt() uint32 {
	if x != nil && x.AllowDecrypt != nil {
		return *x.AllowDecrypt
	}
	return 0
}

func (x *License_KeyContainer_OperatorSessionKeyPermissions) GetAllowSign() uint32 {
	if x != nil && x.AllowSign != nil {
		return *x.AllowSign
	}
	return 0
}

func (x *License_KeyContainer_OperatorSessionKeyPermissions) GetAllowSignatureVerify() uint32 {
	if x != nil && x.AllowSignatureVerify != nil {
		return *x.AllowSignatureVerify
	}
	return 0
}

type License_KeyContainer_VideoResolutionConstraint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinResolutionPixels *uint32                                `protobuf:"varint,1,opt,name=MinResolutionPixels" json:"MinResolutionPixels,omitempty"`
	MaxResolutionPixels *uint32                                `protobuf:"varint,2,opt,name=MaxResolutionPixels" json:"MaxResolutionPixels,omitempty"`
	RequiredProtection  *License_KeyContainer_OutputProtection `protobuf:"bytes,3,opt,name=RequiredProtection" json:"RequiredProtection,omitempty"`
}

func (x *License_KeyContainer_VideoResolutionConstraint) Reset() {
	*x = License_KeyContainer_VideoResolutionConstraint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[34]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *License_KeyContainer_VideoResolutionConstraint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*License_KeyContainer_VideoResolutionConstraint) ProtoMessage() {}

func (x *License_KeyContainer_VideoResolutionConstraint) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[34]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use License_KeyContainer_VideoResolutionConstraint.ProtoReflect.Descriptor instead.
func (*License_KeyContainer_VideoResolutionConstraint) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{6, 1, 3}
}

func (x *License_KeyContainer_VideoResolutionConstraint) GetMinResolutionPixels() uint32 {
	if x != nil && x.MinResolutionPixels != nil {
		return *x.MinResolutionPixels
	}
	return 0
}

func (x *License_KeyContainer_VideoResolutionConstraint) GetMaxResolutionPixels() uint32 {
	if x != nil && x.MaxResolutionPixels != nil {
		return *x.MaxResolutionPixels
	}
	return 0
}

func (x *License_KeyContainer_VideoResolutionConstraint) GetRequiredProtection() *License_KeyContainer_OutputProtection {
	if x != nil {
		return x.RequiredProtection
	}
	return nil
}

type LicenseRequest_ContentIdentification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CencId  *LicenseRequest_ContentIdentification_CENC            `protobuf:"bytes,1,opt,name=CencId" json:"CencId,omitempty"`
	WebmId  *LicenseRequest_ContentIdentification_WebM            `protobuf:"bytes,2,opt,name=WebmId" json:"WebmId,omitempty"`
	License *LicenseRequest_ContentIdentification_ExistingLicense `protobuf:"bytes,3,opt,name=License" json:"License,omitempty"`
}

func (x *LicenseRequest_ContentIdentification) Reset() {
	*x = LicenseRequest_ContentIdentification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[35]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LicenseRequest_ContentIdentification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LicenseRequest_ContentIdentification) ProtoMessage() {}

func (x *LicenseRequest_ContentIdentification) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[35]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LicenseRequest_ContentIdentification.ProtoReflect.Descriptor instead.
func (*LicenseRequest_ContentIdentification) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{8, 0}
}

func (x *LicenseRequest_ContentIdentification) GetCencId() *LicenseRequest_ContentIdentification_CENC {
	if x != nil {
		return x.CencId
	}
	return nil
}

func (x *LicenseRequest_ContentIdentification) GetWebmId() *LicenseRequest_ContentIdentification_WebM {
	if x != nil {
		return x.WebmId
	}
	return nil
}

func (x *LicenseRequest_ContentIdentification) GetLicense() *LicenseRequest_ContentIdentification_ExistingLicense {
	if x != nil {
		return x.License
	}
	return nil
}

type LicenseRequest_ContentIdentification_CENC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//optional bytes Pssh = 1; // the client's definition is opaque, it doesn't care about the contents, but the PSSH has a clear definition that is understood and requested by the server, thus I'll replace it with:
	Pssh        *WidevineCencHeader `protobuf:"bytes,1,opt,name=Pssh" json:"Pssh,omitempty"`
	LicenseType *LicenseType        `protobuf:"varint,2,opt,name=LicenseType,enum=LicenseType" json:"LicenseType,omitempty"` // unfortunately the LicenseType symbols are not present, acceptable value seems to only be 1 (is this persist/don't persist? look into it!)
	RequestId   []byte              `protobuf:"bytes,3,opt,name=RequestId" json:"RequestId,omitempty"`
}

func (x *LicenseRequest_ContentIdentification_CENC) Reset() {
	*x = LicenseRequest_ContentIdentification_CENC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[36]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LicenseRequest_ContentIdentification_CENC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LicenseRequest_ContentIdentification_CENC) ProtoMessage() {}

func (x *LicenseRequest_ContentIdentification_CENC) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[36]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LicenseRequest_ContentIdentification_CENC.ProtoReflect.Descriptor instead.
func (*LicenseRequest_ContentIdentification_CENC) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{8, 0, 0}
}

func (x *LicenseRequest_ContentIdentification_CENC) GetPssh() *WidevineCencHeader {
	if x != nil {
		return x.Pssh
	}
	return nil
}

func (x *LicenseRequest_ContentIdentification_CENC) GetLicenseType() LicenseType {
	if x != nil && x.LicenseType != nil {
		return *x.LicenseType
	}
	return LicenseType_ZERO
}

func (x *LicenseRequest_ContentIdentification_CENC) GetRequestId() []byte {
	if x != nil {
		return x.RequestId
	}
	return nil
}

type LicenseRequest_ContentIdentification_WebM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header      []byte       `protobuf:"bytes,1,opt,name=Header" json:"Header,omitempty"` // identical to CENC, aside from PSSH and the parent field number used
	LicenseType *LicenseType `protobuf:"varint,2,opt,name=LicenseType,enum=LicenseType" json:"LicenseType,omitempty"`
	RequestId   []byte       `protobuf:"bytes,3,opt,name=RequestId" json:"RequestId,omitempty"`
}

func (x *LicenseRequest_ContentIdentification_WebM) Reset() {
	*x = LicenseRequest_ContentIdentification_WebM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[37]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LicenseRequest_ContentIdentification_WebM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LicenseRequest_ContentIdentification_WebM) ProtoMessage() {}

func (x *LicenseRequest_ContentIdentification_WebM) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[37]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LicenseRequest_ContentIdentification_WebM.ProtoReflect.Descriptor instead.
func (*LicenseRequest_ContentIdentification_WebM) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{8, 0, 1}
}

func (x *LicenseRequest_ContentIdentification_WebM) GetHeader() []byte {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *LicenseRequest_ContentIdentification_WebM) GetLicenseType() LicenseType {
	if x != nil && x.LicenseType != nil {
		return *x.LicenseType
	}
	return LicenseType_ZERO
}

func (x *LicenseRequest_ContentIdentification_WebM) GetRequestId() []byte {
	if x != nil {
		return x.RequestId
	}
	return nil
}

type LicenseRequest_ContentIdentification_ExistingLicense struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LicenseId              *LicenseIdentification `protobuf:"bytes,1,opt,name=LicenseId" json:"LicenseId,omitempty"`
	SecondsSinceStarted    *uint32                `protobuf:"varint,2,opt,name=SecondsSinceStarted" json:"SecondsSinceStarted,omitempty"`
	SecondsSinceLastPlayed *uint32                `protobuf:"varint,3,opt,name=SecondsSinceLastPlayed" json:"SecondsSinceLastPlayed,omitempty"`
	SessionUsageTableEntry []byte                 `protobuf:"bytes,4,opt,name=SessionUsageTableEntry" json:"SessionUsageTableEntry,omitempty"` // interesting! try to figure out the connection between the usage table blob and KCB!
}

func (x *LicenseRequest_ContentIdentification_ExistingLicense) Reset() {
	*x = LicenseRequest_ContentIdentification_ExistingLicense{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[38]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LicenseRequest_ContentIdentification_ExistingLicense) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LicenseRequest_ContentIdentification_ExistingLicense) ProtoMessage() {}

func (x *LicenseRequest_ContentIdentification_ExistingLicense) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[38]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LicenseRequest_ContentIdentification_ExistingLicense.ProtoReflect.Descriptor instead.
func (*LicenseRequest_ContentIdentification_ExistingLicense) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{8, 0, 2}
}

func (x *LicenseRequest_ContentIdentification_ExistingLicense) GetLicenseId() *LicenseIdentification {
	if x != nil {
		return x.LicenseId
	}
	return nil
}

func (x *LicenseRequest_ContentIdentification_ExistingLicense) GetSecondsSinceStarted() uint32 {
	if x != nil && x.SecondsSinceStarted != nil {
		return *x.SecondsSinceStarted
	}
	return 0
}

func (x *LicenseRequest_ContentIdentification_ExistingLicense) GetSecondsSinceLastPlayed() uint32 {
	if x != nil && x.SecondsSinceLastPlayed != nil {
		return *x.SecondsSinceLastPlayed
	}
	return 0
}

func (x *LicenseRequest_ContentIdentification_ExistingLicense) GetSessionUsageTableEntry() []byte {
	if x != nil {
		return x.SessionUsageTableEntry
	}
	return nil
}

type LicenseRequestRaw_ContentIdentification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CencId  *LicenseRequestRaw_ContentIdentification_CENC            `protobuf:"bytes,1,opt,name=CencId" json:"CencId,omitempty"`
	WebmId  *LicenseRequestRaw_ContentIdentification_WebM            `protobuf:"bytes,2,opt,name=WebmId" json:"WebmId,omitempty"`
	License *LicenseRequestRaw_ContentIdentification_ExistingLicense `protobuf:"bytes,3,opt,name=License" json:"License,omitempty"`
}

func (x *LicenseRequestRaw_ContentIdentification) Reset() {
	*x = LicenseRequestRaw_ContentIdentification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[39]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LicenseRequestRaw_ContentIdentification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LicenseRequestRaw_ContentIdentification) ProtoMessage() {}

func (x *LicenseRequestRaw_ContentIdentification) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[39]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LicenseRequestRaw_ContentIdentification.ProtoReflect.Descriptor instead.
func (*LicenseRequestRaw_ContentIdentification) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{9, 0}
}

func (x *LicenseRequestRaw_ContentIdentification) GetCencId() *LicenseRequestRaw_ContentIdentification_CENC {
	if x != nil {
		return x.CencId
	}
	return nil
}

func (x *LicenseRequestRaw_ContentIdentification) GetWebmId() *LicenseRequestRaw_ContentIdentification_WebM {
	if x != nil {
		return x.WebmId
	}
	return nil
}

func (x *LicenseRequestRaw_ContentIdentification) GetLicense() *LicenseRequestRaw_ContentIdentification_ExistingLicense {
	if x != nil {
		return x.License
	}
	return nil
}

type LicenseRequestRaw_ContentIdentification_CENC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pssh []byte `protobuf:"bytes,1,opt,name=Pssh" json:"Pssh,omitempty"` // the client's definition is opaque, it doesn't care about the contents, but the PSSH has a clear definition that is understood and requested by the server, thus I'll replace it with:
	//optional WidevineCencHeader Pssh = 1;
	LicenseType *LicenseType `protobuf:"varint,2,opt,name=LicenseType,enum=LicenseType" json:"LicenseType,omitempty"` // unfortunately the LicenseType symbols are not present, acceptable value seems to only be 1 (is this persist/don't persist? look into it!)
	RequestId   []byte       `protobuf:"bytes,3,opt,name=RequestId" json:"RequestId,omitempty"`
}

func (x *LicenseRequestRaw_ContentIdentification_CENC) Reset() {
	*x = LicenseRequestRaw_ContentIdentification_CENC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[40]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LicenseRequestRaw_ContentIdentification_CENC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LicenseRequestRaw_ContentIdentification_CENC) ProtoMessage() {}

func (x *LicenseRequestRaw_ContentIdentification_CENC) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[40]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LicenseRequestRaw_ContentIdentification_CENC.ProtoReflect.Descriptor instead.
func (*LicenseRequestRaw_ContentIdentification_CENC) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{9, 0, 0}
}

func (x *LicenseRequestRaw_ContentIdentification_CENC) GetPssh() []byte {
	if x != nil {
		return x.Pssh
	}
	return nil
}

func (x *LicenseRequestRaw_ContentIdentification_CENC) GetLicenseType() LicenseType {
	if x != nil && x.LicenseType != nil {
		return *x.LicenseType
	}
	return LicenseType_ZERO
}

func (x *LicenseRequestRaw_ContentIdentification_CENC) GetRequestId() []byte {
	if x != nil {
		return x.RequestId
	}
	return nil
}

type LicenseRequestRaw_ContentIdentification_WebM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header      []byte       `protobuf:"bytes,1,opt,name=Header" json:"Header,omitempty"` // identical to CENC, aside from PSSH and the parent field number used
	LicenseType *LicenseType `protobuf:"varint,2,opt,name=LicenseType,enum=LicenseType" json:"LicenseType,omitempty"`
	RequestId   []byte       `protobuf:"bytes,3,opt,name=RequestId" json:"RequestId,omitempty"`
}

func (x *LicenseRequestRaw_ContentIdentification_WebM) Reset() {
	*x = LicenseRequestRaw_ContentIdentification_WebM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[41]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LicenseRequestRaw_ContentIdentification_WebM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LicenseRequestRaw_ContentIdentification_WebM) ProtoMessage() {}

func (x *LicenseRequestRaw_ContentIdentification_WebM) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[41]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LicenseRequestRaw_ContentIdentification_WebM.ProtoReflect.Descriptor instead.
func (*LicenseRequestRaw_ContentIdentification_WebM) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{9, 0, 1}
}

func (x *LicenseRequestRaw_ContentIdentification_WebM) GetHeader() []byte {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *LicenseRequestRaw_ContentIdentification_WebM) GetLicenseType() LicenseType {
	if x != nil && x.LicenseType != nil {
		return *x.LicenseType
	}
	return LicenseType_ZERO
}

func (x *LicenseRequestRaw_ContentIdentification_WebM) GetRequestId() []byte {
	if x != nil {
		return x.RequestId
	}
	return nil
}

type LicenseRequestRaw_ContentIdentification_ExistingLicense struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LicenseId              *LicenseIdentification `protobuf:"bytes,1,opt,name=LicenseId" json:"LicenseId,omitempty"`
	SecondsSinceStarted    *uint32                `protobuf:"varint,2,opt,name=SecondsSinceStarted" json:"SecondsSinceStarted,omitempty"`
	SecondsSinceLastPlayed *uint32                `protobuf:"varint,3,opt,name=SecondsSinceLastPlayed" json:"SecondsSinceLastPlayed,omitempty"`
	SessionUsageTableEntry []byte                 `protobuf:"bytes,4,opt,name=SessionUsageTableEntry" json:"SessionUsageTableEntry,omitempty"` // interesting! try to figure out the connection between the usage table blob and KCB!
}

func (x *LicenseRequestRaw_ContentIdentification_ExistingLicense) Reset() {
	*x = LicenseRequestRaw_ContentIdentification_ExistingLicense{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[42]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LicenseRequestRaw_ContentIdentification_ExistingLicense) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LicenseRequestRaw_ContentIdentification_ExistingLicense) ProtoMessage() {}

func (x *LicenseRequestRaw_ContentIdentification_ExistingLicense) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[42]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LicenseRequestRaw_ContentIdentification_ExistingLicense.ProtoReflect.Descriptor instead.
func (*LicenseRequestRaw_ContentIdentification_ExistingLicense) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{9, 0, 2}
}

func (x *LicenseRequestRaw_ContentIdentification_ExistingLicense) GetLicenseId() *LicenseIdentification {
	if x != nil {
		return x.LicenseId
	}
	return nil
}

func (x *LicenseRequestRaw_ContentIdentification_ExistingLicense) GetSecondsSinceStarted() uint32 {
	if x != nil && x.SecondsSinceStarted != nil {
		return *x.SecondsSinceStarted
	}
	return 0
}

func (x *LicenseRequestRaw_ContentIdentification_ExistingLicense) GetSecondsSinceLastPlayed() uint32 {
	if x != nil && x.SecondsSinceLastPlayed != nil {
		return *x.SecondsSinceLastPlayed
	}
	return 0
}

func (x *LicenseRequestRaw_ContentIdentification_ExistingLicense) GetSessionUsageTableEntry() []byte {
	if x != nil {
		return x.SessionUsageTableEntry
	}
	return nil
}

type FileHashes_Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename    *string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	TestSigning *bool   `protobuf:"varint,2,opt,name=test_signing,json=testSigning" json:"test_signing,omitempty"` //0 - release, 1 - testing
	SHA512Hash  []byte  `protobuf:"bytes,3,opt,name=SHA512Hash" json:"SHA512Hash,omitempty"`
	MainExe     *bool   `protobuf:"varint,4,opt,name=main_exe,json=mainExe" json:"main_exe,omitempty"` //0 for dlls, 1 for exe, this is field 3 in file
	Signature   []byte  `protobuf:"bytes,5,opt,name=signature" json:"signature,omitempty"`
}

func (x *FileHashes_Signature) Reset() {
	*x = FileHashes_Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wv_proto2_proto_msgTypes[43]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileHashes_Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileHashes_Signature) ProtoMessage() {}

func (x *FileHashes_Signature) ProtoReflect() protoreflect.Message {
	mi := &file_wv_proto2_proto_msgTypes[43]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileHashes_Signature.ProtoReflect.Descriptor instead.
func (*FileHashes_Signature) Descriptor() ([]byte, []int) {
	return file_wv_proto2_proto_rawDescGZIP(), []int{26, 0}
}

func (x *FileHashes_Signature) GetFilename() string {
	if x != nil && x.Filename != nil {
		return *x.Filename
	}
	return ""
}

func (x *FileHashes_Signature) GetTestSigning() bool {
	if x != nil && x.TestSigning != nil {
		return *x.TestSigning
	}
	return false
}

func (x *FileHashes_Signature) GetSHA512Hash() []byte {
	if x != nil {
		return x.SHA512Hash
	}
	return nil
}

func (x *FileHashes_Signature) GetMainExe() bool {
	if x != nil && x.MainExe != nil {
		return *x.MainExe
	}
	return false
}

func (x *FileHashes_Signature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_wv_proto2_proto protoreflect.FileDescriptor

var file_wv_proto2_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x77, 0x76, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xae, 0x07, 0x0a, 0x14, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x2e, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x3f, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x30, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x4c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x59, 0x0a, 0x13, 0x5f, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x52, 0x12, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x0b, 0x5f, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x61,
	0x73, 0x68, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x52, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73,
	0x68, 0x65, 0x73, 0x1a, 0x35, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x80, 0x03, 0x0a, 0x12, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3e, 0x0a, 0x1a, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1a, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x5c, 0x0a, 0x0e, 0x4d, 0x61, 0x78, 0x48, 0x64,
	0x63, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x34, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x48, 0x64, 0x63, 0x70, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x4d, 0x61, 0x78, 0x48, 0x64, 0x63, 0x70, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x13, 0x4f, 0x65, 0x6d, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x13, 0x4f, 0x65, 0x6d, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x41, 0x70, 0x69,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x54, 0x0a, 0x0b, 0x48, 0x64, 0x63, 0x70, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x44, 0x43, 0x50, 0x5f, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x44, 0x43, 0x50, 0x5f, 0x56, 0x31,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x44, 0x43, 0x50, 0x5f, 0x56, 0x32, 0x10, 0x02, 0x12,
	0x0d, 0x0a, 0x09, 0x48, 0x44, 0x43, 0x50, 0x5f, 0x56, 0x32, 0x5f, 0x31, 0x10, 0x03, 0x12, 0x0d,
	0x0a, 0x09, 0x48, 0x44, 0x43, 0x50, 0x5f, 0x56, 0x32, 0x5f, 0x32, 0x10, 0x04, 0x22, 0x53, 0x0a,
	0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4b, 0x45,
	0x59, 0x42, 0x4f, 0x58, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45,
	0x5f, 0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x22,
	0x0a, 0x1e, 0x52, 0x45, 0x4d, 0x4f, 0x54, 0x45, 0x5f, 0x41, 0x54, 0x54, 0x45, 0x53, 0x54, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x45,
	0x10, 0x02, 0x22, 0xfa, 0x02, 0x0a, 0x11, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x64,
	0x12, 0x32, 0x0a, 0x14, 0x54, 0x65, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65,
	0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14,
	0x54, 0x65, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x22, 0x4b, 0x0a, 0x0f, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x4f, 0x4f, 0x54, 0x10, 0x00, 0x12,
	0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x54, 0x45, 0x10,
	0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45,
	0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x03, 0x22,
	0xe6, 0x01, 0x0a, 0x17, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x53,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0c, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x42, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2a, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x36, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2b, 0x0a, 0x11, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x09, 0x0a, 0x05, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52,
	0x45, 0x56, 0x4f, 0x4b, 0x45, 0x44, 0x10, 0x01, 0x22, 0x97, 0x01, 0x0a, 0x1b, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x46, 0x0a, 0x11, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x11, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x97, 0x02, 0x0a, 0x1d, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x46, 0x0a, 0x1e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x1e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x45, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x11, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x65, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x49, 0x76, 0x18,
	0x04, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x13, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x49, 0x76, 0x12, 0x30, 0x0a, 0x13, 0x45, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x4b, 0x65,
	0x79, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x13, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x65, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x4b, 0x65, 0x79, 0x22, 0xe3, 0x01, 0x0a,
	0x15, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0c, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x32,
	0x0a, 0x14, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x14, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x8d, 0x13, 0x0a, 0x07, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x4c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x02, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x07, 0x5f, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x27, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x10, 0x4c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x10, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x19, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x19, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x13, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x10, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x65, 0x1a, 0x96, 0x04, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x43, 0x61, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x43,
	0x61, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x6e, 0x50, 0x65, 0x72,
	0x73, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x43, 0x61, 0x6e, 0x50,
	0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x6e, 0x52, 0x65, 0x6e,
	0x65, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x43, 0x61, 0x6e, 0x52, 0x65, 0x6e,
	0x65, 0x77, 0x12, 0x34, 0x0a, 0x15, 0x52, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x15, 0x52, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x38, 0x0a, 0x17, 0x50, 0x6c, 0x61, 0x79,
	0x62, 0x61, 0x63, 0x6b, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x17, 0x50, 0x6c, 0x61, 0x79, 0x62,
	0x61, 0x63, 0x6b, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x12, 0x36, 0x0a, 0x16, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x16, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x46, 0x0a, 0x1e, 0x52, 0x65,
	0x6e, 0x65, 0x77, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x1e, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x52, 0x65,
	0x6e, 0x65, 0x77, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x30,
	0x0a, 0x13, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x52, 0x65, 0x6e,
	0x65, 0x77, 0x61, 0x6c, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x12, 0x40, 0x0a, 0x1b, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x52, 0x65, 0x74, 0x72, 0x79,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1b, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x52, 0x65,
	0x74, 0x72, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x57, 0x69, 0x74, 0x68, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x52, 0x65, 0x6e, 0x65,
	0x77, 0x57, 0x69, 0x74, 0x68, 0x55, 0x73, 0x61, 0x67, 0x65, 0x1a, 0xa5, 0x0c, 0x0a, 0x0c, 0x4b,
	0x65, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x49, 0x76, 0x12, 0x10, 0x0a, 0x03, 0x4b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x31, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x4c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x2e, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x39, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x23, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x56, 0x0a, 0x12, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x2e, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x12, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x58, 0x0a, 0x13, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a,
	0x0b, 0x5f, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x4b, 0x65, 0x79,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x52, 0x0a, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x12, 0x7a, 0x0a, 0x1e, 0x5f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b,
	0x65, 0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x1d, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65,
	0x79, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x6f, 0x0a, 0x1a,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2f, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x74, 0x52, 0x1a, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0xec, 0x01,
	0x0a, 0x10, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x48, 0x0a, 0x04, 0x48, 0x64, 0x63, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x34, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61,
	0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x48, 0x64, 0x63, 0x70, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x48, 0x64, 0x63, 0x70, 0x12, 0x49, 0x0a, 0x09,
	0x43, 0x67, 0x6d, 0x73, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2b, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x47, 0x4d, 0x53, 0x52, 0x09, 0x43, 0x67,
	0x6d, 0x73, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x22, 0x43, 0x0a, 0x04, 0x43, 0x47, 0x4d, 0x53, 0x12,
	0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x50, 0x59, 0x5f, 0x46, 0x52, 0x45, 0x45, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x43, 0x4f, 0x50, 0x59, 0x5f, 0x4f, 0x4e, 0x43, 0x45, 0x10, 0x02, 0x12, 0x0e, 0x0a,
	0x0a, 0x43, 0x4f, 0x50, 0x59, 0x5f, 0x4e, 0x45, 0x56, 0x45, 0x52, 0x10, 0x03, 0x12, 0x0d, 0x0a,
	0x09, 0x43, 0x47, 0x4d, 0x53, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x2a, 0x1a, 0x46, 0x0a, 0x0a,
	0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x4b, 0x65,
	0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x0c, 0x52, 0x0f, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x76, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0c,
	0x52, 0x02, 0x49, 0x76, 0x1a, 0xb9, 0x01, 0x0a, 0x1d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x69, 0x67, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x69, 0x67, 0x6e, 0x12, 0x32, 0x0a, 0x14,
	0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x41, 0x6c, 0x6c, 0x6f,
	0x77, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x1a, 0xd7, 0x01, 0x0a, 0x19, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x30,
	0x0a, 0x13, 0x4d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x69, 0x78, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x4d, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x69, 0x78, 0x65, 0x6c, 0x73,
	0x12, 0x30, 0x0a, 0x13, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x69, 0x78, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x4d,
	0x61, 0x78, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x69, 0x78, 0x65,
	0x6c, 0x73, 0x12, 0x56, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x50, 0x72,
	0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4a, 0x0a, 0x07, 0x4b, 0x65,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x49, 0x47, 0x4e, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12,
	0x0f, 0x0a, 0x0b, 0x4b, 0x45, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x10, 0x03,
	0x12, 0x14, 0x0a, 0x10, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x53, 0x45, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x22, 0x7a, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x57, 0x5f, 0x53, 0x45,
	0x43, 0x55, 0x52, 0x45, 0x5f, 0x43, 0x52, 0x59, 0x50, 0x54, 0x4f, 0x10, 0x01, 0x12, 0x14, 0x0a,
	0x10, 0x53, 0x57, 0x5f, 0x53, 0x45, 0x43, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x45, 0x43, 0x4f, 0x44,
	0x45, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x48, 0x57, 0x5f, 0x53, 0x45, 0x43, 0x55, 0x52, 0x45,
	0x5f, 0x43, 0x52, 0x59, 0x50, 0x54, 0x4f, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x48, 0x57, 0x5f,
	0x53, 0x45, 0x43, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x45, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x04, 0x12,
	0x11, 0x0a, 0x0d, 0x48, 0x57, 0x5f, 0x53, 0x45, 0x43, 0x55, 0x52, 0x45, 0x5f, 0x41, 0x4c, 0x4c,
	0x10, 0x05, 0x22, 0xa3, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x31, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x09, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x60, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x1e, 0x0a, 0x1a, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43,
	0x45, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12,
	0x1e, 0x0a, 0x1a, 0x52, 0x45, 0x56, 0x4f, 0x4b, 0x45, 0x44, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43,
	0x45, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12,
	0x17, 0x0a, 0x13, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x41, 0x56, 0x41,
	0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x03, 0x22, 0xcb, 0x09, 0x0a, 0x0e, 0x4c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x43,
	0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1b, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x19, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x19, 0x4b, 0x65, 0x79, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x3a, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x28, 0x0a, 0x0f, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4e, 0x6f,
	0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x4b, 0x65, 0x79, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x11, 0x45, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65,
	0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0xc9, 0x05, 0x0a, 0x15, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x06, 0x43, 0x65, 0x6e, 0x63, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x45, 0x4e, 0x43, 0x52, 0x06,
	0x43, 0x65, 0x6e, 0x63, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x06, 0x57, 0x65, 0x62, 0x6d, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x57, 0x65,
	0x62, 0x4d, 0x52, 0x06, 0x57, 0x65, 0x62, 0x6d, 0x49, 0x64, 0x12, 0x4f, 0x0a, 0x07, 0x4c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x4c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x52, 0x07, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x1a, 0x7d, 0x0a, 0x04, 0x43,
	0x45, 0x4e, 0x43, 0x12, 0x27, 0x0a, 0x04, 0x50, 0x73, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x57, 0x69, 0x64, 0x65, 0x76, 0x69, 0x6e, 0x65, 0x43, 0x65, 0x6e, 0x63,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x04, 0x50, 0x73, 0x73, 0x68, 0x12, 0x2e, 0x0a, 0x0b,
	0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0b, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x6c, 0x0a, 0x04, 0x57, 0x65,
	0x62, 0x4d, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x0b, 0x4c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0c, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x4c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x1a, 0xe9, 0x01, 0x0a, 0x0f, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09,
	0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x49, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x53, 0x69, 0x6e,
	0x63, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x13, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x12, 0x36, 0x0a, 0x16, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x53,
	0x69, 0x6e, 0x63, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x53, 0x69, 0x6e,
	0x63, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x12, 0x36, 0x0a, 0x16,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x16, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x22, 0x30, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x45, 0x57, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x52, 0x45, 0x4e, 0x45, 0x57, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x4c,
	0x45, 0x41, 0x53, 0x45, 0x10, 0x03, 0x22, 0xc8, 0x09, 0x0a, 0x11, 0x4c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x61, 0x77, 0x12, 0x31, 0x0a, 0x08,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x46, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x61, 0x77, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x61, 0x77, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3c, 0x0a,
	0x19, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4e, 0x6f, 0x6e, 0x63, 0x65,
	0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x19, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4e, 0x6f, 0x6e, 0x63,
	0x65, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x3a, 0x0a, 0x0f, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x4b, 0x65, 0x79, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0f, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4e, 0x6f, 0x6e, 0x63,
	0x65, 0x12, 0x4c, 0x0a, 0x11, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x45, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x1a,
	0xbd, 0x05, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x06, 0x43, 0x65, 0x6e,
	0x63, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x4c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x61, 0x77, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x45, 0x4e, 0x43, 0x52, 0x06, 0x43, 0x65, 0x6e, 0x63, 0x49, 0x64,
	0x12, 0x45, 0x0a, 0x06, 0x57, 0x65, 0x62, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x61, 0x77, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x57, 0x65, 0x62, 0x4d, 0x52,
	0x06, 0x57, 0x65, 0x62, 0x6d, 0x49, 0x64, 0x12, 0x52, 0x0a, 0x07, 0x4c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x61, 0x77, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x52, 0x07, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x1a, 0x68, 0x0a, 0x04, 0x43,
	0x45, 0x4e, 0x43, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x73, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x50, 0x73, 0x73, 0x68, 0x12, 0x2e, 0x0a, 0x0b, 0x4c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x4c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x4c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x6c, 0x0a, 0x04, 0x57, 0x65, 0x62, 0x4d, 0x12, 0x16, 0x0a,
	0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x0b, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x4c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x1a, 0xe9, 0x01, 0x0a, 0x0f, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x4c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x4c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x64, 0x12, 0x30, 0x0a,
	0x13, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12,
	0x36, 0x0a, 0x16, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x4c,
	0x61, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x16, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x4c, 0x61, 0x73,
	0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x12, 0x36, 0x0a, 0x16, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x16, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22,
	0x30, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07,
	0x0a, 0x03, 0x4e, 0x45, 0x57, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x4e, 0x45, 0x57,
	0x41, 0x4c, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x4c, 0x45, 0x41, 0x53, 0x45, 0x10,
	0x03, 0x22, 0xfc, 0x02, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65,
	0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x6f, 0x63, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x6f, 0x63, 0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x61, 0x6e,
	0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x59, 0x65, 0x61, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x59, 0x65, 0x61,
	0x72, 0x12, 0x4c, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x57, 0x76, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x52, 0x0d, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x1e, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22,
	0x4f, 0x0a, 0x0f, 0x57, 0x76, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x45, 0x56,
	0x45, 0x4c, 0x5f, 0x31, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f,
	0x32, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x33, 0x10, 0x03,
	0x22, 0x15, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x16,
	0x0a, 0x14, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0b,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x53, 0x61, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x61,
	0x6c, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x69, 0x74, 0x22,
	0x0e, 0x0a, 0x0c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22,
	0x1d, 0x0a, 0x1b, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xac,
	0x01, 0x0a, 0x17, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x41, 0x0a, 0x12, 0x5f, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x11, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x53,
	0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x22, 0x1b, 0x0a,
	0x19, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xd0, 0x02, 0x0a, 0x0d, 0x53,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x1c,
	0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x40, 0x0a, 0x11,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7d,
	0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a,
	0x0f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x10, 0x02, 0x12,
	0x12, 0x0a, 0x0e, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53,
	0x45, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x43,
	0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f,
	0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x45, 0x10, 0x05, 0x22, 0xd3, 0x03,
	0x0a, 0x12, 0x57, 0x69, 0x64, 0x65, 0x76, 0x69, 0x6e, 0x65, 0x43, 0x65, 0x6e, 0x63, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x57, 0x69, 0x64, 0x65, 0x76, 0x69,
	0x6e, 0x65, 0x43, 0x65, 0x6e, 0x63, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68,
	0x6d, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x70,
	0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x2e, 0x0a, 0x13, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x27, 0x0a, 0x0f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x65,
	0x64, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x28, 0x0a, 0x09, 0x41, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x45, 0x4e, 0x43, 0x52,
	0x59, 0x50, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x45, 0x53, 0x43, 0x54,
	0x52, 0x10, 0x01, 0x22, 0xef, 0x02, 0x0a, 0x14, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b,
	0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4b, 0x65, 0x79, 0x12, 0x40, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x74,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x74, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7d, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x49,
	0x43, 0x45, 0x4e, 0x53, 0x45, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41,
	0x54, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13,
	0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43,
	0x41, 0x54, 0x45, 0x10, 0x05, 0x22, 0xf8, 0x02, 0x0a, 0x17, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x61,
	0x77, 0x12, 0x38, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x24, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x61, 0x77, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x03, 0x4d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x61, 0x77, 0x52, 0x03, 0x4d, 0x73,
	0x67, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0a, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12,
	0x40, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x7d, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45,
	0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x53, 0x50,
	0x4f, 0x4e, 0x53, 0x45, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43,
	0x45, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x45, 0x52, 0x56, 0x49,
	0x43, 0x45, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x45, 0x10, 0x05,
	0x22, 0xda, 0x02, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1a, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1a, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x1c,
	0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x40, 0x0a, 0x11,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7d,
	0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a,
	0x0f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x10, 0x02, 0x12,
	0x12, 0x0a, 0x0e, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53,
	0x45, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x43,
	0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f,
	0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x45, 0x10, 0x05, 0x22, 0x80, 0x03,
	0x0a, 0x18, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x03, 0x4d, 0x73,
	0x67, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0a, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12,
	0x40, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x7d, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45,
	0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x53, 0x50,
	0x4f, 0x4e, 0x53, 0x45, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43,
	0x45, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x45, 0x52, 0x56, 0x49,
	0x43, 0x45, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x45, 0x10, 0x05,
	0x22, 0x81, 0x02, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x1a, 0xa3,
	0x01, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x65, 0x73, 0x74,
	0x5f, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x74, 0x65, 0x73, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x53,
	0x48, 0x41, 0x35, 0x31, 0x32, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x53, 0x48, 0x41, 0x35, 0x31, 0x32, 0x48, 0x61, 0x73, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x61, 0x69, 0x6e, 0x5f, 0x65, 0x78, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6d,
	0x61, 0x69, 0x6e, 0x45, 0x78, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x2a, 0x31, 0x0a, 0x0b, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x5a, 0x45, 0x52, 0x4f, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x46,
	0x46, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x02, 0x2a, 0x1e, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x55,
	0x52, 0x52, 0x45, 0x4e, 0x54, 0x10, 0x15, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x3b, 0x67, 0x65, 0x6e,
}

var (
	file_wv_proto2_proto_rawDescOnce sync.Once
	file_wv_proto2_proto_rawDescData = file_wv_proto2_proto_rawDesc
)

func file_wv_proto2_proto_rawDescGZIP() []byte {
	file_wv_proto2_proto_rawDescOnce.Do(func() {
		file_wv_proto2_proto_rawDescData = protoimpl.X.CompressGZIP(file_wv_proto2_proto_rawDescData)
	})
	return file_wv_proto2_proto_rawDescData
}

var file_wv_proto2_proto_enumTypes = make([]protoimpl.EnumInfo, 19)
var file_wv_proto2_proto_msgTypes = make([]protoimpl.MessageInfo, 44)
var file_wv_proto2_proto_goTypes = []interface{}{
	(LicenseType)(0),                    // 0: LicenseType
	(ProtocolVersion)(0),                // 1: ProtocolVersion
	(ClientIdentification_TokenType)(0), // 2: ClientIdentification.TokenType
	(ClientIdentification_ClientCapabilities_HdcpVersion)(0),        // 3: ClientIdentification.ClientCapabilities.HdcpVersion
	(DeviceCertificate_CertificateType)(0),                          // 4: DeviceCertificate.CertificateType
	(DeviceCertificateStatus_CertificateStatus)(0),                  // 5: DeviceCertificateStatus.CertificateStatus
	(License_KeyContainer_KeyType)(0),                               // 6: License.KeyContainer.KeyType
	(License_KeyContainer_SecurityLevel)(0),                         // 7: License.KeyContainer.SecurityLevel
	(License_KeyContainer_OutputProtection_CGMS)(0),                 // 8: License.KeyContainer.OutputProtection.CGMS
	(LicenseError_Error)(0),                                         // 9: LicenseError.Error
	(LicenseRequest_RequestType)(0),                                 // 10: LicenseRequest.RequestType
	(LicenseRequestRaw_RequestType)(0),                              // 11: LicenseRequestRaw.RequestType
	(ProvisionedDeviceInfo_WvSecurityLevel)(0),                      // 12: ProvisionedDeviceInfo.WvSecurityLevel
	(SignedMessage_MessageType)(0),                                  // 13: SignedMessage.MessageType
	(WidevineCencHeader_Algorithm)(0),                               // 14: WidevineCencHeader.Algorithm
	(SignedLicenseRequest_MessageType)(0),                           // 15: SignedLicenseRequest.MessageType
	(SignedLicenseRequestRaw_MessageType)(0),                        // 16: SignedLicenseRequestRaw.MessageType
	(SignedLicense_MessageType)(0),                                  // 17: SignedLicense.MessageType
	(SignedServiceCertificate_MessageType)(0),                       // 18: SignedServiceCertificate.MessageType
	(*ClientIdentification)(nil),                                    // 19: ClientIdentification
	(*DeviceCertificate)(nil),                                       // 20: DeviceCertificate
	(*DeviceCertificateStatus)(nil),                                 // 21: DeviceCertificateStatus
	(*DeviceCertificateStatusList)(nil),                             // 22: DeviceCertificateStatusList
	(*EncryptedClientIdentification)(nil),                           // 23: EncryptedClientIdentification
	(*LicenseIdentification)(nil),                                   // 24: LicenseIdentification
	(*License)(nil),                                                 // 25: License
	(*LicenseError)(nil),                                            // 26: LicenseError
	(*LicenseRequest)(nil),                                          // 27: LicenseRequest
	(*LicenseRequestRaw)(nil),                                       // 28: LicenseRequestRaw
	(*ProvisionedDeviceInfo)(nil),                                   // 29: ProvisionedDeviceInfo
	(*ProvisioningOptions)(nil),                                     // 30: ProvisioningOptions
	(*ProvisioningRequest)(nil),                                     // 31: ProvisioningRequest
	(*ProvisioningResponse)(nil),                                    // 32: ProvisioningResponse
	(*RemoteAttestation)(nil),                                       // 33: RemoteAttestation
	(*SessionInit)(nil),                                             // 34: SessionInit
	(*SessionState)(nil),                                            // 35: SessionState
	(*SignedCertificateStatusList)(nil),                             // 36: SignedCertificateStatusList
	(*SignedDeviceCertificate)(nil),                                 // 37: SignedDeviceCertificate
	(*SignedProvisioningMessage)(nil),                               // 38: SignedProvisioningMessage
	(*SignedMessage)(nil),                                           // 39: SignedMessage
	(*WidevineCencHeader)(nil),                                      // 40: WidevineCencHeader
	(*SignedLicenseRequest)(nil),                                    // 41: SignedLicenseRequest
	(*SignedLicenseRequestRaw)(nil),                                 // 42: SignedLicenseRequestRaw
	(*SignedLicense)(nil),                                           // 43: SignedLicense
	(*SignedServiceCertificate)(nil),                                // 44: SignedServiceCertificate
	(*FileHashes)(nil),                                              // 45: FileHashes
	(*ClientIdentification_NameValue)(nil),                          // 46: ClientIdentification.NameValue
	(*ClientIdentification_ClientCapabilities)(nil),                 // 47: ClientIdentification.ClientCapabilities
	(*License_Policy)(nil),                                          // 48: License.Policy
	(*License_KeyContainer)(nil),                                    // 49: License.KeyContainer
	(*License_KeyContainer_OutputProtection)(nil),                   // 50: License.KeyContainer.OutputProtection
	(*License_KeyContainer_KeyControl)(nil),                         // 51: License.KeyContainer.KeyControl
	(*License_KeyContainer_OperatorSessionKeyPermissions)(nil),      // 52: License.KeyContainer.OperatorSessionKeyPermissions
	(*License_KeyContainer_VideoResolutionConstraint)(nil),          // 53: License.KeyContainer.VideoResolutionConstraint
	(*LicenseRequest_ContentIdentification)(nil),                    // 54: LicenseRequest.ContentIdentification
	(*LicenseRequest_ContentIdentification_CENC)(nil),               // 55: LicenseRequest.ContentIdentification.CENC
	(*LicenseRequest_ContentIdentification_WebM)(nil),               // 56: LicenseRequest.ContentIdentification.WebM
	(*LicenseRequest_ContentIdentification_ExistingLicense)(nil),    // 57: LicenseRequest.ContentIdentification.ExistingLicense
	(*LicenseRequestRaw_ContentIdentification)(nil),                 // 58: LicenseRequestRaw.ContentIdentification
	(*LicenseRequestRaw_ContentIdentification_CENC)(nil),            // 59: LicenseRequestRaw.ContentIdentification.CENC
	(*LicenseRequestRaw_ContentIdentification_WebM)(nil),            // 60: LicenseRequestRaw.ContentIdentification.WebM
	(*LicenseRequestRaw_ContentIdentification_ExistingLicense)(nil), // 61: LicenseRequestRaw.ContentIdentification.ExistingLicense
	(*FileHashes_Signature)(nil),                                    // 62: FileHashes.Signature
}
var file_wv_proto2_proto_depIdxs = []int32{
	2,  // 0: ClientIdentification.Type:type_name -> ClientIdentification.TokenType
	37, // 1: ClientIdentification.Token:type_name -> SignedDeviceCertificate
	46, // 2: ClientIdentification.ClientInfo:type_name -> ClientIdentification.NameValue
	47, // 3: ClientIdentification._ClientCapabilities:type_name -> ClientIdentification.ClientCapabilities
	45, // 4: ClientIdentification._FileHashes:type_name -> FileHashes
	4,  // 5: DeviceCertificate.Type:type_name -> DeviceCertificate.CertificateType
	5,  // 6: DeviceCertificateStatus.Status:type_name -> DeviceCertificateStatus.CertificateStatus
	29, // 7: DeviceCertificateStatus.DeviceInfo:type_name -> ProvisionedDeviceInfo
	21, // 8: DeviceCertificateStatusList.CertificateStatus:type_name -> DeviceCertificateStatus
	0,  // 9: LicenseIdentification.Type:type_name -> LicenseType
	24, // 10: License.Id:type_name -> LicenseIdentification
	48, // 11: License._Policy:type_name -> License.Policy
	49, // 12: License.Key:type_name -> License.KeyContainer
	9,  // 13: LicenseError.ErrorCode:type_name -> LicenseError.Error
	19, // 14: LicenseRequest.ClientId:type_name -> ClientIdentification
	54, // 15: LicenseRequest.ContentId:type_name -> LicenseRequest.ContentIdentification
	10, // 16: LicenseRequest.Type:type_name -> LicenseRequest.RequestType
	1,  // 17: LicenseRequest.ProtocolVersion:type_name -> ProtocolVersion
	23, // 18: LicenseRequest.EncryptedClientId:type_name -> EncryptedClientIdentification
	19, // 19: LicenseRequestRaw.ClientId:type_name -> ClientIdentification
	58, // 20: LicenseRequestRaw.ContentId:type_name -> LicenseRequestRaw.ContentIdentification
	11, // 21: LicenseRequestRaw.Type:type_name -> LicenseRequestRaw.RequestType
	1,  // 22: LicenseRequestRaw.ProtocolVersion:type_name -> ProtocolVersion
	23, // 23: LicenseRequestRaw.EncryptedClientId:type_name -> EncryptedClientIdentification
	12, // 24: ProvisionedDeviceInfo.SecurityLevel:type_name -> ProvisionedDeviceInfo.WvSecurityLevel
	23, // 25: RemoteAttestation.Certificate:type_name -> EncryptedClientIdentification
	20, // 26: SignedDeviceCertificate._DeviceCertificate:type_name -> DeviceCertificate
	37, // 27: SignedDeviceCertificate.Signer:type_name -> SignedDeviceCertificate
	13, // 28: SignedMessage.Type:type_name -> SignedMessage.MessageType
	33, // 29: SignedMessage.RemoteAttestation:type_name -> RemoteAttestation
	14, // 30: WidevineCencHeader.algorithm:type_name -> WidevineCencHeader.Algorithm
	15, // 31: SignedLicenseRequest.Type:type_name -> SignedLicenseRequest.MessageType
	27, // 32: SignedLicenseRequest.Msg:type_name -> LicenseRequest
	33, // 33: SignedLicenseRequest.RemoteAttestation:type_name -> RemoteAttestation
	16, // 34: SignedLicenseRequestRaw.Type:type_name -> SignedLicenseRequestRaw.MessageType
	28, // 35: SignedLicenseRequestRaw.Msg:type_name -> LicenseRequestRaw
	33, // 36: SignedLicenseRequestRaw.RemoteAttestation:type_name -> RemoteAttestation
	17, // 37: SignedLicense.Type:type_name -> SignedLicense.MessageType
	25, // 38: SignedLicense.Msg:type_name -> License
	33, // 39: SignedLicense.RemoteAttestation:type_name -> RemoteAttestation
	18, // 40: SignedServiceCertificate.Type:type_name -> SignedServiceCertificate.MessageType
	37, // 41: SignedServiceCertificate.Msg:type_name -> SignedDeviceCertificate
	33, // 42: SignedServiceCertificate.RemoteAttestation:type_name -> RemoteAttestation
	62, // 43: FileHashes.signatures:type_name -> FileHashes.Signature
	3,  // 44: ClientIdentification.ClientCapabilities.MaxHdcpVersion:type_name -> ClientIdentification.ClientCapabilities.HdcpVersion
	6,  // 45: License.KeyContainer.Type:type_name -> License.KeyContainer.KeyType
	7,  // 46: License.KeyContainer.Level:type_name -> License.KeyContainer.SecurityLevel
	50, // 47: License.KeyContainer.RequiredProtection:type_name -> License.KeyContainer.OutputProtection
	50, // 48: License.KeyContainer.RequestedProtection:type_name -> License.KeyContainer.OutputProtection
	51, // 49: License.KeyContainer._KeyControl:type_name -> License.KeyContainer.KeyControl
	52, // 50: License.KeyContainer._OperatorSessionKeyPermissions:type_name -> License.KeyContainer.OperatorSessionKeyPermissions
	53, // 51: License.KeyContainer.VideoResolutionConstraints:type_name -> License.KeyContainer.VideoResolutionConstraint
	3,  // 52: License.KeyContainer.OutputProtection.Hdcp:type_name -> ClientIdentification.ClientCapabilities.HdcpVersion
	8,  // 53: License.KeyContainer.OutputProtection.CgmsFlags:type_name -> License.KeyContainer.OutputProtection.CGMS
	50, // 54: License.KeyContainer.VideoResolutionConstraint.RequiredProtection:type_name -> License.KeyContainer.OutputProtection
	55, // 55: LicenseRequest.ContentIdentification.CencId:type_name -> LicenseRequest.ContentIdentification.CENC
	56, // 56: LicenseRequest.ContentIdentification.WebmId:type_name -> LicenseRequest.ContentIdentification.WebM
	57, // 57: LicenseRequest.ContentIdentification.License:type_name -> LicenseRequest.ContentIdentification.ExistingLicense
	40, // 58: LicenseRequest.ContentIdentification.CENC.Pssh:type_name -> WidevineCencHeader
	0,  // 59: LicenseRequest.ContentIdentification.CENC.LicenseType:type_name -> LicenseType
	0,  // 60: LicenseRequest.ContentIdentification.WebM.LicenseType:type_name -> LicenseType
	24, // 61: LicenseRequest.ContentIdentification.ExistingLicense.LicenseId:type_name -> LicenseIdentification
	59, // 62: LicenseRequestRaw.ContentIdentification.CencId:type_name -> LicenseRequestRaw.ContentIdentification.CENC
	60, // 63: LicenseRequestRaw.ContentIdentification.WebmId:type_name -> LicenseRequestRaw.ContentIdentification.WebM
	61, // 64: LicenseRequestRaw.ContentIdentification.License:type_name -> LicenseRequestRaw.ContentIdentification.ExistingLicense
	0,  // 65: LicenseRequestRaw.ContentIdentification.CENC.LicenseType:type_name -> LicenseType
	0,  // 66: LicenseRequestRaw.ContentIdentification.WebM.LicenseType:type_name -> LicenseType
	24, // 67: LicenseRequestRaw.ContentIdentification.ExistingLicense.LicenseId:type_name -> LicenseIdentification
	68, // [68:68] is the sub-list for method output_type
	68, // [68:68] is the sub-list for method input_type
	68, // [68:68] is the sub-list for extension type_name
	68, // [68:68] is the sub-list for extension extendee
	0,  // [0:68] is the sub-list for field type_name
}

func init() { file_wv_proto2_proto_init() }
func file_wv_proto2_proto_init() {
	if File_wv_proto2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wv_proto2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientIdentification); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceCertificate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceCertificateStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceCertificateStatusList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptedClientIdentification); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LicenseIdentification); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*License); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LicenseError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LicenseRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LicenseRequestRaw); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProvisionedDeviceInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProvisioningOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProvisioningRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProvisioningResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteAttestation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionInit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedCertificateStatusList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedDeviceCertificate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[19].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedProvisioningMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[20].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[21].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WidevineCencHeader); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[22].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedLicenseRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[23].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedLicenseRequestRaw); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[24].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedLicense); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[25].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedServiceCertificate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[26].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileHashes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[27].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientIdentification_NameValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[28].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientIdentification_ClientCapabilities); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[29].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*License_Policy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[30].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*License_KeyContainer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[31].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*License_KeyContainer_OutputProtection); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[32].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*License_KeyContainer_KeyControl); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[33].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*License_KeyContainer_OperatorSessionKeyPermissions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[34].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*License_KeyContainer_VideoResolutionConstraint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[35].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LicenseRequest_ContentIdentification); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[36].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LicenseRequest_ContentIdentification_CENC); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[37].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LicenseRequest_ContentIdentification_WebM); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[38].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LicenseRequest_ContentIdentification_ExistingLicense); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[39].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LicenseRequestRaw_ContentIdentification); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[40].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LicenseRequestRaw_ContentIdentification_CENC); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[41].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LicenseRequestRaw_ContentIdentification_WebM); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[42].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LicenseRequestRaw_ContentIdentification_ExistingLicense); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wv_proto2_proto_msgTypes[43].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileHashes_Signature); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wv_proto2_proto_rawDesc,
			NumEnums:      19,
			NumMessages:   44,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wv_proto2_proto_goTypes,
		DependencyIndexes: file_wv_proto2_proto_depIdxs,
		EnumInfos:         file_wv_proto2_proto_enumTypes,
		MessageInfos:      file_wv_proto2_proto_msgTypes,
	}.Build()
	File_wv_proto2_proto = out.File
	file_wv_proto2_proto_rawDesc = nil
	file_wv_proto2_proto_goTypes = nil
	file_wv_proto2_proto_depIdxs = nil
}
