//go:build !js

package main

import (
	"github.com/spf13/cobra"
)

var editStageSetOut string

var editStageSetCmd = &cobra.Command{
	Use:     "edit-stageset [data/stageset.go]",
	Aliases: []string{"stageset", "edit-multistage"},
	Short:   "Edit a multi-stage StageSet file (temporary command)",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			unmarshallFromCode = args[0]
			marshallOnCommit = args[0]
		}
		if editStageSetOut != "" {
			marshallOnCommit = editStageSetOut
		}
		executeServerStageSet()
	},
}

func init() {
	rootCmd.AddCommand(editStageSetCmd)
	editStageSetCmd.Flags().BoolVar(&embeddedDiagrams, "embedded-diagrams", false, "parse/analysis go/models and go/embeddedDiagrams")
	editStageSetCmd.Flags().IntVar(&port, "port", 8080, "port server")
	editStageSetCmd.Flags().StringVar(&editStageSetOut, "out", "", "specify a different file to save commits to")
}
