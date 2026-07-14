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
	half := thicknessTorus / 2.0

	var bottomEdgesWavy, topEdgesWavy, innerEdgesWavy, outerEdgesWavy [][2]*Vector3

	for i := 0; i <= 500; i++ {
		t := float64(i) / 500.0
		angle := t * math.Pi * 2
		x := math.Cos(angle) * radius
		z := math.Sin(angle) * radius
		y := math.Sin(angle*waves) * amplitude + 5.0 // Add the Y:5 position offset here

		// Stable coordinate system to prevent twisting:
		// Outward vector
		oX, oY, oZ := math.Cos(angle), 0.0, math.Sin(angle)
		// Up vector
		uX, uY, uZ := 0.0, 1.0, 0.0

		// Bottom Left (Inner Bottom)
		blX := x - half*oX - half*uX
		blY := y - half*oY - half*uY
		blZ := z - half*oZ - half*uZ

		// Bottom Right (Outer Bottom)
		brX := x + half*oX - half*uX
		brY := y + half*oY - half*uY
		brZ := z + half*oZ - half*uZ

		// Top Left (Inner Top)
		tlX := x - half*oX + half*uX
		tlY := y - half*oY + half*uY
		tlZ := z - half*oZ + half*uZ

		// Top Right (Outer Top)
		trX := x + half*oX + half*uX
		trY := y + half*oY + half*uY
		trZ := z + half*oZ + half*uZ

		vBL := (&Vector3{Name: "Wavy BL " + strconv.Itoa(i), X: blX, Y: blY, Z: blZ}).Stage(stage)
		vBR := (&Vector3{Name: "Wavy BR " + strconv.Itoa(i), X: brX, Y: brY, Z: brZ}).Stage(stage)
		vTL := (&Vector3{Name: "Wavy TL " + strconv.Itoa(i), X: tlX, Y: tlY, Z: tlZ}).Stage(stage)
		vTR := (&Vector3{Name: "Wavy TR " + strconv.Itoa(i), X: trX, Y: trY, Z: trZ}).Stage(stage)

		bottomEdgesWavy = append(bottomEdgesWavy, [2]*Vector3{vBL, vBR})
		topEdgesWavy = append(topEdgesWavy, [2]*Vector3{vTL, vTR})
		innerEdgesWavy = append(innerEdgesWavy, [2]*Vector3{vBL, vTL})
		outerEdgesWavy = append(outerEdgesWavy, [2]*Vector3{vBR, vTR})
	}

	createFaceMeshWavy := func(name string, color string, edges [][2]*Vector3, reverseWinding bool) *Mesh {
		geom := (&BufferGeometry{
			Name: "Wavy Torus " + name + " BufferGeometry",
		}).Stage(stage)

		for i := 0; i < len(edges); i++ {
			p1_src := edges[i][0]
			p2_src := edges[i][1]

			p1 := (&Vector3{
				Name: p1_src.Name + " " + name + " " + strconv.Itoa(i),
				X:    p1_src.X,
				Y:    p1_src.Y,
				Z:    p1_src.Z,
			}).Stage(stage)

			p2 := (&Vector3{
				Name: p2_src.Name + " " + name + " " + strconv.Itoa(i),
				X:    p2_src.X,
				Y:    p2_src.Y,
				Z:    p2_src.Z,
			}).Stage(stage)

			geom.Vertices = append(geom.Vertices, p1, p2)

			if i < len(edges)-1 {
				idx := i * 2
				v1_t1, v2_t1, v3_t1 := idx, idx+1, idx+2
				v1_t2, v2_t2, v3_t2 := idx+1, idx+3, idx+2

				if reverseWinding {
					v2_t1, v3_t1 = v3_t1, v2_t1
					v2_t2, v3_t2 = v3_t2, v2_t2
				}

				t1 := (&Triangle{
					Name: "Wavy T1 " + strconv.Itoa(i),
					V1:   v1_t1,
					V2:   v2_t1,
					V3:   v3_t1,
				}).Stage(stage)

				t2 := (&Triangle{
					Name: "Wavy T2 " + strconv.Itoa(i),
					V1:   v1_t2,
					V2:   v2_t2,
					V3:   v3_t2,
				}).Stage(stage)

				geom.Faces = append(geom.Faces, t1, t2)
			}
		}

		return (&Mesh{
			Name:            "Wavy Torus " + name + " Mesh",
			Position:        Position{X: 0, Y: 0, Z: 0},
			BufferGeometry:  geom,
			MeshPhysicalMaterial: (&MeshPhysicalMaterial{
				Name:                 "Wavy Torus " + name + " Material",
				Transparent:          false,
				Opacity:              1.0,
				MeshMaterialAbstract: MeshMaterialAbstract{Color: color},
			}).Stage(stage),
		}).Stage(stage)
	}

	canvas.Meshs = append(canvas.Meshs,
		createFaceMeshWavy("Bottom", "#1f77b4", bottomEdgesWavy, false), // blue
		createFaceMeshWavy("Top", "#d62728", topEdgesWavy, true),        // red
		createFaceMeshWavy("Inner", "#ff7f0e", innerEdgesWavy, true),    // orange
		createFaceMeshWavy("Outer", "#2ca02c", outerEdgesWavy, false),   // green
	)

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
	// dx2d := 1.147078 * scale
	dy2d := 9.933993 * scale
	thicknessPhylla := 1.5

	createFaceMeshPhylla := func(name string, color string, edges [][2]*Vector3, reverseWinding bool) *Mesh {
		geom := (&BufferGeometry{
			Name: "Phylla Torus " + name + " BufferGeometry",
		}).Stage(stage)

		// Create vertices and faces
		for i := 0; i < len(edges); i++ {
			p1_src := edges[i][0]
			p2_src := edges[i][1]
			
			p1 := (&Vector3{
				Name: p1_src.Name + " " + name + " " + strconv.Itoa(i),
				X:    p1_src.X,
				Y:    p1_src.Y,
				Z:    p1_src.Z,
			}).Stage(stage)

			p2 := (&Vector3{
				Name: p2_src.Name + " " + name + " " + strconv.Itoa(i),
				X:    p2_src.X,
				Y:    p2_src.Y,
				Z:    p2_src.Z,
			}).Stage(stage)
			
			// Copy pointers
			geom.Vertices = append(geom.Vertices, p1, p2)

			if i < len(edges)-1 {
				// Base indices for this segment
				idx := i * 2

				v1_t1, v2_t1, v3_t1 := idx, idx+1, idx+2
				v1_t2, v2_t2, v3_t2 := idx+1, idx+3, idx+2

				if reverseWinding {
					v2_t1, v3_t1 = v3_t1, v2_t1
					v2_t2, v3_t2 = v3_t2, v2_t2
				}

				// Triangle 1
				t1 := (&Triangle{
					Name: "T1 " + strconv.Itoa(i),
					V1:   v1_t1,
					V2:   v2_t1,
					V3:   v3_t1,
				}).Stage(stage)
				
				// Triangle 2
				t2 := (&Triangle{
					Name: "T2 " + strconv.Itoa(i),
					V1:   v1_t2,
					V2:   v2_t2,
					V3:   v3_t2,
				}).Stage(stage)

				geom.Faces = append(geom.Faces, t1, t2)
			}
		}

		return (&Mesh{
			Name:            "Phylla Torus " + name + " Mesh",
			Position:        Position{X: 0, Y: 0, Z: 0},
			BufferGeometry:  geom,
			MeshPhysicalMaterial: (&MeshPhysicalMaterial{
				Name:                 "Phylla Torus " + name + " Material",
				Transparent:          true,
				Opacity:              0.5,
				MeshMaterialAbstract: MeshMaterialAbstract{Color: color},
			}).Stage(stage),
		}).Stage(stage)
	}

	var bottomEdges, topEdges, innerEdges, outerEdges [][2]*Vector3

	for i := 0; i < len(curvePhylla.Points); i++ {
		p := curvePhylla.Points[i]
		
		// The point is (x3d, y3d, z3d) which is globalR * cos(theta), y, globalR * sin(theta)
		// We can compute theta:
		theta := math.Atan2(p.Z, p.X)

		// Compute the 4 corners of the cross section at this angle
		// We enforce perfectly vertical walls by matching X and Z radially
		
		xBL, yBL, zBL := p.X, p.Y, p.Z
		xBR, yBR, zBR := p.X + thicknessPhylla*math.Cos(theta), p.Y, p.Z + thicknessPhylla*math.Sin(theta)
		
		xTL, yTL, zTL := p.X, p.Y + dy2d, p.Z
		xTR, yTR, zTR := p.X + thicknessPhylla*math.Cos(theta), p.Y + dy2d, p.Z + thicknessPhylla*math.Sin(theta)

		vBL := (&Vector3{Name: "BL", X: xBL, Y: yBL, Z: zBL}).Stage(stage)
		vBR := (&Vector3{Name: "BR", X: xBR, Y: yBR, Z: zBR}).Stage(stage)
		vTL := (&Vector3{Name: "TL", X: xTL, Y: yTL, Z: zTL}).Stage(stage)
		vTR := (&Vector3{Name: "TR", X: xTR, Y: yTR, Z: zTR}).Stage(stage)

		bottomEdges = append(bottomEdges, [2]*Vector3{vBL, vBR})
		topEdges = append(topEdges, [2]*Vector3{vTL, vTR})
		innerEdges = append(innerEdges, [2]*Vector3{vBL, vTL})
		outerEdges = append(outerEdges, [2]*Vector3{vBR, vTR})
	}

	canvas.Meshs = append(canvas.Meshs,
		createFaceMeshPhylla("Bottom", "magenta", bottomEdges, false),
		createFaceMeshPhylla("Top", "yellow", topEdges, true),
		createFaceMeshPhylla("Inner", "cyan", innerEdges, true),
		createFaceMeshPhylla("Outer", "lime", outerEdges, false),
	)

	// --- Render the 4 original curves as thin tubes so the user can see them ---
	createTube := func(name string, color string, edges [][2]*Vector3, useLeft bool, tubeRadius float64) *Mesh {
		curve := (&Curve{
			Name: "Curve " + name,
		}).Stage(stage)

		for i := 0; i < len(edges); i++ {
			p := edges[i][0]
			if !useLeft {
				p = edges[i][1]
			}
			curve.Points = append(curve.Points, (&Vector3{
				Name: "CurvePoint " + name + " " + strconv.Itoa(i),
				X:    p.X,
				Y:    p.Y,
				Z:    p.Z,
			}).Stage(stage))
		}

		tubeGeometry := (&TubeGeometry{
			Name:            "TubeGeom " + name,
			Path:            curve,
			TubularSegments: 500,
			Radius:          tubeRadius,
			RadialSegments:  8,
			Closed:          false,
		}).Stage(stage)

		return (&Mesh{
			Name:              "TubeMesh " + name,
			Position:          Position{X: 0, Y: 0, Z: 0},
			TubeGeometry:      tubeGeometry,
			MeshMaterialBasic: (&MeshMaterialBasic{Name: name + " Material", MeshMaterialAbstract: MeshMaterialAbstract{Color: color}}).Stage(stage),
		}).Stage(stage)
	}

	outerRadius := 0.05
	innerRadius := outerRadius * 0.85 // 15% smaller

	canvas.Meshs = append(canvas.Meshs,
		createTube("BottomInner", "black", bottomEdges, true, innerRadius),
		createTube("BottomOuter", "gray", bottomEdges, false, outerRadius),
		createTube("TopInner", "darkgray", topEdges, true, innerRadius),
		createTube("TopOuter", "lightgray", topEdges, false, outerRadius),
	)

	stage.Commit()
}
