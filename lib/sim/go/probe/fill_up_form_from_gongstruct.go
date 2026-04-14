// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/sim/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Command:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Command Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CommandFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DummyAgent:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "DummyAgent Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DummyAgentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Engine:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Engine Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EngineFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Event:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Event Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EventFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Status:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Status Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StatusFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.UpdateState:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "UpdateState Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__UpdateStateFormCallback(
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
