// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/ssg/go/models"
)

// ux_form updates the current form if there is one
func (probe *Probe) ux_form() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *ChapterFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Chapter", true)
			} else {
				FillUpFormFromGongstruct(onSave.chapter, probe)
			}
		case *ContentFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Content", true)
			} else {
				FillUpFormFromGongstruct(onSave.content, probe)
			}
		case *JpgImageFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "JpgImage", true)
			} else {
				FillUpFormFromGongstruct(onSave.jpgimage, probe)
			}
		case *PageFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Page", true)
			} else {
				FillUpFormFromGongstruct(onSave.page, probe)
			}
		case *PngImageFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "PngImage", true)
			} else {
				FillUpFormFromGongstruct(onSave.pngimage, probe)
			}
		case *SectionFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Section", true)
			} else {
				FillUpFormFromGongstruct(onSave.section, probe)
			}
		case *SvgImageFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "SvgImage", true)
			} else {
				FillUpFormFromGongstruct(onSave.svgimage, probe)
			}
		}
	}
}

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
	case "Chapter":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Chapter Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ChapterFormCallback(
			nil,
			probe,
			formGroup,
		)
		chapter := new(models.Chapter)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(chapter, formGroup, probe)
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
	case "Page":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Page Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PageFormCallback(
			nil,
			probe,
			formGroup,
		)
		page := new(models.Page)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(page, formGroup, probe)
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
	case "Section":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Section Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SectionFormCallback(
			nil,
			probe,
			formGroup,
		)
		section := new(models.Section)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(section, formGroup, probe)
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
