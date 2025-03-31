// generated boilerplate code
// edit the file for adding other stages
package main

import (
	"github.com/gin-gonic/gin"

	// insertion point for models import
	probe_models "github.com/fullstack-lang/gong/lib/probe/go/models"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

type Stager struct {
	stage      *probe_models.Stage
	splitStage *split.Stage // root
}

func NewStager(r *gin.Engine, stage *probe_models.Stage, probe *probe_models.Probe2) (stager *Stager) {

	stager = new(Stager)

	stager.stage = stage

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	stack := split_stack.NewStack(r, "", "", "", "", false, true)
	stager.splitStage = stack.Stage

	(&split.View{
		Name: "Probe 2",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Size: 100,
				Split: (&split.Split{
					StackName: stage.GetName(),
				}).Stage(stager.splitStage),
			}).Stage(stager.splitStage),
		},
	}).Stage(stager.splitStage)

	(&split.View{
		Name: "Probe of Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Size: 100,
				Split: (&split.Split{
					StackName: stage.GetProbeSplitStageName(),
				}).Stage(stager.splitStage),
			}).Stage(stager.splitStage),
		},
	}).Stage(stager.splitStage)

	(&split.View{
		Name: "Probe of Split stage of the Probe 2",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Size: 100,
				Split: (&split.Split{
					StackName: probe.GetSplitStage().GetProbeSplitStageName(),
				}).Stage(stager.splitStage),
			}).Stage(stager.splitStage),
		},
	}).Stage(stager.splitStage)

	stager.splitStage.Commit()
	stack.Probe.Refresh()

	return
}
