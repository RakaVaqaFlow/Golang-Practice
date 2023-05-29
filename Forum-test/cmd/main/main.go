package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	postgres "github.com/storm5758/Forum-test/internal/app/repository/postgres"
	"github.com/storm5758/Forum-test/internal/app/server"
	services "github.com/storm5758/Forum-test/internal/app/services"
	"github.com/storm5758/Forum-test/internal/pkg/database"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// connection string
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DBname)

	db, err := database.NewPostgres(ctx, psqlConn, "pgx")
	if err != nil {
		log.Fatal("ping database error", err)
	}
	defer db.Close()

	// ceate repository
	repo := postgres.NewRepository(db)

	// create server
	srv, err := server.New(server.Services{
		Admin:  services.NewAdminService(),
		User:   services.NewUserService(repo),
		Forum:  services.NewForumService(),
		Post:   services.NewPostService(),
		Thread: services.NewThreadService(),
	})
	if err != nil {
		log.Fatalf("can't create server: %s", err.Error())
	}
	go func() {
		http.ListenAndServe(":8024", nil)
	}()
	// run server
	if err := srv.Run(ctx); err != nil {
		log.Println(err)
	}
}
