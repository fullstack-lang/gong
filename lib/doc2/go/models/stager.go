// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"github.com/gin-gonic/gin"

	split "github.com/fullstack-lang/gong/lib/split/go/models"

	tree_models "github.com/fullstack-lang/gong/lib/tree/go/models"

	svg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

type Stager struct {
	stage      *Stage
	splitStage *split.Stage
	treeStage  *tree_models.Stage
	svgStage   *svg_models.Stage
}

func NewStager(
	r *gin.Engine,
	stage *Stage,
	splitStage *split.Stage,
	treeStage *tree_models.Stage,
	svgStage *svg_models.Stage,
) (stager *Stager) {

	stager = new(Stager)

	stager.stage = stage
	stager.splitStage = splitStage
	stager.treeStage = treeStage
	stager.svgStage = svgStage

	// StageBranch will stage on the the first argument
	// all instances related to the second argument
	split.StageBranch(stager.splitStage, &split.View{
		Name: "Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Size: 50,
				AsSplit: (&split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Size: 25,
							Tree: &split.Tree{
								StackName: stager.treeStage.GetName(),
							},
						},
						{
							Size: 75,
							Svg: &split.Svg{
								StackName: stager.svgStage.GetName(),
							},
						},
					},
				}),
			},
			{
				Size: 50,
				Split: (&split.Split{
					StackName: stage.GetProbeSplitStageName(),
				}),
			},
		},
	})

	stager.splitStage.Commit()

	return
}
