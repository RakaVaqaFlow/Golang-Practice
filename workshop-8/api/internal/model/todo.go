package model

type CreateInput struct {
	UserID string
	Text   string
}

type Todo struct {
	ID     int
	UserID string
	Text   string
}

type Pagination struct {
	Page  int
	Limit int
}
