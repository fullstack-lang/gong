package models

import split "github.com/fullstack-lang/gong/lib/split/go/models"

func (stager *Stager) createViews() {
	split.StageBranch(stager.splitStage, &split.View{
		Name: "tree & diagram",

		RootAsSplitAreas: []*split.AsSplitArea{
			{
				AsSplit: &split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Size: 20,
							AsSplit: &split.AsSplit{
								Direction: split.Vertical,
								AsSplitAreas: []*split.AsSplitArea{
									{
										Size: 90,
										Tree: (&split.Tree{
											StackName: stager.treeStage.GetName(),
										}),
									},
									{
										Size: 10,
										Button: &split.Button{
											StackName: stager.buttonStage.GetName(),
										},
									},
									{
										Size: 0,
										Load: &split.Load{
											StackName: stager.loadStage.GetName(),
										},
									},
								},
							},
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
		Name: "Load / Download / Buttons",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "Load  Buttons Download",
				ShowNameInHeader: false,
				Size:             35,
				AsSplit: &split.AsSplit{
					Name:      "as split",
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Size: 50,
							Load: &split.Load{
								StackName: stager.loadStage.GetName(),
							},
						},
						{
							Size: 50,
							Button: &split.Button{
								StackName: stager.buttonStage.GetName(),
							},
						},
					},
				},
			},
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
}
