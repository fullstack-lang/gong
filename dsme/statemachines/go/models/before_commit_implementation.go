package models

type BeforeCommitImplementation struct {
	stager *Stager
}

func (c *BeforeCommitImplementation) BeforeCommit(stage *Stage) {
	c.stager.computeConsistency()
	c.stager.treeObjectSimulation()
	c.stager.updateButtonsStage()
	c.stager.updateExportXLButtonStage()
	c.stager.svg()
	c.stager.updateTreeDiagramStage()
}
