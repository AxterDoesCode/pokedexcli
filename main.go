package main

import (
  "time"
	"github.com/axterdoescode/pokedexcli/internal/pokeapi"
  "internal/pokeapi"
)

func main() {
  pokeClient := pokeapi.NewClient(time.Second * 5)
	cfg := &config{
    pokeapiClient: pokeClient
  }
	startRepl(cfg)
}
