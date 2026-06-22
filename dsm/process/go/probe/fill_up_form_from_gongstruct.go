// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/process/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.AllocatedProcessShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "AllocatedProcessShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AllocatedProcessShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
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
	case *models.DiagramProcess:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "DiagramProcess",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramProcessFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ExternalParticipantShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ExternalParticipantShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ExternalParticipantShapeFormCallback(
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
	case *models.NoteTaskShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "NoteTaskShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteTaskShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Participant:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Participant",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParticipantFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ParticipantShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ParticipantShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParticipantShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Process:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Process",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ProcessFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ProcessShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "ProcessShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ProcessShapeFormCallback(
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
	case *models.Task:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Task",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TaskShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "TaskShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskShapeFormCallback(
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
