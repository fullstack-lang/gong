package models

import (
	"log"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

type svgProxy struct {
	stager                 *Stager
	diagram                *Diagram
	map_SvgRect_StateShape map[*svg.Rect]*StateShape
}

// SVGUpdated implements SVGImplInterface.
func (p *svgProxy) SVGUpdated(updatedSVG *svg.SVG) {
	// log.Println("SVG updated", updatedSVG.GetName())

	// if the user is link two actor state shapes, get them
	if updatedSVG.DrawingState == svg.DRAWING_LINK {
		startStateShape, ok := p.map_SvgRect_StateShape[updatedSVG.StartRect]
		if !ok {
			log.Panic(updatedSVG.StartRect.Name + "unknow actor state shape")
		}

		endStateShape, ok := p.map_SvgRect_StateShape[updatedSVG.EndRect]
		if !ok {
			log.Panic(updatedSVG.StartRect.Name + "unknow actor state shape")
		}

		transition := new(Transition).Stage(p.stager.stage)
		transition.Name = startStateShape.State.Name + " to " + endStateShape.State.Name
		transition.Start = startStateShape.State
		transition.End = endStateShape.State

		transitionShape := new(Transition_Shape).Stage(p.stager.stage)
		transitionShape.Name = startStateShape.State.Name + " to " + endStateShape.State.Name
		transitionShape.Transition = transition
		transitionShape.StartOrientation = ORIENTATION_HORIZONTAL
		transitionShape.EndOrientation = ORIENTATION_HORIZONTAL
		transitionShape.CornerOffsetRatio = 1.2
		transitionShape.StartRatio = 0.5
		transitionShape.EndRatio = 0.5

		p.diagram.Transition_Shapes =
			append(p.diagram.Transition_Shapes, transitionShape)

		p.stager.probeForm.FillUpFormFromGongstruct(transition, "Transition")

		// commit to encode the result, this will generate a new SVG generation
		p.stager.stage.Commit()
	} else {
		// in any cases, have the form editor set up with the instance
		p.stager.probeForm.FillUpFormFromGongstruct(p.diagram, "Diagram")
	}
}

func (p *svgProxy) OnAfterUpdate(stage *svg.Stage, stagedSVG, frontSVG *svg.SVG) {

	// log.Println("SVG updated", stagedSVG.GetName())

}
