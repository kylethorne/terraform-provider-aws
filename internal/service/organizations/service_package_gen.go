// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package organizations

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	organizations_sdkv1 "github.com/aws/aws-sdk-go/service/organizations"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceDelegatedAdministrators,
			TypeName: "aws_organizations_delegated_administrators",
		},
		{
			Factory:  DataSourceDelegatedServices,
			TypeName: "aws_organizations_delegated_services",
		},
		{
			Factory:  DataSourceOrganization,
			TypeName: "aws_organizations_organization",
		},
		{
			Factory:  DataSourceOrganizationalUnitChildAccounts,
			TypeName: "aws_organizations_organizational_unit_child_accounts",
		},
		{
			Factory:  DataSourceOrganizationalUnitDescendantAccounts,
			TypeName: "aws_organizations_organizational_unit_descendant_accounts",
		},
		{
			Factory:  DataSourceOrganizationalUnits,
			TypeName: "aws_organizations_organizational_units",
		},
		{
			Factory:  DataSourceOrganizationalUnit,
			TypeName: "aws_organizations_organizational_unit",
		},
		{
			Factory:  DataSourcePolicies,
			TypeName: "aws_organizations_policies",
		},
		{
			Factory:  DataSourcePoliciesForTarget,
			TypeName: "aws_organizations_policies_for_target",
		},
		{
			Factory:  DataSourcePolicy,
			TypeName: "aws_organizations_policy",
		},
		{
			Factory:  DataSourceResourceTags,
			TypeName: "aws_organizations_resource_tags",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAccount,
			TypeName: "aws_organizations_account",
			Name:     "Account",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceDelegatedAdministrator,
			TypeName: "aws_organizations_delegated_administrator",
		},
		{
			Factory:  ResourceOrganization,
			TypeName: "aws_organizations_organization",
		},
		{
			Factory:  ResourceOrganizationalUnit,
			TypeName: "aws_organizations_organizational_unit",
			Name:     "Organizational Unit",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourcePolicy,
			TypeName: "aws_organizations_policy",
			Name:     "Policy",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourcePolicyAttachment,
			TypeName: "aws_organizations_policy_attachment",
		},
		{
			Factory:  ResourceResourcePolicy,
			TypeName: "aws_organizations_resource_policy",
			Name:     "Resource Policy",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Organizations
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*organizations_sdkv1.Organizations, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return organizations_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
