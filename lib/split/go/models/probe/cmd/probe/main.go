//go:build !js

// generated code - do not edit
package main

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/fullstack-lang/gong/lib/split/go/stack"
	"github.com/fullstack-lang/gong/lib/split/go/static"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

var (
	embeddedDiagrams   bool
	port               int
	unmarshallFromCode string
	marshallOnCommit   string
	editOut            string
)

var editCmd = &cobra.Command{
	Use:     "edit [data/stage.go]",
	Aliases: []string{"edit-stageset", "stageset", "edit-multistage"},
	Short:   "Edit a stage file",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			unmarshallFromCode = args[0]
			marshallOnCommit = args[0]
		}
		if editOut != "" {
			marshallOnCommit = editOut
		}
		executeServer()
	},
}

func executeServer() {
	// setup the static file server and get the controller
	r := static.ServeStaticFiles(false)

	// setup model stack with its probe
	stack := stack.NewStack(r, "split", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

	// Create root split stage for the probe
	rootSplitStage := split_stack.NewStack(r, "", "", "", "", false, false).Stage

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
	err := static.RunServer(r, ":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

var rootCmd = &cobra.Command{
	Use:   "probe",
	Short: "probe CLI for split",
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
	editCmd.Flags().StringVar(&editOut, "out", "", "specify a different file to save commits to")
	rootCmd.AddCommand(editCmd)
	rootCmd.PersistentFlags().BoolVar(&embeddedDiagrams, "embedded-diagrams", true, "parse/analysis go/models and go/embeddedDiagrams")
	rootCmd.PersistentFlags().IntVar(&port, "port", 8080, "port server")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
