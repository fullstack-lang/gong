// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/dsm/xsd/tests/books/go/models"
)

// to avoid errors when time and slices packages are not used in the generated code
const _ = time.Nanosecond

var _ = slices.Delete([]string{"a"}, 0, 1)

var _ = log.Panicf

// insertion point
func __gong__New__A_booksFormCallback(
	a_books *models.A_books,
	probe *Probe,
	formGroup *table.FormGroup,
) (a_booksFormCallback *A_booksFormCallback) {
	a_booksFormCallback = new(A_booksFormCallback)
	a_booksFormCallback.probe = probe
	a_booksFormCallback.a_books = a_books
	a_booksFormCallback.formGroup = formGroup

	a_booksFormCallback.CreationMode = (a_books == nil)

	return
}

type A_booksFormCallback struct {
	a_books *models.A_books

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (a_booksFormCallback *A_booksFormCallback) OnSave() {
	a_booksFormCallback.probe.stageOfInterest.Lock()
	defer a_booksFormCallback.probe.stageOfInterest.Unlock()

	// log.Println("A_booksFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	a_booksFormCallback.probe.formStage.Checkout()

	if a_booksFormCallback.a_books == nil {
		a_booksFormCallback.a_books = new(models.A_books).Stage(a_booksFormCallback.probe.stageOfInterest)
	}
	a_books_ := a_booksFormCallback.a_books
	_ = a_books_

	for _, formDiv := range a_booksFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(a_books_.Name), formDiv)
		case "Book":
			instanceSet := *models.GetGongstructInstancesSetFromPointerType[*models.BookType](a_booksFormCallback.probe.stageOfInterest)
			instanceSlice := make([]*models.BookType, 0)

			// make a map of all instances by their ID
			map_id_instances := make(map[uint]*models.BookType)

			for instance := range instanceSet {
				id := models.GetOrderPointerGongstruct(
					a_booksFormCallback.probe.stageOfInterest,
					instance,
				)
				map_id_instances[id] = instance
			}

			rowIDs, err := DecodeStringToIntSlice(formDiv.FormEditAssocButton.AssociationStorage)

			if err != nil {
				log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage)
			}
			map_RowID_ID := GetMap_RowID_ID[*models.BookType](a_booksFormCallback.probe.stageOfInterest)

			for _, rowID := range rowIDs {
				if id, ok := map_RowID_ID[int(rowID)]; ok {
					instanceSlice = append(instanceSlice, map_id_instances[id])
				} else {
					log.Panic("not a good storage", formDiv.FormEditAssocButton.AssociationStorage, "unkown row id", rowID)
				}
			}
			a_books_.Book = instanceSlice

		}
	}

	// manage the suppress operation
	if a_booksFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_books_.Unstage(a_booksFormCallback.probe.stageOfInterest)
	}

	a_booksFormCallback.probe.stageOfInterest.Commit()
	updateProbeTable[*models.A_books](
		a_booksFormCallback.probe,
	)

	// display a new form by reset the form stage
	if a_booksFormCallback.CreationMode || a_booksFormCallback.formGroup.HasSuppressButtonBeenPressed {
		a_booksFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: FormName,
		}).Stage(a_booksFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__A_booksFormCallback(
			nil,
			a_booksFormCallback.probe,
			newFormGroup,
		)
		a_books := new(models.A_books)
		FillUpForm(a_books, newFormGroup, a_booksFormCallback.probe)
		a_booksFormCallback.probe.formStage.Commit()
	}

	a_booksFormCallback.probe.ux_tree()
}
func __gong__New__BookTypeFormCallback(
	booktype *models.BookType,
	probe *Probe,
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
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

		case "A_books:Book":
			// WARNING : this form deals with the N-N association "A_books.Book []*BookType" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of BookType). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.A_books
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "A_books"
				rf.Fieldname = "Book"
				formerAssociationSource := booktype_.GongGetReverseFieldOwner(
					booktypeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.A_books)
					if !ok {
						log.Fatalln("Source of A_books.Book []*BookType, is not an A_books instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Book, booktype_)
					formerSource.Book = slices.Delete(formerSource.Book, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.A_books
			for _a_books := range *models.GetGongstructInstancesSet[models.A_books](booktypeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _a_books.GetName() == newSourceName.GetName() {
					newSource = _a_books // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of A_books.Book []*BookType, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Book = append(newSource.Book, booktype_)
		case "Books:Book":
			// WARNING : this form deals with the N-N association "Books.Book []*BookType" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of BookType). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Books
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Books"
				rf.Fieldname = "Book"
				formerAssociationSource := booktype_.GongGetReverseFieldOwner(
					booktypeFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Books)
					if !ok {
						log.Fatalln("Source of Books.Book []*BookType, is not an Books instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Book, booktype_)
					formerSource.Book = slices.Delete(formerSource.Book, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Books
			for _books := range *models.GetGongstructInstancesSet[models.Books](booktypeFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _books.GetName() == newSourceName.GetName() {
					newSource = _books // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Books.Book []*BookType, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Book = append(newSource.Book, booktype_)
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
		newFormGroup := (&table.FormGroup{
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
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
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
		newFormGroup := (&table.FormGroup{
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
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
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

		case "Credit_words":
			FormDivBasicFieldToField(&(credit_.Credit_words), formDiv)
		case "Credit_symbol":
			FormDivBasicFieldToField(&(credit_.Credit_symbol), formDiv)
		case "BookType:Credit":
			// WARNING : this form deals with the N-N association "BookType.Credit []*Credit" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Credit). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.BookType
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "BookType"
				rf.Fieldname = "Credit"
				formerAssociationSource := credit_.GongGetReverseFieldOwner(
					creditFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.BookType)
					if !ok {
						log.Fatalln("Source of BookType.Credit []*Credit, is not an BookType instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Credit, credit_)
					formerSource.Credit = slices.Delete(formerSource.Credit, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.BookType
			for _booktype := range *models.GetGongstructInstancesSet[models.BookType](creditFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _booktype.GetName() == newSourceName.GetName() {
					newSource = _booktype // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of BookType.Credit []*Credit, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Credit = append(newSource.Credit, credit_)
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
		newFormGroup := (&table.FormGroup{
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
	formGroup *table.FormGroup,
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

	formGroup *table.FormGroup
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
			// WARNING : this form deals with the N-N association "Credit.Link []*Link" but
			// it work only for 1-N associations (TODO: #660, enable this form only for field with //gong:1_N magic code)
			//
			// In many use cases, for instance tree structures, the assocation is semanticaly a 1-N
			// association. For those use cases, it is handy to set the source of the assocation with
			// the form of the target source (when editing an instance of Link). Setting up a value
			// will discard the former value is there is one.
			//
			// Therefore, the forms works only in ONE particular case:
			// - there was no association to this target
			var formerSource *models.Credit
			{
				var rf models.ReverseField
				_ = rf
				rf.GongstructName = "Credit"
				rf.Fieldname = "Link"
				formerAssociationSource := link_.GongGetReverseFieldOwner(
					linkFormCallback.probe.stageOfInterest,
					&rf)

				var ok bool
				if formerAssociationSource != nil {
					formerSource, ok = formerAssociationSource.(*models.Credit)
					if !ok {
						log.Fatalln("Source of Credit.Link []*Link, is not an Credit instance")
					}
				}
			}

			newSourceName := formDiv.FormFields[0].FormFieldSelect.Value

			// case when the user set empty for the source value
			if newSourceName == nil {
				// That could mean we clear the assocation for all source instances
				if formerSource != nil {
					idx := slices.Index(formerSource.Link, link_)
					formerSource.Link = slices.Delete(formerSource.Link, idx, idx+1)
				}
				break // nothing else to do for this field
			}

			// the former source is not empty. the new value could
			// be different but there mught more that one source thet
			// points to this target
			if formerSource != nil {
				break // nothing else to do for this field
			}

			// (2) find the source
			var newSource *models.Credit
			for _credit := range *models.GetGongstructInstancesSet[models.Credit](linkFormCallback.probe.stageOfInterest) {

				// the match is base on the name
				if _credit.GetName() == newSourceName.GetName() {
					newSource = _credit // we have a match
					break
				}
			}
			if newSource == nil {
				log.Println("Source of Credit.Link []*Link, with name", newSourceName, ", does not exist")
				break
			}

			// (3) append the new value to the new source field
			newSource.Link = append(newSource.Link, link_)
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
		newFormGroup := (&table.FormGroup{
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
