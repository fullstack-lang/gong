package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) tree() {
	stager.treeStage.Reset()

	rootLibrary := stager.rootLibrary
	_ = rootLibrary

	treeInstance := &tree.Tree{Name: "Library Tree"}

	stager.probeForm.AddCommitNavigationNode(func(gni GongNodeIF) {
		treeInstance.RootNodes = append(treeInstance.RootNodes, gni.(*tree.Node))
	})

	stager.treeLibrary(rootLibrary, &treeInstance.RootNodes)

	tree.StageBranch(stager.treeStage, treeInstance)

	stager.treeStage.Commit()
}
