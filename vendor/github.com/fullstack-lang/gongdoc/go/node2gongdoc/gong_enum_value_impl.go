package node2gongdoc

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

type GongEnumValueImpl struct {
	gongEnumValue *gong_models.GongEnumValue
	NodeImpl
}

func (enumValueImpl *GongEnumValueImpl) OnAfterUpdate(
	stage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {

	// find classdiagram
	classdiagram := enumValueImpl.nodeCb.GetSelectedClassdiagram()

	// find the parent node to find the gongenum to find the gongstructshape
	// the node is field, one needs to find the gongenum that contains it
	// get the parent node
	parentNode := enumValueImpl.nodeCb.map_Children_Parent[stagedNode]

	gongEnumImpl := parentNode.Impl.(*GongEnumImpl)
	gongEnum := gongEnumImpl.gongEnum

	// find the classhape in the classdiagram
	foundGongEnumShape := false
	var gongEnumShape *gongdoc_models.GongEnumShape
	for _, _gongstructshape := range classdiagram.GongEnumShapes {
		if gongdoc_models.IdentifierToGongObjectName(_gongstructshape.Identifier) == gongEnum.Name && !foundGongEnumShape {
			gongEnumShape = _gongstructshape
		}
	}
	_ = gongEnumShape

	// insert the value at the correct spot in the classhape
	map_Value_rankInEnum := make(map[gong_models.GongEnumValue]int, 0)
	map_ValueName_Value := make(map[string]gong_models.GongEnumValue, 0)

	// what is the index of the field to insert in the gong struct ?
	rankkInEnum := 0

	if !stagedNode.IsChecked && frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		stage.Checkout()

		gongEnumShape.Heigth = gongEnumShape.Heigth + 15

		var gongEnumValueEntry gongdoc_models.GongEnumValueEntry
		gongEnumValueEntry.Name = stagedNode.Name
		gongEnumValueEntry.Identifier = gongdoc_models.GongstructAndFieldnameToFieldIdentifier(gongEnum.Name, stagedNode.Name)

		for idx, gongEnum := range gongEnum.GongEnumValues {

			map_Value_rankInEnum[*gongEnum] = idx
			map_ValueName_Value[gongEnum.GetName()] = *gongEnum

			if gongEnum.GetName() == gongEnumValueEntry.Name {
				rankkInEnum = idx
			}
		}

		// compute insertionIndex (index where to insert the field to display)
		insertionIndex := 0
		for idx, field := range gongEnumShape.GongEnumValueEntrys {
			value := map_ValueName_Value[gongdoc_models.IdentifierToFieldName(field.Identifier)]
			_rankInEnum := map_Value_rankInEnum[value]
			if rankkInEnum > _rankInEnum {
				insertionIndex = idx + 1
			}
		}

		// append the filed to display in the right index
		if insertionIndex == len(gongEnumShape.GongEnumValueEntrys) {
			gongEnumShape.GongEnumValueEntrys = append(gongEnumShape.GongEnumValueEntrys, &gongEnumValueEntry)
		} else {
			gongEnumShape.GongEnumValueEntrys = append(gongEnumShape.GongEnumValueEntrys[:insertionIndex+1],
				gongEnumShape.GongEnumValueEntrys[insertionIndex:]...)
			gongEnumShape.GongEnumValueEntrys[insertionIndex] = &gongEnumValueEntry
		}
		gongEnumValueEntry.Stage()
		gongdoc_models.Stage.Commit()

	}
	if stagedNode.IsChecked && !frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		stage.Checkout()

		{
			var gongEnumValueEntry *gongdoc_models.GongEnumValueEntry

			for _, _field := range gongEnumShape.GongEnumValueEntrys {
				if gongdoc_models.IdentifierToFieldName(_field.Identifier) == stagedNode.Name {
					gongEnumValueEntry = _field
				}
			}
			if gongEnumValueEntry != nil {
				gongEnumShape.GongEnumValueEntrys = remove(gongEnumShape.GongEnumValueEntrys, gongEnumValueEntry)
				gongEnumShape.Heigth = gongEnumShape.Heigth - 15
				gongEnumValueEntry.Unstage()
			}
		}

		gongdoc_models.Stage.Commit()
	}
}

func (EnumValueImpl *GongEnumValueImpl) OnAfterDelete(
	stage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {
}
