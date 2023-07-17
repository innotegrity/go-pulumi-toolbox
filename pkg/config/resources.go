package config

import (
	"encoding/json"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource defines the configuration for a particular resource to be provisioned.
type Resource struct {
	Type   string          `json:"type"`
	Params json.RawMessage `json:"params"`
}

// ResourceProvisioner defines the interface all resource objects must implement in order to be provisioned.
type ResourceProvisioner interface {
	Provision(*pulumi.Context, StackReferences) error
}

// NewResourceProvisionerFn defines the function interface for creating a new ResourceProvisioner object.
type NewResourceProvisionerFn func() ResourceProvisioner

// ResourceTypes maps a resource type to the function to use for creating a corresponding ResourceProvisioner object.
type ResourceTypes map[string]NewResourceProvisionerFn
