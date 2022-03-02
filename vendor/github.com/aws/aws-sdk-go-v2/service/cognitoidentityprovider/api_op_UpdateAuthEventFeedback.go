// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides the feedback for an authentication event whether it was from a valid
// user or not. This feedback is used for improving the risk evaluation decision
// for the user pool as part of Amazon Cognito advanced security.
func (c *Client) UpdateAuthEventFeedback(ctx context.Context, params *UpdateAuthEventFeedbackInput, optFns ...func(*Options)) (*UpdateAuthEventFeedbackOutput, error) {
	if params == nil {
		params = &UpdateAuthEventFeedbackInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAuthEventFeedback", params, optFns, addOperationUpdateAuthEventFeedbackMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAuthEventFeedbackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAuthEventFeedbackInput struct {

	// The event ID.
	//
	// This member is required.
	EventId *string

	// The feedback token.
	//
	// This member is required.
	FeedbackToken *string

	// The authentication event feedback value.
	//
	// This member is required.
	FeedbackValue types.FeedbackValueType

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// The user pool username.
	//
	// This member is required.
	Username *string
}

type UpdateAuthEventFeedbackOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateAuthEventFeedbackMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateAuthEventFeedback{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateAuthEventFeedback{}, middleware.After)
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
	if err = addOpUpdateAuthEventFeedbackValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAuthEventFeedback(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAuthEventFeedback(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "UpdateAuthEventFeedback",
	}
}