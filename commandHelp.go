package main

import (
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
	for _, x := range getCommands() {
		fmt.Printf("%s : %s\n", x.name, x.description)
	}
	return nil
}
