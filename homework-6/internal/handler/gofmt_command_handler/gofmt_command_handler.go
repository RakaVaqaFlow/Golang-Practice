package gofmt_command_handler

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"unicode"
)

type Handler interface {
	GetListOfCommands() []string
	IsMyCommand(command string) bool
	HandleCommand(ctx context.Context, command string)
	GetGoalOfHandler() string
}

type GoFmtHandler struct {
	// key - command, value - description
	commands map[string]string
}

func CreateGofmtHandler() Handler {
	gofmtHandler := GoFmtHandler{}
	gofmtHandler.commands = map[string]string{
		"gofmt <file.txt>": "to format file.txt",
	}
	return gofmtHandler
}

func (handler GoFmtHandler) GetListOfCommands() []string {
	var commands []string
	for key, value := range handler.commands {
		commands = append(commands, fmt.Sprintf("'%s' %s", key, value))
	}
	return commands
}

// check that command starts with "gofmt" and has file name with '.txt' extension
func (handler GoFmtHandler) IsMyCommand(command string) bool {
	startsWithGofmt := command[:5] == "gofmt"
	hasTxtExtension := command[len(command)-4:] == ".txt"
	return startsWithGofmt && hasTxtExtension
}

func (handler GoFmtHandler) HandleCommand(ctx context.Context, command string) {
	if len(command) < 6 {
		fmt.Println("File name is not specified.")
		return
	}
	fileName := command[6:]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}

	// Format text
	formattedLines := formatText(lines)

	// Check for errors.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}

	// Create file with the same name
	newFile, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer newFile.Close()

	// Create a new Writer for the file.
	writer := bufio.NewWriter(newFile)
	for _, line := range formattedLines {
		fmt.Fprintln(writer, line)
	}

	// Flush the writer.
	writer.Flush()

	fmt.Println("File formatted successfully.")
}

func (handler GoFmtHandler) GetGoalOfHandler() string {
	return "to format some files: add tabs before each paragraph and put dot at the end of each sentence"
}

func formatText(lines []string) []string {
	var formattedLines []string
	var prevLineEmpty bool = true
	for idx, line := range lines {
		if len(line) == 0 {
			formattedLines = append(formattedLines, "")
			prevLineEmpty = true
			continue
		}

		// Check if we need to add tab before the line
		if prevLineEmpty {
			formattedLines = append(formattedLines, "\t"+line)
		} else {
			formattedLines = append(formattedLines, line)
		}
		prevLineEmpty = false

		// Check if we need to add dot at the end of the line
		if len(line) > 0 && !unicode.IsPunct(rune(line[len(line)-1])) {
			// if it the end of the file
			if idx == len(lines)-1 {
				formattedLines[len(formattedLines)-1] += "."
			}
			// if next line is empty
			if idx+1 < len(lines) && len(lines[idx+1]) == 0 {
				formattedLines[len(formattedLines)-1] += "."
			}
			// if next line is not empty and does not start with a capital letter
			if idx+1 < len(lines) && len(lines[idx+1]) > 0 && !unicode.IsUpper(rune(lines[idx+1][0])) {
				formattedLines[len(formattedLines)-1] += "."
			}
		}
	}
	if !prevLineEmpty && len(formattedLines) > 0 {
		// Add dot at the end of the file if it is not there
		lastLine := formattedLines[len(formattedLines)-1]
		if len(lastLine) > 0 && !unicode.IsPunct(rune(lastLine[len(lastLine)-1])) {
			formattedLines[len(formattedLines)-1] += "."
		}
	}
	return formattedLines
}
