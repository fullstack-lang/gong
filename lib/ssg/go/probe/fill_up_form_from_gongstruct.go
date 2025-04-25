// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/ssg/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *gongtable.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Chapter:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Chapter Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ChapterFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Content:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Content Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ContentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Page:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Page Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PageFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
