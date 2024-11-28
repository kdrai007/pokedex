package main

import (
	"fmt"
	"time"

	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("You must provide name of pokemon!")
	}

	name := args[0]
	pokemonResp, err := cfg.pokeapiClient.FetchPokemon(&name)
	if err != nil {
		return err
	}
	randValue := getRandomNumber(50, pokemonResp.BaseExperience)
	catchPercent := (randValue * 100) / pokemonResp.BaseExperience

	fmt.Printf("Throwing a Pokeball at %s\n", name)
	time.Sleep(1 * time.Second)
	if catchPercent > 60 {
		cfg.catchedPokemon[pokemonResp.Name] = pokemonResp
		fmt.Printf("%s was caught!\n", name)
	} else {
		fmt.Printf("%s escaped!!\n", name)
	}

	return nil
}

func getRandomNumber(min, max int) int {
	// Ensure the seed is different each time the program runs
	return rand.Intn(max-min+1) + min
}
