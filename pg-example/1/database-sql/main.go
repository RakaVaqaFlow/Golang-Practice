package main

import (
	"context"
	"database/sql" // https://go.dev/src/database/sql/doc.txt
	"encoding/json"
	"fmt"
	"log"
	"time"

	// database/sql — это набор интерфейсов для работы с базой
	// Чтобы эти интерфейсы работали, для них нужна реализация. Именно за реализацию и отвечают драйверы.

	_ "github.com/lib/pq" // импортируем драйвер для postgres
	// Обратите внимание, что мы загружаем драйвер анонимно, присвоив его квалификатору пакета псевдоним, _ ,
	// чтобы ни одно из его экспортированных имен не было видно нашему коду.
	// Под капотом драйвер регистрирует себя как доступный для пакета database/sql с помощью функции init()
)

const (
	// название регистрируемоего драйвера github.com/lib/pq
	stdPostgresDriverName = "postgres"
	/*
		PostgreSQL:
			* github.com/lib/pq -> postgres
			* github.com/jackc/pgx -> pgx
		MySQL:
			* github.com/go-sql-driver/mysql -> mysql
		SQLite3:
			* github.com/mattn/go-sqlite3 -> sqlite3
		Oracle:
			* github.com/godror/godror -> godror
		MS SQL:
			* github.com/denisenkom/go-mssqldb -> sqlserver

		See more drivers: https://zchee.github.io/golang-wiki/SQLDrivers/
	*/
)

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "test"
	dbname   = "test"
)

func main() {
	// connection string
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open(stdPostgresDriverName, psqlConn) // returns *sql.DB, error
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Обязательно при завершении работы приложения мы должны освободить все ресурсы, иначе соединения к базе останутся висеть.

	/*
		sql.DB - не является соединением с базой данных! Это абстракция интерфейса.

		sql.DB выполняет некоторые важные задачи для вас за кулисами:
		 	* открывает и закрывает соединения с фактической базовой базой данных через драйвер.
			* управляет пулом соединений по мере необходимости.

		Абстракция sql.DB предназначена для того, чтобы вы не беспокоились о том, как управлять одновременным
		доступом к базовому хранилищу данных. Соединение помечается как используемое, когда вы используете
		его для выполнения задачи, а затем возвращается в доступный пул, когда оно больше не используется.
	*/

	// После установления соединеия пингуем базу. Проверяем, что она отвечает нашему приложению.
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection with database successfully established!")

	/* Настройка пула соединений */
	db.SetConnMaxIdleTime(time.Minute) // время, в течение которого соединение может быть бездействующим.
	db.SetConnMaxLifetime(time.Hour)   // время, в течение которого соединение может быть повторно использовано.
	db.SetMaxIdleConns(2)              // максимум 2 простаивающих соединения
	db.SetMaxOpenConns(4)              // максимум 4 открытых соединений с БД

	/* статистика пула соединений */
	statistics := db.Stats()
	bytes, err := json.Marshal(statistics)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("db connection statistics: %s\n", string(bytes))

	/* примеры работы c БД */
	exampleQueryRow(db)
	exampleQueryRowNoRows(db)
	exampleQuery(db)
	exampleExec(db)

	/* примеры работы c БД c контекстом */
	// Совет: используйте запросы с контекстом
	ctx := context.Background()

	exampleQueryRowContext(ctx, db)
	exampleQueryContext(ctx, db)
	exampleExecContext(ctx, db)

	/* примеры работы c транзакциями */
	exampleTransaction(ctx, db)

	/* пример работы с nullable полями*/
	exampleWithNullableFields(ctx, db)

	// Более подробный туториал (правда там с MySQL, но суть та же)
	// Go database/sql tutorial: http://go-database-sql.org/
}
