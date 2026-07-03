package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		if scanner.Scan() {
			input := scanner.Text()
			first := cleanInput(input)[0]
			fmt.Printf("Your command was: %s\n", first)
		}
	}
}
