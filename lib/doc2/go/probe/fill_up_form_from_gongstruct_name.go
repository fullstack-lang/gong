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
	case "FieldShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "FieldShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FieldShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		fieldshape := new(models.AttributeShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(fieldshape, formGroup, probe)
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
	case "GongEnumValueEntry":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "GongEnumValueEntry Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongEnumValueEntryFormCallback(
			nil,
			probe,
			formGroup,
		)
		gongenumvalueentry := new(models.GongEnumValueEntry)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(gongenumvalueentry, formGroup, probe)
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
	case "NoteShape":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "NoteShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteShapeFormCallback(
			nil,
			probe,
			formGroup,
		)
		noteshape := new(models.NoteShape)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(noteshape, formGroup, probe)
	case "NoteShapeLink":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "NoteShapeLink Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteShapeLinkFormCallback(
			nil,
			probe,
			formGroup,
		)
		noteshapelink := new(models.NoteShapeLink)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(noteshapelink, formGroup, probe)
	case "Position":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Position Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PositionFormCallback(
			nil,
			probe,
			formGroup,
		)
		position := new(models.Position)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(position, formGroup, probe)
	case "UmlState":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "UmlState Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__UmlStateFormCallback(
			nil,
			probe,
			formGroup,
		)
		umlstate := new(models.UmlState)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(umlstate, formGroup, probe)
	case "Umlsc":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Umlsc Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__UmlscFormCallback(
			nil,
			probe,
			formGroup,
		)
		umlsc := new(models.Umlsc)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(umlsc, formGroup, probe)
	case "Vertice":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Vertice Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__VerticeFormCallback(
			nil,
			probe,
			formGroup,
		)
		vertice := new(models.Vertice)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(vertice, formGroup, probe)
	}
	formStage.Commit()
}
