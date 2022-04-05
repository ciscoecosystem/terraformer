// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about the products for the specified portfolio or all products.
func (c *Client) SearchProductsAsAdmin(ctx context.Context, params *SearchProductsAsAdminInput, optFns ...func(*Options)) (*SearchProductsAsAdminOutput, error) {
	if params == nil {
		params = &SearchProductsAsAdminInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchProductsAsAdmin", params, optFns, addOperationSearchProductsAsAdminMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchProductsAsAdminOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchProductsAsAdminInput struct {

	// The language code.
	//
	// * en - English (default)
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string

	// The search filters. If no search filters are specified, the output includes all
	// products to which the administrator has access.
	Filters map[string][]string

	// The maximum number of items to return with this call.
	PageSize int32

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string

	// The portfolio identifier.
	PortfolioId *string

	// Access level of the source of the product.
	ProductSource types.ProductSource

	// The sort field. If no value is specified, the results are not sorted.
	SortBy types.ProductViewSortBy

	// The sort order. If no value is specified, the results are not sorted.
	SortOrder types.SortOrder
}

type SearchProductsAsAdminOutput struct {

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string

	// Information about the product views.
	ProductViewDetails []types.ProductViewDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSearchProductsAsAdminMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchProductsAsAdmin{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchProductsAsAdmin{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchProductsAsAdmin(options.Region), middleware.Before); err != nil {
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

// SearchProductsAsAdminAPIClient is a client that implements the
// SearchProductsAsAdmin operation.
type SearchProductsAsAdminAPIClient interface {
	SearchProductsAsAdmin(context.Context, *SearchProductsAsAdminInput, ...func(*Options)) (*SearchProductsAsAdminOutput, error)
}

var _ SearchProductsAsAdminAPIClient = (*Client)(nil)

// SearchProductsAsAdminPaginatorOptions is the paginator options for
// SearchProductsAsAdmin
type SearchProductsAsAdminPaginatorOptions struct {
	// The maximum number of items to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchProductsAsAdminPaginator is a paginator for SearchProductsAsAdmin
type SearchProductsAsAdminPaginator struct {
	options   SearchProductsAsAdminPaginatorOptions
	client    SearchProductsAsAdminAPIClient
	params    *SearchProductsAsAdminInput
	nextToken *string
	firstPage bool
}

// NewSearchProductsAsAdminPaginator returns a new SearchProductsAsAdminPaginator
func NewSearchProductsAsAdminPaginator(client SearchProductsAsAdminAPIClient, params *SearchProductsAsAdminInput, optFns ...func(*SearchProductsAsAdminPaginatorOptions)) *SearchProductsAsAdminPaginator {
	if params == nil {
		params = &SearchProductsAsAdminInput{}
	}

	options := SearchProductsAsAdminPaginatorOptions{}
	if params.PageSize != 0 {
		options.Limit = params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchProductsAsAdminPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchProductsAsAdminPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next SearchProductsAsAdmin page.
func (p *SearchProductsAsAdminPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchProductsAsAdminOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PageToken = p.nextToken

	params.PageSize = p.options.Limit

	result, err := p.client.SearchProductsAsAdmin(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextPageToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opSearchProductsAsAdmin(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "SearchProductsAsAdmin",
	}
}
