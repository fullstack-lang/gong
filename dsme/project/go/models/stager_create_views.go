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
		Name:      "All",
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
}
