package commands

import (
	"github.com/spf13/cobra"
)

var readPort int
var readEmbeddedDiagrams bool

var readCmd = &cobra.Command{
	Use:   "read [file]",
	Short: "Read a stage file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]

		executeServer(file, "", readPort, readEmbeddedDiagrams, logGINFlag)
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
	readCmd.Flags().IntVarP(&readPort, "port", "p", 8080, "port server")
	readCmd.Flags().BoolVar(&readEmbeddedDiagrams, "embedded-diagrams", false, "parse/analysis go/models and go/embeddedDiagrams")
}
