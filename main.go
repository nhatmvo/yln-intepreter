package main

import (
	"fmt"
	"os"
	"os/user"
	"yln/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome %s!\n", user.Username)
	fmt.Printf("Feel free to enter any command!\n")
	repl.Start(os.Stdin, os.Stdout)
}
