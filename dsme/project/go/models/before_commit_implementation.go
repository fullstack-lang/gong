package models

type BeforeCommitImplementation struct {
	stager *Stager
}

func (c *BeforeCommitImplementation) BeforeCommit(stage *Stage) {
	c.stager.enforceSemantic()
	c.stager.tree()
}
