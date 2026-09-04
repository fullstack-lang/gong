package models

import (
	"embed"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

var DataFS *embed.FS

func (stager *Stager) ux_tree() {
	stager.treeStage2D.Reset()
	stager.treeStage3D.Reset()

	rootLibrary := stager.getRootLibrary()
	_ = rootLibrary

	currentView := VIEW_PLANT_2D
	plant := stager.GetCurrentPlant()
	if plant != nil && plant.CurrentView != "" {
		currentView = plant.CurrentView
	}

	treeInstance := &tree.Tree{
		Name:       "Library Tree",
		HaveSearch: true,
	}
	stager.probeForm.AddCommitNavigationNode(func(gni GongNodeIF) {
		treeInstance.RootNodes = append(treeInstance.RootNodes, gni.(*tree.Node))
	})
	stager.treeLibrary(treeInstance, rootLibrary, &treeInstance.RootNodes, currentView)
	tree.StageBranch(stager.treeStage2D, treeInstance)
	stager.treeStage2D.Commit()
}
