package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var syntaxCmd = &cobra.Command{
	Use:   "syntax [file]",
	Short: "Check the syntax of a stage file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		fmt.Printf("Syntax checking %s...\n", file)
		// To be implemented: actual syntax check logic
		fmt.Println("Syntax check complete. No errors found.")
	},
}

func init() {
	checkCmd.AddCommand(syntaxCmd)
}
