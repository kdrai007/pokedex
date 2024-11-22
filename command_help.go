package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range fieldsValue() {
		fmt.Printf("%s: %s \n", cmd.name, cmd.description)
	}
	return nil
}
