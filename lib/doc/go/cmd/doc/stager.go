// generated boilerplate code
// edit the file for adding other stages
package main

import (
	"github.com/gin-gonic/gin"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"

	button "github.com/fullstack-lang/gong/lib/button/go/models"
	button_stack "github.com/fullstack-lang/gong/lib/button/go/stack"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"

	doc_models "github.com/fullstack-lang/gong/lib/doc/go/models"

	probe "github.com/fullstack-lang/gong/lib/doc/go/probe"
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
	stage *doc_models.Stage

	splitStage *split.Stage

	buttonStage *button.Stage

	probe *probe.Probe
}

func NewStager(r *gin.Engine, stage *doc_models.Stage, probe *probe.Probe) (stager *Stager) {

	stager = new(Stager)

	stager.stage = stage
	stager.probe = probe

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	stager.splitStage = split_stack.NewStack(r, "", "", "", "", false, false).Stage
	stager.buttonStage = button_stack.NewStack(r, "", "", "", "", false, false).Stage

	button.StageBranch(stager.buttonStage,
		&button.Layout{
			Groups: []*button.Group{
				&button.Group{
					Buttons: []*button.Button{
						button.NewButton(
							// stager is the target of the button. stager implements interface method OnAfterUpdateButton()
							&ExportAllDiagramsButtonProxy{
								stager: stager,
							},
							"Export All diagrams",
							string(buttons.BUTTON_web),
							"Export All diagrams",
						),
					},
				},
			},
		},
	)

	stager.buttonStage.Commit()

	split.StageBranch(stager.splitStage,
		&split.View{
			Name: "Probe",
			RootAsSplitAreas: []*split.AsSplitArea{
				(&split.AsSplitArea{
					Size: 10,
					Button: (&split.Button{
						StackName: stager.buttonStage.GetName(),
					}),
				}),
				(&split.AsSplitArea{
					Size: 90,
					Split: (&split.Split{
						StackName: stage.GetProbeSplitStageName(),
					}),
				}),
			},
		})

	stager.splitStage.Commit()

	return
}

type ExportAllDiagramsButtonProxy struct {
	stager *Stager
}

// GetButtonsStage implements models.Target.
func (e *ExportAllDiagramsButtonProxy) GetButtonsStage() *button.Stage {
	return e.stager.buttonStage
}

// OnAfterUpdateButton implements models.Target.
func (e *ExportAllDiagramsButtonProxy) OnAfterUpdateButton() {
	e.stager.probe.GeneratesDiagrams("generatedSvg")
}
