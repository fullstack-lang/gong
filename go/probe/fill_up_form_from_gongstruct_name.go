// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/go/models"
)

func FillUpFormFromGongstructName(
	playground *Playground,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := playground.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = "New"
	} else {
		prefix = "Update"
	}

	switch gongstructName {
	// insertion point
	case "GongBasicField":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " GongBasicField Form",
			OnSave: __gong__New__GongBasicFieldFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		gongbasicfield := new(models.GongBasicField)
		FillUpForm(gongbasicfield, formGroup, playground)
	case "GongEnum":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " GongEnum Form",
			OnSave: __gong__New__GongEnumFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		gongenum := new(models.GongEnum)
		FillUpForm(gongenum, formGroup, playground)
	case "GongEnumValue":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " GongEnumValue Form",
			OnSave: __gong__New__GongEnumValueFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		gongenumvalue := new(models.GongEnumValue)
		FillUpForm(gongenumvalue, formGroup, playground)
	case "GongLink":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " GongLink Form",
			OnSave: __gong__New__GongLinkFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		gonglink := new(models.GongLink)
		FillUpForm(gonglink, formGroup, playground)
	case "GongNote":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " GongNote Form",
			OnSave: __gong__New__GongNoteFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		gongnote := new(models.GongNote)
		FillUpForm(gongnote, formGroup, playground)
	case "GongStruct":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " GongStruct Form",
			OnSave: __gong__New__GongStructFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		gongstruct := new(models.GongStruct)
		FillUpForm(gongstruct, formGroup, playground)
	case "GongTimeField":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " GongTimeField Form",
			OnSave: __gong__New__GongTimeFieldFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		gongtimefield := new(models.GongTimeField)
		FillUpForm(gongtimefield, formGroup, playground)
	case "Meta":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Meta Form",
			OnSave: __gong__New__MetaFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		meta := new(models.Meta)
		FillUpForm(meta, formGroup, playground)
	case "MetaReference":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " MetaReference Form",
			OnSave: __gong__New__MetaReferenceFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		metareference := new(models.MetaReference)
		FillUpForm(metareference, formGroup, playground)
	case "ModelPkg":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " ModelPkg Form",
			OnSave: __gong__New__ModelPkgFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		modelpkg := new(models.ModelPkg)
		FillUpForm(modelpkg, formGroup, playground)
	case "PointerToGongStructField":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " PointerToGongStructField Form",
			OnSave: __gong__New__PointerToGongStructFieldFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		pointertogongstructfield := new(models.PointerToGongStructField)
		FillUpForm(pointertogongstructfield, formGroup, playground)
	case "SliceOfPointerToGongStructField":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " SliceOfPointerToGongStructField Form",
			OnSave: __gong__New__SliceOfPointerToGongStructFieldFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		sliceofpointertogongstructfield := new(models.SliceOfPointerToGongStructField)
		FillUpForm(sliceofpointertogongstructfield, formGroup, playground)
	}
	formStage.Commit()
}
