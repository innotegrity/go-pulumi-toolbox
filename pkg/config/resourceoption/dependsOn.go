package resourceoption

import (
	"encoding/json"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"go.innotegrity.dev/pulumi-toolbox/pkg/config"
)

type DependsOnOption struct{}

func NewDependsOnOption() config.ResourceOptionLoader {
	return &DependsOnOption{}
}

func (o *DependsOnOption) Load(ctx *pulumi.Context, data json.RawMessage) (pulumi.ResourceOption, error) {
	return nil, nil
}
