package resourcetype

import (
	"go.innotegrity.dev/pulumi-toolbox/pkg/config"
	"go.innotegrity.dev/pulumi-toolbox/pkg/provider/digitalocean"
	"go.innotegrity.dev/pulumi-toolbox/pkg/provider/gcp"
	"golang.org/x/exp/maps"
)

// GetAll returns all of the available resource types that are available for provisioning.
func GetAll() config.ResourceTypes {
	types := config.ResourceTypes{}
	maps.Copy(types, digitalocean.GetResourceTypes())
	maps.Copy(types, gcp.GetResourceTypes())
	return types
}
