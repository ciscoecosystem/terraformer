// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackage

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing OriginEndpoint.
func (c *Client) UpdateOriginEndpoint(ctx context.Context, params *UpdateOriginEndpointInput, optFns ...func(*Options)) (*UpdateOriginEndpointOutput, error) {
	if params == nil {
		params = &UpdateOriginEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateOriginEndpoint", params, optFns, c.addOperationUpdateOriginEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateOriginEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Configuration parameters used to update an existing OriginEndpoint.
type UpdateOriginEndpointInput struct {

	// The ID of the OriginEndpoint to update.
	//
	// This member is required.
	Id *string

	// CDN Authorization credentials
	Authorization *types.Authorization

	// A Common Media Application Format (CMAF) packaging configuration.
	CmafPackage *types.CmafPackageCreateOrUpdateParameters

	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	DashPackage *types.DashPackage

	// A short text description of the OriginEndpoint.
	Description *string

	// An HTTP Live Streaming (HLS) packaging configuration.
	HlsPackage *types.HlsPackage

	// A short string that will be appended to the end of the Endpoint URL.
	ManifestName *string

	// A Microsoft Smooth Streaming (MSS) packaging configuration.
	MssPackage *types.MssPackage

	// Control whether origination of video is allowed for this OriginEndpoint. If set
	// to ALLOW, the OriginEndpoint may by requested, pursuant to any other form of
	// access control. If set to DENY, the OriginEndpoint may not be requested. This
	// can be helpful for Live to VOD harvesting, or for temporarily disabling
	// origination
	Origination types.Origination

	// Maximum duration (in seconds) of content to retain for startover playback. If
	// not specified, startover playback will be disabled for the OriginEndpoint.
	StartoverWindowSeconds int32

	// Amount of delay (in seconds) to enforce on the playback of live content. If not
	// specified, there will be no time delay in effect for the OriginEndpoint.
	TimeDelaySeconds int32

	// A list of source IP CIDR blocks that will be allowed to access the
	// OriginEndpoint.
	Whitelist []string

	noSmithyDocumentSerde
}

type UpdateOriginEndpointOutput struct {

	// The Amazon Resource Name (ARN) assigned to the OriginEndpoint.
	Arn *string

	// CDN Authorization credentials
	Authorization *types.Authorization

	// The ID of the Channel the OriginEndpoint is associated with.
	ChannelId *string

	// A Common Media Application Format (CMAF) packaging configuration.
	CmafPackage *types.CmafPackage

	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	DashPackage *types.DashPackage

	// A short text description of the OriginEndpoint.
	Description *string

	// An HTTP Live Streaming (HLS) packaging configuration.
	HlsPackage *types.HlsPackage

	// The ID of the OriginEndpoint.
	Id *string

	// A short string appended to the end of the OriginEndpoint URL.
	ManifestName *string

	// A Microsoft Smooth Streaming (MSS) packaging configuration.
	MssPackage *types.MssPackage

	// Control whether origination of video is allowed for this OriginEndpoint. If set
	// to ALLOW, the OriginEndpoint may by requested, pursuant to any other form of
	// access control. If set to DENY, the OriginEndpoint may not be requested. This
	// can be helpful for Live to VOD harvesting, or for temporarily disabling
	// origination
	Origination types.Origination

	// Maximum duration (seconds) of content to retain for startover playback. If not
	// specified, startover playback will be disabled for the OriginEndpoint.
	StartoverWindowSeconds int32

	// A collection of tags associated with a resource
	Tags map[string]string

	// Amount of delay (seconds) to enforce on the playback of live content. If not
	// specified, there will be no time delay in effect for the OriginEndpoint.
	TimeDelaySeconds int32

	// The URL of the packaged OriginEndpoint for consumption.
	Url *string

	// A list of source IP CIDR blocks that will be allowed to access the
	// OriginEndpoint.
	Whitelist []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateOriginEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateOriginEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateOriginEndpoint{}, middleware.After)
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
	if err = addOpUpdateOriginEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateOriginEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateOriginEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage",
		OperationName: "UpdateOriginEndpoint",
	}
}
