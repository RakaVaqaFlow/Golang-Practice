package main

import (
	"context"
	"homework/internal/handler"
	"homework/internal/pkg/db"
	dbname "homework/internal/pkg/repository/postgres"
	"homework/internal/pkg/server"
	"log"
	"net/http"
)

const (
	serverPort = ":9000"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// database connection
	database, err := db.NewDB(ctx)
	if err != nil {
		return
	}
	defer database.GetPool(ctx).Close()
	tasks := dbname.NewTasks(database)
	users := dbname.NewUsers(database)

	// start server
	log.Printf("Server started on localhost%s\n", serverPort)
	serverMux := server.CreateServer(ctx, users, tasks)
	go func() {
		err := http.ListenAndServe(serverPort, serverMux)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// start command handler
	log.Printf("Command handler started\n")
	handler.CommandHandler(ctx, users, tasks)
}
