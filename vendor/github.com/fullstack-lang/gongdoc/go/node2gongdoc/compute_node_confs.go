package node2gongdoc

import (
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

func computeNodeConfs(
	gongtreeStage *gongtree_models.StageStruct,
	gongdocStage *gongdoc_models.StageStruct,
	diagramPackageNode *gongtree_models.Node,
	diagramPackage *gongdoc_models.DiagramPackage,
	treeOfGongObjects *gongtree_models.Tree,
) {

	computeClassdiagramNodesConfigurations(
		diagramPackageNode,
		diagramPackage,
		gongdocStage,
		gongtreeStage)

	classdiagram := diagramPackage.SelectedClassdiagram

	// if no diagram is selected, all gong nodes are to be disabled
	isCheckboxDisabled := true

	// is the classdiagram is not in drawing mode, all gong nodes are to be disabled
	// otherwise gong nodes are enabled by default
	if classdiagram != nil {
		isCheckboxDisabled = !classdiagram.IsInDrawMode
	}

	// now, compute wether each gong node to be checked / disabled
	for _, _node := range treeOfGongObjects.RootNodes {
		applyGongNodesConfRecursively(_node, isCheckboxDisabled, false)
	}

	// no selected diagram yet
	if classdiagram == nil {
		gongdocStage.Commit()
		return
	}

	computeGongNodesConfigurations(gongdocStage, classdiagram, treeOfGongObjects)

	// log.Println("UpdateNodeStates, before commit, nb ", stage.BackRepo.GetLastCommitFromBackNb())
	gongdocStage.Commit()
}
