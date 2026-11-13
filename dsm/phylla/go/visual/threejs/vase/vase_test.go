package vase

import (
	"math"
	"path/filepath"
	"strings"
	"testing"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func TestInspectBasePlateAndRing(t *testing.T) {
	stagePath := filepath.Join("..", "..", "..", "cmd", "phylla", "data", "stage.go")
	stage := models.NewStage("test_stageset")
	err := stage.ParseAstFile(stagePath, true)
	if err != nil {
		t.Fatalf("failed to parse stage.go: %v", err)
	}
	stage.ComputeReverseMaps()
	stage.ComputeInstancesNb()
	stage.ComputeReferenceAndOrders()

	var targetPlant *models.PlantAbstract
	for plant := range stage.PlantAbstracts {
		if plant.Name == "Vase Trapeze" {
			plant.IsSelected = true
			targetPlant = plant
		} else {
			plant.IsSelected = false
		}
	}
	stager := models.NewStagerForTest(stage)
	stager.EnforceSemanticForTest()
	stager.SetCurrentPlant(targetPlant)
	for d := range stage.TubeVase3DDiagrams {
		if d.Name == "Vase Trapeze-TubeVase3DDiagram" {
			d.IsChecked = true
			d.IsHiddenTiledFloor3DShape = true
			d.IsHiddenStackOfRotatedCarvedOutVaseTrapezeRingsShape = false
			d.IsHiddenVaseTrapezeBasePlateShape = false
		}
	}
	threejsStage := threejs.NewStage("threejs")
	stager.SetThreejsStage(threejsStage)

	updater := NewThreeJSStageUpdater()
	updater.UpdateThreeJSStage(stager)

	threejsStage = stager.GetThreejsStage()
	for c := range threejsStage.Curves {
		minY, maxY := math.MaxFloat64, -math.MaxFloat64
		for _, p := range c.Points {
			if p.Y < minY {
				minY = p.Y
			}
			if p.Y > maxY {
				maxY = p.Y
			}
		}
		t.Logf("Curve %s: %d points, Y in [%.2f, %.2f]", c.Name, len(c.Points), minY, maxY)
	}
	for mesh := range threejsStage.Meshs {
		t.Logf("Found Mesh: %q (Visible: %v)", mesh.Name, mesh.BufferGeometry != nil)
		if mesh.BufferGeometry != nil {
			t.Logf("Mesh %s: %d vertices, %d faces", mesh.Name, len(mesh.BufferGeometry.Vertices), len(mesh.BufferGeometry.Faces))
			if len(mesh.BufferGeometry.Vertices) > 0 {
				minY := math.MaxFloat64
				maxY := -math.MaxFloat64
				minR := math.MaxFloat64
				maxR := -math.MaxFloat64
				for _, v := range mesh.BufferGeometry.Vertices {
					if v.Y < minY {
						minY = v.Y
					}
					if v.Y > maxY {
						maxY = v.Y
					}
					r := math.Hypot(v.X, v.Z)
					if r < minR {
						minR = r
					}
					if r > maxR {
						maxR = r
					}
				}
				t.Logf("   Y range: [%.2f, %.2f], R range: [%.2f, %.2f]", minY, maxY, minR, maxR)
			}
		}
	}

	// Verify that h0 is not generated in Stack Of Rotated Carved Out Vase Trapeze Rings,
	// and that Vase Trapeze Base Plate includes the bottom ring (spanning from floorY=0 up to p1H).
	var ringH0Mesh, basePlateMesh *threejs.Mesh
	for mesh := range threejsStage.Meshs {
		if mesh.Name == "Vase Trapeze Stack Of Rotated Carved Out Vase Trapeze Rings h0 Mesh" {
			ringH0Mesh = mesh
		}
		if mesh.Name == "Vase Trapeze Vase Trapeze Base Plate Mesh" {
			basePlateMesh = mesh
		}
	}
	if ringH0Mesh != nil {
		t.Errorf("expected bottom ring (h0) not to be displayed in Stack Of Rotated Carved Out Vase Trapeze Rings")
	}
	if basePlateMesh == nil {
		t.Fatalf("expected Vase Trapeze Base Plate Mesh to be present")
	}
	minBaseY, maxBaseY := math.MaxFloat64, -math.MaxFloat64
	for _, v := range basePlateMesh.BufferGeometry.Vertices {
		if v.Y < minBaseY {
			minBaseY = v.Y
		}
		if v.Y > maxBaseY {
			maxBaseY = v.Y
		}
	}
	expectedP1H := (targetPlant.TubeVaseAbstract.Z_Ribbon + targetPlant.TubeVaseAbstract.TopPlaneHeight*targetPlant.TubeVaseAbstract.RibbonVerticalScale) * targetPlant.RhombusSideLength
	if math.Abs(minBaseY) > 1e-4 || math.Abs(maxBaseY-expectedP1H) > 1e-2 {
		t.Errorf("expected BasePlateMesh Y range [0.00, %.2f] (including bottom ring), got [%.2f, %.2f]", expectedP1H, minBaseY, maxBaseY)
	}

	// Verify STL exports for Vase Trapeze (one ring & base+first ring)
	oneRingSTL := models.GenerateVaseTrapezeOneRingSTL(targetPlant)
	if len(oneRingSTL) == 0 {
		t.Fatalf("expected non-empty one ring STL")
	}
	baseFirstRingSTL := models.GenerateVaseTrapezeBaseAndFirstRingSTL(targetPlant)
	if len(baseFirstRingSTL) == 0 {
		t.Fatalf("expected non-empty base+first ring STL")
	}
	glbBytes := models.GenerateVaseTrapezeGLB(targetPlant)
	if len(glbBytes) < 28 || string(glbBytes[0:4]) != "glTF" {
		t.Fatalf("expected valid glTF 2.0 binary (.glb), got %d bytes", len(glbBytes))
	}
	t.Logf("Generated Vase Trapeze One Ring STL (%d bytes), Base+First Ring STL (%d bytes), and Full Vase GLB (%d bytes)",
		len(oneRingSTL), len(baseFirstRingSTL), len(glbBytes))
}

func TestVaseTrapezeConstant3DThickness(t *testing.T) {
	stagePath := filepath.Join("..", "..", "..", "cmd", "phylla", "data", "stage.go")
	stage := models.NewStage("test_thickness")
	if err := stage.ParseAstFile(stagePath, true); err != nil {
		t.Fatalf("failed to parse stage.go: %v", err)
	}
	stage.ComputeReverseMaps()
	stage.ComputeInstancesNb()
	stage.ComputeReferenceAndOrders()

	var targetPlant *models.PlantAbstract
	for plant := range stage.PlantAbstracts {
		if plant.Name == "Vase Trapeze" {
			plant.IsSelected = true
			targetPlant = plant
		} else {
			plant.IsSelected = false
		}
	}
	if targetPlant == nil {
		t.Fatalf("Vase Trapeze not found")
	}
	targetPlant.N = 1
	targetPlant.M = 3
	targetPlant.TubeVaseAbstract.RadialRepetitions = 3
	targetPlant.TubeVaseAbstract.RotationRatio = 1.0
	targetPlant.TubeVaseAbstract.HasConstantThickness = true

	stager := models.NewStagerForTest(stage)
	stager.EnforceSemanticForTest()
	stager.SetCurrentPlant(targetPlant)
	for d := range stage.TubeVase3DDiagrams {
		if d.Name == "Vase Trapeze-TubeVase3DDiagram" {
			d.IsChecked = true
			d.IsHiddenTiledFloor3DShape = true
			d.IsHiddenStackOfRotatedCarvedOutVaseTrapezeRingsShape = false
		}
	}
	threejsStage := threejs.NewStage("threejs")
	stager.SetThreejsStage(threejsStage)

	updater := NewThreeJSStageUpdater()
	updater.UpdateThreeJSStage(stager)

	var baseBottom, baseTop *threejs.Curve
	var topP2h0, botP2h0, topP1h0, botP2h1 *threejs.Curve
	for c := range threejsStage.Curves {
		switch c.Name {
		case "Base Bottom Resampled":
			baseBottom = c
		case "Base Top Resampled":
			baseTop = c
		case "Vase Trapeze Rotated Carved Stack Top P2 h0":
			topP2h0 = c
		case "Vase Trapeze Rotated Carved Stack Bottom P2 h0":
			botP2h0 = c
		case "Vase Trapeze Rotated Carved Stack Top P1 h0":
			topP1h0 = c
		case "Vase Trapeze Rotated Carved Stack Bottom P2 h1":
			botP2h1 = c
		}
	}
	if baseBottom == nil || baseTop == nil {
		t.Fatalf("missing Resampled Base Bottom or Base Top curves")
	}

	vase := targetPlant.TubeVaseAbstract
	scaleY := vase.RibbonVerticalScale
	projAngleRad := -vase.ProjectionAngle * math.Pi / 180.0
	tanAlpha := math.Tan(projAngleRad)
	cosAlpha := math.Cos(projAngleRad)
	globalR := updater.computeGlobalRadius(targetPlant)
	p1H := (vase.Z_Ribbon + vase.TopPlaneHeight*scaleY) * targetPlant.RhombusSideLength
	p2H := (vase.Z_Ribbon + vase.BottomPlaneHeight*scaleY) * targetPlant.RhombusSideLength
	yRef := 0.5 * (p1H + p2H)

	minY, maxY := math.MaxFloat64, -math.MaxFloat64
	for _, pt := range baseBottom.Points {
		if pt.Y < minY {
			minY = pt.Y
		}
		if pt.Y > maxY {
			maxY = pt.Y
		}
	}
	rRef := globalR + (yRef-0.5*(minY+maxY))*tanAlpha

	K := len(baseBottom.Points) - 1
	dTheta := (2.0 * math.Pi / float64(vase.RadialRepetitions)) / float64(K)

	minT3D, maxT3D := math.MaxFloat64, -math.MaxFloat64
	maxD2Bot, maxD2Top := 0.0, 0.0
	for i := 0; i < K; i++ {
		iPrev := (i - 1 + K) % K
		iNext := (i + 1) % K
		slope := (baseBottom.Points[iNext].Y - baseBottom.Points[iPrev].Y) / (2.0 * dTheta)
		d2Bot := baseBottom.Points[iNext].Y - 2*baseBottom.Points[i].Y + baseBottom.Points[iPrev].Y
		d2Top := baseTop.Points[iNext].Y - 2*baseTop.Points[i].Y + baseTop.Points[iPrev].Y
		if math.Abs(d2Bot) > maxD2Bot {
			maxD2Bot = math.Abs(d2Bot)
		}
		if math.Abs(d2Top) > maxD2Top {
			maxD2Top = math.Abs(d2Top)
		}
		if math.Abs(d2Bot) > 1.0 || math.Abs(d2Top) > 1.5 {
			t.Logf("High d2 at i=%d (theta=%.2f deg): d2Bot=%.3f, d2Top=%.3f, botY=%.2f, topY=%.2f",
				i, float64(i)*dTheta*180/math.Pi, d2Bot, d2Top, baseBottom.Points[i].Y, baseTop.Points[i].Y)
		}
		dy := baseTop.Points[i].Y - baseBottom.Points[i].Y
		dHoriz := (dy * tanAlpha) / math.Hypot(1.0, (slope*tanAlpha)/rRef)
		d3D := dHoriz * cosAlpha
		if d3D < minT3D {
			minT3D = d3D
		}
		if d3D > maxT3D {
			maxT3D = d3D
		}
	}
	t.Logf("3D Wall Thickness across sector: min=%.4f mm, max=%.4f mm (ratio=%.6f), maxD2Bot=%.2f, maxD2Top=%.2f",
		minT3D, maxT3D, maxT3D/minT3D, maxD2Bot, maxD2Top)
	if maxT3D/minT3D > 1.005 {
		t.Errorf("expected constant 3D wall thickness (ratio <= 1.005), got min=%.4f, max=%.4f, ratio=%.6f", minT3D, maxT3D, maxT3D/minT3D)
	}
	if maxD2Bot > 1.5 || maxD2Top > 2.5 {
		t.Errorf("expected smooth curves without squirks (maxD2Bot <= 1.5, maxD2Top <= 2.5), got maxD2Bot=%.2f, maxD2Top=%.2f", maxD2Bot, maxD2Top)
	}

	// Verify rotated fit between Ring 0 inner wall (baseTop) and Ring 1 outer wall (baseBottom shifted by rotatedStepDX, rotatedStepDY)
	var growthVectorX, growthVectorY float64
	if targetPlant.GrowthVectorShape != nil {
		growthVectorX = targetPlant.GrowthVectorShape.X
		growthVectorY = targetPlant.GrowthVectorShape.Y
	}
	vFirst := targetPlant.PerpendicularVectorGrid.PerpendicularVectors[0]
	vx := vFirst.EndX - vFirst.StartX
	vy := vFirst.EndY - vFirst.StartY
	vLen := math.Hypot(vx, vy)
	vx, vy = vx/vLen, vy/vLen
	verticalThickness := vase.RelativeVerticalThickness * targetPlant.RhombusSideLength
	deltaYV := (verticalThickness / vy) * scaleY
	rotStepDY := growthVectorY*scaleY + deltaYV + vase.RelativeRotatedTorusSeparation*targetPlant.RhombusSideLength*scaleY
	rotStepTheta := growthVectorX / globalR
	thetaSector := 2.0 * math.Pi / float64(vase.RadialRepetitions)

	// Interpolate baseTop(theta) periodically
	evalCurveY := func(c *threejs.Curve, theta float64) float64 {
		for theta < 0 {
			theta += thetaSector
		}
		for theta >= thetaSector {
			theta -= thetaSector
		}
		pos := theta / dTheta
		idx := int(math.Floor(pos))
		if idx < 0 {
			idx = 0
		}
		if idx >= K {
			idx = K - 1
		}
		frac := pos - float64(idx)
		return c.Points[idx].Y*(1.0-frac) + c.Points[idx+1].Y*frac
	}

	// On M1 (theta in [61.5 deg, 90.0 deg], i in [123..180]) and M2 (theta in [91.0 deg, 119.5 deg], i in [182..239]),
	// Ring 1's outer wall at theta + rotStepTheta has height baseBottom(theta) + rotStepDY,
	// which touches Ring 0's inner wall at theta + rotStepTheta with height baseTop(theta + rotStepTheta).
	maxRadialGapM1M2 := 0.0
	maxRadialGapM2M3 := 0.0
	for i := 124; i <= 179; i++ {
		theta := float64(i) * dTheta
		yOuterRing1 := baseBottom.Points[i].Y + rotStepDY
		yInnerRing0 := evalCurveY(baseTop, theta+rotStepTheta)
		radialGap := math.Abs(yOuterRing1-yInnerRing0) * tanAlpha
		if radialGap > 0.05 {
			t.Logf("M1->M2 gap at i=%d (theta=%.2f deg, shifted=%.2f deg): yOuter1=%.4f, yInner0=%.4f, radialGap=%.4f mm",
				i, theta*180/math.Pi, (theta+rotStepTheta)*180/math.Pi, yOuterRing1, yInnerRing0, radialGap)
		}
		if radialGap > maxRadialGapM1M2 {
			maxRadialGapM1M2 = radialGap
		}
	}
	for i := 183; i <= 238; i++ {
		theta := float64(i) * dTheta
		yOuterRing1 := baseBottom.Points[i].Y + rotStepDY
		yInnerRing0 := evalCurveY(baseTop, theta+rotStepTheta)
		radialGap := math.Abs(yOuterRing1-yInnerRing0) * tanAlpha
		if radialGap > maxRadialGapM2M3 {
			maxRadialGapM2M3 = radialGap
		}
	}
	t.Logf("Max rotated radial contact gap: M1->M2 = %.4f mm, M2->M3 = %.4f mm", maxRadialGapM1M2, maxRadialGapM2M3)
	if maxRadialGapM1M2 > 0.05 || maxRadialGapM2M3 > 0.05 {
		t.Errorf("expected rotated radial contact gap <= 0.05 mm on touching M-segments, got M1->M2=%.4f mm, M2->M3=%.4f mm", maxRadialGapM1M2, maxRadialGapM2M3)
	}
	if topP2h0 != nil && botP2h0 != nil && topP1h0 != nil && botP2h1 != nil {
		t.Logf("Projected curves found: topP2h0=%d pts, botP2h1=%d pts", len(topP2h0.Points), len(botP2h1.Points))
	}

	// Now verify that when HasConstantThickness is false, the original variable thickness is used
	targetPlant.TubeVaseAbstract.HasConstantThickness = false
	threejsStageVar := threejs.NewStage("threejs_var")
	stager.SetThreejsStage(threejsStageVar)
	updater.UpdateThreeJSStage(stager)

	var varBottom, varTop *threejs.Curve
	for c := range threejsStageVar.Curves {
		switch c.Name {
		case "Base Bottom Resampled":
			varBottom = c
		case "Base Top Resampled":
			varTop = c
		}
	}
	if varBottom == nil || varTop == nil {
		t.Fatalf("missing Resampled Base Bottom or Base Top curves in variable thickness mode")
	}
	minVarT3D, maxVarT3D := math.MaxFloat64, -math.MaxFloat64
	for i := 0; i < K; i++ {
		iPrev := (i - 1 + K) % K
		iNext := (i + 1) % K
		slope := (varBottom.Points[iNext].Y - varBottom.Points[iPrev].Y) / (2.0 * dTheta)
		dy := varTop.Points[i].Y - varBottom.Points[i].Y
		dHoriz := (dy * tanAlpha) / math.Hypot(1.0, (slope*tanAlpha)/rRef)
		d3D := dHoriz * cosAlpha
		if d3D < minVarT3D {
			minVarT3D = d3D
		}
		if d3D > maxVarT3D {
			maxVarT3D = d3D
		}
	}
	t.Logf("Variable 3D Wall Thickness across sector (HasConstantThickness=false): min=%.4f mm, max=%.4f mm (ratio=%.6f)",
		minVarT3D, maxVarT3D, maxVarT3D/minVarT3D)
	if maxVarT3D/minVarT3D <= 1.05 {
		t.Errorf("expected variable 3D wall thickness when HasConstantThickness=false (ratio > 1.05), got min=%.4f, max=%.4f, ratio=%.6f",
			minVarT3D, maxVarT3D, maxVarT3D/minVarT3D)
	}
}

func TestVaseColorCombinations(t *testing.T) {
	for comb := 1; comb <= 10; comb++ {
		c0 := models.GetVaseRingColor(comb, 0)
		c1 := models.GetVaseRingColor(comb, 1)
		c2 := models.GetVaseRingColor(comb, 2)
		c3 := models.GetVaseRingColor(comb, 3)
		if c0 == "" || c1 == "" || c2 == "" || c3 == "" {
			t.Fatalf("combination %d returned empty color", comb)
		}
		if comb <= 4 {
			if c0 != c1 || c1 != c2 || c2 != c3 {
				t.Errorf("expected combination %d to be homogeneous, got %s, %s, %s, %s", comb, c0, c1, c2, c3)
			}
		} else {
			if c0 == c1 {
				t.Errorf("expected combination %d to be heterogeneous across adjacent rings, got %s and %s", comb, c0, c1)
			}
		}
	}
}

func TestVaseTrapezeUnoccupiedInteriorSurfaceProportion(t *testing.T) {
	stagePath := filepath.Join("..", "..", "..", "cmd", "phylla", "data", "stage.go")
	stage := models.NewStage("test_unoccupied_surface")
	if err := stage.ParseAstFile(stagePath, true); err != nil {
		t.Fatalf("failed to parse stage.go: %v", err)
	}
	stage.ComputeReverseMaps()
	stage.ComputeInstancesNb()
	stage.ComputeReferenceAndOrders()

	var targetPlant *models.PlantAbstract
	for plant := range stage.PlantAbstracts {
		if plant.Name == "Vase Trapeze" {
			plant.IsSelected = true
			targetPlant = plant
		} else {
			plant.IsSelected = false
		}
	}
	if targetPlant == nil {
		t.Fatalf("Vase Trapeze not found")
	}

	stager := models.NewStagerForTest(stage)
	stager.EnforceSemanticForTest()

	targetPlant.TubeVaseAbstract.CarvedOutTopRingsParameter = 0.0
	unoccupied0, interior0, prop0, ok0 := models.ComputeVaseTrapezeUnoccupiedInteriorSurfaceProportion(targetPlant)
	if !ok0 {
		t.Fatalf("expected ComputeVaseTrapezeUnoccupiedInteriorSurfaceProportion to succeed")
	}
	if interior0 <= 0 || unoccupied0 <= 0 || prop0 <= 0 || prop0 >= 1 {
		t.Errorf("expected valid positive surfaces and proportion in (0, 1), got unoccupied=%.2f, interior=%.2f, prop=%.4f",
			unoccupied0, interior0, prop0)
	}
	t.Logf("Vase Trapeze unoccupied interior surface at ring top: %.4f (%.2f%%) (unoccupied: %.2f mm², interior: %.2f mm²)",
		prop0, prop0*100.0, unoccupied0, interior0)

	// Verify independence from CarvedOutTopRingsParameter (always computed as if CarvedOutTopRingsParameter == 0)
	targetPlant.TubeVaseAbstract.CarvedOutTopRingsParameter = 0.75
	unoccupied1, interior1, prop1, ok1 := models.ComputeVaseTrapezeUnoccupiedInteriorSurfaceProportion(targetPlant)
	if !ok1 {
		t.Fatalf("expected ComputeVaseTrapezeUnoccupiedInteriorSurfaceProportion to succeed with CarvedOutTopRingsParameter=0.75")
	}
	if math.Abs(unoccupied0-unoccupied1) > 1e-9 || math.Abs(interior0-interior1) > 1e-9 || math.Abs(prop0-prop1) > 1e-12 {
		t.Errorf("expected result to be independent of CarvedOutTopRingsParameter, got (%.2f, %.2f, %.6f) vs (%.2f, %.2f, %.6f)",
			unoccupied0, interior0, prop0, unoccupied1, interior1, prop1)
	}
}

func TestVaseTrapezeRingBottom(t *testing.T) {
	stagePath := filepath.Join("..", "..", "..", "cmd", "phylla", "data", "stage.go")
	stage := models.NewStage("test_ring_bottom")
	if err := stage.ParseAstFile(stagePath, true); err != nil {
		t.Fatalf("failed to parse stage.go: %v", err)
	}
	stage.ComputeReverseMaps()
	stage.ComputeInstancesNb()
	stage.ComputeReferenceAndOrders()

	var targetPlant *models.PlantAbstract
	for plant := range stage.PlantAbstracts {
		if plant.Name == "Vase Trapeze" {
			plant.IsSelected = true
			targetPlant = plant
		} else {
			plant.IsSelected = false
		}
	}
	if targetPlant == nil {
		t.Fatalf("Vase Trapeze not found")
	}

	stager := models.NewStagerForTest(stage)
	stager.EnforceSemanticForTest()
	stager.SetCurrentPlant(targetPlant)
	for d := range stage.TubeVase3DDiagrams {
		if d.Name == "Vase Trapeze-TubeVase3DDiagram" {
			d.IsChecked = true
			d.IsHiddenTiledFloor3DShape = true
			d.IsHiddenVaseTrapezeRingShape = false
			d.IsHiddenCarvedOutVaseTrapezeRingShape = false
			d.IsHiddenStackOfRotatedCarvedOutVaseTrapezeRingsShape = false
		}
	}

	targetPlant.TubeVaseAbstract.HasRingBottom = false
	threejsStageNoBot := threejs.NewStage("threejs_no_bot")
	stager.SetThreejsStage(threejsStageNoBot)
	updater := NewThreeJSStageUpdater()
	updater.UpdateThreeJSStage(stager)

	var ringNoBotFaces int
	for mesh := range threejsStageNoBot.Meshs {
		if mesh.Name == "Vase Trapeze Vase Trapeze Ring Mesh" && mesh.BufferGeometry != nil {
			ringNoBotFaces = len(mesh.BufferGeometry.Faces)
		}
	}
	stlNoBot := models.GenerateVaseTrapezeOneRingSTL(targetPlant)
	glbNoBot := models.GenerateVaseTrapezeGLB(targetPlant)

	targetPlant.TubeVaseAbstract.HasRingBottom = true
	threejsStageBot := threejs.NewStage("threejs_bot")
	stager.SetThreejsStage(threejsStageBot)
	updater.UpdateThreeJSStage(stager)

	var ringBotMesh *threejs.Mesh
	for mesh := range threejsStageBot.Meshs {
		if mesh.Name == "Vase Trapeze Vase Trapeze Ring Mesh" {
			ringBotMesh = mesh
		}
	}
	if ringBotMesh == nil || ringBotMesh.BufferGeometry == nil {
		t.Fatalf("expected Vase Trapeze Ring Mesh with HasRingBottom=true")
	}
	if len(ringBotMesh.BufferGeometry.Faces) != ringNoBotFaces {
		t.Errorf("expected same face count (inner+outer walls + 2 disk fans vs 2 annular quad strips = 8M triangles), got %d vs %d",
			len(ringBotMesh.BufferGeometry.Faces), ringNoBotFaces)
	}

	// Verify Plancher Center and Floor Center vertices exist at expected heights
	vase := targetPlant.TubeVaseAbstract
	scaleY := vase.RibbonVerticalScale
	sideLength := targetPlant.RhombusSideLength
	p2H := (vase.Z_Ribbon + vase.BottomPlaneHeight*scaleY) * sideLength
	expectedPlancherY := p2H + vase.BasePlateHeight*sideLength

	var foundPlancherCenter, foundFloorCenter bool
	for _, v := range ringBotMesh.BufferGeometry.Vertices {
		if v.Name == "Vase Trapeze Vase Trapeze Ring Plancher Center" {
			foundPlancherCenter = true
			if math.Abs(v.Y-expectedPlancherY) > 1e-3 {
				t.Errorf("expected Plancher Center Y=%.2f, got %.2f", expectedPlancherY, v.Y)
			}
		}
		if v.Name == "Vase Trapeze Vase Trapeze Ring Floor Center" {
			foundFloorCenter = true
			if math.Abs(v.Y-p2H) > 1e-3 {
				t.Errorf("expected Floor Center Y=%.2f, got %.2f", p2H, v.Y)
			}
		}
	}
	if !foundPlancherCenter || !foundFloorCenter {
		t.Errorf("expected Plancher Center (%v) and Floor Center (%v) vertices in ring mesh", foundPlancherCenter, foundFloorCenter)
	}

	stlBot := models.GenerateVaseTrapezeOneRingSTL(targetPlant)
	glbBot := models.GenerateVaseTrapezeGLB(targetPlant)
	if stlBot == stlNoBot {
		t.Errorf("expected OneRing STL with HasRingBottom=true to differ from HasRingBottom=false")
	}
	if len(glbBot) <= len(glbNoBot) {
		t.Errorf("expected GLB with HasRingBottom=true (%d bytes) to have dedicated cap vertices and be larger than HasRingBottom=false (%d bytes)",
			len(glbBot), len(glbNoBot))
	}
}

func TestVaseTrapezeSideLengthProportionalScaling(t *testing.T) {
	stagePath := filepath.Join("..", "..", "..", "cmd", "phylla", "data", "stage.go")
	stage := models.NewStage("test_proportional_scaling")
	if err := stage.ParseAstFile(stagePath, true); err != nil {
		t.Fatalf("failed to parse stage.go: %v", err)
	}
	stage.ComputeReverseMaps()
	stage.ComputeInstancesNb()
	stage.ComputeReferenceAndOrders()

	var targetPlant *models.PlantAbstract
	for plant := range stage.PlantAbstracts {
		if plant.Name == "Vase Trapeze" {
			plant.IsSelected = true
			targetPlant = plant
		} else {
			plant.IsSelected = false
		}
	}
	if targetPlant == nil {
		t.Fatalf("Vase Trapeze not found")
	}

	for d := range stage.TubeVase3DDiagrams {
		if d.Name == "Vase Trapeze-TubeVase3DDiagram" {
			d.IsChecked = true
			d.IsHiddenTiledFloor3DShape = true
			d.IsHiddenStackOfRotatedCarvedOutVaseTrapezeRingsShape = false
			d.IsHiddenVaseTrapezeBasePlateShape = false
		}
	}

	measureBounds := func(sideLen float64) (minY, maxY, minR, maxR, prop float64) {
		targetPlant.RhombusSideLength = sideLen
		stager := models.NewStagerForTest(stage)
		stager.EnforceSemanticForTest()
		stager.SetCurrentPlant(targetPlant)
		threejsStage := threejs.NewStage("threejs_scale")
		stager.SetThreejsStage(threejsStage)
		updater := NewThreeJSStageUpdater()
		updater.UpdateThreeJSStage(stager)

		minY, maxY = math.MaxFloat64, -math.MaxFloat64
		minR, maxR = math.MaxFloat64, -math.MaxFloat64
		for mesh := range threejsStage.Meshs {
			if mesh.BufferGeometry == nil {
				continue
			}
			for _, v := range mesh.BufferGeometry.Vertices {
				if v.Y < minY {
					minY = v.Y
				}
				if v.Y > maxY {
					maxY = v.Y
				}
				r := math.Hypot(v.X, v.Z)
				if r < minR {
					minR = r
				}
				if r > maxR {
					maxR = r
				}
			}
		}
		_, _, prop, _ = models.ComputeVaseTrapezeUnoccupiedInteriorSurfaceProportion(targetPlant)
		return
	}

	_, maxY100, _, maxR100, prop100 := measureBounds(100.0)
	_, maxY200, _, maxR200, prop200 := measureBounds(200.0)

	if math.Abs(maxY200/maxY100-2.0) > 1e-4 {
		t.Errorf("expected maxY to scale by 2.0 when SideLength doubles (100->200), got %.2f -> %.2f (ratio %.6f)",
			maxY100, maxY200, maxY200/maxY100)
	}
	if math.Abs(maxR200/maxR100-2.0) > 1e-4 {
		t.Errorf("expected maxR to scale by 2.0 when SideLength doubles (100->200), got %.2f -> %.2f (ratio %.6f)",
			maxR100, maxR200, maxR200/maxR100)
	}
	if math.Abs(prop200-prop100) > 1e-6 {
		t.Errorf("expected unoccupied interior surface proportion to be scale-invariant, got %.6f vs %.6f", prop100, prop200)
	}
}

func TestVaseTrapezeCurvedWall(t *testing.T) {
	stagePath := filepath.Join("..", "..", "..", "cmd", "phylla", "data", "stage.go")
	stage := models.NewStage("test_curved_wall")
	if err := stage.ParseAstFile(stagePath, true); err != nil {
		t.Fatalf("failed to parse stage.go: %v", err)
	}
	stage.ComputeReverseMaps()
	stage.ComputeInstancesNb()
	stage.ComputeReferenceAndOrders()

	var targetPlant *models.PlantAbstract
	for plant := range stage.PlantAbstracts {
		if plant.Name == "Vase Trapeze" {
			plant.IsSelected = true
			targetPlant = plant
		} else {
			plant.IsSelected = false
		}
	}
	if targetPlant == nil || targetPlant.TubeVaseAbstract == nil {
		t.Fatalf("Vase Trapeze not found")
	}
	vase := targetPlant.TubeVaseAbstract

	for d := range stage.TubeVase3DDiagrams {
		if d.Name == "Vase Trapeze-TubeVase3DDiagram" {
			d.IsChecked = true
			d.IsHiddenTiledFloor3DShape = true
			d.IsHiddenVaseTrapezeRingShape = false
			d.IsHiddenCarvedOutVaseTrapezeRingShape = false
			d.IsHiddenStackOfRotatedCarvedOutVaseTrapezeRingsShape = false
			d.IsHiddenVaseTrapezeBasePlateShape = false
		}
	}

	buildScene := func() map[string]*threejs.Mesh {
		stager := models.NewStagerForTest(stage)
		stager.EnforceSemanticForTest()
		stager.SetCurrentPlant(targetPlant)
		threejsStage := threejs.NewStage("threejs_curved_wall")
		stager.SetThreejsStage(threejsStage)
		updater := NewThreeJSStageUpdater()
		updater.UpdateThreeJSStage(stager)

		meshes := make(map[string]*threejs.Mesh)
		for m := range threejsStage.Meshs {
			meshes[m.Name] = m
		}
		return meshes
	}

	// Setup uncarved, zero-separation rotated stack for exact contour conservation checks
	vase.CarvedOutTopRingsParameter = 0.0
	vase.RelativeRotatedTorusSeparation = 0.0
	vase.HasRingBottom = false
	vase.RotationRatio = 1.0

	// 1. Parity when HasCurvedWall == false (even if WallCurvature != 0), when WallCurvature == 0,
	//    and enforcement that WallCurvature cannot be above 0 (clamped to 0).
	vase.HasCurvedWall = false
	vase.WallCurvature = 0.0
	meshesStraight := buildScene()
	stlOneStraight := models.GenerateVaseTrapezeOneRingSTL(targetPlant)
	stlBaseStraight := models.GenerateVaseTrapezeBaseAndFirstRingSTL(targetPlant)
	glbStraight := models.GenerateVaseTrapezeGLB(targetPlant)

	vase.HasCurvedWall = false
	vase.WallCurvature = -0.6
	meshesDisabled := buildScene()
	if len(meshesDisabled["Vase Trapeze Vase Trapeze Ring Mesh"].BufferGeometry.Faces) != len(meshesStraight["Vase Trapeze Vase Trapeze Ring Mesh"].BufferGeometry.Faces) {
		t.Fatalf("expected HasCurvedWall=false to ignore WallCurvature")
	}
	if models.GenerateVaseTrapezeOneRingSTL(targetPlant) != stlOneStraight {
		t.Errorf("expected OneRingSTL with HasCurvedWall=false to match straight STL")
	}
	if len(models.GenerateVaseTrapezeGLB(targetPlant)) != len(glbStraight) {
		t.Errorf("expected GLB with HasCurvedWall=false to match straight GLB length")
	}

	vase.HasCurvedWall = true
	vase.WallCurvature = 0.0
	if models.VaseTrapezeWallSubdivisions(targetPlant) != 1 {
		t.Errorf("expected sWall=1 when WallCurvature=0")
	}
	if models.GenerateVaseTrapezeOneRingSTL(targetPlant) != stlOneStraight {
		t.Errorf("expected OneRingSTL with WallCurvature=0 to match straight STL")
	}
	if models.GenerateVaseTrapezeBaseAndFirstRingSTL(targetPlant) != stlBaseStraight {
		t.Errorf("expected BaseAndFirstRingSTL with WallCurvature=0 to match straight STL")
	}

	// Positive WallCurvature (> 0) must be inactive and semantically clamped to 0.0
	vase.HasCurvedWall = true
	vase.WallCurvature = 0.5
	if models.IsVaseTrapezeCurvedWallActive(targetPlant) {
		t.Errorf("expected positive WallCurvature=0.5 to be inactive")
	}
	if models.VaseTrapezeWallSubdivisions(targetPlant) != 1 {
		t.Errorf("expected sWall=1 for positive WallCurvature=0.5")
	}
	meshesClampedPos := buildScene()
	if vase.WallCurvature != 0.0 {
		t.Errorf("expected semantic check to clamp WallCurvature=0.5 to 0.0, got %.6f", vase.WallCurvature)
	}
	if len(meshesClampedPos["Vase Trapeze Vase Trapeze Ring Mesh"].BufferGeometry.Faces) != len(meshesStraight["Vase Trapeze Vase Trapeze Ring Mesh"].BufferGeometry.Faces) {
		t.Errorf("expected clamped positive WallCurvature to produce straight ring mesh")
	}
	if models.GenerateVaseTrapezeOneRingSTL(targetPlant) != stlOneStraight {
		t.Errorf("expected clamped positive WallCurvature to match straight OneRingSTL")
	}

	// Compute geometric reference heights
	sideLen := targetPlant.RhombusSideLength
	scaleY := vase.RibbonVerticalScale
	if scaleY == 0 {
		scaleY = 1.0
	}
	zStart := vase.Z_Ribbon * sideLen
	p1H := zStart + vase.TopPlaneHeight*sideLen*scaleY
	p2H := zStart + vase.BottomPlaneHeight*sideLen*scaleY
	p2HBase := zStart + vase.BaseRingBottomPlaneHeight*sideLen*scaleY
	deltaYContact := models.ComputeVaseTrapezeContactStepDY(targetPlant, vase.HasConstantThickness)
	yContact := p1H - deltaYContact
	yMidArch := p1H - 0.5*deltaYContact

	meanOuterRadiusNearY := func(mesh *threejs.Mesh, prefix string, targetY, tol float64) (float64, int) {
		var sumR float64
		var count int
		for _, v := range mesh.BufferGeometry.Vertices {
			if strings.Contains(v.Name, prefix) && math.Abs(v.Y-targetY) <= tol {
				sumR += math.Hypot(v.X, v.Z)
				count++
			}
		}
		if count == 0 {
			return 0, 0
		}
		return sumR / float64(count), count
	}

	// Helper to fit a 2D circle through 3 points (r1,y1), (r2,y2), (r3,y3)
	circleThrough3Pts := func(r1, y1, r2, y2, r3, y3 float64) (rc, yc, rad float64, ok bool) {
		det := 2.0 * (r1*(y2-y3) + r2*(y3-y1) + r3*(y1-y2))
		if math.Abs(det) < 1e-9 {
			return 0, 0, 0, false
		}
		n1 := r1*r1 + y1*y1
		n2 := r2*r2 + y2*y2
		n3 := r3*r3 + y3*y3
		rc = (n1*(y2-y3) + n2*(y3-y1) + n3*(y1-y2)) / det
		yc = (n1*(r3-r2) + n2*(r1-r3) + n3*(r2-r1)) / det
		rad = math.Hypot(r1-rc, y1-yc)
		return rc, yc, rad, true
	}

	// Reference straight radii on Ring 1 (at dy = deltaYContact)
	ring1Straight := meshesStraight["Vase Trapeze Stack Of Rotated Carved Out Vase Trapeze Rings h1 Mesh"]
	if ring1Straight == nil {
		t.Fatalf("Ring 1 mesh not found in straight scene")
	}
	straightTopR, nTopStraight := meanOuterRadiusNearY(ring1Straight, "Plane 1 Face A", p1H+deltaYContact, 1e-4)
	straightBotR, _ := meanOuterRadiusNearY(ring1Straight, "Plane 2 Face B", p2H+deltaYContact, 1e-4)
	if nTopStraight == 0 {
		t.Fatalf("expected Plane 1 Face A vertices at top of Ring 1")
	}
	uContactOnRing1 := (p1H - (p2H + deltaYContact)) / (p1H - p2H)
	straightContactR := (1.0-uContactOnRing1)*straightBotR + uContactOnRing1*straightTopR
	uMidOnRing1 := ((yMidArch + deltaYContact) - (p2H + deltaYContact)) / (p1H - p2H)
	straightMidR := (1.0-uMidOnRing1)*straightBotR + uMidOnRing1*straightTopR

	// 2. Convex curvatures: Contour Conservation + Constant Curvature Above & Below Contact Line + Non-Collision
	tanAlpha := math.Tan(-vase.ProjectionAngle * math.Pi / 180.0)
	for _, curv := range []float64{-0.34, -0.5, -1.0} {
		vase.HasCurvedWall = true
		vase.WallCurvature = curv
		meshesCurved := buildScene()
		ring1Curved := meshesCurved["Vase Trapeze Stack Of Rotated Carved Out Vase Trapeze Rings h1 Mesh"]
		if ring1Curved == nil {
			t.Fatalf("Ring 1 mesh not found for curvature %.2f", curv)
		}

		// Top contour at y = p1H + deltaYContact must be conserved
		curvedTopR, nTop := meanOuterRadiusNearY(ring1Curved, "Plane 1 Face A", p1H+deltaYContact, 1e-4)
		if nTop != nTopStraight || math.Abs(curvedTopR-straightTopR) > 1e-5 {
			t.Errorf("curv=%.2f: expected conserved top contour meanR %.6f, got %.6f (n=%d)",
				curv, straightTopR, curvedTopR, nTop)
		}

		// Contact line contour at y = p1H on Ring 1 (row s8 of Outer Wall Face) must be conserved
		curvedContactR, nContact := meanOuterRadiusNearY(ring1Curved, "Outer Wall Face s8 ", p1H, 1e-4)
		if nContact != nTopStraight || math.Abs(curvedContactR-straightContactR) > 1e-4 {
			t.Errorf("curv=%.2f: expected conserved contact contour meanR %.6f at y=p1H, got %.6f (n=%d)",
				curv, straightContactR, curvedContactR, nContact)
		}

		// Mid-arch bulge direction check (s16 is midpoint between s8 at yContact and s24 at p1H+dy)
		curvedMidR, nMid := meanOuterRadiusNearY(ring1Curved, "Outer Wall Face s16 ", yMidArch+deltaYContact, 1e-3)
		if nMid == 0 {
			t.Fatalf("curv=%.2f: expected s16 vertices at mid-arch height", curv)
		}
		if !(curvedMidR > straightMidR+5.0) {
			t.Errorf("convex WallCurvature=%.2f should bulge outward: straightMidR=%.2f, curvedMidR=%.2f",
				curv, straightMidR, curvedMidR)
		}

		// Constant Euclidean curvature check across the ENTIRE ring [p2H, p1H] (both above and below yContact)
		rAt := func(localY float64) float64 {
			rStraight := 300.0 + (localY-p2H)*tanAlpha
			return rStraight + models.EvalVaseTrapezeRingWallDeltaR(targetPlant, localY, p1H, p2H, deltaYContact)
		}
		rc, yc, radCircle, ok := circleThrough3Pts(rAt(p2H), p2H, rAt(yContact), yContact, rAt(p1H), p1H)
		if !ok || radCircle <= 0 {
			t.Fatalf("curv=%.2f: failed to fit circle through p2H, yContact, p1H", curv)
		}
		const nSamples = 30
		for k := 0; k <= nSamples; k++ {
			ly := p2H + (float64(k)/float64(nSamples))*(p1H-p2H)
			ry := rAt(ly)
			dist := math.Hypot(ry-rc, ly-yc)
			if math.Abs(dist-radCircle) > 1e-6 {
				t.Errorf("curv=%.2f localY=%.2f: meridian point (%.4f, %.4f) deviates from constant-curvature circle R=%.4f by %.2e",
					curv, ly, ry, ly, radCircle, math.Abs(dist-radCircle))
			}
			dr := models.EvalVaseTrapezeRingWallDeltaR(targetPlant, ly, p1H, p2H, deltaYContact)
			if ly > yContact+1e-3 && ly < p1H-1e-3 && dr <= 0 {
				t.Errorf("curv=%.2f localY=%.2f (above contact): expected dr > 0, got %.6f", curv, ly, dr)
			}
			if ly < yContact-1e-3 && dr >= 0 {
				t.Errorf("curv=%.2f localY=%.2f (below contact): expected dr < 0 (smooth inward arc), got %.6f", curv, ly, dr)
			}
		}

		// C1 slope continuity across yContact
		eps := 1e-4
		slopeBelow := (models.EvalVaseTrapezeRingWallDeltaR(targetPlant, yContact, p1H, p2H, deltaYContact) -
			models.EvalVaseTrapezeRingWallDeltaR(targetPlant, yContact-eps, p1H, p2H, deltaYContact)) / eps
		slopeAbove := (models.EvalVaseTrapezeRingWallDeltaR(targetPlant, yContact+eps, p1H, p2H, deltaYContact) -
			models.EvalVaseTrapezeRingWallDeltaR(targetPlant, yContact, p1H, p2H, deltaYContact)) / eps
		if math.Abs(slopeAbove-slopeBelow) > 1e-4 {
			t.Errorf("curv=%.2f: slope discontinuity across contact line: below=%.6f vs above=%.6f",
				curv, slopeBelow, slopeAbove)
		}

		// Verify actual 3D mesh meridian vertices (azimuth index 0 across rows s=1..23 on standalone Ring) lie on a single circle
		standaloneCurved := meshesCurved["Vase Trapeze Vase Trapeze Ring Mesh"]
		if standaloneCurved == nil {
			t.Fatalf("standalone Vase Trapeze Ring Mesh not found")
		}
		var meridianPts [][2]float64
		for _, v := range standaloneCurved.BufferGeometry.Vertices {
			if strings.HasSuffix(v.Name, " 0") && strings.Contains(v.Name, "Vase Trapeze Vase Trapeze Ring Outer Wall Face s") {
				meridianPts = append(meridianPts, [2]float64{math.Hypot(v.X, v.Z), v.Y})
			}
		}
		if len(meridianPts) < 20 {
			t.Fatalf("curv=%.2f: expected >= 20 meridian vertices on standalone ring outer wall, got %d", curv, len(meridianPts))
		}
		mrc, myc, mrad, mok := circleThrough3Pts(
			meridianPts[0][0], meridianPts[0][1],
			meridianPts[len(meridianPts)/2][0], meridianPts[len(meridianPts)/2][1],
			meridianPts[len(meridianPts)-1][0], meridianPts[len(meridianPts)-1][1],
		)
		if !mok {
			t.Fatalf("curv=%.2f: failed to fit circle to standalone ring mesh meridian vertices", curv)
		}
		for idx, pt := range meridianPts {
			if dErr := math.Abs(math.Hypot(pt[0]-mrc, pt[1]-myc) - mrad); dErr > 1e-5 {
				t.Errorf("curv=%.2f mesh row %d: vertex (r=%.4f, y=%.4f) deviates from constant-curvature circle R=%.4f by %.2e",
					curv, idx, pt[0], pt[1], mrad, dErr)
			}
		}

		// Overlap band [p2H + deltaYContact, p1H] clearance between Ring 0 / Ring k-1 and Ring 1 / Ring k
		yOverlapStart := p2H + deltaYContact
		for k := 0; k <= 20; k++ {
			y := yOverlapStart + (float64(k)/20.0)*(p1H-yOverlapStart)
			dr0 := models.EvalVaseTrapezeRingWallDeltaR(targetPlant, y, p1H, p2H, deltaYContact)
			dr1 := models.EvalVaseTrapezeRingWallDeltaR(targetPlant, y-deltaYContact, p1H, p2H, deltaYContact)
			if dr0+1e-9 < dr1 {
				t.Errorf("curv=%.2f y=%.2f: Ring k outer offset (%.4f) penetrates Ring k-1 inner offset (%.4f)",
					curv, y, dr1, dr0)
			}
			baseInDR := models.EvalVaseTrapezeBaseRingInnerDeltaR(targetPlant, y, p1H, p2HBase, p2H, deltaYContact, true)
			if baseInDR+1e-9 < dr1 {
				t.Errorf("curv=%.2f y=%.2f: Ring 1 outer offset (%.4f) penetrates Ring 0 inner offset (%.4f)",
					curv, y, dr1, baseInDR)
			}
		}

		// Verify Base_Plate C1 junction at p2HBase: deltaR == 0 and positive slope
		drBot, slopeBot := models.EvalVaseTrapezeBaseRingOuterDeltaRAndSlope(targetPlant, p2HBase, p1H, p2HBase, deltaYContact)
		if math.Abs(drBot) > 1e-9 {
			t.Errorf("curv=%.2f: expected zero radial offset at p2HBase, got %.6f", curv, drBot)
		}
		if slopeBot <= 0 {
			t.Errorf("convex curv=%.2f: expected positive d(deltaR)/dy at p2HBase+, got %.6f", curv, slopeBot)
		}
	}

	// 3. Carved-out top + HasRingBottom (plancher) + STL & GLB export validation
	vase.HasCurvedWall = true
	vase.WallCurvature = -0.5
	vase.CarvedOutTopRingsParameter = 0.5
	vase.HasRingBottom = false
	meshesNoBot := buildScene()
	facesNoBot := len(meshesNoBot["Vase Trapeze Carved Out Vase Trapeze Ring Mesh"].BufferGeometry.Faces)
	stlCurvedNoBot := models.GenerateVaseTrapezeOneRingSTL(targetPlant)
	stlBaseCurved := models.GenerateVaseTrapezeBaseAndFirstRingSTL(targetPlant)
	glbCurvedNoBot := models.GenerateVaseTrapezeGLB(targetPlant)

	vase.HasRingBottom = true
	meshesWithBot := buildScene()
	facesWithBot := len(meshesWithBot["Vase Trapeze Carved Out Vase Trapeze Ring Mesh"].BufferGeometry.Faces)
	stlCurvedWithBot := models.GenerateVaseTrapezeOneRingSTL(targetPlant)
	glbCurvedWithBot := models.GenerateVaseTrapezeGLB(targetPlant)

	if facesWithBot != facesNoBot || facesWithBot <= len(meshesStraight["Vase Trapeze Carved Out Vase Trapeze Ring Mesh"].BufferGeometry.Faces) {
		t.Errorf("expected subdivided equal face counts for curved ring with/without bottom, got %d vs %d", facesWithBot, facesNoBot)
	}
	if !strings.Contains(stlCurvedNoBot, "vertex ") || stlCurvedNoBot == stlCurvedWithBot {
		t.Errorf("expected valid distinct OneRing STLs for HasRingBottom=false vs true")
	}
	if !strings.Contains(stlBaseCurved, "vertex ") || stlBaseCurved == stlBaseStraight {
		t.Errorf("expected curved BaseAndFirstRingSTL to be valid and distinct from straight")
	}
	if len(glbCurvedNoBot) < 12 || string(glbCurvedNoBot[:4]) != "glTF" || len(glbCurvedWithBot) <= len(glbCurvedNoBot) {
		t.Errorf("expected valid curved GLB exports (noBot=%d bytes, withBot=%d bytes)", len(glbCurvedNoBot), len(glbCurvedWithBot))
	}

	// 4. Proportional scaling with curved walls enabled (RhombusSideLength 100 -> 200)
	measureMaxBounds := func(side float64) (maxY, maxR float64) {
		targetPlant.RhombusSideLength = side
		meshes := buildScene()
		for _, m := range meshes {
			if m.BufferGeometry == nil {
				continue
			}
			for _, v := range m.BufferGeometry.Vertices {
				if v.Y > maxY {
					maxY = v.Y
				}
				if r := math.Hypot(v.X, v.Z); r > maxR {
					maxR = r
				}
			}
		}
		return
	}
	maxY100, maxR100 := measureMaxBounds(100.0)
	maxY200, maxR200 := measureMaxBounds(200.0)
	if math.Abs(maxY200/maxY100-2.0) > 1e-4 || math.Abs(maxR200/maxR100-2.0) > 1e-4 {
		t.Errorf("expected 2.0x scaling with curved walls (100->200): Y %.2f->%.2f, R %.2f->%.2f",
			maxY100, maxY200, maxR100, maxR200)
	}
}
