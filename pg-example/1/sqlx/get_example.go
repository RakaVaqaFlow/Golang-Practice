package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func exampleGet(ctx context.Context, db *sqlx.DB) {
	// sqlx.DB.Get эквивалент sql.DB.QueryRowContext

	var st Student
	// Get записывает в st данные из первой строки
	if err := db.GetContext(ctx, &st, "SELECT first_name, last_name,age FROM students LIMIT 1"); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Printf("no students")
		} else {
			log.Fatal(err) // handle the error here
		}
	}

	fmt.Printf("student: %v\n", st)
}
