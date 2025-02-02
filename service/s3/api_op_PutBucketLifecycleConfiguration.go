// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new lifecycle configuration for the bucket or replaces an existing
// lifecycle configuration. For information about lifecycle configuration, see
// Managing your storage lifecycle
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html).
// Bucket lifecycle configuration now supports specifying a lifecycle rule using an
// object key name prefix, one or more object tags, or a combination of both.
// Accordingly, this section describes the latest API. The previous version of the
// API supported filtering based only on an object key name prefix, which is
// supported for backward compatibility. For the related API description, see
// PutBucketLifecycle
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketLifecycle.html).
// Rules You specify the lifecycle configuration in your request body. The
// lifecycle configuration is specified as XML consisting of one or more rules.
// Each rule consists of the following:
//
// * Filter identifying a subset of objects
// to which the rule applies. The filter can be based on a key name prefix, object
// tags, or a combination of both.
//
// * Status whether the rule is in effect.
//
// * One
// or more lifecycle transition and expiration actions that you want Amazon S3 to
// perform on the objects identified by the filter. If the state of your bucket is
// versioning-enabled or versioning-suspended, you can have many versions of the
// same object (one current version and zero or more noncurrent versions). Amazon
// S3 provides predefined actions that you can specify for current and noncurrent
// object versions.
//
// For more information, see Object Lifecycle Management
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) and
// Lifecycle Configuration Elements
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/intro-lifecycle-rules.html).
// Permissions By default, all Amazon S3 resources are private, including buckets,
// objects, and related subresources (for example, lifecycle configuration and
// website configuration). Only the resource owner (that is, the Amazon Web
// Services account that created it) can access the resource. The resource owner
// can optionally grant access permissions to others by writing an access policy.
// For this operation, a user must get the s3:PutLifecycleConfiguration permission.
// You can also explicitly deny permissions. Explicit deny also supersedes any
// other permissions. If you want to block users or accounts from removing or
// deleting objects from your bucket, you must deny them permissions for the
// following actions:
//
// * s3:DeleteObject
//
// * s3:DeleteObjectVersion
//
// *
// s3:PutLifecycleConfiguration
//
// For more information about permissions, see
// Managing Access Permissions to Your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html).
// The following are related to PutBucketLifecycleConfiguration:
//
// * Examples of
// Lifecycle Configuration
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/lifecycle-configuration-examples.html)
//
// *
// GetBucketLifecycleConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLifecycleConfiguration.html)
//
// *
// DeleteBucketLifecycle
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketLifecycle.html)
func (c *Client) PutBucketLifecycleConfiguration(ctx context.Context, params *PutBucketLifecycleConfigurationInput, optFns ...func(*Options)) (*PutBucketLifecycleConfigurationOutput, error) {
	if params == nil {
		params = &PutBucketLifecycleConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketLifecycleConfiguration", params, optFns, c.addOperationPutBucketLifecycleConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketLifecycleConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketLifecycleConfigurationInput struct {

	// The name of the bucket for which to set the configuration.
	//
	// This member is required.
	Bucket *string

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string

	// Container for lifecycle rules. You can add as many as 1,000 rules.
	LifecycleConfiguration *types.BucketLifecycleConfiguration

	noSmithyDocumentSerde
}

type PutBucketLifecycleConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutBucketLifecycleConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketLifecycleConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketLifecycleConfiguration{}, middleware.After)
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
	if err = swapWithCustomHTTPSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutBucketLifecycleConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketLifecycleConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addPutBucketLifecycleConfigurationUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutBucketLifecycleConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketLifecycleConfiguration",
	}
}

// getPutBucketLifecycleConfigurationBucketMember returns a pointer to string
// denoting a provided bucket member valueand a boolean indicating if the input has
// a modeled bucket name,
func getPutBucketLifecycleConfigurationBucketMember(input interface{}) (*string, bool) {
	in := input.(*PutBucketLifecycleConfigurationInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addPutBucketLifecycleConfigurationUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getPutBucketLifecycleConfigurationBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseDualstack:                   options.UseDualstack,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}
