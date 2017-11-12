package goshell

import (
	"fmt"
	"errors"
	"bufio"
	"strings"
	"os"
)

// Command have a name (the thing you put in the cli)
// and a handler (the thing is executed)
type Command struct {
	Name string
	Handler func(parameters ...string) error	
}

var errCommandMissing = errors.New("Command not found")

// Shell is the cli. Have a custom prompt, an exit sequence
// to quit the cli and a list of commands.
type Shell struct {
	prompt string
	exitSequence string
	commands []Command
}

func NewShell(prompt string, maxCommands int, exitSequence string) Shell {
	return Shell{
		prompt: prompt,
		commands: make([]Command, maxCommands),
		exitSequence: exitSequence,
	}
}

func NewDefaultShell() Shell {
	return NewShell("~>", 5, "exit")
}

func(shell *Shell) RegistrerCommand(newCommand Command) {
	shell.commands = append(shell.commands, newCommand)
}

func(shell Shell) DispatchCommand(name string, params []string) error {
	for _, command := range(shell.commands) {
		if(command.Name == name) {
			return command.Handler(params...)
		}
	}
	return errCommandMissing
}

func readCommand(scanner *bufio.Scanner) (string, []string) {
	var (
		line string
		parts []string
	)
	scanner.Scan()
	line = scanner.Text()
	parts = strings.Split(line, " ")
	if len(parts) > 0 {
		return parts[0], parts[1:]
	}
	return "", []string{}
}

func(shell Shell) Run(endNotification chan<- bool) {
	var (
		commandName string
		parameters []string
		scanner *bufio.Scanner
	)
	defer close(endNotification)
	scanner = bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s ", shell.prompt)
		commandName, parameters = readCommand(scanner)
		if commandName == shell.exitSequence {
			break
		}
		if len(commandName) > 0 {
			err := shell.DispatchCommand(commandName, parameters)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	if(endNotification != nil) {
		endNotification <- true
	}
}