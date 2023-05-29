package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func exampleQuery(ctx context.Context, pool *pgxpool.Pool) {
	// API библиотеки pgx схож с API стандартной библиотеки database/sql
	// pgxpool.Pool.Query эквивалент sql.DB.QueryContext

	const query = `SELECT first_name, last_name, age FROM students`

	rows, err := pool.Query(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var s Student
		if err = rows.Scan(&s.FirstName, &s.LastName, &s.Age); err != nil {
			log.Fatal(err)
		}
		students = append(students, s)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(students)
}
