// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"log"

	"github.com/gin-gonic/gin"

	split_fullstack "github.com/fullstack-lang/gong/lib/split/go/fullstack"
	split "github.com/fullstack-lang/gong/lib/split/go/models"

	button_fullstack "github.com/fullstack-lang/gong/lib/button/go/fullstack"
	button "github.com/fullstack-lang/gong/lib/button/go/models"
)

type Stager struct {
	stage      *Stage
	splitStage *split.Stage

	buttonStage *button.Stage
}

// GetButtonsStage implements models.Target.
func (stager *Stager) GetButtonsStage() *button.Stage {
	return stager.buttonStage
}

// OnAfterUpdateButton implements models.Target.
func (stager *Stager) OnAfterUpdateButton() {
	log.Println("button updated")
	stager.stage.Generation()
}

func NewStager(r *gin.Engine, stage *Stage) (stager *Stager) {

	stager = new(Stager)

	stager.stage = stage

	name := stage.GetName()

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	stager.splitStage, _ = split_fullstack.NewStackInstance(r, "")

	stager.buttonStage, _ = button_fullstack.NewStackInstance(r, name)

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
