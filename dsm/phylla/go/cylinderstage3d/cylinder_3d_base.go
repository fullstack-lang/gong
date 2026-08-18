package cylinderstage3d

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

type Cylinder3DParams struct {
	NamePrefix            string
	CanvasName            string
	RadialRepetitions     int
	Transparency          float64
	RelativeTubeDiameter  float64
	RelativeHeight3DTorus float64
	VerticalScale         float64
	RelativeHeight        float64
	ProjectionAngle       float64

	// Visibility flags
	IsHiddenTorus3DShape                bool
	IsHiddenRotatedTorusShape           bool
	IsHiddenTopCurveShape               bool
	IsHiddenRotatedTopCurveShape        bool
	IsHiddenSampledPoints3DShape        bool
	IsHiddenRotatedSampledPoints3DShape bool

	// Features toggle
	HasRotatedShapes bool

	// Camera persistence
	Rendered3DShape *models.Rendered3DShape
}

type Cylinder3DBaseResult struct {
	Canvas             *threejs.Canvas
	ResampledBaseCurve *threejs.Curve
	GlobalR            float64
	RadialRepetitions  int
	TubeRadius         float64
	Opacity            float64
	TopHeight          float64
	TorusHeight        float64
	VertScale          float64
	GrowthVectorX      float64
	GrowthVectorY      float64
	ProjAngleRad       float64
	FloorMinY          float64
	ExpectedDegrees    float64
	TargetAngles       []float64
	SortedAngles       []float64
	SortedPoints       []*threejs.Vector3
	RotTopPoints       []*threejs.Vector3
	RotTopCurve        *threejs.Curve
}

func RenderCylinder3DBase(
	stage3d *threejs.Stage,
	stager *models.Stager,
	plant *models.PlantAbstract,
	params Cylinder3DParams,
) *Cylinder3DBaseResult {
	if stage3d == nil {
		return nil
	}

	stage3d.Reset()

	canvas := (&threejs.Canvas{
		Name: params.CanvasName,
	}).Stage(stage3d)

	// Circumference and cylinder radius
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

	radialRepetitions := params.RadialRepetitions
	if radialRepetitions < 1 {
		radialRepetitions = 1
	}

	globalR := circumference * float64(radialRepetitions) / (2 * math.Pi)

	// Directional and ambient lights positioned relative to scene scale
	lightScale := math.Max(globalR, 50.0)
	dirLight1 := (&threejs.DirectionalLight{
		Name:             "Directional Light 1 (Key)",
		Position:         threejs.Position{X: lightScale * 2.0, Y: lightScale * 2.5, Z: lightScale * 2.0},
		LightAbstract:    threejs.LightAbstract{Intensity: 1.2},
		IsWithCastShadow: true,
	}).Stage(stage3d)

	dirLight2 := (&threejs.DirectionalLight{
		Name:             "Directional Light 2 (Fill)",
		Position:         threejs.Position{X: -lightScale * 2.0, Y: lightScale * 1.5, Z: -lightScale * 2.0},
		LightAbstract:    threejs.LightAbstract{Intensity: 0.6},
		IsWithCastShadow: false,
	}).Stage(stage3d)

	dirLight3 := (&threejs.DirectionalLight{
		Name:             "Directional Light 3 (Rim)",
		Position:         threejs.Position{X: 0, Y: lightScale * 3.5, Z: -lightScale * 2.5},
		LightAbstract:    threejs.LightAbstract{Intensity: 0.8},
		IsWithCastShadow: false,
	}).Stage(stage3d)

	canvas.DirectionalLights = append(canvas.DirectionalLights, dirLight1, dirLight2, dirLight3)

	ambiantLight := (&threejs.AmbiantLight{
		Name:          "Ambiant Light",
		LightAbstract: threejs.LightAbstract{Intensity: 0.4},
	}).Stage(stage3d)
	canvas.AmbiantLight = ambiantLight

	// Camera setup directly from Rendered3DShape
	rendered3DShape := params.Rendered3DShape
	if rendered3DShape != nil && (rendered3DShape.ViewX != 0 || rendered3DShape.ViewY != 0 || rendered3DShape.ViewZ != 0) {
		fov := rendered3DShape.Fov
		if fov == 0 {
			fov = 50
		}
		canvas.Camera = (&threejs.Camera{
			Name: "Camera",
			Position: threejs.Position{
				X: rendered3DShape.ViewX,
				Y: rendered3DShape.ViewY,
				Z: rendered3DShape.ViewZ,
			},
			TargetX: rendered3DShape.TargetX,
			TargetY: rendered3DShape.TargetY,
			TargetZ: rendered3DShape.TargetZ,
			Fov:     fov,
		}).Stage(stage3d)
	} else {
		camDist := globalR * 2.5
		if camDist < 30 {
			camDist = 30
		}
		canvas.Camera = (&threejs.Camera{
			Name: "Camera",
			Position: threejs.Position{
				X: camDist,
				Y: camDist * 0.8,
				Z: camDist,
			},
			TargetY: globalR * 0.5,
			Fov:     50,
		}).Stage(stage3d)
	}

	canvas.Camera.OnUpdate = func(updatedCamera *threejs.Camera) {
		if rendered3DShape != nil {
			rendered3DShape.ViewX = updatedCamera.X
			rendered3DShape.ViewY = updatedCamera.Y
			rendered3DShape.ViewZ = updatedCamera.Z
			rendered3DShape.TargetX = updatedCamera.TargetX
			rendered3DShape.TargetY = updatedCamera.TargetY
			rendered3DShape.TargetZ = updatedCamera.TargetZ
			rendered3DShape.Fov = updatedCamera.Fov
			stager.GetStage().CommitWithSuspendedCallbacks()
		}
	}

	floorMinY := math.MaxFloat64

	res := &Cylinder3DBaseResult{
		Canvas:            canvas,
		GlobalR:           globalR,
		RadialRepetitions: radialRepetitions,
		FloorMinY:         floorMinY,
	}

	// Extract GrowthCurve2D arcs from PlantAbstract (StartArcShapeGrid and EndArcShapeGrid)
	if plant.StartArcShapeGrid != nil && len(plant.StartArcShapeGrid.StartArcShapes) > 0 {
		startArcs := plant.StartArcShapeGrid.StartArcShapes
		var endArcs []*models.EndArcShape
		if plant.EndArcShapeGrid != nil {
			endArcs = plant.EndArcShapeGrid.EndArcShapes
		}

		baseCurve := (&threejs.Curve{
			Name: fmt.Sprintf("%s Base Curve", params.NamePrefix),
		}).Stage(stage3d)

		for i := 0; i < len(startArcs); i++ {
			sa := startArcs[i]
			AppendArcPointsCylinder(stage3d, baseCurve, sa.StartX, sa.StartY, sa.EndX, sa.EndY, sa.RadiusX, !sa.SweepFlag, sa.LargeArcFlag, globalR, 0.0, &floorMinY)

			if i < len(endArcs) {
				ea := endArcs[i]
				AppendArcPointsCylinder(stage3d, baseCurve, ea.EndX, ea.EndY, ea.StartX, ea.StartY, ea.RadiusX, ea.SweepFlag, ea.LargeArcFlag, globalR, 0.0, &floorMinY)
			}
		}

		sortedAngles, sortedPoints := UnwrapAngles(baseCurve)
		res.SortedAngles = sortedAngles
		res.SortedPoints = sortedPoints

		degreeInterval := 0.5
		expectedDegrees := 360.0 / float64(radialRepetitions)
		res.ExpectedDegrees = expectedDegrees
		radInterval := degreeInterval * math.Pi / 180.0
		nbPoints := int(math.Round(expectedDegrees / degreeInterval))

		var targetAngles []float64
		for i := 0; i <= nbPoints; i++ {
			targetAngles = append(targetAngles, float64(i)*radInterval)
		}
		res.TargetAngles = targetAngles

		resampledBaseCurve := ResampleCurveAtAngles(stage3d, sortedAngles, sortedPoints, targetAngles, params.NamePrefix, expectedDegrees)

		vertScale := params.VerticalScale
		if vertScale <= 0.0 {
			vertScale = 1.0
		}
		for _, pt := range resampledBaseCurve.Points {
			pt.Y = pt.Y * vertScale
		}
		res.ResampledBaseCurve = resampledBaseCurve
		res.VertScale = vertScale

		relDiameter := params.RelativeTubeDiameter
		if relDiameter <= 0 {
			relDiameter = 0.01
		}
		tubeRadius := (relDiameter * plant.RhombusSideLength) / 2.0
		if tubeRadius <= 0 {
			tubeRadius = 2.0
		}
		res.TubeRadius = tubeRadius

		transparency := params.Transparency
		opacity := 1.0 - transparency
		if opacity < 0.0 {
			opacity = 0.0
		}
		if opacity > 1.0 {
			opacity = 1.0
		}
		res.Opacity = opacity

		topHeight := params.RelativeHeight * plant.RhombusSideLength
		torusHeight := params.RelativeHeight3DTorus * plant.RhombusSideLength
		res.TopHeight = topHeight
		res.TorusHeight = torusHeight

		createTorusLayer := func(dx, dy float64, namePrefix string, color string) {
			thetaOffset := dx / globalR

			layerCurve := (&threejs.Curve{
				Name: fmt.Sprintf("%s Curve", namePrefix),
			}).Stage(stage3d)

			for k := 0; k < radialRepetitions; k++ {
				baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(radialRepetitions)
				totalThetaOffset := baseThetaOffset + thetaOffset

				for _, pt := range resampledBaseCurve.Points {
					origTheta := math.Atan2(pt.Z, pt.X)
					r := math.Hypot(pt.X, pt.Z)
					newTheta := origTheta + totalThetaOffset

					layerPtY := pt.Y + dy + torusHeight
					if layerPtY < floorMinY {
						floorMinY = layerPtY
					}

					layerCurve.Points = append(layerCurve.Points, (&threejs.Vector3{
						Name: fmt.Sprintf("%s Point k%d %.1f", namePrefix, k, newTheta*180.0/math.Pi),
						X:    r * math.Cos(newTheta),
						Y:    layerPtY,
						Z:    r * math.Sin(newTheta),
					}).Stage(stage3d))
				}
			}

			numSegments := len(layerCurve.Points)
			if numSegments < 2 {
				numSegments = 2
			}

			tGeom := (&threejs.TubeGeometry{
				Name:            fmt.Sprintf("%s TubeGeom", namePrefix),
				Path:            layerCurve,
				TubularSegments: numSegments,
				Radius:          tubeRadius,
				RadialSegments:  16,
				Closed:          true,
			}).Stage(stage3d)

			tubeMesh := (&threejs.Mesh{
				Name:         fmt.Sprintf("%s Mesh", namePrefix),
				Position:     threejs.Position{X: 0, Y: 0, Z: 0},
				TubeGeometry: tGeom,
				MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
					Name:                 fmt.Sprintf("%s Material", namePrefix),
					MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: color},
					Transparent:          true,
					Opacity:              opacity,
				}).Stage(stage3d),
			}).Stage(stage3d)

			canvas.Meshs = append(canvas.Meshs, tubeMesh)
		}

		// Base 3D Torus
		if !params.IsHiddenTorus3DShape {
			createTorusLayer(0.0, 0.0, fmt.Sprintf("%s Base Torus", params.NamePrefix), "darkgreen")
		}

		var growthVectorX, growthVectorY float64
		if plant.GrowthVectorShape != nil {
			growthVectorX = plant.GrowthVectorShape.X
			growthVectorY = plant.GrowthVectorShape.Y * vertScale
		}
		res.GrowthVectorX = growthVectorX
		res.GrowthVectorY = growthVectorY

		// Rotated Torus (if enabled)
		if params.HasRotatedShapes && !params.IsHiddenRotatedTorusShape {
			createTorusLayer(growthVectorX, growthVectorY, fmt.Sprintf("%s Partially Rotated Torus", params.NamePrefix), "darkgreen")
		}

		projAngleRad := -params.ProjectionAngle * math.Pi / 180.0
		res.ProjAngleRad = projAngleRad

		// Top 2D Projected Curve on the horizontal top plane (from Base Torus)
		if !params.IsHiddenTopCurveShape {
			topCurve := (&threejs.Curve{
				Name: fmt.Sprintf("%s Top Curve", params.NamePrefix),
			}).Stage(stage3d)

			for k := 0; k < radialRepetitions; k++ {
				baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(radialRepetitions)

				for _, pt := range resampledBaseCurve.Points {
					origTheta := math.Atan2(pt.Z, pt.X)
					r := math.Hypot(pt.X, pt.Z)
					newTheta := origTheta + baseThetaOffset

					ptY := pt.Y + torusHeight
					deltaY := topHeight - ptY
					rProj := r + deltaY*math.Tan(projAngleRad)

					topCurve.Points = append(topCurve.Points, (&threejs.Vector3{
						Name: fmt.Sprintf("%s Top Point k%d %.1f", params.NamePrefix, k, newTheta*180.0/math.Pi),
						X:    rProj * math.Cos(newTheta),
						Y:    topHeight,
						Z:    rProj * math.Sin(newTheta),
					}).Stage(stage3d))
				}
			}

			numSegments := len(topCurve.Points)
			if numSegments < 2 {
				numSegments = 2
			}

			stGeom := (&threejs.TubeGeometry{
				Name:            fmt.Sprintf("%s Top TubeGeom", params.NamePrefix),
				Path:            topCurve,
				TubularSegments: numSegments,
				Radius:          tubeRadius,
				RadialSegments:  16,
				Closed:          true,
			}).Stage(stage3d)

			stMesh := (&threejs.Mesh{
				Name:         fmt.Sprintf("%s Top Mesh", params.NamePrefix),
				Position:     threejs.Position{X: 0, Y: 0, Z: 0},
				TubeGeometry: stGeom,
				MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
					Name:                 fmt.Sprintf("%s Top Material", params.NamePrefix),
					MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "royalblue"},
					Transparent:          true,
					Opacity:              opacity,
				}).Stage(stage3d),
			}).Stage(stage3d)

			canvas.Meshs = append(canvas.Meshs, stMesh)
		}

		if params.HasRotatedShapes {
			var rotTopPoints []*threejs.Vector3
			thetaOffset := growthVectorX / globalR

			for k := 0; k < radialRepetitions; k++ {
				baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(radialRepetitions)
				totalThetaOffset := baseThetaOffset + thetaOffset

				for _, pt := range resampledBaseCurve.Points {
					origTheta := math.Atan2(pt.Z, pt.X)
					r := math.Hypot(pt.X, pt.Z)
					newTheta := origTheta + totalThetaOffset

					ptY := pt.Y + growthVectorY + torusHeight

					deltaYTop := topHeight - ptY
					rProjTop := r + deltaYTop*math.Tan(projAngleRad)

					rotTopPoints = append(rotTopPoints, (&threejs.Vector3{
						Name: fmt.Sprintf("Rotated %s Top Point k%d %.1f", params.NamePrefix, k, newTheta*180.0/math.Pi),
						X:    rProjTop * math.Cos(newTheta),
						Y:    topHeight,
						Z:    rProjTop * math.Sin(newTheta),
					}).Stage(stage3d))
				}
			}
			res.RotTopPoints = rotTopPoints

			if !params.IsHiddenRotatedTopCurveShape {
				rotTopCurve := (&threejs.Curve{
					Name:   fmt.Sprintf("%s Partially Rotated Top Curve", params.NamePrefix),
					Points: rotTopPoints,
				}).Stage(stage3d)
				res.RotTopCurve = rotTopCurve

				numSegments := len(rotTopCurve.Points)
				if numSegments < 2 {
					numSegments = 2
				}

				rotStGeom := (&threejs.TubeGeometry{
					Name:            fmt.Sprintf("%s Partially Rotated Top TubeGeom", params.NamePrefix),
					Path:            rotTopCurve,
					TubularSegments: numSegments,
					Radius:          tubeRadius,
					RadialSegments:  16,
					Closed:          true,
				}).Stage(stage3d)

				rotStMesh := (&threejs.Mesh{
					Name:         fmt.Sprintf("%s Partially Rotated Top Mesh", params.NamePrefix),
					Position:     threejs.Position{X: 0, Y: 0, Z: 0},
					TubeGeometry: rotStGeom,
					MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
						Name:                 fmt.Sprintf("%s Partially Rotated Top Material", params.NamePrefix),
						MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "darkorange"},
						Transparent:          true,
						Opacity:              opacity,
					}).Stage(stage3d),
				}).Stage(stage3d)

				canvas.Meshs = append(canvas.Meshs, rotStMesh)
			}
		}

		// Sampled points visualization
		if !params.IsHiddenSampledPoints3DShape {
			numPointsPerRep := len(resampledBaseCurve.Points)
			var basePoints []*threejs.Vector3
			for k := 0; k < radialRepetitions; k++ {
				baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(radialRepetitions)
				for _, pt := range resampledBaseCurve.Points {
					origTheta := math.Atan2(pt.Z, pt.X)
					r := math.Hypot(pt.X, pt.Z)
					newTheta := origTheta + baseThetaOffset
					basePoints = append(basePoints, (&threejs.Vector3{
						Name: fmt.Sprintf("%s Sampled Point k%d %.1f", params.NamePrefix, k, newTheta*180.0/math.Pi),
						X:    r * math.Cos(newTheta),
						Y:    pt.Y + torusHeight,
						Z:    r * math.Sin(newTheta),
					}).Stage(stage3d))
				}
			}
			AddPointSpheres(stage3d, basePoints, "red", canvas, fmt.Sprintf("%s Sampled", params.NamePrefix), 0, numPointsPerRep)
		}

		// Rotated Sampled points visualization (if enabled)
		if params.HasRotatedShapes && !params.IsHiddenRotatedSampledPoints3DShape {
			numPointsPerRep := len(resampledBaseCurve.Points)
			var rotPoints []*threejs.Vector3
			thetaOffset := growthVectorX / globalR

			for k := 0; k < radialRepetitions; k++ {
				baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(radialRepetitions)
				totalThetaOffset := baseThetaOffset + thetaOffset

				for _, pt := range resampledBaseCurve.Points {
					origTheta := math.Atan2(pt.Z, pt.X)
					r := math.Hypot(pt.X, pt.Z)
					newTheta := origTheta + totalThetaOffset
					rotPoints = append(rotPoints, (&threejs.Vector3{
						Name: fmt.Sprintf("%s Rotated Sampled Point k%d %.1f", params.NamePrefix, k, newTheta*180.0/math.Pi),
						X:    r * math.Cos(newTheta),
						Y:    pt.Y + growthVectorY + torusHeight,
						Z:    r * math.Sin(newTheta),
					}).Stage(stage3d))
				}
			}
			AddPointSpheres(stage3d, rotPoints, "orange", canvas, fmt.Sprintf("%s Rotated Sampled", params.NamePrefix), 0, numPointsPerRep)
		}

		res.FloorMinY = floorMinY
	}

	return res
}

func AppendArcPointsCylinder(stool3dStage *threejs.Stage, targetCurve *threejs.Curve, x1, y1, x2, y2, r float64, sweepFlag bool, largeArcFlag bool, globalR float64, baseThetaOffset float64, floorMinY *float64) {
	dx := (x1 - x2) / 2.0
	dy := (y1 - y2) / 2.0
	d2 := dx*dx + dy*dy
	var cx, cy float64
	if d2 == 0 || r*r < d2 {
		cx = (x1 + x2) / 2.0
		cy = (y1 + y2) / 2.0
		r = math.Sqrt(d2)
	} else {
		root := math.Sqrt(r*r/d2 - 1.0)
		if largeArcFlag == sweepFlag {
			root = -root
		}
		cx = (x1+x2)/2.0 + root*dy
		cy = (y1+y2)/2.0 - root*dx
	}

	startAngle := math.Atan2(y1-cy, x1-cx)
	endAngle := math.Atan2(y2-cy, x2-cx)

	if sweepFlag {
		for endAngle < startAngle {
			endAngle += 2 * math.Pi
		}
	} else {
		for endAngle > startAngle {
			endAngle -= 2 * math.Pi
		}
	}

	steps := 50
	for i := 0; i <= steps; i++ {
		if i == 0 && len(targetCurve.Points) > 0 {
			continue
		}

		t := float64(i) / float64(steps)
		angle := startAngle + t*(endAngle-startAngle)
		x2d := cx + r*math.Cos(angle)
		y2d := cy + r*math.Sin(angle)

		theta := x2d/globalR + baseThetaOffset
		x3d := globalR * math.Cos(theta)
		z3d := globalR * math.Sin(theta)
		y3d := y2d

		if y3d < *floorMinY {
			*floorMinY = y3d
		}

		vec := (&threejs.Vector3{
			Name: fmt.Sprintf("Base Point %d", len(targetCurve.Points)),
			X:    x3d,
			Y:    y3d,
			Z:    z3d,
		}).Stage(stool3dStage)
		targetCurve.Points = append(targetCurve.Points, vec)
	}
}

// UnwrapAngles processes a 3D curve to extract a strictly monotonic,
// duplicate-free mapping of points to their cylindrical angle (theta).
func UnwrapAngles(curve *threejs.Curve) (angles []float64, points []*threejs.Vector3) {
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

// ResampleCurveAtAngles forces an existing 3D curve to conform to a specific set of target angles.
func ResampleCurveAtAngles(
	stool3dStage *threejs.Stage,
	sortedAngles []float64,
	sortedPoints []*threejs.Vector3,
	targetAngles []float64,
	namePrefix string,
	expectedDegrees float64,
) *threejs.Curve {
	resampled := (&threejs.Curve{
		Name: fmt.Sprintf("%s Resampled", namePrefix),
	}).Stage(stool3dStage)

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

		// Use binary search to find the first index where sortedAngles[i] >= evalTarget
		searchIdx := sort.SearchFloat64s(sortedAngles, evalTarget)

		if searchIdx > 0 && searchIdx < len(sortedAngles) {
			idx = searchIdx - 1
		} else if searchIdx == 0 && sortedAngles[0] == evalTarget {
			idx = 0
		}

		isExtrapolating := false
		if idx == -1 {
			isExtrapolating = true
			if searchIdx == 0 {
				idx = 0
			} else {
				idx = len(sortedAngles) - 2
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
			y = p1.Y + t*(p2.Y-p1.Y)
			x = r * math.Cos(target)
			z = r * math.Sin(target)
		}

		pt := (&threejs.Vector3{
			Name: fmt.Sprintf("%s %.1f", namePrefix, target*180.0/math.Pi),
			X:    x,
			Y:    y,
			Z:    z,
		}).Stage(stool3dStage)

		resampled.Points = append(resampled.Points, pt)
	}

	return resampled
}

func AddPointSpheres(stool3dStage *threejs.Stage, points []*threejs.Vector3, color string, canvas *threejs.Canvas, namePrefix string, dy float64, numPointsPerRep int) {
	for i, pt := range points {
		sphereColor := color
		radius := 2.0

		localIdx := i
		if numPointsPerRep > 0 {
			localIdx = i % numPointsPerRep
		}

		if localIdx%20 == 0 {
			sphereColor = "yellow"
			radius = 4.0
		}

		sphere := (&threejs.Mesh{
			Name: fmt.Sprintf("%s Sphere %d", namePrefix, i),
			Position: threejs.Position{
				X: pt.X,
				Y: pt.Y + dy,
				Z: pt.Z,
			},
			SphereGeometry: (&threejs.SphereGeometry{
				Name:   fmt.Sprintf("%s SphereGeom %d", namePrefix, i),
				Radius: radius,
			}).Stage(stool3dStage),
			MeshMaterialBasic: (&threejs.MeshMaterialBasic{
				Name: fmt.Sprintf("%s SphereMat %d", namePrefix, i),
				MeshMaterialAbstract: threejs.MeshMaterialAbstract{
					Color: sphereColor,
				},
			}).Stage(stool3dStage),
		}).Stage(stool3dStage)
		canvas.Meshs = append(canvas.Meshs, sphere)
	}
}

func AddFloorTiles(stage3d *threejs.Stage, canvas *threejs.Canvas, globalR float64, floorMinY float64) {
	if floorMinY > 0.0 {
		floorMinY = 0.0
	}
	floorMinY = floorMinY - 2.0

	floorSize := globalR * 3.0
	if floorSize < 200 {
		floorSize = 200
	}
	gridSize := 20
	tileSize := floorSize / float64(gridSize)

	for i := -gridSize / 2; i < gridSize/2; i++ {
		for j := -gridSize / 2; j < gridSize/2; j++ {
			color := "white"
			if (i+j)%2 != 0 {
				color = "black"
			}

			tileMesh := (&threejs.Mesh{
				Name: "Floor Tile " + strconv.Itoa(i) + "-" + strconv.Itoa(j),
				Position: threejs.Position{
					X: float64(i)*tileSize + tileSize/2,
					Y: floorMinY - 0.05,
					Z: float64(j)*tileSize + tileSize/2,
				},
				BoxGeometry: (&threejs.BoxGeometry{
					Name:           "Tile Geometry",
					Width:          tileSize,
					Height:         0.1,
					Depth:          tileSize,
					WidthSegments:  1,
					HeightSegments: 1,
					DepthSegments:  1,
				}).Stage(stage3d),
				MeshMaterialBasic: (&threejs.MeshMaterialBasic{
					Name:                 "Tile Material " + color,
					MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: color},
				}).Stage(stage3d),
			}).Stage(stage3d)

			canvas.Meshs = append(canvas.Meshs, tileMesh)
		}
	}
}
