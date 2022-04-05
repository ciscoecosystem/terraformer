// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List all versions for a document.
func (c *Client) ListDocumentVersions(ctx context.Context, params *ListDocumentVersionsInput, optFns ...func(*Options)) (*ListDocumentVersionsOutput, error) {
	if params == nil {
		params = &ListDocumentVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDocumentVersions", params, optFns, addOperationListDocumentVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDocumentVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDocumentVersionsInput struct {

	// The name of the document. You can specify an Amazon Resource Name (ARN).
	//
	// This member is required.
	Name *string

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
}

type ListDocumentVersionsOutput struct {

	// The document versions.
	DocumentVersions []types.DocumentVersionInfo

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDocumentVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDocumentVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDocumentVersions{}, middleware.After)
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
	if err = addOpListDocumentVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDocumentVersions(options.Region), middleware.Before); err != nil {
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

// ListDocumentVersionsAPIClient is a client that implements the
// ListDocumentVersions operation.
type ListDocumentVersionsAPIClient interface {
	ListDocumentVersions(context.Context, *ListDocumentVersionsInput, ...func(*Options)) (*ListDocumentVersionsOutput, error)
}

var _ ListDocumentVersionsAPIClient = (*Client)(nil)

// ListDocumentVersionsPaginatorOptions is the paginator options for
// ListDocumentVersions
type ListDocumentVersionsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDocumentVersionsPaginator is a paginator for ListDocumentVersions
type ListDocumentVersionsPaginator struct {
	options   ListDocumentVersionsPaginatorOptions
	client    ListDocumentVersionsAPIClient
	params    *ListDocumentVersionsInput
	nextToken *string
	firstPage bool
}

// NewListDocumentVersionsPaginator returns a new ListDocumentVersionsPaginator
func NewListDocumentVersionsPaginator(client ListDocumentVersionsAPIClient, params *ListDocumentVersionsInput, optFns ...func(*ListDocumentVersionsPaginatorOptions)) *ListDocumentVersionsPaginator {
	if params == nil {
		params = &ListDocumentVersionsInput{}
	}

	options := ListDocumentVersionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDocumentVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDocumentVersionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListDocumentVersions page.
func (p *ListDocumentVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDocumentVersionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListDocumentVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDocumentVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "ListDocumentVersions",
	}
}
