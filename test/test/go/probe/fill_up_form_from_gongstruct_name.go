// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/test/test/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "Astruct":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Astruct Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AstructFormCallback(
			nil,
			probe,
			formGroup,
		)
		astruct := new(models.Astruct)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(astruct, formGroup, probe)
	case "AstructBstruct2Use":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "AstructBstruct2Use Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AstructBstruct2UseFormCallback(
			nil,
			probe,
			formGroup,
		)
		astructbstruct2use := new(models.AstructBstruct2Use)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(astructbstruct2use, formGroup, probe)
	case "AstructBstructUse":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "AstructBstructUse Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AstructBstructUseFormCallback(
			nil,
			probe,
			formGroup,
		)
		astructbstructuse := new(models.AstructBstructUse)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(astructbstructuse, formGroup, probe)
	case "Bstruct":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Bstruct Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BstructFormCallback(
			nil,
			probe,
			formGroup,
		)
		bstruct := new(models.Bstruct)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(bstruct, formGroup, probe)
	case "Dstruct":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Dstruct Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DstructFormCallback(
			nil,
			probe,
			formGroup,
		)
		dstruct := new(models.Dstruct)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(dstruct, formGroup, probe)
	case "F0123456789012345678901234567890":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "F0123456789012345678901234567890 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__F0123456789012345678901234567890FormCallback(
			nil,
			probe,
			formGroup,
		)
		f0123456789012345678901234567890 := new(models.F0123456789012345678901234567890)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(f0123456789012345678901234567890, formGroup, probe)
	case "Gstruct":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Gstruct Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GstructFormCallback(
			nil,
			probe,
			formGroup,
		)
		gstruct := new(models.Gstruct)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(gstruct, formGroup, probe)
	}
	formStage.Commit()
}
