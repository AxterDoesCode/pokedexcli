package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	data, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = data.Next
	cfg.prevLocationURL = data.Previous

	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationURL == nil {
		return errors.New("You are on the first page")
	}

	data, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = data.Next
	cfg.prevLocationURL = data.Previous

	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
