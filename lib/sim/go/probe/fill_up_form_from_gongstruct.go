// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/sim/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *gongtable.StageStruct, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Command:
		formGroup := (&gongtable.FormGroup{
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
		formGroup := (&gongtable.FormGroup{
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
		formGroup := (&gongtable.FormGroup{
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
		formGroup := (&gongtable.FormGroup{
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
		formGroup := (&gongtable.FormGroup{
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
		formGroup := (&gongtable.FormGroup{
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
	formStage.Commit()
}
