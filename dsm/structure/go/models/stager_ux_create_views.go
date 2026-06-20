package models

import split "github.com/fullstack-lang/gong/lib/split/go/models"

func (stager *Stager) createViews() {
	stage := stager.stage
	split.StageBranch(stager.splitStage, &split.View{
		Name:           "Edit Structure",
		Direction:      split.Horizontal,
		IsSelectedView: true,
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "Sidebar with tree",
				ShowNameInHeader: false,
				Size:             80,
				AsSplit: &split.AsSplit{
					Name:      "as split",
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Size: 30,
							AsSplit: &split.AsSplit{
								Direction: split.Vertical,
								AsSplitAreas: []*split.AsSplitArea{
									{
										Size:             80,
										ShowNameInHeader: false,
										Tree: &split.Tree{
											StackName: stager.treeStage.GetName(),
										},
									},
									{
										Size: 10,
										Load: &split.Load{
											StackName: stager.loadStage.GetName(),
										},
									},
									{
										Size: 10,
										Button: &split.Button{
											StackName: stager.buttonStage.GetName(),
										},
									},
								},
							},
						},
						{
							Size: 70,
							Svg: &split.Svg{
								StackName: stager.svgStage.GetName(),
							},
						},
					},
				},
			},
			{
				Size: 20,
				Form: &split.Form{
					StackName: stager.probeForm.GetFormStage().GetName(),
				},
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "Data Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stage.GetProbeSplitStageName(),
				},
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name:            "Tree Product Probe",
		IsSecondaryView: true,
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stager.treeStage.GetProbeSplitStageName(),
				},
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name:            "Svg Probe",
		IsSecondaryView: true,
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stager.svgStage.GetProbeSplitStageName(),
				},
			},
		},
	})

	stager.splitStage.Commit()
}
