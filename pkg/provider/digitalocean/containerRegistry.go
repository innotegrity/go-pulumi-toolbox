package digitalocean

import (
	"fmt"
	"strings"

	do "github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"go.innotegrity.dev/pulumi-toolbox/pkg/config"
	"go.innotegrity.dev/pulumi-toolbox/pkg/state"
)

// ContainerRegistry is used for creating a container registry resource in DigitalOcean.
type ContainerRegistry struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Region           string `json:"region"`
	SubscriptionTier string `json:"subscriptionTier"`
}

// NewContainerRegistry returns a new, empty object.
func NewContainerRegistry() config.ResourceTypeProvisioner {
	return &ContainerRegistry{}
}

// GetID returns the ID of the object.
func (cr ContainerRegistry) GetID() string {
	return cr.ID
}

// GetOutput retrieves an output value for the given property on the given resource.
func (cr ContainerRegistry) GetOutput(ctx *pulumi.Context, res pulumi.Resource, property string) (
	pulumi.Output, error) {

	// validate the resource
	resource, ok := res.(*do.ContainerRegistry)
	if !ok {
		return nil, fmt.Errorf("resource provided is not a '%s' resource type", cr.GetType())
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
func (cr ContainerRegistry) GetType() string {
	return CONTAINER_REGISTRY_RESOURCE_TYPE
}

// Provision handles provisioning and management of the resource.
func (cr ContainerRegistry) Provision(ctx *pulumi.Context, opts ...pulumi.ResourceOption) (
	pulumi.Resource, error) {

	resId := state.MakeResourceID(ctx, cr.GetType(), cr.GetID())

	ctx.Log.Debug(fmt.Sprintf("provisioning %s", resId), nil)
	res, err := do.NewContainerRegistry(ctx, cr.ID, &do.ContainerRegistryArgs{
		Name:                 pulumi.String(cr.Name),
		Region:               pulumi.String(strings.ToLower(cr.Region)),
		SubscriptionTierSlug: pulumi.String(cr.SubscriptionTier),
	}, opts...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
