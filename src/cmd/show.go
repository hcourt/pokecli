package cmd

import (
	"fmt"
	"log"

	"github.com/hcourt/pokecli/src/simplestructs"
	"github.com/mtslzr/pokeapi-go"
	"github.com/spf13/cobra"
)

type showFlags struct {
	entityFlags
}

type showFunc func(id string) error

var (
	supportedShowFuncs = map[string]showFunc{
		"pokemon": showPokemon,
		"move":    showMove,
	}
)

func init() {
	flags := &showFlags{}
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Show information about an entity",
		RunE:  show,
		Args:  cobra.ExactArgs(1),
	}
	if err := flags.addToCmd(cmd); err != nil {
		log.Fatal(err)
	}
	rootFlags.showFlags = flags
	rootCmd.AddCommand(cmd)
}

func show(cmd *cobra.Command, entities []string) error {
	cmd.SilenceUsage = true
	entity := entities[0]
	showFunc := supportedShowFuncs[rootFlags.showFlags.entityType]
	if showFunc == nil {
		return fmt.Errorf("entity type not supported (supported types: %s)", supportedEntityTypes)
	}

	return showFunc(entity)
}

func showPokemon(id string) error {
	result, err := pokeapi.Pokemon(id)
	if err != nil {
		return err
	}
	simple := simplestructs.SimplePokemon(result)
	fmt.Println(&simple)
	return nil
}

func showMove(id string) error {
	result, err := pokeapi.Move(id)
	if err != nil {
		return err
	}
	simple := simplestructs.SimpleMove(result)
	fmt.Println(&simple)
	return nil
}
