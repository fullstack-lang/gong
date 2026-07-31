package models

import (
	"fmt"
	"log"
	"math"

	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func (stager *Stager) generateRibbonLayer(
	h int, // The vertical stack layer index (0 is bottom)
	dx float64, // Horizontal offset applied to the curve for this layer
	dy float64, // Vertical offset applied to the curve for this layer
	thetaOffset float64, // Angular rotation offset applied to this layer
	baseNamePrefix string, // Prefix for naming the generated 3D meshes
	plant *Plant, // The Plant domain object containing geometry parameters
	checkedDiagram *PlantDiagram, // The current UI diagram state and visibility flags
	curve *threejs.Curve, // The 3D points defining the bottom edge of the ribbon
	topCurve *threejs.Curve, // The 3D points defining the top edge of the ribbon
	thickness float64, // Thickness of the ribbon (inner to outer radius difference)
	globalR float64, // The base global radius of the cylindrical projection
	canvas *threejs.Canvas, // The Three.js canvas where generated meshes are appended
) {
	stackHeight := plant.StackHeight
	threeDModulo := plant.ThreeDModulo

	for k := 0; k < threeDModulo; k++ {
		baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(threeDModulo)

		namePrefix := fmt.Sprintf("%s Layer %d", baseNamePrefix, h)
		japanesePaperColor := "#fdf6e3" // Off-white cream color for Washi paper

		var hasHole bool
		var thL, thR, yB, yT float64
		if !checkedDiagram.IsHiddenKeyHole3DShape && plant.KeyHoleShape != nil {
			hasHole = true
			x_left := plant.OffsetKeyX - plant.WidthKey/2.0
			x_right := plant.OffsetKeyX + plant.WidthKey/2.0
			y_bottom := plant.OffsetKeyY - plant.HeightKey/2.0
			y_top := plant.OffsetKeyY + plant.HeightKey/2.0

			thL = (x_left+dx)/globalR + baseThetaOffset
			thR = (x_right+dx)/globalR + baseThetaOffset
			yB = y_bottom + dy
			yT = y_top + dy
		}

		// Precompute inHole for all segments
		inHoleArr := make([]bool, len(curve.Points)-1)
		if hasHole {
			for i := 0; i < len(curve.Points)-1; i++ {
				p1 := curve.Points[i]
				p2 := curve.Points[i+1]
				thetaBase := math.Atan2((p1.Z+p2.Z)/2.0, (p1.X+p2.X)/2.0)
				th_mid := thetaBase + thetaOffset + baseThetaOffset
				width := thR - thL
				for width < 0 {
					width += 2 * math.Pi
				}
				diff := math.Mod(th_mid-thL, 2*math.Pi)
				if diff < 0 {
					diff += 2 * math.Pi
				}
				if diff <= width {
					inHoleArr[i] = true
				}
			}
		}

		// We generate the 3D meshes by constructing BufferGeometry
		// To share vertices and get smooth shading, we add all vertices first, then construct faces.
		// There are 4 bands vertically for each point:
		// 0: yBase
		// 1: yB (if in hole, clamped)
		// 2: yT (if in hole, clamped)
		// 3: yBaseTop

		geomInner := (&threejs.BufferGeometry{Name: namePrefix + " Inner BufferGeometry"}).Stage(stager.threejsStage)
		geomOuter := (&threejs.BufferGeometry{Name: namePrefix + " Outer BufferGeometry"}).Stage(stager.threejsStage)
		geomHoleWalls := (&threejs.BufferGeometry{Name: namePrefix + " HoleWalls BufferGeometry"}).Stage(stager.threejsStage)

		var bottomEdges, topEdges [][2]*threejs.Vector3

		// First, add all vertices
		for i := 0; i < len(curve.Points) && i < len(topCurve.Points); i++ {
			p := curve.Points[i]
			pTop := topCurve.Points[i]

			thetaBase := math.Atan2(p.Z, p.X)
			theta := thetaBase + thetaOffset + baseThetaOffset

			thetaBaseTop := math.Atan2(pTop.Z, pTop.X)
			thetaTop := thetaBaseTop + thetaOffset + baseThetaOffset

			yBase := p.Y + dy
			yBaseTop := pTop.Y + dy

			rBase := math.Sqrt(p.X*p.X + p.Z*p.Z)
			rOuter := rBase + thickness

			rBaseTop := math.Sqrt(pTop.X*pTop.X + pTop.Z*pTop.Z)
			rOuterTop := rBaseTop + thickness

			// The 4 heights
			y0 := yBase
			y1 := yB
			if y1 < yBase {
				y1 = yBase
			}
			if y1 > yBaseTop {
				y1 = yBaseTop
			}

			y2 := yT
			if y2 < yBase {
				y2 = yBase
			}
			if y2 > yBaseTop {
				y2 = yBaseTop
			}

			y3 := yBaseTop

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
					Name: fmt.Sprintf("%s Inner k%d i%d j%d", namePrefix, k, i, j),
					X:    xIn, Y: yVal, Z: zIn,
				}).Stage(stager.threejsStage)
				geomInner.Vertices = append(geomInner.Vertices, vIn)

				rOut_interp := rOuter + ratio*(rOuterTop-rOuter)
				xOut := rOut_interp * math.Cos(th_interp)
				zOut := rOut_interp * math.Sin(th_interp)

				vOut := (&threejs.Vector3{
					Name: fmt.Sprintf("%s Outer k%d i%d j%d", namePrefix, k, i, j),
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



		for i := 0; i < len(curve.Points)-1; i++ {
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
				// Bottom wall (only if i > 0, because we assume hole doesn't touch ends normally)
				if i > 0 {
					vB1 := geomInner.Vertices[i*4]
					vB2 := geomOuter.Vertices[i*4]
					vB3 := geomInner.Vertices[(i+1)*4]
					vB4 := geomOuter.Vertices[(i+1)*4]
					stager.addWallQuad(geomHoleWalls, vB1, vB2, vB3, vB4, "bottom_wall", false)
				}

				// Top wall
				if i > 0 {
					vT1 := geomInner.Vertices[i*4+2]
					vT2 := geomOuter.Vertices[i*4+2]
					vT3 := geomInner.Vertices[(i+1)*4+2]
					vT4 := geomOuter.Vertices[(i+1)*4+2]
					stager.addWallQuad(geomHoleWalls, vT1, vT2, vT3, vT4, "top_wall", true)
				}

				// Left wall (only if previous segment was NOT in hole)
				if i == 0 || !inHoleArr[i-1] {
					w1 := geomInner.Vertices[i*4+1]
					w2 := geomInner.Vertices[i*4+2]
					w3 := geomOuter.Vertices[i*4+1]
					w4 := geomOuter.Vertices[i*4+2]
					stager.addWallQuad(geomHoleWalls, w1, w2, w3, w4, "left_wall", true)
				}

				// Right wall (only if next segment is NOT in hole)
				if i == len(curve.Points)-2 || !inHoleArr[i+1] {
					w1 := geomInner.Vertices[(i+1)*4+1]
					w2 := geomInner.Vertices[(i+1)*4+2]
					w3 := geomOuter.Vertices[(i+1)*4+1]
					w4 := geomOuter.Vertices[(i+1)*4+2]
					stager.addWallQuad(geomHoleWalls, w1, w2, w3, w4, "right_wall", false)
				}
			}
		}

		opacity := 1.0 - plant.Transparency
		if opacity < 0.0 {
			opacity = 0.0
		}
		if opacity > 1.0 {
			opacity = 1.0
		}

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

		bottomFace := stager.createFaceMesh(namePrefix+" Bottom", japanesePaperColor, bottomEdges, false, plant.Transparency)
		topFace := stager.createFaceMesh(namePrefix+" Top", japanesePaperColor, topEdges, true, plant.Transparency)
		canvas.Meshs = append(canvas.Meshs, bottomFace, topFace)

		outerRadius := 0.1
		innerRadius := outerRadius * 0.85
		bambooColor := "#4a3623"

		if !checkedDiagram.IsHiddenTorusEdge3DShape {
			canvas.Meshs = append(canvas.Meshs,
				stager.createTorusEdgeMesh(namePrefix+" BottomInner", bambooColor, bottomEdges, true, innerRadius),
				stager.createTorusEdgeMesh(namePrefix+" BottomOuter", bambooColor, bottomEdges, false, outerRadius),
				stager.createTorusEdgeMesh(namePrefix+" TopInner", bambooColor, topEdges, true, innerRadius),
				stager.createTorusEdgeMesh(namePrefix+" TopOuter", bambooColor, topEdges, false, outerRadius),
			)
		}

		if !checkedDiagram.IsHiddenPointsAndLines3DShape && h < stackHeight-1 && (plant.ChosenP1P2PairShape != nil || plant.PxShape != nil) {
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
				trajOffsetX := plant.RelativeTrajectoryOffsetX * circLen
				trajOffsetY := plant.RelativeTrajectoryOffsetY * circLen

				var r_h1 float64
				if stackHeight > 1 {
					numSteps := stackHeight - 1
					totalProgress := plant.RotationRatio * float64(numSteps)
					kStep := float64(numSteps - h)
					if totalProgress >= kStep {
						r_h1 = 1.0
					} else if totalProgress <= kStep-1.0 {
						r_h1 = 0.0
					} else {
						r_h1 = totalProgress - (kStep - 1.0)
					}
				} else {
					r_h1 = plant.RotationRatio
				}

				_, dyStep, currentDXStep := ComputePartiallyGrowthCurveDYForRatio(plant, r_h1)
				pxx = baseShape.BottomStartX + currentDXStep + trajOffsetX
				pxy = baseShape.BottomStartY + dyStep + trajOffsetY
			}

			rSurf := globalR

			get3DPt := func(ptX, ptY float64, ptName string) *threejs.Vector3 {
				th := (ptX+dx)/globalR + baseThetaOffset
				return (&threejs.Vector3{
					Name: fmt.Sprintf("%s %s k%d h%d", ptName, namePrefix, k, h),
					X:    rSurf * math.Cos(th),
					Y:    ptY + dy,
					Z:    rSurf * math.Sin(th),
				}).Stage(stager.threejsStage)
			}

			vPx_3d := get3DPt(pxx, pxy, "Px")

			sphereRad := thickness * 0.4
			if sphereRad < 0.3 {
				sphereRad = 0.3
			}

			createPointSphere := func(ptName string, color string, vec *threejs.Vector3) *threejs.Mesh {
				return (&threejs.Mesh{
					Name: fmt.Sprintf("Sphere %s %s k%d h%d", ptName, namePrefix, k, h),
					Position: threejs.Position{
						X: vec.X,
						Y: vec.Y,
						Z: vec.Z,
					},
					SphereGeometry: (&threejs.SphereGeometry{
						Name:           fmt.Sprintf("SphereGeom %s %s k%d h%d", ptName, namePrefix, k, h),
						Radius:         sphereRad,
						WidthSegments:  16,
						HeightSegments: 16,
					}).Stage(stager.threejsStage),
					MeshMaterialBasic: (&threejs.MeshMaterialBasic{
						Name:                 fmt.Sprintf("Material %s %s k%d h%d", ptName, namePrefix, k, h),
						MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: color},
					}).Stage(stager.threejsStage),
				}).Stage(stager.threejsStage)
			}

			createPairTube := func(lineName string, color string, pA, pB *threejs.Vector3) *threejs.Mesh {
				crv := (&threejs.Curve{
					Name:   fmt.Sprintf("Curve %s %s k%d h%d", lineName, namePrefix, k, h),
					Points: []*threejs.Vector3{pA, pB},
				}).Stage(stager.threejsStage)

				tGeom := (&threejs.TubeGeometry{
					Name:            fmt.Sprintf("TubeGeom %s %s k%d h%d", lineName, namePrefix, k, h),
					Path:            crv,
					TubularSegments: 8,
					Radius:          sphereRad * 0.25,
					RadialSegments:  8,
					Closed:          false,
				}).Stage(stager.threejsStage)

				return (&threejs.Mesh{
					Name:         fmt.Sprintf("TubeMesh %s %s k%d h%d", lineName, namePrefix, k, h),
					Position:     threejs.Position{X: 0, Y: 0, Z: 0},
					TubeGeometry: tGeom,
					MeshMaterialBasic: (&threejs.MeshMaterialBasic{
						Name:                 fmt.Sprintf("Material %s %s k%d h%d", lineName, namePrefix, k, h),
						MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: color},
					}).Stage(stager.threejsStage),
				}).Stage(stager.threejsStage)
			}

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

				dx1 := vP1_3d.X - vPx_3d.X
				dy1 := vP1_3d.Y - vPx_3d.Y
				dz1 := vP1_3d.Z - vPx_3d.Z
				distP1Px_3d := math.Sqrt(dx1*dx1 + dy1*dy1 + dz1*dz1)

				dx2 := vP2_3d.X - vPx_3d.X
				dy2 := vP2_3d.Y - vPx_3d.Y
				dz2 := vP2_3d.Z - vPx_3d.Z
				distP2Px_3d := math.Sqrt(dx2*dx2 + dy2*dy2 + dz2*dz2)

				distSum_3d := distP1Px_3d + distP2Px_3d

				if h == 0 && k == 0 {
					log.Printf("[3D Distance] %s (Layer %d, Rep %d) | P1-Px: %.4f, P2-Px: %.4f | 3D Sum: %.4f",
						baseNamePrefix, h, k, distP1Px_3d, distP2Px_3d, distSum_3d)
				}
			}
		}
	}
}
