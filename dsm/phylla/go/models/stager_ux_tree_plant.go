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

		if !plant.IsSelected {
			for p := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
				p.IsSelected = false
			}
			plant.IsSelected = true
		}

		stager.selectedPlant = plant
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

	confPlants := ItemButtonConfiguration[
		PlantDiagram, *PlantDiagram, // AT, PAT (Added Element)
		PlantAbstract, *PlantAbstract, // ParentAT, PParentAT (Parent Element)
	]{
		parentNode:                         plantNode,
		sliceForNewAddedItem:               &plant.PlantDiagrams,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &plant.IsExpanded,
		IsButtonInMenu:                     true,
	}
	addCreateItemButton(stager, confPlants)

	plantNode.Menu.Buttons = append(plantNode.Menu.Buttons, downloadBtn)

	// Add Plant Diagram Button
	diagramsNode := &tree.Node{
		Name:            "Plant Diagrams",
		FontStyle:       tree.ITALIC,
		IsExpanded:      plant.IsPlantDiagramsNodeExpanded,
		IsNodeClickable: true,
	}
	plantNode.Children = append(plantNode.Children, diagramsNode)
	diagramsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&plant.IsPlantDiagramsNodeExpanded)
	diagramsNode.OnClick = onNodeClicked(stager, plant)

	for _, plantDiagram := range plant.PlantDiagrams {
		stager.treePlantDiagram(plant, plantDiagram, &diagramsNode.Children, is3DView)
	}
}
