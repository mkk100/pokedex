package main

import (
	"bufio"
	"fmt"
	"os"
)

var cliName string = "pokedex"
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommand() map[string]cliCommand{
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message\n",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex\n",
			callback:    commandExit,
		},
	}
}


func printPrompt() {
	fmt.Print(cliName," > ")
}
func commandHelp() error{
	commands := getCommand()
	for _, cmd := range commands {
		fmt.Println(cmd.name," : ",cmd.description)
	}
	fmt.Println()
	return nil
}
func commandExit() error{
	os.Exit(0)
	return nil
}

func main(){
	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for reader.Scan(){
		commands := getCommand()
		text := reader.Text()
		if command,ok := commands[text]; ok {
			fmt.Println(command.callback())
		}
		printPrompt()
	}
}