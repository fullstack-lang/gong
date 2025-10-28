package models

import tree "github.com/fullstack-lang/gong/lib/tree/go/models"

type ButtonToggleShowLinkNamesProxy struct {
	stager *Stager
}

func (proxy *ButtonToggleShowLinkNamesProxy) ButtonUpdated(
	treeStage *tree.Stage,
	staged, front *tree.Button) {

	proxy.stager.hideLinkNames = !proxy.stager.hideLinkNames

	proxy.stager.UpdateAndCommitSVGStage()
	proxy.stager.UpdateAndCommitTreeStage()
	proxy.stager.stage.Commit()
}
