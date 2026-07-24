package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *config) error {
	locResponse, err := cfg.pokeapiClient.ListLocations(cfg.next)
	if err != nil {
		return err
	}
	cfg.next = locResponse.Next
	cfg.previous = locResponse.Previous
	for _, location := range locResponse.Locations {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previous == nil {
		return errors.New("you're on the first page")
	}
	locResponse, err := cfg.pokeapiClient.ListLocations(cfg.previous)
	if err != nil {
		return err
	}
	cfg.next = locResponse.Next
	cfg.previous = locResponse.Previous
	for _, location := range locResponse.Locations {
		fmt.Println(location.Name)
	}
	return nil
}
