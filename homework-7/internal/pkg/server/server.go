package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"homework/internal/pkg/repository"
)

// serverUser - struct for user data
type serverUser struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// serverTask - struct for task data
type serverTask struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type server struct {
	userRepo repository.UsersRepo
	taskRepo repository.TasksRepo
}

func CreateServer(ctx context.Context, ur repository.UsersRepo, tr repository.TasksRepo) *http.ServeMux {
	serv := server{
		userRepo: ur,
		taskRepo: tr,
	}
	serveMux := http.NewServeMux()

	// user handler
	serveMux.HandleFunc("/user", func(res http.ResponseWriter, req *http.Request) {
		if req == nil {
			return
		}

		switch req.Method {
		case http.MethodGet:
			data, status := serv.getUser(ctx, req)
			res.WriteHeader(status)
			res.Write(data)
		case http.MethodPost:
			_, status := serv.createUser(ctx, req)
			res.WriteHeader(status)
		case http.MethodPut:
			_, status := serv.updateUser(ctx, req)
			res.WriteHeader(status)
		case http.MethodDelete:
			_, status := serv.deleteUser(ctx, req)
			res.WriteHeader(status)
		default:
			fmt.Printf("unsupported method: [%s]", req.Method)
			res.WriteHeader(http.StatusNotImplemented)
		}
	})

	// task handler
	serveMux.HandleFunc("/task", func(res http.ResponseWriter, req *http.Request) {
		if req == nil {
			return
		}

		switch req.Method {
		case http.MethodGet:
			data, status := serv.getTask(ctx, req)
			res.WriteHeader(status)
			res.Write(data)
		case http.MethodPost:
			_, status := serv.createTask(ctx, req)
			res.WriteHeader(status)
		case http.MethodPut:
			_, status := serv.updateTask(ctx, req)
			res.WriteHeader(status)
		case http.MethodDelete:
			_, status := serv.deleteTask(ctx, req)
			res.WriteHeader(status)
		default:
			fmt.Printf("unsupported method: [%s]", req.Method)
			res.WriteHeader(http.StatusNotImplemented)
		}
	})

	return serveMux
}

func (s *server) getUser(cxt context.Context, req *http.Request) ([]byte, int) {
	id, err := getUserID(req.URL)
	if err != nil {
		fmt.Errorf("Can't parse id: %s", err)
		return nil, http.StatusBadRequest
	}
	var user *repository.User
	user, err = s.userRepo.GetById(cxt, int64(id))
	if err != nil {
		fmt.Errorf("Can't find user with id: %s", err)
		return nil, http.StatusInternalServerError
	}

	su := &serverUser{}
	su.ID = user.ID
	su.Name = user.Name
	su.Email = user.Email

	data, err := json.Marshal(su)
	if err != nil {
		fmt.Errorf("Can't marshal user with id: %d. Error: %s", id, err)
		return nil, http.StatusInternalServerError
	}

	return data, http.StatusOK
}

func (s *server) createUser(cxt context.Context, req *http.Request) (uint, int) {
	user, err := getUserData(req.Body)
	if err != nil {
		fmt.Println(err)
		return 0, http.StatusBadRequest
	}
	repUser := &repository.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	id, err := s.userRepo.Add(cxt, repUser)
	if err != nil {
		fmt.Println(err)
		return 0, http.StatusInternalServerError
	}

	return uint(id), http.StatusOK
}

func getUserData(reader io.ReadCloser) (serverUser, error) {
	body, err := io.ReadAll(reader)
	if err != nil {
		return serverUser{}, err
	}

	data := serverUser{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}

	if len(data.Name) == 0 || len(data.Email) == 0 || len(data.Password) == 0 {
		return data, errors.New("invalid data")
	}

	return data, nil
}

func getUserID(reqUrl *url.URL) (uint64, error) {
	idStr := reqUrl.Query().Get("id")
	if len(idStr) == 0 {
		return 0, errors.New("can't get id")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("can't parse id: %s", err)
	}

	return uint64(id), nil
}

func (s *server) updateUser(cxt context.Context, req *http.Request) (uint, int) {
	id, err := getUserID(req.URL)
	if err != nil {
		fmt.Errorf("Can't parse id: %s", err)
		return 0, http.StatusBadRequest
	}

	user, err := getUserData(req.Body)
	if err != nil {
		fmt.Println(err)
		return 0, http.StatusBadRequest
	}

	repUser := &repository.User{
		ID:       int64(id),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	status, err := s.userRepo.Update(cxt, repUser)
	if err != nil {
		fmt.Println(err)
		return 0, http.StatusInternalServerError
	}

	if !status {
		fmt.Println("Can't update user")
		return 0, http.StatusInternalServerError
	}
	return uint(id), http.StatusOK

}

func (s *server) deleteUser(cxt context.Context, req *http.Request) (uint, int) {
	id, err := getUserID(req.URL)
	if err != nil {
		fmt.Errorf("Can't parse id: %s", err)
		return 0, http.StatusBadRequest
	}

	status, err := s.userRepo.Delete(cxt, int64(id))
	if err != nil {
		fmt.Println(err)
		return 0, http.StatusInternalServerError
	}

	if !status {
		fmt.Println("Can't delete user")
		return 0, http.StatusInternalServerError
	}
	return uint(id), http.StatusOK
}

func (s *server) getTask(cxt context.Context, req *http.Request) ([]byte, int) {
	id, err := getTaskID(req.URL)
	if err != nil {
		fmt.Errorf("Can't parse id: %s", err)
		return nil, http.StatusBadRequest
	}
	var task *repository.Task
	task, err = s.taskRepo.GetById(cxt, int64(id))
	if err != nil {
		fmt.Errorf("Can't find task with id: %s", err)
		return nil, http.StatusInternalServerError
	}

	st := &serverTask{
		ID:          task.ID,
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
	}
	data, err := json.Marshal(st)
	if err != nil {
		fmt.Errorf("Can't marshal task with id: %d. Error: %s", id, err)
		return nil, http.StatusInternalServerError
	}

	return data, http.StatusOK
}

func getTaskID(reqUrl *url.URL) (uint64, error) {
	idStr := reqUrl.Query().Get("id")
	if len(idStr) == 0 {
		return 0, errors.New("can't get id")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("can't parse id: %s", err)
	}

	return uint64(id), nil
}

func getTaskData(reader io.ReadCloser) (serverTask, error) {
	body, err := io.ReadAll(reader)
	if err != nil {
		return serverTask{}, err
	}

	data := serverTask{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}

	if len(data.Title) == 0 || len(data.Description) == 0 {
		return data, errors.New("invalid data")
	}

	return data, nil
}

func (s *server) createTask(cxt context.Context, req *http.Request) (uint, int) {
	task, err := getTaskData(req.Body)
	if err != nil {
		fmt.Println(err)
		return 0, http.StatusBadRequest
	}
	repTask := &repository.Task{
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
	}
	id, err := s.taskRepo.Add(cxt, repTask)
	if err != nil {
		fmt.Println(err)
		return 0, http.StatusInternalServerError
	}

	return uint(id), http.StatusOK
}

func (s *server) updateTask(cxt context.Context, req *http.Request) (uint, int) {
	id, err := getTaskID(req.URL)
	if err != nil {
		fmt.Errorf("Can't parse id: %s", err)
		return 0, http.StatusBadRequest
	}

	task, err := getTaskData(req.Body)
	if err != nil {
		fmt.Println(err)
		return 0, http.StatusBadRequest
	}

	repTask := &repository.Task{
		ID:          int64(id),
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
	}
	status, err := s.taskRepo.Update(cxt, repTask)
	if err != nil {
		fmt.Println(err)
		return 0, http.StatusInternalServerError
	}

	if !status {
		fmt.Println("Can't update task")
		return 0, http.StatusInternalServerError
	}
	return uint(id), http.StatusOK
}

func (s *server) deleteTask(cxt context.Context, req *http.Request) (uint, int) {
	id, err := getTaskID(req.URL)
	if err != nil {
		fmt.Errorf("Can't parse id: %s", err)
		return 0, http.StatusBadRequest
	}

	status, err := s.taskRepo.Delete(cxt, int64(id))
	if err != nil {
		fmt.Println(err)
		return 0, http.StatusInternalServerError
	}

	if !status {
		fmt.Println("Can't delete task")
		return 0, http.StatusInternalServerError
	}
	return uint(id), http.StatusOK
}
