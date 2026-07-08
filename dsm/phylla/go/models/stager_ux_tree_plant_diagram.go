package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treePlantDiagram(
	plantDiagram *PlantDiagram,
	parentNodes *[]*tree.Node,
	plantDiagramsWhoseNodeIsExpanded *[]*PlantDiagram,
) {
	plantDiagramNode := &tree.Node{
		Name:            plantDiagram.Name,
		IsExpanded:      slices.Contains(*plantDiagramsWhoseNodeIsExpanded, plantDiagram),
		IsNodeClickable: true,
		IsInEditMode:    plantDiagram.isInRenameMode,
	}
	*parentNodes = append(*parentNodes, plantDiagramNode)

	plantDiagramNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&plantDiagram.IsExpanded)
	plantDiagramNode.OnNameChange = stager.onNameChange(plantDiagram)
	plantDiagramNode.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(plantDiagram, GetPointerToGongstructName[*PlantDiagram]())
		stager.stage.Commit()
	}

}
