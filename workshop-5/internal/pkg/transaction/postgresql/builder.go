package postgresql

import (
	"context"

	"gitlab.ozon.dev/workshop/internal/pkg/db"
	"gitlab.ozon.dev/workshop/internal/pkg/repository/postgresql"
	"gitlab.ozon.dev/workshop/internal/pkg/transaction"
)

type ServiceTxBuidler struct {
	db db.PGX
}

func NewServiceTxBuilder(db db.PGX) *ServiceTxBuidler {
	return &ServiceTxBuidler{db: db}
}

func (f *ServiceTxBuidler) ServiceTx(ctx context.Context) (*transaction.ServiceTx, error) {
	tx, err := f.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &transaction.ServiceTx{
		Tx:    tx,
		Users: postgresql.NewUsers(tx),
	}, nil
}
