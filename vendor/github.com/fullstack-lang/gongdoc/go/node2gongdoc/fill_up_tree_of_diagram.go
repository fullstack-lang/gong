package node2gongdoc

import (
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

func FillUpTreeOfDiagramNodes(
	gongdocStage *gongdoc_models.StageStruct,
	gongtreeStage *gongtree_models.StageStruct,
	diagramPackage *gongdoc_models.DiagramPackage,
	treeOfGongObjects *gongtree_models.Tree,
) (rootOfClassdiagramsNode *gongtree_models.Node) {

	// create a tree for classdiagrams
	gongdocTree := (&gongtree_models.Tree{Name: "gongdoc"}).Stage(gongtreeStage)

	// add the root of class diagrams
	rootOfClassdiagramsNode = (&gongtree_models.Node{Name: "Class diagrams"}).Stage(gongtreeStage)
	rootOfClassdiagramsNode.IsExpanded = true
	gongdocTree.RootNodes = append(gongdocTree.RootNodes, rootOfClassdiagramsNode)

	// add add button
	addDocButton := (&gongtree_models.Button{
		Name: diagramPackage.Name + " " + string(BUTTON_add),
		Icon: string(BUTTON_add)}).Stage(gongtreeStage)
	rootOfClassdiagramsNode.Buttons = append(rootOfClassdiagramsNode.Buttons, addDocButton)
	addDocButton.Impl = NewButtonImplRootOfClassdiagrams(
		diagramPackage,
		rootOfClassdiagramsNode,
		treeOfGongObjects,
		BUTTON_add,
	)

	// add one node for each diagram
	for classdiagram := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.Classdiagram](gongdocStage) {
		classdiagramNode := NewClassdiagramNode(
			gongtreeStage,
			classdiagram,
			diagramPackage,
			rootOfClassdiagramsNode,
			treeOfGongObjects)
		rootOfClassdiagramsNode.Children = append(rootOfClassdiagramsNode.Children, classdiagramNode)
	}

	return
}
