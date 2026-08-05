package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/yohany99/pokedex/internal/pokeapi"
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
	pkmEncountersResp, err := cfg.pokeapiClient.ListPokemon(url)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon: ")
	for _, pkm := range pkmEncountersResp.Encounters {
		fmt.Println(pkm.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide a pokemon name")
	}
	pokemon, err := cfg.pokeapiClient.CatchPokemon(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	randomNumber := rand.IntN(pokemon.Exp) + 1
	if randomNumber <= 40 {
		fmt.Printf("%s was caught!\n", args[0])
		cfg.pokemonCaught[args[0]] = pokemon

	} else {
		fmt.Printf("%s escaped!\n", args[0])
	}
	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide a pokemon name")
	}
	if pokemon, ok := cfg.pokemonCaught[args[0]]; ok {
		pokeapi.PrintStats(pokemon)
		return nil
	} else {
		return errors.New("you have not caught that pokemon")
	}
}
