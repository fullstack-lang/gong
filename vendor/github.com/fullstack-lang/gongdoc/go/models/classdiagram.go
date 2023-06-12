package models

import (
	"math/rand"
)

// Classdiagram diagram struct store a class diagram
// temporary here
// swagger:model Classdiagram
type Classdiagram struct {
	Name string

	// list of gongstructshapes in the diagram
	GongStructShapes []*GongStructShape

	GongEnumShapes []*GongEnumShape

	// list of notes in the diagram
	NoteShapes []*NoteShape

	// IsInDrawMode indicates the the drawing can be edited (in development mode)
	// or not (in production mode)
	IsInDrawMode bool
}

func (classdiagram *Classdiagram) RemoveGongStructShape(stage *StageStruct, gongstructshapeName string) {

	foundGongStructShape := false
	var gongstructshape *GongStructShape
	for _, _gongstructshape := range classdiagram.GongStructShapes {

		// strange behavior when the gongstructshape is remove within the loop
		if IdentifierToGongObjectName(_gongstructshape.Identifier) == gongstructshapeName && !foundGongStructShape {
			gongstructshape = _gongstructshape
		}
	}

	classdiagram.GongStructShapes = remove(classdiagram.GongStructShapes, gongstructshape)
	gongstructshape.Position.Unstage(stage)
	gongstructshape.Unstage(stage)

	// remove links that go from this gongstructshape
	for _, link := range gongstructshape.Links {
		link.Middlevertice.Unstage(stage)
		link.Unstage(stage)
	}
	gongstructshape.Links = []*Link{}

	// remove association links that go to this gongstructshape
	for _, fromGongStructShape := range classdiagram.GongStructShapes {

		newSliceOfLinks := make([]*Link, 0)
		for _, link := range fromGongStructShape.Links {
			typeOfTheField := IdentifierToGongObjectName(gongstructshape.Identifier)
			typeOfTheLink := IdentifierToGongObjectName(link.Fieldtypename)
			if typeOfTheLink == typeOfTheField {
				link.Middlevertice.Unstage(stage)
				link.Unstage(stage)
			} else {
				newSliceOfLinks = append(newSliceOfLinks, link)
			}
		}
		fromGongStructShape.Links = newSliceOfLinks
	}

	// remove fields of the gongstructshape
	for _, field := range gongstructshape.Fields {
		field.Unstage(stage)
	}

	//
	// remove documentation links that go this gongstructshape
	//
	// generate the map to navigate from children to parents
	fieldName := GetAssociationName[NoteShape]().NoteShapeLinks[0].Name
	map_NoteShapeLink_NodeShape := GetSliceOfPointersReverseMap[NoteShape, NoteShapeLink](fieldName, stage)
	for noteShapeLink := range *GetGongstructInstancesSet[NoteShapeLink](stage) {
		if noteShapeLink.Name == gongstructshapeName {

			// get the note shape
			noteShape := map_NoteShapeLink_NodeShape[noteShapeLink]

			// remove it from the slice of links
			noteShape.NoteShapeLinks = remove(noteShape.NoteShapeLinks, noteShapeLink)

			noteShapeLink.Unstage(stage)
		}
	}

	// log.Println("RemoveGongStructShape, before commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())
	stage.Commit()
	// log.Println("RemoveGongStructShape, after commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())
}

func (classdiagram *Classdiagram) AddGongStructShape(stage *StageStruct, diagramPackage *DiagramPackage, gongstructshapeName string) {

	var gongstructshape GongStructShape
	gongstructshape.Name = classdiagram.Name + "-" + gongstructshapeName
	gongstructshape.Identifier = GongStructNameToIdentifier(gongstructshapeName)
	gongstructshape.Width = 240
	gongstructshape.Heigth = 63

	// attach GongStruct to gongstructshape
	nbInstances, ok := diagramPackage.Map_Identifier_NbInstances[gongstructshape.Identifier]
	if ok {
		gongstructshape.ShowNbInstances = true
		gongstructshape.NbInstances = nbInstances
	}
	gongstructshape.Stage(stage)

	var position Position
	position.Name = "Pos-" + gongstructshape.Name
	position.X = float64(int(rand.Float32()*100) + 10)
	position.Y = float64(int(rand.Float32()*100) + 10)
	gongstructshape.Position = &position
	position.Stage(stage)

	classdiagram.GongStructShapes = append(classdiagram.GongStructShapes, &gongstructshape)

	// log.Println("AddGongStructShape, before commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())
	stage.Commit()
	// log.Println("AddGongStructShape, after commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())

}

func (classdiagram *Classdiagram) AddGongEnumShape(stage *StageStruct, diagramPackage *DiagramPackage, enumshapeName string) {

	var enumshape GongEnumShape
	enumshape.Name = classdiagram.Name + "-" + enumshapeName
	enumshape.Identifier = GongStructNameToIdentifier(enumshapeName)
	enumshape.Width = 240
	enumshape.Heigth = 63

	enumshape.Stage(stage)

	var position Position
	position.Name = "Pos-" + enumshape.Name
	position.X = float64(int(rand.Float32()*100) + 10)
	position.Y = float64(int(rand.Float32()*100) + 10)
	enumshape.Position = &position
	position.Stage(stage)

	classdiagram.GongEnumShapes = append(classdiagram.GongEnumShapes, &enumshape)

	stage.Commit()
}

func (classdiagram *Classdiagram) RemoveGongEnumShape(stage *StageStruct, gongenumshapeName string) {

	foundGongEnumShape := false
	var gongenumshape *GongEnumShape
	for _, _gongenumshape := range classdiagram.GongEnumShapes {

		// strange behavior when the gongenumshape is remove within the loop
		if IdentifierToGongObjectName(_gongenumshape.Identifier) == gongenumshapeName && !foundGongEnumShape {
			gongenumshape = _gongenumshape
		}
	}

	classdiagram.GongEnumShapes = remove(classdiagram.GongEnumShapes, gongenumshape)
	gongenumshape.Position.Unstage(stage)
	gongenumshape.Unstage(stage)

	// remove fields of the gongenumshape
	for _, gongEnumValueEntry := range gongenumshape.GongEnumValueEntrys {
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

	stage.Commit()
}
