package digitalocean

import (
	"fmt"
	"strings"

	do "github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"go.innotegrity.dev/pulumi-toolbox/pkg/config"
)

// VPC_HANDLER_TYPE defines the type name of this resource.
const VPC_HANDLER_TYPE = "digitalocean::VPC"

// VPC is used for creating a VPC resource in DigitalOcean.
type VPC struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IpRange     string `json:"ip-range"`
	Region      string `json:"region"`
}

// NewVPC returns a new, empty object.
func NewVPC() config.ResourceProvisioner {
	return &VPC{}
}

// Provision handles provisioning and management of the resource.
func (v VPC) Provision(ctx *pulumi.Context, refs config.StackReferences) error {
	ctx.Log.Debug(fmt.Sprintf("provisioning %s", VPC_RESOURCE_TYPE), nil)

	if _, err := do.NewVpc(ctx, v.ID, &do.VpcArgs{
		Name:        pulumi.String(v.Name),
		Description: pulumi.String(v.Description),
		IpRange:     pulumi.String(v.IpRange),
		Region:      pulumi.String(strings.ToLower(v.Region)),
	}, pulumi.DeleteBeforeReplace(true)); err != nil {
		return err
	}
	return nil
}
