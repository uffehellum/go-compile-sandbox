package main

import (
	"encoding/json"
	"fmt"

	//hjson "github.com/hjson/hjson-go"
	_ "github.com/robertkrimen/otto/underscore"
	"github.com/uffehellum/go-compile-sandbox/sandbox"
)

const pipelinesource = `{
	"actions":[
		{
			"script":  	[
				"data = event.Data ;data.k *= properties.factor "	,		
				""],
			"properties": {
				"factor": 7
			}
		},
		{
			"script": ["data = event.Data; data.k += properties.addend"],
			"properties": {
				"addend": 11
			}
		},
		{
			"script": [
				"data = event.Data",
				"data.k += properties.addend"
				],
			"properties": {
				"addend": -11
			}
		},
		{
			"script": ["event.Data.k /= properties.factor"],
			"properties": {
				"factor": 7
			}
		},
		{
			"script": ["event.Data.k += properties.addend"],
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
	pipeline, err := readPipeline(pipelinesource)
	if err != nil {
		fmt.Println(err)
		return
	}
	events, err := readEvents(eventsource)
	if err != nil {
		panic(err)
	}
	vm := sandbox.New()
	for i := 0; i < 10000; i++ {
		vm.ExecuteBatch(events, pipeline)
	}
	for _, e := range events {
		fmt.Println(e.Data["k"])
	}
}

func readEvents(source string) (events []sandbox.Event, err error) {
	events = make([]sandbox.Event, 0)
	err = readJSONObject(source, &events)
	return
}

func readPipeline(source string) (pipeline sandbox.Pipeline, err error) {
	err = readJSONObject(source, &pipeline)
	return
}

func readJSONObject(source string, v interface{}) error {
	return json.Unmarshal([]byte(source), v)
}

// func readHjsonObject(source string, v interface{}) error {
// 	var m interface{}
// 	err := hjson.Unmarshal([]byte(source), &m)
// 	if err != nil {
// 		return err
// 	}
// 	bytes, err := json.Marshal(m)
// 	if err != nil {
// 		return err
// 	}
// 	err = json.Unmarshal(bytes, v)
// 	return err
// }
