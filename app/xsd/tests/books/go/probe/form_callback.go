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

// insertion point
func __gong__New__BookTypeFormCallback(
	booktype *models.BookType,
	probe *Probe,
	formGroup *form.FormGroup,
) (booktypeFormCallback *BookTypeFormCallback) {
	booktypeFormCallback = new(BookTypeFormCallback)
	booktypeFormCallback.probe = probe
	booktypeFormCallback.booktype = booktype
	booktypeFormCallback.formGroup = formGroup

	booktypeFormCallback.CreationMode = (booktype == nil)

	return
}

type BookTypeFormCallback struct {
	booktype *models.BookType

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (booktypeFormCallback *BookTypeFormCallback) OnSave() {
	booktypeFormCallback.probe.stageOfInterest.Lock()
	defer booktypeFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("BookTypeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	booktypeFormCallback.probe.formStage.Checkout()

	if booktypeFormCallback.booktype == nil {
		booktypeFormCallback.booktype = new(models.BookType).Stage(booktypeFormCallback.probe.stageOfInterest)
	}
	booktype_ := booktypeFormCallback.booktype
	_ = booktype_

	for _, formDiv := range booktypeFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(booktype_.Name), formDiv)
		case "Edition":
			FormDivBasicFieldToField(&(booktype_.Edition), formDiv)
		case "Isbn":
			FormDivBasicFieldToField(&(booktype_.Isbn), formDiv)
		case "Bestseller":
			FormDivBasicFieldToField(&(booktype_.Bestseller), formDiv)
		case "Title":
			FormDivBasicFieldToField(&(booktype_.Title), formDiv)
		case "Author":
			FormDivBasicFieldToField(&(booktype_.Author), formDiv)
		case "Year":
			FormDivBasicFieldToField(&(booktype_.Year), formDiv)
		case "Format":
			FormDivBasicFieldToField(&(booktype_.Format), formDiv)
		case "Credit":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Credit](booktypeFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Credit, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Credit)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					booktypeFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Credit](booktypeFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			booktype_.Credit = instanceSlice
			booktypeFormCallback.probe.UpdateSliceOfPointersCallback(booktype_, "Credit", &booktype_.Credit)

		case "Books:Book":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Books instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Books instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Books](booktypeFormCallback.probe.stageOfInterest)
			targetBooksIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetBooksIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Books instances and update their Book slice
			for _books := range *models.GetGongstructInstancesSetFromPointerType[*models.Books](booktypeFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(booktypeFormCallback.probe.stageOfInterest, _books)

				// if Books is selected
				if targetBooksIDs[id] {
					// ensure booktype_ is in _books.Book
					found := false
					for _, _b := range _books.Book {
						if _b == booktype_ {
							found = true
							break
						}
					}
					if !found {
						_books.Book = append(_books.Book, booktype_)
						booktypeFormCallback.probe.UpdateSliceOfPointersCallback(_books, "Book", &_books.Book)
					}
				} else {
					// ensure booktype_ is NOT in _books.Book
					idx := slices.Index(_books.Book, booktype_)
					if idx != -1 {
						_books.Book = slices.Delete(_books.Book, idx, idx+1)
						booktypeFormCallback.probe.UpdateSliceOfPointersCallback(_books, "Book", &_books.Book)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if booktypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		booktype_.Unstage(booktypeFormCallback.probe.stageOfInterest)
	}

	booktypeFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.BookType](
		booktypeFormCallback.probe,
	)

	// display a new form by reset the form stage
	if booktypeFormCallback.CreationMode || booktypeFormCallback.formGroup.HasSuppressButtonBeenPressed {
		booktypeFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(booktypeFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BookTypeFormCallback(
			nil,
			booktypeFormCallback.probe,
			newFormGroup,
		)
		booktype := new(models.BookType)
		FillUpForm(booktype, newFormGroup, booktypeFormCallback.probe)
		booktypeFormCallback.probe.formStage.Commit()
	}

	booktypeFormCallback.probe.ux_tree()
}
func __gong__New__BooksFormCallback(
	books *models.Books,
	probe *Probe,
	formGroup *form.FormGroup,
) (booksFormCallback *BooksFormCallback) {
	booksFormCallback = new(BooksFormCallback)
	booksFormCallback.probe = probe
	booksFormCallback.books = books
	booksFormCallback.formGroup = formGroup

	booksFormCallback.CreationMode = (books == nil)

	return
}

type BooksFormCallback struct {
	books *models.Books

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (booksFormCallback *BooksFormCallback) OnSave() {
	booksFormCallback.probe.stageOfInterest.Lock()
	defer booksFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("BooksFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	booksFormCallback.probe.formStage.Checkout()

	if booksFormCallback.books == nil {
		booksFormCallback.books = new(models.Books).Stage(booksFormCallback.probe.stageOfInterest)
	}
	books_ := booksFormCallback.books
	_ = books_

	for _, formDiv := range booksFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(books_.Name), formDiv)
		case "Book":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.BookType](booksFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.BookType, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.BookType)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					booksFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.BookType](booksFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			books_.Book = instanceSlice
			booksFormCallback.probe.UpdateSliceOfPointersCallback(books_, "Book", &books_.Book)

		}
	}

	// manage the suppress operation
	if booksFormCallback.formGroup.HasSuppressButtonBeenPressed {
		books_.Unstage(booksFormCallback.probe.stageOfInterest)
	}

	booksFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Books](
		booksFormCallback.probe,
	)

	// display a new form by reset the form stage
	if booksFormCallback.CreationMode || booksFormCallback.formGroup.HasSuppressButtonBeenPressed {
		booksFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(booksFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BooksFormCallback(
			nil,
			booksFormCallback.probe,
			newFormGroup,
		)
		books := new(models.Books)
		FillUpForm(books, newFormGroup, booksFormCallback.probe)
		booksFormCallback.probe.formStage.Commit()
	}

	booksFormCallback.probe.ux_tree()
}
func __gong__New__CreditFormCallback(
	credit *models.Credit,
	probe *Probe,
	formGroup *form.FormGroup,
) (creditFormCallback *CreditFormCallback) {
	creditFormCallback = new(CreditFormCallback)
	creditFormCallback.probe = probe
	creditFormCallback.credit = credit
	creditFormCallback.formGroup = formGroup

	creditFormCallback.CreationMode = (credit == nil)

	return
}

type CreditFormCallback struct {
	credit *models.Credit

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (creditFormCallback *CreditFormCallback) OnSave() {
	creditFormCallback.probe.stageOfInterest.Lock()
	defer creditFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("CreditFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	creditFormCallback.probe.formStage.Checkout()

	if creditFormCallback.credit == nil {
		creditFormCallback.credit = new(models.Credit).Stage(creditFormCallback.probe.stageOfInterest)
	}
	credit_ := creditFormCallback.credit
	_ = credit_

	for _, formDiv := range creditFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(credit_.Name), formDiv)
		case "Page":
			FormDivBasicFieldToField(&(credit_.Page), formDiv)
		case "Credit_type":
			FormDivBasicFieldToField(&(credit_.Credit_type), formDiv)
		case "Link":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.Link](creditFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.Link, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.Link)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					creditFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.Link](creditFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			credit_.Link = instanceSlice
			creditFormCallback.probe.UpdateSliceOfPointersCallback(credit_, "Link", &credit_.Link)

		case "Credit_words":
			FormDivBasicFieldToField(&(credit_.Credit_words), formDiv)
		case "Credit_symbol":
			FormDivBasicFieldToField(&(credit_.Credit_symbol), formDiv)
		case "BookType:Credit":
			// 1. Decode the AssociationStorage which contains the rowIDs of the BookType instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target BookType instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.BookType](creditFormCallback.probe.stageOfInterest)
			targetBookTypeIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetBookTypeIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all BookType instances and update their Credit slice
			for _booktype := range *models.GetGongstructInstancesSetFromPointerType[*models.BookType](creditFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(creditFormCallback.probe.stageOfInterest, _booktype)

				// if BookType is selected
				if targetBookTypeIDs[id] {
					// ensure credit_ is in _booktype.Credit
					found := false
					for _, _b := range _booktype.Credit {
						if _b == credit_ {
							found = true
							break
						}
					}
					if !found {
						_booktype.Credit = append(_booktype.Credit, credit_)
						creditFormCallback.probe.UpdateSliceOfPointersCallback(_booktype, "Credit", &_booktype.Credit)
					}
				} else {
					// ensure credit_ is NOT in _booktype.Credit
					idx := slices.Index(_booktype.Credit, credit_)
					if idx != -1 {
						_booktype.Credit = slices.Delete(_booktype.Credit, idx, idx+1)
						creditFormCallback.probe.UpdateSliceOfPointersCallback(_booktype, "Credit", &_booktype.Credit)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if creditFormCallback.formGroup.HasSuppressButtonBeenPressed {
		credit_.Unstage(creditFormCallback.probe.stageOfInterest)
	}

	creditFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Credit](
		creditFormCallback.probe,
	)

	// display a new form by reset the form stage
	if creditFormCallback.CreationMode || creditFormCallback.formGroup.HasSuppressButtonBeenPressed {
		creditFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(creditFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__CreditFormCallback(
			nil,
			creditFormCallback.probe,
			newFormGroup,
		)
		credit := new(models.Credit)
		FillUpForm(credit, newFormGroup, creditFormCallback.probe)
		creditFormCallback.probe.formStage.Commit()
	}

	creditFormCallback.probe.ux_tree()
}
func __gong__New__LinkFormCallback(
	link *models.Link,
	probe *Probe,
	formGroup *form.FormGroup,
) (linkFormCallback *LinkFormCallback) {
	linkFormCallback = new(LinkFormCallback)
	linkFormCallback.probe = probe
	linkFormCallback.link = link
	linkFormCallback.formGroup = formGroup

	linkFormCallback.CreationMode = (link == nil)

	return
}

type LinkFormCallback struct {
	link *models.Link

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *form.FormGroup
}

func (linkFormCallback *LinkFormCallback) OnSave() {
	linkFormCallback.probe.stageOfInterest.Lock()
	defer linkFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("LinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linkFormCallback.probe.formStage.Checkout()

	if linkFormCallback.link == nil {
		linkFormCallback.link = new(models.Link).Stage(linkFormCallback.probe.stageOfInterest)
	}
	link_ := linkFormCallback.link
	_ = link_

	for _, formDiv := range linkFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(link_.Name), formDiv)
		case "NameXSD":
			FormDivBasicFieldToField(&(link_.NameXSD), formDiv)
		case "EnclosedText":
			FormDivBasicFieldToField(&(link_.EnclosedText), formDiv)
		case "Credit:Link":
			// 1. Decode the AssociationStorage which contains the rowIDs of the Credit instances
			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)
			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}

			// 2. Build a map of target Credit instances by their ID
			map_RowID_ID := GetMap_RowID_ID[*models.Credit](linkFormCallback.probe.stageOfInterest)
			targetCreditIDs := make(map[uint]bool)
			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					targetCreditIDs[id] = true
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unknown row id", rowID)
				}
			}

			// 3. Iterate over all Credit instances and update their Link slice
			for _credit := range *models.GetGongstructInstancesSetFromPointerType[*models.Credit](linkFormCallback.probe.stageOfInterest) {
				id := models.GetOrderPointerGongstruct(linkFormCallback.probe.stageOfInterest, _credit)

				// if Credit is selected
				if targetCreditIDs[id] {
					// ensure link_ is in _credit.Link
					found := false
					for _, _b := range _credit.Link {
						if _b == link_ {
							found = true
							break
						}
					}
					if !found {
						_credit.Link = append(_credit.Link, link_)
						linkFormCallback.probe.UpdateSliceOfPointersCallback(_credit, "Link", &_credit.Link)
					}
				} else {
					// ensure link_ is NOT in _credit.Link
					idx := slices.Index(_credit.Link, link_)
					if idx != -1 {
						_credit.Link = slices.Delete(_credit.Link, idx, idx+1)
						linkFormCallback.probe.UpdateSliceOfPointersCallback(_credit, "Link", &_credit.Link)
					}
				}
			}
		}
	}

	// manage the suppress operation
	if linkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		link_.Unstage(linkFormCallback.probe.stageOfInterest)
	}

	linkFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.Link](
		linkFormCallback.probe,
	)

	// display a new form by reset the form stage
	if linkFormCallback.CreationMode || linkFormCallback.formGroup.HasSuppressButtonBeenPressed {
		linkFormCallback.probe.formStage.Reset()
		newFormGroup := (&form.FormGroup{
			Name: FormName,
		}).Stage(linkFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__LinkFormCallback(
			nil,
			linkFormCallback.probe,
			newFormGroup,
		)
		link := new(models.Link)
		FillUpForm(link, newFormGroup, linkFormCallback.probe)
		linkFormCallback.probe.formStage.Commit()
	}

	linkFormCallback.probe.ux_tree()
}
