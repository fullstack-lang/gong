// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"github.com/gin-gonic/gin"

	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"

	tree_stack "github.com/fullstack-lang/gong/lib/tree/go/stack"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	svg_stack "github.com/fullstack-lang/gong/lib/svg/go/stack"
)

type Stager struct {
	stage      *Stage
	splitStage *split.Stage
	probeForm  ProbeIF

	root *Root

	treeProductsStage *tree.Stage
	treeTasksStage    *tree.Stage

	svgStage *svg.Stage
}

func NewStager(
	r *gin.Engine,
	stage *Stage,
	probeForm ProbeIF,
) (stager *Stager) {

	stager = new(Stager)

	stager.stage = stage
	stager.probeForm = probeForm

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	stager.splitStage = split_stack.NewStack(r, "", "", "", "", false, false).Stage
	stager.treeProductsStage = tree_stack.NewStack(r, "products", "", "", "", true, true).Stage
	stager.treeTasksStage = tree_stack.NewStack(r, "tasks", "", "", "", true, true).Stage

	stager.svgStage = svg_stack.NewStack(r, "", "", "", "", true, true).Stage

	split.StageBranch(stager.splitStage, &split.View{
		Name:      "Data Probe & Data Model",
		Direction: split.Horizontal,
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "Sidebar with both trees",
				ShowNameInHeader: false,
				Size:             62,
				AsSplit: &split.AsSplit{
					Name:      "as split",
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Size: 38,
							AsSplit: &split.AsSplit{
								Direction: split.Vertical,
								AsSplitAreas: []*split.AsSplitArea{
									{
										Name:             "Bottom",
										Size:             50,
										ShowNameInHeader: false,
										Tree: &split.Tree{
											StackName: stager.treeProductsStage.GetName(),
										},
									},
									{
										Name:             "Top",
										Size:             50,
										ShowNameInHeader: false,
										Tree: &split.Tree{
											StackName: stager.treeTasksStage.GetName(),
										},
									},
								},
							},
						},
						{
							Size: 62,
							Svg: &split.Svg{
								StackName: stager.svgStage.GetName(),
							},
						},
					},
				},
			},
			{
				Size: 38,
				Split: &split.Split{
					StackName: stage.GetProbeSplitStageName(),
				},
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "Tree Product Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stager.treeProductsStage.GetProbeSplitStageName(),
				},
			},
		},
	})

	stager.splitStage.Commit()

	callbacks := &BeforeCommitImplementation{
		stager: stager,
	}
	stager.stage.OnInitCommitFromBackCallback = callbacks
	callbacks.BeforeCommit(stage)

	return
}
