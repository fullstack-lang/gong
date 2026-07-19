package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) ux_tree() {
	stager.treeStage2D.Reset()
	stager.treeStage3D.Reset()

	rootLibrary := stager.getRootLibrary()
	_ = rootLibrary

	treeInstance2D := &tree.Tree{
		Name:       "Library Tree",
		HaveSearch: true,
	}
	stager.probeForm.AddCommitNavigationNode(func(gni GongNodeIF) {
		treeInstance2D.RootNodes = append(treeInstance2D.RootNodes, gni.(*tree.Node))
	})
	stager.treeLibrary(treeInstance2D, rootLibrary, &treeInstance2D.RootNodes, false)
	tree.StageBranch(stager.treeStage2D, treeInstance2D)
	stager.treeStage2D.Commit()

	treeInstance3D := &tree.Tree{
		Name:       "Library Tree",
		HaveSearch: true,
	}
	stager.probeForm.AddCommitNavigationNode(func(gni GongNodeIF) {
		treeInstance3D.RootNodes = append(treeInstance3D.RootNodes, gni.(*tree.Node))
	})
	stager.treeLibrary(treeInstance3D, rootLibrary, &treeInstance3D.RootNodes, true)
	tree.StageBranch(stager.treeStage3D, treeInstance3D)
	stager.treeStage3D.Commit()
}
