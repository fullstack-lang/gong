// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/pkg/docx/go/models"
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
	case "Body":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Body Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BodyFormCallback(
			nil,
			probe,
			formGroup,
		)
		body := new(models.Body)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(body, formGroup, probe)
	case "Document":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Document Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DocumentFormCallback(
			nil,
			probe,
			formGroup,
		)
		document := new(models.Document)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(document, formGroup, probe)
	case "Docx":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Docx Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DocxFormCallback(
			nil,
			probe,
			formGroup,
		)
		docx := new(models.Docx)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(docx, formGroup, probe)
	case "File":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "File Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FileFormCallback(
			nil,
			probe,
			formGroup,
		)
		file := new(models.File)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(file, formGroup, probe)
	case "Node":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Node Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NodeFormCallback(
			nil,
			probe,
			formGroup,
		)
		node := new(models.Node)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(node, formGroup, probe)
	case "Paragraph":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Paragraph Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParagraphFormCallback(
			nil,
			probe,
			formGroup,
		)
		paragraph := new(models.Paragraph)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(paragraph, formGroup, probe)
	case "ParagraphProperties":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ParagraphProperties Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParagraphPropertiesFormCallback(
			nil,
			probe,
			formGroup,
		)
		paragraphproperties := new(models.ParagraphProperties)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(paragraphproperties, formGroup, probe)
	case "ParagraphStyle":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "ParagraphStyle Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ParagraphStyleFormCallback(
			nil,
			probe,
			formGroup,
		)
		paragraphstyle := new(models.ParagraphStyle)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(paragraphstyle, formGroup, probe)
	case "Rune":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Rune Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RuneFormCallback(
			nil,
			probe,
			formGroup,
		)
		rune := new(models.Rune)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(rune, formGroup, probe)
	case "RuneProperties":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "RuneProperties Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RunePropertiesFormCallback(
			nil,
			probe,
			formGroup,
		)
		runeproperties := new(models.RuneProperties)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(runeproperties, formGroup, probe)
	case "Table":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Table Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TableFormCallback(
			nil,
			probe,
			formGroup,
		)
		table := new(models.Table)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(table, formGroup, probe)
	case "TableColumn":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TableColumn Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TableColumnFormCallback(
			nil,
			probe,
			formGroup,
		)
		tablecolumn := new(models.TableColumn)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tablecolumn, formGroup, probe)
	case "TableProperties":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TableProperties Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TablePropertiesFormCallback(
			nil,
			probe,
			formGroup,
		)
		tableproperties := new(models.TableProperties)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tableproperties, formGroup, probe)
	case "TableRow":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TableRow Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TableRowFormCallback(
			nil,
			probe,
			formGroup,
		)
		tablerow := new(models.TableRow)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tablerow, formGroup, probe)
	case "TableStyle":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "TableStyle Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TableStyleFormCallback(
			nil,
			probe,
			formGroup,
		)
		tablestyle := new(models.TableStyle)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tablestyle, formGroup, probe)
	case "Text":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Text Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TextFormCallback(
			nil,
			probe,
			formGroup,
		)
		text := new(models.Text)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(text, formGroup, probe)
	}
	formStage.Commit()
}
