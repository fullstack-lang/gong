package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDiagramStructure(diagram *DiagramStructure, parentNode *tree.Node, expandedSlice *[]*DiagramStructure) {
	diagNode := &tree.Node{
		Name:            diagram.Name,
		IsExpanded:      diagram.IsExpanded,
		IsNodeClickable: true,
		IsInEditMode:    diagram.isInRenameMode,
		IsChecked:       diagram.IsChecked,
	}
	if diagNode.Name == "" {
		diagNode.Name = "Diagram"
	}
	parentNode.Children = append(parentNode.Children, diagNode)

	addRenameButton(diagram, diagNode, stager)

	diagCopy := diagram
	diagNode.OnClick = func(frontNode *tree.Node) {
		// uncheck all diagrams
		for d := range *GetGongstructInstancesSetFromPointerType[*DiagramStructure](stager.stage) {
			d.IsChecked = false
		}
		// check this diagram
		diagCopy.IsChecked = true
		stager.diagram = diagCopy
		stager.stage.Commit()
	}
}
