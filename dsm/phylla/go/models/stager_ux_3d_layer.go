package models

import (
	"fmt"
	"math"

	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func (stager *Stager) generateLayerWithModulo(
	h int, dx, dy, thetaOffset float64, namePrefix string,
	plant *Plant, checkedDiagram *PlantDiagram,
	resampledBaseBottom *threejs.Curve, resampledBaseTop *threejs.Curve,
	thickness float64, globalR float64,
	canvas *threejs.Canvas,
) {
	radialRepetition := plant.RadialRepetitions

	massiveBottomCurve := (&threejs.Curve{Name: fmt.Sprintf("%s Massive Bottom h%d", namePrefix, h)}).Stage(stager.threejsStage)
	massiveTopCurve := (&threejs.Curve{Name: fmt.Sprintf("%s Massive Top h%d", namePrefix, h)}).Stage(stager.threejsStage)

	for k := 0; k < radialRepetition; k++ {
		baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(radialRepetition)
		totalThetaOffset := thetaOffset + baseThetaOffset

		localBottomCurve := stager.cloneAndRotateCurve(resampledBaseBottom, totalThetaOffset)
		localTopCurve := stager.cloneAndRotateCurve(resampledBaseTop, totalThetaOffset)

		for i := 0; i < len(localBottomCurve.Points); i++ {
			massiveBottomCurve.Points = append(massiveBottomCurve.Points, localBottomCurve.Points[i])
		}
		for i := 0; i < len(localTopCurve.Points); i++ {
			massiveTopCurve.Points = append(massiveTopCurve.Points, localTopCurve.Points[i])
		}
	}

	if !checkedDiagram.IsHiddenSampledPoints3DShape {
		numPointsPerRep := len(resampledBaseBottom.Points)
		stager.addPointSpheres(massiveBottomCurve.Points, "red", canvas, namePrefix+" Bottom", dy, numPointsPerRep)
		stager.addPointSpheres(massiveTopCurve.Points, "blue", canvas, namePrefix+" Top", dy, numPointsPerRep)
	}

	stager.generateRibbonMesh(h, thetaOffset, namePrefix, plant, checkedDiagram, massiveBottomCurve, massiveTopCurve, dy, thickness, globalR, canvas)
}
