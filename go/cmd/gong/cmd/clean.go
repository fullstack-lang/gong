package cmd

import (
	"log"

	"github.com/fullstack-lang/gong/go/golang"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean [path to models package]",
	Short: "Removes generated files",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.SetPrefix("gong: ")
		log.SetFlags(0)

		var pkgPath string
		if len(args) > 0 {
			pkgPath = args[0]
		} else {
			pkgPath = "."
		}

		log.Println("Cleaning generated files in", pkgPath)
		golang.RemoveGeneratedGongFilesButDocs(pkgPath)
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
