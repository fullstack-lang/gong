package node2gongdoc

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

// computeGongNodesConfiguration set up all gong nodes according to the classdiagram
func (nodeCb *NodeCB) computeGongNodesConfiguration(stage *gongdoc_models.StageStruct, classdiagram *gongdoc_models.Classdiagram) {

	//
	// compute maps of displayed gong objects
	//
	namesOfDisplayedGongstructs := make(map[string]bool)
	namesOfDisplayedGongfields := make(map[string]bool)
	namesOfDisplayedGongenums := make(map[string]bool)
	namesOfDisplayedGongenumValues := make(map[string]bool)
	namesOfDisplayedGongnotes := make(map[string]bool)
	namesOfDisplayedGongnoteLinks := make(map[string]bool)

	for _, gongStructShape := range classdiagram.GongStructShapes {
		gongStructName := gongdoc_models.IdentifierToGongObjectName(gongStructShape.Identifier)
		namesOfDisplayedGongstructs[gongStructName] = true

		for _, gongFieldShape := range gongStructShape.Links {
			gongFieldName := gongdoc_models.IdentifierToGongObjectName(gongFieldShape.Identifier)
			namesOfDisplayedGongfields[gongFieldName] = true
		}
		for _, gongField := range gongStructShape.Fields {
			gongFieldName := gongdoc_models.IdentifierToGongObjectName(gongField.Identifier)
			namesOfDisplayedGongfields[gongFieldName] = true
		}
	}

	for _, gongEnumShape := range classdiagram.GongEnumShapes {
		gongStructName := gongdoc_models.IdentifierToGongObjectName(gongEnumShape.Identifier)
		namesOfDisplayedGongenums[gongStructName] = true

		for _, gongEnumValueShape := range gongEnumShape.GongEnumValueEntrys {
			namesOfDisplayedGongenumValues[gongEnumValueShape.Name] = true
		}
	}

	for _, noteShape := range classdiagram.NoteShapes {
		gongNoteName := gongdoc_models.IdentifierToGongObjectName(noteShape.Identifier)
		namesOfDisplayedGongnotes[gongNoteName] = true

		for _, noteShapeLink := range noteShape.NoteShapeLinks {
			gongShapeLinkName := gongdoc_models.IdentifierToGongObjectName(noteShapeLink.Identifier)

			// in this case, one need to add the gongNoteName
			namesOfDisplayedGongnoteLinks[gongNoteName+"."+gongShapeLinkName] = true
		}
	}

	// parses all DSL nodes (that is gong identifiers).
	//
	// For each node
	//
	// 1. if a field / link nodes, disable when
	// - the DSL is not displayed
	// - the gongstruct type of the field is not displayed
	//
	// 2. check is DSL shape is present in the diagram
	//
	for _, _node := range nodeCb.treeOfGongObjects.RootNodes {
		for _, _node := range _node.Children {

			isDSLShapePresent :=
				namesOfDisplayedGongstructs[_node.Name] ||
					namesOfDisplayedGongenums[_node.Name] ||
					namesOfDisplayedGongnotes[_node.Name]

			_node.IsChecked = isDSLShapePresent
			parentNodeName := _node.GetName()

			for _, _node := range _node.Children {

				// checkbox of the field/link is disabled if the classdiagram
				_node.IsCheckboxDisabled =
					!isDSLShapePresent || // the parent shape is not present OR
						!classdiagram.IsInDrawMode // the diagram is not editable
				switch nodeImpl := _node.Impl.(type) {
				case *FieldImpl:
					gongField := nodeImpl.field

					fieldUniqueName := parentNodeName + "." + gongField.GetName()
					if namesOfDisplayedGongfields[fieldUniqueName] {
						_node.IsChecked = true
					}

					switch fieldReal := gongField.(type) {
					case *gong_models.PointerToGongStructField:
						if ok := namesOfDisplayedGongstructs[fieldReal.GongStruct.Name]; !ok {
							nodeImpl.DisableNodeCheckbox()
						}
					case *gong_models.SliceOfPointerToGongStructField:
						if ok := namesOfDisplayedGongstructs[fieldReal.GongStruct.Name]; !ok {
							nodeImpl.DisableNodeCheckbox()
						}
					default:
					}
				case *GongLinkImpl:
					gongLink := nodeImpl.gongLink

					// first case is when the gong link points to a shape
					if gongLink.Recv == "" {

						fieldUniqueName := parentNodeName + "." + gongLink.GetName()
						if namesOfDisplayedGongnoteLinks[fieldUniqueName] {
							_node.IsChecked = true
						}

						gongStructShapeWithThisNameIsPresent := namesOfDisplayedGongstructs[gongLink.Name]
						gongEnumShapeWithThisNameIsPresent := namesOfDisplayedGongenums[gongLink.Name]

						if !gongStructShapeWithThisNameIsPresent && !gongEnumShapeWithThisNameIsPresent {

							// no corresponding gong struct shape, therefore, disable the node
							_node.IsCheckboxDisabled = true
						}
					} else // the other case (Recv != "") is when the gonglink points to a link
					{
						fieldName := gongLink.Recv + "." + gongLink.Name

						fieldUniqueName := parentNodeName + "." + fieldName
						if namesOfDisplayedGongnoteLinks[fieldUniqueName] {
							_node.IsChecked = true
						}

						if ok := namesOfDisplayedGongfields[fieldName]; !ok {
							_node.IsCheckboxDisabled = true
						}
					}
				case *GongEnumValueImpl:
					gongEnumValue := nodeImpl.gongEnumValue

					if namesOfDisplayedGongenumValues[gongEnumValue.Name] {
						_node.IsChecked = true
					}

				default:
					_ = nodeImpl
					log.Panic("No known implementation")
				}
			}
		}
	}
}
