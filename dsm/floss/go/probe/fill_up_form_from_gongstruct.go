// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/floss/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.DiagramFloss:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "DiagramFloss",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramFlossFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Library:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "Library",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LibraryFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.System:
		formGroup := (&form.FormGroup{
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "System",
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
			Name:      formName,
			Label:     instancesTyped.GetName(),
			TypeLabel: "SystemShape",
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
