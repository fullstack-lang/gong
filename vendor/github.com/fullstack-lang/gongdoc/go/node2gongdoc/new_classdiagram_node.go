package node2gongdoc

import gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

func NewClassdiagramNode(
	classdiagram *gongdoc_models.Classdiagram,
	diagramPackage *gongdoc_models.DiagramPackage,
	rootOfClassdiagramsNode *gongdoc_models.Node,
	treeOfGongObjects *gongdoc_models.Tree,
) (classdiagramNode *gongdoc_models.Node) {
	classdiagramNode = (&gongdoc_models.Node{Name: classdiagram.Name}).Stage(diagramPackage.Stage_)

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
