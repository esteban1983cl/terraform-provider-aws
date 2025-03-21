// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package elbv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory:  newDataSourceListenerRule,
			TypeName: "aws_lb_listener_rule",
			Name:     "Listener Rule",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceLoadBalancer,
			TypeName: "aws_alb",
			Name:     "Load Balancer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  dataSourceListener,
			TypeName: "aws_alb_listener",
			Name:     "Listener",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  dataSourceTargetGroup,
			TypeName: "aws_alb_target_group",
			Name:     "Target Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  dataSourceLoadBalancer,
			TypeName: "aws_lb",
			Name:     "Load Balancer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  dataSourceHostedZoneID,
			TypeName: "aws_lb_hosted_zone_id",
			Name:     "Hosted Zone ID",
		},
		{
			Factory:  dataSourceListener,
			TypeName: "aws_lb_listener",
			Name:     "Listener",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  dataSourceTargetGroup,
			TypeName: "aws_lb_target_group",
			Name:     "Target Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  dataSourceTrustStore,
			TypeName: "aws_lb_trust_store",
			Name:     "Trust Store",
		},
		{
			Factory:  dataSourceLoadBalancers,
			TypeName: "aws_lbs",
			Name:     "Load Balancers",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceLoadBalancer,
			TypeName: "aws_alb",
			Name:     "Load Balancer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceListener,
			TypeName: "aws_alb_listener",
			Name:     "Listener",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceListenerCertificate,
			TypeName: "aws_alb_listener_certificate",
			Name:     "Listener Certificate",
		},
		{
			Factory:  resourceListenerRule,
			TypeName: "aws_alb_listener_rule",
			Name:     "Listener Rule",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceTargetGroup,
			TypeName: "aws_alb_target_group",
			Name:     "Target Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceTargetGroupAttachment,
			TypeName: "aws_alb_target_group_attachment",
			Name:     "Target Group Attachment",
		},
		{
			Factory:  resourceLoadBalancer,
			TypeName: "aws_lb",
			Name:     "Load Balancer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceListener,
			TypeName: "aws_lb_listener",
			Name:     "Listener",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceListenerCertificate,
			TypeName: "aws_lb_listener_certificate",
			Name:     "Listener Certificate",
		},
		{
			Factory:  resourceListenerRule,
			TypeName: "aws_lb_listener_rule",
			Name:     "Listener Rule",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceTargetGroup,
			TypeName: "aws_lb_target_group",
			Name:     "Target Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceTargetGroupAttachment,
			TypeName: "aws_lb_target_group_attachment",
			Name:     "Target Group Attachment",
		},
		{
			Factory:  resourceTrustStore,
			TypeName: "aws_lb_trust_store",
			Name:     "Trust Store",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceTrustStoreRevocation,
			TypeName: "aws_lb_trust_store_revocation",
			Name:     "Trust Store Revocation",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ELBV2
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*elasticloadbalancingv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*elasticloadbalancingv2.Options){
		elasticloadbalancingv2.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return elasticloadbalancingv2.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*elasticloadbalancingv2.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*elasticloadbalancingv2.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *elasticloadbalancingv2.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*elasticloadbalancingv2.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
