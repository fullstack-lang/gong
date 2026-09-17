// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/floss/go/models"
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
	return &FormCallback[T]{
		Instance:     instance,
		CreationMode: any(instance) == nil,
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

	if any(cb.Instance) == nil {
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
func __gong__New__CompareAnalysisFormCallback(
	_instance *models.CompareAnalysis,
	probe *Probe,
	formGroup *form.FormGroup,
) (compareanalysisFormCallback *FormCallback[*models.CompareAnalysis]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCompareAnalysisFields,
	)
}

type CompareAnalysisFormCallback = FormCallback[*models.CompareAnalysis]

func saveCompareAnalysisFields(
	_instance *models.CompareAnalysis,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "FromSystem":
			FormDivSelectFieldToField(&(_instance.FromSystem), probe.stageOfInterest, formDiv)
		case "ToSystem":
			FormDivSelectFieldToField(&(_instance.ToSystem), probe.stageOfInterest, formDiv)
		case "Mu":
			FormDivBasicFieldToField(&(_instance.Mu), formDiv)
		case "Epsilon":
			FormDivBasicFieldToField(&(_instance.Epsilon), formDiv)
		case "DiagramFlossEquations":
			FormDivSliceOfPointersToField(_instance, "DiagramFlossEquations", &(_instance.DiagramFlossEquations), formDiv, probe)
		case "DiagramFlossEquationsWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "DiagramFlossEquationsWhoseNodeIsExpanded", &(_instance.DiagramFlossEquationsWhoseNodeIsExpanded), formDiv, probe)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "Library:RootCompareAnalysis":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootCompareAnalysis", func(owner *models.Library) *[]*models.CompareAnalysis { return &owner.RootCompareAnalysis })
		case "Library:CompareAnalysisWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "CompareAnalysisWhoseNodeIsExpanded", func(owner *models.Library) *[]*models.CompareAnalysis { return &owner.CompareAnalysisWhoseNodeIsExpanded })
		}
	}
}

func __gong__New__ComplexityFormCallback(
	_instance *models.Complexity,
	probe *Probe,
	formGroup *form.FormGroup,
) (complexityFormCallback *FormCallback[*models.Complexity]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveComplexityFields,
	)
}

type ComplexityFormCallback = FormCallback[*models.Complexity]

func saveComplexityFields(
	_instance *models.Complexity,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Strength":
			FormDivBasicFieldToField(&(_instance.Strength), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(_instance.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "DiagramFlossEquation:ComplexitysWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ComplexitysWhoseNodeIsExpanded", func(owner *models.DiagramFlossEquation) *[]*models.Complexity { return &owner.ComplexitysWhoseNodeIsExpanded })
		case "Library:RootComplexitys":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootComplexitys", func(owner *models.Library) *[]*models.Complexity { return &owner.RootComplexitys })
		case "Library:ComplexitysWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ComplexitysWhoseNodeIsExpanded", func(owner *models.Library) *[]*models.Complexity { return &owner.ComplexitysWhoseNodeIsExpanded })
		case "Note:Complexities":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Complexities", func(owner *models.Note) *[]*models.Complexity { return &owner.Complexities })
		case "System:Complexities":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Complexities", func(owner *models.System) *[]*models.Complexity { return &owner.Complexities })
		case "System:ComplexitysWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "ComplexitysWhoseNodeIsExpanded", func(owner *models.System) *[]*models.Complexity { return &owner.ComplexitysWhoseNodeIsExpanded })
		}
	}
}

func __gong__New__DiagramFlossEquationFormCallback(
	_instance *models.DiagramFlossEquation,
	probe *Probe,
	formGroup *form.FormGroup,
) (diagramflossequationFormCallback *FormCallback[*models.DiagramFlossEquation]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDiagramFlossEquationFields,
	)
}

type DiagramFlossEquationFormCallback = FormCallback[*models.DiagramFlossEquation]

func saveDiagramFlossEquationFields(
	_instance *models.DiagramFlossEquation,
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
		case "Scale":
			FormDivBasicFieldToField(&(_instance.Scale), formDiv)
		case "FontSize":
			FormDivEnumStringFieldToField(&(_instance.FontSize), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(_instance.IsChecked), formDiv)
		case "IsEditable_":
			FormDivBasicFieldToField(&(_instance.IsEditable_), formDiv)
		case "IsInDelta3ColumnsMode":
			FormDivBasicFieldToField(&(_instance.IsInDelta3ColumnsMode), formDiv)
		case "AreQuantitativeElementsVisible":
			FormDivBasicFieldToField(&(_instance.AreQuantitativeElementsVisible), formDiv)
		case "AreSubsystemsVisible":
			FormDivBasicFieldToField(&(_instance.AreSubsystemsVisible), formDiv)
		case "AreCommonElementsHidden":
			FormDivBasicFieldToField(&(_instance.AreCommonElementsHidden), formDiv)
		case "AreCPEArrowsVisible":
			FormDivBasicFieldToField(&(_instance.AreCPEArrowsVisible), formDiv)
		case "AreColumnTitlesVisible":
			FormDivBasicFieldToField(&(_instance.AreColumnTitlesVisible), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(_instance.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(_instance.Height), formDiv)
		case "DefaultBoxWidth":
			FormDivBasicFieldToField(&(_instance.DefaultBoxWidth), formDiv)
		case "DefaultBoxHeigth":
			FormDivBasicFieldToField(&(_instance.DefaultBoxHeigth), formDiv)
		case "Note_Shapes":
			FormDivSliceOfPointersToField(_instance, "Note_Shapes", &(_instance.Note_Shapes), formDiv, probe)
		case "NoteComplexityShapes":
			FormDivSliceOfPointersToField(_instance, "NoteComplexityShapes", &(_instance.NoteComplexityShapes), formDiv, probe)
		case "NotePerformanceShapes":
			FormDivSliceOfPointersToField(_instance, "NotePerformanceShapes", &(_instance.NotePerformanceShapes), formDiv, probe)
		case "NoteEffortShapes":
			FormDivSliceOfPointersToField(_instance, "NoteEffortShapes", &(_instance.NoteEffortShapes), formDiv, probe)
		case "IsNotesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsNotesNodeExpanded), formDiv)
		case "NotesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "NotesWhoseNodeIsExpanded", &(_instance.NotesWhoseNodeIsExpanded), formDiv, probe)
		case "IsComplexitysNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsComplexitysNodeExpanded), formDiv)
		case "ComplexitysWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ComplexitysWhoseNodeIsExpanded", &(_instance.ComplexitysWhoseNodeIsExpanded), formDiv, probe)
		case "IsPerformancesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsPerformancesNodeExpanded), formDiv)
		case "PerformancesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "PerformancesWhoseNodeIsExpanded", &(_instance.PerformancesWhoseNodeIsExpanded), formDiv, probe)
		case "IsEffortsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsEffortsNodeExpanded), formDiv)
		case "EffortsWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "EffortsWhoseNodeIsExpanded", &(_instance.EffortsWhoseNodeIsExpanded), formDiv, probe)
		case "CompareAnalysis:DiagramFlossEquations":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DiagramFlossEquations", func(owner *models.CompareAnalysis) *[]*models.DiagramFlossEquation { return &owner.DiagramFlossEquations })
		case "CompareAnalysis:DiagramFlossEquationsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DiagramFlossEquationsWhoseNodeIsExpanded", func(owner *models.CompareAnalysis) *[]*models.DiagramFlossEquation { return &owner.DiagramFlossEquationsWhoseNodeIsExpanded })
		case "System:DiagramFlossEquations":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DiagramFlossEquations", func(owner *models.System) *[]*models.DiagramFlossEquation { return &owner.DiagramFlossEquations })
		case "System:DiagramFlossEquationsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "DiagramFlossEquationsWhoseNodeIsExpanded", func(owner *models.System) *[]*models.DiagramFlossEquation { return &owner.DiagramFlossEquationsWhoseNodeIsExpanded })
		}
	}
}

func __gong__New__EffortFormCallback(
	_instance *models.Effort,
	probe *Probe,
	formGroup *form.FormGroup,
) (effortFormCallback *FormCallback[*models.Effort]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveEffortFields,
	)
}

type EffortFormCallback = FormCallback[*models.Effort]

func saveEffortFields(
	_instance *models.Effort,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Strength":
			FormDivBasicFieldToField(&(_instance.Strength), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(_instance.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "DiagramFlossEquation:EffortsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "EffortsWhoseNodeIsExpanded", func(owner *models.DiagramFlossEquation) *[]*models.Effort { return &owner.EffortsWhoseNodeIsExpanded })
		case "Library:RootEfforts":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootEfforts", func(owner *models.Library) *[]*models.Effort { return &owner.RootEfforts })
		case "Library:EffortsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "EffortsWhoseNodeIsExpanded", func(owner *models.Library) *[]*models.Effort { return &owner.EffortsWhoseNodeIsExpanded })
		case "Note:Efforts":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Efforts", func(owner *models.Note) *[]*models.Effort { return &owner.Efforts })
		case "System:Efforts":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Efforts", func(owner *models.System) *[]*models.Effort { return &owner.Efforts })
		case "System:EffortsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "EffortsWhoseNodeIsExpanded", func(owner *models.System) *[]*models.Effort { return &owner.EffortsWhoseNodeIsExpanded })
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
		case "Description":
			FormDivBasicFieldToField(&(_instance.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "SubLibraries":
			FormDivSliceOfPointersToField(_instance, "SubLibraries", &(_instance.SubLibraries), formDiv, probe)
		case "RootSystems":
			FormDivSliceOfPointersToField(_instance, "RootSystems", &(_instance.RootSystems), formDiv, probe)
		case "RootComplexitys":
			FormDivSliceOfPointersToField(_instance, "RootComplexitys", &(_instance.RootComplexitys), formDiv, probe)
		case "RootPerformances":
			FormDivSliceOfPointersToField(_instance, "RootPerformances", &(_instance.RootPerformances), formDiv, probe)
		case "RootEfforts":
			FormDivSliceOfPointersToField(_instance, "RootEfforts", &(_instance.RootEfforts), formDiv, probe)
		case "RootCompareAnalysis":
			FormDivSliceOfPointersToField(_instance, "RootCompareAnalysis", &(_instance.RootCompareAnalysis), formDiv, probe)
		case "RootNotes":
			FormDivSliceOfPointersToField(_instance, "RootNotes", &(_instance.RootNotes), formDiv, probe)
		case "IsRootLibrary":
			FormDivBasicFieldToField(&(_instance.IsRootLibrary), formDiv)
		case "IsSubLibrariesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsSubLibrariesNodeExpanded), formDiv)
		case "SubLibrariesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "SubLibrariesWhoseNodeIsExpanded", &(_instance.SubLibrariesWhoseNodeIsExpanded), formDiv, probe)
		case "NbPixPerCharacter":
			FormDivBasicFieldToField(&(_instance.NbPixPerCharacter), formDiv)
		case "LogoSVGFile":
			FormDivBasicFieldToField(&(_instance.LogoSVGFile), formDiv)
		case "IsSystemsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsSystemsNodeExpanded), formDiv)
		case "SystemsWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "SystemsWhoseNodeIsExpanded", &(_instance.SystemsWhoseNodeIsExpanded), formDiv, probe)
		case "IsComplexitysNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsComplexitysNodeExpanded), formDiv)
		case "ComplexitysWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ComplexitysWhoseNodeIsExpanded", &(_instance.ComplexitysWhoseNodeIsExpanded), formDiv, probe)
		case "IsPerformancesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsPerformancesNodeExpanded), formDiv)
		case "PerformancesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "PerformancesWhoseNodeIsExpanded", &(_instance.PerformancesWhoseNodeIsExpanded), formDiv, probe)
		case "IsEffortsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsEffortsNodeExpanded), formDiv)
		case "EffortsWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "EffortsWhoseNodeIsExpanded", &(_instance.EffortsWhoseNodeIsExpanded), formDiv, probe)
		case "IsCompareAnalysisNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsCompareAnalysisNodeExpanded), formDiv)
		case "CompareAnalysisWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "CompareAnalysisWhoseNodeIsExpanded", &(_instance.CompareAnalysisWhoseNodeIsExpanded), formDiv, probe)
		case "IsNotesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsNotesNodeExpanded), formDiv)
		case "NotesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "NotesWhoseNodeIsExpanded", &(_instance.NotesWhoseNodeIsExpanded), formDiv, probe)
		case "IsExpandedTmp":
			FormDivBasicFieldToField(&(_instance.IsExpandedTmp), formDiv)
		case "Library:SubLibraries":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubLibraries", func(owner *models.Library) *[]*models.Library { return &owner.SubLibraries })
		case "Library:SubLibrariesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubLibrariesWhoseNodeIsExpanded", func(owner *models.Library) *[]*models.Library { return &owner.SubLibrariesWhoseNodeIsExpanded })
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
		case "Description":
			FormDivBasicFieldToField(&(_instance.Description), formDiv)
		case "Complexities":
			FormDivSliceOfPointersToField(_instance, "Complexities", &(_instance.Complexities), formDiv, probe)
		case "Performances":
			FormDivSliceOfPointersToField(_instance, "Performances", &(_instance.Performances), formDiv, probe)
		case "Efforts":
			FormDivSliceOfPointersToField(_instance, "Efforts", &(_instance.Efforts), formDiv, probe)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "IsComplexitysNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsComplexitysNodeExpanded), formDiv)
		case "IsPerformancesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsPerformancesNodeExpanded), formDiv)
		case "IsEffortsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsEffortsNodeExpanded), formDiv)
		case "DiagramFlossEquation:NotesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "NotesWhoseNodeIsExpanded", func(owner *models.DiagramFlossEquation) *[]*models.Note { return &owner.NotesWhoseNodeIsExpanded })
		case "Library:RootNotes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootNotes", func(owner *models.Library) *[]*models.Note { return &owner.RootNotes })
		case "Library:NotesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "NotesWhoseNodeIsExpanded", func(owner *models.Library) *[]*models.Note { return &owner.NotesWhoseNodeIsExpanded })
		}
	}
}

func __gong__New__NoteComplexityShapeFormCallback(
	_instance *models.NoteComplexityShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (notecomplexityshapeFormCallback *FormCallback[*models.NoteComplexityShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNoteComplexityShapeFields,
	)
}

type NoteComplexityShapeFormCallback = FormCallback[*models.NoteComplexityShape]

func saveNoteComplexityShapeFields(
	_instance *models.NoteComplexityShape,
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
		case "Complexity":
			FormDivSelectFieldToField(&(_instance.Complexity), probe.stageOfInterest, formDiv)
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
		case "DiagramFlossEquation:NoteComplexityShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "NoteComplexityShapes", func(owner *models.DiagramFlossEquation) *[]*models.NoteComplexityShape { return &owner.NoteComplexityShapes })
		}
	}
}

func __gong__New__NoteEffortShapeFormCallback(
	_instance *models.NoteEffortShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (noteeffortshapeFormCallback *FormCallback[*models.NoteEffortShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNoteEffortShapeFields,
	)
}

type NoteEffortShapeFormCallback = FormCallback[*models.NoteEffortShape]

func saveNoteEffortShapeFields(
	_instance *models.NoteEffortShape,
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
		case "Effort":
			FormDivSelectFieldToField(&(_instance.Effort), probe.stageOfInterest, formDiv)
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
		case "DiagramFlossEquation:NoteEffortShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "NoteEffortShapes", func(owner *models.DiagramFlossEquation) *[]*models.NoteEffortShape { return &owner.NoteEffortShapes })
		}
	}
}

func __gong__New__NotePerformanceShapeFormCallback(
	_instance *models.NotePerformanceShape,
	probe *Probe,
	formGroup *form.FormGroup,
) (noteperformanceshapeFormCallback *FormCallback[*models.NotePerformanceShape]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNotePerformanceShapeFields,
	)
}

type NotePerformanceShapeFormCallback = FormCallback[*models.NotePerformanceShape]

func saveNotePerformanceShapeFields(
	_instance *models.NotePerformanceShape,
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
		case "Performance":
			FormDivSelectFieldToField(&(_instance.Performance), probe.stageOfInterest, formDiv)
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
		case "DiagramFlossEquation:NotePerformanceShapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "NotePerformanceShapes", func(owner *models.DiagramFlossEquation) *[]*models.NotePerformanceShape { return &owner.NotePerformanceShapes })
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
		case "DiagramFlossEquation:Note_Shapes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Note_Shapes", func(owner *models.DiagramFlossEquation) *[]*models.NoteShape { return &owner.Note_Shapes })
		}
	}
}

func __gong__New__PerformanceFormCallback(
	_instance *models.Performance,
	probe *Probe,
	formGroup *form.FormGroup,
) (performanceFormCallback *FormCallback[*models.Performance]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		savePerformanceFields,
	)
}

type PerformanceFormCallback = FormCallback[*models.Performance]

func savePerformanceFields(
	_instance *models.Performance,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Strength":
			FormDivBasicFieldToField(&(_instance.Strength), formDiv)
		case "Description":
			FormDivBasicFieldToField(&(_instance.Description), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "DiagramFlossEquation:PerformancesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "PerformancesWhoseNodeIsExpanded", func(owner *models.DiagramFlossEquation) *[]*models.Performance { return &owner.PerformancesWhoseNodeIsExpanded })
		case "Library:RootPerformances":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootPerformances", func(owner *models.Library) *[]*models.Performance { return &owner.RootPerformances })
		case "Library:PerformancesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "PerformancesWhoseNodeIsExpanded", func(owner *models.Library) *[]*models.Performance { return &owner.PerformancesWhoseNodeIsExpanded })
		case "Note:Performances":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Performances", func(owner *models.Note) *[]*models.Performance { return &owner.Performances })
		case "System:Performances":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Performances", func(owner *models.System) *[]*models.Performance { return &owner.Performances })
		case "System:PerformancesWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "PerformancesWhoseNodeIsExpanded", func(owner *models.System) *[]*models.Performance { return &owner.PerformancesWhoseNodeIsExpanded })
		}
	}
}

func __gong__New__SystemFormCallback(
	_instance *models.System,
	probe *Probe,
	formGroup *form.FormGroup,
) (systemFormCallback *FormCallback[*models.System]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveSystemFields,
	)
}

type SystemFormCallback = FormCallback[*models.System]

func saveSystemFields(
	_instance *models.System,
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
		case "Complexities":
			FormDivSliceOfPointersToField(_instance, "Complexities", &(_instance.Complexities), formDiv, probe)
		case "Performances":
			FormDivSliceOfPointersToField(_instance, "Performances", &(_instance.Performances), formDiv, probe)
		case "Efforts":
			FormDivSliceOfPointersToField(_instance, "Efforts", &(_instance.Efforts), formDiv, probe)
		case "SubSystems":
			FormDivSliceOfPointersToField(_instance, "SubSystems", &(_instance.SubSystems), formDiv, probe)
		case "AreCPEsCompoundedFromSubSystems":
			FormDivBasicFieldToField(&(_instance.AreCPEsCompoundedFromSubSystems), formDiv)
		case "ComputedPrefix":
			FormDivBasicFieldToField(&(_instance.ComputedPrefix), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(_instance.IsExpanded), formDiv)
		case "SVG_Path":
			FormDivBasicFieldToField(&(_instance.SVG_Path), formDiv)
		case "InverseAppliedScaling":
			FormDivBasicFieldToField(&(_instance.InverseAppliedScaling), formDiv)
		case "DiagramFlossEquations":
			FormDivSliceOfPointersToField(_instance, "DiagramFlossEquations", &(_instance.DiagramFlossEquations), formDiv, probe)
		case "DiagramFlossEquationsWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "DiagramFlossEquationsWhoseNodeIsExpanded", &(_instance.DiagramFlossEquationsWhoseNodeIsExpanded), formDiv, probe)
		case "IsSubSystemNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsSubSystemNodeExpanded), formDiv)
		case "IsComplexitysNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsComplexitysNodeExpanded), formDiv)
		case "ComplexitysWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "ComplexitysWhoseNodeIsExpanded", &(_instance.ComplexitysWhoseNodeIsExpanded), formDiv, probe)
		case "IsPerformancesNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsPerformancesNodeExpanded), formDiv)
		case "PerformancesWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "PerformancesWhoseNodeIsExpanded", &(_instance.PerformancesWhoseNodeIsExpanded), formDiv, probe)
		case "IsEffortsNodeExpanded":
			FormDivBasicFieldToField(&(_instance.IsEffortsNodeExpanded), formDiv)
		case "EffortsWhoseNodeIsExpanded":
			FormDivSliceOfPointersToField(_instance, "EffortsWhoseNodeIsExpanded", &(_instance.EffortsWhoseNodeIsExpanded), formDiv, probe)
		case "Library:RootSystems":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "RootSystems", func(owner *models.Library) *[]*models.System { return &owner.RootSystems })
		case "Library:SystemsWhoseNodeIsExpanded":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SystemsWhoseNodeIsExpanded", func(owner *models.Library) *[]*models.System { return &owner.SystemsWhoseNodeIsExpanded })
		case "System:SubSystems":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "SubSystems", func(owner *models.System) *[]*models.System { return &owner.SubSystems })
		}
	}
}

