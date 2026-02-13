package models

type BeforeCommitImplementation struct {
	stager *Stager
}

func (c *BeforeCommitImplementation) BeforeCommit(stage *Stage) {
	c.stager.enforceSemantic()

	var selectedDiagram *Diagram
	for _, diagram := range GetGongstrucsSorted[*Diagram](stage) {
		if diagram.IsEditable() {
			selectedDiagram = diagram
			break
		}
	}
	if selectedDiagram != nil {
		TaskOutputShapes := selectedDiagram.TaskOutputShapes
		_ = TaskOutputShapes
	}

	c.stager.tree()
	c.stager.svg()
}
