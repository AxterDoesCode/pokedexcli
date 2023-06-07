package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AxterDoesCode/pokedexcli/internal/pokeapi"
	"github.com/AxterDoesCode/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	caughtPokemon   map[string]pokeapi.Pokemon
	cache           pokecache.Cache
	pokeapiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},

		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
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

		"explore": {
			name:        "explore",
			description: "Explore a location in the Pokemon world",
			callback:    commandExplore,
		},

		"catch": {
			name:        "catch",
			description: "Catch a Pokemon",
			callback:    commandCatch,
		},
	}
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		reader.Scan()

		words := strings.Fields(
			reader.Text(),
		) //.Text() returns a string of the entire input into the reader, strings.Fields splits the string

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]

		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
