package digitalocean

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"go.innotegrity.dev/pulumi-toolbox/pkg/config"
)

// DatabaseCluster is used for creating a database cluster resource in DigitalOcean.
type DatabaseCluster struct {
	ID                 string                             `json:"id"`
	Name               string                             `json:"name"`
	Engine             string                             `json:"engine"`
	Version            string                             `json:"version"`
	NodeCount          uint                               `json:"node-count"`
	Size               string                             `json:"size"`
	MaintenanceWindows []databaseClusterMaintenanceWindow `json:"maintenance-windows"`
	PrivateNetworkID   string                             `json:"vpc-id"`
	ProjectID          string                             `json:"project-id"`
	Region             string                             `json:"region"`
	Tags               []string                           `json:"tags"`
}

// databaseClusterMaintenanceWindow defines the maintenance window for a DB cluster.
type databaseClusterMaintenanceWindow struct {
	Day  string `json:"day"`
	Hour string `json:"hour"`
}

// NewDatabaseCluster returns a new, empty object.
func NewDatabaseCluster() config.ResourceProvisioner {
	return &DatabaseCluster{}
}

// Provision handles provisioning and management of the resource.
func (dc DatabaseCluster) Provision(ctx *pulumi.Context, refs config.StackReferences) error {
	ctx.Log.Debug(fmt.Sprintf("provisioning %s", DATABASE_CLUSTER_RESOURCE_TYPE), nil)

	// lookup project ID, if supplied

	// lookup private network ID, if supplied

	// provision the cluster

	return nil
}
