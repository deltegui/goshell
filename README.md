# goshell
![Shell](https://etc.usf.edu/clipart/87700/87746/87746_conch-shell_md.gif)

[![Build Status](https://travis-ci.org/frikiman34/goshell.svg?branch=master)](https://travis-ci.org/frikiman34/goshell)
[![BCH compliance](https://bettercodehub.com/edge/badge/frikiman34/FuenteovejunaFinderApi?branch=master)](https://bettercodehub.com/)

Small cli for Go

##Example of use
```go
package main

import "goshell"
import "fmt"

func echo(params ...string) error {
	fmt.Println(params)
	return nil
}

func hello(params ...string) error {
	fmt.Print("Hi")
	if len(params) >= 1 {
		fmt.Printf(" %s", params[0])
	}
	fmt.Printf("!\n")
	return nil
}

func main() {
	var (
		shell goshell.Shell
		notifyEnd chan bool
	)

	notifyEnd = make(chan bool)
	shell = goshell.NewDefaultShell()
	shell.RegistrerCommand("echo", echo)
	shell.RegistrerCommand("hello", hello)
	go goshell.Run(shell, notifyEnd)
	<-notifyEnd
}
```
