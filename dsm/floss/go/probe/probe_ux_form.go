// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/floss/go/models"
)

// ux_form updates the current form if there is one
func (probe *Probe) ux_form() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *CompareAnalysisFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "CompareAnalysis", true)
			} else {
				FillUpFormFromGongstruct(onSave.compareanalysis, probe)
			}
		case *ComplexityFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Complexity", true)
			} else {
				FillUpFormFromGongstruct(onSave.complexity, probe)
			}
		case *DiagramFlossEquationFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "DiagramFlossEquation", true)
			} else {
				FillUpFormFromGongstruct(onSave.diagramflossequation, probe)
			}
		case *EffortFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Effort", true)
			} else {
				FillUpFormFromGongstruct(onSave.effort, probe)
			}
		case *LibraryFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Library", true)
			} else {
				FillUpFormFromGongstruct(onSave.library, probe)
			}
		case *NoteFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Note", true)
			} else {
				FillUpFormFromGongstruct(onSave.note, probe)
			}
		case *NoteComplexityShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "NoteComplexityShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.notecomplexityshape, probe)
			}
		case *NoteEffortShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "NoteEffortShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.noteeffortshape, probe)
			}
		case *NotePerformanceShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "NotePerformanceShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.noteperformanceshape, probe)
			}
		case *NoteShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "NoteShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.noteshape, probe)
			}
		case *PerformanceFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Performance", true)
			} else {
				FillUpFormFromGongstruct(onSave.performance, probe)
			}
		case *SystemFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "System", true)
			} else {
				FillUpFormFromGongstruct(onSave.system, probe)
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
	case "CompareAnalysis":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "CompareAnalysis Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CompareAnalysisFormCallback(
			nil,
			probe,
			formGroup,
		)
		compareanalysis := new(models.CompareAnalysis)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(compareanalysis, formGroup, probe)
	case "Complexity":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Complexity Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ComplexityFormCallback(
			nil,
			probe,
			formGroup,
		)
		complexity := new(models.Complexity)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(complexity, formGroup, probe)
	case "DiagramFlossEquation":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DiagramFlossEquation Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramFlossEquationFormCallback(
			nil,
			probe,
			formGroup,
		)
		diagramflossequation := new(models.DiagramFlossEquation)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(diagramflossequation, formGroup, probe)
	case "Effort":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Effort Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EffortFormCallback(
			nil,
			probe,
			formGroup,
		)
		effort := new(models.Effort)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(effort, formGroup, probe)
	case "Library":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Library Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LibraryFormCallback(
			nil,
			probe,
			formGroup,
		)
		library := new(models.Library)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(library, formGroup, probe)
	case "Note":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Note Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteFormCallback(
			nil,
			probe,
			formGroup,
		)
		note := new(models.Note)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(note, formGroup, probe)
	case "NoteComplexityShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "NoteComplexityShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteComplexityShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		notecomplexityshape := new(models.NoteComplexityShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(notecomplexityshape, formGroup, probe)
	case "NoteEffortShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "NoteEffortShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteEffortShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		noteeffortshape := new(models.NoteEffortShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(noteeffortshape, formGroup, probe)
	case "NotePerformanceShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "NotePerformanceShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NotePerformanceShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		noteperformanceshape := new(models.NotePerformanceShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(noteperformanceshape, formGroup, probe)
	case "NoteShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "NoteShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		noteshape := new(models.NoteShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(noteshape, formGroup, probe)
	case "Performance":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Performance Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PerformanceFormCallback(
			nil,
			probe,
			formGroup,
		)
		performance := new(models.Performance)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(performance, formGroup, probe)
	case "System":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "System Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SystemFormCallback(
			nil,
			probe,
			formGroup,
		)
		system := new(models.System)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(system, formGroup, probe)
	}
	formStage.Commit()
}
