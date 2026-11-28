package stool

import (
	"fmt"
	"math"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

const (
	ghLegBodySegments   = 64
	ghLegFilletSegments = 16
	ghSeatFilletRows    = 8
)

type pt2D struct {
	X, Z float64
}

type grasshopperLegContour struct {
	// Closed loop of length 2*(ghLegBodySegments + ghLegFilletSegments):
	//   [0 .. Nb]               : exterior curve (left tangent -> right tangent)
	//   [Nb+1 .. Nb+Nf-1]       : right tip fillet arc
	//   [Nb+Nf .. 2*Nb+Nf]      : interior curve (right tangent -> left tangent)
	//   [2*Nb+Nf+1 .. 2*(Nb+Nf)-1]: left tip fillet arc
	Loop []pt2D
	Nb   int
	Nf   int
	Y    float64
}

type grasshopperContoursResult struct {
	SeatTopRing     []pt2D
	SeatBottomRing  []pt2D
	StoolTopHeight  float64
	SeatBottomY     float64
	LegTopContour   *grasshopperLegContour
	LegBotContour   *grasshopperLegContour
	Eye3DPoints     []*threejs.Vector3
	RadialReps      int
	GrowthThetaStep float64
	GrowthYStep     float64
	SeatFilletR     float64
}

func dist2D(a, b pt2D) float64 {
	return math.Hypot(b.X-a.X, b.Z-a.Z)
}

func lerp2D(a, b pt2D, t float64) pt2D {
	return pt2D{
		X: a.X + t*(b.X-a.X),
		Z: a.Z + t*(b.Z-a.Z),
	}
}

func polylineNormalsOutward(poly []pt2D, refDir pt2D, useRefAtMid bool) []pt2D {
	n := len(poly)
	normals := make([]pt2D, n)
	if n == 0 {
		return normals
	}
	if n == 1 {
		l := math.Hypot(refDir.X, refDir.Z)
		if l > 1e-9 {
			normals[0] = pt2D{X: refDir.X / l, Z: refDir.Z / l}
		} else {
			normals[0] = pt2D{X: 1, Z: 0}
		}
		return normals
	}

	for i := range n {
		var tx, tz float64
		switch i {
		case 0:
			tx = poly[1].X - poly[0].X
			tz = poly[1].Z - poly[0].Z
		case n - 1:
			tx = poly[n-1].X - poly[n-2].X
			tz = poly[n-1].Z - poly[n-2].Z
		default:
			tx = poly[i+1].X - poly[i-1].X
			tz = poly[i+1].Z - poly[i-1].Z
		}
		l := math.Hypot(tx, tz)
		if l > 1e-12 {
			// For increasing polar angle in (X=r cos θ, Z=r sin θ), (tz, -tx) points radially outward.
			normals[i] = pt2D{X: tz / l, Z: -tx / l}
		} else if i > 0 {
			normals[i] = normals[i-1]
		} else {
			normals[i] = pt2D{X: 1, Z: 0}
		}
	}

	sign := 1.0
	mid := n / 2
	if useRefAtMid {
		if normals[mid].X*refDir.X+normals[mid].Z*refDir.Z < 0 {
			sign = -1.0
		}
	} else {
		if normals[mid].X*poly[mid].X+normals[mid].Z*poly[mid].Z < 0 {
			sign = -1.0
		}
	}
	if sign < 0 {
		for i := range normals {
			normals[i].X = -normals[i].X
			normals[i].Z = -normals[i].Z
		}
	}
	return normals
}

func offsetOpenPolyline(poly []pt2D, dist float64, towardsDir pt2D) []pt2D {
	n := len(poly)
	out := make([]pt2D, n)
	if n == 0 {
		return out
	}
	normals := polylineNormalsOutward(poly, towardsDir, true)
	for i := range n {
		out[i] = pt2D{
			X: poly[i].X + dist*normals[i].X,
			Z: poly[i].Z + dist*normals[i].Z,
		}
	}
	return out
}

// intersectSegments2D checks whether segment [a0, a1] intersects [b0, b1] and returns
// (pt, ta, tb, true) with ta, tb in [0, 1].
func intersectSegments2D(a0, a1, b0, b1 pt2D) (pt2D, float64, float64, bool) {
	rx := a1.X - a0.X
	rz := a1.Z - a0.Z
	sx := b1.X - b0.X
	sz := b1.Z - b0.Z

	denom := rx*sz - rz*sx
	if math.Abs(denom) < 1e-12 {
		return pt2D{}, 0, 0, false
	}

	qpx := b0.X - a0.X
	qpz := b0.Z - a0.Z

	ta := (qpx*sz - qpz*sx) / denom
	tb := (qpx*rz - qpz*rx) / denom

	const eps = 1e-7
	if ta >= -eps && ta <= 1.0+eps && tb >= -eps && tb <= 1.0+eps {
		ta = math.Max(0.0, math.Min(1.0, ta))
		tb = math.Max(0.0, math.Min(1.0, tb))
		return lerp2D(a0, a1, ta), ta, tb, true
	}
	return pt2D{}, 0, 0, false
}

// findPolylineIntersectionLeft searches from midA/midB down toward 0 for the first crossing.
func findPolylineIntersectionLeft(polyA, polyB []pt2D) (pt pt2D, paramA, paramB float64, found bool) {
	nA := len(polyA)
	nB := len(polyB)
	if nA < 2 || nB < 2 {
		return pt2D{}, 0, 0, false
	}
	midA := (nA - 1) / 2
	midB := (nB - 1) / 2
	for i := midA - 1; i >= 0; i-- {
		for j := midB - 1; j >= 0; j-- {
			if p, ta, tb, ok := intersectSegments2D(polyA[i], polyA[i+1], polyB[j], polyB[j+1]); ok {
				return p, float64(i) + ta, float64(j) + tb, true
			}
		}
	}
	return pt2D{}, 0, 0, false
}

// findPolylineIntersectionRight searches from midA/midB up toward end for the first crossing.
func findPolylineIntersectionRight(polyA, polyB []pt2D) (pt pt2D, paramA, paramB float64, found bool) {
	nA := len(polyA)
	nB := len(polyB)
	if nA < 2 || nB < 2 {
		return pt2D{}, 0, 0, false
	}
	midA := (nA - 1) / 2
	midB := (nB - 1) / 2
	for i := midA; i < nA-1; i++ {
		for j := midB; j < nB-1; j++ {
			if p, ta, tb, ok := intersectSegments2D(polyA[i], polyA[i+1], polyB[j], polyB[j+1]); ok {
				return p, float64(i) + ta, float64(j) + tb, true
			}
		}
	}
	return pt2D{}, 0, 0, false
}

// trimPolyline extracts the sub-polyline between continuous index parameters [pStart, pEnd].
func trimPolyline(poly []pt2D, pStart, pEnd float64) []pt2D {
	n := len(poly)
	if n == 0 {
		return nil
	}
	if pStart < 0 {
		pStart = 0
	}
	maxP := float64(n - 1)
	if pEnd > maxP {
		pEnd = maxP
	}
	if pEnd <= pStart+1e-9 {
		pt := evalPolylineAtParam(poly, pStart)
		return []pt2D{pt, pt}
	}

	var res []pt2D
	res = append(res, evalPolylineAtParam(poly, pStart))
	firstIdx := int(math.Floor(pStart)) + 1
	lastIdx := int(math.Ceil(pEnd)) - 1
	for k := firstIdx; k <= lastIdx; k++ {
		if k >= 0 && k < n {
			if dist2D(res[len(res)-1], poly[k]) > 1e-7 {
				res = append(res, poly[k])
			}
		}
	}
	endPt := evalPolylineAtParam(poly, pEnd)
	if dist2D(res[len(res)-1], endPt) > 1e-7 || len(res) == 1 {
		res = append(res, endPt)
	}
	return res
}

func evalPolylineAtParam(poly []pt2D, param float64) pt2D {
	n := len(poly)
	if n == 0 {
		return pt2D{}
	}
	if param <= 0 {
		return poly[0]
	}
	if param >= float64(n-1) {
		return poly[n-1]
	}
	idx := int(math.Floor(param))
	if idx >= n-1 {
		return poly[n-1]
	}
	t := param - float64(idx)
	return lerp2D(poly[idx], poly[idx+1], t)
}

func closestPointOnPolylineRange(poly []pt2D, q pt2D, iMin, iMax int) (bestPt pt2D, bestParam float64) {
	n := len(poly)
	if n == 0 {
		return pt2D{}, 0
	}
	if n == 1 {
		return poly[0], 0
	}
	if iMin < 0 {
		iMin = 0
	}
	if iMax > n-1 {
		iMax = n - 1
	}
	if iMax <= iMin {
		return poly[iMin], float64(iMin)
	}

	bestDistSq := math.MaxFloat64
	for i := iMin; i < iMax; i++ {
		a := poly[i]
		b := poly[i+1]
		dx := b.X - a.X
		dz := b.Z - a.Z
		lenSq := dx*dx + dz*dz
		t := 0.0
		if lenSq > 1e-18 {
			t = ((q.X-a.X)*dx + (q.Z-a.Z)*dz) / lenSq
			t = math.Max(0.0, math.Min(1.0, t))
		}
		proj := pt2D{X: a.X + t*dx, Z: a.Z + t*dz}
		dsq := (q.X-proj.X)*(q.X-proj.X) + (q.Z-proj.Z)*(q.Z-proj.Z)
		if dsq < bestDistSq {
			bestDistSq = dsq
			bestPt = proj
			bestParam = float64(i) + t
		}
	}
	return bestPt, bestParam
}

func resamplePolylineUniform(poly []pt2D, numSegments int) []pt2D {
	if numSegments < 1 {
		numSegments = 1
	}
	if len(poly) == 0 {
		return make([]pt2D, numSegments+1)
	}
	if len(poly) == 1 {
		out := make([]pt2D, numSegments+1)
		for i := range out {
			out[i] = poly[0]
		}
		return out
	}

	cum := make([]float64, len(poly))
	total := 0.0
	for i := 1; i < len(poly); i++ {
		total += dist2D(poly[i-1], poly[i])
		cum[i] = total
	}
	if total <= 1e-12 {
		out := make([]pt2D, numSegments+1)
		for i := range out {
			out[i] = poly[0]
		}
		return out
	}

	out := make([]pt2D, numSegments+1)
	out[0] = poly[0]
	out[numSegments] = poly[len(poly)-1]
	segIdx := 0
	for k := 1; k < numSegments; k++ {
		target := float64(k) * total / float64(numSegments)
		for segIdx < len(cum)-2 && cum[segIdx+1] < target {
			segIdx++
		}
		segLen := cum[segIdx+1] - cum[segIdx]
		t := 0.0
		if segLen > 1e-12 {
			t = (target - cum[segIdx]) / segLen
		}
		out[k] = lerp2D(poly[segIdx], poly[segIdx+1], t)
	}
	return out
}

// sampleTipArc samples internal points (1 .. numSegments-1) of the circular arc centered at center
// from pStart to pEnd, bulging outward toward tipVertex.
func sampleTipArc(center, pStart, pEnd, tipVertex pt2D, numSegments int) []pt2D {
	if numSegments <= 1 {
		return nil
	}
	a0 := math.Atan2(pStart.Z-center.Z, pStart.X-center.X)
	a1 := math.Atan2(pEnd.Z-center.Z, pEnd.X-center.X)
	r0 := dist2D(center, pStart)
	r1 := dist2D(center, pEnd)

	// Normalize delta into (-pi, pi] first (short arc).
	delta := a1 - a0
	for delta <= -math.Pi {
		delta += 2 * math.Pi
	}
	for delta > math.Pi {
		delta -= 2 * math.Pi
	}

	// Verify the midpoint of (a0, a0+delta) bulges toward tipVertex rather than inward.
	midAngle := a0 + 0.5*delta
	rmid := 0.5 * (r0 + r1)
	arcMid := pt2D{X: center.X + rmid*math.Cos(midAngle), Z: center.Z + rmid*math.Sin(midAngle)}
	toTip := pt2D{X: tipVertex.X - center.X, Z: tipVertex.Z - center.Z}
	toMid := pt2D{X: arcMid.X - center.X, Z: arcMid.Z - center.Z}
	if toTip.X*toMid.X+toTip.Z*toMid.Z < 0 {
		if delta > 0 {
			delta -= 2 * math.Pi
		} else {
			delta += 2 * math.Pi
		}
	}

	pts := make([]pt2D, 0, numSegments-1)
	for k := 1; k < numSegments; k++ {
		u := float64(k) / float64(numSegments)
		ang := a0 + u*delta
		r := r0 + u*(r1-r0)
		pts = append(pts, pt2D{
			X: center.X + r*math.Cos(ang),
			Z: center.Z + r*math.Sin(ang),
		})
	}
	return pts
}

// buildGrasshopperCrescentOnPlane constructs the 2D crescent contour at horizontal plane yPlane:
//  1. Projects outer & inner cylinder curves onto yPlane.
//  2. Offsets interior curve toward exterior curve by innerOffsetDist and trims at their two crossings.
//  3. Fillets both sharp tips with rolling-ball circular arcs of radius tipFilletRadius.
func buildGrasshopperCrescentOnPlane(
	angles []float64,
	yOuter []float64,
	yInner []float64,
	globalR float64,
	projAngleRad float64,
	yPlane float64,
	innerOffsetDist float64,
	tipFilletRadius float64,
) *grasshopperLegContour {
	n := len(angles)
	if n < 5 {
		return nil
	}

	extRaw := make([]pt2D, n)
	intRaw := make([]pt2D, n)
	tanProj := math.Tan(projAngleRad)

	mid := n / 2
	r1Mid := globalR + (yPlane-yOuter[mid])*tanProj
	r2Mid := globalR + (yPlane-yInner[mid])*tanProj
	swap := r2Mid > r1Mid

	for i := range n {
		yo, yi := yOuter[i], yInner[i]
		if swap {
			yo, yi = yi, yo
		}
		rExt := math.Max(1.0, globalR+(yPlane-yo)*tanProj)
		rInt := math.Max(1.0, globalR+(yPlane-yi)*tanProj)
		c, s := math.Cos(angles[i]), math.Sin(angles[i])
		extRaw[i] = pt2D{X: rExt * c, Z: rExt * s}
		intRaw[i] = pt2D{X: rInt * c, Z: rInt * s}
	}

	// Ensure exact closure at both cusps before offsetting.
	leftCusp := lerp2D(extRaw[0], intRaw[0], 0.5)
	rightCusp := lerp2D(extRaw[n-1], intRaw[n-1], 0.5)
	extRaw[0], intRaw[0] = leftCusp, leftCusp
	extRaw[n-1], intRaw[n-1] = rightCusp, rightCusp

	midThickness := dist2D(extRaw[mid], intRaw[mid])
	if midThickness <= 1e-5 {
		return nil
	}

	inToOutDir := pt2D{
		X: extRaw[mid].X - intRaw[mid].X,
		Z: extRaw[mid].Z - intRaw[mid].Z,
	}

	effOffset := math.Max(0.0, math.Min(innerOffsetDist, 0.45*midThickness))
	extTrim := extRaw
	intTrim := intRaw

	if effOffset > 1e-6 {
		intOff := offsetOpenPolyline(intRaw, effOffset, inToOutDir)
		pL, sL, tL, okL := findPolylineIntersectionLeft(extRaw, intOff)
		pR, sR, tR, okR := findPolylineIntersectionRight(extRaw, intOff)
		if okL && okR && sR > sL+1e-4 && tR > tL+1e-4 {
			extTrim = trimPolyline(extRaw, sL, sR)
			intTrim = trimPolyline(intOff, tL, tR)
			extTrim[0], intTrim[0] = pL, pL
			extTrim[len(extTrim)-1], intTrim[len(intTrim)-1] = pR, pR
		}
	}

	if len(extTrim) < 3 || len(intTrim) < 3 {
		return nil
	}

	midExt := extTrim[len(extTrim)/2]
	midInt := intTrim[len(intTrim)/2]
	trimmedWidth := dist2D(midExt, midInt)
	if trimmedWidth <= 1e-5 {
		return nil
	}

	outToInDir := pt2D{X: midInt.X - midExt.X, Z: midInt.Z - midExt.Z}
	inToOutTrimDir := pt2D{X: midExt.X - midInt.X, Z: midExt.Z - midInt.Z}

	effFillet := math.Max(0.0, math.Min(tipFilletRadius, 0.40*trimmedWidth))

	pTipLeft := extTrim[0]
	pTipRight := extTrim[len(extTrim)-1]

	var extBody, intBodyRev []pt2D
	var rightArc, leftArc []pt2D

	if effFillet > 1e-5 {
		extIn := offsetOpenPolyline(extTrim, effFillet, outToInDir)
		intIn := offsetOpenPolyline(intTrim, effFillet, inToOutTrimDir)
		cL, _, _, okCL := findPolylineIntersectionLeft(extIn, intIn)
		cR, _, _, okCR := findPolylineIntersectionRight(extIn, intIn)

		if okCL && okCR {
			tLExtPt, sLExt := closestPointOnPolylineRange(extTrim, cL, 0, len(extTrim)/2)
			tLIntPt, sLInt := closestPointOnPolylineRange(intTrim, cL, 0, len(intTrim)/2)
			tRExtPt, sRExt := closestPointOnPolylineRange(extTrim, cR, len(extTrim)/2, len(extTrim)-1)
			tRIntPt, sRInt := closestPointOnPolylineRange(intTrim, cR, len(intTrim)/2, len(intTrim)-1)

			if sRExt > sLExt+1e-4 && sRInt > sLInt+1e-4 {
				extSub := trimPolyline(extTrim, sLExt, sRExt)
				extSub[0] = tLExtPt
				extSub[len(extSub)-1] = tRExtPt
				extBody = resamplePolylineUniform(extSub, ghLegBodySegments)

				intSub := trimPolyline(intTrim, sLInt, sRInt)
				intSub[0] = tLIntPt
				intSub[len(intSub)-1] = tRIntPt
				intSampled := resamplePolylineUniform(intSub, ghLegBodySegments)
				intBodyRev = make([]pt2D, len(intSampled))
				for i := range intSampled {
					intBodyRev[i] = intSampled[len(intSampled)-1-i]
				}

				rightArc = sampleTipArc(cR, tRExtPt, tRIntPt, pTipRight, ghLegFilletSegments)
				leftArc = sampleTipArc(cL, tLIntPt, tLExtPt, pTipLeft, ghLegFilletSegments)
			}
		}
	}

	// Fallback if fillet radius is 0 or tiny relative to discretization.
	if len(extBody) == 0 {
		// Trim a tiny fraction at the tips so left/right chords remain well-defined.
		nE := len(extTrim)
		nI := len(intTrim)
		epsE := 0.01 * float64(nE-1)
		epsI := 0.01 * float64(nI-1)
		extSub := trimPolyline(extTrim, epsE, float64(nE-1)-epsE)
		intSub := trimPolyline(intTrim, epsI, float64(nI-1)-epsI)
		extBody = resamplePolylineUniform(extSub, ghLegBodySegments)
		intSampled := resamplePolylineUniform(intSub, ghLegBodySegments)
		intBodyRev = make([]pt2D, len(intSampled))
		for i := range intSampled {
			intBodyRev[i] = intSampled[len(intSampled)-1-i]
		}
		rightArc = make([]pt2D, ghLegFilletSegments-1)
		for k := 1; k < ghLegFilletSegments; k++ {
			u := float64(k) / float64(ghLegFilletSegments)
			rightArc[k-1] = lerp2D(extBody[ghLegBodySegments], intBodyRev[0], u)
		}
		leftArc = make([]pt2D, ghLegFilletSegments-1)
		for k := 1; k < ghLegFilletSegments; k++ {
			u := float64(k) / float64(ghLegFilletSegments)
			leftArc[k-1] = lerp2D(intBodyRev[ghLegBodySegments], extBody[0], u)
		}
	}

	loop := make([]pt2D, 0, 2*(ghLegBodySegments+ghLegFilletSegments))
	loop = append(loop, extBody...)    // [0 .. Nb]
	loop = append(loop, rightArc...)   // [Nb+1 .. Nb+Nf-1]
	loop = append(loop, intBodyRev...) // [Nb+Nf .. 2*Nb+Nf]
	loop = append(loop, leftArc...)    // [2*Nb+Nf+1 .. 2*(Nb+Nf)-1]

	return &grasshopperLegContour{
		Loop: loop,
		Nb:   ghLegBodySegments,
		Nf:   ghLegFilletSegments,
		Y:    yPlane,
	}
}

func computeGrasshopperContours(
	stage *threejs.Stage,
	plant *models.PlantAbstract,
	resampledBaseCurve *threejs.Curve,
	radialRepetitions int,
	globalR float64,
	stoolTopHeight float64,
	torusHeight float64,
	growthVectorX float64,
	growthVectorY float64,
	projAngleRad float64,
	expectedDegrees float64,
) *grasshopperContoursResult {
	if plant == nil || plant.StoolAbstract == nil || resampledBaseCurve == nil || len(resampledBaseCurve.Points) < 3 {
		return nil
	}

	sa := plant.StoolAbstract
	expectedRad := expectedDegrees * math.Pi / 180.0
	numBaseIntervals := len(resampledBaseCurve.Points) - 1
	if numBaseIntervals < 2 || expectedRad <= 0 {
		return nil
	}
	baseStepRad := expectedRad / float64(numBaseIntervals)

	evalBaseY := func(angle float64) float64 {
		a := math.Mod(angle, expectedRad)
		if a < 0 {
			a += expectedRad
		}
		f := a / baseStepRad
		idx := int(math.Floor(f))
		if idx < 0 {
			idx = 0
		}
		if idx >= numBaseIntervals {
			idx = numBaseIntervals - 1
		}
		t := f - float64(idx)
		y0 := resampledBaseCurve.Points[idx].Y
		y1 := resampledBaseCurve.Points[idx+1].Y
		return y0 + t*(y1-y0)
	}

	// Seat thickness in Grasshopper mode is locked to the vertical stacking rise (growthVectorY)
	// so stacked stools rest flush seat-on-seat, falling back to RelativeSeatThickness if needed.
	seatThickness := growthVectorY
	if seatThickness <= 1e-4 {
		seatThickness = sa.RelativeSeatThickness * plant.RhombusSideLength
	}
	if seatThickness > 0.85*stoolTopHeight {
		seatThickness = 0.85 * stoolTopHeight
	}
	seatBottomY := stoolTopHeight - seatThickness

	thetaOffset := growthVectorX / globalR
	tanProj := math.Tan(projAngleRad)

	// Build deduplicated 360° Seat Top and Seat Bottom rings (no duplicate seam vertices).
	samplesPerRep := numBaseIntervals
	totalSeatPts := radialRepetitions * samplesPerRep
	seatTopRing := make([]pt2D, 0, totalSeatPts)
	seatBottomRing := make([]pt2D, 0, totalSeatPts)

	for k := range radialRepetitions {
		baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(radialRepetitions)
		totalThetaOffset := baseThetaOffset + thetaOffset
		for i := range samplesPerRep {
			pt := resampledBaseCurve.Points[i]
			origTheta := math.Atan2(pt.Z, pt.X)
			r := math.Hypot(pt.X, pt.Z)
			newTheta := origTheta + totalThetaOffset
			ptY := pt.Y + growthVectorY + torusHeight

			rTop := r + (stoolTopHeight-ptY)*tanProj
			rBot := r + (seatBottomY-ptY)*tanProj
			c, s := math.Cos(newTheta), math.Sin(newTheta)

			seatTopRing = append(seatTopRing, pt2D{X: rTop * c, Z: rTop * s})
			seatBottomRing = append(seatBottomRing, pt2D{X: rBot * c, Z: rBot * s})
		}
	}

	// Fine angular sampling over one period to extract the exact crescent between yRot and yBase.
	const subSteps = 4
	fineCount := numBaseIntervals * subSteps
	fineStep := expectedRad / float64(fineCount)

	yBaseFine := make([]float64, fineCount)
	yRotFine := make([]float64, fineCount)
	gapFine := make([]float64, fineCount)

	minIdx := 0
	minGap := math.MaxFloat64
	for m := range fineCount {
		alpha := float64(m) * fineStep
		yb := evalBaseY(alpha) + torusHeight
		yr := evalBaseY(alpha-thetaOffset) + growthVectorY + torusHeight
		g := math.Abs(yr - yb)
		yBaseFine[m] = yb
		yRotFine[m] = yr
		gapFine[m] = g
		if g < minGap {
			minGap = g
			minIdx = m
		}
	}

	// Unwrap starting at minIdx (middle of contact zone) so the crescent never straddles array ends.
	uAngles := make([]float64, fineCount+1)
	uBase := make([]float64, fineCount+1)
	uRot := make([]float64, fineCount+1)
	uGap := make([]float64, fineCount+1)

	startAlpha := float64(minIdx) * fineStep
	peakJ := 0
	maxGap := -1.0
	for j := 0; j <= fineCount; j++ {
		idx := (minIdx + j) % fineCount
		uAngles[j] = startAlpha + float64(j)*fineStep
		uBase[j] = yBaseFine[idx]
		uRot[j] = yRotFine[idx]
		uGap[j] = gapFine[idx]
		if uGap[j] > maxGap {
			maxGap = uGap[j]
			peakJ = j
		}
	}

	touchTol := minGap + math.Max(1e-3*maxGap, 1e-4)
	jStart := peakJ
	for jStart > 0 && uGap[jStart] > touchTol {
		jStart--
	}
	jEnd := peakJ
	for jEnd < fineCount && uGap[jEnd] > touchTol {
		jEnd++
	}

	cresAngles := uAngles[jStart : jEnd+1]
	cresOuter := make([]float64, len(cresAngles))
	cresInner := make([]float64, len(cresAngles))
	for i := range cresAngles {
		yb := uBase[jStart+i]
		yr := uRot[jStart+i]
		if yr >= yb {
			cresOuter[i] = yr
			cresInner[i] = yb
		} else {
			cresOuter[i] = yb
			cresInner[i] = yr
		}
	}

	// Build 3D eye loop on cylinder for optional debug visualization.
	var eye3DPts []*threejs.Vector3
	if stage != nil && len(cresAngles) >= 2 {
		stepVis := max(len(cresAngles)/80, 1)
		for i := 0; i < len(cresAngles); i += stepVis {
			a := cresAngles[i]
			eye3DPts = append(eye3DPts, (&threejs.Vector3{
				Name: fmt.Sprintf("GH Eye Bottom %.1f", a*180.0/math.Pi),
				X:    globalR * math.Cos(a),
				Y:    cresInner[i],
				Z:    globalR * math.Sin(a),
			}).Stage(stage))
		}
		for i := len(cresAngles) - 1; i >= 0; i -= stepVis {
			a := cresAngles[i]
			eye3DPts = append(eye3DPts, (&threejs.Vector3{
				Name: fmt.Sprintf("GH Eye Top %.1f", a*180.0/math.Pi),
				X:    globalR * math.Cos(a),
				Y:    cresOuter[i],
				Z:    globalR * math.Sin(a),
			}).Stage(stage))
		}
	}

	innerOffsetDist := sa.RelativeInnerLegOffset * plant.RhombusSideLength
	tipFilletRadius := sa.RelativeLegTipFilletRadius * plant.RhombusSideLength
	seatFilletRadius := sa.RelativeSeatTopFilletRadius * plant.RhombusSideLength

	legTopY := seatBottomY
	if sa.IsLegsReachTop {
		legTopY = stoolTopHeight
	}
	legBotY := 0.0

	legTopContour := buildGrasshopperCrescentOnPlane(
		cresAngles, cresOuter, cresInner, globalR, projAngleRad, legTopY, innerOffsetDist, tipFilletRadius,
	)
	legBotContour := buildGrasshopperCrescentOnPlane(
		cresAngles, cresOuter, cresInner, globalR, projAngleRad, legBotY, innerOffsetDist, tipFilletRadius,
	)

	return &grasshopperContoursResult{
		SeatTopRing:     seatTopRing,
		SeatBottomRing:  seatBottomRing,
		StoolTopHeight:  stoolTopHeight,
		SeatBottomY:     seatBottomY,
		LegTopContour:   legTopContour,
		LegBotContour:   legBotContour,
		Eye3DPoints:     eye3DPts,
		RadialReps:      radialRepetitions,
		GrowthThetaStep: thetaOffset,
		GrowthYStep:     growthVectorY,
		SeatFilletR:     seatFilletRadius,
	}
}

func appendGrasshopperSeatToGeom(
	stage *threejs.Stage,
	geom *threejs.BufferGeometry,
	geomName string,
	res *grasshopperContoursResult,
	deltaTheta float64,
	deltaY float64,
) {
	if res == nil || len(res.SeatTopRing) < 3 || len(res.SeatTopRing) != len(res.SeatBottomRing) {
		return
	}

	N := len(res.SeatTopRing)
	cosO := math.Cos(deltaTheta)
	sinO := math.Sin(deltaTheta)
	rotPt := func(p pt2D) pt2D {
		return pt2D{
			X: p.X*cosO - p.Z*sinO,
			Z: p.X*sinO + p.Z*cosO,
		}
	}

	topRing := make([]pt2D, N)
	botRing := make([]pt2D, N)
	var sumTopX, sumTopZ, sumBotX, sumBotZ float64
	minTopR := math.MaxFloat64
	for i := range N {
		tp := rotPt(res.SeatTopRing[i])
		bp := rotPt(res.SeatBottomRing[i])
		topRing[i] = tp
		botRing[i] = bp
		sumTopX += tp.X
		sumTopZ += tp.Z
		sumBotX += bp.X
		sumBotZ += bp.Z
		if r := math.Hypot(tp.X, tp.Z); r < minTopR {
			minTopR = r
		}
	}

	topY := res.StoolTopHeight + deltaY
	botY := res.SeatBottomY + deltaY
	seatH := topY - botY
	if seatH <= 1e-6 {
		return
	}

	normals2D := polylineNormalsOutward(topRing, pt2D{}, false)
	// Enforce periodic wrap-around normal at endpoints 0 and N-1.
	{
		for _, idx := range []int{0, N - 1} {
			prev := (idx - 1 + N) % N
			next := (idx + 1) % N
			tx := topRing[next].X - topRing[prev].X
			tz := topRing[next].Z - topRing[prev].Z
			if l := math.Hypot(tx, tz); l > 1e-12 {
				nx, nz := tz/l, -tx/l
				if nx*topRing[idx].X+nz*topRing[idx].Z < 0 {
					nx, nz = -nx, -nz
				}
				normals2D[idx] = pt2D{X: nx, Z: nz}
			}
		}
	}

	rf := math.Max(0.0, res.SeatFilletR)
	rf = math.Min(rf, 0.42*seatH)
	rf = math.Min(rf, 0.25*minTopR)

	numFilletRows := 0
	if rf > 1e-4 {
		numFilletRows = ghSeatFilletRows
	}

	// Build smooth top-fillet-wall vertex rings:
	//   row 0 .. numFilletRows : fillet rows from Q_top (on topY) to Q_wall (on sloped wall)
	//   row numFilletRows + 1  : bottom of sloped wall at botY
	numSmoothRows := numFilletRows + 2
	smoothBaseIdx := len(geom.Vertices)

	if numFilletRows == 0 {
		for i := range N {
			geom.Vertices = append(geom.Vertices, (&threejs.Vector3{
				Name: fmt.Sprintf("%s Seat Top %d", geomName, i),
				X:    topRing[i].X, Y: topY, Z: topRing[i].Z,
			}).Stage(stage))
		}
	} else {
		gammas := make([]float64, N)
		tanBetas := make([]float64, N)
		dInsets := make([]float64, N)
		maxRatio := 1.0
		for i := range N {
			dx := botRing[i].X - topRing[i].X
			dz := botRing[i].Z - topRing[i].Z
			w := dx*normals2D[i].X + dz*normals2D[i].Z
			tb := w / seatH
			beta := math.Atan(tb)
			gamma := 0.5*math.Pi - beta
			if gamma < 0.05 {
				gamma = 0.05
			}
			if gamma > math.Pi-0.05 {
				gamma = math.Pi - 0.05
			}
			gammas[i] = gamma
			tanBetas[i] = tb
			d := rf * math.Tan(0.5*gamma)
			dInsets[i] = d
			drop := rf * (1.0 - math.Cos(gamma))
			if drop > 0.45*seatH {
				if r := (0.45 * seatH) / drop; r < maxRatio {
					maxRatio = r
				}
			}
		}
		if maxRatio < 1.0 {
			rf *= maxRatio
			for i := range N {
				dInsets[i] *= maxRatio
			}
		}

		for s := 0; s <= numFilletRows; s++ {
			frac := float64(s) / float64(numFilletRows)
			for i := range N {
				psi := frac * gammas[i]
				dropY := rf * (1.0 - math.Cos(psi))
				uWall := dropY / seatH
				wx := topRing[i].X + uWall*(botRing[i].X-topRing[i].X)
				wz := topRing[i].Z + uWall*(botRing[i].Z-topRing[i].Z)
				inwardShift := dInsets[i] - rf*math.Sin(psi) + dropY*tanBetas[i]
				if s == numFilletRows {
					inwardShift = 0.0
				}
				geom.Vertices = append(geom.Vertices, (&threejs.Vector3{
					Name: fmt.Sprintf("%s Seat Fillet r%d %d", geomName, s, i),
					X:    wx - inwardShift*normals2D[i].X,
					Y:    topY - dropY,
					Z:    wz - inwardShift*normals2D[i].Z,
				}).Stage(stage))
			}
		}
	}

	// Bottom row of outer sloped wall (at botY).
	for i := range N {
		geom.Vertices = append(geom.Vertices, (&threejs.Vector3{
			Name: fmt.Sprintf("%s Seat Wall Bot %d", geomName, i),
			X:    botRing[i].X, Y: botY, Z: botRing[i].Z,
		}).Stage(stage))
	}

	// Top center vertex (shares row 0 of smoothBaseIdx since fillet meets top tangentially).
	topCenterIdx := len(geom.Vertices)
	geom.Vertices = append(geom.Vertices, (&threejs.Vector3{
		Name: fmt.Sprintf("%s Seat Top Center", geomName),
		X:    sumTopX / float64(N), Y: topY, Z: sumTopZ / float64(N),
	}).Stage(stage))

	// Separate bottom cap ring + bottom center for crisp 90° crease normals in Three.js,
	// with bitwise-identical (X, Y, Z) coordinates to Seat Wall Bot for watertight STL export.
	botCapBaseIdx := len(geom.Vertices)
	for i := range N {
		geom.Vertices = append(geom.Vertices, (&threejs.Vector3{
			Name: fmt.Sprintf("%s Seat Cap Bot %d", geomName, i),
			X:    botRing[i].X, Y: botY, Z: botRing[i].Z,
		}).Stage(stage))
	}
	botCenterIdx := len(geom.Vertices)
	geom.Vertices = append(geom.Vertices, (&threejs.Vector3{
		Name: fmt.Sprintf("%s Seat Bottom Center", geomName),
		X:    sumBotX / float64(N), Y: botY, Z: sumBotZ / float64(N),
	}).Stage(stage))

	// 1. Top flat cap (facing +Y): (topCenter, nextI, i) on row 0.
	for i := range N {
		nextI := (i + 1) % N
		geom.Faces = append(geom.Faces, (&threejs.Triangle{
			Name: fmt.Sprintf("%s Seat Top Face %d", geomName, i),
			V1:   topCenterIdx,
			V2:   smoothBaseIdx + nextI,
			V3:   smoothBaseIdx + i,
		}).Stage(stage))
	}

	// 2. Fillet strips + sloped outer wall quads across consecutive rows r -> r+1.
	for r := range numSmoothRows - 1 {
		rowTop := smoothBaseIdx + r*N
		rowBot := smoothBaseIdx + (r+1)*N
		for i := range N {
			nextI := (i + 1) % N
			tI := rowTop + i
			tNext := rowTop + nextI
			bI := rowBot + i
			bNext := rowBot + nextI

			geom.Faces = append(geom.Faces,
				(&threejs.Triangle{
					Name: fmt.Sprintf("%s Seat Strip r%d T1 %d", geomName, r, i),
					V1:   bI, V2: tI, V3: tNext,
				}).Stage(stage),
				(&threejs.Triangle{
					Name: fmt.Sprintf("%s Seat Strip r%d T2 %d", geomName, r, i),
					V1:   bI, V2: tNext, V3: bNext,
				}).Stage(stage),
			)
		}
	}

	// 3. Bottom flat cap (facing -Y): (botCenter, botCapI, botCapNextI).
	for i := range N {
		nextI := (i + 1) % N
		geom.Faces = append(geom.Faces, (&threejs.Triangle{
			Name: fmt.Sprintf("%s Seat Bot Face %d", geomName, i),
			V1:   botCenterIdx,
			V2:   botCapBaseIdx + i,
			V3:   botCapBaseIdx + nextI,
		}).Stage(stage))
	}
}

// appendCrescentCapFaces triangulates a closed crescent loop of layout:
//   ext:      0 .. Nb
//   rightTip: Nb .. Nb+Nf
//   intRev:   Nb+Nf .. 2*Nb+Nf  (index Nb+Nf+m opposes ext[Nb-m])
//   leftTip:  2*Nb+Nf .. M      (where M wraps to 0)
// When facingUp is true, normals point +Y; otherwise -Y.
func appendCrescentCapFaces(
	stage *threejs.Stage,
	geom *threejs.BufferGeometry,
	namePrefix string,
	baseIdx int,
	nb, nf int,
	facingUp bool,
) {
	M := 2 * (nb + nf)
	addTri := func(a, b, c int) {
		if !facingUp {
			b, c = c, b
		}
		geom.Faces = append(geom.Faces, (&threejs.Triangle{
			Name: namePrefix,
			V1:   baseIdx + (a % M),
			V2:   baseIdx + (b % M),
			V3:   baseIdx + (c % M),
		}).Stage(stage))
	}

	// 1. Middle crescent strip: quads between ext[k..k+1] and opposing intRev points.
	// Opposing index for ext[k] is idxInt(k) = 2*nb + nf - k.
	for k := range nb {
		e0 := k
		e1 := k + 1
		i0 := 2*nb + nf - k
		i1 := 2*nb + nf - (k + 1)
		// (e0, i0, e1) and (e1, i0, i1) have +Y normal in (X, Y, Z).
		addTri(e0, i0, e1)
		addTri(e1, i0, i1)
	}

	// 2. Right convex circular tip cap: vertices R_j = nb + j for j = 0 .. nf (R_0 = ext[nb], R_nf = intRev[0]).
	r0 := nb
	for j := 1; j < nf; j++ {
		addTri(r0, nb+j+1, nb+j)
	}

	// 3. Left convex circular tip cap: vertices L_j = 2*nb + nf + j for j = 0 .. nf (L_0 = intRev[nb], L_nf = M -> 0).
	l0 := 2*nb + nf
	for j := 1; j < nf; j++ {
		addTri(l0, l0+j+1, l0+j)
	}
}

func appendGrasshopperLegsToGeom(
	stage *threejs.Stage,
	geom *threejs.BufferGeometry,
	geomName string,
	res *grasshopperContoursResult,
	deltaTheta float64,
	deltaY float64,
) {
	if res == nil || res.LegTopContour == nil || res.LegBotContour == nil {
		return
	}
	topC := res.LegTopContour
	botC := res.LegBotContour
	M := len(topC.Loop)
	if M < 6 || len(botC.Loop) != M {
		return
	}

	topY := topC.Y + deltaY
	botY := botC.Y + deltaY

	for k := range res.RadialReps {
		rot := float64(k)*2.0*math.Pi/float64(res.RadialReps) + deltaTheta
		cosK := math.Cos(rot)
		sinK := math.Sin(rot)

		topPts := make([]pt2D, M)
		botPts := make([]pt2D, M)
		for i := range M {
			tp := topC.Loop[i]
			bp := botC.Loop[i]
			topPts[i] = pt2D{X: tp.X*cosK - tp.Z*sinK, Z: tp.X*sinK + tp.Z*cosK}
			botPts[i] = pt2D{X: bp.X*cosK - bp.Z*sinK, Z: bp.X*sinK + bp.Z*cosK}
		}

		// 1. Smooth side wall vertices (top ring + bottom ring).
		wallTopIdx := len(geom.Vertices)
		for i := range M {
			geom.Vertices = append(geom.Vertices, (&threejs.Vector3{
				Name: fmt.Sprintf("%s Leg Wall Top k%d %d", geomName, k, i),
				X:    topPts[i].X, Y: topY, Z: topPts[i].Z,
			}).Stage(stage))
		}
		wallBotIdx := len(geom.Vertices)
		for i := range M {
			geom.Vertices = append(geom.Vertices, (&threejs.Vector3{
				Name: fmt.Sprintf("%s Leg Wall Bot k%d %d", geomName, k, i),
				X:    botPts[i].X, Y: botY, Z: botPts[i].Z,
			}).Stage(stage))
		}

		for i := range M {
			nextI := (i + 1) % M
			tI := wallTopIdx + i
			tNext := wallTopIdx + nextI
			bI := wallBotIdx + i
			bNext := wallBotIdx + nextI

			geom.Faces = append(geom.Faces,
				(&threejs.Triangle{
					Name: fmt.Sprintf("%s Leg Wall T1 k%d %d", geomName, k, i),
					V1:   bI, V2: tI, V3: tNext,
				}).Stage(stage),
				(&threejs.Triangle{
					Name: fmt.Sprintf("%s Leg Wall T2 k%d %d", geomName, k, i),
					V1:   bI, V2: tNext, V3: bNext,
				}).Stage(stage),
			)
		}

		// 2. Separate top cap & bottom cap vertex rings for sharp crease normals,
		// sharing exact (X, Y, Z) coordinates with wall rings for watertight STL.
		capTopIdx := len(geom.Vertices)
		for i := range M {
			geom.Vertices = append(geom.Vertices, (&threejs.Vector3{
				Name: fmt.Sprintf("%s Leg Cap Top k%d %d", geomName, k, i),
				X:    topPts[i].X, Y: topY, Z: topPts[i].Z,
			}).Stage(stage))
		}
		appendCrescentCapFaces(stage, geom, fmt.Sprintf("%s Leg Top Cap k%d", geomName, k), capTopIdx, topC.Nb, topC.Nf, true)

		capBotIdx := len(geom.Vertices)
		for i := range M {
			geom.Vertices = append(geom.Vertices, (&threejs.Vector3{
				Name: fmt.Sprintf("%s Leg Cap Bot k%d %d", geomName, k, i),
				X:    botPts[i].X, Y: botY, Z: botPts[i].Z,
			}).Stage(stage))
		}
		appendCrescentCapFaces(stage, geom, fmt.Sprintf("%s Leg Bot Cap k%d", geomName, k), capBotIdx, botC.Nb, botC.Nf, false)
	}
}

func buildGrasshopperSeatGeom(
	stage *threejs.Stage,
	geomName string,
	res *grasshopperContoursResult,
	deltaTheta float64,
	deltaY float64,
) *threejs.BufferGeometry {
	geom := (&threejs.BufferGeometry{Name: geomName}).Stage(stage)
	appendGrasshopperSeatToGeom(stage, geom, geomName, res, deltaTheta, deltaY)
	return geom
}

func buildGrasshopperLegsGeom(
	stage *threejs.Stage,
	geomName string,
	res *grasshopperContoursResult,
	deltaTheta float64,
	deltaY float64,
) *threejs.BufferGeometry {
	geom := (&threejs.BufferGeometry{Name: geomName}).Stage(stage)
	appendGrasshopperLegsToGeom(stage, geom, geomName, res, deltaTheta, deltaY)
	return geom
}

func buildGrasshopperSeatAndLegsGeom(
	stage *threejs.Stage,
	geomName string,
	res *grasshopperContoursResult,
	deltaTheta float64,
	deltaY float64,
) *threejs.BufferGeometry {
	geom := (&threejs.BufferGeometry{Name: geomName}).Stage(stage)
	appendGrasshopperSeatToGeom(stage, geom, geomName, res, deltaTheta, deltaY)
	appendGrasshopperLegsToGeom(stage, geom, geomName, res, deltaTheta, deltaY)
	return geom
}
