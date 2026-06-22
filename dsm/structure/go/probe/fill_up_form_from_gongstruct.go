// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/structure/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.AllocatedResourceShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "AllocatedResourceShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AllocatedResourceShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.AllocatedSystemShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "AllocatedSystemShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AllocatedSystemShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ControlFlow:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ControlFlow",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ControlFlowFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ControlFlowShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ControlFlowShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ControlFlowShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Data:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Data",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DataFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DataFlow:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "DataFlow",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DataFlowFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DataFlowShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "DataFlowShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DataFlowShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DataShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "DataShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DataShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DiagramStructure:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "DiagramStructure",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramStructureFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ExternalPartShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ExternalPartShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ExternalPartShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Library:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Library",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LibraryFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Note:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Note",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.NotePortShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "NotePortShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NotePortShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.NoteShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "NoteShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Part:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Part",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PartShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "PartShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PartShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Port:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Port",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PortFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PortShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "PortShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PortShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Resource:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Resource",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ResourceFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.System:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "System",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SystemFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SystemShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "SystemShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SystemShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}

	if probe.GetCommitMode() {
		formStage.Commit()
	}
}
