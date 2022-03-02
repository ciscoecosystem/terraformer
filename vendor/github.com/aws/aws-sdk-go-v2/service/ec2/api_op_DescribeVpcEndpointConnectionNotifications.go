// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the connection notifications for VPC endpoints and VPC endpoint
// services.
func (c *Client) DescribeVpcEndpointConnectionNotifications(ctx context.Context, params *DescribeVpcEndpointConnectionNotificationsInput, optFns ...func(*Options)) (*DescribeVpcEndpointConnectionNotificationsOutput, error) {
	if params == nil {
		params = &DescribeVpcEndpointConnectionNotificationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVpcEndpointConnectionNotifications", params, optFns, addOperationDescribeVpcEndpointConnectionNotificationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVpcEndpointConnectionNotificationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVpcEndpointConnectionNotificationsInput struct {

	// The ID of the notification.
	ConnectionNotificationId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// One or more filters.
	//
	// * connection-notification-arn - The ARN of the SNS topic
	// for the notification.
	//
	// * connection-notification-id - The ID of the
	// notification.
	//
	// * connection-notification-state - The state of the notification
	// (Enabled | Disabled).
	//
	// * connection-notification-type - The type of notification
	// (Topic).
	//
	// * service-id - The ID of the endpoint service.
	//
	// * vpc-endpoint-id -
	// The ID of the VPC endpoint.
	Filters []types.Filter

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another request with the returned NextToken value.
	MaxResults int32

	// The token to request the next page of results.
	NextToken *string
}

type DescribeVpcEndpointConnectionNotificationsOutput struct {

	// One or more notifications.
	ConnectionNotificationSet []types.ConnectionNotification

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeVpcEndpointConnectionNotificationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeVpcEndpointConnectionNotifications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVpcEndpointConnectionNotifications{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcEndpointConnectionNotifications(options.Region), middleware.Before); err != nil {
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

// DescribeVpcEndpointConnectionNotificationsAPIClient is a client that implements
// the DescribeVpcEndpointConnectionNotifications operation.
type DescribeVpcEndpointConnectionNotificationsAPIClient interface {
	DescribeVpcEndpointConnectionNotifications(context.Context, *DescribeVpcEndpointConnectionNotificationsInput, ...func(*Options)) (*DescribeVpcEndpointConnectionNotificationsOutput, error)
}

var _ DescribeVpcEndpointConnectionNotificationsAPIClient = (*Client)(nil)

// DescribeVpcEndpointConnectionNotificationsPaginatorOptions is the paginator
// options for DescribeVpcEndpointConnectionNotifications
type DescribeVpcEndpointConnectionNotificationsPaginatorOptions struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another request with the returned NextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeVpcEndpointConnectionNotificationsPaginator is a paginator for
// DescribeVpcEndpointConnectionNotifications
type DescribeVpcEndpointConnectionNotificationsPaginator struct {
	options   DescribeVpcEndpointConnectionNotificationsPaginatorOptions
	client    DescribeVpcEndpointConnectionNotificationsAPIClient
	params    *DescribeVpcEndpointConnectionNotificationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeVpcEndpointConnectionNotificationsPaginator returns a new
// DescribeVpcEndpointConnectionNotificationsPaginator
func NewDescribeVpcEndpointConnectionNotificationsPaginator(client DescribeVpcEndpointConnectionNotificationsAPIClient, params *DescribeVpcEndpointConnectionNotificationsInput, optFns ...func(*DescribeVpcEndpointConnectionNotificationsPaginatorOptions)) *DescribeVpcEndpointConnectionNotificationsPaginator {
	if params == nil {
		params = &DescribeVpcEndpointConnectionNotificationsInput{}
	}

	options := DescribeVpcEndpointConnectionNotificationsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeVpcEndpointConnectionNotificationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeVpcEndpointConnectionNotificationsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeVpcEndpointConnectionNotifications page.
func (p *DescribeVpcEndpointConnectionNotificationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeVpcEndpointConnectionNotificationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeVpcEndpointConnectionNotifications(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeVpcEndpointConnectionNotifications(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVpcEndpointConnectionNotifications",
	}
}