package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func exampleQueryx(ctx context.Context, db *sqlx.DB) {
	const minAge = 18

	rows, err := db.QueryxContext(ctx, "SELECT first_name, last_name, age FROM students WHERE age >= $1", minAge)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() // Обязательно закрываем иначе соединение с БД повиснет

	var students []Student
	for rows.Next() {
		var st Student

		if err := rows.StructScan(&st); err != nil {
			log.Fatal(err)
		}

		students = append(students, st)
	}

	if err = rows.Err(); err != nil {
		// handle the error here
		log.Fatal(err)
	}

	fmt.Printf("students: %v\n", students)
}
