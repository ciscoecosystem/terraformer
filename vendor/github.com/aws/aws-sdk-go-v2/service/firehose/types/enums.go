// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type CompressionFormat string

// Enum values for CompressionFormat
const (
	CompressionFormatUncompressed CompressionFormat = "UNCOMPRESSED"
	CompressionFormatGzip         CompressionFormat = "GZIP"
	CompressionFormatZip          CompressionFormat = "ZIP"
	CompressionFormatSnappy       CompressionFormat = "Snappy"
	CompressionFormatHadoopSnappy CompressionFormat = "HADOOP_SNAPPY"
)

// Values returns all known values for CompressionFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CompressionFormat) Values() []CompressionFormat {
	return []CompressionFormat{
		"UNCOMPRESSED",
		"GZIP",
		"ZIP",
		"Snappy",
		"HADOOP_SNAPPY",
	}
}

type ContentEncoding string

// Enum values for ContentEncoding
const (
	ContentEncodingNone ContentEncoding = "NONE"
	ContentEncodingGzip ContentEncoding = "GZIP"
)

// Values returns all known values for ContentEncoding. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ContentEncoding) Values() []ContentEncoding {
	return []ContentEncoding{
		"NONE",
		"GZIP",
	}
}

type DeliveryStreamEncryptionStatus string

// Enum values for DeliveryStreamEncryptionStatus
const (
	DeliveryStreamEncryptionStatusEnabled         DeliveryStreamEncryptionStatus = "ENABLED"
	DeliveryStreamEncryptionStatusEnabling        DeliveryStreamEncryptionStatus = "ENABLING"
	DeliveryStreamEncryptionStatusEnablingFailed  DeliveryStreamEncryptionStatus = "ENABLING_FAILED"
	DeliveryStreamEncryptionStatusDisabled        DeliveryStreamEncryptionStatus = "DISABLED"
	DeliveryStreamEncryptionStatusDisabling       DeliveryStreamEncryptionStatus = "DISABLING"
	DeliveryStreamEncryptionStatusDisablingFailed DeliveryStreamEncryptionStatus = "DISABLING_FAILED"
)

// Values returns all known values for DeliveryStreamEncryptionStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (DeliveryStreamEncryptionStatus) Values() []DeliveryStreamEncryptionStatus {
	return []DeliveryStreamEncryptionStatus{
		"ENABLED",
		"ENABLING",
		"ENABLING_FAILED",
		"DISABLED",
		"DISABLING",
		"DISABLING_FAILED",
	}
}

type DeliveryStreamFailureType string

// Enum values for DeliveryStreamFailureType
const (
	DeliveryStreamFailureTypeRetireKmsGrantFailed      DeliveryStreamFailureType = "RETIRE_KMS_GRANT_FAILED"
	DeliveryStreamFailureTypeCreateKmsGrantFailed      DeliveryStreamFailureType = "CREATE_KMS_GRANT_FAILED"
	DeliveryStreamFailureTypeKmsAccessDenied           DeliveryStreamFailureType = "KMS_ACCESS_DENIED"
	DeliveryStreamFailureTypeDisabledKmsKey            DeliveryStreamFailureType = "DISABLED_KMS_KEY"
	DeliveryStreamFailureTypeInvalidKmsKey             DeliveryStreamFailureType = "INVALID_KMS_KEY"
	DeliveryStreamFailureTypeKmsKeyNotFound            DeliveryStreamFailureType = "KMS_KEY_NOT_FOUND"
	DeliveryStreamFailureTypeKmsOptInRequired          DeliveryStreamFailureType = "KMS_OPT_IN_REQUIRED"
	DeliveryStreamFailureTypeCreateEniFailed           DeliveryStreamFailureType = "CREATE_ENI_FAILED"
	DeliveryStreamFailureTypeDeleteEniFailed           DeliveryStreamFailureType = "DELETE_ENI_FAILED"
	DeliveryStreamFailureTypeSubnetNotFound            DeliveryStreamFailureType = "SUBNET_NOT_FOUND"
	DeliveryStreamFailureTypeSecurityGroupNotFound     DeliveryStreamFailureType = "SECURITY_GROUP_NOT_FOUND"
	DeliveryStreamFailureTypeEniAccessDenied           DeliveryStreamFailureType = "ENI_ACCESS_DENIED"
	DeliveryStreamFailureTypeSubnetAccessDenied        DeliveryStreamFailureType = "SUBNET_ACCESS_DENIED"
	DeliveryStreamFailureTypeSecurityGroupAccessDenied DeliveryStreamFailureType = "SECURITY_GROUP_ACCESS_DENIED"
	DeliveryStreamFailureTypeUnknownError              DeliveryStreamFailureType = "UNKNOWN_ERROR"
)

// Values returns all known values for DeliveryStreamFailureType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeliveryStreamFailureType) Values() []DeliveryStreamFailureType {
	return []DeliveryStreamFailureType{
		"RETIRE_KMS_GRANT_FAILED",
		"CREATE_KMS_GRANT_FAILED",
		"KMS_ACCESS_DENIED",
		"DISABLED_KMS_KEY",
		"INVALID_KMS_KEY",
		"KMS_KEY_NOT_FOUND",
		"KMS_OPT_IN_REQUIRED",
		"CREATE_ENI_FAILED",
		"DELETE_ENI_FAILED",
		"SUBNET_NOT_FOUND",
		"SECURITY_GROUP_NOT_FOUND",
		"ENI_ACCESS_DENIED",
		"SUBNET_ACCESS_DENIED",
		"SECURITY_GROUP_ACCESS_DENIED",
		"UNKNOWN_ERROR",
	}
}

type DeliveryStreamStatus string

// Enum values for DeliveryStreamStatus
const (
	DeliveryStreamStatusCreating       DeliveryStreamStatus = "CREATING"
	DeliveryStreamStatusCreatingFailed DeliveryStreamStatus = "CREATING_FAILED"
	DeliveryStreamStatusDeleting       DeliveryStreamStatus = "DELETING"
	DeliveryStreamStatusDeletingFailed DeliveryStreamStatus = "DELETING_FAILED"
	DeliveryStreamStatusActive         DeliveryStreamStatus = "ACTIVE"
)

// Values returns all known values for DeliveryStreamStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeliveryStreamStatus) Values() []DeliveryStreamStatus {
	return []DeliveryStreamStatus{
		"CREATING",
		"CREATING_FAILED",
		"DELETING",
		"DELETING_FAILED",
		"ACTIVE",
	}
}

type DeliveryStreamType string

// Enum values for DeliveryStreamType
const (
	DeliveryStreamTypeDirectPut             DeliveryStreamType = "DirectPut"
	DeliveryStreamTypeKinesisStreamAsSource DeliveryStreamType = "KinesisStreamAsSource"
)

// Values returns all known values for DeliveryStreamType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeliveryStreamType) Values() []DeliveryStreamType {
	return []DeliveryStreamType{
		"DirectPut",
		"KinesisStreamAsSource",
	}
}

type ElasticsearchIndexRotationPeriod string

// Enum values for ElasticsearchIndexRotationPeriod
const (
	ElasticsearchIndexRotationPeriodNoRotation ElasticsearchIndexRotationPeriod = "NoRotation"
	ElasticsearchIndexRotationPeriodOneHour    ElasticsearchIndexRotationPeriod = "OneHour"
	ElasticsearchIndexRotationPeriodOneDay     ElasticsearchIndexRotationPeriod = "OneDay"
	ElasticsearchIndexRotationPeriodOneWeek    ElasticsearchIndexRotationPeriod = "OneWeek"
	ElasticsearchIndexRotationPeriodOneMonth   ElasticsearchIndexRotationPeriod = "OneMonth"
)

// Values returns all known values for ElasticsearchIndexRotationPeriod. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ElasticsearchIndexRotationPeriod) Values() []ElasticsearchIndexRotationPeriod {
	return []ElasticsearchIndexRotationPeriod{
		"NoRotation",
		"OneHour",
		"OneDay",
		"OneWeek",
		"OneMonth",
	}
}

type ElasticsearchS3BackupMode string

// Enum values for ElasticsearchS3BackupMode
const (
	ElasticsearchS3BackupModeFailedDocumentsOnly ElasticsearchS3BackupMode = "FailedDocumentsOnly"
	ElasticsearchS3BackupModeAllDocuments        ElasticsearchS3BackupMode = "AllDocuments"
)

// Values returns all known values for ElasticsearchS3BackupMode. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ElasticsearchS3BackupMode) Values() []ElasticsearchS3BackupMode {
	return []ElasticsearchS3BackupMode{
		"FailedDocumentsOnly",
		"AllDocuments",
	}
}

type HECEndpointType string

// Enum values for HECEndpointType
const (
	HECEndpointTypeRaw   HECEndpointType = "Raw"
	HECEndpointTypeEvent HECEndpointType = "Event"
)

// Values returns all known values for HECEndpointType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (HECEndpointType) Values() []HECEndpointType {
	return []HECEndpointType{
		"Raw",
		"Event",
	}
}

type HttpEndpointS3BackupMode string

// Enum values for HttpEndpointS3BackupMode
const (
	HttpEndpointS3BackupModeFailedDataOnly HttpEndpointS3BackupMode = "FailedDataOnly"
	HttpEndpointS3BackupModeAllData        HttpEndpointS3BackupMode = "AllData"
)

// Values returns all known values for HttpEndpointS3BackupMode. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (HttpEndpointS3BackupMode) Values() []HttpEndpointS3BackupMode {
	return []HttpEndpointS3BackupMode{
		"FailedDataOnly",
		"AllData",
	}
}

type KeyType string

// Enum values for KeyType
const (
	KeyTypeAwsOwnedCmk        KeyType = "AWS_OWNED_CMK"
	KeyTypeCustomerManagedCmk KeyType = "CUSTOMER_MANAGED_CMK"
)

// Values returns all known values for KeyType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (KeyType) Values() []KeyType {
	return []KeyType{
		"AWS_OWNED_CMK",
		"CUSTOMER_MANAGED_CMK",
	}
}

type NoEncryptionConfig string

// Enum values for NoEncryptionConfig
const (
	NoEncryptionConfigNoEncryption NoEncryptionConfig = "NoEncryption"
)

// Values returns all known values for NoEncryptionConfig. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NoEncryptionConfig) Values() []NoEncryptionConfig {
	return []NoEncryptionConfig{
		"NoEncryption",
	}
}

type OrcCompression string

// Enum values for OrcCompression
const (
	OrcCompressionNone   OrcCompression = "NONE"
	OrcCompressionZlib   OrcCompression = "ZLIB"
	OrcCompressionSnappy OrcCompression = "SNAPPY"
)

// Values returns all known values for OrcCompression. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OrcCompression) Values() []OrcCompression {
	return []OrcCompression{
		"NONE",
		"ZLIB",
		"SNAPPY",
	}
}

type OrcFormatVersion string

// Enum values for OrcFormatVersion
const (
	OrcFormatVersionV011 OrcFormatVersion = "V0_11"
	OrcFormatVersionV012 OrcFormatVersion = "V0_12"
)

// Values returns all known values for OrcFormatVersion. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OrcFormatVersion) Values() []OrcFormatVersion {
	return []OrcFormatVersion{
		"V0_11",
		"V0_12",
	}
}

type ParquetCompression string

// Enum values for ParquetCompression
const (
	ParquetCompressionUncompressed ParquetCompression = "UNCOMPRESSED"
	ParquetCompressionGzip         ParquetCompression = "GZIP"
	ParquetCompressionSnappy       ParquetCompression = "SNAPPY"
)

// Values returns all known values for ParquetCompression. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ParquetCompression) Values() []ParquetCompression {
	return []ParquetCompression{
		"UNCOMPRESSED",
		"GZIP",
		"SNAPPY",
	}
}

type ParquetWriterVersion string

// Enum values for ParquetWriterVersion
const (
	ParquetWriterVersionV1 ParquetWriterVersion = "V1"
	ParquetWriterVersionV2 ParquetWriterVersion = "V2"
)

// Values returns all known values for ParquetWriterVersion. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ParquetWriterVersion) Values() []ParquetWriterVersion {
	return []ParquetWriterVersion{
		"V1",
		"V2",
	}
}

type ProcessorParameterName string

// Enum values for ProcessorParameterName
const (
	ProcessorParameterNameLambdaArn               ProcessorParameterName = "LambdaArn"
	ProcessorParameterNameLambdaNumberOfRetries   ProcessorParameterName = "NumberOfRetries"
	ProcessorParameterNameRoleArn                 ProcessorParameterName = "RoleArn"
	ProcessorParameterNameBufferSizeInMb          ProcessorParameterName = "BufferSizeInMBs"
	ProcessorParameterNameBufferIntervalInSeconds ProcessorParameterName = "BufferIntervalInSeconds"
)

// Values returns all known values for ProcessorParameterName. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ProcessorParameterName) Values() []ProcessorParameterName {
	return []ProcessorParameterName{
		"LambdaArn",
		"NumberOfRetries",
		"RoleArn",
		"BufferSizeInMBs",
		"BufferIntervalInSeconds",
	}
}

type ProcessorType string

// Enum values for ProcessorType
const (
	ProcessorTypeLambda ProcessorType = "Lambda"
)

// Values returns all known values for ProcessorType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ProcessorType) Values() []ProcessorType {
	return []ProcessorType{
		"Lambda",
	}
}

type RedshiftS3BackupMode string

// Enum values for RedshiftS3BackupMode
const (
	RedshiftS3BackupModeDisabled RedshiftS3BackupMode = "Disabled"
	RedshiftS3BackupModeEnabled  RedshiftS3BackupMode = "Enabled"
)

// Values returns all known values for RedshiftS3BackupMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RedshiftS3BackupMode) Values() []RedshiftS3BackupMode {
	return []RedshiftS3BackupMode{
		"Disabled",
		"Enabled",
	}
}

type S3BackupMode string

// Enum values for S3BackupMode
const (
	S3BackupModeDisabled S3BackupMode = "Disabled"
	S3BackupModeEnabled  S3BackupMode = "Enabled"
)

// Values returns all known values for S3BackupMode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (S3BackupMode) Values() []S3BackupMode {
	return []S3BackupMode{
		"Disabled",
		"Enabled",
	}
}

type SplunkS3BackupMode string

// Enum values for SplunkS3BackupMode
const (
	SplunkS3BackupModeFailedEventsOnly SplunkS3BackupMode = "FailedEventsOnly"
	SplunkS3BackupModeAllEvents        SplunkS3BackupMode = "AllEvents"
)

// Values returns all known values for SplunkS3BackupMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SplunkS3BackupMode) Values() []SplunkS3BackupMode {
	return []SplunkS3BackupMode{
		"FailedEventsOnly",
		"AllEvents",
	}
}
