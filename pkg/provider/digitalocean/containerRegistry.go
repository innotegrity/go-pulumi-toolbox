package digitalocean

import (
	"fmt"
	"strings"

	do "github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"go.innotegrity.dev/pulumi-toolbox/pkg/config"
)

// ContainerRegistry is used for creating a container registry resource in DigitalOcean.
type ContainerRegistry struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Region           string `json:"region"`
	SubscriptionTier string `json:"subscription-tier"`
}

// NewContainerRegistry returns a new, empty object.
func NewContainerRegistry() config.ResourceProvisioner {
	return &ContainerRegistry{}
}

// Provision handles provisioning and management of the resource.
func (cr ContainerRegistry) Provision(ctx *pulumi.Context, refs config.StackReferences) error {
	ctx.Log.Debug(fmt.Sprintf("provisioning %s", CONTAINER_REGISTRY_RESOURCE_TYPE), nil)

	if _, err := do.NewContainerRegistry(ctx, cr.ID, &do.ContainerRegistryArgs{
		Name:                 pulumi.String(cr.Name),
		Region:               pulumi.String(strings.ToLower(cr.Region)),
		SubscriptionTierSlug: pulumi.String(cr.SubscriptionTier),
	}, pulumi.DeleteBeforeReplace(true)); err != nil {
		return err
	}
	return nil
}
