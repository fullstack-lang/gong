package cmd

const PackageMainSemantic = `//go:build !js

package main

import (
	"log"

	"github.com/spf13/cobra"
)

var semanticCmd = &cobra.Command{
	Use:   "semantic",
	Short: "Check semantic",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Check semantic")
	},
}

func init() {
	checkCmd.AddCommand(semanticCmd)
}
`
