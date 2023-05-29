package memcache

import (
	"context"
	"fmt"
	"strconv"

	"github.com/bradfitz/gomemcache/memcache"
	"gitlab.ozon.dev/workshop/internal/pkg/repository"
)

type UsersRepo struct {
	cli *memcache.Client
}

func NewCachedRepo(cli *memcache.Client) *UsersRepo {
	return &UsersRepo{cli: cli}
}

func (u *UsersRepo) Add(_ context.Context, user *repository.User) error {
	err := u.cli.Set(&memcache.Item{
		Key:        strconv.FormatInt(user.ID, 10),
		Value:      []byte("bar"),
		Expiration: 0,
	})
	return err
}

func (u *UsersRepo) Get(ctx context.Context, id int64) (*repository.User, error) {
	it, err := u.cli.Get(strconv.FormatInt(id, 10))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &repository.User{Name: string(it.Value)}, nil

}

func (u *UsersRepo) GetMulti(ctx context.Context, ids []*int64) ([]*repository.User, error) {
	it, err := u.cli.GetMulti([]string{"1", "2"})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(it)
	return nil, nil
}
