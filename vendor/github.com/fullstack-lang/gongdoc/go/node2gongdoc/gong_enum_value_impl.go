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

	// find the parent node to find the gongenum to find the classshape
	// the node is field, one needs to find the gongenum that contains it
	// get the parent node
	parentNode := enumValueImpl.nodeCb.map_Children_Parent[stagedNode]

	gongEnumImpl := parentNode.Impl.(*GongEnumImpl)
	gongEnum := gongEnumImpl.gongEnum

	// find the classhape in the classdiagram
	foundClassshape := false
	var classshape *gongdoc_models.GongEnumShape
	for _, _classshape := range classdiagram.GongEnumShapes {
		if gongdoc_models.IdentifierToGongStructName(_classshape.Identifier) == gongEnum.Name && !foundClassshape {
			classshape = _classshape
		}
	}
	_ = classshape

	// insert the value at the correct spot in the classhape
	map_Value_rankInEnum := make(map[gong_models.GongEnumValue]int, 0)
	map_ValueName_Value := make(map[string]gong_models.GongEnumValue, 0)

	// what is the index of the field to insert in the gong struct ?
	rankkInEnum := 0

	if !stagedNode.IsChecked && frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		stage.Checkout()

		classshape.Heigth = classshape.Heigth + 15

		var field gongdoc_models.Field
		field.Name = stagedNode.Name
		field.Identifier = gongdoc_models.GongstructAndFieldnameToFieldIdentifier(gongEnum.Name, stagedNode.Name)

		for idx, gongEnum := range gongEnum.GongEnumValues {

			map_Value_rankInEnum[*gongEnum] = idx
			map_ValueName_Value[gongEnum.GetName()] = *gongEnum

			if gongEnum.GetName() == field.Name {
				rankkInEnum = idx
			}
		}

		// compute insertionIndex (index where to insert the field to display)
		insertionIndex := 0
		for idx, field := range classshape.Fields {
			value := map_ValueName_Value[gongdoc_models.IdentifierToFieldName(field.Identifier)]
			_rankInEnum := map_Value_rankInEnum[value]
			if rankkInEnum > _rankInEnum {
				insertionIndex = idx + 1
			}
		}

		// append the filed to display in the right index
		if insertionIndex == len(classshape.Fields) {
			classshape.Fields = append(classshape.Fields, &field)
		} else {
			classshape.Fields = append(classshape.Fields[:insertionIndex+1],
				classshape.Fields[insertionIndex:]...)
			classshape.Fields[insertionIndex] = &field
		}
		field.Stage()
		gongdoc_models.Stage.Commit()

	}
	if stagedNode.IsChecked && !frontNode.IsChecked {

		// get the latest version of the diagram before modifying it
		stage.Checkout()

		{
			var field *gongdoc_models.Field

			for _, _field := range classshape.Fields {
				if gongdoc_models.IdentifierToFieldName(_field.Identifier) == stagedNode.Name {
					field = _field
				}
			}
			if field != nil {
				classshape.Fields = remove(classshape.Fields, field)
				classshape.Heigth = classshape.Heigth - 15
				field.Unstage()
			}
		}

		gongdoc_models.Stage.Commit()
	}
}

func (EnumValueImpl *GongEnumValueImpl) OnAfterDelete(
	stage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {
}
