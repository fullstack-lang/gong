// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/gantt/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Arrow:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Arrow",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArrowFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Bar:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Bar",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BarFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Gantt:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Gantt",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GanttFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Group:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Group",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GroupFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Lane:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Lane",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LaneFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.LaneUse:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "LaneUse",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LaneUseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Milestone:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Milestone",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MilestoneFormCallback(
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
