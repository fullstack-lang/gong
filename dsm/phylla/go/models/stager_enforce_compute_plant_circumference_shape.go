package models

import (
	"math"
)

// enforceComputePlantCircumferenceShape computes the AngleDegree and Length of the PlantCircumferenceShape
// for each PlantDiagram based on its owning Plant's properties (N, M, InsideAngle, SideLength).
func (stager *Stager) enforceComputePlantCircumferenceShape() (needCommit bool) {
	stage := stager.stage

	for plant := range *GetGongstructInstancesSetFromPointerType[*Plant](stage) {
		// The lattice of the plant surface is formed by rhombuses.
		// There are two main grid directions (parastichies):
		// 1. Up and to the right
		// 2. Up and to the left

		// The angle of these paths relative to the horizontal axis is half the InsideAngle.
		insideAngleRad := plant.RhombusInsideAngle * math.Pi / 180.0
		sinHalfInsideAngle := math.Sin(insideAngleRad / 2.0)
		cosHalfInsideAngle := math.Cos(insideAngleRad / 2.0)
		sideLength := plant.RhombusSideLength

		// Y movement: moving N steps along the up-right path and M steps along the up-left path.
		// Both paths go up, so we add the Y components together.
		y := float64(plant.N)*sideLength*sinHalfInsideAngle +
			float64(plant.M)*sideLength*sinHalfInsideAngle

		// X movement: moving N steps along the up-right path (positive X)
		// and M steps along the up-left path (negative X).
		// So we subtract the M component from the N component.
		x := float64(plant.N)*sideLength*cosHalfInsideAngle -
			float64(plant.M)*sideLength*cosHalfInsideAngle

		// Convert the final Cartesian coordinate into a polar vector (Angle and Length)
		computedAngleDegree := math.Atan2(y, x) * 180.0 / math.Pi
		computedLength := math.Sqrt(x*x + y*y)

		if plant.RhombusStuff.PlantCircumferenceShape != nil {
			if plant.RhombusStuff.PlantCircumferenceShape.AngleDegree != computedAngleDegree ||
				plant.RhombusStuff.PlantCircumferenceShape.Length != computedLength {

				plant.RhombusStuff.PlantCircumferenceShape.AngleDegree = computedAngleDegree
				plant.RhombusStuff.PlantCircumferenceShape.Length = computedLength

				needCommit = true
			}
		}
		if plant.RhombusStuff.RotatedPlantCircumferenceShape != nil {
			if plant.RhombusStuff.RotatedPlantCircumferenceShape.AngleDegree != 0 ||
				plant.RhombusStuff.RotatedPlantCircumferenceShape.Length != computedLength {

				plant.RhombusStuff.RotatedPlantCircumferenceShape.AngleDegree = 0
				plant.RhombusStuff.RotatedPlantCircumferenceShape.Length = computedLength

				needCommit = true
			}
		}
	}

	return
}
