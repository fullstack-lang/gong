package models

import tree "github.com/fullstack-lang/gong/lib/tree/go/models"

type ButtonToggleShowNbInstancesProxy struct {
	stager *Stager
}

func (proxy *ButtonToggleShowNbInstancesProxy) ButtonUpdated(
	treeStage *tree.Stage,
	staged, front *tree.Button) {

	proxy.stager.hideNbInstances = !proxy.stager.hideNbInstances

	proxy.stager.UpdateAndCommitSVGStage()
	proxy.stager.UpdateAndCommitTreeStage()
	proxy.stager.stage.Commit()
}
