// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/xlsx/go/models"
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
	case "DisplaySelection":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DisplaySelection Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DisplaySelectionFormCallback(
			nil,
			probe,
			formGroup,
		)
		displayselection := new(models.DisplaySelection)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(displayselection, formGroup, probe)
	case "XLCell":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "XLCell Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__XLCellFormCallback(
			nil,
			probe,
			formGroup,
		)
		xlcell := new(models.XLCell)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xlcell, formGroup, probe)
	case "XLFile":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "XLFile Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__XLFileFormCallback(
			nil,
			probe,
			formGroup,
		)
		xlfile := new(models.XLFile)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xlfile, formGroup, probe)
	case "XLRow":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "XLRow Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__XLRowFormCallback(
			nil,
			probe,
			formGroup,
		)
		xlrow := new(models.XLRow)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xlrow, formGroup, probe)
	case "XLSheet":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "XLSheet Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__XLSheetFormCallback(
			nil,
			probe,
			formGroup,
		)
		xlsheet := new(models.XLSheet)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xlsheet, formGroup, probe)
	}
	formStage.Commit()
}
