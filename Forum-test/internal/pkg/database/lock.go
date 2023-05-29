package database

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type lockType int

// LockTypeOfficeEvents - константа для блокировки таблицы office events
const (
	_ lockType = iota
	LockTypeOfficeEvents
)

// AcquireLock берёт рекомендательную блокировку, которая снимается при завершении транзакции (xact)
func AcquireLock(ctx context.Context, tx *sqlx.Tx, lockID lockType) error {
	_, err := tx.ExecContext(ctx, fmt.Sprintf("select pg_advisory_xact_lock(%d)", lockID))
	return err
}
