package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func exampleNamedQuery(ctx context.Context, db *sqlx.DB) {
	st := Student{
		FirstName: "Bob",
		LastName:  "Brown",
	}

	const query = `
		SELECT 
			first_name, 
			last_name, 
			age 
		FROM students 
		WHERE first_name=:first_name 
		   OR last_name=:last_name
	`
	rows, err := db.NamedQueryContext(ctx, query, st)
	if err != nil {
		log.Fatal(err) // handle the error here
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var st Student
		if err := rows.StructScan(&st); err != nil {
			log.Fatal(err) // handle the error here
		}

		students = append(students, st)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err) // handle the error here
	}

	fmt.Printf("students: %v\n", students)
}
