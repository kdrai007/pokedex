package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("You must provide a location name")
	}
	exploreStr := args[0]
	data, err := cfg.pokeapiClient.ListPokemons(&exploreStr)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	fmt.Println()
	fmt.Printf("Pokemon according to area %s....\n", data.Name)
	for _, pokemon := range data.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	fmt.Println()
	return nil
}
