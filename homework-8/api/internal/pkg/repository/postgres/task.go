package postgresql

import (
	"context"
	"database/sql"

	"homework/internal/pkg/repository"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type dbOps interface {
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	ExecQueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	GetPool(ctx context.Context) *pgxpool.Pool
}

type TasksRepo struct {
	db dbOps
}

func NewTasks(db dbOps) *TasksRepo {
	return &TasksRepo{db: db}
}

// Add specific task
func (r *TasksRepo) Add(ctx context.Context, task *repository.Task) (int64, error) {
	tr := otel.Tracer("CreateTask")
	ctx, span := tr.Start(ctx, "database layer")
	span.SetAttributes(attribute.Key("params").String(task.ToString()))
	defer span.End()

	var id int64
	err := r.db.ExecQueryRow(ctx, `INSERT INTO tasks(user_id, title, description) VALUES ($1, $2, $3) RETURNING id`, task.UserID, task.Title, task.Description).Scan(&id)
	return id, err
}

// Get info about task by id
func (r *TasksRepo) GetById(ctx context.Context, id int64) (*repository.Task, error) {
	var t repository.Task
	err := r.db.Get(ctx, &t, "SELECT id,user_id,title,description,created_at,updated_at FROM tasks WHERE id=$1", id)
	if err == sql.ErrNoRows {
		return nil, repository.ErrObjectNotFound
	}
	return &t, err
}

// Get list of all tasks
func (r *TasksRepo) List(ctx context.Context) ([]*repository.Task, error) {
	tasks := make([]*repository.Task, 0)
	err := r.db.Select(ctx, &tasks, "SELECT id,user_id,title,description,created_at,updated_at FROM tasks")
	return tasks, err
}

// Update task info
func (r *TasksRepo) Update(ctx context.Context, task *repository.Task) (bool, error) {
	tr := otel.Tracer("UpdateTask")
	ctx, span := tr.Start(ctx, "database layer")
	span.SetAttributes(attribute.Key("params").String(task.ToString()))
	defer span.End()

	result, err := r.db.Exec(ctx,
		"UPDATE tasks SET user_id = $1, title = $2, description = $3, updated_at = now()  WHERE id = $4", task.UserID, task.Title, task.Description, task.ID)
	return result.RowsAffected() > 0, err
}

// Delete task by id
func (r *TasksRepo) Delete(ctx context.Context, id int64) (bool, error) {
	tr := otel.Tracer("DeleteTask")
	ctx, span := tr.Start(ctx, "database layer")
	span.SetAttributes(attribute.Key("params").Int64(id))
	defer span.End()

	result, err := r.db.Exec(ctx, "DELETE FROM tasks WHERE id = $1", id)
	return result.RowsAffected() > 0, err
}
