package main

import (
	"time"

	"github.com/kdrai007/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		catchedPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient:  pokeClient,
	}
	startRepl(cfg)
}
