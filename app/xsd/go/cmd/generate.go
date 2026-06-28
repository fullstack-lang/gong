/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// generateCmd represents the read command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "read an xsd file and generate the gong models",
	Args:  cobra.ExactArgs(1),
	Long:  `read an xsd file and generate the gong models`,
	Run: func(cmd *cobra.Command, args []string) {

		process(args)

	},
}
