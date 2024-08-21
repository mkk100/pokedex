package main

import (
	"fmt"
	"os"
)

func commandHelp() error {
	commands := getCommand()
	for _, cmd := range commands {
		fmt.Println(cmd.name, " : ", cmd.description)
	}
	fmt.Println()
	return nil
}
func commandExit() error {
	os.Exit(0)
	return nil
}
