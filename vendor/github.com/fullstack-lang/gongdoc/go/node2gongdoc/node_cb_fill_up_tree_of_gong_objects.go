package node2gongdoc

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

func (nodeCb *NodeCB) FillUpTreeOfGongObjects() {

	// set up the gongTree to display elements
	gongTree := (&gongdoc_models.Tree{Name: "gong"}).Stage(nodeCb.diagramPackage.Stage_)
	nodeCb.treeOfGongObjects = gongTree

	gongstructRootNode := (&gongdoc_models.Node{Name: "gongstructs"}).Stage(nodeCb.diagramPackage.Stage_)
	gongstructRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongstructRootNode)
	for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct]() {

		nodeGongstruct := (&gongdoc_models.Node{Name: gongStruct.Name}).Stage(nodeCb.diagramPackage.Stage_)

		nodeGongstruct.HasCheckboxButton = true
		nodeGongstruct.IsExpanded = false

		// set up the back pointer from the shape to the node
		gongStructImpl := new(GongStructImpl)
		gongStructImpl.node = nodeGongstruct
		gongStructImpl.gongStruct = gongStruct
		gongStructImpl.nodeCb = nodeCb
		nodeGongstruct.Impl = gongStructImpl

		// append to the tree
		gongstructRootNode.Children = append(gongstructRootNode.Children, nodeGongstruct)

		SetNodeBackPointer(gongStruct, gongStructImpl)

		for _, field := range gongStruct.Fields {
			nodeGongField := (&gongdoc_models.Node{Name: field.GetName()}).Stage(nodeCb.diagramPackage.Stage_)

			nodeGongField.HasCheckboxButton = true

			fieldImpl := new(FieldImpl)
			fieldImpl.node = nodeGongstruct
			fieldImpl.field = field
			fieldImpl.nodeCb = nodeCb
			nodeGongField.Impl = fieldImpl

			// append to tree
			nodeGongstruct.Children = append(nodeGongstruct.Children, nodeGongField)

			switch fieldReal := field.(type) {
			case *gong_models.GongBasicField:
				SetNodeBackPointer(fieldReal, fieldImpl)
			case *gong_models.GongTimeField:
				SetNodeBackPointer(fieldReal, fieldImpl)
			case *gong_models.PointerToGongStructField:
				SetNodeBackPointer(fieldReal, fieldImpl)
			case *gong_models.SliceOfPointerToGongStructField:
				SetNodeBackPointer(fieldReal, fieldImpl)
			default:
			}
		}
	}

	gongenumRootNode := (&gongdoc_models.Node{Name: "gongenums"}).Stage(nodeCb.diagramPackage.Stage_)
	gongenumRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongenumRootNode)
	for gongEnum := range *gong_models.GetGongstructInstancesSet[gong_models.GongEnum]() {

		nodeGongEnum := (&gongdoc_models.Node{Name: gongEnum.Name}).Stage(nodeCb.diagramPackage.Stage_)
		nodeGongEnum.HasCheckboxButton = true
		nodeGongEnum.IsExpanded = false

		gongEnumImpl := new(GongEnumImpl)
		gongEnumImpl.node = nodeGongEnum
		gongEnumImpl.gongEnum = gongEnum
		gongEnumImpl.nodeCb = nodeCb
		nodeGongEnum.Impl = gongEnumImpl

		// append to tree
		gongenumRootNode.Children = append(gongenumRootNode.Children, nodeGongEnum)
		SetNodeBackPointer(gongEnum, gongEnumImpl)

		for _, gongEnumValue := range gongEnum.GongEnumValues {
			nodeGongEnumValue := (&gongdoc_models.Node{Name: gongEnumValue.GetName()}).Stage(nodeCb.diagramPackage.Stage_)

			nodeGongEnumValue.HasCheckboxButton = true

			gongEnumValueImpl := new(GongEnumValueImpl)
			gongEnumValueImpl.node = nodeGongEnum
			gongEnumValueImpl.gongEnumValue = gongEnumValue
			gongEnumValueImpl.nodeCb = nodeCb
			nodeGongEnumValue.Impl = gongEnumValueImpl

			// append to tree
			nodeGongEnum.Children = append(nodeGongEnum.Children, nodeGongEnumValue)
			SetNodeBackPointer(gongEnumValue, gongEnumValueImpl)
		}
	}

	gongNotesRootNode := (&gongdoc_models.Node{Name: "notes"}).Stage(nodeCb.diagramPackage.Stage_)
	gongNotesRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongNotesRootNode)
	for gongNote := range *gong_models.GetGongstructInstancesSet[gong_models.GongNote]() {

		node := (&gongdoc_models.Node{Name: gongNote.Name}).Stage(nodeCb.diagramPackage.Stage_)
		node.HasCheckboxButton = true

		node.IsExpanded = true

		gongNoteImpl := new(GongNoteImpl)
		gongNoteImpl.node = node
		gongNoteImpl.gongNote = gongNote
		gongNoteImpl.nodeCb = nodeCb
		node.Impl = gongNoteImpl

		// append to tree
		gongNotesRootNode.Children = append(gongNotesRootNode.Children, node)

		SetNodeBackPointer(gongNote, gongNoteImpl)

		for _, gongLink := range gongNote.Links {

			gongLinkName := gongLink.Name

			if gongLink.Recv != "" {
				gongLinkName = gongLink.Recv + "." + gongLinkName
			}

			nodeGongLink := (&gongdoc_models.Node{Name: gongLinkName}).Stage(nodeCb.diagramPackage.Stage_)
			nodeGongLink.HasCheckboxButton = true

			gongLinkImpl := new(GongLinkImpl)
			gongLinkImpl.node = node
			gongLinkImpl.gongLink = gongLink
			gongLinkImpl.nodeCb = nodeCb
			nodeGongLink.Impl = gongLinkImpl

			// append to tree
			node.Children = append(node.Children, nodeGongLink)

			SetNodeBackPointer(gongLink, gongLinkImpl)
		}
	}

	// generate the map to navigate from children to parents
	fieldName := gongdoc_models.GetAssociationName[gongdoc_models.Node]().Children[0].Name
	nodeCb.map_Children_Parent = gongdoc_models.GetSliceOfPointersReverseMap[gongdoc_models.Node, gongdoc_models.Node](fieldName)
}
