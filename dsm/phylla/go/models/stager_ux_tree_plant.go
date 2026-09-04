package models

import (
	"encoding/base64"
	"fmt"
	"time"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treePlant(plant *PlantAbstract, parentNodes *[]*tree.Node, currentView ViewType) {
	nodeName := plant.Name
	if !plant.isInRenameMode && plant.PlantType == Vase {
		h0 := plant.VaseAbstract.heightAtRotRatio0
		h1 := plant.VaseAbstract.heightAtRotRatio1
		var ratio float64
		if h0 > 0 {
			ratio = h1 / h0
		}
		nodeName = fmt.Sprintf("%s (h(0): %.2f, h(1): %.2f, ratio: %.2f)", plant.Name, h0, h1, ratio)
	}

	plantNode := &tree.Node{
		Name:            nodeName,
		IsExpanded:      plant.IsExpanded,
		IsNodeClickable: true,
		IsInEditMode:    plant.isInRenameMode,
	}
	*parentNodes = append(*parentNodes, plantNode)

	addRenameButton(plant, plantNode, stager)

	plantNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&plant.IsExpanded)
	plantNode.OnNameChange = stager.onNameChange(plant)
	plantNode.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(plant, GetPointerToGongstructName[*PlantAbstract]())

		for p := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
			p.IsSelected = (p == plant)
		}

		stager.selectedPlant = plant
		if plant.PlantType == Stool {
			if plant.CurrentView != VIEW_PLANT_2D && plant.CurrentView != VIEW_PLANT_3D && plant.CurrentView != VIEW_STOOL_3D && plant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
				plant.CurrentView = VIEW_STOOL_3D
			}
		} else if plant.PlantType == Clock {
			if plant.CurrentView != VIEW_PLANT_2D && plant.CurrentView != VIEW_PLANT_3D && plant.CurrentView != VIEW_CLOCK_3D && plant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
				plant.CurrentView = VIEW_CLOCK_3D
			}
		} else if plant.PlantType == Vase {
			if plant.CurrentView != VIEW_PLANT_2D && plant.CurrentView != VIEW_PLANT_3D && plant.CurrentView != VIEW_VASE_FORM && plant.CurrentView != VIEW_VASE_2D && plant.CurrentView != VIEW_VASE_3D && plant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
				plant.CurrentView = VIEW_VASE_3D
			}
		} else if plant.PlantType == Music {
			if plant.CurrentView != VIEW_PLANT_2D && plant.CurrentView != VIEW_PLANT_3D && plant.CurrentView != VIEW_MUSIC_SCORE && plant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
				plant.CurrentView = VIEW_MUSIC_SCORE
			}
		} else {
			if plant.CurrentView != VIEW_PLANT_2D && plant.CurrentView != VIEW_PLANT_3D && plant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
				plant.CurrentView = VIEW_PLANT_2D
			}
		}

		// Ensure one diagram is checked by default if none are checked
		hasCheckedDiagram := false
		if !hasCheckedDiagram {
			for _, d := range plant.Plant2DDiagrams {
				if d.IsChecked {
					hasCheckedDiagram = true
					break
				}
			}
		}
		if !hasCheckedDiagram {
			for _, d := range plant.Plant3DDiagrams {
				if d.IsChecked {
					hasCheckedDiagram = true
					break
				}
			}
		}
		if !hasCheckedDiagram {
			for _, d := range plant.Vase2DDiagrams {
				if d.IsChecked {
					hasCheckedDiagram = true
					break
				}
			}
		}
		if !hasCheckedDiagram {
			for _, d := range plant.Vase3DDiagrams {
				if d.IsChecked {
					hasCheckedDiagram = true
					break
				}
			}
		}
		if !hasCheckedDiagram {
			for _, d := range plant.Stool2DDiagrams {
				if d.IsChecked {
					hasCheckedDiagram = true
					break
				}
			}
		}
		if !hasCheckedDiagram {
			for _, d := range plant.Stool3DDiagrams {
				if d.IsChecked {
					hasCheckedDiagram = true
					break
				}
			}
		}
		if !hasCheckedDiagram {
			for _, d := range plant.Clock2DDiagrams {
				if d.IsChecked {
					hasCheckedDiagram = true
					break
				}
			}
		}
		if !hasCheckedDiagram {
			for _, d := range plant.Clock3DDiagrams {
				if d.IsChecked {
					hasCheckedDiagram = true
					break
				}
			}
		}
		if !hasCheckedDiagram && plant.PlantType == Music && plant.MusicAbstract != nil {
			if plant.MusicAbstract.IsChecked {
				hasCheckedDiagram = true
			}
		}

		if !hasCheckedDiagram {
			// Uncheck all globally and select the first one if available
			uncheckAllDiagrams(stager)
			if plant.CurrentView == VIEW_PLANT_3D && len(plant.Plant3DDiagrams) > 0 {
				plant.Plant3DDiagrams[0].IsChecked = true
				plant.Plant3DDiagrams[0].IsExpanded = true
			} else if plant.PlantType == Vase {
				if plant.CurrentView == VIEW_VASE_2D && len(plant.Vase2DDiagrams) > 0 {
					plant.Vase2DDiagrams[0].IsChecked = true
				} else if len(plant.Vase3DDiagrams) > 0 {
					plant.Vase3DDiagrams[0].IsChecked = true
					plant.Vase3DDiagrams[0].IsExpanded = true
					plant.CurrentView = VIEW_VASE_3D
				} else if len(plant.Vase2DDiagrams) > 0 {
					plant.Vase2DDiagrams[0].IsChecked = true
					plant.CurrentView = VIEW_VASE_2D
				}
			} else if plant.PlantType == Stool {
				if len(plant.Stool3DDiagrams) > 0 {
					plant.Stool3DDiagrams[0].IsChecked = true
					plant.CurrentView = VIEW_STOOL_3D
				} else if len(plant.Stool2DDiagrams) > 0 {
					plant.Stool2DDiagrams[0].IsChecked = true
					plant.CurrentView = VIEW_PLANT_2D
				}
			} else if plant.PlantType == Clock {
				if len(plant.Clock3DDiagrams) > 0 {
					plant.Clock3DDiagrams[0].IsChecked = true
					plant.CurrentView = VIEW_CLOCK_3D
				} else if len(plant.Clock2DDiagrams) > 0 {
					plant.Clock2DDiagrams[0].IsChecked = true
					plant.CurrentView = VIEW_PLANT_2D
				}
			} else if plant.PlantType == Music {
				if plant.MusicAbstract != nil {
					plant.MusicAbstract.IsChecked = true
					plant.MusicAbstract.IsComposerNodeExpanded = true
				}
				plant.CurrentView = VIEW_MUSIC_SCORE
			} else {
				if len(plant.Plant2DDiagrams) > 0 {
					plant.Plant2DDiagrams[0].IsChecked = true
					plant.CurrentView = VIEW_PLANT_2D
				} else if len(plant.Plant3DDiagrams) > 0 {
					plant.Plant3DDiagrams[0].IsChecked = true
					plant.CurrentView = VIEW_PLANT_3D
				}
			}
		}

		stager.stage.Commit()
	}

	if plant.PlantType == Vase {
		topRingBtn := &tree.Button{
			Name:            "export Top ring STL",
			Icon:            string(buttons.BUTTON_vertical_align_top),
			ToolTipText:     "export Top ring STL",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				stager.loadStage.Reset()
				fileToDownload := new(load.FileToDownload).Stage(stager.loadStage)

				stlContent := GenerateTopRingSTL(plant)

				fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString([]byte(stlContent))
				fileToDownload.Name = time.Now().Format("20060102 1504 ") + "phylla-" + stager.stage.GetName() + "-" + plant.Name + "-top-ring.stl"
				stager.loadStage.Commit()

				time.Sleep(1 * time.Second)
				stager.load()
			},
		}
		plantNode.Menu.Buttons = append(plantNode.Menu.Buttons, topRingBtn)

		oneRingBtn := &tree.Button{
			Name:            "export one ring STL",
			Icon:            string(buttons.BUTTON_vertical_align_center),
			ToolTipText:     "export one ring STL",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				stager.loadStage.Reset()
				fileToDownload := new(load.FileToDownload).Stage(stager.loadStage)

				stlContent := GenerateOneRingSTL(plant)

				fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString([]byte(stlContent))
				fileToDownload.Name = time.Now().Format("20060102 1504 ") + "phylla-" + stager.stage.GetName() + "-" + plant.Name + "-one-ring.stl"
				stager.loadStage.Commit()

				time.Sleep(1 * time.Second)
				stager.load()
			},
		}
		plantNode.Menu.Buttons = append(plantNode.Menu.Buttons, oneRingBtn)

		bottomRingBtn := &tree.Button{
			Name:            "export bottom ring STL",
			Icon:            string(buttons.BUTTON_vertical_align_bottom),
			ToolTipText:     "export bottom ring STL",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				stager.loadStage.Reset()
				fileToDownload := new(load.FileToDownload).Stage(stager.loadStage)

				stlContent := GenerateBottomRingSTL(plant)

				fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString([]byte(stlContent))
				fileToDownload.Name = time.Now().Format("20060102 1504 ") + "phylla-" + stager.stage.GetName() + "-" + plant.Name + "-bottom-ring.stl"
				stager.loadStage.Commit()

				time.Sleep(1 * time.Second)
				stager.load()
			},
		}
		plantNode.Menu.Buttons = append(plantNode.Menu.Buttons, bottomRingBtn)
	} else {
		downloadBtn := &tree.Button{
			Name:            "Download STL",
			Icon:            string(buttons.BUTTON_download),
			ToolTipText:     "Download STL",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				stager.loadStage.Reset()
				fileToDownload := new(load.FileToDownload).Stage(stager.loadStage)

				// Generate the actual STL content
				stlContent := GenerateSTL(plant)

				fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString([]byte(stlContent))
				fileToDownload.Name = time.Now().Format("20060102 1504 ") + "phylla-" + stager.stage.GetName() + "-" + plant.Name + ".stl"
				stager.loadStage.Commit()

				time.Sleep(1 * time.Second) // Sleep to ensure the client has time to start the download before we reset the stage.
				stager.load()
			},
		}
		plantNode.Menu.Buttons = append(plantNode.Menu.Buttons, downloadBtn)
	}

	switch currentView {
	case VIEW_PLANT_2D:
		confPlant2D := ItemButtonConfiguration[Plant2DDiagram, *Plant2DDiagram, PlantAbstract, *PlantAbstract]{
			parentNode:                         plantNode,
			sliceForNewAddedItem:               &plant.Plant2DDiagrams,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &plant.IsExpanded,
			IsButtonInMenu:                     true,
		}
		addCreateItemButton(stager, confPlant2D)
		if len(plantNode.Menu.Buttons) > 0 {
			btn := plantNode.Menu.Buttons[len(plantNode.Menu.Buttons)-1]
			btn.Name = "Add Plant 2D Diagram"
			btn.ToolTipText = "Add a Plant 2D Diagram"
		}
		for _, diag := range plant.Plant2DDiagrams {
			stager.treePlant2DDiagram(plant, diag, &plantNode.Children, false)
		}
	case VIEW_PLANT_3D:
		confPlant3D := ItemButtonConfiguration[Plant3DDiagram, *Plant3DDiagram, PlantAbstract, *PlantAbstract]{
			parentNode:                         plantNode,
			sliceForNewAddedItem:               &plant.Plant3DDiagrams,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &plant.IsExpanded,
			IsButtonInMenu:                     true,
		}
		addCreateItemButton(stager, confPlant3D)
		if len(plantNode.Menu.Buttons) > 0 {
			btn := plantNode.Menu.Buttons[len(plantNode.Menu.Buttons)-1]
			btn.Name = "Add Plant 3D Diagram"
			btn.ToolTipText = "Add a Plant 3D Diagram"
		}
		for _, diag := range plant.Plant3DDiagrams {
			stager.treePlant3DDiagram(plant, diag, &plantNode.Children, true)
		}
	case VIEW_VASE_2D, VIEW_VASE_FORM:
		if plant.PlantType == Vase {
			confVase2D := ItemButtonConfiguration[Vase2DDiagram, *Vase2DDiagram, PlantAbstract, *PlantAbstract]{
				parentNode:                         plantNode,
				sliceForNewAddedItem:               &plant.Vase2DDiagrams,
				isParentNodeExpandedByAddOperation: true,
				parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
				parentNodeExpansionBooleanValue:    &plant.IsExpanded,
				IsButtonInMenu:                     true,
			}
			addCreateItemButton(stager, confVase2D)
			if len(plantNode.Menu.Buttons) > 0 {
				btn := plantNode.Menu.Buttons[len(plantNode.Menu.Buttons)-1]
				btn.Name = "Add Vase 2D Diagram"
				btn.ToolTipText = "Add a Vase 2D Diagram"
			}
			for _, diag := range plant.Vase2DDiagrams {
				stager.treeVase2DDiagram(plant, diag, &plantNode.Children, false)
			}
		}
	case VIEW_VASE_3D:
		if plant.PlantType == Vase {
			confVase3D := ItemButtonConfiguration[Vase3DDiagram, *Vase3DDiagram, PlantAbstract, *PlantAbstract]{
				parentNode:                         plantNode,
				sliceForNewAddedItem:               &plant.Vase3DDiagrams,
				isParentNodeExpandedByAddOperation: true,
				parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
				parentNodeExpansionBooleanValue:    &plant.IsExpanded,
				IsButtonInMenu:                     true,
			}
			addCreateItemButton(stager, confVase3D)
			if len(plantNode.Menu.Buttons) > 0 {
				btn := plantNode.Menu.Buttons[len(plantNode.Menu.Buttons)-1]
				btn.Name = "Add Vase 3D Diagram"
				btn.ToolTipText = "Add a Vase 3D Diagram"
			}
			for _, diag := range plant.Vase3DDiagrams {
				stager.treeVase3DDiagram(plant, diag, &plantNode.Children, true)
			}
		}
	case VIEW_STOOL_3D:
		if plant.PlantType == Stool {
			confStool3D := ItemButtonConfiguration[Stool3DDiagram, *Stool3DDiagram, PlantAbstract, *PlantAbstract]{
				parentNode:                         plantNode,
				sliceForNewAddedItem:               &plant.Stool3DDiagrams,
				isParentNodeExpandedByAddOperation: true,
				parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
				parentNodeExpansionBooleanValue:    &plant.IsExpanded,
				IsButtonInMenu:                     true,
			}
			addCreateItemButton(stager, confStool3D)
			if len(plantNode.Menu.Buttons) > 0 {
				btn := plantNode.Menu.Buttons[len(plantNode.Menu.Buttons)-1]
				btn.Name = "Add Stool 3D Diagram"
				btn.ToolTipText = "Add a Stool 3D Diagram"
			}
			for _, diag := range plant.Stool3DDiagrams {
				stager.treeStool3DDiagram(plant, diag, &plantNode.Children, true)
			}
		}
	case VIEW_CLOCK_3D:
		if plant.PlantType == Clock {
			confClock3D := ItemButtonConfiguration[Clock3DDiagram, *Clock3DDiagram, PlantAbstract, *PlantAbstract]{
				parentNode:                         plantNode,
				sliceForNewAddedItem:               &plant.Clock3DDiagrams,
				isParentNodeExpandedByAddOperation: true,
				parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
				parentNodeExpansionBooleanValue:    &plant.IsExpanded,
				IsButtonInMenu:                     true,
			}
			addCreateItemButton(stager, confClock3D)
			if len(plantNode.Menu.Buttons) > 0 {
				btn := plantNode.Menu.Buttons[len(plantNode.Menu.Buttons)-1]
				btn.Name = "Add Clock 3D Diagram"
				btn.ToolTipText = "Add a Clock 3D Diagram"
			}
			for _, diag := range plant.Clock3DDiagrams {
				stager.treeClock3DDiagram(plant, diag, &plantNode.Children, true)
			}
		}
	case VIEW_MUSIC_SCORE:
		if plant.PlantType == Music && plant.MusicAbstract != nil {
			ma := plant.MusicAbstract
			composerNode := &tree.Node{
				Name:              "Composer",
				IsExpanded:        ma.IsComposerNodeExpanded,
				IsNodeClickable:   true,
				HasCheckboxButton: true,
				IsChecked:         ma.IsChecked,
				HasToolTip:        true,
				ToolTipPosition:   tree.Right,
				ToolTipText:       "Check to select the Music Score view",
			}
			plantNode.Children = append(plantNode.Children, composerNode)
			composerNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&ma.IsComposerNodeExpanded)
			composerNode.OnIsCheckedChanged = func(isChecked bool) {
				if isChecked {
					uncheckAllDiagrams(stager)
					for p := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
						p.IsSelected = (p == plant)
					}
					stager.selectedPlant = plant
					ma.IsChecked = true
					plant.CurrentView = VIEW_MUSIC_SCORE
					stager.stage.Commit()
				} else {
					ma.IsChecked = false
					stager.stage.Commit()
				}
			}
			composerNode.OnClick = func(frontNode *tree.Node) {
				stager.probeForm.FillUpFormFromGongstruct(ma, GetPointerToGongstructName[*MusicAbstract]())
				uncheckAllDiagrams(stager)
				for p := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
					p.IsSelected = (p == plant)
				}
				stager.selectedPlant = plant
				ma.IsChecked = true
				plant.CurrentView = VIEW_MUSIC_SCORE
				stager.stage.Commit()
			}

			addVoiceShowToggle := func(name string, showRef *bool) {
				childNode := &tree.Node{
					Name:              name,
					HasCheckboxButton: false,
					IsNodeClickable:   true,
				}
				btn := &tree.Button{
					Name:            "Hide",
					Icon:            string(buttons.BUTTON_visibility_off),
					ToolTipText:     "Hide from score",
					HasToolTip:      true,
					ToolTipPosition: tree.Right,
				}
				if !*showRef {
					btn.Icon = string(buttons.BUTTON_visibility)
					btn.Name = "Show"
					btn.ToolTipText = "Show on score"
				}
				btn.OnClick = func() {
					*showRef = !*showRef
					if *showRef {
						btn.Icon = string(buttons.BUTTON_visibility_off)
						btn.Name = "Hide"
						btn.ToolTipText = "Hide from score"
					} else {
						btn.Icon = string(buttons.BUTTON_visibility)
						btn.Name = "Show"
						btn.ToolTipText = "Show on score"
					}
					stager.stage.CommitWithSuspendedCallbacks()
					stager.treeStage2D.Commit()
					stager.treeStage3D.Commit()
					stager.ux_svg_music()
				}
				childNode.Buttons = append(childNode.Buttons, btn)
				composerNode.Children = append(composerNode.Children, childNode)
			}

			addVoiceShowToggle("First Voice", &ma.ShowFirstVoice)
			addVoiceShowToggle("First Voice Shift Right", &ma.ShowFirstVoiceShiftRight)
			addVoiceShowToggle("2nd Voice", &ma.ShowSecondVoice)
			addVoiceShowToggle("2nd voice shifted right", &ma.ShowSecondVoiceShiftRight)
			addVoiceShowToggle("First Voice notes", &ma.ShowFirstVoiceNotes)
			addVoiceShowToggle("First Voice note shifted right", &ma.ShowFirstVoiceNotesShiftRight)
			addVoiceShowToggle("Second Voice notes", &ma.ShowSecondVoiceNotes)
			addVoiceShowToggle("Second Voice Notes Shift Right", &ma.ShowSecondVoiceNotesShiftRight)
		}
	case VIEW_ABOUT_SPIRAL_PLANTS:
		// In about view, no diagram nodes are attached to plant
	}
}

func uncheckAllDiagrams(stager *Stager) {
	for d := range *GetGongstructInstancesSetFromPointerType[*Plant2DDiagram](stager.stage) {
		d.IsChecked = false
	}
	for d := range *GetGongstructInstancesSetFromPointerType[*Plant3DDiagram](stager.stage) {
		d.IsChecked = false
	}
	for d := range *GetGongstructInstancesSetFromPointerType[*Vase2DDiagram](stager.stage) {
		d.IsChecked = false
	}
	for d := range *GetGongstructInstancesSetFromPointerType[*Vase3DDiagram](stager.stage) {
		d.IsChecked = false
	}
	for d := range *GetGongstructInstancesSetFromPointerType[*Stool2DDiagram](stager.stage) {
		d.IsChecked = false
	}
	for d := range *GetGongstructInstancesSetFromPointerType[*Stool3DDiagram](stager.stage) {
		d.IsChecked = false
	}
	for d := range *GetGongstructInstancesSetFromPointerType[*Clock2DDiagram](stager.stage) {
		d.IsChecked = false
	}
	for d := range *GetGongstructInstancesSetFromPointerType[*Clock3DDiagram](stager.stage) {
		d.IsChecked = false
	}
	for ma := range *GetGongstructInstancesSetFromPointerType[*MusicAbstract](stager.stage) {
		ma.IsChecked = false
	}
}
