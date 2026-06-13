package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type Diagram_Tree_DiagramProxy struct {
	stager           *Stager
	diagramHierarchy *DiagramHierarchy
}

// OnAfterUpdate implements models.NodeImplInterface.
func (p *Diagram_Tree_DiagramProxy) OnAfterUpdate(stage *tree.Stage, staged *tree.Node, front *tree.Node) {
	if front.IsChecked && !staged.IsChecked {

		// reset all ddiagram selection
		for diagram_ := range *GetGongstructInstancesSet[DiagramHierarchy](p.stager.stage) {
			diagram_.IsChecked = false
		}
		p.diagramHierarchy.IsChecked = true
		staged.IsChecked = front.IsChecked
		p.stager.stage.Commit()
		return
	}
	if !front.IsChecked && staged.IsChecked {
		p.diagramHierarchy.IsChecked = false
		staged.IsChecked = front.IsChecked

		// reset all ddiagram selection
		for diagram_ := range *GetGongstructInstancesSet[DiagramHierarchy](p.stager.stage) {
			diagram_.IsChecked = false
		}
		p.stager.stage.Commit()
		return
	}
	if front.IsExpanded && !staged.IsExpanded {
		staged.IsExpanded = front.IsExpanded
		p.diagramHierarchy.IsExpanded = true
		p.stager.stage.Commit()
		return
	}
	if !front.IsExpanded && staged.IsExpanded {
		staged.IsExpanded = front.IsExpanded
		p.diagramHierarchy.IsExpanded = false
		p.stager.stage.Commit()
		return
	}

	if front.Name != staged.Name {
		p.diagramHierarchy.Name = front.Name
		p.diagramHierarchy.isInRenameMode = false

		p.stager.stage.Commit()
		return
	}

	p.stager.probeForm.FillUpFormFromGongstruct(p.diagramHierarchy, "Diagram")
}
