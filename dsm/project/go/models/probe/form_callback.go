// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/project/go/models"
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
func __gong__New__DiagramFormCallback(
	_instance *models.Diagram,
	probe *Probe,
	formGroup *form.FormGroup,
) (diagramFormCallback *FormCallback[*models.Diagram]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDiagramFields,
	)
}

type DiagramFormCallback = FormCallback[*models.Diagram]

func saveDiagramFields(
	_instance *models.Diagram,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "DefaultBoxWidth":
			FormDivBasicFieldToField(&(_instance.DefaultBoxWidth), formDiv)
		case "DefaultBoxHeigth":
			FormDivBasicFieldToField(&(_instance.DefaultBoxHeigth), formDiv)
		case "DateFormat":
			FormDivBasicFieldToField(&(_instance.DateFormat), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "IsTimeDiagram":
			FormDivBasicFieldToField(&(_instance.IsTimeDiagram), formDiv)
		case "ComputedStart":
			FormDivTimeFieldToField(&(_instance.ComputedStart), formDiv, false)
		case "ComputedEnd":
			FormDivTimeFieldToField(&(_instance.ComputedEnd), formDiv, false)
		case "ComputedDuration":
			FormDivBasicFieldToField(&(_instance.ComputedDuration), formDiv)
		case "UseManualStartAndEndDates":
			FormDivBasicFieldToField(&(_instance.UseManualStartAndEndDates), formDiv)
		case "ManualStart":
			FormDivTimeFieldToField(&(_instance.ManualStart), formDiv, false)
		case "ManualEnd":
			FormDivTimeFieldToField(&(_instance.ManualEnd), formDiv, false)
		case "TimeStep":
			FormDivBasicFieldToField(&(_instance.TimeStep), formDiv)
		case "TimeStepScale":
			FormDivEnumStringFieldToField(&(_instance.TimeStepScale), formDiv)
		case "LaneHeight":
			FormDivBasicFieldToField(&(_instance.LaneHeight), formDiv)
		case "RatioBarToLaneHeight":
			FormDivBasicFieldToField(&(_instance.RatioBarToLaneHeight), formDiv)
		case "YTopMargin":
			FormDivBasicFieldToField(&(_instance.YTopMargin), formDiv)
		case "XLeftText":
			FormDivBasicFieldToField(&(_instance.XLeftText), formDiv)
		case "TextHeight":
			FormDivBasicFieldToField(&(_instance.TextHeight), formDiv)
		case "XLeftLanes":
			FormDivBasicFieldToField(&(_instance.XLeftLanes), formDiv)
		case "XRightMargin":
			FormDivBasicFieldToField(&(_instance.XRightMargin), formDiv)
		case "ArrowLengthToTheRightOfStartBar":
			FormDivBasicFieldToField(&(_instance.ArrowLengthToTheRightOfStartBar), formDiv)
		case "ArrowTipLenght":
			FormDivBasicFieldToField(&(_instance.ArrowTipLenght), formDiv)
		case "TimeLine_Color":
			FormDivBasicFieldToField(&(_instance.TimeLine_Color), formDiv)
		case "TimeLine_FillOpacity":
			FormDivBasicFieldToField(&(_instance.TimeLine_FillOpacity), formDiv)
		case "TimeLine_Stroke":
			FormDivBasicFieldToField(&(_instance.TimeLine_Stroke), formDiv)
		case "TimeLine_StrokeWidth":
			FormDivBasicFieldToField(&(_instance.TimeLine_StrokeWidth), formDiv)
		case "DrawVerticalTimeLines":
			FormDivBasicFieldToField(&(_instance.DrawVerticalTimeLines), formDiv)
		case "Group_Stroke":
			FormDivBasicFieldToField(&(_instance.Group_Stroke), formDiv)
		case "Group_StrokeWidth":
			FormDivBasicFieldToField(&(_instance.Group_StrokeWidth), formDiv)
		case "Group_StrokeDashArray":
			FormDivBasicFieldToField(&(_instance.Group_StrokeDashArray), formDiv)
		case "DateYOffset":
			FormDivBasicFieldToField(&(_instance.DateYOffset), formDiv)
		case "AlignOnStartEndOnYearStart":
			FormDivBasicFieldToField(&(_instance.AlignOnStartEndOnYearStart), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(_instance.IsChecked), formDiv)
		case "IsEditable_":
			FormDivBasicFieldToField(&(_instance.IsEditable_), formDiv)
		case "IsShowPrefix":
			FormDivBasicFieldToField(&(_instance.IsShowPrefix), formDiv)
		case "IsInAutoLayoutMode":
			FormDivBasicFieldToField(&(_instance.IsInAutoLayoutMode), formDiv)
		case "Product_Shapes":
			FormDivSliceOfPointersToField(_instance, "Product_Shapes", &(_instance.Product_Shapes), formDiv, probe)
		case "ProductsWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ProductsWhoseNodeIsExpanded", &(_instance.ProductsWhoseNodeIsExpanded), formDiv, probe)
		case "IsPBSNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsPBSNodeExpanded), formDiv)
		case "ProductComposition_Shapes":
			FormDivSliceOfPointersToField(_instance, "ProductComposition_Shapes", &(_instance.ProductComposition_Shapes), formDiv, probe)
		case "ProductReference_Shapes":
			FormDivSliceOfPointersToField(_instance, "ProductReference_Shapes", &(_instance.ProductReference_Shapes), formDiv, probe)
		case "IsWBSNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsWBSNodeExpanded), formDiv)
		case "Task_Shapes":
			FormDivSliceOfPointersToField(_instance, "Task_Shapes", &(_instance.Task_Shapes), formDiv, probe)
		case "TasksWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "TasksWhoseNodeIsExpanded", &(_instance.TasksWhoseNodeIsExpanded), formDiv, probe)
		case "TasksWhoseInputNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "TasksWhoseInputNodeIsExpanded", &(_instance.TasksWhoseInputNodeIsExpanded), formDiv, probe)
		case "TasksWhoseOutputNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "TasksWhoseOutputNodeIsExpanded", &(_instance.TasksWhoseOutputNodeIsExpanded), formDiv, probe)
		case "TasksWhosePredecessorNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "TasksWhosePredecessorNodeIsExpanded", &(_instance.TasksWhosePredecessorNodeIsExpanded), formDiv, probe)
		case "IsTaskGroupsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsTaskGroupsNodeExpanded), formDiv)
		case "TaskGroupShapes":
			FormDivSliceOfPointersToField(_instance, "TaskGroupShapes", &(_instance.TaskGroupShapes), formDiv, probe)
		case "TaskGroupsWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "TaskGroupsWhoseNodeIsExpanded", &(_instance.TaskGroupsWhoseNodeIsExpanded), formDiv, probe)
		case "TaskComposition_Shapes":
			FormDivSliceOfPointersToField(_instance, "TaskComposition_Shapes", &(_instance.TaskComposition_Shapes), formDiv, probe)
		case "TaskInputShapes":
			FormDivSliceOfPointersToField(_instance, "TaskInputShapes", &(_instance.TaskInputShapes), formDiv, probe)
		case "TaskOutputShapes":
			FormDivSliceOfPointersToField(_instance, "TaskOutputShapes", &(_instance.TaskOutputShapes), formDiv, probe)
		case "TaskPredecessorShapes":
			FormDivSliceOfPointersToField(_instance, "TaskPredecessorShapes", &(_instance.TaskPredecessorShapes), formDiv, probe)
		case "Note_Shapes":
			FormDivSliceOfPointersToField(_instance, "Note_Shapes", &(_instance.Note_Shapes), formDiv, probe)
		case "NotesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "NotesWhoseNodeIsExpanded", &(_instance.NotesWhoseNodeIsExpanded), formDiv, probe)
		case "IsNotesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsNotesNodeExpanded), formDiv)
		case "NoteProductShapes":
			FormDivSliceOfPointersToField(_instance, "NoteProductShapes", &(_instance.NoteProductShapes), formDiv, probe)
		case "NoteTaskShapes":
			FormDivSliceOfPointersToField(_instance, "NoteTaskShapes", &(_instance.NoteTaskShapes), formDiv, probe)
		case "NoteResourceShapes":
			FormDivSliceOfPointersToField(_instance, "NoteResourceShapes", &(_instance.NoteResourceShapes), formDiv, probe)
		case "Resource_Shapes":
			FormDivSliceOfPointersToField(_instance, "Resource_Shapes", &(_instance.Resource_Shapes), formDiv, probe)
		case "ResourcesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ResourcesWhoseNodeIsExpanded", &(_instance.ResourcesWhoseNodeIsExpanded), formDiv, probe)
		case "IsResourcesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsResourcesNodeExpanded), formDiv)
		case "ResourceComposition_Shapes":
			FormDivSliceOfPointersToField(_instance, "ResourceComposition_Shapes", &(_instance.ResourceComposition_Shapes), formDiv, probe)
		case "ResourceTaskShapes":
			FormDivSliceOfPointersToField(_instance, "ResourceTaskShapes", &(_instance.ResourceTaskShapes), formDiv, probe)
		case "Library:Diagrams":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Diagrams", func(owner *models.Library) *[]*models.Diagram { return &owner.Diagrams })
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
		case "RootProducts":
			FormDivSliceOfPointersToField(_instance, "RootProducts", &(_instance.RootProducts), formDiv, probe)
		case "RootTasks":
			FormDivSliceOfPointersToField(_instance, "RootTasks", &(_instance.RootTasks), formDiv, probe)
		case "RootTaskGroups":
			FormDivSliceOfPointersToField(_instance, "RootTaskGroups", &(_instance.RootTaskGroups), formDiv, probe)
		case "RootResources":
			FormDivSliceOfPointersToField(_instance, "RootResources", &(_instance.RootResources), formDiv, probe)
		case "Notes":
			FormDivSliceOfPointersToField(_instance, "Notes", &(_instance.Notes), formDiv, probe)
		case "Diagrams":
			FormDivSliceOfPointersToField(_instance, "Diagrams", &(_instance.Diagrams), formDiv, probe)
		case "Library:SubLibraries":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubLibraries", func(owner *models.Library) *[]*models.Library { return &owner.SubLibraries })
		}
	}
}

func __gong__New__NoteFormCallback(
	_instance *models.Note,
	probe *Probe,
	formGroup *form.FormGroup,
) (noteFormCallback *FormCallback[*models.Note]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNoteFields,
	)
}

type NoteFormCallback = FormCallback[*models.Note]

func saveNoteFields(
	_instance *models.Note,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(_instance.LayoutDirection), formDiv)
		case "Products":
			FormDivSliceOfPointersToField(_instance, "Products", &(_instance.Products), formDiv, probe)
		case "Tasks":
			FormDivSliceOfPointersToField(_instance, "Tasks", &(_instance.Tasks), formDiv, probe)
		case "Resources":
			FormDivSliceOfPointersToField(_instance, "Resources", &(_instance.Resources), formDiv, probe)
		case "Diagram:NotesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "NotesWhoseNodeIsExpanded", func(owner *models.Diagram) *[]*models.Note { return &owner.NotesWhoseNodeIsExpanded })
		case "Library:Notes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Notes", func(owner *models.Library) *[]*models.Note { return &owner.Notes })
		}
	}
}

func __gong__New__NoteProductShapeFormCallback(
	_instance *models.NoteProductShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (noteproductshapeFormCallback *FormCallback[*models.NoteProductShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNoteProductShapeFields,
	)
}

type NoteProductShapeFormCallback = FormCallback[*models.NoteProductShape]

func saveNoteProductShapeFields(
	_instance *models.NoteProductShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(_instance.Note), probe.stageOfInterest, formDiv)
		case "Product":
			FormDivSelectFieldToField(&(_instance.Product), probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(_instance.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(_instance.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(_instance.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(_instance.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(_instance.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "Diagram:NoteProductShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "NoteProductShapes", func(owner *models.Diagram) *[]*models.NoteProductShape { return &owner.NoteProductShapes })
		}
	}
}

func __gong__New__NoteResourceShapeFormCallback(
	_instance *models.NoteResourceShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (noteresourceshapeFormCallback *FormCallback[*models.NoteResourceShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNoteResourceShapeFields,
	)
}

type NoteResourceShapeFormCallback = FormCallback[*models.NoteResourceShape]

func saveNoteResourceShapeFields(
	_instance *models.NoteResourceShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(_instance.Note), probe.stageOfInterest, formDiv)
		case "Resource":
			FormDivSelectFieldToField(&(_instance.Resource), probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(_instance.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(_instance.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(_instance.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(_instance.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(_instance.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "Diagram:NoteResourceShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "NoteResourceShapes", func(owner *models.Diagram) *[]*models.NoteResourceShape { return &owner.NoteResourceShapes })
		}
	}
}

func __gong__New__NoteShapeFormCallback(
	_instance *models.NoteShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (noteshapeFormCallback *FormCallback[*models.NoteShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNoteShapeFields,
	)
}

type NoteShapeFormCallback = FormCallback[*models.NoteShape]

func saveNoteShapeFields(
	_instance *models.NoteShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(_instance.Note), probe.stageOfInterest, formDiv)
		case "OverideLayoutDirection":
			FormDivBasicFieldToField(&(_instance.OverideLayoutDirection), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(_instance.LayoutDirection), formDiv)
		case "X":
			FormDivBasicFieldToField(&(_instance.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(_instance.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "Diagram:Note_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Note_Shapes", func(owner *models.Diagram) *[]*models.NoteShape { return &owner.Note_Shapes })
		}
	}
}

func __gong__New__NoteTaskShapeFormCallback(
	_instance *models.NoteTaskShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (notetaskshapeFormCallback *FormCallback[*models.NoteTaskShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNoteTaskShapeFields,
	)
}

type NoteTaskShapeFormCallback = FormCallback[*models.NoteTaskShape]

func saveNoteTaskShapeFields(
	_instance *models.NoteTaskShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Note":
			FormDivSelectFieldToField(&(_instance.Note), probe.stageOfInterest, formDiv)
		case "Task":
			FormDivSelectFieldToField(&(_instance.Task), probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(_instance.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(_instance.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(_instance.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(_instance.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(_instance.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "Diagram:NoteTaskShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "NoteTaskShapes", func(owner *models.Diagram) *[]*models.NoteTaskShape { return &owner.NoteTaskShapes })
		}
	}
}

func __gong__New__ProductFormCallback(
	_instance *models.Product,
	probe *Probe,
	formGroup *form.FormGroup,
) (productFormCallback *FormCallback[*models.Product]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveProductFields,
	)
}

type ProductFormCallback = FormCallback[*models.Product]

func saveProductFields(
	_instance *models.Product,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(_instance.Description), formDiv)
		case "SubProducts":
			FormDivSliceOfPointersToField(_instance, "SubProducts", &(_instance.SubProducts), formDiv, probe)
		case "IsProducersNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsProducersNodeExpanded), formDiv)
		case "IsConsumersNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsConsumersNodeExpanded), formDiv)
		case "IsImport":
			FormDivBasicFieldToField(&(_instance.IsImport), formDiv)
		case "ReferencedProduct":
			FormDivSelectFieldToField(&(_instance.ReferencedProduct), probe.stageOfInterest, formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(_instance.LayoutDirection), formDiv)
		case "Diagram:ProductsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ProductsWhoseNodeIsExpanded", func(owner *models.Diagram) *[]*models.Product { return &owner.ProductsWhoseNodeIsExpanded })
		case "Library:RootProducts":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootProducts", func(owner *models.Library) *[]*models.Product { return &owner.RootProducts })
		case "Note:Products":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Products", func(owner *models.Note) *[]*models.Product { return &owner.Products })
		case "Product:SubProducts":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubProducts", func(owner *models.Product) *[]*models.Product { return &owner.SubProducts })
		case "Task:Inputs":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Inputs", func(owner *models.Task) *[]*models.Product { return &owner.Inputs })
		case "Task:Outputs":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Outputs", func(owner *models.Task) *[]*models.Product { return &owner.Outputs })
		}
	}
}

func __gong__New__ProductCompositionShapeFormCallback(
	_instance *models.ProductCompositionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (productcompositionshapeFormCallback *FormCallback[*models.ProductCompositionShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveProductCompositionShapeFields,
	)
}

type ProductCompositionShapeFormCallback = FormCallback[*models.ProductCompositionShape]

func saveProductCompositionShapeFields(
	_instance *models.ProductCompositionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Product":
			FormDivSelectFieldToField(&(_instance.Product), probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(_instance.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(_instance.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(_instance.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(_instance.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(_instance.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "Diagram:ProductComposition_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ProductComposition_Shapes", func(owner *models.Diagram) *[]*models.ProductCompositionShape { return &owner.ProductComposition_Shapes })
		}
	}
}

func __gong__New__ProductReferenceShapeFormCallback(
	_instance *models.ProductReferenceShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (productreferenceshapeFormCallback *FormCallback[*models.ProductReferenceShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveProductReferenceShapeFields,
	)
}

type ProductReferenceShapeFormCallback = FormCallback[*models.ProductReferenceShape]

func saveProductReferenceShapeFields(
	_instance *models.ProductReferenceShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Product":
			FormDivSelectFieldToField(&(_instance.Product), probe.stageOfInterest, formDiv)
		case "ReferencedProduct":
			FormDivSelectFieldToField(&(_instance.ReferencedProduct), probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(_instance.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(_instance.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(_instance.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(_instance.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(_instance.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "Diagram:ProductReference_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ProductReference_Shapes", func(owner *models.Diagram) *[]*models.ProductReferenceShape { return &owner.ProductReference_Shapes })
		}
	}
}

func __gong__New__ProductShapeFormCallback(
	_instance *models.ProductShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (productshapeFormCallback *FormCallback[*models.ProductShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveProductShapeFields,
	)
}

type ProductShapeFormCallback = FormCallback[*models.ProductShape]

func saveProductShapeFields(
	_instance *models.ProductShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Product":
			FormDivSelectFieldToField(&(_instance.Product), probe.stageOfInterest, formDiv)
		case "IsShowType":
			FormDivBasicFieldToField(&(_instance.IsShowType), formDiv)
		case "OverideLayoutDirection":
			FormDivBasicFieldToField(&(_instance.OverideLayoutDirection), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(_instance.LayoutDirection), formDiv)
		case "X":
			FormDivBasicFieldToField(&(_instance.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(_instance.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "Diagram:Product_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Product_Shapes", func(owner *models.Diagram) *[]*models.ProductShape { return &owner.Product_Shapes })
		}
	}
}

func __gong__New__ResourceFormCallback(
	_instance *models.Resource,
	probe *Probe,
	formGroup *form.FormGroup,
) (resourceFormCallback *FormCallback[*models.Resource]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveResourceFields,
	)
}

type ResourceFormCallback = FormCallback[*models.Resource]

func saveResourceFields(
	_instance *models.Resource,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(_instance.Description), formDiv)
		case "Tasks":
			FormDivSliceOfPointersToField(_instance, "Tasks", &(_instance.Tasks), formDiv, probe)
		case "SubResources":
			FormDivSliceOfPointersToField(_instance, "SubResources", &(_instance.SubResources), formDiv, probe)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(_instance.LayoutDirection), formDiv)
		case "IsImport":
			FormDivBasicFieldToField(&(_instance.IsImport), formDiv)
		case "ReferencedResource":
			FormDivSelectFieldToField(&(_instance.ReferencedResource), probe.stageOfInterest, formDiv)
		case "Diagram:ResourcesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ResourcesWhoseNodeIsExpanded", func(owner *models.Diagram) *[]*models.Resource { return &owner.ResourcesWhoseNodeIsExpanded })
		case "Library:RootResources":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootResources", func(owner *models.Library) *[]*models.Resource { return &owner.RootResources })
		case "Note:Resources":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Resources", func(owner *models.Note) *[]*models.Resource { return &owner.Resources })
		case "Resource:SubResources":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubResources", func(owner *models.Resource) *[]*models.Resource { return &owner.SubResources })
		}
	}
}

func __gong__New__ResourceCompositionShapeFormCallback(
	_instance *models.ResourceCompositionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (resourcecompositionshapeFormCallback *FormCallback[*models.ResourceCompositionShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveResourceCompositionShapeFields,
	)
}

type ResourceCompositionShapeFormCallback = FormCallback[*models.ResourceCompositionShape]

func saveResourceCompositionShapeFields(
	_instance *models.ResourceCompositionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Resource":
			FormDivSelectFieldToField(&(_instance.Resource), probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(_instance.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(_instance.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(_instance.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(_instance.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(_instance.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "Diagram:ResourceComposition_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ResourceComposition_Shapes", func(owner *models.Diagram) *[]*models.ResourceCompositionShape { return &owner.ResourceComposition_Shapes })
		}
	}
}

func __gong__New__ResourceShapeFormCallback(
	_instance *models.ResourceShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (resourceshapeFormCallback *FormCallback[*models.ResourceShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveResourceShapeFields,
	)
}

type ResourceShapeFormCallback = FormCallback[*models.ResourceShape]

func saveResourceShapeFields(
	_instance *models.ResourceShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Resource":
			FormDivSelectFieldToField(&(_instance.Resource), probe.stageOfInterest, formDiv)
		case "OverideLayoutDirection":
			FormDivBasicFieldToField(&(_instance.OverideLayoutDirection), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(_instance.LayoutDirection), formDiv)
		case "X":
			FormDivBasicFieldToField(&(_instance.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(_instance.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "Diagram:Resource_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Resource_Shapes", func(owner *models.Diagram) *[]*models.ResourceShape { return &owner.Resource_Shapes })
		}
	}
}

func __gong__New__ResourceTaskShapeFormCallback(
	_instance *models.ResourceTaskShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (resourcetaskshapeFormCallback *FormCallback[*models.ResourceTaskShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveResourceTaskShapeFields,
	)
}

type ResourceTaskShapeFormCallback = FormCallback[*models.ResourceTaskShape]

func saveResourceTaskShapeFields(
	_instance *models.ResourceTaskShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Resource":
			FormDivSelectFieldToField(&(_instance.Resource), probe.stageOfInterest, formDiv)
		case "Task":
			FormDivSelectFieldToField(&(_instance.Task), probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(_instance.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(_instance.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(_instance.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(_instance.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(_instance.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "Diagram:ResourceTaskShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ResourceTaskShapes", func(owner *models.Diagram) *[]*models.ResourceTaskShape { return &owner.ResourceTaskShapes })
		}
	}
}

func __gong__New__TaskFormCallback(
	_instance *models.Task,
	probe *Probe,
	formGroup *form.FormGroup,
) (taskFormCallback *FormCallback[*models.Task]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTaskFields,
	)
}

type TaskFormCallback = FormCallback[*models.Task]

func saveTaskFields(
	_instance *models.Task,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(_instance.Description), formDiv)
		case "Start":
			FormDivTimeFieldToField(&(_instance.Start), formDiv, true)
		case "End":
			FormDivTimeFieldToField(&(_instance.End), formDiv, true)
		case "IsMilestone":
			FormDivBasicFieldToField(&(_instance.IsMilestone), formDiv)
		case "Predecessors":
			FormDivSliceOfPointersToField(_instance, "Predecessors", &(_instance.Predecessors), formDiv, probe)
		case "DependencyType":
			FormDivEnumStringFieldToField(&(_instance.DependencyType), formDiv)
		case "DependencyDurationYears":
			FormDivBasicFieldToField(&(_instance.DependencyDurationYears), formDiv)
		case "DependencyDurationMonths":
			FormDivBasicFieldToField(&(_instance.DependencyDurationMonths), formDiv)
		case "DependencyDurationWeeks":
			FormDivBasicFieldToField(&(_instance.DependencyDurationWeeks), formDiv)
		case "DependencyDurationDays":
			FormDivBasicFieldToField(&(_instance.DependencyDurationDays), formDiv)
		case "DependencyDurationHours":
			FormDivBasicFieldToField(&(_instance.DependencyDurationHours), formDiv)
		case "DurationYears":
			FormDivBasicFieldToField(&(_instance.DurationYears), formDiv)
		case "DurationMonths":
			FormDivBasicFieldToField(&(_instance.DurationMonths), formDiv)
		case "DurationWeeks":
			FormDivBasicFieldToField(&(_instance.DurationWeeks), formDiv)
		case "DurationDays":
			FormDivBasicFieldToField(&(_instance.DurationDays), formDiv)
		case "DurationHours":
			FormDivBasicFieldToField(&(_instance.DurationHours), formDiv)
		case "IsEndDateComputedFromDuration":
			FormDivBasicFieldToField(&(_instance.IsEndDateComputedFromDuration), formDiv)
		case "Inputs":
			FormDivSliceOfPointersToField(_instance, "Inputs", &(_instance.Inputs), formDiv, probe)
		case "Outputs":
			FormDivSliceOfPointersToField(_instance, "Outputs", &(_instance.Outputs), formDiv, probe)
		case "SubTasks":
			FormDivSliceOfPointersToField(_instance, "SubTasks", &(_instance.SubTasks), formDiv, probe)
		case "IsWithCompletion":
			FormDivBasicFieldToField(&(_instance.IsWithCompletion), formDiv)
		case "Completion":
			FormDivEnumStringFieldToField(&(_instance.Completion), formDiv)
		case "DisplayVerticalBar":
			FormDivBasicFieldToField(&(_instance.DisplayVerticalBar), formDiv)
		case "TaskGroupsToDisplay":
			FormDivSliceOfPointersToField(_instance, "TaskGroupsToDisplay", &(_instance.TaskGroupsToDisplay), formDiv, probe)
		case "TextPosition":
			FormDivEnumStringFieldToField(&(_instance.TextPosition), formDiv)
		case "XOffset":
			FormDivBasicFieldToField(&(_instance.XOffset), formDiv)
		case "YOffset":
			FormDivBasicFieldToField(&(_instance.YOffset), formDiv)
		case "IsImport":
			FormDivBasicFieldToField(&(_instance.IsImport), formDiv)
		case "ReferencedTask":
			FormDivSelectFieldToField(&(_instance.ReferencedTask), probe.stageOfInterest, formDiv)
		case "IsInputsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsInputsNodeExpanded), formDiv)
		case "IsOutputsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsOutputsNodeExpanded), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(_instance.LayoutDirection), formDiv)
		case "Diagram:TasksWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TasksWhoseNodeIsExpanded", func(owner *models.Diagram) *[]*models.Task { return &owner.TasksWhoseNodeIsExpanded })
		case "Diagram:TasksWhoseInputNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TasksWhoseInputNodeIsExpanded", func(owner *models.Diagram) *[]*models.Task { return &owner.TasksWhoseInputNodeIsExpanded })
		case "Diagram:TasksWhoseOutputNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TasksWhoseOutputNodeIsExpanded", func(owner *models.Diagram) *[]*models.Task { return &owner.TasksWhoseOutputNodeIsExpanded })
		case "Diagram:TasksWhosePredecessorNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TasksWhosePredecessorNodeIsExpanded", func(owner *models.Diagram) *[]*models.Task { return &owner.TasksWhosePredecessorNodeIsExpanded })
		case "Library:RootTasks":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootTasks", func(owner *models.Library) *[]*models.Task { return &owner.RootTasks })
		case "Note:Tasks":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Tasks", func(owner *models.Note) *[]*models.Task { return &owner.Tasks })
		case "Resource:Tasks":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Tasks", func(owner *models.Resource) *[]*models.Task { return &owner.Tasks })
		case "Task:Predecessors":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Predecessors", func(owner *models.Task) *[]*models.Task { return &owner.Predecessors })
		case "Task:SubTasks":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubTasks", func(owner *models.Task) *[]*models.Task { return &owner.SubTasks })
		case "TaskGroup:Tasks":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Tasks", func(owner *models.TaskGroup) *[]*models.Task { return &owner.Tasks })
		}
	}
}

func __gong__New__TaskCompositionShapeFormCallback(
	_instance *models.TaskCompositionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (taskcompositionshapeFormCallback *FormCallback[*models.TaskCompositionShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTaskCompositionShapeFields,
	)
}

type TaskCompositionShapeFormCallback = FormCallback[*models.TaskCompositionShape]

func saveTaskCompositionShapeFields(
	_instance *models.TaskCompositionShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Task":
			FormDivSelectFieldToField(&(_instance.Task), probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(_instance.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(_instance.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(_instance.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(_instance.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(_instance.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "Diagram:TaskComposition_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TaskComposition_Shapes", func(owner *models.Diagram) *[]*models.TaskCompositionShape { return &owner.TaskComposition_Shapes })
		}
	}
}

func __gong__New__TaskGroupFormCallback(
	_instance *models.TaskGroup,
	probe *Probe,
	formGroup *form.FormGroup,
) (taskgroupFormCallback *FormCallback[*models.TaskGroup]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTaskGroupFields,
	)
}

type TaskGroupFormCallback = FormCallback[*models.TaskGroup]

func saveTaskGroupFields(
	_instance *models.TaskGroup,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "Tasks":
			FormDivSliceOfPointersToField(_instance, "Tasks", &(_instance.Tasks), formDiv, probe)
		case "Diagram:TaskGroupsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TaskGroupsWhoseNodeIsExpanded", func(owner *models.Diagram) *[]*models.TaskGroup { return &owner.TaskGroupsWhoseNodeIsExpanded })
		case "Library:RootTaskGroups":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootTaskGroups", func(owner *models.Library) *[]*models.TaskGroup { return &owner.RootTaskGroups })
		case "Task:TaskGroupsToDisplay":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TaskGroupsToDisplay", func(owner *models.Task) *[]*models.TaskGroup { return &owner.TaskGroupsToDisplay })
		}
	}
}

func __gong__New__TaskGroupShapeFormCallback(
	_instance *models.TaskGroupShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (taskgroupshapeFormCallback *FormCallback[*models.TaskGroupShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTaskGroupShapeFields,
	)
}

type TaskGroupShapeFormCallback = FormCallback[*models.TaskGroupShape]

func saveTaskGroupShapeFields(
	_instance *models.TaskGroupShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "TaskGroup":
			FormDivSelectFieldToField(&(_instance.TaskGroup), probe.stageOfInterest, formDiv)
		case "X":
			FormDivBasicFieldToField(&(_instance.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(_instance.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "Diagram:TaskGroupShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TaskGroupShapes", func(owner *models.Diagram) *[]*models.TaskGroupShape { return &owner.TaskGroupShapes })
		}
	}
}

func __gong__New__TaskInputShapeFormCallback(
	_instance *models.TaskInputShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (taskinputshapeFormCallback *FormCallback[*models.TaskInputShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTaskInputShapeFields,
	)
}

type TaskInputShapeFormCallback = FormCallback[*models.TaskInputShape]

func saveTaskInputShapeFields(
	_instance *models.TaskInputShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Product":
			FormDivSelectFieldToField(&(_instance.Product), probe.stageOfInterest, formDiv)
		case "Task":
			FormDivSelectFieldToField(&(_instance.Task), probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(_instance.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(_instance.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(_instance.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(_instance.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(_instance.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "Diagram:TaskInputShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TaskInputShapes", func(owner *models.Diagram) *[]*models.TaskInputShape { return &owner.TaskInputShapes })
		}
	}
}

func __gong__New__TaskOutputShapeFormCallback(
	_instance *models.TaskOutputShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (taskoutputshapeFormCallback *FormCallback[*models.TaskOutputShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTaskOutputShapeFields,
	)
}

type TaskOutputShapeFormCallback = FormCallback[*models.TaskOutputShape]

func saveTaskOutputShapeFields(
	_instance *models.TaskOutputShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Task":
			FormDivSelectFieldToField(&(_instance.Task), probe.stageOfInterest, formDiv)
		case "Product":
			FormDivSelectFieldToField(&(_instance.Product), probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(_instance.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(_instance.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(_instance.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(_instance.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(_instance.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "Diagram:TaskOutputShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TaskOutputShapes", func(owner *models.Diagram) *[]*models.TaskOutputShape { return &owner.TaskOutputShapes })
		}
	}
}

func __gong__New__TaskPredecessorShapeFormCallback(
	_instance *models.TaskPredecessorShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (taskpredecessorshapeFormCallback *FormCallback[*models.TaskPredecessorShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTaskPredecessorShapeFields,
	)
}

type TaskPredecessorShapeFormCallback = FormCallback[*models.TaskPredecessorShape]

func saveTaskPredecessorShapeFields(
	_instance *models.TaskPredecessorShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Predecessor":
			FormDivSelectFieldToField(&(_instance.Predecessor), probe.stageOfInterest, formDiv)
		case "Task":
			FormDivSelectFieldToField(&(_instance.Task), probe.stageOfInterest, formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(_instance.StartRatio), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(_instance.EndRatio), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(_instance.StartOrientation), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(_instance.EndOrientation), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(_instance.CornerOffsetRatio), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "Diagram:TaskPredecessorShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TaskPredecessorShapes", func(owner *models.Diagram) *[]*models.TaskPredecessorShape { return &owner.TaskPredecessorShapes })
		}
	}
}

func __gong__New__TaskShapeFormCallback(
	_instance *models.TaskShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (taskshapeFormCallback *FormCallback[*models.TaskShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTaskShapeFields,
	)
}

type TaskShapeFormCallback = FormCallback[*models.TaskShape]

func saveTaskShapeFields(
	_instance *models.TaskShape,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Task":
			FormDivSelectFieldToField(&(_instance.Task), probe.stageOfInterest, formDiv)
		case "IsShowDate":
			FormDivBasicFieldToField(&(_instance.IsShowDate), formDiv)
		case "OverideLayoutDirection":
			FormDivBasicFieldToField(&(_instance.OverideLayoutDirection), formDiv)
		case "LayoutDirection":
			FormDivEnumIntFieldToField(&(_instance.LayoutDirection), formDiv)
		case "X":
			FormDivBasicFieldToField(&(_instance.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(_instance.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "IsHidden":
			FormDivBasicFieldToField(&(_instance.IsHidden), formDiv)
		case "Diagram:Task_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Task_Shapes", func(owner *models.Diagram) *[]*models.TaskShape { return &owner.Task_Shapes })
		}
	}
}

