package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next     *string
	previous *string
}

type locResponse struct {
	Next      *string    `json:"next"`
	Previous  *string    `json:"previous"`
	Locations []location `json:"results"`
}

type location struct {
	Name string `json:"name"`
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
	}
}

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
	var res *http.Response
	var err error
	if cfg.next == nil {
		res, err = http.Get("https://pokeapi.co/api/v2/location-area/")
		if err != nil {
			return err
		}
	} else {
		res, err = http.Get(*cfg.next)
		if err != nil {
			return err
		}
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}
	locResponse := locResponse{}
	err = json.Unmarshal(body, &locResponse)
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
		fmt.Println("You're on the first page")
	} else {
		res, err := http.Get(*cfg.previous)
		if err != nil {
			return err
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return fmt.Errorf("response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			return err
		}
		locResponse := locResponse{}
		err = json.Unmarshal(body, &locResponse)
		if err != nil {
			return err
		}
		cfg.next = locResponse.Next
		cfg.previous = locResponse.Previous
		for _, location := range locResponse.Locations {
			fmt.Println(location.Name)
		}
	}
	return nil
}
