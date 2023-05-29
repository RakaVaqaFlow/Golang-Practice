package main

import (
	"context"
	commandManager "homework/internal/command_manager"
	databaseHandler "homework/internal/handler/database_command_handler"
	gofmtHandler "homework/internal/handler/gofmt_command_handler"
	spellHandler "homework/internal/handler/spell_command_handler"
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
	databaseCommandHandler := databaseHandler.CreateDataBaseCommandHandler(users, tasks)

	// handlers for commands
	spellCommandHandler := spellHandler.CreateSpellHandler()
	gofmtCommandHandler := gofmtHandler.CreateGofmtHandler()

	// command manager
	Manager := commandManager.CreateManager(databaseCommandHandler, spellCommandHandler, gofmtCommandHandler)
	Manager.Start(ctx)
}
