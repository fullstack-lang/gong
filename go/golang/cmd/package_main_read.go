package cmd

const PackageMainRead = `//go:build !js

package main

import (
	"log"

	"github.com/spf13/cobra"
)

var readDir bool

var readCmd = &cobra.Command{
	Use:   "read [data/stage.go]",
	Short: "Read a stage file or directory",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			unmarshallFromCode = args[0]
		}
		log.Println("Read a stage file", unmarshallFromCode)

		// executeServer(args)
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
	readCmd.Flags().BoolVarP(&readDir, "dir", "d", false, "read stage data from a Go module directory")
}
`
