// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *gongtable.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.AttributeShape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "AttributeShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AttributeShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Classdiagram:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Classdiagram Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ClassdiagramFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.DiagramPackage:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "DiagramPackage Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramPackageFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongEnumShape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "GongEnumShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongEnumShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongEnumValueShape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "GongEnumValueShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongEnumValueShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongNoteLinkShape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "GongNoteLinkShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongNoteLinkShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongNoteShape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "GongNoteShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongNoteShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongStructShape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "GongStructShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongStructShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.LinkShape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "LinkShape Form",
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
	formStage.Commit()
}
