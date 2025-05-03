package models

import (
	"slices"

	gong "github.com/fullstack-lang/gong/go/models"
)

type GongEnumValueShape struct {
	// Name of the enum value
	Name string

	// Identifier is the identifier of the enum value referenced by the shape in the modeled package
	//gong:ident
	Identifier string
}

const HeightBetween2AttributeShapes = 20

// AddToDiagram implements diagrammer.ElementNode.
func (classdiagram *Classdiagram) AddGongEnumValueShapeToDiagram(
	stage *Stage,
	gongEnumShape *GongEnumShape,
	gongEnum *gong.GongEnum,
	gongEnumValue *gong.GongEnumValue,
) {

	// insert the value at the correct spot in the classhape
	map_Value_rankInEnum := make(map[gong.GongEnumValue]int, 0)
	map_ValueName_Value := make(map[string]gong.GongEnumValue, 0)

	// what is the index of the field to insert in the gong struct ?
	rankInEnum := 0

	gongEnumShape.Height = gongEnumShape.Height + HeightBetween2AttributeShapes

	gongEnumValueShape := (&GongEnumValueShape{
		Name: gongEnumValue.GetName(),
		Identifier: GongstructAndFieldnameToFieldIdentifier(
			gongEnum.Name, gongEnumValue.GetName()),
	}).Stage(stage)

	for idx, gongEnumValue := range gongEnum.GongEnumValues {

		map_Value_rankInEnum[*gongEnumValue] = idx
		map_ValueName_Value[gongEnumValue.GetName()] = *gongEnumValue

		if gongEnumValue.GetName() == gongEnumValueShape.Name {
			rankInEnum = idx
		}
	}

	// compute insertionIndex (index where to insert the field to display)
	insertionIndex := 0
	for idx, field := range gongEnumShape.GongEnumValueShapes {
		value := map_ValueName_Value[IdentifierToFieldName(field.Identifier)]
		_rankInEnum := map_Value_rankInEnum[value]
		if rankInEnum > _rankInEnum {
			insertionIndex = idx + 1
		}
	}

	// append the filed to display in the right index
	if insertionIndex == len(gongEnumShape.GongEnumValueShapes) {
		gongEnumShape.GongEnumValueShapes = append(gongEnumShape.GongEnumValueShapes, gongEnumValueShape)
	} else {
		gongEnumShape.GongEnumValueShapes = append(gongEnumShape.GongEnumValueShapes[:insertionIndex+1],
			gongEnumShape.GongEnumValueShapes[insertionIndex:]...)
		gongEnumShape.GongEnumValueShapes[insertionIndex] = gongEnumValueShape
	}
}

// RemoveFromDiagram implements diagrammer.ModelElementNode.
func (classdiagram *Classdiagram) RemoveGongEnumValueShapeFromDiagram(
	stage *Stage,
	gongEnumShape *GongEnumShape,
	gongEnumValue *gong.GongEnumValue,
) {

	var gongEnumValueShape *GongEnumValueShape

	for _, _field := range gongEnumShape.GongEnumValueShapes {
		if IdentifierToFieldName(_field.Identifier) == gongEnumValue.GetName() {
			gongEnumValueShape = _field
		}
	}
	if gongEnumValueShape != nil {
		idx := slices.Index(gongEnumShape.GongEnumValueShapes, gongEnumValueShape)
		gongEnumShape.GongEnumValueShapes = slices.Delete(gongEnumShape.GongEnumValueShapes, idx, idx+1)
		gongEnumShape.Height = gongEnumShape.Height - HeightBetween2AttributeShapes
		gongEnumValueShape.Unstage(stage)
	}
}
