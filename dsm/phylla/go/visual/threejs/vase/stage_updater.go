package vase

import (
	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

type ThreeJSStageUpdater struct {
	// movie recording state
	isRecording             bool
	recordingRot            float64
	recordingFrameCount     int
	recordingPlant          *models.PlantAbstract
	savedInitCommitCallback models.OnInitCommitInterface

	// active cutoff plane and scene bounds for fast slider updates
	clippingPlane *threejs.ClippingPlane
	boundsMinX    float64
	boundsMaxX    float64
	boundsMinY    float64
	boundsMaxY    float64
	boundsMinZ    float64
	boundsMaxZ    float64
	hasBounds     bool
}

func NewThreeJSStageUpdater() *ThreeJSStageUpdater {
	return &ThreeJSStageUpdater{}
}

func (u *ThreeJSStageUpdater) UpdateThreeJSStage(stager *models.Stager) {
	u.ux_3d_plant_diagram(stager)
}

func (u *ThreeJSStageUpdater) UpdateCutoffPlane(stager *models.Stager) {
	u.updateCutoffPlane(stager)
}

func (u *ThreeJSStageUpdater) StartMovieRecordingTubeVase3D(stager *models.Stager, plant *models.PlantAbstract, diagram *models.TubeVase3DDiagram) {
	u.startMovieRecording(stager, plant, diagram)
}

func (u *ThreeJSStageUpdater) StopMovieRecording(stager *models.Stager) {
	u.stopMovieRecording(stager)
}

func (u *ThreeJSStageUpdater) IsMovieRecording() bool {
	return u.isRecording
}

func (u *ThreeJSStageUpdater) GetMovieRecordingFrameCount() int {
	return u.recordingFrameCount
}
