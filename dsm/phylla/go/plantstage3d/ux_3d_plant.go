package plantstage3d

import (
	"fmt"
	"math"

	"github.com/fullstack-lang/gong/dsm/phylla/go/cylinderstage3d"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

type Plant3DStageUpdater struct{}

func NewPlant3DStageUpdater() *Plant3DStageUpdater {
	return &Plant3DStageUpdater{}
}

func (u *Plant3DStageUpdater) UpdatePlant3DStage(stager *models.Stager) {
	u.ux_3d_plant(stager)
}

func (u *Plant3DStageUpdater) ux_3d_plant(stager *models.Stager) {
	plant3dStage := stager.GetPlant3dStage()
	if plant3dStage == nil {
		return
	}

	plant := stager.GetCurrentPlant()
	if plant == nil {
		plant3dStage.Reset()
		plant3dStage.Commit()
		return
	}

	var checkedDiagram *models.Plant3DDiagram
	for _, d := range plant.Plant3DDiagrams {
		if d.IsChecked {
			checkedDiagram = d
			break
		}
	}
	if checkedDiagram == nil && len(plant.Plant3DDiagrams) > 0 {
		checkedDiagram = plant.Plant3DDiagrams[0]
	}

	plant3dStage.Reset()

	canvas := (&threejs.Canvas{
		Name: "Plant 3D Canvas",
	}).Stage(plant3dStage)

	N := plant.N
	M := plant.M
	if N < 1 {
		N = 1
	}
	if M < 1 {
		M = 1
	}

	insideAngleRad := plant.RhombusInsideAngle * math.Pi / 180.0
	halfAngle := insideAngleRad / 2.0
	sideLength := plant.RhombusSideLength
	if sideLength <= 0 {
		sideLength = 10.0
	}

	v1_x := sideLength * math.Cos(halfAngle)
	v1_y := sideLength * math.Sin(halfAngle)
	v2_x := -sideLength * math.Cos(halfAngle)
	v2_y := sideLength * math.Sin(halfAngle)

	Wx := float64(N)*v1_x + float64(M)*v2_x
	Wy := float64(N)*v1_y + float64(M)*v2_y
	C := math.Sqrt(Wx*Wx + Wy*Wy)
	if C <= 0 {
		C = 10.0
	}

	globalR := C / (2.0 * math.Pi)

	phi := math.Atan2(Wy, Wx)
	cosRot := math.Cos(-phi)
	sinRot := math.Sin(-phi)

	v1_rot_x := v1_x*cosRot - v1_y*sinRot
	v1_rot_y := v1_x*sinRot + v1_y*cosRot
	v2_rot_x := v2_x*cosRot - v2_y*sinRot
	v2_rot_y := v2_x*sinRot + v2_y*cosRot

	if math.Abs(v1_rot_y) < 1e-6 {
		v1_rot_y = 1e-6
	}
	if math.Abs(v2_rot_y) < 1e-6 {
		v2_rot_y = 1e-6
	}

	var H float64
	if plant.StackHeight > 0 && plant.GrowthVectorShape != nil && plant.GrowthVectorShape.Y > 0 {
		H = float64(plant.StackHeight) * plant.GrowthVectorShape.Y
	}
	if H < 2.0*globalR {
		H = 2.5 * globalR
	}

	tubeRadius := math.Max(globalR*0.015, 1.2)
	rTubeCenter := globalR + tubeRadius*0.25

	// Lights
	lightScale := math.Max(globalR, H)
	dirLight1 := (&threejs.DirectionalLight{
		Name:             "Directional Light 1 (Key)",
		Position:         threejs.Position{X: lightScale * 2.0, Y: lightScale * 2.5, Z: lightScale * 2.0},
		LightAbstract:    threejs.LightAbstract{Intensity: 1.2},
		IsWithCastShadow: true,
	}).Stage(plant3dStage)

	dirLight2 := (&threejs.DirectionalLight{
		Name:             "Directional Light 2 (Fill)",
		Position:         threejs.Position{X: -lightScale * 2.0, Y: lightScale * 1.5, Z: -lightScale * 2.0},
		LightAbstract:    threejs.LightAbstract{Intensity: 0.6},
		IsWithCastShadow: false,
	}).Stage(plant3dStage)

	dirLight3 := (&threejs.DirectionalLight{
		Name:             "Directional Light 3 (Rim)",
		Position:         threejs.Position{X: 0, Y: lightScale * 3.5, Z: -lightScale * 2.5},
		LightAbstract:    threejs.LightAbstract{Intensity: 0.8},
		IsWithCastShadow: false,
	}).Stage(plant3dStage)

	canvas.DirectionalLights = append(canvas.DirectionalLights, dirLight1, dirLight2, dirLight3)

	ambiantLight := (&threejs.AmbiantLight{
		Name:          "Ambiant Light",
		LightAbstract: threejs.LightAbstract{Intensity: 0.45},
	}).Stage(plant3dStage)
	canvas.AmbiantLight = ambiantLight

	// Camera setup and persistence
	rendered3DShape := (*models.Rendered3DShape)(nil)
	if checkedDiagram != nil {
		rendered3DShape = checkedDiagram.Rendered3DShape
	}

	if rendered3DShape != nil && (rendered3DShape.ViewX != 0 || rendered3DShape.ViewY != 0 || rendered3DShape.ViewZ != 0) {
		fov := rendered3DShape.Fov
		if fov == 0 {
			fov = 50
		}
		canvas.Camera = (&threejs.Camera{
			Name: "Camera",
			Position: threejs.Position{
				X: rendered3DShape.ViewX,
				Y: rendered3DShape.ViewY,
				Z: rendered3DShape.ViewZ,
			},
			TargetX: rendered3DShape.TargetX,
			TargetY: rendered3DShape.TargetY,
			TargetZ: rendered3DShape.TargetZ,
			Fov:     fov,
		}).Stage(plant3dStage)
	} else {
		camDist := globalR * 2.8
		if camDist < 30 {
			camDist = 30
		}
		canvas.Camera = (&threejs.Camera{
			Name: "Camera",
			Position: threejs.Position{
				X: camDist,
				Y: H * 0.7,
				Z: camDist,
			},
			TargetY: H * 0.5,
			Fov:     50,
		}).Stage(plant3dStage)
	}

	canvas.Camera.OnUpdate = func(updatedCamera *threejs.Camera) {
		if rendered3DShape != nil {
			rendered3DShape.ViewX = updatedCamera.X
			rendered3DShape.ViewY = updatedCamera.Y
			rendered3DShape.ViewZ = updatedCamera.Z
			rendered3DShape.TargetX = updatedCamera.TargetX
			rendered3DShape.TargetY = updatedCamera.TargetY
			rendered3DShape.TargetZ = updatedCamera.TargetZ
			rendered3DShape.Fov = updatedCamera.Fov
			stager.GetStage().CommitWithSuspendedCallbacks()
		}
	}

	// 1. Stem Cylinder Surface
	if checkedDiagram == nil || !checkedDiagram.IsHiddenStemCylinder3DShape {
		cylGeom := (&threejs.CylinderGeometry{
			Name:           "Stem Cylinder Geom",
			RadiusTop:      globalR,
			RadiusBottom:   globalR,
			Height:         H,
			RadialSegments: 64,
			HeightSegments: 1,
			OpenEnded:      true,
		}).Stage(plant3dStage)

		cylMesh := (&threejs.Mesh{
			Name:             "Stem Cylinder Mesh",
			Position:         threejs.Position{X: 0, Y: H / 2.0, Z: 0},
			CylinderGeometry: cylGeom,
			MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
				Name:                 "Stem Cylinder Material",
				MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "#f1f5f9"},
				Transparent:          true,
				Opacity:              0.65,
			}).Stage(plant3dStage),
		}).Stage(plant3dStage)

		canvas.Meshs = append(canvas.Meshs, cylMesh)
	}

	// 2. Circumference Rings (Top and Bottom)
	if checkedDiagram == nil || !checkedDiagram.IsHiddenCircumference3DShape {
		createHorizontalRing := func(yRing float64, name string, color string) {
			ringCurve := (&threejs.Curve{
				Name: fmt.Sprintf("%s Curve", name),
			}).Stage(plant3dStage)

			ringSegments := 64
			for i := 0; i < ringSegments; i++ {
				theta := float64(i) * 2.0 * math.Pi / float64(ringSegments)
				ringCurve.Points = append(ringCurve.Points, (&threejs.Vector3{
					Name: fmt.Sprintf("%s Pt %d", name, i),
					X:    rTubeCenter * math.Cos(theta),
					Y:    yRing,
					Z:    rTubeCenter * math.Sin(theta),
				}).Stage(plant3dStage))
			}

			ringGeom := (&threejs.TubeGeometry{
				Name:            fmt.Sprintf("%s Geom", name),
				Path:            ringCurve,
				TubularSegments: ringSegments,
				Radius:          tubeRadius * 1.1,
				RadialSegments:  16,
				Closed:          true,
			}).Stage(plant3dStage)

			ringMesh := (&threejs.Mesh{
				Name:         fmt.Sprintf("%s Mesh", name),
				TubeGeometry: ringGeom,
				MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
					Name:                 fmt.Sprintf("%s Material", name),
					MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: color},
					Transparent:          true,
					Opacity:              0.9,
				}).Stage(plant3dStage),
			}).Stage(plant3dStage)

			canvas.Meshs = append(canvas.Meshs, ringMesh)
		}

		createHorizontalRing(0.0, "Bottom Circumference Ring", "#64748b")
		createHorizontalRing(H, "Top Circumference Ring", "#64748b")
	}

	// 3. Cut Line along the height of the cylinder (theta = 0)
	if checkedDiagram == nil || !checkedDiagram.IsHiddenCutLine3DShape {
		cutCurve := (&threejs.Curve{
			Name: "Cut Line Curve",
		}).Stage(plant3dStage)

		cutSteps := 32
		for s := 0; s <= cutSteps; s++ {
			yPt := float64(s) / float64(cutSteps) * H
			cutCurve.Points = append(cutCurve.Points, (&threejs.Vector3{
				Name: fmt.Sprintf("Cut Line Pt %d", s),
				X:    rTubeCenter,
				Y:    yPt,
				Z:    0,
			}).Stage(plant3dStage))
		}

		cutGeom := (&threejs.TubeGeometry{
			Name:            "Cut Line Geom",
			Path:            cutCurve,
			TubularSegments: cutSteps,
			Radius:          tubeRadius * 1.05,
			RadialSegments:  12,
			Closed:          false,
		}).Stage(plant3dStage)

		cutMesh := (&threejs.Mesh{
			Name:         "Cut Line Mesh",
			TubeGeometry: cutGeom,
			MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
				Name:                 "Cut Line Material",
				MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "#db2777"}, // Deep pink/magenta
				Transparent:          true,
				Opacity:              1.0,
			}).Stage(plant3dStage),
		}).Stage(plant3dStage)

		canvas.Meshs = append(canvas.Meshs, cutMesh)
	}

	// 4. Parastichy N Curves (Blue spirals)
	if checkedDiagram == nil || !checkedDiagram.IsHiddenParastichyNCurves3DShape {
		steps := 120
		for k := 0; k < N; k++ {
			theta0 := float64(k) * 2.0 * math.Pi / float64(N)

			curve := (&threejs.Curve{
				Name: fmt.Sprintf("Parastichy N Curve %d", k),
			}).Stage(plant3dStage)

			for s := 0; s <= steps; s++ {
				yPt := float64(s) / float64(steps) * H
				theta := theta0 + (yPt/v1_rot_y)*(v1_rot_x/globalR)

				curve.Points = append(curve.Points, (&threejs.Vector3{
					Name: fmt.Sprintf("Parastichy N Pt %d-%d", k, s),
					X:    rTubeCenter * math.Cos(theta),
					Y:    yPt,
					Z:    rTubeCenter * math.Sin(theta),
				}).Stage(plant3dStage))
			}

			geom := (&threejs.TubeGeometry{
				Name:            fmt.Sprintf("Parastichy N Geom %d", k),
				Path:            curve,
				TubularSegments: steps,
				Radius:          tubeRadius,
				RadialSegments:  16,
				Closed:          false,
			}).Stage(plant3dStage)

			mesh := (&threejs.Mesh{
				Name:         fmt.Sprintf("Parastichy N Mesh %d", k),
				TubeGeometry: geom,
				MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
					Name:                 fmt.Sprintf("Parastichy N Material %d", k),
					MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "#2563eb"}, // Vibrant Blue
					Transparent:          true,
					Opacity:              1.0,
				}).Stage(plant3dStage),
			}).Stage(plant3dStage)

			canvas.Meshs = append(canvas.Meshs, mesh)
		}
	}

	// 5. Parastichy M Curves (Orange spirals)
	if checkedDiagram == nil || !checkedDiagram.IsHiddenParastichyMCurves3DShape {
		steps := 120
		for m := 0; m < M; m++ {
			theta0 := float64(m) * 2.0 * math.Pi / float64(M)

			curve := (&threejs.Curve{
				Name: fmt.Sprintf("Parastichy M Curve %d", m),
			}).Stage(plant3dStage)

			for s := 0; s <= steps; s++ {
				yPt := float64(s) / float64(steps) * H
				theta := theta0 + (yPt/v2_rot_y)*(v2_rot_x/globalR)

				curve.Points = append(curve.Points, (&threejs.Vector3{
					Name: fmt.Sprintf("Parastichy M Pt %d-%d", m, s),
					X:    rTubeCenter * math.Cos(theta),
					Y:    yPt,
					Z:    rTubeCenter * math.Sin(theta),
				}).Stage(plant3dStage))
			}

			geom := (&threejs.TubeGeometry{
				Name:            fmt.Sprintf("Parastichy M Geom %d", m),
				Path:            curve,
				TubularSegments: steps,
				Radius:          tubeRadius,
				RadialSegments:  16,
				Closed:          false,
			}).Stage(plant3dStage)

			mesh := (&threejs.Mesh{
				Name:         fmt.Sprintf("Parastichy M Mesh %d", m),
				TubeGeometry: geom,
				MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
					Name:                 fmt.Sprintf("Parastichy M Material %d", m),
					MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "#ea580c"}, // Vibrant Orange
					Transparent:          true,
					Opacity:              1.0,
				}).Stage(plant3dStage),
			}).Stage(plant3dStage)

			canvas.Meshs = append(canvas.Meshs, mesh)
		}
	}

	// 6. Tiled Floor
	if checkedDiagram == nil || !checkedDiagram.IsHiddenTiledFloor3DShape {
		cylinderstage3d.AddFloorTiles(plant3dStage, canvas, globalR, 0.0)
	}

	plant3dStage.Commit()
}
