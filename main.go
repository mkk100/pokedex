package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var cliName string = "pokedex"
type cliCommand struct {
	name        string
	description string
	callback    func() error
}
type PokemonCities struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
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
		"map": {
			name: "map",
			description: "Explore the world of Pokemon\n",
			callback: commandMap,
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
type City struct {
    Name string
    URL  string
}
func commandMap() error{
	resp, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		fmt.Println("yikes")
	}
	defer resp.Body.Close()
	readResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("yikes")
	}

	var p PokemonCities
	err = json.Unmarshal(readResp, &p)
	if err != nil {
		fmt.Println("yikes")
	}
	for _, city := range p.Results{
		fmt.Println(city.Name)
	}
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