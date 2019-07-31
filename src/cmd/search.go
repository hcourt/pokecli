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
	entityType string
}

var (
	supportedTypes = []string{
		"pokemon",
		"move",
	}
	flags *searchFlags
)

func init() {
	flags = &searchFlags{}
	var cmd = &cobra.Command{
		Use:   "search",
		Short: "Show information about an entity",
		RunE:  search,
		Args:  cobra.ExactArgs(1),
	}
	if err := flags.addToCmd(cmd); err != nil {
		log.Fatal(err)
	}
	rootCmd.AddCommand(cmd)
}

func (f *searchFlags) addToCmd(cmd *cobra.Command) error {
	const typeFlag = "type"
	cmd.Flags().StringVarP(
		&f.entityType,
		typeFlag,
		"t",
		supportedTypes[0],
		fmt.Sprintf("The type of entity to search for.  Supported types: %s", supportedTypes))
	return cmd.MarkFlagRequired(typeFlag)
}

// search is a dictionary search for one or more entities by type and name.
// This only supports a max of 9999 results due to limitations in pokeapi-go.
func search(cmd *cobra.Command, entities []string) error {
	cmd.SilenceUsage = true
	var errs error
	var results []structs.Result
	for _, entity := range entities {
		found, err := pokeapi.Search(flags.entityType, entity)
		if err != nil {
			// Fail quickly for I/O errors
			return err
		}
		if found.Count == 0 {
			errs = multierr.Append(errs, fmt.Errorf("entity not found: %s (type %s)", entity, flags.entityType))
		}
		results = append(results, found.Results...)
	}
	if errs != nil {
		return errs
	}
	printResults(results)
	return nil
}
