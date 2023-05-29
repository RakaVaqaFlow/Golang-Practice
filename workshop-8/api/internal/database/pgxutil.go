package database

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

// pgxUtil обертка для упрощенной работы с pgx.
type pgxUtil struct {
	db *sqlx.DB
}

func NewPGX(db *sqlx.DB) PGX {
	return &pgxUtil{db: db}
}

func (p *pgxUtil) Exec(ctx context.Context, sqlizer Sqlizer) error {
	query, args, err := sqlizer.ToSql()
	if err != nil {
		return fmt.Errorf("ToSql: %w", err)
	}

	_, err = p.db.ExecContext(ctx, query, args...)
	return err
}

func (p *pgxUtil) Select(ctx context.Context, dst interface{}, sqlizer Sqlizer) error {
	query, args, err := sqlizer.ToSql()
	if err != nil {
		return fmt.Errorf("ToSql: %w", err)
	}

	return p.db.SelectContext(ctx, dst, query, args...)
}

func (p *pgxUtil) Get(ctx context.Context, dst interface{}, sqlizer Sqlizer) error {
	query, args, err := sqlizer.ToSql()
	if err != nil {
		return fmt.Errorf("ToSql: %w", err)
	}

	return p.db.GetContext(ctx, dst, query, args...)
}
