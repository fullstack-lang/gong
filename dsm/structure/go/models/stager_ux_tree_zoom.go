package models

import tree "github.com/fullstack-lang/gong/lib/tree/go/models"

func (stager *Stager) treeZoom() {
	stager.zoomTreeStage.Reset()

	treeInstance := &tree.Tree{Name: "Zoom Tree"}

	stager.probeForm.AddCommitNavigationNode(func(gni GongNodeIF) {
		treeInstance.RootNodes = append(treeInstance.RootNodes, gni.(*tree.Node))
	})

	var diagramStructure *DiagramStructure
	for diagramsystem_ := range *GetGongstructInstancesSet[DiagramStructure](stager.stage) {
		if diagramsystem_.IsChecked {
			diagramStructure = diagramsystem_
			break
		}
	}

	if diagramStructure != nil && diagramStructure.owningSystem != nil {
		dummyNode := &tree.Node{}
		stager.treeDiagramStructure(diagramStructure.owningSystem, diagramStructure, dummyNode)
		treeInstance.RootNodes = append(treeInstance.RootNodes, dummyNode.Children...)
	}

	tree.StageBranch(stager.zoomTreeStage, treeInstance)

	stager.zoomTreeStage.Commit()
}
