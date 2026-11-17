package models_test

import (
	"embed"
	"net/http"
	"testing"

	"github.com/fullstack-lang/gong/dsm/floss/go/models"
	"github.com/fullstack-lang/gong/dsm/floss/go/models/probe"
	form "github.com/fullstack-lang/gong/lib/form/go/models"
)

func TestFlossMergedFormSlidersAndAccordions(t *testing.T) {
	var emptyFS embed.FS
	mux := http.NewServeMux()
	stage := models.NewStage("floss-test")
	prb := probe.NewProbe(mux, emptyFS, emptyFS, true, stage)

	lib := (&models.Library{
		Name:          "Root",
		IsRootLibrary: true,
	}).Stage(stage)

	comp := (&models.Complexity{
		Name:     "C1",
		Strength: 10.0,
	}).Stage(stage)

	subComp := (&models.Complexity{
		Name:     "SubC",
		Strength: 5.0,
	}).Stage(stage)

	perf := (&models.Performance{
		Name:     "P1",
		Strength: 20.0,
	}).Stage(stage)

	eff := (&models.Effort{
		Name:     "E1",
		Strength: 15.0,
	}).Stage(stage)

	diag := (&models.DiagramFlossEquation{
		Name:                 "Eq1",
		IsChecked:            true,
		Scale:                2.0,
		AreSubsystemsVisible: false,
	}).Stage(stage)

	subSys := (&models.System{
		Name:         "SubSys1",
		Complexities: []*models.Complexity{subComp},
	}).Stage(stage)

	sys := (&models.System{
		Name:                            "Sys1",
		Complexities:                    []*models.Complexity{comp},
		Performances:                    []*models.Performance{perf},
		Efforts:                         []*models.Effort{eff},
		SubSystems:                      []*models.System{subSys},
		AreCPEsCompoundedFromSubSystems: true,
		DiagramFlossEquations:           []*models.DiagramFlossEquation{diag},
	}).Stage(stage)
	lib.RootSystems = append(lib.RootSystems, sys)

	stager := models.NewStager(mux, stage, prb)

	formGroups := stager.GetSliderStage().GetInstancesSorted[*form.FormGroup]()
	if len(formGroups) != 1 {
		t.Fatalf("expected 1 FormGroup in floss sliderStage, got %d", len(formGroups))
	}
	fg := formGroups[0]
	if !fg.HideSubmitButton {
		t.Fatalf("expected HideSubmitButton=true on floss FormGroup")
	}

	accordions := map[string]bool{}
	var c1Field *form.FormField
	var scaleField *form.FormFieldFloat64
	for _, div := range fg.FormDivs {
		if div.IsAStartAccordionGroup {
			accordions[div.AccordionGroupName] = true
		}
		for _, ff := range div.FormFields {
			if ff.Name == "C1" {
				c1Field = ff
			}
			if ff.Name == "Scale" && ff.FormFieldFloat64 != nil {
				scaleField = ff.FormFieldFloat64
			}
		}
	}

	for _, acc := range []string{"Diagram", "Complexities", "Performances", "Efforts"} {
		if !accordions[acc] {
			t.Fatalf("expected accordion %q in floss controls, got %v", acc, accordions)
		}
	}
	if c1Field == nil || !c1Field.IsDisabled {
		t.Fatalf("expected C1 slider to exist and be IsDisabled=true when compounded and subsystems hidden, got %+v", c1Field)
	}
	if scaleField == nil || !scaleField.IsSlider {
		t.Fatalf("expected Scale float64 slider")
	}

	// Simulate live slider update on Scale
	frontScale := *scaleField
	frontScale.Value = 4.5
	stager.GetSliderStage().OnAfterUpdateFromFront(scaleField, &frontScale)
	if diag.Scale != 4.5 {
		t.Fatalf("expected diag.Scale=4.5 after live slider update, got %f", diag.Scale)
	}
}
