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

	a1 := (&models.A{Name: "A1"}).Stage(stage)
	a1.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	a1.FloatValue = 14.5
	a1.IntValue = 100
	a1.Duration = -446400000
	a1.EnumString = models.EnumTypeString_Value1
	a1.EnumInt = models.EnumTypeInt_Value2

	b0 := (&models.B{Name: `B0`}).Stage(stage)
	_ = b0
	b1 := (&models.B{Name: `B1`}).Stage(stage)
	_ = b1
	stage.Commit()

	a1.IntValue = 150
	// a0.IntValue = 200
	// a0.B = b0
	// a0.EnumString = ""
	// a0.Bs = append(a0.Bs, b0)
	// a0.Bs = append(a0.Bs, b1)

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
