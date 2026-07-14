package models

import (
	"fmt"
	"math"
	"strconv"

	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func (stager *Stager) ux_3d_plant_diagram() {

	if true {
		stager.threejsStage.Reset()

		plant := stager.GetCurrentPlant()
		if plant == nil {
			stager.threejsStage.Commit()
			return
		}

		canvas := (&threejs.Canvas{
			Name: "Plant 3D Canvas",
		}).Stage(stager.threejsStage)

		var checkedDiagram *PlantDiagram
		for _, diagram := range plant.PlantDiagrams {
			if diagram.IsChecked {
				checkedDiagram = diagram
				break
			}
		}

		if checkedDiagram != nil && checkedDiagram.Rendered3DShape != nil {
			canvas.Camera = (&threejs.Camera{
				Name: "Camera",
				Position: threejs.Position{
					X: checkedDiagram.Rendered3DShape.ViewX,
					Y: checkedDiagram.Rendered3DShape.ViewY,
					Z: checkedDiagram.Rendered3DShape.ViewZ,
				},
			}).Stage(stager.threejsStage)
			circumference := 10.0
			if plant.PlantCircumferenceShape != nil && plant.PlantCircumferenceShape.Length > 0 {
				circumference = plant.PlantCircumferenceShape.Length
			} else if pGrid := plant.PerpendicularVectorGrid; pGrid != nil && len(pGrid.PerpendicularVectors) > 0 {
				first := pGrid.PerpendicularVectors[0]
				last := pGrid.PerpendicularVectors[len(pGrid.PerpendicularVectors)-1]
				circumference = last.StartX - first.StartX
			}
			if circumference <= 0 {
				circumference = 10.0
			}
			R := circumference / (2 * math.Pi)
			camDist := R * 2.5
			if camDist < 15 {
				camDist = 15
			}

			canvas.Camera = (&threejs.Camera{
				Name: "Default Camera",
				Position: threejs.Position{
					X: camDist,
					Y: camDist * 0.8,
					Z: camDist,
				},
				TargetY: R,
			}).Stage(stager.threejsStage)
		}

		// lights
		dirLight1 := (&threejs.DirectionalLight{
			Name:             "Directional Light 1 (Key)",
			Position:         threejs.Position{X: 15, Y: 20, Z: 15},
			LightAbstract:    threejs.LightAbstract{Intensity: 1.2},
			IsWithCastShadow: true,
		}).Stage(stager.threejsStage)

		dirLight2 := (&threejs.DirectionalLight{
			Name:             "Directional Light 2 (Fill)",
			Position:         threejs.Position{X: -15, Y: 10, Z: -15},
			LightAbstract:    threejs.LightAbstract{Intensity: 0.6},
			IsWithCastShadow: false,
		}).Stage(stager.threejsStage)

		dirLight3 := (&threejs.DirectionalLight{
			Name:             "Directional Light 3 (Rim)",
			Position:         threejs.Position{X: 0, Y: 30, Z: -20},
			LightAbstract:    threejs.LightAbstract{Intensity: 0.8},
			IsWithCastShadow: false,
		}).Stage(stager.threejsStage)

		canvas.DirectionalLights = append(canvas.DirectionalLights, dirLight1, dirLight2, dirLight3)

		ambiantLight := (&threejs.AmbiantLight{
			Name:          "Ambiant Light",
			LightAbstract: threejs.LightAbstract{Intensity: 0.3},
		}).Stage(stager.threejsStage)
		canvas.AmbiantLight = ambiantLight

		var globalR float64
		{
			circumference := 10.0
			if plant.PlantCircumferenceShape != nil && plant.PlantCircumferenceShape.Length > 0 {
				circumference = plant.PlantCircumferenceShape.Length
			} else if pGrid := plant.PerpendicularVectorGrid; pGrid != nil && len(pGrid.PerpendicularVectors) > 0 {
				first := pGrid.PerpendicularVectors[0]
				last := pGrid.PerpendicularVectors[len(pGrid.PerpendicularVectors)-1]
				circumference = last.StartX - first.StartX
			}
			if circumference <= 0 {
				circumference = 10.0
			}
			globalR = circumference / (2 * math.Pi)
		}

		if checkedDiagram != nil && checkedDiagram.Rendered3DShape != nil {
			canvas.Camera = (&threejs.Camera{
				Name: "Camera",
				Position: threejs.Position{
					X: checkedDiagram.Rendered3DShape.ViewX,
					Y: checkedDiagram.Rendered3DShape.ViewY,
					Z: checkedDiagram.Rendered3DShape.ViewZ,
				},
				TargetX: checkedDiagram.Rendered3DShape.TargetX,
				TargetY: checkedDiagram.Rendered3DShape.TargetY,
				TargetZ: checkedDiagram.Rendered3DShape.TargetZ,
				Fov:     checkedDiagram.Rendered3DShape.Fov,
			}).Stage(stager.threejsStage)
		} else {
			camDist := globalR * 2.5
			if camDist < 15 {
				camDist = 15
			}

			canvas.Camera = (&threejs.Camera{
				Name: "Camera",
				Position: threejs.Position{
					X: camDist,
					Y: camDist * 0.8,
					Z: camDist,
				},
				TargetY: globalR,
			}).Stage(stager.threejsStage)
		}

		floorMinY := math.MaxFloat64

		// Torus generated from GrowthCurve2D and TopGrowthCurve2D
		if plant.GrowthCurve2D != nil && plant.TopGrowthCurve2D != nil &&
			plant.GrowthCurve2D.StartArcShapeGrid != nil && len(plant.GrowthCurve2D.StartArcShapeGrid.StartArcShapes) > 0 &&
			plant.TopGrowthCurve2D.TopStartArcShapeGrid != nil && len(plant.TopGrowthCurve2D.TopStartArcShapeGrid.TopStartArcShapes) > 0 {

			gc := plant.GrowthCurve2D
			tgc := plant.TopGrowthCurve2D

			startArcs := gc.StartArcShapeGrid.StartArcShapes
			endArcs := gc.EndArcShapeGrid.EndArcShapes

			topStartArcs := tgc.TopStartArcShapeGrid.TopStartArcShapes

			// dx2d := topStartArcs[0].StartX - startArcs[0].StartX
			dy2d := topStartArcs[0].StartY - startArcs[0].StartY

			thickness := plant.RadialThickness
			if thickness == 0 {
				thickness = 5.0
			}

			// We will use BufferGeometry directly, no need for 2D shapes

			curve := (&threejs.Curve{
				Name: "Torus Continuous Curve",
			}).Stage(stager.threejsStage)

			appendArcPoints := func(x1, y1, x2, y2, r float64, sweepFlag bool, largeArcFlag bool) {
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
					// avoid duplicating the exact same point if it's not the first point of the curve
					if i == 0 && len(curve.Points) > 0 {
						continue
					}

					t := float64(i) / float64(steps)
					angle := startAngle + t*(endAngle-startAngle)
					x2d := cx + r*math.Cos(angle)
					y2d := cy + r*math.Sin(angle)

					theta := x2d / globalR
					x3d := globalR * math.Cos(theta)
					z3d := globalR * math.Sin(theta)
					y3d := y2d

					if y3d < floorMinY {
						floorMinY = y3d
					}

					vec := (&threejs.Vector3{
						Name: fmt.Sprintf("Point %d", len(curve.Points)),
						X:    x3d,
						Y:    y3d,
						Z:    z3d,
					}).Stage(stager.threejsStage)
					curve.Points = append(curve.Points, vec)
				}
			}

			for i := 0; i < len(startArcs); i++ {
				sa := startArcs[i]
				// Cartesian sweep is the inverse of SVG sweep due to Y-axis mirroring
				appendArcPoints(sa.StartX, sa.StartY, sa.EndX, sa.EndY, sa.RadiusX, !sa.SweepFlag, sa.LargeArcFlag)

				if i < len(endArcs) {
					ea := endArcs[i]
					// Cartesian sweep is !ea.SweepFlag.
					// Traversing backwards inverts it again, so we pass ea.SweepFlag.
					appendArcPoints(ea.EndX, ea.EndY, ea.StartX, ea.StartY, ea.RadiusX, ea.SweepFlag, ea.LargeArcFlag)
				}
			}

			if len(curve.Points) > 1 {
				createFaceMesh := func(faceName string, color string, edges [][2]*threejs.Vector3, reverseWinding bool) *threejs.Mesh {
					geom := (&threejs.BufferGeometry{
						Name: fmt.Sprintf("Torus Continuous %s BufferGeometry", faceName),
					}).Stage(stager.threejsStage)

					for i := 0; i < len(edges); i++ {
						p1_src := edges[i][0]
						p2_src := edges[i][1]

						p1 := (&threejs.Vector3{
							Name: fmt.Sprintf("%s %s %d", p1_src.Name, faceName, i),
							X:    p1_src.X,
							Y:    p1_src.Y,
							Z:    p1_src.Z,
						}).Stage(stager.threejsStage)

						p2 := (&threejs.Vector3{
							Name: fmt.Sprintf("%s %s %d", p2_src.Name, faceName, i),
							X:    p2_src.X,
							Y:    p2_src.Y,
							Z:    p2_src.Z,
						}).Stage(stager.threejsStage)

						geom.Vertices = append(geom.Vertices, p1, p2)

						if i < len(edges)-1 {
							idx := i * 2

							v1_t1, v2_t1, v3_t1 := idx, idx+1, idx+2
							v1_t2, v2_t2, v3_t2 := idx+1, idx+3, idx+2

							if reverseWinding {
								v2_t1, v3_t1 = v3_t1, v2_t1
								v2_t2, v3_t2 = v3_t2, v2_t2
							}

							t1 := (&threejs.Triangle{
								Name: fmt.Sprintf("T1 %d", i),
								V1:   v1_t1,
								V2:   v2_t1,
								V3:   v3_t1,
							}).Stage(stager.threejsStage)

							t2 := (&threejs.Triangle{
								Name: fmt.Sprintf("T2 %d", i),
								V1:   v1_t2,
								V2:   v2_t2,
								V3:   v3_t2,
							}).Stage(stager.threejsStage)

							geom.Faces = append(geom.Faces, t1, t2)
						}
					}

					return (&threejs.Mesh{
						Name:            fmt.Sprintf("Torus Continuous %s Mesh", faceName),
						Position:        threejs.Position{X: 0, Y: 0, Z: 0},
						BufferGeometry:  geom,
						MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
							Name:                 fmt.Sprintf("Torus Continuous %s Material", faceName),
							MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: color},
						}).Stage(stager.threejsStage),
					}).Stage(stager.threejsStage)
				}

				var bottomEdges, topEdges, innerEdges, outerEdges [][2]*threejs.Vector3

				for i := 0; i < len(curve.Points); i++ {
					p := curve.Points[i]
					
					theta := math.Atan2(p.Z, p.X)

					xBL, yBL, zBL := p.X, p.Y, p.Z
					xBR, yBR, zBR := p.X + thickness*math.Cos(theta), p.Y, p.Z + thickness*math.Sin(theta)
					
					xTL, yTL, zTL := p.X, p.Y + dy2d, p.Z
					xTR, yTR, zTR := p.X + thickness*math.Cos(theta), p.Y + dy2d, p.Z + thickness*math.Sin(theta)

					vBL := (&threejs.Vector3{Name: "BL", X: xBL, Y: yBL, Z: zBL}).Stage(stager.threejsStage)
					vBR := (&threejs.Vector3{Name: "BR", X: xBR, Y: yBR, Z: zBR}).Stage(stager.threejsStage)
					vTL := (&threejs.Vector3{Name: "TL", X: xTL, Y: yTL, Z: zTL}).Stage(stager.threejsStage)
					vTR := (&threejs.Vector3{Name: "TR", X: xTR, Y: yTR, Z: zTR}).Stage(stager.threejsStage)

					bottomEdges = append(bottomEdges, [2]*threejs.Vector3{vBL, vBR})
					topEdges = append(topEdges, [2]*threejs.Vector3{vTL, vTR})
					innerEdges = append(innerEdges, [2]*threejs.Vector3{vBL, vTL})
					outerEdges = append(outerEdges, [2]*threejs.Vector3{vBR, vTR})
				}

				bottomFace := createFaceMesh("Bottom", "#1f77b4", bottomEdges, false) // blue
				topFace := createFaceMesh("Top", "#d62728", topEdges, true) // red
				innerFace := createFaceMesh("Inner", "#ff7f0e", innerEdges, true) // orange
				outerFace := createFaceMesh("Outer", "#2ca02c", outerEdges, false) // green

				canvas.Meshs = append(canvas.Meshs, bottomFace, topFace, innerFace, outerFace)
			}
		}

		if floorMinY == math.MaxFloat64 {
			floorMinY = 0.0
		} else {
			thickness := plant.VerticalThickness
			if thickness == 0 {
				thickness = 1.0
			}
			floorMinY = floorMinY - (thickness / 2.0)
		}

		floorSize := globalR * 2.5
		if floorSize < 100 {
			floorSize = 100
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
					}).Stage(stager.threejsStage),
					MeshMaterialBasic: (&threejs.MeshMaterialBasic{
						Name:                 "Tile Material " + color,
						MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: color},
					}).Stage(stager.threejsStage),
				}).Stage(stager.threejsStage)

				canvas.Meshs = append(canvas.Meshs, tileMesh)
			}
		}

		stager.threejsStage.Commit()
	} else {
		threejs.GenerateReferenceScene(stager.threejsStage)
	}
}
