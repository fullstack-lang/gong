// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"github.com/gin-gonic/gin"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	svg_stack "github.com/fullstack-lang/gong/lib/svg/go/stack"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	tree_stack "github.com/fullstack-lang/gong/lib/tree/go/stack"
)

type Stager struct {
	stage       *Stage
	splitStage  *split.Stage
	asSplitArea *split.AsSplitArea

	svgStage  *svg.Stage
	treeStage *tree.Stage

	desk *Desk
}

func NewStager(
	r *gin.Engine,
	stage *Stage,
	probeForm ProbeIF,
) (stager *Stager) {
	stager = new(Stager)

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	stager.splitStage = split_stack.NewStack(r, "", "", "", "", false, false).Stage

	stager.stage = stage

	stager.svgStage = svg_stack.NewStack(r, stage.GetName(), "", "", "", true, true).Stage
	stager.treeStage = tree_stack.NewStack(r, stage.GetName(), "", "", "", true, true).Stage

	split.StageBranch(stager.splitStage, &split.View{
		Name: "tree & diagram",

		RootAsSplitAreas: []*split.AsSplitArea{
			{
				AsSplit: &split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Size: 20,
							Tree: (&split.Tree{
								StackName: stager.treeStage.GetName(),
							}),
						},
						{
							Size: 40,
							Svg: (&split.Svg{
								StackName: stager.svgStage.GetName(),
							}),
						},
						{
							Size: 40,
							Split: (&split.Split{
								StackName: stager.stage.GetProbeSplitStageName(),
							}),
						},
					},
				},
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.stage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "svg probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.svgStage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	stager.splitStage.Commit()
	stage.OnInitCommitFromBackCallback = &BeforeCommitImplementation{stager: stager}
	stage.Commit()

	return
}

func (stager *Stager) GetAsSplitArea() (asSplitArea *split.AsSplitArea) {
	asSplitArea = stager.asSplitArea
	return
}

func (stager *Stager) GetGongtreeStage() *tree.Stage {
	return stager.treeStage
}

type BeforeCommitImplementation struct {
	stager *Stager
}

func (c *BeforeCommitImplementation) BeforeCommit(stage *Stage) {
	c.stager.ComputeConsistency()
	c.stager.SvgStageUpdate()
	c.stager.TreeStageUpdate()
}
