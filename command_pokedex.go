package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.catchedPokemon) == 0 {
		fmt.Println("No pokemon yet!!")
	}

	fmt.Println("Your Pokedex: ")
	for _, pokemon := range cfg.catchedPokemon {
		fmt.Println("-", pokemon.Name)
	}
	fmt.Println()
	return nil
}
