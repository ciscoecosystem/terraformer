// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates a metric stream. Metric streams can automatically stream
// CloudWatch metrics to AWS destinations including Amazon S3 and to many
// third-party solutions. For more information, see  Using Metric Streams
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Metric-Streams.html).
// To create a metric stream, you must be logged on to an account that has the
// iam:PassRole permission and either the CloudWatchFullAccess policy or the
// cloudwatch:PutMetricStream permission. When you create or update a metric
// stream, you choose one of the following:
//
// * Stream metrics from all metric
// namespaces in the account.
//
// * Stream metrics from all metric namespaces in the
// account, except for the namespaces that you list in ExcludeFilters.
//
// * Stream
// metrics from only the metric namespaces that you list in IncludeFilters.
//
// When
// you use PutMetricStream to create a new metric stream, the stream is created in
// the running state. If you use it to update an existing stream, the state of the
// stream is not changed.
func (c *Client) PutMetricStream(ctx context.Context, params *PutMetricStreamInput, optFns ...func(*Options)) (*PutMetricStreamOutput, error) {
	if params == nil {
		params = &PutMetricStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutMetricStream", params, optFns, addOperationPutMetricStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutMetricStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutMetricStreamInput struct {

	// The ARN of the Amazon Kinesis Firehose delivery stream to use for this metric
	// stream. This Amazon Kinesis Firehose delivery stream must already exist and must
	// be in the same account as the metric stream.
	//
	// This member is required.
	FirehoseArn *string

	// If you are creating a new metric stream, this is the name for the new stream.
	// The name must be different than the names of other metric streams in this
	// account and Region. If you are updating a metric stream, specify the name of
	// that stream here. Valid characters are A-Z, a-z, 0-9, "-" and "_".
	//
	// This member is required.
	Name *string

	// The output format for the stream. Valid values are json and opentelemetry0.7.
	// For more information about metric stream output formats, see  Metric streams
	// output formats
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats.html).
	//
	// This member is required.
	OutputFormat types.MetricStreamOutputFormat

	// The ARN of an IAM role that this metric stream will use to access Amazon Kinesis
	// Firehose resources. This IAM role must already exist and must be in the same
	// account as the metric stream. This IAM role must include the following
	// permissions:
	//
	// * firehose:PutRecord
	//
	// * firehose:PutRecordBatch
	//
	// This member is required.
	RoleArn *string

	// If you specify this parameter, the stream sends metrics from all metric
	// namespaces except for the namespaces that you specify here. You cannot include
	// ExcludeFilters and IncludeFilters in the same operation.
	ExcludeFilters []types.MetricStreamFilter

	// If you specify this parameter, the stream sends only the metrics from the metric
	// namespaces that you specify here. You cannot include IncludeFilters and
	// ExcludeFilters in the same operation.
	IncludeFilters []types.MetricStreamFilter

	// A list of key-value pairs to associate with the metric stream. You can associate
	// as many as 50 tags with a metric stream. Tags can help you organize and
	// categorize your resources. You can also use them to scope user permissions by
	// granting a user permission to access or change only resources with certain tag
	// values.
	Tags []types.Tag
}

type PutMetricStreamOutput struct {

	// The ARN of the metric stream.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutMetricStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPutMetricStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPutMetricStream{}, middleware.After)
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
	if err = addOpPutMetricStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutMetricStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutMetricStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "PutMetricStream",
	}
}
