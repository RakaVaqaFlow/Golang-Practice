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
}

type UsersRepoCached interface {
	Add(ctx context.Context, user *User) error
	Get(ctx context.Context, id int64) (*User, error)
	GetMulti(ctx context.Context, ids []*int64) ([]*User, error)
}

type UserAbstract interface {
	SomeFunc()
}

type User1 struct {
}

func (u *User1) SomeFunc() {

}

func (u *User2) SomeFunc() {

}

type User2 struct {
}
