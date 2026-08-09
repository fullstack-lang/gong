package stage3d

import (
	"fmt"
	"log"
	"math"
	"sort"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

// unwrapAngles processes a 3D curve to extract a strictly monotonic,
// duplicate-free mapping of points to their cylindrical angle (theta).
func unwrapAngles(curve *threejs.Curve) (angles []float64, points []*threejs.Vector3) {
	angleToPoint := make(map[float64]*threejs.Vector3)
	if len(curve.Points) == 0 {
		return nil, nil
	}

	firstP := curve.Points[0]
	lastTheta := math.Atan2(firstP.Z, firstP.X)

	accumulated := lastTheta
	for accumulated < 0 {
		accumulated += 2 * math.Pi
	}
	for accumulated >= 2*math.Pi {
		accumulated -= 2 * math.Pi
	}

	angleToPoint[accumulated] = firstP
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

		if diff < -1e-7 {
			log.Printf("overlapping segment detected: curve goes backwards at index %d (diff: %f)", i, diff)
			diff = 0
		}

		accumulated += diff

		angleToPoint[accumulated] = p
		lastTheta = theta
	}

	for a := range angleToPoint {
		angles = append(angles, a)
	}
	sort.Float64s(angles)

	for _, a := range angles {
		points = append(points, angleToPoint[a])
	}

	return angles, points
}

func (u *ThreeJSStageUpdater) getTargetAngles(
	originalBottom *threejs.Curve,
	originalTop *threejs.Curve,
	degreeInterval float64,
	radialRepetitions int,
) (targetAngles []float64, bottomAngles []float64, bottomPoints []*threejs.Vector3, topAngles []float64, topPoints []*threejs.Vector3) {
	if len(originalBottom.Points) == 0 || len(originalTop.Points) == 0 {
		return nil, nil, nil, nil, nil
	}

	bottomAngles, bottomPoints = unwrapAngles(originalBottom)
	topAngles, topPoints = unwrapAngles(originalTop)

	radInterval := degreeInterval * math.Pi / 180.0

	// We expect the curve to cover exactly 1/RadialRepetitions of the full circle.
	if radialRepetitions < 1 {
		radialRepetitions = 1
	}
	expectedDegrees := 360.0 / float64(radialRepetitions)

	// Generate a strictly monotonic sequence of target angles spaced by degreeInterval
	nbPoints := int(math.Round(expectedDegrees / degreeInterval))
	for i := 0; i <= nbPoints; i++ {
		targetAngles = append(targetAngles, float64(i)*radInterval)
	}

	return targetAngles, bottomAngles, bottomPoints, topAngles, topPoints
}

// resampleCurveAtAngles forces an existing 3D curve to conform to a specific set of target angles.
func (u *ThreeJSStageUpdater) resampleCurveAtAngles(
	stager *models.Stager,
	sortedAngles []float64,
	sortedPoints []*threejs.Vector3,
	targetAngles []float64,
	namePrefix string,
	expectedDegrees float64,
) *threejs.Curve {
	threejsStage := stager.GetThreejsStage()

	resampled := (&threejs.Curve{
		Name: fmt.Sprintf("%s Resampled", namePrefix),
	}).Stage(threejsStage)

	if len(sortedAngles) == 0 {
		return resampled
	}

	for _, target := range targetAngles {
		evalTarget := target
		if len(sortedAngles) > 0 && expectedDegrees > 0 {
			minA := sortedAngles[0]
			maxA := sortedAngles[len(sortedAngles)-1]
			expectedRad := expectedDegrees * math.Pi / 180.0

			for evalTarget < minA {
				evalTarget += expectedRad
			}
			for evalTarget > maxA {
				evalTarget -= expectedRad
			}
			if evalTarget < minA {
				evalTarget = minA
			}
			if evalTarget > maxA {
				evalTarget = maxA
			}
		}

		idx := -1

		// Use binary search (O(log n)) to find the first index where sortedAngles[i] >= evalTarget
		searchIdx := sort.SearchFloat64s(sortedAngles, evalTarget)

		if searchIdx > 0 && searchIdx < len(sortedAngles) {
			// Target is enclosed perfectly between searchIdx-1 and searchIdx
			idx = searchIdx - 1
		} else if searchIdx == 0 && sortedAngles[0] == evalTarget {
			// Target perfectly matches the very first angle
			idx = 0
		}

		isExtrapolating := false
		if idx == -1 {
			isExtrapolating = true

			// Extrapolating outside bounds: snap to the absolute closest end
			if searchIdx == 0 {
				idx = 0 // Closer to the start
			} else {
				idx = len(sortedAngles) - 2 // Closer to the end
			}
		}

		a1 := sortedAngles[idx]
		a2 := sortedAngles[idx+1]
		p1 := sortedPoints[idx]
		p2 := sortedPoints[idx+1]

		t := 0.0
		if a1 != a2 {
			t = (evalTarget - a1) / (a2 - a1)
		}
		if t < 0 {
			t = 0
		}
		if t > 1 {
			t = 1
		}

		var x, y, z float64
		if isExtrapolating {
			var r float64
			if t == 0 {
				r = math.Hypot(p1.X, p1.Z)
				y = p1.Y
			} else {
				r = math.Hypot(p2.X, p2.Z)
				y = p2.Y
			}
			x = r * math.Cos(target)
			z = r * math.Sin(target)
		} else {
			r1 := math.Hypot(p1.X, p1.Z)
			r2 := math.Hypot(p2.X, p2.Z)
			r := r1 + t*(r2-r1)

			// interpolate the height between the two points related to those origin angles
			y = p1.Y + t*(p2.Y-p1.Y)

			x = r * math.Cos(target)
			z = r * math.Sin(target)
		}

		pt := (&threejs.Vector3{
			Name: fmt.Sprintf("%s %.1f", namePrefix, target*180.0/math.Pi),
			X:    x,
			Y:    y,
			Z:    z,
		}).Stage(threejsStage)

		resampled.Points = append(resampled.Points, pt)
	}

	return resampled
}
