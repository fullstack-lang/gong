// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/tests/books/go/models"
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
func __gong__New__BookTypeFormCallback(
	_instance *models.BookType,
	probe *Probe,
	formGroup *form.FormGroup,
) (booktypeFormCallback *FormCallback[*models.BookType]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBookTypeFields,
	)
}

type BookTypeFormCallback = FormCallback[*models.BookType]

func saveBookTypeFields(
	_instance *models.BookType,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Edition":
			FormDivBasicFieldToField(&(_instance.Edition), formDiv)
		case "Isbn":
			FormDivBasicFieldToField(&(_instance.Isbn), formDiv)
		case "Bestseller":
			FormDivBasicFieldToField(&(_instance.Bestseller), formDiv)
		case "Title":
			FormDivBasicFieldToField(&(_instance.Title), formDiv)
		case "Author":
			FormDivBasicFieldToField(&(_instance.Author), formDiv)
		case "Year":
			FormDivBasicFieldToField(&(_instance.Year), formDiv)
		case "Format":
			FormDivBasicFieldToField(&(_instance.Format), formDiv)
		case "Credit":
			FormDivSliceOfPointersToField(_instance, "Credit", &(_instance.Credit), formDiv, probe)
		case "Books:Book":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Book", func(owner *models.Books) *[]*models.BookType { return &owner.Book })
		}
	}
}

func __gong__New__BooksFormCallback(
	_instance *models.Books,
	probe *Probe,
	formGroup *form.FormGroup,
) (booksFormCallback *FormCallback[*models.Books]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveBooksFields,
	)
}

type BooksFormCallback = FormCallback[*models.Books]

func saveBooksFields(
	_instance *models.Books,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Book":
			FormDivSliceOfPointersToField(_instance, "Book", &(_instance.Book), formDiv, probe)
		}
	}
}

func __gong__New__CreditFormCallback(
	_instance *models.Credit,
	probe *Probe,
	formGroup *form.FormGroup,
) (creditFormCallback *FormCallback[*models.Credit]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveCreditFields,
	)
}

type CreditFormCallback = FormCallback[*models.Credit]

func saveCreditFields(
	_instance *models.Credit,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "Page":
			FormDivBasicFieldToField(&(_instance.Page), formDiv)
		case "Credit_type":
			FormDivBasicFieldToField(&(_instance.Credit_type), formDiv)
		case "Link":
			FormDivSliceOfPointersToField(_instance, "Link", &(_instance.Link), formDiv, probe)
		case "Credit_words":
			FormDivBasicFieldToField(&(_instance.Credit_words), formDiv)
		case "Credit_symbol":
			FormDivBasicFieldToField(&(_instance.Credit_symbol), formDiv)
		case "BookType:Credit":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Credit", func(owner *models.BookType) *[]*models.Credit { return &owner.Credit })
		}
	}
}

func __gong__New__LinkFormCallback(
	_instance *models.Link,
	probe *Probe,
	formGroup *form.FormGroup,
) (linkFormCallback *FormCallback[*models.Link]) {
	return NewFormCallback(
		_instance,
		probe,
		formGroup,
		saveLinkFields,
	)
}

type LinkFormCallback = FormCallback[*models.Link]

func saveLinkFields(
	_instance *models.Link,
	probe *Probe,
	formGroup *form.FormGroup,
) {
	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(_instance.Name), formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(_instance.NameXSD), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(_instance.EnclosedText), formDiv)
		case "Credit:Link":
			FormDivReverseSliceOfPointersToField(_instance, formDiv, probe, "Link", func(owner *models.Credit) *[]*models.Link { return &owner.Link })
		}
	}
}

