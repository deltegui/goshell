package goshell

import (
	"fmt"
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func command(...string) error {
	return nil
}

func commandTwo(...string) error {
	return nil
}

func spyCommand(t *testing.T) CommandHandler {
	return func(c ...string) error {
		if !assert.Equal(t, c, []string{"test"}) {
			t.FailNow()
		}
		return nil
	}
}

func TestShellConstructor(t *testing.T) {
	s := NewShell("-", "pipo", os.Stdout, os.Stdin)
	expected := Shell{
		prompt: s.prompt,
		exitSequence: s.exitSequence,
		reader: s.reader,
		writer: s.writer,
		commands: s.commands,
	}
	if !assert.ObjectsAreEqualValues(expected, s) {
		t.FailNow()
	}
}

func TestDefaultShellConstructor(t *testing.T) {
	s := NewDefaultShell()
	expected := Shell{
		prompt: s.prompt,
		exitSequence: s.exitSequence,
		reader: s.reader,
		writer: s.writer,
		commands: s.commands,
	}
	if !assert.ObjectsAreEqualValues(expected, s) {
		t.FailNow()
	}
}

func TestDefaultShellConstructorCreatesOneCommand(t *testing.T) {
	s := NewDefaultShell()
	if len(s.commands) != 1 {
		t.FailNow()
	}
}

func TestRegistrerCommandShouldInsertCommand(t *testing.T) {
	s := NewShell("-", "e", os.Stdout, os.Stdin)
	s.RegistrerCommand("test", command)
	if len(s.commands) != 1 {
		t.FailNow()
	}
}

func TestRegistrerAllCommandsShouldInsertCommands(t *testing.T) {
	s := NewShell("-", "e", os.Stdout, os.Stdin)
	commandMap := make(map[string]CommandHandler)
	commandMap["c"] = command
	commandMap["c1"] = commandTwo
	s.RegistrerAllCommands(commandMap)
	if len(s.commands) != 2 {
		fmt.Fprint(os.Stdout, "Command map have not two elements: ", s.commands, " len: ", len(s.commands))
		t.FailNow()
	}
}

func TestRegistrerAllCommandsShouldKeepEmptyIfEmptyMap(t *testing.T) {
	s := NewShell("-", "e", os.Stdout, os.Stdin)
	commandMap := make(map[string]CommandHandler)
	s.RegistrerAllCommands(commandMap)
	if len(s.commands) != 0 {
		fmt.Fprint(os.Stdout, "Command map have not two elements: ", s.commands, " len: ", len(s.commands))
		t.FailNow()
	}
}

func TestShellShouldNotStoreSameCommands(t *testing.T) {
	s := NewShell("-", "e", os.Stdout, os.Stdin)
	commandMap := make(map[string]CommandHandler)
	commandMap["c"] = command
	commandMap["c"] = command
	s.RegistrerAllCommands(commandMap)
	if len(s.commands) != 1 {
		fmt.Fprint(os.Stdout, "Command map have not two elements: ", s.commands, " len: ", len(s.commands))
		t.FailNow()
	}
}

func TestDispatchCommandShouldCallACommand(t *testing.T) {
	s := NewDefaultShell()
	c := spyCommand(t)
	s.RegistrerCommand("c", c)
	s.dispatchCommand("c", []string{"test"})
}