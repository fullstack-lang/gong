package models_test

import (
	"embed"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models/probe"
	form "github.com/fullstack-lang/gong/lib/form/go/models"
)

func TestPhyllaMergedFormSlidersCheckboxesSelectAndAccordions(t *testing.T) {
	var emptyFS embed.FS
	mux := http.NewServeMux()
	stage := models.NewStage("phylla-test")
	prb := probe.NewProbe(mux, emptyFS, emptyFS, true, stage)

	lib := (&models.Library{
		Name:          "Root",
		IsRootLibrary: true,
	}).Stage(stage)

	vase := (&models.TubeVaseAbstract{
		Name:                  "Trapeze",
		HasConstantThickness:  false,
		HasCurvedWall:         false,
		WallCurvature:         -0.25,
		LipHeight:             0.5,
		LipWidth:              0.2,
		RadialRepetitions:     2,
		ColorCombination:      1,
		RelativeRadialThickness: 0.1,
		RelativeVerticalThickness: 0.1,
	}).Stage(stage)

	plant := (&models.PlantAbstract{
		Name:               "P1",
		N:                  5,
		M:                  8,
		RhombusSideLength:  50,
		RhombusInsideAngle: 60,
		StackHeight:        10,
		PlantType:          models.VaseTrapeze,
		TubeVaseAbstract:   vase,
		CurrentView:        models.VIEW_PLANT_2D,
	}).Stage(stage)
	lib.Plants = append(lib.Plants, plant)

	stager := models.NewStager(mux, stage, prb, "", nil, nil, nil, nil)
	stager.SetCurrentPlant(plant)

	// 1. In VIEW_PLANT_2D, verify PlantType select + Plant Lattice accordion are staged in sliderStage
	stager.UxSlider()
	formGroups := stager.GetSliderStage().GetInstancesSorted[*form.FormGroup]()
	if len(formGroups) != 1 {
		t.Fatalf("expected 1 FormGroup in sliderStage, got %d", len(formGroups))
	}
	fg2D := formGroups[0]
	if !fg2D.HideSubmitButton {
		t.Fatalf("expected HideSubmitButton=true on phylla FormGroup")
	}

	var foundPlantTypeSelect bool
	for _, div := range fg2D.FormDivs {
		if div.Name == "PlantType" && len(div.FormFields) == 1 && div.FormFields[0].FormFieldSelect != nil {
			foundPlantTypeSelect = true
		}
	}
	if !foundPlantTypeSelect {
		t.Fatalf("expected PlantType select in VIEW_PLANT_2D FormGroup")
	}

	// 2. Switch to VIEW_VASE_TRAPEZE_3D and verify VaseTrapeze accordions + real CheckBoxes (no 0..1 bool sliders!)
	plant.CurrentView = models.VIEW_VASE_TRAPEZE_3D
	stager.UxSlider()

	formGroups = stager.GetSliderStage().GetInstancesSorted[*form.FormGroup]()
	fg3D := formGroups[0]

	accordions := map[string]bool{}
	var curvedWallCB *form.CheckBox
	var nSlider *form.FormFieldInt
	var ambientSlider *form.FormFieldFloat64
	for _, div := range fg3D.FormDivs {
		if div.IsAStartAccordionGroup {
			accordions[div.AccordionGroupName] = true
		}
		for _, cb := range div.CheckBoxs {
			if cb.Name == "Curved Wall" {
				curvedWallCB = cb
			}
		}
		for _, ff := range div.FormFields {
			if ff.Name == "N" && ff.FormFieldInt != nil && ff.FormFieldInt.IsSlider {
				nSlider = ff.FormFieldInt
			}
			if ff.Name == "Ambient Light" && ff.FormFieldFloat64 != nil && ff.FormFieldFloat64.IsSlider {
				ambientSlider = ff.FormFieldFloat64
			}
		}
	}

	for _, expectedAcc := range []string{"Plant Lattice", "Vase Core", "Ribbon & Planes", "Bulbous, Wall & Lip", "Fill Line & Drainage", "Lighting", "Cutoff Plane"} {
		if !accordions[expectedAcc] {
			t.Fatalf("expected accordion %q in VaseTrapeze 3D controls, got %v", expectedAcc, accordions)
		}
	}

	var fillLineCB, drainageHoleCB, footRimCB *form.CheckBox
	var holeRadiusSlider, footHeightSlider, footWidthSlider *form.FormFieldFloat64
	for _, div := range fg3D.FormDivs {
		for _, cb := range div.CheckBoxs {
			switch cb.Name {
			case "Fill Level Line":
				fillLineCB = cb
			case "Drainage Hole":
				drainageHoleCB = cb
			case "Foot Rim":
				footRimCB = cb
			}
		}
		for _, ff := range div.FormFields {
			if ff.FormFieldFloat64 != nil && ff.FormFieldFloat64.IsSlider {
				switch ff.Name {
				case "Drainage Hole Radius":
					holeRadiusSlider = ff.FormFieldFloat64
				case "Foot Rim Height":
					footHeightSlider = ff.FormFieldFloat64
				case "Foot Rim Width":
					footWidthSlider = ff.FormFieldFloat64
				}
			}
		}
	}
	if fillLineCB == nil || drainageHoleCB == nil || footRimCB == nil ||
		holeRadiusSlider == nil || footHeightSlider == nil || footWidthSlider == nil {
		t.Fatalf("missing Fill Line & Drainage controls: fill=%v hole=%v foot=%v r=%v h=%v w=%v",
			fillLineCB != nil, drainageHoleCB != nil, footRimCB != nil,
			holeRadiusSlider != nil, footHeightSlider != nil, footWidthSlider != nil)
	}
	if curvedWallCB == nil {
		t.Fatalf("expected real CheckBox for 'Curved Wall' instead of 0..1 slider")
	}
	if nSlider == nil {
		t.Fatalf("expected int slider for 'N'")
	}
	if ambientSlider == nil {
		t.Fatalf("expected float64 slider for 'Ambient Light'")
	}
	if ambientSlider.Value != models.DefaultTubeVase3DAmbientLightLevel {
		t.Fatalf("expected default Ambient Light slider value %.2f, got %.2f", models.DefaultTubeVase3DAmbientLightLevel, ambientSlider.Value)
	}

	frontAmbient := *ambientSlider
	frontAmbient.Value = 0.25
	stager.GetSliderStage().OnAfterUpdateFromFront(ambientSlider, &frontAmbient)
	if plant.TubeVase3DDiagrams[0].AmbientLightLevel != 0.25 {
		t.Fatalf("expected diagram.AmbientLightLevel=0.25 after live slider update, got %.2f", plant.TubeVase3DDiagrams[0].AmbientLightLevel)
	}

	// Simulate live frontend toggle on Curved Wall checkbox & N slider
	frontCB := *curvedWallCB
	frontCB.Value = true
	stager.GetSliderStage().OnAfterUpdateFromFront(curvedWallCB, &frontCB)
	if !vase.HasCurvedWall {
		t.Fatalf("expected vase.HasCurvedWall to update to true via live CheckBox proxy")
	}

	// Setting N=13 when M=8 triggers OnAfterUpdateSliderElement -> enforceSemantic, which swaps N<=M to N=8, M=13
	frontN := *nSlider
	frontN.Value = 13
	stager.GetSliderStage().OnAfterUpdateFromFront(nSlider, &frontN)
	if plant.N != 8 || plant.M != 13 {
		t.Fatalf("expected plant.N=8 and plant.M=13 after live slider update + semantic enforcement, got N=%d M=%d", plant.N, plant.M)
	}
}

func TestTubeVase3DCutoffSliderControlsAndRoundtrip(t *testing.T) {
	var emptyFS embed.FS
	mux := http.NewServeMux()
	stage := models.NewStage("phylla-cutoff-slider-test")
	prb := probe.NewProbe(mux, emptyFS, emptyFS, true, stage)

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
		Name:               "TV1",
		N:                  5,
		M:                  8,
		RhombusSideLength:  50,
		RhombusInsideAngle: 60,
		StackHeight:        5,
		PlantType:          models.TubeVase,
		TubeVaseAbstract:   vase,
		CurrentView:        models.VIEW_PLANT_3D,
	}).Stage(stage)
	lib.Plants = append(lib.Plants, plant)

	stager := models.NewStager(mux, stage, prb, "", nil, nil, nil, nil)
	stager.SetCurrentPlant(plant)

	// Not shown in VIEW_PLANT_3D
	stager.UxSlider()
	fg := stager.GetSliderStage().GetInstancesSorted[*form.FormGroup]()[0]
	for _, div := range fg.FormDivs {
		if div.IsAStartAccordionGroup && div.AccordionGroupName == "Cutoff Plane" {
			t.Fatalf("did not expect Cutoff Plane accordion in VIEW_PLANT_3D")
		}
	}

	// Shown in VIEW_TUBE_VASE_3D
	plant.CurrentView = models.VIEW_TUBE_VASE_3D
	stager.UxSlider()
	fg = stager.GetSliderStage().GetInstancesSorted[*form.FormGroup]()[0]

	var enableCB, helperCB, flipCB *form.CheckBox
	var axisSelect *form.FormFieldSelect
	var offsetSlider *form.FormFieldFloat64

	for _, div := range fg.FormDivs {
		for _, cb := range div.CheckBoxs {
			switch cb.Name {
			case "Enable Cutoff":
				enableCB = cb
			case "Show Cutoff Plane":
				helperCB = cb
			case "Flip Cutoff":
				flipCB = cb
			}
		}
		for _, ff := range div.FormFields {
			if ff.Name == "Cutoff Axis" && ff.FormFieldSelect != nil {
				axisSelect = ff.FormFieldSelect
			}
			if ff.Name == "Cutoff Offset" && ff.FormFieldFloat64 != nil {
				offsetSlider = ff.FormFieldFloat64
			}
		}
	}

	if enableCB == nil || helperCB == nil || flipCB == nil || axisSelect == nil || offsetSlider == nil {
		t.Fatalf("missing Cutoff Plane form controls: enable=%v helper=%v flip=%v axis=%v offset=%v",
			enableCB != nil, helperCB != nil, flipCB != nil, axisSelect != nil, offsetSlider != nil)
	}

	diagram := plant.TubeVase3DDiagrams[0]

	frontEnable := *enableCB
	frontEnable.Value = true
	stager.GetSliderStage().OnAfterUpdateFromFront(enableCB, &frontEnable)

	frontHelper := *helperCB
	frontHelper.Value = true
	stager.GetSliderStage().OnAfterUpdateFromFront(helperCB, &frontHelper)

	frontFlip := *flipCB
	frontFlip.Value = true
	stager.GetSliderStage().OnAfterUpdateFromFront(flipCB, &frontFlip)

	var yOpt *form.Option
	for opt := range *stager.GetSliderStage().GetInstancesSet[*form.Option]() {
		if opt.Name == "Y" {
			yOpt = opt
			break
		}
	}
	if yOpt == nil {
		t.Fatalf("expected Option 'Y' in Cutoff Axis select")
	}
	frontSelect := *axisSelect
	frontSelect.Value = yOpt
	stager.GetSliderStage().OnAfterUpdateFromFront(axisSelect, &frontSelect)

	frontOffset := *offsetSlider
	frontOffset.Value = -0.35
	stager.GetSliderStage().OnAfterUpdateFromFront(offsetSlider, &frontOffset)

	if !diagram.IsCutoffEnabled || !diagram.IsCutoffHelperVisible || !diagram.IsCutoffFlipped ||
		diagram.CutoffAxis != models.CutoffAxisY || diagram.CutoffOffset != -0.35 {
		t.Fatalf("unexpected diagram cutoff state: %+v", diagram)
	}

	// Verify AST Marshall/Parse roundtrip preserves cutoff and ambient light fields
	diagram.AmbientLightLevel = 1.65
	tmpPath := filepath.Join(t.TempDir(), "stage.go")
	stage.MarshallFile(tmpPath, "main")

	stage2 := models.NewStage("phylla-cutoff-roundtrip")
	if err := stage2.ParseAstFile(tmpPath, true); err != nil {
		t.Fatalf("failed to parse marshalled stage: %v", err)
	}
	for d2 := range stage2.TubeVase3DDiagrams {
		if !d2.IsCutoffEnabled || !d2.IsCutoffHelperVisible || !d2.IsCutoffFlipped ||
			d2.CutoffAxis != models.CutoffAxisY || d2.CutoffOffset != -0.35 || d2.AmbientLightLevel != 1.65 {
			t.Fatalf("roundtrip lost cutoff or ambient light state: %+v", d2)
		}
	}
}
