// generated code - do not edit
package x

import (
	"github.com/fullstack-lang/gong/test2/go/models/x"
	gongtable "github.com/fullstack-lang/gongtable/go/models"
)

func FillUpFormFromGongstruct[T x.Gongstruct](instance *T, playground *Playground) {
	formStage := playground.formStage
	formStage.Reset()
	formStage.Commit()

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *x.A:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update A Form",
			OnSave: __gong__New__AFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
