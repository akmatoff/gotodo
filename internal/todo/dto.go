package todo

type CreateTodoRequest struct {
	Text string `json:"text"`
}

type UpdateTodoRequest struct {
	Text string `json:"text"`
}
