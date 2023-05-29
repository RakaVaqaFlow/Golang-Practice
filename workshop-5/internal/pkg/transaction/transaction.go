package transaction

import (
	"context"

	"gitlab.ozon.dev/workshop/internal/pkg/db"
	"gitlab.ozon.dev/workshop/internal/pkg/repository"
)

type ServiceTxBuilder interface {
	ServiceTx(ctx context.Context) (*ServiceTx, error)
}

type ServiceTx struct {
	db.Tx
	Users repository.UsersRepo
}
