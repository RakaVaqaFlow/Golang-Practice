//go:build unit

package server

import (
	"bytes"
	"context"
	"encoding/json"
	"homework/internal/pkg/repository"
	mock_repository "homework/internal/pkg/repository/mocks"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

		s := server{userRepo: m}

		req, err := http.NewRequest(http.MethodGet, "user?id=1", bytes.NewReader([]byte{}))
		require.NoError(t, err)

		resultUser := &repository.User{ID: 1,
			Name:     "Test",
			Email:    "test@test.com",
			Password: "test",
		}
		m.EXPECT().GetById(gomock.Any(), int64(id)).Return(resultUser, nil)

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

func Test_getTask(t *testing.T) {
	var (
		ctx = context.Background()
		id  = 1
	)
	t.Run("success,db", func(t *testing.T) {
		t.Parallel()

		// arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := mock_repository.NewMockTasksRepo(ctrl)

		s := server{taskRepo: m}

		req, err := http.NewRequest(http.MethodGet, "task?id=1", bytes.NewReader([]byte{}))
		require.NoError(t, err)

		resultTask := &repository.Task{ID: 1,
			UserID:      1,
			Title:       "Test",
			Description: "test",
		}

		m.EXPECT().GetById(gomock.Any(), int64(id)).Return(resultTask, nil)
		// act

		_, status := s.getTask(ctx, req)
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
				&url.URL{RawQuery: "task?id"},
				false,
			},
			{
				"wrong id",
				&url.URL{RawQuery: "task?id=asdasd"},
				false,
			},
			{
				"empty",
				&url.URL{RawQuery: ""},
				false,
			},
			{
				"ok",
				&url.URL{RawQuery: "task?id=1"},
				true,
			},
		}
		for _, tc := range tt {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				id, err := getTaskID(tc.request)
				if !tc.isOk {
					assert.EqualError(t, err, "can't get id")
				} else {
					assert.Equal(t, uint64(0), id)
				}
			})
		}

	})
}

func Test_createUser(t *testing.T) {
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

		s := server{userRepo: m}

		resultUser := &repository.User{
			Name:     "Test",
			Email:    "test@test.com",
			Password: "test",
		}

		requestBody, err := json.Marshal(resultUser)
		require.NoError(t, err)
		req, err := http.NewRequest(http.MethodPost, "user", bytes.NewReader(requestBody))
		require.NoError(t, err)

		m.EXPECT().Add(gomock.Any(), resultUser).Return(int64(id), nil)

		// act
		_, status := s.createUser(ctx, req)
		// assert
		require.Equal(t, http.StatusOK, status)
	})
	t.Run("fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Parallel()
		tt := []struct {
			name        string
			requestBody io.ReadCloser
			isOk        bool
			expectedErr string
		}{
			{
				"without body",
				ioutil.NopCloser(bytes.NewReader([]byte{})),
				false,
				"unexpected end of JSON input",
			},
			{
				"wrong json format",
				ioutil.NopCloser(bytes.NewReader([]byte("asdasd"))),
				false,
				"invalid character 'a' looking for beginning of value",
			},
			{
				"ok",
				ioutil.NopCloser(bytes.NewReader([]byte(`{"name":"Test","email":"test@test.ru", "password":"test"}`))),
				true,
				"",
			},
		}
		for _, tc := range tt {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				_, err := getUserData(tc.requestBody)
				if !tc.isOk {
					assert.EqualError(t, err, tc.expectedErr)
				} else {
					assert.NoError(t, err)
				}
			})
		}

	})

}

func Test_createTask(t *testing.T) {
	var (
		ctx = context.Background()
		id  = 0
	)
	t.Run("success,db", func(t *testing.T) {
		t.Parallel()

		// arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := mock_repository.NewMockTasksRepo(ctrl)

		s := server{taskRepo: m}

		resultTask := &repository.Task{
			UserID:      0,
			Title:       "Test",
			Description: "test",
		}

		requestBody, err := json.Marshal(resultTask)
		require.NoError(t, err)
		req, err := http.NewRequest(http.MethodPost, "task", bytes.NewReader(requestBody))
		require.NoError(t, err)

		m.EXPECT().Add(gomock.Any(), resultTask).Return(int64(id), nil)

		// act
		_, status := s.createTask(ctx, req)
		// assert
		require.Equal(t, http.StatusOK, status)
	})
	t.Run("fail", func(t *testing.T) {

		t.Parallel()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tt := []struct {
			name        string
			requestBody io.ReadCloser
			isOk        bool
			expectedErr string
		}{
			{
				"without body",
				ioutil.NopCloser(bytes.NewReader([]byte{})),
				false,
				"unexpected end of JSON input",
			},
			{
				"wrong json format",
				ioutil.NopCloser(bytes.NewReader([]byte("asdasd"))),
				false,
				"invalid character 'a' looking for beginning of value",
			},
			{
				"ok",
				ioutil.NopCloser(bytes.NewReader([]byte(`{"user_id":0,"title":"Test","description":"test"}`))),
				true,
				"",
			},
		}

		for _, tc := range tt {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				_, err := getTaskData(tc.requestBody)
				if !tc.isOk {
					assert.EqualError(t, err, tc.expectedErr)
				} else {
					assert.NoError(t, err)
				}
			})
		}
	})

}

func Test_updateUser(t *testing.T) {
	var (
		ctx = context.Background()
	)
	t.Run("success,db", func(t *testing.T) {
		//arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := mock_repository.NewMockUsersRepo(ctrl)

		s := server{userRepo: m}

		resultUser := &repository.User{
			ID:       1,
			Name:     "Test",
			Email:    "test@test.com",
			Password: "test",
		}

		requestBody, err := json.Marshal(resultUser)
		require.NoError(t, err)
		req, err := http.NewRequest(http.MethodPost, "user?id=1", bytes.NewReader(requestBody))
		require.NoError(t, err)

		m.EXPECT().Update(gomock.Any(), resultUser).Return(true, nil)
		//act
		_, status := s.updateUser(ctx, req)
		//assert
		require.Equal(t, http.StatusOK, status)
	})

}

func Test_updateTask(t *testing.T) {
	var (
		ctx = context.Background()
	)
	t.Run("success,db", func(t *testing.T) {
		//arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := mock_repository.NewMockTasksRepo(ctrl)

		s := server{taskRepo: m}

		resultTask := &repository.Task{
			ID:          1,
			UserID:      0,
			Title:       "Test",
			Description: "test",
		}

		requestBody, err := json.Marshal(resultTask)
		require.NoError(t, err)
		req, err := http.NewRequest(http.MethodPost, "task?id=1", bytes.NewReader(requestBody))
		require.NoError(t, err)

		m.EXPECT().Update(gomock.Any(), resultTask).Return(true, nil)
		//act
		_, status := s.updateTask(ctx, req)
		//assert
		require.Equal(t, http.StatusOK, status)
	})

}

func Test_deleteUser(t *testing.T) {
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

		s := server{userRepo: m}

		req, err := http.NewRequest(http.MethodDelete, "user?id=1", bytes.NewReader([]byte{}))
		require.NoError(t, err)

		m.EXPECT().Delete(gomock.Any(), int64(id)).Return(true, nil)

		// act
		_, status := s.deleteUser(ctx, req)
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
				&url.URL{RawQuery: "task?id"},
				false,
			},
			{
				"wrong id",
				&url.URL{RawQuery: "task?id=asdasd"},
				false,
			},
			{
				"empty",
				&url.URL{RawQuery: ""},
				false,
			},
			{
				"ok",
				&url.URL{RawQuery: "task?id=1"},
				true,
			},
		}
		for _, tc := range tt {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				id, err := getTaskID(tc.request)
				if !tc.isOk {
					assert.EqualError(t, err, "can't get id")
				} else {
					assert.Equal(t, uint64(0), id)
				}
			})
		}

	})

}

func Test_deleteTask(t *testing.T) {
	var (
		ctx = context.Background()
		id  = 1
	)
	t.Run("success,db", func(t *testing.T) {
		t.Parallel()

		// arrange
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := mock_repository.NewMockTasksRepo(ctrl)

		s := server{taskRepo: m}

		req, err := http.NewRequest(http.MethodDelete, "task?id=1", bytes.NewReader([]byte{}))
		require.NoError(t, err)

		m.EXPECT().Delete(gomock.Any(), int64(id)).Return(true, nil)

		// act
		_, status := s.deleteTask(ctx, req)
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
				&url.URL{RawQuery: "task?id"},
				false,
			},
			{
				"wrong id",
				&url.URL{RawQuery: "task?id=asdasd"},
				false,
			},
			{
				"empty",
				&url.URL{RawQuery: ""},
				false,
			},
			{
				"ok",
				&url.URL{RawQuery: "task?id=1"},
				true,
			},
		}
		for _, tc := range tt {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				id, err := getTaskID(tc.request)
				if !tc.isOk {
					assert.EqualError(t, err, "can't get id")
				} else {
					assert.Equal(t, uint64(0), id)
				}
			})
		}

	})
}
