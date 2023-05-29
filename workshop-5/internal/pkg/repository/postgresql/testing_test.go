package postgresql

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_database "gitlab.ozon.dev/workshop/internal/pkg/db/mocks"
	"gitlab.ozon.dev/workshop/internal/pkg/repository"
)

type usersRepoFixture struct {
	ctrl   *gomock.Controller
	repo   repository.UsersRepo
	mockDb *mock_database.MockPGX
}

func setUp(t *testing.T) usersRepoFixture {
	ctrl := gomock.NewController(t)

	mockDb := mock_database.NewMockPGX(ctrl)
	repo := NewUsers(mockDb)
	return usersRepoFixture{repo: repo, ctrl: ctrl, mockDb: mockDb}

}

func (u *usersRepoFixture) tearDown() {
	u.ctrl.Finish()
}
