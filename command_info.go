package main

import (
	"fmt"
)

func commandInfo(cfg *config, args ...string) error {
	pokemonName := args[0]
	pokemonData, err := cfg.pokeapiClient.FetchPokemon(&pokemonName)
	if err != nil {
		return fmt.Errorf("No pokemon found!! \n error: %v", err)
	}
	fmt.Printf("Looking at stats of %s...\n", pokemonName)
	fmt.Println()
	fmt.Println("Name: ", pokemonData.Name)
	fmt.Println("Height: ", pokemonData.Height)
	fmt.Println("Weight: ", pokemonData.Weight)
	fmt.Println("stats: ")
	for _, stat := range pokemonData.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types: ")
	for _, typ := range pokemonData.Types {
		fmt.Printf(" - %s\n", typ.Type.Name)
	}
	fmt.Println()

	return nil
}
