package models

import "encoding/json"

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) ToString() string {
	st, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(st)
}
