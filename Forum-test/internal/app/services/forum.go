package service

import (
	"context"

	"github.com/storm5758/Forum-test/pkg/api"
	"github.com/storm5758/Forum-test/pkg/api/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type forumService struct {
	api.UnimplementedForumServer
}

func NewForumService() api.ForumServer {
	return &forumService{}
}

// Создание форума
//
// Создание нового форума.
func (s *forumService) ForumCreate(context.Context, *api.ForumCreateRequest) (*models.Forum, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForumCreate not implemented")
}

// Получение информации о форуме
//
// Получение информации о форуме по его идентификаторе.
func (s *forumService) ForumGetOne(context.Context, *api.ForumGetOneRequest) (*models.Forum, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForumGetOne not implemented")
}

// Список ветвей обсужления форума
//
// Получение списка ветвей обсужления данного форума.
//
// Ветви обсуждения выводятся отсортированные по дате создания.
func (s *forumService) ForumGetThreads(context.Context, *api.ForumGetThreadsRequest) (*models.Thread, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForumGetThreads not implemented")
}

// Пользователи данного форума
//
// Получение списка пользователей, у которых есть пост или ветка обсуждения в данном форуме.
//
// Пользователи выводятся отсортированные по nickname в порядке возрастания.
// Порядок сотрировки должен соответсвовать побайтовому сравнение в нижнем регистре.
func (s *forumService) ForumGetUsers(context.Context, *api.ForumGetUsersRequest) (*api.ForumGetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForumGetUsers not implemented")
}
