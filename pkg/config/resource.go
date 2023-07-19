package config

import (
	"encoding/json"
)

// Resource defines the configuration for a particular resource to be provisioned.
type Resource struct {
	Type    string                     `json:"type"`
	Args    json.RawMessage            `json:"args"`
	Options map[string]json.RawMessage `json:"options"` // TODO: implement resource options
	Exports []string                   `json:"exports"` // TODO: implement exports
}
