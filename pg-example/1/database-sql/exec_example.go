package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

func exampleExec(db *sql.DB) {
	// Ex. 1:
	const notExistedStudentID = 1234567
	result, err := db.Exec("UPDATE students SET age = age+1 WHERE id = $1", notExistedStudentID)
	if err != nil {
		log.Fatal(err)
	}

	var (
		rowsAffected, lastInsertId int64
	)

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		fmt.Println("sql.Result.RowsAffected():", err) // ok
	}
	lastInsertId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("sql.Result.LastInsertId():", err) // LastInsertId is not supported by "postgres" driver
	}
	fmt.Printf("rows affected: %d, last insert id: %d\n", rowsAffected, lastInsertId)
}

func exampleExecContext(ctx context.Context, db *sql.DB) {
	const studentID = 1
	result, err := db.ExecContext(ctx, "UPDATE students SET age = age+1 WHERE id = $1", studentID)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("sql.Result.RowsAffected():", err)
	}
	fmt.Printf("rows affected: %d\n", rowsAffected)
}
