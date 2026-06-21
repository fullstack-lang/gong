package commands

import (
	"log"

	"github.com/fullstack-lang/gong/test/test3/go/level1stack"
	"github.com/spf13/cobra"
)

var syntaxCmd = &cobra.Command{
	Use:   "syntax [file]",
	Short: "Check the syntax of a stage file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		log.Printf("Checking syntax of %s...", file)

		// To check syntax, we simply load the stack.
		// If there are parsing errors, ParseAstFile will log them.
		// We disable probe, no port, no embedded diagrams for a pure syntax check.
		_ = level1stack.NewLevel1StackDelta("test3", file, "", false, false, false)

		log.Println("Syntax check completed.")
	},
}

func init() {
	checkCmd.AddCommand(syntaxCmd)
}
