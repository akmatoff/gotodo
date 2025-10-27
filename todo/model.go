package todo

import "time"

type Todo struct {
	Id          int       `json:"id"`
	Text        string    `json:"text"`
	IsCompleted bool      `json:"isCompleted"`
	CreatedAt   time.Time `json:"createdAt"`
	CompletedAt time.Time `json:"completedAt"`
}

func (t *Todo) Complete() {
	t.IsCompleted = true
	t.CompletedAt = time.Now()
}

func (t *Todo) Uncomplete() {
	t.IsCompleted = false
	t.CompletedAt = time.Time{}
}
