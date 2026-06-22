// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.AttributeShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "AttributeShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AttributeShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Classdiagram:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "Classdiagram",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ClassdiagramFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DiagramPackage:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "DiagramPackage",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramPackageFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongEnumShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "GongEnumShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongEnumShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongEnumValueShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "GongEnumValueShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongEnumValueShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongNoteLinkShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "GongNoteLinkShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongNoteLinkShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongNoteShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "GongNoteShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongNoteShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongStructShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "GongStructShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongStructShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.LinkShape:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: "LinkShape",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinkShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}

	if probe.GetCommitMode() {
		formStage.Commit()
	}
}
