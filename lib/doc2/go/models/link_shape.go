package models

import (
	"log"
	"strings"

	gong "github.com/fullstack-lang/gong/go/models"
)

// LinkShape represent the UML LinkShape in any diagram
// uni-directional https://en.wikipedia.org/wiki/Association_(object-oriented_programming)
// More specificaly It is a 0..1 ---> 0..1
// swagger:model LinkShape
type LinkShape struct {
	Name string

	// Identifier is the identifier of the struct field referenced by the
	// UML field of the classshape in the modeled package
	//gong:ident
	Identifier string

	// for storing the reference as a renaming target for gopls
	// for instance 'ref_models.Astruct.Anarrayofb'
	//gong:meta
	IdentifierMeta any

	//gong:meta
	FieldTypeIdentifierMeta any

	// position of the field text relative to the
	// arrow end
	FieldOffsetX float64
	FieldOffsetY float64

	TargetMultiplicity MultiplicityType

	// position of the TargetMultiplicity text relative to the
	// arrow end
	TargetMultiplicityOffsetX float64
	TargetMultiplicityOffsetY float64

	SourceMultiplicity MultiplicityType

	// position of the SourceMultiplicity text relative to the
	// arrow end
	SourceMultiplicityOffsetX float64
	SourceMultiplicityOffsetY float64

	// for links with floating anchors
	// if link type is floating orthogonal ratio, from 0 to 1,
	// where the anchor starts on the edge (horizontal / vertical)
	// Vertice in the middle
	X                float64
	Y                float64
	StartOrientation OrientationType
	StartRatio       float64
	EndOrientation   OrientationType
	EndRatio         float64

	// in case StartOrientation is the same as EndOrientation,
	// there is a perpendicular line that reach the corner at
	// CornerOffsetRatio
	CornerOffsetRatio float64
}

func (classdiagram *Classdiagram) AddLinkShape(
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
			if IdentifierMetaToGongStructName(_gongstructshape.IdentifierMeta) == targetStructName && !targetSourceGongStructShape {
				targetSourceGongStructShape = true
				targetGongStructShape = _gongstructshape
			}
		}
		if !targetSourceGongStructShape {
			log.Panicf("GongStructShape %s of field not present ", targetStructName)
		}
		_ = targetGongStructShape

		linkShape := new(LinkShape).Stage(stage)
		linkShape.Name = field.GetName()
		linkShape.SourceMultiplicity = sourceMultiplicity
		linkShape.SourceMultiplicityOffsetX = 0
		linkShape.SourceMultiplicityOffsetY = 0

		linkShape.TargetMultiplicity = targetMultiplicity
		linkShape.TargetMultiplicityOffsetX = 0
		linkShape.TargetMultiplicityOffsetY = 0

		linkShape.FieldOffsetX = 0
		linkShape.FieldOffsetY = 0

		linkShape.Identifier =
			GongstructAndFieldnameToFieldIdentifier(gongStruct.Name, field.GetName())
		fieldIdentifier := GongstructAndFieldnameToFieldIdentifier(gongStruct.Name, field.GetName())
		linkShape.IdentifierMeta = moveStructLiteralToType(fieldIdentifier)
		linkShape.FieldTypeIdentifierMeta = GongStructNameToIdentifier(targetStructName) + "{}"

		gongStructShape.LinkShapes = append(gongStructShape.LinkShapes, linkShape)

		linkShape.X = (gongStructShape.X+targetGongStructShape.X)/2.0 +
			gongStructShape.Width*1.5
		linkShape.Y = (gongStructShape.Y+targetGongStructShape.Y)/2.0 +
			gongStructShape.Height/2.0

		linkShape.StartOrientation = ORIENTATION_HORIZONTAL
		linkShape.StartRatio = 0.5
		linkShape.EndOrientation = ORIENTATION_HORIZONTAL
		linkShape.EndRatio = 0.5
		linkShape.CornerOffsetRatio = 1.38
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
