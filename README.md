# goshell
Small cli for Go
![Shell](https://www.google.es/search?hl=es&q=conch+shell+clip+art&tbm=isch&tbs=simg:CAQSmQEJlVzya1Itp6kajQELEKjU2AQaBggUCAMICgwLELCMpwgaYgpgCAMSKIIUxgj4CZQDrQiSA-wTkwPHCO0ThTiMOIk4hSrGN8U3_1SH-IYsp3SoaMM06MRhFUJrri8UoddsySGHa6dPBQoHI0wRKIrERkfb7qI5vNozrMdhfFFfQcCaZNiAEDAsQjq7-CBoKCggIARIEZJm63Qw,isz:s&sa=X&ved=0ahUKEwiUyf_ekrnXAhVFBBoKHTWPC7kQ2A4IJigC&biw=1187&bih=703&dpr=2#imgrc=8yxSmmLHP9_9rM:)

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