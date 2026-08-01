package models

import (
	"fmt"
	"log"
	"math"
	"sort"

	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

// unwrapAngles processes a 3D curve to extract a strictly monotonic,
// duplicate-free mapping of points to their cylindrical angle (theta).
//
// Algorithm:
//  1. Iterates through the points in the curve.
//  2. Computes the horizontal angle `theta = math.Atan2(Z, X)` for each point.
//  3. Normalizes all angles strictly to the range [0, 2π).
//  4. Stores points in a map keyed by their normalized angle.
//     This intrinsically removes any overlapping segments (e.g. if the curve wraps past 2π,
//     those points merge with the existing points in the 0-2π space).
//  5. Extracts the unique keys and sorts them.
//  6. Returns a perfectly sorted, strictly increasing array of angles along with
//     their corresponding 3D points.
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
			log.Fatalf("overlapping segment detected: curve goes backwards at index %d (diff: %f)", i, diff)
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

// getTargetAngles analyzes the Top and Bottom curves to establish a unified set of target
// angles for resampling. It guarantees that both the Top and Bottom ribbons will be generated
// with the exact same number of horizontal points, perfectly vertically aligned.
func (stager *Stager) getTargetAngles(
	originalBottom *threejs.Curve,
	originalTop *threejs.Curve,
	degreeInterval float64,
) (targetAngles []float64, bottomAngles []float64, bottomPoints []*threejs.Vector3, topAngles []float64, topPoints []*threejs.Vector3) {
	if len(originalBottom.Points) == 0 || len(originalTop.Points) == 0 {
		return nil, nil, nil, nil, nil
	}

	bottomAngles, bottomPoints = unwrapAngles(originalBottom)
	topAngles, topPoints = unwrapAngles(originalTop)

	// Determine the absolute minimum and maximum angles present across both curves
	minAngle := math.Min(bottomAngles[0], topAngles[0])
	maxAngle := math.Max(bottomAngles[len(bottomAngles)-1], topAngles[len(topAngles)-1])

	radInterval := degreeInterval * math.Pi / 180.0

	// Generate a strictly monotonic sequence of target angles spaced by degreeInterval
	startTarget := math.Ceil((minAngle-1e-7)/radInterval) * radInterval
	for target := startTarget; target <= maxAngle+1e-7; target += radInterval {
		targetAngles = append(targetAngles, target)
	}

	return targetAngles, bottomAngles, bottomPoints, topAngles, topPoints
}

// resampleCurveAtAngles forces an existing 3D curve to conform to a specific set of target angles.
// It iterates through the target angles, finds the closest bounding points from the original sorted curve,
// and linearly interpolates the height (Y) and radius (R) to generate a new smoothly spaced 3D vertex.
//
// If a target angle falls strictly outside the boundaries of the original curve (extrapolation),
// it safely clamps to the exact 3D coordinate of the nearest endpoint, allowing the mesh
// to taper seamlessly without generating erratic circular sweeps.
func (stager *Stager) resampleCurveAtAngles(
	sortedAngles []float64,
	sortedPoints []*threejs.Vector3,
	targetAngles []float64,
	namePrefix string,
) *threejs.Curve {
	resampled := (&threejs.Curve{
		Name: fmt.Sprintf("%s Resampled", namePrefix),
	}).Stage(stager.threejsStage)

	if len(sortedAngles) == 0 {
		return resampled
	}

	for _, target := range targetAngles {
		idx := -1
		
		// Use binary search (O(log n)) to find the first index where sortedAngles[i] >= target
		searchIdx := sort.SearchFloat64s(sortedAngles, target)
		
		if searchIdx > 0 && searchIdx < len(sortedAngles) {
			// Target is enclosed perfectly between searchIdx-1 and searchIdx
			idx = searchIdx - 1
		} else if searchIdx == 0 && sortedAngles[0] == target {
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
			t = (target - a1) / (a2 - a1)
		}
		if t < 0 {
			t = 0
		}
		if t > 1 {
			t = 1
		}

		var x, y, z float64
		if isExtrapolating {
			if t == 0 {
				x, y, z = p1.X, p1.Y, p1.Z
			} else {
				x, y, z = p2.X, p2.Y, p2.Z
			}
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
		}).Stage(stager.threejsStage)

		resampled.Points = append(resampled.Points, pt)
	}

	return resampled
}
