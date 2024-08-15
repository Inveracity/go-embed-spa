package events

type Msg struct {
	Msg       string `json:"msg"`
	EventType string `json:"eventtype"`
}
