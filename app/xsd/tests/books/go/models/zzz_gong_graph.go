// generated code - do not edit
package models

import "fmt"

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (booktype *BookType) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.BookTypes[booktype]

	return
}

func (stage *Stage) IsStagedBookType(booktype *BookType) (ok bool) {

	return booktype.GongIsStaged(stage)
}

func (books *Books) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Bookss[books]

	return
}

func (stage *Stage) IsStagedBooks(books *Books) (ok bool) {

	return books.GongIsStaged(stage)
}

func (credit *Credit) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Credits[credit]

	return
}

func (stage *Stage) IsStagedCredit(credit *Credit) (ok bool) {

	return credit.GongIsStaged(stage)
}

func (link *Link) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Links[link]

	return
}

func (stage *Stage) IsStagedLink(link *Link) (ok bool) {

	return link.GongIsStaged(stage)
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (booktype *BookType) GongStageBranch(stage *Stage) {
	stage.StageBranchBookType(booktype)
}

func (stage *Stage) StageBranchBookType(booktype *BookType) {

	// check if instance is already staged
	if stage.IsStaged(booktype) {
		return
	}

	booktype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _credit := range booktype.Credit {
		stage.StageBranch(_credit)
	}

}

func (books *Books) GongStageBranch(stage *Stage) {
	stage.StageBranchBooks(books)
}

func (stage *Stage) StageBranchBooks(books *Books) {

	// check if instance is already staged
	if stage.IsStaged(books) {
		return
	}

	books.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _booktype := range books.Book {
		stage.StageBranch(_booktype)
	}

}

func (credit *Credit) GongStageBranch(stage *Stage) {
	stage.StageBranchCredit(credit)
}

func (stage *Stage) StageBranchCredit(credit *Credit) {

	// check if instance is already staged
	if stage.IsStaged(credit) {
		return
	}

	credit.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _link := range credit.Link {
		stage.StageBranch(_link)
	}

}

func (link *Link) GongStageBranch(stage *Stage) {
	stage.StageBranchLink(link)
}

func (stage *Stage) StageBranchLink(link *Link) {

	// check if instance is already staged
	if stage.IsStaged(link) {
		return
	}

	link.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *BookType:
		toT := GongCopyBranchBookType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Books:
		toT := GongCopyBranchBooks(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Credit:
		toT := GongCopyBranchCredit(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Link:
		toT := GongCopyBranchLink(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchBookType(mapOrigCopy map[any]any, booktypeFrom *BookType) (booktypeTo *BookType) {

	// booktypeFrom has already been copied
	if _booktypeTo, ok := mapOrigCopy[booktypeFrom]; ok {
		booktypeTo = _booktypeTo.(*BookType)
		return
	}

	booktypeTo = new(BookType)
	mapOrigCopy[booktypeFrom] = booktypeTo
	booktypeFrom.GongCopyBasicFields(booktypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _credit := range booktypeFrom.Credit {
		booktypeTo.Credit = append(booktypeTo.Credit, GongCopyBranchCredit(mapOrigCopy, _credit))
	}

	return
}

func GongCopyBranchBooks(mapOrigCopy map[any]any, booksFrom *Books) (booksTo *Books) {

	// booksFrom has already been copied
	if _booksTo, ok := mapOrigCopy[booksFrom]; ok {
		booksTo = _booksTo.(*Books)
		return
	}

	booksTo = new(Books)
	mapOrigCopy[booksFrom] = booksTo
	booksFrom.GongCopyBasicFields(booksTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _booktype := range booksFrom.Book {
		booksTo.Book = append(booksTo.Book, GongCopyBranchBookType(mapOrigCopy, _booktype))
	}

	return
}

func GongCopyBranchCredit(mapOrigCopy map[any]any, creditFrom *Credit) (creditTo *Credit) {

	// creditFrom has already been copied
	if _creditTo, ok := mapOrigCopy[creditFrom]; ok {
		creditTo = _creditTo.(*Credit)
		return
	}

	creditTo = new(Credit)
	mapOrigCopy[creditFrom] = creditTo
	creditFrom.GongCopyBasicFields(creditTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _link := range creditFrom.Link {
		creditTo.Link = append(creditTo.Link, GongCopyBranchLink(mapOrigCopy, _link))
	}

	return
}

func GongCopyBranchLink(mapOrigCopy map[any]any, linkFrom *Link) (linkTo *Link) {

	// linkFrom has already been copied
	if _linkTo, ok := mapOrigCopy[linkFrom]; ok {
		linkTo = _linkTo.(*Link)
		return
	}

	linkTo = new(Link)
	mapOrigCopy[linkFrom] = linkTo
	linkFrom.GongCopyBasicFields(linkTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// insertion point for unstage branch per struct
func (booktype *BookType) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchBookType(booktype)
}

func (stage *Stage) UnstageBranchBookType(booktype *BookType) {

	// check if instance is already staged
	if !stage.IsStaged(booktype) {
		return
	}

	booktype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _credit := range booktype.Credit {
		stage.UnstageBranch(_credit)
	}

}

func (books *Books) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchBooks(books)
}

func (stage *Stage) UnstageBranchBooks(books *Books) {

	// check if instance is already staged
	if !stage.IsStaged(books) {
		return
	}

	books.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _booktype := range books.Book {
		stage.UnstageBranch(_booktype)
	}

}

func (credit *Credit) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchCredit(credit)
}

func (stage *Stage) UnstageBranchCredit(credit *Credit) {

	// check if instance is already staged
	if !stage.IsStaged(credit) {
		return
	}

	credit.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _link := range credit.Link {
		stage.UnstageBranch(_link)
	}

}

func (link *Link) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchLink(link)
}

func (stage *Stage) UnstageBranchLink(link *Link) {

	// check if instance is already staged
	if !stage.IsStaged(link) {
		return
	}

	link.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *BookType) GongReconstructPointersFromReferences(stage *Stage, instance *BookType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Credit = reference.Credit[:0]
	for _, _b := range instance.Credit {
		reference.Credit = append(reference.Credit, stage.Credits_reference[_b])
	}
}

func (reference *Books) GongReconstructPointersFromReferences(stage *Stage, instance *Books) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Book = reference.Book[:0]
	for _, _b := range instance.Book {
		reference.Book = append(reference.Book, stage.BookTypes_reference[_b])
	}
}

func (reference *Credit) GongReconstructPointersFromReferences(stage *Stage, instance *Credit) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Link = reference.Link[:0]
	for _, _b := range instance.Link {
		reference.Link = append(reference.Link, stage.Links_reference[_b])
	}
}

func (reference *Link) GongReconstructPointersFromReferences(stage *Stage, instance *Link) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *BookType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Credit []*Credit
	for _, _reference := range reference.Credit {
		if _instance, ok := stage.Credits_instance[_reference]; ok {
			_Credit = append(_Credit, _instance)
		}
	}
	reference.Credit = _Credit
}

func (reference *Books) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Book []*BookType
	for _, _reference := range reference.Book {
		if _instance, ok := stage.BookTypes_instance[_reference]; ok {
			_Book = append(_Book, _instance)
		}
	}
	reference.Book = _Book
}

func (reference *Credit) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Link []*Link
	for _, _reference := range reference.Link {
		if _instance, ok := stage.Links_instance[_reference]; ok {
			_Link = append(_Link, _instance)
		}
	}
	reference.Link = _Link
}

func (reference *Link) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (booktype *BookType) GongDiff(stage *Stage, booktypeOther *BookType) (diffs []string) {
	// insertion point for field diffs
	if booktype.Name != booktypeOther.Name {
		diffs = append(diffs, booktype.GongMarshallField(stage, "Name"))
	}
	if booktype.Edition != booktypeOther.Edition {
		diffs = append(diffs, booktype.GongMarshallField(stage, "Edition"))
	}
	if booktype.Isbn != booktypeOther.Isbn {
		diffs = append(diffs, booktype.GongMarshallField(stage, "Isbn"))
	}
	if booktype.Bestseller != booktypeOther.Bestseller {
		diffs = append(diffs, booktype.GongMarshallField(stage, "Bestseller"))
	}
	if booktype.Title != booktypeOther.Title {
		diffs = append(diffs, booktype.GongMarshallField(stage, "Title"))
	}
	if booktype.Author != booktypeOther.Author {
		diffs = append(diffs, booktype.GongMarshallField(stage, "Author"))
	}
	if booktype.Year != booktypeOther.Year {
		diffs = append(diffs, booktype.GongMarshallField(stage, "Year"))
	}
	if booktype.Format != booktypeOther.Format {
		diffs = append(diffs, booktype.GongMarshallField(stage, "Format"))
	}
	CreditDifferent := false
	if len(booktype.Credit) != len(booktypeOther.Credit) {
		CreditDifferent = true
	} else {
		for i := range booktype.Credit {
			if (booktype.Credit[i] == nil) != (booktypeOther.Credit[i] == nil) {
				CreditDifferent = true
				break
			} else if booktype.Credit[i] != nil && booktypeOther.Credit[i] != nil {
				// this is a pointer comparaison
				if booktype.Credit[i] != booktypeOther.Credit[i] {
					CreditDifferent = true
					break
				}
			}
		}
	}
	if CreditDifferent {
		ops := stage.Diff(
			booktype,
			"Credit",
			len(booktypeOther.Credit),
			len(booktype.Credit),
			func(i, j int) bool {
				return booktypeOther.Credit[i] == booktype.Credit[j]
			},
			func(j int) string {
				return booktype.Credit[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (books *Books) GongDiff(stage *Stage, booksOther *Books) (diffs []string) {
	// insertion point for field diffs
	if books.Name != booksOther.Name {
		diffs = append(diffs, books.GongMarshallField(stage, "Name"))
	}
	BookDifferent := false
	if len(books.Book) != len(booksOther.Book) {
		BookDifferent = true
	} else {
		for i := range books.Book {
			if (books.Book[i] == nil) != (booksOther.Book[i] == nil) {
				BookDifferent = true
				break
			} else if books.Book[i] != nil && booksOther.Book[i] != nil {
				// this is a pointer comparaison
				if books.Book[i] != booksOther.Book[i] {
					BookDifferent = true
					break
				}
			}
		}
	}
	if BookDifferent {
		ops := stage.Diff(
			books,
			"Book",
			len(booksOther.Book),
			len(books.Book),
			func(i, j int) bool {
				return booksOther.Book[i] == books.Book[j]
			},
			func(j int) string {
				return books.Book[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (credit *Credit) GongDiff(stage *Stage, creditOther *Credit) (diffs []string) {
	// insertion point for field diffs
	if credit.Name != creditOther.Name {
		diffs = append(diffs, credit.GongMarshallField(stage, "Name"))
	}
	if credit.Page != creditOther.Page {
		diffs = append(diffs, credit.GongMarshallField(stage, "Page"))
	}
	if credit.Credit_type != creditOther.Credit_type {
		diffs = append(diffs, credit.GongMarshallField(stage, "Credit_type"))
	}
	LinkDifferent := false
	if len(credit.Link) != len(creditOther.Link) {
		LinkDifferent = true
	} else {
		for i := range credit.Link {
			if (credit.Link[i] == nil) != (creditOther.Link[i] == nil) {
				LinkDifferent = true
				break
			} else if credit.Link[i] != nil && creditOther.Link[i] != nil {
				// this is a pointer comparaison
				if credit.Link[i] != creditOther.Link[i] {
					LinkDifferent = true
					break
				}
			}
		}
	}
	if LinkDifferent {
		ops := stage.Diff(
			credit,
			"Link",
			len(creditOther.Link),
			len(credit.Link),
			func(i, j int) bool {
				return creditOther.Link[i] == credit.Link[j]
			},
			func(j int) string {
				return credit.Link[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if credit.Credit_words != creditOther.Credit_words {
		diffs = append(diffs, credit.GongMarshallField(stage, "Credit_words"))
	}
	if credit.Credit_symbol != creditOther.Credit_symbol {
		diffs = append(diffs, credit.GongMarshallField(stage, "Credit_symbol"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (link *Link) GongDiff(stage *Stage, linkOther *Link) (diffs []string) {
	// insertion point for field diffs
	if link.Name != linkOther.Name {
		diffs = append(diffs, link.GongMarshallField(stage, "Name"))
	}
	if link.NameXSD != linkOther.NameXSD {
		diffs = append(diffs, link.GongMarshallField(stage, "NameXSD"))
	}
	if link.EnclosedText != linkOther.EnclosedText {
		diffs = append(diffs, link.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}
