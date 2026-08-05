package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/yohany99/pokedex/internal/pokeapi"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
		pokemonCaught: map[string]pokeapi.Pokemon{},
	}
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			args := cleanInput(input)
			if len(args) == 0 {
				continue
			}
			first := args[0]
			if val, ok := getCommands()[first]; ok {
				err := val.callback(cfg, args[1:]...)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
