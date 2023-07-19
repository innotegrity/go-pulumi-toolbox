package state

import (
	"fmt"
	"strings"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"go.innotegrity.dev/pulumi-toolbox/pkg/config"
)

// State holds the current state of the stack as resources are provisioned.
type State struct {
	resources map[string]stateResource
}

type stateResource struct {
	resource pulumi.Resource
	object   config.ResourceTypeProvisioner
}

// MakeResourceID creates a resource ID for the given resource type with the given unique ID.
func MakeResourceID(ctx *pulumi.Context, resType, id string) string {
	return fmt.Sprintf("%s/%s/%s/%s::%s", ctx.Organization(), ctx.Project(), ctx.Stack(), resType, id)
}

// Add inserts a provisioned stack resource object into the current state for later retrieval.
func (s *State) Add(ctx *pulumi.Context, obj config.ResourceTypeProvisioner, res pulumi.Resource) {
	resId := MakeResourceID(ctx, obj.GetType(), obj.GetID())
	ctx.Log.Debug(fmt.Sprintf("%s: storing resource", resId), nil)
	s.resources[resId] = stateResource{
		resource: res,
		object:   obj,
	}
}

// GetResource returns a provisioned stack resource object from the current state.
func (s *State) GetResource(ctx *pulumi.Context, resType, resId string) (pulumi.Resource, error) {
	r, ok := s.resources[MakeResourceID(ctx, resType, resId)]
	if !ok {
		return nil, fmt.Errorf("'%s': resource not found in the current state", resId)
	}
	return r.resource, nil
}

// GetOutput returns an output object from either a resource in the current state or another stack.
func (s *State) GetOutput(ctx *pulumi.Context, outputPropertyId string) (pulumi.Output, error) {
	parts := strings.SplitN(outputPropertyId, "/", 4)
	numParts := len(parts)

	switch numParts {
	case 4: // retrieve output from another stack or this state using fully-qualified stack name
		return s.getStackOutput(ctx, fmt.Sprintf("%s/%s/%s", parts[0], parts[1], parts[2]), parts[3])
	case 2: // retrieve output from another stack or this state using a stack alias
		if parts[0] == "." {
			return s.getLocalOutput(ctx, parts[1])
		}
		aliases := config.Get().StackAliases
		stack, ok := aliases[parts[0]]
		if !ok {
			return nil, fmt.Errorf("'%s': unknown stack alias", parts[0])
		}
		return s.getStackOutput(ctx, stack, parts[1])
	case 1: // retrieve output from this state
		return s.getLocalOutput(ctx, outputPropertyId)
	default:
		return nil, fmt.Errorf("'%s': invalid output identifier", outputPropertyId)
	}
}

// getLocalOutput retrieves a resource from the local state and returns an output.
func (s *State) getLocalOutput(ctx *pulumi.Context, outputPropertyId string) (pulumi.Output, error) {
	parts := strings.Split(outputPropertyId, "::")
	if len(parts) != 4 {
		return nil, fmt.Errorf("'%s': local output property ID should have 4 parts", outputPropertyId)
	}
	resId := MakeResourceID(ctx, fmt.Sprintf("%s::%s", parts[0], parts[1]), parts[2])

	outputProperty := parts[3]
	r, ok := s.resources[resId]
	if !ok {
		return nil, fmt.Errorf("'%s': resource could not be found", resId)
	}
	return r.object.GetOutput(ctx, r.resource, outputProperty)
}

// getStackOutput retrieves a resource from a referenced stack's output.
func (s *State) getStackOutput(ctx *pulumi.Context, stack string, outputPropertyId string) (pulumi.Output, error) {
	thisStack := fmt.Sprintf("%s/%s/%s", ctx.Organization(), ctx.Project(), ctx.Stack())
	if stack == thisStack {
		return s.getLocalOutput(ctx, outputPropertyId)
	}
	stackRef, err := pulumi.NewStackReference(ctx, stack, nil)
	if err != nil {
		return nil, err
	}
	return stackRef.GetOutput(pulumi.String(outputPropertyId)), nil
}
