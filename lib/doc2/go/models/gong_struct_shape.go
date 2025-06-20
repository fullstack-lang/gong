package models

import (
	"log"
	"math/rand"
)

const GongStructShapeDefaultWidth = 240.0
const GongStructShapeDefaultHeight = 48.0

// GongStructShape mirrors joint.shapes.uml.Class
// swagger:model GongStructShape
type GongStructShape struct {
	Name string

	X float64
	Y float64

	// Identifier is the identifier of the struct referenced by the shape in the modeled package
	//gong:meta
	IdentifierMeta any

	// models of the composition of Field
	AttributeShapes []*AttributeShape

	// models of the composition of Link
	LinkShapes []*LinkShape

	// with and height of the shape when they are rendered on SVG or with jointjs
	// They are optional fields. they can be computed when empty
	Width, Height float64

	// this is always false in the backend, but it can be set to true by the front end
	// this means it is selected by the user
	IsSelected bool
}

func (classdiagram *Classdiagram) HasGongStructShape(gongstructName string) (foundGongStructShape bool, gongstructshape *GongStructShape) {

	for _, _gongstructshape := range classdiagram.GongStructShapes {

		// strange behavior when the gongstructshape is remove within the loop
		if IdentifierMetaToGongStructName(_gongstructshape.IdentifierMeta) == gongstructName && !foundGongStructShape {
			foundGongStructShape = true
			gongstructshape = _gongstructshape
		}
	}
	return
}

func (classdiagram *Classdiagram) RemoveGongStructShape(stage *Stage, gongstructName string) {

	foundGongStructShape, gongstructshape := classdiagram.HasGongStructShape(gongstructName)
	if !foundGongStructShape {
		log.Fatalln("Shape not found", gongstructName)
	}
	classdiagram.GongStructShapes = remove(classdiagram.GongStructShapes, gongstructshape)
	gongstructshape.Unstage(stage)

	// remove links that go from this gongstructshape
	for _, link := range gongstructshape.LinkShapes {
		link.Unstage(stage)
	}
	gongstructshape.LinkShapes = []*LinkShape{}

	// remove association links that go to this gongstructshape
	for _, fromGongStructShape := range classdiagram.GongStructShapes {

		newSliceOfLinks := make([]*LinkShape, 0)
		for _, linkShape := range fromGongStructShape.LinkShapes {
			typeOfTheField := IdentifierMetaToGongStructName(gongstructshape.IdentifierMeta)
			typeOfTheLink := IdentifierMetaToGongStructName(linkShape.FieldTypeIdentifierMeta)
			if typeOfTheLink == typeOfTheField {
				linkShape.Unstage(stage)
			} else {
				newSliceOfLinks = append(newSliceOfLinks, linkShape)
			}
		}
		fromGongStructShape.LinkShapes = newSliceOfLinks
	}

	// remove fields of the gongstructshape
	for _, field := range gongstructshape.AttributeShapes {
		field.Unstage(stage)
	}

	//
	// remove documentation links that go this gongstructshape
	//
	// generate the map to navigate from children to parents
	fieldName := GetAssociationName[GongNoteShape]().GongNoteLinkShapes[0].Name
	map_NoteShapeLink_NodeShape := GetSliceOfPointersReverseMap[GongNoteShape, GongNoteLinkShape](fieldName, stage)
	for noteShapeLink := range *GetGongstructInstancesSet[GongNoteLinkShape](stage) {
		if noteShapeLink.Name == gongstructName {

			// get the note shape
			noteShapes := map_NoteShapeLink_NodeShape[noteShapeLink]

			// remove it from the slice of links
			noteShapes[0].GongNoteLinkShapes = remove(noteShapes[0].GongNoteLinkShapes, noteShapeLink)

			noteShapeLink.Unstage(stage)
		}
	}

	// log.Println("RemoveGongStructShape, before commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())
	stage.Commit()
	// log.Println("RemoveGongStructShape, after commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())
}

func (classdiagram *Classdiagram) AddGongStructShape(stage *Stage, diagramPackage *DiagramPackage, gongStructShapeName string) {

	var gongStructShape GongStructShape
	gongStructShape.Name = classdiagram.Name + "-" + gongStructShapeName
	// gongStructShape.Identifier = GongStructNameToIdentifier(gongStructShapeName)

	// for instanciation of the struct ref_models.Astruct{}
	gongStructShape.IdentifierMeta = GongStructNameToIdentifier(gongStructShapeName) + "{}"
	gongStructShape.Width = 240
	gongStructShape.Height = 63

	// attach GongStruct to gongstructshape
	// nbInstances, ok := diagramPackage.Map_Identifier_NbInstances[gongStructShapeName]

	gongStructShape.Stage(stage)

	gongStructShape.X = float64(int(rand.Float32()*100) + 10)
	gongStructShape.Y = float64(int(rand.Float32()*100) + 10)

	classdiagram.GongStructShapes = append(classdiagram.GongStructShapes, &gongStructShape)

	// log.Println("AddGongStructShape, before commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())
	stage.Commit()
	// log.Println("AddGongStructShape, after commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())

}
