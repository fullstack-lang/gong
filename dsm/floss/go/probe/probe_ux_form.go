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
		case *ComplexityFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Complexity", true)
			} else {
				FillUpFormFromGongstruct(onSave.complexity, probe)
			}
		case *DiagramFlossFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "DiagramFloss", true)
			} else {
				FillUpFormFromGongstruct(onSave.diagramfloss, probe)
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
		case *SystemShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "SystemShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.systemshape, probe)
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
	case "DiagramFloss":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DiagramFloss Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramFlossFormCallback(
			nil,
			probe,
			formGroup,
		)
		diagramfloss := new(models.DiagramFloss)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(diagramfloss, formGroup, probe)
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
	case "SystemShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SystemShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SystemShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		systemshape := new(models.SystemShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(systemshape, formGroup, probe)
	}
	formStage.Commit()
}
