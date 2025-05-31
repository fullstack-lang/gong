package models

import (
	"log"
	"slices"
	"strings"

	gong "github.com/fullstack-lang/gong/go/models"
)

// AttributeShape
type AttributeShape struct {
	Name string

	//gong:ident
	Identifier string

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
	gongStructShape.Height = gongStructShape.Height - HeightBetween2AttributeShapes
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

		concreteField.Identifier = fieldIdentifier

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

		concreteField.Structname = IdentifierToGongObjectName(gongStructShape.Identifier)
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
			gongField := map_Name_Field[IdentifierToFieldName(field.Identifier)]
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

// AddToDiagram implements diagrammer.ElementNode.
func (classdiagram *Classdiagram) AddLinkFieldShape(
	stage *Stage,
	gongStage *gong.Stage,
	gongStruct *gong.GongStruct,
	field gong.FieldInterface,
	gongStructShape *GongStructShape) {
	diagramPackage := getTheDiagramPackage(stage)

	switch field.(type) {
	case *gong.PointerToGongStructField, *gong.SliceOfPointerToGongStructField:

		var targetStructName string
		var sourceMultiplicity MultiplicityType
		var targetMultiplicity MultiplicityType

		switch realField := field.(type) {
		case *gong.PointerToGongStructField:
			targetStructName = realField.GongStruct.Name
			sourceMultiplicity = MANY
			targetMultiplicity = ZERO_ONE
		case *gong.SliceOfPointerToGongStructField:
			targetStructName = realField.GongStruct.Name
			sourceMultiplicity = MANY
			targetMultiplicity = MANY
		}
		targetSourceGongStructShape := false
		var targetGongStructShape *GongStructShape
		for _, _gongstructshape := range diagramPackage.SelectedClassdiagram.GongStructShapes {

			// strange behavior when the gongstructshape is remove within the loop
			if IdentifierToGongObjectName(_gongstructshape.Identifier) == targetStructName && !targetSourceGongStructShape {
				targetSourceGongStructShape = true
				targetGongStructShape = _gongstructshape
			}
		}
		if !targetSourceGongStructShape {
			log.Panicf("GongStructShape %s of field not present ", targetStructName)
		}
		_ = targetGongStructShape

		link := new(LinkShape).Stage(stage)
		link.Name = field.GetName()
		link.SourceMultiplicity = sourceMultiplicity
		link.SourceMultiplicityOffsetX = 0
		link.SourceMultiplicityOffsetY = 0

		link.TargetMultiplicity = targetMultiplicity
		link.TargetMultiplicityOffsetX = 0
		link.TargetMultiplicityOffsetY = 0

		link.FieldOffsetX = 0
		link.FieldOffsetY = 0

		link.Identifier =
			GongstructAndFieldnameToFieldIdentifier(gongStruct.Name, field.GetName())
		link.Fieldtypename = GongStructNameToIdentifier(targetStructName)

		gongStructShape.LinkShapes = append(gongStructShape.LinkShapes, link)

		link.X = (gongStructShape.X+targetGongStructShape.X)/2.0 +
			gongStructShape.Width*1.5
		link.Y = (gongStructShape.Y+targetGongStructShape.Y)/2.0 +
			gongStructShape.Height/2.0

		link.StartOrientation = ORIENTATION_HORIZONTAL
		link.StartRatio = 0.5
		link.EndOrientation = ORIENTATION_HORIZONTAL
		link.EndRatio = 0.5
		link.CornerOffsetRatio = 1.38
	}
}

// moveStructLiteralToType converts "Type.Field{}" to "Type{}.Field"
func moveStructLiteralToType(input string) string {

	// Find the last dot to separate type from field
	lastDotIndex := strings.LastIndex(input, ".")
	if lastDotIndex == -1 {
		// No dot found, return as is
		return input
	}

	// Split into type part and field part
	typePart := input[:lastDotIndex]    // Everything before the last dot
	fieldPart := input[lastDotIndex+1:] // Everything after the last dot

	// Rebuild as "Type{}.Field" and add quotes back
	return typePart + "{}." + fieldPart
}
