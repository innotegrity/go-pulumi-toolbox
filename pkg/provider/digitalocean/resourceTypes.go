package digitalocean

import "go.innotegrity.dev/pulumi-toolbox/pkg/config"

// Resource type strings
const (
	CONTAINER_REGISTRY_RESOURCE_TYPE = "digitalocean::ContainerRegistry"
	PROJECT_RESOURCE_TYPE            = "digitalocean::Project"
	VPC_RESOURCE_TYPE                = "digitalocean::VPC"
	DATABASE_CLUSTER_RESOURCE_TYPE   = "digitalocean::DatabaseCluster"
	DATABASE_FIREWALL_RESOURCE_TYPE  = "digitalocean::DatabaseFirewall"
	PROJECT_RESOURCES_RESOURCE_TYPE  = "digitalocean::ProjectResources"
	KUBERNETES_CLUSTER_RESOURCE_TYPE = "digitalocean::KubernetesCluster"
)

// GetResourceTypes is responsible for defining the various resource types that are available for provisioning by
// this provider.
func GetResourceTypes() config.ResourceTypes {
	types := config.ResourceTypes{
		CONTAINER_REGISTRY_RESOURCE_TYPE: NewContainerRegistry,
		PROJECT_RESOURCE_TYPE:            NewProject,
		VPC_RESOURCE_TYPE:                NewVPC,
		DATABASE_CLUSTER_RESOURCE_TYPE:   NewDatabaseCluster,
		DATABASE_FIREWALL_RESOURCE_TYPE:  NewDatabaseFirewall,
		/*
			PROJECT_RESOURCES_RESOURCE_TYPE:  NewProjectResources,
			KUBERNETES_CLUSTER_RESOURCE_TYPE: NewKubernetesCluster,
		*/
	}
	return types
}
