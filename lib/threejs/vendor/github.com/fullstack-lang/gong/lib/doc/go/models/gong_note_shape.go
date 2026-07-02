package models

import (
	"log"
	"math/rand"

	gong "github.com/fullstack-lang/gong/go/models"
)

// GongNoteShape is a UML note in a class diagram
type GongNoteShape struct {
	Name string

	// Identifier of the note object in the model package
	// a note is a special coment syntax https://pkg.go.dev/go/doc#Note
	// see vendor/github.com/fullstack-lang/gong/go/models/gong_note.go
	// since a note name is not an identifier, it cannot be renamed by gopls
	// unless it is understood to be the name of the identifier immediatly
	// following the comment
	//
	// Gongdoc makes this option a wrtting condition for note:
	// a comment text in gong containing a
	// 'gongnote' have to be followed by a constant identifier.
	//
	//    // GONGNOTE(ShortNodeOnModels): this is an example of a short note
	//    // It uses the DocLink convention for referencing Identifiers
	//    // In this case [Line], [Point] and [Line.Start]
	//    const ShortNodeOnModels = ""
	//
	// the following directive directs gong to manage this field as an identifier field
	//
	//gong:ident
	Identifier string

	Body string

	BodyHTML string

	X, Y          float64
	Width, Height float64
	Matched       bool // if a note with the same name has been found

	GongNoteLinkShapes []*GongNoteLinkShape

	IsExpanded bool
}

func (classdiagram *Classdiagram) HasGongNoteShape(gongstructName string) (foundGongNoteShape bool, gongNoteShape *GongNoteShape) {

	for _, _gongNoteShape := range classdiagram.GongNoteShapes {

		// strange behavior when the gongNoteShape is remove within the loop
		if IdentifierToGongStructName(_gongNoteShape.Identifier) == gongstructName && !foundGongNoteShape {
			foundGongNoteShape = true
			gongNoteShape = _gongNoteShape
		}
	}
	return
}

func (classdiagram *Classdiagram) RemoveGongNoteShape(stage *Stage, gongNoteName string) {

	foundGongNoteShape, gongNoteShape := classdiagram.HasGongNoteShape(gongNoteName)
	if !foundGongNoteShape {
		log.Fatalln("Shape not found", gongNoteName)
	}
	classdiagram.GongNoteShapes = remove(classdiagram.GongNoteShapes, gongNoteShape)
	gongNoteShape.Unstage(stage)

	// remove links that go from this gongNoteShape
	for _, link := range gongNoteShape.GongNoteLinkShapes {
		link.Unstage(stage)
	}
	gongNoteShape.GongNoteLinkShapes = []*GongNoteLinkShape{}

	// remove association links that go to this gongNoteShape
	for _, fromGongNoteShape := range classdiagram.GongNoteShapes {

		newSliceOfLinks := make([]*GongNoteLinkShape, 0)
		for _, noteShapeLink := range fromGongNoteShape.GongNoteLinkShapes {
			typeOfTheField := IdentifierToGongStructName(gongNoteShape.Identifier)
			typeOfTheLink := IdentifierToGongStructName(noteShapeLink.Identifier)
			if typeOfTheLink == typeOfTheField {
				noteShapeLink.Unstage(stage)
			} else {
				newSliceOfLinks = append(newSliceOfLinks, noteShapeLink)
			}
		}
		fromGongNoteShape.GongNoteLinkShapes = newSliceOfLinks
	}

	//
	// remove documentation links that go this gongNoteShape
	//
	// generate the map to navigate from children to parents
	fieldName := GetAssociationName[GongNoteShape]().GongNoteLinkShapes[0].Name
	map_NoteShapeLink_NoteShape := GetSliceOfPointersReverseMap[GongNoteShape, GongNoteLinkShape](fieldName, stage)
	for noteShapeLink := range *GetGongstructInstancesSet[GongNoteLinkShape](stage) {
		if noteShapeLink.Name == gongNoteName {

			// get the note shape
			noteShapes := map_NoteShapeLink_NoteShape[noteShapeLink]

			// remove it from the slice of links
			noteShapes[0].GongNoteLinkShapes = remove(noteShapes[0].GongNoteLinkShapes, noteShapeLink)

			noteShapeLink.Unstage(stage)
		}
	}

	// log.Println("RemoveGongNoteShape, before commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())
	stage.Commit()
	// log.Println("RemoveGongNoteShape, after commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())
}

func (classdiagram *Classdiagram) AddGongNoteShape(stage *Stage, gongNote *gong.GongNote, diagramPackage *DiagramPackage, gongNoteShapeName string) {

	var gongNoteShape GongNoteShape
	gongNoteShape.Name = classdiagram.Name + "-" + gongNoteShapeName
	gongNoteShape.Identifier = GongNoteNameToIdentifier(gongNoteShapeName)
	gongNoteShape.Width = 240
	gongNoteShape.Height = 63

	gongNoteShape.Body = gongNote.Body
	gongNoteShape.BodyHTML = gongNote.BodyHTML

	gongNoteShape.Stage(stage)

	gongNoteShape.X = float64(int(rand.Float32()*100) + 10)
	gongNoteShape.Y = float64(int(rand.Float32()*100) + 10)

	classdiagram.GongNoteShapes = append(classdiagram.GongNoteShapes, &gongNoteShape)

	// log.Println("AddGongNoteShape, before commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())
	stage.Commit()
	// log.Println("AddGongNoteShape, after commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())

}
