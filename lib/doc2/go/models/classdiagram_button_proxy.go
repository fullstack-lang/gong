package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type ButtonType int

const (
	DUPLICATE ButtonType = iota
	EDIT_CANCEL
	EDIT
	REMOVE
	RENAME_CANCEL
	RENAME
	SAVE
)

type ClassDiagramButtonProxy struct {
	stager       *Stager
	classdiagram *Classdiagram
	node         *tree.Node

	ButtonType
}

func NewClassDiagramButtonProxy(
	stager *Stager,
	classdiagram *Classdiagram,
	node *tree.Node,
	buttonType ButtonType,
) (proxy *ClassDiagramButtonProxy) {

	proxy = new(ClassDiagramButtonProxy)

	proxy.node = node
	proxy.stager = stager
	proxy.classdiagram = classdiagram
	proxy.ButtonType = buttonType

	return
}

func (proxy *ClassDiagramButtonProxy) ButtonUpdated(
	treeStage *tree.Stage,
	stageButton, front *tree.Button) {

	diagramPackage := getTheDiagramPackage(proxy.stager.stage)

	switch proxy.ButtonType {
	case DUPLICATE:
		duplicateDiagram := proxy.classdiagram.DuplicateDiagram()
		duplicateDiagram.Name += " Copy"

		proxy.stager.treeStage.Commit()

	// case EDIT_CANCEL:
	// 	map_ModelNode_Shape := proxy.portfolioDiagramNode.CancelEdit()

	// 	proxy.diagrammer.generatePortfolioNodesStatusAndButtons()
	// 	proxy.diagrammer.computeModelNodeStatus(map_ModelNode_Shape)
	// 	proxy.stager.treeStage.Commit()
	// case EDIT:
	// 	map_ModelNode_Shape := proxy.portfolioDiagramNode.DrawDiagram()

	// 	proxy.diagrammer.generatePortfolioNodesStatusAndButtons()
	// 	proxy.diagrammer.computeModelNodeStatus(map_ModelNode_Shape)
	// 	proxy.stager.treeStage.Commit()
	case REMOVE:

		// remove the classdiagram node from the pkg element node
		idx := slices.Index(diagramPackage.Classdiagrams, proxy.classdiagram)
		diagramPackage.Classdiagrams = slices.Delete(diagramPackage.Classdiagrams, idx, idx+1)
		UnstageBranch(proxy.stager.stage, proxy.classdiagram)

		if diagramPackage.SelectedClassdiagram == proxy.classdiagram {
			diagramPackage.SelectedClassdiagram = nil
		}

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitSVGStage()
		proxy.stager.stage.Commit()

	case RENAME:
		proxy.classdiagram.IsInRenameMode = true

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.stage.Commit()

	case RENAME_CANCEL:
		proxy.classdiagram.IsInRenameMode = false

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.stage.Commit()

	case SAVE:
		proxy.classdiagram.IsInRenameMode = false
		proxy.classdiagram.Name = front.GetName()

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.stage.Commit()

	}

}
