//go:build !js

// generated code - do not edit
package main

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/fullstack-lang/gong/test/test2/go/level1stack"
	"github.com/fullstack-lang/gong/test/test2/go/models"

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

var editCmd = &cobra.Command{
	Use:   "edit [data/stage.go]",
	Short: "Edit a stage file",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			unmarshallFromCode = args[0]
			marshallOnCommit = args[0]
		}
		executeServer()
	},
}

func executeServer() {
	// setup
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1Stack("test2", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams)

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

	// Create root split stage for the probe
	rootSplitStage := split_stack.NewStack(stack.R, "", "", "", "", false, false).Stage

	split.StageBranch(rootSplitStage, &split.View{
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

var editStageSetCmd = &cobra.Command{
	Use:     "edit-stageset [data/stage.go]",
	Aliases: []string{"stageset", "edit-multistage"},
	Short:   "Edit a multi-stage StageSet file (temporary command)",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			unmarshallFromCode = args[0]
			marshallOnCommit = args[0]
		}
		executeServerStageSet()
	},
}

func executeServerStageSet() {
	stack := level1stack.NewLevel1StackStageSet("test2", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams)

	stack.Probe.Refresh()
	if stack.StageSetProbe != nil {
		stack.StageSetProbe.Refresh()
	}

	rootSplitStage := split_stack.NewStack(stack.R, "", "", "", "", false, false).Stage

	if stack.StageSet != nil {
		split.StageBranch(rootSplitStage, &split.View{
			Name: "StageSet Probe",
			RootAsSplitAreas: []*split.AsSplitArea{
				{
					Split: &split.Split{
						StackName: stack.StageSet.GetProbeSplitStageName(),
					},
				},
			},
		})
	}

	split.StageBranch(rootSplitStage, &split.View{
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

var rootCmd = &cobra.Command{
	Use:   "probe",
	Short: "probe CLI for test2",
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
	rootCmd.AddCommand(editCmd)
	rootCmd.AddCommand(editStageSetCmd)
	migrateCmd.Flags().StringVar(&migrateOut, "out", "", "output file path (default: data/stageset.go)")
	rootCmd.AddCommand(migrateCmd)
	rootCmd.PersistentFlags().BoolVar(&embeddedDiagrams, "embedded-diagrams", true, "parse/analysis go/models and go/embeddedDiagrams")
	rootCmd.PersistentFlags().IntVar(&port, "port", 8080, "port server")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
