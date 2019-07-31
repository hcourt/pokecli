package cmd

import (
	"fmt"

	"github.com/mtslzr/pokeapi-go/structs"
)

// printResults prints the names of all results in a simple format
// TODO: support formats
func printResults(results []structs.Result) {
	fmt.Println("Found results:")
	for _, r := range results {
		fmt.Println(r.Name)
	}
}
