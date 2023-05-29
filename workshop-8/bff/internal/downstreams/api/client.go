package api

import (
	"context"
	"google.golang.org/grpc"
	pb "workshop-8-3/bff/internal/pb/api"
)

// Client ...
type Client struct {
	grpc pb.TodoServiceClient
	conn *grpc.ClientConn
}

// NewClient создает Client.
func NewClient(ctx context.Context, target string) (*Client, error) {
	conn, err := grpc.DialContext(ctx, target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{
		grpc: pb.NewTodoServiceClient(conn),
		conn: conn,
	}, nil
}

// Close закрывает соединение.
func (c *Client) Close() error {
	return c.conn.Close()
}
