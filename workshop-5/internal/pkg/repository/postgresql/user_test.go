package postgresql

import (
	"context"
	"database/sql"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_database "gitlab.ozon.dev/workshop/internal/pkg/db/mocks"
)

func TestUsersRepo_GetById(t *testing.T) {
	var (
		ctx = context.Background()
		id  = 1
	)
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		// arrange
		s := setUp(t)
		defer s.tearDown()

		s.mockDb.EXPECT().Get(gomock.Any(), gomock.Any(), "SELECT id,name,created_at,updated_at FROM users WHERE id=$1", gomock.Any()).Return(nil)

		// act
		user, err := s.repo.GetById(ctx, int64(id))

		// assert
		assert.NoError(t, err)
		assert.Equal(t, int64(0), user.ID)
	})
	t.Run("fail", func(t *testing.T) {
		t.Parallel()

		t.Run("internal error", func(t *testing.T) {
			t.Parallel()

			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().Get(gomock.Any(), gomock.Any(), "SELECT id,name,created_at,updated_at FROM users WHERE id=$1", gomock.Any()).Return(assert.AnError)
			// act
			user, err := s.repo.GetById(ctx, int64(id))
			// assert

			assert.EqualError(t, err, "assert.AnError general error for testing")
			assert.NotNil(t, user)
		})

		t.Run("not found", func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockDb := mock_database.NewMockPGX(ctrl)
			repo := NewUsers(mockDb)

			mockDb.EXPECT().Get(gomock.Any(), gomock.Any(), "SELECT id,name,created_at,updated_at FROM users WHERE id=$1", gomock.Any()).Return(sql.ErrNoRows)
			// act
			user, err := repo.GetById(ctx, int64(id))
			// assert

			assert.EqualError(t, err, "object not found")
			assert.Nil(t, user)
		})

	})

}
