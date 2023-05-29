package server

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/workshop/internal/pkg/repository"
	in_memory "gitlab.ozon.dev/workshop/internal/pkg/repository/in-memory"
	mock_repository "gitlab.ozon.dev/workshop/internal/pkg/repository/mocks"
)

func Test_getUser(t *testing.T) {
	var (
		ctx = context.Background()
		id  = 1
	)
	t.Run("success,db", func(t *testing.T) {
		t.Parallel()

		// arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := mock_repository.NewMockUsersRepo(ctrl)

		mc := in_memory.NewUsers()
		s := server{userRepo: m, userCachedRepo: mc}

		req, err := http.NewRequest(http.MethodGet, "user?id=1", bytes.NewReader([]byte{}))
		require.NoError(t, err)

		m.EXPECT().GetById(gomock.Any(), int64(id)).Return(&repository.User{ID: 1, Name: "asd"}, nil)
		// act

		_, status := s.getUser(ctx, req)
		// assert
		require.Equal(t, http.StatusOK, status)
	})

	t.Run("fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Parallel()
		tt := []struct {
			name    string
			request *url.URL
			isOk    bool
		}{
			{
				"without id",
				&url.URL{RawQuery: "user?id"},
				false,
			},
			{
				"wrong id",
				&url.URL{RawQuery: "user?id=asdasd"},
				false,
			},
			{
				"empty",
				&url.URL{RawQuery: ""},
				false,
			},
			{
				"ok",
				&url.URL{RawQuery: "user?id=1"},
				true,
			},
		}
		for _, tc := range tt {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				id, err := getUserID(tc.request)
				if !tc.isOk {
					assert.EqualError(t, err, "can't get id")
				} else {
					assert.Equal(t, uint64(0), id)
				}
			})
		}

	})
}
