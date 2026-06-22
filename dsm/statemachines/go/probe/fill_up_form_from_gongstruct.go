// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/statemachines/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Action:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Action",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ActionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Activities:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Activities",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ActivitiesFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Architecture:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Architecture",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArchitectureFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Diagram:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Diagram",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Guard:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Guard",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GuardFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Kill:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Kill",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__KillFormCallback(
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
	case *models.Message:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Message",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MessageFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.MessageType:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "MessageType",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MessageTypeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Object:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Object",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ObjectFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Role:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Role",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RoleFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.State:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "State",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StateFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StateMachine:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "StateMachine",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StateMachineFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StateShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "StateShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StateShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Transition:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Transition",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TransitionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Transition_Shape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Transition_Shape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Transition_ShapeFormCallback(
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
