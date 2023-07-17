package provisioner

import (
	"go.innotegrity.dev/pulumi-toolbox/pkg/config"
	"go.innotegrity.dev/pulumi-toolbox/pkg/provider/digitalocean"
	"golang.org/x/exp/maps"
)

// initResourceTypes loads all of the available resource types that are available for provisioning.
func initResourceTypes() config.ResourceTypes {
	types := config.ResourceTypes{}
	maps.Copy(types, digitalocean.InitResourceTypes())
	return types
}
