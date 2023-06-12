package node2gongdoc

import (
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

func FillUpNodeTree(gongdocStage *gongdoc_models.StageStruct, diagramPackage *gongdoc_models.DiagramPackage) {

	treeOfGongObjects := FillUpTreeOfGongObjects(gongdocStage, diagramPackage)

	// create a tree for classdiagrams
	gongdocTree := (&gongdoc_models.Tree{Name: "gongdoc"}).Stage(gongdocStage)

	// add the root of class diagrams
	rootOfClassdiagramsNode := (&gongdoc_models.Node{Name: "Class diagrams"}).Stage(gongdocStage)
	rootOfClassdiagramsNode.IsExpanded = true
	gongdocTree.RootNodes = append(gongdocTree.RootNodes, rootOfClassdiagramsNode)

	// add add button
	addButton := (&gongdoc_models.Button{
		Name: diagramPackage.Name + " " + string(BUTTON_add),
		Icon: string(BUTTON_add)}).Stage(gongdocStage)
	_ = addButton
	rootOfClassdiagramsNode.Buttons = append(rootOfClassdiagramsNode.Buttons, addButton)
	addButton.Impl = NewButtonImplRootOfClassdiagrams(
		diagramPackage,
		rootOfClassdiagramsNode,
		treeOfGongObjects,
		BUTTON_add,
	)

	// add one node for each diagram
	for classdiagram := range *gongdoc_models.GetGongstructInstancesSet[gongdoc_models.Classdiagram](gongdocStage) {
		classdiagramNode := NewClassdiagramNode(classdiagram, diagramPackage, rootOfClassdiagramsNode, treeOfGongObjects)
		rootOfClassdiagramsNode.Children = append(rootOfClassdiagramsNode.Children, classdiagramNode)
	}

	computeNodeConfs(gongdocStage,
		rootOfClassdiagramsNode,
		diagramPackage,
		treeOfGongObjects)
}
