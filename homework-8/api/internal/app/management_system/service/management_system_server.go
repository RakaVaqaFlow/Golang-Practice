package service

import (
	"context"
	"homework/internal/pb"
	"homework/internal/pkg/repository"

	"homework/internal"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type Implementation struct {
	pb.UnimplementedManagementSystemSeviceServer

	userRepo repository.UsersRepo
	taskRepo repository.TasksRepo
}

func NewImplementation(userRepo repository.UsersRepo, taskRepo repository.TasksRepo) *Implementation {
	return &Implementation{
		userRepo: userRepo,
		taskRepo: taskRepo,
	}
}

func (i *Implementation) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	tr := otel.Tracer("CreateUser")
	ctx, span := tr.Start(ctx, "received request")
	span.SetAttributes(attribute.Key("params").String(request.String()))
	defer span.End()

	user := &repository.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	id, err := i.userRepo.Add(ctx, user)
	if err != nil {
		return nil, err
	}

	internal.RegUserCounter.Add(1)

	return &pb.CreateUserResponse{Id: id}, nil
}

func (i *Implementation) CreateTask(ctx context.Context, request *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	tr := otel.Tracer("CreateTask")
	ctx, span := tr.Start(ctx, "received request")
	span.SetAttributes(attribute.Key("params").String(request.String()))
	defer span.End()

	task := &repository.Task{
		UserID:      request.UserId,
		Title:       request.Title,
		Description: request.Description,
	}
	id, err := i.taskRepo.Add(ctx, task)
	if err != nil {
		return nil, err
	}

	internal.RegTaskCounter.Add(1)

	return &pb.CreateTaskResponse{Id: id}, nil
}

func (i *Implementation) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	id := request.Id
	user, err := i.userRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{User: &pb.User{
		Id:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}}, nil
}

func (i *Implementation) GetTask(ctx context.Context, request *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	id := request.Id
	task, err := i.taskRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &pb.GetTaskResponse{Task: &pb.Task{
		Id:          task.ID,
		UserId:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
	}}, nil
}

func (i *Implementation) ListUsers(ctx context.Context, request *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	users, err := i.userRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*pb.User, 0, len(users))
	for _, user := range users {
		result = append(result, &pb.User{
			Id:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		})
	}
	return &pb.GetUsersResponse{Users: result}, nil
}

func (i *Implementation) ListTasks(ctx context.Context, request *pb.GetTasksRequest) (*pb.GetTasksResponse, error) {
	tasks, err := i.taskRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*pb.Task, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, &pb.Task{
			Id:          task.ID,
			UserId:      task.UserID,
			Title:       task.Title,
			Description: task.Description,
		})
	}
	return &pb.GetTasksResponse{Tasks: result}, nil
}

func (i *Implementation) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	tr := otel.Tracer("UpdateUser")
	ctx, span := tr.Start(ctx, "received request")
	span.SetAttributes(attribute.Key("params").String(request.String()))
	defer span.End()

	user := &repository.User{
		ID:       request.Id,
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	status, err := i.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	internal.UpdatedUserCounter.Add(1)

	return &pb.UpdateUserResponse{Ok: status}, nil
}

func (i *Implementation) UpdateTask(ctx context.Context, request *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	tr := otel.Tracer("UpdateTask")
	ctx, span := tr.Start(ctx, "received request")
	span.SetAttributes(attribute.Key("params").String(request.String()))
	defer span.End()

	task := &repository.Task{
		ID:          request.Id,
		UserID:      request.UserId,
		Title:       request.Title,
		Description: request.Description,
	}
	status, err := i.taskRepo.Update(ctx, task)
	if err != nil {
		return nil, err
	}

	internal.UpdatedTaskCounter.Add(1)
	return &pb.UpdateTaskResponse{Ok: status}, nil
}

func (i *Implementation) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	tr := otel.Tracer("DeleteUser")
	ctx, span := tr.Start(ctx, "received request")
	span.SetAttributes(attribute.Key("params").String(request.String()))
	defer span.End()

	id := request.Id
	status, err := i.userRepo.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	internal.DeletedUserCounter.Add(1)

	return &pb.DeleteUserResponse{Ok: status}, nil
}

func (i *Implementation) DeleteTask(ctx context.Context, request *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	tr := otel.Tracer("DeleteTask")
	ctx, span := tr.Start(ctx, "received request")
	span.SetAttributes(attribute.Key("params").String(request.String()))
	defer span.End()

	id := request.Id
	status, err := i.taskRepo.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	internal.DeletedTaskCounter.Add(1)

	return &pb.DeleteTaskResponse{Ok: status}, nil
}
