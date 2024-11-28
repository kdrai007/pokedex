package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kdrai007/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationUrl *string
	prevLocationUrl *string
	catchedPokemon  map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()

		words := clearInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		command := words[0]
		commandType, exists := fieldsValue()[command]

		args := []string{}

		if len(words) > 1 {
			args = words[1:]
		}
		if exists {
			err := commandType.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func clearInput(input string) []string {
	output := strings.ToLower(input)
	words := strings.Fields(output)
	return words
}

type cliValue struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func fieldsValue() map[string]cliValue {
	return map[string]cliValue{
		"pokedex": {
			name:        "pokedex",
			description: "See all your Pokemons",
			callback:    commandPokedex,
		},
		"catch": {
			name:        "catch",
			description: "Catch pokemon using this command",
			callback:    commandCatch,
		},
		"info": {
			name:        "info",
			description: "Look pokemon stats",
			callback:    commandInfo,
		},
		"explore": {
			name:        "explore",
			description: "find pokemon in given location",
			callback:    commandExplore,
		},
		"map": {
			name:        "map",
			description: "Display area names",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous area names",
			callback:    commandMapB,
		},
		"help": {
			name:        "help",
			description: "Displayes a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
