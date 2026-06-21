package commands

import (
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check operations on the stage",
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
