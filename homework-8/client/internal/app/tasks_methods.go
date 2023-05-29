package app

import (
	pb "client/internal/pb/api"
	"client/internal/pkg/models"
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func (c *Client) CreateTask(ctx context.Context, req models.Task) (int64, error) {
	tr := otel.Tracer("CreateTask")
	ctx, span := tr.Start(ctx, "client layer")
	span.SetAttributes(attribute.Key("params").String(req.ToString()))
	defer span.End()

	resp, err := c.grpc.CreateTask(ctx, &pb.CreateTaskRequest{
		UserId:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		return 0, err
	}
	return resp.Id, nil
}

func (c *Client) GetTask(ctx context.Context, req int64) (*models.Task, error) {
	resp, err := c.grpc.GetTask(ctx, &pb.GetTaskRequest{Id: req})
	if err != nil {
		return nil, err
	}
	return &models.Task{
		ID:          resp.Task.Id,
		UserID:      resp.Task.UserId,
		Title:       resp.Task.Title,
		Description: resp.Task.Description,
	}, nil
}

func (c *Client) ListTasks(ctx context.Context) ([]*models.Task, error) {
	resp, err := c.grpc.ListTasks(ctx, &pb.GetTasksRequest{})
	if err != nil {
		return nil, err
	}
	tasks := make([]*models.Task, 0, len(resp.Tasks))
	for _, task := range resp.Tasks {
		tasks = append(tasks, &models.Task{
			ID:          task.Id,
			UserID:      task.UserId,
			Title:       task.Title,
			Description: task.Description,
		})
	}
	return tasks, nil
}

func (c *Client) UpdateTask(ctx context.Context, req models.Task) (bool, error) {
	tr := otel.Tracer("UpdateTask")
	ctx, span := tr.Start(ctx, "client layer")
	span.SetAttributes(attribute.Key("params").String(req.ToString()))
	defer span.End()

	resp, err := c.grpc.UpdateTask(ctx, &pb.UpdateTaskRequest{
		Id:          req.ID,
		UserId:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		return false, err
	}
	return resp.Ok, nil
}

func (c *Client) DeleteTask(ctx context.Context, req int64) (bool, error) {
	tr := otel.Tracer("DeleteTask")
	ctx, span := tr.Start(ctx, "client layer")
	span.SetAttributes(attribute.Key("params").Int64(req))
	defer span.End()

	resp, err := c.grpc.DeleteTask(ctx, &pb.DeleteTaskRequest{Id: req})
	if err != nil {
		return false, err
	}
	return resp.Ok, nil
}
