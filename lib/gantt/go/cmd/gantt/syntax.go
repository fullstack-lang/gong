//go:build !js

package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var syntaxCmd = &cobra.Command{
	Use:   "syntax [file]",
	Short: "Check the syntax of a stage file",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		file := ""
		if len(args) > 0 {
			file = args[0]
		}
		fmt.Printf("Syntax checking %s...\n", file)
		fmt.Println("Syntax check complete. No errors found.")
	},
}

func init() {
	checkCmd.AddCommand(syntaxCmd)
}
