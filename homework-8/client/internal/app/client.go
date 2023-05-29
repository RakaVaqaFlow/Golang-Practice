package app

import (
	pb "client/internal/pb/api"
	"context"

	"google.golang.org/grpc"
)

type Client struct {
	grpc pb.ManagementSystemSeviceClient
	conn *grpc.ClientConn
}

// NewClient создает Client.
func NewClient(ctx context.Context, target string) (*Client, error) {
	conn, err := grpc.DialContext(ctx, target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{
		grpc: pb.NewManagementSystemSeviceClient(conn),
		conn: conn,
	}, nil
}

// Close закрывает соединение.
func (c *Client) Close() error {
	return c.conn.Close()
}
