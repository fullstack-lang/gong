package models

import (
	"fmt"
	"math"

	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func (stager *Stager) generateLayerWithModulo(
	h int, stackHeight int, dx, dy, thetaOffset float64, namePrefix string,
	plant *PlantAbstract, checkedDiagram *PlantDiagram,
	resampledBaseBottom *threejs.Curve, resampledBaseTop *threejs.Curve,
	thickness float64, globalR float64,
	canvas *threejs.Canvas,
) {
	radialRepetition := 1
	h_horiz := 0.0
	if plant.VaseAbstract != nil {
		radialRepetition = plant.VaseAbstract.RadialRepetitions
		h_horiz = plant.VaseAbstract.RelativeHorizontalRingsHeight * plant.VaseAbstract.RhombusSideLength
	}

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

	stager.generateRibbonMesh(h, stackHeight, thetaOffset, namePrefix, plant, checkedDiagram, massiveBottomCurve, massiveTopCurve, dy, thickness, globalR, canvas)

	if h_horiz > 0 && h == 0 && len(massiveBottomCurve.Points) > 0 {
		minY_bottom := math.MaxFloat64
		for _, p := range massiveBottomCurve.Points {
			yVal := p.Y + dy
			if yVal < minY_bottom {
				minY_bottom = yVal
			}
		}

		horizBottomCurve := (&threejs.Curve{Name: fmt.Sprintf("%s Horiz Bottom h%d", namePrefix, h)}).Stage(stager.threejsStage)
		horizTopCurve := (&threejs.Curve{Name: fmt.Sprintf("%s Horiz Top h%d", namePrefix, h)}).Stage(stager.threejsStage)
		for _, p := range massiveBottomCurve.Points {
			horizBottomCurve.Points = append(horizBottomCurve.Points, (&threejs.Vector3{
				X: p.X,
				Y: minY_bottom - dy,
				Z: p.Z,
			}).Stage(stager.threejsStage))
			horizTopCurve.Points = append(horizTopCurve.Points, (&threejs.Vector3{
				X: p.X,
				Y: (minY_bottom + h_horiz) - dy,
				Z: p.Z,
			}).Stage(stager.threejsStage))
		}
		stager.generateRibbonMesh(h, stackHeight, thetaOffset, namePrefix+" Horiz Bottom", plant, checkedDiagram, horizBottomCurve, horizTopCurve, dy, thickness, globalR, canvas)
	}

	if h_horiz > 0 && h == stackHeight-1 && len(massiveTopCurve.Points) > 0 {
		maxY_top := -math.MaxFloat64
		for _, p := range massiveTopCurve.Points {
			yVal := p.Y + dy
			if yVal > maxY_top {
				maxY_top = yVal
			}
		}

		horizBottomCurve := (&threejs.Curve{Name: fmt.Sprintf("%s Horiz Top Bottom h%d", namePrefix, h)}).Stage(stager.threejsStage)
		horizTopCurve := (&threejs.Curve{Name: fmt.Sprintf("%s Horiz Top Top h%d", namePrefix, h)}).Stage(stager.threejsStage)
		for _, p := range massiveTopCurve.Points {
			horizBottomCurve.Points = append(horizBottomCurve.Points, (&threejs.Vector3{
				X: p.X,
				Y: (maxY_top - h_horiz) - dy,
				Z: p.Z,
			}).Stage(stager.threejsStage))
			horizTopCurve.Points = append(horizTopCurve.Points, (&threejs.Vector3{
				X: p.X,
				Y: maxY_top - dy,
				Z: p.Z,
			}).Stage(stager.threejsStage))
		}
		stager.generateRibbonMesh(h, stackHeight, thetaOffset, namePrefix+" Horiz Top", plant, checkedDiagram, horizBottomCurve, horizTopCurve, dy, thickness, globalR, canvas)
	}
}
