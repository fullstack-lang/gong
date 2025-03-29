// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/tree/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *gongtable.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Button:
		formGroup := (&gongtable.FormGroup{
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
	case *models.Node:
		formGroup := (&gongtable.FormGroup{
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
		formGroup := (&gongtable.FormGroup{
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
		formGroup := (&gongtable.FormGroup{
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
	formStage.Commit()
}
