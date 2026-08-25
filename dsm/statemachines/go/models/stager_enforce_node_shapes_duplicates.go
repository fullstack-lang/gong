package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceNodeShapeDuplicates() (needCommit bool) {
	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {
		needCommit = removeDuplicateStateShapes(stager, diagram) || needCommit
		needCommit = removeDuplicateTransitionShapes(stager, diagram) || needCommit
		needCommit = removeDuplicateNoteShapes(stager, diagram) || needCommit
		needCommit = removeDuplicateNoteStateShapes(stager, diagram) || needCommit
		needCommit = removeDuplicateNoteTransitionShapes(stager, diagram) || needCommit
	}
	return
}

func removeDuplicateStateShapes(stager *Stager, diagram *Diagram) (needCommit bool) {
	seen := make(map[*State]bool)
	var newShapes []*StateShape

	for _, shape := range diagram.State_Shapes {
		if shape.State != nil {
			if seen[shape.State] {
				shape.Unstage(stager.stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged duplicate StateShape \"%s\"", shape.Name))
				}
				needCommit = true
				continue
			}
			seen[shape.State] = true
		}
		newShapes = append(newShapes, shape)
	}
	if needCommit {
		diagram.State_Shapes = newShapes
	}
	return
}

func removeDuplicateTransitionShapes(stager *Stager, diagram *Diagram) (needCommit bool) {
	seen := make(map[*Transition]bool)
	var newShapes []*Transition_Shape

	for _, shape := range diagram.Transition_Shapes {
		if shape.Transition != nil {
			if seen[shape.Transition] {
				shape.Unstage(stager.stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged duplicate Transition_Shape \"%s\"", shape.Name))
				}
				needCommit = true
				continue
			}
			seen[shape.Transition] = true
		}
		newShapes = append(newShapes, shape)
	}
	if needCommit {
		diagram.Transition_Shapes = newShapes
	}
	return
}

func removeDuplicateNoteShapes(stager *Stager, diagram *Diagram) (needCommit bool) {
	seen := make(map[*Note]bool)
	var newShapes []*NoteShape

	for _, shape := range diagram.Note_Shapes {
		if shape.Note != nil {
			if seen[shape.Note] {
				shape.Unstage(stager.stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged duplicate NoteShape \"%s\"", shape.Name))
				}
				needCommit = true
				continue
			}
			seen[shape.Note] = true
		}
		newShapes = append(newShapes, shape)
	}
	if needCommit {
		diagram.Note_Shapes = newShapes
	}
	return
}

type noteStateKey struct {
	note  *Note
	state *State
}

func removeDuplicateNoteStateShapes(stager *Stager, diagram *Diagram) (needCommit bool) {
	seen := make(map[noteStateKey]bool)
	var newShapes []*NoteStateShape

	for _, shape := range diagram.NoteState_Shapes {
		if shape.Note != nil && shape.State != nil {
			key := noteStateKey{note: shape.Note, state: shape.State}
			if seen[key] {
				shape.Unstage(stager.stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged duplicate NoteStateShape \"%s\"", shape.Name))
				}
				needCommit = true
				continue
			}
			seen[key] = true
		}
		newShapes = append(newShapes, shape)
	}
	if needCommit {
		diagram.NoteState_Shapes = newShapes
	}
	return
}

type noteTransitionKey struct {
	note       *Note
	transition *Transition
}

func removeDuplicateNoteTransitionShapes(stager *Stager, diagram *Diagram) (needCommit bool) {
	seen := make(map[noteTransitionKey]bool)
	var newShapes []*NoteTransitionShape

	for _, shape := range diagram.NoteTransition_Shapes {
		if shape.Note != nil && shape.Transition != nil {
			key := noteTransitionKey{note: shape.Note, transition: shape.Transition}
			if seen[key] {
				shape.Unstage(stager.stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged duplicate NoteTransitionShape \"%s\"", shape.Name))
				}
				needCommit = true
				continue
			}
			seen[key] = true
		}
		newShapes = append(newShapes, shape)
	}
	if needCommit {
		diagram.NoteTransition_Shapes = newShapes
	}
	return
}
