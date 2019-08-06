package queue

//Task ...
type Task struct {
	QueueName string  `json:"queue_name"`
	Message   Message `json:"message"`
}

//Message ...
type Message struct {
	Priority uint8             `json:"priority"`
	Token    string            `json:"token"`
	Name     string            `json:"name"`
	Params   map[string]string `json:"params"`
}
