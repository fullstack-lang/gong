// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/tree/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Button:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Button Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ButtonFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Menu:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Menu Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MenuFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Node:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Node Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NodeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SVGIcon:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "SVGIcon Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SVGIconFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Tree:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Tree Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TreeFormCallback(
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
