package api

import (
	"context"
	pb "workshop-8-3/bff/internal/pb/api"
)

// DeleteByID ...
func (c *Client) DeleteByID(ctx context.Context, id int64) (bool, error) {
	todoResponse, err := c.grpc.DeleteTodo(ctx, &pb.DeleteTodoRequest{Id: uint32(id)})
	if err != nil {
		return false, err
	}
	return todoResponse.Ok, nil
}
