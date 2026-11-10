//go:build !js

package main

import (
	"log"

	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read [data/stage.go]",
	Short: "Read a stage file",
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
}
