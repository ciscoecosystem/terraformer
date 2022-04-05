// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the list of all policies in an organization of a specified type.
// Always check the NextToken response parameter for a null value when calling a
// List* operation. These operations can occasionally return an empty set of
// results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display. This
// operation can be called only from the organization's management account or by a
// member account that is a delegated administrator for an AWS service.
func (c *Client) ListPolicies(ctx context.Context, params *ListPoliciesInput, optFns ...func(*Options)) (*ListPoliciesOutput, error) {
	if params == nil {
		params = &ListPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPolicies", params, optFns, addOperationListPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPoliciesInput struct {

	// Specifies the type of policy that you want to include in the response. You must
	// specify one of the following values:
	//
	// * AISERVICES_OPT_OUT_POLICY
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_ai-opt-out.html)
	//
	// *
	// BACKUP_POLICY
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_backup.html)
	//
	// *
	// SERVICE_CONTROL_POLICY
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scp.html)
	//
	// *
	// TAG_POLICY
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies.html)
	//
	// This member is required.
	Filter types.PolicyType

	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific to
	// the operation. If additional items exist beyond the maximum you specify, the
	// NextToken response element is present and has a value (is not null). Include
	// that value as the NextToken request parameter in the next call to the operation
	// to get the next part of the results. Note that Organizations might return fewer
	// results than the maximum even when there are more results available. You should
	// check NextToken after every operation to ensure that you receive all of the
	// results.
	MaxResults *int32

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string
}

type ListPoliciesOutput struct {

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null.
	NextToken *string

	// A list of policies that match the filter criteria in the request. The output
	// list doesn't include the policy contents. To see the content for a policy, see
	// DescribePolicy.
	Policies []types.PolicySummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPolicies{}, middleware.After)
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
	if err = addOpListPoliciesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPolicies(options.Region), middleware.Before); err != nil {
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

// ListPoliciesAPIClient is a client that implements the ListPolicies operation.
type ListPoliciesAPIClient interface {
	ListPolicies(context.Context, *ListPoliciesInput, ...func(*Options)) (*ListPoliciesOutput, error)
}

var _ ListPoliciesAPIClient = (*Client)(nil)

// ListPoliciesPaginatorOptions is the paginator options for ListPolicies
type ListPoliciesPaginatorOptions struct {
	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific to
	// the operation. If additional items exist beyond the maximum you specify, the
	// NextToken response element is present and has a value (is not null). Include
	// that value as the NextToken request parameter in the next call to the operation
	// to get the next part of the results. Note that Organizations might return fewer
	// results than the maximum even when there are more results available. You should
	// check NextToken after every operation to ensure that you receive all of the
	// results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPoliciesPaginator is a paginator for ListPolicies
type ListPoliciesPaginator struct {
	options   ListPoliciesPaginatorOptions
	client    ListPoliciesAPIClient
	params    *ListPoliciesInput
	nextToken *string
	firstPage bool
}

// NewListPoliciesPaginator returns a new ListPoliciesPaginator
func NewListPoliciesPaginator(client ListPoliciesAPIClient, params *ListPoliciesInput, optFns ...func(*ListPoliciesPaginatorOptions)) *ListPoliciesPaginator {
	if params == nil {
		params = &ListPoliciesInput{}
	}

	options := ListPoliciesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPoliciesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPoliciesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListPolicies page.
func (p *ListPoliciesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPoliciesOutput, error) {
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

	result, err := p.client.ListPolicies(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "ListPolicies",
	}
}
