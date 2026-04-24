// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/process/go/models"
)

// ux_form updates the current form if there is one
func (probe *Probe) ux_form() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *DiagramFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Diagram", true)
			} else {
				FillUpFormFromGongstruct(onSave.diagram, probe)
			}
		case *LibraryFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Library", true)
			} else {
				FillUpFormFromGongstruct(onSave.library, probe)
			}
		case *ProcessFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Process", true)
			} else {
				FillUpFormFromGongstruct(onSave.process, probe)
			}
		case *ProcessCompositionShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ProcessCompositionShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.processcompositionshape, probe)
			}
		case *ProcessShapeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ProcessShape", true)
			} else {
				FillUpFormFromGongstruct(onSave.processshape, probe)
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
	case "Diagram":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Diagram Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramFormCallback(
			nil,
			probe,
			formGroup,
		)
		diagram := new(models.Diagram)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(diagram, formGroup, probe)
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
	case "Process":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Process Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ProcessFormCallback(
			nil,
			probe,
			formGroup,
		)
		process := new(models.Process)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(process, formGroup, probe)
	case "ProcessCompositionShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ProcessCompositionShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ProcessCompositionShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		processcompositionshape := new(models.ProcessCompositionShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(processcompositionshape, formGroup, probe)
	case "ProcessShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ProcessShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ProcessShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		processshape := new(models.ProcessShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(processshape, formGroup, probe)
	}
	formStage.Commit()
}
