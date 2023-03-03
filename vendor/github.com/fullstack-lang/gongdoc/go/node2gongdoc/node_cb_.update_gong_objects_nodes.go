package node2gongdoc

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

func (nodeCb *NodeCB) updateGongObjectsNodes(stage *gongdoc_models.StageStruct, classdiagram *gongdoc_models.Classdiagram) {

	listOfGongStructShapes := make(map[string]bool)
	for _, gongStructShape := range classdiagram.GongStructShapes {
		gongStructName := gongdoc_models.IdentifierToGongObjectName(gongStructShape.Identifier)
		listOfGongStructShapes[gongStructName] = true
	}
	listOfGongEnumShapes := make(map[string]bool)
	for _, gongEnumShape := range classdiagram.GongEnumShapes {
		gongStructName := gongdoc_models.IdentifierToGongObjectName(gongEnumShape.Identifier)
		listOfGongEnumShapes[gongStructName] = true
	}

	listOfGongFields := make(map[string]bool)
	for _, gongStructShape := range classdiagram.GongStructShapes {

		for _, gongFieldShape := range gongStructShape.Links {
			gongFieldName := gongdoc_models.IdentifierToGongObjectName(gongFieldShape.Identifier)
			listOfGongFields[gongFieldName] = true
		}
	}

	// disable some nodes
	for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct]() {
		for _, gongField := range gongStruct.Fields {
			switch fieldReal := gongField.(type) {
			case *gong_models.PointerToGongStructField:
				impl := GetNodeBackPointer(fieldReal)
				if ok := listOfGongStructShapes[fieldReal.GongStruct.Name]; !ok {
					impl.SetHasToBeDisabledValue(true)
				}
			case *gong_models.SliceOfPointerToGongStructField:
				impl := GetNodeBackPointer(fieldReal)
				if ok := listOfGongStructShapes[fieldReal.GongStruct.Name]; !ok {
					impl.SetHasToBeDisabledValue(true)
				}
			default:
			}
		}
	}

	for gongNote := range *gong_models.GetGongstructInstancesSet[gong_models.GongNote]() {
		for _, gongLink := range gongNote.Links {
			impl := GetNodeBackPointer(gongLink)
			if gongLink.Recv == "" {
				gongStructShapeWithThisNameIsPresent := listOfGongStructShapes[gongLink.Name]
				gongEnumShapeWithThisNameIsPresent := listOfGongEnumShapes[gongLink.Name]

				if !gongStructShapeWithThisNameIsPresent && !gongEnumShapeWithThisNameIsPresent {

					// no corresponding gong struct shape, therefore, disable the node
					impl.SetHasToBeDisabledValue(true)
				}
			} else {
				fieldName := gongLink.Recv + "." + gongLink.Name
				if ok := listOfGongFields[fieldName]; !ok {
					impl.SetHasToBeDisabledValue(true)
				}
			}
		}
	}

	// parse the tree of diagram elements and
	// get the corresponding gong object
	// and from the gong object, go to its node implementation and
	// set up that it has a diagram element and that is has to be checked
	for _, gongStructShape := range classdiagram.GongStructShapes {

		// get the gong struct related to the gong struct, and set t
		gongStructName := gongdoc_models.IdentifierToGongObjectName(gongStructShape.Identifier)
		gongStruct := (*gong_models.GetGongstructInstancesMap[gong_models.GongStruct]())[gongStructName]
		gongStructImpl := GetNodeBackPointer(gongStruct)
		gongStructImpl.SetHasToBeCheckedValue(true)

		for _, field := range gongStructShape.Fields {
			_ = field
			fieldName := gongdoc_models.IdentifierToFieldName(field.Identifier)

			// range over gongStruct fields (to be redone)
			for _, gongField := range gongStruct.Fields {
				if gongField.GetName() == fieldName {
					switch fieldReal := gongField.(type) {
					case *gong_models.GongBasicField:
						GetNodeBackPointer(fieldReal).SetHasToBeCheckedValue(true)
					case *gong_models.GongTimeField:
						GetNodeBackPointer(fieldReal).SetHasToBeCheckedValue(true)
					default:
					}
					continue
				}
			}
		}
		for _, link := range gongStructShape.Links {
			_ = link
			linkName := gongdoc_models.IdentifierToFieldName(link.Identifier)

			// range over gongStruct fields (to be redone)
			for _, gongField := range gongStruct.Fields {
				if gongField.GetName() == linkName {
					switch fieldReal := gongField.(type) {
					case *gong_models.PointerToGongStructField:
						impl := GetNodeBackPointer(fieldReal)
						impl.SetHasToBeCheckedValue(true)
					case *gong_models.SliceOfPointerToGongStructField:
						impl := GetNodeBackPointer(fieldReal)
						impl.SetHasToBeCheckedValue(true)
					default:
					}
				}
				// compute if the node for pointer or slice of pointer has to be disabled
			}
		}
	}
	for _, gongEnumShape := range classdiagram.GongEnumShapes {

		// get the gong struct related to the gong struct, and set t
		gongEnumName := gongdoc_models.IdentifierToGongObjectName(gongEnumShape.Identifier)
		gongEnum := (*gong_models.GetGongstructInstancesMap[gong_models.GongEnum]())[gongEnumName]
		gongEnumImpl := GetNodeBackPointer(gongEnum)
		gongEnumImpl.SetHasToBeCheckedValue(true)

		for _, gongEnumValueEntry := range gongEnumShape.GongEnumValueEntrys {
			_ = gongEnumValueEntry
			gongEnumValueName := gongdoc_models.IdentifierToFieldName(gongEnumValueEntry.Identifier)

			// range over gongStruct fields (to be redone)
			for _, gongEnumValue := range gongEnum.GongEnumValues {
				if gongEnumValue.GetName() == gongEnumValueName {
					GetNodeBackPointer(gongEnumValue).SetHasToBeCheckedValue(true)
				}
				continue
			}
		}
	}
	for _, noteShape := range classdiagram.NoteShapes {

		// get the gong struct related to the gong struct, and set t
		noteShapeName := gongdoc_models.IdentifierToGongObjectName(noteShape.Identifier)
		gongNote := (*gong_models.GetGongstructInstancesMap[gong_models.GongNote]())[noteShapeName]
		gongNodeImpl := GetNodeBackPointer(gongNote)
		gongNodeImpl.SetHasToBeCheckedValue(true)

		for _, nodeShapeLink := range noteShape.NoteShapeLinks {

			switch nodeShapeLink.Type {
			case gongdoc_models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE:
				targetShapeGongStructName := gongdoc_models.IdentifierToGongObjectName(nodeShapeLink.Identifier)
				// range over gongStruct fields (to be redone)
				for _, link := range gongNote.Links {
					if link.GetName() == targetShapeGongStructName {
						impl := GetNodeBackPointer(link)
						impl.SetHasToBeCheckedValue(true)
					}
					continue
				}
			case gongdoc_models.NOTE_SHAPE_LINK_TO_GONG_FIELD:
				receiverName, fieldName :=
					gongdoc_models.IdentifierToReceiverAndFieldName(nodeShapeLink.Identifier)
					// range over gongStruct fields (to be redone)
				for _, link := range gongNote.Links {
					if link.GetName() == fieldName && link.Recv == receiverName {
						impl := GetNodeBackPointer(link)
						impl.SetHasToBeCheckedValue(true)
					}
					continue
				}
			}

		}
	}

	// now, parse again the tree of nodes and set the IsChecked according to the
	// the node implementation
	for _, rootNode := range nodeCb.treeOfGongObjects.RootNodes {
		for _, node := range rootNode.Children {
			node.IsCheckboxDisabled = node.Impl.HasToBeDisabled()
			node.IsChecked = node.Impl.HasToBeChecked()

			for _, node := range node.Children {
				node.IsCheckboxDisabled = node.Impl.HasToBeDisabled()
				node.IsChecked = node.Impl.HasToBeChecked()
			}
		}
	}
}
