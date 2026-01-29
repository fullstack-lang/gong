// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/ssg/go/models"
)

// updateFillUpForm updates the current form if there is one
func (probe *Probe) updateFillUpForm() {
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
		case *PageFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Page", true)
			} else {
				FillUpFormFromGongstruct(onSave.page, probe)
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
	}
	formStage.Commit()
}
