// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an Amazon EMR Studio configuration, including attributes such as name,
// description, and subnets.
func (c *Client) UpdateStudio(ctx context.Context, params *UpdateStudioInput, optFns ...func(*Options)) (*UpdateStudioOutput, error) {
	if params == nil {
		params = &UpdateStudioInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateStudio", params, optFns, addOperationUpdateStudioMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateStudioOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateStudioInput struct {

	// The ID of the Amazon EMR Studio to update.
	//
	// This member is required.
	StudioId *string

	// The Amazon S3 location to back up Workspaces and notebook files for the Amazon
	// EMR Studio.
	DefaultS3Location *string

	// A detailed description to assign to the Amazon EMR Studio.
	Description *string

	// A descriptive name for the Amazon EMR Studio.
	Name *string

	// A list of subnet IDs to associate with the Amazon EMR Studio. The list can
	// include new subnet IDs, but must also include all of the subnet IDs previously
	// associated with the Studio. The list order does not matter. A Studio can have a
	// maximum of 5 subnets. The subnets must belong to the same VPC as the Studio.
	SubnetIds []string
}

type UpdateStudioOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateStudioMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateStudio{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateStudio{}, middleware.After)
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
	if err = addOpUpdateStudioValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateStudio(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateStudio(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "UpdateStudio",
	}
}
