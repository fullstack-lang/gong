package commands

import (
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Validation commands (syntax, semantic, etc.)",
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
