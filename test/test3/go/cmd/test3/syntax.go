//go:build !js

package main

import (
	"log"

	"github.com/spf13/cobra"
)

var syntaxCmd = &cobra.Command{
	Use:   "syntax",
	Short: "Check syntax",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Check syntax")
	},
}

func init() {
	checkCmd.AddCommand(syntaxCmd)
}
