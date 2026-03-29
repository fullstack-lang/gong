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

	startStateShape, ok := stager.map_SvgRect_StateShape[frontSVG.StartRect]
	if !ok {
		log.Panic(frontSVG.StartRect.Name + "unknow actor state shape")
	}

	endStateShape, ok := stager.map_SvgRect_StateShape[frontSVG.EndRect]
	if !ok {
		log.Panic(frontSVG.StartRect.Name + "unknow actor state shape")
	}

	transition := new(Transition).Stage(stager.stage)
	transition.Name = startStateShape.State.Name + " to " + endStateShape.State.Name
	transition.Start = startStateShape.State
	transition.End = endStateShape.State

	transitionShape := new(Transition_Shape).Stage(stager.stage)
	transitionShape.Name = startStateShape.State.Name + " to " + endStateShape.State.Name
	transitionShape.Transition = transition
	transitionShape.StartOrientation = ORIENTATION_HORIZONTAL
	transitionShape.EndOrientation = ORIENTATION_HORIZONTAL
	transitionShape.CornerOffsetRatio = 1.2
	transitionShape.StartRatio = 0.5
	transitionShape.EndRatio = 0.5

	stager.diagram.Transition_Shapes =
		append(stager.diagram.Transition_Shapes, transitionShape)

	stager.probeForm.FillUpFormFromGongstruct(transition, "Transition")

	// commit to encode the result, this will generate a new SVG generation
	stager.stage.Commit()

}
