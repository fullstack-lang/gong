// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"github.com/gin-gonic/gin"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	svg_stack "github.com/fullstack-lang/gong/lib/svg/go/stack"
)

type Stager struct {
	stage      *Stage
	splitStage *split.Stage
	svgStage   *svg.Stage
}

func NewStager(r *gin.Engine, stage *Stage, name string) (stager *Stager) {

	stager = new(Stager)

	stager.stage = stage

	svgPersistanceFile := "svg.go"
	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	stager.splitStage = split_stack.NewStack(r, "", "", "", "", false, false).Stage
	stager.svgStage = svg_stack.NewStack(r, name, svgPersistanceFile, svgPersistanceFile, "", true, true).Stage

	split.StageBranch(stager.splitStage, &split.View{
		Name: "Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Size: 50,
				Svg: (&split.Svg{
					StackName: stager.svgStage.GetName(),
				}),
			}),
			(&split.AsSplitArea{
				Size: 25,
				Split: (&split.Split{
					StackName: stage.GetProbeSplitStageName(),
				}),
			}),
			(&split.AsSplitArea{
				Size: 25,
				Split: (&split.Split{
					StackName: stager.svgStage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.svgStage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	stager.splitStage.Commit()

	return
}
