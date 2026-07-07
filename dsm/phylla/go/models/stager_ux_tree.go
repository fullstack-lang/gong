package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) ux_tree() {
	stager.treeStage.Reset()

	rootLibrary := stager.getRootLibrary()
	_ = rootLibrary

	treeInstance := &tree.Tree{
		Name:       "Library Tree",
		HaveSearch: true,
	}

	stager.probeForm.AddCommitNavigationNode(func(gni GongNodeIF) {
		treeInstance.RootNodes = append(treeInstance.RootNodes, gni.(*tree.Node))
	})

	tree.StageBranch(stager.treeStage, treeInstance)

	stager.treeStage.Commit()
}
