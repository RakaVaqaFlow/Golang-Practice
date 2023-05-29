//go:generate mockgen -source ./repository.go -destination=./mocks/repository.go -package=mock_repository
package repository

import (
	"context"
	"errors"
)

var (
	ErrObjectNotFound = errors.New("object not found")
)

type UsersRepo interface {
	Add(ctx context.Context, user *User) (int64, error)
	GetById(ctx context.Context, id int64) (*User, error)
	List(ctx context.Context) ([]*User, error)
	Update(ctx context.Context, user *User) (bool, error)
	Delete(ctx context.Context, id int64) (bool, error)
}

type TasksRepo interface {
	Add(ctx context.Context, task *Task) (int64, error)
	GetById(ctx context.Context, id int64) (*Task, error)
	List(ctx context.Context) ([]*Task, error)
	Update(ctx context.Context, task *Task) (bool, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
