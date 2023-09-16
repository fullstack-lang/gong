package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/test/go/models"
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
	case "Astruct":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Astruct Form",
			OnSave: NewAstructFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		astruct := new(models.Astruct)
		FillUpForm(astruct, formGroup, playground)
	case "AstructBstruct2Use":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " AstructBstruct2Use Form",
			OnSave: NewAstructBstruct2UseFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		astructbstruct2use := new(models.AstructBstruct2Use)
		FillUpForm(astructbstruct2use, formGroup, playground)
	case "AstructBstructUse":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " AstructBstructUse Form",
			OnSave: NewAstructBstructUseFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		astructbstructuse := new(models.AstructBstructUse)
		FillUpForm(astructbstructuse, formGroup, playground)
	case "Bstruct":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Bstruct Form",
			OnSave: NewBstructFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		bstruct := new(models.Bstruct)
		FillUpForm(bstruct, formGroup, playground)
	case "Dstruct":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Dstruct Form",
			OnSave: NewDstructFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		dstruct := new(models.Dstruct)
		FillUpForm(dstruct, formGroup, playground)
	}
	formStage.Commit()
}
