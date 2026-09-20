// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

type FormCallbackIF interface {
	GetCreationMode() bool
	GetInstance() any
	GetGongstructName() string
	OnSave()
}

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
func __gong__New__AnimateFormCallback(
	_instance *models.Animate,
	probe *Probe,
	formGroup *form.FormGroup,
) (animateFormCallback *FormCallback[*models.Animate]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAnimateFields,
	)
}

type AnimateFormCallback = FormCallback[*models.Animate]

func saveAnimateFields(
	_instance *models.Animate,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "AttributeName":
			FormDivBasicFieldToField(&(_instance.AttributeName), formDiv)
		case "Values":
			FormDivBasicFieldToField(&(_instance.Values), formDiv)
		case "From":
			FormDivBasicFieldToField(&(_instance.From), formDiv)
		case "To":
			FormDivBasicFieldToField(&(_instance.To), formDiv)
		case "Dur":
			FormDivBasicFieldToField(&(_instance.Dur), formDiv)
		case "RepeatCount":
			FormDivBasicFieldToField(&(_instance.RepeatCount), formDiv)
		case "Circle:Animations":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Animations", func(owner *models.Circle) *[]*models.Animate { return &owner.Animations })
		case "Ellipse:Animates":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Animates", func(owner *models.Ellipse) *[]*models.Animate { return &owner.Animates })
		case "Line:Animates":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Animates", func(owner *models.Line) *[]*models.Animate { return &owner.Animates })
		case "LinkAnchoredText:Animates":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Animates", func(owner *models.LinkAnchoredText) *[]*models.Animate { return &owner.Animates })
		case "Path:Animates":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Animates", func(owner *models.Path) *[]*models.Animate { return &owner.Animates })
		case "Polygone:Animates":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Animates", func(owner *models.Polygone) *[]*models.Animate { return &owner.Animates })
		case "Polyline:Animates":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Animates", func(owner *models.Polyline) *[]*models.Animate { return &owner.Animates })
		case "Rect:Animations":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Animations", func(owner *models.Rect) *[]*models.Animate { return &owner.Animations })
		case "RectAnchoredText:Animates":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Animates", func(owner *models.RectAnchoredText) *[]*models.Animate { return &owner.Animates })
		case "Text:Animates":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Animates", func(owner *models.Text) *[]*models.Animate { return &owner.Animates })
		}
	}
}

func __gong__New__CircleFormCallback(
	_instance *models.Circle,
	probe *Probe,
	formGroup *form.FormGroup,
) (circleFormCallback *FormCallback[*models.Circle]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCircleFields,
	)
}

type CircleFormCallback = FormCallback[*models.Circle]

func saveCircleFields(
	_instance *models.Circle,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "CX":
			FormDivBasicFieldToField(&(_instance.CX), formDiv)
		case "CY":
			FormDivBasicFieldToField(&(_instance.CY), formDiv)
		case "Radius":
			FormDivBasicFieldToField(&(_instance.Radius), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(_instance.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(_instance.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(_instance.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(_instance.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(_instance.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(_instance.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(_instance.Transform), formDiv)
		case "Animations":
			FormDivSliceOfPointersToField(_instance, "Animations", &(_instance.Animations), formDiv, probe)
		case "Layer:Circles":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Circles", func(owner *models.Layer) *[]*models.Circle { return &owner.Circles })
		}
	}
}

func __gong__New__ConditionFormCallback(
	_instance *models.Condition,
	probe *Probe,
	formGroup *form.FormGroup,
) (conditionFormCallback *FormCallback[*models.Condition]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveConditionFields,
	)
}

type ConditionFormCallback = FormCallback[*models.Condition]

func saveConditionFields(
	_instance *models.Condition,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Rect:HoveringTrigger":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "HoveringTrigger", func(owner *models.Rect) *[]*models.Condition { return &owner.HoveringTrigger })
		case "Rect:DisplayConditions":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DisplayConditions", func(owner *models.Rect) *[]*models.Condition { return &owner.DisplayConditions })
		}
	}
}

func __gong__New__ControlPointFormCallback(
	_instance *models.ControlPoint,
	probe *Probe,
	formGroup *form.FormGroup,
) (controlpointFormCallback *FormCallback[*models.ControlPoint]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveControlPointFields,
	)
}

type ControlPointFormCallback = FormCallback[*models.ControlPoint]

func saveControlPointFields(
	_instance *models.ControlPoint,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "X_Relative":
			FormDivBasicFieldToField(&(_instance.X_Relative), formDiv)
		case "Y_Relative":
			FormDivBasicFieldToField(&(_instance.Y_Relative), formDiv)
		case "ClosestRect":
			FormDivSelectFieldToField(&(_instance.ClosestRect), probe.stageOfInterest, formDiv)
		case "Link:ControlPoints":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ControlPoints", func(owner *models.Link) *[]*models.ControlPoint { return &owner.ControlPoints })
		}
	}
}

func __gong__New__EllipseFormCallback(
	_instance *models.Ellipse,
	probe *Probe,
	formGroup *form.FormGroup,
) (ellipseFormCallback *FormCallback[*models.Ellipse]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEllipseFields,
	)
}

type EllipseFormCallback = FormCallback[*models.Ellipse]

func saveEllipseFields(
	_instance *models.Ellipse,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "CX":
			FormDivBasicFieldToField(&(_instance.CX), formDiv)
		case "CY":
			FormDivBasicFieldToField(&(_instance.CY), formDiv)
		case "RX":
			FormDivBasicFieldToField(&(_instance.RX), formDiv)
		case "RY":
			FormDivBasicFieldToField(&(_instance.RY), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(_instance.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(_instance.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(_instance.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(_instance.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(_instance.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(_instance.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(_instance.Transform), formDiv)
		case "Animates":
			FormDivSliceOfPointersToField(_instance, "Animates", &(_instance.Animates), formDiv, probe)
		case "Layer:Ellipses":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Ellipses", func(owner *models.Layer) *[]*models.Ellipse { return &owner.Ellipses })
		}
	}
}

func __gong__New__FileToDownloadFormCallback(
	_instance *models.FileToDownload,
	probe *Probe,
	formGroup *form.FormGroup,
) (filetodownloadFormCallback *FormCallback[*models.FileToDownload]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFileToDownloadFields,
	)
}

type FileToDownloadFormCallback = FormCallback[*models.FileToDownload]

func saveFileToDownloadFields(
	_instance *models.FileToDownload,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Base64EncodedContent":
			FormDivBasicFieldToField(&(_instance.Base64EncodedContent), formDiv)
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
		case "Rects":
			FormDivSliceOfPointersToField(_instance, "Rects", &(_instance.Rects), formDiv, probe)
		case "Texts":
			FormDivSliceOfPointersToField(_instance, "Texts", &(_instance.Texts), formDiv, probe)
		case "Circles":
			FormDivSliceOfPointersToField(_instance, "Circles", &(_instance.Circles), formDiv, probe)
		case "Lines":
			FormDivSliceOfPointersToField(_instance, "Lines", &(_instance.Lines), formDiv, probe)
		case "Ellipses":
			FormDivSliceOfPointersToField(_instance, "Ellipses", &(_instance.Ellipses), formDiv, probe)
		case "Polylines":
			FormDivSliceOfPointersToField(_instance, "Polylines", &(_instance.Polylines), formDiv, probe)
		case "Polygones":
			FormDivSliceOfPointersToField(_instance, "Polygones", &(_instance.Polygones), formDiv, probe)
		case "Paths":
			FormDivSliceOfPointersToField(_instance, "Paths", &(_instance.Paths), formDiv, probe)
		case "Links":
			FormDivSliceOfPointersToField(_instance, "Links", &(_instance.Links), formDiv, probe)
		case "RectLinkLinks":
			FormDivSliceOfPointersToField(_instance, "RectLinkLinks", &(_instance.RectLinkLinks), formDiv, probe)
		case "SVG:Layers":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Layers", func(owner *models.SVG) *[]*models.Layer { return &owner.Layers })
		}
	}
}

func __gong__New__LineFormCallback(
	_instance *models.Line,
	probe *Probe,
	formGroup *form.FormGroup,
) (lineFormCallback *FormCallback[*models.Line]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLineFields,
	)
}

type LineFormCallback = FormCallback[*models.Line]

func saveLineFields(
	_instance *models.Line,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "X1":
			FormDivBasicFieldToField(&(_instance.X1), formDiv)
		case "Y1":
			FormDivBasicFieldToField(&(_instance.Y1), formDiv)
		case "X2":
			FormDivBasicFieldToField(&(_instance.X2), formDiv)
		case "Y2":
			FormDivBasicFieldToField(&(_instance.Y2), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(_instance.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(_instance.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(_instance.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(_instance.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(_instance.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(_instance.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(_instance.Transform), formDiv)
		case "Animates":
			FormDivSliceOfPointersToField(_instance, "Animates", &(_instance.Animates), formDiv, probe)
		case "MouseClickX":
			FormDivBasicFieldToField(&(_instance.MouseClickX), formDiv)
		case "MouseClickY":
			FormDivBasicFieldToField(&(_instance.MouseClickY), formDiv)
		case "Layer:Lines":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Lines", func(owner *models.Layer) *[]*models.Line { return &owner.Lines })
		}
	}
}

func __gong__New__LinkFormCallback(
	_instance *models.Link,
	probe *Probe,
	formGroup *form.FormGroup,
) (linkFormCallback *FormCallback[*models.Link]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLinkFields,
	)
}

type LinkFormCallback = FormCallback[*models.Link]

func saveLinkFields(
	_instance *models.Link,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(_instance.Type), formDiv)
		case "IsBezierCurve":
			FormDivBasicFieldToField(&(_instance.IsBezierCurve), formDiv)
		case "Start":
			FormDivSelectFieldToField(&(_instance.Start), probe.stageOfInterest, formDiv)
		case "StartAnchorType":
			FormDivEnumStringFieldToField(&(_instance.StartAnchorType), formDiv)
		case "End":
			FormDivSelectFieldToField(&(_instance.End), probe.stageOfInterest, formDiv)
		case "EndAnchorType":
			FormDivEnumStringFieldToField(&(_instance.EndAnchorType), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(_instance.StartOrientation), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(_instance.StartRatio), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(_instance.EndOrientation), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(_instance.EndRatio), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(_instance.CornerOffsetRatio), formDiv)
		case "CornerRadius":
			FormDivBasicFieldToField(&(_instance.CornerRadius), formDiv)
		case "HasEndArrow":
			FormDivBasicFieldToField(&(_instance.HasEndArrow), formDiv)
		case "EndArrowSize":
			FormDivBasicFieldToField(&(_instance.EndArrowSize), formDiv)
		case "EndArrowOffset":
			FormDivBasicFieldToField(&(_instance.EndArrowOffset), formDiv)
		case "HasStartArrow":
			FormDivBasicFieldToField(&(_instance.HasStartArrow), formDiv)
		case "StartArrowSize":
			FormDivBasicFieldToField(&(_instance.StartArrowSize), formDiv)
		case "StartArrowOffset":
			FormDivBasicFieldToField(&(_instance.StartArrowOffset), formDiv)
		case "TextAtArrowStart":
			FormDivSliceOfPointersToField(_instance, "TextAtArrowStart", &(_instance.TextAtArrowStart), formDiv, probe)
		case "TextAtArrowEnd":
			FormDivSliceOfPointersToField(_instance, "TextAtArrowEnd", &(_instance.TextAtArrowEnd), formDiv, probe)
		case "TextAtCorner":
			FormDivSliceOfPointersToField(_instance, "TextAtCorner", &(_instance.TextAtCorner), formDiv, probe)
		case "PathAtArrowStart":
			FormDivSliceOfPointersToField(_instance, "PathAtArrowStart", &(_instance.PathAtArrowStart), formDiv, probe)
		case "PathAtArrowEnd":
			FormDivSliceOfPointersToField(_instance, "PathAtArrowEnd", &(_instance.PathAtArrowEnd), formDiv, probe)
		case "PathAtCorner":
			FormDivSliceOfPointersToField(_instance, "PathAtCorner", &(_instance.PathAtCorner), formDiv, probe)
		case "ControlPoints":
			FormDivSliceOfPointersToField(_instance, "ControlPoints", &(_instance.ControlPoints), formDiv, probe)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(_instance.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(_instance.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(_instance.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(_instance.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(_instance.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(_instance.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(_instance.Transform), formDiv)
		case "MouseX":
			FormDivBasicFieldToField(&(_instance.MouseX), formDiv)
		case "MouseY":
			FormDivBasicFieldToField(&(_instance.MouseY), formDiv)
		case "MouseEventKey":
			FormDivEnumStringFieldToField(&(_instance.MouseEventKey), formDiv)
		case "Layer:Links":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Links", func(owner *models.Layer) *[]*models.Link { return &owner.Links })
		}
	}
}

func __gong__New__LinkAnchoredPathFormCallback(
	_instance *models.LinkAnchoredPath,
	probe *Probe,
	formGroup *form.FormGroup,
) (linkanchoredpathFormCallback *FormCallback[*models.LinkAnchoredPath]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLinkAnchoredPathFields,
	)
}

type LinkAnchoredPathFormCallback = FormCallback[*models.LinkAnchoredPath]

func saveLinkAnchoredPathFields(
	_instance *models.LinkAnchoredPath,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Definition":
			FormDivBasicFieldToField(&(_instance.Definition), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(_instance.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(_instance.Y_Offset), formDiv)
		case "ScalePropotionnally":
			FormDivBasicFieldToField(&(_instance.ScalePropotionnally), formDiv)
		case "AppliedScaling":
			FormDivBasicFieldToField(&(_instance.AppliedScaling), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(_instance.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(_instance.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(_instance.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(_instance.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(_instance.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(_instance.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(_instance.Transform), formDiv)
		case "Link:PathAtArrowStart":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "PathAtArrowStart", func(owner *models.Link) *[]*models.LinkAnchoredPath { return &owner.PathAtArrowStart })
		case "Link:PathAtArrowEnd":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "PathAtArrowEnd", func(owner *models.Link) *[]*models.LinkAnchoredPath { return &owner.PathAtArrowEnd })
		case "Link:PathAtCorner":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "PathAtCorner", func(owner *models.Link) *[]*models.LinkAnchoredPath { return &owner.PathAtCorner })
		}
	}
}

func __gong__New__LinkAnchoredTextFormCallback(
	_instance *models.LinkAnchoredText,
	probe *Probe,
	formGroup *form.FormGroup,
) (linkanchoredtextFormCallback *FormCallback[*models.LinkAnchoredText]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLinkAnchoredTextFields,
	)
}

type LinkAnchoredTextFormCallback = FormCallback[*models.LinkAnchoredText]

func saveLinkAnchoredTextFields(
	_instance *models.LinkAnchoredText,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(_instance.Content), formDiv)
		case "AutomaticLayout":
			FormDivBasicFieldToField(&(_instance.AutomaticLayout), formDiv)
		case "LinkAnchorType":
			FormDivEnumStringFieldToField(&(_instance.LinkAnchorType), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(_instance.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(_instance.Y_Offset), formDiv)
		case "FontWeight":
			FormDivBasicFieldToField(&(_instance.FontWeight), formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&(_instance.FontSize), formDiv)
		case "FontStyle":
			FormDivBasicFieldToField(&(_instance.FontStyle), formDiv)
		case "LetterSpacing":
			FormDivBasicFieldToField(&(_instance.LetterSpacing), formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&(_instance.FontFamily), formDiv)
		case "WhiteSpace":
			FormDivEnumStringFieldToField(&(_instance.WhiteSpace), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(_instance.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(_instance.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(_instance.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(_instance.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(_instance.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(_instance.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(_instance.Transform), formDiv)
		case "Animates":
			FormDivSliceOfPointersToField(_instance, "Animates", &(_instance.Animates), formDiv, probe)
		case "Link:TextAtArrowStart":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TextAtArrowStart", func(owner *models.Link) *[]*models.LinkAnchoredText { return &owner.TextAtArrowStart })
		case "Link:TextAtArrowEnd":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TextAtArrowEnd", func(owner *models.Link) *[]*models.LinkAnchoredText { return &owner.TextAtArrowEnd })
		case "Link:TextAtCorner":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TextAtCorner", func(owner *models.Link) *[]*models.LinkAnchoredText { return &owner.TextAtCorner })
		}
	}
}

func __gong__New__PathFormCallback(
	_instance *models.Path,
	probe *Probe,
	formGroup *form.FormGroup,
) (pathFormCallback *FormCallback[*models.Path]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePathFields,
	)
}

type PathFormCallback = FormCallback[*models.Path]

func savePathFields(
	_instance *models.Path,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Definition":
			FormDivBasicFieldToField(&(_instance.Definition), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(_instance.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(_instance.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(_instance.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(_instance.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(_instance.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(_instance.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(_instance.Transform), formDiv)
		case "Animates":
			FormDivSliceOfPointersToField(_instance, "Animates", &(_instance.Animates), formDiv, probe)
		case "Layer:Paths":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Paths", func(owner *models.Layer) *[]*models.Path { return &owner.Paths })
		}
	}
}

func __gong__New__PointFormCallback(
	_instance *models.Point,
	probe *Probe,
	formGroup *form.FormGroup,
) (pointFormCallback *FormCallback[*models.Point]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePointFields,
	)
}

type PointFormCallback = FormCallback[*models.Point]

func savePointFields(
	_instance *models.Point,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(_instance.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(_instance.Y), formDiv)
		}
	}
}

func __gong__New__PolygoneFormCallback(
	_instance *models.Polygone,
	probe *Probe,
	formGroup *form.FormGroup,
) (polygoneFormCallback *FormCallback[*models.Polygone]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePolygoneFields,
	)
}

type PolygoneFormCallback = FormCallback[*models.Polygone]

func savePolygoneFields(
	_instance *models.Polygone,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Points":
			FormDivBasicFieldToField(&(_instance.Points), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(_instance.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(_instance.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(_instance.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(_instance.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(_instance.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(_instance.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(_instance.Transform), formDiv)
		case "Animates":
			FormDivSliceOfPointersToField(_instance, "Animates", &(_instance.Animates), formDiv, probe)
		case "Layer:Polygones":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Polygones", func(owner *models.Layer) *[]*models.Polygone { return &owner.Polygones })
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
		case "Points":
			FormDivBasicFieldToField(&(_instance.Points), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(_instance.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(_instance.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(_instance.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(_instance.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(_instance.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(_instance.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(_instance.Transform), formDiv)
		case "Animates":
			FormDivSliceOfPointersToField(_instance, "Animates", &(_instance.Animates), formDiv, probe)
		case "Layer:Polylines":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Polylines", func(owner *models.Layer) *[]*models.Polyline { return &owner.Polylines })
		}
	}
}

func __gong__New__RectFormCallback(
	_instance *models.Rect,
	probe *Probe,
	formGroup *form.FormGroup,
) (rectFormCallback *FormCallback[*models.Rect]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRectFields,
	)
}

type RectFormCallback = FormCallback[*models.Rect]

func saveRectFields(
	_instance *models.Rect,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(_instance.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(_instance.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "RX":
			FormDivBasicFieldToField(&(_instance.RX), formDiv)
		case "Peers":
			FormDivSliceOfPointersToField(_instance, "Peers", &(_instance.Peers), formDiv, probe)
		case "EnclosingRect":
			FormDivSelectFieldToField(&(_instance.EnclosingRect), probe.stageOfInterest, formDiv)
		case "Obstacles":
			FormDivSliceOfPointersToField(_instance, "Obstacles", &(_instance.Obstacles), formDiv, probe)
		case "AnchoredTo":
			FormDivSelectFieldToField(&(_instance.AnchoredTo), probe.stageOfInterest, formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(_instance.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(_instance.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(_instance.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(_instance.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(_instance.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(_instance.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(_instance.Transform), formDiv)
		case "HoveringTrigger":
			FormDivSliceOfPointersToField(_instance, "HoveringTrigger", &(_instance.HoveringTrigger), formDiv, probe)
		case "DisplayConditions":
			FormDivSliceOfPointersToField(_instance, "DisplayConditions", &(_instance.DisplayConditions), formDiv, probe)
		case "Animations":
			FormDivSliceOfPointersToField(_instance, "Animations", &(_instance.Animations), formDiv, probe)
		case "IsSelectable":
			FormDivBasicFieldToField(&(_instance.IsSelectable), formDiv)
		case "IsSelected":
			FormDivBasicFieldToField(&(_instance.IsSelected), formDiv)
		case "CanHaveLeftHandle":
			FormDivBasicFieldToField(&(_instance.CanHaveLeftHandle), formDiv)
		case "HasLeftHandle":
			FormDivBasicFieldToField(&(_instance.HasLeftHandle), formDiv)
		case "CanHaveRightHandle":
			FormDivBasicFieldToField(&(_instance.CanHaveRightHandle), formDiv)
		case "HasRightHandle":
			FormDivBasicFieldToField(&(_instance.HasRightHandle), formDiv)
		case "CanHaveTopHandle":
			FormDivBasicFieldToField(&(_instance.CanHaveTopHandle), formDiv)
		case "HasTopHandle":
			FormDivBasicFieldToField(&(_instance.HasTopHandle), formDiv)
		case "IsScalingProportionally":
			FormDivBasicFieldToField(&(_instance.IsScalingProportionally), formDiv)
		case "CanHaveBottomHandle":
			FormDivBasicFieldToField(&(_instance.CanHaveBottomHandle), formDiv)
		case "HasBottomHandle":
			FormDivBasicFieldToField(&(_instance.HasBottomHandle), formDiv)
		case "CanMoveHorizontaly":
			FormDivBasicFieldToField(&(_instance.CanMoveHorizontaly), formDiv)
		case "CanMoveVerticaly":
			FormDivBasicFieldToField(&(_instance.CanMoveVerticaly), formDiv)
		case "RectAnchoredTexts":
			FormDivSliceOfPointersToField(_instance, "RectAnchoredTexts", &(_instance.RectAnchoredTexts), formDiv, probe)
		case "RectAnchoredRects":
			FormDivSliceOfPointersToField(_instance, "RectAnchoredRects", &(_instance.RectAnchoredRects), formDiv, probe)
		case "RectAnchoredPaths":
			FormDivSliceOfPointersToField(_instance, "RectAnchoredPaths", &(_instance.RectAnchoredPaths), formDiv, probe)
		case "RectAnchoredPngImages":
			FormDivSliceOfPointersToField(_instance, "RectAnchoredPngImages", &(_instance.RectAnchoredPngImages), formDiv, probe)
		case "ChangeColorWhenHovered":
			FormDivBasicFieldToField(&(_instance.ChangeColorWhenHovered), formDiv)
		case "ColorWhenHovered":
			FormDivBasicFieldToField(&(_instance.ColorWhenHovered), formDiv)
		case "OriginalColor":
			FormDivBasicFieldToField(&(_instance.OriginalColor), formDiv)
		case "FillOpacityWhenHovered":
			FormDivBasicFieldToField(&(_instance.FillOpacityWhenHovered), formDiv)
		case "OriginalFillOpacity":
			FormDivBasicFieldToField(&(_instance.OriginalFillOpacity), formDiv)
		case "HasToolTip":
			FormDivBasicFieldToField(&(_instance.HasToolTip), formDiv)
		case "ToolTipText":
			FormDivBasicFieldToField(&(_instance.ToolTipText), formDiv)
		case "ToolTipPosition":
			FormDivEnumStringFieldToField(&(_instance.ToolTipPosition), formDiv)
		case "MouseX":
			FormDivBasicFieldToField(&(_instance.MouseX), formDiv)
		case "MouseY":
			FormDivBasicFieldToField(&(_instance.MouseY), formDiv)
		case "MouseEventKey":
			FormDivEnumStringFieldToField(&(_instance.MouseEventKey), formDiv)
		case "URLPath":
			FormDivBasicFieldToField(&(_instance.URLPath), formDiv)
		case "URLTarget":
			FormDivEnumStringFieldToField(&(_instance.URLTarget), formDiv)
		case "Layer:Rects":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Rects", func(owner *models.Layer) *[]*models.Rect { return &owner.Rects })
		case "Rect:Peers":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Peers", func(owner *models.Rect) *[]*models.Rect { return &owner.Peers })
		case "Rect:Obstacles":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Obstacles", func(owner *models.Rect) *[]*models.Rect { return &owner.Obstacles })
		}
	}
}

func __gong__New__RectAnchoredPathFormCallback(
	_instance *models.RectAnchoredPath,
	probe *Probe,
	formGroup *form.FormGroup,
) (rectanchoredpathFormCallback *FormCallback[*models.RectAnchoredPath]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRectAnchoredPathFields,
	)
}

type RectAnchoredPathFormCallback = FormCallback[*models.RectAnchoredPath]

func saveRectAnchoredPathFields(
	_instance *models.RectAnchoredPath,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Definition":
			FormDivBasicFieldToField(&(_instance.Definition), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(_instance.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(_instance.Y_Offset), formDiv)
		case "RectAnchorType":
			FormDivEnumStringFieldToField(&(_instance.RectAnchorType), formDiv)
		case "ScalePropotionnally":
			FormDivBasicFieldToField(&(_instance.ScalePropotionnally), formDiv)
		case "AppliedScaling":
			FormDivBasicFieldToField(&(_instance.AppliedScaling), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(_instance.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(_instance.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(_instance.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(_instance.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(_instance.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(_instance.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(_instance.Transform), formDiv)
		case "Rect:RectAnchoredPaths":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RectAnchoredPaths", func(owner *models.Rect) *[]*models.RectAnchoredPath { return &owner.RectAnchoredPaths })
		}
	}
}

func __gong__New__RectAnchoredPngImageFormCallback(
	_instance *models.RectAnchoredPngImage,
	probe *Probe,
	formGroup *form.FormGroup,
) (rectanchoredpngimageFormCallback *FormCallback[*models.RectAnchoredPngImage]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRectAnchoredPngImageFields,
	)
}

type RectAnchoredPngImageFormCallback = FormCallback[*models.RectAnchoredPngImage]

func saveRectAnchoredPngImageFields(
	_instance *models.RectAnchoredPngImage,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(_instance.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(_instance.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "RX":
			FormDivBasicFieldToField(&(_instance.RX), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(_instance.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(_instance.Y_Offset), formDiv)
		case "RectAnchorType":
			FormDivEnumStringFieldToField(&(_instance.RectAnchorType), formDiv)
		case "Base64Content":
			FormDivBasicFieldToField(&(_instance.Base64Content), formDiv)
		case "Rect:RectAnchoredPngImages":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RectAnchoredPngImages", func(owner *models.Rect) *[]*models.RectAnchoredPngImage { return &owner.RectAnchoredPngImages })
		}
	}
}

func __gong__New__RectAnchoredRectFormCallback(
	_instance *models.RectAnchoredRect,
	probe *Probe,
	formGroup *form.FormGroup,
) (rectanchoredrectFormCallback *FormCallback[*models.RectAnchoredRect]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRectAnchoredRectFields,
	)
}

type RectAnchoredRectFormCallback = FormCallback[*models.RectAnchoredRect]

func saveRectAnchoredRectFields(
	_instance *models.RectAnchoredRect,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(_instance.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(_instance.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "RX":
			FormDivBasicFieldToField(&(_instance.RX), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(_instance.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(_instance.Y_Offset), formDiv)
		case "RectAnchorType":
			FormDivEnumStringFieldToField(&(_instance.RectAnchorType), formDiv)
		case "WidthFollowRect":
			FormDivBasicFieldToField(&(_instance.WidthFollowRect), formDiv)
		case "HeightFollowRect":
			FormDivBasicFieldToField(&(_instance.HeightFollowRect), formDiv)
		case "HasToolTip":
			FormDivBasicFieldToField(&(_instance.HasToolTip), formDiv)
		case "ToolTipText":
			FormDivBasicFieldToField(&(_instance.ToolTipText), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(_instance.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(_instance.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(_instance.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(_instance.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(_instance.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(_instance.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(_instance.Transform), formDiv)
		case "Rect:RectAnchoredRects":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RectAnchoredRects", func(owner *models.Rect) *[]*models.RectAnchoredRect { return &owner.RectAnchoredRects })
		}
	}
}

func __gong__New__RectAnchoredTextFormCallback(
	_instance *models.RectAnchoredText,
	probe *Probe,
	formGroup *form.FormGroup,
) (rectanchoredtextFormCallback *FormCallback[*models.RectAnchoredText]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRectAnchoredTextFields,
	)
}

type RectAnchoredTextFormCallback = FormCallback[*models.RectAnchoredText]

func saveRectAnchoredTextFields(
	_instance *models.RectAnchoredText,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(_instance.Content), formDiv)
		case "FontWeight":
			FormDivBasicFieldToField(&(_instance.FontWeight), formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&(_instance.FontSize), formDiv)
		case "FontStyle":
			FormDivBasicFieldToField(&(_instance.FontStyle), formDiv)
		case "LetterSpacing":
			FormDivBasicFieldToField(&(_instance.LetterSpacing), formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&(_instance.FontFamily), formDiv)
		case "WhiteSpace":
			FormDivEnumStringFieldToField(&(_instance.WhiteSpace), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(_instance.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(_instance.Y_Offset), formDiv)
		case "RectAnchorType":
			FormDivEnumStringFieldToField(&(_instance.RectAnchorType), formDiv)
		case "TextAnchorType":
			FormDivEnumStringFieldToField(&(_instance.TextAnchorType), formDiv)
		case "DominantBaseline":
			FormDivEnumStringFieldToField(&(_instance.DominantBaseline), formDiv)
		case "WritingMode":
			FormDivEnumStringFieldToField(&(_instance.WritingMode), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(_instance.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(_instance.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(_instance.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(_instance.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(_instance.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(_instance.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(_instance.Transform), formDiv)
		case "Animates":
			FormDivSliceOfPointersToField(_instance, "Animates", &(_instance.Animates), formDiv, probe)
		case "URLPath":
			FormDivBasicFieldToField(&(_instance.URLPath), formDiv)
		case "URLTarget":
			FormDivEnumStringFieldToField(&(_instance.URLTarget), formDiv)
		case "Rect:RectAnchoredTexts":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RectAnchoredTexts", func(owner *models.Rect) *[]*models.RectAnchoredText { return &owner.RectAnchoredTexts })
		}
	}
}

func __gong__New__RectLinkLinkFormCallback(
	_instance *models.RectLinkLink,
	probe *Probe,
	formGroup *form.FormGroup,
) (rectlinklinkFormCallback *FormCallback[*models.RectLinkLink]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRectLinkLinkFields,
	)
}

type RectLinkLinkFormCallback = FormCallback[*models.RectLinkLink]

func saveRectLinkLinkFields(
	_instance *models.RectLinkLink,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Start":
			FormDivSelectFieldToField(&(_instance.Start), probe.stageOfInterest, formDiv)
		case "End":
			FormDivSelectFieldToField(&(_instance.End), probe.stageOfInterest, formDiv)
		case "TargetAnchorPosition":
			FormDivBasicFieldToField(&(_instance.TargetAnchorPosition), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(_instance.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(_instance.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(_instance.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(_instance.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(_instance.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(_instance.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(_instance.Transform), formDiv)
		case "Layer:RectLinkLinks":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RectLinkLinks", func(owner *models.Layer) *[]*models.RectLinkLink { return &owner.RectLinkLinks })
		}
	}
}

func __gong__New__SVGFormCallback(
	_instance *models.SVG,
	probe *Probe,
	formGroup *form.FormGroup,
) (svgFormCallback *FormCallback[*models.SVG]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSVGFields,
	)
}

type SVGFormCallback = FormCallback[*models.SVG]

func saveSVGFields(
	_instance *models.SVG,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Layers":
			FormDivSliceOfPointersToField(_instance, "Layers", &(_instance.Layers), formDiv, probe)
		case "DrawingState":
			FormDivEnumStringFieldToField(&(_instance.DrawingState), formDiv)
		case "StartRect":
			FormDivSelectFieldToField(&(_instance.StartRect), probe.stageOfInterest, formDiv)
		case "EndRect":
			FormDivSelectFieldToField(&(_instance.EndRect), probe.stageOfInterest, formDiv)
		case "IsEditable":
			FormDivBasicFieldToField(&(_instance.IsEditable), formDiv)
		case "IsSVGFrontEndFileGenerated":
			FormDivBasicFieldToField(&(_instance.IsSVGFrontEndFileGenerated), formDiv)
		case "IsSVGBackEndFileGenerated":
			FormDivBasicFieldToField(&(_instance.IsSVGBackEndFileGenerated), formDiv)
		case "DefaultDirectoryForGeneratedImages":
			FormDivBasicFieldToField(&(_instance.DefaultDirectoryForGeneratedImages), formDiv)
		case "IsControlBannerHidden":
			FormDivBasicFieldToField(&(_instance.IsControlBannerHidden), formDiv)
		case "PanX":
			FormDivBasicFieldToField(&(_instance.PanX), formDiv)
		case "PanY":
			FormDivBasicFieldToField(&(_instance.PanY), formDiv)
		case "Zoom":
			FormDivBasicFieldToField(&(_instance.Zoom), formDiv)
		case "OverrideWidth":
			FormDivBasicFieldToField(&(_instance.OverrideWidth), formDiv)
		case "OverriddenWidth":
			FormDivBasicFieldToField(&(_instance.OverriddenWidth), formDiv)
		case "OverrideHeight":
			FormDivBasicFieldToField(&(_instance.OverrideHeight), formDiv)
		case "OverriddenHeight":
			FormDivBasicFieldToField(&(_instance.OverriddenHeight), formDiv)
		}
	}
}

func __gong__New__SvgTextFormCallback(
	_instance *models.SvgText,
	probe *Probe,
	formGroup *form.FormGroup,
) (svgtextFormCallback *FormCallback[*models.SvgText]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSvgTextFields,
	)
}

type SvgTextFormCallback = FormCallback[*models.SvgText]

func saveSvgTextFields(
	_instance *models.SvgText,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Text":
			FormDivBasicFieldToField(&(_instance.Text), formDiv)
		}
	}
}

func __gong__New__TextFormCallback(
	_instance *models.Text,
	probe *Probe,
	formGroup *form.FormGroup,
) (textFormCallback *FormCallback[*models.Text]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTextFields,
	)
}

type TextFormCallback = FormCallback[*models.Text]

func saveTextFields(
	_instance *models.Text,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(_instance.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(_instance.Y), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(_instance.Content), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(_instance.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(_instance.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(_instance.Stroke), formDiv)
		case "StrokeOpacity":
			FormDivBasicFieldToField(&(_instance.StrokeOpacity), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(_instance.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(_instance.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(_instance.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(_instance.Transform), formDiv)
		case "FontWeight":
			FormDivBasicFieldToField(&(_instance.FontWeight), formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&(_instance.FontSize), formDiv)
		case "FontStyle":
			FormDivBasicFieldToField(&(_instance.FontStyle), formDiv)
		case "LetterSpacing":
			FormDivBasicFieldToField(&(_instance.LetterSpacing), formDiv)
		case "FontFamily":
			FormDivBasicFieldToField(&(_instance.FontFamily), formDiv)
		case "WhiteSpace":
			FormDivEnumStringFieldToField(&(_instance.WhiteSpace), formDiv)
		case "Animates":
			FormDivSliceOfPointersToField(_instance, "Animates", &(_instance.Animates), formDiv, probe)
		case "Layer:Texts":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Texts", func(owner *models.Layer) *[]*models.Text { return &owner.Texts })
		}
	}
}

