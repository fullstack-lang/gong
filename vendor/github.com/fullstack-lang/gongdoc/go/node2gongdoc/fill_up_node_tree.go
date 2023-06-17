package node2gongdoc

import (
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

func FillUpNodeTree(
	gongdocStage *gongdoc_models.StageStruct,
	gongtreeStage *gongtree_models.StageStruct,
	diagramPackage *gongdoc_models.DiagramPackage,
) {

	treeOfGongObjects := FillUpTreeOfGongObjects(gongdocStage, gongtreeStage, diagramPackage)
	rootOfClassdiagramsNode := FillUpTreeOfDiagramNodes(gongdocStage, gongtreeStage, diagramPackage, treeOfGongObjects)

	computeNodeConfs(
		gongtreeStage,
		gongdocStage,
		rootOfClassdiagramsNode,
		diagramPackage,
		treeOfGongObjects)

	gongdocStage.Commit()
	gongtreeStage.Commit()
}
