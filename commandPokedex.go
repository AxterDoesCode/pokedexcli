package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("No Pokemon in your Pokedex")
	}
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("  -%s\n", pokemon.Name)
	}
	return nil
}
