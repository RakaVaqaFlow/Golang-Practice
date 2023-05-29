package api

import (
	"context"
	pb "workshop-8-3/bff/internal/pb/api"
)

type CreateTodoRequest struct {
	Text   string
	UserID string
}

func (c *Client) CreateTodo(ctx context.Context, req CreateTodoRequest) (int, error) {
	resp, err := c.grpc.CreateTodo(ctx, &pb.CreateTodoRequest{
		Text:   req.Text,
		UserId: req.UserID,
	})
	if err != nil {
		return 0, err
	}
	return int(resp.GetId()), nil
}
