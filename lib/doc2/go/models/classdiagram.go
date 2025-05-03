package models

import (
	"log"
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

	// IsInRenameMode means the user can edit the node
	IsInRenameMode bool

	// IsExpanded is true if the corresponding node is expanded
	IsExpanded bool

	NodeGongStructsIsExpanded bool

	// encoding of the the expanded status per named struct node
	// the bit position is the NodeNamedStruct position
	// beyond 63, there is no storage of expanded status
	//
	// bit 0 at 0 means first gong struct is not encoded
	// bit 1 at 0 means second gong struct is not encoded
	// etc...
	NodeGongStructNodeExpansionBinaryEncoding int

	NodeGongEnumsIsExpanded                 bool
	NodeGongEnumNodeExpansionBinaryEncoding int

	NodeGongNotesIsExpanded                 bool
	NodeGongNoteNodeExpansionBinaryEncoding int
}

func (classdiagram *Classdiagram) HasGongStructShape(gongstructName string) (foundGongStructShape bool, gongstructshape *GongStructShape) {

	for _, _gongstructshape := range classdiagram.GongStructShapes {

		// strange behavior when the gongstructshape is remove within the loop
		if IdentifierToGongObjectName(_gongstructshape.Identifier) == gongstructName && !foundGongStructShape {
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
		for _, link := range fromGongStructShape.LinkShapes {
			typeOfTheField := IdentifierToGongObjectName(gongstructshape.Identifier)
			typeOfTheLink := IdentifierToGongObjectName(link.Fieldtypename)
			if typeOfTheLink == typeOfTheField {
				link.Unstage(stage)
			} else {
				newSliceOfLinks = append(newSliceOfLinks, link)
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
	fieldName := GetAssociationName[NoteShape]().NoteShapeLinks[0].Name
	map_NoteShapeLink_NodeShape := GetSliceOfPointersReverseMap[NoteShape, NoteShapeLink](fieldName, stage)
	for noteShapeLink := range *GetGongstructInstancesSet[NoteShapeLink](stage) {
		if noteShapeLink.Name == gongstructName {

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

func (classdiagram *Classdiagram) AddGongStructShape(stage *Stage, diagramPackage *DiagramPackage, gongstructshapeName string) {

	var gongstructshape GongStructShape
	gongstructshape.Name = classdiagram.Name + "-" + gongstructshapeName
	gongstructshape.Identifier = GongStructNameToIdentifier(gongstructshapeName)
	gongstructshape.Width = 240
	gongstructshape.Height = 63

	// attach GongStruct to gongstructshape
	nbInstances, ok := diagramPackage.Map_Identifier_NbInstances[gongstructshape.Identifier]
	if ok {
		gongstructshape.ShowNbInstances = true
		gongstructshape.NbInstances = nbInstances
	}
	gongstructshape.Stage(stage)

	gongstructshape.X = float64(int(rand.Float32()*100) + 10)
	gongstructshape.Y = float64(int(rand.Float32()*100) + 10)

	classdiagram.GongStructShapes = append(classdiagram.GongStructShapes, &gongstructshape)

	// log.Println("AddGongStructShape, before commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())
	stage.Commit()
	// log.Println("AddGongStructShape, after commit, nb ", Stage.BackRepo.GetLastCommitFromBackNb())

}

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

	stage.Commit()
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

// DuplicateDiagram generates a new diagram with duplicated shapes
func (classdiagram *Classdiagram) DuplicateDiagram() (newClassdiagram *Classdiagram) {

	newClassdiagram = CopyBranch(classdiagram)

	return
}

func IsNodeExpanded(binaryEncoding, nodeRank int) bool {
	if nodeRank < 0 || nodeRank > 63 {
		// Rank must be between 0 and 63 for a 64-bit integer
		return false
	}
	// Check if the rank-th note is played
	return binaryEncoding&(1<<nodeRank) != 0
}

func ToggleNodeExpanded(binaryEncoding *int, nodeRank int) {
	if nodeRank < 0 || nodeRank > 63 {
		return // Ignore invalid ranks
	}

	*binaryEncoding ^= 1 << nodeRank
}
