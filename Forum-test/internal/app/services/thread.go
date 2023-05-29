package service

import (
	"context"

	"github.com/storm5758/Forum-test/pkg/api"
	"github.com/storm5758/Forum-test/pkg/api/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type threadService struct {
	api.UnimplementedThreadServer
}

func NewThreadService() api.ThreadServer {
	return &threadService{}
}

// Создание ветки
//
// Добавление новой ветки обсуждения на форум.
func (s *threadService) ThreadCreate(context.Context, *api.ThreadCreateRequest) (*models.Thread, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ThreadCreate not implemented")
}

// Получение информации о ветке обсуждения
//
// Получение информации о ветке обсуждения по его имени.
func (s *threadService) ThreadGetOne(context.Context, *api.ThreadGetOneRequest) (*models.Thread, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ThreadGetOne not implemented")
}

// Сообщения данной ветви обсуждения
//
// Получение списка сообщений в данной ветке форуме.
//
// Сообщения выводятся отсортированные по дате создания.
func (s *threadService) ThreadGetPosts(context.Context, *api.ThreadGetPostsRequest) (*models.Thread, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ThreadGetPosts not implemented")
}

// Обновление ветки
//
// Обновление ветки обсуждения на форуме.
func (s *threadService) ThreadUpdate(context.Context, *api.ThreadUpdateRequest) (*models.Thread, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ThreadUpdate not implemented")
}

// Проголосовать за ветвь обсуждения
//
// Изменение голоса за ветвь обсуждения.
//
// Один пользователь учитывается только один раз и может изменить своё
// мнение.
func (s *threadService) ThreadVote(context.Context, *api.ThreadVoteRequest) (*models.Thread, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ThreadVote not implemented")
}
