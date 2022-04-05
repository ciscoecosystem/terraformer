// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists your Device Defender audit listings.
func (c *Client) ListAuditSuppressions(ctx context.Context, params *ListAuditSuppressionsInput, optFns ...func(*Options)) (*ListAuditSuppressionsOutput, error) {
	if params == nil {
		params = &ListAuditSuppressionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAuditSuppressions", params, optFns, addOperationListAuditSuppressionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAuditSuppressionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAuditSuppressionsInput struct {

	// Determines whether suppressions are listed in ascending order by expiration date
	// or not. If parameter isn't provided, ascendingOrder=true.
	AscendingOrder bool

	// An audit check name. Checks must be enabled for your account. (Use
	// DescribeAccountAuditConfiguration to see the list of all checks, including those
	// that are enabled or use UpdateAccountAuditConfiguration to select which checks
	// are enabled.)
	CheckName *string

	// The maximum number of results to return at one time. The default is 25.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	// Information that identifies the noncompliant resource.
	ResourceIdentifier *types.ResourceIdentifier
}

type ListAuditSuppressionsOutput struct {

	// A token that can be used to retrieve the next set of results, or null if there
	// are no additional results.
	NextToken *string

	// List of audit suppressions.
	Suppressions []types.AuditSuppression

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAuditSuppressionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAuditSuppressions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAuditSuppressions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAuditSuppressions(options.Region), middleware.Before); err != nil {
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

// ListAuditSuppressionsAPIClient is a client that implements the
// ListAuditSuppressions operation.
type ListAuditSuppressionsAPIClient interface {
	ListAuditSuppressions(context.Context, *ListAuditSuppressionsInput, ...func(*Options)) (*ListAuditSuppressionsOutput, error)
}

var _ ListAuditSuppressionsAPIClient = (*Client)(nil)

// ListAuditSuppressionsPaginatorOptions is the paginator options for
// ListAuditSuppressions
type ListAuditSuppressionsPaginatorOptions struct {
	// The maximum number of results to return at one time. The default is 25.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAuditSuppressionsPaginator is a paginator for ListAuditSuppressions
type ListAuditSuppressionsPaginator struct {
	options   ListAuditSuppressionsPaginatorOptions
	client    ListAuditSuppressionsAPIClient
	params    *ListAuditSuppressionsInput
	nextToken *string
	firstPage bool
}

// NewListAuditSuppressionsPaginator returns a new ListAuditSuppressionsPaginator
func NewListAuditSuppressionsPaginator(client ListAuditSuppressionsAPIClient, params *ListAuditSuppressionsInput, optFns ...func(*ListAuditSuppressionsPaginatorOptions)) *ListAuditSuppressionsPaginator {
	if params == nil {
		params = &ListAuditSuppressionsInput{}
	}

	options := ListAuditSuppressionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAuditSuppressionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAuditSuppressionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListAuditSuppressions page.
func (p *ListAuditSuppressionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAuditSuppressionsOutput, error) {
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

	result, err := p.client.ListAuditSuppressions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAuditSuppressions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListAuditSuppressions",
	}
}
