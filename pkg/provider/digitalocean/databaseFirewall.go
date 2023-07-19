package digitalocean

import (
	"fmt"

	do "github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"go.innotegrity.dev/pulumi-toolbox/pkg/config"
	"go.innotegrity.dev/pulumi-toolbox/pkg/state"
)

// DatabaseFirewall is used for creating a database firewall resource in DigitalOcean.
type DatabaseFirewall struct {
	ID        string                 `json:"id"`
	ClusterID string                 `json:"cluster-id"`
	Rules     []databaseFirewallRule `json:"rules"`
}

// databaseFirewallRule defines a single firewall rule.
type databaseFirewallRule struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// NewDatabaseFirewall returns a new, empty object.
func NewDatabaseFirewall() config.ResourceTypeProvisioner {
	return &DatabaseFirewall{}
}

// GetID returns the ID of the object.
func (df DatabaseFirewall) GetID() string {
	return df.ID
}

// GetOutput retrieves an output value for the given property on the given resource.
func (df DatabaseFirewall) GetOutput(ctx *pulumi.Context, res pulumi.Resource, property string) (
	pulumi.Output, error) {

	// validate the resource
	resource, ok := res.(*do.DatabaseFirewall)
	if !ok {
		return nil, fmt.Errorf("resource provided is not a '%s' resource type", df.GetType())
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
func (df DatabaseFirewall) GetType() string {
	return DATABASE_FIREWALL_RESOURCE_TYPE
}

// Provision handles provisioning and management of the resource.
func (df DatabaseFirewall) Provision(ctx *pulumi.Context, opts ...pulumi.ResourceOption) (
	pulumi.Resource, error) {

	resId := state.MakeResourceID(ctx, df.GetType(), df.GetID())
	stackState := state.Get()

	// lookup cluster ID
	if df.ClusterID == "" {
		return nil, fmt.Errorf("cluster ID is required")
	}
	id, err := stackState.GetOutput(ctx, df.ClusterID)
	if err != nil {
		return nil, err
	}
	doClusterId := id.ApplyT(func(id string) string {
		ctx.Log.Debug(fmt.Sprintf("'%s': project ID retrieved", df.ClusterID), nil)
		return id
	}).(pulumi.StringOutput)

	// parse the rules
	rules := do.DatabaseFirewallRuleArray{}
	for _, rule := range df.Rules {
		switch rule.Type {
		case "ip_addr", "tag":
			rules = append(rules, do.DatabaseFirewallRuleArgs{
				Type:  pulumi.String(rule.Type),
				Value: pulumi.String(rule.Value),
			})
		case "k8s", "droplet", "app":
			// TODO: support k8s, droplet and app
			ctx.Log.Warn(" support coming soon", nil)
		}
	}

	// provision the firewall
	ctx.Log.Debug(fmt.Sprintf("provisioning %s", resId), nil)
	res, err := do.NewDatabaseFirewall(ctx, df.ID, &do.DatabaseFirewallArgs{
		ClusterId: doClusterId,
		Rules:     rules,
	}, opts...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
