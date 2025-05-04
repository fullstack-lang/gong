// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	split "github.com/fullstack-lang/gong/lib/split/go/models"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"

	gong "github.com/fullstack-lang/gong/go/models"
)

type Stager struct {
	stage     *Stage
	treeStage *tree.Stage
	svgStage  *svg.Stage
	gongStage *gong.Stage

	sidebarTree *tree.Tree

	embeddedDiagrams bool
}

func NewStager(
	r *gin.Engine,
	receivingAsSplitArea *split.AsSplitArea,
	stage *Stage,
	treeStage *tree.Stage,
	svgStage *svg.Stage,
	gongStage *gong.Stage,

	embeddedDiagrams bool,
) (stager *Stager) {

	stager = new(Stager)

	stager.stage = stage
	stager.treeStage = treeStage
	stager.svgStage = svgStage
	stager.gongStage = gongStage

	stager.embeddedDiagrams = embeddedDiagrams

	// StageBranch will stage on the the first argument
	// all instances related to the second argument
	receivingAsSplitArea.AsSplit = &split.AsSplit{
		Direction: split.Horizontal,
		AsSplitAreas: []*split.AsSplitArea{
			{
				Size: 50,
				AsSplit: (&split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Size: 25,
							Tree: &split.Tree{
								StackName: stager.treeStage.GetName(),
								TreeName:  stager.stage.GetProbeTreeSidebarStageName(),
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
	}

	// if no diagram package is present, creates one
	diagramPackages := *GetGongstructInstancesSet[DiagramPackage](stage)
	var diagramPackage *DiagramPackage
	for k, _ := range diagramPackages {
		diagramPackage = k
	}
	if diagramPackage == nil {
		diagramPackage = (&DiagramPackage{
			Name: fmt.Sprintf("Diagram Package created the %s", time.Now().Local().UTC().Format(time.RFC3339)),
		}).Stage(stage)
		stage.Commit()
	}

	stager.UpdateAndCommitSVGStage()
	stager.UpdateAndCommitTreeStage()

	return
}
