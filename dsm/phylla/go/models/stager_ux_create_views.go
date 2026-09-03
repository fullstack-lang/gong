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
	isMusic := (plant != nil && plant.PlantType == Music)

	var isPlant2DChecked, isPlant3DChecked, isVase2DChecked, isVase3DChecked, isStool2DChecked, isStool3DChecked, isClock2DChecked, isClock3DChecked, isMusicScoreChecked bool
	if plant != nil {
		for _, d := range plant.Plant2DDiagrams {
			if d.IsChecked {
				isPlant2DChecked = true
			}
		}
		for _, d := range plant.Plant3DDiagrams {
			if d.IsChecked {
				isPlant3DChecked = true
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
		if plant.MusicAbstract != nil && plant.MusicAbstract.IsChecked {
			isMusicScoreChecked = true
		}
	}

	if plant != nil {
		if plant.PlantType == Plant && plant.CurrentView != VIEW_PLANT_2D && plant.CurrentView != VIEW_PLANT_3D && plant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
			plant.CurrentView = VIEW_PLANT_2D
		} else if plant.PlantType == Stool && plant.CurrentView != VIEW_PLANT_2D && plant.CurrentView != VIEW_PLANT_3D && plant.CurrentView != VIEW_STOOL_3D && plant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
			plant.CurrentView = VIEW_PLANT_2D
		} else if plant.PlantType == Clock && plant.CurrentView != VIEW_PLANT_2D && plant.CurrentView != VIEW_PLANT_3D && plant.CurrentView != VIEW_CLOCK_3D && plant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
			plant.CurrentView = VIEW_PLANT_2D
		} else if plant.PlantType == Music && plant.CurrentView != VIEW_PLANT_2D && plant.CurrentView != VIEW_PLANT_3D && plant.CurrentView != VIEW_MUSIC_SCORE && plant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
			plant.CurrentView = VIEW_MUSIC_SCORE
		} else if plant.PlantType == Vase && plant.CurrentView != VIEW_PLANT_2D && plant.CurrentView != VIEW_PLANT_3D && plant.CurrentView != VIEW_VASE_FORM && plant.CurrentView != VIEW_VASE_2D && plant.CurrentView != VIEW_VASE_3D && plant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
			plant.CurrentView = VIEW_PLANT_2D
		}
		if isMusicScoreChecked {
			plant.CurrentView = VIEW_MUSIC_SCORE
		}
		if plant.CurrentView != "" {
			currentView = plant.CurrentView
		} else {
			plant.CurrentView = VIEW_PLANT_2D
		}
	}

	view0Name := string(VIEW_PLANT_2D)
	viewPlant3DName := string(VIEW_PLANT_3D)
	view1Name := string(VIEW_VASE_FORM)
	view2Name := string(VIEW_VASE_2D)
	view3Name := string(VIEW_VASE_3D)
	viewStool3DName := string(VIEW_STOOL_3D)
	viewClock3DName := string(VIEW_CLOCK_3D)
	viewMusicScoreName := string(VIEW_MUSIC_SCORE)
	viewAboutSpiralPlantsName := string(VIEW_ABOUT_SPIRAL_PLANTS)

	isView0Selected := (currentView == VIEW_PLANT_2D)
	isViewPlant3DSelected := (currentView == VIEW_PLANT_3D)
	isView1Selected := (currentView == VIEW_VASE_FORM)
	isView2Selected := (currentView == VIEW_VASE_2D)
	isView3Selected := (currentView == VIEW_VASE_3D)
	isViewStool3DSelected := (currentView == VIEW_STOOL_3D)
	isViewClock3DSelected := (currentView == VIEW_CLOCK_3D)
	isViewMusicScoreSelected := (currentView == VIEW_MUSIC_SCORE)
	isViewAboutSpiralPlantsSelected := (currentView == VIEW_ABOUT_SPIRAL_PLANTS)

	if isMusicScoreChecked {
		isViewMusicScoreSelected = true
		isView0Selected = false
		isViewPlant3DSelected = false
	} else if !isView0Selected && !isViewPlant3DSelected && !isView1Selected && !isView2Selected && !isView3Selected && !isViewStool3DSelected && !isViewClock3DSelected && !isViewMusicScoreSelected && !isViewAboutSpiralPlantsSelected {
		if isMusic {
			isViewMusicScoreSelected = true
		} else {
			isView0Selected = true
		}
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

	vPlant3D := &split.View{
		Name:           viewPlant3DName,
		Direction:      split.Horizontal,
		IsSizeInPixel:  true,
		IsSelectedView: isViewPlant3DSelected,
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
								StackName: stager.plant3dStage.GetName(),
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
	split.StageBranch(stager.splitStage, vPlant3D)
	vPlant3D.OnClick = func() {
		plant := stager.GetCurrentPlant()
		if plant != nil {
			if plant.CurrentView != VIEW_PLANT_3D {
				plant.CurrentView = VIEW_PLANT_3D
			}
			if !isPlant3DChecked && len(plant.Plant3DDiagrams) > 0 {
				uncheckAllDiagrams(stager)
				plant.Plant3DDiagrams[0].IsChecked = true
				plant.Plant3DDiagrams[0].IsExpanded = true
				plant.IsPlant3DDiagramsNodeExpanded = true
			}
			stager.stage.Commit()
		}
		stager.UpdatePlant3DStage()
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
		split.StageBranch(stager.splitStage, v1)
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
		split.StageBranch(stager.splitStage, v2)
		v2.OnClick = func() {
			plant := stager.GetCurrentPlant()
			if plant != nil {
				if plant.CurrentView != VIEW_VASE_2D {
					plant.CurrentView = VIEW_VASE_2D
				}
				if !isVase2DChecked && len(plant.Vase2DDiagrams) > 0 {
					uncheckAllDiagrams(stager)
					plant.Vase2DDiagrams[0].IsChecked = true
				}
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
		split.StageBranch(stager.splitStage, v3)
		v3.OnClick = func() {
			plant := stager.GetCurrentPlant()
			if plant != nil {
				if plant.CurrentView != VIEW_VASE_3D {
					plant.CurrentView = VIEW_VASE_3D
				}
				if !isVase3DChecked && len(plant.Vase3DDiagrams) > 0 {
					uncheckAllDiagrams(stager)
					plant.Vase3DDiagrams[0].IsChecked = true
					plant.Vase3DDiagrams[0].IsExpanded = true
					plant.IsVase3DDiagramsNodeExpanded = true
				}
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

	if isMusic {
		vMusicScore := &split.View{
			Name:           viewMusicScoreName,
			Direction:      split.Horizontal,
			IsSizeInPixel:  true,
			IsSelectedView: isViewMusicScoreSelected,
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
								IsAny:    true,
								HasDiv:   true,
								DivStyle: "position: relative; width: 100%; height: 100%;",
								Svg: &split.Svg{
									StackName: stager.svgMusicStage.GetName(),
									Style:     "position: absolute;top: 0;left: 0;width: 100%;height: 100%;z-index: 1;",
								},
								Cursor: &split.Cursor{
									StackName: stager.cursorStage.GetName(),
									Style:     "position: absolute;top: 0;left: 0;width: 100%;height: 100%;z-index: 2;pointer-events: none;",
								},
							},
						},
					},
				},
				{
					Size: 585,
					AsSplit: &split.AsSplit{
						Direction: split.Vertical,
						AsSplitAreas: []*split.AsSplitArea{
							{
								Size: 82,
								Slider: &split.Slider{
									StackName: stager.sliderMusicStage.GetName(),
								},
							},
							{
								Size: 9,
								Button: &split.Button{
									StackName: stager.buttonMusicStage.GetName(),
								},
							},
							{
								Size: 9,
								Tone: &split.Tone{
									StackName: stager.toneStage.GetName(),
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
			},
		}
		split.StageBranch(stager.splitStage, vMusicScore)
		vMusicScore.OnClick = func() {
			plant := stager.GetCurrentPlant()
			if plant != nil && plant.CurrentView != VIEW_MUSIC_SCORE {
				plant.CurrentView = VIEW_MUSIC_SCORE
				stager.stage.Commit()
			}
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
		Name:            "plant3d Probe",
		IsSecondaryView: true,
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stager.plant3dStage.GetProbeSplitStageName(),
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
