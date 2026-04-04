// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"log"

	"github.com/gin-gonic/gin"

	split_fullstack "github.com/fullstack-lang/gong/lib/split/go/fullstack"
	split "github.com/fullstack-lang/gong/lib/split/go/models"
	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"

	button_fullstack "github.com/fullstack-lang/gong/lib/button/go/fullstack"
	button "github.com/fullstack-lang/gong/lib/button/go/models"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
	load_stack "github.com/fullstack-lang/gong/lib/load/go/stack"
)

type Stager struct {
	stage      *Stage
	splitStage *split.Stage

	buttonStage *button.Stage
	loadStage   *load.Stage
}

func NewStager(r *gin.Engine, stage *Stage) (stager *Stager) {

	stager = new(Stager)

	stager.stage = stage

	name := stage.GetName()

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	stager.splitStage, _ = split_fullstack.NewStackInstance(r, "")

	stager.buttonStage, _ = button_fullstack.NewStackInstance(r, name)
	stager.loadStage = load_stack.NewStack(r, "", "", "", "", true, true).Stage

	// StageBranch will stage on the the first argument
	// all instances related to the second argument
	split.StageBranch(stager.splitStage,
		&split.View{
			Name: "Main view",
			RootAsSplitAreas: []*split.AsSplitArea{
				(&split.AsSplitArea{
					Size:             6,
					ShowNameInHeader: false,
					Button: (&split.Button{
						StackName: stager.buttonStage.GetName(),
					}),
				}),
				(&split.AsSplitArea{
					Size:             94,
					ShowNameInHeader: false,
					Split: (&split.Split{
						StackName: stage.GetProbeSplitStageName(),
					}),
				}),
			},
		})

	stager.UpdateAndCommitButtonStage()
	stager.splitStage.Commit()

	return
}

func (stager *Stager) UpdateAndCommitButtonStage() {
	stager.buttonStage.Reset()

	layout := new(button.Layout).Stage(stager.buttonStage)

	group1 := new(button.Group).Stage(stager.buttonStage)
	group1.Percentage = 100
	layout.Groups = append(layout.Groups, group1)

	btn := (&button.Button{
		Name:  "Generate static site",
		Icon:  string(buttons.BUTTON_web_asset),
		Label: "Generate static site",
		OnUpdate: func() {
			log.Println("button updated")
			stager.stage.Generation()
		},
	}).Stage(stager.buttonStage)

	group1.Buttons = append(group1.Buttons, btn)

	stager.buttonStage.Commit()
}
