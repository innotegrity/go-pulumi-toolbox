package gcp

import "go.innotegrity.dev/pulumi-toolbox/pkg/config"

// Resource type strings
const (
	NETWORK_RESOURCE_TYPE    = "gcp::Network"
	SUBNETWORK_RESOURCE_TYPE = "gcp::Subnetwork"
)

// GetResourceTypes is responsible for defining the various resource types that are available for provisioning by
// this provider.
func GetResourceTypes() config.ResourceTypes {
	types := config.ResourceTypes{
		NETWORK_RESOURCE_TYPE:    NewNetwork,
		SUBNETWORK_RESOURCE_TYPE: NewSubnetwork,
	}
	return types
}
