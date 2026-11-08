// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/maplibre/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Coordinate:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Longitude", instanceWithInferedType.Longitude, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Latitude", instanceWithInferedType.Latitude, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Altitude", instanceWithInferedType.Altitude, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		AssociationReverseSliceToForm(
			"Polygon",
			"Coordinates",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Polygon) []*models.Coordinate {
				return owner.Coordinates
			})
		AssociationReverseSliceToForm(
			"Polyline",
			"Coordinates",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Polyline) []*models.Coordinate {
				return owner.Coordinates
			})

	case *models.Layer:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsVisible", instanceWithInferedType.IsVisible, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Polygons", instanceWithInferedType, &instanceWithInferedType.Polygons, formGroup, probe)
		AssociationSliceToForm("Polylines", instanceWithInferedType, &instanceWithInferedType.Polylines, formGroup, probe)
		AssociationSliceToForm("Markers", instanceWithInferedType, &instanceWithInferedType.Markers, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		AssociationReverseSliceToForm(
			"MapView",
			"Layers",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.MapView) []*models.Layer {
				return owner.Layers
			})

	case *models.MapView:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StyleURL", instanceWithInferedType.StyleURL, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Longitude", instanceWithInferedType.Longitude, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Latitude", instanceWithInferedType.Latitude, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Zoom", instanceWithInferedType.Zoom, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Pitch", instanceWithInferedType.Pitch, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Bearing", instanceWithInferedType.Bearing, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Layers", instanceWithInferedType, &instanceWithInferedType.Layers, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Marker:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("PopupHTML", instanceWithInferedType.PopupHTML, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Longitude", instanceWithInferedType.Longitude, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Latitude", instanceWithInferedType.Latitude, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Altitude", instanceWithInferedType.Altitude, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("SVG", instanceWithInferedType.SVG, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, true, 300, false)
		BasicFieldtoForm("Color", instanceWithInferedType.Color, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Scale", instanceWithInferedType.Scale, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Rotation", instanceWithInferedType.Rotation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("RotationAlignment", instanceWithInferedType.RotationAlignment, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("PitchAlignment", instanceWithInferedType.PitchAlignment, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeStringToForm("Anchor", instanceWithInferedType.Anchor, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("IsDraggable", instanceWithInferedType.IsDraggable, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		AssociationReverseSliceToForm(
			"Layer",
			"Markers",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Layer) []*models.Marker {
				return owner.Markers
			})

	case *models.Polygon:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("PopupHTML", instanceWithInferedType.PopupHTML, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Coordinates", instanceWithInferedType, &instanceWithInferedType.Coordinates, formGroup, probe)
		BasicFieldtoForm("FillColor", instanceWithInferedType.FillColor, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("FillOpacity", instanceWithInferedType.FillOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StrokeColor", instanceWithInferedType.StrokeColor, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IsExtruded", instanceWithInferedType.IsExtruded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("BaseAltitude", instanceWithInferedType.BaseAltitude, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ExtrusionHeight", instanceWithInferedType.ExtrusionHeight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		AssociationReverseSliceToForm(
			"Layer",
			"Polygons",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Layer) []*models.Polygon {
				return owner.Polygons
			})

	case *models.Polyline:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("PopupHTML", instanceWithInferedType.PopupHTML, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("Coordinates", instanceWithInferedType, &instanceWithInferedType.Coordinates, formGroup, probe)
		BasicFieldtoForm("StrokeColor", instanceWithInferedType.StrokeColor, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StrokeOpacity", instanceWithInferedType.StrokeOpacity, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StrokeWidth", instanceWithInferedType.StrokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("StrokeDashArray", instanceWithInferedType.StrokeDashArray, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		AssociationReverseSliceToForm(
			"Layer",
			"Polylines",
			instanceWithInferedType,
			formGroup,
			probe,
			func(owner *models.Layer) []*models.Polyline {
				return owner.Polylines
			})

	default:
		_ = instanceWithInferedType
	}
}
