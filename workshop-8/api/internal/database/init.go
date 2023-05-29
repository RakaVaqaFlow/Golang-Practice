package database

import (
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func New(uri string) (*sqlx.DB, error) {
	config, err := pgx.ParseURI(uri)
	if err != nil {
		return nil, err
	}
	nativeDB := stdlib.OpenDB(config)
	db := sqlx.NewDb(nativeDB, "pgx")

	return db, db.Ping()
}
