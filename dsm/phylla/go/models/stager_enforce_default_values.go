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

	for _, plant := range GetGongstrucsSorted[*Plant](stager.stage) {
		if plant.N == 0 {
			needCommit = true
			plant.N = 1
		}
		if plant.M == 0 {
			needCommit = true
			plant.M = 1
		}
		if plant.RhombusSideLength == 0.0 {
			needCommit = true
			plant.RhombusSideLength = 100.0
		}
		if plant.RhombusInsideAngle == 0.0 {
			needCommit = true
			plant.RhombusInsideAngle = 65.0
		}
		if plant.Thickness == 0.0 {
			needCommit = true
			plant.Thickness = 10.0
		}
		if plant.Name == "" {
			needCommit = true
			plant.Name = "New Plant"
		}
	}
	for _, plantDiagram := range GetGongstrucsSorted[*PlantDiagram](stager.stage) {
		if plantDiagram.Name == "" {
			needCommit = true
			plantDiagram.Name = "New Plant Diagram"
		}
		if plantDiagram.OriginX == 0.0 {
			needCommit = true
			plantDiagram.OriginX = 280.000000
		}
		if plantDiagram.OriginY == 0.0 {
			needCommit = true
			plantDiagram.OriginY = 950.000000
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
