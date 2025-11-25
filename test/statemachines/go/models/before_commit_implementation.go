package models

type BeforeCommitImplementation struct {
	stager *Stager
}

func (c *BeforeCommitImplementation) BeforeCommit(stage *Stage) {
	c.stager.computeConsistency()
	c.stager.updateObjectTreeStage()
	c.stager.updateButtonsStage()
	c.stager.updateExportXLButtonStage()
	c.stager.updateSvgStage()
	c.stager.updateTreeDiagramStage()
}
