package stage3d

import (
	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
)

type ThreeJSStageUpdater struct {
	// movie recording state
	isRecording             bool
	recordingRot            float64
	recordingFrameCount     int
	recordingPlant          *models.PlantAbstract
	savedInitCommitCallback models.OnInitCommitInterface
}

func NewThreeJSStageUpdater() *ThreeJSStageUpdater {
	return &ThreeJSStageUpdater{}
}

func (u *ThreeJSStageUpdater) UpdateThreeJSStage(stager *models.Stager) {
	u.ux_3d_plant_diagram(stager)
}

func (u *ThreeJSStageUpdater) StartMovieRecording(stager *models.Stager, plant *models.PlantAbstract, plantDiagram *models.PlantDiagram) {
	u.startMovieRecording(stager, plant, plantDiagram)
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
