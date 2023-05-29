package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func exampleQueryRow(db *sql.DB) {
	row := db.QueryRow("SELECT count(*) FROM students")

	var totalStudents uint
	if err := row.Scan(&totalStudents); err != nil { // Обязательно передаем адрес переменной, куда будем сканировать значение.
		log.Fatal(err)
	}

	fmt.Printf("total students: %d\n", totalStudents)
}

func exampleQueryRowContext(ctx context.Context, db *sql.DB) {
	row := db.QueryRowContext(ctx, "SELECT count(*) FROM students")

	var totalStudents uint
	if err := row.Scan(&totalStudents); err != nil { // Обязательно передаем адрес переменной, куда будем сканировать значение.
		log.Fatal(err)
	}

	fmt.Printf("total students: %d\n", totalStudents)
}

func exampleQueryRowNoRows(db *sql.DB) {
	var studentID int64
	_ = fmt.Sprintf("SELECT id FROM students WHERE age = %d", 1000)    // так нельзя
	_ = "SELECT id FROM students WHERE age = %d" + "1000 "             // и так делать не стоит
	row := db.QueryRow("SELECT id FROM students WHERE age = $1", 1000) // такого "долгожителя" в нашей таблице может не быть
	if errors.Is(row.Err(), sql.ErrNoRows) {
		fmt.Println("Не найден в БД студент с age > 10000")
	}

	if err := row.Scan(&studentID); err != nil { // мы тут получим ошибку, так как нам ничего не вернулось из БД
		fmt.Println("db.QueryRow.Scan():", err) // нам вернется ошибка sql.ErrNoRows
		if errors.Is(err, sql.ErrNoRows) {      // при использовании QueryRow не забывайте обрабатывать ошибку на sql.ErrNoRows, так как отстуствие результата может быть стандартным кейсом
			fmt.Println("Не найден в БД студент с age > 10000")
		}
	}
}
