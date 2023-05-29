package todo

import (
	"context"
	"fmt"
	"workshop-8-3/api/internal/database"
	"workshop-8-3/api/internal/model"
)

type Searcher struct{}

func NewSearcher() *Searcher {
	return &Searcher{}
}

func (*Searcher) Search(
	ctx context.Context,
	q database.Queryable,
	pagination model.Pagination,
) ([]*model.Todo, error) {

	query := database.PSQL.
		Select("id",
			"user_id",
			"text",
		).
		Limit(uint64(pagination.Limit)).
		Offset(uint64((pagination.Page - 1) * pagination.Limit)).
		From(database.TableTodo)

	var dto []*todoDTO
	err := q.Select(ctx, &dto, query)
	if err != nil {
		return nil, fmt.Errorf("получение todo: %w", err)
	}
	result := make([]*model.Todo, 0, len(dto))
	for _, o := range dto {
		result = append(result, &model.Todo{
			ID:     int(o.Id),
			UserID: o.UserID,
			Text:   o.Text,
		})
	}

	return result, nil
}
