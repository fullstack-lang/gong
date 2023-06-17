package node2gongdoc

import (
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

func NewClassdiagramNode(
	gongtreeStage *gongtree_models.StageStruct,
	classdiagram *gongdoc_models.Classdiagram,
	diagramPackage *gongdoc_models.DiagramPackage,
	rootOfClassdiagramsNode *gongtree_models.Node,
	treeOfGongObjects *gongtree_models.Tree,
) (classdiagramNode *gongtree_models.Node) {
	classdiagramNode = (&gongtree_models.Node{Name: classdiagram.Name}).Stage(gongtreeStage)

	rootOfClassdiagramsNode.Children = append(rootOfClassdiagramsNode.Children, classdiagramNode)

	classdiagramNode.HasCheckboxButton = true

	nodeImplClassdiagram := NewNodeImplClasssiagram(
		diagramPackage,
		classdiagram,
		rootOfClassdiagramsNode,
		treeOfGongObjects,
	)
	classdiagramNode.Impl = nodeImplClassdiagram

	return
}
