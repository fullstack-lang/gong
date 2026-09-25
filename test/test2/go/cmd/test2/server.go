//go:build !js

package main

import (
	"log"
	"strconv"

	"github.com/fullstack-lang/gong/test/test2/go/level1stack"
	"github.com/fullstack-lang/gong/test/test2/go/models"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

func executeServer() {

	// setup
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1Stack("test2", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams)

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

	// initiates the UX loop
	models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
	)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := stack.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func executeServerStageSet() {

	// setup
	// - model level1 stack with its probe and stageset probe
	// - unmarshall/marshall go file with multi-stage data
	stack := level1stack.NewLevel1StackStageSet("test2", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams)

	// refresh the probes, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()
	if stack.StageSetProbe != nil {
		stack.StageSetProbe.Refresh()
	}

	rootSplitStage := split_stack.NewStack(stack.R, "", "", "", "", false, false).Stage

	if stack.StageSet != nil {
		rootSplitStage.StageBranch(&split.View{
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
