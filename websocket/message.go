package websocket

type Message struct {
	Task string `json:"taskId"`
	Status string `json:"status"`
}