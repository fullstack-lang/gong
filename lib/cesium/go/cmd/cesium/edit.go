//go:build !js

package main

import (
	"github.com/spf13/cobra"
)

var editOut string

var editCmd = &cobra.Command{
	Use:     "edit [data/stage.go]",
	Aliases: []string{"edit-stageset", "stageset", "edit-multistage"},
	Short:   "Edit a stage file",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			unmarshallFromCode = args[0]
			marshallOnCommit = args[0]
		}
		if editOut != "" {
			marshallOnCommit = editOut
		}
		executeServer()
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().BoolVar(&embeddedDiagrams, "embedded-diagrams", false, "parse/analysis go/models and go/embeddedDiagrams")
	editCmd.Flags().IntVar(&port, "port", 8080, "port server")
	editCmd.Flags().StringVar(&editOut, "out", "", "specify a different file to save commits to")
}
