// generated code - do not edit
package x

import (
	"github.com/fullstack-lang/gong/test2/go/models/x"
	form "github.com/fullstack-lang/gongtable/go/models"
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
	case "A":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " A Form",
			OnSave: __gong__New__AFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		a := new(x.A)
		FillUpForm(a, formGroup, playground)
	}
	formStage.Commit()
}
