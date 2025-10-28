package models

import tree "github.com/fullstack-lang/gong/lib/tree/go/models"

type ButtonToggleShowMultiplicityProxy struct {
	stager *Stager
}

func (proxy *ButtonToggleShowMultiplicityProxy) ButtonUpdated(
	treeStage *tree.Stage,
	staged, front *tree.Button) {

	proxy.stager.hideMultiplicity = !proxy.stager.hideMultiplicity

	proxy.stager.UpdateAndCommitSVGStage()
	proxy.stager.UpdateAndCommitTreeStage()
	proxy.stager.stage.Commit()
}
