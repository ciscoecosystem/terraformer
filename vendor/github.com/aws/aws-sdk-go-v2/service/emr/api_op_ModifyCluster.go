// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the number of steps that can be executed concurrently for the cluster
// specified using ClusterID.
func (c *Client) ModifyCluster(ctx context.Context, params *ModifyClusterInput, optFns ...func(*Options)) (*ModifyClusterOutput, error) {
	if params == nil {
		params = &ModifyClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyCluster", params, optFns, addOperationModifyClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyClusterInput struct {

	// The unique identifier of the cluster.
	//
	// This member is required.
	ClusterId *string

	// The number of steps that can be executed concurrently. You can specify a minimum
	// of 1 step and a maximum of 256 steps.
	StepConcurrencyLevel *int32
}

type ModifyClusterOutput struct {

	// The number of steps that can be executed concurrently.
	StepConcurrencyLevel *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpModifyCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyCluster{}, middleware.After)
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
	if err = addOpModifyClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "ModifyCluster",
	}
}
