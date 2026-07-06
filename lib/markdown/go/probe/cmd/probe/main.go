//go:build !js

package main

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/fullstack-lang/gong/lib/markdown/go/stack"
	"github.com/fullstack-lang/gong/lib/markdown/go/static"

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
	// setup the static file server and get the controller
	r := static.ServeStaticFiles(false)

	// setup model stack with its probe
	stack := stack.NewStack(r, "markdown", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

	// Create root split stage for the probe
	rootSplitStage := split_stack.NewStack(r, "", "", "", "", false, false).Stage

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
	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

var rootCmd = &cobra.Command{
	Use:   "probe",
	Short: "probe CLI for markdown",
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
	rootCmd.PersistentFlags().BoolVar(&embeddedDiagrams, "embeddedDiagrams", true, "parse/analysis go/models and go/embeddedDiagrams")
	rootCmd.PersistentFlags().IntVar(&port, "port", 8080, "port server")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
