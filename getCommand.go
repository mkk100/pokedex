package main

func getCommand() map[string]cliCommand {
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
			name:        "map",
			description: "Explore the world of Pokemon\n",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Explore the world of Pokemon backwards\n",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Explore the pokemons",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect current pokemons",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "displays all the pokemons you have",
			callback:    commandPokedex,
		},
	}
}
