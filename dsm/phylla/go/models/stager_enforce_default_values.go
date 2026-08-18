package models

import (
	"fmt"
	"time"
)

// enforceDefaultValues enforce defaut values when there are not suitable
func (stager *Stager) enforceDefaultValues() (needCommit bool) {
	const (
		defaultBoxWidth  = 250.0
		defaultBoxHeigth = 70.0
	)

	root := stager.getRootLibrary()
	if root.NbPixPerCharacter == 0 {
		root.NbPixPerCharacter = 8
		needCommit = true

		stager.probeForm.AddNotification(time.Now(),
			fmt.Sprintf("Root: setting nbPixPerCharacter to %f", root.NbPixPerCharacter))
	}

	for _, plant := range GetGongstrucsSorted[*PlantAbstract](stager.stage) {
		if plant.N == 0 {
			needCommit = true
			plant.N = 1
		}
		if plant.M == 0 {
			needCommit = true
			plant.M = 1
		}
		if plant.RhombusInsideAngle == 0.0 {
			needCommit = true
			plant.RhombusInsideAngle = 65.0
		}
		if plant.PlantType == "" {
			needCommit = true
			plant.PlantType = Plant
		}
		if plant.RhombusSideLength == 0.0 {
			needCommit = true
			plant.RhombusSideLength = 100.0
		}
		if plant.Name == "" {
			needCommit = true
			plant.Name = "New Plant"
		}
		if vase := plant.VaseAbstract; vase != nil {
			if vase.RelativeVerticalThickness == 0.0 {
				needCommit = true
				vase.RelativeVerticalThickness = 0.1
			}
			if vase.RadialRepetitions < 1 {
				needCommit = true
				vase.RadialRepetitions = 1
			}
			if vase.NbStepP1P2 <= 0 {
				needCommit = true
				vase.NbStepP1P2 = 10
			}
			if vase.WidthKey == 0.0 {
				needCommit = true
				vase.WidthKey = 30.0
			}
			if vase.HeightKey == 0.0 {
				needCommit = true
				vase.HeightKey = 50.0
			}
			if vase.RelativeKeySize == 0.0 {
				needCommit = true
				vase.RelativeKeySize = 0.20
			}
			if vase.MovieNbFrames <= 0 {
				needCommit = true
				vase.MovieNbFrames = 1000
			}
		}
		if stool := plant.StoolAbstract; stool != nil {
			if stool.RadialRepetitions < 1 {
				needCommit = true
				stool.RadialRepetitions = 1
			}
			if stool.RelativeTubeDiameter == 0.0 {
				needCommit = true
				stool.RelativeTubeDiameter = 0.01
			}
			if stool.RelativeHeight3DTorus == 0.0 {
				needCommit = true
				stool.RelativeHeight3DTorus = 1.0
			}
			if stool.RelativeHeight == 0.0 {
				needCommit = true
				stool.RelativeHeight = 1.28
			}
			if stool.RelativeSeatThickness == 0.0 {
				needCommit = true
				stool.RelativeSeatThickness = 0.15
			}
			if stool.RelativeEyeSeparationCriteria == 0.0 {
				needCommit = true
				stool.RelativeEyeSeparationCriteria = 0.05
			}
			if stool.RelativeEyeCornerControlVectorStrength == 0.0 {
				needCommit = true
				stool.RelativeEyeCornerControlVectorStrength = 0.55
			}
			if stool.StoolTorusVerticalScale == 0.0 {
				needCommit = true
				stool.StoolTorusVerticalScale = 1.0
			}
		}
	}

	for _, plant2DDiagram := range GetGongstrucsSorted[*Plant2DDiagram](stager.stage) {
		if plant2DDiagram.Name == "" {
			needCommit = true
			plant2DDiagram.Name = "New Plant Diagram"
		}
		if plant2DDiagram.OriginX == 0.0 {
			needCommit = true
			plant2DDiagram.OriginX = 280.000000
		}
		if plant2DDiagram.OriginY == 0.0 {
			needCommit = true
			plant2DDiagram.OriginY = 950.000000
		}
		if plant2DDiagram.Zoom == 0.0 {
			needCommit = true
			plant2DDiagram.Zoom = 1.0
		}
	}
	for _, vase2DDiagram := range GetGongstrucsSorted[*Vase2DDiagram](stager.stage) {
		if vase2DDiagram.Zoom == 0.0 {
			needCommit = true
			vase2DDiagram.Zoom = 1.0
		}
	}
	for _, stool2DDiagram := range GetGongstrucsSorted[*Stool2DDiagram](stager.stage) {
		if stool2DDiagram.Zoom == 0.0 {
			needCommit = true
			stool2DDiagram.Zoom = 1.0
		}
	}
	for _, clock2DDiagram := range GetGongstrucsSorted[*Clock2DDiagram](stager.stage) {
		if clock2DDiagram.Zoom == 0.0 {
			needCommit = true
			clock2DDiagram.Zoom = 1.0
		}
	}
	for _, axesShape := range GetGongstrucsSorted[*AxesShape](stager.stage) {
		if axesShape.LengthX == 0.0 {
			needCommit = true
			axesShape.LengthX = 200.0
		}
		if axesShape.LengthY == 0.0 {
			needCommit = true
			axesShape.LengthY = 200.0
		}
	}

	return
}
