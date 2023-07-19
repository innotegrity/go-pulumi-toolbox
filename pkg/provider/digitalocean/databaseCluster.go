package digitalocean

import (
	"fmt"
	"sort"
	"strings"

	do "github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"go.innotegrity.dev/pulumi-toolbox/pkg/config"
	"go.innotegrity.dev/pulumi-toolbox/pkg/state"
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
func NewDatabaseCluster() config.ResourceTypeProvisioner {
	return &DatabaseCluster{}
}

// GetID returns the ID of the object.
func (dc DatabaseCluster) GetID() string {
	return dc.ID
}

// GetOutput retrieves an output value for the given property on the given resource.
func (dc DatabaseCluster) GetOutput(ctx *pulumi.Context, res pulumi.Resource, property string) (
	pulumi.Output, error) {

	// validate the resource
	resource, ok := res.(*do.DatabaseCluster)
	if !ok {
		return nil, fmt.Errorf("resource provided is not a '%s' resource type", dc.GetType())
	}

	// return the output for the given property
	switch property {
	case "ID":
		return resource.ID(), nil
	default:
		return nil, fmt.Errorf("'%s': unknown resource property", property)
	}
}

// GetType returns the type of the object.
func (dc DatabaseCluster) GetType() string {
	return DATABASE_CLUSTER_RESOURCE_TYPE
}

// Provision handles provisioning and management of the resource.
func (dc DatabaseCluster) Provision(ctx *pulumi.Context, opts ...pulumi.ResourceOption) (
	pulumi.Resource, error) {

	resId := state.MakeResourceID(ctx, dc.GetType(), dc.GetID())
	stackState := state.Get()

	// lookup project ID, if supplied
	var doProjectID pulumi.StringOutput
	if dc.ProjectID != "" {
		id, err := stackState.GetOutput(ctx, dc.ProjectID)
		if err != nil {
			return nil, err
		}
		doProjectID = id.ApplyT(func(id string) string {
			ctx.Log.Debug(fmt.Sprintf("'%s': project ID retrieved", dc.ProjectID), nil)
			return id
		}).(pulumi.StringOutput)
	}

	// lookup VPC ID, if supplied
	var doPrivateNetworkUUID pulumi.StringOutput
	if dc.PrivateNetworkID != "" {
		id, err := stackState.GetOutput(ctx, dc.PrivateNetworkID)
		if err != nil {
			return nil, err
		}
		doPrivateNetworkUUID = id.ApplyT(func(id string) string {
			ctx.Log.Debug(fmt.Sprintf("'%s': VPC ID retrieved", dc.PrivateNetworkID), nil)
			return id
		}).(pulumi.StringOutput)
	}

	// set up maintenance windows and sort tags (for consistency)
	maintenanceWindows := do.DatabaseClusterMaintenanceWindowArray{}
	for _, mw := range dc.MaintenanceWindows {
		maintenanceWindows = append(maintenanceWindows, do.DatabaseClusterMaintenanceWindowArgs{
			Hour: pulumi.String(mw.Hour),
			Day:  pulumi.String(mw.Day),
		})
	}
	sort.Strings(dc.Tags)

	// provision the cluster
	ctx.Log.Debug(fmt.Sprintf("provisioning %s", resId), nil)
	res, err := do.NewDatabaseCluster(ctx, dc.ID, &do.DatabaseClusterArgs{
		Name:               pulumi.String(dc.Name),
		Engine:             pulumi.String(dc.Engine),
		Version:            pulumi.String(dc.Version),
		NodeCount:          pulumi.Int(dc.NodeCount),
		Size:               pulumi.String(dc.Size),
		MaintenanceWindows: maintenanceWindows,
		PrivateNetworkUuid: doPrivateNetworkUUID,
		ProjectId:          doProjectID,
		Region:             pulumi.String(strings.ToLower(dc.Region)),
		// TODO: Tags appears to be broken for now and causes constant state updates
		//Tags:               pulumi.ToStringArray(dc.Tags),
	}, opts...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
