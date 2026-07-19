package models

import "fmt"

func (stager *Stager) enforcePlantDiagramDiscreteTorusStackShape() bool {
	modified := false

	for plant := range stager.stage.Plants {
		circumference := 10.0
		if plant.RhombusStuff != nil && plant.RhombusStuff.PlantCircumferenceShape != nil && plant.RhombusStuff.PlantCircumferenceShape.Length > 0 {
			circumference = plant.RhombusStuff.PlantCircumferenceShape.Length
		} else if pGrid := plant.PerpendicularVectorGrid; pGrid != nil && len(pGrid.PerpendicularVectors) > 0 {
			first := pGrid.PerpendicularVectors[0]
			last := pGrid.PerpendicularVectors[len(pGrid.PerpendicularVectors)-1]
			circumference = last.StartX - first.StartX
		}
		if circumference <= 0 {
			circumference = 10.0
		}
		
		R := circumference / (2 * 3.14159265358979323846)

		tubeRadius := plant.RelativeVerticalThickness * plant.RhombusSideLength / 2.0
		if tubeRadius <= 0 {
			tubeRadius = 0.5
		}

		for _, plantDiagram := range plant.PlantDiagrams {
			if plantDiagram.DiscreteTorusStackShape == nil {
				shape := (&DiscreteTorusStackShape{
					Name:    plantDiagram.Name + "-DiscreteTorusStackShape",
				}).Stage(stager.stage)
				plantDiagram.DiscreteTorusStackShape = shape
				modified = true
			}

			// Synchronize DiscreteTorusShapes
			stackShape := plantDiagram.DiscreteTorusStackShape
			expectedNum := plant.StackHeight
			if expectedNum < 0 {
				expectedNum = 0
			}

			if len(stackShape.DiscreteTorusShapes) != expectedNum {
				// Clear old shapes
				for _, t := range stackShape.DiscreteTorusShapes {
					t.Unstage(stager.stage)
				}
				stackShape.DiscreteTorusShapes = make([]*DiscreteTorusShape, expectedNum)
				for i := 0; i < expectedNum; i++ {
					t := (&DiscreteTorusShape{
						Name: stackShape.Name + fmt.Sprintf("-Torus-%d", i),
					}).Stage(stager.stage)
					stackShape.DiscreteTorusShapes[i] = t
				}
				modified = true
			}

			for i, t := range stackShape.DiscreteTorusShapes {
				centerY := float64(i) * plant.RelativeCuttedStackFloorHeight * plant.RhombusSideLength
				if t.CenterY != centerY || t.Radius != R || t.TubeRadius != tubeRadius || t.Color != "blue" {
					t.CenterY = centerY
					t.Radius = R
					t.TubeRadius = tubeRadius
					t.Color = "blue"
					modified = true
				}
			}
		}
	}

	return modified
}
