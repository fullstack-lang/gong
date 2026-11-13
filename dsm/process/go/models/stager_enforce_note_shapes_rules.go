package models

import (
	"fmt"
	"slices"
	"time"
)

func (stager *Stager) enforceNoteShapesRules() (needCommit bool) {
	for _, diagram := range stager.stage.GetInstancesSorted[*DiagramProcess]() {
		seenTaskKeys := make(map[noteTaskKey]bool)
		var validNoteTaskShapes []*NoteTaskShape
		for _, shape := range diagram.NoteTaskShapes {
			key := noteTaskKey{Note: shape.Note, Task: shape.Task}
			isValid := shape.Note != nil &&
				shape.Task != nil &&
				slices.Contains(shape.Note.Tasks, shape.Task) &&
				!seenTaskKeys[key]
			if isValid {
				seenTaskKeys[key] = true
				validNoteTaskShapes = append(validNoteTaskShapes, shape)
			} else {
				shape.UnstageVoid(stager.stage)
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged note task shape %s (invalid or duplicate link)", shape.GetName()))
				}
			}
		}
		if len(validNoteTaskShapes) != len(diagram.NoteTaskShapes) {
			diagram.NoteTaskShapes = validNoteTaskShapes
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

		seenControlFlowKeys := make(map[noteControlFlowKey]bool)
		var validNoteControlFlowShapes []*NoteControlFlowShape
		for _, shape := range diagram.NoteControlFlowShapes {
			key := noteControlFlowKey{Note: shape.Note, ControlFlow: shape.ControlFlow}
			isValid := shape.Note != nil &&
				shape.ControlFlow != nil &&
				slices.Contains(shape.Note.ControlFlows, shape.ControlFlow) &&
				!seenControlFlowKeys[key]
			if isValid {
				seenControlFlowKeys[key] = true
				validNoteControlFlowShapes = append(validNoteControlFlowShapes, shape)
			} else {
				shape.UnstageVoid(stager.stage)
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged note control flow shape %s (invalid or duplicate link)", shape.GetName()))
				}
			}
		}
		if len(validNoteControlFlowShapes) != len(diagram.NoteControlFlowShapes) {
			diagram.NoteControlFlowShapes = validNoteControlFlowShapes
			needCommit = true
		}
	}

	return
}
