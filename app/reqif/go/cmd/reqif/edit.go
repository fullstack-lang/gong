//go:build !js

package main

import (
	"github.com/spf13/cobra"
)

var editOut string

var editCmd = &cobra.Command{
	Use:   "edit [file]",
	Short: "Edit a stage file",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			unmarshallFromCode = args[0]
			marshallOnCommit = args[0]
		}

		if editOut != "" {
			marshallOnCommit = editOut
		}

		executeServer(args)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().IntVarP(&port, "port", "p", 8080, "port server")
	editCmd.Flags().BoolVar(&embeddedDiagrams, "embedded-diagrams", false, "parse/analysis go/models and go/embeddedDiagrams")
	editCmd.Flags().StringVarP(&editOut, "out", "o", "", "specify a different file to save commits to")
}
