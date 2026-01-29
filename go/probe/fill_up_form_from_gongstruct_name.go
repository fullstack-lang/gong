// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/go/models"
)

// updateFillUpForm updates the current form if there is one
func (probe *Probe) updateFillUpForm() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *GongBasicFieldFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "GongBasicField", true)
			} else {
				FillUpFormFromGongstruct(onSave.gongbasicfield, probe)
			}
		case *GongEnumFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "GongEnum", true)
			} else {
				FillUpFormFromGongstruct(onSave.gongenum, probe)
			}
		case *GongEnumValueFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "GongEnumValue", true)
			} else {
				FillUpFormFromGongstruct(onSave.gongenumvalue, probe)
			}
		case *GongLinkFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "GongLink", true)
			} else {
				FillUpFormFromGongstruct(onSave.gonglink, probe)
			}
		case *GongNoteFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "GongNote", true)
			} else {
				FillUpFormFromGongstruct(onSave.gongnote, probe)
			}
		case *GongStructFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "GongStruct", true)
			} else {
				FillUpFormFromGongstruct(onSave.gongstruct, probe)
			}
		case *GongTimeFieldFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "GongTimeField", true)
			} else {
				FillUpFormFromGongstruct(onSave.gongtimefield, probe)
			}
		case *MetaReferenceFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "MetaReference", true)
			} else {
				FillUpFormFromGongstruct(onSave.metareference, probe)
			}
		case *ModelPkgFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "ModelPkg", true)
			} else {
				FillUpFormFromGongstruct(onSave.modelpkg, probe)
			}
		case *PointerToGongStructFieldFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "PointerToGongStructField", true)
			} else {
				FillUpFormFromGongstruct(onSave.pointertogongstructfield, probe)
			}
		case *SliceOfPointerToGongStructFieldFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "SliceOfPointerToGongStructField", true)
			} else {
				FillUpFormFromGongstruct(onSave.sliceofpointertogongstructfield, probe)
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
	case "GongBasicField":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GongBasicField Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongBasicFieldFormCallback(
			nil,
			probe,
			formGroup,
		)
		gongbasicfield := new(models.GongBasicField)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(gongbasicfield, formGroup, probe)
	case "GongEnum":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GongEnum Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongEnumFormCallback(
			nil,
			probe,
			formGroup,
		)
		gongenum := new(models.GongEnum)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(gongenum, formGroup, probe)
	case "GongEnumValue":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GongEnumValue Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongEnumValueFormCallback(
			nil,
			probe,
			formGroup,
		)
		gongenumvalue := new(models.GongEnumValue)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(gongenumvalue, formGroup, probe)
	case "GongLink":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GongLink Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongLinkFormCallback(
			nil,
			probe,
			formGroup,
		)
		gonglink := new(models.GongLink)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(gonglink, formGroup, probe)
	case "GongNote":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GongNote Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongNoteFormCallback(
			nil,
			probe,
			formGroup,
		)
		gongnote := new(models.GongNote)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(gongnote, formGroup, probe)
	case "GongStruct":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GongStruct Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongStructFormCallback(
			nil,
			probe,
			formGroup,
		)
		gongstruct := new(models.GongStruct)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(gongstruct, formGroup, probe)
	case "GongTimeField":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GongTimeField Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongTimeFieldFormCallback(
			nil,
			probe,
			formGroup,
		)
		gongtimefield := new(models.GongTimeField)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(gongtimefield, formGroup, probe)
	case "MetaReference":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "MetaReference Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MetaReferenceFormCallback(
			nil,
			probe,
			formGroup,
		)
		metareference := new(models.MetaReference)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(metareference, formGroup, probe)
	case "ModelPkg":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ModelPkg Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ModelPkgFormCallback(
			nil,
			probe,
			formGroup,
		)
		modelpkg := new(models.ModelPkg)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(modelpkg, formGroup, probe)
	case "PointerToGongStructField":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "PointerToGongStructField Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PointerToGongStructFieldFormCallback(
			nil,
			probe,
			formGroup,
		)
		pointertogongstructfield := new(models.PointerToGongStructField)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(pointertogongstructfield, formGroup, probe)
	case "SliceOfPointerToGongStructField":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "SliceOfPointerToGongStructField Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SliceOfPointerToGongStructFieldFormCallback(
			nil,
			probe,
			formGroup,
		)
		sliceofpointertogongstructfield := new(models.SliceOfPointerToGongStructField)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(sliceofpointertogongstructfield, formGroup, probe)
	}
	formStage.Commit()
}
