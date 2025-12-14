package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type Diagram_Tree_DiagramProxy struct {
	stager  *Stager
	diagram *Diagram
}

// OnAfterUpdate implements models.NodeImplInterface.
func (p *Diagram_Tree_DiagramProxy) OnAfterUpdate(stage *tree.Stage, staged *tree.Node, front *tree.Node) {
	if front.IsChecked && !staged.IsChecked {

		// reset all ddiagram selection
		for diagram_ := range *GetGongstructInstancesSet[Diagram](p.stager.stage) {
			diagram_.IsChecked = false
		}
		p.diagram.IsChecked = true
		staged.IsChecked = front.IsChecked
		p.stager.stage.Commit()
		return
	}
	if !front.IsChecked && staged.IsChecked {
		p.diagram.IsChecked = false
		staged.IsChecked = front.IsChecked

		// reset all ddiagram selection
		for diagram_ := range *GetGongstructInstancesSet[Diagram](p.stager.stage) {
			diagram_.IsChecked = false
		}
		p.stager.stage.Commit()
		return
	}
	if front.IsExpanded && !staged.IsExpanded {
		staged.IsExpanded = front.IsExpanded
		p.diagram.IsExpanded = true
		p.stager.stage.Commit()
		return
	}
	if !front.IsExpanded && staged.IsExpanded {
		staged.IsExpanded = front.IsExpanded
		p.diagram.IsExpanded = false
		p.stager.stage.Commit()
		return
	}

	if front.Name != staged.Name {
		p.diagram.Name = front.Name
		p.diagram.IsInRenameMode = false

		p.stager.stage.Commit()
		return
	}

	p.stager.probeForm.FillUpFormFromGongstruct(p.diagram, "Diagram")
}
