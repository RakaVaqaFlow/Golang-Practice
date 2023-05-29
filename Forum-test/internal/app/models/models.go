package models

type User struct {
	Nickname string `db:"nickname"`
	Email    string `db:"email"`
	Fullname string `db:"full_name"`
	About    string `db:"about"`
}

type UserUpdate struct {
	About    string `json:"about"    db:"about"`
	Email    string `json:"email"    db:"email"`
	Fullname string `json:"fullname" db:"fullname"`
}

type NewForum struct {
	Slug  string `json:"slug"     db:"slug"`
	Title string `json:"title"    db:"title"`
	User  string `json:"user"     db:"user_nick"`
}

type Forum struct {
	Posts   int64  `json:"posts"    db:"posts"`
	Slug    string `json:"slug"     db:"slug"`
	Threads int32  `json:"threads"  db:"threads"`
	Title   string `json:"title"    db:"title"`
	User    string `json:"user"     db:"user_nick"`
}

type Thread struct {
	Author  string `json:"author"  db:"author"`
	Created string `json:"created" db:"created"`
	Forum   string `json:"forum"   db:"forum"`
	Id      int32  `json:"id"      db:"id"`
	Message string `json:"message" db:"message"`
	Slug    string `json:"slug,omitempty"    db:"slug"`
	Title   string `json:"title"   db:"title"`
	Votes   int32  `json:"votes"   db:"votes"`
}

type Post struct {
	Author   string `json:"author"   db:"author"`
	Created  string `json:"created"  db:"created"`
	Forum    string `json:"forum"    db:"forum"`
	Id       int64  `json:"id"       db:"id"`
	IsEdited bool   `json:"isEdited" db:"isEdited"`
	Message  string `json:"message"  db:"message"`
	Parent   int64  `json:"parent"   db:"parent"`
	Thread   int32  `json:"thread"   db:"thread"`
}

type PostAccount struct {
	Author *User   `json:"author,omitempty"`
	Forum  *Forum  `json:"forum,omitempty"`
	Post   Post    `json:"post"`
	Thread *Thread `json:"thread,omitempty"`
}

type Status struct {
	Forum  int64 `json:"forum"`
	Post   int64 `json:"post"`
	Thread int64 `json:"thread"`
	User   int64 `json:"user"`
}

type Vote struct {
	Nickname string `json:"nickname"`
	Voice    int    `json:"voice"`
}
