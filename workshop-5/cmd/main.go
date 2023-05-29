package main

import (
	"context"
	"fmt"

	"gitlab.ozon.dev/workshop/internal/pkg/db"
	"gitlab.ozon.dev/workshop/internal/pkg/repository"
	"gitlab.ozon.dev/workshop/internal/pkg/transaction/postgresql"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	database, err := db.NewDB(ctx)
	if err != nil {
		return
	}
	defer database.GetPool(ctx).Close()
	serviceTxBuilder := postgresql.NewServiceTxBuilder(database)
	//
	tx, err := serviceTxBuilder.ServiceTx(ctx)
	if err != nil {
		return
	}
	user := &repository.User{Name: "name", ID: 5}

	isOk, err := tx.Users.Update(ctx, user)
	if err != nil {
		tx.Rollback(ctx)
		return
	}
	fmt.Println(isOk)
	tx.Commit(ctx)
}

//	type someHandler struct {
//		userR   repository.UsersRepo
//		txModel postgresql.ServiceTxBuidler
//	}
//
// func (s someHandler)  Handle(){
//
//		go func() {
//			s.txModel.ServiceTx()
//		}()
//
//		go func() {
//			s.userR.Add()
//		}()
//	}
func SetUser(ctx context.Context, cached repository.UsersRepoCached, usersRepo repository.UsersRepo) {
	user := &repository.User{Name: "asd"}
	id, err := usersRepo.Add(ctx, user)
	if err != nil {
		fmt.Print(err)
	}
	user.ID = id

	cached.Add(ctx, user)
	if err != nil {
		fmt.Print(err)
	}

}

func GetUser(ctx context.Context, id int64, cached repository.UsersRepoCached, usersRepo repository.UsersRepo) (*repository.User, error) {
	var user *repository.User
	user, err := cached.Get(ctx, id)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	if user.Name == "bar" {
		return user, nil
	}
	user, err = usersRepo.GetById(ctx, id)
	if err != nil {
		fmt.Print(err)
	}
	return user, nil
}

//func waitForSignal() {
//	ch := make(chan os.Signal, 1)
//	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
//	for {
//		s := <-ch
//
//		switch s {
//		case syscall.SIGINT:
//			fallthrough
//		case syscall.SIGTERM:
//			logger.Get(logger.AppLoggerKey).Println(fmt.Sprintf("Got signal: %v, exiting.", s))
//			return
//		case syscall.SIGHUP:
//			logger.Get(logger.AppLoggerKey).Println("Reload config")
//			config.LoadConfig()
//			logger.ReInitLoggers()
//			logger.Get(logger.AppLoggerKey).Println("Reload config complete")
//		}
//	}
//}
