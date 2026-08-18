package models

import (
	split "github.com/fullstack-lang/gong/lib/split/go/models"
)

func getPersistanceFile(stager *Stager) string {
	if stager.stage.OnInitCommitCallback != nil {
		return stager.persistanceFile
	} else {
		return "no persistance"
	}
}

func (stager *Stager) updateSelectedViewFromPlant(plant *PlantAbstract) {
	if plant == nil || plant.CurrentView == "" {
		return
	}
	modified := false
	for view := range *split.GetGongstructInstancesSetFromPointerType[*split.View](stager.splitStage) {
		isSelected := (view.Name == string(plant.CurrentView))
		if view.IsSelectedView != isSelected {
			view.IsSelectedView = isSelected
			modified = true
		}
	}
	if modified {
		stager.splitStage.Commit()
	}
}

func (stager *Stager) createViews() {
	stager.splitStage.Reset()

	split.SetOrchestratorOnAfterUpdate[split.View](stager.splitStage)

	tabTitle := &split.Title{
		Name: "Phylla (" + getPersistanceFile(stager) + ")",
	}
	tabTitle.Stage(stager.splitStage)

	plant := stager.GetCurrentPlant()
	currentView := VIEW_PLANT_2D
	isVase := (plant != nil && plant.PlantType == Vase)
	isStool := (plant != nil && plant.PlantType == Stool)
	isClock := (plant != nil && plant.PlantType == Clock)

	var isPlant2DChecked, isVase2DChecked, isVase3DChecked, isStool2DChecked, isStool3DChecked, isClock2DChecked, isClock3DChecked bool
	if plant != nil {
		for _, d := range plant.Plant2DDiagrams {
			if d.IsChecked {
				isPlant2DChecked = true
			}
		}
		for _, d := range plant.Vase2DDiagrams {
			if d.IsChecked {
				isVase2DChecked = true
			}
		}
		for _, d := range plant.Vase3DDiagrams {
			if d.IsChecked {
				isVase3DChecked = true
			}
		}
		for _, d := range plant.Stool2DDiagrams {
			if d.IsChecked {
				isStool2DChecked = true
			}
		}
		for _, d := range plant.Stool3DDiagrams {
			if d.IsChecked {
				isStool3DChecked = true
			}
		}
		for _, d := range plant.Clock2DDiagrams {
			if d.IsChecked {
				isClock2DChecked = true
			}
		}
		for _, d := range plant.Clock3DDiagrams {
			if d.IsChecked {
				isClock3DChecked = true
			}
		}
	}

	if plant != nil {
		if plant.PlantType == Plant && plant.CurrentView != VIEW_PLANT_2D && plant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
			plant.CurrentView = VIEW_PLANT_2D
		} else if plant.PlantType == Stool && plant.CurrentView != VIEW_PLANT_2D && plant.CurrentView != VIEW_STOOL_3D && plant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
			plant.CurrentView = VIEW_PLANT_2D
		} else if plant.PlantType == Clock && plant.CurrentView != VIEW_PLANT_2D && plant.CurrentView != VIEW_CLOCK_3D && plant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
			plant.CurrentView = VIEW_PLANT_2D
		} else if plant.PlantType == Vase && plant.CurrentView != VIEW_PLANT_2D && plant.CurrentView != VIEW_VASE_FORM && plant.CurrentView != VIEW_VASE_2D && plant.CurrentView != VIEW_VASE_3D && plant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
			plant.CurrentView = VIEW_PLANT_2D
		}
		if plant.CurrentView != "" {
			currentView = plant.CurrentView
		} else {
			plant.CurrentView = VIEW_PLANT_2D
		}
	}

	view0Name := string(VIEW_PLANT_2D)
	view1Name := string(VIEW_VASE_FORM)
	view2Name := string(VIEW_VASE_2D)
	view3Name := string(VIEW_VASE_3D)
	viewStool3DName := string(VIEW_STOOL_3D)
	viewClock3DName := string(VIEW_CLOCK_3D)
	viewAboutSpiralPlantsName := string(VIEW_ABOUT_SPIRAL_PLANTS)

	isView0Selected := (currentView == VIEW_PLANT_2D)
	isView1Selected := (currentView == VIEW_VASE_FORM)
	isView2Selected := (currentView == VIEW_VASE_2D)
	isView3Selected := (currentView == VIEW_VASE_3D)
	isViewStool3DSelected := (currentView == VIEW_STOOL_3D)
	isViewClock3DSelected := (currentView == VIEW_CLOCK_3D)
	isViewAboutSpiralPlantsSelected := (currentView == VIEW_ABOUT_SPIRAL_PLANTS)

	if !isView0Selected && !isView1Selected && !isView2Selected && !isView3Selected && !isViewStool3DSelected && !isViewClock3DSelected && !isViewAboutSpiralPlantsSelected {
		isView0Selected = true
	}

	v0 := &split.View{
		Name:           view0Name,
		Direction:      split.Horizontal,
		IsSizeInPixel:  true,
		IsSelectedView: isView0Selected,
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
											StackName: stager.treeStage2D.GetName(),
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
								StackName: stager.svgPlantStage.GetName(),
							},
						},
					},
				},
			},
			{
				Size: 525,
				AsSplit: &split.AsSplit{
					Direction: split.Vertical,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Size: 60,
							Slider: &split.Slider{
								StackName: stager.sliderStage.GetName(),
							},
						},
						{
							Size: 40,
							Form: &split.Form{
								StackName: stager.plantFormStage.GetName(),
							},
						},
					},
				},
			},
		},
	}
	if isPlant2DChecked || isStool2DChecked || isClock2DChecked || (!isVase2DChecked && !isVase3DChecked && !isStool3DChecked && !isClock3DChecked) {
		split.StageBranch(stager.splitStage, v0)
	}
	v0.OnClick = func() {
		plant := stager.GetCurrentPlant()
		if plant != nil && plant.CurrentView != VIEW_PLANT_2D {
			plant.CurrentView = VIEW_PLANT_2D
			stager.stage.Commit()
		}
	}

	if isVase {
		v1 := &split.View{
			Name:           view1Name,
			Direction:      split.Horizontal,
			IsSizeInPixel:  true,
			IsSelectedView: isView1Selected,
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
												StackName: stager.treeStage2D.GetName(),
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
									StackName: stager.svgVaseStage.GetName(),
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
		}
		if isVase2DChecked {
			split.StageBranch(stager.splitStage, v1)
		}
		v1.OnClick = func() {
			plant := stager.GetCurrentPlant()
			if plant != nil && plant.CurrentView != VIEW_VASE_FORM {
				plant.CurrentView = VIEW_VASE_FORM
				stager.stage.Commit()
			}
		}

		v2 := &split.View{
			Name:           view2Name,
			Direction:      split.Horizontal,
			IsSizeInPixel:  true,
			IsSelectedView: isView2Selected,
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
												StackName: stager.treeStage2D.GetName(),
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
									StackName: stager.svgVaseStage.GetName(),
								},
							},
						},
					},
				},
				{
					Size: 525,
					Slider: &split.Slider{
						StackName: stager.sliderStage.GetName(),
					},
				},
			},
		}
		if isVase2DChecked {
			split.StageBranch(stager.splitStage, v2)
		}
		v2.OnClick = func() {
			plant := stager.GetCurrentPlant()
			if plant != nil && plant.CurrentView != VIEW_VASE_2D {
				plant.CurrentView = VIEW_VASE_2D
				stager.stage.Commit()
			}
		}

		v3 := &split.View{
			Name:           view3Name,
			Direction:      split.Horizontal,
			IsSizeInPixel:  true,
			IsSelectedView: isView3Selected,
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
												StackName: stager.treeStage3D.GetName(),
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
								Threejs: &split.Threejs{
									StackName: stager.threejsStage.GetName(),
								},
							},
						},
					},
				},
				{
					Size: 585,
					Slider: &split.Slider{
						StackName: stager.sliderStage.GetName(),
					},
				},
			},
		}
		if isVase3DChecked {
			split.StageBranch(stager.splitStage, v3)
		}
		v3.OnClick = func() {
			plant := stager.GetCurrentPlant()
			if plant != nil && plant.CurrentView != VIEW_VASE_3D {
				plant.CurrentView = VIEW_VASE_3D
				stager.stage.Commit()
			}
			stager.UpdateThreeJSStage()
		}
	}

	if isStool {
		vStool := &split.View{
			Name:           viewStool3DName,
			Direction:      split.Horizontal,
			IsSizeInPixel:  true,
			IsSelectedView: isViewStool3DSelected,
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
												StackName: stager.treeStage3D.GetName(),
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
								Threejs: &split.Threejs{
									StackName: stager.stool3dStage.GetName(),
								},
							},
						},
					},
				},
				{
					Size: 585,
					Slider: &split.Slider{
						StackName: stager.sliderStoolStage.GetName(),
					},
				},
			},
		}
		if isStool3DChecked {
			split.StageBranch(stager.splitStage, vStool)
		}
		vStool.OnClick = func() {
			plant := stager.GetCurrentPlant()
			if plant != nil && plant.CurrentView != VIEW_STOOL_3D {
				plant.CurrentView = VIEW_STOOL_3D
				stager.stage.Commit()
			}
			stager.UpdateStool3DStage()
		}
	}

	if isClock {
		vClock := &split.View{
			Name:           viewClock3DName,
			Direction:      split.Horizontal,
			IsSizeInPixel:  true,
			IsSelectedView: isViewClock3DSelected,
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
												StackName: stager.treeStage3D.GetName(),
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
								Threejs: &split.Threejs{
									StackName: stager.clock3dStage.GetName(),
								},
							},
						},
					},
				},
				{
					Size: 585,
					Slider: &split.Slider{
						StackName: stager.sliderClockStage.GetName(),
					},
				},
			},
		}
		if isClock3DChecked {
			split.StageBranch(stager.splitStage, vClock)
		}
		vClock.OnClick = func() {
			plant := stager.GetCurrentPlant()
			if plant != nil && plant.CurrentView != VIEW_CLOCK_3D {
				plant.CurrentView = VIEW_CLOCK_3D
				stager.stage.Commit()
			}
			stager.UpdateClock3DStage()
		}
	}

	vAbout := &split.View{
		Name:           viewAboutSpiralPlantsName,
		Direction:      split.Horizontal,
		IsSizeInPixel:  true,
		IsSelectedView: isViewAboutSpiralPlantsSelected,
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Name:             "Sidebar with both trees",
				ShowNameInHeader: false,
				Size:             525,
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
											StackName: stager.treeStage2D.GetName(),
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
					},
				},
			},
			{
				IsAny: true,
				Markdown: &split.Markdown{
					StackName: stager.markdownStage.GetName(),
				},
			},
		},
	}
	split.StageBranch(stager.splitStage, vAbout)
	vAbout.OnClick = func() {
		plant := stager.GetCurrentPlant()
		if plant != nil && plant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
			plant.CurrentView = VIEW_ABOUT_SPIRAL_PLANTS
			stager.stage.Commit()
		}
	}

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
					StackName: stager.treeStage2D.GetProbeSplitStageName(),
				},
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name:            "Svg Plant Probe",
		IsSecondaryView: true,
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stager.svgPlantStage.GetProbeSplitStageName(),
				},
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name:            "Svg Vase Probe",
		IsSecondaryView: true,
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stager.svgVaseStage.GetProbeSplitStageName(),
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

	split.StageBranch(stager.splitStage, &split.View{
		Name:            "threejs Probe",
		IsSecondaryView: true,
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stager.threejsStage.GetProbeSplitStageName(),
				},
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name:            "stool3d Probe",
		IsSecondaryView: true,
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stager.stool3dStage.GetProbeSplitStageName(),
				},
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name:            "sliderStool Probe",
		IsSecondaryView: true,
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stager.sliderStoolStage.GetProbeSplitStageName(),
				},
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name:            "clock3d Probe",
		IsSecondaryView: true,
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stager.clock3dStage.GetProbeSplitStageName(),
				},
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name:            "sliderClock Probe",
		IsSecondaryView: true,
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stager.sliderClockStage.GetProbeSplitStageName(),
				},
			},
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name:            "markdown probe",
		IsSecondaryView: true,
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stager.markdownStage.GetProbeSplitStageName(),
				},
			},
		},
	})
	stager.splitStage.Commit()
}
