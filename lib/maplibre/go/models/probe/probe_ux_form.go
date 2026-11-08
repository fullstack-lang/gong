// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/maplibre/go/models"
)

// ux_form updates the current form if there is one
func (probe *Probe) ux_form() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		if onSave, ok := formGroup.OnSave.(FormCallbackIF); ok {
			if onSave.GetCreationMode() {
				FillUpFormFromGongstructName(probe, onSave.GetGongstructName(), true)
			} else {
				FillUpFormFromGongstruct(onSave.GetInstance(), probe)
			}
		}
	}
}

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "Coordinate":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Coordinate Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CoordinateFormCallback(
			nil,
			probe,
			formGroup,
		)
		coordinate := new(models.Coordinate)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(coordinate, formGroup, probe)
	case "Layer":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Layer Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LayerFormCallback(
			nil,
			probe,
			formGroup,
		)
		layer := new(models.Layer)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(layer, formGroup, probe)
	case "MapView":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "MapView Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MapViewFormCallback(
			nil,
			probe,
			formGroup,
		)
		mapview := new(models.MapView)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(mapview, formGroup, probe)
	case "Marker":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Marker Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MarkerFormCallback(
			nil,
			probe,
			formGroup,
		)
		marker := new(models.Marker)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(marker, formGroup, probe)
	case "Polygon":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Polygon Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PolygonFormCallback(
			nil,
			probe,
			formGroup,
		)
		polygon := new(models.Polygon)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(polygon, formGroup, probe)
	case "Polyline":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Polyline Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PolylineFormCallback(
			nil,
			probe,
			formGroup,
		)
		polyline := new(models.Polyline)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(polyline, formGroup, probe)
	}
	formStage.Commit()
}
