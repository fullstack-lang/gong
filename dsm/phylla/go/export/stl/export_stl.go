package stl

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// Vector3 represents a 3D coordinate for mesh computation and STL export.
type Vector3 struct {
	X, Y, Z float64
}

func (v Vector3) Sub(other Vector3) Vector3 {
	return Vector3{v.X - other.X, v.Y - other.Y, v.Z - other.Z}
}

func (v Vector3) Cross(other Vector3) Vector3 {
	return Vector3{
		X: v.Y*other.Z - v.Z*other.Y,
		Y: v.Z*other.X - v.X*other.Z,
		Z: v.X*other.Y - v.Y*other.X,
	}
}

func (v Vector3) Normalize() Vector3 {
	length := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	if length == 0 {
		return Vector3{0, 0, 0}
	}
	return Vector3{v.X / length, v.Y / length, v.Z / length}
}

// WriteFacet writes a single triangular facet with calculated normal into the string builder.
func WriteFacet(sb *strings.Builder, v1, v2, v3 Vector3) {
	normal := v2.Sub(v1).Cross(v3.Sub(v1)).Normalize()
	sb.WriteString(fmt.Sprintf("  facet normal %e %e %e\n", normal.X, normal.Y, normal.Z))
	sb.WriteString("    outer loop\n")
	sb.WriteString(fmt.Sprintf("      vertex %e %e %e\n", v1.X, v1.Y, v1.Z))
	sb.WriteString(fmt.Sprintf("      vertex %e %e %e\n", v2.X, v2.Y, v2.Z))
	sb.WriteString(fmt.Sprintf("      vertex %e %e %e\n", v3.X, v3.Y, v3.Z))
	sb.WriteString("    endloop\n")
	sb.WriteString("  endfacet\n")
}

// UnwrapAngles unwraps points ordered by azimuth angle theta.
func UnwrapAngles(pts []Vector3) (angles []float64, points []Vector3) {
	if len(pts) == 0 {
		return nil, nil
	}
	angleToPoint := make(map[float64]Vector3)

	firstP := pts[0]
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

	for i := 1; i < len(pts); i++ {
		p := pts[i]
		theta := math.Atan2(p.Z, p.X)
		diff := theta - lastTheta

		for diff < -math.Pi {
			diff += 2 * math.Pi
		}
		for diff > math.Pi {
			diff -= 2 * math.Pi
		}

		if diff < -1e-7 {
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
