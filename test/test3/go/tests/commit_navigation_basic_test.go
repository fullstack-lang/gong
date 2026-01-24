package tests

import (
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/fullstack-lang/gong/test/test3/go/level1stack"
	"github.com/fullstack-lang/gong/test/test3/go/models"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

func TestBasicCommitNavigation(t *testing.T) {

	// setup
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1Stack("test3", "", "", true, true)
	stage := stack.Stage

	stage.SetDeltaMode(true)

	a0 := (&models.A{Name: "A0"}).Stage(stage)
	a0.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	a0.FloatValue = 14.5
	a0.IntValue = 100
	a0.Duration = -446400000
	a0.EnumString = models.EnumTypeString_Value1
	a0.EnumInt = models.EnumTypeInt_Value2
	_ = a0

	b0 := (&models.B{Name: `B0`}).Stage(stage)
	stage.Commit()

	a0.IntValue = 200
	a0.B = b0

	stage.Commit()
	stack.Probe.Refresh()

	// 	err := stage.ApplyBackwardCommit()
	// if err != nil {
	// 	t.Errorf("failed to apply backward commit: %v", err)
	// }
	// stack.Probe.Refresh()

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
