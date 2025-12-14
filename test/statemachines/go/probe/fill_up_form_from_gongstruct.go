// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/test/statemachines/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *gongtable.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Action:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Action Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ActionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Activities:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Activities Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ActivitiesFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Architecture:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Architecture Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArchitectureFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Diagram:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Diagram Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Guard:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Guard Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GuardFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Kill:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Kill Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__KillFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Message:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Message Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MessageFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.MessageType:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "MessageType Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MessageTypeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Object:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Object Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ObjectFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Role:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Role Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RoleFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.State:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "State Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StateFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StateMachine:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "StateMachine Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StateMachineFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.StateShape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "StateShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StateShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Transition:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Transition Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TransitionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Transition_Shape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Transition_Shape Form",
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
	formStage.Commit()
}
