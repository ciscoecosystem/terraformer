// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes properties of scheduled actions.
func (c *Client) DescribeScheduledActions(ctx context.Context, params *DescribeScheduledActionsInput, optFns ...func(*Options)) (*DescribeScheduledActionsOutput, error) {
	if params == nil {
		params = &DescribeScheduledActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeScheduledActions", params, optFns, c.addOperationDescribeScheduledActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeScheduledActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeScheduledActionsInput struct {

	// If true, retrieve only active scheduled actions. If false, retrieve only
	// disabled scheduled actions.
	Active *bool

	// The end time in UTC of the scheduled action to retrieve. Only active scheduled
	// actions that have invocations before this time are retrieved.
	EndTime *time.Time

	// List of scheduled action filters.
	Filters []types.ScheduledActionFilter

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeScheduledActions request exceed
	// the value specified in MaxRecords, Amazon Web Services returns a value in the
	// Marker field of the response. You can retrieve the next set of response records
	// by providing the returned marker value in the Marker parameter and retrying the
	// request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int32

	// The name of the scheduled action to retrieve.
	ScheduledActionName *string

	// The start time in UTC of the scheduled actions to retrieve. Only active
	// scheduled actions that have invocations after this time are retrieved.
	StartTime *time.Time

	// The type of the scheduled actions to retrieve.
	TargetActionType types.ScheduledActionTypeValues

	noSmithyDocumentSerde
}

type DescribeScheduledActionsOutput struct {

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeScheduledActions request exceed
	// the value specified in MaxRecords, Amazon Web Services returns a value in the
	// Marker field of the response. You can retrieve the next set of response records
	// by providing the returned marker value in the Marker parameter and retrying the
	// request.
	Marker *string

	// List of retrieved scheduled actions.
	ScheduledActions []types.ScheduledAction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeScheduledActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeScheduledActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeScheduledActions{}, middleware.After)
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
	if err = addOpDescribeScheduledActionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeScheduledActions(options.Region), middleware.Before); err != nil {
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

// DescribeScheduledActionsAPIClient is a client that implements the
// DescribeScheduledActions operation.
type DescribeScheduledActionsAPIClient interface {
	DescribeScheduledActions(context.Context, *DescribeScheduledActionsInput, ...func(*Options)) (*DescribeScheduledActionsOutput, error)
}

var _ DescribeScheduledActionsAPIClient = (*Client)(nil)

// DescribeScheduledActionsPaginatorOptions is the paginator options for
// DescribeScheduledActions
type DescribeScheduledActionsPaginatorOptions struct {
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeScheduledActionsPaginator is a paginator for DescribeScheduledActions
type DescribeScheduledActionsPaginator struct {
	options   DescribeScheduledActionsPaginatorOptions
	client    DescribeScheduledActionsAPIClient
	params    *DescribeScheduledActionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeScheduledActionsPaginator returns a new
// DescribeScheduledActionsPaginator
func NewDescribeScheduledActionsPaginator(client DescribeScheduledActionsAPIClient, params *DescribeScheduledActionsInput, optFns ...func(*DescribeScheduledActionsPaginatorOptions)) *DescribeScheduledActionsPaginator {
	if params == nil {
		params = &DescribeScheduledActionsInput{}
	}

	options := DescribeScheduledActionsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeScheduledActionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeScheduledActionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeScheduledActions page.
func (p *DescribeScheduledActionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeScheduledActionsOutput, error) {
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

	result, err := p.client.DescribeScheduledActions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeScheduledActions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeScheduledActions",
	}
}
