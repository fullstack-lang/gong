// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"fmt"

	"github.com/gin-gonic/gin"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
)

type Stager struct {
	stage       *Stage
	splitStage  *split.Stage
	asSplitArea *split.AsSplitArea
}

func NewStager(r *gin.Engine, stage *Stage, splitStage *split.Stage) (stager *Stager) {

	stager = new(Stager)

	stager.stage = stage
	stager.splitStage = splitStage

	stager.asSplitArea = &split.AsSplitArea{}

	// one for the probe of the
	split.StageBranch(splitStage, &split.View{
		Name: stage.GetName() + "with Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Size: 50,
				AsSplit: (&split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						stager.GetAsSplitArea(),
					},
				}),
			}),
			(&split.AsSplitArea{
				Size: 50,
				Split: (&split.Split{
					StackName: stage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	// commit the split stage (this will initiate the front components)
	splitStage.Commit()

	// While there are less than 5000 instances of Bs
	nbBs := len(*GetGongstructInstancesSetFromPointerType[*B](stage))
	for nbBs < 5000 {
		(&B{Name: fmt.Sprintf("%d", nbBs)}).Stage(stage)
		nbBs = len(*GetGongstructInstancesSetFromPointerType[*B](stage))
	}

	return
}

func (stager *Stager) GetAsSplitArea() (asSplitArea *split.AsSplitArea) {
	asSplitArea = stager.asSplitArea
	return
}
