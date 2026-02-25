package model

type Task struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Payload string `json:"payload"` //data from client for server to process
}

//each field is serialized as "id, status" etc in JSON responses
