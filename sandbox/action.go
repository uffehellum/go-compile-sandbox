package sandbox

// Action is one script with arguments
type Action struct {
	Script     string                 `json:"script"`
	Properties map[string]interface{} `json:"properties"`
}
