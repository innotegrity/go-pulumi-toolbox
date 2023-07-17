package digitalocean

import (
	"fmt"

	do "github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"go.innotegrity.dev/pulumi-toolbox/pkg/config"
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
func NewProject() config.ResourceProvisioner {
	return &Project{}
}

// Provision handles provisioning and management of the resource.
func (p Project) Provision(ctx *pulumi.Context, refs config.StackReferences) error {
	ctx.Log.Debug(fmt.Sprintf("provisioning %s", PROJECT_RESOURCE_TYPE), nil)

	if _, err := do.NewProject(ctx, p.ID, &do.ProjectArgs{
		Name:        pulumi.String(p.Name),
		Description: pulumi.String(p.Description),
		Environment: pulumi.String(p.Environment),
		Purpose:     pulumi.String(p.Purpose),
	}, pulumi.DeleteBeforeReplace(true)); err != nil {
		return err
	}
	return nil
}
