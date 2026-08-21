package models

import (
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// SVGUpdated implements SVGImplInterface.
func (stager *Stager) onUpdateSVG(frontSVG *svg.SVG) {
	diagramFloss := stager.diagramFloss
	svgObjectDiagramFloss := stager.svgObjectDiagramFloss

	if svgObjectDiagramFloss.DrawingState == frontSVG.DrawingState {
		// in any cases, have the form editor set up with the instance
		stager.probeForm.FillUpFormFromGongstruct(diagramFloss, "Diagram")
		return
	}

	// commit to encode the result, this will generate a new SVG generation
	stager.stage.Commit()
}
