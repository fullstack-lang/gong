// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
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
				probe,
			),
		}).Stage(formStage)
		gongbasicfield := new(models.GongBasicField)
		FillUpForm(gongbasicfield, formGroup, probe)
	case "GongEnum":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " GongEnum Form",
			OnSave: __gong__New__GongEnumFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		gongenum := new(models.GongEnum)
		FillUpForm(gongenum, formGroup, probe)
	case "GongEnumValue":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " GongEnumValue Form",
			OnSave: __gong__New__GongEnumValueFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		gongenumvalue := new(models.GongEnumValue)
		FillUpForm(gongenumvalue, formGroup, probe)
	case "GongLink":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " GongLink Form",
			OnSave: __gong__New__GongLinkFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		gonglink := new(models.GongLink)
		FillUpForm(gonglink, formGroup, probe)
	case "GongNote":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " GongNote Form",
			OnSave: __gong__New__GongNoteFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		gongnote := new(models.GongNote)
		FillUpForm(gongnote, formGroup, probe)
	case "GongStruct":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " GongStruct Form",
			OnSave: __gong__New__GongStructFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		gongstruct := new(models.GongStruct)
		FillUpForm(gongstruct, formGroup, probe)
	case "GongTimeField":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " GongTimeField Form",
			OnSave: __gong__New__GongTimeFieldFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		gongtimefield := new(models.GongTimeField)
		FillUpForm(gongtimefield, formGroup, probe)
	case "Meta":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Meta Form",
			OnSave: __gong__New__MetaFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		meta := new(models.Meta)
		FillUpForm(meta, formGroup, probe)
	case "MetaReference":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " MetaReference Form",
			OnSave: __gong__New__MetaReferenceFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		metareference := new(models.MetaReference)
		FillUpForm(metareference, formGroup, probe)
	case "ModelPkg":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " ModelPkg Form",
			OnSave: __gong__New__ModelPkgFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		modelpkg := new(models.ModelPkg)
		FillUpForm(modelpkg, formGroup, probe)
	case "PointerToGongStructField":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " PointerToGongStructField Form",
			OnSave: __gong__New__PointerToGongStructFieldFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		pointertogongstructfield := new(models.PointerToGongStructField)
		FillUpForm(pointertogongstructfield, formGroup, probe)
	case "SliceOfPointerToGongStructField":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " SliceOfPointerToGongStructField Form",
			OnSave: __gong__New__SliceOfPointerToGongStructFieldFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		sliceofpointertogongstructfield := new(models.SliceOfPointerToGongStructField)
		FillUpForm(sliceofpointertogongstructfield, formGroup, probe)
	}
	formStage.Commit()
}
