package main

import (
	"time"

	"github.com/AxterDoesCode/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Second*5, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
