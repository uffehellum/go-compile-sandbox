package main

import (
	"fmt"

	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
)

//Event is in pipeline
type Event struct {
	Message interface{}
	Data    map[string]interface{}
}

// Properties is in a plugin
type Properties struct {
	script string,
	scriptProperties map[string]interface{},
}

func main() {
	data := map[string]interface{}{
		"a": "value a",
		"b": "value b",
		"k": 42,
	}
	event := Event{Data: data, Message: "message string"}
	properties := Properties {
		script: `
			console.log('hej otto - ' + event.Data.a)
			event.Data.k *= properties.factor
			data = event.Data
			data.name = 'bob'
		`,
		scriptProperties: map[string]interface{}{
			"factor": 3,
		},
	}
	vm := otto.New()
	vm.Set("event", event)
	vm.Set("properties", properties.scriptProperties)
	vm.Run(properties.script)
	fmt.Println("Hej " + data["name"].(string))
	fmt.Printf("Answer = %d\n", int(data["k"].(float64)))
}
