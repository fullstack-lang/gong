// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"github.com/gin-gonic/gin"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

type Stager struct {
	stage      *Stage
	splitStage *split.Stage
}

func NewStager(r *gin.Engine, stage *Stage) (stager *Stager) {

	stager = new(Stager)

	stager.stage = stage

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	stager.splitStage = split_stack.NewStack(r, "", "", "", "", false, false).Stage

	(&split.View{
		Name: "Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Size: 20,
				Button: (&split.Button{
					StackName: stage.GetName(),
				}).Stage(stager.splitStage),
			}).Stage(stager.splitStage),
			(&split.AsSplitArea{
				Size: 80,
				Split: (&split.Split{
					StackName: stage.GetProbeSplitStageName(),
				}).Stage(stager.splitStage),
			}).Stage(stager.splitStage),
		},
	}).Stage(stager.splitStage)

	stager.splitStage.Commit()

	return
}
