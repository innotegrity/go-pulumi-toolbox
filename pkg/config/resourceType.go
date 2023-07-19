package config

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

// NewResourceTypeProvisionerFn defines the function interface for creating a new ResourceTypeProvisioner object.
type NewResourceTypeProvisionerFn func() ResourceTypeProvisioner

// ResourceTypes maps a resource type to the function to use for creating a corresponding ResourceTypeProvisioner
// object.
type ResourceTypes map[string]NewResourceTypeProvisionerFn

// ResourceTypeProvisioner defines the interface all resource objects must implement in order to be provisioned.
type ResourceTypeProvisioner interface {
	GetID() string
	GetOutput(*pulumi.Context, pulumi.Resource, string) (pulumi.Output, error)
	GetType() string
	Provision(*pulumi.Context, ...pulumi.ResourceOption) (pulumi.Resource, error)
}
