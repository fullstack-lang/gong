// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/tests/books/go/models"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func StageSetFillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *StageSetProbe,
) {
	switch inst := instance.(type) {
	case *models.BookType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Edition", inst.Edition, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Isbn", inst.Isbn, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Bestseller", inst.Bestseller, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Title", inst.Title, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Author", inst.Author, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Year", inst.Year, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Format", inst.Format, probe.formStage, formGroup)

		{
			// Slice of pointers: Credit
			div := (&form.FormDiv{Name: "Credit"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Credit {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Credit",
				Label: "Credit",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Bookss {
				for _, target := range src.Book {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Books", "Book", refNames, formGroup, probe.formStage)
		}
	case *models.Books:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: Book
			div := (&form.FormDiv{Name: "Book"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Book {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Book",
				Label: "Book",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
	case *models.Credit:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Page", inst.Page, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Credit_type", inst.Credit_type, probe.formStage, formGroup)

		{
			// Slice of pointers: Link
			div := (&form.FormDiv{Name: "Link"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Link {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Link",
				Label: "Link",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("Credit_words", inst.Credit_words, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Credit_symbol", inst.Credit_symbol, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.BookTypes {
				for _, target := range src.Credit {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.BookType", "Credit", refNames, formGroup, probe.formStage)
		}
	case *models.Link:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NameXSD", inst.NameXSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Credits {
				for _, target := range src.Link {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Credit", "Link", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
