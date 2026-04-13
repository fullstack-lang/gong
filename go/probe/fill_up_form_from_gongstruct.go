// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.GongBasicField:
		formGroup := (&form.FormGroup{
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
		formGroup := (&form.FormGroup{
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
		formGroup := (&form.FormGroup{
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
		formGroup := (&form.FormGroup{
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
		formGroup := (&form.FormGroup{
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
		formGroup := (&form.FormGroup{
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
		formGroup := (&form.FormGroup{
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
	case *models.MetaReference:
		formGroup := (&form.FormGroup{
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
		formGroup := (&form.FormGroup{
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
		formGroup := (&form.FormGroup{
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
		formGroup := (&form.FormGroup{
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

	if probe.GetCommitMode() {
		formStage.Commit()
	}
}
