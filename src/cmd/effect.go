package cmd

import (
	"fmt"
	"log"

	"github.com/hcourt/pokecli/src/simplestructs"
	"github.com/mtslzr/pokeapi-go"
	"github.com/spf13/cobra"
)

const (
	pokemonFlag          = "pokemon"
	moveFlag             = "move"
	nonDamagingMoveClass = "status"
)

type effectFlags struct {
	pokemonName string
	moveName    string
}

func (f *effectFlags) addToCmd(cmd *cobra.Command) error {
	cmd.Flags().StringVarP(
		&f.pokemonName,
		pokemonFlag,
		"p",
		"",
		"Name of the defending pokemon.",
	)
	if err := cmd.MarkFlagRequired(pokemonFlag); err != nil {
		return err
	}
	cmd.Flags().StringVarP(
		&f.moveName,
		moveFlag,
		"m",
		"",
		"Name of the attacking move.",
	)
	return cmd.MarkFlagRequired(moveFlag)
}

func init() {
	flags := &effectFlags{}
	cmd := &cobra.Command{
		Use:   "effect",
		Short: "Check type effectiveness of an attack move against a pokemon.",
		RunE:  effect,
		Args:  cobra.NoArgs,
	}
	if err := flags.addToCmd(cmd); err != nil {
		log.Fatal(err)
	}
	rootFlags.effectFlags = flags
	rootCmd.AddCommand(cmd)
}

// effect calculates the damage multiplier of one move against a defending
// pokemon, and prints a summary message.
func effect(cmd *cobra.Command, _ []string) error {
	move, err := pokeapi.Move(rootFlags.moveName)
	if err != nil {
		return err
	}

	if move.DamageClass.Name == nonDamagingMoveClass {
		fmt.Printf("Move is a %s move and will not cause typed damage.", nonDamagingMoveClass)
		return nil
	}

	poke, err := pokeapi.Pokemon(rootFlags.pokemonName)
	if err != nil {
		return err
	}

	moveType := simplestructs.SimpleType{Name: move.Type.Name}
	var pokeTypes []*simplestructs.SimpleType
	for _, t := range poke.Types {
		pokeType, err := pokeapi.Type(t.Type.Name)
		if err != nil {
			return err
		}
		simpleType := simplestructs.SimpleType(pokeType)
		pokeTypes = append(pokeTypes, &simpleType)
	}

	result := moveType.EffectMulti(pokeTypes)
	if logVerbose {
		printEffect(&moveType, pokeTypes, &result)
	} else {
		fmt.Printf("%v", &result)
	}
	return nil
}
