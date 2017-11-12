package goshell

import "fmt"

func(shell Shell) listCommands(params ...string) error {
	for name := range(shell.commands) {
		fmt.Println(name)
	}
	fmt.Printf("Type '%s' to exit.\n", shell.exitSequence)
	return nil
}