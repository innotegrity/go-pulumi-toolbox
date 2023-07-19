package provisioner

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"go.innotegrity.dev/pulumi-toolbox/pkg/config"
	"go.innotegrity.dev/pulumi-toolbox/pkg/config/resourceoption"
	"go.innotegrity.dev/pulumi-toolbox/pkg/config/resourcetype"
	"go.innotegrity.dev/pulumi-toolbox/pkg/state"
)

// Run parses the Pulumi configuration data and provisions resources that are defined.
func Run(ctx *pulumi.Context, namespace string) error {
	resourceTypes := resourcetype.GetAll()
	resourceOptions := resourceoption.GetAll()

	// load configuration data
	if err := config.Load(ctx, namespace); err != nil {
		return err
	}

	// provision resources
	for _, res := range config.Get().Resources {
		// make sure the resource type is valid
		if _, ok := resourceTypes[res.Type]; !ok {
			return fmt.Errorf("'%s': not a supported resource type", res.Type)
		}
		obj := resourceTypes[res.Type]()
		if err := json.Unmarshal(res.Args, obj); err != nil {
			return err
		}

		// configure any options for the provisioning
		stackState := state.Get()
		opts := []pulumi.ResourceOption{}
		for opt, args := range res.Options {
			if _, ok := resourceOptions[opt]; !ok {
				return fmt.Errorf("'%s': not a supported resource option", opt)
			}
			obj := resourceOptions[opt]()
			option, err := obj.Load(ctx, args)
			if err != nil {
				return err
			}
			opts = append(opts, option)
		}

		// provision the resources
		resource, err := obj.Provision(ctx, opts...)
		if err != nil {
			return err
		}
		stackState.Add(ctx, obj, resource)
	}
	return nil
}
