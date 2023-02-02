package models

import (
	"math/rand"
)

// Classdiagram diagram struct store a class diagram
// temporary here
// swagger:model Classdiagram
type Classdiagram struct {
	Name string

	// list of classshapes in the diagram
	GongStructShapes []*GongStructShape

	GongEnumShapes []*GongEnumShape

	// list of notes in the diagram
	NoteShapes []*NoteShape

	// IsInDrawMode indicates the the drawing can be edited (in development mode)
	// or not (in production mode)
	IsInDrawMode bool
}

func (classdiagram *Classdiagram) RemoveClassshape(classshapeName string) {

	foundClassshape := false
	var classshape *GongStructShape
	for _, _classshape := range classdiagram.GongStructShapes {

		// strange behavior when the classshape is remove within the loop
		if IdentifierToGongStructName(_classshape.Identifier) == classshapeName && !foundClassshape {
			classshape = _classshape
		}
	}

	classdiagram.GongStructShapes = remove(classdiagram.GongStructShapes, classshape)
	classshape.Position.Unstage()
	classshape.Unstage()

	// remove links that go from this classshape
	for _, link := range classshape.Links {
		link.Middlevertice.Unstage()
		link.Unstage()
	}
	classshape.Links = []*Link{}

	// remove association links that go to this classshape
	for _, fromClassshape := range classdiagram.GongStructShapes {

		newSliceOfLinks := make([]*Link, 0)
		for _, link := range fromClassshape.Links {
			if link.Fieldtypename == IdentifierToGongStructName(classshape.Identifier) {
				link.Middlevertice.Unstage()
				link.Unstage()
			} else {
				newSliceOfLinks = append(newSliceOfLinks, link)
			}
		}
		fromClassshape.Links = newSliceOfLinks
	}

	// remove fields of the classshape
	for _, field := range classshape.Fields {
		field.Unstage()
	}

	//
	// remove documentation links that go this classshape
	//
	// generate the map to navigate from children to parents
	fieldName := GetAssociationName[NoteShape]().NoteShapeLinks[0].Name
	map_NoteShapeLink_NodeShape := GetSliceOfPointersReverseMap[NoteShape, NoteShapeLink](fieldName)
	for noteShapeLink := range *GetGongstructInstancesSet[NoteShapeLink]() {
		if noteShapeLink.Name == classshapeName {

			// get the note shape
			noteShape := map_NoteShapeLink_NodeShape[noteShapeLink]

			// remove it from the slice of links
			noteShape.NoteShapeLinks = remove(noteShape.NoteShapeLinks, noteShapeLink)

			noteShapeLink.DeleteStageAndCommit()
		}
	}

	// log.Println("RemoveClassshape, before commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())
	Stage.Commit()
	// log.Println("RemoveClassshape, after commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())
}

func (classdiagram *Classdiagram) AddClassshape(classshapeName string) {

	var classshape GongStructShape
	classshape.Name = classdiagram.Name + "-" + classshapeName
	classshape.Identifier = GongStructNameToIdentifier(classshapeName)
	classshape.Width = 240
	classshape.Heigth = 63

	// attach GongStruct to classshape
	nbInstances, ok := Map_Identifier_NbInstances[classshape.Identifier]
	if ok {
		classshape.ShowNbInstances = true
		classshape.NbInstances = nbInstances
	}
	classshape.Stage()

	var position Position
	position.Name = "Pos-" + classshape.Name
	position.X = float64(int(rand.Float32()*100) + 10)
	position.Y = float64(int(rand.Float32()*100) + 10)
	classshape.Position = &position
	position.Stage()

	classdiagram.GongStructShapes = append(classdiagram.GongStructShapes, &classshape)

	// log.Println("AddClassshape, before commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())
	Stage.Commit()
	// log.Println("AddClassshape, after commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())

}
