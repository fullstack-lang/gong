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
				Name: "Camera",
				Position: threejs.Position{
					X: camDist,
					Y: camDist * 0.8,
					Z: camDist,
				},
				TargetY: R,
			}).Stage(stager.threejsStage)
		}
		tileSize := 10.0
		gridSize := 10

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
						Y: -0.05,
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



		// Torus generated from StartArcShapeV2Grid & EndArcShapeV2Grid
		if plant.StartArcShapeV2Grid != nil && plant.EndArcShapeV2Grid != nil &&
			len(plant.StartArcShapeV2Grid.StartArcShapesV2) > 0 {

			circumference := 10.0
			if plant.PlantCircumferenceShape != nil && plant.PlantCircumferenceShape.Length > 0 {
				circumference = plant.PlantCircumferenceShape.Length
			}
			R := circumference / (2 * math.Pi)

			thickness := plant.Thickness
			if thickness == 0 {
				thickness = 1.0
			}
			e := thickness * 0.05 // thickness of the shell
			half := thickness / 2.0

			createShape := func(name string, xMin, xMax, yMin, yMax float64) *threejs.Shape {
				shape := (&threejs.Shape{
					Name: "Torus " + name + " Shape",
				}).Stage(stager.threejsStage)
				pts := [][2]float64{
					{xMin, yMin},
					{xMax, yMin},
					{xMax, yMax},
					{xMin, yMax},
					{xMin, yMin},
				}
				for i, p := range pts {
					v := (&threejs.Vector2{
						Name: fmt.Sprintf("SquarePoint %s %d", name, i),
						X:    p[0],
						Y:    p[1],
					}).Stage(stager.threejsStage)
					shape.Points = append(shape.Points, v)
				}
				return shape
			}

			bottomShape := createShape("Bottom", -half, half, -half, -half+e)
			topShape := createShape("Top", -half, half, half-e, half)
			innerShape := createShape("Inner", -half, -half+e, -half, half)
			outerShape := createShape("Outer", half-e, half, -half, half)

			createArcMeshes := func(x1, y1, x2, y2, r float64, sweepFlag bool, name string) {
				dx := (x1 - x2) / 2.0
				dy := (y1 - y2) / 2.0
				d2 := dx*dx + dy*dy
				var cx, cy float64
				if d2 == 0 || r*r < d2 {
					cx = (x1 + x2) / 2.0
					cy = (y1 + y2) / 2.0
					r = math.Sqrt(d2)
				} else {
					h := math.Sqrt(r*r/d2 - 1.0)
					if !sweepFlag {
						h = -h
					}
					cx = (x1+x2)/2.0 + h*dy
					cy = (y1+y2)/2.0 - h*dx
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

				curve := (&threejs.Curve{
					Name: name + " Curve",
				}).Stage(stager.threejsStage)

				steps := 50
				for i := 0; i <= steps; i++ {
					t := float64(i) / float64(steps)
					angle := startAngle + t*(endAngle-startAngle)
					x2d := cx + r*math.Cos(angle)
					y2d := cy + r*math.Sin(angle)

					theta := x2d / R
					x3d := R * math.Cos(theta)
					z3d := R * math.Sin(theta)
					y3d := y2d

					vec := (&threejs.Vector3{
						Name: fmt.Sprintf("%s Point %d", name, i),
						X:    x3d,
						Y:    y3d,
						Z:    z3d,
					}).Stage(stager.threejsStage)
					curve.Points = append(curve.Points, vec)
				}

				if len(curve.Points) > 1 {
					createFaceMesh := func(faceName string, color string, shape *threejs.Shape) *threejs.Mesh {
						extrudeGeom := (&threejs.ExtrudeGeometry{
							Name:        fmt.Sprintf("Torus %s %s Extrude", name, faceName),
							ExtrudePath: curve,
							Shape:       shape,
							Steps:       len(curve.Points) * 2,
						}).Stage(stager.threejsStage)

						return (&threejs.Mesh{
							Name:                 fmt.Sprintf("Torus %s %s Mesh", name, faceName),
							Position:             threejs.Position{X: 0, Y: 0, Z: 0},
							ExtrudeGeometry:      extrudeGeom,
							MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
								Name:                 fmt.Sprintf("Torus %s %s Material", name, faceName),
								MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: color},
							}).Stage(stager.threejsStage),
						}).Stage(stager.threejsStage)
					}

					bottomFace := createFaceMesh("Bottom", "#1f77b4", bottomShape)
					topFace := createFaceMesh("Top", "#d62728", topShape)
					innerFace := createFaceMesh("Inner", "#ff7f0e", innerShape)
					outerFace := createFaceMesh("Outer", "#2ca02c", outerShape)

					canvas.Meshs = append(canvas.Meshs, bottomFace, topFace, innerFace, outerFace)
				}
			}

			startArcs := plant.StartArcShapeV2Grid.StartArcShapesV2
			endArcs := plant.EndArcShapeV2Grid.EndArcShapesV2

			for i := 0; i < len(startArcs); i++ {
				sa := startArcs[i]
				createArcMeshes(sa.StartX, sa.StartY, sa.EndX, sa.EndY, sa.RadiusX, sa.SweepFlag, fmt.Sprintf("StartArc %d", i))

				if i < len(endArcs) {
					ea := endArcs[i]
					createArcMeshes(ea.StartX, ea.StartY, ea.EndX, ea.EndY, ea.RadiusX, ea.SweepFlag, fmt.Sprintf("EndArc %d", i))
				}
			}
		}

		stager.threejsStage.Commit()
	} else {
		threejs.GenerateReferenceScene(stager.threejsStage)
	}
}
