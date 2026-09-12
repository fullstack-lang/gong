// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"net/http"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
)

type Stager struct {
	stage       *Stage
	splitStage  *split.Stage
	asSplitArea *split.AsSplitArea
}

func NewStager(r *http.ServeMux, stage *Stage, splitStage *split.Stage) (stager *Stager) {

	stager = new(Stager)

	stager.stage = stage
	stager.splitStage = splitStage

	stager.asSplitArea = &split.AsSplitArea{}

	return
}

func (stager *Stager) GetAsSplitArea() (asSplitArea *split.AsSplitArea) {
	asSplitArea = stager.asSplitArea
	return
}
