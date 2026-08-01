package models

import (
	"fmt"
	"math"
	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func (stager *Stager) resampleCurvesByAngle(
	originalBottom *threejs.Curve,
	originalTop *threejs.Curve,
	degreeInterval float64,
	namePrefix string,
) (*threejs.Curve, *threejs.Curve) {
	resampledBottom := (&threejs.Curve{
		Name: fmt.Sprintf("%s Resampled Bottom", namePrefix),
	}).Stage(stager.threejsStage)
	resampledTop := (&threejs.Curve{
		Name: fmt.Sprintf("%s Resampled Top", namePrefix),
	}).Stage(stager.threejsStage)

	if len(originalBottom.Points) == 0 || len(originalTop.Points) == 0 || len(originalBottom.Points) != len(originalTop.Points) {
		return resampledBottom, resampledTop
	}

	unwrapAngles := func(curve *threejs.Curve) []float64 {
		angles := make([]float64, len(curve.Points))
		firstP := curve.Points[0]
		lastTheta := math.Atan2(firstP.Z, firstP.X)
		accumulated := lastTheta
		for accumulated < 0 {
			accumulated += 2 * math.Pi
		}
		for accumulated >= 2*math.Pi {
			accumulated -= 2 * math.Pi
		}
		angles[0] = accumulated
		lastTheta = math.Atan2(firstP.Z, firstP.X)

		for i := 1; i < len(curve.Points); i++ {
			p := curve.Points[i]
			theta := math.Atan2(p.Z, p.X)
			diff := theta - lastTheta
			for diff < -math.Pi {
				diff += 2 * math.Pi
			}
			for diff > math.Pi {
				diff -= 2 * math.Pi
			}
			accumulated += diff
			angles[i] = accumulated
			lastTheta = theta
		}
		return angles
	}

	anglesBottom := unwrapAngles(originalBottom)
	anglesTop := unwrapAngles(originalTop)

	// Align phases! If they start at different multiples of 2pi (e.g. 0.0 vs 6.28)
	// target - a1T will be massively off, causing tT to be clamped to 0 or 1.
	for (anglesTop[0] - anglesBottom[0]) > math.Pi {
		for i := range anglesTop {
			anglesTop[i] -= 2 * math.Pi
		}
	}
	for (anglesTop[0] - anglesBottom[0]) < -math.Pi {
		for i := range anglesTop {
			anglesTop[i] += 2 * math.Pi
		}
	}

	radInterval := degreeInterval * math.Pi / 180.0
	
	lastTarget := -math.MaxFloat64
	lastIsForward := true

	for i := 0; i < len(originalBottom.Points)-1; i++ {
		a1B := anglesBottom[i]
		a2B := anglesBottom[i+1]

		isForward := a1B <= a2B

		if i > 0 && isForward != lastIsForward {
			lastTarget = -math.MaxFloat64
		}
		lastIsForward = isForward

		var targetAngles []float64
		if isForward {
			startTarget := math.Ceil((a1B - 1e-7) / radInterval) * radInterval
			for target := startTarget; target <= a2B + 1e-7; target += radInterval {
				targetAngles = append(targetAngles, target)
			}
		} else {
			startTarget := math.Floor((a1B + 1e-7) / radInterval) * radInterval
			for target := startTarget; target >= a2B - 1e-7; target -= radInterval {
				targetAngles = append(targetAngles, target)
			}
		}

		p1B := originalBottom.Points[i]
		p2B := originalBottom.Points[i+1]
		p1T := originalTop.Points[i]
		p2T := originalTop.Points[i+1]

		for _, target := range targetAngles {
			if math.Abs(target - lastTarget) < 1e-6 {
				continue
			}
			lastTarget = target

			t := 0.0
			if a1B != a2B {
				t = (target - a1B) / (a2B - a1B)
			}

			ptBottom := (&threejs.Vector3{
				Name: fmt.Sprintf("%s Bottom %.1f", namePrefix, target*180.0/math.Pi),
				X:    p1B.X + t*(p2B.X - p1B.X),
				Y:    p1B.Y + t*(p2B.Y - p1B.Y),
				Z:    p1B.Z + t*(p2B.Z - p1B.Z),
			}).Stage(stager.threejsStage)

			ptTop := (&threejs.Vector3{
				Name: fmt.Sprintf("%s Top %.1f", namePrefix, target*180.0/math.Pi),
				X:    p1T.X + t*(p2T.X - p1T.X),
				Y:    p1T.Y + t*(p2T.Y - p1T.Y),
				Z:    p1T.Z + t*(p2T.Z - p1T.Z),
			}).Stage(stager.threejsStage)

			resampledBottom.Points = append(resampledBottom.Points, ptBottom)
			resampledTop.Points = append(resampledTop.Points, ptTop)
		}
	}

	return resampledBottom, resampledTop
}

func (stager *Stager) addPointSpheres(points []*threejs.Vector3, color string, canvas *threejs.Canvas, namePrefix string) {
	for i, pt := range points {
		sphereColor := color
		radius := 2.0
		if i % 40 == 0 {
			sphereColor = "yellow"
			radius = 4.0
		}

		sphere := (&threejs.Mesh{
			Name: fmt.Sprintf("%s Sphere %d", namePrefix, i),
			Position: threejs.Position{
				X: pt.X,
				Y: pt.Y,
				Z: pt.Z,
			},
			SphereGeometry: (&threejs.SphereGeometry{
				Name:   fmt.Sprintf("%s SphereGeom %d", namePrefix, i),
				Radius: radius,
			}).Stage(stager.threejsStage),
			MeshMaterialBasic: (&threejs.MeshMaterialBasic{
				Name: fmt.Sprintf("%s SphereMat %d", namePrefix, i),
				MeshMaterialAbstract: threejs.MeshMaterialAbstract{
					Color: sphereColor,
				},
			}).Stage(stager.threejsStage),
		}).Stage(stager.threejsStage)
		canvas.Meshs = append(canvas.Meshs, sphere)
	}
}
