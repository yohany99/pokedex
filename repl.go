package main

import (
	"strings"

	"github.com/yohany99/pokedex/internal/pokeapi"
)

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lower := strings.ToLower(trimmed)
	return strings.Fields(lower)
}

type config struct {
	pokeapiClient *pokeapi.Client
	nextURL       *string
	previousURL   *string
	pokemonCaught map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the names of the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of the previous 20 locations areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location>",
			description: "List the names of all the Pokemon located in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "See details about a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Look at your pokedex",
			callback:    commandPokedex,
		},
	}
}
