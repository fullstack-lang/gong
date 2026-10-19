//go:build !js

package main

import (
	"log"
	"strconv"

	"github.com/fullstack-lang/gong/test/test2/go/level1stack"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

func executeServer() {

	// setup
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with multi-stage data
	stack := level1stack.NewLevel1Stack("test2", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams)

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

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
	err := stack.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
