package models

import (
	"encoding/base64"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treePlant(plant *Plant, parentNodes *[]*tree.Node) {
	plantNode := &tree.Node{
		Name:            plant.Name,
		IsExpanded:      plant.IsExpanded,
		IsNodeClickable: true,
		IsInEditMode:    plant.isInRenameMode,
	}
	*parentNodes = append(*parentNodes, plantNode)

	plantNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&plant.IsExpanded)
	plantNode.OnNameChange = stager.onNameChange(plant)
	plantNode.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(plant, GetPointerToGongstructName[*Plant]())

		if !plant.IsSelected {
			for p := range stager.stage.Plants {
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
			
			// Mockup of STL file generation in memory
			stlContent := "solid plant\n" +
				"  facet normal 0 0 0\n" +
				"    outer loop\n" +
				"      vertex 0 0 0\n" +
				"      vertex 1 0 0\n" +
				"      vertex 0 1 0\n" +
				"    endloop\n" +
				"  endfacet\n" +
				"endsolid plant\n"
			
			fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString([]byte(stlContent))
			fileToDownload.Name = plant.Name + ".stl"
			stager.loadStage.Commit()
		},
	}
	plantNode.Buttons = append(plantNode.Buttons, downloadBtn)

	confPlants := ItemButtonConfiguration[
		PlantDiagram, *PlantDiagram, // AT, PAT (Added Element)
		Plant, *Plant, // ParentAT, PParentAT (Parent Element)
	]{
		parentNode:                         plantNode,
		sliceForNewAddedItem:               &plant.PlantDiagrams,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &plant.IsExpanded,
		IsButtonInMenu:                     true,
	}
	addCreateItemButton(stager, confPlants)

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
		stager.treePlantDiagram(plant, plantDiagram, &diagramsNode.Children)
	}
}
