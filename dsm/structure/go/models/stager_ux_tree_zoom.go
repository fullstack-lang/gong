package models

import tree "github.com/fullstack-lang/gong/lib/tree/go/models"

func (stager *Stager) treeZoom() {
	stager.zoomTreeStage.Reset()

	treeInstance := &tree.Tree{Name: "Zoom Tree"}

	stager.probeForm.AddCommitNavigationNode(func(gni GongNodeIF) {
		treeInstance.RootNodes = append(treeInstance.RootNodes, gni.(*tree.Node))
	})

	var diagramProcess *DiagramProcess
	for diagramprocess_ := range *GetGongstructInstancesSet[DiagramProcess](stager.stage) {
		if diagramprocess_.IsChecked {
			diagramProcess = diagramprocess_
			break
		}
	}

	if diagramProcess != nil && diagramProcess.owningProcess != nil {
		dummyNode := &tree.Node{}
		stager.treeDiagramProcess(diagramProcess.owningProcess, diagramProcess, dummyNode)
		treeInstance.RootNodes = append(treeInstance.RootNodes, dummyNode.Children...)
	}

	tree.StageBranch(stager.zoomTreeStage, treeInstance)

	stager.zoomTreeStage.Commit()
}
