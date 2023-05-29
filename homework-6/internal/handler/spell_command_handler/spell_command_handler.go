package handler

import (
	"context"
	"fmt"
)

type Handler interface {
	GetListOfCommands() []string
	IsMyCommand(command string) bool
	HandleCommand(ctx context.Context, command string)
	GetGoalOfHandler() string
}

type SpellCommandHandler struct {
	// key - command, value - description
	commands map[string]string
}

func CreateSpellHandler() Handler {
	spellCommandHandler := SpellCommandHandler{}
	spellCommandHandler.commands = map[string]string{
		"spell <word>": "to print word with spaces between letters",
	}
	return spellCommandHandler
}

func (handler SpellCommandHandler) GetListOfCommands() []string {
	var commands []string
	for key, value := range handler.commands {
		commands = append(commands, fmt.Sprintf("'%s' %s", key, value))
	}
	return commands
}

func (handler SpellCommandHandler) IsMyCommand(command string) bool {
	startsWithSpell := command[:5] == "spell"
	return startsWithSpell
}

func (handler SpellCommandHandler) HandleCommand(ctx context.Context, command string) {
	if len(command) < 7 {
		fmt.Println("Please, enter word after 'spell' command")
		return
	}
	word := command[6:]
	for _, letter := range word {
		fmt.Printf("%c ", letter)
	}
	fmt.Println()
}

func (handler SpellCommandHandler) GetGoalOfHandler() string {
	return "to print word with spaces between letters"
}
