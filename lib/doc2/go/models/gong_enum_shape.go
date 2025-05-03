package models

import "math/rand"

const GongEnumShapeDefaultWidth = 240.0
const GongEnumShapeDefaultHeight = 48.0

// GongEnumShape mirrors joint.shapes.uml.Class
// swagger:model GongEnumShape
type GongEnumShape struct {
	Name string

	X, Y float64

	// Identifier is the identifier of the struct referenced by the shape in the modeled package
	//gong:ident
	Identifier string

	GongEnumValueShapes []*GongEnumValueShape

	// with and height of the shape when they are rendered on SVG or with jointjs
	// They are optional fields. they can be computed when empty
	Width, Height float64

	IsExpanded bool
}

type GongEnumShapeType int

const (
	Int GongEnumShapeType = iota
	String
)

func (classdiagram *Classdiagram) AddGongEnumShape(stage *Stage, diagramPackage *DiagramPackage, enumshapeName string) {

	var enumshape GongEnumShape
	enumshape.Name = classdiagram.Name + "-" + enumshapeName
	enumshape.Identifier = GongStructNameToIdentifier(enumshapeName)
	enumshape.Width = 240
	enumshape.Height = 63

	enumshape.Stage(stage)

	enumshape.X = float64(int(rand.Float32()*100) + 10)
	enumshape.Y = float64(int(rand.Float32()*100) + 10)

	classdiagram.GongEnumShapes = append(classdiagram.GongEnumShapes, &enumshape)
}

func (classdiagram *Classdiagram) RemoveGongEnumShape(stage *Stage, gongenumshapeName string) {

	foundGongEnumShape := false
	var gongenumshape *GongEnumShape
	for _, _gongenumshape := range classdiagram.GongEnumShapes {

		// strange behavior when the gongenumshape is remove within the loop
		if IdentifierToGongObjectName(_gongenumshape.Identifier) == gongenumshapeName && !foundGongEnumShape {
			gongenumshape = _gongenumshape
		}
	}

	classdiagram.GongEnumShapes = remove(classdiagram.GongEnumShapes, gongenumshape)
	gongenumshape.Unstage(stage)

	// remove fields of the gongenumshape
	for _, gongEnumValueEntry := range gongenumshape.GongEnumValueShapes {
		gongEnumValueEntry.Unstage(stage)
	}

	//
	// remove documentation links that go this gongenumshape
	//
	// generate the map to navigate from children to parents
	fieldName := GetAssociationName[NoteShape]().NoteShapeLinks[0].Name
	map_NoteShapeLink_NodeShape := GetSliceOfPointersReverseMap[NoteShape, NoteShapeLink](fieldName, stage)
	for noteShapeLink := range *GetGongstructInstancesSet[NoteShapeLink](stage) {
		if noteShapeLink.Name == gongenumshapeName {

			// get the note shape
			noteShape := map_NoteShapeLink_NodeShape[noteShapeLink]

			// remove it from the slice of links
			noteShape.NoteShapeLinks = remove(noteShape.NoteShapeLinks, noteShapeLink)

			noteShapeLink.Unstage(stage)
		}
	}
}
