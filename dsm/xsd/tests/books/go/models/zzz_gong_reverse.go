// generated code - do not edit
package models

// insertion point
func (inst *A_books) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *BookType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "A_books":
		switch reverseField.Fieldname {
		case "Book":
			if _a_books, ok := stage.A_books_Book_reverseMap[inst]; ok {
				res = _a_books.Name
			}
		}
	case "Books":
		switch reverseField.Fieldname {
		case "Book":
			if _books, ok := stage.Books_Book_reverseMap[inst]; ok {
				res = _books.Name
			}
		}
	}
	return
}

func (inst *Books) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Credit) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "BookType":
		switch reverseField.Fieldname {
		case "Credit":
			if _booktype, ok := stage.BookType_Credit_reverseMap[inst]; ok {
				res = _booktype.Name
			}
		}
	}
	return
}

func (inst *Link) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Credit":
		switch reverseField.Fieldname {
		case "Link":
			if _credit, ok := stage.Credit_Link_reverseMap[inst]; ok {
				res = _credit.Name
			}
		}
	}
	return
}

// insertion point
func (inst *A_books) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *BookType) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "A_books":
		switch reverseField.Fieldname {
		case "Book":
			res = stage.A_books_Book_reverseMap[inst]
		}
	case "Books":
		switch reverseField.Fieldname {
		case "Book":
			res = stage.Books_Book_reverseMap[inst]
		}
	}
	return res
}

func (inst *Books) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Credit) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "BookType":
		switch reverseField.Fieldname {
		case "Credit":
			res = stage.BookType_Credit_reverseMap[inst]
		}
	}
	return res
}

func (inst *Link) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Credit":
		switch reverseField.Fieldname {
		case "Link":
			res = stage.Credit_Link_reverseMap[inst]
		}
	}
	return res
}
