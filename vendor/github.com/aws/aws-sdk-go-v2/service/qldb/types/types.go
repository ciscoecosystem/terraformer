// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// The information about an Amazon QLDB journal stream, including the Amazon
// Resource Name (ARN), stream name, creation time, current status, and the
// parameters of your original stream creation request.
type JournalKinesisStreamDescription struct {

	// The configuration settings of the Amazon Kinesis Data Streams destination for
	// your QLDB journal stream.
	//
	// This member is required.
	KinesisConfiguration *KinesisConfiguration

	// The name of the ledger.
	//
	// This member is required.
	LedgerName *string

	// The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions for
	// a journal stream to write data records to a Kinesis Data Streams resource.
	//
	// This member is required.
	RoleArn *string

	// The current state of the QLDB journal stream.
	//
	// This member is required.
	Status StreamStatus

	// The unique ID that QLDB assigns to each QLDB journal stream.
	//
	// This member is required.
	StreamId *string

	// The user-defined name of the QLDB journal stream.
	//
	// This member is required.
	StreamName *string

	// The Amazon Resource Name (ARN) of the QLDB journal stream.
	Arn *string

	// The date and time, in epoch time format, when the QLDB journal stream was
	// created. (Epoch time format is the number of seconds elapsed since 12:00:00 AM
	// January 1, 1970 UTC.)
	CreationTime *time.Time

	// The error message that describes the reason that a stream has a status of
	// IMPAIRED or FAILED. This is not applicable to streams that have other status
	// values.
	ErrorCause ErrorCause

	// The exclusive date and time that specifies when the stream ends. If this
	// parameter is blank, the stream runs indefinitely until you cancel it.
	ExclusiveEndTime *time.Time

	// The inclusive start date and time from which to start streaming journal data.
	InclusiveStartTime *time.Time
}

// The information about a journal export job, including the ledger name, export
// ID, when it was created, current status, and its start and end time export
// parameters.
type JournalS3ExportDescription struct {

	// The exclusive end date and time for the range of journal contents that are
	// specified in the original export request.
	//
	// This member is required.
	ExclusiveEndTime *time.Time

	// The date and time, in epoch time format, when the export job was created. (Epoch
	// time format is the number of seconds elapsed since 12:00:00 AM January 1, 1970
	// UTC.)
	//
	// This member is required.
	ExportCreationTime *time.Time

	// The unique ID of the journal export job.
	//
	// This member is required.
	ExportId *string

	// The inclusive start date and time for the range of journal contents that are
	// specified in the original export request.
	//
	// This member is required.
	InclusiveStartTime *time.Time

	// The name of the ledger.
	//
	// This member is required.
	LedgerName *string

	// The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions for
	// a journal export job to do the following:
	//
	// * Write objects into your Amazon
	// Simple Storage Service (Amazon S3) bucket.
	//
	// * (Optional) Use your customer
	// master key (CMK) in AWS Key Management Service (AWS KMS) for server-side
	// encryption of your exported data.
	//
	// This member is required.
	RoleArn *string

	// The Amazon Simple Storage Service (Amazon S3) bucket location in which a journal
	// export job writes the journal contents.
	//
	// This member is required.
	S3ExportConfiguration *S3ExportConfiguration

	// The current state of the journal export job.
	//
	// This member is required.
	Status ExportStatus
}

// The configuration settings of the Amazon Kinesis Data Streams destination for
// your Amazon QLDB journal stream.
type KinesisConfiguration struct {

	// The Amazon Resource Name (ARN) of the Kinesis data stream resource.
	//
	// This member is required.
	StreamArn *string

	// Enables QLDB to publish multiple data records in a single Kinesis Data Streams
	// record. To learn more, see KPL Key Concepts
	// (https://docs.aws.amazon.com/streams/latest/dev/kinesis-kpl-concepts.html) in
	// the Amazon Kinesis Data Streams Developer Guide.
	AggregationEnabled *bool
}

// Information about a ledger, including its name, state, and when it was created.
type LedgerSummary struct {

	// The date and time, in epoch time format, when the ledger was created. (Epoch
	// time format is the number of seconds elapsed since 12:00:00 AM January 1, 1970
	// UTC.)
	CreationDateTime *time.Time

	// The name of the ledger.
	Name *string

	// The current status of the ledger.
	State LedgerState
}

// The encryption settings that are used by a journal export job to write data in
// an Amazon Simple Storage Service (Amazon S3) bucket.
type S3EncryptionConfiguration struct {

	// The Amazon S3 object encryption type. To learn more about server-side encryption
	// options in Amazon S3, see Protecting Data Using Server-Side Encryption
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/serv-side-encryption.html) in
	// the Amazon S3 Developer Guide.
	//
	// This member is required.
	ObjectEncryptionType S3ObjectEncryptionType

	// The Amazon Resource Name (ARN) for a symmetric customer master key (CMK) in AWS
	// Key Management Service (AWS KMS). Amazon QLDB does not support asymmetric CMKs.
	// You must provide a KmsKeyArn if you specify SSE_KMS as the ObjectEncryptionType.
	// KmsKeyArn is not required if you specify SSE_S3 as the ObjectEncryptionType.
	KmsKeyArn *string
}

// The Amazon Simple Storage Service (Amazon S3) bucket location in which a journal
// export job writes the journal contents.
type S3ExportConfiguration struct {

	// The Amazon S3 bucket name in which a journal export job writes the journal
	// contents. The bucket name must comply with the Amazon S3 bucket naming
	// conventions. For more information, see Bucket Restrictions and Limitations
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html) in the
	// Amazon S3 Developer Guide.
	//
	// This member is required.
	Bucket *string

	// The encryption settings that are used by a journal export job to write data in
	// an Amazon S3 bucket.
	//
	// This member is required.
	EncryptionConfiguration *S3EncryptionConfiguration

	// The prefix for the Amazon S3 bucket in which a journal export job writes the
	// journal contents. The prefix must comply with Amazon S3 key naming rules and
	// restrictions. For more information, see Object Key and Metadata
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html) in the
	// Amazon S3 Developer Guide. The following are examples of valid Prefix values:
	//
	// *
	// JournalExports-ForMyLedger/Testing/
	//
	// * JournalExports
	//
	// * My:Tests/
	//
	// This member is required.
	Prefix *string
}

// A structure that can contain a value in multiple encoding formats.
type ValueHolder struct {

	// An Amazon Ion plaintext value contained in a ValueHolder structure.
	IonText *string
}
