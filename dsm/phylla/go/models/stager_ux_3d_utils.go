package models

import (
	"fmt"
	"math"

	"strconv"

	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func (stager *Stager) addLights(canvas *threejs.Canvas) {
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
}

func (stager *Stager) preserveCamera(hasPreservedCamera bool, preservedFov float64, canvas *threejs.Canvas, preservedX float64, preservedY float64, preservedZ float64, preservedTargetX float64, preservedTargetY float64, preservedTargetZ float64, checkedDiagram *PlantDiagram, globalR float64) {
	if hasPreservedCamera {
		if preservedFov == 0 {
			preservedFov = 50
		}
		canvas.Camera = (&threejs.Camera{
			Name: "Camera",
			Position: threejs.Position{
				X: preservedX,
				Y: preservedY,
				Z: preservedZ,
			},
			TargetX: preservedTargetX,
			TargetY: preservedTargetY,
			TargetZ: preservedTargetZ,
			Fov:     preservedFov,
		}).Stage(stager.threejsStage)
	} else if checkedDiagram != nil && checkedDiagram.Rendered3DShape != nil {
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

	canvas.Camera.OnUpdate = func(updatedCamera *threejs.Camera) {
		if checkedDiagram != nil && checkedDiagram.Rendered3DShape != nil {
			checkedDiagram.Rendered3DShape.ViewX = updatedCamera.X
			checkedDiagram.Rendered3DShape.ViewY = updatedCamera.Y
			checkedDiagram.Rendered3DShape.ViewZ = updatedCamera.Z
			checkedDiagram.Rendered3DShape.TargetX = updatedCamera.TargetX
			checkedDiagram.Rendered3DShape.TargetY = updatedCamera.TargetY
			checkedDiagram.Rendered3DShape.TargetZ = updatedCamera.TargetZ
			checkedDiagram.Rendered3DShape.Fov = updatedCamera.Fov
			stager.stage.CommitWithSuspendedCallbacks()
			stager.ux_slider()
		}
	}
}

func (*Stager) computeGlobalRadius(plant *Plant) (globalR float64) {
	circumference := 10.0
	if plant.RhombusStuff.PlantCircumferenceShape.Length > 0 {
		circumference = plant.RhombusStuff.PlantCircumferenceShape.Length
	} else if pGrid := plant.PerpendicularVectorGrid; len(pGrid.PerpendicularVectors) > 0 {
		first := pGrid.PerpendicularVectors[0]
		last := pGrid.PerpendicularVectors[len(pGrid.PerpendicularVectors)-1]
		circumference = last.StartX - first.StartX
	}
	if circumference <= 0 {
		circumference = 10.0
	}
	threeDModulo := plant.RadialRepetitions
	if threeDModulo < 1 {
		threeDModulo = 1
	}
	globalR = circumference * float64(threeDModulo) / (2 * math.Pi)
	return globalR
}
func (stager *Stager) addFloorTiles(floorMinY float64, plant *Plant, globalR float64, canvas *threejs.Canvas) {
	if floorMinY == math.MaxFloat64 {
		floorMinY = 0.0
	} else {
		thickness := plant.RelativeVerticalThickness * plant.RhombusSideLength
		if thickness == 0 {
			thickness = 5.0
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
}

func (stager *Stager) appendArcPoints(targetCurve *threejs.Curve, x1, y1, x2, y2, r float64, sweepFlag bool, largeArcFlag bool, globalR float64, floorMinY *float64) {
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
		if i == 0 && len(targetCurve.Points) > 0 {
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

		if y3d < *floorMinY {
			*floorMinY = y3d
		}

		vec := (&threejs.Vector3{
			Name: fmt.Sprintf("Base Point %d", len(targetCurve.Points)),
			X:    x3d,
			Y:    y3d,
			Z:    z3d,
		}).Stage(stager.threejsStage)
		targetCurve.Points = append(targetCurve.Points, vec)
	}
}

func (stager *Stager) createFaceMesh(faceName string, color string, edges [][2]*threejs.Vector3, reverseWinding bool, transparency float64) *threejs.Mesh {
	geom := (&threejs.BufferGeometry{
		Name: fmt.Sprintf("%s BufferGeometry", faceName),
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

	opacity := 1.0 - transparency
	if opacity < 0.0 {
		opacity = 0.0
	}
	if opacity > 1.0 {
		opacity = 1.0
	}

	return (&threejs.Mesh{
		Name:           fmt.Sprintf("%s Mesh", faceName),
		Position:       threejs.Position{X: 0, Y: 0, Z: 0},
		BufferGeometry: geom,
		MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
			Name:                 fmt.Sprintf("%s Material", faceName),
			MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: color},
			Transparent:          true,
			Opacity:              opacity,
		}).Stage(stager.threejsStage),
	}).Stage(stager.threejsStage)
}

func (stager *Stager) createTorusEdgeMesh(name string, color string, edges [][2]*threejs.Vector3, useLeft bool, tubeRadius float64) *threejs.Mesh {
	curve := (&threejs.Curve{
		Name: "Curve " + name,
	}).Stage(stager.threejsStage)

	for i := 0; i < len(edges); i++ {
		p := edges[i][0]
		if !useLeft {
			p = edges[i][1]
		}
		curve.Points = append(curve.Points, (&threejs.Vector3{
			Name: "CurvePoint " + name + " " + strconv.Itoa(i),
			X:    p.X,
			Y:    p.Y,
			Z:    p.Z,
		}).Stage(stager.threejsStage))
	}

	tubeGeometry := (&threejs.TubeGeometry{
		Name:            "TubeGeom " + name,
		Path:            curve,
		TubularSegments: 500,
		Radius:          tubeRadius,
		RadialSegments:  8,
		Closed:          false,
	}).Stage(stager.threejsStage)

	return (&threejs.Mesh{
		Name:              "TubeMesh " + name,
		Position:          threejs.Position{X: 0, Y: 0, Z: 0},
		TubeGeometry:      tubeGeometry,
		MeshMaterialBasic: (&threejs.MeshMaterialBasic{Name: name + " Material", MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: color}}).Stage(stager.threejsStage),
	}).Stage(stager.threejsStage)
}

func (stager *Stager) createKeyHole3DTubeMesh(tubeName string, pA, pB *threejs.Vector3, tubeRadius float64) *threejs.Mesh {
	crv := (&threejs.Curve{
		Name:   fmt.Sprintf("Curve %s", tubeName),
		Points: []*threejs.Vector3{pA, pB},
	}).Stage(stager.threejsStage)

	tGeom := (&threejs.TubeGeometry{
		Name:            fmt.Sprintf("TubeGeom %s", tubeName),
		Path:            crv,
		TubularSegments: 8,
		Radius:          tubeRadius,
		RadialSegments:  8,
		Closed:          false,
	}).Stage(stager.threejsStage)

	return (&threejs.Mesh{
		Name:         fmt.Sprintf("TubeMesh %s", tubeName),
		Position:     threejs.Position{X: 0, Y: 0, Z: 0},
		TubeGeometry: tGeom,
		MeshMaterialBasic: (&threejs.MeshMaterialBasic{
			Name:                 fmt.Sprintf("Material %s", tubeName),
			MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "darkred"},
		}).Stage(stager.threejsStage),
	}).Stage(stager.threejsStage)
}

func (stager *Stager) get3DPtHK(ptX, ptY float64, ptName string, dx_h, dy_h, globalR, baseThetaOffset float64, h, k int) *threejs.Vector3 {
	th := (ptX+dx_h)/globalR + baseThetaOffset
	return (&threejs.Vector3{
		Name: fmt.Sprintf("KeyHole3D %s h%d k%d", ptName, h, k),
		X:    globalR * math.Cos(th),
		Y:    ptY + dy_h,
		Z:    globalR * math.Sin(th),
	}).Stage(stager.threejsStage)
}

func (stager *Stager) cloneAndRotateCurve(source *threejs.Curve, thetaOffset float64) *threejs.Curve {
	clone := (&threejs.Curve{
		Name: source.Name + fmt.Sprintf(" Rotated %.2f", thetaOffset),
	}).Stage(stager.threejsStage)

	for _, p := range source.Points {
		thetaBase := math.Atan2(p.Z, p.X)
		r := math.Sqrt(p.X*p.X + p.Z*p.Z)
		newTheta := thetaBase + thetaOffset

		newP := (&threejs.Vector3{
			Name: p.Name + " Rotated",
			X:    r * math.Cos(newTheta),
			Y:    p.Y,
			Z:    r * math.Sin(newTheta),
		}).Stage(stager.threejsStage)
		clone.Points = append(clone.Points, newP)
	}

	return clone
}
