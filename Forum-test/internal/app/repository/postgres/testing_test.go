package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/storm5758/Forum-test/internal/app/repository"
)

type usersRepoFixtures struct {
	usersRepo repository.User
	db        *sqlx.DB
	dbMock    sqlmock.Sqlmock
}

func setUp(t *testing.T) usersRepoFixtures {
	var fixture usersRepoFixtures

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("somerror")
	}
	fixture.db = sqlx.NewDb(db, "sqlmock")
	fixture.dbMock = mock
	fixture.usersRepo = NewRepository(fixture.db)
	return fixture
}

func (f *usersRepoFixtures) tearDown() {
	f.db.Close()
}
