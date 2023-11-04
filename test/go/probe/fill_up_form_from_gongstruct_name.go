// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/test/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
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
	case "Astruct":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Astruct Form",
			OnSave: __gong__New__AstructFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		astruct := new(models.Astruct)
		FillUpForm(astruct, formGroup, probe)
	case "AstructBstruct2Use":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " AstructBstruct2Use Form",
			OnSave: __gong__New__AstructBstruct2UseFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		astructbstruct2use := new(models.AstructBstruct2Use)
		FillUpForm(astructbstruct2use, formGroup, probe)
	case "AstructBstructUse":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " AstructBstructUse Form",
			OnSave: __gong__New__AstructBstructUseFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		astructbstructuse := new(models.AstructBstructUse)
		FillUpForm(astructbstructuse, formGroup, probe)
	case "Bstruct":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Bstruct Form",
			OnSave: __gong__New__BstructFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		bstruct := new(models.Bstruct)
		FillUpForm(bstruct, formGroup, probe)
	case "Dstruct":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Dstruct Form",
			OnSave: __gong__New__DstructFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		dstruct := new(models.Dstruct)
		FillUpForm(dstruct, formGroup, probe)
	case "Fstruct":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Fstruct Form",
			OnSave: __gong__New__FstructFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		fstruct := new(models.Fstruct)
		FillUpForm(fstruct, formGroup, probe)
	}
	formStage.Commit()
}
