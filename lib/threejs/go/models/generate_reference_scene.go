package models

import (
	"math"
	"strconv"
)

func GenerateReferenceScene(stage *Stage) {
	stage.Reset()

	canvas := (&Canvas{
		Name: "Singloton",
	}).Stage(stage)

	canvas.Camera = (&Camera{
		Name:     "Camera",
		Position: Position{X: 30, Y: 30, Z: 30},
		TargetX:  0,
		TargetY:  5,
		TargetZ:  0,
		Fov:      45,
	}).Stage(stage)

	canvas.AmbiantLight = (&AmbiantLight{
		Name: "Ambiant Light",
		LightAbstract: LightAbstract{
			Intensity: 0.3,
		},
	}).Stage(stage)

	dirLight1 := (&DirectionalLight{
		Name:     "Directional Light 1 (Key)",
		Position: Position{X: 15, Y: 20, Z: 15},
		LightAbstract: LightAbstract{
			Intensity: 1.2,
		},
		IsWithCastShadow: true,
	}).Stage(stage)

	dirLight2 := (&DirectionalLight{
		Name:     "Directional Light 2 (Fill)",
		Position: Position{X: -15, Y: 10, Z: -15},
		LightAbstract: LightAbstract{
			Intensity: 0.6,
		},
		IsWithCastShadow: false,
	}).Stage(stage)

	dirLight3 := (&DirectionalLight{
		Name:     "Directional Light 3 (Rim)",
		Position: Position{X: 0, Y: 30, Z: -20},
		LightAbstract: LightAbstract{
			Intensity: 0.8,
		},
		IsWithCastShadow: false,
	}).Stage(stage)

	canvas.DirectionalLights = append(canvas.DirectionalLights, dirLight1, dirLight2, dirLight3)

	tileSize := 10.0
	tilesCount := 10

	boxGeom := (&BoxGeometry{
		Name:           "Floor Tile Geometry",
		Width:          tileSize,
		Height:         0.1,
		Depth:          tileSize,
		WidthSegments:  1,
		HeightSegments: 1,
		DepthSegments:  1,
	}).Stage(stage)

	color1 := (&MeshMaterialBasic{Name: "TileColor1", MeshMaterialAbstract: MeshMaterialAbstract{Color: "white"}}).Stage(stage)
	color2 := (&MeshMaterialBasic{Name: "TileColor2", MeshMaterialAbstract: MeshMaterialAbstract{Color: "black"}}).Stage(stage)

	for i := 0; i < tilesCount; i++ {
		for j := 0; j < tilesCount; j++ {
			var mat *MeshMaterialBasic
			if (i+j)%2 == 0 {
				mat = color1
			} else {
				mat = color2
			}

			tileMesh := (&Mesh{
				Name: "Floor Tile " + strconv.Itoa(i) + "-" + strconv.Itoa(j),
				Position: Position{
					X: (float64(i) - float64(tilesCount)/2.0 + 0.5) * tileSize,
					Y: -0.05,
					Z: (float64(j) - float64(tilesCount)/2.0 + 0.5) * tileSize,
				},
				BoxGeometry:       boxGeom,
				MeshMaterialBasic: mat,
			}).Stage(stage)

			canvas.Meshs = append(canvas.Meshs, tileMesh)
		}
	}

	planeMesh := (&Mesh{
		Name:     "Wall",
		Position: Position{X: 0, Y: 0, Z: -30},
		PlaneGeometry: (&PlaneGeometry{
			Name:           "Wall Plane",
			Width:          200,
			Height:         200,
			WidthSegments:  1,
			HeightSegments: 1,
		}).Stage(stage),
		MeshMaterialBasic: (&MeshMaterialBasic{Name: "Gray", MeshMaterialAbstract: MeshMaterialAbstract{Color: "gray"}}).Stage(stage),
	}).Stage(stage)
	canvas.Meshs = append(canvas.Meshs, planeMesh)

	sphereMesh := (&Mesh{
		Name:     "Sphere",
		Position: Position{X: 5, Y: 5, Z: -5},
		SphereGeometry: (&SphereGeometry{
			Name:           "Sphere Geometry",
			Radius:         2,
			WidthSegments:  32,
			HeightSegments: 16,
			PhiLength:      math.Pi * 2,
			ThetaLength:    math.Pi,
		}).Stage(stage),
		MeshMaterialBasic: (&MeshMaterialBasic{Name: "Cyan", MeshMaterialAbstract: MeshMaterialAbstract{Color: "cyan"}}).Stage(stage),
	}).Stage(stage)
	canvas.Meshs = append(canvas.Meshs, sphereMesh)

	curve := (&Curve{
		Name: "Helix Path",
	}).Stage(stage)

	turns := 4.0
	radius := 2.0
	height := 6.0
	for i := 0; i <= 200; i++ {
		t := float64(i) / 200.0
		angle := t * math.Pi * 2 * turns
		x := math.Cos(angle) * radius
		z := math.Sin(angle) * radius
		y := t * height

		vec := (&Vector3{
			Name: "Point " + strconv.Itoa(i),
			X:    x,
			Y:    y,
			Z:    z,
		}).Stage(stage)
		curve.Points = append(curve.Points, vec)
	}

	tubeGeometry := (&TubeGeometry{
		Name:            "Helix Tube Geometry",
		Path:            curve,
		TubularSegments: 200,
		Radius:          0.4,
		RadialSegments:  16,
		Closed:          false,
	}).Stage(stage)

	tubeMesh := (&Mesh{
		Name:              "Helix Mesh",
		Position:          Position{X: 0, Y: 3.5, Z: -10},
		TubeGeometry:      tubeGeometry,
		MeshMaterialBasic: (&MeshMaterialBasic{Name: "Silver", MeshMaterialAbstract: MeshMaterialAbstract{Color: "silver"}}).Stage(stage),
	}).Stage(stage)

	canvas.Meshs = append(canvas.Meshs, tubeMesh)

	extrudePath := (&Curve{
		Name: "Torus Like Curve",
	}).Stage(stage)

	radius = 10.0
	waves := 5.0
	amplitude := 2.0
	for i := 0; i <= 500; i++ {
		t := float64(i) / 500.0
		angle := t * math.Pi * 2
		x := math.Cos(angle) * radius
		z := math.Sin(angle) * radius
		y := math.Sin(angle*waves) * amplitude

		vec := (&Vector3{
			Name: "TorusPoint " + strconv.Itoa(i),
			X:    x,
			Y:    y,
			Z:    z,
		}).Stage(stage)
		extrudePath.Points = append(extrudePath.Points, vec)
	}

	thicknessTorus := 1.5 // 2 * size (0.75)
	eTorus := thicknessTorus * 0.05

	createFaceMesh := func(name string, color string, xMin, xMax, yMin, yMax float64) *Mesh {
		shape := (&Shape{
			Name: "Torus " + name + " Shape",
		}).Stage(stage)

		pts := [][2]float64{
			{xMin, yMin},
			{xMax, yMin},
			{xMax, yMax},
			{xMin, yMax},
			{xMin, yMin},
		}
		for i, p := range pts {
			v := (&Vector2{
				Name: "SquarePoint " + name + " " + strconv.Itoa(i),
				X:    p[0],
				Y:    p[1],
			}).Stage(stage)
			shape.Points = append(shape.Points, v)
		}

		extrudeGeom := (&ExtrudeGeometry{
			Name:        "Torus " + name + " Extrude Geometry",
			ExtrudePath: extrudePath,
			Shape:       shape,
			Steps:       500,
		}).Stage(stage)

		return (&Mesh{
			Name:            "Torus " + name + " Mesh",
			Position:        Position{X: 0, Y: 5, Z: 0},
			ExtrudeGeometry: extrudeGeom,
			MeshPhysicalMaterial: (&MeshPhysicalMaterial{
				Name:                 "Torus " + name + " Material",
				MeshMaterialAbstract: MeshMaterialAbstract{Color: color},
			}).Stage(stage),
		}).Stage(stage)
	}

	// Create 4 faces centered at 0,0
	half := thicknessTorus / 2.0
	bottomFace := createFaceMesh("Bottom", "#1f77b4", -half, half, -half, -half+eTorus)
	topFace := createFaceMesh("Top", "#d62728", -half, half, half-eTorus, half)
	innerFace := createFaceMesh("Inner", "#ff7f0e", -half, -half+eTorus, -half, half)
	outerFace := createFaceMesh("Outer", "#2ca02c", half-eTorus, half, -half, half)

	canvas.Meshs = append(canvas.Meshs, bottomFace, topFace, innerFace, outerFace)

	curvePhylla := (&Curve{
		Name: "Phylla Like Curve",
	}).Stage(stage)

	// Hardcoded arcs extracted from phylla's GrowthCurve2D
	type Arc struct {
		StartX, StartY, EndX, EndY, Radius float64
		SweepFlag, LargeArcFlag            bool
	}
	startArcs := []Arc{
		{0.000000, 0.000000, 91.766294, 39.735971, 100.000000, false, false},
		{183.532587, 79.471941, 263.828094, 19.867985, 99.999999, true, false},
		{344.123600, -39.735971, 435.889894, -0.000000, 100.000001, false, false},
		{527.656188, 39.735970, 607.951695, -19.867986, 99.999999, true, false},
		{688.247201, -79.471942, 780.013495, -39.735971, 100.000000, false, false},
	}
	endArcs := []Arc{
		{183.532587, 79.471941, 91.766294, 39.735971, 100.000000, false, false},
		{344.123600, -39.735971, 263.828094, 19.867985, 99.999999, true, false},
		{527.656188, 39.735970, 435.889894, -0.000000, 100.000001, false, false},
		{688.247201, -79.471942, 607.951695, -19.867986, 99.999999, true, false},
		{871.779788, 0.000000, 780.013495, -39.735971, 100.000000, false, false},
	}

	circumference := 871.779788
	globalR := circumference / (2 * math.Pi)
	
	// Scale factor to shrink down the 138-radius shape into a ~10-radius shape
	scale := 10.0 / globalR
	yOffset := 10.0 // Raise it above the first torus

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
			if i == 0 && len(curvePhylla.Points) > 0 {
				continue
			}

			t := float64(i) / float64(steps)
			angle := startAngle + t*(endAngle-startAngle)
			x2d := cx + r*math.Cos(angle)
			y2d := cy + r*math.Sin(angle)

			theta := x2d / globalR
			x3d := globalR * math.Cos(theta)
			z3d := globalR * math.Sin(theta)

			vec := (&Vector3{
				Name: "PhyllaPoint " + strconv.Itoa(len(curvePhylla.Points)),
				X:    x3d * scale,
				Y:    y2d * scale + yOffset,
				Z:    z3d * scale,
			}).Stage(stage)
			curvePhylla.Points = append(curvePhylla.Points, vec)
		}
	}

	for i := 0; i < len(startArcs); i++ {
		sa := startArcs[i]
		appendArcPoints(sa.StartX, sa.StartY, sa.EndX, sa.EndY, sa.Radius, !sa.SweepFlag, sa.LargeArcFlag)

		if i < len(endArcs) {
			ea := endArcs[i]
			appendArcPoints(ea.EndX, ea.EndY, ea.StartX, ea.StartY, ea.Radius, ea.SweepFlag, ea.LargeArcFlag)
		}
	}

	// Offset between GrowthCurve2D and TopGrowthCurve2D
	dx2d := 1.147078 * scale
	dy2d := 9.933993 * scale
	thicknessPhylla := 1.5 // A custom thickness scaled nicely for radius 10.0

	createFaceMeshPhylla := func(name string, color string, p1, p2, p3, p4 [2]float64) *Mesh {
		shape := (&Shape{
			Name: "Phylla Torus " + name + " Shape",
		}).Stage(stage)

		pts := [][2]float64{p1, p2, p3, p4, p1}
		for i, p := range pts {
			v := (&Vector2{
				Name: "Phylla SquarePoint " + name + " " + strconv.Itoa(i),
				X:    p[0],
				Y:    p[1],
			}).Stage(stage)
			shape.Points = append(shape.Points, v)
		}

		extrudeGeom := (&ExtrudeGeometry{
			Name:        "Phylla Torus " + name + " Extrude Geometry",
			ExtrudePath: curvePhylla,
			Shape:       shape,
			Steps:       len(curvePhylla.Points) * 2,
		}).Stage(stage)

		return (&Mesh{
			Name:            "Phylla Torus " + name + " Mesh",
			Position:        Position{X: 0, Y: 0, Z: 0},
			ExtrudeGeometry: extrudeGeom,
			MeshPhysicalMaterial: (&MeshPhysicalMaterial{
				Name:                 "Phylla Torus " + name + " Material",
				MeshMaterialAbstract: MeshMaterialAbstract{Color: color},
			}).Stage(stage),
		}).Stage(stage)
	}

	ePhylla := 0.1

	// Points of the skewed rectangle matching the exact GrowthCurve2D -> TopGrowthCurve2D offsets
	// Bottom-Left (GrowthCurve2D): 0, 0
	// Bottom-Right: thickness, 0
	// Top-Left (TopGrowthCurve2D): dx2d, dy2d
	// Top-Right: dx2d + thickness, dy2d

	pBL := [2]float64{0, 0}
	pBR := [2]float64{thicknessPhylla, 0}
	pTL := [2]float64{dx2d, dy2d}
	pTR := [2]float64{dx2d + thicknessPhylla, dy2d}

	// Bottom Face: thickness ePhylla
	bottomFacePhylla := createFaceMeshPhylla("Bottom", "magenta",
		[2]float64{pBL[0], pBL[1]},
		[2]float64{pBR[0], pBR[1]},
		[2]float64{pBR[0], pBR[1] + ePhylla},
		[2]float64{pBL[0], pBL[1] + ePhylla})

	topFacePhylla := createFaceMeshPhylla("Top", "yellow",
		[2]float64{pTL[0], pTL[1] - ePhylla},
		[2]float64{pTR[0], pTR[1] - ePhylla},
		[2]float64{pTR[0], pTR[1]},
		[2]float64{pTL[0], pTL[1]})

	innerFacePhylla := createFaceMeshPhylla("Inner", "cyan",
		[2]float64{pBL[0], pBL[1]},
		[2]float64{pBL[0] + ePhylla, pBL[1]},
		[2]float64{pTL[0] + ePhylla, pTL[1]},
		[2]float64{pTL[0], pTL[1]})

	outerFacePhylla := createFaceMeshPhylla("Outer", "lime",
		[2]float64{pBR[0] - ePhylla, pBR[1]},
		[2]float64{pBR[0], pBR[1]},
		[2]float64{pTR[0], pTR[1]},
		[2]float64{pTR[0] - ePhylla, pTR[1]})

	canvas.Meshs = append(canvas.Meshs, bottomFacePhylla, topFacePhylla, innerFacePhylla, outerFacePhylla)

	stage.Commit()
}
