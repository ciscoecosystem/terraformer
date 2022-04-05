// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the specified metrics. You can use the returned metrics with GetMetricData
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricData.html)
// or GetMetricStatistics
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html)
// to obtain statistical data. Up to 500 results are returned for any one call. To
// retrieve additional results, use the returned token with subsequent calls. After
// you create a metric, allow up to 15 minutes before the metric appears. You can
// see statistics about the metric sooner by using GetMetricData
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricData.html)
// or GetMetricStatistics
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html).
// ListMetrics doesn't return information about metrics if those metrics haven't
// reported data in the past two weeks. To retrieve those metrics, use
// GetMetricData
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricData.html)
// or GetMetricStatistics
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html).
func (c *Client) ListMetrics(ctx context.Context, params *ListMetricsInput, optFns ...func(*Options)) (*ListMetricsOutput, error) {
	if params == nil {
		params = &ListMetricsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMetrics", params, optFns, addOperationListMetricsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMetricsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMetricsInput struct {

	// The dimensions to filter against. Only the dimensions that match exactly will be
	// returned.
	Dimensions []types.DimensionFilter

	// The name of the metric to filter against. Only the metrics with names that match
	// exactly will be returned.
	MetricName *string

	// The metric namespace to filter against. Only the namespace that matches exactly
	// will be returned.
	Namespace *string

	// The token returned by a previous call to indicate that there is more data
	// available.
	NextToken *string

	// To filter the results to show only metrics that have had data points published
	// in the past three hours, specify this parameter with a value of PT3H. This is
	// the only valid value for this parameter. The results that are returned are an
	// approximation of the value you specify. There is a low probability that the
	// returned results include metrics with last published data as much as 40 minutes
	// more than the specified time interval.
	RecentlyActive types.RecentlyActive
}

type ListMetricsOutput struct {

	// The metrics that match your request.
	Metrics []types.Metric

	// The token that marks the start of the next batch of returned results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListMetricsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListMetrics{}, middleware.After)
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
	if err = addOpListMetricsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMetrics(options.Region), middleware.Before); err != nil {
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

// ListMetricsAPIClient is a client that implements the ListMetrics operation.
type ListMetricsAPIClient interface {
	ListMetrics(context.Context, *ListMetricsInput, ...func(*Options)) (*ListMetricsOutput, error)
}

var _ ListMetricsAPIClient = (*Client)(nil)

// ListMetricsPaginatorOptions is the paginator options for ListMetrics
type ListMetricsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMetricsPaginator is a paginator for ListMetrics
type ListMetricsPaginator struct {
	options   ListMetricsPaginatorOptions
	client    ListMetricsAPIClient
	params    *ListMetricsInput
	nextToken *string
	firstPage bool
}

// NewListMetricsPaginator returns a new ListMetricsPaginator
func NewListMetricsPaginator(client ListMetricsAPIClient, params *ListMetricsInput, optFns ...func(*ListMetricsPaginatorOptions)) *ListMetricsPaginator {
	if params == nil {
		params = &ListMetricsInput{}
	}

	options := ListMetricsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMetricsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMetricsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListMetrics page.
func (p *ListMetricsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMetricsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListMetrics(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMetrics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "ListMetrics",
	}
}
