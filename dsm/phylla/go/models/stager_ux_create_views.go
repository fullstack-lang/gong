package models

import split "github.com/fullstack-lang/gong/lib/split/go/models"

func getPersistanceFile(stager *Stager) string {
	if stager.stage.OnInitCommitCallback != nil {
		return stager.persistanceFile
	} else {
		return "no persistance"
	}
}

func (stager *Stager) createViews() {
	stager.splitStage.Reset()

	tabTitle := &split.Title{
		Name: "Project (" + getPersistanceFile(stager) + ")",
	}
	tabTitle.Stage(stager.splitStage)

	split.StageBranch(stager.splitStage, &split.View{
		Name:           "Edit PBS/WBS (" + getPersistanceFile(stager) + ")",
		Direction:      split.Horizontal,
		IsSizeInPixel:  true,
		IsSelectedView: true,
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "Sidebar with both trees",
				ShowNameInHeader: false,
				IsAny:            true,
				AsSplit: &split.AsSplit{
					Name:          "as split",
					IsSizeInPixel: true,
					Direction:     split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Size: 525,
							AsSplit: &split.AsSplit{
								Direction: split.Vertical,
								AsSplitAreas: []*split.AsSplitArea{
									{
										Name:             "Libraries",
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
							IsAny: true,
							Svg: &split.Svg{
								StackName: stager.svgStage.GetName(),
							},
						},
					},
				},
			},
			{
				Size: 525,
				Form: &split.Form{
					StackName: stager.probeForm.GetFormStage().GetName(),
				},
			},
		},
	})

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
		Name:            "Tree Probe",
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

	split.StageBranch(stager.splitStage, &split.View{
		Name:            "ssg Probe",
		IsSecondaryView: true,
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stager.ssgStage.GetProbeSplitStageName(),
				},
			},
		},
	})

	stager.splitStage.Commit()
}
