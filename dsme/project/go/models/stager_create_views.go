package models

import split "github.com/fullstack-lang/gong/lib/split/go/models"

func createViews(stager *Stager, stage *Stage) {

	split.StageBranch(stager.splitStage, &split.View{
		Name:      "Edit PBS/WBS",
		Direction: split.Horizontal,
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "Sidebar with both trees",
				ShowNameInHeader: false,
				Size:             82,
				AsSplit: &split.AsSplit{
					Name:      "as split",
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Size: 26,
							AsSplit: &split.AsSplit{
								Direction: split.Vertical,
								AsSplitAreas: []*split.AsSplitArea{
									{
										Name:             "Projects",
										Size:             100,
										ShowNameInHeader: false,
										Tree: &split.Tree{
											StackName: stager.treeStage.GetName(),
										},
									},
								},
							},
						},
						{
							Size: 74,
							Svg: &split.Svg{
								StackName: stager.svgStage.GetName(),
							},
						},
					},
				},
			},
			{
				Size: 18,
				Form: &split.Form{
					StackName: stager.probeForm.GetFormStage().GetName(),
				},
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name:      "Edit PBS/WBS",
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
										Name:             "Projects",
										Size:             100,
										ShowNameInHeader: false,
										Tree: &split.Tree{
											StackName: stager.treeStage.GetName(),
										},
									},
								},
							},
						},
						{
							Size: 62,
							Form: &split.Form{
								StackName: stager.probeForm.GetFormStage().GetName(),
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
		Name:           "All",
		Direction:      split.Horizontal,
		IsSelectedView: true,
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "Sidebar with both trees",
				ShowNameInHeader: false,
				Size:             35,
				AsSplit: &split.AsSplit{
					Name:      "as split",
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Size: 50,
							AsSplit: &split.AsSplit{
								Direction: split.Vertical,
								AsSplitAreas: []*split.AsSplitArea{
									{
										Name:             "Bottom",
										Size:             100,
										ShowNameInHeader: false,
										Tree: &split.Tree{
											StackName: stager.treeStage.GetName(),
										},
									},
								},
							},
						},
						{
							Size: 50,
							Svg: &split.Svg{
								StackName: stager.svgStage.GetName(),
							},
						},
					},
				},
			},
			{
				Size: 65,
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
					StackName: stager.treeStage.GetProbeSplitStageName(),
				},
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "Svg Probe",
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
