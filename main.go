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
		},
		{
			"script": "data = event.Data; data.k += properties.addend",
			"properties": {
				"addend": -11
			}
		},
		{
			"script": "data = event.Data; data.k /= properties.factor",
			"properties": {
				"factor": 7
			}
		},
		{
			"script": "data = event.Data; data.k += properties.addend",
			"properties": {
				"addend": 0.0001
			}
		}
	]
}`

const eventsource = `[
	{
		"data": {
			"a": "value a",
			"b": "value b",
			"k": 1
		},
		"message":""
	},	
	{
		"data": {
			"a": "value a",
			"b": "value b",
			"k": 2
		},
		"message":""
	},	
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
	for i := 0; i < 10000; i++ {
		sandbox.ExecuteBatch(events, pipeline)
	}
	for _, e := range events {
		fmt.Println(e.Data["k"])
	}
}
