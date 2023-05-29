package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

// StatementBuilder глобальная переменная с сконфигурированным плейсхолдером для pgsql
var StatementBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

// NewPostgres returns DB
func NewPostgres(ctx context.Context, dsn, driver string) (*sqlx.DB, error) {
	db, err := sqlx.Open(driver, dsn)
	if err != nil {

		return nil, err
	}

	if err = db.Ping(); err != nil {

		return nil, err
	}

	return db, nil
}
