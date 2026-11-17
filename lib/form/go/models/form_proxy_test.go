package models_test

import (
	"net/http"
	"testing"

	"github.com/fullstack-lang/gong/lib/form/go/fullstack"
	form "github.com/fullstack-lang/gong/lib/form/go/models"
)

func TestMergedFormSlidersCheckboxesSelectAndAccordions(t *testing.T) {
	mux := http.NewServeMux()
	stage, _ := fullstack.NewStackInstance(mux, "test-merged-form")

	intVal := 5
	floatVal := 0.25
	boolVal := false
	radioVal := true
	selectedEnum := "Plant"

	updateCount := 0
	onUpdate := func() {
		updateCount++
	}

	fg := form.NewLiveFormGroup(stage, "Test Controls")
	if !fg.HideSubmitButton {
		t.Fatalf("expected HideSubmitButton to be true on live FormGroup")
	}

	accDiv := form.StartAccordion(stage, fg, "Geometry", true)
	if !accDiv.IsAStartAccordionGroup || accDiv.AccordionGroupName != "Geometry" || !accDiv.IsExpandedInitially {
		t.Fatalf("unexpected accordion start div state: %+v", accDiv)
	}

	intField := form.AppendSlider(stage, fg, "N", 1, 20, 1, &intVal, onUpdate)
	if intField.FormFieldInt == nil || !intField.FormFieldInt.IsSlider || intField.FormFieldInt.Step != 1 || !intField.FormFieldInt.IsLiveUpdate {
		t.Fatalf("unexpected int slider field: %+v", intField.FormFieldInt)
	}

	floatField := form.AppendSlider(stage, fg, "Thickness", 0.01, 1.0, 0.05, &floatVal, onUpdate)
	floatField.IsDisabled = true
	if floatField.FormFieldFloat64 == nil || !floatField.FormFieldFloat64.IsSlider || floatField.FormFieldFloat64.Step != 0.05 || !floatField.IsDisabled {
		t.Fatalf("unexpected float64 slider field: %+v", floatField.FormFieldFloat64)
	}

	cb := form.AppendCheckBox(stage, fg, "Curved Wall", &boolVal, onUpdate)
	if cb.Value != false || !cb.IsLiveUpdate {
		t.Fatalf("unexpected checkbox state: %+v", cb)
	}

	radioCb := form.AppendRadioCheckBox(stage, fg, "Scale", "Minor", "Major", &radioVal, onUpdate)
	if radioCb.LabelForTrue != "Minor" || radioCb.LabelForFalse != "Major" || !radioCb.Value {
		t.Fatalf("unexpected radio checkbox state: %+v", radioCb)
	}

	selField := form.AppendSelect(stage, fg, "PlantType", "Plant Type", []string{"Plant", "TubeVase", "Stool"}, selectedEnum, func(sel string) {
		selectedEnum = sel
		updateCount++
	})
	if selField.FormFieldSelect == nil || selField.FormFieldSelect.Value == nil || selField.FormFieldSelect.Value.Name != "Plant" {
		t.Fatalf("unexpected select field state: %+v", selField.FormFieldSelect)
	}

	endDiv := form.EndAccordion(stage, fg)
	if !endDiv.IsAEndAccordionGroup {
		t.Fatalf("expected IsAEndAccordionGroup to be true")
	}

	stage.Commit()

	// Simulate frontend updates via OnAfterUpdateFromFront
	frontInt := *intField.FormFieldInt
	frontInt.Value = 13
	stage.OnAfterUpdateFromFront(intField.FormFieldInt, &frontInt)
	if intVal != 13 || updateCount != 1 {
		t.Fatalf("expected intVal=13 and updateCount=1, got intVal=%d updateCount=%d", intVal, updateCount)
	}

	frontFloat := *floatField.FormFieldFloat64
	frontFloat.Value = 0.75
	stage.OnAfterUpdateFromFront(floatField.FormFieldFloat64, &frontFloat)
	if floatVal != 0.75 || updateCount != 2 {
		t.Fatalf("expected floatVal=0.75 and updateCount=2, got floatVal=%f updateCount=%d", floatVal, updateCount)
	}

	frontCb := *cb
	frontCb.Value = true
	stage.OnAfterUpdateFromFront(cb, &frontCb)
	if !boolVal || updateCount != 3 {
		t.Fatalf("expected boolVal=true and updateCount=3, got boolVal=%v updateCount=%d", boolVal, updateCount)
	}

	frontRadio := *radioCb
	frontRadio.Value = false
	stage.OnAfterUpdateFromFront(radioCb, &frontRadio)
	if radioVal || updateCount != 4 {
		t.Fatalf("expected radioVal=false and updateCount=4, got radioVal=%v updateCount=%d", radioVal, updateCount)
	}

	frontSel := *selField.FormFieldSelect
	frontSel.Value = selField.FormFieldSelect.Options[2] // "Stool"
	stage.OnAfterUpdateFromFront(selField.FormFieldSelect, &frontSel)
	if selectedEnum != "Stool" || updateCount != 5 {
		t.Fatalf("expected selectedEnum=Stool and updateCount=5, got selectedEnum=%q updateCount=%d", selectedEnum, updateCount)
	}
}
