package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a Pokemon Name")
	}

	pokemonName := strings.Join(args, " ")

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	rand.Seed(time.Now().UnixNano())

	catch := rand.Intn(pokemon.BaseExperience) < 40
	if catch {
		fmt.Printf("You caught %s\n", pokemon.Name)
		cfg.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("Failed to catch %s\n", pokemon.Name)
	}

	return nil
}
