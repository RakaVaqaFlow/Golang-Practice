package command_manager

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

type Handler interface {
	GetListOfCommands() []string
	IsMyCommand(command string) bool
	HandleCommand(ctx context.Context, command string)
	GetGoalOfHandler() string
}

type HandlersManager struct {
	handlers []Handler
}

func CreateManager(handlers ...Handler) HandlersManager {
	handlersManager := HandlersManager{}
	handlersManager.handlers = handlers
	return handlersManager
}

func (handlersManager *HandlersManager) AddHandler(serviceHandler Handler) {
	handlersManager.handlers = append(handlersManager.handlers, serviceHandler)
}

func (handlersManager *HandlersManager) Start(ctx context.Context) {
	handlersManager.printListOfCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter command: ")
		var command string
		if scanner.Scan() {
			command = scanner.Text()
		}
		// to exit from program
		if command == "exit" {
			fmt.Println("Bye!")
			break
		}
		// to print list of commands from handlers
		if command == "help" {
			handlersManager.printListOfCommands()
			continue
		}
		// to handle command from handlers
		isHandled := false
		for _, handler := range handlersManager.handlers {
			if handler.IsMyCommand(command) {
				handler.HandleCommand(ctx, command)
				isHandled = true
			}
		}
		// if command is not found
		if !isHandled {
			fmt.Println("Unknown command, type 'help' to print list of commands")
		}
	}
}

func (handlersManager *HandlersManager) printListOfCommands() {
	fmt.Println("Available commands:")
	fmt.Println(`	'help' to print list of commands`)
	fmt.Println(`	'exit'`)
	for _, handler := range handlersManager.handlers {
		fmt.Println("  " + handler.GetGoalOfHandler() + `:`)
		commands := handler.GetListOfCommands()
		for _, command := range commands {
			fmt.Println("\t" + command)
		}
	}
}
