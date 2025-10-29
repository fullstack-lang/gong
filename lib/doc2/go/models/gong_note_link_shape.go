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

func (classdiagram *Classdiagram) RemoveGongNoteLinkShape(
	stage *Stage,
	gongNoteLinkShape *GongNoteLinkShape,
	gongNoteShape *GongNoteShape) {

	idx := slices.Index(gongNoteShape.GongNoteLinkShapes, gongNoteLinkShape)
	gongNoteShape.GongNoteLinkShapes = slices.Delete(gongNoteShape.GongNoteLinkShapes, idx, idx+1)
	gongNoteLinkShape.Unstage(stage)

}

func (classdiagram *Classdiagram) AddGongNoteLinkShape(
	stage *Stage,
	gongStage *gong.Stage,
	gongNote *gong.GongNote,
	gongNoteLink *gong.GongLink,
	gongNoteShape *GongNoteShape) {

	gongNoteLinkShape := (&GongNoteLinkShape{
		Name:       gongNoteLink.Name,
		Identifier: gongNoteLink.ImportPath + "." + gongNoteLink.Name,
	}).Stage(stage)

	if gongNoteLink.Recv == "" {
		gongNoteLinkShape.Type = NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE
	} else {
		gongNoteLinkShape.Type = NOTE_SHAPE_LINK_TO_GONG_FIELD
	}

	gongNoteShape.GongNoteLinkShapes = append(gongNoteShape.GongNoteLinkShapes, gongNoteLinkShape)

}
