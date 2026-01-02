package models

import (
	"fmt"
	"time"
)

// enforceNoteRelatedShapes ensures that NoteShape instances have a valid Note reference.
// It removes NoteShapes with a nil Note, unstages them, and notifies the user.
func (stager *Stager) enforceNoteRelatedShapes() (needCommit bool) {
	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {
		// Filter out NoteShapes that have a nil Note
		var validNoteShapes []*NoteShape
		for _, noteShape := range diagram.Note_Shapes {
			if noteShape.Note == nil {
				noteShape.UnstageVoid(stager.stage)
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Removed NoteShape %s with missing Note reference", noteShape.GetName()))
				needCommit = true
			} else {
				validNoteShapes = append(validNoteShapes, noteShape)
			}
		}
		if needCommit {
			diagram.Note_Shapes = validNoteShapes
		}

		for _, noteProductShape := range diagram.NoteProductShapes {
			if noteProductShape.Note == nil {
				noteProductShape.UnstageVoid(stager.stage)
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Removed noteProductShape \"%s\" with missing Note reference", noteProductShape.GetName()))
				needCommit = true
			}
			if noteProductShape.Product == nil {
				noteProductShape.Unstage(stager.stage)
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Removed noteProductShape \"%s\" with missing Product reference", noteProductShape.GetName()))
				needCommit = true
			}
		}

		for _, noteTaskShape := range diagram.NoteTaskShapes {
			if noteTaskShape.Note == nil {
				noteTaskShape.UnstageVoid(stager.stage)
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Removed noteTaskShape \"%s\" with missing Note reference", noteTaskShape.GetName()))
				needCommit = true
			}
			if noteTaskShape.Task == nil {
				noteTaskShape.Unstage(stager.stage)
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Removed noteTaskShape \"%s\" with missing Task reference", noteTaskShape.GetName()))
				needCommit = true
			}
		}

	}
	return
}
