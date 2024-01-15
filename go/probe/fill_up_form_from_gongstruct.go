// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	FillUpNamedFormFromGongstruct[T](instance, probe, formStage, gongtable.FormGroupDefaultName.ToString())

}

func FillUpNamedFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe, formStage *gongtable.StageStruct, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.GongBasicField:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "GongBasicField Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongBasicFieldFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongEnum:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "GongEnum Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongEnumFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongEnumValue:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "GongEnumValue Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongEnumValueFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongLink:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "GongLink Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongLinkFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongNote:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "GongNote Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongNoteFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongStruct:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "GongStruct Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongStructFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongTimeField:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "GongTimeField Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongTimeFieldFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Meta:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Meta Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MetaFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.MetaReference:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "MetaReference Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MetaReferenceFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ModelPkg:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ModelPkg Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ModelPkgFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PointerToGongStructField:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "PointerToGongStructField Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PointerToGongStructFieldFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SliceOfPointerToGongStructField:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "SliceOfPointerToGongStructField Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SliceOfPointerToGongStructFieldFormCallback(
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
