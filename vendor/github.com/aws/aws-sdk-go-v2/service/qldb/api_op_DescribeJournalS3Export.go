// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about a journal export job, including the ledger name,
// export ID, when it was created, current status, and its start and end time
// export parameters. This action does not return any expired export jobs. For more
// information, see Export Job Expiration
// (https://docs.aws.amazon.com/qldb/latest/developerguide/export-journal.request.html#export-journal.request.expiration)
// in the Amazon QLDB Developer Guide. If the export job with the given ExportId
// doesn't exist, then throws ResourceNotFoundException. If the ledger with the
// given Name doesn't exist, then throws ResourceNotFoundException.
func (c *Client) DescribeJournalS3Export(ctx context.Context, params *DescribeJournalS3ExportInput, optFns ...func(*Options)) (*DescribeJournalS3ExportOutput, error) {
	if params == nil {
		params = &DescribeJournalS3ExportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeJournalS3Export", params, optFns, addOperationDescribeJournalS3ExportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeJournalS3ExportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeJournalS3ExportInput struct {

	// The unique ID of the journal export job that you want to describe.
	//
	// This member is required.
	ExportId *string

	// The name of the ledger.
	//
	// This member is required.
	Name *string
}

type DescribeJournalS3ExportOutput struct {

	// Information about the journal export job returned by a DescribeJournalS3Export
	// request.
	//
	// This member is required.
	ExportDescription *types.JournalS3ExportDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeJournalS3ExportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeJournalS3Export{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeJournalS3Export{}, middleware.After)
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
	if err = addOpDescribeJournalS3ExportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeJournalS3Export(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeJournalS3Export(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "DescribeJournalS3Export",
	}
}