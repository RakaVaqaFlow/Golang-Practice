package repository

import (
	"database/sql"
	"encoding/json"
	"time"
)

type User struct {
	ID        int64        `db:"id" json:"id"`
	Name      string       `db:"name" json:"name"`
	Email     string       `db:"email" json:"email"`
	Password  string       `db:"password" json:"password"`
	CreatedAt time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at" json:"updated_at"`
}

type Task struct {
	ID          int64        `db:"id" json:"id"`
	UserID      int64        `db:"user_id" json:"user_id"`
	Title       string       `db:"title" json:"title"`
	Description string       `db:"description" json:"description"`
	CreatedAt   time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at" json:"updated_at"`
}

func (t *Task) ToString() string {
	st, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return string(st)
}

func (u *User) ToString() string {
	st, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(st)
}
