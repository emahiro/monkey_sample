package main

import (
	"fmt"
	"os"
	"os/user"

	"monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("hello %s!. this is the Monkey programming language. \n", user.Username)
	fmt.Printf("feel free ato type in command following. \n")
	repl.Start(os.Stdin, os.Stdout)
}
