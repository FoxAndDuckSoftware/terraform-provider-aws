// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newLifecyclePolicyResource,
			Name:    "Lifecycle Policy",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceComponent,
			TypeName: "aws_imagebuilder_component",
			Name:     "Component",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceComponents,
			TypeName: "aws_imagebuilder_components",
			Name:     "Components",
		},
		{
			Factory:  dataSourceContainerRecipe,
			TypeName: "aws_imagebuilder_container_recipe",
			Name:     "Container Recipe",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceContainerRecipes,
			TypeName: "aws_imagebuilder_container_recipes",
			Name:     "Container Recipes",
		},
		{
			Factory:  dataSourceDistributionConfiguration,
			TypeName: "aws_imagebuilder_distribution_configuration",
			Name:     "Distribution Configuration",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceDistributionConfigurations,
			TypeName: "aws_imagebuilder_distribution_configurations",
			Name:     "Distribution Configurations",
		},
		{
			Factory:  dataSourceImage,
			TypeName: "aws_imagebuilder_image",
			Name:     "Image",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceImagePipeline,
			TypeName: "aws_imagebuilder_image_pipeline",
			Name:     "Image Pipeline",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceImagePipelines,
			TypeName: "aws_imagebuilder_image_pipelines",
			Name:     "Image Pipelines",
		},
		{
			Factory:  dataSourceImageRecipe,
			TypeName: "aws_imagebuilder_image_recipe",
			Name:     "Image Recipe",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceImageRecipes,
			TypeName: "aws_imagebuilder_image_recipes",
			Name:     "Image Recipes",
		},
		{
			Factory:  dataSourceInfrastructureConfiguration,
			TypeName: "aws_imagebuilder_infrastructure_configuration",
			Name:     "Infrastructure Configuration",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceInfrastructureConfigurations,
			TypeName: "aws_imagebuilder_infrastructure_configurations",
			Name:     "Infrastructure Configurations",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceComponent,
			TypeName: "aws_imagebuilder_component",
			Name:     "Component",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceContainerRecipe,
			TypeName: "aws_imagebuilder_container_recipe",
			Name:     "Container Recipe",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceDistributionConfiguration,
			TypeName: "aws_imagebuilder_distribution_configuration",
			Name:     "Distribution Configuration",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceImage,
			TypeName: "aws_imagebuilder_image",
			Name:     "Image",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceImagePipeline,
			TypeName: "aws_imagebuilder_image_pipeline",
			Name:     "Image Pipeline",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceImageRecipe,
			TypeName: "aws_imagebuilder_image_recipe",
			Name:     "Image Recipe",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceInfrastructureConfiguration,
			TypeName: "aws_imagebuilder_infrastructure_configuration",
			Name:     "Infrastructure Configuration",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceWorkflow,
			TypeName: "aws_imagebuilder_workflow",
			Name:     "Workflow",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ImageBuilder
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*imagebuilder.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))

	return imagebuilder.NewFromConfig(cfg,
		imagebuilder.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
