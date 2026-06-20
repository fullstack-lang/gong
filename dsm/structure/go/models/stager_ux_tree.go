package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) ux_tree() {
	stager.treeStage.Reset()

	treeInstance := &tree.Tree{Name: "Library Tree"}

	for library := range *GetGongstructInstancesSetFromPointerType[*Library](stager.stage) {
		if library.IsRootLibrary {
			stager.treeLibrary(library, &treeInstance.RootNodes)
		}
	}
	
	tree.StageBranch(stager.treeStage, treeInstance)
	stager.treeStage.Commit()
}
