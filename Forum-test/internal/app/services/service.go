package service

import (
	"context"

	"github.com/storm5758/Forum-test/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Implementation struct {
	api.UnimplementedAdminServer
}

func NewAdminService() *Implementation {
	return &Implementation{}
}

// Очистка всех данных в базе
//
// Безвозвратное удаление всей пользовательской информации из базы данных.
func (s *Implementation) Clear(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clear not implemented")
}

// Получение инфомарции о базе данных
//
// Получение инфомарции о базе данных.
func (s *Implementation) Status(context.Context, *emptypb.Empty) (*api.StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}

func SomeFunc(a chan int) {

}
