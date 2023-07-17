package provisioner

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pulumiConfig "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
	"go.innotegrity.dev/pulumi-toolbox/pkg/config"
)

// Run parses the Pulumi configuration data and provisions resources that are defined.
func Run(ctx *pulumi.Context, namespace string) error {
	// perform initialization steps
	resTypes := initResourceTypes()

	// load configuration data
	var stackConfig config.Stack
	c := pulumiConfig.New(ctx, namespace)
	if err := c.TryObject("stack", &stackConfig); err != nil {
		return err
	}

	// provision resources
	for _, res := range stackConfig.Resources {
		if _, ok := resTypes[res.Type]; !ok {
			return fmt.Errorf("'%s' is not a supported resource type", res.Type)
		}
		obj := resTypes[res.Type]()
		if err := json.Unmarshal(res.Params, obj); err != nil {
			return err
		}
		if err := obj.Provision(ctx, stackConfig.References); err != nil {
			return err
		}
	}
	return nil
}
