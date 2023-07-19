package digitalocean

import (
	"fmt"

	do "github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"go.innotegrity.dev/pulumi-toolbox/pkg/config"
	"go.innotegrity.dev/pulumi-toolbox/pkg/state"
)

// Project is used for creating a project resource in DigitalOcean.
type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Environment string `json:"environment"`
	Purpose     string `json:"purpose"`
}

// NewProject returns a new, empty object.
func NewProject() config.ResourceTypeProvisioner {
	return &Project{}
}

// GetID returns the ID of the object.
func (p Project) GetID() string {
	return p.ID
}

// GetOutput retrieves an output value for the given property on the given resource.
func (p Project) GetOutput(ctx *pulumi.Context, res pulumi.Resource, property string) (
	pulumi.Output, error) {

	// validate the resource
	resource, ok := res.(*do.Project)
	if !ok {
		return nil, fmt.Errorf("resource provided is not a '%s' resource type", p.GetType())
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
func (p Project) GetType() string {
	return PROJECT_RESOURCE_TYPE
}

// Provision handles provisioning and management of the resource.
func (p Project) Provision(ctx *pulumi.Context, opts ...pulumi.ResourceOption) (
	pulumi.Resource, error) {

	resId := state.MakeResourceID(ctx, p.GetType(), p.GetID())

	ctx.Log.Debug(fmt.Sprintf("provisioning %s", resId), nil)
	res, err := do.NewProject(ctx, p.ID, &do.ProjectArgs{
		Name:        pulumi.String(p.Name),
		Description: pulumi.String(p.Description),
		Environment: pulumi.String(p.Environment),
		Purpose:     pulumi.String(p.Purpose),
	}, opts...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
