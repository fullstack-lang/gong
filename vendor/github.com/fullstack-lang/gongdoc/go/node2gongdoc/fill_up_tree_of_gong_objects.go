package node2gongdoc

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

func FillUpTreeOfGongObjects(
	gongdocStage *gongdoc_models.StageStruct,
	gongtreeStage *gongtree_models.StageStruct,
	diagramPackage *gongdoc_models.DiagramPackage,
) (treeOfGongObjects *gongtree_models.Tree) {

	// set up the gongTree to display elements
	treeOfGongObjects = (&gongtree_models.Tree{Name: "gong"}).Stage(gongtreeStage)

	gongstructRootNode := (&gongtree_models.Node{Name: "gongstructs"}).Stage(gongtreeStage)
	gongstructRootNode.IsExpanded = true
	treeOfGongObjects.RootNodes = append(treeOfGongObjects.RootNodes, gongstructRootNode)
	for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](diagramPackage.ModelPkg.GetStage()) {

		nodeGongstruct := (&gongtree_models.Node{Name: gongStruct.Name}).Stage(gongtreeStage)
		gongstructRootNode.Children = append(gongstructRootNode.Children, nodeGongstruct)

		nodeGongstruct.HasCheckboxButton = true
		nodeGongstruct.IsExpanded = false

		// set up inversion control
		nodeGongstruct.Impl = NewNodeImplGongstruct(
			gongStruct, NewNodeImplGongObjectAbstract(diagramPackage, treeOfGongObjects))

		for _, field := range gongStruct.Fields {
			nodeGongField := (&gongtree_models.Node{Name: field.GetName()}).Stage(gongtreeStage)
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeGongField)

			nodeGongField.HasCheckboxButton = true
			nodeGongField.Impl = NewNodeImplField(
				gongStruct,
				field,
				nodeGongstruct,
				nodeGongField,
				NewNodeImplGongObjectAbstract(diagramPackage, treeOfGongObjects))
		}
	}

	gongenumRootNode := (&gongtree_models.Node{Name: "gongenums"}).Stage(gongtreeStage)
	gongenumRootNode.IsExpanded = true
	treeOfGongObjects.RootNodes = append(treeOfGongObjects.RootNodes, gongenumRootNode)
	for gongEnum := range *gong_models.GetGongstructInstancesSet[gong_models.GongEnum](diagramPackage.ModelPkg.GetStage()) {

		nodeGongEnum := (&gongtree_models.Node{Name: gongEnum.Name}).Stage(gongtreeStage)
		nodeGongEnum.HasCheckboxButton = true
		nodeGongEnum.IsExpanded = false

		// append to tree
		gongenumRootNode.Children = append(gongenumRootNode.Children, nodeGongEnum)
		nodeGongEnum.Impl = NewNodeImplGongEnum(gongEnum,
			NewNodeImplGongObjectAbstract(diagramPackage, treeOfGongObjects))

		for _, gongEnumValue := range gongEnum.GongEnumValues {
			nodeGongEnumValue := (&gongtree_models.Node{Name: gongEnumValue.GetName()}).Stage(gongtreeStage)

			nodeGongEnumValue.HasCheckboxButton = true
			nodeGongEnum.Children = append(nodeGongEnum.Children, nodeGongEnumValue)

			nodeGongEnumValue.Impl = NewNodeImplEnumValue(gongEnum, gongEnumValue, nodeGongEnum, nodeGongEnumValue,
				NewNodeImplGongObjectAbstract(diagramPackage, treeOfGongObjects))
		}
	}

	gongNotesRootNode := (&gongtree_models.Node{Name: "notes"}).Stage(gongtreeStage)
	gongNotesRootNode.IsExpanded = true
	treeOfGongObjects.RootNodes = append(treeOfGongObjects.RootNodes, gongNotesRootNode)
	for gongNote := range *gong_models.GetGongstructInstancesSet[gong_models.GongNote](diagramPackage.ModelPkg.GetStage()) {

		nodeGongNote := (&gongtree_models.Node{Name: gongNote.Name}).Stage(gongtreeStage)
		nodeGongNote.HasCheckboxButton = true
		nodeGongNote.IsExpanded = true
		gongNotesRootNode.Children = append(gongNotesRootNode.Children, nodeGongNote)

		nodeGongNote.Impl = NewNodeImplGongnote(gongNote,
			NewNodeImplGongObjectAbstract(diagramPackage, treeOfGongObjects))

		for _, gongLink := range gongNote.Links {

			gongLinkName := gongLink.Name

			if gongLink.Recv != "" {
				gongLinkName = gongLink.Recv + "." + gongLinkName
			}

			nodeGongLink := (&gongtree_models.Node{Name: gongLinkName}).Stage(gongtreeStage)
			nodeGongLink.HasCheckboxButton = true

			// append to tree
			nodeGongNote.Children = append(nodeGongNote.Children, nodeGongLink)

			nodeGongLink.Impl = NewNodeImplLink(gongNote, gongLink, nodeGongNote, nodeGongLink,
				NewNodeImplGongObjectAbstract(diagramPackage, treeOfGongObjects))
		}
	}
	return
}
