package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range fieldsValue() {
		fmt.Printf("%s: %s \n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
