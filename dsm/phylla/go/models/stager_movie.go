package models

import (
	"encoding/base64"
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func (stager *Stager) startMovieRecording(plant *PlantAbstract, plantDiagram *PlantDiagram) {
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
	stager.isRecording = true
	stager.recordingRot = 0.0
	stager.recordingFrameCount = 0
	stager.recordingPlant = plant

	// Save and disable automatic stage marshalling (disk writes) during frame recording
	stager.savedInitCommitCallback = stager.stage.OnInitCommitCallback
	stager.stage.OnInitCommitCallback = nil

	if plant.PlantType == Vase {
		plant.VaseAbstract.RotationRatio = 0.0
	}
	stager.stage.Commit()
}

func (stager *Stager) stopMovieRecording() {
	if !stager.isRecording {
		return
	}
	log.Printf("[Movie Recorder] Stopping movie recording. Total frames saved: %d", stager.recordingFrameCount)
	stager.isRecording = false
	if stager.savedInitCommitCallback != nil {
		stager.stage.OnInitCommitCallback = stager.savedInitCommitCallback
		stager.savedInitCommitCallback = nil
	}
	stager.stage.Commit()
}

func (stager *Stager) onCanvasFrameCaptured(canvas *threejs.Canvas) {
	if !stager.isRecording || stager.recordingPlant == nil {
		return
	}

	currentRot := 0.0
	if stager.recordingPlant.PlantType == Vase {
		currentRot = stager.recordingPlant.VaseAbstract.RotationRatio
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
			log.Printf("[Movie Recorder] Error decoding base64 image for frame %d: %v", stager.recordingFrameCount, err)
		} else {
			fileName := fmt.Sprintf("movie_frames/frame_%04d%s", stager.recordingFrameCount, ext)
			if err := os.WriteFile(fileName, imgBytes, 0644); err != nil {
				log.Printf("[Movie Recorder] Error writing frame %s: %v", fileName, err)
			} else {
				log.Printf("[Movie Recorder] Saved frame %d (rot: %.3f) -> %s (%d bytes)",
					stager.recordingFrameCount, currentRot, fileName, len(imgBytes))
			}
		}
	} else {
		log.Printf("[Movie Recorder] Warning: frame %d (rot: %.3f) received empty Frame64BitsEncoded",
			stager.recordingFrameCount, currentRot)
	}

	stager.recordingFrameCount++

	nbFrames := 1000
	if stager.recordingPlant.PlantType == Vase && stager.recordingPlant.VaseAbstract.MovieNbFrames > 0 {
		nbFrames = stager.recordingPlant.VaseAbstract.MovieNbFrames
	}
	rotIncrement := 1.0 / float64(nbFrames)
	stager.recordingRot += rotIncrement
	nextRot := math.Round(stager.recordingRot/rotIncrement) * rotIncrement

	if stager.recordingFrameCount < nbFrames {
		if stager.recordingPlant.PlantType == Vase {
			stager.recordingPlant.VaseAbstract.RotationRatio = nextRot
		}
		stager.enforceSemantic()
		stager.ux_3d_plant_diagram()
	} else {
		stager.stopMovieRecording()
	}
}
