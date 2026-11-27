//go:build !js

// generated code - do not edit
package main

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/fullstack-lang/gong/lib/xlsx/go/level1stack"
	"github.com/fullstack-lang/gong/lib/xlsx/go/models"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
	split_static "github.com/fullstack-lang/gong/lib/split/go/static"
)

var (
	embeddedDiagrams   bool
	port               int
	unmarshallFromCode string
	marshallOnCommit   string
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

func executeServer() {
	// setup
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1Stack("xlsx", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams)

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

	// Create root split stage for the probe
	rootSplitStage := split_stack.NewStack(stack.R, "", "", "", "", false, false).Stage

	rootSplitStage.StageBranch(&split.View{
		Name: "Data Probe & Data Model",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stack.Stage.GetProbeSplitStageName(),
				},
			},
		},
	})
	rootSplitStage.Commit()

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := split_static.RunServer(stack.R, ":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

var (
	migrateOut string
	migrateDir bool
)

var migrateCmd = &cobra.Command{
	Use:   "migrate [data/stage.go]",
	Short: "Migrate a single-stage data file to multi-stage format or directory module",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := "data/stage.go"
		if len(args) > 0 {
			inputFile = args[0]
		}
		outputFile := migrateOut
		if outputFile == "" {
			if migrateDir {
				outputFile = "data_dir"
			} else {
				outputFile = "data/stageset.go"
			}
		}

		stage := models.NewStage("")
		if err := stage.ParseAstFile(inputFile, true); err != nil {
			log.Fatalf("failed to parse input stage file: %v", err)
		}
		if migrateDir {
			if err := os.MkdirAll(outputFile, 0755); err != nil {
				log.Fatalf("failed to create output directory: %v", err)
			}
			if err := stage.MarshallDirectory(outputFile); err != nil {
				log.Fatalf("failed to marshall stage directory: %v", err)
			}
		} else {
			stage.MarshallFile(outputFile, "main")
		}
		log.Printf("Successfully migrated %s to %s", inputFile, outputFile)
	},
}

var rootCmd = &cobra.Command{
	Use:   "probe",
	Short: "probe CLI for xlsx",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			unmarshallFromCode = args[0]
			marshallOnCommit = args[0]
		}
		executeServer()
	},
}

func main() {
	editCmd.Flags().StringVarP(&editOut, "out", "o", "", "specify a different file or directory to save commits to")
	editCmd.Flags().BoolVarP(&editDir, "dir", "d", false, "persist stage data as a Go module directory")
	rootCmd.AddCommand(editCmd)
	migrateCmd.Flags().StringVarP(&migrateOut, "out", "o", "", "output file or directory path (default: data/stageset.go or data_dir)")
	migrateCmd.Flags().BoolVarP(&migrateDir, "dir", "d", false, "migrate into a Go module directory")
	rootCmd.AddCommand(migrateCmd)
	rootCmd.PersistentFlags().BoolVar(&embeddedDiagrams, "embedded-diagrams", true, "parse/analysis go/models and go/embeddedDiagrams")
	rootCmd.PersistentFlags().IntVar(&port, "port", 8080, "port server")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
