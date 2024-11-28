package main

import (
	"errors"
	"fmt"
)

type locationData struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ApiResponse struct {
	Count   int            `json:"count"`
	Next    string         `json:"next"`
	Prev    string         `json:"previous"`
	Results []locationData `json:"results"`
}

var apiResponse ApiResponse

func commandMap(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationUrl = &locationsResp.Next
	cfg.prevLocationUrl = &locationsResp.Prev

	fmt.Println()
	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	fmt.Println()
	return nil
}

func commandMapB(cfg *config, args ...string) error {
	if cfg.prevLocationUrl == nil {
		return errors.New("!!No previous locations")
	}
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationUrl = &locationsResp.Next
	cfg.prevLocationUrl = &locationsResp.Prev

	fmt.Println()
	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	fmt.Println()
	return nil
}
