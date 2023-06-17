package node2gongdoc

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

func computeGongNodesConfigurations(
	stage *gongdoc_models.StageStruct,
	classdiagram *gongdoc_models.Classdiagram,
	treeOfGongObjects *gongtree_models.Tree) {

	//
	// compute maps of displayed gong objects because we need to
	// access them when analysing associations
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

	// now, compute wether each gong node to be checked / disabled
	isCheckboxDisabled := !classdiagram.IsInDrawMode
	for _, _node := range treeOfGongObjects.RootNodes {
		applyGongNodesConfRecursively(_node, isCheckboxDisabled, false)
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
	for _, _nodeForIdentifierCategory := range treeOfGongObjects.RootNodes {
		for _, _nodeForCategoryInstance := range _nodeForIdentifierCategory.Children {

			isDSLShapePresent :=
				namesOfDisplayedGongstructs[_nodeForCategoryInstance.Name] ||
					namesOfDisplayedGongenums[_nodeForCategoryInstance.Name] ||
					namesOfDisplayedGongnotes[_nodeForCategoryInstance.Name]

			_nodeForCategoryInstance.IsChecked = isDSLShapePresent
			parentNodeName := _nodeForCategoryInstance.GetName()

			for _, _nodeForProperty := range _nodeForCategoryInstance.Children {

				// checkbox of the field/link is disabled if the classdiagram
				_nodeForProperty.IsCheckboxDisabled =
					!isDSLShapePresent || // the parent shape is not present OR
						!classdiagram.IsInDrawMode // the diagram is not editable
				switch nodeImpl := _nodeForProperty.Impl.(type) {
				case *NodeImplField:
					nodeImplField := nodeImpl
					field := nodeImplField.field

					fieldUniqueName := parentNodeName + "." + field.GetName()
					if namesOfDisplayedGongfields[fieldUniqueName] {
						_nodeForProperty.IsChecked = true
					}

					switch fieldReal := field.(type) {
					case *gong_models.PointerToGongStructField:
						if ok := namesOfDisplayedGongstructs[fieldReal.GongStruct.Name]; !ok {
							// if the target type is not present, it is not possible
							// for the user to add the association to the diagram
							nodeImpl.nodeOfField.IsCheckboxDisabled = true
						}
					case *gong_models.SliceOfPointerToGongStructField:
						if ok := namesOfDisplayedGongstructs[fieldReal.GongStruct.Name]; !ok {
							// if the target type is not present, it is not possible
							// for the user to add the association to the diagram
							nodeImpl.nodeOfField.IsCheckboxDisabled = true
						}
					default:
					}
				case *NodeImplLink:
					nodeImplGongLink := nodeImpl
					gongLink := nodeImplGongLink.gongLink

					// first case is when the gong link points to a shape
					if gongLink.Recv == "" {

						fieldUniqueName := parentNodeName + "." + gongLink.GetName()
						if namesOfDisplayedGongnoteLinks[fieldUniqueName] {
							_nodeForProperty.IsChecked = true
						}

						gongStructShapeWithThisNameIsPresent := namesOfDisplayedGongstructs[gongLink.Name]
						gongEnumShapeWithThisNameIsPresent := namesOfDisplayedGongenums[gongLink.Name]

						if !gongStructShapeWithThisNameIsPresent && !gongEnumShapeWithThisNameIsPresent {

							// no corresponding gong struct shape, therefore, disable the node
							_nodeForProperty.IsCheckboxDisabled = true
						}
					} else // the other case (Recv != "") is when the gonglink points to a link
					{
						fieldName := gongLink.Recv + "." + gongLink.Name

						fieldUniqueName := parentNodeName + "." + fieldName
						if namesOfDisplayedGongnoteLinks[fieldUniqueName] {
							_nodeForProperty.IsChecked = true
						}

						if ok := namesOfDisplayedGongfields[fieldName]; !ok {
							_nodeForProperty.IsCheckboxDisabled = true
						}
					}
				case *NodeImplEnumValue:
					gongEnumValue := nodeImpl.gongEnumValue

					if namesOfDisplayedGongenumValues[gongEnumValue.Name] {
						_nodeForProperty.IsChecked = true
					}
				default:
					_ = nodeImpl
					log.Panic("No known implementation")
				}
			}
		}
	}
}
