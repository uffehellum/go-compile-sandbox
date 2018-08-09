package sandbox

import (
	"strings"

	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore" // give all scripts access to lodash
)

// Sandbox is a separate ES5 vm
type Sandbox struct {
	otto *otto.Otto
}

// New creates a sandbox
func New() *Sandbox {
	return &Sandbox{
		otto: otto.New(),
	}
}

// ExecuteBatch runs all pipelines on all events
func (vm *Sandbox) ExecuteBatch(events []Event, pipeline Pipeline) {
	for _, event := range events {
		vm.otto.Set("event", event)
		for _, action := range pipeline.Actions {
			vm.otto.Set("properties", action.Properties)
			vm.otto.Run(strings.Join(action.Script, "\n"))
		}
	}
}
