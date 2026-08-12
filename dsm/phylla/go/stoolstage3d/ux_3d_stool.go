package stoolstage3d

import (
	"fmt"
	"math"
	"sort"

	"github.com/fullstack-lang/gong/dsm/phylla/go/cylinderstage3d"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func (u *Stool3DStageUpdater) ux_3d_stool(stager *models.Stager) {
	stool3dStage := stager.GetStool3dStage()
	if stool3dStage == nil {
		return
	}

	plant := stager.GetCurrentPlant()
	if plant == nil || plant.PlantType != models.Stool || plant.StoolAbstract == nil {
		stool3dStage.Reset()
		stool3dStage.Commit()
		return
	}

	var checkedDiagram *models.PlantDiagram
	for _, d := range plant.PlantDiagrams {
		if d.IsChecked {
			checkedDiagram = d
			break
		}
	}

	params := cylinderstage3d.Cylinder3DParams{
		NamePrefix:                          "Stool",
		CanvasName:                          "Stool 3D Canvas",
		RadialRepetitions:                   plant.StoolAbstract.RadialRepetitions,
		Transparency:                        plant.StoolAbstract.Transparency,
		RelativeTubeDiameter:                plant.StoolAbstract.RelativeTubeDiameter,
		RelativeHeight3DTorus:               plant.StoolAbstract.RelativeHeight3DTorus,
		VerticalScale:                       plant.StoolAbstract.StoolTorusVerticalScale,
		RelativeHeight:                      plant.StoolAbstract.RelativeHeight,
		ProjectionAngle:                     plant.StoolAbstract.ProjectionAngle,
		HasRotatedShapes:                    true,
	}

	if checkedDiagram != nil && checkedDiagram.StoolDiagram != nil {
		params.Rendered3DShape = checkedDiagram.StoolDiagram.Rendered3DShape
		params.IsHiddenTorus3DShape = checkedDiagram.StoolDiagram.IsHiddenTorus3DShape
		params.IsHiddenRotatedTorusShape = checkedDiagram.StoolDiagram.IsHiddenRotatedTorusShape
		params.IsHiddenTopCurveShape = checkedDiagram.StoolDiagram.IsHiddenSeatTopCurveShape
		params.IsHiddenRotatedTopCurveShape = checkedDiagram.StoolDiagram.IsHiddenRotatedSeatTopCurveShape
		params.IsHiddenSampledPoints3DShape = checkedDiagram.StoolDiagram.IsHiddenSampledPoints3DShape
		params.IsHiddenRotatedSampledPoints3DShape = checkedDiagram.StoolDiagram.IsHiddenRotatedSampledPoints3DShape
	}

	base := cylinderstage3d.RenderCylinder3DBase(stool3dStage, stager, plant, params)
	if base == nil || base.ResampledBaseCurve == nil {
		stool3dStage.Commit()
		return
	}

	canvas := base.Canvas
	resampledBaseCurve := base.ResampledBaseCurve
	radialRepetitions := base.RadialRepetitions
	globalR := base.GlobalR
	tubeRadius := base.TubeRadius
	opacity := base.Opacity
	stoolTopHeight := base.TopHeight
	torusHeight := base.TorusHeight
	growthVectorX := base.GrowthVectorX
	growthVectorY := base.GrowthVectorY
	projAngleRad := base.ProjAngleRad
	vertScale := base.VertScale
	floorMinY := base.FloorMinY
	expectedDegrees := base.ExpectedDegrees
	targetAngles := base.TargetAngles
	rotSeatTopPoints := base.RotTopPoints
	_ = vertScale
	_ = floorMinY

	if plant.StartArcShapeGrid != nil && len(plant.StartArcShapeGrid.StartArcShapes) > 0 {
		seatThickness := plant.StoolAbstract.RelativeSeatThickness * plant.RhombusSideLength
		seatBottomHeight := stoolTopHeight - seatThickness

		var rotSeatBottomPoints []*threejs.Vector3
		thetaOffset := growthVectorX / globalR

		for k := 0; k < radialRepetitions; k++ {
			baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(radialRepetitions)
			totalThetaOffset := baseThetaOffset + thetaOffset

			for _, pt := range resampledBaseCurve.Points {
				origTheta := math.Atan2(pt.Z, pt.X)
				r := math.Hypot(pt.X, pt.Z)
				newTheta := origTheta + totalThetaOffset

				ptY := pt.Y + growthVectorY + torusHeight

				deltaYBottom := seatBottomHeight - ptY
				rProjBottom := r + deltaYBottom*math.Tan(projAngleRad)

				rotSeatBottomPoints = append(rotSeatBottomPoints, (&threejs.Vector3{
					Name: fmt.Sprintf("Rotated Seat Bottom Point k%d %.1f", k, newTheta*180.0/math.Pi),
					X:    rProjBottom * math.Cos(newTheta),
					Y:    seatBottomHeight,
					Z:    rProjBottom * math.Sin(newTheta),
				}).Stage(stool3dStage))
			}
		}

		// 10. Seat Bottom 2D Projected Curve on the horizontal stool seat bottom plane (from Base Torus)
		if checkedDiagram != nil && checkedDiagram.StoolDiagram != nil && !checkedDiagram.StoolDiagram.IsHiddenSeatBottomCurveShape {
			seatBottomCurve := (&threejs.Curve{
				Name: "Stool Seat Bottom Curve",
			}).Stage(stool3dStage)

			for k := 0; k < radialRepetitions; k++ {
				baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(radialRepetitions)

				for _, pt := range resampledBaseCurve.Points {
					origTheta := math.Atan2(pt.Z, pt.X)
					r := math.Hypot(pt.X, pt.Z)
					newTheta := origTheta + baseThetaOffset

					ptY := pt.Y + torusHeight
					deltaY := seatBottomHeight - ptY
					rProj := r + deltaY*math.Tan(projAngleRad)

					seatBottomCurve.Points = append(seatBottomCurve.Points, (&threejs.Vector3{
						Name: fmt.Sprintf("Seat Bottom Point k%d %.1f", k, newTheta*180.0/math.Pi),
						X:    rProj * math.Cos(newTheta),
						Y:    seatBottomHeight,
						Z:    rProj * math.Sin(newTheta),
					}).Stage(stool3dStage))
				}
			}

			numSegments := len(seatBottomCurve.Points)
			if numSegments < 2 {
				numSegments = 2
			}

			sbGeom := (&threejs.TubeGeometry{
				Name:            "Stool Seat Bottom TubeGeom",
				Path:            seatBottomCurve,
				TubularSegments: numSegments,
				Radius:          tubeRadius,
				RadialSegments:  16,
				Closed:          true,
			}).Stage(stool3dStage)

			sbMesh := (&threejs.Mesh{
				Name:         "Stool Seat Bottom Mesh",
				Position:     threejs.Position{X: 0, Y: 0, Z: 0},
				TubeGeometry: sbGeom,
				MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
					Name:                 "Stool Seat Bottom Material",
					MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "dodgerblue"},
					Transparent:          true,
					Opacity:              opacity,
				}).Stage(stool3dStage),
			}).Stage(stool3dStage)

			canvas.Meshs = append(canvas.Meshs, sbMesh)
		}

		// 11. Rotated Seat Bottom 2D Projected Curve (from Partially Rotated Torus)
		if checkedDiagram != nil && checkedDiagram.StoolDiagram != nil && !checkedDiagram.StoolDiagram.IsHiddenRotatedSeatBottomCurveShape {
			rotSeatBottomCurve := (&threejs.Curve{
				Name:   "Stool Rotated Seat Bottom Curve",
				Points: rotSeatBottomPoints,
			}).Stage(stool3dStage)

			numSegments := len(rotSeatBottomCurve.Points)
			if numSegments < 2 {
				numSegments = 2
			}

			rotSbGeom := (&threejs.TubeGeometry{
				Name:            "Stool Rotated Seat Bottom TubeGeom",
				Path:            rotSeatBottomCurve,
				TubularSegments: numSegments,
				Radius:          tubeRadius,
				RadialSegments:  16,
				Closed:          true,
			}).Stage(stool3dStage)

			rotSbMesh := (&threejs.Mesh{
				Name:         "Stool Rotated Seat Bottom Mesh",
				Position:     threejs.Position{X: 0, Y: 0, Z: 0},
				TubeGeometry: rotSbGeom,
				MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
					Name:                 "Stool Rotated Seat Bottom Material",
					MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "coral"},
					Transparent:          true,
					Opacity:              opacity,
				}).Stage(stool3dStage),
			}).Stage(stool3dStage)

			canvas.Meshs = append(canvas.Meshs, rotSbMesh)
		}

		// 11b. 3D Seat Volume Mesh between Rotated Seat Top and Rotated Seat Bottom
		if checkedDiagram == nil || checkedDiagram.StoolDiagram == nil || !checkedDiagram.StoolDiagram.IsHiddenSeat3DShape {
			if len(rotSeatTopPoints) >= 3 && len(rotSeatTopPoints) == len(rotSeatBottomPoints) {
				N := len(rotSeatTopPoints)
				seatGeom := (&threejs.BufferGeometry{
					Name: "Stool Seat BufferGeometry",
				}).Stage(stool3dStage)

				var sumTopX, sumTopZ, sumBottomX, sumBottomZ float64
				for i := 0; i < N; i++ {
					topV := (&threejs.Vector3{
						Name: fmt.Sprintf("Seat Top V %d", i),
						X:    rotSeatTopPoints[i].X,
						Y:    rotSeatTopPoints[i].Y,
						Z:    rotSeatTopPoints[i].Z,
					}).Stage(stool3dStage)
					seatGeom.Vertices = append(seatGeom.Vertices, topV)
					sumTopX += topV.X
					sumTopZ += topV.Z
				}

				for i := 0; i < N; i++ {
					botV := (&threejs.Vector3{
						Name: fmt.Sprintf("Seat Bottom V %d", i),
						X:    rotSeatBottomPoints[i].X,
						Y:    rotSeatBottomPoints[i].Y,
						Z:    rotSeatBottomPoints[i].Z,
					}).Stage(stool3dStage)
					seatGeom.Vertices = append(seatGeom.Vertices, botV)
					sumBottomX += botV.X
					sumBottomZ += botV.Z
				}

				topCenterIdx := len(seatGeom.Vertices)
				topCenterV := (&threejs.Vector3{
					Name: "Seat Top Center",
					X:    sumTopX / float64(N),
					Y:    stoolTopHeight,
					Z:    sumTopZ / float64(N),
				}).Stage(stool3dStage)
				seatGeom.Vertices = append(seatGeom.Vertices, topCenterV)

				botCenterIdx := len(seatGeom.Vertices)
				botCenterV := (&threejs.Vector3{
					Name: "Seat Bottom Center",
					X:    sumBottomX / float64(N),
					Y:    seatBottomHeight,
					Z:    sumBottomZ / float64(N),
				}).Stage(stool3dStage)
				seatGeom.Vertices = append(seatGeom.Vertices, botCenterV)

				// 1. Top face (facing +Y): (topCenter, nextI, i)
				for i := 0; i < N; i++ {
					nextI := (i + 1) % N
					seatGeom.Faces = append(seatGeom.Faces, (&threejs.Triangle{
						Name: fmt.Sprintf("Seat Top Face %d", i),
						V1:   topCenterIdx,
						V2:   nextI,
						V3:   i,
					}).Stage(stool3dStage))
				}

				// 2. Bottom face (facing -Y): (botCenter, botI, botNextI)
				for i := 0; i < N; i++ {
					nextI := (i + 1) % N
					botI := N + i
					botNextI := N + nextI
					seatGeom.Faces = append(seatGeom.Faces, (&threejs.Triangle{
						Name: fmt.Sprintf("Seat Bottom Face %d", i),
						V1:   botCenterIdx,
						V2:   botI,
						V3:   botNextI,
					}).Stage(stool3dStage))
				}

				// 3. Side wall quads between Top and Bottom:
				for i := 0; i < N; i++ {
					nextI := (i + 1) % N
					topI := i
					topNextI := nextI
					botI := N + i
					botNextI := N + nextI

					// Triangle 1: (botI, topI, topNextI)
					seatGeom.Faces = append(seatGeom.Faces, (&threejs.Triangle{
						Name: fmt.Sprintf("Seat Wall T1 %d", i),
						V1:   botI,
						V2:   topI,
						V3:   topNextI,
					}).Stage(stool3dStage))

					// Triangle 2: (botI, topNextI, botNextI)
					seatGeom.Faces = append(seatGeom.Faces, (&threejs.Triangle{
						Name: fmt.Sprintf("Seat Wall T2 %d", i),
						V1:   botI,
						V2:   topNextI,
						V3:   botNextI,
					}).Stage(stool3dStage))
				}

				seatMesh := (&threejs.Mesh{
					Name:           "Stool Seat 3D Mesh",
					Position:       threejs.Position{X: 0, Y: 0, Z: 0},
					BufferGeometry: seatGeom,
					MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
						Name:                 "Stool Seat Material",
						MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "peru"},
						Transparent:          true,
						Opacity:              opacity,
					}).Stage(stool3dStage),
				}).Stage(stool3dStage)

				canvas.Meshs = append(canvas.Meshs, seatMesh)
			}
		}

		// Compute Eye geometry, transitions, and Bézier corners for first repetition (k=0)
		thetaOffset = growthVectorX / globalR
		eyeCriteria := plant.StoolAbstract.RelativeEyeSeparationCriteria * plant.RhombusSideLength * vertScale
		expectedRad := expectedDegrees * math.Pi / 180.0
		radInterval := 0.5 * math.Pi / 180.0
		dStep := globalR * radInterval
		if dStep <= 0 {
			dStep = 1.0
		}

		getYAtAngle := func(evalAngle float64) float64 {
			if len(resampledBaseCurve.Points) == 0 {
				return 0
			}
			for evalAngle < 0 {
				evalAngle += expectedRad
			}
			for evalAngle > expectedRad {
				evalAngle -= expectedRad
			}
			idx := int(math.Floor(evalAngle / radInterval))
			if idx < 0 {
				idx = 0
			}
			if idx >= len(resampledBaseCurve.Points)-1 {
				return resampledBaseCurve.Points[len(resampledBaseCurve.Points)-1].Y
			}
			t := (evalAngle - float64(idx)*radInterval) / radInterval
			y0 := resampledBaseCurve.Points[idx].Y
			y1 := resampledBaseCurve.Points[idx+1].Y
			return y0 + t*(y1-y0)
		}

		// Identify transitions into and out of eye regions
		numPts := len(resampledBaseCurve.Points)
		inEye := make([]bool, numPts)
		yBaseList := make([]float64, numPts)
		yRotList := make([]float64, numPts)

		for i, pt := range resampledBaseCurve.Points {
			alpha := targetAngles[i]
			yBaseList[i] = pt.Y + torusHeight
			yRotList[i] = getYAtAngle(alpha-thetaOffset) + growthVectorY + torusHeight
			dist := math.Abs(yBaseList[i] - yRotList[i])
			inEye[i] = (dist > eyeCriteria)
		}

		type point2D struct {
			U, Y float64
		}

		evalBezier := func(c0, c1, c2, c3 point2D, t float64) point2D {
			omt := 1.0 - t
			omt2 := omt * omt
			omt3 := omt2 * omt
			t2 := t * t
			t3 := t2 * t
			return point2D{
				U: omt3*c0.U + 3*omt2*t*c1.U + 3*omt*t2*c2.U + t3*c3.U,
				Y: omt3*c0.Y + 3*omt2*t*c1.Y + 3*omt*t2*c2.Y + t3*c3.Y,
			}
		}

		getBezierPoints := func(c0, c1, c2, c3 point2D, cornerName string) []*threejs.Vector3 {
			const numSub = 200
			pts := make([]point2D, numSub+1)
			cumLen := make([]float64, numSub+1)
			totalLen := 0.0
			pts[0] = c0
			cumLen[0] = 0.0

			for k := 1; k <= numSub; k++ {
				t := float64(k) / float64(numSub)
				pts[k] = evalBezier(c0, c1, c2, c3, t)
				du := pts[k].U - pts[k-1].U
				dy := pts[k].Y - pts[k-1].Y
				totalLen += math.Hypot(du, dy)
				cumLen[k] = totalLen
			}

			n := int(math.Round(totalLen / dStep))
			if n < 2 {
				n = 2
			}

			var result []*threejs.Vector3
			for j := 1; j < n; j++ {
				targetDist := float64(j) * (totalLen / float64(n))
				searchIdx := sort.SearchFloat64s(cumLen, targetDist)
				if searchIdx <= 0 {
					searchIdx = 1
				}
				if searchIdx > numSub {
					searchIdx = numSub
				}
				segL := cumLen[searchIdx] - cumLen[searchIdx-1]
				segT := 0.0
				if segL > 0 {
					segT = (targetDist - cumLen[searchIdx-1]) / segL
				}
				uVal := pts[searchIdx-1].U + segT*(pts[searchIdx].U-pts[searchIdx-1].U)
				yVal := pts[searchIdx-1].Y + segT*(pts[searchIdx].Y-pts[searchIdx-1].Y)

				theta := uVal / globalR
				x := globalR * math.Cos(theta)
				z := globalR * math.Sin(theta)

				result = append(result, (&threejs.Vector3{
					Name: fmt.Sprintf("%s Bezier Point %d/%.1f", cornerName, j, theta*180.0/math.Pi),
					X:    x,
					Y:    yVal,
					Z:    z,
				}).Stage(stool3dStage))
			}
			return result
		}

		normalize2D := func(du, dy float64, defU, defY float64) (float64, float64) {
			lenV := math.Hypot(du, dy)
			if lenV > 1e-7 {
				return du / lenV, dy / lenV
			}
			return defU, defY
		}

		controlStrength := plant.StoolAbstract.RelativeEyeCornerControlVectorStrength
		if controlStrength <= 0.0 {
			controlStrength = 0.55
		}

		// Find eye segment endpoints
		iStart := -1
		iEnd := -1
		for i := 0; i < numPts; i++ {
			if inEye[i] {
				if iStart == -1 {
					iStart = i
				}
				iEnd = i
			}
		}

		var leftCornerPts []*threejs.Vector3
		var rightCornerPts []*threejs.Vector3

		if iStart != -1 && iEnd != -1 {
			// Left corner (connects Top -> Bottom at iStart)
			{
				i0 := iStart
				i1 := iStart + 1
				if i1 >= numPts {
					i1 = i0
				}

				pTop0 := point2D{U: targetAngles[i0] * globalR, Y: yRotList[i0]}
				pTop1 := point2D{U: targetAngles[i1] * globalR, Y: yRotList[i1]}
				tTopExitU, tTopExitY := normalize2D(pTop0.U-pTop1.U, pTop0.Y-pTop1.Y, -1.0, 0.0)

				pBottom0 := point2D{U: targetAngles[i0] * globalR, Y: yBaseList[i0]}
				pBottom1 := point2D{U: targetAngles[i1] * globalR, Y: yBaseList[i1]}
				tBottomEnterU, tBottomEnterY := normalize2D(pBottom1.U-pBottom0.U, pBottom1.Y-pBottom0.Y, 1.0, 0.0)

				c0 := pTop0
				c3 := pBottom0
				chord := math.Abs(c3.Y - c0.Y)
				handleLen := controlStrength * chord

				c1 := point2D{U: c0.U + handleLen*tTopExitU, Y: c0.Y + handleLen*tTopExitY}
				c2 := point2D{U: c3.U - handleLen*tBottomEnterU, Y: c3.Y - handleLen*tBottomEnterY}

				leftCornerPts = getBezierPoints(c0, c1, c2, c3, "Left Corner")
			}

			// Right corner (connects Bottom -> Top at iEnd)
			{
				i0 := iEnd
				i1 := iEnd - 1
				if i1 < 0 {
					i1 = 0
				}

				pBottom0 := point2D{U: targetAngles[i0] * globalR, Y: yBaseList[i0]}
				pBottom1 := point2D{U: targetAngles[i1] * globalR, Y: yBaseList[i1]}
				tBottomExitU, tBottomExitY := normalize2D(pBottom0.U-pBottom1.U, pBottom0.Y-pBottom1.Y, 1.0, 0.0)

				pTop0 := point2D{U: targetAngles[i0] * globalR, Y: yRotList[i0]}
				pTop1 := point2D{U: targetAngles[i1] * globalR, Y: yRotList[i1]}
				tTopEnterU, tTopEnterY := normalize2D(pTop1.U-pTop0.U, pTop1.Y-pTop0.Y, -1.0, 0.0)

				c0 := pBottom0
				c3 := pTop0
				chord := math.Abs(c3.Y - c0.Y)
				handleLen := controlStrength * chord

				c1 := point2D{U: c0.U + handleLen*tBottomExitU, Y: c0.Y + handleLen*tBottomExitY}
				c2 := point2D{U: c3.U - handleLen*tTopEnterU, Y: c3.Y - handleLen*tTopEnterY}

				rightCornerPts = getBezierPoints(c0, c1, c2, c3, "Right Corner")
			}
		}

		// 14. Add 3D Eye Sampled Points visualization for first repetition if toggled on
		if checkedDiagram != nil && checkedDiagram.StoolDiagram != nil && !checkedDiagram.StoolDiagram.IsHiddenEyeSampledPoints3DShape {
			var eyePoints []*threejs.Vector3
			for i, inE := range inEye {
				if inE {
					alpha := targetAngles[i]
					x := globalR * math.Cos(alpha)
					z := globalR * math.Sin(alpha)

					eyePoints = append(eyePoints, (&threejs.Vector3{
						Name: fmt.Sprintf("Eye Base Point %.1f", alpha*180.0/math.Pi),
						X:    x,
						Y:    yBaseList[i],
						Z:    z,
					}).Stage(stool3dStage))

					eyePoints = append(eyePoints, (&threejs.Vector3{
						Name: fmt.Sprintf("Eye Rotated Point %.1f", alpha*180.0/math.Pi),
						X:    x,
						Y:    yRotList[i],
						Z:    z,
					}).Stage(stool3dStage))
				}
			}
			cylinderstage3d.AddPointSpheres(stool3dStage, eyePoints, "magenta", canvas, "Stool Eye Sampled", 0, 0)
		}

		// 15. Render 3D Eye Corners Sampled Points
		if checkedDiagram != nil && checkedDiagram.StoolDiagram != nil && !checkedDiagram.StoolDiagram.IsHiddenEyeCornersSampledPoints3DShape {
			var cornerPoints []*threejs.Vector3
			cornerPoints = append(cornerPoints, leftCornerPts...)
			cornerPoints = append(cornerPoints, rightCornerPts...)
			cylinderstage3d.AddPointSpheres(stool3dStage, cornerPoints, "cyan", canvas, "Stool Eye Corners Sampled", 0, 0)
		}

		var eye3DPoints []*threejs.Vector3
		if iStart != -1 && iEnd != -1 {
			// A. Bottom curve: from iStart to iEnd
			for i := iStart; i <= iEnd; i++ {
				alpha := targetAngles[i]
				x := globalR * math.Cos(alpha)
				z := globalR * math.Sin(alpha)
				eye3DPoints = append(eye3DPoints, (&threejs.Vector3{
					Name: fmt.Sprintf("Eye Loop Bottom %.1f", alpha*180.0/math.Pi),
					X:    x,
					Y:    yBaseList[i],
					Z:    z,
				}).Stage(stool3dStage))
			}

			// B. Right corner Bézier: Bottom -> Top
			eye3DPoints = append(eye3DPoints, rightCornerPts...)

			// C. Top curve: from iEnd down to iStart (reversed)
			for i := iEnd; i >= iStart; i-- {
				alpha := targetAngles[i]
				x := globalR * math.Cos(alpha)
				z := globalR * math.Sin(alpha)
				eye3DPoints = append(eye3DPoints, (&threejs.Vector3{
					Name: fmt.Sprintf("Eye Loop Top %.1f", alpha*180.0/math.Pi),
					X:    x,
					Y:    yRotList[i],
					Z:    z,
				}).Stage(stool3dStage))
			}

			// D. Left corner Bézier: Top -> Bottom
			eye3DPoints = append(eye3DPoints, leftCornerPts...)
		}

		// 16. Add 3D Eye (Continuous interpolated closed tube)
		if checkedDiagram == nil || checkedDiagram.StoolDiagram == nil || !checkedDiagram.StoolDiagram.IsHiddenEye3DShape {
			if len(eye3DPoints) > 0 {
				eyeLoopCurve := (&threejs.Curve{
					Name:   "Stool Eye Continuous Loop Curve",
					Points: eye3DPoints,
				}).Stage(stool3dStage)

				numSegments := len(eyeLoopCurve.Points)
				if numSegments < 2 {
					numSegments = 2
				}

				eyeGeom := (&threejs.TubeGeometry{
					Name:            "Stool Eye TubeGeom",
					Path:            eyeLoopCurve,
					TubularSegments: numSegments,
					Radius:          tubeRadius,
					RadialSegments:  16,
					Closed:          true,
				}).Stage(stool3dStage)

				eyeMesh := (&threejs.Mesh{
					Name:         "Stool Eye Mesh",
					Position:     threejs.Position{X: 0, Y: 0, Z: 0},
					TubeGeometry: eyeGeom,
					MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
						Name:                 "Stool Eye Material",
						MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "deeppink"},
						Transparent:          true,
						Opacity:              opacity,
					}).Stage(stool3dStage),
				}).Stage(stool3dStage)

				canvas.Meshs = append(canvas.Meshs, eyeMesh)
			}
		}

		var projSeatBottomEyePoints []*threejs.Vector3
		var projStoolBottomEyePoints []*threejs.Vector3

		if len(eye3DPoints) > 0 {
			stoolBottomHeight := 0.0
			for _, pt := range eye3DPoints {
				origTheta := math.Atan2(pt.Z, pt.X)
				r := math.Hypot(pt.X, pt.Z)

				// Seat Bottom projection
				deltaYSB := seatBottomHeight - pt.Y
				rProjSB := r + deltaYSB*math.Tan(projAngleRad)
				projSeatBottomEyePoints = append(projSeatBottomEyePoints, (&threejs.Vector3{
					Name: fmt.Sprintf("Seat Bottom Eye Point %.1f", origTheta*180.0/math.Pi),
					X:    rProjSB * math.Cos(origTheta),
					Y:    seatBottomHeight,
					Z:    rProjSB * math.Sin(origTheta),
				}).Stage(stool3dStage))

				// Stool Bottom projection
				deltaYStool := stoolBottomHeight - pt.Y
				rProjStool := r + deltaYStool*math.Tan(projAngleRad)
				projStoolBottomEyePoints = append(projStoolBottomEyePoints, (&threejs.Vector3{
					Name: fmt.Sprintf("Stool Bottom Eye Point %.1f", origTheta*180.0/math.Pi),
					X:    rProjStool * math.Cos(origTheta),
					Y:    stoolBottomHeight,
					Z:    rProjStool * math.Sin(origTheta),
				}).Stage(stool3dStage))
			}
		}

		// 17. Seat Bottom Eye 2D Projected Curve on horizontal seat bottom plane
		if checkedDiagram != nil && checkedDiagram.StoolDiagram != nil && !checkedDiagram.StoolDiagram.IsHiddenEyeSeatBottomCurveShape {
			if len(projSeatBottomEyePoints) > 0 {
				numSegments := len(projSeatBottomEyePoints)
				if numSegments < 2 {
					numSegments = 2
				}

				for k := 0; k < radialRepetitions; k++ {
					rotAngle := float64(k) * 2.0 * math.Pi / float64(radialRepetitions)
					cosRot := math.Cos(rotAngle)
					sinRot := math.Sin(rotAngle)

					projSeatBottomCurve := (&threejs.Curve{
						Name: fmt.Sprintf("Stool Seat Bottom Eye Curve k%d", k),
					}).Stage(stool3dStage)

					for _, pt := range projSeatBottomEyePoints {
						rx := pt.X*cosRot - pt.Z*sinRot
						rz := pt.X*sinRot + pt.Z*cosRot
						projSeatBottomCurve.Points = append(projSeatBottomCurve.Points, (&threejs.Vector3{
							Name: fmt.Sprintf("Seat Bottom Eye Point k%d", k),
							X:    rx,
							Y:    pt.Y,
							Z:    rz,
						}).Stage(stool3dStage))
					}

					sbEyeGeom := (&threejs.TubeGeometry{
						Name:            fmt.Sprintf("Stool Seat Bottom Eye TubeGeom k%d", k),
						Path:            projSeatBottomCurve,
						TubularSegments: numSegments,
						Radius:          tubeRadius,
						RadialSegments:  16,
						Closed:          true,
					}).Stage(stool3dStage)

					sbEyeMesh := (&threejs.Mesh{
						Name:         fmt.Sprintf("Stool Seat Bottom Eye Mesh k%d", k),
						Position:     threejs.Position{X: 0, Y: 0, Z: 0},
						TubeGeometry: sbEyeGeom,
						MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
							Name:                 fmt.Sprintf("Stool Seat Bottom Eye Material k%d", k),
							MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "mediumvioletred"},
							Transparent:          true,
							Opacity:              opacity,
						}).Stage(stool3dStage),
					}).Stage(stool3dStage)

					canvas.Meshs = append(canvas.Meshs, sbEyeMesh)
				}
			}
		}

		// 18. Stool Bottom Eye 2D Projected Curve on horizontal stool bottom / floor plane (Y = 0)
		if checkedDiagram != nil && checkedDiagram.StoolDiagram != nil && !checkedDiagram.StoolDiagram.IsHiddenEyeStoolBottomCurveShape {
			if len(projStoolBottomEyePoints) > 0 {
				numSegments := len(projStoolBottomEyePoints)
				if numSegments < 2 {
					numSegments = 2
				}

				for k := 0; k < radialRepetitions; k++ {
					rotAngle := float64(k) * 2.0 * math.Pi / float64(radialRepetitions)
					cosRot := math.Cos(rotAngle)
					sinRot := math.Sin(rotAngle)

					projStoolBottomCurve := (&threejs.Curve{
						Name: fmt.Sprintf("Stool Bottom Eye Curve k%d", k),
					}).Stage(stool3dStage)

					for _, pt := range projStoolBottomEyePoints {
						rx := pt.X*cosRot - pt.Z*sinRot
						rz := pt.X*sinRot + pt.Z*cosRot
						projStoolBottomCurve.Points = append(projStoolBottomCurve.Points, (&threejs.Vector3{
							Name: fmt.Sprintf("Stool Bottom Eye Point k%d", k),
							X:    rx,
							Y:    pt.Y,
							Z:    rz,
						}).Stage(stool3dStage))
					}

					stoolBottomEyeGeom := (&threejs.TubeGeometry{
						Name:            fmt.Sprintf("Stool Bottom Eye TubeGeom k%d", k),
						Path:            projStoolBottomCurve,
						TubularSegments: numSegments,
						Radius:          tubeRadius,
						RadialSegments:  16,
						Closed:          true,
					}).Stage(stool3dStage)

					stoolBottomEyeMesh := (&threejs.Mesh{
						Name:         fmt.Sprintf("Stool Bottom Eye Mesh k%d", k),
						Position:     threejs.Position{X: 0, Y: 0, Z: 0},
						TubeGeometry: stoolBottomEyeGeom,
						MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
							Name:                 fmt.Sprintf("Stool Bottom Eye Material k%d", k),
							MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "darkviolet"},
							Transparent:          true,
							Opacity:              opacity,
						}).Stage(stool3dStage),
					}).Stage(stool3dStage)

					canvas.Meshs = append(canvas.Meshs, stoolBottomEyeMesh)
				}
			}
		}

		// 19. 3D Eye Volume Mesh between Seat Bottom Eye Curve and Stool Bottom Eye Curve (repeated across radialRepetitions)
		if checkedDiagram == nil || checkedDiagram.StoolDiagram == nil || !checkedDiagram.StoolDiagram.IsHiddenEyeVolume3DShape {
			if len(projSeatBottomEyePoints) >= 3 && len(projSeatBottomEyePoints) == len(projStoolBottomEyePoints) {
				M := len(projSeatBottomEyePoints)

				for k := 0; k < radialRepetitions; k++ {
					rotAngle := float64(k) * 2.0 * math.Pi / float64(radialRepetitions)
					cosRot := math.Cos(rotAngle)
					sinRot := math.Sin(rotAngle)

					eyeVolGeom := (&threejs.BufferGeometry{
						Name: fmt.Sprintf("Stool Eye Volume BufferGeometry k%d", k),
					}).Stage(stool3dStage)

					var sumTopX, sumTopZ, sumBottomX, sumBottomZ float64
					for i := 0; i < M; i++ {
						origTop := projSeatBottomEyePoints[i]
						rx := origTop.X*cosRot - origTop.Z*sinRot
						rz := origTop.X*sinRot + origTop.Z*cosRot
						topV := (&threejs.Vector3{
							Name: fmt.Sprintf("Eye Top V k%d %d", k, i),
							X:    rx,
							Y:    origTop.Y,
							Z:    rz,
						}).Stage(stool3dStage)
						eyeVolGeom.Vertices = append(eyeVolGeom.Vertices, topV)
						sumTopX += topV.X
						sumTopZ += topV.Z
					}

					for i := 0; i < M; i++ {
						origBot := projStoolBottomEyePoints[i]
						rx := origBot.X*cosRot - origBot.Z*sinRot
						rz := origBot.X*sinRot + origBot.Z*cosRot
						botV := (&threejs.Vector3{
							Name: fmt.Sprintf("Eye Bottom V k%d %d", k, i),
							X:    rx,
							Y:    origBot.Y,
							Z:    rz,
						}).Stage(stool3dStage)
						eyeVolGeom.Vertices = append(eyeVolGeom.Vertices, botV)
						sumBottomX += botV.X
						sumBottomZ += botV.Z
					}

					topCenterIdx := len(eyeVolGeom.Vertices)
					topCenterV := (&threejs.Vector3{
						Name: fmt.Sprintf("Eye Top Center k%d", k),
						X:    sumTopX / float64(M),
						Y:    seatBottomHeight,
						Z:    sumTopZ / float64(M),
					}).Stage(stool3dStage)
					eyeVolGeom.Vertices = append(eyeVolGeom.Vertices, topCenterV)

					botCenterIdx := len(eyeVolGeom.Vertices)
					botCenterV := (&threejs.Vector3{
						Name: fmt.Sprintf("Eye Bottom Center k%d", k),
						X:    sumBottomX / float64(M),
						Y:    0.0,
						Z:    sumBottomZ / float64(M),
					}).Stage(stool3dStage)
					eyeVolGeom.Vertices = append(eyeVolGeom.Vertices, botCenterV)

					// 1. Top face (facing +Y): (topCenter, nextI, i)
					for i := 0; i < M; i++ {
						nextI := (i + 1) % M
						eyeVolGeom.Faces = append(eyeVolGeom.Faces, (&threejs.Triangle{
							Name: fmt.Sprintf("Eye Top Face k%d %d", k, i),
							V1:   topCenterIdx,
							V2:   nextI,
							V3:   i,
						}).Stage(stool3dStage))
					}

					// 2. Bottom face (facing -Y): (botCenter, botI, botNextI)
					for i := 0; i < M; i++ {
						nextI := (i + 1) % M
						botI := M + i
						botNextI := M + nextI
						eyeVolGeom.Faces = append(eyeVolGeom.Faces, (&threejs.Triangle{
							Name: fmt.Sprintf("Eye Bottom Face k%d %d", k, i),
							V1:   botCenterIdx,
							V2:   botI,
							V3:   botNextI,
						}).Stage(stool3dStage))
					}

					// 3. Side wall quads between Top and Bottom:
					for i := 0; i < M; i++ {
						nextI := (i + 1) % M
						topI := i
						topNextI := nextI
						botI := M + i
						botNextI := M + nextI

						// Triangle 1: (botI, topI, topNextI)
						eyeVolGeom.Faces = append(eyeVolGeom.Faces, (&threejs.Triangle{
							Name: fmt.Sprintf("Eye Wall T1 k%d %d", k, i),
							V1:   botI,
							V2:   topI,
							V3:   topNextI,
						}).Stage(stool3dStage))

						// Triangle 2: (botI, topNextI, botNextI)
						eyeVolGeom.Faces = append(eyeVolGeom.Faces, (&threejs.Triangle{
							Name: fmt.Sprintf("Eye Wall T2 k%d %d", k, i),
							V1:   botI,
							V2:   topNextI,
							V3:   botNextI,
						}).Stage(stool3dStage))
					}

					eyeVolMesh := (&threejs.Mesh{
						Name:           fmt.Sprintf("Stool Eye Volume 3D Mesh k%d", k),
						Position:       threejs.Position{X: 0, Y: 0, Z: 0},
						BufferGeometry: eyeVolGeom,
						MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
							Name:                 fmt.Sprintf("Stool Eye Volume Material k%d", k),
							MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "mediumorchid"},
							Transparent:          true,
							Opacity:              opacity,
						}).Stage(stool3dStage),
					}).Stage(stool3dStage)

					canvas.Meshs = append(canvas.Meshs, eyeVolMesh)
				}
			}
		}

		buildSeatAndLegsGeom := func(geomName string, deltaTheta float64, deltaY float64) *threejs.BufferGeometry {
			geom := (&threejs.BufferGeometry{
				Name: geomName,
			}).Stage(stool3dStage)

			cosOffset := math.Cos(deltaTheta)
			sinOffset := math.Sin(deltaTheta)

			transformPoint := func(pt *threejs.Vector3, addY float64) (float64, float64, float64) {
				tx := pt.X*cosOffset - pt.Z*sinOffset
				tz := pt.X*sinOffset + pt.Z*cosOffset
				ty := pt.Y + addY
				return tx, ty, tz
			}

			// --- A. Seat Volume ---
			if len(rotSeatTopPoints) >= 3 && len(rotSeatTopPoints) == len(rotSeatBottomPoints) {
				N := len(rotSeatTopPoints)
				seatBaseIdx := len(geom.Vertices)

				var sumTopX, sumTopZ, sumBottomX, sumBottomZ float64
				for i := 0; i < N; i++ {
					tx, ty, tz := transformPoint(rotSeatTopPoints[i], deltaY)
					topV := (&threejs.Vector3{
						Name: fmt.Sprintf("%s Seat Top V %d", geomName, i),
						X:    tx,
						Y:    ty,
						Z:    tz,
					}).Stage(stool3dStage)
					geom.Vertices = append(geom.Vertices, topV)
					sumTopX += tx
					sumTopZ += tz
				}

				for i := 0; i < N; i++ {
					tx, ty, tz := transformPoint(rotSeatBottomPoints[i], deltaY)
					botV := (&threejs.Vector3{
						Name: fmt.Sprintf("%s Seat Bottom V %d", geomName, i),
						X:    tx,
						Y:    ty,
						Z:    tz,
					}).Stage(stool3dStage)
					geom.Vertices = append(geom.Vertices, botV)
					sumBottomX += tx
					sumBottomZ += tz
				}

				topCenterIdx := len(geom.Vertices)
				topCenterV := (&threejs.Vector3{
					Name: fmt.Sprintf("%s Seat Top Center", geomName),
					X:    sumTopX / float64(N),
					Y:    stoolTopHeight + deltaY,
					Z:    sumTopZ / float64(N),
				}).Stage(stool3dStage)
				geom.Vertices = append(geom.Vertices, topCenterV)

				botCenterIdx := len(geom.Vertices)
				botCenterV := (&threejs.Vector3{
					Name: fmt.Sprintf("%s Seat Bottom Center", geomName),
					X:    sumBottomX / float64(N),
					Y:    seatBottomHeight + deltaY,
					Z:    sumBottomZ / float64(N),
				}).Stage(stool3dStage)
				geom.Vertices = append(geom.Vertices, botCenterV)

				// 1. Top face (facing +Y): (topCenter, nextI, i)
				for i := 0; i < N; i++ {
					nextI := (i + 1) % N
					geom.Faces = append(geom.Faces, (&threejs.Triangle{
						Name: fmt.Sprintf("%s Seat Top Face %d", geomName, i),
						V1:   topCenterIdx,
						V2:   seatBaseIdx + nextI,
						V3:   seatBaseIdx + i,
					}).Stage(stool3dStage))
				}

				// 2. Bottom face (facing -Y): (botCenter, botI, botNextI)
				for i := 0; i < N; i++ {
					nextI := (i + 1) % N
					botI := seatBaseIdx + N + i
					botNextI := seatBaseIdx + N + nextI
					geom.Faces = append(geom.Faces, (&threejs.Triangle{
						Name: fmt.Sprintf("%s Seat Bottom Face %d", geomName, i),
						V1:   botCenterIdx,
						V2:   botI,
						V3:   botNextI,
					}).Stage(stool3dStage))
				}

				// 3. Side wall quads between Top and Bottom:
				for i := 0; i < N; i++ {
					nextI := (i + 1) % N
					topI := seatBaseIdx + i
					topNextI := seatBaseIdx + nextI
					botI := seatBaseIdx + N + i
					botNextI := seatBaseIdx + N + nextI

					geom.Faces = append(geom.Faces, (&threejs.Triangle{
						Name: fmt.Sprintf("%s Seat Wall T1 %d", geomName, i),
						V1:   botI,
						V2:   topI,
						V3:   topNextI,
					}).Stage(stool3dStage))

					geom.Faces = append(geom.Faces, (&threejs.Triangle{
						Name: fmt.Sprintf("%s Seat Wall T2 %d", geomName, i),
						V1:   botI,
						V2:   topNextI,
						V3:   botNextI,
					}).Stage(stool3dStage))
				}
			}

			// --- B. Legs (Eye Volumes across radialRepetitions) ---
			if len(projSeatBottomEyePoints) >= 3 && len(projSeatBottomEyePoints) == len(projStoolBottomEyePoints) {
				M := len(projSeatBottomEyePoints)

				for k := 0; k < radialRepetitions; k++ {
					baseRot := float64(k) * 2.0 * math.Pi / float64(radialRepetitions)
					totalRot := baseRot + deltaTheta
					cosK := math.Cos(totalRot)
					sinK := math.Sin(totalRot)

					legBaseIdx := len(geom.Vertices)

					var sumTopX, sumTopZ, sumBottomX, sumBottomZ float64
					for i := 0; i < M; i++ {
						origTop := projSeatBottomEyePoints[i]
						rx := origTop.X*cosK - origTop.Z*sinK
						rz := origTop.X*sinK + origTop.Z*cosK
						topV := (&threejs.Vector3{
							Name: fmt.Sprintf("%s Leg Top V k%d %d", geomName, k, i),
							X:    rx,
							Y:    origTop.Y + deltaY,
							Z:    rz,
						}).Stage(stool3dStage)
						geom.Vertices = append(geom.Vertices, topV)
						sumTopX += rx
						sumTopZ += rz
					}

					for i := 0; i < M; i++ {
						origBot := projStoolBottomEyePoints[i]
						rx := origBot.X*cosK - origBot.Z*sinK
						rz := origBot.X*sinK + origBot.Z*cosK
						botV := (&threejs.Vector3{
							Name: fmt.Sprintf("%s Leg Bottom V k%d %d", geomName, k, i),
							X:    rx,
							Y:    origBot.Y + deltaY,
							Z:    rz,
						}).Stage(stool3dStage)
						geom.Vertices = append(geom.Vertices, botV)
						sumBottomX += rx
						sumBottomZ += rz
					}

					topCenterIdx := len(geom.Vertices)
					topCenterV := (&threejs.Vector3{
						Name: fmt.Sprintf("%s Leg Top Center k%d", geomName, k),
						X:    sumTopX / float64(M),
						Y:    seatBottomHeight + deltaY,
						Z:    sumTopZ / float64(M),
					}).Stage(stool3dStage)
					geom.Vertices = append(geom.Vertices, topCenterV)

					botCenterIdx := len(geom.Vertices)
					botCenterV := (&threejs.Vector3{
						Name: fmt.Sprintf("%s Leg Bottom Center k%d", geomName, k),
						X:    sumBottomX / float64(M),
						Y:    0.0 + deltaY,
						Z:    sumBottomZ / float64(M),
					}).Stage(stool3dStage)
					geom.Vertices = append(geom.Vertices, botCenterV)

					// 1. Top face (facing +Y): (topCenter, nextI, i)
					for i := 0; i < M; i++ {
						nextI := (i + 1) % M
						geom.Faces = append(geom.Faces, (&threejs.Triangle{
							Name: fmt.Sprintf("%s Leg Top Face k%d %d", geomName, k, i),
							V1:   topCenterIdx,
							V2:   legBaseIdx + nextI,
							V3:   legBaseIdx + i,
						}).Stage(stool3dStage))
					}

					// 2. Bottom face (facing -Y): (botCenter, botI, botNextI)
					for i := 0; i < M; i++ {
						nextI := (i + 1) % M
						botI := legBaseIdx + M + i
						botNextI := legBaseIdx + M + nextI
						geom.Faces = append(geom.Faces, (&threejs.Triangle{
							Name: fmt.Sprintf("%s Leg Bottom Face k%d %d", geomName, k, i),
							V1:   botCenterIdx,
							V2:   botI,
							V3:   botNextI,
						}).Stage(stool3dStage))
					}

					// 3. Side wall quads between Top and Bottom:
					for i := 0; i < M; i++ {
						nextI := (i + 1) % M
						topI := legBaseIdx + i
						topNextI := legBaseIdx + nextI
						botI := legBaseIdx + M + i
						botNextI := legBaseIdx + M + nextI

						geom.Faces = append(geom.Faces, (&threejs.Triangle{
							Name: fmt.Sprintf("%s Leg Wall T1 k%d %d", geomName, k, i),
							V1:   botI,
							V2:   topI,
							V3:   topNextI,
						}).Stage(stool3dStage))

						geom.Faces = append(geom.Faces, (&threejs.Triangle{
							Name: fmt.Sprintf("%s Leg Wall T2 k%d %d", geomName, k, i),
							V1:   botI,
							V2:   topNextI,
							V3:   botNextI,
						}).Stage(stool3dStage))
					}
				}
			}

			return geom
		}

		// 20. 3D Seat and Legs (Union of seat and all legs)
		if checkedDiagram == nil || checkedDiagram.StoolDiagram == nil || !checkedDiagram.StoolDiagram.IsHiddenSeatAndLegs3DShape {
			seatAndLegsGeom := buildSeatAndLegsGeom("Stool Seat and Legs", 0.0, 0.0)
			if len(seatAndLegsGeom.Faces) > 0 {
				seatAndLegsMesh := (&threejs.Mesh{
					Name:           "Stool Seat and Legs Mesh",
					Position:       threejs.Position{X: 0, Y: 0, Z: 0},
					BufferGeometry: seatAndLegsGeom,
					MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
						Name:                 "Stool Seat and Legs Material",
						MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "saddlebrown"},
						Transparent:          true,
						Opacity:              opacity,
					}).Stage(stool3dStage),
				}).Stage(stool3dStage)

				canvas.Meshs = append(canvas.Meshs, seatAndLegsMesh)
			}
		}

		// 21. 3D Rotated Seat and Legs (Transformed by growthVector)
		if checkedDiagram != nil && checkedDiagram.StoolDiagram != nil && !checkedDiagram.StoolDiagram.IsHiddenRotatedSeatAndLegs3DShape {
			rotSeatAndLegsGeom := buildSeatAndLegsGeom("Stool Rotated Seat and Legs", growthVectorX/globalR, growthVectorY)
			if len(rotSeatAndLegsGeom.Faces) > 0 {
				rotSeatAndLegsMesh := (&threejs.Mesh{
					Name:           "Stool Rotated Seat and Legs Mesh",
					Position:       threejs.Position{X: 0, Y: 0, Z: 0},
					BufferGeometry: rotSeatAndLegsGeom,
					MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
						Name:                 "Stool Rotated Seat and Legs Material",
						MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "darkgoldenrod"},
						Transparent:          true,
						Opacity:              opacity,
					}).Stage(stool3dStage),
				}).Stage(stool3dStage)

				canvas.Meshs = append(canvas.Meshs, rotSeatAndLegsMesh)
			}
		}
	}

	if checkedDiagram == nil || checkedDiagram.StoolDiagram == nil || !checkedDiagram.StoolDiagram.IsHiddenTiledFloor3DShape {
		cylinderstage3d.AddFloorTiles(stool3dStage, canvas, globalR, floorMinY)
	}

	stool3dStage.Commit()
}
