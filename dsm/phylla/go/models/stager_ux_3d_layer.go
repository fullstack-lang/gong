package models

import (
	"fmt"
	"math"

	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func (stager *Stager) generateLayerWithModulo(
	h int, dx, dy, thetaOffset float64, namePrefix string,
	plant *Plant, checkedDiagram *PlantDiagram,
	curve, topCurve *threejs.Curve,
	thickness, globalR float64, canvas *threejs.Canvas,
) {
	radialRepetition := plant.RadialRepetitions

	massiveBottomCurve := (&threejs.Curve{Name: fmt.Sprintf("%s Massive Bottom h%d", namePrefix, h)}).Stage(stager.threejsStage)
	massiveTopCurve := (&threejs.Curve{Name: fmt.Sprintf("%s Massive Top h%d", namePrefix, h)}).Stage(stager.threejsStage)

	for k := 0; k < radialRepetition; k++ {
		baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(radialRepetition)
		totalThetaOffset := thetaOffset + baseThetaOffset

		localBottomCurve := stager.cloneAndRotateCurve(curve, totalThetaOffset)
		localTopCurve := stager.cloneAndRotateCurve(topCurve, totalThetaOffset)

		for i := 0; i < len(localBottomCurve.Points); i++ {
			massiveBottomCurve.Points = append(massiveBottomCurve.Points, localBottomCurve.Points[i])
		}
		for i := 0; i < len(localTopCurve.Points); i++ {
			massiveTopCurve.Points = append(massiveTopCurve.Points, localTopCurve.Points[i])
		}
	}

	stager.generateRibbonMesh(h, thetaOffset, namePrefix, plant, checkedDiagram, massiveBottomCurve, massiveTopCurve, dy, thickness, globalR, canvas)
}
