package digitalocean

import (
	"fmt"
	"strings"

	do "github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"go.innotegrity.dev/pulumi-toolbox/pkg/config"
	"go.innotegrity.dev/pulumi-toolbox/pkg/state"
)

// VPC is used for creating a VPC resource in DigitalOcean.
type VPC struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IpRange     string `json:"ip-range"`
	Region      string `json:"region"`
}

// NewVPC returns a new, empty object.
func NewVPC() config.ResourceTypeProvisioner {
	return &VPC{}
}

// GetID returns the ID of the object.
func (v VPC) GetID() string {
	return v.ID
}

// GetOutput retrieves an output value for the given property on the given resource.
func (v VPC) GetOutput(ctx *pulumi.Context, res pulumi.Resource, property string) (
	pulumi.Output, error) {

	// validate the resource
	resource, ok := res.(*do.Vpc)
	if !ok {
		return nil, fmt.Errorf("resource provided is not a '%s' resource type", v.GetType())
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
func (v VPC) GetType() string {
	return VPC_RESOURCE_TYPE
}

// Provision handles provisioning and management of the resource.
func (v VPC) Provision(ctx *pulumi.Context, opts ...pulumi.ResourceOption) (
	pulumi.Resource, error) {

	resId := state.MakeResourceID(ctx, v.GetType(), v.GetID())

	ctx.Log.Debug(fmt.Sprintf("provisioning %s", resId), nil)
	res, err := do.NewVpc(ctx, v.ID, &do.VpcArgs{
		Name:        pulumi.String(v.Name),
		Description: pulumi.String(v.Description),
		IpRange:     pulumi.String(v.IpRange),
		Region:      pulumi.String(strings.ToLower(v.Region)),
	}, opts...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
