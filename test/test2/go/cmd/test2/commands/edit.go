package commands

import (
	"github.com/spf13/cobra"
)

var editPort int
var editEmbeddedDiagrams bool
var editOut string

var editCmd = &cobra.Command{
	Use:   "edit [file]",
	Short: "Edit a stage file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]

		marshallOnCommit := file
		if editOut != "" {
			marshallOnCommit = editOut
		}

		executeServer(file, marshallOnCommit, editPort, editEmbeddedDiagrams, logGINFlag)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().IntVarP(&editPort, "port", "p", 8080, "port server")
	editCmd.Flags().BoolVar(&editEmbeddedDiagrams, "embedded-diagrams", false, "parse/analysis go/models and go/embeddedDiagrams")
	editCmd.Flags().StringVarP(&editOut, "out", "o", "", "specify a different file to save commits to")
}
