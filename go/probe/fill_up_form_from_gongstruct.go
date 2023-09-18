// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, playground *Playground) {
	formStage := playground.formStage
	formStage.Reset()
	formStage.Commit()

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.GongBasicField:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update GongBasicField Form",
			OnSave: NewGongBasicFieldFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.GongEnum:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update GongEnum Form",
			OnSave: NewGongEnumFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.GongEnumValue:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update GongEnumValue Form",
			OnSave: NewGongEnumValueFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.GongLink:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update GongLink Form",
			OnSave: NewGongLinkFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.GongNote:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update GongNote Form",
			OnSave: NewGongNoteFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.GongStruct:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update GongStruct Form",
			OnSave: NewGongStructFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.GongTimeField:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update GongTimeField Form",
			OnSave: NewGongTimeFieldFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.Meta:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Meta Form",
			OnSave: NewMetaFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.MetaReference:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update MetaReference Form",
			OnSave: NewMetaReferenceFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.ModelPkg:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update ModelPkg Form",
			OnSave: NewModelPkgFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.PointerToGongStructField:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update PointerToGongStructField Form",
			OnSave: NewPointerToGongStructFieldFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.SliceOfPointerToGongStructField:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update SliceOfPointerToGongStructField Form",
			OnSave: NewSliceOfPointerToGongStructFieldFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
