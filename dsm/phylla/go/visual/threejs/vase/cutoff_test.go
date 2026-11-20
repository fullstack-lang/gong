package vase

import (
	"math"
	"strings"
	"testing"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func TestCutoffPlaneAxisAlignedAndFloorUnclipped(t *testing.T) {
	stage := models.NewStage("test_cutoff")
	lib := (&models.Library{
		Name:          "Root",
		IsRootLibrary: true,
	}).Stage(stage)

	vase := (&models.TubeVaseAbstract{
		Name:                      "TubeVase",
		RadialRepetitions:         2,
		ColorCombination:          1,
		RelativeRadialThickness:   0.1,
		RelativeVerticalThickness: 0.1,
	}).Stage(stage)

	plant := (&models.PlantAbstract{
		Name:               "TestVase",
		N:                  3,
		M:                  5,
		RhombusSideLength:  80,
		RhombusInsideAngle: 120,
		StackHeight:        3,
		PlantType:          models.TubeVase,
		TubeVaseAbstract:   vase,
		CurrentView:        models.VIEW_TUBE_VASE_3D,
		IsSelected:         true,
	}).Stage(stage)
	lib.Plants = append(lib.Plants, plant)

	stager := models.NewStagerForTest(stage)
	stager.EnforceSemanticForTest()
	stager.SetCurrentPlant(plant)

	if len(plant.TubeVase3DDiagrams) == 0 {
		t.Fatalf("expected TubeVase3DDiagram after semantic enforcement")
	}
	diagram := plant.TubeVase3DDiagrams[0]
	diagram.IsChecked = true
	diagram.IsHiddenTiledFloor3DShape = false

	// Default CutoffAxis should be X
	if diagram.CutoffAxis != models.CutoffAxisX {
		t.Fatalf("expected default CutoffAxis=X, got %q", diagram.CutoffAxis)
	}

	threejsStage := threejs.NewStage("threejs")
	stager.SetThreejsStage(threejsStage)

	updater := NewThreeJSStageUpdater()
	updater.UpdateThreeJSStage(stager)

	canvases := threejsStage.GetInstancesSorted[*threejs.Canvas]()
	if len(canvases) != 1 {
		t.Fatalf("expected 1 Canvas, got %d", len(canvases))
	}
	canvas := canvases[0]
	if len(canvas.ClippingPlanes) != 1 {
		t.Fatalf("expected 1 ClippingPlane on Canvas, got %d", len(canvas.ClippingPlanes))
	}
	cp := canvas.ClippingPlanes[0]
	if cp.IsEnabled {
		t.Fatalf("expected ClippingPlane disabled initially")
	}

	// Verify floor tiles have IsNotClipped=true while vase meshes have IsNotClipped=false
	floorTileCount := 0
	clippedMeshCount := 0
	var sampleMesh *threejs.Mesh
	for _, m := range canvas.Meshs {
		if strings.HasPrefix(m.Name, "Floor Tile ") {
			floorTileCount++
			if !m.IsNotClipped {
				t.Fatalf("expected floor tile %q to have IsNotClipped=true", m.Name)
			}
		} else {
			clippedMeshCount++
			if m.IsNotClipped {
				t.Fatalf("expected vase mesh %q to have IsNotClipped=false", m.Name)
			}
			if sampleMesh == nil {
				sampleMesh = m
			}
		}
	}
	if floorTileCount == 0 {
		t.Fatalf("expected floor tiles in scene")
	}
	if clippedMeshCount == 0 {
		t.Fatalf("expected clipped vase meshes in scene")
	}

	// Enable X cutoff at offset 0.0 via fast-path UpdateCutoffPlane
	diagram.IsCutoffEnabled = true
	diagram.IsCutoffHelperVisible = true
	diagram.CutoffAxis = models.CutoffAxisX
	diagram.CutoffOffset = 0.0
	diagram.IsCutoffFlipped = false
	updater.UpdateCutoffPlane(stager)

	// Fast-path must preserve existing Mesh instances without rebuilding scene
	if _, ok := threejsStage.Meshs[sampleMesh]; !ok {
		t.Fatalf("expected fast-path UpdateCutoffPlane to preserve existing staged meshes")
	}
	if !cp.IsEnabled || !cp.IsHelperVisible {
		t.Fatalf("expected ClippingPlane IsEnabled=true and IsHelperVisible=true")
	}
	if cp.NormalX != -1.0 || cp.NormalY != 0.0 || cp.NormalZ != 0.0 || math.Abs(cp.Constant) > 1e-9 {
		t.Fatalf("unexpected X plane at offset 0: normal=(%v,%v,%v) constant=%v", cp.NormalX, cp.NormalY, cp.NormalZ, cp.Constant)
	}

	// Flip X cutoff at positive offset
	diagram.IsCutoffFlipped = true
	diagram.CutoffOffset = 0.5
	updater.UpdateCutoffPlane(stager)
	if cp.NormalX != 1.0 || cp.NormalY != 0.0 || cp.NormalZ != 0.0 || cp.Constant >= 0 {
		t.Fatalf("expected flipped X plane at +0.5 to have NormalX=+1 and Constant<0, got normal=(%v,%v,%v) constant=%v",
			cp.NormalX, cp.NormalY, cp.NormalZ, cp.Constant)
	}

	// Switch to Y axis (horizontal slice) unflipped at offset 0.0 (mid-height > 0)
	diagram.IsCutoffFlipped = false
	diagram.CutoffAxis = models.CutoffAxisY
	diagram.CutoffOffset = 0.0
	updater.UpdateCutoffPlane(stager)
	if cp.NormalX != 0.0 || cp.NormalY != -1.0 || cp.NormalZ != 0.0 {
		t.Fatalf("expected Y plane normal=(0,-1,0), got (%v,%v,%v)", cp.NormalX, cp.NormalY, cp.NormalZ)
	}
	if cp.Constant != cp.HelperCenterY {
		t.Fatalf("expected Y plane at offset 0 to sit at HelperCenterY=%v, got Constant=%v", cp.HelperCenterY, cp.Constant)
	}

	// Switch to Z axis
	diagram.CutoffAxis = models.CutoffAxisZ
	diagram.CutoffOffset = -0.25
	updater.UpdateCutoffPlane(stager)
	if cp.NormalX != 0.0 || cp.NormalY != 0.0 || cp.NormalZ != -1.0 || cp.Constant >= 0 {
		t.Fatalf("expected Z plane at -0.25 to have NormalZ=-1 and Constant<0, got normal=(%v,%v,%v) constant=%v",
			cp.NormalX, cp.NormalY, cp.NormalZ, cp.Constant)
	}

	// Verify lighting provides sufficient ambient and interior/plancher illumination (Three.js r182 PBR units)
	if canvas.AmbiantLight == nil || canvas.AmbiantLight.Intensity < 0.8 {
		t.Fatalf("expected AmbiantLight intensity >= 0.8 for readable shadow/interior surfaces, got %+v", canvas.AmbiantLight)
	}
	testNormals := [][3]float64{
		{0, 1, 0},   // Top rim & interior plancher (+Y)
		{1, 0, 0},   // +X wall
		{-1, 0, 0},  // -X wall
		{0, 0, 1},   // +Z wall
		{0, 0, -1},  // -Z wall
		{0, -1, 0},  // Underside (-Y)
	}
	for _, n := range testNormals {
		irr := canvas.AmbiantLight.Intensity
		for _, dl := range canvas.DirectionalLights {
			lLen := math.Sqrt(dl.X*dl.X + dl.Y*dl.Y + dl.Z*dl.Z)
			if lLen > 1e-9 {
				ndotl := (n[0]*dl.X + n[1]*dl.Y + n[2]*dl.Z) / lLen
				if ndotl > 0 {
					irr += dl.Intensity * ndotl
				}
			}
		}
		diffuseFactor := irr / math.Pi
		if diffuseFactor < 0.40 {
			t.Errorf("insufficient diffuse factor %.3f for normal %v (irradiance=%.3f)", diffuseFactor, n, irr)
		}
	}

	// Verify fast-path UpdateAmbientLight doses AmbiantLight.Intensity without rebuilding meshes
	diagram.AmbientLightLevel = 0.25
	updater.UpdateAmbientLight(stager)
	if _, ok := threejsStage.Meshs[sampleMesh]; !ok {
		t.Fatalf("expected fast-path UpdateAmbientLight to preserve existing staged meshes")
	}
	if math.Abs(canvas.AmbiantLight.Intensity-0.25) > 1e-9 {
		t.Fatalf("expected AmbiantLight.Intensity=0.25 after fast-path update, got %v", canvas.AmbiantLight.Intensity)
	}

	// Verify full UpdateThreeJSStage rebuild also applies diagram.AmbientLightLevel
	diagram.AmbientLightLevel = 1.85
	updater.UpdateThreeJSStage(stager)
	canvases = threejsStage.GetInstancesSorted[*threejs.Canvas]()
	if len(canvases) != 1 || canvases[0].AmbiantLight == nil || math.Abs(canvases[0].AmbiantLight.Intensity-1.85) > 1e-9 {
		t.Fatalf("expected AmbiantLight.Intensity=1.85 after full rebuild, got %+v", canvases[0].AmbiantLight)
	}
}

func TestVaseTrapezeAmbientLightParameter(t *testing.T) {
	stage := models.NewStage("test_vase_trapeze_ambient")
	lib := (&models.Library{
		Name:          "Root",
		IsRootLibrary: true,
	}).Stage(stage)

	vase := (&models.TubeVaseAbstract{
		Name:                      "Vase Trapeze-TubeVaseAbstract",
		Z_Ribbon:                  7.5,
		RibbonVerticalScale:       6.0,
		TopPlaneHeight:            0.6,
		BottomPlaneHeight:         0.2,
		BaseRingBottomPlaneHeight: 0.25,
		ProjectionAngle:           -15.0,
		BasePlateHeight:           0.25,
		HasConstantThickness:      true,
		HasRingBottom:             true,
		HasCurvedWall:             true,
		WallCurvature:             -0.18,
		HasLip:                    true,
		RelativeRadialThickness:   0.02,
		RelativeVerticalThickness: 0.06,
		RotationRatio:             0.99,
		RadialRepetitions:         1,
		ColorCombination:          4,
	}).Stage(stage)

	plant := (&models.PlantAbstract{
		Name:               "Vase Trapeze",
		N:                  8,
		M:                  13,
		RhombusSideLength:  100,
		RhombusInsideAngle: 65,
		StackHeight:        3,
		PlantType:          models.VaseTrapeze,
		TubeVaseAbstract:   vase,
		CurrentView:        models.VIEW_VASE_TRAPEZE_3D,
		IsSelected:         true,
	}).Stage(stage)
	lib.Plants = append(lib.Plants, plant)

	stager := models.NewStagerForTest(stage)
	stager.EnforceSemanticForTest()
	stager.SetCurrentPlant(plant)

	diagram := plant.TubeVase3DDiagrams[0]
	diagram.IsChecked = true
	if math.Abs(diagram.AmbientLightLevel-models.DefaultTubeVase3DAmbientLightLevel) > 1e-9 {
		t.Fatalf("expected default AmbientLightLevel=%.2f, got %v", models.DefaultTubeVase3DAmbientLightLevel, diagram.AmbientLightLevel)
	}

	threejsStage := threejs.NewStage("threejs_trapeze")
	stager.SetThreejsStage(threejsStage)

	updater := NewThreeJSStageUpdater()
	stager.SetThreeJSUpdater(updater)
	stager.UpdateThreeJSStage()

	canvases := threejsStage.GetInstancesSorted[*threejs.Canvas]()
	if len(canvases) != 1 || canvases[0].AmbiantLight == nil {
		t.Fatalf("expected Canvas with AmbiantLight in VaseTrapeze 3D scene")
	}
	if math.Abs(canvases[0].AmbiantLight.Intensity-models.DefaultTubeVase3DAmbientLightLevel) > 1e-9 {
		t.Fatalf("expected initial AmbiantLight.Intensity=%.2f, got %v",
			models.DefaultTubeVase3DAmbientLightLevel, canvases[0].AmbiantLight.Intensity)
	}

	diagram.AmbientLightLevel = 0.20
	stager.UpdateAmbientLight()
	if math.Abs(canvases[0].AmbiantLight.Intensity-0.20) > 1e-9 {
		t.Fatalf("expected AmbiantLight.Intensity=0.20 after UpdateAmbientLight, got %v", canvases[0].AmbiantLight.Intensity)
	}
}
