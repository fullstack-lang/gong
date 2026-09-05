package stoolstage3d

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/fullstack-lang/gong/dsm/phylla/go/cylinderstage3d"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

type StoolGeometriesResult struct {
	SeatGeom     *threejs.BufferGeometry
	LegsGeom     *threejs.BufferGeometry
	CombinedGeom *threejs.BufferGeometry
}

func buildSeatGeom(
	stage *threejs.Stage,
	geomName string,
	deltaTheta float64,
	deltaY float64,
	rotSeatTopPoints []*threejs.Vector3,
	rotSeatBottomPoints []*threejs.Vector3,
	stoolTopHeight float64,
	seatBottomHeight float64,
) *threejs.BufferGeometry {
	geom := (&threejs.BufferGeometry{
		Name: geomName,
	}).Stage(stage)

	if len(rotSeatTopPoints) < 3 || len(rotSeatTopPoints) != len(rotSeatBottomPoints) {
		return geom
	}

	cosOffset := math.Cos(deltaTheta)
	sinOffset := math.Sin(deltaTheta)

	transformPoint := func(pt *threejs.Vector3, addY float64) (float64, float64, float64) {
		tx := pt.X*cosOffset - pt.Z*sinOffset
		tz := pt.X*sinOffset + pt.Z*cosOffset
		ty := pt.Y + addY
		return tx, ty, tz
	}

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
		}).Stage(stage)
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
		}).Stage(stage)
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
	}).Stage(stage)
	geom.Vertices = append(geom.Vertices, topCenterV)

	botCenterIdx := len(geom.Vertices)
	botCenterV := (&threejs.Vector3{
		Name: fmt.Sprintf("%s Seat Bottom Center", geomName),
		X:    sumBottomX / float64(N),
		Y:    seatBottomHeight + deltaY,
		Z:    sumBottomZ / float64(N),
	}).Stage(stage)
	geom.Vertices = append(geom.Vertices, botCenterV)

	// 1. Top face (facing +Y): (topCenter, nextI, i)
	for i := 0; i < N; i++ {
		nextI := (i + 1) % N
		geom.Faces = append(geom.Faces, (&threejs.Triangle{
			Name: fmt.Sprintf("%s Seat Top Face %d", geomName, i),
			V1:   topCenterIdx,
			V2:   seatBaseIdx + nextI,
			V3:   seatBaseIdx + i,
		}).Stage(stage))
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
		}).Stage(stage))
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
		}).Stage(stage))

		geom.Faces = append(geom.Faces, (&threejs.Triangle{
			Name: fmt.Sprintf("%s Seat Wall T2 %d", geomName, i),
			V1:   botI,
			V2:   topNextI,
			V3:   botNextI,
		}).Stage(stage))
	}

	return geom
}

func buildLegsGeom(
	stage *threejs.Stage,
	geomName string,
	deltaTheta float64,
	deltaY float64,
	projSeatBottomEyePoints []*threejs.Vector3,
	projStoolBottomEyePoints []*threejs.Vector3,
	seatBottomHeight float64,
	radialRepetitions int,
) *threejs.BufferGeometry {
	geom := (&threejs.BufferGeometry{
		Name: geomName,
	}).Stage(stage)

	if len(projSeatBottomEyePoints) < 3 || len(projSeatBottomEyePoints) != len(projStoolBottomEyePoints) {
		return geom
	}

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
			}).Stage(stage)
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
			}).Stage(stage)
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
		}).Stage(stage)
		geom.Vertices = append(geom.Vertices, topCenterV)

		botCenterIdx := len(geom.Vertices)
		botCenterV := (&threejs.Vector3{
			Name: fmt.Sprintf("%s Leg Bottom Center k%d", geomName, k),
			X:    sumBottomX / float64(M),
			Y:    0.0 + deltaY,
			Z:    sumBottomZ / float64(M),
		}).Stage(stage)
		geom.Vertices = append(geom.Vertices, botCenterV)

		// 1. Top face (facing +Y): (topCenter, nextI, i)
		for i := 0; i < M; i++ {
			nextI := (i + 1) % M
			geom.Faces = append(geom.Faces, (&threejs.Triangle{
				Name: fmt.Sprintf("%s Leg Top Face k%d %d", geomName, k, i),
				V1:   topCenterIdx,
				V2:   legBaseIdx + nextI,
				V3:   legBaseIdx + i,
			}).Stage(stage))
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
			}).Stage(stage))
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
			}).Stage(stage))

			geom.Faces = append(geom.Faces, (&threejs.Triangle{
				Name: fmt.Sprintf("%s Leg Wall T2 k%d %d", geomName, k, i),
				V1:   botI,
				V2:   topNextI,
				V3:   botNextI,
			}).Stage(stage))
		}
	}

	return geom
}

func buildSeatAndLegsGeom(
	stage *threejs.Stage,
	geomName string,
	deltaTheta float64,
	deltaY float64,
	rotSeatTopPoints []*threejs.Vector3,
	rotSeatBottomPoints []*threejs.Vector3,
	stoolTopHeight float64,
	seatBottomHeight float64,
	projSeatBottomEyePoints []*threejs.Vector3,
	projStoolBottomEyePoints []*threejs.Vector3,
	radialRepetitions int,
) *threejs.BufferGeometry {
	geom := (&threejs.BufferGeometry{
		Name: geomName,
	}).Stage(stage)

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
			}).Stage(stage)
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
			}).Stage(stage)
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
		}).Stage(stage)
		geom.Vertices = append(geom.Vertices, topCenterV)

		botCenterIdx := len(geom.Vertices)
		botCenterV := (&threejs.Vector3{
			Name: fmt.Sprintf("%s Seat Bottom Center", geomName),
			X:    sumBottomX / float64(N),
			Y:    seatBottomHeight + deltaY,
			Z:    sumBottomZ / float64(N),
		}).Stage(stage)
		geom.Vertices = append(geom.Vertices, botCenterV)

		// 1. Top face (facing +Y): (topCenter, nextI, i)
		for i := 0; i < N; i++ {
			nextI := (i + 1) % N
			geom.Faces = append(geom.Faces, (&threejs.Triangle{
				Name: fmt.Sprintf("%s Seat Top Face %d", geomName, i),
				V1:   topCenterIdx,
				V2:   seatBaseIdx + nextI,
				V3:   seatBaseIdx + i,
			}).Stage(stage))
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
			}).Stage(stage))
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
			}).Stage(stage))

			geom.Faces = append(geom.Faces, (&threejs.Triangle{
				Name: fmt.Sprintf("%s Seat Wall T2 %d", geomName, i),
				V1:   botI,
				V2:   topNextI,
				V3:   botNextI,
			}).Stage(stage))
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
				}).Stage(stage)
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
				}).Stage(stage)
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
			}).Stage(stage)
			geom.Vertices = append(geom.Vertices, topCenterV)

			botCenterIdx := len(geom.Vertices)
			botCenterV := (&threejs.Vector3{
				Name: fmt.Sprintf("%s Leg Bottom Center k%d", geomName, k),
				X:    sumBottomX / float64(M),
				Y:    0.0 + deltaY,
				Z:    sumBottomZ / float64(M),
			}).Stage(stage)
			geom.Vertices = append(geom.Vertices, botCenterV)

			// 1. Top face (facing +Y): (topCenter, nextI, i)
			for i := 0; i < M; i++ {
				nextI := (i + 1) % M
				geom.Faces = append(geom.Faces, (&threejs.Triangle{
					Name: fmt.Sprintf("%s Leg Top Face k%d %d", geomName, k, i),
					V1:   topCenterIdx,
					V2:   legBaseIdx + nextI,
					V3:   legBaseIdx + i,
				}).Stage(stage))
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
				}).Stage(stage))
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
				}).Stage(stage))

				geom.Faces = append(geom.Faces, (&threejs.Triangle{
					Name: fmt.Sprintf("%s Leg Wall T2 k%d %d", geomName, k, i),
					V1:   botI,
					V2:   topNextI,
					V3:   botNextI,
				}).Stage(stage))
			}
		}
	}

	return geom
}

func ComputeStoolGeometries(
	stage *threejs.Stage,
	stager *models.Stager,
	plant *models.PlantAbstract,
) *StoolGeometriesResult {
	if plant == nil || plant.PlantType != models.Stool || plant.StoolAbstract == nil {
		return nil
	}

	if stage == nil {
		stage = threejs.NewStage("stool_computation")
	}

	params := cylinderstage3d.Cylinder3DParams{
		NamePrefix:            "Stool",
		CanvasName:            "Stool 3D Canvas",
		RadialRepetitions:     plant.StoolAbstract.RadialRepetitions,
		Transparency:          plant.StoolAbstract.Transparency,
		RelativeTubeDiameter:  plant.StoolAbstract.RelativeTubeDiameter,
		RelativeHeight3DTorus: plant.StoolAbstract.RelativeHeight3DTorus,
		VerticalScale:         plant.StoolAbstract.StoolTorusVerticalScale,
		RelativeHeight:        plant.StoolAbstract.RelativeHeight,
		ProjectionAngle:       plant.StoolAbstract.ProjectionAngle,
		HasRotatedShapes:      true,
	}

	base := cylinderstage3d.RenderCylinder3DBase(stage, stager, plant, params)
	if base == nil || base.ResampledBaseCurve == nil || plant.StartArcShapeGrid == nil || len(plant.StartArcShapeGrid.StartArcShapes) == 0 {
		return nil
	}

	resampledBaseCurve := base.ResampledBaseCurve
	radialRepetitions := base.RadialRepetitions
	globalR := base.GlobalR
	stoolTopHeight := base.TopHeight
	torusHeight := base.TorusHeight
	growthVectorX := base.GrowthVectorX
	growthVectorY := base.GrowthVectorY
	projAngleRad := base.ProjAngleRad
	vertScale := base.VertScale
	expectedDegrees := base.ExpectedDegrees
	targetAngles := base.TargetAngles
	rotSeatTopPoints := base.RotTopPoints

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
			}).Stage(stage))
		}
	}

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
			}).Stage(stage))
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
			}).Stage(stage))
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
			}).Stage(stage))
		}

		// D. Left corner Bézier: Top -> Bottom
		eye3DPoints = append(eye3DPoints, leftCornerPts...)
	}

	var projSeatBottomEyePoints []*threejs.Vector3
	var projStoolBottomEyePoints []*threejs.Vector3

	if len(eye3DPoints) > 0 {
		stoolBottomHeight := 0.0
		for _, pt := range eye3DPoints {
			origTheta := math.Atan2(pt.Z, pt.X)
			r := math.Hypot(pt.X, pt.Z)

			deltaYSB := seatBottomHeight - pt.Y
			rProjSB := r + deltaYSB*math.Tan(projAngleRad)
			projSeatBottomEyePoints = append(projSeatBottomEyePoints, (&threejs.Vector3{
				Name: fmt.Sprintf("Seat Bottom Eye Point %.1f", origTheta*180.0/math.Pi),
				X:    rProjSB * math.Cos(origTheta),
				Y:    seatBottomHeight,
				Z:    rProjSB * math.Sin(origTheta),
			}).Stage(stage))

			deltaYStool := stoolBottomHeight - pt.Y
			rProjStool := r + deltaYStool*math.Tan(projAngleRad)
			projStoolBottomEyePoints = append(projStoolBottomEyePoints, (&threejs.Vector3{
				Name: fmt.Sprintf("Stool Bottom Eye Point %.1f", origTheta*180.0/math.Pi),
				X:    rProjStool * math.Cos(origTheta),
				Y:    stoolBottomHeight,
				Z:    rProjStool * math.Sin(origTheta),
			}).Stage(stage))
		}
	}

	seatGeom := buildSeatGeom(stage, "Stool Seat", 0.0, 0.0, rotSeatTopPoints, rotSeatBottomPoints, stoolTopHeight, seatBottomHeight)
	legsGeom := buildLegsGeom(stage, "Stool Legs", 0.0, 0.0, projSeatBottomEyePoints, projStoolBottomEyePoints, seatBottomHeight, radialRepetitions)
	combinedGeom := buildSeatAndLegsGeom(stage, "Stool Seat and Legs", 0.0, 0.0, rotSeatTopPoints, rotSeatBottomPoints, stoolTopHeight, seatBottomHeight, projSeatBottomEyePoints, projStoolBottomEyePoints, radialRepetitions)

	return &StoolGeometriesResult{
		SeatGeom:     seatGeom,
		LegsGeom:     legsGeom,
		CombinedGeom: combinedGeom,
	}
}

func GenerateStoolOBJ(stager *models.Stager, plant *models.PlantAbstract) string {
	if plant == nil || plant.PlantType != models.Stool || plant.StoolAbstract == nil {
		return ""
	}

	stage := threejs.NewStage("export_obj")
	res := ComputeStoolGeometries(stage, stager, plant)
	if res == nil {
		return ""
	}

	var sb strings.Builder
	sb.WriteString("# Wavefront OBJ file generated by Phylla\n")
	sb.WriteString(fmt.Sprintf("# Plant: %s\n", plant.Name))
	sb.WriteString(fmt.Sprintf("# RadialRepetitions: %d\n\n", plant.StoolAbstract.RadialRepetitions))

	vertexOffset := 0

	writeGeom := func(objectName string, geom *threejs.BufferGeometry) {
		if geom == nil || len(geom.Vertices) == 0 || len(geom.Faces) == 0 {
			return
		}

		sb.WriteString(fmt.Sprintf("o %s\n", objectName))
		for _, v := range geom.Vertices {
			sb.WriteString(fmt.Sprintf("v %.6f %.6f %.6f\n", v.X, v.Y, v.Z))
		}
		for _, f := range geom.Faces {
			// OBJ 1-based indexing
			v1 := f.V1 + vertexOffset + 1
			v2 := f.V2 + vertexOffset + 1
			v3 := f.V3 + vertexOffset + 1
			sb.WriteString(fmt.Sprintf("f %d %d %d\n", v1, v2, v3))
		}
		vertexOffset += len(geom.Vertices)
		sb.WriteString("\n")
	}

	writeGeom("Stool_Seat", res.SeatGeom)
	writeGeom("Stool_Legs", res.LegsGeom)

	return sb.String()
}

func GenerateStoolSTL(stager *models.Stager, plant *models.PlantAbstract) string {
	if plant == nil || plant.PlantType != models.Stool || plant.StoolAbstract == nil {
		return ""
	}

	stage := threejs.NewStage("export_stl")
	res := ComputeStoolGeometries(stage, stager, plant)
	if res == nil || res.CombinedGeom == nil {
		return ""
	}

	geom := res.CombinedGeom
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("solid %s\n", plant.Name))

	writeFacet := func(v1, v2, v3 *threejs.Vector3) {
		e1x := v2.X - v1.X
		e1y := v2.Y - v1.Y
		e1z := v2.Z - v1.Z
		e2x := v3.X - v1.X
		e2y := v3.Y - v1.Y
		e2z := v3.Z - v1.Z

		nx := e1y*e2z - e1z*e2y
		ny := e1z*e2x - e1x*e2z
		nz := e1x*e2y - e1y*e2x
		l := math.Sqrt(nx*nx + ny*ny + nz*nz)
		if l > 0 {
			nx /= l
			ny /= l
			nz /= l
		}

		sb.WriteString(fmt.Sprintf("  facet normal %e %e %e\n", nx, ny, nz))
		sb.WriteString("    outer loop\n")
		sb.WriteString(fmt.Sprintf("      vertex %e %e %e\n", v1.X, v1.Y, v1.Z))
		sb.WriteString(fmt.Sprintf("      vertex %e %e %e\n", v2.X, v2.Y, v2.Z))
		sb.WriteString(fmt.Sprintf("      vertex %e %e %e\n", v3.X, v3.Y, v3.Z))
		sb.WriteString("    endloop\n")
		sb.WriteString("  endfacet\n")
	}

	for _, f := range geom.Faces {
		if f.V1 < len(geom.Vertices) && f.V2 < len(geom.Vertices) && f.V3 < len(geom.Vertices) {
			writeFacet(geom.Vertices[f.V1], geom.Vertices[f.V2], geom.Vertices[f.V3])
		}
	}

	sb.WriteString(fmt.Sprintf("endsolid %s\n", plant.Name))
	return sb.String()
}
