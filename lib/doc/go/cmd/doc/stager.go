// generated boilerplate code
// edit the file for adding other stages
package main

import (
	"github.com/gin-gonic/gin"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"

	doc_models "github.com/fullstack-lang/gong/lib/doc/go/models"
)

// Stager is usualy in the models package.
//
// for doc, it is not possible because the doc models package imports the split stack package
// which imports the doc models package (not allowed)
//
// imports github.com/fullstack-lang/gong/lib/doc/go/models from main.go
// imports github.com/fullstack-lang/gong/lib/split/go/stack from stager.go
// imports github.com/fullstack-lang/gong/lib/split/go/probe from stack.go
// imports github.com/fullstack-lang/gong/lib/doc/go/load from probe.go
// imports github.com/fullstack-lang/gong/lib/doc/go/adapter from load.go
// imports github.com/fullstack-lang/gong/lib/doc/go/doc2svg from class_diagram_node.go
// imports github.com/fullstack-lang/gong/lib/doc/go/models from anchored_text_impl_link_field_name.go: import cycle not allowed
type Stager struct {
	stage      *doc_models.Stage
	splitStage *split.Stage
}

func NewStager(r *gin.Engine, stage *doc_models.Stage) (stager *Stager) {

	stager = new(Stager)

	stager.stage = stage

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	stager.splitStage = split_stack.NewStack(r, "", "", "", "", false, false).Stage

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
