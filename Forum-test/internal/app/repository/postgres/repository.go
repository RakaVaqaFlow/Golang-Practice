package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/storm5758/Forum-test/internal/app/repository"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) repository.User {
	return &Repository{db: db}
}
