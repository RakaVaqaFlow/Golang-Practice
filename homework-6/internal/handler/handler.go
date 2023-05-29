package handler

import (
	"context"
)

type Handler interface {
	GetListOfCommands() []string
	IsMyCommand(command string) bool
	HandleCommand(ctx context.Context, command string)
	GetGoalOfHandler() string
}
