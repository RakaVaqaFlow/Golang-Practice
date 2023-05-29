package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

func exampleTransaction(ctx context.Context, db *sql.DB) {
	// создаем транзакцию
	tx, err := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted, // указываем в опциях уровень изоляции
		ReadOnly:  false,                  // можем указать, что транзакции только для чтения
	})
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback() // если на любом из этапов произойдет ошибка, то мы откатим изменения при выходе из функции.
	// После вызова tx.Commit() вызов tx.Rollback() ничего уже не откатит, а просто вернет ошибку sql.ErrTxDone

	rows, err := tx.QueryContext(ctx, "SELECT 1") // у sql.Tx все те же селекторы что и у sql.DB
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	_, err = tx.ExecContext(ctx, "SELECT 1")
	if err != nil {
		log.Fatal(err)
	}

	// коммит транзакции
	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("transaction is commited")
}
