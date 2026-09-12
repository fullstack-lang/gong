// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"net/http"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

type Stager struct {
	stage      *Stage
	splitStage *split.Stage
}

func NewStager(r *http.ServeMux, stage *Stage) (stager *Stager) {

	stager = new(Stager)

	stager.stage = stage
	stager.splitStage = split_stack.NewStack(r, "", "", "", "", false, false).Stage

	(&split.View{
		Name: "Tone",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Tone: (&split.Tone{
					StackName: stage.GetName(),
				}).Stage(stager.splitStage),
			}).Stage(stager.splitStage),
		},
	}).Stage(stager.splitStage)

	(&split.View{
		Name: "Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stage.GetProbeSplitStageName(),
				}).Stage(stager.splitStage),
			}).Stage(stager.splitStage),
		},
	}).Stage(stager.splitStage)

	stager.splitStage.Commit()

	return
}
