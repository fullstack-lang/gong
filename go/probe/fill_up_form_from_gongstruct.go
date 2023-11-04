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

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.GongBasicField:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update GongBasicField Form",
			OnSave: __gong__New__GongBasicFieldFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongEnum:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update GongEnum Form",
			OnSave: __gong__New__GongEnumFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongEnumValue:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update GongEnumValue Form",
			OnSave: __gong__New__GongEnumValueFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongLink:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update GongLink Form",
			OnSave: __gong__New__GongLinkFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongNote:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update GongNote Form",
			OnSave: __gong__New__GongNoteFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongStruct:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update GongStruct Form",
			OnSave: __gong__New__GongStructFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.GongTimeField:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update GongTimeField Form",
			OnSave: __gong__New__GongTimeFieldFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Meta:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Meta Form",
			OnSave: __gong__New__MetaFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.MetaReference:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update MetaReference Form",
			OnSave: __gong__New__MetaReferenceFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ModelPkg:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update ModelPkg Form",
			OnSave: __gong__New__ModelPkgFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.PointerToGongStructField:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update PointerToGongStructField Form",
			OnSave: __gong__New__PointerToGongStructFieldFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.SliceOfPointerToGongStructField:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update SliceOfPointerToGongStructField Form",
			OnSave: __gong__New__SliceOfPointerToGongStructFieldFormCallback(
				instancesTyped,
				probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
