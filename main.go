package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var cliName string = "pokedex"
var runOrNot bool = false
var url string = "https://pokeapi.co/api/v2/location-area/"
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
	
type Pokemon struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
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
		"mapb":{
			name: "mapb",
			description: "Explore the world of Pokemon backwards\n",
			callback: commandMapB,
		},
		"explore":{
			name: "explore",
			description: "Explore the pokemons",
			callback: commandExplore,
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
func apiCall(url string) PokemonCities {
	resp, err := http.Get(url)
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
	return p
}
func pokemonApiCall(url string, place string) Pokemon {
	resp, err := http.Get(url + strings.TrimSpace(place))
	if err != nil {
		fmt.Println("yikes1")
	}
	defer resp.Body.Close()
	readResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("yikes2")
	}
	var p Pokemon
	err = json.Unmarshal(readResp, &p)
	if err != nil {
		fmt.Println("yikes3")
	}
	return p
}
func commandMap() error{
	
	p := apiCall(url)
	if !runOrNot {
	for _, city := range p.Results{
		fmt.Println(city.Name)
	}
	runOrNot = true
	} else {
		n := apiCall(p.Next)
		for _, city := range n.Results{
			fmt.Println(city.Name)
		}
		url = p.Next
	}
	return nil
}
func commandExplore() error {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	p := pokemonApiCall(url, text)
	for _, pokemon := range p.PokemonEncounters{
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
	
}
func commandMapB() error{
	p := apiCall(url)
	if !runOrNot {
	for _, city := range p.Results{
		fmt.Println(city.Name)
	}
	runOrNot = true
	} else {
		n := apiCall(p.Previous.(string))
		for _, city := range n.Results{
			fmt.Println(city.Name)
		}
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