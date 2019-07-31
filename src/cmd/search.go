package cmd

import (
	"fmt"
	"log"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/spf13/cobra"
	"go.uber.org/multierr"
)

type searchFlags struct {
	entityFlags
}

func init() {
	flags := &searchFlags{}
	cmd := &cobra.Command{
		Use:   "search",
		Short: "Search for entities by type and name",
		RunE:  search,
		Args:  cobra.MinimumNArgs(1),
	}
	if err := flags.addToCmd(cmd); err != nil {
		log.Fatal(err)
	}
	rootFlags.searchFlags = flags
	rootCmd.AddCommand(cmd)
}

// search is a dictionary search for one or more entities by type and name.
// This only supports a max of 9999 results due to limitations in pokeapi-go.
func search(cmd *cobra.Command, entities []string) error {
	cmd.SilenceUsage = true
	var errs error
	var results []structs.Result
	for _, entity := range entities {
		found, err := pokeapi.Search(rootFlags.searchFlags.entityType, entity)
		if err != nil {
			// Fail quickly for I/O errors
			return err
		}
		if found.Count == 0 {
			errs = multierr.Append(errs, fmt.Errorf("entity not found: %s (type %s)", entity, rootFlags.searchFlags.entityType))
		}
		results = append(results, found.Results...)
	}
	if errs != nil {
		return errs
	}
	printResults(results)
	return nil
}
