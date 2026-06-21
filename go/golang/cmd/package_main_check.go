package cmd

const PackageMainCheck = `//go:build !js

package main

import (
	"log"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check operations on the stage",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Check a stage file")
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
`
