package cmd

import(
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "pokecli",
}

// Execute runs the cli tool
func Execute() error {
	fmt.Println("Hello world")
	return rootCmd.Execute()
}

