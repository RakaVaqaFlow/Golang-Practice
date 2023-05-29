package postgresql

import (
	"context"
	"database/sql"

	"gitlab.ozon.dev/workshop/internal/pkg/db"
	"gitlab.ozon.dev/workshop/internal/pkg/repository"
)

type UsersRepo struct {
	db db.DBops
}

func NewUsers(db db.DBops) *UsersRepo {
	return &UsersRepo{db: db}
}

// Add specific user
func (r *UsersRepo) Add(ctx context.Context, user *repository.User) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx, `INSERT INTO users(name) VALUES ($1) RETURNING id`, user.Name).Scan(&id)
	return id, err
}

func (r *UsersRepo) GetById(ctx context.Context, id int64) (*repository.User, error) {
	var u repository.User
	err := r.db.Get(ctx, &u, "SELECT id,name,created_at,updated_at FROM users WHERE id=$1", id)
	if err == sql.ErrNoRows {
		return nil, repository.ErrObjectNotFound
	}
	return &u, err
}

func (r *UsersRepo) List(ctx context.Context) ([]*repository.User, error) {
	users := make([]*repository.User, 0)
	err := r.db.Select(ctx, &users, "SELECT id,name,created_at,updated_at FROM users")
	return users, err
}

func (r *UsersRepo) Update(ctx context.Context, user *repository.User) (bool, error) {
	result, err := r.db.Exec(ctx,
		"UPDATE users SET name = $1 WHERE id = $2", user.Name, user.ID)
	return result.RowsAffected() > 0, err
}
