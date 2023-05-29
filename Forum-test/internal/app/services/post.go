package service

import (
	"context"

	"github.com/storm5758/Forum-test/pkg/api"
	"github.com/storm5758/Forum-test/pkg/api/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type postService struct {
	api.UnimplementedPostServer
}

func NewPostService() api.PostServer {
	return &postService{}
}

// Создание новых постов
//
// Добавление новых постов в ветку обсуждения на форум.
//
// Все посты, созданные в рамках одного вызова данного метода должны иметь одинаковую дату создания (Post.Created).
func (s *postService) PostsCreate(context.Context, *api.PostsCreateRequest) (*models.Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostsCreate not implemented")
}

// Получение информации о ветке обсуждения
//
// Получение информации о ветке обсуждения по его имени.
func (s *postService) PostGetOne(context.Context, *api.PostGetOneRequest) (*models.PostFull, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostGetOne not implemented")
}

// Изменение сообщения
//
// Изменение сообщения на форуме.
//
// Если сообщение поменяло текст, то оно должно получить отметку `isEdited`.
func (s *postService) PostUpdate(context.Context, *api.PostUpdateRequest) (*models.Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostUpdate not implemented")
}
