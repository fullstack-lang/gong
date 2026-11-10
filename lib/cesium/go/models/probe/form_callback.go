// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"

	"github.com/fullstack-lang/gong/lib/cesium/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

type FormCallbackIF = gongprobe.FormCallbackIF


type FormCallback[T models.PointerToGongstruct] struct {
	Instance     T
	CreationMode bool
	probe        *Probe
	formGroup    *form.FormGroup
	saveFields   func(instance T, probe *Probe, formGroup *form.FormGroup)
}

func NewFormCallback[T models.PointerToGongstruct](
	instance T,
	probe *Probe,
	formGroup *form.FormGroup,
	saveFields func(instance T, probe *Probe, formGroup *form.FormGroup),
) *FormCallback[T] {
	var zero T
	return &FormCallback[T]{
		Instance:     instance,
		CreationMode: instance == zero,
		probe:        probe,
		formGroup:    formGroup,
		saveFields:   saveFields,
	}
}

func (cb *FormCallback[T]) GetCreationMode() bool     { return cb.CreationMode }
func (cb *FormCallback[T]) GetInstance() any           { return cb.Instance }
func (cb *FormCallback[T]) GetGongstructName() string { return models.GetPointerToGongstructName[T]() }

func (cb *FormCallback[T]) OnSave() {
	cb.probe.stageOfInterest.Lock()
	defer cb.probe.stageOfInterest.Unlock()

	cb.probe.formStage.Checkout()

	var zero T
	if cb.Instance == zero {
		cb.Instance = cb.probe.stageOfInterest.GongNewInstance[T]()
	}

	cb.saveFields(cb.Instance, cb.probe, cb.formGroup)

	if cb.formGroup.HasSuppressButtonBeenPressed {
		cb.Instance.UnstageVoid(cb.probe.stageOfInterest)
	}

	cb.probe.stageOfInterest.Commit()
	updateProbeTable[T](cb.probe)

	if cb.CreationMode || cb.formGroup.HasSuppressButtonBeenPressed {
		cb.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(cb.probe.formStage)
		newFormGroup.OnSave = NewFormCallback[T](
			*new(T),
			cb.probe,
			newFormGroup,
			cb.saveFields,
		)
		newInstance := models.GongNewInstance[T]()
		FillUpForm(newInstance, newFormGroup, cb.probe)
		cb.probe.formStage.Commit()
	}

	cb.probe.ux_tree()
}

// insertion point
func __gong__New__CoordinateFormCallback(
	_instance *models.Coordinate,
	probe *Probe,
	formGroup *form.FormGroup,
) (coordinateFormCallback *FormCallback[*models.Coordinate]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCoordinateFields,
	)
}

type CoordinateFormCallback = FormCallback[*models.Coordinate]

func saveCoordinateFields(
	_instance *models.Coordinate,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Longitude":
			FormDivBasicFieldToField(&(_instance.Longitude), formDiv)
		case "Latitude":
			FormDivBasicFieldToField(&(_instance.Latitude), formDiv)
		case "Altitude":
			FormDivBasicFieldToField(&(_instance.Altitude), formDiv)
		case "Polygon:Coordinates":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Coordinates", func(owner *models.Polygon) *[]*models.Coordinate { return &owner.Coordinates })
		case "Polyline:Coordinates":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Coordinates", func(owner *models.Polyline) *[]*models.Coordinate { return &owner.Coordinates })
		}
	}
}

func __gong__New__LayerFormCallback(
	_instance *models.Layer,
	probe *Probe,
	formGroup *form.FormGroup,
) (layerFormCallback *FormCallback[*models.Layer]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLayerFields,
	)
}

type LayerFormCallback = FormCallback[*models.Layer]

func saveLayerFields(
	_instance *models.Layer,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "IsVisible":
			FormDivBasicFieldToField(&(_instance.IsVisible), formDiv)
		case "Polygons":
			FormDivSliceOfPointersToField(_instance, "Polygons", &(_instance.Polygons), formDiv, probe)
		case "Polylines":
			FormDivSliceOfPointersToField(_instance, "Polylines", &(_instance.Polylines), formDiv, probe)
		case "Markers":
			FormDivSliceOfPointersToField(_instance, "Markers", &(_instance.Markers), formDiv, probe)
		case "MapView:Layers":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Layers", func(owner *models.MapView) *[]*models.Layer { return &owner.Layers })
		}
	}
}

func __gong__New__MapViewFormCallback(
	_instance *models.MapView,
	probe *Probe,
	formGroup *form.FormGroup,
) (mapviewFormCallback *FormCallback[*models.MapView]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMapViewFields,
	)
}

type MapViewFormCallback = FormCallback[*models.MapView]

func saveMapViewFields(
	_instance *models.MapView,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StyleURL":
			FormDivBasicFieldToField(&(_instance.StyleURL), formDiv)
		case "Longitude":
			FormDivBasicFieldToField(&(_instance.Longitude), formDiv)
		case "Latitude":
			FormDivBasicFieldToField(&(_instance.Latitude), formDiv)
		case "Zoom":
			FormDivBasicFieldToField(&(_instance.Zoom), formDiv)
		case "Pitch":
			FormDivBasicFieldToField(&(_instance.Pitch), formDiv)
		case "Bearing":
			FormDivBasicFieldToField(&(_instance.Bearing), formDiv)
		case "Layers":
			FormDivSliceOfPointersToField(_instance, "Layers", &(_instance.Layers), formDiv, probe)
		}
	}
}

func __gong__New__MarkerFormCallback(
	_instance *models.Marker,
	probe *Probe,
	formGroup *form.FormGroup,
) (markerFormCallback *FormCallback[*models.Marker]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMarkerFields,
	)
}

type MarkerFormCallback = FormCallback[*models.Marker]

func saveMarkerFields(
	_instance *models.Marker,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(_instance.Label), formDiv)
		case "PopupHTML":
			FormDivBasicFieldToField(&(_instance.PopupHTML), formDiv)
		case "Longitude":
			FormDivBasicFieldToField(&(_instance.Longitude), formDiv)
		case "Latitude":
			FormDivBasicFieldToField(&(_instance.Latitude), formDiv)
		case "Altitude":
			FormDivBasicFieldToField(&(_instance.Altitude), formDiv)
		case "SVG":
			FormDivBasicFieldToField(&(_instance.SVG), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "Scale":
			FormDivBasicFieldToField(&(_instance.Scale), formDiv)
		case "Rotation":
			FormDivBasicFieldToField(&(_instance.Rotation), formDiv)
		case "RotationAlignment":
			FormDivEnumStringFieldToField(&(_instance.RotationAlignment), formDiv)
		case "PitchAlignment":
			FormDivEnumStringFieldToField(&(_instance.PitchAlignment), formDiv)
		case "Anchor":
			FormDivEnumStringFieldToField(&(_instance.Anchor), formDiv)
		case "IsDraggable":
			FormDivBasicFieldToField(&(_instance.IsDraggable), formDiv)
		case "Layer:Markers":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Markers", func(owner *models.Layer) *[]*models.Marker { return &owner.Markers })
		}
	}
}

func __gong__New__PolygonFormCallback(
	_instance *models.Polygon,
	probe *Probe,
	formGroup *form.FormGroup,
) (polygonFormCallback *FormCallback[*models.Polygon]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePolygonFields,
	)
}

type PolygonFormCallback = FormCallback[*models.Polygon]

func savePolygonFields(
	_instance *models.Polygon,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(_instance.Label), formDiv)
		case "PopupHTML":
			FormDivBasicFieldToField(&(_instance.PopupHTML), formDiv)
		case "Coordinates":
			FormDivSliceOfPointersToField(_instance, "Coordinates", &(_instance.Coordinates), formDiv, probe)
		case "FillColor":
			FormDivBasicFieldToField(&(_instance.FillColor), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(_instance.FillOpacity), formDiv)
		case "StrokeColor":
			FormDivBasicFieldToField(&(_instance.StrokeColor), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(_instance.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(_instance.StrokeWidth), formDiv)
		case "IsExtruded":
			FormDivBasicFieldToField(&(_instance.IsExtruded), formDiv)
		case "BaseAltitude":
			FormDivBasicFieldToField(&(_instance.BaseAltitude), formDiv)
		case "ExtrusionHeight":
			FormDivBasicFieldToField(&(_instance.ExtrusionHeight), formDiv)
		case "Layer:Polygons":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Polygons", func(owner *models.Layer) *[]*models.Polygon { return &owner.Polygons })
		}
	}
}

func __gong__New__PolylineFormCallback(
	_instance *models.Polyline,
	probe *Probe,
	formGroup *form.FormGroup,
) (polylineFormCallback *FormCallback[*models.Polyline]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePolylineFields,
	)
}

type PolylineFormCallback = FormCallback[*models.Polyline]

func savePolylineFields(
	_instance *models.Polyline,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(_instance.Label), formDiv)
		case "PopupHTML":
			FormDivBasicFieldToField(&(_instance.PopupHTML), formDiv)
		case "Coordinates":
			FormDivSliceOfPointersToField(_instance, "Coordinates", &(_instance.Coordinates), formDiv, probe)
		case "StrokeColor":
			FormDivBasicFieldToField(&(_instance.StrokeColor), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(_instance.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(_instance.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(_instance.StrokeDashArray), formDiv)
		case "Layer:Polylines":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Polylines", func(owner *models.Layer) *[]*models.Polyline { return &owner.Polylines })
		}
	}
}

