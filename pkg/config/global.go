package config

import (
	"sync"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pulumiConfig "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var (
	_mutex sync.Mutex
	_stack *Stack
)

// Load simply loads the configuration settings into the global configuration object.
func Load(ctx *pulumi.Context, namespace string) error {
	if _stack == nil {
		_mutex.Lock()
		defer _mutex.Unlock()

		// load the configuration
		_stack = &Stack{}
		if err := pulumiConfig.New(ctx, namespace).TryObject("stack", _stack); err != nil {
			_stack = nil
			return err
		}
	}
	return nil
}

// Get returns the stack configuration.
//
// You must call Load() prior to making any calls to this function, otherwise it will return nil.
func Get() *Stack {
	return _stack
}
