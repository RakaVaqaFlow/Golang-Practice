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

	"gitlab.ozon.dev/workshop/internal/pkg/repository"
)

type serverUser struct {
	ID    uint64
	Name  *string
	Email *string
}

type server struct {
	userRepo       repository.UsersRepo
	userCachedRepo repository.UsersRepo
}

func (s *server) getUser(cxt context.Context, req *http.Request) ([]byte, int) {
	id, err := getUserID(req.URL)
	if err != nil {
		fmt.Errorf("can't parse id: %s", err)
		return nil, http.StatusBadRequest
	}
	var user *repository.User
	user, _ = s.userCachedRepo.GetById(cxt, int64(id))
	if user == nil {
		user, err = s.userRepo.GetById(cxt, int64(id))
		if err != nil {
			fmt.Errorf("can't parse id: %s", err)
			return nil, http.StatusInternalServerError
		}
	}

	su := &serverUser{}
	su.ID = uint64(user.ID)
	su.Name = &user.Name

	data, err := json.Marshal(su)
	if err != nil {
		fmt.Errorf("can't marshal user with id: %d. Error: %s", id, err)
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
	id, err := s.userRepo.Add(cxt, &repository.User{Name: *user.Name})
	if err != nil {
		fmt.Println(err)
		return 0, http.StatusInternalServerError
	}

	return uint(id), http.StatusOK
}

func CreateServer(ctx context.Context, ur repository.UsersRepo) *http.ServeMux {
	serv := server{
		userRepo: ur,
	}
	serveMux := http.NewServeMux()

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

		default:
			fmt.Printf("unsupported method: [%s]", req.Method)
			res.WriteHeader(http.StatusNotImplemented)
		}
	})

	return serveMux
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
