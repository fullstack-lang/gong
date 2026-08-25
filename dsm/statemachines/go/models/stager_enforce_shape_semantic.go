package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceShapeSemantic() (needCommit bool) {
	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {
		for _, stateShape := range diagram.State_Shapes {
			if stateShape.State == nil {
				stateShape.Unstage(stager.stage)
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), "Unstaged StateShape with nil State")
				}
				continue
			}

			if stateShape.Name != stateShape.State.Name {
				stateShape.Name = stateShape.State.Name
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(),
						fmt.Sprintf("Synchronized StateShape name to \"%s\"", stateShape.Name))
				}
			}
		}

		for _, transitionShape := range diagram.Transition_Shapes {
			if transitionShape.Transition == nil {
				transitionShape.Unstage(stager.stage)
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), "Unstaged Transition_Shape with nil Transition")
				}
			}
		}

		for _, noteShape := range diagram.Note_Shapes {
			if noteShape.Note == nil {
				noteShape.Unstage(stager.stage)
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), "Unstaged NoteShape with nil Note")
				}
				continue
			}

			expectedName := noteShape.Note.Name + "-" + diagram.Name
			if noteShape.Name != expectedName {
				noteShape.Name = expectedName
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(),
						fmt.Sprintf("Synchronized NoteShape name to \"%s\"", expectedName))
				}
			}
		}

		for _, noteStateShape := range diagram.NoteState_Shapes {
			if noteStateShape.Note == nil || noteStateShape.State == nil {
				noteStateShape.Unstage(stager.stage)
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), "Unstaged NoteStateShape with nil Note or State")
				}
			}
		}

		for _, noteTransitionShape := range diagram.NoteTransition_Shapes {
			if noteTransitionShape.Note == nil || noteTransitionShape.Transition == nil {
				noteTransitionShape.Unstage(stager.stage)
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), "Unstaged NoteTransitionShape with nil Note or Transition")
				}
			}
		}
	}

	return
}
