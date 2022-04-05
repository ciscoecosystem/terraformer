// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a cache for the GraphQL API.
func (c *Client) CreateApiCache(ctx context.Context, params *CreateApiCacheInput, optFns ...func(*Options)) (*CreateApiCacheOutput, error) {
	if params == nil {
		params = &CreateApiCacheInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateApiCache", params, optFns, addOperationCreateApiCacheMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateApiCacheOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateApiCache operation.
type CreateApiCacheInput struct {

	// Caching behavior.
	//
	// * FULL_REQUEST_CACHING: All requests are fully cached.
	//
	// *
	// PER_RESOLVER_CACHING: Individual resolvers that you specify are cached.
	//
	// This member is required.
	ApiCachingBehavior types.ApiCachingBehavior

	// The GraphQL API Id.
	//
	// This member is required.
	ApiId *string

	// TTL in seconds for cache entries. Valid values are between 1 and 3600 seconds.
	//
	// This member is required.
	Ttl int64

	// The cache instance type. Valid values are
	//
	// * SMALL
	//
	// * MEDIUM
	//
	// * LARGE
	//
	// *
	// XLARGE
	//
	// * LARGE_2X
	//
	// * LARGE_4X
	//
	// * LARGE_8X (not available in all regions)
	//
	// *
	// LARGE_12X
	//
	// Historically, instance types were identified by an EC2-style value.
	// As of July 2020, this is deprecated, and the generic identifiers above should be
	// used. The following legacy instance types are available, but their use is
	// discouraged:
	//
	// * T2_SMALL: A t2.small instance type.
	//
	// * T2_MEDIUM: A t2.medium
	// instance type.
	//
	// * R4_LARGE: A r4.large instance type.
	//
	// * R4_XLARGE: A r4.xlarge
	// instance type.
	//
	// * R4_2XLARGE: A r4.2xlarge instance type.
	//
	// * R4_4XLARGE: A
	// r4.4xlarge instance type.
	//
	// * R4_8XLARGE: A r4.8xlarge instance type.
	//
	// This member is required.
	Type types.ApiCacheType

	// At rest encryption flag for cache. This setting cannot be updated after
	// creation.
	AtRestEncryptionEnabled bool

	// Transit encryption flag when connecting to cache. This setting cannot be updated
	// after creation.
	TransitEncryptionEnabled bool
}

// Represents the output of a CreateApiCache operation.
type CreateApiCacheOutput struct {

	// The ApiCache object.
	ApiCache *types.ApiCache

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateApiCacheMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateApiCache{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateApiCache{}, middleware.After)
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
	if err = addOpCreateApiCacheValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApiCache(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateApiCache(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "CreateApiCache",
	}
}
