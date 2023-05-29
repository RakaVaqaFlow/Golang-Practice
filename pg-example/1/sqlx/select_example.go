package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func exampleSelect(ctx context.Context, db *sqlx.DB) {
	// sqlx.DB.Select эквивалент sql.DB.QueryContext (но сокращенный)

	const (
		minAge = 18
		query  = "SELECT first_name, last_name, age FROM students WHERE age >= $1"
	)

	var students []Student
	// Select записывает в students массив полученных строк.
	if err := db.SelectContext(ctx, &students, query, minAge); err != nil {
		log.Fatal(err) // handle the error here
	}

	fmt.Printf("students: %v\n", students)
}
