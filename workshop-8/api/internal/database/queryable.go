package database

import (
	"context"
)

// PGX содержит основные операции для работы с базой данных.
type PGX interface {
	Queryable
}

// Queryable содержит основные операции для query-инга db.
type Queryable interface {
	Exec(ctx context.Context, sqlizer Sqlizer) error
	Get(ctx context.Context, dst interface{}, sqlizer Sqlizer) error
	Select(ctx context.Context, dst interface{}, sqlizer Sqlizer) error
}

// Sqlizer возвращает sql-запрос и его аргументы.
type Sqlizer interface {
	ToSql() (sql string, args []interface{}, err error)
}
