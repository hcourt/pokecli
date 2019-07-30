package main

import (
	"log"

	"github.com/hcourt/pokecli/src/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
