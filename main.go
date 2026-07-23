package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/yohany99/pokedex/internal/pokecache"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{}
	cache := pokecache.NewCache(5 * time.Second)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			first := cleanInput(input)[0]
			if val, ok := getCommands()[first]; ok {
				err := val.callback(cfg)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}
	}
}
