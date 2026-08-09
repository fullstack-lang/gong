package vasestage3d

import (
	"encoding/base64"
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func (u *ThreeJSStageUpdater) startMovieRecording(stager *models.Stager, plant *models.PlantAbstract, plantDiagram *models.PlantDiagram) {
	if plant == nil {
		plant = stager.GetCurrentPlant()
	}
	if plant == nil {
		log.Println("[Movie Recorder] Cannot start recording: no plant selected")
		return
	}

	outputDir := "movie_frames"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Printf("[Movie Recorder] Failed to create directory %s: %v", outputDir, err)
		return
	}

	log.Printf("[Movie Recorder] Starting movie recording for plant %s from rot 0.0 to 1.0 (0.001 increment)...", plant.Name)
	u.isRecording = true
	u.recordingRot = 0.0
	u.recordingFrameCount = 0
	u.recordingPlant = plant

	// Save and disable automatic stage marshalling (disk writes) during frame recording
	u.savedInitCommitCallback = stager.GetStage().OnInitCommitCallback
	stager.GetStage().OnInitCommitCallback = nil

	if plant.PlantType == models.Vase {
		plant.VaseAbstract.RotationRatio = 0.0
	}
	stager.GetStage().Commit()
}

func (u *ThreeJSStageUpdater) stopMovieRecording(stager *models.Stager) {
	if !u.isRecording {
		return
	}
	log.Printf("[Movie Recorder] Stopping movie recording. Total frames saved: %d", u.recordingFrameCount)
	u.isRecording = false
	if u.savedInitCommitCallback != nil {
		stager.GetStage().OnInitCommitCallback = u.savedInitCommitCallback
		u.savedInitCommitCallback = nil
	}
	stager.GetStage().Commit()
}

func (u *ThreeJSStageUpdater) onCanvasFrameCaptured(stager *models.Stager, canvas *threejs.Canvas) {
	if !u.isRecording || u.recordingPlant == nil {
		return
	}

	currentRot := 0.0
	if u.recordingPlant.PlantType == models.Vase {
		currentRot = u.recordingPlant.VaseAbstract.RotationRatio
	}

	b64Data := canvas.Frame64BitsEncoded
	ext := ".png"
	if strings.HasPrefix(b64Data, "data:image/jpeg") {
		ext = ".jpg"
	}
	if idx := strings.Index(b64Data, ","); idx != -1 {
		b64Data = b64Data[idx+1:]
	}

	if len(b64Data) > 0 {
		imgBytes, err := base64.StdEncoding.DecodeString(b64Data)
		if err != nil {
			log.Printf("[Movie Recorder] Error decoding base64 image for frame %d: %v", u.recordingFrameCount, err)
		} else {
			fileName := fmt.Sprintf("movie_frames/frame_%04d%s", u.recordingFrameCount, ext)
			if err := os.WriteFile(fileName, imgBytes, 0644); err != nil {
				log.Printf("[Movie Recorder] Error writing frame %s: %v", fileName, err)
			} else {
				log.Printf("[Movie Recorder] Saved frame %d (rot: %.3f) -> %s (%d bytes)",
					u.recordingFrameCount, currentRot, fileName, len(imgBytes))
			}
		}
	} else {
		log.Printf("[Movie Recorder] Warning: frame %d (rot: %.3f) received empty Frame64BitsEncoded",
			u.recordingFrameCount, currentRot)
	}

	u.recordingFrameCount++

	nbFrames := 1000
	if u.recordingPlant.PlantType == models.Vase && u.recordingPlant.VaseAbstract.MovieNbFrames > 0 {
		nbFrames = u.recordingPlant.VaseAbstract.MovieNbFrames
	}
	rotIncrement := 1.0 / float64(nbFrames)
	u.recordingRot += rotIncrement
	nextRot := math.Round(u.recordingRot/rotIncrement) * rotIncrement

	if u.recordingFrameCount < nbFrames {
		if u.recordingPlant.PlantType == models.Vase {
			u.recordingPlant.VaseAbstract.RotationRatio = nextRot
		}
		stager.EnforceSemantic()
		u.ux_3d_plant_diagram(stager)
	} else {
		u.stopMovieRecording(stager)
	}
}
