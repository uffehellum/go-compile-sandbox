package sandbox

import (
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore" // give all scripts access to lodash
)

var vm *otto.Otto

// ExecuteBatch runs all pipelines on all events
func ExecuteBatch(events []Event, pipeline Pipeline) {
	if vm == nil {
		vm = otto.New()
	}
	// defer otto.Close(vm)
	for _, event := range events {
		vm.Set("event", event)
		for _, action := range pipeline.Actions {
			vm.Set("properties", action.Properties)
			vm.Run(action.Script)
		}
	}
}
