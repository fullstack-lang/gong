//go:build !js

package main

import (
	"github.com/spf13/cobra"
)

var readDir bool

var readCmd = &cobra.Command{
	Use:   "read [file]",
	Short: "Read a stage file or directory",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			unmarshallFromCode = args[0]
		}

		executeServer(args)
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
	readCmd.Flags().IntVarP(&port, "port", "p", 8080, "port server")
	readCmd.Flags().BoolVar(&embeddedDiagrams, "embedded-diagrams", false, "parse/analysis go/models and go/embeddedDiagrams")
	readCmd.Flags().BoolVarP(&readDir, "dir", "d", false, "read stage data from a Go module directory")
}
