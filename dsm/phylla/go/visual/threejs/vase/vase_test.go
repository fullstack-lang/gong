package vase

import (
	"math"
	"path/filepath"
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
