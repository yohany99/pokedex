package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
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

func commandMap(cfg *config, args ...string) error {
	locResp, err := cfg.pokeapiClient.ListLocations(cfg.nextURL)
	if err != nil {
		return err
	}
	cfg.nextURL = locResp.NextURL
	cfg.previousURL = locResp.PreviousURL
	for _, location := range locResp.Locations {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousURL == nil {
		return errors.New("you're on the first page")
	}
	locResp, err := cfg.pokeapiClient.ListLocations(cfg.previousURL)
	if err != nil {
		return err
	}
	cfg.nextURL = locResp.NextURL
	cfg.previousURL = locResp.PreviousURL
	for _, location := range locResp.Locations {
		fmt.Println(location.Name)
	}
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide a location name")
	}
	url := "https://pokeapi.co/api/v2/location-area/"
	url += args[0]
	pokemonResp, err := cfg.pokeapiClient.ListPokemon(url)
	if err != nil {
		return err
	}
	for _, pkm := range pokemonResp.PokemonEncounters {
		fmt.Println(pkm.Pokemon.Name)
	}
	return nil
}
