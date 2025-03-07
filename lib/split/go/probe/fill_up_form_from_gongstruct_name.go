// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/split/go/models"
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
	case "AsSplit":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "AsSplit Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AsSplitFormCallback(
			nil,
			probe,
			formGroup,
		)
		assplit := new(models.AsSplit)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(assplit, formGroup, probe)
	case "AsSplitArea":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "AsSplitArea Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AsSplitAreaFormCallback(
			nil,
			probe,
			formGroup,
		)
		assplitarea := new(models.AsSplitArea)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(assplitarea, formGroup, probe)
	case "Table":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Table Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TableFormCallback(
			nil,
			probe,
			formGroup,
		)
		table := new(models.Table)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(table, formGroup, probe)
	case "Tree":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Tree Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TreeFormCallback(
			nil,
			probe,
			formGroup,
		)
		tree := new(models.Tree)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tree, formGroup, probe)
	case "View":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "View Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ViewFormCallback(
			nil,
			probe,
			formGroup,
		)
		view := new(models.View)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(view, formGroup, probe)
	}
	formStage.Commit()
}
