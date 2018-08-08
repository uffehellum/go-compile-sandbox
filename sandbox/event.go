package sandbox

//Event is in pipeline
type Event struct {
	Message interface{}            `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
