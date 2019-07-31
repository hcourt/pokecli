package cmd

import (
	"fmt"

	"github.com/hcourt/pokecli/src/simplestructs"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/spf13/cobra"
)

const typeFlag = "type"

var (
	supportedEntityTypes = []string{
		"pokemon",
		"move",
	}
)

type entityFlags struct {
	entityType string
}

func (f *entityFlags) addToCmd(cmd *cobra.Command) error {
	cmd.Flags().StringVarP(
		&f.entityType,
		typeFlag,
		"t",
		supportedEntityTypes[0],
		fmt.Sprintf("The type of entity to search for.  Supported types: %s", supportedEntityTypes),
	)
	return cmd.MarkFlagRequired(typeFlag)
}

// printResults prints the names of all results in a simple format
// TODO: support formats
func printResults(results []structs.Result) {
	for _, r := range results {
		fmt.Println(r.Name)
	}
}

func printEffect(attackType *simplestructs.SimpleType, defendTypes []*simplestructs.SimpleType, effect *simplestructs.DamageEffect) {
	fmt.Printf("If a %s move attacks a %v pokemon, the damage is %s.", attackType, defendTypes, effect)
}
