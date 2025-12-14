// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"os"

	"github.com/gin-gonic/gin"

	button_stack "github.com/fullstack-lang/gong/lib/button/go/stack"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
	svg_stack "github.com/fullstack-lang/gong/lib/svg/go/stack"
	tree_stack "github.com/fullstack-lang/gong/lib/tree/go/stack"

	// tmp
	_ "github.com/fullstack-lang/gong/lib/cursor/go/stack"
	_ "github.com/fullstack-lang/gong/lib/load/go/stack"
	_ "github.com/fullstack-lang/gong/lib/slider/go/stack"
	_ "github.com/fullstack-lang/gong/lib/tone/go/stack"

	button "github.com/fullstack-lang/gong/lib/button/go/models"
	split "github.com/fullstack-lang/gong/lib/split/go/models"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type Stager struct {
	stage      *Stage
	splitStage *split.Stage

	treeObjectsStage       *tree.Stage
	treeDiagramStage       *tree.Stage
	svgStage               *svg.Stage
	buttonTransitionsStage *button.Stage
	buttonExportXLStage    *button.Stage

	// maps

	// singloton architecture
	architecture *Architecture

	set_StartStates map[*State]struct{}

	map_state_nextStates map[*State][]*State

	map_stateMachine_objects map[*StateMachine][]*Object
	map_state_stateMachine   map[*State]*StateMachine
	map_diagram_stateMachine map[*Diagram]*StateMachine

	probeForm ProbeIF
}

func (stager *Stager) GetStage() *Stage {
	return stager.stage
}

// OnAfterUpdateButton implements models.Target.
func (stager *Stager) OnAfterUpdateButton() {

}

func (stager *Stager) GetButtonsStage() *button.Stage {
	return stager.buttonTransitionsStage
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

	stackName := "statemachines"
	stager.treeObjectsStage = tree_stack.NewStack(r, stackName+"-objects", "", "", "", true, true).Stage
	stager.treeDiagramStage = tree_stack.NewStack(r, stackName+"-diagrams", "", "", "", true, true).Stage
	stager.svgStage = svg_stack.NewStack(r, stackName, "", "", "", true, true).Stage
	stager.buttonTransitionsStage = button_stack.NewStack(r, stackName+"-transitions", "", "", "", false, false).Stage
	stager.buttonExportXLStage = button_stack.NewStack(r, stackName+"-exportXL", "", "", "", true, true).Stage

	split.StageBranch(stager.splitStage, &split.View{
		Name:           "Edit view",
		IsSelectedView: true,
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "Main view area",
				Size:             100,
				ShowNameInHeader: false,
				AsSplit: &split.AsSplit{
					Direction: split.Vertical,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Name:             "Main view area",
							Size:             100,
							ShowNameInHeader: false,
							AsSplit: &split.AsSplit{
								Direction: split.Horizontal,
								AsSplitAreas: []*split.AsSplitArea{
									// Sidebar
									{
										Name:             "Sidebar with both trees",
										ShowNameInHeader: false,
										Size:             25,
										AsSplit: &split.AsSplit{
											Name:      "as split",
											Direction: split.Vertical,
											AsSplitAreas: []*split.AsSplitArea{
												{
													Name:             "Top",
													Size:             100,
													ShowNameInHeader: false,

													AsSplit: &split.AsSplit{
														Direction: split.Vertical,
														AsSplitAreas: []*split.AsSplitArea{
															{
																Name:             "Model tree",
																Size:             90,
																ShowNameInHeader: false,
																Tree: &split.Tree{
																	StackName: stager.treeDiagramStage.GetName(),
																},
															},

															//
															{
																Name:             "Export XL button",
																Size:             10,
																ShowNameInHeader: false,
																Button: &split.Button{
																	StackName: stager.buttonExportXLStage.GetName(),
																},
															},
															//
														}},
												},
											},
										},
									},

									// SVG area
									{
										Name:             "svg diagram",
										ShowNameInHeader: false,
										Size:             55,
										Svg: &split.Svg{
											StackName: stager.svgStage.GetName(),
										},
									},
									{
										Name:             "form",
										ShowNameInHeader: false,
										Size:             20,
										Form: &split.Form{
											StackName: stager.probeForm.GetFormStage().GetName(),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "Model view",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "Main view area",
				Size:             100,
				ShowNameInHeader: false,
				AsSplit: &split.AsSplit{
					Direction: split.Vertical,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Name:             "Main view area",
							Size:             50,
							ShowNameInHeader: false,
							AsSplit: &split.AsSplit{
								Direction: split.Horizontal,
								AsSplitAreas: []*split.AsSplitArea{
									// Sidebar
									{
										Name:             "Sidebar with both trees",
										ShowNameInHeader: false,
										Size:             25,
										AsSplit: &split.AsSplit{
											Name:      "as split",
											Direction: split.Vertical,
											AsSplitAreas: []*split.AsSplitArea{
												{
													Name:             "Top",
													Size:             100,
													ShowNameInHeader: false,

													AsSplit: &split.AsSplit{
														Direction: split.Vertical,
														AsSplitAreas: []*split.AsSplitArea{
															{
																Name:             "Model tree",
																Size:             90,
																ShowNameInHeader: false,
																Tree: &split.Tree{
																	StackName: stager.treeDiagramStage.GetName(),
																},
															},

															//
															{
																Name:             "Export XL button",
																Size:             10,
																ShowNameInHeader: false,
																Button: &split.Button{
																	StackName: stager.buttonExportXLStage.GetName(),
																},
															},
															//
														}},
												},
											},
										},
									},

									// SVG area
									{
										Name:             "svg diagram",
										ShowNameInHeader: false,
										Size:             75,
										Svg: &split.Svg{
											StackName: stager.svgStage.GetName(),
										},
									},
								},
							},
						},
						{
							Size:    50,
							AsSplit: probeForm.GetDataEditor(),
						},
					},
				},
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "Main view",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "Main view area",
				Size:             100,
				ShowNameInHeader: false,
				AsSplit: &split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Name:             "Sidebar with both trees",
							ShowNameInHeader: false,
							Size:             25,
							AsSplit: &split.AsSplit{
								Name:      "as split",
								Direction: split.Vertical,
								AsSplitAreas: []*split.AsSplitArea{
									{
										Name:             "Top",
										Size:             50,
										ShowNameInHeader: false,
										Tree: &split.Tree{
											StackName: stager.treeDiagramStage.GetName(),
										},
									},
									{
										Name:             "Bottom",
										Size:             50,
										ShowNameInHeader: false,
										Tree: &split.Tree{
											StackName: stager.treeObjectsStage.GetName(),
										},
									},
								},
							},
						},
						{
							Name:             "Actions",
							ShowNameInHeader: false,
							Size:             20,
							Button: &split.Button{
								StackName: stager.buttonTransitionsStage.GetName(),
							},
						},
						{
							Name:             "svg diagram",
							ShowNameInHeader: false,
							Size:             55,
							Svg: &split.Svg{
								StackName: stager.svgStage.GetName(),
							},
						},
					},
				},
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "Sim. view",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "Main view area",
				Size:             100,
				ShowNameInHeader: false,
				AsSplit: &split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Name:             "Actions",
							ShowNameInHeader: false,
							Size:             20,
							Button: &split.Button{
								StackName: stager.buttonTransitionsStage.GetName(),
							},
						},
						{
							Name:             "svg diagram",
							ShowNameInHeader: false,
							Size:             80,
							Svg: &split.Svg{
								StackName: stager.svgStage.GetName(),
							},
						},
					},
				},
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "Data Probe & Data Model",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stage.GetProbeSplitStageName(),
				},
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "Data Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				AsSplit: probeForm.GetDataEditor(),
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "Data Model",
		RootAsSplitAreas: []*split.AsSplitArea{
			probeForm.GetDiagramEditor(),
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "svg Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stager.svgStage.GetProbeSplitStageName(),
				},
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "Diagram Tree Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stager.treeDiagramStage.GetProbeSplitStageName(),
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

	// hook the stage on a kill command
	// 	curl --request POST \
	//   --url http://localhost:8000/api/gongstatesmachines/go/v1/kills?Name="gongstatesmachines" \
	//   --header 'content-type: application/json' \
	//   --data '{"Name": "Die !"}'
	stage.OnAfterKillCreateCallback = &OnAfterKillCreateCallback{}

	return
}

type Kill struct {
	Name string
}

type OnAfterKillCreateCallback struct {
}

// OnAfterCreate implements OnAfterCreateInterface.
func (o *OnAfterKillCreateCallback) OnAfterCreate(stage *Stage, instance *Kill) {
	os.Exit(0)
}

type OnInitCommitFromBackCallback struct {
	stager *Stager
}
