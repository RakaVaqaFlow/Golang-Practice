package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

func exampleQuery(db *sql.DB) {
	const minAge = 18

	rows, err := db.Query("SELECT first_name, last_name, age FROM students WHERE age >= $1", minAge)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() // Обязательно закрываем иначе соединение с БД повиснет

	type Student struct {
		FirstName string
		LastName  string
		Age       uint
	}

	var students []Student
	for rows.Next() {
		var st Student
		if err := rows.Scan(&st.FirstName, &st.LastName, &st.Age); err != nil {
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

func exampleQueryContext(ctx context.Context, db *sql.DB) {
	const minAge = 18

	rows, err := db.QueryContext(ctx, "SELECT first_name, last_name, age FROM students WHERE age >= $1", minAge)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() // Обязательно закрываем иначе соединение с БД повиснет

	type Student struct {
		FirstName string
		LastName  string
		Age       uint
	}

	var students []Student
	for rows.Next() {
		var st Student
		if err := rows.Scan(&st.FirstName, &st.LastName, &st.Age); err != nil {
			log.Fatal(err)
		}
		students = append(students, st)
	}
	// Внутри драйвера мы получаем данные, накапливая их в буфер размером 4KB.
	// rows.Next() порождает поход в сеть и наполняет буфер. Если буфера не хватает,
	// то мы идём в сеть за оставшимися данными. Больше походов в сеть – меньше скорость обработки.

	fmt.Printf("students: %v\n", students)
}
