package cmd

const PackageMainEdit = `//go:build !js

package main

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	editOut string
	editDir bool
)

var editCmd = &cobra.Command{
	Use:     "edit [data/stage.go]",
	Aliases: []string{"edit-stageset", "stageset", "edit-multistage"},
	Short:   "Edit a stage file or directory",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			unmarshallFromCode = args[0]
			marshallOnCommit = args[0]
		}
		if editOut != "" {
			marshallOnCommit = editOut
		}
		if editDir && marshallOnCommit != "" {
			_ = os.MkdirAll(marshallOnCommit, 0755)
		}
		executeServer()
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().BoolVar(&embeddedDiagrams, "embedded-diagrams", false, "parse/analysis go/models and go/embeddedDiagrams")
	editCmd.Flags().IntVar(&port, "port", 8080, "port server")
	editCmd.Flags().StringVarP(&editOut, "out", "o", "", "specify a different file or directory to save commits to")
	editCmd.Flags().BoolVarP(&editDir, "dir", "d", false, "persist stage data as a Go module directory")
	editCmd.Flags().BoolVar(&skipGoMod, "skip-go-mod", false, "do not generate go.mod in directory persistence mode")
}
`
