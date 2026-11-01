package models

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type glbMeshData struct {
	name      string
	colorHex  string
	positions []vector3
	normals   []vector3
	indices   [][3]uint32
}

func (m *glbMeshData) addQuadStrip(curveA, curveB []vector3) {
	n := min(len(curveA), len(curveB))
	if n < 3 {
		return
	}
	baseIdx := uint32(len(m.positions))
	for i := range n {
		m.positions = append(m.positions, curveA[i], curveB[i])
	}
	for i := range n {
		nextI := (i + 1) % n
		idxA := baseIdx + uint32(2*i)
		idxB := baseIdx + uint32(2*i+1)
		idxANext := baseIdx + uint32(2*nextI)
		idxBNext := baseIdx + uint32(2*nextI+1)

		m.indices = append(m.indices,
			[3]uint32{idxA, idxB, idxBNext},
			[3]uint32{idxA, idxBNext, idxANext},
		)
	}
}

func (m *glbMeshData) computeNormals() {
	m.normals = make([]vector3, len(m.positions))
	for _, tri := range m.indices {
		i0, i1, i2 := tri[0], tri[1], tri[2]
		v0 := m.positions[i0]
		v1 := m.positions[i1]
		v2 := m.positions[i2]
		fn := v1.sub(v0).cross(v2.sub(v0))
		m.normals[i0] = vector3{
			X: m.normals[i0].X + fn.X,
			Y: m.normals[i0].Y + fn.Y,
			Z: m.normals[i0].Z + fn.Z,
		}
		m.normals[i1] = vector3{
			X: m.normals[i1].X + fn.X,
			Y: m.normals[i1].Y + fn.Y,
			Z: m.normals[i1].Z + fn.Z,
		}
		m.normals[i2] = vector3{
			X: m.normals[i2].X + fn.X,
			Y: m.normals[i2].Y + fn.Y,
			Z: m.normals[i2].Z + fn.Z,
		}
	}
	for i := range m.normals {
		n := m.normals[i].normalize()
		if n.X == 0 && n.Y == 0 && n.Z == 0 {
			n = vector3{X: 0, Y: 1, Z: 0}
		}
		m.normals[i] = n
	}
}

func srgbToLinear(c float64) float64 {
	if c <= 0.04045 {
		return c / 12.92
	}
	return math.Pow((c+0.055)/1.055, 2.4)
}

func parseHexColorLinear(hexStr string) (float64, float64, float64) {
	s := strings.TrimPrefix(strings.TrimSpace(hexStr), "#")
	if len(s) != 6 {
		return 0.8, 0.8, 0.8
	}
	r64, errR := strconv.ParseUint(s[0:2], 16, 8)
	g64, errG := strconv.ParseUint(s[2:4], 16, 8)
	b64, errB := strconv.ParseUint(s[4:6], 16, 8)
	if errR != nil || errG != nil || errB != nil {
		return 0.8, 0.8, 0.8
	}
	return srgbToLinear(float64(r64) / 255.0),
		srgbToLinear(float64(g64) / 255.0),
		srgbToLinear(float64(b64) / 255.0)
}

// GenerateVaseTrapezeGLB generates a binary glTF 2.0 (.glb) file for Blender
// containing the full assembled VaseTrapeze (Base_Plate and all stacked/rotated rings
// Ring_0 .. Ring_{StackHeight-1}) with smooth vertex normals and PBR materials
// matching ColorCombination (1..10) and Transparency.
func GenerateVaseTrapezeGLB(plant *PlantAbstract) []byte {
	if plant == nil || plant.PlantType != VaseTrapeze || plant.TubeVaseAbstract == nil {
		return nil
	}
	vase := plant.TubeVaseAbstract
	radialRepetitions := max(vase.RadialRepetitions, 1)
	sideLength := plant.RhombusSideLength
	relativeVerticalThickness := vase.RelativeVerticalThickness
	relativeRotatedTorusSeparation := vase.RelativeRotatedTorusSeparation
	rotRatio := vase.RotationRatio
	stackHeight := max(plant.StackHeight, 1)

	circumference := 10.0
	if plant.RhombusStuff != nil && plant.RhombusStuff.PlantCircumferenceShape != nil && plant.RhombusStuff.PlantCircumferenceShape.Length > 0 {
		circumference = plant.RhombusStuff.PlantCircumferenceShape.Length
	} else if plant.PerpendicularVectorGrid != nil && len(plant.PerpendicularVectorGrid.PerpendicularVectors) > 0 {
		pGrid := plant.PerpendicularVectorGrid
		first := pGrid.PerpendicularVectors[0]
		last := pGrid.PerpendicularVectors[len(pGrid.PerpendicularVectors)-1]
		circumference = last.StartX - first.StartX
	}
	if circumference <= 0 {
		circumference = 10.0
	}
	globalR := circumference * float64(radialRepetitions) / (2 * math.Pi)

	gc := plant.GrowthCurve2D
	tgc := vase.TopGrowthCurve2D
	if gc == nil || tgc == nil ||
		gc.StartHalfwayArcShapeGrid == nil ||
		tgc.TopStartHalfwayArcShapeGrid == nil ||
		len(gc.StartHalfwayArcShapeGrid.StartHalfwayArcShapes) == 0 ||
		len(tgc.TopStartHalfwayArcShapeGrid.TopStartHalfwayArcShapes) == 0 {
		return nil
	}

	startArcs := gc.StartHalfwayArcShapeGrid.StartHalfwayArcShapes
	var endArcs []*EndHalfwayArcShape
	if gc.EndHalfwayArcShapeGrid != nil {
		endArcs = gc.EndHalfwayArcShapeGrid.EndHalfwayArcShapes
	}
	topStartArcs := tgc.TopStartHalfwayArcShapeGrid.TopStartHalfwayArcShapes
	var topEndArcs []*TopEndHalfwayArcShape
	if tgc.TopEndHalfwayArcShapeGrid != nil {
		topEndArcs = tgc.TopEndHalfwayArcShapeGrid.TopEndHalfwayArcShapes
	}

	var curvePoints, topCurvePoints []vector3
	appendArcPoints := func(targetCurve *[]vector3, x1, y1, x2, y2, r float64, sweepFlag bool, largeArcFlag bool) {
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
		const steps = 50
		for i := 0; i <= steps; i++ {
			if i == 0 && len(*targetCurve) > 0 {
				continue
			}
			t := float64(i) / float64(steps)
			angle := startAngle + t*(endAngle-startAngle)
			x2d := cx + r*math.Cos(angle)
			y2d := cy + r*math.Sin(angle)
			theta := x2d / globalR
			*targetCurve = append(*targetCurve, vector3{
				X: globalR * math.Cos(theta),
				Y: y2d,
				Z: globalR * math.Sin(theta),
			})
		}
	}

	for i := range startArcs {
		sa := startArcs[i]
		appendArcPoints(&curvePoints, sa.StartX, sa.StartY, sa.EndX, sa.EndY, sa.RadiusX, !sa.SweepFlag, sa.LargeArcFlag)
		if i < len(endArcs) {
			ea := endArcs[i]
			appendArcPoints(&curvePoints, ea.StartX, ea.StartY, ea.EndX, ea.EndY, ea.RadiusX, !ea.SweepFlag, ea.LargeArcFlag)
		}
	}
	for i := range topStartArcs {
		tsa := topStartArcs[i]
		appendArcPoints(&topCurvePoints, tsa.StartX, tsa.StartY, tsa.EndX, tsa.EndY, tsa.RadiusX, !tsa.SweepFlag, tsa.LargeArcFlag)
		if i < len(topEndArcs) {
			ea := topEndArcs[i]
			appendArcPoints(&topCurvePoints, ea.StartX, ea.StartY, ea.EndX, ea.EndY, ea.RadiusX, !ea.SweepFlag, ea.LargeArcFlag)
		}
	}

	scaleY := vase.RibbonVerticalScale
	if scaleY == 0 {
		scaleY = 1.0
	}
	if len(curvePoints) > 0 {
		y0 := curvePoints[0].Y
		for i := range curvePoints {
			curvePoints[i].Y = vase.Z_Ribbon + scaleY*(curvePoints[i].Y-y0)
		}
		for i := range topCurvePoints {
			topCurvePoints[i].Y = vase.Z_Ribbon + scaleY*(topCurvePoints[i].Y-y0)
		}
	}

	targetAngles, anglesBottom, bottomPoints, anglesTop, topPoints := getTargetAnglesSTL(curvePoints, topCurvePoints, 0.5, radialRepetitions)
	if len(targetAngles) < 4 {
		return nil
	}
	expectedDegrees := 360.0 / float64(radialRepetitions)
	resampledBaseBottom := resampleCurveAtAnglesSTL(anglesBottom, bottomPoints, targetAngles, expectedDegrees)
	resampledBaseTop := resampleCurveAtAnglesSTL(anglesTop, topPoints, targetAngles, expectedDegrees)

	var trapezeUnrotatedStepDY, trapezeRotatedStepDX, trapezeRotatedStepDY float64
	var hasTrapezeConstantThickness bool
	if len(resampledBaseBottom) == len(targetAngles) && len(resampledBaseTop) == len(targetAngles) {
		baseBottomY := make([]float64, len(targetAngles))
		for i, pt := range resampledBaseBottom {
			baseBottomY[i] = pt.Y
		}
		if newBottomY, newTopY, unrotDY, rotDX, rotDY, ok := ComputeVaseTrapezeConstant3DThickness(plant, globalR, targetAngles, baseBottomY); ok {
			for i := range targetAngles {
				resampledBaseBottom[i].Y = newBottomY[i]
				resampledBaseTop[i].Y = newTopY[i]
			}
			trapezeUnrotatedStepDY = unrotDY
			trapezeRotatedStepDX = rotDX
			trapezeRotatedStepDY = rotDY
			hasTrapezeConstantThickness = true
		}
	}

	p1H := vase.Z_Ribbon + vase.TopPlaneHeight*scaleY
	p2H := vase.Z_Ribbon + vase.BottomPlaneHeight*scaleY
	p2HBase := vase.Z_Ribbon + vase.BaseRingBottomPlaneHeight*scaleY
	projAngleRad := -vase.ProjectionAngle * math.Pi / 180.0
	carvedParam := math.Max(0.0, math.Min(1.0, vase.CarvedOutTopRingsParameter))

	projectCurve := func(srcCurve []vector3, planeHeight float64, dy float64) []vector3 {
		var projCurve []vector3
		for _, pt := range srcCurve {
			origTheta := math.Atan2(pt.Z, pt.X)
			r := math.Hypot(pt.X, pt.Z)
			deltaY := (planeHeight - dy) - pt.Y
			rProj := r + deltaY*math.Tan(projAngleRad)
			newPt := vector3{
				X: rProj * math.Cos(origTheta),
				Y: planeHeight,
				Z: rProj * math.Sin(origTheta),
			}
			if len(projCurve) > 0 {
				prev := projCurve[len(projCurve)-1]
				if math.Hypot(newPt.X-prev.X, newPt.Z-prev.Z) < 1e-3 {
					continue
				}
			}
			projCurve = append(projCurve, newPt)
		}
		if len(projCurve) > 2 {
			first := projCurve[0]
			last := projCurve[len(projCurve)-1]
			if math.Hypot(last.X-first.X, last.Z-first.Z) < 1e-3 {
				projCurve = projCurve[:len(projCurve)-1]
			}
		}
		return projCurve
	}

	computeCarvedOutCurve := func(cTop, cBottom []vector3, param float64) []vector3 {
		M := min(len(cTop), len(cBottom))
		if M == 0 {
			return nil
		}
		rMin := math.MaxFloat64
		rMax := -math.MaxFloat64
		for i := range M {
			r := math.Hypot(cTop[i].X, cTop[i].Z)
			if r < rMin {
				rMin = r
			}
			if r > rMax {
				rMax = r
			}
		}
		rRange := rMax - rMin
		carvedCurve := make([]vector3, 0, M)
		for i := range M {
			ptTop := cTop[i]
			ptBottom := cBottom[i]
			r := math.Hypot(ptTop.X, ptTop.Z)
			coeff := 0.0
			if rRange > 1e-6 {
				coeff = param * (rMax - r) / rRange
			}
			carvedCurve = append(carvedCurve, vector3{
				X: (1.0-coeff)*ptTop.X + coeff*ptBottom.X,
				Y: (1.0-coeff)*ptTop.Y + coeff*ptBottom.Y,
				Z: (1.0-coeff)*ptTop.Z + coeff*ptBottom.Z,
			})
		}
		return carvedCurve
	}

	// Compute stack offsets for h = 0 .. stackHeight-1
	dxs := make([]float64, stackHeight)
	dys := make([]float64, stackHeight)
	if stackHeight > 1 {
		var growthVectorX, growthVectorY float64
		if plant.GrowthVectorShape != nil {
			growthVectorX = plant.GrowthVectorShape.X
			growthVectorY = plant.GrowthVectorShape.Y
		}
		var vx, vy float64
		if plant.PerpendicularVectorGrid != nil && len(plant.PerpendicularVectorGrid.PerpendicularVectors) > 0 {
			vFirst := plant.PerpendicularVectorGrid.PerpendicularVectors[0]
			vx = vFirst.EndX - vFirst.StartX
			vy = vFirst.EndY - vFirst.StartY
			if vLen := math.Hypot(vx, vy); vLen > 0 {
				vx, vy = vx/vLen, vy/vLen
			}
		}
		verticalThickness := relativeVerticalThickness * sideLength
		rotatedSeparation := relativeRotatedTorusSeparation * sideLength

		_, unrotatedStepDY, _ := ComputePartiallyGrowthCurveDYForRatio(plant, 0.0)
		unrotatedStepDY *= scaleY
		rotatedStepDX := growthVectorX + verticalThickness*vx
		rotatedStepDY := (growthVectorY + verticalThickness*vy + rotatedSeparation) * scaleY
		if hasTrapezeConstantThickness {
			unrotatedStepDY = trapezeUnrotatedStepDY
			rotatedStepDX = trapezeRotatedStepDX
			rotatedStepDY = trapezeRotatedStepDY
		}

		numSteps := stackHeight - 1
		var cumDX, cumDY float64
		for k := 1; k <= numSteps; k++ {
			stepOrder := numSteps - k + 1
			threshold := float64(stepOrder) / float64(stackHeight)
			if rotRatio+1e-9 >= threshold {
				cumDX += rotatedStepDX
				cumDY += rotatedStepDY
			} else {
				cumDY += unrotatedStepDY
			}
			dxs[k] = cumDX
			dys[k] = cumDY
		}
	}

	var meshes []*glbMeshData

	// 1. Build Base_Plate mesh (connected to h=0 at p2HBase)
	baseBottom0, baseTop0 := buildBaseRingSTL(resampledBaseBottom, resampledBaseTop, radialRepetitions, 0.0)
	cTopP1_0 := projectCurve(baseTop0, p1H, 0)
	cBottomP1_0 := projectCurve(baseBottom0, p1H, 0)
	cTopP2Base := projectCurve(baseTop0, p2HBase, 0)
	cBottomP2Base := projectCurve(baseBottom0, p2HBase, 0)

	const floorY = 0.0
	if p2HBase > floorY && len(cTopP1_0) >= 3 && len(cBottomP1_0) >= 3 && len(cTopP2Base) >= 3 && len(cBottomP2Base) >= 3 {
		M := min(len(cTopP1_0), len(cBottomP1_0), len(cTopP2Base), len(cBottomP2Base))
		var sumRTop, sumRBottom float64
		for i := range M {
			sumRTop += math.Hypot(cTopP2Base[i].X, cTopP2Base[i].Z)
			sumRBottom += math.Hypot(cBottomP2Base[i].X, cBottomP2Base[i].Z)
		}
		cOuter := cTopP2Base[:M]
		cInner := cBottomP2Base[:M]
		cOuterP1 := cTopP1_0[:M]
		if sumRBottom > sumRTop {
			cOuter = cBottomP2Base[:M]
			cInner = cTopP2Base[:M]
			cOuterP1 = cBottomP1_0[:M]
		}

		pBottomH := p2HBase
		hPlancher := math.Max(0.0, math.Min(pBottomH-floorY, vase.BasePlateHeight))
		startMag := vase.BulbousStartTangentMagnitude
		endAngleDeg := 30.0
		if vase.BulbousEndAngle != 0.0 {
			endAngleDeg = vase.BulbousEndAngle
		}
		endMag := vase.BulbousEndTangentMagnitude

		S := 24
		if startMag == 0 && endMag == 0 {
			S = 1
		}

		baseMesh := &glbMeshData{
			name:     "Base_Plate",
			colorHex: GetVaseRingColor(vase.ColorCombination, 0),
		}

		idxOuter := make([][]uint32, S+1)
		idxInner := make([][]uint32, S+1)
		for s := range S + 1 {
			idxOuter[s] = make([]uint32, M)
			idxInner[s] = make([]uint32, M)
		}

		var sumOuterX, sumOuterZ, sumInnerX, sumInnerZ float64
		plancherY := floorY + hPlancher
		verticalScale := (pBottomH - plancherY) / (pBottomH - floorY)

		beta := endAngleDeg * math.Pi / 180.0
		v1R := endMag * (-math.Sin(beta))
		v1Y := endMag * (-math.Cos(beta))

		for i := range M {
			ptOut := cOuter[i]
			ptOutP1 := cOuterP1[i]
			ptIn := cInner[i]

			sumOuterX += ptOut.X
			sumOuterZ += ptOut.Z
			sumInnerX += ptIn.X
			sumInnerZ += ptIn.Z

			r0 := math.Hypot(ptOut.X, ptOut.Z)
			rTop := math.Hypot(ptOutP1.X, ptOutP1.Z)
			theta := math.Atan2(ptOut.Z, ptOut.X)
			cosTheta, sinTheta := math.Cos(theta), math.Sin(theta)

			rIn0 := math.Hypot(ptIn.X, ptIn.Z)
			thetaIn := math.Atan2(ptIn.Z, ptIn.X)
			cosThetaIn, sinThetaIn := math.Cos(thetaIn), math.Sin(thetaIn)
			radialScale := 1.0
			if r0 > 1e-6 {
				radialScale = rIn0 / r0
			}

			deltaR := r0 - rTop
			deltaY := pBottomH - p1H
			lRing := math.Hypot(deltaR, deltaY)
			t0R, t0Y := 0.0, -1.0
			if lRing > 1e-6 {
				t0R, t0Y = deltaR/lRing, deltaY/lRing
			}
			v0R, v0Y := startMag*t0R, startMag*t0Y

			p0R, p0Y := r0, pBottomH
			p1R, p1Y := r0+v0R/3.0, pBottomH+v0Y/3.0
			p2R, p2Y := r0-v1R/3.0, floorY-v1Y/3.0
			p3R, p3Y := r0, floorY

			for s := 0; s <= S; s++ {
				t := float64(s) / float64(S)
				omt := 1.0 - t
				b0 := omt * omt * omt
				b1 := 3.0 * omt * omt * t
				b2 := 3.0 * omt * t * t
				b3 := t * t * t

				rs := b0*p0R + b1*p1R + b2*p2R + b3*p3R
				ys := b0*p0Y + b1*p1Y + b2*p2Y + b3*p3Y
				if s == 0 {
					rs, ys = r0, pBottomH
				} else if s == S {
					rs, ys = r0, floorY
				}

				idxOuter[s][i] = uint32(len(baseMesh.positions))
				baseMesh.positions = append(baseMesh.positions, vector3{
					X: rs * cosTheta,
					Y: ys,
					Z: rs * sinTheta,
				})

				rsIn := rs * radialScale
				ysIn := plancherY + (ys-floorY)*verticalScale
				var inPt vector3
				if s == 0 {
					inPt = vector3{X: ptIn.X, Y: pBottomH, Z: ptIn.Z}
				} else if s == S {
					inPt = vector3{X: ptIn.X, Y: plancherY, Z: ptIn.Z}
				} else {
					inPt = vector3{X: rsIn * cosThetaIn, Y: ysIn, Z: rsIn * sinThetaIn}
				}
				idxInner[s][i] = uint32(len(baseMesh.positions))
				baseMesh.positions = append(baseMesh.positions, inPt)
			}
		}

		for s := range S {
			for i := range M {
				nextI := (i + 1) % M
				// Outer bulbous wall
				baseMesh.indices = append(baseMesh.indices,
					[3]uint32{idxOuter[s][i], idxOuter[s][nextI], idxOuter[s+1][nextI]},
					[3]uint32{idxOuter[s][i], idxOuter[s+1][nextI], idxOuter[s+1][i]},
				)
				// Inner curved wall
				baseMesh.indices = append(baseMesh.indices,
					[3]uint32{idxInner[s][i], idxInner[s+1][nextI], idxInner[s][nextI]},
					[3]uint32{idxInner[s][i], idxInner[s+1][i], idxInner[s+1][nextI]},
				)
			}
		}

		// Dedicated cap vertices so flat caps do not distort smooth wall normals
		idxPlancherCenter := uint32(len(baseMesh.positions))
		baseMesh.positions = append(baseMesh.positions, vector3{
			X: sumInnerX / float64(M), Y: plancherY, Z: sumInnerZ / float64(M),
		})
		idxFloorCenter := uint32(len(baseMesh.positions))
		baseMesh.positions = append(baseMesh.positions, vector3{
			X: sumOuterX / float64(M), Y: floorY, Z: sumOuterZ / float64(M),
		})

		idxCapOuterTop := make([]uint32, M)
		idxCapInnerTop := make([]uint32, M)
		idxCapInnerFloor := make([]uint32, M)
		idxCapOuterFloor := make([]uint32, M)
		for i := range M {
			idxCapOuterTop[i] = uint32(len(baseMesh.positions))
			baseMesh.positions = append(baseMesh.positions, baseMesh.positions[idxOuter[0][i]])

			idxCapInnerTop[i] = uint32(len(baseMesh.positions))
			baseMesh.positions = append(baseMesh.positions, baseMesh.positions[idxInner[0][i]])

			idxCapInnerFloor[i] = uint32(len(baseMesh.positions))
			baseMesh.positions = append(baseMesh.positions, baseMesh.positions[idxInner[S][i]])

			idxCapOuterFloor[i] = uint32(len(baseMesh.positions))
			baseMesh.positions = append(baseMesh.positions, baseMesh.positions[idxOuter[S][i]])
		}

		for i := range M {
			nextI := (i + 1) % M
			baseMesh.indices = append(baseMesh.indices,
				[3]uint32{idxCapInnerTop[i], idxCapOuterTop[nextI], idxCapOuterTop[i]},
				[3]uint32{idxCapInnerTop[i], idxCapInnerTop[nextI], idxCapOuterTop[nextI]},
				[3]uint32{idxPlancherCenter, idxCapInnerFloor[nextI], idxCapInnerFloor[i]},
				[3]uint32{idxFloorCenter, idxCapOuterFloor[i], idxCapOuterFloor[nextI]},
			)
		}

		baseMesh.computeNormals()
		meshes = append(meshes, baseMesh)
	}

	// 2. Build Ring_0 .. Ring_{stackHeight-1}
	for h := range stackHeight {
		dx := dxs[h]
		dy := dys[h]
		thetaOffset := dx / globalR

		curBottom, curTop := buildBaseRingSTL(resampledBaseBottom, resampledBaseTop, radialRepetitions, thetaOffset)
		curP2H := p2H
		if h == 0 {
			curP2H = p2HBase
		}

		ringTopP1 := projectCurve(curTop, p1H+dy, dy)
		ringBottomP1 := projectCurve(curBottom, p1H+dy, dy)
		ringTopP2 := projectCurve(curTop, curP2H+dy, dy)
		ringBottomP2 := projectCurve(curBottom, curP2H+dy, dy)

		ringCarvedTopP1 := computeCarvedOutCurve(ringTopP1, ringTopP2, carvedParam)
		ringCarvedBottomP1 := computeCarvedOutCurve(ringBottomP1, ringBottomP2, carvedParam)

		M := min(len(ringCarvedTopP1), len(ringCarvedBottomP1), len(ringTopP2), len(ringBottomP2))
		if M < 3 {
			continue
		}
		ringCarvedTopP1 = ringCarvedTopP1[:M]
		ringCarvedBottomP1 = ringCarvedBottomP1[:M]
		ringTopP2 = ringTopP2[:M]
		ringBottomP2 = ringBottomP2[:M]

		var sumRTop, sumRBottom float64
		for i := range M {
			sumRTop += math.Hypot(ringTopP2[i].X, ringTopP2[i].Z)
			sumRBottom += math.Hypot(ringBottomP2[i].X, ringBottomP2[i].Z)
		}

		cOuterP1 := ringCarvedTopP1
		cInnerP1 := ringCarvedBottomP1
		cOuterP2 := ringTopP2
		cInnerP2 := ringBottomP2
		if sumRBottom > sumRTop {
			cOuterP1 = ringCarvedBottomP1
			cInnerP1 = ringCarvedTopP1
			cOuterP2 = ringBottomP2
			cInnerP2 = ringTopP2
		}

		ringMesh := &glbMeshData{
			name:     fmt.Sprintf("Ring_%d", h),
			colorHex: GetVaseRingColor(vase.ColorCombination, h),
		}
		// 1. Top carved rim face (normal +Y)
		ringMesh.addQuadStrip(cOuterP1, cInnerP1)
		// 2. Bottom annular face (normal -Y)
		ringMesh.addQuadStrip(cInnerP2, cOuterP2)
		// 3. Outer wall (normal +r)
		ringMesh.addQuadStrip(cOuterP2, cOuterP1)
		// 4. Inner wall (normal -r)
		ringMesh.addQuadStrip(cInnerP1, cInnerP2)

		ringMesh.computeNormals()
		meshes = append(meshes, ringMesh)
	}

	if len(meshes) == 0 {
		return nil
	}

	return packGLB(meshes, vase.Transparency)
}

func packGLB(meshes []*glbMeshData, transparency float64) []byte {
	const mmToMeters = 0.001

	type bufferViewInfo struct {
		byteOffset int
		byteLength int
		target     int
	}
	type accessorInfo struct {
		bufferView    int
		componentType int
		count         int
		accType       string
		hasMinMax     bool
		min           [3]float32
		max           [3]float32
	}

	var binBuf []byte
	var bufferViews []bufferViewInfo
	var accessors []accessorInfo

	appendUint32 := func(v uint32) {
		var b [4]byte
		binary.LittleEndian.PutUint32(b[:], v)
		binBuf = append(binBuf, b[:]...)
	}
	appendFloat32 := func(v float32) {
		var b [4]byte
		binary.LittleEndian.PutUint32(b[:], math.Float32bits(v))
		binBuf = append(binBuf, b[:]...)
	}

	for _, m := range meshes {
		// 1. POSITION bufferView & accessor
		posOffset := len(binBuf)
		minPos := [3]float32{math.MaxFloat32, math.MaxFloat32, math.MaxFloat32}
		maxPos := [3]float32{-math.MaxFloat32, -math.MaxFloat32, -math.MaxFloat32}
		for _, p := range m.positions {
			px := float32(p.X * mmToMeters)
			py := float32(p.Y * mmToMeters)
			pz := float32(p.Z * mmToMeters)
			if px < minPos[0] {
				minPos[0] = px
			}
			if py < minPos[1] {
				minPos[1] = py
			}
			if pz < minPos[2] {
				minPos[2] = pz
			}
			if px > maxPos[0] {
				maxPos[0] = px
			}
			if py > maxPos[1] {
				maxPos[1] = py
			}
			if pz > maxPos[2] {
				maxPos[2] = pz
			}
			appendFloat32(px)
			appendFloat32(py)
			appendFloat32(pz)
		}
		posBV := len(bufferViews)
		bufferViews = append(bufferViews, bufferViewInfo{
			byteOffset: posOffset,
			byteLength: len(binBuf) - posOffset,
			target:     34962, // ARRAY_BUFFER
		})
		accessors = append(accessors, accessorInfo{
			bufferView:    posBV,
			componentType: 5126, // FLOAT
			count:         len(m.positions),
			accType:       "VEC3",
			hasMinMax:     true,
			min:           minPos,
			max:           maxPos,
		})

		// 2. NORMAL bufferView & accessor
		normOffset := len(binBuf)
		for _, n := range m.normals {
			appendFloat32(float32(n.X))
			appendFloat32(float32(n.Y))
			appendFloat32(float32(n.Z))
		}
		normBV := len(bufferViews)
		bufferViews = append(bufferViews, bufferViewInfo{
			byteOffset: normOffset,
			byteLength: len(binBuf) - normOffset,
			target:     34962, // ARRAY_BUFFER
		})
		accessors = append(accessors, accessorInfo{
			bufferView:    normBV,
			componentType: 5126, // FLOAT
			count:         len(m.normals),
			accType:       "VEC3",
		})

		// 3. INDICES bufferView & accessor
		idxOffset := len(binBuf)
		for _, tri := range m.indices {
			appendUint32(tri[0])
			appendUint32(tri[1])
			appendUint32(tri[2])
		}
		idxBV := len(bufferViews)
		bufferViews = append(bufferViews, bufferViewInfo{
			byteOffset: idxOffset,
			byteLength: len(binBuf) - idxOffset,
			target:     34963, // ELEMENT_ARRAY_BUFFER
		})
		accessors = append(accessors, accessorInfo{
			bufferView:    idxBV,
			componentType: 5125, // UNSIGNED_INT
			count:         len(m.indices) * 3,
			accType:       "SCALAR",
		})
	}

	for len(binBuf)%4 != 0 {
		binBuf = append(binBuf, 0)
	}

	var jb strings.Builder
	jb.WriteString(`{"asset":{"version":"2.0","generator":"Phylla Gong GLB Exporter"}`)
	if transparency > 0 {
		jb.WriteString(`,"extensionsUsed":["KHR_materials_transmission","KHR_materials_ior"]`)
	}
	jb.WriteString(`,"scene":0,"scenes":[{"name":"VaseTrapeze","nodes":[`)
	for i := range meshes {
		if i > 0 {
			jb.WriteByte(',')
		}
		jb.WriteString(strconv.Itoa(i))
	}
	jb.WriteString(`]}],"nodes":[`)
	for i, m := range meshes {
		if i > 0 {
			jb.WriteByte(',')
		}
		jb.WriteString(fmt.Sprintf(`{"name":%q,"mesh":%d}`, m.name, i))
	}
	jb.WriteString(`],"meshes":[`)
	for i, m := range meshes {
		if i > 0 {
			jb.WriteByte(',')
		}
		posAcc := 3 * i
		normAcc := 3*i + 1
		idxAcc := 3*i + 2
		jb.WriteString(fmt.Sprintf(
			`{"name":%q,"primitives":[{"attributes":{"POSITION":%d,"NORMAL":%d},"indices":%d,"material":%d}]}`,
			m.name, posAcc, normAcc, idxAcc, i,
		))
	}
	jb.WriteString(`],"materials":[`)
	for i, m := range meshes {
		if i > 0 {
			jb.WriteByte(',')
		}
		rLin, gLin, bLin := parseHexColorLinear(m.colorHex)
		alpha := 1.0 - math.Max(0.0, math.Min(0.95, transparency))
		roughness := 0.32
		if transparency > 0 {
			roughness = 0.06
		}
		jb.WriteString(fmt.Sprintf(
			`{"name":%q,"doubleSided":true,"pbrMetallicRoughness":{"baseColorFactor":[%.6f,%.6f,%.6f,%.4f],"metallicFactor":0.0,"roughnessFactor":%.3f}`,
			m.name+"_Material", rLin, gLin, bLin, alpha, roughness,
		))
		if transparency > 0 {
			jb.WriteString(fmt.Sprintf(
				`,"alphaMode":"BLEND","extensions":{"KHR_materials_transmission":{"transmissionFactor":%.4f},"KHR_materials_ior":{"ior":1.52}}`,
				math.Max(0.0, math.Min(1.0, transparency)),
			))
		}
		jb.WriteByte('}')
	}
	jb.WriteString(`],"accessors":[`)
	for i, a := range accessors {
		if i > 0 {
			jb.WriteByte(',')
		}
		if a.hasMinMax {
			jb.WriteString(fmt.Sprintf(
				`{"bufferView":%d,"componentType":%d,"count":%d,"type":%q,"min":[%.6f,%.6f,%.6f],"max":[%.6f,%.6f,%.6f]}`,
				a.bufferView, a.componentType, a.count, a.accType,
				a.min[0], a.min[1], a.min[2],
				a.max[0], a.max[1], a.max[2],
			))
		} else {
			jb.WriteString(fmt.Sprintf(
				`{"bufferView":%d,"componentType":%d,"count":%d,"type":%q}`,
				a.bufferView, a.componentType, a.count, a.accType,
			))
		}
	}
	jb.WriteString(`],"bufferViews":[`)
	for i, bv := range bufferViews {
		if i > 0 {
			jb.WriteByte(',')
		}
		jb.WriteString(fmt.Sprintf(
			`{"buffer":0,"byteOffset":%d,"byteLength":%d,"target":%d}`,
			bv.byteOffset, bv.byteLength, bv.target,
		))
	}
	jb.WriteString(fmt.Sprintf(`],"buffers":[{"byteLength":%d}]}`, len(binBuf)))

	jsonBytes := []byte(jb.String())
	for len(jsonBytes)%4 != 0 {
		jsonBytes = append(jsonBytes, 0x20) // pad with ASCII space per glTF 2.0 spec
	}

	totalLen := 12 + 8 + len(jsonBytes) + 8 + len(binBuf)
	out := make([]byte, totalLen)

	// 12-byte GLB header
	binary.LittleEndian.PutUint32(out[0:4], 0x46546C67) // "glTF"
	binary.LittleEndian.PutUint32(out[4:8], 2)          // version 2
	binary.LittleEndian.PutUint32(out[8:12], uint32(totalLen))

	// JSON chunk header + payload
	binary.LittleEndian.PutUint32(out[12:16], uint32(len(jsonBytes)))
	binary.LittleEndian.PutUint32(out[16:20], 0x4E4F534A) // "JSON"
	copy(out[20:20+len(jsonBytes)], jsonBytes)

	// BIN chunk header + payload
	binHeaderOffset := 20 + len(jsonBytes)
	binary.LittleEndian.PutUint32(out[binHeaderOffset:binHeaderOffset+4], uint32(len(binBuf)))
	binary.LittleEndian.PutUint32(out[binHeaderOffset+4:binHeaderOffset+8], 0x004E4942) // "BIN\0"
	copy(out[binHeaderOffset+8:], binBuf)

	return out
}

