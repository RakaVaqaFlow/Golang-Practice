package todo

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"workshop-8-3/api/internal/database"
	"workshop-8-3/api/internal/model"
)

// Creator записывает в базу транзакции
type Creator struct{}

// NewCreator создает новый экземпляр Creator
func NewCreator() *Creator {
	return &Creator{}
}

// Create записывает в базу транзакции
func (*Creator) Create(
	ctx context.Context,
	q database.Queryable,
	input model.CreateInput,
) (int64, error) {
	tr := otel.Tracer("database")
	_, span := tr.Start(ctx, "entering database layer")
	span.SetAttributes(attribute.Key("params").String(input.Text))
	defer span.End()

	query := database.PSQL.
		Insert(database.TableTodo).
		Columns(
			"user_id",
			"text",
		).
		Values(
			input.UserID,
			input.Text,
		).
		Suffix("RETURNING id")

	var id int64
	err := q.Get(ctx, &id, query)
	if err != nil {
		return id, fmt.Errorf("запись todo: %w", err)
	}

	return id, nil
}
