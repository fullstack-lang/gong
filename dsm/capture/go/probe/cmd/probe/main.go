//go:build !js

package main

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/fullstack-lang/gong/dsm/capture/go/level1stack"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
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
	stack := level1stack.NewLevel1Stack("capture", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams)

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
	err := stack.R.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

var rootCmd = &cobra.Command{
	Use:   "probe",
	Short: "probe CLI for capture",
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
