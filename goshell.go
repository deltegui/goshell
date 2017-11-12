// Package goshell is a small library to create easilly CLIs for Golang
package goshell

import (
	"fmt"
	"errors"
	"bufio"
	"strings"
	"os"
)

// CommandHandler is a function that is called
// when a command is entered. The library call it
// with the parameters the user introduces in the cli.
// You can return an error that will be printed in the 
// screen.
type CommandHandler func(...string) error

var errCommandMissing = errors.New("Command not found")

// Shell is the cli. Have a custom prompt, an exit sequence
// to quit the cli and a list of commands.
type Shell struct {
	prompt string
	exitSequence string
	commands map[string]CommandHandler
}

// NewShell creates a custom shell with a prompt and a exitSequence.
// The exitSequence is a string that the user should write in the cli
// to exit.
func NewShell(prompt string, exitSequence string) Shell {
	s := Shell{
		prompt: prompt,
		commands: make(map[string]CommandHandler),
		exitSequence: exitSequence,
	}
	s.RegistrerCommand("command-list", s.listCommands)
	return s
}

// NewDefaultShell creates a new shell with default options.
func NewDefaultShell() Shell {
	return NewShell("~>", "exit")
}

// RegistrerCommand registrers a new command. You should pass a command name
// (thats is the user should type in the cli to call yor command), and the
// handler.
func(shell *Shell) RegistrerCommand(name string, newCommand CommandHandler) {
	shell.commands[name] = newCommand
}

// RegistrerAllCommands registrers a map of commands.
func(shell *Shell) RegistrerAllCommands(newCommands map[string]CommandHandler) {
	for name, command := range(newCommands) {
		shell.RegistrerCommand(name, command)
	}
}

func(shell Shell) dispatchCommand(target string, params []string) error {
	for name, handler := range(shell.commands) {
		if(name == target) {
			return handler(params...)
		}
	}
	return errCommandMissing
}

// Run starts a your cli. You pass your shell strcture and a channel to
// notify when the cli exited. If you are not using goroutines, you can
// pass nil.
func Run(shell Shell, endNotification chan<- bool) {
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
			err := shell.dispatchCommand(commandName, parameters)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	if(endNotification != nil) {
		endNotification <- true
	}
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
	return "", nil
}