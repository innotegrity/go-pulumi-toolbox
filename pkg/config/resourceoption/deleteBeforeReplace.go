package resourceoption

import (
	"encoding/json"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"go.innotegrity.dev/pulumi-toolbox/pkg/config"
)

type DeleteBeforeReplaceOption struct{}

func NewDeleteBeforeReplaceOption() config.ResourceOptionLoader {
	return &DeleteBeforeReplaceOption{}
}

func (o *DeleteBeforeReplaceOption) Load(ctx *pulumi.Context, data json.RawMessage) (pulumi.ResourceOption, error) {
	return nil, nil
}
