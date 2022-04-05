// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists completed exports within the past 90 days.
func (c *Client) ListExports(ctx context.Context, params *ListExportsInput, optFns ...func(*Options)) (*ListExportsOutput, error) {
	if params == nil {
		params = &ListExportsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListExports", params, optFns, addOperationListExportsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListExportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListExportsInput struct {

	// Maximum number of results to return per page.
	MaxResults *int32

	// An optional string that, if supplied, must be copied from the output of a
	// previous call to ListExports. When provided in this manner, the API fetches the
	// next page of results.
	NextToken *string

	// The Amazon Resource Name (ARN) associated with the exported table.
	TableArn *string
}

type ListExportsOutput struct {

	// A list of ExportSummary objects.
	ExportSummaries []types.ExportSummary

	// If this value is returned, there are additional results to be displayed. To
	// retrieve them, call ListExports again, with NextToken set to this value.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListExportsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListExports{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListExports{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListExports(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addValidateResponseChecksum(stack, options); err != nil {
		return err
	}
	if err = addAcceptEncodingGzip(stack, options); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// ListExportsAPIClient is a client that implements the ListExports operation.
type ListExportsAPIClient interface {
	ListExports(context.Context, *ListExportsInput, ...func(*Options)) (*ListExportsOutput, error)
}

var _ ListExportsAPIClient = (*Client)(nil)

// ListExportsPaginatorOptions is the paginator options for ListExports
type ListExportsPaginatorOptions struct {
	// Maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListExportsPaginator is a paginator for ListExports
type ListExportsPaginator struct {
	options   ListExportsPaginatorOptions
	client    ListExportsAPIClient
	params    *ListExportsInput
	nextToken *string
	firstPage bool
}

// NewListExportsPaginator returns a new ListExportsPaginator
func NewListExportsPaginator(client ListExportsAPIClient, params *ListExportsInput, optFns ...func(*ListExportsPaginatorOptions)) *ListExportsPaginator {
	if params == nil {
		params = &ListExportsInput{}
	}

	options := ListExportsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListExportsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListExportsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListExports page.
func (p *ListExportsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListExportsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListExports(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListExports(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "ListExports",
	}
}
