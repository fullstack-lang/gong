package models

import (
	"log"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) onUpdateSVG(frontSVG *svg.SVG) {
	// log.Println("SVG updated", updatedSVG.GetName())
	diagram := stager.diagram
	svgObject := stager.svgObject

	if svgObject.DrawingState == frontSVG.DrawingState {
		// in any cases, have the form editor set up with the instance
		stager.probeForm.FillUpFormFromGongstruct(diagram, "Diagram")
		return
	}

	// IMPORTANT : we are only interested when the updateSVG has finished drawing the connexion
	if frontSVG.DrawingState == svg.DRAWING_LINK {
		return
	}

	startStateShape, isStartState := stager.map_SvgRect_StateShape[frontSVG.StartRect]
	startNoteShape, isStartNote := stager.map_SvgRect_NoteShape[frontSVG.StartRect]

	if !isStartState && !isStartNote {
		log.Panic(frontSVG.StartRect.Name + " unknown actor state or note shape")
	}

	endStateShape, isEndState := stager.map_SvgRect_StateShape[frontSVG.EndRect]
	if !isEndState {
		// Only supporting links ending in States for now via drag and drop
		log.Panic(frontSVG.EndRect.Name + " unknown actor state shape")
	}

	if isStartState && isEndState {
		transition := new(Transition).Stage(stager.stage)

		transition.Name = ""
		if stateMachine, ok := stager.map_diagram_stateMachine[stager.diagram]; ok {
			if stateMachine.IsWithTransitionNameAutonamticalyGenerated {
				transition.Name = startStateShape.State.Name + " to " + endStateShape.State.Name
			}
		}

		transition.Start = startStateShape.State
		transition.End = endStateShape.State

		transitionShape := new(Transition_Shape).Stage(stager.stage)
		transitionShape.Name = transition.Name + "-" + stager.diagram.Name
		transitionShape.Transition = transition
		transitionShape.StartOrientation = ORIENTATION_HORIZONTAL
		transitionShape.EndOrientation = ORIENTATION_HORIZONTAL
		transitionShape.CornerOffsetRatio = 1.2
		transitionShape.StartRatio = 0.5
		transitionShape.EndRatio = 0.5

		stager.diagram.Transition_Shapes =
			append(stager.diagram.Transition_Shapes, transitionShape)

		stager.probeForm.FillUpFormFromGongstruct(transition, "Transition")
	} else if isStartNote && isEndState {
		noteStateShape := new(NoteStateShape).Stage(stager.stage)
		noteStateShape.Name = startNoteShape.Note.Name + " to " + endStateShape.State.Name
		noteStateShape.Note = startNoteShape.Note
		noteStateShape.State = endStateShape.State
		noteStateShape.StartOrientation = ORIENTATION_HORIZONTAL
		noteStateShape.EndOrientation = ORIENTATION_HORIZONTAL
		noteStateShape.CornerOffsetRatio = 1.2
		noteStateShape.StartRatio = 0.5
		noteStateShape.EndRatio = 0.5

		stager.diagram.NoteState_Shapes =
			append(stager.diagram.NoteState_Shapes, noteStateShape)

		// also add the state to the note's state links
		startNoteShape.Note.States = append(startNoteShape.Note.States, endStateShape.State)

		stager.probeForm.FillUpFormFromGongstruct(noteStateShape, "NoteStateShape")
	}

	// commit to encode the result, this will generate a new SVG generation
	stager.stage.Commit()

}
