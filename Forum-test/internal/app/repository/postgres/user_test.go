package repository

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/magiconair/properties/assert"
	"github.com/storm5758/Forum-test/internal/app/models"
	"github.com/stretchr/testify/require"
)

func TestNewUsers(t *testing.T) {
	var (
		ctx = context.Background()
	)
	t.Run("success", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()
		// arrange
		queryStore := regexp.QuoteMeta(`INSERT INTO users (nickname, email, full_name, about) 
VALUES ($1,$2,$3,$4) RETURNING nickname`)
		rows := sqlmock.NewRows([]string{"nickname"}).AddRow("asd")
		f.dbMock.ExpectQuery(queryStore).
			WithArgs("asd", "asd@gmail.com", "ivan ivanov", "").
			WillReturnRows(rows)
		// act
		result, err := f.usersRepo.CreateUser(ctx, models.User{
			About:    "",
			Email:    "asd@gmail.com",
			Fullname: "ivan ivanov",
			Nickname: "asd",
		})

		// assert
		require.NoError(t, err)

		assert.Equal(t, result, models.User{
			About:    "",
			Email:    "asd@gmail.com",
			Fullname: "ivan ivanov",
			Nickname: "asd",
		})
	})
}
