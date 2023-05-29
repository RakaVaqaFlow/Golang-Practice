package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // импортируем драйвер для postgres
)

const (
	// название регистрируемоего драйвера github.com/lib/pq
	stdPostgresDriverName = "postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "test"
	dbname   = "test"
)

func clear(c io.Closer) {
	if err := c.Close(); err != nil {
		fmt.Println("close:", err)
	}
}

func main() {
	// connection string
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// 1
	sqlDB, err := sql.Open(stdPostgresDriverName, psqlConn) // returns *sql.DB, error
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	// Можем обернуть существующий *sql.DB в *sqlx.DB; нужно обязательно передать тот же драйвер, иначе не заведется
	sqlxDB := sqlx.NewDb(sqlDB, stdPostgresDriverName)
	clear(sqlxDB) // дальше не использую, поэтому просто для примера

	// 2
	// А можем сразу создавать *sqlx.DB с помощью функции sqlx.Connect
	// sqlx.Connect = sql.Open + sql.Ping (2 в одном :))
	db, err := sqlx.Connect(stdPostgresDriverName, psqlConn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // так же не забываем освободить ресурсы

	// 3
	// А если мы ленивые и не хотим обрабатывать ошибку, хотим сразу падать
	sqlxDB = sqlx.MustConnect(stdPostgresDriverName, psqlConn)
	clear(sqlxDB) // дальше не использую

	fmt.Println("Connection with database successfully established!")
	/*
		sqlx реализует стандартные функции: database/sql (ведут себя точно так же как стандартные, нет никаких преимуществ):
			* Exec(...) (sql.Result, error) - unchanged from database/sql
			* Query(...) (*sql.Rows, error) - unchanged from database/sql
			* QueryRow(...) *sql.Row - unchanged from database/sql

		Расширенные стандартные функции (позволяют использовать фичи sqlx):
			* MustExec() sql.Result -- Exec, but panic on error // не надо использовать в проде
			* Queryx(...) (*sqlx.Rows, error) - Query, but return an sqlx.Rows
			* QueryRowx(...) *sqlx.Row -- QueryRow, but return an sqlx.Row

		Новые:
			* Get(dest interface{}, ...) error
			* Select(dest interface{}, ...) error
			* NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	*/
	ctx := context.Background()

	exampleQueryx(ctx, db)
	exampleGet(ctx, db)
	exampleSelect(ctx, db)
	exampleNamedQuery(ctx, db)

	// See more: https://jmoiron.github.io/sqlx/
}
