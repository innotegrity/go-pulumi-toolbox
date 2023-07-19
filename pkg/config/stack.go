package config

// Stack holds the configuration for the entire stack.
type Stack struct {
	StackAliases StackAliases `json:"stack-aliases"`
	Resources    []Resource   `json:"resources"`
}

// StackAliases are used to shorten stack reference paths to alias names.
type StackAliases map[string]string
