package api

import (
	"context"
	pb "workshop-8-3/bff/internal/pb/api"
)

type ListTodoRequest struct {
	Pagination struct {
		Page  int
		Limit int
	}
	UserID string
}

type Todo struct {
	ID   int
	Text string
}

func (c *Client) ListTodo(ctx context.Context, req ListTodoRequest) ([]*Todo, error) {
	resp, err := c.grpc.ListTodo(ctx, &pb.ListTodoRequest{
		UserId: req.UserID,
		Pagination: &pb.ListTodoRequest_Pagination{
			Page:  uint32(req.Pagination.Page),
			Limit: uint32(req.Pagination.Limit),
		},
	})
	if err != nil {
		return nil, err
	}
	todos := make([]*Todo, 0, len(resp.Todos))
	for _, todo := range resp.Todos {
		todos = append(todos, &Todo{
			ID:   int(todo.Id),
			Text: todo.Text,
		})
	}

	return todos, nil
}
