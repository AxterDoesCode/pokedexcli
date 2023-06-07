package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a pokemon name to inspect")
	}
	name := args[0]
	if pokemon, ok := cfg.caughtPokemon[name]; ok {
		fmt.Printf("Name : %s\n", pokemon.Name)
		fmt.Printf("Height : %d\n", pokemon.Height)
		fmt.Printf("Weight : %d\n", pokemon.Weight)
		fmt.Println("Stats :")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}

		fmt.Println("Types:")
		for _, typeInfo := range pokemon.Types {
			fmt.Println("  -", typeInfo.Type.Name)
		}

	} else {
		return errors.New("You have not caught that Pokemon yet!")
	}
	return nil
}
