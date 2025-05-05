// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
)

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
	case "AttributeShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "AttributeShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AttributeShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		attributeshape := new(models.AttributeShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attributeshape, formGroup, probe)
	case "Classdiagram":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Classdiagram Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ClassdiagramFormCallback(
			nil,
			probe,
			formGroup,
		)
		classdiagram := new(models.Classdiagram)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(classdiagram, formGroup, probe)
	case "DiagramPackage":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "DiagramPackage Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DiagramPackageFormCallback(
			nil,
			probe,
			formGroup,
		)
		diagrampackage := new(models.DiagramPackage)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(diagrampackage, formGroup, probe)
	case "GongEnumShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GongEnumShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongEnumShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		gongenumshape := new(models.GongEnumShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(gongenumshape, formGroup, probe)
	case "GongEnumValueShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GongEnumValueShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongEnumValueShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		gongenumvalueshape := new(models.GongEnumValueShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(gongenumvalueshape, formGroup, probe)
	case "GongNoteLinkShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GongNoteLinkShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongNoteLinkShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		gongnotelinkshape := new(models.GongNoteLinkShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(gongnotelinkshape, formGroup, probe)
	case "GongNoteShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GongNoteShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongNoteShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		gongnoteshape := new(models.GongNoteShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(gongnoteshape, formGroup, probe)
	case "GongStructShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GongStructShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongStructShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		gongstructshape := new(models.GongStructShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(gongstructshape, formGroup, probe)
	case "LinkShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "LinkShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinkShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		linkshape := new(models.LinkShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(linkshape, formGroup, probe)
	}
	formStage.Commit()
}
