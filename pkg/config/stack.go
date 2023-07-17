package config

// Stack holds the configuration settings for the stack.
type Stack struct {
	References StackReferences `json:"stack-references"`
	Resources  []Resource      `json:"resources"`
}

// StackReferences maps a short name to an actual stack name for use in referencing resources in other stacks.
type StackReferences map[string]string
