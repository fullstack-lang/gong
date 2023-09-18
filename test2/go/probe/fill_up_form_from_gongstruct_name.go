// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/test2/go/models"
)

func FillUpFormFromGongstructName(
	playground *Playground,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := playground.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = "New"
	} else {
		prefix = "Update"
	}

	switch gongstructName {
	// insertion point
	case "Dummy":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Dummy Form",
			OnSave: NewDummyFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		dummy := new(models.Dummy)
		FillUpForm(dummy, formGroup, playground)
	}
	formStage.Commit()
}
