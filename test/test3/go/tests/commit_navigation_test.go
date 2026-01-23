package tests

import (
	"log"
	"strconv"
	"testing"

	"github.com/fullstack-lang/gong/test/test3/go/level1stack"
	"github.com/fullstack-lang/gong/test/test3/go/models"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

func TestCommitNavigation(t *testing.T) {

	// setup
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1Stack("test3", "", "", true, true)
	stage := stack.Stage

	stage.SetDeltaMode(true)

	// load initial data into stage
	err := models.ParseAstFile2(
		stage,
		"./stage.go",
		false)
	if err != nil {
		t.Errorf("failed to parse stage.go: %v", err)
	}
	stage.Commit()

	// 1. creates A and B instances in the stage, commit
	a1 := (&models.A{Name: "A1"}).Stage(stage)
	b1 := (&models.B{Name: "B1"}).Stage(stage)
	stage.Commit()

	// 2. creates/deletes/updates A and B instances, commit
	// Update
	a1.Name = "A1 Updated"
	// Create
	a2 := (&models.A{Name: "A2"}).Stage(stage)
	// Delete
	b1.Unstage(stage)
	stage.Commit()

	// 3. creates/deletes/updates A and B instances, commit
	// Update
	a2.Name = "A2 Updated"
	// Create
	b2 := (&models.B{Name: "B2"}).Stage(stage)
	// Delete
	a1.Unstage(stage)
	stage.Commit()

	// 4. creates/deletes/updates A and B instances, commit
	// Update
	b2.Name = "B2 Updated"
	// Create
	a3 := (&models.A{Name: "A3"}).Stage(stage)
	_ = a3
	// Delete
	a2.Unstage(stage)
	stage.Commit()

	// navigate backward the commits one by one
	for i := 0; i < stage.GetNbBackwardCommits(); i++ {
		err := stage.ApplyBackwardCommit()
		if err != nil {
			t.Errorf("failed to apply backward commit: %v", err)
		}
	}

	// vanilla setup of the stager to be able to run the server
	splitStage := split_stack.NewStack(stack.R, "", "", "", "", false, false).Stage

	split.StageBranch(splitStage, &split.View{
		Name: "Data Probe & Data Model",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stage.GetProbeSplitStageName(),
				},
			},
		},
	})

	splitStage.Commit()

	log.Println("Server ready serve on localhost:" + strconv.Itoa(8080))
	err = stack.R.Run(":" + strconv.Itoa(8080))
	if err != nil {
		t.Errorf("failed to run server: %v", err)
	}
}
