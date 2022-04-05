// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the available option groups.
func (c *Client) DescribeOptionGroups(ctx context.Context, params *DescribeOptionGroupsInput, optFns ...func(*Options)) (*DescribeOptionGroupsOutput, error) {
	if params == nil {
		params = &DescribeOptionGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOptionGroups", params, optFns, c.addOperationDescribeOptionGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOptionGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeOptionGroupsInput struct {

	// Filters the list of option groups to only include groups associated with a
	// specific database engine. Valid Values:
	//
	// * mariadb
	//
	// * mysql
	//
	// * oracle-ee
	//
	// *
	// oracle-ee-cdb
	//
	// * oracle-se2
	//
	// * oracle-se2-cdb
	//
	// * postgres
	//
	// * sqlserver-ee
	//
	// *
	// sqlserver-se
	//
	// * sqlserver-ex
	//
	// * sqlserver-web
	EngineName *string

	// This parameter isn't currently supported.
	Filters []types.Filter

	// Filters the list of option groups to only include groups associated with a
	// specific database engine version. If specified, then EngineName must also be
	// specified.
	MajorEngineVersion *string

	// An optional pagination token provided by a previous DescribeOptionGroups
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// The name of the option group to describe. Can't be supplied together with
	// EngineName or MajorEngineVersion.
	OptionGroupName *string

	noSmithyDocumentSerde
}

// List of option groups.
type DescribeOptionGroupsOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// List of option groups.
	OptionGroupsList []types.OptionGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeOptionGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeOptionGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeOptionGroups{}, middleware.After)
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
	if err = addOpDescribeOptionGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOptionGroups(options.Region), middleware.Before); err != nil {
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

// DescribeOptionGroupsAPIClient is a client that implements the
// DescribeOptionGroups operation.
type DescribeOptionGroupsAPIClient interface {
	DescribeOptionGroups(context.Context, *DescribeOptionGroupsInput, ...func(*Options)) (*DescribeOptionGroupsOutput, error)
}

var _ DescribeOptionGroupsAPIClient = (*Client)(nil)

// DescribeOptionGroupsPaginatorOptions is the paginator options for
// DescribeOptionGroups
type DescribeOptionGroupsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeOptionGroupsPaginator is a paginator for DescribeOptionGroups
type DescribeOptionGroupsPaginator struct {
	options   DescribeOptionGroupsPaginatorOptions
	client    DescribeOptionGroupsAPIClient
	params    *DescribeOptionGroupsInput
	nextToken *string
	firstPage bool
}

// NewDescribeOptionGroupsPaginator returns a new DescribeOptionGroupsPaginator
func NewDescribeOptionGroupsPaginator(client DescribeOptionGroupsAPIClient, params *DescribeOptionGroupsInput, optFns ...func(*DescribeOptionGroupsPaginatorOptions)) *DescribeOptionGroupsPaginator {
	if params == nil {
		params = &DescribeOptionGroupsInput{}
	}

	options := DescribeOptionGroupsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeOptionGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeOptionGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeOptionGroups page.
func (p *DescribeOptionGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeOptionGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeOptionGroups(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeOptionGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeOptionGroups",
	}
}
