package main

import (
	"time"
	
	"github.com/sabrek15/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startrepl(cfg)
}
