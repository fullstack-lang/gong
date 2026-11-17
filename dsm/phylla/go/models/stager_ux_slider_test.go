package models_test

import (
	"embed"
	"net/http"
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
		}
	}

	for _, expectedAcc := range []string{"Plant Lattice", "Vase Core", "Ribbon & Planes", "Bulbous, Wall & Lip"} {
		if !accordions[expectedAcc] {
			t.Fatalf("expected accordion %q in VaseTrapeze 3D controls, got %v", expectedAcc, accordions)
		}
	}
	if curvedWallCB == nil {
		t.Fatalf("expected real CheckBox for 'Curved Wall' instead of 0..1 slider")
	}
	if nSlider == nil {
		t.Fatalf("expected int slider for 'N'")
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
