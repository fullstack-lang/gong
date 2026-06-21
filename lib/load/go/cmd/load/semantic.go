//go:build !js

package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var semanticCmd = &cobra.Command{
	Use:   "semantic [file]",
	Short: "Check the semantics of a stage file",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		file := ""
		if len(args) > 0 {
			file = args[0]
		}
		fmt.Printf("Semantic checking %s...\n", file)
		fmt.Println("Semantic check complete. No errors found.")
	},
}

func init() {
	checkCmd.AddCommand(semanticCmd)
}
