package models

import (
	split "github.com/fullstack-lang/gong/lib/split/go/models"
)

func getFileName(stager *Stager) string {
	if stager.stage.OnInitCommitCallback != nil {
		return stager.fileName
	} else {
		return "no persistance"
	}
}

func (stager *Stager) createViews() {
	stager.splitStage.Reset()

	tabTitle := &split.Title{
		Name: "State Machines (" + getFileName(stager) + ")",
	}
	tabTitle.Stage(stager.splitStage)

	split.StageBranch(stager.splitStage, &split.View{
		Name:           "Edit view (" + getFileName(stager) + ")",
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
																Name:             "Export XL button",
																Size:             10,
																ShowNameInHeader: false,
																Button: &split.Button{
																	StackName: stager.buttonStage.GetName(),
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
																	StackName: stager.treeStage.GetName(),
																},
															},

															//
															{
																Name:             "Export XL button",
																Size:             10,
																ShowNameInHeader: false,
																Button: &split.Button{
																	StackName: stager.buttonStage.GetName(),
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
							AsSplit: stager.probeForm.GetDataEditor(),
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
											StackName: stager.treeStage.GetName(),
										},
									},
									{
										Name:             "Bottom",
										Size:             50,
										ShowNameInHeader: false,
										Tree: &split.Tree{
											StackName: stager.treeObjectsSimulationStage.GetName(),
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
					StackName: stager.stage.GetProbeSplitStageName(),
				},
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "Data Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				AsSplit: stager.probeForm.GetDataEditor(),
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "Data Model",
		RootAsSplitAreas: []*split.AsSplitArea{
			stager.probeForm.GetDiagramEditor(),
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
					StackName: stager.treeStage.GetProbeSplitStageName(),
				},
			},
		},
	})

	stager.splitStage.Commit()
}
