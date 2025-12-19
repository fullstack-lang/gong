// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/pkg/docx/go/models"
)

func FillUpFormFromGongstruct(instance any, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()

	FillUpNamedFormFromGongstruct(instance, probe, formStage, FormName)

}

func FillUpNamedFormFromGongstruct(instance any, probe *Probe, formStage *gongtable.Stage, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Body:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Body Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BodyFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Document:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Document Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DocumentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Docx:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Docx Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DocxFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.File:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "File Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FileFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Node:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Node Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NodeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Paragraph:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Paragraph Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParagraphFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ParagraphProperties:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ParagraphProperties Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParagraphPropertiesFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.ParagraphStyle:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "ParagraphStyle Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParagraphStyleFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Rune:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Rune Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RuneFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.RuneProperties:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "RuneProperties Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RunePropertiesFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Table:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Table Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TableFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TableColumn:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "TableColumn Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TableColumnFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TableProperties:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "TableProperties Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TablePropertiesFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TableRow:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "TableRow Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TableRowFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.TableStyle:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "TableStyle Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TableStyleFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Text:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Text Form",
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
	formStage.Commit()
}
