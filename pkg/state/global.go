package state

import (
	"sync"
)

var (
	_once  sync.Once
	_state *State
)

// Get returns the current state of the stack as resources are provisioned.
func Get() *State {
	_once.Do(func() {
		_state = &State{
			resources: map[string]stateResource{},
		}
	})
	return _state
}
