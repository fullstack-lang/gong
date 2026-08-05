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

func (stager *Stager) startMovieRecording(plant *Plant, plantDiagram *PlantDiagram) {
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

	plant.RotationRatio = 0.0
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

	b64Data := canvas.Frame64BitsEncoded
	if idx := strings.Index(b64Data, ","); idx != -1 {
		b64Data = b64Data[idx+1:]
	}

	if len(b64Data) > 0 {
		pngBytes, err := base64.StdEncoding.DecodeString(b64Data)
		if err != nil {
			log.Printf("[Movie Recorder] Error decoding base64 image for frame %d: %v", stager.recordingFrameCount, err)
		} else {
			fileName := fmt.Sprintf("movie_frames/frame_%04d.png", stager.recordingFrameCount)
			if err := os.WriteFile(fileName, pngBytes, 0644); err != nil {
				log.Printf("[Movie Recorder] Error writing frame %s: %v", fileName, err)
			} else {
				log.Printf("[Movie Recorder] Saved frame %d (rot: %.3f) -> %s (%d bytes)",
					stager.recordingFrameCount, stager.recordingPlant.RotationRatio, fileName, len(pngBytes))
			}
		}
	} else {
		log.Printf("[Movie Recorder] Warning: frame %d (rot: %.3f) received empty Frame64BitsEncoded",
			stager.recordingFrameCount, stager.recordingPlant.RotationRatio)
	}

	stager.recordingFrameCount++
	stager.recordingRot += 0.001
	nextRot := math.Round(stager.recordingRot*1000.0) / 1000.0

	if nextRot <= 1.0000001 {
		stager.recordingPlant.RotationRatio = nextRot
		stager.enforceSemantic()
		stager.ux_3d_plant_diagram()
	} else {
		stager.stopMovieRecording()
	}
}
