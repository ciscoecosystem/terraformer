// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of report groups that are shared with other AWS accounts or users.
func (c *Client) ListSharedReportGroups(ctx context.Context, params *ListSharedReportGroupsInput, optFns ...func(*Options)) (*ListSharedReportGroupsOutput, error) {
	if params == nil {
		params = &ListSharedReportGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSharedReportGroups", params, optFns, addOperationListSharedReportGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSharedReportGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSharedReportGroupsInput struct {

	// The maximum number of paginated shared report groups per response. Use nextToken
	// to iterate pages in the list of returned ReportGroup objects. The default value
	// is 100.
	MaxResults *int32

	// During a previous call, the maximum number of items that can be returned is the
	// value specified in maxResults. If there more items in the list, then a unique
	// string called a nextToken is returned. To get the next batch of items in the
	// list, call this operation again, adding the next token to the call. To get all
	// of the items in the list, keep calling this operation with each subsequent next
	// token that is returned, until no more next tokens are returned.
	NextToken *string

	// The criterion to be used to list report groups shared with the current AWS
	// account or user. Valid values include:
	//
	// * ARN: List based on the ARN.
	//
	// *
	// MODIFIED_TIME: List based on when information about the shared report group was
	// last changed.
	SortBy types.SharedResourceSortByType

	// The order in which to list shared report groups. Valid values include:
	//
	// *
	// ASCENDING: List in ascending order.
	//
	// * DESCENDING: List in descending order.
	SortOrder types.SortOrderType
}

type ListSharedReportGroupsOutput struct {

	// During a previous call, the maximum number of items that can be returned is the
	// value specified in maxResults. If there more items in the list, then a unique
	// string called a nextToken is returned. To get the next batch of items in the
	// list, call this operation again, adding the next token to the call. To get all
	// of the items in the list, keep calling this operation with each subsequent next
	// token that is returned, until no more next tokens are returned.
	NextToken *string

	// The list of ARNs for the report groups shared with the current AWS account or
	// user.
	ReportGroups []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListSharedReportGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSharedReportGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSharedReportGroups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSharedReportGroups(options.Region), middleware.Before); err != nil {
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

// ListSharedReportGroupsAPIClient is a client that implements the
// ListSharedReportGroups operation.
type ListSharedReportGroupsAPIClient interface {
	ListSharedReportGroups(context.Context, *ListSharedReportGroupsInput, ...func(*Options)) (*ListSharedReportGroupsOutput, error)
}

var _ ListSharedReportGroupsAPIClient = (*Client)(nil)

// ListSharedReportGroupsPaginatorOptions is the paginator options for
// ListSharedReportGroups
type ListSharedReportGroupsPaginatorOptions struct {
	// The maximum number of paginated shared report groups per response. Use nextToken
	// to iterate pages in the list of returned ReportGroup objects. The default value
	// is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSharedReportGroupsPaginator is a paginator for ListSharedReportGroups
type ListSharedReportGroupsPaginator struct {
	options   ListSharedReportGroupsPaginatorOptions
	client    ListSharedReportGroupsAPIClient
	params    *ListSharedReportGroupsInput
	nextToken *string
	firstPage bool
}

// NewListSharedReportGroupsPaginator returns a new ListSharedReportGroupsPaginator
func NewListSharedReportGroupsPaginator(client ListSharedReportGroupsAPIClient, params *ListSharedReportGroupsInput, optFns ...func(*ListSharedReportGroupsPaginatorOptions)) *ListSharedReportGroupsPaginator {
	if params == nil {
		params = &ListSharedReportGroupsInput{}
	}

	options := ListSharedReportGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSharedReportGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSharedReportGroupsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListSharedReportGroups page.
func (p *ListSharedReportGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSharedReportGroupsOutput, error) {
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

	result, err := p.client.ListSharedReportGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSharedReportGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "ListSharedReportGroups",
	}
}
