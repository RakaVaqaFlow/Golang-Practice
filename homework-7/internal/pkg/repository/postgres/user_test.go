//go:build integration
// +build integration

package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"homework/internal/pkg/db"
	"homework/internal/pkg/repository"
)

var testDB *sql.DB

func setupUsers() (*UsersRepo, error) {
	ctx := context.Background()
	db, err := db.NewDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating test database: %w", err)
	}
	return NewUsers(db), nil
}

func TestAddUserAndGetById(t *testing.T) {
	var (
		ctx = context.Background()
	)
	t.Run("success", func(t *testing.T) {
		// Successfully add a user and get it by id

		//arange
		t.Parallel()
		r, err := setupUsers()
		require.NoError(t, err)
		user := &repository.User{
			Name:     "Test",
			Email:    "some@test.ru",
			Password: "test",
		}

		//act
		id, err := r.Add(ctx, user)
		assert.NoError(t, err)

		user, err = r.GetById(ctx, id)

		//assert
		assert.NoError(t, err)
		assert.Equal(t, "Test", user.Name)
		assert.Equal(t, "some@test.ru", user.Email)
		assert.Equal(t, "test", user.Password)
	})

	t.Run("fail", func(t *testing.T) {
		// Fail to get a user by id due to non-existent id

		//arange
		t.Parallel()
		r, err := setupUsers()
		require.NoError(t, err)

		//act
		_, err = r.GetById(ctx, int64(-1))

		//assert
		assert.Error(t, err)
	})

}

func TestUpdateUser(t *testing.T) {
	var (
		ctx = context.Background()
	)
	t.Run("success", func(t *testing.T) {
		// Successfully update a user
		//arange
		t.Parallel()
		r, err := setupUsers()
		assert.NoError(t, err)
		user := &repository.User{
			Name:     "Jaden Smith",
			Email:    "jadensmith@example.com",
			Password: "secret",
		}
		id, err := r.Add(ctx, user)
		assert.NoError(t, err)
		user.Password = "password"
		user.ID = id

		// act
		updated, err := r.Update(ctx, user)

		// assert
		assert.NoError(t, err)
		assert.True(t, updated)
		user, err = r.GetById(ctx, id)
		assert.NoError(t, err)
		assert.Equal(t, "password", user.Password)
	})

	t.Run("nothing to update", func(t *testing.T) {
		// Trying to update a non-existent user
		// arange
		t.Parallel()
		r, err := setupUsers()
		assert.NoError(t, err)

		// act
		updated, err := r.Update(ctx, &repository.User{ID: -1})

		// assert
		assert.False(t, updated)
	})
}

func TestDeleteUser(t *testing.T) {
	var (
		ctx = context.Background()
	)
	t.Run("success", func(t *testing.T) {
		// Successfully delete a user
		// arange
		t.Parallel()
		r, err := setupUsers()
		assert.NoError(t, err)

		user := &repository.User{
			Name:     "Will Smith",
			Email:    "willsmith@example.com",
			Password: "secret",
		}
		id, err := r.Add(ctx, user)
		assert.NoError(t, err)

		// act
		deleted, err := r.Delete(ctx, id)

		// assert
		assert.NoError(t, err)
		assert.True(t, deleted)

		_, err = r.GetById(ctx, id)
		assert.Error(t, err)
	})

	t.Run("nothing to delete", func(t *testing.T) {
		// Trying to delete a non-existent user
		// arange
		t.Parallel()
		r, err := setupUsers()
		assert.NoError(t, err)

		// act
		deleted, err := r.Delete(ctx, int64(-1))

		// assert
		assert.NoError(t, err)
		assert.False(t, deleted)
	})
}
