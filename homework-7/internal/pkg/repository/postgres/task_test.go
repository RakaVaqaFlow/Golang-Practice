//go:build integration
// +build integration

package postgresql

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"homework/internal/pkg/db"
	"homework/internal/pkg/repository"
)

func setupTasks() (*TasksRepo, error) {
	ctx := context.Background()
	db, err := db.NewDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating test database: %w", err)
	}
	return NewTasks(db), nil
}

func TestAddTaskAndGetById(t *testing.T) {
	var (
		ctx = context.Background()
	)
	t.Run("success", func(t *testing.T) {
		// Successfully add a task and get it by id
		// arrange
		t.Parallel()
		rt, err := setupTasks()
		require.NoError(t, err)
		ru, err := setupUsers()
		require.NoError(t, err)
		id, err := ru.Add(ctx, &repository.User{})
		require.NoError(t, err)
		task := &repository.Task{
			UserID:      id,
			Title:       "Test",
			Description: "test",
		}
		// act
		id, err = rt.Add(ctx, task)

		assert.NoError(t, err)
		task, err = rt.GetById(ctx, id)
		// assert
		assert.NoError(t, err)
		assert.Equal(t, "Test", task.Title)
		assert.Equal(t, "test", task.Description)
	})

	t.Run("non-existed", func(t *testing.T) {
		// Fail to get a task by id due to non-existent id
		// arrange
		t.Parallel()
		rt, err := setupTasks()
		require.NoError(t, err)
		// act
		_, err = rt.GetById(ctx, int64(-1))
		// assert
		assert.Error(t, err)
	})

	t.Run("invalid userId", func(t *testing.T) {
		// Fail to add a task due to invalid userId
		// arrange
		t.Parallel()
		rt, err := setupTasks()
		require.NoError(t, err)
		task := &repository.Task{
			UserID:      int64(-1),
			Title:       "Test",
			Description: "test",
		}
		// act
		_, err = rt.Add(ctx, task)
		// assert
		assert.Error(t, err)
	})

}

func TestUpdateTask(t *testing.T) {
	var (
		ctx = context.Background()
	)
	t.Run("success", func(t *testing.T) {
		// Successfully update a task
		// arrange
		t.Parallel()
		rt, err := setupTasks()
		require.NoError(t, err)
		ru, err := setupUsers()
		require.NoError(t, err)
		id, err := ru.Add(ctx, &repository.User{Name: "Test", Email: "test@test.ru", Password: "test"})
		require.NoError(t, err)
		task := &repository.Task{
			UserID:      id,
			Title:       "Test",
			Description: "test",
		}
		id, err = rt.Add(ctx, task)
		require.NoError(t, err)
		task.ID = id
		task.Title = "Test2"
		task.Description = "test2"
		// act
		updated, err := rt.Update(ctx, task)
		// assert
		assert.NoError(t, err)
		assert.True(t, updated)
		task, err = rt.GetById(ctx, id)
		assert.NoError(t, err)
		assert.Equal(t, "Test2", task.Title)
		assert.Equal(t, "test2", task.Description)
	})

	t.Run("non-existed", func(t *testing.T) {
		// Fail to update a task due to non-existent id
		// arrange
		t.Parallel()
		rt, err := setupTasks()
		require.NoError(t, err)
		task := &repository.Task{
			ID:          int64(-1),
			Title:       "Test",
			Description: "test",
		}
		// act
		updated, err := rt.Update(ctx, task)
		// assert
		assert.False(t, updated)
	})
}

func TestDeleteTask(t *testing.T) {
	var (
		ctx = context.Background()
	)

	t.Run("success", func(t *testing.T) {
		// Successfully delete a task
		// arrange
		t.Parallel()
		rt, err := setupTasks()
		require.NoError(t, err)
		ru, err := setupUsers()
		require.NoError(t, err)
		id, err := ru.Add(ctx, &repository.User{Name: "Test", Email: "test@some.ru", Password: "test"})
		require.NoError(t, err)
		task := &repository.Task{
			UserID:      id,
			Title:       "Test",
			Description: "test",
		}
		id, err = rt.Add(ctx, task)
		require.NoError(t, err)
		// act
		deleted, err := rt.Delete(ctx, id)
		// assert
		assert.NoError(t, err)
		assert.True(t, deleted)
	})

	t.Run("non-existed", func(t *testing.T) {
		// Fail to delete a task due to non-existent id
		// arrange
		t.Parallel()
		rt, err := setupTasks()
		require.NoError(t, err)
		// act
		deleted, err := rt.Delete(ctx, int64(-1))
		// assert
		assert.False(t, deleted)
	})
}
