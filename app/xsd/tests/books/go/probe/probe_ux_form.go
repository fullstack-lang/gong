// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/tests/books/go/models"
)

// ux_form updates the current form if there is one
func (probe *Probe) ux_form() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *BookTypeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "BookType", true)
			} else {
				FillUpFormFromGongstruct(onSave.booktype, probe)
			}
		case *BooksFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Books", true)
			} else {
				FillUpFormFromGongstruct(onSave.books, probe)
			}
		case *CreditFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Credit", true)
			} else {
				FillUpFormFromGongstruct(onSave.credit, probe)
			}
		case *LinkFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Link", true)
			} else {
				FillUpFormFromGongstruct(onSave.link, probe)
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
	case "BookType":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "BookType Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BookTypeFormCallback(
			nil,
			probe,
			formGroup,
		)
		booktype := new(models.BookType)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(booktype, formGroup, probe)
	case "Books":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Books Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BooksFormCallback(
			nil,
			probe,
			formGroup,
		)
		books := new(models.Books)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(books, formGroup, probe)
	case "Credit":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Credit Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CreditFormCallback(
			nil,
			probe,
			formGroup,
		)
		credit := new(models.Credit)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(credit, formGroup, probe)
	case "Link":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Link Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinkFormCallback(
			nil,
			probe,
			formGroup,
		)
		link := new(models.Link)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(link, formGroup, probe)
	}
	formStage.Commit()
}
