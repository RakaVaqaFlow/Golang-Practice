package todo

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"workshop-8-3/api/internal/database"
)

type Deleter struct{}

func NewDeleter() *Deleter {
	return &Deleter{}
}

func (*Deleter) Delete(
	ctx context.Context,
	q database.Queryable,
	id int,
) (bool, error) {
	tr := otel.Tracer("DeleteTodo")
	ctx, span := tr.Start(ctx, "database layer")
	span.SetAttributes(attribute.Key("params").Int(id))
	defer span.End()

	query := database.PSQL.Delete(database.TableTodo).Where(squirrel.Eq{"id": id})

	err := q.Exec(ctx, query)
	if err != nil {
		return false, fmt.Errorf("удаление todo: %w", err)
	}

	return true, nil
}
