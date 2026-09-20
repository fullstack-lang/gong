//go:build !js

package main

import (
	"log"

	"github.com/fullstack-lang/gong/test/test2/go/models"
	"github.com/spf13/cobra"
)

var migrateOut string

var migrateCmd = &cobra.Command{
	Use:   "migrate [data/stage.go]",
	Short: "Migrate a single-stage data file to multiple-stage StageSet format",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := "data/stage.go"
		if len(args) > 0 {
			inputFile = args[0]
		}
		outputFile := migrateOut
		if outputFile == "" {
			outputFile = "data/stageset.go"
		}

		stage := models.NewStage("")
		if err := stage.ParseAstFile(inputFile, true); err != nil {
			log.Fatalf("failed to parse input stage file: %v", err)
		}
		stageSet := models.NewStageSetFromStage(stage)
		stageSet.MarshallFile(outputFile, "main")
		log.Printf("Successfully migrated %s to %s", inputFile, outputFile)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().StringVar(&migrateOut, "out", "", "output file path (default: data/stageset.go)")
}
