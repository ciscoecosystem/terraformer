// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the parameters of a parameter group. For the parameters parameter, it
// can't contain ASCII characters. For more information about parameters and
// parameter groups, go to Amazon Redshift Parameter Groups
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html)
// in the Amazon Redshift Cluster Management Guide.
func (c *Client) ModifyClusterParameterGroup(ctx context.Context, params *ModifyClusterParameterGroupInput, optFns ...func(*Options)) (*ModifyClusterParameterGroupOutput, error) {
	if params == nil {
		params = &ModifyClusterParameterGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyClusterParameterGroup", params, optFns, c.addOperationModifyClusterParameterGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyClusterParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Describes a modify cluster parameter group operation.
type ModifyClusterParameterGroupInput struct {

	// The name of the parameter group to be modified.
	//
	// This member is required.
	ParameterGroupName *string

	// An array of parameters to be modified. A maximum of 20 parameters can be
	// modified in a single request. For each parameter to be modified, you must supply
	// at least the parameter name and parameter value; other name-value pairs of the
	// parameter are optional. For the workload management (WLM) configuration, you
	// must supply all the name-value pairs in the wlm_json_configuration parameter.
	//
	// This member is required.
	Parameters []types.Parameter

	noSmithyDocumentSerde
}

//
type ModifyClusterParameterGroupOutput struct {

	// The name of the cluster parameter group.
	ParameterGroupName *string

	// The status of the parameter group. For example, if you made a change to a
	// parameter group name-value pair, then the change could be pending a reboot of an
	// associated cluster.
	ParameterGroupStatus *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyClusterParameterGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyClusterParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyClusterParameterGroup{}, middleware.After)
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
	if err = addOpModifyClusterParameterGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyClusterParameterGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyClusterParameterGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "ModifyClusterParameterGroup",
	}
}
