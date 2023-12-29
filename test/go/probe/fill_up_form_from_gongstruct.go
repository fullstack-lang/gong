// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/test/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	FillUpNamedFormFromGongstruct[T](instance, probe, formStage, gongtable.FormGroupDefaultName.ToString())

}

func FillUpNamedFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe, formStage *gongtable.StageStruct, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Astruct:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Astruct Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AstructFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.AstructBstruct2Use:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "AstructBstruct2Use Form",
			OnSave: __gong__New__AstructBstruct2UseFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.AstructBstructUse:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "AstructBstructUse Form",
			OnSave: __gong__New__AstructBstructUseFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Bstruct:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Bstruct Form",
			OnSave: __gong__New__BstructFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Dstruct:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Dstruct Form",
			OnSave: __gong__New__DstructFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Fstruct:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Fstruct Form",
			OnSave: __gong__New__FstructFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
