// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/doc/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *gongtable.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
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
	case *models.Field:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Field Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FieldFormCallback(
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
	case *models.GongEnumValueEntry:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "GongEnumValueEntry Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GongEnumValueEntryFormCallback(
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
	case *models.Link:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Link Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinkFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.NoteShape:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "NoteShape Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteShapeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.NoteShapeLink:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "NoteShapeLink Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteShapeLinkFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Position:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Position Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PositionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.UmlState:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "UmlState Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__UmlStateFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Umlsc:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Umlsc Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__UmlscFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Vertice:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Vertice Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__VerticeFormCallback(
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
