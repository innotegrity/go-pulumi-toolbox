package resourceoption

import "go.innotegrity.dev/pulumi-toolbox/pkg/config"

// Option strings
const (
	DELETE_BEFORE_REPLACE = "deleteBeforeReplace"
	DEPENDS_ON            = "dependsOn"
)

// GetAll returns the various resource options that are available during resource provisioning.
func GetAll() config.ResourceOptions {
	options := config.ResourceOptions{
		DELETE_BEFORE_REPLACE: NewDeleteBeforeReplaceOption,
		DEPENDS_ON:            NewDependsOnOption,
	}
	return options
}
