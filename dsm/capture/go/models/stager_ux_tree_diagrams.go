package models

import (
	"cmp"
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDiagramBSinDiagram(currentDiagram *Diagram, library *Library, parentNode *tree.Node) {
	// First, sort and iterate over diagrams in this library
	diagrams := library.Diagrams
	slices.SortFunc(diagrams, func(a, b *Diagram) int {
		return cmp.Compare(a.Name, b.Name)
	})

	for _, targetDiagram := range diagrams {
		// Do not show the current diagram in its own diagram tree, to prevent infinite self-referencing diagram shapes?
		// Or maybe it's fine. Usually, you shouldn't embed a diagram into itself.
		if targetDiagram == currentDiagram {
			continue
		}

		n := &tree.Node{
			Name:                    targetDiagram.GetName(),
			IsExpanded:              true,
			IsNodeClickable:         true,
			HasCheckboxButton:       true,
			CheckboxHasToolTip:      true,
			CheckboxToolTipPosition: tree.Right,
		}
		parentNode.Children = append(parentNode.Children, n)

		diagramShape, ok := currentDiagram.map_Diagram_DiagramShape[targetDiagram]
		n.IsChecked = ok

		if ok {
			n.CheckboxToolTipText = "Uncheck to remove DiagramShape from diagram"
		} else {
			n.CheckboxToolTipText = "Check to add DiagramShape to diagram"
		}

		n.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
			if frontNode.IsChecked && !stagedNode.IsChecked {
				stagedNode.IsChecked = true
				newShapeToDiagram(targetDiagram, currentDiagram, &currentDiagram.Diagram_Shapes, stager.stage)
				stager.stage.Commit()
			}
			if !frontNode.IsChecked && stagedNode.IsChecked {
				stagedNode.IsChecked = false
				diagramShape.UnstageVoid(stager.stage)
				// Remove it from currentDiagram.Diagram_Shapes
				for i, ds := range currentDiagram.Diagram_Shapes {
					if ds == diagramShape {
						currentDiagram.Diagram_Shapes = append(currentDiagram.Diagram_Shapes[:i], currentDiagram.Diagram_Shapes[i+1:]...)
						break
					}
				}
				stager.stage.Commit()
			}
		}
	}

	// Recursively process sub-libraries
	subLibraries := library.SubLibraries
	slices.SortFunc(subLibraries, func(a, b *Library) int {
		return cmp.Compare(a.Name, b.Name)
	})

	for _, subLibrary := range subLibraries {
		libNode := &tree.Node{
			Name:            subLibrary.GetName(),
			IsExpanded:      true, // Expand by default or manage state if needed
			IsNodeClickable: true,
		}
		parentNode.Children = append(parentNode.Children, libNode)

		stager.treeDiagramBSinDiagram(currentDiagram, subLibrary, libNode)
	}
}
