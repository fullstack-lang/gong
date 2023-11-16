// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/test2/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.A:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update A Form",
			OnSave: __gong__New__AFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.B:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update B Form",
			OnSave: __gong__New__BFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
