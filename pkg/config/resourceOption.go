package config

import (
	"encoding/json"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NewResourceOptionLoaderFn defines the function interface for creating a new ResourceOptionLoader object.
type NewResourceOptionLoaderFn func() ResourceOptionLoader

// ResourceOptions maps a resource option to the function to use for creating a corresponding ResourceOptionLoader
// object.
type ResourceOptions map[string]NewResourceOptionLoaderFn

// ResourceOption defines the interface all resource objects must implement in order to be provisioned.
type ResourceOptionLoader interface {
	Load(*pulumi.Context, json.RawMessage) (pulumi.ResourceOption, error)
}
