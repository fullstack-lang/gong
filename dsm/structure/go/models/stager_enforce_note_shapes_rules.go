package models

import (
	"fmt"
	"slices"
	"time"
)

func (stager *Stager) enforceNoteShapesRules() (needCommit bool) {
	for _, diagram := range stager.stage.GetInstancesSorted[*DiagramStructure]() {
		seenPortKeys := make(map[notePortKey]bool)
		var validNotePortShapes []*NotePortShape
		for _, shape := range diagram.NotePortShapes {
			key := notePortKey{Note: shape.Note, Port: shape.Port}
			isValid := shape.Note != nil &&
				shape.Port != nil &&
				slices.Contains(shape.Note.Ports, shape.Port) &&
				!seenPortKeys[key]
			if isValid {
				seenPortKeys[key] = true
				validNotePortShapes = append(validNotePortShapes, shape)
			} else {
				shape.UnstageVoid(stager.stage)
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged note port shape %s (invalid or duplicate link)", shape.GetName()))
				}
			}
		}
		if len(validNotePortShapes) != len(diagram.NotePortShapes) {
			diagram.NotePortShapes = validNotePortShapes
			needCommit = true
		}

		seenPartKeys := make(map[notePartKey]bool)
		var validNotePartShapes []*NotePartShape
		for _, shape := range diagram.NotePartShapes {
			key := notePartKey{Note: shape.Note, Part: shape.Part}
			isValid := shape.Note != nil &&
				shape.Part != nil &&
				slices.Contains(shape.Note.Parts, shape.Part) &&
				!seenPartKeys[key]
			if isValid {
				seenPartKeys[key] = true
				validNotePartShapes = append(validNotePartShapes, shape)
			} else {
				shape.UnstageVoid(stager.stage)
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged note part shape %s (invalid or duplicate link)", shape.GetName()))
				}
			}
		}
		if len(validNotePartShapes) != len(diagram.NotePartShapes) {
			diagram.NotePartShapes = validNotePartShapes
			needCommit = true
		}

		seenDataFlowKeys := make(map[noteDataFlowKey]bool)
		var validNoteDataFlowShapes []*NoteDataFlowShape
		for _, shape := range diagram.NoteDataFlowShapes {
			key := noteDataFlowKey{Note: shape.Note, DataFlow: shape.DataFlow}
			isValid := shape.Note != nil &&
				shape.DataFlow != nil &&
				slices.Contains(shape.Note.DataFlows, shape.DataFlow) &&
				!seenDataFlowKeys[key]
			if isValid {
				seenDataFlowKeys[key] = true
				validNoteDataFlowShapes = append(validNoteDataFlowShapes, shape)
			} else {
				shape.UnstageVoid(stager.stage)
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged note data flow shape %s (invalid or duplicate link)", shape.GetName()))
				}
			}
		}
		if len(validNoteDataFlowShapes) != len(diagram.NoteDataFlowShapes) {
			diagram.NoteDataFlowShapes = validNoteDataFlowShapes
			needCommit = true
		}
	}

	return
}
