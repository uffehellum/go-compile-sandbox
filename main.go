package main

import (
	"encoding/json"
	"fmt"

	_ "github.com/robertkrimen/otto/underscore"
	"github.com/uffehellum/go-compile-sandbox/sandbox"
)

const pipelinesource = `{
	"actions":[
		{
			"script": "data = event.Data; data.k *= properties.factor",
			"properties": {
				"factor": 7
			}
		},
		{
			"script": "data = event.Data; data.k += properties.addend",
			"properties": {
				"addend": 11
			}
		}
	]
}`

const eventsource = `[
	{
		"data": {
			"a": "value a",
			"b": "value b",
			"k": 3
		},
		"message":""
	},
	{
		"data": {
			"a": "value a",
			"b": "value b",
			"k": 4
		},
		"message":""
	}
]`

// Properties is in a plugin
func main() {
	pipeline := sandbox.Pipeline{}
	err := json.Unmarshal([]byte(pipelinesource), &pipeline)
	if err != nil {
		fmt.Println(err)
		return
	}
	events := make([]sandbox.Event, 0)
	err = json.Unmarshal([]byte(eventsource), &events)
	if err != nil {
		panic(err)
	}
	sandbox.ExecuteBatch(events, pipeline)
	for _, e := range events {
		fmt.Println(e.Data["k"])
	}
	// data := map[string]interface{}{}
	// event := Event{Data: data, Message: "message string"}
	// properties := Properties{
	// 	script: `
	// 		// console.log('hej otto - ' + event.Data.a)
	// 		event.Data.k *= properties.factor
	// 		data = event.Data
	// 		data.name = 'bob'
	// 	`,
	// 	scriptProperties: map[string]interface{}{
	// 		"factor": 3,
	// 	},
	// }
	// vm := otto.New()
	// vm.Set("event", event)
	// vm.Set("properties", properties.scriptProperties)
	// for i := 1; i <= 1000; i++ {
	// 	properties.scriptProperties["factor"] = i
	// 	vm.Run(properties.script)
	// }
	// fmt.Println("Hej " + data["name"].(string))
	// fmt.Printf("Answer = %f\n", data["k"])
}
