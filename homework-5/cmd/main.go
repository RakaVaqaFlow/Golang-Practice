package main

import (
	"context"
	"homework/internal/handler"
	"homework/internal/pkg/db"
	dbname "homework/internal/pkg/repository/postgres"
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

	handler.CommandHandler(ctx, users, tasks)

}
