package models

import (
	"slices"

	gong "github.com/fullstack-lang/gong/go/models"
)

// A GongNoteLinkShape is a visual link from a shape to
// a Link or to a Classshape
//
// # The end point of the link is computed from the Type
//
// Because of[issue](https://github.com/golang/go/issues/57559)
// the referenced identifiers in the note comments are not
// renamed. Unlinke for identifiers used for referencing
// stuff, there is workaround this limitation
type GongNoteLinkShape struct {
	Name string

	// Identifier of the target shape / link of the note link
	//gong:ident
	Identifier string

	// Type of the target shape / link of the note link
	Type NoteShapeLinkType
}

func (classdiagram *Classdiagram) AddGongNoteLinkShapeToDiagram(
	stage *Stage,
	gongNoteShape *GongNoteShape,
	gongLink *gong.GongLink,
) {
	var gongNoteLinkShape GongNoteLinkShape
	gongNoteLinkShape.Name = classdiagram.Name + "-" + gongNoteShape.Name + "-" + gongLink.Name
	if gongLink.Recv != "" {
		gongNoteLinkShape.Identifier = GongstructAndFieldnameToFieldIdentifier(gongLink.Recv, gongLink.Name)
		gongNoteLinkShape.Type = NOTE_SHAPE_LINK_TO_GONG_FIELD
	} else {
		gongNoteLinkShape.Identifier = GongStructNameToIdentifier(gongLink.Name)
		gongNoteLinkShape.Type = NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE
	}
	gongNoteLinkShape.Stage(stage)

	gongNoteShape.GongNoteLinkShapes = append(gongNoteShape.GongNoteLinkShapes, &gongNoteLinkShape)
}

func (classdiagram *Classdiagram) RemoveGongNoteLinkShapeFromDiagram(
	stage *Stage,
	gongNoteShape *GongNoteShape,
	gongLink *gong.GongLink,
) {
	var gongNoteLinkShape *GongNoteLinkShape
	var identifier string
	if gongLink.Recv != "" {
		identifier = GongstructAndFieldnameToFieldIdentifier(gongLink.Recv, gongLink.Name)
	} else {
		identifier = GongStructNameToIdentifier(gongLink.Name)
	}

	for _, _gongNoteLinkShape := range gongNoteShape.GongNoteLinkShapes {
		if _gongNoteLinkShape.Identifier == identifier {
			gongNoteLinkShape = _gongNoteLinkShape
		}
	}
	if gongNoteLinkShape != nil {
		idx := slices.Index(gongNoteShape.GongNoteLinkShapes, gongNoteLinkShape)
		gongNoteShape.GongNoteLinkShapes = slices.Delete(gongNoteShape.GongNoteLinkShapes, idx, idx+1)
		gongNoteLinkShape.Unstage(stage)
	}
}
