// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/cesium/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Coordinate:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "Coordinate",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CoordinateFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Layer:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "Layer",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LayerFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.MapView:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "MapView",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MapViewFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Marker:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "Marker",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MarkerFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Polygon:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "Polygon",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PolygonFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Polyline:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "Polyline",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PolylineFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}

	if probe.GetCommitMode() {
		formStage.Commit()
	}
}
