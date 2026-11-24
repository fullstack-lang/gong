package vase

import (
	"math"
	"path/filepath"
	"strconv"
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

	targetPlant.TubeVaseAbstract.HasFillLevelLine = false
	targetPlant.TubeVaseAbstract.HasDrainageHole = false
	targetPlant.TubeVaseAbstract.HasFootRim = false
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
	vase.HasLip = false
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
	vase.HasFillLevelLine = false
	vase.HasDrainageHole = false
	vase.HasFootRim = false
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

func TestVaseTrapezeUnrotatedCurvedWallStacking(t *testing.T) {
	stagePath := filepath.Join("..", "..", "..", "cmd", "phylla", "data", "stage.go")
	stage := models.NewStage("test_unrotated_curved_stacking")
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
			d.IsHiddenStackOfRotatedCarvedOutVaseTrapezeRingsShape = false
			d.IsHiddenVaseTrapezeBasePlateShape = false
		}
	}

	buildScene := func() map[string]*threejs.Mesh {
		stager := models.NewStagerForTest(stage)
		stager.EnforceSemanticForTest()
		stager.SetCurrentPlant(targetPlant)
		threejsStage := threejs.NewStage("threejs_unrot_curved")
		stager.SetThreejsStage(threejsStage)
		updater := NewThreeJSStageUpdater()
		updater.UpdateThreeJSStage(stager)

		meshes := make(map[string]*threejs.Mesh)
		for m := range threejsStage.Meshs {
			meshes[m.Name] = m
		}
		return meshes
	}

	meshYBounds := func(mesh *threejs.Mesh) (minY, maxY float64) {
		minY, maxY = math.MaxFloat64, -math.MaxFloat64
		for _, v := range mesh.BufferGeometry.Vertices {
			if v.Y < minY {
				minY = v.Y
			}
			if v.Y > maxY {
				maxY = v.Y
			}
		}
		return
	}

	// Verify on actual 3D BufferGeometry vertices that at RotationRatio == 0:
	// 1. Ring 1's outer wall does not penetrate Base Plate's inner top rim (at Y = p1H),
	//    nor does Ring 2's outer wall penetrate Ring 1's inner top rim (at Y = p1H + dy1).
	// 2. When expectFull360Flush is true, BOTH minimum AND maximum radial clearance across
	//    all 360° of azimuth are ~0 mm (flush full-circumference contact with zero remaining space).
	checkRimClearanceAndContact := func(label string, meshes map[string]*threejs.Mesh, p1H float64, expectFull360Flush bool) (dy1, dy2 float64) {
		baseMesh := meshes["Vase Trapeze Vase Trapeze Base Plate Mesh"]
		ring1Mesh := meshes["Vase Trapeze Stack Of Rotated Carved Out Vase Trapeze Rings h1 Mesh"]
		ring2Mesh := meshes["Vase Trapeze Stack Of Rotated Carved Out Vase Trapeze Rings h2 Mesh"]
		if baseMesh == nil || ring1Mesh == nil || ring2Mesh == nil {
			t.Fatalf("%s: missing expected meshes (base=%v, h1=%v, h2=%v)", label, baseMesh != nil, ring1Mesh != nil, ring2Mesh != nil)
		}

		_, maxBaseY := meshYBounds(baseMesh)
		_, maxR1Y := meshYBounds(ring1Mesh)
		_, maxR2Y := meshYBounds(ring2Mesh)
		dy1 = maxR1Y - maxBaseY
		dy2 = maxR2Y - maxR1Y

		// Extract Base Plate inner top rim radii by azimuth index i: "Cap Inner Top <i>"
		baseInTopR := make(map[int]float64)
		for _, v := range baseMesh.BufferGeometry.Vertices {
			const marker = "Cap Inner Top "
			if idx := strings.Index(v.Name, marker); idx != -1 {
				var i int
				if _, err := scanIntSuffix(v.Name[idx+len(marker):], &i); err == nil {
					baseInTopR[i] = math.Hypot(v.X, v.Z)
				}
			}
		}
		if len(baseInTopR) < 100 {
			t.Fatalf("%s: expected >= 100 Base Plate Cap Inner Top vertices, got %d", label, len(baseInTopR))
		}

		// Extract Ring 1 outer profile polyline (lower wall + lip collar) per azimuth i and interpolate r_outer_1(i, p1H)
		ring1OuterRows := extractOuterWallAndLipByAzimuth(ring1Mesh)
		minRadialClearance01 := math.MaxFloat64
		maxRadialClearance01 := -math.MaxFloat64
		for i, rIn0 := range baseInTopR {
			pts := ring1OuterRows[i]
			if rOut1AtP1H, ok := interpRAtMonotoneY(pts, p1H); ok {
				clearance := rIn0 - rOut1AtP1H
				if clearance < minRadialClearance01 {
					minRadialClearance01 = clearance
				}
				if clearance > maxRadialClearance01 {
					maxRadialClearance01 = clearance
				}
				if clearance < -0.05 {
					t.Errorf("%s: Ring 1 outer wall penetrates Base Plate inner rim at azimuth %d: rOut1=%.3f > rIn0=%.3f (penetration=%.3f mm)",
						label, i, rOut1AtP1H, rIn0, -clearance)
					break
				}
			}
		}

		// Check Ring 2 outer wall vs Ring 1 inner top rim ("Plane 1 Face B <i>" when bottomIsOuter)
		ring1InTopR := make(map[int]float64)
		for _, v := range ring1Mesh.BufferGeometry.Vertices {
			const marker = "Plane 1 Face B "
			if idx := strings.Index(v.Name, marker); idx != -1 {
				var i int
				if _, err := scanIntSuffix(v.Name[idx+len(marker):], &i); err == nil {
					ring1InTopR[i] = math.Hypot(v.X, v.Z)
				}
			}
		}
		ring2OuterRows := extractOuterWallAndLipByAzimuth(ring2Mesh)
		minRadialClearance12 := math.MaxFloat64
		maxRadialClearance12 := -math.MaxFloat64
		for i, rIn1 := range ring1InTopR {
			pts := ring2OuterRows[i]
			if rOut2AtTop1, ok := interpRAtMonotoneY(pts, maxR1Y); ok {
				clearance := rIn1 - rOut2AtTop1
				if clearance < minRadialClearance12 {
					minRadialClearance12 = clearance
				}
				if clearance > maxRadialClearance12 {
					maxRadialClearance12 = clearance
				}
				if clearance < -0.05 {
					t.Errorf("%s: Ring 2 outer wall penetrates Ring 1 inner rim at azimuth %d: rOut2=%.3f > rIn1=%.3f (penetration=%.3f mm)",
						label, i, rOut2AtTop1, rIn1, -clearance)
					break
				}
			}
		}

		t.Logf("%s: dy1=%.2f mm, dy2=%.2f mm, rimClearance01=[%.4f, %.4f] mm, rimClearance12=[%.4f, %.4f] mm",
			label, dy1, dy2, minRadialClearance01, maxRadialClearance01, minRadialClearance12, maxRadialClearance12)
		if minRadialClearance01 > 0.5 {
			t.Errorf("%s: expected Ring 1 to rest in contact (min clearance <= 0.5 mm), got %.4f mm (floating)", label, minRadialClearance01)
		}
		if minRadialClearance12 > 0.5 {
			t.Errorf("%s: expected Ring 2 to rest in contact (min clearance <= 0.5 mm), got %.4f mm (floating)", label, minRadialClearance12)
		}
		if expectFull360Flush {
			if maxRadialClearance01 > 0.05 {
				t.Errorf("%s: expected flush 360° contact between Base Plate and Ring 1 (max clearance <= 0.05 mm), got %.4f mm gap",
					label, maxRadialClearance01)
			}
			if maxRadialClearance12 > 0.05 {
				t.Errorf("%s: expected flush 360° contact between Ring 1 and Ring 2 (max clearance <= 0.05 mm), got %.4f mm gap",
					label, maxRadialClearance12)
			}
		}
		return dy1, dy2
	}

	// Verify flush 360° unrotated contact across both (N=2,M=3,rep=1) and (N=1,M=3,rep=2)
	for _, cfg := range []struct {
		n, m, reps        int
		bottomPlaneHeight float64
		relVertThickness  float64
	}{
		{n: 2, m: 3, reps: 1, bottomPlaneHeight: 0.22, relVertThickness: 0.06},
		{n: 1, m: 3, reps: 2, bottomPlaneHeight: 0.16, relVertThickness: 0.08},
	} {
		targetPlant.N = cfg.n
		targetPlant.M = cfg.m
		targetPlant.StackHeight = 3
		vase.RadialRepetitions = cfg.reps
		vase.BottomPlaneHeight = cfg.bottomPlaneHeight
		vase.RelativeVerticalThickness = cfg.relVertThickness
		vase.RotationRatio = 0.0
		vase.RelativeRotatedTorusSeparation = 0.0
		vase.CarvedOutTopRingsParameter = 0.0
		vase.HasConstantThickness = true
		vase.HasRingBottom = true
		vase.HasCurvedWall = true
		vase.HasLip = false

		sideLen := targetPlant.RhombusSideLength
		scaleY := vase.RibbonVerticalScale
		p1H := (vase.Z_Ribbon + vase.TopPlaneHeight*scaleY) * sideLen

		var baseDY1 float64
		for idx, curv := range []float64{0.0, -0.15, -0.25, -0.5, -1.0} {
			vase.WallCurvature = curv
			meshes := buildScene()
			label := "N" + strconv.Itoa(cfg.n) + "M" + strconv.Itoa(cfg.m) + "/curv=" + formatFloat2(curv)
			dy1, dy2 := checkRimClearanceAndContact(label, meshes, p1H, true)
			if math.Abs(dy1-dy2) > 0.05 {
				t.Errorf("%s: expected dy1 (%.3f) == dy2 (%.3f) for flush unrotated contact", label, dy1, dy2)
			}
			if idx == 0 {
				baseDY1 = dy1
			} else if math.Abs(dy1-baseDY1) > 0.05 {
				t.Errorf("%s: expected conserved unrotated contact step %.3f mm across curvatures, got %.3f mm",
					label, baseDY1, dy1)
			}

			glb := models.GenerateVaseTrapezeGLB(targetPlant)
			if len(glb) < 28 || string(glb[:4]) != "glTF" {
				t.Fatalf("%s: expected valid GLB at RotationRatio=0", label)
			}
		}
	}

	// Pin (N=1, M=3, RadialRepetitions=2) for transition, plancher, variable-thickness, and lip checks
	targetPlant.N = 1
	targetPlant.M = 3
	targetPlant.StackHeight = 3
	vase.RadialRepetitions = 2
	vase.BottomPlaneHeight = 0.16
	vase.RelativeVerticalThickness = 0.08
	vase.WallCurvature = -0.25
	vase.HasConstantThickness = true
	vase.HasRingBottom = true
	vase.HasCurvedWall = true
	vase.HasLip = false

	sideLen := targetPlant.RhombusSideLength
	scaleY := vase.RibbonVerticalScale
	p1H := (vase.Z_Ribbon + vase.TopPlaneHeight*scaleY) * sideLen

	// Verify smooth intermediate RotationRatio blend (0 < RotationRatio < 1/StackHeight) and STL watertightness
	vase.RotationRatio = 0.15
	wMid := models.VaseTrapezeRotationBlend(targetPlant)
	if !(wMid > 0.1 && wMid < 0.9) {
		t.Errorf("expected intermediate VaseTrapezeRotationBlend in (0.1, 0.9) at RotationRatio=0.15, got %.4f", wMid)
	}
	meshesMid := buildScene()
	checkRimClearanceAndContact("rotRatio=0.15/curv=-0.25", meshesMid, p1H, false)
	assertWatertightASCIISTL(t, "OneRingSTL(rotRatio=0.15)", models.GenerateVaseTrapezeOneRingSTL(targetPlant), true)

	vase.RotationRatio = 0.0
	assertWatertightASCIISTL(t, "OneRingSTL(rotRatio=0.0)", models.GenerateVaseTrapezeOneRingSTL(targetPlant), true)
	assertWatertightASCIISTL(t, "BaseAndFirstRingSTL(rotRatio=0.0)", models.GenerateVaseTrapezeBaseAndFirstRingSTL(targetPlant), false)

	// Verify thick plancher floor contact when BasePlateHeight exceeds wall-wedging step
	vase.WallCurvature = -0.25
	vase.HasRingBottom = true
	savedBasePlateHeight := vase.BasePlateHeight
	vase.BasePlateHeight = 2.0 // 240 mm solid plancher floor inside each ring
	meshesThickBot := buildScene()
	minR2Y, _ := meshYBounds(meshesThickBot["Vase Trapeze Stack Of Rotated Carved Out Vase Trapeze Rings h2 Mesh"])
	minR1Y, _ := meshYBounds(meshesThickBot["Vase Trapeze Stack Of Rotated Carved Out Vase Trapeze Rings h1 Mesh"])
	expectedPlancherThickness := vase.BasePlateHeight * sideLen
	if minR2Y-minR1Y < expectedPlancherThickness-1e-3 {
		t.Errorf("expected Ring 2 bottom (%.2f) not to penetrate Ring 1 solid plancher (minR1Y + %.2f = %.2f)",
			minR2Y, expectedPlancherThickness, minR1Y+expectedPlancherThickness)
	}
	vase.BasePlateHeight = savedBasePlateHeight

	// Verify variable thickness (HasConstantThickness = false) and carved-out top (CarvedOutTopRingsParameter = 0.5)
	vase.HasConstantThickness = false
	vase.CarvedOutTopRingsParameter = 0.5
	vase.WallCurvature = -0.35
	meshesVarCarved := buildScene()
	if meshesVarCarved["Vase Trapeze Stack Of Rotated Carved Out Vase Trapeze Rings h1 Mesh"] == nil {
		t.Fatalf("expected Ring 1 mesh with HasConstantThickness=false and CarvedOutTopRingsParameter=0.5")
	}

	// Verify HasLip = true with curved wall at RotationRatio = 0 (both narrow lip drop-in contact and wide lip-on-lip contact)
	vase.HasConstantThickness = true
	vase.CarvedOutTopRingsParameter = 0.0
	vase.WallCurvature = -0.25
	vase.HasLip = true
	vase.LipHeight = models.DefaultVaseTrapezeLipExtraRatio
	vase.LipWidth = 0.06
	meshesLip := buildScene()
	checkRimClearanceAndContact("curv=-0.25+HasLip", meshesLip, p1H, false)

	vase.LipHeight = 0.10 // Lip extends rotContactDY + 0.10*sideLen, LipWidth wide enough to overhang Base Plate lip top unrotated
	vase.LipWidth = 0.50
	meshesTallLip := buildScene()
	_, maxBaseTallLipY := meshYBounds(meshesTallLip["Vase Trapeze Vase Trapeze Base Plate Mesh"])
	_, maxR1TallLipY := meshYBounds(meshesTallLip["Vase Trapeze Stack Of Rotated Carved Out Vase Trapeze Rings h1 Mesh"])
	expectedLipStep, _, okLip := models.VaseTrapezeLipDimensions(targetPlant)
	if !okLip {
		t.Fatalf("expected active lip dimensions")
	}
	if math.Abs((maxR1TallLipY-maxBaseTallLipY)-expectedLipStep) > 1e-2 {
		t.Errorf("expected tall lip step dy1=%.2f mm (lip-on-lip contact), got %.2f mm",
			expectedLipStep, maxR1TallLipY-maxBaseTallLipY)
	}
}

func extractOuterWallAndLipByAzimuth(mesh *threejs.Mesh) map[int][][2]float64 {
	res := extractWallRowsByAzimuth(mesh, "Outer Wall Face")
	lip := extractWallRowsByAzimuth(mesh, "Lip Collar Face")
	for az, lpts := range lip {
		merged := append(res[az], lpts...)
		for a := 0; a < len(merged); a++ {
			for b := a + 1; b < len(merged); b++ {
				if merged[b][1] < merged[a][1] || (math.Abs(merged[b][1]-merged[a][1]) <= 1e-6 && merged[b][0] < merged[a][0]) {
					merged[a], merged[b] = merged[b], merged[a]
				}
			}
		}
		res[az] = merged
	}
	return res
}

func scanIntSuffix(s string, out *int) (int, error) {
	s = strings.TrimSpace(s)
	val := 0
	if len(s) == 0 {
		return 0, errInvalidInt
	}
	for k := 0; k < len(s); k++ {
		if s[k] < '0' || s[k] > '9' {
			return 0, errInvalidInt
		}
		val = val*10 + int(s[k]-'0')
	}
	*out = val
	return 1, nil
}

type simpleErr string

func (e simpleErr) Error() string { return string(e) }

const errInvalidInt = simpleErr("invalid int")

func extractWallRowsByAzimuth(mesh *threejs.Mesh, stripName string) map[int][][2]float64 {
	res := make(map[int][][2]float64)
	for _, v := range mesh.BufferGeometry.Vertices {
		idx := strings.Index(v.Name, stripName+" ")
		if idx == -1 {
			continue
		}
		rest := v.Name[idx+len(stripName)+1:]
		// Format is either "s<row> <azimuth>" (curved) or "A <azimuth>" / "B <azimuth>" (straight)
		parts := strings.Fields(rest)
		if len(parts) != 2 {
			continue
		}
		if !strings.HasPrefix(parts[0], "s") && parts[0] != "A" && parts[0] != "B" {
			continue
		}
		var az int
		if _, err := scanIntSuffix(parts[1], &az); err != nil {
			continue
		}
		res[az] = append(res[az], [2]float64{math.Hypot(v.X, v.Z), v.Y})
	}
	// Sort each azimuth slice by Y ascending
	for az, pts := range res {
		for a := 0; a < len(pts); a++ {
			for b := a + 1; b < len(pts); b++ {
				if pts[b][1] < pts[a][1] {
					pts[a], pts[b] = pts[b], pts[a]
				}
			}
		}
		res[az] = pts
	}
	return res
}

func interpRAtMonotoneY(pts [][2]float64, targetY float64) (float64, bool) {
	if len(pts) < 2 {
		return 0, false
	}
	if targetY < pts[0][1]-1e-4 || targetY > pts[len(pts)-1][1]+1e-4 {
		return 0, false
	}
	for k := 0; k+1 < len(pts); k++ {
		y0, y1 := pts[k][1], pts[k+1][1]
		if targetY >= y0-1e-4 && targetY <= y1+1e-4 {
			dy := y1 - y0
			if math.Abs(dy) <= 1e-9 {
				return 0.5 * (pts[k][0] + pts[k+1][0]), true
			}
			t := math.Max(0.0, math.Min(1.0, (targetY-y0)/dy))
			return (1.0-t)*pts[k][0] + t*pts[k+1][0], true
		}
	}
	return 0, false
}

func formatFloat2(v float64) string {
	neg := v < 0
	if neg {
		v = -v
	}
	scaled := int(math.Round(v * 100.0))
	whole := scaled / 100
	frac := scaled % 100
	sign := ""
	if neg && scaled > 0 {
		sign = "-"
	}
	d1 := byte('0' + frac/10)
	d2 := byte('0' + frac%10)
	var wholeStr string
	if whole == 0 {
		wholeStr = "0"
	} else {
		var digits []byte
		for whole > 0 {
			digits = append([]byte{byte('0' + whole%10)}, digits...)
			whole /= 10
		}
		wholeStr = string(digits)
	}
	return sign + wholeStr + "." + string([]byte{d1, d2})
}

func TestVaseTrapezeLip(t *testing.T) {
	stagePath := filepath.Join("..", "..", "..", "cmd", "phylla", "data", "stage.go")
	stage := models.NewStage("test_vase_trapeze_lip")
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
		threejsStage := threejs.NewStage("threejs_lip")
		stager.SetThreejsStage(threejsStage)
		updater := NewThreeJSStageUpdater()
		updater.UpdateThreeJSStage(stager)

		meshes := make(map[string]*threejs.Mesh)
		for m := range threejsStage.Meshs {
			meshes[m.Name] = m
		}
		return meshes
	}

	// 1. Default values enforcement when LipHeight / LipWidth start at 0
	vase.HasLip = false
	vase.LipHeight = 0.0
	vase.LipWidth = 0.0
	meshesNoLip := buildScene()
	if math.Abs(vase.LipHeight-models.DefaultVaseTrapezeLipExtraRatio) > 1e-6 || math.Abs(vase.LipWidth-0.12) > 1e-6 {
		t.Fatalf("expected default LipHeight=%.3f and LipWidth=0.12, got LipHeight=%.4f LipWidth=%.4f",
			models.DefaultVaseTrapezeLipExtraRatio, vase.LipHeight, vase.LipWidth)
	}

	stlOneNoLip := models.GenerateVaseTrapezeOneRingSTL(targetPlant)
	stlBaseNoLip := models.GenerateVaseTrapezeBaseAndFirstRingSTL(targetPlant)
	glbNoLip := models.GenerateVaseTrapezeGLB(targetPlant)

	// 2. Parity when HasLip == false even with custom non-zero LipHeight / LipWidth
	vase.HasLip = false
	vase.LipHeight = 0.65
	vase.LipWidth = 0.25
	meshesStillNoLip := buildScene()
	for _, meshName := range []string{
		"Vase Trapeze Vase Trapeze Ring Mesh",
		"Vase Trapeze Carved Out Vase Trapeze Ring Mesh",
		"Vase Trapeze Vase Trapeze Base Plate Mesh",
		"Vase Trapeze Stack Of Rotated Carved Out Vase Trapeze Rings h1 Mesh",
	} {
		m0 := meshesNoLip[meshName]
		m1 := meshesStillNoLip[meshName]
		if m0 == nil || m1 == nil {
			t.Fatalf("expected mesh %q in scene", meshName)
		}
		if len(m0.BufferGeometry.Faces) != len(m1.BufferGeometry.Faces) ||
			len(m0.BufferGeometry.Vertices) != len(m1.BufferGeometry.Vertices) {
			t.Errorf("%s: expected identical geometry when HasLip=false", meshName)
		}
	}
	if models.GenerateVaseTrapezeOneRingSTL(targetPlant) != stlOneNoLip {
		t.Errorf("expected OneRingSTL with HasLip=false to match no-lip baseline")
	}
	if models.GenerateVaseTrapezeBaseAndFirstRingSTL(targetPlant) != stlBaseNoLip {
		t.Errorf("expected BaseAndFirstRingSTL with HasLip=false to match no-lip baseline")
	}
	if len(models.GenerateVaseTrapezeGLB(targetPlant)) != len(glbNoLip) {
		t.Errorf("expected GLB with HasLip=false to match no-lip baseline length")
	}

	// 3. Enable Lip and verify 3D profile, normals, and height/width fine-tuning
	vase.CarvedOutTopRingsParameter = 0.0
	vase.HasRingBottom = true
	vase.HasCurvedWall = true
	vase.WallCurvature = -0.25
	vase.HasLip = false
	baseScene := buildScene()

	sideLen := targetPlant.RhombusSideLength
	scaleY := vase.RibbonVerticalScale
	if scaleY == 0 {
		scaleY = 1.0
	}
	p1H := (vase.Z_Ribbon + vase.TopPlaneHeight*scaleY) * sideLen

	meshMaxR := func(m *threejs.Mesh) float64 {
		maxR := 0.0
		for _, v := range m.BufferGeometry.Vertices {
			if r := math.Hypot(v.X, v.Z); r > maxR {
				maxR = r
			}
		}
		return maxR
	}
	meshMaxY := func(m *threejs.Mesh) float64 {
		maxY := -math.MaxFloat64
		for _, v := range m.BufferGeometry.Vertices {
			if v.Y > maxY {
				maxY = v.Y
			}
		}
		return maxY
	}
	meanRadiusContaining := func(m *threejs.Mesh, substr string) (float64, float64, int) {
		var sumR, sumY float64
		var count int
		for _, v := range m.BufferGeometry.Vertices {
			if strings.Contains(v.Name, substr) {
				sumR += math.Hypot(v.X, v.Z)
				sumY += v.Y
				count++
			}
		}
		if count == 0 {
			return 0, 0, 0
		}
		return sumR / float64(count), sumY / float64(count), count
	}

	noLipRing := baseScene["Vase Trapeze Vase Trapeze Ring Mesh"]
	noLipBase := baseScene["Vase Trapeze Vase Trapeze Base Plate Mesh"]
	noLipH1 := baseScene["Vase Trapeze Stack Of Rotated Carved Out Vase Trapeze Rings h1 Mesh"]
	noLipRingMaxR := meshMaxR(noLipRing)
	noLipBaseMaxR := meshMaxR(noLipBase)
	noLipH1MaxY := meshMaxY(noLipH1)
	noLipInnerBaseR, _, nInnerBase := meanRadiusContaining(noLipBase, "Cap Inner Top ")
	if nInnerBase < 100 {
		t.Fatalf("expected >= 100 Cap Inner Top vertices on Base Plate, got %d", nInnerBase)
	}

	vase.RotationRatio = 1.0
	vase.RelativeRotatedTorusSeparation = 0.0
	vase.HasLip = true
	vase.LipHeight = models.DefaultVaseTrapezeLipExtraRatio // ~5.04 mm at SideLength=120 mm
	vase.LipWidth = 0.15
	lipScene := buildScene()

	lipRing := lipScene["Vase Trapeze Vase Trapeze Ring Mesh"]
	lipBase := lipScene["Vase Trapeze Vase Trapeze Base Plate Mesh"]
	lipH1 := lipScene["Vase Trapeze Stack Of Rotated Carved Out Vase Trapeze Rings h1 Mesh"]

	// With curved walls (sWall = 24), each ring gains (2 + 2*sWall)*M faces (2*M underside + 2*sWall*M subdivided curved collar)
	sWall := models.VaseTrapezeWallSubdivisions(targetPlant)
	expectedExtraFaces := (2 + 2*sWall) * nInnerBase
	if gotExtra := len(lipRing.BufferGeometry.Faces) - len(noLipRing.BufferGeometry.Faces); gotExtra != expectedExtraFaces {
		t.Errorf("standalone ring: expected +%d lip faces, got +%d", expectedExtraFaces, gotExtra)
	}
	if gotExtra := len(lipBase.BufferGeometry.Faces) - len(noLipBase.BufferGeometry.Faces); gotExtra != expectedExtraFaces {
		t.Errorf("base plate: expected +%d lip faces, got +%d", expectedExtraFaces, gotExtra)
	}
	if gotExtra := len(lipH1.BufferGeometry.Faces) - len(noLipH1.BufferGeometry.Faces); gotExtra != expectedExtraFaces {
		t.Errorf("stack ring h1: expected +%d lip faces, got +%d", expectedExtraFaces, gotExtra)
	}

	// Top rim Y is preserved at p1H on Ring 0 & Base Plate while outer radius increases and inner top rim radius stays unchanged
	if math.Abs(meshMaxY(lipRing)-p1H) > 1e-4 || math.Abs(meshMaxY(lipBase)-p1H) > 1e-4 {
		t.Errorf("expected top Y preserved at p1H=%.3f, got ring=%.3f base=%.3f",
			p1H, meshMaxY(lipRing), meshMaxY(lipBase))
	}
	if !(meshMaxR(lipRing) > noLipRingMaxR+vase.LipWidth*sideLen-1.0) {
		t.Errorf("expected standalone ring maxR to grow by at least LipWidth*sideLen (%.2f -> %.2f)",
			noLipRingMaxR, meshMaxR(lipRing))
	}
	if !(meshMaxR(lipBase) > noLipBaseMaxR+vase.LipWidth*sideLen-1.0) {
		t.Errorf("expected base plate maxR to grow by at least LipWidth*sideLen (%.2f -> %.2f)",
			noLipBaseMaxR, meshMaxR(lipBase))
	}
	lipInnerBaseR, _, _ := meanRadiusContaining(lipBase, "Cap Inner Top ")
	if math.Abs(lipInnerBaseR-noLipInnerBaseR) > 1e-6 {
		t.Errorf("expected inner pot rim radius unchanged (%.4f vs %.4f)", noLipInnerBaseR, lipInnerBaseR)
	}

	// Verify Lip starts at top (p1H) and extends down to lower pot edge + ~5mm:
	// Lip Underside Y on standalone ring / Base Plate equals p1H - (rotContactDY + LipHeight*sideLen)
	rotContactDY := models.ComputeVaseTrapezeRotatedContactStepDY(targetPlant, true)
	expectedTotalLipH := rotContactDY + vase.LipHeight*sideLen
	expectedLipBotY := p1H - expectedTotalLipH
	_, ringUnderY, nRingUnder := meanRadiusContaining(lipRing, "Lip Underside Face ")
	_, baseUnderY, nBaseUnder := meanRadiusContaining(lipBase, "Lip Underside ")
	if nRingUnder == 0 || nBaseUnder == 0 {
		t.Fatalf("missing Lip Underside vertices (ring=%d, base=%d)", nRingUnder, nBaseUnder)
	}
	if math.Abs(ringUnderY-expectedLipBotY) > 1e-4 || math.Abs(baseUnderY-expectedLipBotY) > 1e-4 {
		t.Errorf("expected lip underside Y=%.4f (p1H - [rotContactDY=%.2f + extra=%.2f]), got ring=%.4f base=%.4f",
			expectedLipBotY, rotContactDY, vase.LipHeight*sideLen, ringUnderY, baseUnderY)
	}

	// Verify stacked upper pot (h1) sits horizontally on lower pot's top rim (Y == p1H)
	// AND is lifted by LipHeight*sideLen (~5mm) to eliminate wedging stress on the bottom pot!
	_, h1UnderY, nH1Under := meanRadiusContaining(lipH1, "Lip Underside Face ")
	if nH1Under == 0 {
		t.Fatalf("missing Lip Underside Face vertices on stacked Ring h1")
	}
	if math.Abs(h1UnderY-p1H) > 1e-4 {
		t.Errorf("expected stacked Ring h1 Lip Underside to sit flush at lower pot top rim p1H=%.4f, got %.4f", p1H, h1UnderY)
	}
	expectedLift := vase.LipHeight * sideLen
	if gotLift := meshMaxY(lipH1) - noLipH1MaxY; math.Abs(gotLift-expectedLift) > 1e-3 {
		t.Errorf("expected stacked Ring h1 to be lifted by %.3f mm for stress-free wall clearance, got %.3f mm",
			expectedLift, gotLift)
	}

	// Verify face normal orientations on both standalone ring and Base Plate
	checkLipFaceNormals := func(meshLabel string, m *threejs.Mesh, topFaceSubstr string) {
		verts := m.BufferGeometry.Vertices
		var nUnderside, nCollar, nTop int
		for _, f := range m.BufferGeometry.Faces {
			v1, v2, v3 := verts[f.V1], verts[f.V2], verts[f.V3]
			ux, uy, uz := v2.X-v1.X, v2.Y-v1.Y, v2.Z-v1.Z
			wx, wy, wz := v3.X-v1.X, v3.Y-v1.Y, v3.Z-v1.Z
			nx := uy*wz - uz*wy
			ny := uz*wx - ux*wz
			nz := ux*wy - uy*wx
			nLen := math.Sqrt(nx*nx + ny*ny + nz*nz)
			if nLen < 1e-9 {
				t.Errorf("%s: degenerate triangle %q", meshLabel, f.Name)
				continue
			}
			nx, ny, nz = nx/nLen, ny/nLen, nz/nLen
			cx, cz := (v1.X+v2.X+v3.X)/3.0, (v1.Z+v2.Z+v3.Z)/3.0

			switch {
			case strings.Contains(f.Name, "Lip Underside Face"):
				nUnderside++
				if ny > -0.95 {
					t.Errorf("%s %q: expected downward normal ny < -0.95, got %.4f", meshLabel, f.Name, ny)
					return
				}
			case strings.Contains(f.Name, "Lip Collar Face"):
				nCollar++
				if nx*cx+nz*cz <= 0 {
					t.Errorf("%s %q: expected outward radial normal, got dot=%.4f", meshLabel, f.Name, nx*cx+nz*cz)
					return
				}
			case strings.Contains(f.Name, topFaceSubstr):
				nTop++
				if ny < 0.95 {
					t.Errorf("%s %q: expected upward normal ny > 0.95, got %.4f", meshLabel, f.Name, ny)
					return
				}
			}
		}
		if nUnderside == 0 || nCollar == 0 || nTop == 0 {
			t.Errorf("%s: expected non-zero face counts (underside=%d collar=%d top=%d)",
				meshLabel, nUnderside, nCollar, nTop)
		}
	}
	checkLipFaceNormals("Standalone Ring", lipRing, "Plane 1 Face")
	checkLipFaceNormals("Base Plate", lipBase, "Top Face")

	// 4. Fine-tuning LipHeight (extra drop margin) and LipWidth sliders
	vase.LipHeight = 0.03
	vase.LipWidth = 0.10
	sceneSmall := buildScene()
	maxRSmall := meshMaxR(sceneSmall["Vase Trapeze Vase Trapeze Base Plate Mesh"])
	_, underYSmall, _ := meanRadiusContaining(sceneSmall["Vase Trapeze Vase Trapeze Base Plate Mesh"], "Lip Underside ")

	vase.LipHeight = 0.15
	vase.LipWidth = 0.28
	sceneLarge := buildScene()
	maxRLarge := meshMaxR(sceneLarge["Vase Trapeze Vase Trapeze Base Plate Mesh"])
	_, underYLarge, _ := meanRadiusContaining(sceneLarge["Vase Trapeze Vase Trapeze Base Plate Mesh"], "Lip Underside ")

	if !(maxRLarge > maxRSmall+0.17*sideLen) {
		t.Errorf("expected increasing LipWidth (0.10->0.28) to widen outer lip: %.2f -> %.2f", maxRSmall, maxRLarge)
	}
	expectedDrop := (0.15 - 0.03) * sideLen
	if math.Abs((underYSmall-underYLarge)-expectedDrop) > 1e-4 {
		t.Errorf("expected increasing LipHeight extra margin (0.03->0.15) to lower lip underside by %.3f, got %.3f",
			expectedDrop, underYSmall-underYLarge)
	}

	// Verify extreme LipHeight exceeding ring wall height clamps safely without degenerate faces
	vase.LipHeight = 10.0
	sceneClamped := buildScene()
	checkLipFaceNormals("Clamped Tall Lip Ring", sceneClamped["Vase Trapeze Vase Trapeze Ring Mesh"], "Plane 1 Face")
	vase.LipHeight = models.DefaultVaseTrapezeLipExtraRatio
	vase.LipWidth = 0.15

	// 5. Watertight 2-manifold STL verification (with carved top + curved wall + both open & closed ring bottom)
	vase.CarvedOutTopRingsParameter = 0.35
	vase.HasCurvedWall = true
	vase.WallCurvature = -0.30

	vase.HasRingBottom = false
	stlOneOpen := models.GenerateVaseTrapezeOneRingSTL(targetPlant)
	assertWatertightASCIISTL(t, "OneRingSTL(HasRingBottom=false,HasLip=true)", stlOneOpen, true)

	vase.HasRingBottom = true
	stlOneClosed := models.GenerateVaseTrapezeOneRingSTL(targetPlant)
	assertWatertightASCIISTL(t, "OneRingSTL(HasRingBottom=true,HasLip=true)", stlOneClosed, true)

	stlBaseAndFirst := models.GenerateVaseTrapezeBaseAndFirstRingSTL(targetPlant)
	assertWatertightASCIISTL(t, "BaseAndFirstRingSTL(HasLip=true)", stlBaseAndFirst, false)

	// Also verify straight-wall STL watertightness with lip
	vase.HasCurvedWall = false
	assertWatertightASCIISTL(t, "OneRingSTL(Straight,HasLip=true)", models.GenerateVaseTrapezeOneRingSTL(targetPlant), true)
	assertWatertightASCIISTL(t, "BaseAndFirstRingSTL(Straight,HasLip=true)", models.GenerateVaseTrapezeBaseAndFirstRingSTL(targetPlant), false)

	// 6. GLB export & 2.0x proportional scaling with Lip + Curved Wall enabled
	vase.HasCurvedWall = true
	vase.WallCurvature = -0.25
	glbWithLip := models.GenerateVaseTrapezeGLB(targetPlant)
	if len(glbWithLip) <= len(glbNoLip) || string(glbWithLip[:4]) != "glTF" {
		t.Errorf("expected valid larger GLB with HasLip=true (noLip=%d, withLip=%d)", len(glbNoLip), len(glbWithLip))
	}
	if !strings.Contains(string(glbWithLip), "Base_Plate") || !strings.Contains(string(glbWithLip), "Ring_1") {
		t.Errorf("expected GLB with HasLip=true to contain Base_Plate and Ring_1 nodes")
	}

	measureBounds := func(side float64) (maxY, maxR float64) {
		targetPlant.RhombusSideLength = side
		ms := buildScene()
		for _, m := range ms {
			if m.BufferGeometry == nil {
				continue
			}
			if y := meshMaxY(m); y > maxY {
				maxY = y
			}
			if r := meshMaxR(m); r > maxR {
				maxR = r
			}
		}
		return
	}
	y100, r100 := measureBounds(100.0)
	y200, r200 := measureBounds(200.0)
	if math.Abs(y200/y100-2.0) > 1e-4 || math.Abs(r200/r100-2.0) > 1e-4 {
		t.Errorf("expected 2.0x scaling with HasLip=true (100->200): Y %.2f->%.2f, R %.2f->%.2f",
			y100, y200, r100, r200)
	}
}

func TestVaseTrapezeFillLineDrainageAndFootRim(t *testing.T) {
	stagePath := filepath.Join("..", "..", "..", "cmd", "phylla", "data", "stage.go")
	stage := models.NewStage("test_fill_drainage_foot")
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

	// Verify stage.go has the planter details enabled by default
	if !vase.HasFillLevelLine || !vase.HasDrainageHole || !vase.HasFootRim {
		t.Fatalf("expected stage.go Vase Trapeze to enable fill line, drainage hole, and foot rim by default")
	}

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
		threejsStage := threejs.NewStage("threejs_pot_details")
		stager.SetThreejsStage(threejsStage)
		updater := NewThreeJSStageUpdater()
		updater.UpdateThreeJSStage(stager)

		meshes := make(map[string]*threejs.Mesh)
		for m := range threejsStage.Meshs {
			meshes[m.Name] = m
		}
		return meshes
	}

	// 1. Default values enforcement when zeroed
	vase.HasFillLevelLine = false
	vase.HasDrainageHole = false
	vase.HasFootRim = false
	vase.DrainageHoleRadius = 0.0
	vase.FootRimHeight = 0.0
	vase.FootRimWidth = 0.0
	baseOff := buildScene()
	if math.Abs(vase.DrainageHoleRadius-models.DefaultVaseTrapezeDrainageHoleRadiusRatio) > 1e-6 ||
		math.Abs(vase.FootRimHeight-models.DefaultVaseTrapezeFootRimHeightRatio) > 1e-6 ||
		math.Abs(vase.FootRimWidth-models.DefaultVaseTrapezeFootRimWidthRatio) > 1e-6 {
		t.Fatalf("unexpected default planter detail ratios: r=%.4f h=%.4f w=%.4f",
			vase.DrainageHoleRadius, vase.FootRimHeight, vase.FootRimWidth)
	}

	basePlateOff := baseOff["Vase Trapeze Vase Trapeze Base Plate Mesh"]
	stlBaseOff := models.GenerateVaseTrapezeBaseAndFirstRingSTL(targetPlant)

	// 2. Enable all three features on HasRingBottom rings; verify Base Plate stays 100% sealed & unchanged
	vase.HasRingBottom = true
	vase.HasCurvedWall = true
	vase.WallCurvature = -0.18
	vase.HasLip = true
	vase.HasFillLevelLine = true
	vase.HasDrainageHole = true
	vase.HasFootRim = true
	vase.RotationRatio = 1.0
	vase.RelativeRotatedTorusSeparation = 0.0

	sceneOn := buildScene()
	basePlateOn := sceneOn["Vase Trapeze Vase Trapeze Base Plate Mesh"]
	if len(basePlateOn.BufferGeometry.Faces) != len(basePlateOff.BufferGeometry.Faces) ||
		len(basePlateOn.BufferGeometry.Vertices) != len(basePlateOff.BufferGeometry.Vertices) {
		t.Errorf("expected sealed Base Plate mesh to stay unchanged when upper-ring planter details are enabled")
	}
	if models.GenerateVaseTrapezeBaseAndFirstRingSTL(targetPlant) != stlBaseOff {
		t.Errorf("expected BaseAndFirstRingSTL to stay unchanged (sealed Base Plate)")
	}

	// 3. Verify fill line height, independence from RotationRatio & 3D Separation, and face normals
	side := targetPlant.RhombusSideLength
	scaleY := vase.RibbonVerticalScale
	p1H := (vase.Z_Ribbon + vase.TopPlaneHeight*scaleY) * side
	p2H := (vase.Z_Ribbon + vase.BottomPlaneHeight*scaleY) * side
	hRingBottom := vase.BasePlateHeight * side
	yFillExpected, _, _, okFill := models.ComputeVaseTrapezeFillLevelLineParams(targetPlant, p1H, p2H, hRingBottom, true)
	if !okFill {
		t.Fatalf("expected active fill level line params")
	}

	ringMesh := sceneOn["Vase Trapeze Vase Trapeze Ring Mesh"]
	if ringMesh == nil || ringMesh.BufferGeometry == nil {
		t.Fatalf("expected Vase Trapeze Ring Mesh")
	}

	// Verify groove is positioned at yFillExpected and below stacked Ring h1 bottom by ~5mm clearance
	ringH1 := sceneOn["Vase Trapeze Stack Of Rotated Carved Out Vase Trapeze Rings h1 Mesh"]
	minYH1 := math.MaxFloat64
	for _, v := range ringH1.BufferGeometry.Vertices {
		if v.Y < minYH1 {
			minYH1 = v.Y
		}
	}
	expectedClearance := models.DefaultVaseTrapezeFillLineClearanceRatio * side
	if math.Abs((minYH1-yFillExpected)-expectedClearance) > 0.05 {
		t.Errorf("expected fill line to sit %.2f mm below Ring h1 bottom (%.2f vs %.2f, diff=%.2f)",
			expectedClearance, yFillExpected, minYH1, minYH1-yFillExpected)
	}

	// Verify fill line does not drift when scrubbing RotationRatio or 3D Separation
	vase.RotationRatio = 0.0
	vase.RelativeRotatedTorusSeparation = 0.35
	yFillPreview, _, _, okPreview := models.ComputeVaseTrapezeFillLevelLineParams(targetPlant, p1H, p2H, hRingBottom, true)
	if !okPreview || math.Abs(yFillPreview-yFillExpected) > 1e-6 {
		t.Errorf("expected fill line height to remain constant across preview sliders: %.4f vs %.4f", yFillPreview, yFillExpected)
	}
	vase.RotationRatio = 1.0
	vase.RelativeRotatedTorusSeparation = 0.0

	// Verify foot rim resting plane, 3 drainage breather notches, and underside recess
	minStandY := math.MaxFloat64
	maxStandY := -math.MaxFloat64
	nNotchedVerts := 0
	for _, v := range ringMesh.BufferGeometry.Vertices {
		if strings.Contains(v.Name, "Foot Rim Standing Face") {
			if v.Y < minStandY {
				minStandY = v.Y
			}
			if v.Y > maxStandY {
				maxStandY = v.Y
			}
			if v.Y > p2H+0.5 {
				nNotchedVerts++
			}
		}
	}
	expectedFootH := models.ComputeVaseTrapezeFootRimHeight(targetPlant, hRingBottom, true)
	if math.Abs(minStandY-p2H) > 1e-3 {
		t.Errorf("expected foot rim standing surface at p2H=%.3f, got %.3f", p2H, minStandY)
	}
	if maxStandY < p2H+0.80*expectedFootH || nNotchedVerts < 12 {
		t.Errorf("expected 3 drainage breather notches rising to ~%.2f (got maxY=%.2f, notchedVerts=%d)",
			p2H+0.88*expectedFootH, maxStandY, nNotchedVerts)
	}

	// Verify face normals on groove, foot rim, sloped plancher, countersink chamfer, drainage bore, and recessed underside
	verts := ringMesh.BufferGeometry.Vertices
	counts := map[string]int{}
	for _, f := range ringMesh.BufferGeometry.Faces {
		v1, v2, v3 := verts[f.V1], verts[f.V2], verts[f.V3]
		ux, uy, uz := v2.X-v1.X, v2.Y-v1.Y, v2.Z-v1.Z
		wx, wy, wz := v3.X-v1.X, v3.Y-v1.Y, v3.Z-v1.Z
		nx, ny, nz := uy*wz-uz*wy, uz*wx-ux*wz, ux*wy-uy*wx
		nLen := math.Sqrt(nx*nx + ny*ny + nz*nz)
		if nLen < 1e-9 {
			t.Fatalf("degenerate triangle %q", f.Name)
		}
		nx, ny, nz = nx/nLen, ny/nLen, nz/nLen
		switch {
		case strings.Contains(f.Name, "Fill Line Lower Groove Face"):
			counts["fill_lo"]++
			if ny <= 0.1 {
				t.Errorf("%q: expected upward-tilted lower groove normal ny > 0.1, got %.4f", f.Name, ny)
				return
			}
		case strings.Contains(f.Name, "Fill Line Upper Groove Face"):
			counts["fill_hi"]++
			if ny >= -0.1 {
				t.Errorf("%q: expected downward-tilted upper groove normal ny < -0.1, got %.4f", f.Name, ny)
				return
			}
		case strings.Contains(f.Name, "Inner Wall Face"):
			counts["inner_smooth_pocket"]++
		case strings.Contains(f.Name, "Plancher Sloped Face"):
			counts["plancher"]++
			if ny < 0.98 {
				t.Errorf("%q: expected +Y sloped plancher normal ny >= 0.98, got %.4f", f.Name, ny)
				return
			}
		case strings.Contains(f.Name, "Drainage Hole Chamfer Face"):
			counts["chamfer"]++
			if ny < 0.4 || ny > 0.95 {
				t.Errorf("%q: expected ~45° chamfer normal ny in [0.4,0.95], got %.4f", f.Name, ny)
				return
			}
		case strings.Contains(f.Name, "Drainage Hole Bore Face"):
			counts["bore"]++
			if math.Abs(ny) > 1e-3 {
				t.Errorf("%q: expected horizontal bore normal |ny| < 1e-3, got %.4f", f.Name, ny)
				return
			}
		case strings.Contains(f.Name, "Floor Recessed Face"):
			counts["recess"]++
			if ny > -0.99 {
				t.Errorf("%q: expected downward recessed underside normal ny < -0.99, got %.4f", f.Name, ny)
				return
			}
		case strings.Contains(f.Name, "Foot Rim Outer Bevel Face"):
			counts["foot_out"]++
			if ny >= 0 {
				t.Errorf("%q: expected downward-outward foot outer bevel normal ny < 0, got %.4f", f.Name, ny)
				return
			}
		case strings.Contains(f.Name, "Foot Rim Inner Bevel Face"):
			counts["foot_in"]++
			if ny >= 0 {
				t.Errorf("%q: expected downward-inward foot inner bevel normal ny < 0, got %.4f", f.Name, ny)
				return
			}
		case strings.Contains(f.Name, "Foot Rim Standing Face"):
			counts["foot_stand"]++
			if ny > -0.5 {
				t.Errorf("%q: expected downward foot standing normal ny < -0.5, got %.4f", f.Name, ny)
				return
			}
		}
	}
	for _, k := range []string{"fill_lo", "fill_hi", "inner_smooth_pocket", "plancher", "chamfer", "bore", "recess", "foot_out", "foot_in", "foot_stand"} {
		if counts[k] == 0 {
			t.Errorf("expected non-zero face count for %q", k)
		}
	}

	// Verify fill-level groove stops cleanly along the open planting pocket (where there is no pot above):
	// out of M=720 azimuthal columns (1440 triangles for a full-circle strip), the groove should cover the
	// ~65% occupied arc under the upper pot (~850..1050 triangles) and stop completely across the ~30% open pocket.
	if counts["fill_lo"] < 600 || counts["fill_lo"] > 1150 {
		t.Errorf("expected partial-arc fill-level groove stopping in open pocket (600..1150 tris), got %d", counts["fill_lo"])
	}
	if counts["fill_hi"] != counts["fill_lo"] {
		t.Errorf("expected matching lower and upper groove face counts, got %d vs %d", counts["fill_lo"], counts["fill_hi"])
	}

	// Verify every emitted groove vertex sits inside the covered arc (never deep inside the open pocket at ~290°)
	openPocketRad := 290.0 * math.Pi / 180.0
	for _, v := range verts {
		if strings.Contains(v.Name, "Fill Line Lower Groove Face") || strings.Contains(v.Name, "Fill Line Upper Groove Face") {
			th := math.Atan2(v.Z, v.X)
			diff := math.Mod(th-openPocketRad+3.0*math.Pi, 2.0*math.Pi) - math.Pi
			if math.Abs(diff) < 35.0*math.Pi/180.0 {
				t.Fatalf("unexpected fill-level groove vertex %q inside open planting pocket (angle=%.1f°)",
					v.Name, math.Mod(th*180.0/math.Pi+360.0, 360.0))
			}
		}
	}

	// Verify partial groove arc is exact-invariant to 3D Separation and stable (<2°) across RotationRatio morphing
	vase.RelativeRotatedTorusSeparation = 0.35
	sceneSep := buildScene()
	sepFillLo := 0
	for _, f := range sceneSep["Vase Trapeze Vase Trapeze Ring Mesh"].BufferGeometry.Faces {
		if strings.Contains(f.Name, "Fill Line Lower Groove Face") {
			sepFillLo++
		}
	}
	if sepFillLo != counts["fill_lo"] {
		t.Errorf("expected partial fill-level groove count to remain exact-invariant under 3D Separation: %d vs %d",
			sepFillLo, counts["fill_lo"])
	}

	vase.RotationRatio = 0.0
	scenePreview := buildScene()
	previewFillLo := 0
	for _, f := range scenePreview["Vase Trapeze Vase Trapeze Ring Mesh"].BufferGeometry.Faces {
		if strings.Contains(f.Name, "Fill Line Lower Groove Face") {
			previewFillLo++
		}
	}
	if diff := previewFillLo - counts["fill_lo"]; diff < -8 || diff > 8 {
		t.Errorf("expected partial fill-level groove count to stay stable under RotationRatio=0: %d vs %d",
			previewFillLo, counts["fill_lo"])
	}
	vase.RotationRatio = 1.0
	vase.RelativeRotatedTorusSeparation = 0.0

	// 4. Watertight 2-manifold STL checks across all feature combinations
	cases := []struct {
		name             string
		curved, lip      bool
		carved           float64
		fill, hole, foot bool
		holeR, footH, fw float64
	}{
		{"FullCurvedLip", true, true, 0.0, true, true, true, 0.085, 0.035, 0.045},
		{"FullCurvedCarvedLip", true, true, 0.35, true, true, true, 0.085, 0.035, 0.045},
		{"FullStraight", false, false, 0.0, true, true, true, 0.085, 0.035, 0.045},
		{"FillLineOnly", true, true, 0.20, true, false, false, 0.085, 0.035, 0.045},
		{"DrainageHoleOnly", true, true, 0.0, false, true, false, 0.085, 0.035, 0.045},
		{"FootRimOnly", true, true, 0.0, false, false, true, 0.085, 0.035, 0.045},
		{"ExtremeClampedSliders", true, true, 0.40, true, true, true, 0.30, 0.15, 0.20},
	}
	for _, tc := range cases {
		vase.HasCurvedWall = tc.curved
		vase.HasLip = tc.lip
		vase.CarvedOutTopRingsParameter = tc.carved
		vase.HasFillLevelLine = tc.fill
		vase.HasDrainageHole = tc.hole
		vase.HasFootRim = tc.foot
		vase.DrainageHoleRadius = tc.holeR
		vase.FootRimHeight = tc.footH
		vase.FootRimWidth = tc.fw
		stl := models.GenerateVaseTrapezeOneRingSTL(targetPlant)
		assertWatertightASCIISTL(t, "OneRingSTL("+tc.name+")", stl, true)
	}

	// 5. GLB export & 2.0x proportional scaling with all planter details active
	vase.HasCurvedWall = true
	vase.HasLip = true
	vase.CarvedOutTopRingsParameter = 0.0
	vase.HasFillLevelLine = true
	vase.HasDrainageHole = true
	vase.HasFootRim = true
	vase.DrainageHoleRadius = models.DefaultVaseTrapezeDrainageHoleRadiusRatio
	vase.FootRimHeight = models.DefaultVaseTrapezeFootRimHeightRatio
	vase.FootRimWidth = models.DefaultVaseTrapezeFootRimWidthRatio

	glb := models.GenerateVaseTrapezeGLB(targetPlant)
	if len(glb) < 28 || string(glb[:4]) != "glTF" || !strings.Contains(string(glb), "Ring_1") {
		t.Errorf("expected valid GLB with planter details enabled")
	}

	measureBounds := func(sideVal float64) (maxY, maxR float64) {
		targetPlant.RhombusSideLength = sideVal
		ms := buildScene()
		for _, m := range ms {
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
	y100, r100 := measureBounds(100.0)
	y200, r200 := measureBounds(200.0)
	if math.Abs(y200/y100-2.0) > 1e-4 || math.Abs(r200/r100-2.0) > 1e-4 {
		t.Errorf("expected exact 2.0x scaling with planter details (100->200): Y %.2f->%.2f, R %.2f->%.2f",
			y100, y200, r100, r200)
	}
}

func assertWatertightASCIISTL(t *testing.T, label, stl string, strictSingleManifold bool) {
	t.Helper()
	type qVert [3]int64
	type dEdge struct {
		a, b qVert
	}
	quantize := func(x, y, z float64) qVert {
		return qVert{
			int64(math.Round(x * 100.0)),
			int64(math.Round(y * 100.0)),
			int64(math.Round(z * 100.0)),
		}
	}

	edges := make(map[dEdge]int)
	var triVerts []qVert
	var nTris int
	for _, line := range strings.Split(stl, "\n") {
		fields := strings.Fields(line)
		if len(fields) == 4 && fields[0] == "vertex" {
			x, err1 := strconv.ParseFloat(fields[1], 64)
			y, err2 := strconv.ParseFloat(fields[2], 64)
			z, err3 := strconv.ParseFloat(fields[3], 64)
			if err1 != nil || err2 != nil || err3 != nil {
				t.Fatalf("%s: invalid vertex line %q", label, line)
			}
			triVerts = append(triVerts, quantize(x, y, z))
			if len(triVerts) == 3 {
				v0, v1, v2 := triVerts[0], triVerts[1], triVerts[2]
				triVerts = triVerts[:0]
				if v0 == v1 || v1 == v2 || v2 == v0 {
					t.Fatalf("%s: degenerate triangle at %v", label, v0)
				}
				edges[dEdge{v0, v1}]++
				edges[dEdge{v1, v2}]++
				edges[dEdge{v2, v0}]++
				nTris++
			}
		}
	}
	if nTris < 100 {
		t.Fatalf("%s: expected >= 100 triangles in STL, got %d", label, nTris)
	}
	for e, cnt := range edges {
		oppCnt := edges[dEdge{e.b, e.a}]
		if strictSingleManifold {
			if cnt != 1 || oppCnt != 1 {
				t.Fatalf("%s: non-manifold edge %v->%v (count=%d, opposite=%d)", label, e.a, e.b, cnt, oppCnt)
			}
		} else if cnt != oppCnt {
			t.Fatalf("%s: unmatched boundary edge %v->%v (count=%d, opposite=%d)", label, e.a, e.b, cnt, oppCnt)
		}
	}
}
