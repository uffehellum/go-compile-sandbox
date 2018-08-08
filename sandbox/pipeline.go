package sandbox

// Pipeline contains a sequence of actions changing an event
type Pipeline struct {
	Actions []Action `json:"actions"`
}
