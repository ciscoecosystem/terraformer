// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the schedule of a crawler using a cron expression.
func (c *Client) UpdateCrawlerSchedule(ctx context.Context, params *UpdateCrawlerScheduleInput, optFns ...func(*Options)) (*UpdateCrawlerScheduleOutput, error) {
	if params == nil {
		params = &UpdateCrawlerScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCrawlerSchedule", params, optFns, addOperationUpdateCrawlerScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCrawlerScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCrawlerScheduleInput struct {

	// The name of the crawler whose schedule to update.
	//
	// This member is required.
	CrawlerName *string

	// The updated cron expression used to specify the schedule (see Time-Based
	// Schedules for Jobs and Crawlers
	// (https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html).
	// For example, to run something every day at 12:15 UTC, you would specify: cron(15
	// 12 * * ? *).
	Schedule *string
}

type UpdateCrawlerScheduleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateCrawlerScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateCrawlerSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateCrawlerSchedule{}, middleware.After)
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
	if err = addOpUpdateCrawlerScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCrawlerSchedule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCrawlerSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "UpdateCrawlerSchedule",
	}
}