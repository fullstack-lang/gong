// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"github.com/gin-gonic/gin"

	split "github.com/fullstack-lang/gong/lib/split/go/models"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

type Stager struct {
	stage      *Stage
	splitStage *split.Stage
	treeStage  *tree.Stage
	svgStage   *svg.Stage

	sidebarTree *tree.Tree
}

func NewStager(
	r *gin.Engine,
	stage *Stage,
	splitStage *split.Stage,
	treeStage *tree.Stage,
	svgStage *svg.Stage,
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

	split.StageBranch(stager.splitStage, &split.View{
		Name: "Tree Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: (&split.Split{
					StackName: stager.treeStage.GetProbeSplitStageName(),
				}),
			},
		},
	})

	stager.UpdateAndCommitTreeStage()
	stager.splitStage.Commit()

	return
}
