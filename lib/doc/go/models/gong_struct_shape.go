package models

import (
	"log"
	"math/rand"

	svg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

const GongStructShapeDefaultWidth = 240.0
const GongStructShapeDefaultHeight = 48.0

// GongStructShape mirrors joint.shapes.uml.Class
// swagger:model GongStructShape
type GongStructShape struct {
	Name string

	svg_models.RectShape

	// Identifier is the identifier of the struct referenced by the shape in the modeled package
	//gong:meta
	IdentifierMeta any

	// models of the composition of Field
	AttributeShapes []*AttributeShape

	// models of the composition of Link
	LinkShapes []*LinkShape

	// this is always false in the backend, but it can be set to true by the front end
	// this means it is selected by the user
	IsSelected bool
}

func (classdiagram *Classdiagram) HasGongStructShape(gongstructName string) (foundGongStructShape bool, gongstructshape *GongStructShape) {
	return classdiagram.HasGongStructShapeWithPackage("", gongstructName)
}

func (classdiagram *Classdiagram) HasGongStructShapeWithPackage(pkgName, gongstructName string) (foundGongStructShape bool, gongstructshape *GongStructShape) {
	for _, _gongstructshape := range classdiagram.GongStructShapes {
		pkg, name := IdentifierMetaToPackageAndGongStructName(_gongstructshape.IdentifierMeta)
		if (pkgName == "" || pkg == pkgName) && name == gongstructName && !foundGongStructShape {
			foundGongStructShape = true
			gongstructshape = _gongstructshape
		}
	}
	return
}

func (classdiagram *Classdiagram) RemoveGongStructShape(stage *Stage, gongstructName string) {
	classdiagram.RemoveGongStructShapeWithPackage(stage, "", gongstructName)
}

func (classdiagram *Classdiagram) RemoveGongStructShapeWithPackage(stage *Stage, pkgName, gongstructName string) {

	foundGongStructShape, gongstructshape := classdiagram.HasGongStructShapeWithPackage(pkgName, gongstructName)
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
			targetPkg, typeOfTheField := IdentifierMetaToPackageAndGongStructName(gongstructshape.IdentifierMeta)
			linkTargetPkg, typeOfTheLink := IdentifierMetaToPackageAndGongStructName(linkShape.FieldTypeIdentifierMeta)
			if typeOfTheLink == typeOfTheField && (targetPkg == linkTargetPkg || targetPkg == "" || linkTargetPkg == "") {
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

	// remove documentation links that go this gongstructshape
	fieldName := GongGetAssociationName[GongNoteShape]().GongNoteLinkShapes[0].Name
	map_NoteShapeLink_NodeShape := stage.GetSliceOfPointersReverseMap[GongNoteShape, GongNoteLinkShape](fieldName)
	for noteShapeLink := range *stage.GetInstancesSet[*GongNoteLinkShape]() {
		if noteShapeLink.Name == gongstructName {

			// get the note shape
			noteShapes := map_NoteShapeLink_NodeShape[noteShapeLink]

			// remove it from the slice of links
			noteShapes[0].GongNoteLinkShapes = remove(noteShapes[0].GongNoteLinkShapes, noteShapeLink)

			noteShapeLink.Unstage(stage)
		}
	}

	stage.Commit()
}

func (classdiagram *Classdiagram) AddGongStructShape(stage *Stage, diagramPackage *DiagramPackage, gongStructShapeName string) {
	classdiagram.AddGongStructShapeWithPackage(stage, diagramPackage, "models", gongStructShapeName)
}

func (classdiagram *Classdiagram) AddGongStructShapeWithPackage(stage *Stage, diagramPackage *DiagramPackage, pkgName string, gongStructShapeName string) {
	if pkgName == "" {
		pkgName = "models"
	}
	var gongStructShape GongStructShape
	gongStructShape.Name = classdiagram.Name + "-" + gongStructShapeName

	// for instanciation of the struct ref_<pkg>.<Struct>{}
	gongStructShape.IdentifierMeta = GongStructNameToIdentifierWithPackage(pkgName, gongStructShapeName) + "{}"
	gongStructShape.Width = 240
	gongStructShape.Height = 63

	gongStructShape.Stage(stage)

	gongStructShape.X = float64(int(rand.Float32()*100) + 10)
	gongStructShape.Y = float64(int(rand.Float32()*100) + 10)

	classdiagram.GongStructShapes = append(classdiagram.GongStructShapes, &gongStructShape)

	stage.Commit()
}

