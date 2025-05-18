// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"github.com/gin-gonic/gin"

	split_fullstack "github.com/fullstack-lang/gong/lib/split/go/fullstack"
	split "github.com/fullstack-lang/gong/lib/split/go/models"
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
	stager.splitStage, _ = split_fullstack.NewStackInstance(r, "")

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
