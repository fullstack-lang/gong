// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point
func (booktype *BookType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBookTypeCreateCallback != nil {
		stage.OnAfterBookTypeCreateCallback.OnAfterCreate(stage, booktype)
	}
}

func (booktype *BookType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBookTypeUpdateCallback != nil {
		var frontBookType *BookType
		if front != nil {
			frontBookType, _ = front.(*BookType)
		}
		stage.OnAfterBookTypeUpdateCallback.OnAfterUpdate(stage, booktype, frontBookType)
	}
}

func (booktype *BookType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBookTypeDeleteCallback != nil {
		var frontBookType *BookType
		if front != nil {
			frontBookType, _ = front.(*BookType)
		}
		stage.OnAfterBookTypeDeleteCallback.OnAfterDelete(stage, booktype, frontBookType)
	}
}

func (books *Books) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBooksCreateCallback != nil {
		stage.OnAfterBooksCreateCallback.OnAfterCreate(stage, books)
	}
}

func (books *Books) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBooksUpdateCallback != nil {
		var frontBooks *Books
		if front != nil {
			frontBooks, _ = front.(*Books)
		}
		stage.OnAfterBooksUpdateCallback.OnAfterUpdate(stage, books, frontBooks)
	}
}

func (books *Books) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBooksDeleteCallback != nil {
		var frontBooks *Books
		if front != nil {
			frontBooks, _ = front.(*Books)
		}
		stage.OnAfterBooksDeleteCallback.OnAfterDelete(stage, books, frontBooks)
	}
}

func (credit *Credit) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCreditCreateCallback != nil {
		stage.OnAfterCreditCreateCallback.OnAfterCreate(stage, credit)
	}
}

func (credit *Credit) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCreditUpdateCallback != nil {
		var frontCredit *Credit
		if front != nil {
			frontCredit, _ = front.(*Credit)
		}
		stage.OnAfterCreditUpdateCallback.OnAfterUpdate(stage, credit, frontCredit)
	}
}

func (credit *Credit) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCreditDeleteCallback != nil {
		var frontCredit *Credit
		if front != nil {
			frontCredit, _ = front.(*Credit)
		}
		stage.OnAfterCreditDeleteCallback.OnAfterDelete(stage, credit, frontCredit)
	}
}

func (link *Link) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLinkCreateCallback != nil {
		stage.OnAfterLinkCreateCallback.OnAfterCreate(stage, link)
	}
}

func (link *Link) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinkUpdateCallback != nil {
		var frontLink *Link
		if front != nil {
			frontLink, _ = front.(*Link)
		}
		stage.OnAfterLinkUpdateCallback.OnAfterUpdate(stage, link, frontLink)
	}
}

func (link *Link) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinkDeleteCallback != nil {
		var frontLink *Link
		if front != nil {
			frontLink, _ = front.(*Link)
		}
		stage.OnAfterLinkDeleteCallback.OnAfterDelete(stage, link, frontLink)
	}
}

