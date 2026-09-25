package models

import split "github.com/fullstack-lang/gong/lib/split/go/models"

func (stager *Stager) createViews() {
	stager.splitStage.Reset()

	stager.splitStage.StageBranch(&split.View{
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
										Size: 78,
										Tree: (&split.Tree{
											StackName: stager.treeStage.GetName(),
										}),
									},
									{
										Size: 11,
										AsSplit: &split.AsSplit{
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
													Load: &split.Load{
														StackName: stager.loadStageMultistage.GetName(),
													},
												},
											},
										},
									},
									{
										Size: 11,
										Button: &split.Button{
											StackName: stager.buttonStage.GetName(),
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

	stager.splitStage.StageBranch(&split.View{
		Name: "probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.stage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	stager.splitStage.StageBranch(&split.View{
		Name: "About",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "Load  Buttons Download",
				ShowNameInHeader: false,
				AsSplit: &split.AsSplit{
					Name:      "as split",
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Size: 100,
							Markdown: &split.Markdown{
								StackName: stager.markdownStage.GetName(),
							},
						},
					},
				},
			},
		},
	})

	stager.splitStage.StageBranch(&split.View{
		Name:            "svg probe",
		IsSecondaryView: true,
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.svgStage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	stager.splitStage.StageBranch(&split.View{
		Name:            "markdown probe",
		IsSecondaryView: true,
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.markdownStage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	stager.splitStage.Commit()
}
