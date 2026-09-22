// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
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
func __gong__New__Angle0ShapeFormCallback(
	_instance *models.Angle0Shape,
	probe *Probe,
	formGroup *form.FormGroup,
) (angle0shapeFormCallback *FormCallback[*models.Angle0Shape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAngle0ShapeFields,
	)
}

type Angle0ShapeFormCallback = FormCallback[*models.Angle0Shape]

func saveAngle0ShapeFields(
	_instance *models.Angle0Shape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__ArcNormalVectorShapeFormCallback(
	_instance *models.ArcNormalVectorShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (arcnormalvectorshapeFormCallback *FormCallback[*models.ArcNormalVectorShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveArcNormalVectorShapeFields,
	)
}

type ArcNormalVectorShapeFormCallback = FormCallback[*models.ArcNormalVectorShape]

func saveArcNormalVectorShapeFields(
	_instance *models.ArcNormalVectorShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "ArcNormalVectorShapeGrid:ArcNormalVectorShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ArcNormalVectorShapes", func(owner *models.ArcNormalVectorShapeGrid) *[]*models.ArcNormalVectorShape { return &owner.ArcNormalVectorShapes })
		}
	}
}

func __gong__New__ArcNormalVectorShapeGridFormCallback(
	_instance *models.ArcNormalVectorShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (arcnormalvectorshapegridFormCallback *FormCallback[*models.ArcNormalVectorShapeGrid]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveArcNormalVectorShapeGridFields,
	)
}

type ArcNormalVectorShapeGridFormCallback = FormCallback[*models.ArcNormalVectorShapeGrid]

func saveArcNormalVectorShapeGridFields(
	_instance *models.ArcNormalVectorShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ArcNormalVectorShapes":
			FormDivSliceOfPointersToField(_instance, "ArcNormalVectorShapes", &(_instance.ArcNormalVectorShapes), formDiv, probe)
		}
	}
}

func __gong__New__AxesShapeFormCallback(
	_instance *models.AxesShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (axesshapeFormCallback *FormCallback[*models.AxesShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveAxesShapeFields,
	)
}

type AxesShapeFormCallback = FormCallback[*models.AxesShape]

func saveAxesShapeFields(
	_instance *models.AxesShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "LengthX":
			FormDivBasicFieldToField(&(_instance.LengthX), formDiv)
		case "LengthY":
			FormDivBasicFieldToField(&(_instance.LengthY), formDiv)
		case "IsWithHiddenHandle":
			FormDivBasicFieldToField(&(_instance.IsWithHiddenHandle), formDiv)
		}
	}
}

func __gong__New__BaseVectorShapeFormCallback(
	_instance *models.BaseVectorShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (basevectorshapeFormCallback *FormCallback[*models.BaseVectorShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBaseVectorShapeFields,
	)
}

type BaseVectorShapeFormCallback = FormCallback[*models.BaseVectorShape]

func saveBaseVectorShapeFields(
	_instance *models.BaseVectorShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "BaseVectorShapeGrid:BaseVectorShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "BaseVectorShapes", func(owner *models.BaseVectorShapeGrid) *[]*models.BaseVectorShape { return &owner.BaseVectorShapes })
		}
	}
}

func __gong__New__BaseVectorShapeGridFormCallback(
	_instance *models.BaseVectorShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (basevectorshapegridFormCallback *FormCallback[*models.BaseVectorShapeGrid]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBaseVectorShapeGridFields,
	)
}

type BaseVectorShapeGridFormCallback = FormCallback[*models.BaseVectorShapeGrid]

func saveBaseVectorShapeGridFields(
	_instance *models.BaseVectorShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "BaseVectorShapes":
			FormDivSliceOfPointersToField(_instance, "BaseVectorShapes", &(_instance.BaseVectorShapes), formDiv, probe)
		}
	}
}

func __gong__New__ChosenP1P2PairShapeFormCallback(
	_instance *models.ChosenP1P2PairShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (chosenp1p2pairshapeFormCallback *FormCallback[*models.ChosenP1P2PairShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveChosenP1P2PairShapeFields,
	)
}

type ChosenP1P2PairShapeFormCallback = FormCallback[*models.ChosenP1P2PairShape]

func saveChosenP1P2PairShapeFields(
	_instance *models.ChosenP1P2PairShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "P1X":
			FormDivBasicFieldToField(&(_instance.P1X), formDiv)
		case "P1Y":
			FormDivBasicFieldToField(&(_instance.P1Y), formDiv)
		case "P2X":
			FormDivBasicFieldToField(&(_instance.P2X), formDiv)
		case "P2Y":
			FormDivBasicFieldToField(&(_instance.P2Y), formDiv)
		case "PxX":
			FormDivBasicFieldToField(&(_instance.PxX), formDiv)
		case "PxY":
			FormDivBasicFieldToField(&(_instance.PxY), formDiv)
		case "DistanceP1Px":
			FormDivBasicFieldToField(&(_instance.DistanceP1Px), formDiv)
		case "DistanceP2Px":
			FormDivBasicFieldToField(&(_instance.DistanceP2Px), formDiv)
		case "DistanceSum":
			FormDivBasicFieldToField(&(_instance.DistanceSum), formDiv)
		}
	}
}

func __gong__New__CircleGridShapeFormCallback(
	_instance *models.CircleGridShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (circlegridshapeFormCallback *FormCallback[*models.CircleGridShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCircleGridShapeFields,
	)
}

type CircleGridShapeFormCallback = FormCallback[*models.CircleGridShape]

func saveCircleGridShapeFields(
	_instance *models.CircleGridShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__Circumference3DShapeFormCallback(
	_instance *models.Circumference3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (circumference3dshapeFormCallback *FormCallback[*models.Circumference3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCircumference3DShapeFields,
	)
}

type Circumference3DShapeFormCallback = FormCallback[*models.Circumference3DShape]

func saveCircumference3DShapeFields(
	_instance *models.Circumference3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__Clock2DDiagramFormCallback(
	_instance *models.Clock2DDiagram,
	probe *Probe,
	formGroup *form.FormGroup,
) (clock2ddiagramFormCallback *FormCallback[*models.Clock2DDiagram]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveClock2DDiagramFields,
	)
}

type Clock2DDiagramFormCallback = FormCallback[*models.Clock2DDiagram]

func saveClock2DDiagramFields(
	_instance *models.Clock2DDiagram,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Zoom":
			FormDivBasicFieldToField(&(_instance.Zoom), formDiv)
		case "IsHiddenAxesShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenAxesShape), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(_instance.IsChecked), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "PlantAbstract:Clock2DDiagrams":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Clock2DDiagrams", func(owner *models.PlantAbstract) *[]*models.Clock2DDiagram { return &owner.Clock2DDiagrams })
		}
	}
}

func __gong__New__Clock3DDiagramFormCallback(
	_instance *models.Clock3DDiagram,
	probe *Probe,
	formGroup *form.FormGroup,
) (clock3ddiagramFormCallback *FormCallback[*models.Clock3DDiagram]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveClock3DDiagramFields,
	)
}

type Clock3DDiagramFormCallback = FormCallback[*models.Clock3DDiagram]

func saveClock3DDiagramFields(
	_instance *models.Clock3DDiagram,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "IsHiddenClockTopCurveShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenClockTopCurveShape), formDiv)
		case "ClockTopCurveShape":
			FormDivSelectFieldToField(&(_instance.ClockTopCurveShape), probe.stageOfInterest, formDiv)
		case "IsHiddenTorus3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenTorus3DShape), formDiv)
		case "Torus3DShape":
			FormDivSelectFieldToField(&(_instance.Torus3DShape), probe.stageOfInterest, formDiv)
		case "IsHiddenSampledPoints3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenSampledPoints3DShape), formDiv)
		case "SampledPoints3DShape":
			FormDivSelectFieldToField(&(_instance.SampledPoints3DShape), probe.stageOfInterest, formDiv)
		case "IsHiddenTiledFloor3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenTiledFloor3DShape), formDiv)
		case "TiledFloor3DShape":
			FormDivSelectFieldToField(&(_instance.TiledFloor3DShape), probe.stageOfInterest, formDiv)
		case "Rendered3DShape":
			FormDivSelectFieldToField(&(_instance.Rendered3DShape), probe.stageOfInterest, formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(_instance.IsChecked), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "PlantAbstract:Clock3DDiagrams":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Clock3DDiagrams", func(owner *models.PlantAbstract) *[]*models.Clock3DDiagram { return &owner.Clock3DDiagrams })
		}
	}
}

func __gong__New__ClockAbstractFormCallback(
	_instance *models.ClockAbstract,
	probe *Probe,
	formGroup *form.FormGroup,
) (clockabstractFormCallback *FormCallback[*models.ClockAbstract]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveClockAbstractFields,
	)
}

type ClockAbstractFormCallback = FormCallback[*models.ClockAbstract]

func saveClockAbstractFields(
	_instance *models.ClockAbstract,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "RadialRepetitions":
			FormDivBasicFieldToField(&(_instance.RadialRepetitions), formDiv)
		case "Transparency":
			FormDivBasicFieldToField(&(_instance.Transparency), formDiv)
		case "RelativeTubeDiameter":
			FormDivBasicFieldToField(&(_instance.RelativeTubeDiameter), formDiv)
		case "RelativeHeight3DTorus":
			FormDivBasicFieldToField(&(_instance.RelativeHeight3DTorus), formDiv)
		case "ClockTorusVerticalScale":
			FormDivBasicFieldToField(&(_instance.ClockTorusVerticalScale), formDiv)
		case "RelativeHeight":
			FormDivBasicFieldToField(&(_instance.RelativeHeight), formDiv)
		case "ProjectionAngle":
			FormDivBasicFieldToField(&(_instance.ProjectionAngle), formDiv)
		}
	}
}

func __gong__New__ClockTopCurveShapeFormCallback(
	_instance *models.ClockTopCurveShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (clocktopcurveshapeFormCallback *FormCallback[*models.ClockTopCurveShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveClockTopCurveShapeFields,
	)
}

type ClockTopCurveShapeFormCallback = FormCallback[*models.ClockTopCurveShape]

func saveClockTopCurveShapeFields(
	_instance *models.ClockTopCurveShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__CutLine3DShapeFormCallback(
	_instance *models.CutLine3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (cutline3dshapeFormCallback *FormCallback[*models.CutLine3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCutLine3DShapeFields,
	)
}

type CutLine3DShapeFormCallback = FormCallback[*models.CutLine3DShape]

func saveCutLine3DShapeFields(
	_instance *models.CutLine3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__EndArcShapeFormCallback(
	_instance *models.EndArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (endarcshapeFormCallback *FormCallback[*models.EndArcShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEndArcShapeFields,
	)
}

type EndArcShapeFormCallback = FormCallback[*models.EndArcShape]

func saveEndArcShapeFields(
	_instance *models.EndArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(_instance.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(_instance.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(_instance.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(_instance.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(_instance.RadiusY), formDiv)
		case "EndArcShapeGrid:EndArcShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "EndArcShapes", func(owner *models.EndArcShapeGrid) *[]*models.EndArcShape { return &owner.EndArcShapes })
		}
	}
}

func __gong__New__EndArcShapeGridFormCallback(
	_instance *models.EndArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (endarcshapegridFormCallback *FormCallback[*models.EndArcShapeGrid]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEndArcShapeGridFields,
	)
}

type EndArcShapeGridFormCallback = FormCallback[*models.EndArcShapeGrid]

func saveEndArcShapeGridFields(
	_instance *models.EndArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "EndArcShapes":
			FormDivSliceOfPointersToField(_instance, "EndArcShapes", &(_instance.EndArcShapes), formDiv, probe)
		}
	}
}

func __gong__New__EndHalfwayArcShapeFormCallback(
	_instance *models.EndHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (endhalfwayarcshapeFormCallback *FormCallback[*models.EndHalfwayArcShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEndHalfwayArcShapeFields,
	)
}

type EndHalfwayArcShapeFormCallback = FormCallback[*models.EndHalfwayArcShape]

func saveEndHalfwayArcShapeFields(
	_instance *models.EndHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(_instance.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(_instance.RadiusY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(_instance.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(_instance.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(_instance.SweepFlag), formDiv)
		case "EndHalfwayArcShapeGrid:EndHalfwayArcShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "EndHalfwayArcShapes", func(owner *models.EndHalfwayArcShapeGrid) *[]*models.EndHalfwayArcShape { return &owner.EndHalfwayArcShapes })
		}
	}
}

func __gong__New__EndHalfwayArcShapeGridFormCallback(
	_instance *models.EndHalfwayArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (endhalfwayarcshapegridFormCallback *FormCallback[*models.EndHalfwayArcShapeGrid]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEndHalfwayArcShapeGridFields,
	)
}

type EndHalfwayArcShapeGridFormCallback = FormCallback[*models.EndHalfwayArcShapeGrid]

func saveEndHalfwayArcShapeGridFields(
	_instance *models.EndHalfwayArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "EndHalfwayArcShapes":
			FormDivSliceOfPointersToField(_instance, "EndHalfwayArcShapes", &(_instance.EndHalfwayArcShapes), formDiv, probe)
		}
	}
}

func __gong__New__ExplanationTextShapeFormCallback(
	_instance *models.ExplanationTextShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (explanationtextshapeFormCallback *FormCallback[*models.ExplanationTextShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveExplanationTextShapeFields,
	)
}

type ExplanationTextShapeFormCallback = FormCallback[*models.ExplanationTextShape]

func saveExplanationTextShapeFields(
	_instance *models.ExplanationTextShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__Eye3DShapeFormCallback(
	_instance *models.Eye3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (eye3dshapeFormCallback *FormCallback[*models.Eye3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEye3DShapeFields,
	)
}

type Eye3DShapeFormCallback = FormCallback[*models.Eye3DShape]

func saveEye3DShapeFields(
	_instance *models.Eye3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__EyeCornersSampledPoints3DShapeFormCallback(
	_instance *models.EyeCornersSampledPoints3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (eyecornerssampledpoints3dshapeFormCallback *FormCallback[*models.EyeCornersSampledPoints3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEyeCornersSampledPoints3DShapeFields,
	)
}

type EyeCornersSampledPoints3DShapeFormCallback = FormCallback[*models.EyeCornersSampledPoints3DShape]

func saveEyeCornersSampledPoints3DShapeFields(
	_instance *models.EyeCornersSampledPoints3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__EyeSampledPoints3DShapeFormCallback(
	_instance *models.EyeSampledPoints3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (eyesampledpoints3dshapeFormCallback *FormCallback[*models.EyeSampledPoints3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEyeSampledPoints3DShapeFields,
	)
}

type EyeSampledPoints3DShapeFormCallback = FormCallback[*models.EyeSampledPoints3DShape]

func saveEyeSampledPoints3DShapeFields(
	_instance *models.EyeSampledPoints3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__EyeSeatBottomCurveShapeFormCallback(
	_instance *models.EyeSeatBottomCurveShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (eyeseatbottomcurveshapeFormCallback *FormCallback[*models.EyeSeatBottomCurveShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEyeSeatBottomCurveShapeFields,
	)
}

type EyeSeatBottomCurveShapeFormCallback = FormCallback[*models.EyeSeatBottomCurveShape]

func saveEyeSeatBottomCurveShapeFields(
	_instance *models.EyeSeatBottomCurveShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__EyeStoolBottomCurveShapeFormCallback(
	_instance *models.EyeStoolBottomCurveShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (eyestoolbottomcurveshapeFormCallback *FormCallback[*models.EyeStoolBottomCurveShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEyeStoolBottomCurveShapeFields,
	)
}

type EyeStoolBottomCurveShapeFormCallback = FormCallback[*models.EyeStoolBottomCurveShape]

func saveEyeStoolBottomCurveShapeFields(
	_instance *models.EyeStoolBottomCurveShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__EyeVolume3DShapeFormCallback(
	_instance *models.EyeVolume3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (eyevolume3dshapeFormCallback *FormCallback[*models.EyeVolume3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEyeVolume3DShapeFields,
	)
}

type EyeVolume3DShapeFormCallback = FormCallback[*models.EyeVolume3DShape]

func saveEyeVolume3DShapeFields(
	_instance *models.EyeVolume3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__GridPathShapeFormCallback(
	_instance *models.GridPathShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (gridpathshapeFormCallback *FormCallback[*models.GridPathShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGridPathShapeFields,
	)
}

type GridPathShapeFormCallback = FormCallback[*models.GridPathShape]

func saveGridPathShapeFields(
	_instance *models.GridPathShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__GrowthCurve2DFormCallback(
	_instance *models.GrowthCurve2D,
	probe *Probe,
	formGroup *form.FormGroup,
) (growthcurve2dFormCallback *FormCallback[*models.GrowthCurve2D]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGrowthCurve2DFields,
	)
}

type GrowthCurve2DFormCallback = FormCallback[*models.GrowthCurve2D]

func saveGrowthCurve2DFields(
	_instance *models.GrowthCurve2D,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartHalfwayArcShapeGrid":
			FormDivSelectFieldToField(&(_instance.StartHalfwayArcShapeGrid), probe.stageOfInterest, formDiv)
		case "EndHalfwayArcShapeGrid":
			FormDivSelectFieldToField(&(_instance.EndHalfwayArcShapeGrid), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__GrowthCurve2DRibbonFormCallback(
	_instance *models.GrowthCurve2DRibbon,
	probe *Probe,
	formGroup *form.FormGroup,
) (growthcurve2dribbonFormCallback *FormCallback[*models.GrowthCurve2DRibbon]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGrowthCurve2DRibbonFields,
	)
}

type GrowthCurve2DRibbonFormCallback = FormCallback[*models.GrowthCurve2DRibbon]

func saveGrowthCurve2DRibbonFields(
	_instance *models.GrowthCurve2DRibbon,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "GrowthCurve2DRibbonStartShapes":
			FormDivSliceOfPointersToField(_instance, "GrowthCurve2DRibbonStartShapes", &(_instance.GrowthCurve2DRibbonStartShapes), formDiv, probe)
		case "GrowthCurve2DRibbonEndShapes":
			FormDivSliceOfPointersToField(_instance, "GrowthCurve2DRibbonEndShapes", &(_instance.GrowthCurve2DRibbonEndShapes), formDiv, probe)
		}
	}
}

func __gong__New__GrowthCurve2DRibbonEndShapeFormCallback(
	_instance *models.GrowthCurve2DRibbonEndShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (growthcurve2dribbonendshapeFormCallback *FormCallback[*models.GrowthCurve2DRibbonEndShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGrowthCurve2DRibbonEndShapeFields,
	)
}

type GrowthCurve2DRibbonEndShapeFormCallback = FormCallback[*models.GrowthCurve2DRibbonEndShape]

func saveGrowthCurve2DRibbonEndShapeFields(
	_instance *models.GrowthCurve2DRibbonEndShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(_instance.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(_instance.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(_instance.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(_instance.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(_instance.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(_instance.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(_instance.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(_instance.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(_instance.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(_instance.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(_instance.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(_instance.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(_instance.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(_instance.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(_instance.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(_instance.TopSweepFlag), formDiv)
		case "GrowthCurve2DRibbon:GrowthCurve2DRibbonEndShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "GrowthCurve2DRibbonEndShapes", func(owner *models.GrowthCurve2DRibbon) *[]*models.GrowthCurve2DRibbonEndShape { return &owner.GrowthCurve2DRibbonEndShapes })
		}
	}
}

func __gong__New__GrowthCurve2DRibbonStartShapeFormCallback(
	_instance *models.GrowthCurve2DRibbonStartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (growthcurve2dribbonstartshapeFormCallback *FormCallback[*models.GrowthCurve2DRibbonStartShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGrowthCurve2DRibbonStartShapeFields,
	)
}

type GrowthCurve2DRibbonStartShapeFormCallback = FormCallback[*models.GrowthCurve2DRibbonStartShape]

func saveGrowthCurve2DRibbonStartShapeFields(
	_instance *models.GrowthCurve2DRibbonStartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(_instance.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(_instance.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(_instance.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(_instance.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(_instance.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(_instance.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(_instance.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(_instance.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(_instance.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(_instance.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(_instance.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(_instance.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(_instance.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(_instance.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(_instance.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(_instance.TopSweepFlag), formDiv)
		case "GrowthCurve2DRibbon:GrowthCurve2DRibbonStartShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "GrowthCurve2DRibbonStartShapes", func(owner *models.GrowthCurve2DRibbon) *[]*models.GrowthCurve2DRibbonStartShape { return &owner.GrowthCurve2DRibbonStartShapes })
		}
	}
}

func __gong__New__GrowthCurveRhombusGridShapeFormCallback(
	_instance *models.GrowthCurveRhombusGridShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (growthcurverhombusgridshapeFormCallback *FormCallback[*models.GrowthCurveRhombusGridShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGrowthCurveRhombusGridShapeFields,
	)
}

type GrowthCurveRhombusGridShapeFormCallback = FormCallback[*models.GrowthCurveRhombusGridShape]

func saveGrowthCurveRhombusGridShapeFields(
	_instance *models.GrowthCurveRhombusGridShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "GrowthCurveRhombusShapes":
			FormDivSliceOfPointersToField(_instance, "GrowthCurveRhombusShapes", &(_instance.GrowthCurveRhombusShapes), formDiv, probe)
		}
	}
}

func __gong__New__GrowthCurveRhombusShapeFormCallback(
	_instance *models.GrowthCurveRhombusShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (growthcurverhombusshapeFormCallback *FormCallback[*models.GrowthCurveRhombusShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGrowthCurveRhombusShapeFields,
	)
}

type GrowthCurveRhombusShapeFormCallback = FormCallback[*models.GrowthCurveRhombusShape]

func saveGrowthCurveRhombusShapeFields(
	_instance *models.GrowthCurveRhombusShape,
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
		case "GrowthCurveRhombusGridShape:GrowthCurveRhombusShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "GrowthCurveRhombusShapes", func(owner *models.GrowthCurveRhombusGridShape) *[]*models.GrowthCurveRhombusShape { return &owner.GrowthCurveRhombusShapes })
		}
	}
}

func __gong__New__GrowthVectorShapeFormCallback(
	_instance *models.GrowthVectorShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (growthvectorshapeFormCallback *FormCallback[*models.GrowthVectorShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveGrowthVectorShapeFields,
	)
}

type GrowthVectorShapeFormCallback = FormCallback[*models.GrowthVectorShape]

func saveGrowthVectorShapeFields(
	_instance *models.GrowthVectorShape,
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

func __gong__New__InitialRhombusGridShapeFormCallback(
	_instance *models.InitialRhombusGridShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (initialrhombusgridshapeFormCallback *FormCallback[*models.InitialRhombusGridShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveInitialRhombusGridShapeFields,
	)
}

type InitialRhombusGridShapeFormCallback = FormCallback[*models.InitialRhombusGridShape]

func saveInitialRhombusGridShapeFields(
	_instance *models.InitialRhombusGridShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "InitialRhombusShapes":
			FormDivSliceOfPointersToField(_instance, "InitialRhombusShapes", &(_instance.InitialRhombusShapes), formDiv, probe)
		}
	}
}

func __gong__New__InitialRhombusShapeFormCallback(
	_instance *models.InitialRhombusShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (initialrhombusshapeFormCallback *FormCallback[*models.InitialRhombusShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveInitialRhombusShapeFields,
	)
}

type InitialRhombusShapeFormCallback = FormCallback[*models.InitialRhombusShape]

func saveInitialRhombusShapeFields(
	_instance *models.InitialRhombusShape,
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
		case "InitialRhombusGridShape:InitialRhombusShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "InitialRhombusShapes", func(owner *models.InitialRhombusGridShape) *[]*models.InitialRhombusShape { return &owner.InitialRhombusShapes })
		}
	}
}

func __gong__New__Key3DShapeFormCallback(
	_instance *models.Key3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (key3dshapeFormCallback *FormCallback[*models.Key3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveKey3DShapeFields,
	)
}

type Key3DShapeFormCallback = FormCallback[*models.Key3DShape]

func saveKey3DShapeFields(
	_instance *models.Key3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__KeyHole3DShapeFormCallback(
	_instance *models.KeyHole3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (keyhole3dshapeFormCallback *FormCallback[*models.KeyHole3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveKeyHole3DShapeFields,
	)
}

type KeyHole3DShapeFormCallback = FormCallback[*models.KeyHole3DShape]

func saveKeyHole3DShapeFields(
	_instance *models.KeyHole3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__KeyHoleShapeFormCallback(
	_instance *models.KeyHoleShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (keyholeshapeFormCallback *FormCallback[*models.KeyHoleShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveKeyHoleShapeFields,
	)
}

type KeyHoleShapeFormCallback = FormCallback[*models.KeyHoleShape]

func saveKeyHoleShapeFields(
	_instance *models.KeyHoleShape,
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
		}
	}
}

func __gong__New__Leaves3DShapeFormCallback(
	_instance *models.Leaves3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (leaves3dshapeFormCallback *FormCallback[*models.Leaves3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLeaves3DShapeFields,
	)
}

type Leaves3DShapeFormCallback = FormCallback[*models.Leaves3DShape]

func saveLeaves3DShapeFields(
	_instance *models.Leaves3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__LibraryFormCallback(
	_instance *models.Library,
	probe *Probe,
	formGroup *form.FormGroup,
) (libraryFormCallback *FormCallback[*models.Library]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLibraryFields,
	)
}

type LibraryFormCallback = FormCallback[*models.Library]

func saveLibraryFields(
	_instance *models.Library,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Plants":
			FormDivSliceOfPointersToField(_instance, "Plants", &(_instance.Plants), formDiv, probe)
		case "SubLibraries":
			FormDivSliceOfPointersToField(_instance, "SubLibraries", &(_instance.SubLibraries), formDiv, probe)
		case "NbPixPerCharacter":
			FormDivBasicFieldToField(&(_instance.NbPixPerCharacter), formDiv)
		case "LogoSVGFile":
			FormDivBasicFieldToField(&(_instance.LogoSVGFile), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "IsRootLibrary":
			FormDivBasicFieldToField(&(_instance.IsRootLibrary), formDiv)
		case "Library:SubLibraries":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubLibraries", func(owner *models.Library) *[]*models.Library { return &owner.SubLibraries })
		}
	}
}

func __gong__New__MidArcVectorShapeFormCallback(
	_instance *models.MidArcVectorShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (midarcvectorshapeFormCallback *FormCallback[*models.MidArcVectorShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMidArcVectorShapeFields,
	)
}

type MidArcVectorShapeFormCallback = FormCallback[*models.MidArcVectorShape]

func saveMidArcVectorShapeFields(
	_instance *models.MidArcVectorShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "MidArcVectorShapeGrid:MidArcVectorShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "MidArcVectorShapes", func(owner *models.MidArcVectorShapeGrid) *[]*models.MidArcVectorShape { return &owner.MidArcVectorShapes })
		}
	}
}

func __gong__New__MidArcVectorShapeGridFormCallback(
	_instance *models.MidArcVectorShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (midarcvectorshapegridFormCallback *FormCallback[*models.MidArcVectorShapeGrid]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMidArcVectorShapeGridFields,
	)
}

type MidArcVectorShapeGridFormCallback = FormCallback[*models.MidArcVectorShapeGrid]

func saveMidArcVectorShapeGridFields(
	_instance *models.MidArcVectorShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "MidArcVectorShapes":
			FormDivSliceOfPointersToField(_instance, "MidArcVectorShapes", &(_instance.MidArcVectorShapes), formDiv, probe)
		}
	}
}

func __gong__New__MusicAbstractFormCallback(
	_instance *models.MusicAbstract,
	probe *Probe,
	formGroup *form.FormGroup,
) (musicabstractFormCallback *FormCallback[*models.MusicAbstract]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveMusicAbstractFields,
	)
}

type MusicAbstractFormCallback = FormCallback[*models.MusicAbstract]

func saveMusicAbstractFields(
	_instance *models.MusicAbstract,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(_instance.IsChecked), formDiv)
		case "PitchHeight":
			FormDivBasicFieldToField(&(_instance.PitchHeight), formDiv)
		case "NbOfBeatsInTheme":
			FormDivBasicFieldToField(&(_instance.NbOfBeatsInTheme), formDiv)
		case "BeatsPerSecond":
			FormDivBasicFieldToField(&(_instance.BeatsPerSecond), formDiv)
		case "FirstVoiceShiftX":
			FormDivBasicFieldToField(&(_instance.FirstVoiceShiftX), formDiv)
		case "FirstVoiceShiftY":
			FormDivBasicFieldToField(&(_instance.FirstVoiceShiftY), formDiv)
		case "PitchDifference":
			FormDivBasicFieldToField(&(_instance.PitchDifference), formDiv)
		case "Level":
			FormDivBasicFieldToField(&(_instance.Level), formDiv)
		case "ActualBeatsTemporalShift":
			FormDivBasicFieldToField(&(_instance.ActualBeatsTemporalShift), formDiv)
		case "IsMinor":
			FormDivBasicFieldToField(&(_instance.IsMinor), formDiv)
		case "ThemeBinaryEncoding":
			FormDivBasicFieldToField(&(_instance.ThemeBinaryEncoding), formDiv)
		case "BezierControlLengthRatio":
			FormDivBasicFieldToField(&(_instance.BezierControlLengthRatio), formDiv)
		case "NbPitchLines":
			FormDivBasicFieldToField(&(_instance.NbPitchLines), formDiv)
		case "NbBeatLines":
			FormDivBasicFieldToField(&(_instance.NbBeatLines), formDiv)
		case "OriginX":
			FormDivBasicFieldToField(&(_instance.OriginX), formDiv)
		case "OriginY":
			FormDivBasicFieldToField(&(_instance.OriginY), formDiv)
		case "ScoreScale":
			FormDivBasicFieldToField(&(_instance.ScoreScale), formDiv)
		case "ShowFirstVoice":
			FormDivBasicFieldToField(&(_instance.ShowFirstVoice), formDiv)
		case "ShowFirstVoiceShiftRight":
			FormDivBasicFieldToField(&(_instance.ShowFirstVoiceShiftRight), formDiv)
		case "ShowSecondVoice":
			FormDivBasicFieldToField(&(_instance.ShowSecondVoice), formDiv)
		case "ShowSecondVoiceShiftRight":
			FormDivBasicFieldToField(&(_instance.ShowSecondVoiceShiftRight), formDiv)
		case "ShowFirstVoiceNotes":
			FormDivBasicFieldToField(&(_instance.ShowFirstVoiceNotes), formDiv)
		case "ShowFirstVoiceNotesShiftRight":
			FormDivBasicFieldToField(&(_instance.ShowFirstVoiceNotesShiftRight), formDiv)
		case "ShowSecondVoiceNotes":
			FormDivBasicFieldToField(&(_instance.ShowSecondVoiceNotes), formDiv)
		case "ShowSecondVoiceNotesShiftRight":
			FormDivBasicFieldToField(&(_instance.ShowSecondVoiceNotesShiftRight), formDiv)
		case "IsComposerNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsComposerNodeExpanded), formDiv)
		}
	}
}

func __gong__New__OriginalPoints3DShapeFormCallback(
	_instance *models.OriginalPoints3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (originalpoints3dshapeFormCallback *FormCallback[*models.OriginalPoints3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveOriginalPoints3DShapeFields,
	)
}

type OriginalPoints3DShapeFormCallback = FormCallback[*models.OriginalPoints3DShape]

func saveOriginalPoints3DShapeFields(
	_instance *models.OriginalPoints3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__ParastichyMCurves3DShapeFormCallback(
	_instance *models.ParastichyMCurves3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (parastichymcurves3dshapeFormCallback *FormCallback[*models.ParastichyMCurves3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveParastichyMCurves3DShapeFields,
	)
}

type ParastichyMCurves3DShapeFormCallback = FormCallback[*models.ParastichyMCurves3DShape]

func saveParastichyMCurves3DShapeFields(
	_instance *models.ParastichyMCurves3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__ParastichyNCurves3DShapeFormCallback(
	_instance *models.ParastichyNCurves3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (parastichyncurves3dshapeFormCallback *FormCallback[*models.ParastichyNCurves3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveParastichyNCurves3DShapeFields,
	)
}

type ParastichyNCurves3DShapeFormCallback = FormCallback[*models.ParastichyNCurves3DShape]

func saveParastichyNCurves3DShapeFields(
	_instance *models.ParastichyNCurves3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__PartiallyGrowthCurve2DRibbonFormCallback(
	_instance *models.PartiallyGrowthCurve2DRibbon,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallygrowthcurve2dribbonFormCallback *FormCallback[*models.PartiallyGrowthCurve2DRibbon]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePartiallyGrowthCurve2DRibbonFields,
	)
}

type PartiallyGrowthCurve2DRibbonFormCallback = FormCallback[*models.PartiallyGrowthCurve2DRibbon]

func savePartiallyGrowthCurve2DRibbonFields(
	_instance *models.PartiallyGrowthCurve2DRibbon,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "PartiallyGrowthCurve2DRibbonStartShapes":
			FormDivSliceOfPointersToField(_instance, "PartiallyGrowthCurve2DRibbonStartShapes", &(_instance.PartiallyGrowthCurve2DRibbonStartShapes), formDiv, probe)
		case "PartiallyGrowthCurve2DRibbonEndShapes":
			FormDivSliceOfPointersToField(_instance, "PartiallyGrowthCurve2DRibbonEndShapes", &(_instance.PartiallyGrowthCurve2DRibbonEndShapes), formDiv, probe)
		}
	}
}

func __gong__New__PartiallyGrowthCurve2DRibbonEndShapeFormCallback(
	_instance *models.PartiallyGrowthCurve2DRibbonEndShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallygrowthcurve2dribbonendshapeFormCallback *FormCallback[*models.PartiallyGrowthCurve2DRibbonEndShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePartiallyGrowthCurve2DRibbonEndShapeFields,
	)
}

type PartiallyGrowthCurve2DRibbonEndShapeFormCallback = FormCallback[*models.PartiallyGrowthCurve2DRibbonEndShape]

func savePartiallyGrowthCurve2DRibbonEndShapeFields(
	_instance *models.PartiallyGrowthCurve2DRibbonEndShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(_instance.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(_instance.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(_instance.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(_instance.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(_instance.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(_instance.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(_instance.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(_instance.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(_instance.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(_instance.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(_instance.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(_instance.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(_instance.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(_instance.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(_instance.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(_instance.TopSweepFlag), formDiv)
		case "PartiallyGrowthCurve2DRibbon:PartiallyGrowthCurve2DRibbonEndShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "PartiallyGrowthCurve2DRibbonEndShapes", func(owner *models.PartiallyGrowthCurve2DRibbon) *[]*models.PartiallyGrowthCurve2DRibbonEndShape { return &owner.PartiallyGrowthCurve2DRibbonEndShapes })
		}
	}
}

func __gong__New__PartiallyGrowthCurve2DRibbonStartShapeFormCallback(
	_instance *models.PartiallyGrowthCurve2DRibbonStartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallygrowthcurve2dribbonstartshapeFormCallback *FormCallback[*models.PartiallyGrowthCurve2DRibbonStartShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePartiallyGrowthCurve2DRibbonStartShapeFields,
	)
}

type PartiallyGrowthCurve2DRibbonStartShapeFormCallback = FormCallback[*models.PartiallyGrowthCurve2DRibbonStartShape]

func savePartiallyGrowthCurve2DRibbonStartShapeFields(
	_instance *models.PartiallyGrowthCurve2DRibbonStartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(_instance.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(_instance.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(_instance.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(_instance.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(_instance.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(_instance.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(_instance.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(_instance.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(_instance.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(_instance.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(_instance.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(_instance.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(_instance.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(_instance.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(_instance.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(_instance.TopSweepFlag), formDiv)
		case "PartiallyGrowthCurve2DRibbon:PartiallyGrowthCurve2DRibbonStartShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "PartiallyGrowthCurve2DRibbonStartShapes", func(owner *models.PartiallyGrowthCurve2DRibbon) *[]*models.PartiallyGrowthCurve2DRibbonStartShape { return &owner.PartiallyGrowthCurve2DRibbonStartShapes })
		}
	}
}

func __gong__New__PartiallyGrowthCurve2DTrajectoryFormCallback(
	_instance *models.PartiallyGrowthCurve2DTrajectory,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallygrowthcurve2dtrajectoryFormCallback *FormCallback[*models.PartiallyGrowthCurve2DTrajectory]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePartiallyGrowthCurve2DTrajectoryFields,
	)
}

type PartiallyGrowthCurve2DTrajectoryFormCallback = FormCallback[*models.PartiallyGrowthCurve2DTrajectory]

func savePartiallyGrowthCurve2DTrajectoryFields(
	_instance *models.PartiallyGrowthCurve2DTrajectory,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "PartiallyGrowthCurve2DTrajectoryShapes":
			FormDivSliceOfPointersToField(_instance, "PartiallyGrowthCurve2DTrajectoryShapes", &(_instance.PartiallyGrowthCurve2DTrajectoryShapes), formDiv, probe)
		}
	}
}

func __gong__New__PartiallyGrowthCurve2DTrajectoryP1CurveShapeFormCallback(
	_instance *models.PartiallyGrowthCurve2DTrajectoryP1CurveShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallygrowthcurve2dtrajectoryp1curveshapeFormCallback *FormCallback[*models.PartiallyGrowthCurve2DTrajectoryP1CurveShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePartiallyGrowthCurve2DTrajectoryP1CurveShapeFields,
	)
}

type PartiallyGrowthCurve2DTrajectoryP1CurveShapeFormCallback = FormCallback[*models.PartiallyGrowthCurve2DTrajectoryP1CurveShape]

func savePartiallyGrowthCurve2DTrajectoryP1CurveShapeFields(
	_instance *models.PartiallyGrowthCurve2DTrajectoryP1CurveShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "PartiallyGrowthCurve2DTrajectoryP1P2:P1CurveShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "P1CurveShapes", func(owner *models.PartiallyGrowthCurve2DTrajectoryP1P2) *[]*models.PartiallyGrowthCurve2DTrajectoryP1CurveShape { return &owner.P1CurveShapes })
		}
	}
}

func __gong__New__PartiallyGrowthCurve2DTrajectoryP1P2FormCallback(
	_instance *models.PartiallyGrowthCurve2DTrajectoryP1P2,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallygrowthcurve2dtrajectoryp1p2FormCallback *FormCallback[*models.PartiallyGrowthCurve2DTrajectoryP1P2]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePartiallyGrowthCurve2DTrajectoryP1P2Fields,
	)
}

type PartiallyGrowthCurve2DTrajectoryP1P2FormCallback = FormCallback[*models.PartiallyGrowthCurve2DTrajectoryP1P2]

func savePartiallyGrowthCurve2DTrajectoryP1P2Fields(
	_instance *models.PartiallyGrowthCurve2DTrajectoryP1P2,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "P1PointShapes":
			FormDivSliceOfPointersToField(_instance, "P1PointShapes", &(_instance.P1PointShapes), formDiv, probe)
		case "P2PointShapes":
			FormDivSliceOfPointersToField(_instance, "P2PointShapes", &(_instance.P2PointShapes), formDiv, probe)
		case "P1CurveShapes":
			FormDivSliceOfPointersToField(_instance, "P1CurveShapes", &(_instance.P1CurveShapes), formDiv, probe)
		case "P2CurveShapes":
			FormDivSliceOfPointersToField(_instance, "P2CurveShapes", &(_instance.P2CurveShapes), formDiv, probe)
		case "P1P2PairLineShapes":
			FormDivSliceOfPointersToField(_instance, "P1P2PairLineShapes", &(_instance.P1P2PairLineShapes), formDiv, probe)
		}
	}
}

func __gong__New__PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeFormCallback(
	_instance *models.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallygrowthcurve2dtrajectoryp1p2pairlineshapeFormCallback *FormCallback[*models.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeFields,
	)
}

type PartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeFormCallback = FormCallback[*models.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape]

func savePartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeFields(
	_instance *models.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "PartiallyGrowthCurve2DTrajectoryP1P2:P1P2PairLineShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "P1P2PairLineShapes", func(owner *models.PartiallyGrowthCurve2DTrajectoryP1P2) *[]*models.PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape { return &owner.P1P2PairLineShapes })
		}
	}
}

func __gong__New__PartiallyGrowthCurve2DTrajectoryP1PointShapeFormCallback(
	_instance *models.PartiallyGrowthCurve2DTrajectoryP1PointShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallygrowthcurve2dtrajectoryp1pointshapeFormCallback *FormCallback[*models.PartiallyGrowthCurve2DTrajectoryP1PointShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePartiallyGrowthCurve2DTrajectoryP1PointShapeFields,
	)
}

type PartiallyGrowthCurve2DTrajectoryP1PointShapeFormCallback = FormCallback[*models.PartiallyGrowthCurve2DTrajectoryP1PointShape]

func savePartiallyGrowthCurve2DTrajectoryP1PointShapeFields(
	_instance *models.PartiallyGrowthCurve2DTrajectoryP1PointShape,
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
		case "PartiallyGrowthCurve2DTrajectoryP1P2:P1PointShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "P1PointShapes", func(owner *models.PartiallyGrowthCurve2DTrajectoryP1P2) *[]*models.PartiallyGrowthCurve2DTrajectoryP1PointShape { return &owner.P1PointShapes })
		}
	}
}

func __gong__New__PartiallyGrowthCurve2DTrajectoryP2CurveShapeFormCallback(
	_instance *models.PartiallyGrowthCurve2DTrajectoryP2CurveShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallygrowthcurve2dtrajectoryp2curveshapeFormCallback *FormCallback[*models.PartiallyGrowthCurve2DTrajectoryP2CurveShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePartiallyGrowthCurve2DTrajectoryP2CurveShapeFields,
	)
}

type PartiallyGrowthCurve2DTrajectoryP2CurveShapeFormCallback = FormCallback[*models.PartiallyGrowthCurve2DTrajectoryP2CurveShape]

func savePartiallyGrowthCurve2DTrajectoryP2CurveShapeFields(
	_instance *models.PartiallyGrowthCurve2DTrajectoryP2CurveShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "PartiallyGrowthCurve2DTrajectoryP1P2:P2CurveShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "P2CurveShapes", func(owner *models.PartiallyGrowthCurve2DTrajectoryP1P2) *[]*models.PartiallyGrowthCurve2DTrajectoryP2CurveShape { return &owner.P2CurveShapes })
		}
	}
}

func __gong__New__PartiallyGrowthCurve2DTrajectoryP2PointShapeFormCallback(
	_instance *models.PartiallyGrowthCurve2DTrajectoryP2PointShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallygrowthcurve2dtrajectoryp2pointshapeFormCallback *FormCallback[*models.PartiallyGrowthCurve2DTrajectoryP2PointShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePartiallyGrowthCurve2DTrajectoryP2PointShapeFields,
	)
}

type PartiallyGrowthCurve2DTrajectoryP2PointShapeFormCallback = FormCallback[*models.PartiallyGrowthCurve2DTrajectoryP2PointShape]

func savePartiallyGrowthCurve2DTrajectoryP2PointShapeFields(
	_instance *models.PartiallyGrowthCurve2DTrajectoryP2PointShape,
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
		case "PartiallyGrowthCurve2DTrajectoryP1P2:P2PointShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "P2PointShapes", func(owner *models.PartiallyGrowthCurve2DTrajectoryP1P2) *[]*models.PartiallyGrowthCurve2DTrajectoryP2PointShape { return &owner.P2PointShapes })
		}
	}
}

func __gong__New__PartiallyGrowthCurve2DTrajectoryShapeFormCallback(
	_instance *models.PartiallyGrowthCurve2DTrajectoryShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallygrowthcurve2dtrajectoryshapeFormCallback *FormCallback[*models.PartiallyGrowthCurve2DTrajectoryShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePartiallyGrowthCurve2DTrajectoryShapeFields,
	)
}

type PartiallyGrowthCurve2DTrajectoryShapeFormCallback = FormCallback[*models.PartiallyGrowthCurve2DTrajectoryShape]

func savePartiallyGrowthCurve2DTrajectoryShapeFields(
	_instance *models.PartiallyGrowthCurve2DTrajectoryShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "PartiallyGrowthCurve2DTrajectory:PartiallyGrowthCurve2DTrajectoryShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "PartiallyGrowthCurve2DTrajectoryShapes", func(owner *models.PartiallyGrowthCurve2DTrajectory) *[]*models.PartiallyGrowthCurve2DTrajectoryShape { return &owner.PartiallyGrowthCurve2DTrajectoryShapes })
		}
	}
}

func __gong__New__PartiallyRotatedSeatBottomCurveShapeFormCallback(
	_instance *models.PartiallyRotatedSeatBottomCurveShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallyrotatedseatbottomcurveshapeFormCallback *FormCallback[*models.PartiallyRotatedSeatBottomCurveShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePartiallyRotatedSeatBottomCurveShapeFields,
	)
}

type PartiallyRotatedSeatBottomCurveShapeFormCallback = FormCallback[*models.PartiallyRotatedSeatBottomCurveShape]

func savePartiallyRotatedSeatBottomCurveShapeFields(
	_instance *models.PartiallyRotatedSeatBottomCurveShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__PartiallyRotatedSeatTopCurveShapeFormCallback(
	_instance *models.PartiallyRotatedSeatTopCurveShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallyrotatedseattopcurveshapeFormCallback *FormCallback[*models.PartiallyRotatedSeatTopCurveShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePartiallyRotatedSeatTopCurveShapeFields,
	)
}

type PartiallyRotatedSeatTopCurveShapeFormCallback = FormCallback[*models.PartiallyRotatedSeatTopCurveShape]

func savePartiallyRotatedSeatTopCurveShapeFields(
	_instance *models.PartiallyRotatedSeatTopCurveShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__PartiallyRotatedTorusShapeFormCallback(
	_instance *models.PartiallyRotatedTorusShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (partiallyrotatedtorusshapeFormCallback *FormCallback[*models.PartiallyRotatedTorusShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePartiallyRotatedTorusShapeFields,
	)
}

type PartiallyRotatedTorusShapeFormCallback = FormCallback[*models.PartiallyRotatedTorusShape]

func savePartiallyRotatedTorusShapeFields(
	_instance *models.PartiallyRotatedTorusShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__PerpendicularVectorFormCallback(
	_instance *models.PerpendicularVector,
	probe *Probe,
	formGroup *form.FormGroup,
) (perpendicularvectorFormCallback *FormCallback[*models.PerpendicularVector]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePerpendicularVectorFields,
	)
}

type PerpendicularVectorFormCallback = FormCallback[*models.PerpendicularVector]

func savePerpendicularVectorFields(
	_instance *models.PerpendicularVector,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "PerpendicularVectorGrid:PerpendicularVectors":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "PerpendicularVectors", func(owner *models.PerpendicularVectorGrid) *[]*models.PerpendicularVector { return &owner.PerpendicularVectors })
		}
	}
}

func __gong__New__PerpendicularVectorGridFormCallback(
	_instance *models.PerpendicularVectorGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (perpendicularvectorgridFormCallback *FormCallback[*models.PerpendicularVectorGrid]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePerpendicularVectorGridFields,
	)
}

type PerpendicularVectorGridFormCallback = FormCallback[*models.PerpendicularVectorGrid]

func savePerpendicularVectorGridFields(
	_instance *models.PerpendicularVectorGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "PerpendicularVectors":
			FormDivSliceOfPointersToField(_instance, "PerpendicularVectors", &(_instance.PerpendicularVectors), formDiv, probe)
		}
	}
}

func __gong__New__PerpendicularVectorGridHalfwayFormCallback(
	_instance *models.PerpendicularVectorGridHalfway,
	probe *Probe,
	formGroup *form.FormGroup,
) (perpendicularvectorgridhalfwayFormCallback *FormCallback[*models.PerpendicularVectorGridHalfway]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePerpendicularVectorGridHalfwayFields,
	)
}

type PerpendicularVectorGridHalfwayFormCallback = FormCallback[*models.PerpendicularVectorGridHalfway]

func savePerpendicularVectorGridHalfwayFields(
	_instance *models.PerpendicularVectorGridHalfway,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "PerpendicularVectorHalfways":
			FormDivSliceOfPointersToField(_instance, "PerpendicularVectorHalfways", &(_instance.PerpendicularVectorHalfways), formDiv, probe)
		}
	}
}

func __gong__New__PerpendicularVectorHalfwayFormCallback(
	_instance *models.PerpendicularVectorHalfway,
	probe *Probe,
	formGroup *form.FormGroup,
) (perpendicularvectorhalfwayFormCallback *FormCallback[*models.PerpendicularVectorHalfway]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePerpendicularVectorHalfwayFields,
	)
}

type PerpendicularVectorHalfwayFormCallback = FormCallback[*models.PerpendicularVectorHalfway]

func savePerpendicularVectorHalfwayFields(
	_instance *models.PerpendicularVectorHalfway,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "PerpendicularVectorGridHalfway:PerpendicularVectorHalfways":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "PerpendicularVectorHalfways", func(owner *models.PerpendicularVectorGridHalfway) *[]*models.PerpendicularVectorHalfway { return &owner.PerpendicularVectorHalfways })
		}
	}
}

func __gong__New__Plant2DDiagramFormCallback(
	_instance *models.Plant2DDiagram,
	probe *Probe,
	formGroup *form.FormGroup,
) (plant2ddiagramFormCallback *FormCallback[*models.Plant2DDiagram]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePlant2DDiagramFields,
	)
}

type Plant2DDiagramFormCallback = FormCallback[*models.Plant2DDiagram]

func savePlant2DDiagramFields(
	_instance *models.Plant2DDiagram,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "OriginX":
			FormDivBasicFieldToField(&(_instance.OriginX), formDiv)
		case "OriginY":
			FormDivBasicFieldToField(&(_instance.OriginY), formDiv)
		case "Zoom":
			FormDivBasicFieldToField(&(_instance.Zoom), formDiv)
		case "IsRhombusNodesExpanded":
			FormDivBasicFieldToField(&(_instance.IsRhombusNodesExpanded), formDiv)
		case "IsArcNodesExpanded":
			FormDivBasicFieldToField(&(_instance.IsArcNodesExpanded), formDiv)
		case "IsHiddenAxesShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenAxesShape), formDiv)
		case "IsHiddenReferenceRhombus":
			FormDivBasicFieldToField(&(_instance.IsHiddenReferenceRhombus), formDiv)
		case "IsHiddenPlantCircumferenceShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenPlantCircumferenceShape), formDiv)
		case "IsHiddenGridPathShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenGridPathShape), formDiv)
		case "IsHiddenRhombusGridShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenRhombusGridShape), formDiv)
		case "IsHiddenExplanationTextShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenExplanationTextShape), formDiv)
		case "IsHiddenRotatedReferenceRhombus":
			FormDivBasicFieldToField(&(_instance.IsHiddenRotatedReferenceRhombus), formDiv)
		case "IsHiddenRotatedPlantCircumferenceShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenRotatedPlantCircumferenceShape), formDiv)
		case "IsHiddenRotatedGridPathShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenRotatedGridPathShape), formDiv)
		case "IsHiddenRotatedRhombusGridShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenRotatedRhombusGridShape), formDiv)
		case "IsHiddenGrowthPathRhombusGridShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenGrowthPathRhombusGridShape), formDiv)
		case "IsHiddenGrowthVectorShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenGrowthVectorShape), formDiv)
		case "IsHiddenPerpendicularVectorGrid":
			FormDivBasicFieldToField(&(_instance.IsHiddenPerpendicularVectorGrid), formDiv)
		case "IsHiddenBaseVectorShapeGrid":
			FormDivBasicFieldToField(&(_instance.IsHiddenBaseVectorShapeGrid), formDiv)
		case "IsHiddenArcNormalVectorShapeGrid":
			FormDivBasicFieldToField(&(_instance.IsHiddenArcNormalVectorShapeGrid), formDiv)
		case "IsHiddenStartArcShapeGrid":
			FormDivBasicFieldToField(&(_instance.IsHiddenStartArcShapeGrid), formDiv)
		case "IsHiddenMidArcVectorShapeGrid":
			FormDivBasicFieldToField(&(_instance.IsHiddenMidArcVectorShapeGrid), formDiv)
		case "IsHiddenEndArcShapeGrid":
			FormDivBasicFieldToField(&(_instance.IsHiddenEndArcShapeGrid), formDiv)
		case "IsHiddenGrowthCurve2D":
			FormDivBasicFieldToField(&(_instance.IsHiddenGrowthCurve2D), formDiv)
		case "IsHiddenStackOfGrowthCurve2DByGrowthVector":
			FormDivBasicFieldToField(&(_instance.IsHiddenStackOfGrowthCurve2DByGrowthVector), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(_instance.IsChecked), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "PlantAbstract:Plant2DDiagrams":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Plant2DDiagrams", func(owner *models.PlantAbstract) *[]*models.Plant2DDiagram { return &owner.Plant2DDiagrams })
		}
	}
}

func __gong__New__Plant3DDiagramFormCallback(
	_instance *models.Plant3DDiagram,
	probe *Probe,
	formGroup *form.FormGroup,
) (plant3ddiagramFormCallback *FormCallback[*models.Plant3DDiagram]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePlant3DDiagramFields,
	)
}

type Plant3DDiagramFormCallback = FormCallback[*models.Plant3DDiagram]

func savePlant3DDiagramFields(
	_instance *models.Plant3DDiagram,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "IsHiddenStemCylinder3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenStemCylinder3DShape), formDiv)
		case "StemCylinder3DShape":
			FormDivSelectFieldToField(&(_instance.StemCylinder3DShape), probe.stageOfInterest, formDiv)
		case "IsHiddenParastichyNCurves3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenParastichyNCurves3DShape), formDiv)
		case "ParastichyNCurves3DShape":
			FormDivSelectFieldToField(&(_instance.ParastichyNCurves3DShape), probe.stageOfInterest, formDiv)
		case "IsHiddenParastichyMCurves3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenParastichyMCurves3DShape), formDiv)
		case "ParastichyMCurves3DShape":
			FormDivSelectFieldToField(&(_instance.ParastichyMCurves3DShape), probe.stageOfInterest, formDiv)
		case "IsHiddenCutLine3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenCutLine3DShape), formDiv)
		case "CutLine3DShape":
			FormDivSelectFieldToField(&(_instance.CutLine3DShape), probe.stageOfInterest, formDiv)
		case "IsHiddenCircumference3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenCircumference3DShape), formDiv)
		case "Circumference3DShape":
			FormDivSelectFieldToField(&(_instance.Circumference3DShape), probe.stageOfInterest, formDiv)
		case "IsHiddenTiledFloor3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenTiledFloor3DShape), formDiv)
		case "TiledFloor3DShape":
			FormDivSelectFieldToField(&(_instance.TiledFloor3DShape), probe.stageOfInterest, formDiv)
		case "IsHiddenLeaves3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenLeaves3DShape), formDiv)
		case "Leaves3DShape":
			FormDivSelectFieldToField(&(_instance.Leaves3DShape), probe.stageOfInterest, formDiv)
		case "Rendered3DShape":
			FormDivSelectFieldToField(&(_instance.Rendered3DShape), probe.stageOfInterest, formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(_instance.IsChecked), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "PlantAbstract:Plant3DDiagrams":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Plant3DDiagrams", func(owner *models.PlantAbstract) *[]*models.Plant3DDiagram { return &owner.Plant3DDiagrams })
		}
	}
}

func __gong__New__PlantAbstractFormCallback(
	_instance *models.PlantAbstract,
	probe *Probe,
	formGroup *form.FormGroup,
) (plantabstractFormCallback *FormCallback[*models.PlantAbstract]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePlantAbstractFields,
	)
}

type PlantAbstractFormCallback = FormCallback[*models.PlantAbstract]

func savePlantAbstractFields(
	_instance *models.PlantAbstract,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "N":
			FormDivBasicFieldToField(&(_instance.N), formDiv)
		case "M":
			FormDivBasicFieldToField(&(_instance.M), formDiv)
		case "StackHeight":
			FormDivBasicFieldToField(&(_instance.StackHeight), formDiv)
		case "RhombusInsideAngle":
			FormDivBasicFieldToField(&(_instance.RhombusInsideAngle), formDiv)
		case "RhombusSideLength":
			FormDivBasicFieldToField(&(_instance.RhombusSideLength), formDiv)
		case "PlantType":
			FormDivEnumStringFieldToField(&(_instance.PlantType), formDiv)
		case "TubeVaseAbstract":
			FormDivSelectFieldToField(&(_instance.TubeVaseAbstract), probe.stageOfInterest, formDiv)
		case "StoolAbstract":
			FormDivSelectFieldToField(&(_instance.StoolAbstract), probe.stageOfInterest, formDiv)
		case "ClockAbstract":
			FormDivSelectFieldToField(&(_instance.ClockAbstract), probe.stageOfInterest, formDiv)
		case "MusicAbstract":
			FormDivSelectFieldToField(&(_instance.MusicAbstract), probe.stageOfInterest, formDiv)
		case "CurrentView":
			FormDivEnumStringFieldToField(&(_instance.CurrentView), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "IsSelected":
			FormDivBasicFieldToField(&(_instance.IsSelected), formDiv)
		case "IsPlant2DDiagramsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsPlant2DDiagramsNodeExpanded), formDiv)
		case "Plant2DDiagrams":
			FormDivSliceOfPointersToField(_instance, "Plant2DDiagrams", &(_instance.Plant2DDiagrams), formDiv, probe)
		case "IsPlant3DDiagramsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsPlant3DDiagramsNodeExpanded), formDiv)
		case "Plant3DDiagrams":
			FormDivSliceOfPointersToField(_instance, "Plant3DDiagrams", &(_instance.Plant3DDiagrams), formDiv, probe)
		case "IsVase2DDiagramsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsVase2DDiagramsNodeExpanded), formDiv)
		case "Vase2DDiagrams":
			FormDivSliceOfPointersToField(_instance, "Vase2DDiagrams", &(_instance.Vase2DDiagrams), formDiv, probe)
		case "IsTubeVase3DDiagramsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsTubeVase3DDiagramsNodeExpanded), formDiv)
		case "TubeVase3DDiagrams":
			FormDivSliceOfPointersToField(_instance, "TubeVase3DDiagrams", &(_instance.TubeVase3DDiagrams), formDiv, probe)
		case "IsStool2DDiagramsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsStool2DDiagramsNodeExpanded), formDiv)
		case "Stool2DDiagrams":
			FormDivSliceOfPointersToField(_instance, "Stool2DDiagrams", &(_instance.Stool2DDiagrams), formDiv, probe)
		case "IsStool3DDiagramsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsStool3DDiagramsNodeExpanded), formDiv)
		case "Stool3DDiagrams":
			FormDivSliceOfPointersToField(_instance, "Stool3DDiagrams", &(_instance.Stool3DDiagrams), formDiv, probe)
		case "IsClock2DDiagramsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsClock2DDiagramsNodeExpanded), formDiv)
		case "Clock2DDiagrams":
			FormDivSliceOfPointersToField(_instance, "Clock2DDiagrams", &(_instance.Clock2DDiagrams), formDiv, probe)
		case "IsClock3DDiagramsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsClock3DDiagramsNodeExpanded), formDiv)
		case "Clock3DDiagrams":
			FormDivSliceOfPointersToField(_instance, "Clock3DDiagrams", &(_instance.Clock3DDiagrams), formDiv, probe)
		case "AxesShape":
			FormDivSelectFieldToField(&(_instance.AxesShape), probe.stageOfInterest, formDiv)
		case "RhombusStuff":
			FormDivSelectFieldToField(&(_instance.RhombusStuff), probe.stageOfInterest, formDiv)
		case "GrowthVectorShape":
			FormDivSelectFieldToField(&(_instance.GrowthVectorShape), probe.stageOfInterest, formDiv)
		case "PerpendicularVectorGrid":
			FormDivSelectFieldToField(&(_instance.PerpendicularVectorGrid), probe.stageOfInterest, formDiv)
		case "BaseVectorShapeGrid":
			FormDivSelectFieldToField(&(_instance.BaseVectorShapeGrid), probe.stageOfInterest, formDiv)
		case "ArcNormalVectorShapeGrid":
			FormDivSelectFieldToField(&(_instance.ArcNormalVectorShapeGrid), probe.stageOfInterest, formDiv)
		case "StartArcShapeGrid":
			FormDivSelectFieldToField(&(_instance.StartArcShapeGrid), probe.stageOfInterest, formDiv)
		case "MidArcVectorShapeGrid":
			FormDivSelectFieldToField(&(_instance.MidArcVectorShapeGrid), probe.stageOfInterest, formDiv)
		case "EndArcShapeGrid":
			FormDivSelectFieldToField(&(_instance.EndArcShapeGrid), probe.stageOfInterest, formDiv)
		case "GrowthCurve2D":
			FormDivSelectFieldToField(&(_instance.GrowthCurve2D), probe.stageOfInterest, formDiv)
		case "StackOfGrowthCurve2DByGrowthVector":
			FormDivSelectFieldToField(&(_instance.StackOfGrowthCurve2DByGrowthVector), probe.stageOfInterest, formDiv)
		case "Library:Plants":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Plants", func(owner *models.Library) *[]*models.PlantAbstract { return &owner.Plants })
		}
	}
}

func __gong__New__PlantCircumferenceShapeFormCallback(
	_instance *models.PlantCircumferenceShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (plantcircumferenceshapeFormCallback *FormCallback[*models.PlantCircumferenceShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePlantCircumferenceShapeFields,
	)
}

type PlantCircumferenceShapeFormCallback = FormCallback[*models.PlantCircumferenceShape]

func savePlantCircumferenceShapeFields(
	_instance *models.PlantCircumferenceShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "AngleDegree":
			FormDivBasicFieldToField(&(_instance.AngleDegree), formDiv)
		case "Length":
			FormDivBasicFieldToField(&(_instance.Length), formDiv)
		}
	}
}

func __gong__New__PointsAndLines3DShapeFormCallback(
	_instance *models.PointsAndLines3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (pointsandlines3dshapeFormCallback *FormCallback[*models.PointsAndLines3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePointsAndLines3DShapeFields,
	)
}

type PointsAndLines3DShapeFormCallback = FormCallback[*models.PointsAndLines3DShape]

func savePointsAndLines3DShapeFields(
	_instance *models.PointsAndLines3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__PxShapeFormCallback(
	_instance *models.PxShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (pxshapeFormCallback *FormCallback[*models.PxShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePxShapeFields,
	)
}

type PxShapeFormCallback = FormCallback[*models.PxShape]

func savePxShapeFields(
	_instance *models.PxShape,
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

func __gong__New__Rendered3DShapeFormCallback(
	_instance *models.Rendered3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (rendered3dshapeFormCallback *FormCallback[*models.Rendered3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRendered3DShapeFields,
	)
}

type Rendered3DShapeFormCallback = FormCallback[*models.Rendered3DShape]

func saveRendered3DShapeFields(
	_instance *models.Rendered3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ViewX":
			FormDivBasicFieldToField(&(_instance.ViewX), formDiv)
		case "ViewY":
			FormDivBasicFieldToField(&(_instance.ViewY), formDiv)
		case "ViewZ":
			FormDivBasicFieldToField(&(_instance.ViewZ), formDiv)
		case "TargetX":
			FormDivBasicFieldToField(&(_instance.TargetX), formDiv)
		case "TargetY":
			FormDivBasicFieldToField(&(_instance.TargetY), formDiv)
		case "TargetZ":
			FormDivBasicFieldToField(&(_instance.TargetZ), formDiv)
		case "Fov":
			FormDivBasicFieldToField(&(_instance.Fov), formDiv)
		}
	}
}

func __gong__New__RhombusShapeFormCallback(
	_instance *models.RhombusShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (rhombusshapeFormCallback *FormCallback[*models.RhombusShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRhombusShapeFields,
	)
}

type RhombusShapeFormCallback = FormCallback[*models.RhombusShape]

func saveRhombusShapeFields(
	_instance *models.RhombusShape,
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

func __gong__New__RhombusStuffFormCallback(
	_instance *models.RhombusStuff,
	probe *Probe,
	formGroup *form.FormGroup,
) (rhombusstuffFormCallback *FormCallback[*models.RhombusStuff]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRhombusStuffFields,
	)
}

type RhombusStuffFormCallback = FormCallback[*models.RhombusStuff]

func saveRhombusStuffFields(
	_instance *models.RhombusStuff,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ReferenceRhombus":
			FormDivSelectFieldToField(&(_instance.ReferenceRhombus), probe.stageOfInterest, formDiv)
		case "PlantCircumferenceShape":
			FormDivSelectFieldToField(&(_instance.PlantCircumferenceShape), probe.stageOfInterest, formDiv)
		case "GridPathShape":
			FormDivSelectFieldToField(&(_instance.GridPathShape), probe.stageOfInterest, formDiv)
		case "InitialRhombusGridShape":
			FormDivSelectFieldToField(&(_instance.InitialRhombusGridShape), probe.stageOfInterest, formDiv)
		case "ExplanationTextShape":
			FormDivSelectFieldToField(&(_instance.ExplanationTextShape), probe.stageOfInterest, formDiv)
		case "RotatedReferenceRhombus":
			FormDivSelectFieldToField(&(_instance.RotatedReferenceRhombus), probe.stageOfInterest, formDiv)
		case "RotatedPlantCircumferenceShape":
			FormDivSelectFieldToField(&(_instance.RotatedPlantCircumferenceShape), probe.stageOfInterest, formDiv)
		case "RotatedGridPathShape":
			FormDivSelectFieldToField(&(_instance.RotatedGridPathShape), probe.stageOfInterest, formDiv)
		case "RotatedRhombusGridShape2":
			FormDivSelectFieldToField(&(_instance.RotatedRhombusGridShape2), probe.stageOfInterest, formDiv)
		case "GrowthCurveRhombusGridShape":
			FormDivSelectFieldToField(&(_instance.GrowthCurveRhombusGridShape), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__RotatedRhombusGridShapeFormCallback(
	_instance *models.RotatedRhombusGridShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (rotatedrhombusgridshapeFormCallback *FormCallback[*models.RotatedRhombusGridShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRotatedRhombusGridShapeFields,
	)
}

type RotatedRhombusGridShapeFormCallback = FormCallback[*models.RotatedRhombusGridShape]

func saveRotatedRhombusGridShapeFields(
	_instance *models.RotatedRhombusGridShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "RotatedRhombusShapes":
			FormDivSliceOfPointersToField(_instance, "RotatedRhombusShapes", &(_instance.RotatedRhombusShapes), formDiv, probe)
		}
	}
}

func __gong__New__RotatedRhombusShapeFormCallback(
	_instance *models.RotatedRhombusShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (rotatedrhombusshapeFormCallback *FormCallback[*models.RotatedRhombusShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRotatedRhombusShapeFields,
	)
}

type RotatedRhombusShapeFormCallback = FormCallback[*models.RotatedRhombusShape]

func saveRotatedRhombusShapeFields(
	_instance *models.RotatedRhombusShape,
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
		case "RotatedRhombusGridShape:RotatedRhombusShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RotatedRhombusShapes", func(owner *models.RotatedRhombusGridShape) *[]*models.RotatedRhombusShape { return &owner.RotatedRhombusShapes })
		}
	}
}

func __gong__New__RotatedSampledPoints3DShapeFormCallback(
	_instance *models.RotatedSampledPoints3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (rotatedsampledpoints3dshapeFormCallback *FormCallback[*models.RotatedSampledPoints3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRotatedSampledPoints3DShapeFields,
	)
}

type RotatedSampledPoints3DShapeFormCallback = FormCallback[*models.RotatedSampledPoints3DShape]

func saveRotatedSampledPoints3DShapeFields(
	_instance *models.RotatedSampledPoints3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__RotatedSeatAndLegs3DShapeFormCallback(
	_instance *models.RotatedSeatAndLegs3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (rotatedseatandlegs3dshapeFormCallback *FormCallback[*models.RotatedSeatAndLegs3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRotatedSeatAndLegs3DShapeFields,
	)
}

type RotatedSeatAndLegs3DShapeFormCallback = FormCallback[*models.RotatedSeatAndLegs3DShape]

func saveRotatedSeatAndLegs3DShapeFields(
	_instance *models.RotatedSeatAndLegs3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__SampledPoints3DShapeFormCallback(
	_instance *models.SampledPoints3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (sampledpoints3dshapeFormCallback *FormCallback[*models.SampledPoints3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSampledPoints3DShapeFields,
	)
}

type SampledPoints3DShapeFormCallback = FormCallback[*models.SampledPoints3DShape]

func saveSampledPoints3DShapeFields(
	_instance *models.SampledPoints3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__Seat3DShapeFormCallback(
	_instance *models.Seat3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (seat3dshapeFormCallback *FormCallback[*models.Seat3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSeat3DShapeFields,
	)
}

type Seat3DShapeFormCallback = FormCallback[*models.Seat3DShape]

func saveSeat3DShapeFields(
	_instance *models.Seat3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__SeatAndLegs3DShapeFormCallback(
	_instance *models.SeatAndLegs3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (seatandlegs3dshapeFormCallback *FormCallback[*models.SeatAndLegs3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSeatAndLegs3DShapeFields,
	)
}

type SeatAndLegs3DShapeFormCallback = FormCallback[*models.SeatAndLegs3DShape]

func saveSeatAndLegs3DShapeFields(
	_instance *models.SeatAndLegs3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__SeatBottomCurveShapeFormCallback(
	_instance *models.SeatBottomCurveShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (seatbottomcurveshapeFormCallback *FormCallback[*models.SeatBottomCurveShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSeatBottomCurveShapeFields,
	)
}

type SeatBottomCurveShapeFormCallback = FormCallback[*models.SeatBottomCurveShape]

func saveSeatBottomCurveShapeFields(
	_instance *models.SeatBottomCurveShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__SeatTopCurveShapeFormCallback(
	_instance *models.SeatTopCurveShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (seattopcurveshapeFormCallback *FormCallback[*models.SeatTopCurveShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSeatTopCurveShapeFields,
	)
}

type SeatTopCurveShapeFormCallback = FormCallback[*models.SeatTopCurveShape]

func saveSeatTopCurveShapeFields(
	_instance *models.SeatTopCurveShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__ShiftedBottomTopStartArcShapeFormCallback(
	_instance *models.ShiftedBottomTopStartArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedbottomtopstartarcshapeFormCallback *FormCallback[*models.ShiftedBottomTopStartArcShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveShiftedBottomTopStartArcShapeFields,
	)
}

type ShiftedBottomTopStartArcShapeFormCallback = FormCallback[*models.ShiftedBottomTopStartArcShape]

func saveShiftedBottomTopStartArcShapeFields(
	_instance *models.ShiftedBottomTopStartArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(_instance.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(_instance.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(_instance.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(_instance.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(_instance.RadiusY), formDiv)
		case "ShiftedBottomTopStartArcShapeGrid:ShiftedBottomTopStartArcShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ShiftedBottomTopStartArcShapes", func(owner *models.ShiftedBottomTopStartArcShapeGrid) *[]*models.ShiftedBottomTopStartArcShape { return &owner.ShiftedBottomTopStartArcShapes })
		}
	}
}

func __gong__New__ShiftedBottomTopStartArcShapeGridFormCallback(
	_instance *models.ShiftedBottomTopStartArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedbottomtopstartarcshapegridFormCallback *FormCallback[*models.ShiftedBottomTopStartArcShapeGrid]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveShiftedBottomTopStartArcShapeGridFields,
	)
}

type ShiftedBottomTopStartArcShapeGridFormCallback = FormCallback[*models.ShiftedBottomTopStartArcShapeGrid]

func saveShiftedBottomTopStartArcShapeGridFields(
	_instance *models.ShiftedBottomTopStartArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ShiftedBottomTopStartArcShapes":
			FormDivSliceOfPointersToField(_instance, "ShiftedBottomTopStartArcShapes", &(_instance.ShiftedBottomTopStartArcShapes), formDiv, probe)
		}
	}
}

func __gong__New__ShiftedLeftGrowthCurve2DRibbonFormCallback(
	_instance *models.ShiftedLeftGrowthCurve2DRibbon,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedleftgrowthcurve2dribbonFormCallback *FormCallback[*models.ShiftedLeftGrowthCurve2DRibbon]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveShiftedLeftGrowthCurve2DRibbonFields,
	)
}

type ShiftedLeftGrowthCurve2DRibbonFormCallback = FormCallback[*models.ShiftedLeftGrowthCurve2DRibbon]

func saveShiftedLeftGrowthCurve2DRibbonFields(
	_instance *models.ShiftedLeftGrowthCurve2DRibbon,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ShiftedLeftGrowthCurve2DRibbonStartShapes":
			FormDivSliceOfPointersToField(_instance, "ShiftedLeftGrowthCurve2DRibbonStartShapes", &(_instance.ShiftedLeftGrowthCurve2DRibbonStartShapes), formDiv, probe)
		case "ShiftedLeftGrowthCurve2DRibbonEndShapes":
			FormDivSliceOfPointersToField(_instance, "ShiftedLeftGrowthCurve2DRibbonEndShapes", &(_instance.ShiftedLeftGrowthCurve2DRibbonEndShapes), formDiv, probe)
		}
	}
}

func __gong__New__ShiftedLeftGrowthCurve2DRibbonEndShapeFormCallback(
	_instance *models.ShiftedLeftGrowthCurve2DRibbonEndShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedleftgrowthcurve2dribbonendshapeFormCallback *FormCallback[*models.ShiftedLeftGrowthCurve2DRibbonEndShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveShiftedLeftGrowthCurve2DRibbonEndShapeFields,
	)
}

type ShiftedLeftGrowthCurve2DRibbonEndShapeFormCallback = FormCallback[*models.ShiftedLeftGrowthCurve2DRibbonEndShape]

func saveShiftedLeftGrowthCurve2DRibbonEndShapeFields(
	_instance *models.ShiftedLeftGrowthCurve2DRibbonEndShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(_instance.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(_instance.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(_instance.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(_instance.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(_instance.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(_instance.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(_instance.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(_instance.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(_instance.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(_instance.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(_instance.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(_instance.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(_instance.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(_instance.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(_instance.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(_instance.TopSweepFlag), formDiv)
		case "ShiftedLeftGrowthCurve2DRibbon:ShiftedLeftGrowthCurve2DRibbonEndShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ShiftedLeftGrowthCurve2DRibbonEndShapes", func(owner *models.ShiftedLeftGrowthCurve2DRibbon) *[]*models.ShiftedLeftGrowthCurve2DRibbonEndShape { return &owner.ShiftedLeftGrowthCurve2DRibbonEndShapes })
		}
	}
}

func __gong__New__ShiftedLeftGrowthCurve2DRibbonStartShapeFormCallback(
	_instance *models.ShiftedLeftGrowthCurve2DRibbonStartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedleftgrowthcurve2dribbonstartshapeFormCallback *FormCallback[*models.ShiftedLeftGrowthCurve2DRibbonStartShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveShiftedLeftGrowthCurve2DRibbonStartShapeFields,
	)
}

type ShiftedLeftGrowthCurve2DRibbonStartShapeFormCallback = FormCallback[*models.ShiftedLeftGrowthCurve2DRibbonStartShape]

func saveShiftedLeftGrowthCurve2DRibbonStartShapeFields(
	_instance *models.ShiftedLeftGrowthCurve2DRibbonStartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(_instance.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(_instance.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(_instance.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(_instance.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(_instance.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(_instance.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(_instance.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(_instance.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(_instance.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(_instance.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(_instance.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(_instance.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(_instance.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(_instance.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(_instance.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(_instance.TopSweepFlag), formDiv)
		case "ShiftedLeftGrowthCurve2DRibbon:ShiftedLeftGrowthCurve2DRibbonStartShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ShiftedLeftGrowthCurve2DRibbonStartShapes", func(owner *models.ShiftedLeftGrowthCurve2DRibbon) *[]*models.ShiftedLeftGrowthCurve2DRibbonStartShape { return &owner.ShiftedLeftGrowthCurve2DRibbonStartShapes })
		}
	}
}

func __gong__New__ShiftedLeftPartiallyGrowthCurve2DRibbonFormCallback(
	_instance *models.ShiftedLeftPartiallyGrowthCurve2DRibbon,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedleftpartiallygrowthcurve2dribbonFormCallback *FormCallback[*models.ShiftedLeftPartiallyGrowthCurve2DRibbon]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveShiftedLeftPartiallyGrowthCurve2DRibbonFields,
	)
}

type ShiftedLeftPartiallyGrowthCurve2DRibbonFormCallback = FormCallback[*models.ShiftedLeftPartiallyGrowthCurve2DRibbon]

func saveShiftedLeftPartiallyGrowthCurve2DRibbonFields(
	_instance *models.ShiftedLeftPartiallyGrowthCurve2DRibbon,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes":
			FormDivSliceOfPointersToField(_instance, "ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes", &(_instance.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes), formDiv, probe)
		case "ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes":
			FormDivSliceOfPointersToField(_instance, "ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes", &(_instance.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes), formDiv, probe)
		}
	}
}

func __gong__New__ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapeFormCallback(
	_instance *models.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedleftpartiallygrowthcurve2dribbonendshapeFormCallback *FormCallback[*models.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveShiftedLeftPartiallyGrowthCurve2DRibbonEndShapeFields,
	)
}

type ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapeFormCallback = FormCallback[*models.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape]

func saveShiftedLeftPartiallyGrowthCurve2DRibbonEndShapeFields(
	_instance *models.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(_instance.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(_instance.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(_instance.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(_instance.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(_instance.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(_instance.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(_instance.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(_instance.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(_instance.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(_instance.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(_instance.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(_instance.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(_instance.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(_instance.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(_instance.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(_instance.TopSweepFlag), formDiv)
		case "ShiftedLeftPartiallyGrowthCurve2DRibbon:ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes", func(owner *models.ShiftedLeftPartiallyGrowthCurve2DRibbon) *[]*models.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape { return &owner.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes })
		}
	}
}

func __gong__New__ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapeFormCallback(
	_instance *models.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedleftpartiallygrowthcurve2dribbonstartshapeFormCallback *FormCallback[*models.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveShiftedLeftPartiallyGrowthCurve2DRibbonStartShapeFields,
	)
}

type ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapeFormCallback = FormCallback[*models.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape]

func saveShiftedLeftPartiallyGrowthCurve2DRibbonStartShapeFields(
	_instance *models.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(_instance.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(_instance.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(_instance.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(_instance.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(_instance.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(_instance.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(_instance.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(_instance.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(_instance.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(_instance.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(_instance.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(_instance.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(_instance.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(_instance.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(_instance.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(_instance.TopSweepFlag), formDiv)
		case "ShiftedLeftPartiallyGrowthCurve2DRibbon:ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes", func(owner *models.ShiftedLeftPartiallyGrowthCurve2DRibbon) *[]*models.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape { return &owner.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes })
		}
	}
}

func __gong__New__ShiftedLeftStackGrowthCurveEndArcShapeFormCallback(
	_instance *models.ShiftedLeftStackGrowthCurveEndArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedleftstackgrowthcurveendarcshapeFormCallback *FormCallback[*models.ShiftedLeftStackGrowthCurveEndArcShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveShiftedLeftStackGrowthCurveEndArcShapeFields,
	)
}

type ShiftedLeftStackGrowthCurveEndArcShapeFormCallback = FormCallback[*models.ShiftedLeftStackGrowthCurveEndArcShape]

func saveShiftedLeftStackGrowthCurveEndArcShapeFields(
	_instance *models.ShiftedLeftStackGrowthCurveEndArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(_instance.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(_instance.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(_instance.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(_instance.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(_instance.RadiusY), formDiv)
		case "ShiftedLeftStackOfGrowthCurve:ShiftedLeftStackGrowthCurveEndArcShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ShiftedLeftStackGrowthCurveEndArcShapes", func(owner *models.ShiftedLeftStackOfGrowthCurve) *[]*models.ShiftedLeftStackGrowthCurveEndArcShape { return &owner.ShiftedLeftStackGrowthCurveEndArcShapes })
		}
	}
}

func __gong__New__ShiftedLeftStackGrowthCurveStartArcShapeFormCallback(
	_instance *models.ShiftedLeftStackGrowthCurveStartArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedleftstackgrowthcurvestartarcshapeFormCallback *FormCallback[*models.ShiftedLeftStackGrowthCurveStartArcShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveShiftedLeftStackGrowthCurveStartArcShapeFields,
	)
}

type ShiftedLeftStackGrowthCurveStartArcShapeFormCallback = FormCallback[*models.ShiftedLeftStackGrowthCurveStartArcShape]

func saveShiftedLeftStackGrowthCurveStartArcShapeFields(
	_instance *models.ShiftedLeftStackGrowthCurveStartArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(_instance.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(_instance.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(_instance.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(_instance.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(_instance.RadiusY), formDiv)
		case "ShiftedLeftStackOfGrowthCurve:ShiftedLeftStackGrowthCurveStartArcShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ShiftedLeftStackGrowthCurveStartArcShapes", func(owner *models.ShiftedLeftStackOfGrowthCurve) *[]*models.ShiftedLeftStackGrowthCurveStartArcShape { return &owner.ShiftedLeftStackGrowthCurveStartArcShapes })
		}
	}
}

func __gong__New__ShiftedLeftStackNormalVectorFormCallback(
	_instance *models.ShiftedLeftStackNormalVector,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedleftstacknormalvectorFormCallback *FormCallback[*models.ShiftedLeftStackNormalVector]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveShiftedLeftStackNormalVectorFields,
	)
}

type ShiftedLeftStackNormalVectorFormCallback = FormCallback[*models.ShiftedLeftStackNormalVector]

func saveShiftedLeftStackNormalVectorFields(
	_instance *models.ShiftedLeftStackNormalVector,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "ShiftedLeftStackOfNormalVector:ShiftedLeftStackNormalVectors":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ShiftedLeftStackNormalVectors", func(owner *models.ShiftedLeftStackOfNormalVector) *[]*models.ShiftedLeftStackNormalVector { return &owner.ShiftedLeftStackNormalVectors })
		}
	}
}

func __gong__New__ShiftedLeftStackOfGrowthCurveFormCallback(
	_instance *models.ShiftedLeftStackOfGrowthCurve,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedleftstackofgrowthcurveFormCallback *FormCallback[*models.ShiftedLeftStackOfGrowthCurve]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveShiftedLeftStackOfGrowthCurveFields,
	)
}

type ShiftedLeftStackOfGrowthCurveFormCallback = FormCallback[*models.ShiftedLeftStackOfGrowthCurve]

func saveShiftedLeftStackOfGrowthCurveFields(
	_instance *models.ShiftedLeftStackOfGrowthCurve,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ShiftedLeftStackGrowthCurveStartArcShapes":
			FormDivSliceOfPointersToField(_instance, "ShiftedLeftStackGrowthCurveStartArcShapes", &(_instance.ShiftedLeftStackGrowthCurveStartArcShapes), formDiv, probe)
		case "ShiftedLeftStackGrowthCurveEndArcShapes":
			FormDivSliceOfPointersToField(_instance, "ShiftedLeftStackGrowthCurveEndArcShapes", &(_instance.ShiftedLeftStackGrowthCurveEndArcShapes), formDiv, probe)
		}
	}
}

func __gong__New__ShiftedLeftStackOfNormalVectorFormCallback(
	_instance *models.ShiftedLeftStackOfNormalVector,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedleftstackofnormalvectorFormCallback *FormCallback[*models.ShiftedLeftStackOfNormalVector]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveShiftedLeftStackOfNormalVectorFields,
	)
}

type ShiftedLeftStackOfNormalVectorFormCallback = FormCallback[*models.ShiftedLeftStackOfNormalVector]

func saveShiftedLeftStackOfNormalVectorFields(
	_instance *models.ShiftedLeftStackOfNormalVector,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ShiftedLeftStackNormalVectors":
			FormDivSliceOfPointersToField(_instance, "ShiftedLeftStackNormalVectors", &(_instance.ShiftedLeftStackNormalVectors), formDiv, probe)
		}
	}
}

func __gong__New__ShiftedRightGrowthCurve2DRibbonFormCallback(
	_instance *models.ShiftedRightGrowthCurve2DRibbon,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedrightgrowthcurve2dribbonFormCallback *FormCallback[*models.ShiftedRightGrowthCurve2DRibbon]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveShiftedRightGrowthCurve2DRibbonFields,
	)
}

type ShiftedRightGrowthCurve2DRibbonFormCallback = FormCallback[*models.ShiftedRightGrowthCurve2DRibbon]

func saveShiftedRightGrowthCurve2DRibbonFields(
	_instance *models.ShiftedRightGrowthCurve2DRibbon,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ShiftedRightGrowthCurve2DRibbonStartShapes":
			FormDivSliceOfPointersToField(_instance, "ShiftedRightGrowthCurve2DRibbonStartShapes", &(_instance.ShiftedRightGrowthCurve2DRibbonStartShapes), formDiv, probe)
		case "ShiftedRightGrowthCurve2DRibbonEndShapes":
			FormDivSliceOfPointersToField(_instance, "ShiftedRightGrowthCurve2DRibbonEndShapes", &(_instance.ShiftedRightGrowthCurve2DRibbonEndShapes), formDiv, probe)
		}
	}
}

func __gong__New__ShiftedRightGrowthCurve2DRibbonEndShapeFormCallback(
	_instance *models.ShiftedRightGrowthCurve2DRibbonEndShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedrightgrowthcurve2dribbonendshapeFormCallback *FormCallback[*models.ShiftedRightGrowthCurve2DRibbonEndShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveShiftedRightGrowthCurve2DRibbonEndShapeFields,
	)
}

type ShiftedRightGrowthCurve2DRibbonEndShapeFormCallback = FormCallback[*models.ShiftedRightGrowthCurve2DRibbonEndShape]

func saveShiftedRightGrowthCurve2DRibbonEndShapeFields(
	_instance *models.ShiftedRightGrowthCurve2DRibbonEndShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(_instance.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(_instance.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(_instance.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(_instance.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(_instance.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(_instance.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(_instance.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(_instance.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(_instance.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(_instance.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(_instance.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(_instance.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(_instance.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(_instance.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(_instance.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(_instance.TopSweepFlag), formDiv)
		case "ShiftedRightGrowthCurve2DRibbon:ShiftedRightGrowthCurve2DRibbonEndShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ShiftedRightGrowthCurve2DRibbonEndShapes", func(owner *models.ShiftedRightGrowthCurve2DRibbon) *[]*models.ShiftedRightGrowthCurve2DRibbonEndShape { return &owner.ShiftedRightGrowthCurve2DRibbonEndShapes })
		}
	}
}

func __gong__New__ShiftedRightGrowthCurve2DRibbonStartShapeFormCallback(
	_instance *models.ShiftedRightGrowthCurve2DRibbonStartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (shiftedrightgrowthcurve2dribbonstartshapeFormCallback *FormCallback[*models.ShiftedRightGrowthCurve2DRibbonStartShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveShiftedRightGrowthCurve2DRibbonStartShapeFields,
	)
}

type ShiftedRightGrowthCurve2DRibbonStartShapeFormCallback = FormCallback[*models.ShiftedRightGrowthCurve2DRibbonStartShape]

func saveShiftedRightGrowthCurve2DRibbonStartShapeFields(
	_instance *models.ShiftedRightGrowthCurve2DRibbonStartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(_instance.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(_instance.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(_instance.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(_instance.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(_instance.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(_instance.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(_instance.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(_instance.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(_instance.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(_instance.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(_instance.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(_instance.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(_instance.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(_instance.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(_instance.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(_instance.TopSweepFlag), formDiv)
		case "ShiftedRightGrowthCurve2DRibbon:ShiftedRightGrowthCurve2DRibbonStartShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ShiftedRightGrowthCurve2DRibbonStartShapes", func(owner *models.ShiftedRightGrowthCurve2DRibbon) *[]*models.ShiftedRightGrowthCurve2DRibbonStartShape { return &owner.ShiftedRightGrowthCurve2DRibbonStartShapes })
		}
	}
}

func __gong__New__StackGrowthCurve2DEndHalfwayArcShapeFormCallback(
	_instance *models.StackGrowthCurve2DEndHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackgrowthcurve2dendhalfwayarcshapeFormCallback *FormCallback[*models.StackGrowthCurve2DEndHalfwayArcShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStackGrowthCurve2DEndHalfwayArcShapeFields,
	)
}

type StackGrowthCurve2DEndHalfwayArcShapeFormCallback = FormCallback[*models.StackGrowthCurve2DEndHalfwayArcShape]

func saveStackGrowthCurve2DEndHalfwayArcShapeFields(
	_instance *models.StackGrowthCurve2DEndHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(_instance.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(_instance.RadiusY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(_instance.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(_instance.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(_instance.SweepFlag), formDiv)
		case "StackOfGrowthCurve2D:StackGrowthCurve2DEndHalfwayArcShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "StackGrowthCurve2DEndHalfwayArcShapes", func(owner *models.StackOfGrowthCurve2D) *[]*models.StackGrowthCurve2DEndHalfwayArcShape { return &owner.StackGrowthCurve2DEndHalfwayArcShapes })
		}
	}
}

func __gong__New__StackGrowthCurve2DRibbonEndShapeFormCallback(
	_instance *models.StackGrowthCurve2DRibbonEndShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackgrowthcurve2dribbonendshapeFormCallback *FormCallback[*models.StackGrowthCurve2DRibbonEndShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStackGrowthCurve2DRibbonEndShapeFields,
	)
}

type StackGrowthCurve2DRibbonEndShapeFormCallback = FormCallback[*models.StackGrowthCurve2DRibbonEndShape]

func saveStackGrowthCurve2DRibbonEndShapeFields(
	_instance *models.StackGrowthCurve2DRibbonEndShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(_instance.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(_instance.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(_instance.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(_instance.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(_instance.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(_instance.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(_instance.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(_instance.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(_instance.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(_instance.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(_instance.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(_instance.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(_instance.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(_instance.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(_instance.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(_instance.TopSweepFlag), formDiv)
		case "StackOfGrowthCurve2DRibbon:StackGrowthCurve2DRibbonEndShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "StackGrowthCurve2DRibbonEndShapes", func(owner *models.StackOfGrowthCurve2DRibbon) *[]*models.StackGrowthCurve2DRibbonEndShape { return &owner.StackGrowthCurve2DRibbonEndShapes })
		}
	}
}

func __gong__New__StackGrowthCurve2DRibbonStartShapeFormCallback(
	_instance *models.StackGrowthCurve2DRibbonStartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackgrowthcurve2dribbonstartshapeFormCallback *FormCallback[*models.StackGrowthCurve2DRibbonStartShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStackGrowthCurve2DRibbonStartShapeFields,
	)
}

type StackGrowthCurve2DRibbonStartShapeFormCallback = FormCallback[*models.StackGrowthCurve2DRibbonStartShape]

func saveStackGrowthCurve2DRibbonStartShapeFields(
	_instance *models.StackGrowthCurve2DRibbonStartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(_instance.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(_instance.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(_instance.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(_instance.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(_instance.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(_instance.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(_instance.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(_instance.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(_instance.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(_instance.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(_instance.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(_instance.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(_instance.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(_instance.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(_instance.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(_instance.TopSweepFlag), formDiv)
		case "StackOfGrowthCurve2DRibbon:StackGrowthCurve2DRibbonStartShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "StackGrowthCurve2DRibbonStartShapes", func(owner *models.StackOfGrowthCurve2DRibbon) *[]*models.StackGrowthCurve2DRibbonStartShape { return &owner.StackGrowthCurve2DRibbonStartShapes })
		}
	}
}

func __gong__New__StackGrowthCurve2DStartHalfwayArcShapeFormCallback(
	_instance *models.StackGrowthCurve2DStartHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackgrowthcurve2dstarthalfwayarcshapeFormCallback *FormCallback[*models.StackGrowthCurve2DStartHalfwayArcShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStackGrowthCurve2DStartHalfwayArcShapeFields,
	)
}

type StackGrowthCurve2DStartHalfwayArcShapeFormCallback = FormCallback[*models.StackGrowthCurve2DStartHalfwayArcShape]

func saveStackGrowthCurve2DStartHalfwayArcShapeFields(
	_instance *models.StackGrowthCurve2DStartHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(_instance.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(_instance.RadiusY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(_instance.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(_instance.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(_instance.SweepFlag), formDiv)
		case "StackOfGrowthCurve2D:StackGrowthCurve2DStartHalfwayArcShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "StackGrowthCurve2DStartHalfwayArcShapes", func(owner *models.StackOfGrowthCurve2D) *[]*models.StackGrowthCurve2DStartHalfwayArcShape { return &owner.StackGrowthCurve2DStartHalfwayArcShapes })
		}
	}
}

func __gong__New__StackOfGrowthCurve2DFormCallback(
	_instance *models.StackOfGrowthCurve2D,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackofgrowthcurve2dFormCallback *FormCallback[*models.StackOfGrowthCurve2D]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStackOfGrowthCurve2DFields,
	)
}

type StackOfGrowthCurve2DFormCallback = FormCallback[*models.StackOfGrowthCurve2D]

func saveStackOfGrowthCurve2DFields(
	_instance *models.StackOfGrowthCurve2D,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StackGrowthCurve2DStartHalfwayArcShapes":
			FormDivSliceOfPointersToField(_instance, "StackGrowthCurve2DStartHalfwayArcShapes", &(_instance.StackGrowthCurve2DStartHalfwayArcShapes), formDiv, probe)
		case "StackGrowthCurve2DEndHalfwayArcShapes":
			FormDivSliceOfPointersToField(_instance, "StackGrowthCurve2DEndHalfwayArcShapes", &(_instance.StackGrowthCurve2DEndHalfwayArcShapes), formDiv, probe)
		}
	}
}

func __gong__New__StackOfGrowthCurve2DByGrowthVectorFormCallback(
	_instance *models.StackOfGrowthCurve2DByGrowthVector,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackofgrowthcurve2dbygrowthvectorFormCallback *FormCallback[*models.StackOfGrowthCurve2DByGrowthVector]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStackOfGrowthCurve2DByGrowthVectorFields,
	)
}

type StackOfGrowthCurve2DByGrowthVectorFormCallback = FormCallback[*models.StackOfGrowthCurve2DByGrowthVector]

func saveStackOfGrowthCurve2DByGrowthVectorFields(
	_instance *models.StackOfGrowthCurve2DByGrowthVector,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__StackOfGrowthCurve2DRibbonFormCallback(
	_instance *models.StackOfGrowthCurve2DRibbon,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackofgrowthcurve2dribbonFormCallback *FormCallback[*models.StackOfGrowthCurve2DRibbon]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStackOfGrowthCurve2DRibbonFields,
	)
}

type StackOfGrowthCurve2DRibbonFormCallback = FormCallback[*models.StackOfGrowthCurve2DRibbon]

func saveStackOfGrowthCurve2DRibbonFields(
	_instance *models.StackOfGrowthCurve2DRibbon,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StackGrowthCurve2DRibbonStartShapes":
			FormDivSliceOfPointersToField(_instance, "StackGrowthCurve2DRibbonStartShapes", &(_instance.StackGrowthCurve2DRibbonStartShapes), formDiv, probe)
		case "StackGrowthCurve2DRibbonEndShapes":
			FormDivSliceOfPointersToField(_instance, "StackGrowthCurve2DRibbonEndShapes", &(_instance.StackGrowthCurve2DRibbonEndShapes), formDiv, probe)
		}
	}
}

func __gong__New__StackOfPartiallyRotatedTorusShapeFormCallback(
	_instance *models.StackOfPartiallyRotatedTorusShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackofpartiallyrotatedtorusshapeFormCallback *FormCallback[*models.StackOfPartiallyRotatedTorusShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStackOfPartiallyRotatedTorusShapeFields,
	)
}

type StackOfPartiallyRotatedTorusShapeFormCallback = FormCallback[*models.StackOfPartiallyRotatedTorusShape]

func saveStackOfPartiallyRotatedTorusShapeFields(
	_instance *models.StackOfPartiallyRotatedTorusShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__StackOfRotatedGrowthCurve2DFormCallback(
	_instance *models.StackOfRotatedGrowthCurve2D,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackofrotatedgrowthcurve2dFormCallback *FormCallback[*models.StackOfRotatedGrowthCurve2D]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStackOfRotatedGrowthCurve2DFields,
	)
}

type StackOfRotatedGrowthCurve2DFormCallback = FormCallback[*models.StackOfRotatedGrowthCurve2D]

func saveStackOfRotatedGrowthCurve2DFields(
	_instance *models.StackOfRotatedGrowthCurve2D,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StackRotatedGrowthCurve2DStartArcShapes":
			FormDivSliceOfPointersToField(_instance, "StackRotatedGrowthCurve2DStartArcShapes", &(_instance.StackRotatedGrowthCurve2DStartArcShapes), formDiv, probe)
		case "StackRotatedGrowthCurve2DEndArcShapes":
			FormDivSliceOfPointersToField(_instance, "StackRotatedGrowthCurve2DEndArcShapes", &(_instance.StackRotatedGrowthCurve2DEndArcShapes), formDiv, probe)
		}
	}
}

func __gong__New__StackOfRotatedGrowthCurve2DRibbonFormCallback(
	_instance *models.StackOfRotatedGrowthCurve2DRibbon,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackofrotatedgrowthcurve2dribbonFormCallback *FormCallback[*models.StackOfRotatedGrowthCurve2DRibbon]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStackOfRotatedGrowthCurve2DRibbonFields,
	)
}

type StackOfRotatedGrowthCurve2DRibbonFormCallback = FormCallback[*models.StackOfRotatedGrowthCurve2DRibbon]

func saveStackOfRotatedGrowthCurve2DRibbonFields(
	_instance *models.StackOfRotatedGrowthCurve2DRibbon,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StackRotatedGrowthCurve2DRibbonStartShapes":
			FormDivSliceOfPointersToField(_instance, "StackRotatedGrowthCurve2DRibbonStartShapes", &(_instance.StackRotatedGrowthCurve2DRibbonStartShapes), formDiv, probe)
		case "StackRotatedGrowthCurve2DRibbonEndShapes":
			FormDivSliceOfPointersToField(_instance, "StackRotatedGrowthCurve2DRibbonEndShapes", &(_instance.StackRotatedGrowthCurve2DRibbonEndShapes), formDiv, probe)
		}
	}
}

func __gong__New__StackRotatedGrowthCurve2DEndArcShapeFormCallback(
	_instance *models.StackRotatedGrowthCurve2DEndArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackrotatedgrowthcurve2dendarcshapeFormCallback *FormCallback[*models.StackRotatedGrowthCurve2DEndArcShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStackRotatedGrowthCurve2DEndArcShapeFields,
	)
}

type StackRotatedGrowthCurve2DEndArcShapeFormCallback = FormCallback[*models.StackRotatedGrowthCurve2DEndArcShape]

func saveStackRotatedGrowthCurve2DEndArcShapeFields(
	_instance *models.StackRotatedGrowthCurve2DEndArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(_instance.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(_instance.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(_instance.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(_instance.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(_instance.RadiusY), formDiv)
		case "StackOfRotatedGrowthCurve2D:StackRotatedGrowthCurve2DEndArcShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "StackRotatedGrowthCurve2DEndArcShapes", func(owner *models.StackOfRotatedGrowthCurve2D) *[]*models.StackRotatedGrowthCurve2DEndArcShape { return &owner.StackRotatedGrowthCurve2DEndArcShapes })
		}
	}
}

func __gong__New__StackRotatedGrowthCurve2DRibbonEndShapeFormCallback(
	_instance *models.StackRotatedGrowthCurve2DRibbonEndShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackrotatedgrowthcurve2dribbonendshapeFormCallback *FormCallback[*models.StackRotatedGrowthCurve2DRibbonEndShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStackRotatedGrowthCurve2DRibbonEndShapeFields,
	)
}

type StackRotatedGrowthCurve2DRibbonEndShapeFormCallback = FormCallback[*models.StackRotatedGrowthCurve2DRibbonEndShape]

func saveStackRotatedGrowthCurve2DRibbonEndShapeFields(
	_instance *models.StackRotatedGrowthCurve2DRibbonEndShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(_instance.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(_instance.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(_instance.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(_instance.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(_instance.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(_instance.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(_instance.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(_instance.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(_instance.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(_instance.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(_instance.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(_instance.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(_instance.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(_instance.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(_instance.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(_instance.TopSweepFlag), formDiv)
		case "StackOfRotatedGrowthCurve2DRibbon:StackRotatedGrowthCurve2DRibbonEndShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "StackRotatedGrowthCurve2DRibbonEndShapes", func(owner *models.StackOfRotatedGrowthCurve2DRibbon) *[]*models.StackRotatedGrowthCurve2DRibbonEndShape { return &owner.StackRotatedGrowthCurve2DRibbonEndShapes })
		}
	}
}

func __gong__New__StackRotatedGrowthCurve2DRibbonStartShapeFormCallback(
	_instance *models.StackRotatedGrowthCurve2DRibbonStartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackrotatedgrowthcurve2dribbonstartshapeFormCallback *FormCallback[*models.StackRotatedGrowthCurve2DRibbonStartShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStackRotatedGrowthCurve2DRibbonStartShapeFields,
	)
}

type StackRotatedGrowthCurve2DRibbonStartShapeFormCallback = FormCallback[*models.StackRotatedGrowthCurve2DRibbonStartShape]

func saveStackRotatedGrowthCurve2DRibbonStartShapeFields(
	_instance *models.StackRotatedGrowthCurve2DRibbonStartShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "BottomStartX":
			FormDivBasicFieldToField(&(_instance.BottomStartX), formDiv)
		case "BottomStartY":
			FormDivBasicFieldToField(&(_instance.BottomStartY), formDiv)
		case "BottomEndX":
			FormDivBasicFieldToField(&(_instance.BottomEndX), formDiv)
		case "BottomEndY":
			FormDivBasicFieldToField(&(_instance.BottomEndY), formDiv)
		case "BottomRadiusX":
			FormDivBasicFieldToField(&(_instance.BottomRadiusX), formDiv)
		case "BottomRadiusY":
			FormDivBasicFieldToField(&(_instance.BottomRadiusY), formDiv)
		case "BottomXAxisRotation":
			FormDivBasicFieldToField(&(_instance.BottomXAxisRotation), formDiv)
		case "BottomLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.BottomLargeArcFlag), formDiv)
		case "BottomSweepFlag":
			FormDivBasicFieldToField(&(_instance.BottomSweepFlag), formDiv)
		case "TopStartX":
			FormDivBasicFieldToField(&(_instance.TopStartX), formDiv)
		case "TopStartY":
			FormDivBasicFieldToField(&(_instance.TopStartY), formDiv)
		case "TopEndX":
			FormDivBasicFieldToField(&(_instance.TopEndX), formDiv)
		case "TopEndY":
			FormDivBasicFieldToField(&(_instance.TopEndY), formDiv)
		case "TopRadiusX":
			FormDivBasicFieldToField(&(_instance.TopRadiusX), formDiv)
		case "TopRadiusY":
			FormDivBasicFieldToField(&(_instance.TopRadiusY), formDiv)
		case "TopXAxisRotation":
			FormDivBasicFieldToField(&(_instance.TopXAxisRotation), formDiv)
		case "TopLargeArcFlag":
			FormDivBasicFieldToField(&(_instance.TopLargeArcFlag), formDiv)
		case "TopSweepFlag":
			FormDivBasicFieldToField(&(_instance.TopSweepFlag), formDiv)
		case "StackOfRotatedGrowthCurve2DRibbon:StackRotatedGrowthCurve2DRibbonStartShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "StackRotatedGrowthCurve2DRibbonStartShapes", func(owner *models.StackOfRotatedGrowthCurve2DRibbon) *[]*models.StackRotatedGrowthCurve2DRibbonStartShape { return &owner.StackRotatedGrowthCurve2DRibbonStartShapes })
		}
	}
}

func __gong__New__StackRotatedGrowthCurve2DStartArcShapeFormCallback(
	_instance *models.StackRotatedGrowthCurve2DStartArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stackrotatedgrowthcurve2dstartarcshapeFormCallback *FormCallback[*models.StackRotatedGrowthCurve2DStartArcShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStackRotatedGrowthCurve2DStartArcShapeFields,
	)
}

type StackRotatedGrowthCurve2DStartArcShapeFormCallback = FormCallback[*models.StackRotatedGrowthCurve2DStartArcShape]

func saveStackRotatedGrowthCurve2DStartArcShapeFields(
	_instance *models.StackRotatedGrowthCurve2DStartArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(_instance.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(_instance.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(_instance.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(_instance.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(_instance.RadiusY), formDiv)
		case "StackOfRotatedGrowthCurve2D:StackRotatedGrowthCurve2DStartArcShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "StackRotatedGrowthCurve2DStartArcShapes", func(owner *models.StackOfRotatedGrowthCurve2D) *[]*models.StackRotatedGrowthCurve2DStartArcShape { return &owner.StackRotatedGrowthCurve2DStartArcShapes })
		}
	}
}

func __gong__New__StartArcShapeFormCallback(
	_instance *models.StartArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (startarcshapeFormCallback *FormCallback[*models.StartArcShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStartArcShapeFields,
	)
}

type StartArcShapeFormCallback = FormCallback[*models.StartArcShape]

func saveStartArcShapeFields(
	_instance *models.StartArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(_instance.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(_instance.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(_instance.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(_instance.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(_instance.RadiusY), formDiv)
		case "StartArcShapeGrid:StartArcShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "StartArcShapes", func(owner *models.StartArcShapeGrid) *[]*models.StartArcShape { return &owner.StartArcShapes })
		}
	}
}

func __gong__New__StartArcShapeGridFormCallback(
	_instance *models.StartArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (startarcshapegridFormCallback *FormCallback[*models.StartArcShapeGrid]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStartArcShapeGridFields,
	)
}

type StartArcShapeGridFormCallback = FormCallback[*models.StartArcShapeGrid]

func saveStartArcShapeGridFields(
	_instance *models.StartArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartArcShapes":
			FormDivSliceOfPointersToField(_instance, "StartArcShapes", &(_instance.StartArcShapes), formDiv, probe)
		}
	}
}

func __gong__New__StartHalfwayArcShapeFormCallback(
	_instance *models.StartHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (starthalfwayarcshapeFormCallback *FormCallback[*models.StartHalfwayArcShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStartHalfwayArcShapeFields,
	)
}

type StartHalfwayArcShapeFormCallback = FormCallback[*models.StartHalfwayArcShape]

func saveStartHalfwayArcShapeFields(
	_instance *models.StartHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(_instance.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(_instance.RadiusY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(_instance.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(_instance.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(_instance.SweepFlag), formDiv)
		case "StartHalfwayArcShapeGrid:StartHalfwayArcShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "StartHalfwayArcShapes", func(owner *models.StartHalfwayArcShapeGrid) *[]*models.StartHalfwayArcShape { return &owner.StartHalfwayArcShapes })
		}
	}
}

func __gong__New__StartHalfwayArcShapeGridFormCallback(
	_instance *models.StartHalfwayArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (starthalfwayarcshapegridFormCallback *FormCallback[*models.StartHalfwayArcShapeGrid]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStartHalfwayArcShapeGridFields,
	)
}

type StartHalfwayArcShapeGridFormCallback = FormCallback[*models.StartHalfwayArcShapeGrid]

func saveStartHalfwayArcShapeGridFields(
	_instance *models.StartHalfwayArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartHalfwayArcShapes":
			FormDivSliceOfPointersToField(_instance, "StartHalfwayArcShapes", &(_instance.StartHalfwayArcShapes), formDiv, probe)
		}
	}
}

func __gong__New__StemCylinder3DShapeFormCallback(
	_instance *models.StemCylinder3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (stemcylinder3dshapeFormCallback *FormCallback[*models.StemCylinder3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStemCylinder3DShapeFields,
	)
}

type StemCylinder3DShapeFormCallback = FormCallback[*models.StemCylinder3DShape]

func saveStemCylinder3DShapeFields(
	_instance *models.StemCylinder3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Transparency":
			FormDivBasicFieldToField(&(_instance.Transparency), formDiv)
		}
	}
}

func __gong__New__Stool2DDiagramFormCallback(
	_instance *models.Stool2DDiagram,
	probe *Probe,
	formGroup *form.FormGroup,
) (stool2ddiagramFormCallback *FormCallback[*models.Stool2DDiagram]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStool2DDiagramFields,
	)
}

type Stool2DDiagramFormCallback = FormCallback[*models.Stool2DDiagram]

func saveStool2DDiagramFields(
	_instance *models.Stool2DDiagram,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Zoom":
			FormDivBasicFieldToField(&(_instance.Zoom), formDiv)
		case "IsHiddenAxesShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenAxesShape), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(_instance.IsChecked), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "PlantAbstract:Stool2DDiagrams":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Stool2DDiagrams", func(owner *models.PlantAbstract) *[]*models.Stool2DDiagram { return &owner.Stool2DDiagrams })
		}
	}
}

func __gong__New__Stool3DDiagramFormCallback(
	_instance *models.Stool3DDiagram,
	probe *Probe,
	formGroup *form.FormGroup,
) (stool3ddiagramFormCallback *FormCallback[*models.Stool3DDiagram]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStool3DDiagramFields,
	)
}

type Stool3DDiagramFormCallback = FormCallback[*models.Stool3DDiagram]

func saveStool3DDiagramFields(
	_instance *models.Stool3DDiagram,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "IsHiddenSeatTopCurveShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenSeatTopCurveShape), formDiv)
		case "SeatTopCurveShape":
			FormDivSelectFieldToField(&(_instance.SeatTopCurveShape), probe.stageOfInterest, formDiv)
		case "IsHiddenRotatedSeatTopCurveShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenRotatedSeatTopCurveShape), formDiv)
		case "RotatedSeatTopCurveShape":
			FormDivSelectFieldToField(&(_instance.RotatedSeatTopCurveShape), probe.stageOfInterest, formDiv)
		case "IsHiddenSeatBottomCurveShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenSeatBottomCurveShape), formDiv)
		case "SeatBottomCurveShape":
			FormDivSelectFieldToField(&(_instance.SeatBottomCurveShape), probe.stageOfInterest, formDiv)
		case "IsHiddenRotatedSeatBottomCurveShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenRotatedSeatBottomCurveShape), formDiv)
		case "RotatedSeatBottomCurveShape":
			FormDivSelectFieldToField(&(_instance.RotatedSeatBottomCurveShape), probe.stageOfInterest, formDiv)
		case "IsHiddenTorus3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenTorus3DShape), formDiv)
		case "Torus3DShape":
			FormDivSelectFieldToField(&(_instance.Torus3DShape), probe.stageOfInterest, formDiv)
		case "IsHiddenRotatedTorusShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenRotatedTorusShape), formDiv)
		case "RotatedTorusShape":
			FormDivSelectFieldToField(&(_instance.RotatedTorusShape), probe.stageOfInterest, formDiv)
		case "IsHiddenSampledPoints3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenSampledPoints3DShape), formDiv)
		case "SampledPoints3DShape":
			FormDivSelectFieldToField(&(_instance.SampledPoints3DShape), probe.stageOfInterest, formDiv)
		case "IsHiddenRotatedSampledPoints3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenRotatedSampledPoints3DShape), formDiv)
		case "RotatedSampledPoints3DShape":
			FormDivSelectFieldToField(&(_instance.RotatedSampledPoints3DShape), probe.stageOfInterest, formDiv)
		case "IsHiddenEyeSampledPoints3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenEyeSampledPoints3DShape), formDiv)
		case "EyeSampledPoints3DShape":
			FormDivSelectFieldToField(&(_instance.EyeSampledPoints3DShape), probe.stageOfInterest, formDiv)
		case "IsHiddenEyeCornersSampledPoints3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenEyeCornersSampledPoints3DShape), formDiv)
		case "EyeCornersSampledPoints3DShape":
			FormDivSelectFieldToField(&(_instance.EyeCornersSampledPoints3DShape), probe.stageOfInterest, formDiv)
		case "IsHiddenEye3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenEye3DShape), formDiv)
		case "Eye3DShape":
			FormDivSelectFieldToField(&(_instance.Eye3DShape), probe.stageOfInterest, formDiv)
		case "IsHiddenEyeSeatBottomCurveShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenEyeSeatBottomCurveShape), formDiv)
		case "EyeSeatBottomCurveShape":
			FormDivSelectFieldToField(&(_instance.EyeSeatBottomCurveShape), probe.stageOfInterest, formDiv)
		case "IsHiddenEyeStoolBottomCurveShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenEyeStoolBottomCurveShape), formDiv)
		case "EyeStoolBottomCurveShape":
			FormDivSelectFieldToField(&(_instance.EyeStoolBottomCurveShape), probe.stageOfInterest, formDiv)
		case "IsHiddenSeat3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenSeat3DShape), formDiv)
		case "Seat3DShape":
			FormDivSelectFieldToField(&(_instance.Seat3DShape), probe.stageOfInterest, formDiv)
		case "IsHiddenEyeVolume3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenEyeVolume3DShape), formDiv)
		case "EyeVolume3DShape":
			FormDivSelectFieldToField(&(_instance.EyeVolume3DShape), probe.stageOfInterest, formDiv)
		case "IsHiddenSeatAndLegs3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenSeatAndLegs3DShape), formDiv)
		case "SeatAndLegs3DShape":
			FormDivSelectFieldToField(&(_instance.SeatAndLegs3DShape), probe.stageOfInterest, formDiv)
		case "IsHiddenRotatedSeatAndLegs3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenRotatedSeatAndLegs3DShape), formDiv)
		case "RotatedSeatAndLegs3DShape":
			FormDivSelectFieldToField(&(_instance.RotatedSeatAndLegs3DShape), probe.stageOfInterest, formDiv)
		case "IsHiddenTiledFloor3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenTiledFloor3DShape), formDiv)
		case "TiledFloor3DShape":
			FormDivSelectFieldToField(&(_instance.TiledFloor3DShape), probe.stageOfInterest, formDiv)
		case "Rendered3DShape":
			FormDivSelectFieldToField(&(_instance.Rendered3DShape), probe.stageOfInterest, formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(_instance.IsChecked), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "PlantAbstract:Stool3DDiagrams":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Stool3DDiagrams", func(owner *models.PlantAbstract) *[]*models.Stool3DDiagram { return &owner.Stool3DDiagrams })
		}
	}
}

func __gong__New__StoolAbstractFormCallback(
	_instance *models.StoolAbstract,
	probe *Probe,
	formGroup *form.FormGroup,
) (stoolabstractFormCallback *FormCallback[*models.StoolAbstract]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveStoolAbstractFields,
	)
}

type StoolAbstractFormCallback = FormCallback[*models.StoolAbstract]

func saveStoolAbstractFields(
	_instance *models.StoolAbstract,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "RadialRepetitions":
			FormDivBasicFieldToField(&(_instance.RadialRepetitions), formDiv)
		case "Transparency":
			FormDivBasicFieldToField(&(_instance.Transparency), formDiv)
		case "RelativeTubeDiameter":
			FormDivBasicFieldToField(&(_instance.RelativeTubeDiameter), formDiv)
		case "RelativeHeight3DTorus":
			FormDivBasicFieldToField(&(_instance.RelativeHeight3DTorus), formDiv)
		case "StoolTorusVerticalScale":
			FormDivBasicFieldToField(&(_instance.StoolTorusVerticalScale), formDiv)
		case "RelativeHeight":
			FormDivBasicFieldToField(&(_instance.RelativeHeight), formDiv)
		case "RelativeSeatThickness":
			FormDivBasicFieldToField(&(_instance.RelativeSeatThickness), formDiv)
		case "ProjectionAngle":
			FormDivBasicFieldToField(&(_instance.ProjectionAngle), formDiv)
		case "RelativeEyeSeparationCriteria":
			FormDivBasicFieldToField(&(_instance.RelativeEyeSeparationCriteria), formDiv)
		case "RelativeEyeCornerControlVectorStrength":
			FormDivBasicFieldToField(&(_instance.RelativeEyeCornerControlVectorStrength), formDiv)
		}
	}
}

func __gong__New__TiledFloor3DShapeFormCallback(
	_instance *models.TiledFloor3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (tiledfloor3dshapeFormCallback *FormCallback[*models.TiledFloor3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTiledFloor3DShapeFields,
	)
}

type TiledFloor3DShapeFormCallback = FormCallback[*models.TiledFloor3DShape]

func saveTiledFloor3DShapeFields(
	_instance *models.TiledFloor3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__TopEndArcShapeFormCallback(
	_instance *models.TopEndArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (topendarcshapeFormCallback *FormCallback[*models.TopEndArcShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTopEndArcShapeFields,
	)
}

type TopEndArcShapeFormCallback = FormCallback[*models.TopEndArcShape]

func saveTopEndArcShapeFields(
	_instance *models.TopEndArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(_instance.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(_instance.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(_instance.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(_instance.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(_instance.RadiusY), formDiv)
		case "TopEndArcShapeGrid:TopEndArcShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TopEndArcShapes", func(owner *models.TopEndArcShapeGrid) *[]*models.TopEndArcShape { return &owner.TopEndArcShapes })
		}
	}
}

func __gong__New__TopEndArcShapeGridFormCallback(
	_instance *models.TopEndArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (topendarcshapegridFormCallback *FormCallback[*models.TopEndArcShapeGrid]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTopEndArcShapeGridFields,
	)
}

type TopEndArcShapeGridFormCallback = FormCallback[*models.TopEndArcShapeGrid]

func saveTopEndArcShapeGridFields(
	_instance *models.TopEndArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "TopEndArcShapes":
			FormDivSliceOfPointersToField(_instance, "TopEndArcShapes", &(_instance.TopEndArcShapes), formDiv, probe)
		}
	}
}

func __gong__New__TopEndHalfwayArcShapeFormCallback(
	_instance *models.TopEndHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (topendhalfwayarcshapeFormCallback *FormCallback[*models.TopEndHalfwayArcShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTopEndHalfwayArcShapeFields,
	)
}

type TopEndHalfwayArcShapeFormCallback = FormCallback[*models.TopEndHalfwayArcShape]

func saveTopEndHalfwayArcShapeFields(
	_instance *models.TopEndHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(_instance.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(_instance.RadiusY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(_instance.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(_instance.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(_instance.SweepFlag), formDiv)
		case "TopEndHalfwayArcShapeGrid:TopEndHalfwayArcShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TopEndHalfwayArcShapes", func(owner *models.TopEndHalfwayArcShapeGrid) *[]*models.TopEndHalfwayArcShape { return &owner.TopEndHalfwayArcShapes })
		}
	}
}

func __gong__New__TopEndHalfwayArcShapeGridFormCallback(
	_instance *models.TopEndHalfwayArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (topendhalfwayarcshapegridFormCallback *FormCallback[*models.TopEndHalfwayArcShapeGrid]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTopEndHalfwayArcShapeGridFields,
	)
}

type TopEndHalfwayArcShapeGridFormCallback = FormCallback[*models.TopEndHalfwayArcShapeGrid]

func saveTopEndHalfwayArcShapeGridFields(
	_instance *models.TopEndHalfwayArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "TopEndHalfwayArcShapes":
			FormDivSliceOfPointersToField(_instance, "TopEndHalfwayArcShapes", &(_instance.TopEndHalfwayArcShapes), formDiv, probe)
		}
	}
}

func __gong__New__TopGrowthCurve2DFormCallback(
	_instance *models.TopGrowthCurve2D,
	probe *Probe,
	formGroup *form.FormGroup,
) (topgrowthcurve2dFormCallback *FormCallback[*models.TopGrowthCurve2D]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTopGrowthCurve2DFields,
	)
}

type TopGrowthCurve2DFormCallback = FormCallback[*models.TopGrowthCurve2D]

func saveTopGrowthCurve2DFields(
	_instance *models.TopGrowthCurve2D,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "TopStartHalfwayArcShapeGrid":
			FormDivSelectFieldToField(&(_instance.TopStartHalfwayArcShapeGrid), probe.stageOfInterest, formDiv)
		case "TopEndHalfwayArcShapeGrid":
			FormDivSelectFieldToField(&(_instance.TopEndHalfwayArcShapeGrid), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__TopMidArcVectorShapeFormCallback(
	_instance *models.TopMidArcVectorShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (topmidarcvectorshapeFormCallback *FormCallback[*models.TopMidArcVectorShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTopMidArcVectorShapeFields,
	)
}

type TopMidArcVectorShapeFormCallback = FormCallback[*models.TopMidArcVectorShape]

func saveTopMidArcVectorShapeFields(
	_instance *models.TopMidArcVectorShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "TopMidArcVectorShapeGrid:TopMidArcVectorShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TopMidArcVectorShapes", func(owner *models.TopMidArcVectorShapeGrid) *[]*models.TopMidArcVectorShape { return &owner.TopMidArcVectorShapes })
		}
	}
}

func __gong__New__TopMidArcVectorShapeGridFormCallback(
	_instance *models.TopMidArcVectorShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (topmidarcvectorshapegridFormCallback *FormCallback[*models.TopMidArcVectorShapeGrid]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTopMidArcVectorShapeGridFields,
	)
}

type TopMidArcVectorShapeGridFormCallback = FormCallback[*models.TopMidArcVectorShapeGrid]

func saveTopMidArcVectorShapeGridFields(
	_instance *models.TopMidArcVectorShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "TopMidArcVectorShapes":
			FormDivSliceOfPointersToField(_instance, "TopMidArcVectorShapes", &(_instance.TopMidArcVectorShapes), formDiv, probe)
		}
	}
}

func __gong__New__TopStackGrowthCurve2DEndHalfwayArcShapeFormCallback(
	_instance *models.TopStackGrowthCurve2DEndHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstackgrowthcurve2dendhalfwayarcshapeFormCallback *FormCallback[*models.TopStackGrowthCurve2DEndHalfwayArcShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTopStackGrowthCurve2DEndHalfwayArcShapeFields,
	)
}

type TopStackGrowthCurve2DEndHalfwayArcShapeFormCallback = FormCallback[*models.TopStackGrowthCurve2DEndHalfwayArcShape]

func saveTopStackGrowthCurve2DEndHalfwayArcShapeFields(
	_instance *models.TopStackGrowthCurve2DEndHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(_instance.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(_instance.RadiusY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(_instance.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(_instance.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(_instance.SweepFlag), formDiv)
		case "TopStackOfGrowthCurve2D:TopStackGrowthCurve2DEndHalfwayArcShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TopStackGrowthCurve2DEndHalfwayArcShapes", func(owner *models.TopStackOfGrowthCurve2D) *[]*models.TopStackGrowthCurve2DEndHalfwayArcShape { return &owner.TopStackGrowthCurve2DEndHalfwayArcShapes })
		}
	}
}

func __gong__New__TopStackGrowthCurve2DStartHalfwayArcShapeFormCallback(
	_instance *models.TopStackGrowthCurve2DStartHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstackgrowthcurve2dstarthalfwayarcshapeFormCallback *FormCallback[*models.TopStackGrowthCurve2DStartHalfwayArcShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTopStackGrowthCurve2DStartHalfwayArcShapeFields,
	)
}

type TopStackGrowthCurve2DStartHalfwayArcShapeFormCallback = FormCallback[*models.TopStackGrowthCurve2DStartHalfwayArcShape]

func saveTopStackGrowthCurve2DStartHalfwayArcShapeFields(
	_instance *models.TopStackGrowthCurve2DStartHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(_instance.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(_instance.RadiusY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(_instance.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(_instance.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(_instance.SweepFlag), formDiv)
		case "TopStackOfGrowthCurve2D:TopStackGrowthCurve2DStartHalfwayArcShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TopStackGrowthCurve2DStartHalfwayArcShapes", func(owner *models.TopStackOfGrowthCurve2D) *[]*models.TopStackGrowthCurve2DStartHalfwayArcShape { return &owner.TopStackGrowthCurve2DStartHalfwayArcShapes })
		}
	}
}

func __gong__New__TopStackOfGrowthCurve2DFormCallback(
	_instance *models.TopStackOfGrowthCurve2D,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstackofgrowthcurve2dFormCallback *FormCallback[*models.TopStackOfGrowthCurve2D]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTopStackOfGrowthCurve2DFields,
	)
}

type TopStackOfGrowthCurve2DFormCallback = FormCallback[*models.TopStackOfGrowthCurve2D]

func saveTopStackOfGrowthCurve2DFields(
	_instance *models.TopStackOfGrowthCurve2D,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "TopStackGrowthCurve2DStartHalfwayArcShapes":
			FormDivSliceOfPointersToField(_instance, "TopStackGrowthCurve2DStartHalfwayArcShapes", &(_instance.TopStackGrowthCurve2DStartHalfwayArcShapes), formDiv, probe)
		case "TopStackGrowthCurve2DEndHalfwayArcShapes":
			FormDivSliceOfPointersToField(_instance, "TopStackGrowthCurve2DEndHalfwayArcShapes", &(_instance.TopStackGrowthCurve2DEndHalfwayArcShapes), formDiv, probe)
		}
	}
}

func __gong__New__TopStackOfRotatedGrowthCurve2DFormCallback(
	_instance *models.TopStackOfRotatedGrowthCurve2D,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstackofrotatedgrowthcurve2dFormCallback *FormCallback[*models.TopStackOfRotatedGrowthCurve2D]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTopStackOfRotatedGrowthCurve2DFields,
	)
}

type TopStackOfRotatedGrowthCurve2DFormCallback = FormCallback[*models.TopStackOfRotatedGrowthCurve2D]

func saveTopStackOfRotatedGrowthCurve2DFields(
	_instance *models.TopStackOfRotatedGrowthCurve2D,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "TopStackOfRotatedGrowthCurve2DStartArcShapes":
			FormDivSliceOfPointersToField(_instance, "TopStackOfRotatedGrowthCurve2DStartArcShapes", &(_instance.TopStackOfRotatedGrowthCurve2DStartArcShapes), formDiv, probe)
		case "TopStackOfRotatedGrowthCurve2DEndArcShapes":
			FormDivSliceOfPointersToField(_instance, "TopStackOfRotatedGrowthCurve2DEndArcShapes", &(_instance.TopStackOfRotatedGrowthCurve2DEndArcShapes), formDiv, probe)
		}
	}
}

func __gong__New__TopStackOfRotatedGrowthCurve2DEndArcShapeFormCallback(
	_instance *models.TopStackOfRotatedGrowthCurve2DEndArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstackofrotatedgrowthcurve2dendarcshapeFormCallback *FormCallback[*models.TopStackOfRotatedGrowthCurve2DEndArcShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTopStackOfRotatedGrowthCurve2DEndArcShapeFields,
	)
}

type TopStackOfRotatedGrowthCurve2DEndArcShapeFormCallback = FormCallback[*models.TopStackOfRotatedGrowthCurve2DEndArcShape]

func saveTopStackOfRotatedGrowthCurve2DEndArcShapeFields(
	_instance *models.TopStackOfRotatedGrowthCurve2DEndArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(_instance.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(_instance.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(_instance.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(_instance.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(_instance.RadiusY), formDiv)
		case "TopStackOfRotatedGrowthCurve2D:TopStackOfRotatedGrowthCurve2DEndArcShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TopStackOfRotatedGrowthCurve2DEndArcShapes", func(owner *models.TopStackOfRotatedGrowthCurve2D) *[]*models.TopStackOfRotatedGrowthCurve2DEndArcShape { return &owner.TopStackOfRotatedGrowthCurve2DEndArcShapes })
		}
	}
}

func __gong__New__TopStackOfRotatedGrowthCurve2DStartArcShapeFormCallback(
	_instance *models.TopStackOfRotatedGrowthCurve2DStartArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstackofrotatedgrowthcurve2dstartarcshapeFormCallback *FormCallback[*models.TopStackOfRotatedGrowthCurve2DStartArcShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTopStackOfRotatedGrowthCurve2DStartArcShapeFields,
	)
}

type TopStackOfRotatedGrowthCurve2DStartArcShapeFormCallback = FormCallback[*models.TopStackOfRotatedGrowthCurve2DStartArcShape]

func saveTopStackOfRotatedGrowthCurve2DStartArcShapeFields(
	_instance *models.TopStackOfRotatedGrowthCurve2DStartArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(_instance.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(_instance.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(_instance.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(_instance.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(_instance.RadiusY), formDiv)
		case "TopStackOfRotatedGrowthCurve2D:TopStackOfRotatedGrowthCurve2DStartArcShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TopStackOfRotatedGrowthCurve2DStartArcShapes", func(owner *models.TopStackOfRotatedGrowthCurve2D) *[]*models.TopStackOfRotatedGrowthCurve2DStartArcShape { return &owner.TopStackOfRotatedGrowthCurve2DStartArcShapes })
		}
	}
}

func __gong__New__TopStartArcShapeFormCallback(
	_instance *models.TopStartArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstartarcshapeFormCallback *FormCallback[*models.TopStartArcShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTopStartArcShapeFields,
	)
}

type TopStartArcShapeFormCallback = FormCallback[*models.TopStartArcShape]

func saveTopStartArcShapeFields(
	_instance *models.TopStartArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(_instance.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(_instance.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(_instance.SweepFlag), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(_instance.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(_instance.RadiusY), formDiv)
		case "TopStartArcShapeGrid:TopStartArcShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TopStartArcShapes", func(owner *models.TopStartArcShapeGrid) *[]*models.TopStartArcShape { return &owner.TopStartArcShapes })
		}
	}
}

func __gong__New__TopStartArcShapeGridFormCallback(
	_instance *models.TopStartArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstartarcshapegridFormCallback *FormCallback[*models.TopStartArcShapeGrid]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTopStartArcShapeGridFields,
	)
}

type TopStartArcShapeGridFormCallback = FormCallback[*models.TopStartArcShapeGrid]

func saveTopStartArcShapeGridFields(
	_instance *models.TopStartArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "TopStartArcShapes":
			FormDivSliceOfPointersToField(_instance, "TopStartArcShapes", &(_instance.TopStartArcShapes), formDiv, probe)
		}
	}
}

func __gong__New__TopStartHalfwayArcShapeFormCallback(
	_instance *models.TopStartHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstarthalfwayarcshapeFormCallback *FormCallback[*models.TopStartHalfwayArcShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTopStartHalfwayArcShapeFields,
	)
}

type TopStartHalfwayArcShapeFormCallback = FormCallback[*models.TopStartHalfwayArcShape]

func saveTopStartHalfwayArcShapeFields(
	_instance *models.TopStartHalfwayArcShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "StartX":
			FormDivBasicFieldToField(&(_instance.StartX), formDiv)
		case "StartY":
			FormDivBasicFieldToField(&(_instance.StartY), formDiv)
		case "EndX":
			FormDivBasicFieldToField(&(_instance.EndX), formDiv)
		case "EndY":
			FormDivBasicFieldToField(&(_instance.EndY), formDiv)
		case "RadiusX":
			FormDivBasicFieldToField(&(_instance.RadiusX), formDiv)
		case "RadiusY":
			FormDivBasicFieldToField(&(_instance.RadiusY), formDiv)
		case "XAxisRotation":
			FormDivBasicFieldToField(&(_instance.XAxisRotation), formDiv)
		case "LargeArcFlag":
			FormDivBasicFieldToField(&(_instance.LargeArcFlag), formDiv)
		case "SweepFlag":
			FormDivBasicFieldToField(&(_instance.SweepFlag), formDiv)
		case "TopStartHalfwayArcShapeGrid:TopStartHalfwayArcShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TopStartHalfwayArcShapes", func(owner *models.TopStartHalfwayArcShapeGrid) *[]*models.TopStartHalfwayArcShape { return &owner.TopStartHalfwayArcShapes })
		}
	}
}

func __gong__New__TopStartHalfwayArcShapeGridFormCallback(
	_instance *models.TopStartHalfwayArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) (topstarthalfwayarcshapegridFormCallback *FormCallback[*models.TopStartHalfwayArcShapeGrid]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTopStartHalfwayArcShapeGridFields,
	)
}

type TopStartHalfwayArcShapeGridFormCallback = FormCallback[*models.TopStartHalfwayArcShapeGrid]

func saveTopStartHalfwayArcShapeGridFields(
	_instance *models.TopStartHalfwayArcShapeGrid,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "TopStartHalfwayArcShapes":
			FormDivSliceOfPointersToField(_instance, "TopStartHalfwayArcShapes", &(_instance.TopStartHalfwayArcShapes), formDiv, probe)
		}
	}
}

func __gong__New__Torus3DShapeFormCallback(
	_instance *models.Torus3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (torus3dshapeFormCallback *FormCallback[*models.Torus3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTorus3DShapeFields,
	)
}

type Torus3DShapeFormCallback = FormCallback[*models.Torus3DShape]

func saveTorus3DShapeFields(
	_instance *models.Torus3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__TorusEdge3DShapeFormCallback(
	_instance *models.TorusEdge3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (torusedge3dshapeFormCallback *FormCallback[*models.TorusEdge3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTorusEdge3DShapeFields,
	)
}

type TorusEdge3DShapeFormCallback = FormCallback[*models.TorusEdge3DShape]

func saveTorusEdge3DShapeFields(
	_instance *models.TorusEdge3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__TorusStackShapeFormCallback(
	_instance *models.TorusStackShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (torusstackshapeFormCallback *FormCallback[*models.TorusStackShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTorusStackShapeFields,
	)
}

type TorusStackShapeFormCallback = FormCallback[*models.TorusStackShape]

func saveTorusStackShapeFields(
	_instance *models.TorusStackShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__TubeVase3DDiagramFormCallback(
	_instance *models.TubeVase3DDiagram,
	probe *Probe,
	formGroup *form.FormGroup,
) (tubevase3ddiagramFormCallback *FormCallback[*models.TubeVase3DDiagram]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTubeVase3DDiagramFields,
	)
}

type TubeVase3DDiagramFormCallback = FormCallback[*models.TubeVase3DDiagram]

func saveTubeVase3DDiagramFields(
	_instance *models.TubeVase3DDiagram,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon":
			FormDivBasicFieldToField(&(_instance.IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon), formDiv)
		case "IsHiddenTorusStackShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenTorusStackShape), formDiv)
		case "IsHiddenVerticalTorusStackShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenVerticalTorusStackShape), formDiv)
		case "IsHiddenPartiallyRotatedTorusShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenPartiallyRotatedTorusShape), formDiv)
		case "IsHiddenStackOfPartiallyRotatedTorusShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenStackOfPartiallyRotatedTorusShape), formDiv)
		case "IsHiddenPointsAndLines3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenPointsAndLines3DShape), formDiv)
		case "IsHiddenKeyHole3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenKeyHole3DShape), formDiv)
		case "IsHiddenKey3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenKey3DShape), formDiv)
		case "IsHiddenVolumeKey3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenVolumeKey3DShape), formDiv)
		case "IsHiddenTorusEdge3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenTorusEdge3DShape), formDiv)
		case "IsHiddenSampledPoints3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenSampledPoints3DShape), formDiv)
		case "IsHiddenOriginalPoints3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenOriginalPoints3DShape), formDiv)
		case "IsHiddenAngle0Shape":
			FormDivBasicFieldToField(&(_instance.IsHiddenAngle0Shape), formDiv)
		case "IsHiddenTiledFloor3DShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenTiledFloor3DShape), formDiv)
		case "Rendered3DShape":
			FormDivSelectFieldToField(&(_instance.Rendered3DShape), probe.stageOfInterest, formDiv)
		case "TorusStackShape":
			FormDivSelectFieldToField(&(_instance.TorusStackShape), probe.stageOfInterest, formDiv)
		case "VerticalTorusStackShape":
			FormDivSelectFieldToField(&(_instance.VerticalTorusStackShape), probe.stageOfInterest, formDiv)
		case "PartiallyRotatedTorusShape":
			FormDivSelectFieldToField(&(_instance.PartiallyRotatedTorusShape), probe.stageOfInterest, formDiv)
		case "StackOfPartiallyRotatedTorusShape":
			FormDivSelectFieldToField(&(_instance.StackOfPartiallyRotatedTorusShape), probe.stageOfInterest, formDiv)
		case "PointsAndLines3DShape":
			FormDivSelectFieldToField(&(_instance.PointsAndLines3DShape), probe.stageOfInterest, formDiv)
		case "SampledPoints3DShape":
			FormDivSelectFieldToField(&(_instance.SampledPoints3DShape), probe.stageOfInterest, formDiv)
		case "OriginalPoints3DShape":
			FormDivSelectFieldToField(&(_instance.OriginalPoints3DShape), probe.stageOfInterest, formDiv)
		case "Angle0Shape":
			FormDivSelectFieldToField(&(_instance.Angle0Shape), probe.stageOfInterest, formDiv)
		case "KeyHole3DShape":
			FormDivSelectFieldToField(&(_instance.KeyHole3DShape), probe.stageOfInterest, formDiv)
		case "Key3DShape":
			FormDivSelectFieldToField(&(_instance.Key3DShape), probe.stageOfInterest, formDiv)
		case "VolumeKey3DShape":
			FormDivSelectFieldToField(&(_instance.VolumeKey3DShape), probe.stageOfInterest, formDiv)
		case "TorusEdge3DShape":
			FormDivSelectFieldToField(&(_instance.TorusEdge3DShape), probe.stageOfInterest, formDiv)
		case "TiledFloor3DShape":
			FormDivSelectFieldToField(&(_instance.TiledFloor3DShape), probe.stageOfInterest, formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(_instance.IsChecked), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "PlantAbstract:TubeVase3DDiagrams":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TubeVase3DDiagrams", func(owner *models.PlantAbstract) *[]*models.TubeVase3DDiagram { return &owner.TubeVase3DDiagrams })
		}
	}
}

func __gong__New__TubeVaseAbstractFormCallback(
	_instance *models.TubeVaseAbstract,
	probe *Probe,
	formGroup *form.FormGroup,
) (tubevaseabstractFormCallback *FormCallback[*models.TubeVaseAbstract]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTubeVaseAbstractFields,
	)
}

type TubeVaseAbstractFormCallback = FormCallback[*models.TubeVaseAbstract]

func saveTubeVaseAbstractFields(
	_instance *models.TubeVaseAbstract,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Z_Ribbon":
			FormDivBasicFieldToField(&(_instance.Z_Ribbon), formDiv)
		case "RelativeVerticalThickness":
			FormDivBasicFieldToField(&(_instance.RelativeVerticalThickness), formDiv)
		case "RelativeRadialThickness":
			FormDivBasicFieldToField(&(_instance.RelativeRadialThickness), formDiv)
		case "RelativeCuttedStackFloorHeight":
			FormDivBasicFieldToField(&(_instance.RelativeCuttedStackFloorHeight), formDiv)
		case "RelativeRotatedTorusSeparation":
			FormDivBasicFieldToField(&(_instance.RelativeRotatedTorusSeparation), formDiv)
		case "RotationRatio":
			FormDivBasicFieldToField(&(_instance.RotationRatio), formDiv)
		case "RadialRepetitions":
			FormDivBasicFieldToField(&(_instance.RadialRepetitions), formDiv)
		case "Transparency":
			FormDivBasicFieldToField(&(_instance.Transparency), formDiv)
		case "HasAlternatingRingColors":
			FormDivBasicFieldToField(&(_instance.HasAlternatingRingColors), formDiv)
		case "RelativeTrajectoryOffsetX":
			FormDivBasicFieldToField(&(_instance.RelativeTrajectoryOffsetX), formDiv)
		case "RelativeTrajectoryOffsetY":
			FormDivBasicFieldToField(&(_instance.RelativeTrajectoryOffsetY), formDiv)
		case "NbStepP1P2":
			FormDivBasicFieldToField(&(_instance.NbStepP1P2), formDiv)
		case "ChosenStep":
			FormDivBasicFieldToField(&(_instance.ChosenStep), formDiv)
		case "RelativeHorizontalRingsHeight":
			FormDivBasicFieldToField(&(_instance.RelativeHorizontalRingsHeight), formDiv)
		case "OffsetKeyX":
			FormDivBasicFieldToField(&(_instance.OffsetKeyX), formDiv)
		case "OffsetKeyY":
			FormDivBasicFieldToField(&(_instance.OffsetKeyY), formDiv)
		case "HeightKey":
			FormDivBasicFieldToField(&(_instance.HeightKey), formDiv)
		case "WidthKey":
			FormDivBasicFieldToField(&(_instance.WidthKey), formDiv)
		case "RelativeKeySize":
			FormDivBasicFieldToField(&(_instance.RelativeKeySize), formDiv)
		case "MovieNbFrames":
			FormDivBasicFieldToField(&(_instance.MovieNbFrames), formDiv)
		case "PerpendicularVectorGridHalfway":
			FormDivSelectFieldToField(&(_instance.PerpendicularVectorGridHalfway), probe.stageOfInterest, formDiv)
		case "TopStartArcShapeGrid":
			FormDivSelectFieldToField(&(_instance.TopStartArcShapeGrid), probe.stageOfInterest, formDiv)
		case "TopEndArcShapeGrid":
			FormDivSelectFieldToField(&(_instance.TopEndArcShapeGrid), probe.stageOfInterest, formDiv)
		case "ShiftedBottomTopStartArcShapeGrid":
			FormDivSelectFieldToField(&(_instance.ShiftedBottomTopStartArcShapeGrid), probe.stageOfInterest, formDiv)
		case "TopMidArcVectorShapeGrid":
			FormDivSelectFieldToField(&(_instance.TopMidArcVectorShapeGrid), probe.stageOfInterest, formDiv)
		case "StartHalfwayArcShapeGrid":
			FormDivSelectFieldToField(&(_instance.StartHalfwayArcShapeGrid), probe.stageOfInterest, formDiv)
		case "TopStartHalfwayArcShapeGrid":
			FormDivSelectFieldToField(&(_instance.TopStartHalfwayArcShapeGrid), probe.stageOfInterest, formDiv)
		case "EndHalfwayArcShapeGrid":
			FormDivSelectFieldToField(&(_instance.EndHalfwayArcShapeGrid), probe.stageOfInterest, formDiv)
		case "TopEndHalfwayArcShapeGrid":
			FormDivSelectFieldToField(&(_instance.TopEndHalfwayArcShapeGrid), probe.stageOfInterest, formDiv)
		case "StackOfRotatedGrowthCurve2D":
			FormDivSelectFieldToField(&(_instance.StackOfRotatedGrowthCurve2D), probe.stageOfInterest, formDiv)
		case "TopStackOfRotatedGrowthCurve2D":
			FormDivSelectFieldToField(&(_instance.TopStackOfRotatedGrowthCurve2D), probe.stageOfInterest, formDiv)
		case "TopGrowthCurve2D":
			FormDivSelectFieldToField(&(_instance.TopGrowthCurve2D), probe.stageOfInterest, formDiv)
		case "StackOfGrowthCurve2D":
			FormDivSelectFieldToField(&(_instance.StackOfGrowthCurve2D), probe.stageOfInterest, formDiv)
		case "TopStackOfGrowthCurve2D":
			FormDivSelectFieldToField(&(_instance.TopStackOfGrowthCurve2D), probe.stageOfInterest, formDiv)
		case "StackOfGrowthCurve2DRibbon":
			FormDivSelectFieldToField(&(_instance.StackOfGrowthCurve2DRibbon), probe.stageOfInterest, formDiv)
		case "StackOfRotatedGrowthCurve2DRibbon":
			FormDivSelectFieldToField(&(_instance.StackOfRotatedGrowthCurve2DRibbon), probe.stageOfInterest, formDiv)
		case "GrowthCurve2DRibbon":
			FormDivSelectFieldToField(&(_instance.GrowthCurve2DRibbon), probe.stageOfInterest, formDiv)
		case "ShiftedRightGrowthCurve2DRibbon":
			FormDivSelectFieldToField(&(_instance.ShiftedRightGrowthCurve2DRibbon), probe.stageOfInterest, formDiv)
		case "ShiftedLeftGrowthCurve2DRibbon":
			FormDivSelectFieldToField(&(_instance.ShiftedLeftGrowthCurve2DRibbon), probe.stageOfInterest, formDiv)
		case "PartiallyGrowthCurve2DRibbon":
			FormDivSelectFieldToField(&(_instance.PartiallyGrowthCurve2DRibbon), probe.stageOfInterest, formDiv)
		case "ShiftedLeftPartiallyGrowthCurve2DRibbon":
			FormDivSelectFieldToField(&(_instance.ShiftedLeftPartiallyGrowthCurve2DRibbon), probe.stageOfInterest, formDiv)
		case "PartiallyGrowthCurve2DTrajectory":
			FormDivSelectFieldToField(&(_instance.PartiallyGrowthCurve2DTrajectory), probe.stageOfInterest, formDiv)
		case "PartiallyGrowthCurve2DTrajectoryP1P2":
			FormDivSelectFieldToField(&(_instance.PartiallyGrowthCurve2DTrajectoryP1P2), probe.stageOfInterest, formDiv)
		case "PxShape":
			FormDivSelectFieldToField(&(_instance.PxShape), probe.stageOfInterest, formDiv)
		case "ChosenP1P2PairShape":
			FormDivSelectFieldToField(&(_instance.ChosenP1P2PairShape), probe.stageOfInterest, formDiv)
		case "KeyHoleShape":
			FormDivSelectFieldToField(&(_instance.KeyHoleShape), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__Vase2DDiagramFormCallback(
	_instance *models.Vase2DDiagram,
	probe *Probe,
	formGroup *form.FormGroup,
) (vase2ddiagramFormCallback *FormCallback[*models.Vase2DDiagram]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveVase2DDiagramFields,
	)
}

type Vase2DDiagramFormCallback = FormCallback[*models.Vase2DDiagram]

func saveVase2DDiagramFields(
	_instance *models.Vase2DDiagram,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Zoom":
			FormDivBasicFieldToField(&(_instance.Zoom), formDiv)
		case "IsVaseArcNodesExpanded":
			FormDivBasicFieldToField(&(_instance.IsVaseArcNodesExpanded), formDiv)
		case "IsVaseClampingNodesExpanded":
			FormDivBasicFieldToField(&(_instance.IsVaseClampingNodesExpanded), formDiv)
		case "IsHiddenAxesShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenAxesShape), formDiv)
		case "IsHiddenBottomStartArcShapeGrid":
			FormDivBasicFieldToField(&(_instance.IsHiddenBottomStartArcShapeGrid), formDiv)
		case "IsHiddenBottomEndArcShapeGrid":
			FormDivBasicFieldToField(&(_instance.IsHiddenBottomEndArcShapeGrid), formDiv)
		case "IsHiddenBottomStackOfGrowthCurve":
			FormDivBasicFieldToField(&(_instance.IsHiddenBottomStackOfGrowthCurve), formDiv)
		case "IsHiddenShiftedLeftStackOfGrowthCurve":
			FormDivBasicFieldToField(&(_instance.IsHiddenShiftedLeftStackOfGrowthCurve), formDiv)
		case "IsHiddenShiftedLeftStackOfNormalVector":
			FormDivBasicFieldToField(&(_instance.IsHiddenShiftedLeftStackOfNormalVector), formDiv)
		case "IsHiddenPerpendicularVectorGridHalfway":
			FormDivBasicFieldToField(&(_instance.IsHiddenPerpendicularVectorGridHalfway), formDiv)
		case "IsHiddenTopStartArcShapeGrid":
			FormDivBasicFieldToField(&(_instance.IsHiddenTopStartArcShapeGrid), formDiv)
		case "IsHiddenShiftedBottomTopStartArcShapeGrid":
			FormDivBasicFieldToField(&(_instance.IsHiddenShiftedBottomTopStartArcShapeGrid), formDiv)
		case "IsHiddenTopMidArcVectorShapeGrid":
			FormDivBasicFieldToField(&(_instance.IsHiddenTopMidArcVectorShapeGrid), formDiv)
		case "IsHiddenStartHalfwayArcShapeGrid":
			FormDivBasicFieldToField(&(_instance.IsHiddenStartHalfwayArcShapeGrid), formDiv)
		case "IsHiddenTopStartHalfwayArcShapeGrid":
			FormDivBasicFieldToField(&(_instance.IsHiddenTopStartHalfwayArcShapeGrid), formDiv)
		case "IsHiddenEndHalfwayArcShapeGrid":
			FormDivBasicFieldToField(&(_instance.IsHiddenEndHalfwayArcShapeGrid), formDiv)
		case "IsHiddenTopEndHalfwayArcShapeGrid":
			FormDivBasicFieldToField(&(_instance.IsHiddenTopEndHalfwayArcShapeGrid), formDiv)
		case "IsHiddenTopEndArcShapeGrid":
			FormDivBasicFieldToField(&(_instance.IsHiddenTopEndArcShapeGrid), formDiv)
		case "IsHiddenStackOfGrowthCurve":
			FormDivBasicFieldToField(&(_instance.IsHiddenStackOfGrowthCurve), formDiv)
		case "IsHiddenTopStackOfGrowthCurve":
			FormDivBasicFieldToField(&(_instance.IsHiddenTopStackOfGrowthCurve), formDiv)
		case "IsHiddenTopGrowthCurve2D":
			FormDivBasicFieldToField(&(_instance.IsHiddenTopGrowthCurve2D), formDiv)
		case "IsHiddenStackOfGrowthCurve2D":
			FormDivBasicFieldToField(&(_instance.IsHiddenStackOfGrowthCurve2D), formDiv)
		case "IsHiddenTopStackOfGrowthCurve2D":
			FormDivBasicFieldToField(&(_instance.IsHiddenTopStackOfGrowthCurve2D), formDiv)
		case "IsHiddenGrowthCurve2DRibbon":
			FormDivBasicFieldToField(&(_instance.IsHiddenGrowthCurve2DRibbon), formDiv)
		case "IsHiddenShiftedRightGrowthCurve2DRibbon":
			FormDivBasicFieldToField(&(_instance.IsHiddenShiftedRightGrowthCurve2DRibbon), formDiv)
		case "IsHiddenShiftedLeftGrowthCurve2DRibbon":
			FormDivBasicFieldToField(&(_instance.IsHiddenShiftedLeftGrowthCurve2DRibbon), formDiv)
		case "IsHiddenStackOfGrowthCurve2DRibbon":
			FormDivBasicFieldToField(&(_instance.IsHiddenStackOfGrowthCurve2DRibbon), formDiv)
		case "IsHiddenStackOfRotatedGrowthCurve2DRibbon":
			FormDivBasicFieldToField(&(_instance.IsHiddenStackOfRotatedGrowthCurve2DRibbon), formDiv)
		case "IsHiddenPartiallyGrowthCurve2DRibbon":
			FormDivBasicFieldToField(&(_instance.IsHiddenPartiallyGrowthCurve2DRibbon), formDiv)
		case "IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon":
			FormDivBasicFieldToField(&(_instance.IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon), formDiv)
		case "IsHiddenPartiallyGrowthCurve2DTrajectory":
			FormDivBasicFieldToField(&(_instance.IsHiddenPartiallyGrowthCurve2DTrajectory), formDiv)
		case "IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2":
			FormDivBasicFieldToField(&(_instance.IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2), formDiv)
		case "IsHiddenPxShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenPxShape), formDiv)
		case "IsHiddenChosenP1P2PairShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenChosenP1P2PairShape), formDiv)
		case "IsHiddenKeyHoleShape":
			FormDivBasicFieldToField(&(_instance.IsHiddenKeyHoleShape), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(_instance.IsChecked), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "PlantAbstract:Vase2DDiagrams":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Vase2DDiagrams", func(owner *models.PlantAbstract) *[]*models.Vase2DDiagram { return &owner.Vase2DDiagrams })
		}
	}
}

func __gong__New__VerticalTorusStackShapeFormCallback(
	_instance *models.VerticalTorusStackShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (verticaltorusstackshapeFormCallback *FormCallback[*models.VerticalTorusStackShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveVerticalTorusStackShapeFields,
	)
}

type VerticalTorusStackShapeFormCallback = FormCallback[*models.VerticalTorusStackShape]

func saveVerticalTorusStackShapeFields(
	_instance *models.VerticalTorusStackShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

func __gong__New__VolumeKey3DShapeFormCallback(
	_instance *models.VolumeKey3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (volumekey3dshapeFormCallback *FormCallback[*models.VolumeKey3DShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveVolumeKey3DShapeFields,
	)
}

type VolumeKey3DShapeFormCallback = FormCallback[*models.VolumeKey3DShape]

func saveVolumeKey3DShapeFields(
	_instance *models.VolumeKey3DShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		}
	}
}

