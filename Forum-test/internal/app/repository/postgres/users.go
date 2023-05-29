package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/storm5758/Forum-test/internal/app/models"
)

func (r *Repository) GetUsersByNicknameOrEmail(ctx context.Context, nickname, email string) ([]models.User, error) {
	query, args, err := squirrel.Select("nickname, email, full_name, about").
		From("users").
		Where(squirrel.Or{
			squirrel.Eq{
				"nickname": strings.ToLower(nickname),
			},
			squirrel.Eq{
				"email": email,
			},
		}).PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("Repository.GetUsersByNiknameOrEmail: to sql: %w", err)
	}

	var users []models.User

	err = r.db.SelectContext(ctx, &users, query, args...)

	if err != nil {
		return nil, errors.Wrap(err, "db.SelectContext()")
	}

	return users, nil
}

func (r *Repository) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	query, args, err := squirrel.Insert("users").
		Columns("nickname, email, full_name, about").
		Values(strings.ToLower(user.Nickname), user.Email, user.Fullname, user.About).
		Suffix("RETURNING nickname").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return models.User{}, fmt.Errorf("Repository.UserCreate: to sql: %w", err)
	}

	queryer := r.getQueryer(nil)
	row := queryer.QueryRowxContext(ctx, query, args...)

	var nickname string
	err = row.Scan(&nickname)

	if err != nil {
		return models.User{}, errors.Wrap(err, "CreateOffice:Scan()")
	}
	return user, nil
}

func (r *Repository) getQueryer(tx *sqlx.Tx) sqlx.QueryerContext {
	if tx == nil {
		return r.db
	}
	return tx
}

func (r *Repository) getExecer(tx *sqlx.Tx) sqlx.ExecerContext {
	if tx == nil {
		return r.db
	}
	return tx
}
