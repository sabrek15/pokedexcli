package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
	"github.com/sabrek15/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string	
	caughtPokemon	map[string]pokeapi.Pokemon
}

func startrepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("Pokedex > ")
			scanned := scanner.Scan()
			if !scanned {
				break
			}
			input := scanner.Text()

			if input == "" {
				continue
			}

			words := cleanInput(input)
			commandName := words[0]
			args := []string{}
			if len(words) > 1 {
				args = words[1:]
			}

			if cmd, exists := getCommands()[commandName]; exists {
				err := cmd.callback(cfg, args...)
				if err != nil {
					fmt.Printf("Error: %v\n", err)
				}
			} else {
				fmt.Println("unknown Command")
			}
		}

}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:		 "explore <location_name>",
			description: "Explore a location",
			callback: 	 commandExplore,
		},
		"catch": {
			name:		 "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback: 	 commandCatch,
		},
		"inspect": {
			name:		 "inspect <pokemon_name>",
			description: "view details of caught pokemon",
			callback: 	 commandInspect,
		},
		"pokedex": {
			name:		 "pokedex",
			description: "view caught pokemon",
			callback: 	 commandPokedex,
		},
	}
}
