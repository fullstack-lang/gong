package models

import (
	"slices"

	gong "github.com/fullstack-lang/gong/go/models"
)

type GongEnumValueShape struct {
	// Name of the enum value
	Name string

	//gong:meta
	IdentifierMeta any
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
		Name:           gongEnumValue.GetName(),
		IdentifierMeta: GongEnumValueToIdentifierMeta(gongEnumValue.GetName()),
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
	for idx, gongEnumValueShape := range gongEnumShape.GongEnumValueShapes {
		value := map_ValueName_Value[GongEnumValueShapeIdentifierMetaToValueName(gongEnumValueShape.IdentifierMeta)]
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

	for _, _gongEnumValueShape := range gongEnumShape.GongEnumValueShapes {
		if GongEnumValueShapeIdentifierMetaToValueName(_gongEnumValueShape.IdentifierMeta) == gongEnumValue.GetName() {
			gongEnumValueShape = _gongEnumValueShape
		}
	}
	if gongEnumValueShape != nil {
		idx := slices.Index(gongEnumShape.GongEnumValueShapes, gongEnumValueShape)
		gongEnumShape.GongEnumValueShapes = slices.Delete(gongEnumShape.GongEnumValueShapes, idx, idx+1)
		gongEnumShape.Height = gongEnumShape.Height - HeightBetween2AttributeShapes
		gongEnumValueShape.Unstage(stage)
	}
}
