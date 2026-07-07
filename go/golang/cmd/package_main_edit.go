package cmd

const PackageMainEdit = `//go:build !js

package main

import (
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [data/stage.go]",
	Short: "Edit a stage file",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			unmarshallFromCode = args[0]
			marshallOnCommit = args[0]
		}
		executeServer()
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().BoolVar(&embeddedDiagrams, "embedded-diagrams", false, "parse/analysis go/models and go/embeddedDiagrams")
	editCmd.Flags().IntVar(&port, "port", 8080, "port server")
}
`
