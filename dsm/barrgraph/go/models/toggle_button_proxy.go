package models

import tree "github.com/fullstack-lang/gong/lib/tree/go/models"

type toggleButtonProxy struct {
	stager      *Stager
	toggleValue *bool
}

func (p *toggleButtonProxy) ButtonUpdated(
	treeStage *tree.Stage,
	staged, front *tree.Button) {

	*p.toggleValue = !*p.toggleValue

	p.stager.stage.Commit()
}
