// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancels (stops) a task run. Machine learning task runs are asynchronous tasks
// that AWS Glue runs on your behalf as part of various machine learning workflows.
// You can cancel a machine learning task run at any time by calling
// CancelMLTaskRun with a task run's parent transform's TransformID and the task
// run's TaskRunId.
func (c *Client) CancelMLTaskRun(ctx context.Context, params *CancelMLTaskRunInput, optFns ...func(*Options)) (*CancelMLTaskRunOutput, error) {
	if params == nil {
		params = &CancelMLTaskRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelMLTaskRun", params, optFns, addOperationCancelMLTaskRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelMLTaskRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelMLTaskRunInput struct {

	// A unique identifier for the task run.
	//
	// This member is required.
	TaskRunId *string

	// The unique identifier of the machine learning transform.
	//
	// This member is required.
	TransformId *string
}

type CancelMLTaskRunOutput struct {

	// The status for this run.
	Status types.TaskStatusType

	// The unique identifier for the task run.
	TaskRunId *string

	// The unique identifier of the machine learning transform.
	TransformId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCancelMLTaskRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCancelMLTaskRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCancelMLTaskRun{}, middleware.After)
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
	if err = addOpCancelMLTaskRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelMLTaskRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelMLTaskRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "CancelMLTaskRun",
	}
}
