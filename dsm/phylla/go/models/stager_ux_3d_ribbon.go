package models

import (
	"fmt"
	"log"
	"math"

	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func (stager *Stager) generateRibbonMesh(
	h int,
	stackHeight int,
	totalThetaOffset float64,
	namePrefix string,
	plant *PlantAbstract,
	checkedDiagram *PlantDiagram,
	localBottomCurve *threejs.Curve,
	localTopCurve *threejs.Curve,
	dy float64,
	thickness float64,
	globalR float64,
	canvas *threejs.Canvas,
) {
	hasAlternatingRingColors := false
	offsetKeyX := 0.0
	widthKey := 0.0
	offsetKeyY := 0.0
	heightKey := 0.0
	radialRepetitions := 1
	transparency := 0.0
	trajOffsetXRel := 0.0
	trajOffsetYRel := 0.0
	rotRatio := 0.0
	if plant.VaseAbstract != nil {
		vase := plant.VaseAbstract
		hasAlternatingRingColors = vase.HasAlternatingRingColors
		offsetKeyX = vase.OffsetKeyX
		widthKey = vase.WidthKey
		offsetKeyY = vase.OffsetKeyY
		heightKey = vase.HeightKey
		radialRepetitions = vase.RadialRepetitions
		transparency = vase.Transparency
		trajOffsetXRel = vase.RelativeTrajectoryOffsetX
		trajOffsetYRel = vase.RelativeTrajectoryOffsetY
		rotRatio = vase.RotationRatio
	}

	japanesePaperColor := "#fdf6e3" // Off-white cream color for Washi paper
	if hasAlternatingRingColors && h%2 != 0 {
		japanesePaperColor = "#8d6e63" // Warm rose brown / wood paper color for alternating rings
	}

	// ============================================================================
	// FIRST PASS: Vertex Generation & Elevation Levels
	// ============================================================================
	// We generate the 3D meshes by manually constructing Three.js BufferGeometry.
	// To achieve perfectly smooth continuous shading (Gouraud/Phong shading),
	// we MUST share vertices between adjacent quad faces.
	// Therefore, we first generate a single continuous 2D grid of vertices,
	// and later (in the second pass) we define faces by connecting these shared vertices by index.
	//
	// For every horizontal point `i` along the ribbon curve, we generate 4 vertically
	// stacked vertices (creating 3 horizontal bands). The y-coordinates for these 4 levels are:
	//   - Level 0 (y0): The absolute bottom edge of the ribbon (yBase)
	//   - Level 1 (y1): Identical to y0 (legacy hole logic removed)
	//   - Level 2 (y2): Identical to y3 (legacy hole logic removed)
	//   - Level 3 (y3): The absolute top edge of the ribbon (yBaseTop)

	// geomInner is the inner surface of the ribbon
	geomInner := (&threejs.BufferGeometry{Name: namePrefix + " Inner BufferGeometry"}).Stage(stager.threejsStage)
	// geomOuter is the outer surface of the ribbon
	geomOuter := (&threejs.BufferGeometry{Name: namePrefix + " Outer BufferGeometry"}).Stage(stager.threejsStage)
	geomHoleWalls := (&threejs.BufferGeometry{Name: namePrefix + " HoleWalls BufferGeometry"}).Stage(stager.threejsStage)

	var bottomEdges, topEdges [][2]*threejs.Vector3

	hasHole := false
	inHoleArr := make([]bool, len(localBottomCurve.Points))
	var y_bottom_abs, y_top_abs float64

	if !checkedDiagram.VaseDiagram.IsHiddenKeyHole3DShape && plant.KeyHoleShape != nil && globalR > 0 && h != 0 {
		x_left := offsetKeyX - widthKey/2.0
		x_right := offsetKeyX + widthKey/2.0
		y_bottom_abs = offsetKeyY - heightKey/2.0 + dy
		y_top_abs = offsetKeyY + heightKey/2.0 + dy

		hasHole = true

		numPointsPerRep := len(localBottomCurve.Points) / radialRepetitions
		if numPointsPerRep == 0 {
			numPointsPerRep = 1
		}

		unwrapAngles := func(curve *threejs.Curve) []float64 {
			angles := make([]float64, len(curve.Points))
			if len(curve.Points) == 0 {
				return angles
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

		th_continuous := unwrapAngles(localBottomCurve)

		for i := 0; i < len(localBottomCurve.Points); i++ {
			k := i / numPointsPerRep
			if k >= radialRepetitions {
				k = radialRepetitions - 1
			}

			baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(radialRepetitions)
			th_left := x_left/globalR + totalThetaOffset + baseThetaOffset
			th_right := x_right/globalR + totalThetaOffset + baseThetaOffset

			diff_left := th_continuous[i] - th_left
			for diff_left > math.Pi {
				diff_left -= 2 * math.Pi
			}
			for diff_left < -math.Pi {
				diff_left += 2 * math.Pi
			}

			diff_right := th_continuous[i] - th_right
			for diff_right > math.Pi {
				diff_right -= 2 * math.Pi
			}
			for diff_right < -math.Pi {
				diff_right += 2 * math.Pi
			}

			if diff_left >= 0 && diff_right <= 0 {
				inHoleArr[i] = true
			}
		}
	}

	// Loop through every point along the curve and push the 4 vertical levels (Inner and Outer)
	for i := 0; i < len(localBottomCurve.Points) && i < len(localTopCurve.Points); i++ {
		p := localBottomCurve.Points[i]
		pTop := localTopCurve.Points[i]

		thetaBase := math.Atan2(p.Z, p.X)
		theta := thetaBase

		thetaBaseTop := math.Atan2(pTop.Z, pTop.X)
		thetaTop := thetaBaseTop

		yBase := p.Y + dy
		yBaseTop := pTop.Y + dy

		rBase := math.Sqrt(p.X*p.X + p.Z*p.Z)
		rOuter := rBase + thickness

		rBaseTop := math.Sqrt(pTop.X*pTop.X + pTop.Z*pTop.Z)
		rOuterTop := rBaseTop + thickness

		// The 4 heights
		y0 := yBase
		y1 := yBase
		y2 := yBaseTop
		y3 := yBaseTop

		if hasHole {
			y1 = math.Max(yBase, math.Min(yBaseTop, y_bottom_abs))
			y2 = math.Max(yBase, math.Min(yBaseTop, y_top_abs))
		}

		// inner 4 points
		for j, yVal := range []float64{y0, y1, y2, y3} {
			// interpolate theta and r based on y
			var ratio float64
			if y3 > y0 {
				ratio = (yVal - y0) / (y3 - y0)
			}
			th_interp := theta + ratio*(thetaTop-theta)
			r_interp := rBase + ratio*(rBaseTop-rBase)

			xIn := r_interp * math.Cos(th_interp)
			zIn := r_interp * math.Sin(th_interp)

			vIn := (&threejs.Vector3{
				Name: fmt.Sprintf("%s Inner i%d j%d", namePrefix, i, j),
				X:    xIn, Y: yVal, Z: zIn,
			}).Stage(stager.threejsStage)
			geomInner.Vertices = append(geomInner.Vertices, vIn)

			rOut_interp := rOuter + ratio*(rOuterTop-rOuter)
			xOut := rOut_interp * math.Cos(th_interp)
			zOut := rOut_interp * math.Sin(th_interp)

			vOut := (&threejs.Vector3{
				Name: fmt.Sprintf("%s Outer i%d j%d", namePrefix, i, j),
				X:    xOut, Y: yVal, Z: zOut,
			}).Stage(stager.threejsStage)
			geomOuter.Vertices = append(geomOuter.Vertices, vOut)
		}

		// for bottom/top face
		vBL := geomInner.Vertices[len(geomInner.Vertices)-4]
		vBR := geomOuter.Vertices[len(geomOuter.Vertices)-4]
		bottomEdges = append(bottomEdges, [2]*threejs.Vector3{vBL, vBR})

		vTL := geomInner.Vertices[len(geomInner.Vertices)-1]
		vTR := geomOuter.Vertices[len(geomOuter.Vertices)-1]
		topEdges = append(topEdges, [2]*threejs.Vector3{vTL, vTR})
	}

	// ============================================================================
	// SECOND PASS: Face (Quad) Generation
	// ============================================================================
	// The ribbon surface is vertically subdivided into 3 horizontal bands formed
	// by the 4 elevation levels (y0, y1, y2, y3) computed during the vertex pass.

	numPointsPerRep := len(localBottomCurve.Points) / radialRepetitions
	if numPointsPerRep == 0 {
		numPointsPerRep = 1
	}

	for i := 0; i < len(localBottomCurve.Points)-1; i++ {
		// Skip the bridging quad between separate radial repetitions
		if (i+1)%numPointsPerRep == 0 {
			continue
		}

		if !hasHole || !inHoleArr[i] {
			// Add full quads for all 3 bands
			stager.addQuad(geomInner, i, 0, true, "inner")
			stager.addQuad(geomInner, i, 1, true, "inner")
			stager.addQuad(geomInner, i, 2, true, "inner")

			stager.addQuad(geomOuter, i, 0, false, "outer")
			stager.addQuad(geomOuter, i, 1, false, "outer")
			stager.addQuad(geomOuter, i, 2, false, "outer")
		} else {
			// Hole segment: only add bottom band (0) and top band (2)
			stager.addQuad(geomInner, i, 0, true, "inner_below")
			stager.addQuad(geomInner, i, 2, true, "inner_above")

			stager.addQuad(geomOuter, i, 0, false, "outer_below")
			stager.addQuad(geomOuter, i, 2, false, "outer_above")

			// Hole walls
			vB1 := geomInner.Vertices[i*4+1]
			vB2 := geomOuter.Vertices[i*4+1]
			vB3 := geomInner.Vertices[(i+1)*4+1]
			vB4 := geomOuter.Vertices[(i+1)*4+1]
			stager.addWallQuad(geomHoleWalls, vB1, vB2, vB3, vB4, "bottom_wall", true)

			vT1 := geomInner.Vertices[i*4+2]
			vT2 := geomOuter.Vertices[i*4+2]
			vT3 := geomInner.Vertices[(i+1)*4+2]
			vT4 := geomOuter.Vertices[(i+1)*4+2]
			stager.addWallQuad(geomHoleWalls, vT1, vT2, vT3, vT4, "top_wall", false)
			if i == 0 || !inHoleArr[i-1] {
				stager.addWallQuad(geomHoleWalls, geomInner.Vertices[i*4+1], geomInner.Vertices[i*4+2], geomOuter.Vertices[i*4+1], geomOuter.Vertices[i*4+2], "left_wall", true)
			}
			if i == len(localBottomCurve.Points)-2 || !inHoleArr[i+1] {
				stager.addWallQuad(geomHoleWalls, geomInner.Vertices[(i+1)*4+1], geomInner.Vertices[(i+1)*4+2], geomOuter.Vertices[(i+1)*4+1], geomOuter.Vertices[(i+1)*4+2], "right_wall", false)
			}
		}
	}

	opacity := 1.0 - transparency
	if opacity < 0.0 {
		opacity = 0.0
	}
	if opacity > 1.0 {
		opacity = 1.0
	}

	// ============================================================================
	// THIRD PASS: Mesh Creation
	// ============================================================================
	// Finally, create the Three.js Mesh objects and add them to the scene.
	// The inner and outer surfaces are created with shared materials.
	// The hole walls are created with the same material.
	innerMesh := (&threejs.Mesh{
		Name:           namePrefix + " Inner Mesh",
		Position:       threejs.Position{X: 0, Y: 0, Z: 0},
		BufferGeometry: geomInner,
		MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
			Name:                 namePrefix + " Inner Material",
			MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: japanesePaperColor},
			Transparent:          true,
			Opacity:              opacity,
		}).Stage(stager.threejsStage),
	}).Stage(stager.threejsStage)

	outerMesh := (&threejs.Mesh{
		Name:           namePrefix + " Outer Mesh",
		Position:       threejs.Position{X: 0, Y: 0, Z: 0},
		BufferGeometry: geomOuter,
		MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
			Name:                 namePrefix + " Outer Material",
			MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: japanesePaperColor},
			Transparent:          true,
			Opacity:              opacity,
		}).Stage(stager.threejsStage),
	}).Stage(stager.threejsStage)

	canvas.Meshs = append(canvas.Meshs, innerMesh, outerMesh)

	if hasHole {
		holeWallsMesh := (&threejs.Mesh{
			Name:           namePrefix + " HoleWalls Mesh",
			Position:       threejs.Position{X: 0, Y: 0, Z: 0},
			BufferGeometry: geomHoleWalls,
			MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
				Name:                 namePrefix + " HoleWalls Material",
				MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: japanesePaperColor},
				Transparent:          true,
				Opacity:              opacity,
			}).Stage(stager.threejsStage),
		}).Stage(stager.threejsStage)
		canvas.Meshs = append(canvas.Meshs, holeWallsMesh)
	}

	bottomFace := stager.createFaceMesh(namePrefix+" Bottom", japanesePaperColor, bottomEdges, false, transparency)
	topFace := stager.createFaceMesh(namePrefix+" Top", japanesePaperColor, topEdges, true, transparency)
	canvas.Meshs = append(canvas.Meshs, bottomFace, topFace)

	outerRadius := 0.1
	innerRadius := outerRadius * 0.85
	bambooColor := "#4a3623"

	if !checkedDiagram.VaseDiagram.IsHiddenTorusEdge3DShape {
		canvas.Meshs = append(canvas.Meshs,
			stager.createTorusEdgeMesh(namePrefix+" BottomInner", bambooColor, bottomEdges, true, innerRadius),
			stager.createTorusEdgeMesh(namePrefix+" BottomOuter", bambooColor, bottomEdges, false, outerRadius),
			stager.createTorusEdgeMesh(namePrefix+" TopInner", bambooColor, topEdges, true, innerRadius),
			stager.createTorusEdgeMesh(namePrefix+" TopOuter", bambooColor, topEdges, false, outerRadius),
		)
	}

	if !checkedDiagram.VaseDiagram.IsHiddenPointsAndLines3DShape && h < stackHeight-1 && (plant.ChosenP1P2PairShape != nil || plant.PxShape != nil) {
		var p1x, p1y, p2x, p2y, pxx, pxy float64
		hasP1P2 := false
		if plant.ChosenP1P2PairShape != nil {
			p1x, p1y = plant.ChosenP1P2PairShape.P1X, plant.ChosenP1P2PairShape.P1Y
			p2x, p2y = plant.ChosenP1P2PairShape.P2X, plant.ChosenP1P2PairShape.P2Y
			pxx, pxy = plant.ChosenP1P2PairShape.PxX, plant.ChosenP1P2PairShape.PxY
			hasP1P2 = true
		} else if plant.PxShape != nil {
			pxx, pxy = plant.PxShape.X, plant.PxShape.Y
		}

		// Recompute Px for layer h based on step h+1's specific rotation ratio
		if plant.StackOfGrowthCurve2DRibbon != nil && len(plant.StackOfGrowthCurve2DRibbon.StackGrowthCurve2DRibbonStartShapes) > 0 && plant.RhombusStuff != nil && plant.RhombusStuff.PlantCircumferenceShape != nil {
			baseShape := plant.StackOfGrowthCurve2DRibbon.StackGrowthCurve2DRibbonStartShapes[0]
			circLen := plant.RhombusStuff.PlantCircumferenceShape.Length
			trajOffsetX := trajOffsetXRel * circLen
			trajOffsetY := trajOffsetYRel * circLen

			var r_h1 float64
			if stackHeight > 1 {
				numSteps := stackHeight - 1
				totalProgress := rotRatio * float64(numSteps)
				kStep := float64(numSteps - h)
				if totalProgress >= kStep {
					r_h1 = 1.0
				} else if totalProgress <= kStep-1.0 {
					r_h1 = 0.0
				} else {
					r_h1 = totalProgress - (kStep - 1.0)
				}
			} else {
				r_h1 = rotRatio
			}

			_, dyStep, currentDXStep := ComputePartiallyGrowthCurveDYForRatio(plant, r_h1)
			pxx = baseShape.BottomStartX + currentDXStep + trajOffsetX
			pxy = baseShape.BottomStartY + dyStep + trajOffsetY
		}

		rSurf := globalR

		sphereRad := thickness * 0.4
		if sphereRad < 0.3 {
			sphereRad = 0.3
		}

		for rep := 0; rep < radialRepetitions; rep++ {
			baseThetaOffset := float64(rep) * 2.0 * math.Pi / float64(radialRepetitions)
			currentThetaOffset := totalThetaOffset + baseThetaOffset

			get3DPt := func(ptX, ptY float64, ptName string) *threejs.Vector3 {
				th := ptX/globalR + currentThetaOffset
				return (&threejs.Vector3{
					Name: fmt.Sprintf("%s %s h%d r%d", ptName, namePrefix, h, rep),
					X:    rSurf * math.Cos(th),
					Y:    ptY + dy,
					Z:    rSurf * math.Sin(th),
				}).Stage(stager.threejsStage)
			}

			createPointSphere := func(ptName string, color string, vec *threejs.Vector3) *threejs.Mesh {
				return (&threejs.Mesh{
					Name: fmt.Sprintf("Sphere %s %s h%d r%d", ptName, namePrefix, h, rep),
					Position: threejs.Position{
						X: vec.X,
						Y: vec.Y,
						Z: vec.Z,
					},
					SphereGeometry: (&threejs.SphereGeometry{
						Name:           fmt.Sprintf("SphereGeom %s %s h%d r%d", ptName, namePrefix, h, rep),
						Radius:         sphereRad,
						WidthSegments:  16,
						HeightSegments: 16,
					}).Stage(stager.threejsStage),
					MeshMaterialBasic: (&threejs.MeshMaterialBasic{
						Name:                 fmt.Sprintf("Material %s %s h%d r%d", ptName, namePrefix, h, rep),
						MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: color},
					}).Stage(stager.threejsStage),
				}).Stage(stager.threejsStage)
			}

			createPairTube := func(lineName string, color string, pA, pB *threejs.Vector3) *threejs.Mesh {
				crv := (&threejs.Curve{
					Name:   fmt.Sprintf("Curve %s %s h%d r%d", lineName, namePrefix, h, rep),
					Points: []*threejs.Vector3{pA, pB},
				}).Stage(stager.threejsStage)

				tGeom := (&threejs.TubeGeometry{
					Name:            fmt.Sprintf("TubeGeom %s %s h%d r%d", lineName, namePrefix, h, rep),
					Path:            crv,
					TubularSegments: 8,
					Radius:          sphereRad * 0.25,
					RadialSegments:  8,
					Closed:          false,
				}).Stage(stager.threejsStage)

				return (&threejs.Mesh{
					Name:         fmt.Sprintf("TubeMesh %s %s h%d r%d", lineName, namePrefix, h, rep),
					Position:     threejs.Position{X: 0, Y: 0, Z: 0},
					TubeGeometry: tGeom,
					MeshMaterialBasic: (&threejs.MeshMaterialBasic{
						Name:                 fmt.Sprintf("Material %s %s h%d r%d", lineName, namePrefix, h, rep),
						MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: color},
					}).Stage(stager.threejsStage),
				}).Stage(stager.threejsStage)
			}

			vPx_3d := get3DPt(pxx, pxy, "Px")
			sPx := createPointSphere("Px", "purple", vPx_3d)
			canvas.Meshs = append(canvas.Meshs, sPx)

			if hasP1P2 {
				vP1_3d := get3DPt(p1x, p1y, "P1")
				vP2_3d := get3DPt(p2x, p2y, "P2")

				sP1 := createPointSphere("P1", "red", vP1_3d)
				sP2 := createPointSphere("P2", "#8d6e63", vP2_3d)

				tP1Px := createPairTube("P1-Px", "purple", vP1_3d, vPx_3d)
				tP2Px := createPairTube("P2-Px", "purple", vP2_3d, vPx_3d)

				canvas.Meshs = append(canvas.Meshs, sP1, sP2, tP1Px, tP2Px)

				if rep == 0 {
					dx1 := vP1_3d.X - vPx_3d.X
					dy1 := vP1_3d.Y - vPx_3d.Y
					dz1 := vP1_3d.Z - vPx_3d.Z
					distP1Px_3d := math.Sqrt(dx1*dx1 + dy1*dy1 + dz1*dz1)

					dx2 := vP2_3d.X - vPx_3d.X
					dy2 := vP2_3d.Y - vPx_3d.Y
					dz2 := vP2_3d.Z - vPx_3d.Z
					distP2Px_3d := math.Sqrt(dx2*dx2 + dy2*dy2 + dz2*dz2)

					distSum_3d := distP1Px_3d + distP2Px_3d

					if h == 0 {
						log.Printf("[3D Distance] %s (Layer %d) | P1-Px: %.4f, P2-Px: %.4f | 3D Sum: %.4f",
							namePrefix, h, distP1Px_3d, distP2Px_3d, distSum_3d)
					}
				}
			}
		}
	}
}


