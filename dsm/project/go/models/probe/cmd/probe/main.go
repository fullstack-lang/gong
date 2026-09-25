//go:build !js

// generated code - do not edit
package main

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/fullstack-lang/gong/dsm/project/go/level1stack"

	splitlite "github.com/fullstack-lang/gong/lib/splitlite/go/models"
	splitlite_stack "github.com/fullstack-lang/gong/lib/splitlite/go/stack"
	splitlite_static "github.com/fullstack-lang/gong/lib/splitlite/go/static"
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
	stack := level1stack.NewLevel1Stack("project", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams)

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

	// Create root split stage for the probe
	rootSplitStage := splitlite_stack.NewStack(stack.R, "", "", "", "", false, false).Stage

	rootSplitStage.StageBranch(&splitlite.View{
		Name: "Data Probe & Data Model",
		RootAsSplitAreas: []*splitlite.AsSplitArea{
			{
				Split: &splitlite.Split{
					StackName: stack.Stage.GetProbeSplitStageName(),
				},
			},
		},
	})
	rootSplitStage.Commit()

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := splitlite_static.RunServer(stack.R, ":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

var rootCmd = &cobra.Command{
	Use:   "probe",
	Short: "probe CLI for project",
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
	rootCmd.PersistentFlags().BoolVar(&embeddedDiagrams, "embedded-diagrams", true, "parse/analysis go/models and go/embeddedDiagrams")
	rootCmd.PersistentFlags().IntVar(&port, "port", 8080, "port server")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
