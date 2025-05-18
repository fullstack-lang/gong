// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"github.com/gin-gonic/gin"

	split "github.com/fullstack-lang/gong/lib/split/go/models"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	svg_stack "github.com/fullstack-lang/gong/lib/svg/go/stack"
)

type Stager struct {
	stage       *Stage
	splitStage  *split.Stage
	svgStage    *svg.Stage
	asSplitArea *split.AsSplitArea
}

func NewStager(r *gin.Engine, stage *Stage, splitStage *split.Stage) (stager *Stager) {

	stager = new(Stager)

	stager.stage = stage
	stager.splitStage = splitStage

	// creates a svg stack and display the front into the split
	svgPersistanceFile := "svg.go"
	stager.svgStage = svg_stack.NewStack(r, stage.GetName(), svgPersistanceFile, svgPersistanceFile, "", true, true).Stage

	stager.asSplitArea = &split.AsSplitArea{}

	return
}

func (stager *Stager) GetAsSplitArea() (asSplitArea *split.AsSplitArea) {
	asSplitArea = stager.asSplitArea
	return
}
