package db

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TxWrapper struct {
	pgxTx pgx.Tx
}

func (db Database) PrepareTX() {

}

func (t *TxWrapper) Commit(ctx context.Context) error {
	return t.pgxTx.Commit(ctx)
}

func (t *TxWrapper) Rollback(ctx context.Context) error {
	return t.pgxTx.Rollback(ctx)
}

func (t *TxWrapper) GetPool(_ context.Context) *pgxpool.Pool {
	return nil
}

func (t *TxWrapper) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, t.pgxTx, dest, query, args...)
}

func (t *TxWrapper) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, t.pgxTx, dest, query, args...)
}

func (t *TxWrapper) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return t.pgxTx.Exec(ctx, query, args...)
}

func (t *TxWrapper) ExecQueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return t.pgxTx.QueryRow(ctx, query, args...)
}
