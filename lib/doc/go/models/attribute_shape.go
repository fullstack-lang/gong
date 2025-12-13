package models

import (
	"slices"
	"strings"

	gong "github.com/fullstack-lang/gong/go/models"
)

// AttributeShape
type AttributeShape struct {
	Name string

	// for storing the reference as a renaming target for gopls
	// for instance 'ref_models.Astruct.IntField'
	//gong:meta
	IdentifierMeta any

	FieldTypeAsString string
	Structname        string
	Fieldtypename     string
}

// RemoveAttributeFieldShape implements diagrammer.ModelElementNode.
func (classdiagram *Classdiagram) RemoveAttributeFieldShape(
	stage *Stage,
	attributeShape *AttributeShape,
	gongStructShape *GongStructShape) {

	idx := slices.Index(gongStructShape.AttributeShapes, attributeShape)
	gongStructShape.AttributeShapes = slices.Delete(gongStructShape.AttributeShapes, idx, idx+1)
	gongStructShape.Height = gongStructShape.Height - HeightBetween2AttributeShapes
	attributeShape.Unstage(stage)

}

// RemoveLinkFieldShape implements diagrammer.ModelElementNode.
func (classdiagram *Classdiagram) RemoveLinkFieldShape(
	stage *Stage,
	linkShape *LinkShape,
	gongStructShape *GongStructShape) {

	idx := slices.Index(gongStructShape.LinkShapes, linkShape)
	gongStructShape.LinkShapes = slices.Delete(gongStructShape.LinkShapes, idx, idx+1)
	linkShape.Unstage(stage)

}

// AddToDiagram implements diagrammer.ElementNode.
func (classdiagram *Classdiagram) AddAttributeFieldShape(
	stage *Stage,
	gongStage *gong.Stage,
	gongStruct *gong.GongStruct,
	field gong.FieldInterface,
	gongStructShape *GongStructShape) {

	switch field.(type) {
	case *gong.GongBasicField, *gong.GongTimeField:

		// concrete in the sense of UML concrete syntax
		var concreteField AttributeShape
		concreteField.Name = field.GetName()

		fieldIdentifier := GongstructAndFieldnameToFieldIdentifier(
			gongStruct.Name, field.GetName())

		// turn ref_models.Button.Name{} into ref_models.Button{}.Name
		concreteField.IdentifierMeta = moveStructLiteralToType(fieldIdentifier)

		switch realField := field.(type) {
		case *gong.GongBasicField:

			// get the type after the "."
			names := strings.Split(realField.DeclaredType, ".")
			fieldTypeName := names[len(names)-1]

			concreteField.Fieldtypename = fieldTypeName
		case *gong.GongTimeField:
			concreteField.Fieldtypename = "Time"
		case *gong.PointerToGongStructField:
		case *gong.SliceOfPointerToGongStructField:
		}

		concreteField.Structname = IdentifierMetaToGongStructName(gongStructShape.IdentifierMeta)
		concreteField.Stage(stage)

		gongStructShape.Height = gongStructShape.Height + HeightBetween2AttributeShapes

		// construct ordered slice of fields
		map_Field_Rank := make(map[gong.FieldInterface]int, 0)
		map_Name_Field := make(map[string]gong.FieldInterface, 0)

		// what is the index of the field to insert in the gong struct ?
		fieldRank := 0

		for idx, gongField := range gongStruct.Fields {

			map_Field_Rank[gongField] = idx
			map_Name_Field[gongField.GetName()] = gongField

			if gongField.GetName() == concreteField.Name {
				fieldRank = idx
			}
		}

		// compute insertionIndex (index where to insert the field to display)
		insertionIndex := 0
		for idx, field := range gongStructShape.AttributeShapes {
			gongField := map_Name_Field[IdentifierMetaToFieldName(field.IdentifierMeta)]
			_fieldRank := map_Field_Rank[gongField]
			if fieldRank > _fieldRank {
				insertionIndex = idx + 1
			}
		}

		// append the filed to display in the right index
		if insertionIndex == len(gongStructShape.AttributeShapes) {
			gongStructShape.AttributeShapes = append(gongStructShape.AttributeShapes, &concreteField)
		} else {
			gongStructShape.AttributeShapes = append(gongStructShape.AttributeShapes[:insertionIndex+1],
				gongStructShape.AttributeShapes[insertionIndex:]...)
			gongStructShape.AttributeShapes[insertionIndex] = &concreteField
		}

	}
}
