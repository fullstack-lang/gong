package models

import (
	"fmt"
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

		stager.logAndNotify(
			fmt.Sprintf("Root: setting nbPixPerCharacter to %f", root.NbPixPerCharacter))
	}

	for _, plant := range stager.stage.GetInstancesSorted[*PlantAbstract]() {
		if plant.N == 0 {
			needCommit = true
			plant.N = 1
			stager.logAndNotify(fmt.Sprintf("Plant %s: default N set to 1", plant.Name))
		}
		if plant.M == 0 {
			needCommit = true
			plant.M = 1
			stager.logAndNotify(fmt.Sprintf("Plant %s: default M set to 1", plant.Name))
		}
		if plant.RhombusInsideAngle == 0.0 {
			needCommit = true
			plant.RhombusInsideAngle = 65.0
			stager.logAndNotify(fmt.Sprintf("Plant %s: default RhombusInsideAngle set to 65.0", plant.Name))
		}
		if plant.PlantType == "" {
			needCommit = true
			plant.PlantType = Plant
			stager.logAndNotify(fmt.Sprintf("Plant %s: default PlantType set to Plant", plant.Name))
		}
		if plant.RhombusSideLength == 0.0 {
			needCommit = true
			plant.RhombusSideLength = 100.0
			stager.logAndNotify(fmt.Sprintf("Plant %s: default RhombusSideLength set to 100.0", plant.Name))
		}
		if plant.StackHeight == 0 {
			needCommit = true
			plant.StackHeight = 1
			stager.logAndNotify(fmt.Sprintf("Plant %s: default StackHeight set to 1", plant.Name))
		}
		if plant.Name == "" {
			needCommit = true
			plant.Name = "New Plant"
			stager.logAndNotify("Plant: empty name set to 'New Plant'")
		}
		if vase := plant.TubeVaseAbstract; vase != nil {
			if vase.RelativeVerticalThickness == 0.0 {
				needCommit = true
				vase.RelativeVerticalThickness = 0.1
				stager.logAndNotify(fmt.Sprintf("Plant %s TubeVase: default RelativeVerticalThickness set to 0.1", plant.Name))
			}
			if vase.RadialRepetitions < 1 {
				needCommit = true
				vase.RadialRepetitions = 1
				stager.logAndNotify(fmt.Sprintf("Plant %s TubeVase: default RadialRepetitions set to 1", plant.Name))
			}
			if vase.NbStepP1P2 <= 0 {
				needCommit = true
				vase.NbStepP1P2 = 10
				stager.logAndNotify(fmt.Sprintf("Plant %s TubeVase: default NbStepP1P2 set to 10", plant.Name))
			}
			if vase.WidthKey == 0.0 {
				needCommit = true
				vase.WidthKey = 30.0
				stager.logAndNotify(fmt.Sprintf("Plant %s TubeVase: default WidthKey set to 30.0", plant.Name))
			}
			if vase.HeightKey == 0.0 {
				needCommit = true
				vase.HeightKey = 50.0
				stager.logAndNotify(fmt.Sprintf("Plant %s TubeVase: default HeightKey set to 50.0", plant.Name))
			}
			if vase.RelativeKeySize == 0.0 {
				needCommit = true
				vase.RelativeKeySize = 0.20
				stager.logAndNotify(fmt.Sprintf("Plant %s TubeVase: default RelativeKeySize set to 0.20", plant.Name))
			}
			if vase.MovieNbFrames <= 0 {
				needCommit = true
				vase.MovieNbFrames = 1000
				stager.logAndNotify(fmt.Sprintf("Plant %s TubeVase: default MovieNbFrames set to 1000", plant.Name))
			}
		}
		if stool := plant.StoolAbstract; stool != nil {
			if stool.RadialRepetitions < 1 {
				needCommit = true
				stool.RadialRepetitions = 1
				stager.logAndNotify(fmt.Sprintf("Plant %s Stool: default RadialRepetitions set to 1", plant.Name))
			}
			if stool.RelativeTubeDiameter == 0.0 {
				needCommit = true
				stool.RelativeTubeDiameter = 0.01
				stager.logAndNotify(fmt.Sprintf("Plant %s Stool: default RelativeTubeDiameter set to 0.01", plant.Name))
			}
			if stool.RelativeHeight3DTorus == 0.0 {
				needCommit = true
				stool.RelativeHeight3DTorus = 1.0
				stager.logAndNotify(fmt.Sprintf("Plant %s Stool: default RelativeHeight3DTorus set to 1.0", plant.Name))
			}
			if stool.RelativeHeight == 0.0 {
				needCommit = true
				stool.RelativeHeight = 1.28
				stager.logAndNotify(fmt.Sprintf("Plant %s Stool: default RelativeHeight set to 1.28", plant.Name))
			}
			if stool.RelativeSeatThickness == 0.0 {
				needCommit = true
				stool.RelativeSeatThickness = 0.15
				stager.logAndNotify(fmt.Sprintf("Plant %s Stool: default RelativeSeatThickness set to 0.15", plant.Name))
			}
			if stool.RelativeEyeSeparationCriteria == 0.0 {
				needCommit = true
				stool.RelativeEyeSeparationCriteria = 0.05
				stager.logAndNotify(fmt.Sprintf("Plant %s Stool: default RelativeEyeSeparationCriteria set to 0.05", plant.Name))
			}
			if stool.RelativeEyeCornerControlVectorStrength == 0.0 {
				needCommit = true
				stool.RelativeEyeCornerControlVectorStrength = 0.55
				stager.logAndNotify(fmt.Sprintf("Plant %s Stool: default RelativeEyeCornerControlVectorStrength set to 0.55", plant.Name))
			}
			if stool.StoolTorusVerticalScale == 0.0 {
				needCommit = true
				stool.StoolTorusVerticalScale = 1.0
				stager.logAndNotify(fmt.Sprintf("Plant %s Stool: default StoolTorusVerticalScale set to 1.0", plant.Name))
			}
		}
	}

	for _, plant2DDiagram := range stager.stage.GetInstancesSorted[*Plant2DDiagram]() {
		if plant2DDiagram.Name == "" {
			needCommit = true
			plant2DDiagram.Name = "New Plant Diagram"
			stager.logAndNotify("Plant2DDiagram: empty name set to 'New Plant Diagram'")
		}
		if plant2DDiagram.OriginX == 0.0 {
			needCommit = true
			plant2DDiagram.OriginX = 280.000000
			stager.logAndNotify(fmt.Sprintf("Plant2DDiagram %s: default OriginX set to 280.0", plant2DDiagram.Name))
		}
		if plant2DDiagram.OriginY == 0.0 {
			needCommit = true
			plant2DDiagram.OriginY = 950.000000
			stager.logAndNotify(fmt.Sprintf("Plant2DDiagram %s: default OriginY set to 950.0", plant2DDiagram.Name))
		}
		if plant2DDiagram.Zoom == 0.0 {
			needCommit = true
			plant2DDiagram.Zoom = 1.0
			stager.logAndNotify(fmt.Sprintf("Plant2DDiagram %s: default Zoom set to 1.0", plant2DDiagram.Name))
		}
	}
	for _, vase2DDiagram := range stager.stage.GetInstancesSorted[*Vase2DDiagram]() {
		if vase2DDiagram.Zoom == 0.0 {
			needCommit = true
			vase2DDiagram.Zoom = 1.0
			stager.logAndNotify(fmt.Sprintf("Vase2DDiagram %s: default Zoom set to 1.0", vase2DDiagram.Name))
		}
	}
	for _, stool2DDiagram := range stager.stage.GetInstancesSorted[*Stool2DDiagram]() {
		if stool2DDiagram.Zoom == 0.0 {
			needCommit = true
			stool2DDiagram.Zoom = 1.0
			stager.logAndNotify(fmt.Sprintf("Stool2DDiagram %s: default Zoom set to 1.0", stool2DDiagram.Name))
		}
	}
	for _, clock2DDiagram := range stager.stage.GetInstancesSorted[*Clock2DDiagram]() {
		if clock2DDiagram.Zoom == 0.0 {
			needCommit = true
			clock2DDiagram.Zoom = 1.0
			stager.logAndNotify(fmt.Sprintf("Clock2DDiagram %s: default Zoom set to 1.0", clock2DDiagram.Name))
		}
	}
	for _, axesShape := range stager.stage.GetInstancesSorted[*AxesShape]() {
		if axesShape.LengthX == 0.0 {
			needCommit = true
			axesShape.LengthX = 200.0
			stager.logAndNotify(fmt.Sprintf("AxesShape %s: default LengthX set to 200.0", axesShape.Name))
		}
		if axesShape.LengthY == 0.0 {
			needCommit = true
			axesShape.LengthY = 200.0
			stager.logAndNotify(fmt.Sprintf("AxesShape %s: default LengthY set to 200.0", axesShape.Name))
		}
	}

	return
}
