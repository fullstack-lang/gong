package models

import gong_models "github.com/fullstack-lang/gong/go/models"

func FillUpTreeOfIdentifiers(pkgelt *DiagramPackage, onNodeCallbackStruct *NodeCallbacksSingloton) {

	// set up the gongTree to display elements
	gongTree := (&Tree{Name: "gong", Type: TREE_OF_IDENTIFIERS}).Stage()
	gongstructRootNode := (&Node{Name: "gongstructs"}).Stage()
	gongstructRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongstructRootNode)
	onNodeCallbackStruct.GongstructsRootNode = gongstructRootNode
	for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct]() {

		node := (&Node{Name: gongStruct.Name}).Stage()
		node.Type = GONG_STRUCT
		node.Gongstruct = gongStruct
		node.HasCheckboxButton = true
		node.IsExpanded = true
		gongstructRootNode.Children = append(gongstructRootNode.Children, node)

		for _, field := range gongStruct.Fields {
			node2 := (&Node{Name: field.GetName()}).Stage()
			node2.Type = GONG_STRUCT_FIELD
			node2.Gongfield = field
			node2.HasCheckboxButton = true
			node.Children = append(node.Children, node2)
		}
	}

	gongenumRootNode := (&Node{Name: "gongenums"}).Stage()
	gongenumRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongenumRootNode)
	onNodeCallbackStruct.GongenumsRootNode = gongenumRootNode
	for gongEnum := range *gong_models.GetGongstructInstancesSet[gong_models.GongEnum]() {

		node := (&Node{Name: gongEnum.Name}).Stage()
		node.HasCheckboxButton = true
		node.IsExpanded = true
		node.Type = GONG_ENUM
		node.GongEnum = gongEnum
		gongenumRootNode.Children = append(gongenumRootNode.Children, node)

		for _, value := range gongEnum.GongEnumValues {
			node2 := (&Node{Name: value.GetName()}).Stage()
			node2.Type = GONG_ENUM_VALUE
			node2.HasCheckboxButton = true
			node.Children = append(node.Children, node2)
		}
	}

	gongNotesRootNode := (&Node{Name: "notes"}).Stage()
	gongNotesRootNode.IsExpanded = true
	gongTree.RootNodes = append(gongTree.RootNodes, gongNotesRootNode)
	onNodeCallbackStruct.GongnotesRootNode = gongNotesRootNode
	for gongNote := range *gong_models.GetGongstructInstancesSet[gong_models.GongNote]() {

		node := (&Node{Name: gongNote.Name}).Stage()
		node.HasCheckboxButton = true
		node.IsExpanded = true
		node.Type = GONG_NOTE
		gongNotesRootNode.Children = append(gongNotesRootNode.Children, node)
	}
}
