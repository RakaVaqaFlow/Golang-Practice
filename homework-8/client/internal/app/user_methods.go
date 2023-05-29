package app

import (
	pb "client/internal/pb/api"
	"client/internal/pkg/models"
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func (c *Client) CreateUser(ctx context.Context, req models.User) (int64, error) {
	tr := otel.Tracer("CreateUser")
	ctx, span := tr.Start(ctx, "client layer")
	span.SetAttributes(attribute.Key("params").String(req.ToString()))
	defer span.End()

	resp, err := c.grpc.CreateUser(ctx, &pb.CreateUserRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return 0, err
	}
	return resp.Id, nil
}

func (c *Client) GetUser(ctx context.Context, req int64) (*models.User, error) {
	resp, err := c.grpc.GetUser(ctx, &pb.GetUserRequest{Id: req})
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:       resp.User.Id,
		Name:     resp.User.Name,
		Email:    resp.User.Email,
		Password: resp.User.Password,
	}, nil
}

func (c *Client) ListUsers(ctx context.Context) ([]*models.User, error) {
	resp, err := c.grpc.ListUsers(ctx, &pb.GetUsersRequest{})
	if err != nil {
		return nil, err
	}
	users := make([]*models.User, 0, len(resp.Users))
	for _, user := range resp.Users {
		users = append(users, &models.User{
			ID:       user.Id,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		})
	}
	return users, nil
}

func (c *Client) UpdateUser(ctx context.Context, req models.User) (bool, error) {
	tr := otel.Tracer("UpdateUser")
	ctx, span := tr.Start(ctx, "client layer")
	span.SetAttributes(attribute.Key("params").String(req.ToString()))
	defer span.End()

	resp, err := c.grpc.UpdateUser(ctx, &pb.UpdateUserRequest{
		Id:       req.ID,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return false, err
	}
	return resp.Ok, nil
}

func (c *Client) DeleteUser(ctx context.Context, req int64) (bool, error) {
	tr := otel.Tracer("DeleteUser")
	ctx, span := tr.Start(ctx, "client layer")
	span.SetAttributes(attribute.Key("params").Int64(req))
	defer span.End()

	resp, err := c.grpc.DeleteUser(ctx, &pb.DeleteUserRequest{Id: req})
	if err != nil {
		return false, err
	}
	return resp.Ok, nil
}
