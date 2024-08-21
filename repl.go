package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printPrompt() {
	fmt.Print(cliName, " > ")
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		printPrompt()
		reader.Scan()
		text := cleanInput(reader.Text())
		if len(text) == 0 {
			continue
		}
		commands := getCommand()
		if command, ok := commands[text[0]]; ok {
			command.callback()
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
