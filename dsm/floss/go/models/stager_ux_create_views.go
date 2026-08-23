package models

import split "github.com/fullstack-lang/gong/lib/split/go/models"

func (stager *Stager) createViews() {
	stager.splitStage.Reset()

	var equationChecked bool
	for d := range *GetGongstructInstancesSet[DiagramFlossEquation](stager.stage) {
		if d.IsChecked {
			equationChecked = true
			break
		}
	}

	// View 1: System Diagram (Tree, SVG, Form ONLY - no sliders)
	split.StageBranch(stager.splitStage, &split.View{
		Name:           "System Diagram",
		Direction:      split.Horizontal,
		IsSelectedView: !equationChecked,
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "Sidebar with tree and SVG",
				ShowNameInHeader: false,
				Size:             75,
				AsSplit: &split.AsSplit{
					Name:      "as split",
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Size: 32,
							AsSplit: &split.AsSplit{
								Direction: split.Vertical,
								AsSplitAreas: []*split.AsSplitArea{
									{
										Size:             86,
										ShowNameInHeader: false,
										Tree: &split.Tree{
											StackName: stager.treeStage.GetName(),
										},
									},
									{
										Size: 7,
										Load: &split.Load{
											StackName: stager.loadStage.GetName(),
										},
									},
									{
										Size: 7,
										Button: &split.Button{
											StackName: stager.buttonStage.GetName(),
										},
									},
								},
							},
						},
						{
							Size: 68,
							Svg: &split.Svg{
								StackName: stager.systemDiagramSvgStage.GetName(),
							},
						},
					},
				},
			},
			{
				Size: 25,
				AsSplit: &split.AsSplit{
					Direction: split.Vertical,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Size: 55,
							Slider: &split.Slider{
								StackName: stager.sliderStage.GetName(),
							},
						},
						{
							Size: 45,
							Form: &split.Form{
								StackName: stager.probeForm.GetFormStage().GetName(),
							},
						},
					},
				},
			},
		},
	})

	// View 4: Probe
	split.StageBranch(stager.splitStage, &split.View{
		Name: "Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stager.stage.GetProbeSplitStageName(),
				},
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name:            "Tree probe",
		IsSecondaryView: true,
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.treeStage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	stager.splitStage.Commit()
}
