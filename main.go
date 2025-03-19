package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	cfg := &config{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanned := scanner.Scan()
		if !scanned {
			break
		}
		input := scanner.Text()

		if input == "" {
			continue
		}

		words := cleanInput(input)

		if cmd, exists := commands[words[0]]; exists {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Println("unknown Command")
		}
	}
}
