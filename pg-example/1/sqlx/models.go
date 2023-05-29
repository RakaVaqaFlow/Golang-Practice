package main

type Student struct {
	FirstName  string `db:"first_name"`
	LastName   string `db:"last_name"`
	Age        uint   `db:"age"`
	OtherField string `db:"-"`
}
