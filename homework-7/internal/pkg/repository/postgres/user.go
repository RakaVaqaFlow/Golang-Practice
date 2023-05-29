package postgresql

import (
	"context"
	"database/sql"

	"homework/internal/pkg/repository"
)

type UsersRepo struct {
	db dbOps
}

func NewUsers(db dbOps) *UsersRepo {
	return &UsersRepo{db: db}
}

// Add specific user
func (r *UsersRepo) Add(ctx context.Context, user *repository.User) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx, `INSERT INTO users(name, email, password) VALUES ($1, $2, $3) RETURNING id`, user.Name, user.Email, user.Password).Scan(&id)
	return id, err
}

// Get info about user by id
func (r *UsersRepo) GetById(ctx context.Context, id int64) (*repository.User, error) {
	var u repository.User
	err := r.db.Get(ctx, &u, "SELECT id,name,email,password,created_at,updated_at FROM users WHERE id=$1", id)
	if err == sql.ErrNoRows {
		return nil, repository.ErrObjectNotFound
	}
	return &u, err
}

// Get list of all users
func (r *UsersRepo) List(ctx context.Context) ([]*repository.User, error) {
	users := make([]*repository.User, 0)
	err := r.db.Select(ctx, &users, "SELECT id,name,email,password,created_at,updated_at FROM users")
	return users, err
}

// Update user info
func (r *UsersRepo) Update(ctx context.Context, user *repository.User) (bool, error) {
	result, err := r.db.Exec(ctx,
		"UPDATE users SET name = $1, email = $2, password = $3, updated_at = now() WHERE id = $4", user.Name, user.Email, user.Password, user.ID)
	return result.RowsAffected() > 0, err
}

// Delete user by id
func (r *UsersRepo) Delete(ctx context.Context, id int64) (bool, error) {
	result, err := r.db.Exec(ctx, "DELETE FROM users WHERE id = $1", id)
	return result.RowsAffected() > 0, err
}
