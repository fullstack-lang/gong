package tests

import (
	"log"
	"strconv"
	"testing"

	"github.com/fullstack-lang/gong/test/test3/go/level1stack"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

func TestCommitNavigation(t *testing.T) {

	// setup
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1StackDelta("test3", "stage.go", "", true, true, true)
	stage := stack.Stage
	stack.Probe.Refresh()

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
	err := stack.R.Run(":" + strconv.Itoa(8080))
	if err != nil {
		t.Errorf("failed to run server: %v", err)
	}
}
