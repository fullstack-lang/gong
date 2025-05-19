// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"github.com/gin-gonic/gin"

	split "github.com/fullstack-lang/gong/lib/split/go/models"

	svg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
	svg_stack "github.com/fullstack-lang/gong/lib/svg/go/stack"
)

type Stager struct {
	stage       *Stage
	splitStage  *split.Stage
	asSplitArea *split.AsSplitArea
	svgStage    *svg_models.Stage
}

func NewStager(r *gin.Engine, stage *Stage, splitStage *split.Stage) (stager *Stager) {

	stager = new(Stager)

	stager.stage = stage
	stager.splitStage = splitStage

	stacksvg := svg_stack.NewStack(r, stage.GetName(), "", "", "", true, true)
	stacksvg.Probe.Refresh()

	stager.svgStage = stacksvg.Stage

	stager.asSplitArea = &split.AsSplitArea{
		Svg: &split.Svg{
			StackName: stager.GetSvgStage().GetName(),
		},
	}

	return
}

func (stager *Stager) GetAsSplitArea() (asSplitArea *split.AsSplitArea) {
	asSplitArea = stager.asSplitArea
	return
}

func (stager *Stager) GetSvgStage() (svgStage *svg_models.Stage) {
	svgStage = stager.svgStage
	return
}
