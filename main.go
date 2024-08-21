package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
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
type Pokedex struct {
	roster map[string]Pokemon
}

var pokedex Pokedex = Pokedex{roster: make(map[string]Pokemon)}

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
func pokemonApiCall(url string, place string) PokemonEncounter {
	resp, err := http.Get(url + strings.TrimSpace(place))
	if err != nil {
		fmt.Println("yikes1")
	}
	defer resp.Body.Close()
	readResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("yikes2")
	}
	var p PokemonEncounter
	err = json.Unmarshal(readResp, &p)
	if err != nil {
		fmt.Println("yikes3")
	}
	return p
}
func commandMap() error {

	p := apiCall(url)
	if !runOrNot {
		for _, city := range p.Results {
			fmt.Println(city.Name)
		}
		runOrNot = true
	} else {
		n := apiCall(*p.Next)
		for _, city := range n.Results {
			fmt.Println(city.Name)
		}
		url = *p.Next
	}
	return nil
}

func commandExplore() error {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	p := pokemonApiCall(url, text)
	for _, pokemon := range p.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
func commandMapB() error {
	p := apiCall(url)
	if !runOrNot {
		for _, city := range p.Results {
			fmt.Println(city.Name)
		}
		runOrNot = true
	} else {
		n := apiCall(p.Previous.(string))
		for _, city := range n.Results {
			fmt.Println(city.Name)
		}
	}
	return nil
}
func pokemonCatchApiCall(pkm string) PokemonAPI {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + strings.TrimSpace(pkm))
	if err != nil {
		fmt.Println("yikes1")
	}
	defer resp.Body.Close()
	readResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("yikes2")
	}
	var p PokemonAPI
	err = json.Unmarshal(readResp, &p)
	if err != nil {
		fmt.Println("yikes3")
	}
	return p
}
func commandCatch() error {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	p := pokemonCatchApiCall(text)

	pokemonType := Pokemon{name: text, height: p.Height, weight: p.Weight, Stats: p.Stats, Types: p.Types}
	for _, stat := range pokemonType.Stats {

		if stat.BaseStat < rand.Intn(100) {
			fmt.Println(text + "caught")
			break
		} else {
			fmt.Println("Didn't catch", text)
			break
		}
	}
	pokedex.roster[pokemonType.name] = pokemonType
	return nil

}
func commandInspect() error {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if pkm, ok := pokedex.roster[text]; ok {
		fmt.Println("Stats: ")
		for _, stat := range pkm.Stats {
			fmt.Println(stat.Stat.Name + " " + string(stat.BaseStat))
		}
		fmt.Println("Types:")
		for _, pkmType := range pkm.Types {
			fmt.Println("- " + pkmType.Type.Name)
		}
	} else {
		fmt.Println("don't have")
	}
	return nil
}
func commandPokedex() error {
	for key := range pokedex.roster {
		fmt.Println(key)
	}
	return nil
}
func main() {
	startRepl()
}
