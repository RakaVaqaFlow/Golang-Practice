package main

type Student struct {
	ID         int64  `db:"id"`
	FirstName  string `db:"first_name"`
	LastName   string `db:"last_name"`
	Age        uint   `db:"age"`
	OtherField string `db:"-"` // '-' не использовать это поле при мапинге полей в запросах
}
