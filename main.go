package main

import (
	"fmt"
	"os"
	"os/user"

	"vichr.me/go/monkey-lang-interpreter/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Write code %s\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
