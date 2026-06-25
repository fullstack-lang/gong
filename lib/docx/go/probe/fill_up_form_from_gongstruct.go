// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/docx/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *form.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Body:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: instancesTyped.GetName() + " : Body",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BodyFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Document:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: instancesTyped.GetName() + " : Document",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DocumentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Docx:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: instancesTyped.GetName() + " : Docx",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DocxFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.File:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: instancesTyped.GetName() + " : File",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FileFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Node:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: instancesTyped.GetName() + " : Node",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NodeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Paragraph:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: instancesTyped.GetName() + " : Paragraph",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParagraphFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ParagraphProperties:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: instancesTyped.GetName() + " : ParagraphProperties",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParagraphPropertiesFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ParagraphStyle:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: instancesTyped.GetName() + " : ParagraphStyle",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParagraphStyleFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Rune:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: instancesTyped.GetName() + " : Rune",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RuneFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.RuneProperties:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: instancesTyped.GetName() + " : RuneProperties",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RunePropertiesFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Table:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: instancesTyped.GetName() + " : Table",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TableFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TableColumn:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: instancesTyped.GetName() + " : TableColumn",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TableColumnFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TableProperties:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: instancesTyped.GetName() + " : TableProperties",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TablePropertiesFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TableRow:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: instancesTyped.GetName() + " : TableRow",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TableRowFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TableStyle:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: instancesTyped.GetName() + " : TableStyle",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TableStyleFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Text:
		formGroup := (&form.FormGroup{
			Name:  formName,
			Label: instancesTyped.GetName() + " : Text",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TextFormCallback(
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
