package in_memory

import (
	"context"

	"gitlab.ozon.dev/workshop/internal/pkg/repository"
)

type UsersRepo struct {
	users map[int64]*repository.User
}

func NewUsers() *UsersRepo {
	mapUsers := make(map[int64]*repository.User)
	return &UsersRepo{users: mapUsers}
}

// Add specific user
func (r *UsersRepo) Add(ctx context.Context, user *repository.User) (int64, error) {
	lastId := int64(len(r.users))
	r.users[lastId+1] = user
	return lastId + 1, nil
}

func (r *UsersRepo) GetById(ctx context.Context, id int64) (*repository.User, error) {
	var u repository.User

	return &u, nil
}

func (r *UsersRepo) List(ctx context.Context) ([]*repository.User, error) {
	users := make([]*repository.User, 0)
	return users, nil
}

func (r *UsersRepo) Update(ctx context.Context, user *repository.User) (bool, error) {

	return true, nil
}
