// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a pull through cache rule. A pull through cache rule provides a way to
// cache images from an upstream registry source in your Amazon ECR private
// registry. For more information, see Using pull through cache rules (https://docs.aws.amazon.com/AmazonECR/latest/userguide/pull-through-cache.html)
// in the Amazon Elastic Container Registry User Guide.
func (c *Client) CreatePullThroughCacheRule(ctx context.Context, params *CreatePullThroughCacheRuleInput, optFns ...func(*Options)) (*CreatePullThroughCacheRuleOutput, error) {
	if params == nil {
		params = &CreatePullThroughCacheRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePullThroughCacheRule", params, optFns, c.addOperationCreatePullThroughCacheRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePullThroughCacheRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePullThroughCacheRuleInput struct {

	// The repository name prefix to use when caching images from the source registry.
	//
	// This member is required.
	EcrRepositoryPrefix *string

	// The registry URL of the upstream public registry to use as the source for the
	// pull through cache rule. The following is the syntax to use for each supported
	// upstream registry.
	//   - Amazon ECR Public ( ecr-public ) - public.ecr.aws
	//   - Docker Hub ( docker-hub ) - registry-1.docker.io
	//   - Quay ( quay ) - quay.io
	//   - Kubernetes ( k8s ) - registry.k8s.io
	//   - GitHub Container Registry ( github-container-registry ) - ghcr.io
	//   - Microsoft Azure Container Registry ( azure-container-registry ) -
	//   .azurecr.io
	//
	// This member is required.
	UpstreamRegistryUrl *string

	// The Amazon Resource Name (ARN) of the Amazon Web Services Secrets Manager
	// secret that identifies the credentials to authenticate to the upstream registry.
	CredentialArn *string

	// The Amazon Web Services account ID associated with the registry to create the
	// pull through cache rule for. If you do not specify a registry, the default
	// registry is assumed.
	RegistryId *string

	// The name of the upstream registry.
	UpstreamRegistry types.UpstreamRegistry

	noSmithyDocumentSerde
}

type CreatePullThroughCacheRuleOutput struct {

	// The date and time, in JavaScript date format, when the pull through cache rule
	// was created.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the Amazon Web Services Secrets Manager
	// secret associated with the pull through cache rule.
	CredentialArn *string

	// The Amazon ECR repository prefix associated with the pull through cache rule.
	EcrRepositoryPrefix *string

	// The registry ID associated with the request.
	RegistryId *string

	// The name of the upstream registry associated with the pull through cache rule.
	UpstreamRegistry types.UpstreamRegistry

	// The upstream registry URL associated with the pull through cache rule.
	UpstreamRegistryUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePullThroughCacheRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePullThroughCacheRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePullThroughCacheRule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePullThroughCacheRule"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreatePullThroughCacheRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePullThroughCacheRule(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreatePullThroughCacheRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePullThroughCacheRule",
	}
}