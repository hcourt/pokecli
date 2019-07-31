package cmd

import (
	"github.com/spf13/cobra"
)

var logVerbose bool

var rootCmd = &cobra.Command{
	Use: "pokecli",
	Short: "pokecli is a command line interface (CLI) for PokéAPI",
	Long: `A simple command line interface wrapper for PokéAPI written by
                hcourt in Go using pokeapi-go.
                Complete documentation is available at https://pokeapi.co`,
}

type allFlags struct {
	*searchFlags
	*showFlags
}

var rootFlags = &allFlags{}

// Execute runs the cli tool
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&logVerbose, "verbose", "v", false, "enable logging")
}
