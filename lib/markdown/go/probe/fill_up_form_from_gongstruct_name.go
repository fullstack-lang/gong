// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/markdown/go/models"
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
	case "Content":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Content Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ContentFormCallback(
			nil,
			probe,
			formGroup,
		)
		content := new(models.Content)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(content, formGroup, probe)
	case "JpgImage":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "JpgImage Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__JpgImageFormCallback(
			nil,
			probe,
			formGroup,
		)
		jpgimage := new(models.JpgImage)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(jpgimage, formGroup, probe)
	case "PngImage":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PngImage Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PngImageFormCallback(
			nil,
			probe,
			formGroup,
		)
		pngimage := new(models.PngImage)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(pngimage, formGroup, probe)
	case "SvgImage":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SvgImage Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SvgImageFormCallback(
			nil,
			probe,
			formGroup,
		)
		svgimage := new(models.SvgImage)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(svgimage, formGroup, probe)
	}
	formStage.Commit()
}
