package gcp

import (
	"fmt"
	"strings"

	googlecompute "github.com/pulumi/pulumi-google-native/sdk/go/google/compute/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"go.innotegrity.dev/pulumi-toolbox/pkg/config"
	"go.innotegrity.dev/pulumi-toolbox/pkg/state"
)

// Network is used for creating a network resource.
type Network struct {
	ID                                    string `json:"id"`
	AutoCreateSubnets                     bool   `json:"autoCreateSubnets"`
	Description                           string `json:"description"`
	EnableUlaInternalIPv6                 bool   `json:"enableUlaInternalIPv6"`
	InternalIPv6Range                     string `json:"internalIPv6Range"`
	MTU                                   int    `json:"mtu"`
	Name                                  string `json:"name"`
	NetworkFirewallPolicyEnforcementOrder string `json:"networkFirewallPolicyEnforcementOrder"`
	Project                               string `json:"project"`
	RequestID                             string `json:"requestID"`
	RoutingMode                           string `json:"routingMode"`
}

// NewNetwork returns a new, empty object.
func NewNetwork() config.ResourceTypeProvisioner {
	return &Network{}
}

// GetID returns the ID of the object.
func (n Network) GetID() string {
	return n.ID
}

// GetOutput retrieves an output value for the given property on the given resource.
func (n Network) GetOutput(ctx *pulumi.Context, res pulumi.Resource, property string) (
	pulumi.Output, error) {

	// validate the resource
	resource, ok := res.(*googlecompute.Network)
	if !ok {
		return nil, fmt.Errorf("resource provided is not a '%s' resource type", n.GetType())
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
func (n Network) GetType() string {
	return NETWORK_RESOURCE_TYPE
}

// Provision handles provisioning and management of the resource.
func (n Network) Provision(ctx *pulumi.Context, opts ...pulumi.ResourceOption) (
	pulumi.Resource, error) {

	resId := state.MakeResourceID(ctx, n.GetType(), n.GetID())

	// initialize known arguments without defaults
	args := googlecompute.NetworkArgs{
		AutoCreateSubnetworks: pulumi.Bool(n.AutoCreateSubnets),
		Description:           pulumi.String(n.Description),
		EnableUlaInternalIpv6: pulumi.Bool(n.EnableUlaInternalIPv6),
		Name:                  pulumi.String(n.Name),
	}

	// only pass arguments which were specified in the config
	if n.InternalIPv6Range != "" {
		args.InternalIpv6Range = pulumi.String(n.InternalIPv6Range)
	}
	if n.MTU != 0 {
		args.Mtu = pulumi.Int(n.MTU)
	}
	if n.NetworkFirewallPolicyEnforcementOrder != "" {
		args.NetworkFirewallPolicyEnforcementOrder = googlecompute.NetworkNetworkFirewallPolicyEnforcementOrder(
			strings.ToUpper(n.NetworkFirewallPolicyEnforcementOrder))
	}
	if n.Project != "" {
		args.Project = pulumi.String(n.Project)
	}
	if n.RequestID != "" {
		args.RequestId = pulumi.String(n.RequestID)
	}
	if n.RoutingMode != "" {
		args.RoutingConfig = googlecompute.NetworkRoutingConfigArgs{
			RoutingMode: googlecompute.NetworkRoutingConfigRoutingMode(strings.ToUpper(n.RoutingMode)),
		}
	}

	// provision the resource
	ctx.Log.Debug(fmt.Sprintf("provisioning %s", resId), nil)
	res, err := googlecompute.NewNetwork(ctx, n.ID, &args, opts...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
