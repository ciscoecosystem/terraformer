// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the settings that specify delivery of log files. Changes to a trail do
// not require stopping the CloudTrail service. Use this action to designate an
// existing bucket for log delivery. If the existing bucket has previously been a
// target for CloudTrail log files, an IAM policy exists for the bucket.
// UpdateTrail must be called from the region in which the trail was created;
// otherwise, an InvalidHomeRegionException is thrown.
func (c *Client) UpdateTrail(ctx context.Context, params *UpdateTrailInput, optFns ...func(*Options)) (*UpdateTrailOutput, error) {
	if params == nil {
		params = &UpdateTrailInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTrail", params, optFns, addOperationUpdateTrailMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTrailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Specifies settings to update for the trail.
type UpdateTrailInput struct {

	// Specifies the name of the trail or trail ARN. If Name is a trail name, the
	// string must meet the following requirements:
	//
	// * Contain only ASCII letters (a-z,
	// A-Z), numbers (0-9), periods (.), underscores (_), or dashes (-)
	//
	// * Start with a
	// letter or number, and end with a letter or number
	//
	// * Be between 3 and 128
	// characters
	//
	// * Have no adjacent periods, underscores or dashes. Names like
	// my-_namespace and my--namespace are invalid.
	//
	// * Not be in IP address format (for
	// example, 192.168.5.4)
	//
	// If Name is a trail ARN, it must be in the format:
	// arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail
	//
	// This member is required.
	Name *string

	// Specifies a log group name using an Amazon Resource Name (ARN), a unique
	// identifier that represents the log group to which CloudTrail logs will be
	// delivered. Not required unless you specify CloudWatchLogsRoleArn.
	CloudWatchLogsLogGroupArn *string

	// Specifies the role for the CloudWatch Logs endpoint to assume to write to a
	// user's log group.
	CloudWatchLogsRoleArn *string

	// Specifies whether log file validation is enabled. The default is false. When you
	// disable log file integrity validation, the chain of digest files is broken after
	// one hour. CloudTrail will not create digest files for log files that were
	// delivered during a period in which log file integrity validation was disabled.
	// For example, if you enable log file integrity validation at noon on January 1,
	// disable it at noon on January 2, and re-enable it at noon on January 10, digest
	// files will not be created for the log files delivered from noon on January 2 to
	// noon on January 10. The same applies whenever you stop CloudTrail logging or
	// delete a trail.
	EnableLogFileValidation *bool

	// Specifies whether the trail is publishing events from global services such as
	// IAM to the log files.
	IncludeGlobalServiceEvents *bool

	// Specifies whether the trail applies only to the current region or to all
	// regions. The default is false. If the trail exists only in the current region
	// and this value is set to true, shadow trails (replications of the trail) will be
	// created in the other regions. If the trail exists in all regions and this value
	// is set to false, the trail will remain in the region where it was created, and
	// its shadow trails in other regions will be deleted. As a best practice, consider
	// using trails that log events in all regions.
	IsMultiRegionTrail *bool

	// Specifies whether the trail is applied to all accounts in an organization in AWS
	// Organizations, or only for the current AWS account. The default is false, and
	// cannot be true unless the call is made on behalf of an AWS account that is the
	// master account for an organization in AWS Organizations. If the trail is not an
	// organization trail and this is set to true, the trail will be created in all AWS
	// accounts that belong to the organization. If the trail is an organization trail
	// and this is set to false, the trail will remain in the current AWS account but
	// be deleted from all member accounts in the organization.
	IsOrganizationTrail *bool

	// Specifies the KMS key ID to use to encrypt the logs delivered by CloudTrail. The
	// value can be an alias name prefixed by "alias/", a fully specified ARN to an
	// alias, a fully specified ARN to a key, or a globally unique identifier.
	// Examples:
	//
	// * alias/MyAliasName
	//
	// *
	// arn:aws:kms:us-east-2:123456789012:alias/MyAliasName
	//
	// *
	// arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012
	//
	// *
	// 12345678-1234-1234-1234-123456789012
	KmsKeyId *string

	// Specifies the name of the Amazon S3 bucket designated for publishing log files.
	// See Amazon S3 Bucket Naming Requirements
	// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/create_trail_naming_policy.html).
	S3BucketName *string

	// Specifies the Amazon S3 key prefix that comes after the name of the bucket you
	// have designated for log file delivery. For more information, see Finding Your
	// CloudTrail Log Files
	// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-find-log-files.html).
	// The maximum length is 200 characters.
	S3KeyPrefix *string

	// Specifies the name of the Amazon SNS topic defined for notification of log file
	// delivery. The maximum length is 256 characters.
	SnsTopicName *string
}

// Returns the objects or data listed below if successful. Otherwise, returns an
// error.
type UpdateTrailOutput struct {

	// Specifies the Amazon Resource Name (ARN) of the log group to which CloudTrail
	// logs will be delivered.
	CloudWatchLogsLogGroupArn *string

	// Specifies the role for the CloudWatch Logs endpoint to assume to write to a
	// user's log group.
	CloudWatchLogsRoleArn *string

	// Specifies whether the trail is publishing events from global services such as
	// IAM to the log files.
	IncludeGlobalServiceEvents *bool

	// Specifies whether the trail exists in one region or in all regions.
	IsMultiRegionTrail *bool

	// Specifies whether the trail is an organization trail.
	IsOrganizationTrail *bool

	// Specifies the KMS key ID that encrypts the logs delivered by CloudTrail. The
	// value is a fully specified ARN to a KMS key in the format:
	// arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012
	KmsKeyId *string

	// Specifies whether log file integrity validation is enabled.
	LogFileValidationEnabled *bool

	// Specifies the name of the trail.
	Name *string

	// Specifies the name of the Amazon S3 bucket designated for publishing log files.
	S3BucketName *string

	// Specifies the Amazon S3 key prefix that comes after the name of the bucket you
	// have designated for log file delivery. For more information, see Finding Your
	// CloudTrail Log Files
	// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-find-log-files.html).
	S3KeyPrefix *string

	// Specifies the ARN of the Amazon SNS topic that CloudTrail uses to send
	// notifications when log files are delivered. The format of a topic ARN is:
	// arn:aws:sns:us-east-2:123456789012:MyTopic
	SnsTopicARN *string

	// This field is no longer in use. Use SnsTopicARN.
	//
	// Deprecated: This member has been deprecated.
	SnsTopicName *string

	// Specifies the ARN of the trail that was updated. The format of a trail ARN is:
	// arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail
	TrailARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateTrailMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateTrail{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateTrail{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateTrailValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTrail(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateTrail(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "UpdateTrail",
	}
}
