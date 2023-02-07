package node2gongdoc

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
)

type GongNoteImpl struct {
	gongNote *gong_models.GongNote
	NodeImpl
}

func (gongNoteImpl *GongNoteImpl) OnAfterUpdate(
	stage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {

	classdiagram := gongNoteImpl.nodeCb.GetSelectedClassdiagram()

	// adding a note shape
	if !stagedNode.IsChecked && frontNode.IsChecked {
		stage.Checkout()

		noteShape := (&gongdoc_models.NoteShape{Name: stagedNode.Name}).Stage()

		mapOfGongNotes := *gong_models.GetGongstructInstancesMap[gong_models.GongNote]()

		gongNote, ok := mapOfGongNotes[noteShape.Name]
		if !ok {
			log.Fatal("Unkown note ", noteShape.Name)
		}

		noteShape.Identifier = gongdoc_models.GongStructNameToIdentifier(noteShape.Name)

		noteShape.Body = gongNote.Body
		noteShape.X = 30
		noteShape.Y = 30
		noteShape.Width = 240
		noteShape.Heigth = 63

		classdiagram.NoteShapes = append(classdiagram.NoteShapes, noteShape)

	}

	// suppression a note
	if stagedNode.IsChecked && !frontNode.IsChecked {
		stage.Checkout()
		foundNote := false
		var noteShape *gongdoc_models.NoteShape
		var _noteShape *gongdoc_models.NoteShape
		for _, _noteShape = range classdiagram.NoteShapes {

			// strange behavior when the note is removed within the loop
			if gongdoc_models.IdentifierToGongObjectName(_noteShape.Identifier) == stagedNode.Name && !foundNote {
				foundNote = true
				noteShape = _noteShape
			}
		}
		if !foundNote {
			log.Panicf("Note %s of field not present ", stagedNode.Name)
		}

		// remove the links of the note shape
		for _, noteLink := range noteShape.NoteShapeLinks {
			noteLink.Unstage()
			noteShape.NoteShapeLinks = remove(noteShape.NoteShapeLinks, noteLink)
		}
		classdiagram.NoteShapes = remove(classdiagram.NoteShapes, noteShape)
		noteShape.Unstage()

	}

}

func (GongNoteImpl *GongNoteImpl) OnAfterDelete(
	stage *gongdoc_models.StageStruct,
	stagedNode, frontNode *gongdoc_models.Node) {
}
