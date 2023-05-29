package models

import "encoding/json"

type Task struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (t *Task) ToString() string {
	st, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return string(st)
}
