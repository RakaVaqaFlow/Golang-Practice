package service

import (
	"context"
	"log"
	"sync"

	"github.com/storm5758/Forum-test/internal/app/models"
	"github.com/storm5758/Forum-test/internal/app/repository"
	"github.com/storm5758/Forum-test/pkg/api"
	api_models "github.com/storm5758/Forum-test/pkg/api/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	api.UnimplementedUserServer
	userRepository repository.User
}

// NewUserService return new instance of Implementation.
func NewUserService(userRepository repository.User) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

// Создание нового пользователя
//
// Создание нового пользователя в базе данных.
func (s *UserService) UserCreate(ctx context.Context, req *api.UserCreateRequest) (*api_models.User, error) {
	nikname := req.GetNickname()
	if len(nikname) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty nickname")
	}
	profile := req.GetProfile()
	if profile == nil {
		return nil, status.Error(codes.InvalidArgument, "empty profile")
	}
	email := profile.GetEmail()
	if len(email) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty email")
	}

	existedUsers, err := s.userRepository.GetUsersByNicknameOrEmail(ctx, nikname, email)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, codes.Internal.String())
	}

	if len(existedUsers) > 0 {
		return nil, status.Error(codes.AlreadyExists, codes.AlreadyExists.String())
	}

	user := models.User{
		Nickname: nikname,
		Email:    email,
		Fullname: profile.GetFullname(),
		About:    profile.GetAbout(),
	}

	createdUser, err := s.userRepository.CreateUser(ctx, user)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, codes.Internal.String())
	}

	return &api_models.User{
		About:    createdUser.About,
		Email:    createdUser.Email,
		Fullname: createdUser.Fullname,
		Nickname: createdUser.Nickname,
	}, nil
}

// Получение информации о пользователе
//
// Получение информации о пользователе форума по его имени.
func (s *UserService) UserGetOne(ctx context.Context, req *api.UserGetOneRequest) (*api_models.User, error) {
	nikname := req.GetNickname()
	if len(nikname) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty nickname")
	}
	user, err := s.userRepository.GetUsersByNicknameOrEmail(ctx, nikname, "")
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(user) == 0 {
		log.Println(err)
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}
	return &api_models.User{
		About:    user[0].About,
		Email:    user[0].Email,
		Fullname: user[0].Fullname,
		Nickname: user[0].Nickname,
	}, nil
}

// Изменение данных о пользователе
//
// Изменение информации в профиле пользователя.
func (s *UserService) UserUpdate(ctx context.Context, req *api.UserUpdateRequest) (*api_models.User, error) {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := s.userRepository.GetUsersByNicknameOrEmail(ctx, req.Nickname, req.Nickname)
			if err != nil {
				return
			}
		}()
	}
	wg.Wait()
	return nil, status.Errorf(codes.Unimplemented, "method UserUpdate not implemented")
}
