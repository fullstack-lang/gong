package models

import (
	"log"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) svg() {
	log.Println("svg")
	stager.processDiagramSvgStage.Reset()

	stager.svgObject = new(svg.SVG).Stage(stager.processDiagramSvgStage)

	stager.svgObject.Name = "SVG"

	stager.processDiagramSvgStage.Commit()
}
