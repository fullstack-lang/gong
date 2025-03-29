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
			Name:  FormName,
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
			Name:  FormName,
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
	case "Button":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Button Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ButtonFormCallback(
			nil,
			probe,
			formGroup,
		)
		button := new(models.Button)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(button, formGroup, probe)
	case "Cursor":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Cursor Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CursorFormCallback(
			nil,
			probe,
			formGroup,
		)
		cursor := new(models.Cursor)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(cursor, formGroup, probe)
	case "Doc":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Doc Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DocFormCallback(
			nil,
			probe,
			formGroup,
		)
		doc := new(models.Doc)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(doc, formGroup, probe)
	case "Form":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Form Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFormCallback(
			nil,
			probe,
			formGroup,
		)
		form := new(models.Form)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(form, formGroup, probe)
	case "Load":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Load Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LoadFormCallback(
			nil,
			probe,
			formGroup,
		)
		load := new(models.Load)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(load, formGroup, probe)
	case "Slider":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Slider Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SliderFormCallback(
			nil,
			probe,
			formGroup,
		)
		slider := new(models.Slider)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(slider, formGroup, probe)
	case "Split":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Split Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SplitFormCallback(
			nil,
			probe,
			formGroup,
		)
		split := new(models.Split)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(split, formGroup, probe)
	case "Svg":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Svg Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SvgFormCallback(
			nil,
			probe,
			formGroup,
		)
		svg := new(models.Svg)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(svg, formGroup, probe)
	case "Table":
		formGroup := (&form.FormGroup{
			Name:  FormName,
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
	case "Tone":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Tone Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ToneFormCallback(
			nil,
			probe,
			formGroup,
		)
		tone := new(models.Tone)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tone, formGroup, probe)
	case "Tree":
		formGroup := (&form.FormGroup{
			Name:  FormName,
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
			Name:  FormName,
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
	case "Xlsx":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Xlsx Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__XlsxFormCallback(
			nil,
			probe,
			formGroup,
		)
		xlsx := new(models.Xlsx)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(xlsx, formGroup, probe)
	}
	formStage.Commit()
}
