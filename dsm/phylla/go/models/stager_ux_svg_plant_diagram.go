package models

import (
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) ux_svg_plant_diagram() {

	stager.svgStage.Reset()

	var plantDiagram *PlantDiagram
	{
		for d_ := range *GetGongstructInstancesSet[PlantDiagram](stager.stage) {
			if d_.IsChecked {
				plantDiagram = d_
			}
		}
	}

	if plantDiagram == nil {
		stager.svgStage.Commit()
		return
	}
	svgObject := stager.generateSvgObject(plantDiagram)

	svg.StageBranch(stager.svgStage, svgObject)
	stager.svgObject = svgObject

	stager.svgStage.Commit()
}

func (stager *Stager) generateSvgObject(plantDiagram *PlantDiagram) (svg_ *svg.SVG) {

	return
}
