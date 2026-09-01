package models

import (
	"encoding/base64"
	"fmt"
	"time"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treePlant(plant *PlantAbstract, parentNodes *[]*tree.Node, is3DView bool) {
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
			if plant.CurrentView != VIEW_PLANT_2D && plant.CurrentView != VIEW_STOOL_3D && plant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
				plant.CurrentView = VIEW_STOOL_3D
			}
		} else if plant.PlantType == Clock {
			if plant.CurrentView != VIEW_PLANT_2D && plant.CurrentView != VIEW_CLOCK_3D && plant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
				plant.CurrentView = VIEW_CLOCK_3D
			}
		} else if plant.PlantType == Vase {
			if plant.CurrentView != VIEW_PLANT_2D && plant.CurrentView != VIEW_VASE_FORM && plant.CurrentView != VIEW_VASE_2D && plant.CurrentView != VIEW_VASE_3D && plant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
				plant.CurrentView = VIEW_VASE_3D
			}
		} else {
			if plant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
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

		if !hasCheckedDiagram {
			// Uncheck all globally and select the first one if available
			uncheckAllDiagrams(stager)
			if plant.PlantType == Vase {
				if len(plant.Vase3DDiagrams) > 0 {
					plant.Vase3DDiagrams[0].IsChecked = true
					plant.Vase3DDiagrams[0].IsExpanded = true
					plant.IsVase3DDiagramsNodeExpanded = true
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
			} else {
				if len(plant.Plant2DDiagrams) > 0 {
					plant.Plant2DDiagrams[0].IsChecked = true
					plant.CurrentView = VIEW_PLANT_2D
				}
			}
		}

		stager.stage.Commit()
	}

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

	diagramsNode := &tree.Node{
		Name:            "Plant 2D Diagrams",
		FontStyle:       tree.ITALIC,
		IsExpanded:      plant.IsPlant2DDiagramsNodeExpanded,
		IsNodeClickable: true,
	}
	plantNode.Children = append(plantNode.Children, diagramsNode)
	diagramsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&plant.IsPlant2DDiagramsNodeExpanded)
	diagramsNode.OnClick = onNodeClicked(stager, plant)
	for _, diag := range plant.Plant2DDiagrams {
		stager.treePlant2DDiagram(plant, diag, &diagramsNode.Children, is3DView)
	}

	if plant.PlantType == Vase { // Vase 2D Diagrams
		vase2DNode := &tree.Node{
			Name: "Vase 2D Diagrams", FontStyle: tree.ITALIC, IsExpanded: plant.IsVase2DDiagramsNodeExpanded, IsNodeClickable: true,
		}
		plantNode.Children = append(plantNode.Children, vase2DNode)
		vase2DNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&plant.IsVase2DDiagramsNodeExpanded)
		vase2DNode.OnClick = onNodeClicked(stager, plant)

		confVase2D := ItemButtonConfiguration[Vase2DDiagram, *Vase2DDiagram, PlantAbstract, *PlantAbstract]{
			parentNode: vase2DNode, sliceForNewAddedItem: &plant.Vase2DDiagrams,
			isParentNodeExpandedByAddOperation: true, parentNodeExpansionType: parentNodeExpansionTypeByBooleanValue, parentNodeExpansionBooleanValue: &plant.IsVase2DDiagramsNodeExpanded, IsButtonInMenu: true,
		}
		addCreateItemButton(stager, confVase2D)
		if len(vase2DNode.Menu.Buttons) > 0 {
			btn := vase2DNode.Menu.Buttons[len(vase2DNode.Menu.Buttons)-1]
			btn.Name = "Add Vase 2D Diagram"
			btn.ToolTipText = "Add a Vase 2D Diagram"
		}
		for _, diag := range plant.Vase2DDiagrams {
			stager.treeVase2DDiagram(plant, diag, &vase2DNode.Children, is3DView)
		} // Vase 3D Diagrams
		vase3DNode := &tree.Node{
			Name: "Vase 3D Diagrams", FontStyle: tree.ITALIC, IsExpanded: plant.IsVase3DDiagramsNodeExpanded, IsNodeClickable: true,
		}
		plantNode.Children = append(plantNode.Children, vase3DNode)
		vase3DNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&plant.IsVase3DDiagramsNodeExpanded)
		vase3DNode.OnClick = onNodeClicked(stager, plant)

		confVase3D := ItemButtonConfiguration[Vase3DDiagram, *Vase3DDiagram, PlantAbstract, *PlantAbstract]{
			parentNode: vase3DNode, sliceForNewAddedItem: &plant.Vase3DDiagrams,
			isParentNodeExpandedByAddOperation: true, parentNodeExpansionType: parentNodeExpansionTypeByBooleanValue, parentNodeExpansionBooleanValue: &plant.IsVase3DDiagramsNodeExpanded, IsButtonInMenu: true,
		}
		addCreateItemButton(stager, confVase3D)
		if len(vase3DNode.Menu.Buttons) > 0 {
			btn := vase3DNode.Menu.Buttons[len(vase3DNode.Menu.Buttons)-1]
			btn.Name = "Add Vase 3D Diagram"
			btn.ToolTipText = "Add a Vase 3D Diagram"
		}
		for _, diag := range plant.Vase3DDiagrams {
			stager.treeVase3DDiagram(plant, diag, &vase3DNode.Children, is3DView)
		}
	} else if plant.PlantType == Stool { // Stool 2D Diagrams
		stool2DNode := &tree.Node{
			Name: "Stool 2D Diagrams", FontStyle: tree.ITALIC, IsExpanded: plant.IsStool2DDiagramsNodeExpanded, IsNodeClickable: true,
		}
		plantNode.Children = append(plantNode.Children, stool2DNode)
		stool2DNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&plant.IsStool2DDiagramsNodeExpanded)
		stool2DNode.OnClick = onNodeClicked(stager, plant)

		confStool2D := ItemButtonConfiguration[Stool2DDiagram, *Stool2DDiagram, PlantAbstract, *PlantAbstract]{
			parentNode: stool2DNode, sliceForNewAddedItem: &plant.Stool2DDiagrams,
			isParentNodeExpandedByAddOperation: true, parentNodeExpansionType: parentNodeExpansionTypeByBooleanValue, parentNodeExpansionBooleanValue: &plant.IsStool2DDiagramsNodeExpanded, IsButtonInMenu: true,
		}
		addCreateItemButton(stager, confStool2D)
		if len(stool2DNode.Menu.Buttons) > 0 {
			btn := stool2DNode.Menu.Buttons[len(stool2DNode.Menu.Buttons)-1]
			btn.Name = "Add Stool 2D Diagram"
			btn.ToolTipText = "Add a Stool 2D Diagram"
		}
		for _, diag := range plant.Stool2DDiagrams {
			stager.treeStool2DDiagram(plant, diag, &stool2DNode.Children, is3DView)
		} // Stool 3D Diagrams
		stool3DNode := &tree.Node{
			Name: "Stool 3D Diagrams", FontStyle: tree.ITALIC, IsExpanded: plant.IsStool3DDiagramsNodeExpanded, IsNodeClickable: true,
		}
		plantNode.Children = append(plantNode.Children, stool3DNode)
		stool3DNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&plant.IsStool3DDiagramsNodeExpanded)
		stool3DNode.OnClick = onNodeClicked(stager, plant)

		confStool3D := ItemButtonConfiguration[Stool3DDiagram, *Stool3DDiagram, PlantAbstract, *PlantAbstract]{
			parentNode: stool3DNode, sliceForNewAddedItem: &plant.Stool3DDiagrams,
			isParentNodeExpandedByAddOperation: true, parentNodeExpansionType: parentNodeExpansionTypeByBooleanValue, parentNodeExpansionBooleanValue: &plant.IsStool3DDiagramsNodeExpanded, IsButtonInMenu: true,
		}
		addCreateItemButton(stager, confStool3D)
		if len(stool3DNode.Menu.Buttons) > 0 {
			btn := stool3DNode.Menu.Buttons[len(stool3DNode.Menu.Buttons)-1]
			btn.Name = "Add Stool 3D Diagram"
			btn.ToolTipText = "Add a Stool 3D Diagram"
		}
		for _, diag := range plant.Stool3DDiagrams {
			stager.treeStool3DDiagram(plant, diag, &stool3DNode.Children, is3DView)
		}
	} else if plant.PlantType == Clock { // Clock 2D Diagrams
		clock2DNode := &tree.Node{
			Name: "Clock 2D Diagrams", FontStyle: tree.ITALIC, IsExpanded: plant.IsClock2DDiagramsNodeExpanded, IsNodeClickable: true,
		}
		plantNode.Children = append(plantNode.Children, clock2DNode)
		clock2DNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&plant.IsClock2DDiagramsNodeExpanded)
		clock2DNode.OnClick = onNodeClicked(stager, plant)

		confClock2D := ItemButtonConfiguration[Clock2DDiagram, *Clock2DDiagram, PlantAbstract, *PlantAbstract]{
			parentNode: clock2DNode, sliceForNewAddedItem: &plant.Clock2DDiagrams,
			isParentNodeExpandedByAddOperation: true, parentNodeExpansionType: parentNodeExpansionTypeByBooleanValue, parentNodeExpansionBooleanValue: &plant.IsClock2DDiagramsNodeExpanded, IsButtonInMenu: true,
		}
		addCreateItemButton(stager, confClock2D)
		if len(clock2DNode.Menu.Buttons) > 0 {
			btn := clock2DNode.Menu.Buttons[len(clock2DNode.Menu.Buttons)-1]
			btn.Name = "Add Clock 2D Diagram"
			btn.ToolTipText = "Add a Clock 2D Diagram"
		}
		for _, diag := range plant.Clock2DDiagrams {
			stager.treeClock2DDiagram(plant, diag, &clock2DNode.Children, is3DView)
		} // Clock 3D Diagrams
		clock3DNode := &tree.Node{
			Name: "Clock 3D Diagrams", FontStyle: tree.ITALIC, IsExpanded: plant.IsClock3DDiagramsNodeExpanded, IsNodeClickable: true,
		}
		plantNode.Children = append(plantNode.Children, clock3DNode)
		clock3DNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&plant.IsClock3DDiagramsNodeExpanded)
		clock3DNode.OnClick = onNodeClicked(stager, plant)

		confClock3D := ItemButtonConfiguration[Clock3DDiagram, *Clock3DDiagram, PlantAbstract, *PlantAbstract]{
			parentNode: clock3DNode, sliceForNewAddedItem: &plant.Clock3DDiagrams,
			isParentNodeExpandedByAddOperation: true, parentNodeExpansionType: parentNodeExpansionTypeByBooleanValue, parentNodeExpansionBooleanValue: &plant.IsClock3DDiagramsNodeExpanded, IsButtonInMenu: true,
		}
		addCreateItemButton(stager, confClock3D)
		if len(clock3DNode.Menu.Buttons) > 0 {
			btn := clock3DNode.Menu.Buttons[len(clock3DNode.Menu.Buttons)-1]
			btn.Name = "Add Clock 3D Diagram"
			btn.ToolTipText = "Add a Clock 3D Diagram"
		}
		for _, diag := range plant.Clock3DDiagrams {
			stager.treeClock3DDiagram(plant, diag, &clock3DNode.Children, is3DView)
		}
	}
}

func uncheckAllDiagrams(stager *Stager) {
	for d := range *GetGongstructInstancesSetFromPointerType[*Plant2DDiagram](stager.stage) {
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
}
