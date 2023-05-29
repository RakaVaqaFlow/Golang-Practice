package todo

type todoDTO struct {
	Id     int64  `db:"id"`
	UserID string `db:"user_id"`
	Text   string `db:"text"`
}
