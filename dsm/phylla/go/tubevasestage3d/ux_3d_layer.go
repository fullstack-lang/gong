package tubevasestage3d

import (
	"fmt"
	"math"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func (u *ThreeJSStageUpdater) computeMassiveCurves(
	stager *models.Stager,
	thetaOffset float64, namePrefix string,
	plant *models.PlantAbstract,
	resampledBaseBottom *threejs.Curve, resampledBaseTop *threejs.Curve,
) (*threejs.Curve, *threejs.Curve) {
	threejsStage := stager.GetThreejsStage()

	radialRepetition := 1
	if (plant.PlantType == models.TubeVase || plant.PlantType == models.VaseTrapeze) && plant.TubeVaseAbstract != nil {
		radialRepetition = plant.TubeVaseAbstract.RadialRepetitions
	}

	massiveBottomCurve := (&threejs.Curve{Name: fmt.Sprintf("%s Massive Bottom", namePrefix)}).Stage(threejsStage)
	massiveTopCurve := (&threejs.Curve{Name: fmt.Sprintf("%s Massive Top", namePrefix)}).Stage(threejsStage)

	for k := 0; k < radialRepetition; k++ {
		baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(radialRepetition)
		totalThetaOffset := thetaOffset + baseThetaOffset

		localBottomCurve := u.cloneAndRotateCurve(stager, resampledBaseBottom, totalThetaOffset)
		localTopCurve := u.cloneAndRotateCurve(stager, resampledBaseTop, totalThetaOffset)

		for i := 0; i < len(localBottomCurve.Points); i++ {
			massiveBottomCurve.Points = append(massiveBottomCurve.Points, localBottomCurve.Points[i])
		}
		for i := 0; i < len(localTopCurve.Points); i++ {
			massiveTopCurve.Points = append(massiveTopCurve.Points, localTopCurve.Points[i])
		}
	}

	return massiveBottomCurve, massiveTopCurve
}

func (u *ThreeJSStageUpdater) generateLayerWithModulo(
	stager *models.Stager,
	h int, stackHeight int, dx, dy, thetaOffset float64, namePrefix string,
	plant *models.PlantAbstract, checkedDiagram *models.TubeVase3DDiagram,
	resampledBaseBottom *threejs.Curve, resampledBaseTop *threejs.Curve,
	thickness float64, globalR float64,
	canvas *threejs.Canvas,
) (*threejs.Curve, *threejs.Curve) {
	threejsStage := stager.GetThreejsStage()

	h_horiz := 0.0
	if (plant.PlantType == models.TubeVase || plant.PlantType == models.VaseTrapeze) && plant.TubeVaseAbstract != nil {
		h_horiz = plant.TubeVaseAbstract.RelativeHorizontalRingsHeight * plant.RhombusSideLength
	}

	massiveBottomCurve, massiveTopCurve := u.computeMassiveCurves(stager, thetaOffset, fmt.Sprintf("%s h%d", namePrefix, h), plant, resampledBaseBottom, resampledBaseTop)

	if !checkedDiagram.IsHiddenSampledPoints3DShape {
		numPointsPerRep := len(resampledBaseBottom.Points)
		u.addPointSpheres(stager, massiveBottomCurve.Points, "red", canvas, namePrefix+" Bottom", dy, numPointsPerRep)
		u.addPointSpheres(stager, massiveTopCurve.Points, "blue", canvas, namePrefix+" Top", dy, numPointsPerRep)
	}

	u.generateRibbonMesh(stager, h, stackHeight, thetaOffset, namePrefix, plant, checkedDiagram, massiveBottomCurve, massiveTopCurve, dy, thickness, globalR, canvas)

	if h_horiz > 0 && h == 0 && len(massiveBottomCurve.Points) > 0 {
		minY_bottom := math.MaxFloat64
		for _, p := range massiveBottomCurve.Points {
			yVal := p.Y + dy
			if yVal < minY_bottom {
				minY_bottom = yVal
			}
		}

		horizBottomCurve := (&threejs.Curve{Name: fmt.Sprintf("%s Horiz Bottom h%d", namePrefix, h)}).Stage(threejsStage)
		horizTopCurve := (&threejs.Curve{Name: fmt.Sprintf("%s Horiz Top h%d", namePrefix, h)}).Stage(threejsStage)
		for _, p := range massiveBottomCurve.Points {
			horizBottomCurve.Points = append(horizBottomCurve.Points, (&threejs.Vector3{
				X: p.X,
				Y: minY_bottom - dy,
				Z: p.Z,
			}).Stage(threejsStage))
			horizTopCurve.Points = append(horizTopCurve.Points, (&threejs.Vector3{
				X: p.X,
				Y: (minY_bottom + h_horiz) - dy,
				Z: p.Z,
			}).Stage(threejsStage))
		}
		u.generateRibbonMesh(stager, h, stackHeight, thetaOffset, namePrefix+" Horiz Bottom", plant, checkedDiagram, horizBottomCurve, horizTopCurve, dy, thickness, globalR, canvas)
	}

	if h_horiz > 0 && h == stackHeight-1 && len(massiveTopCurve.Points) > 0 {
		maxY_top := -math.MaxFloat64
		for _, p := range massiveTopCurve.Points {
			yVal := p.Y + dy
			if yVal > maxY_top {
				maxY_top = yVal
			}
		}

		horizBottomCurve := (&threejs.Curve{Name: fmt.Sprintf("%s Horiz Top Bottom h%d", namePrefix, h)}).Stage(threejsStage)
		horizTopCurve := (&threejs.Curve{Name: fmt.Sprintf("%s Horiz Top Top h%d", namePrefix, h)}).Stage(threejsStage)
		for _, p := range massiveTopCurve.Points {
			horizBottomCurve.Points = append(horizBottomCurve.Points, (&threejs.Vector3{
				X: p.X,
				Y: (maxY_top - h_horiz) - dy,
				Z: p.Z,
			}).Stage(threejsStage))
			horizTopCurve.Points = append(horizTopCurve.Points, (&threejs.Vector3{
				X: p.X,
				Y: maxY_top - dy,
				Z: p.Z,
			}).Stage(threejsStage))
		}
		u.generateRibbonMesh(stager, h, stackHeight, thetaOffset, namePrefix+" Horiz Top", plant, checkedDiagram, horizBottomCurve, horizTopCurve, dy, thickness, globalR, canvas)
	}

	return massiveBottomCurve, massiveTopCurve
}
