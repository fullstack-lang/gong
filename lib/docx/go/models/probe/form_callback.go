// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/docx/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

type FormCallbackIF interface {
	GetCreationMode() bool
	GetInstance() any
	GetGongstructName() string
	OnSave()
}

type FormCallback[T models.PointerToGongstruct] struct {
	Instance     T
	CreationMode bool
	probe        *Probe
	formGroup    *form.FormGroup
	saveFields   func(instance T, probe *Probe, formGroup *form.FormGroup)
}

func NewFormCallback[T models.PointerToGongstruct](
	instance T,
	probe *Probe,
	formGroup *form.FormGroup,
	saveFields func(instance T, probe *Probe, formGroup *form.FormGroup),
) *FormCallback[T] {
	return &FormCallback[T]{
		Instance:     instance,
		CreationMode: any(instance) == nil,
		probe:        probe,
		formGroup:    formGroup,
		saveFields:   saveFields,
	}
}

func (cb *FormCallback[T]) GetCreationMode() bool     { return cb.CreationMode }
func (cb *FormCallback[T]) GetInstance() any           { return cb.Instance }
func (cb *FormCallback[T]) GetGongstructName() string { return models.GetPointerToGongstructName[T]() }

func (cb *FormCallback[T]) OnSave() {
	cb.probe.stageOfInterest.Lock()
	defer cb.probe.stageOfInterest.Unlock()

	cb.probe.formStage.Checkout()

	if any(cb.Instance) == nil {
		cb.Instance = cb.probe.stageOfInterest.GongNewInstance[T]()
	}

	cb.saveFields(cb.Instance, cb.probe, cb.formGroup)

	if cb.formGroup.HasSuppressButtonBeenPressed {
		cb.Instance.UnstageVoid(cb.probe.stageOfInterest)
	}

	cb.probe.stageOfInterest.Commit()
	updateProbeTable[T](cb.probe)

	if cb.CreationMode || cb.formGroup.HasSuppressButtonBeenPressed {
		cb.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(cb.probe.formStage)
		newFormGroup.OnSave = NewFormCallback[T](
			*new(T),
			cb.probe,
			newFormGroup,
			cb.saveFields,
		)
		newInstance := models.GongNewInstance[T]()
		FillUpForm(newInstance, newFormGroup, cb.probe)
		cb.probe.formStage.Commit()
	}

	cb.probe.ux_tree()
}

// insertion point
func __gong__New__BodyFormCallback(
	_instance *models.Body,
	probe *Probe,
	formGroup *form.FormGroup,
) (bodyFormCallback *FormCallback[*models.Body]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBodyFields,
	)
}

type BodyFormCallback = FormCallback[*models.Body]

func saveBodyFields(
	_instance *models.Body,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Paragraphs":
			FormDivSliceOfPointersToField(_instance, "Paragraphs", &(_instance.Paragraphs), formDiv, probe)
		case "Tables":
			FormDivSliceOfPointersToField(_instance, "Tables", &(_instance.Tables), formDiv, probe)
		case "LastParagraph":
			FormDivSelectFieldToField(&(_instance.LastParagraph), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__DocumentFormCallback(
	_instance *models.Document,
	probe *Probe,
	formGroup *form.FormGroup,
) (documentFormCallback *FormCallback[*models.Document]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDocumentFields,
	)
}

type DocumentFormCallback = FormCallback[*models.Document]

func saveDocumentFields(
	_instance *models.Document,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "File":
			FormDivSelectFieldToField(&(_instance.File), probe.stageOfInterest, formDiv)
		case "Root":
			FormDivSelectFieldToField(&(_instance.Root), probe.stageOfInterest, formDiv)
		case "Body":
			FormDivSelectFieldToField(&(_instance.Body), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__DocxFormCallback(
	_instance *models.Docx,
	probe *Probe,
	formGroup *form.FormGroup,
) (docxFormCallback *FormCallback[*models.Docx]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveDocxFields,
	)
}

type DocxFormCallback = FormCallback[*models.Docx]

func saveDocxFields(
	_instance *models.Docx,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Files":
			FormDivSliceOfPointersToField(_instance, "Files", &(_instance.Files), formDiv, probe)
		case "Document":
			FormDivSelectFieldToField(&(_instance.Document), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__FileFormCallback(
	_instance *models.File,
	probe *Probe,
	formGroup *form.FormGroup,
) (fileFormCallback *FormCallback[*models.File]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveFileFields,
	)
}

type FileFormCallback = FormCallback[*models.File]

func saveFileFields(
	_instance *models.File,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Docx:Files":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Files", func(owner *models.Docx) *[]*models.File { return &owner.Files })
		}
	}
}

func __gong__New__NodeFormCallback(
	_instance *models.Node,
	probe *Probe,
	formGroup *form.FormGroup,
) (nodeFormCallback *FormCallback[*models.Node]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveNodeFields,
	)
}

type NodeFormCallback = FormCallback[*models.Node]

func saveNodeFields(
	_instance *models.Node,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Nodes":
			FormDivSliceOfPointersToField(_instance, "Nodes", &(_instance.Nodes), formDiv, probe)
		case "Node:Nodes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Nodes", func(owner *models.Node) *[]*models.Node { return &owner.Nodes })
		}
	}
}

func __gong__New__ParagraphFormCallback(
	_instance *models.Paragraph,
	probe *Probe,
	formGroup *form.FormGroup,
) (paragraphFormCallback *FormCallback[*models.Paragraph]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveParagraphFields,
	)
}

type ParagraphFormCallback = FormCallback[*models.Paragraph]

func saveParagraphFields(
	_instance *models.Paragraph,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(_instance.Content), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(_instance.Node), probe.stageOfInterest, formDiv)
		case "ParagraphProperties":
			FormDivSelectFieldToField(&(_instance.ParagraphProperties), probe.stageOfInterest, formDiv)
		case "Runes":
			FormDivSliceOfPointersToField(_instance, "Runes", &(_instance.Runes), formDiv, probe)
		case "CollatedText":
			FormDivBasicFieldToField(&(_instance.CollatedText), formDiv)
		case "Next":
			FormDivSelectFieldToField(&(_instance.Next), probe.stageOfInterest, formDiv)
		case "Previous":
			FormDivSelectFieldToField(&(_instance.Previous), probe.stageOfInterest, formDiv)
		case "EnclosingBody":
			FormDivSelectFieldToField(&(_instance.EnclosingBody), probe.stageOfInterest, formDiv)
		case "EnclosingTableColumn":
			FormDivSelectFieldToField(&(_instance.EnclosingTableColumn), probe.stageOfInterest, formDiv)
		case "Body:Paragraphs":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Paragraphs", func(owner *models.Body) *[]*models.Paragraph { return &owner.Paragraphs })
		case "TableColumn:Paragraphs":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Paragraphs", func(owner *models.TableColumn) *[]*models.Paragraph { return &owner.Paragraphs })
		}
	}
}

func __gong__New__ParagraphPropertiesFormCallback(
	_instance *models.ParagraphProperties,
	probe *Probe,
	formGroup *form.FormGroup,
) (paragraphpropertiesFormCallback *FormCallback[*models.ParagraphProperties]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveParagraphPropertiesFields,
	)
}

type ParagraphPropertiesFormCallback = FormCallback[*models.ParagraphProperties]

func saveParagraphPropertiesFields(
	_instance *models.ParagraphProperties,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(_instance.Content), formDiv)
		case "ParagraphStyle":
			FormDivSelectFieldToField(&(_instance.ParagraphStyle), probe.stageOfInterest, formDiv)
		case "Node":
			FormDivSelectFieldToField(&(_instance.Node), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__ParagraphStyleFormCallback(
	_instance *models.ParagraphStyle,
	probe *Probe,
	formGroup *form.FormGroup,
) (paragraphstyleFormCallback *FormCallback[*models.ParagraphStyle]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveParagraphStyleFields,
	)
}

type ParagraphStyleFormCallback = FormCallback[*models.ParagraphStyle]

func saveParagraphStyleFields(
	_instance *models.ParagraphStyle,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(_instance.Node), probe.stageOfInterest, formDiv)
		case "Content":
			FormDivBasicFieldToField(&(_instance.Content), formDiv)
		case "ValAttr":
			FormDivBasicFieldToField(&(_instance.ValAttr), formDiv)
		}
	}
}

func __gong__New__RuneFormCallback(
	_instance *models.Rune,
	probe *Probe,
	formGroup *form.FormGroup,
) (runeFormCallback *FormCallback[*models.Rune]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRuneFields,
	)
}

type RuneFormCallback = FormCallback[*models.Rune]

func saveRuneFields(
	_instance *models.Rune,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(_instance.Content), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(_instance.Node), probe.stageOfInterest, formDiv)
		case "Text":
			FormDivSelectFieldToField(&(_instance.Text), probe.stageOfInterest, formDiv)
		case "RuneProperties":
			FormDivSelectFieldToField(&(_instance.RuneProperties), probe.stageOfInterest, formDiv)
		case "EnclosingParagraph":
			FormDivSelectFieldToField(&(_instance.EnclosingParagraph), probe.stageOfInterest, formDiv)
		case "Paragraph:Runes":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Runes", func(owner *models.Paragraph) *[]*models.Rune { return &owner.Runes })
		}
	}
}

func __gong__New__RunePropertiesFormCallback(
	_instance *models.RuneProperties,
	probe *Probe,
	formGroup *form.FormGroup,
) (runepropertiesFormCallback *FormCallback[*models.RuneProperties]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveRunePropertiesFields,
	)
}

type RunePropertiesFormCallback = FormCallback[*models.RuneProperties]

func saveRunePropertiesFields(
	_instance *models.RuneProperties,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(_instance.Node), probe.stageOfInterest, formDiv)
		case "IsBold":
			FormDivBasicFieldToField(&(_instance.IsBold), formDiv)
		case "IsStrike":
			FormDivBasicFieldToField(&(_instance.IsStrike), formDiv)
		case "IsItalic":
			FormDivBasicFieldToField(&(_instance.IsItalic), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(_instance.Content), formDiv)
		}
	}
}

func __gong__New__TableFormCallback(
	_instance *models.Table,
	probe *Probe,
	formGroup *form.FormGroup,
) (tableFormCallback *FormCallback[*models.Table]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTableFields,
	)
}

type TableFormCallback = FormCallback[*models.Table]

func saveTableFields(
	_instance *models.Table,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(_instance.Node), probe.stageOfInterest, formDiv)
		case "Content":
			FormDivBasicFieldToField(&(_instance.Content), formDiv)
		case "TableProperties":
			FormDivSelectFieldToField(&(_instance.TableProperties), probe.stageOfInterest, formDiv)
		case "TableRows":
			FormDivSliceOfPointersToField(_instance, "TableRows", &(_instance.TableRows), formDiv, probe)
		case "Body:Tables":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Tables", func(owner *models.Body) *[]*models.Table { return &owner.Tables })
		}
	}
}

func __gong__New__TableColumnFormCallback(
	_instance *models.TableColumn,
	probe *Probe,
	formGroup *form.FormGroup,
) (tablecolumnFormCallback *FormCallback[*models.TableColumn]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTableColumnFields,
	)
}

type TableColumnFormCallback = FormCallback[*models.TableColumn]

func saveTableColumnFields(
	_instance *models.TableColumn,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(_instance.Content), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(_instance.Node), probe.stageOfInterest, formDiv)
		case "Paragraphs":
			FormDivSliceOfPointersToField(_instance, "Paragraphs", &(_instance.Paragraphs), formDiv, probe)
		case "TableRow:TableColumns":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TableColumns", func(owner *models.TableRow) *[]*models.TableColumn { return &owner.TableColumns })
		}
	}
}

func __gong__New__TablePropertiesFormCallback(
	_instance *models.TableProperties,
	probe *Probe,
	formGroup *form.FormGroup,
) (tablepropertiesFormCallback *FormCallback[*models.TableProperties]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTablePropertiesFields,
	)
}

type TablePropertiesFormCallback = FormCallback[*models.TableProperties]

func saveTablePropertiesFields(
	_instance *models.TableProperties,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(_instance.Node), probe.stageOfInterest, formDiv)
		case "Content":
			FormDivBasicFieldToField(&(_instance.Content), formDiv)
		case "TableStyle":
			FormDivSelectFieldToField(&(_instance.TableStyle), probe.stageOfInterest, formDiv)
		}
	}
}

func __gong__New__TableRowFormCallback(
	_instance *models.TableRow,
	probe *Probe,
	formGroup *form.FormGroup,
) (tablerowFormCallback *FormCallback[*models.TableRow]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTableRowFields,
	)
}

type TableRowFormCallback = FormCallback[*models.TableRow]

func saveTableRowFields(
	_instance *models.TableRow,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(_instance.Content), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(_instance.Node), probe.stageOfInterest, formDiv)
		case "TableColumns":
			FormDivSliceOfPointersToField(_instance, "TableColumns", &(_instance.TableColumns), formDiv, probe)
		case "Table:TableRows":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "TableRows", func(owner *models.Table) *[]*models.TableRow { return &owner.TableRows })
		}
	}
}

func __gong__New__TableStyleFormCallback(
	_instance *models.TableStyle,
	probe *Probe,
	formGroup *form.FormGroup,
) (tablestyleFormCallback *FormCallback[*models.TableStyle]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTableStyleFields,
	)
}

type TableStyleFormCallback = FormCallback[*models.TableStyle]

func saveTableStyleFields(
	_instance *models.TableStyle,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(_instance.Node), probe.stageOfInterest, formDiv)
		case "Content":
			FormDivBasicFieldToField(&(_instance.Content), formDiv)
		case "Val":
			FormDivBasicFieldToField(&(_instance.Val), formDiv)
		}
	}
}

func __gong__New__TextFormCallback(
	_instance *models.Text,
	probe *Probe,
	formGroup *form.FormGroup,
) (textFormCallback *FormCallback[*models.Text]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveTextFields,
	)
}

type TextFormCallback = FormCallback[*models.Text]

func saveTextFields(
	_instance *models.Text,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(_instance.Content), formDiv)
		case "Node":
			FormDivSelectFieldToField(&(_instance.Node), probe.stageOfInterest, formDiv)
		case "PreserveWhiteSpace":
			FormDivBasicFieldToField(&(_instance.PreserveWhiteSpace), formDiv)
		case "EnclosingRune":
			FormDivSelectFieldToField(&(_instance.EnclosingRune), probe.stageOfInterest, formDiv)
		}
	}
}

