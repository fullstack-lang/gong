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
	probeForm  ProbeIF
}

func NewStager(
	r *http.ServeMux,
	stage *Stage,
	probeForm ProbeIF,
) (stager *Stager) {

	stager = new(Stager)

	stager.stage = stage
	stager.probeForm = probeForm

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	stager.splitStage = split_stack.NewStack(r, "", "", "", "", false, false).Stage

	split.StageBranch(stager.splitStage, &split.View{
		Name: "Data Probe & Data Model",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stage.GetProbeSplitStageName(),
				},
			},
		},
	})

	stager.splitStage.Commit()

	callbacks := &BeforeCommitImplementation{
		stager: stager,
	}
	stager.stage.OnInitCommitFromBackCallback = callbacks
	callbacks.BeforeCommit(stage)

	return
}

type BeforeCommitImplementation struct {
	stager *Stager
}

func (c *BeforeCommitImplementation) BeforeCommit(stage *Stage) {

}
