package models

import (
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func asciiSTLVolumeCm3(t *testing.T, stl string) float64 {
	t.Helper()
	var vol6 float64
	var tri [3][3]float64
	vi := 0
	for _, line := range strings.Split(stl, "\n") {
		f := strings.Fields(line)
		if len(f) == 4 && f[0] == "vertex" {
			x, err1 := strconv.ParseFloat(f[1], 64)
			y, err2 := strconv.ParseFloat(f[2], 64)
			z, err3 := strconv.ParseFloat(f[3], 64)
			if err1 != nil || err2 != nil || err3 != nil {
				t.Fatalf("invalid STL vertex %q", line)
			}
			tri[vi] = [3]float64{x, y, z}
			vi++
			if vi == 3 {
				v0, v1, v2 := tri[0], tri[1], tri[2]
				vol6 += v0[0]*(v1[1]*v2[2]-v1[2]*v2[1]) +
					v0[1]*(v1[2]*v2[0]-v1[0]*v2[2]) +
					v0[2]*(v1[0]*v2[1]-v1[1]*v2[0])
				vi = 0
			}
		}
	}
	return math.Abs(vol6) / 6000.0
}

func TestVaseTrapezePotWeightAndThicknessInsights(t *testing.T) {
	stagePath := filepath.Join("..", "cmd", "phylla", "data", "stage.go")
	if _, err := os.Stat(stagePath); os.IsNotExist(err) {
		t.Skipf("stage.go not found at %s", stagePath)
	}

	stage := NewStage("test_vase_trapeze_insights")
	if err := stage.ParseAstFile(stagePath, true); err != nil {
		t.Fatalf("failed to parse stage.go: %v", err)
	}
	stage.ComputeReverseMaps()
	stage.ComputeInstancesNb()
	stage.ComputeReferenceAndOrders()

	var plant *PlantAbstract
	for p := range stage.PlantAbstracts {
		if p.Name == "Vase Trapeze" {
			p.IsSelected = true
			plant = p
		} else {
			p.IsSelected = false
		}
	}
	if plant == nil || plant.TubeVaseAbstract == nil {
		t.Fatalf("Vase Trapeze not found in stage.go")
	}
	vase := plant.TubeVaseAbstract

	stager := NewStagerForTest(stage)
	stager.enforceSemantic()

	weightKg, volCm3, outerDiamMM, potHeightMM, hasBot, ok := ComputeVaseTrapezePotWeightInsight(plant)
	if !ok || !hasBot {
		t.Fatalf("expected valid closed-bottom pot weight insight, got ok=%v hasBot=%v", ok, hasBot)
	}

	// 1. Exact volume parity with GenerateVaseTrapezeOneRingSTL at seated rotated geometry
	origRot := vase.RotationRatio
	vase.RotationRatio = 1.0
	stlVolCm3 := asciiSTLVolumeCm3(t, GenerateVaseTrapezeOneRingSTL(plant))
	vase.RotationRatio = origRot

	if math.Abs(volCm3-stlVolCm3) > 1e-2 {
		t.Errorf("expected insight solid volume (%.4f cm³) to match OneRing STL volume (%.4f cm³)", volCm3, stlVolCm3)
	}
	expectedWeight := volCm3 * StandardTerracottaDensityGPerCm3 / 1000.0
	if math.Abs(weightKg-expectedWeight) > 1e-9 {
		t.Errorf("expected weight %.6f kg, got %.6f kg", expectedWeight, weightKg)
	}
	if weightKg < 8.0 || weightKg > 20.0 {
		t.Errorf("expected realistic terracotta terrace pot weight in [8, 20] kg, got %.2f kg", weightKg)
	}
	t.Logf("Pot weight: %.2f kg (solid: %.2f cm³, Øout: %.1f mm, H: %.1f mm)", weightKg, volCm3, outerDiamMM, potHeightMM)

	// 2. Thickness insights on current stage.go Vase Trapeze
	lipThick, stdLip, lipDiff, lipPct,
		wallThick, lipOverhang, lipCollarH,
		botSlab, stdBot, botDiff, botPct,
		botDrain, botGross,
		hasLip, hasRingBot, okThick := ComputeVaseTrapezePotThicknessInsights(plant)
	if !okThick || !hasLip || !hasRingBot {
		t.Fatalf("expected valid thickness insights with lip and bottom enabled")
	}
	t.Logf("Lip: %.2f mm vs %.2f mm std (%+.2f mm, %+.1f%%) [wall=%.2f, overhang=%.2f, collarH=%.1f]",
		lipThick, stdLip, lipDiff, lipPct, wallThick, lipOverhang, lipCollarH)
	t.Logf("Bottom: %.2f mm vs %.2f mm std (%+.2f mm, %+.1f%%) [drain=%.2f, gross=%.2f]",
		botSlab, stdBot, botDiff, botPct, botDrain, botGross)

	if lipDiff <= 0 || botDiff <= 5.0 {
		t.Errorf("expected current stage.go lip (%+.2f mm) and bottom (%+.2f mm) to exceed standard thickness", lipDiff, botDiff)
	}
	if botDrain >= botSlab || botGross <= botSlab {
		t.Errorf("expected drain thickness (%.2f) < net slab (%.2f) < gross bottom (%.2f)", botDrain, botSlab, botGross)
	}

	// 3. Preview slider invariance (RotationRatio and RelativeRotatedTorusSeparation)
	vase.RotationRatio = 0.0
	vase.RelativeRotatedTorusSeparation = 0.40
	wPreview, vPreview, _, _, _, okPrev := ComputeVaseTrapezePotWeightInsight(plant)
	if !okPrev || math.Abs(wPreview-weightKg) > 1e-9 || math.Abs(vPreview-volCm3) > 1e-6 {
		t.Errorf("expected pot weight/volume to stay invariant across preview sliders: %.4f/%.2f vs %.4f/%.2f",
			wPreview, vPreview, weightKg, volCm3)
	}
	vase.RotationRatio = origRot
	vase.RelativeRotatedTorusSeparation = 0.0

	// 4. Monotonic response when thinning bottom and lip, and exact L^3 mass scaling
	origBaseH := vase.BasePlateHeight
	vase.BasePlateHeight = 0.16
	wThinBot, _, _, _, _, _ := ComputeVaseTrapezePotWeightInsight(plant)
	_, _, _, _, _, _, _, botSlabThin, _, botDiffThin, _, _, _, _, _, _ := ComputeVaseTrapezePotThicknessInsights(plant)
	vase.BasePlateHeight = origBaseH
	if wThinBot >= weightKg || botSlabThin >= botSlab || math.Abs(botDiffThin) >= math.Abs(botDiff) {
		t.Errorf("expected thinner BasePlateHeight=0.16 to reduce pot weight (%.2f -> %.2f kg) and move closer to std bottom (%.2f -> %.2f mm)",
			weightKg, wThinBot, botDiff, botDiffThin)
	}

	// Missing bottom edge case
	vase.HasRingBottom = false
	_, _, _, _, hasBotOff, okOff := ComputeVaseTrapezePotWeightInsight(plant)
	_, _, _, _, _, _, _, botSlabOff, _, botDiffOff, botPctOff, _, _, _, hasRingBotOff, _ := ComputeVaseTrapezePotThicknessInsights(plant)
	vase.HasRingBottom = true
	if !okOff || hasBotOff || hasRingBotOff || botSlabOff != 0 || botDiffOff >= 0 || botPctOff != -100.0 {
		t.Errorf("expected HasRingBottom=false to report 0 mm floor (-100%% vs std), got slab=%.2f diff=%.2f pct=%.1f",
			botSlabOff, botDiffOff, botPctOff)
	}

	// Exact 8.0x volume/weight scaling when RhombusSideLength doubles (100 -> 200)
	origSide := plant.RhombusSideLength
	plant.RhombusSideLength = 100.0
	stager.enforceSemantic()
	w100, v100, _, _, _, _ := ComputeVaseTrapezePotWeightInsight(plant)

	plant.RhombusSideLength = 200.0
	stager.enforceSemantic()
	w200, v200, _, _, _, _ := ComputeVaseTrapezePotWeightInsight(plant)

	plant.RhombusSideLength = origSide
	stager.enforceSemantic()

	if math.Abs(v200/v100-8.0) > 1e-3 || math.Abs(w200/w100-8.0) > 1e-3 {
		t.Errorf("expected 8.0x cubic scaling of pot volume & weight when SideLength doubles (100->200), got vRatio=%.5f wRatio=%.5f",
			v200/v100, w200/w100)
	}
}
