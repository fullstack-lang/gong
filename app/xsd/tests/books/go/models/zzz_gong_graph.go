// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *BookType:
		ok = stage.IsStagedBookType(target)

	case *Books:
		ok = stage.IsStagedBooks(target)

	case *Credit:
		ok = stage.IsStagedCredit(target)

	case *Link:
		ok = stage.IsStagedLink(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *BookType:
		ok = stage.IsStagedBookType(target)

	case *Books:
		ok = stage.IsStagedBooks(target)

	case *Credit:
		ok = stage.IsStagedCredit(target)

	case *Link:
		ok = stage.IsStagedLink(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedBookType(booktype *BookType) (ok bool) {

	_, ok = stage.BookTypes[booktype]

	return
}

func (stage *Stage) IsStagedBooks(books *Books) (ok bool) {

	_, ok = stage.Bookss[books]

	return
}

func (stage *Stage) IsStagedCredit(credit *Credit) (ok bool) {

	_, ok = stage.Credits[credit]

	return
}

func (stage *Stage) IsStagedLink(link *Link) (ok bool) {

	_, ok = stage.Links[link]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *BookType:
		stage.StageBranchBookType(target)

	case *Books:
		stage.StageBranchBooks(target)

	case *Credit:
		stage.StageBranchCredit(target)

	case *Link:
		stage.StageBranchLink(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchBookType(booktype *BookType) {

	// check if instance is already staged
	if IsStaged(stage, booktype) {
		return
	}

	booktype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _credit := range booktype.Credit {
		StageBranch(stage, _credit)
	}

}

func (stage *Stage) StageBranchBooks(books *Books) {

	// check if instance is already staged
	if IsStaged(stage, books) {
		return
	}

	books.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _booktype := range books.Book {
		StageBranch(stage, _booktype)
	}

}

func (stage *Stage) StageBranchCredit(credit *Credit) {

	// check if instance is already staged
	if IsStaged(stage, credit) {
		return
	}

	credit.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _link := range credit.Link {
		StageBranch(stage, _link)
	}

}

func (stage *Stage) StageBranchLink(link *Link) {

	// check if instance is already staged
	if IsStaged(stage, link) {
		return
	}

	link.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *BookType:
		toT := CopyBranchBookType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Books:
		toT := CopyBranchBooks(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Credit:
		toT := CopyBranchCredit(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Link:
		toT := CopyBranchLink(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchBookType(mapOrigCopy map[any]any, booktypeFrom *BookType) (booktypeTo *BookType) {

	// booktypeFrom has already been copied
	if _booktypeTo, ok := mapOrigCopy[booktypeFrom]; ok {
		booktypeTo = _booktypeTo.(*BookType)
		return
	}

	booktypeTo = new(BookType)
	mapOrigCopy[booktypeFrom] = booktypeTo
	booktypeFrom.CopyBasicFields(booktypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _credit := range booktypeFrom.Credit {
		booktypeTo.Credit = append(booktypeTo.Credit, CopyBranchCredit(mapOrigCopy, _credit))
	}

	return
}

func CopyBranchBooks(mapOrigCopy map[any]any, booksFrom *Books) (booksTo *Books) {

	// booksFrom has already been copied
	if _booksTo, ok := mapOrigCopy[booksFrom]; ok {
		booksTo = _booksTo.(*Books)
		return
	}

	booksTo = new(Books)
	mapOrigCopy[booksFrom] = booksTo
	booksFrom.CopyBasicFields(booksTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _booktype := range booksFrom.Book {
		booksTo.Book = append(booksTo.Book, CopyBranchBookType(mapOrigCopy, _booktype))
	}

	return
}

func CopyBranchCredit(mapOrigCopy map[any]any, creditFrom *Credit) (creditTo *Credit) {

	// creditFrom has already been copied
	if _creditTo, ok := mapOrigCopy[creditFrom]; ok {
		creditTo = _creditTo.(*Credit)
		return
	}

	creditTo = new(Credit)
	mapOrigCopy[creditFrom] = creditTo
	creditFrom.CopyBasicFields(creditTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _link := range creditFrom.Link {
		creditTo.Link = append(creditTo.Link, CopyBranchLink(mapOrigCopy, _link))
	}

	return
}

func CopyBranchLink(mapOrigCopy map[any]any, linkFrom *Link) (linkTo *Link) {

	// linkFrom has already been copied
	if _linkTo, ok := mapOrigCopy[linkFrom]; ok {
		linkTo = _linkTo.(*Link)
		return
	}

	linkTo = new(Link)
	mapOrigCopy[linkFrom] = linkTo
	linkFrom.CopyBasicFields(linkTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *BookType:
		stage.UnstageBranchBookType(target)

	case *Books:
		stage.UnstageBranchBooks(target)

	case *Credit:
		stage.UnstageBranchCredit(target)

	case *Link:
		stage.UnstageBranchLink(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchBookType(booktype *BookType) {

	// check if instance is already staged
	if !IsStaged(stage, booktype) {
		return
	}

	booktype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _credit := range booktype.Credit {
		UnstageBranch(stage, _credit)
	}

}

func (stage *Stage) UnstageBranchBooks(books *Books) {

	// check if instance is already staged
	if !IsStaged(stage, books) {
		return
	}

	books.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _booktype := range books.Book {
		UnstageBranch(stage, _booktype)
	}

}

func (stage *Stage) UnstageBranchCredit(credit *Credit) {

	// check if instance is already staged
	if !IsStaged(stage, credit) {
		return
	}

	credit.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _link := range credit.Link {
		UnstageBranch(stage, _link)
	}

}

func (stage *Stage) UnstageBranchLink(link *Link) {

	// check if instance is already staged
	if !IsStaged(stage, link) {
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
		ops := Diff(stage, booktype, booktypeOther, "Credit", booktypeOther.Credit, booktype.Credit)
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
		ops := Diff(stage, books, booksOther, "Book", booksOther.Book, books.Book)
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
		ops := Diff(stage, credit, creditOther, "Link", creditOther.Link, credit.Link)
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

// Diff returns the sequence of operations to transform oldSlice into newSlice.
// It requires type T to be comparable (e.g., pointers, ints, strings).
func Diff[T1, T2 PointerToGongstruct](stage *Stage, a, b T1, fieldName string, oldSlice, newSlice []T2) (ops string) {
	m, n := len(oldSlice), len(newSlice)

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if oldSlice[i] == newSlice[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if oldSlice[i-1] == newSlice[j-1] {
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

	// Create a temporary view of what's left after deletions for tracking matches
	var currentLCS []T2
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			currentLCS = append(currentLCS, oldSlice[k])
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k, targetVal := range newSlice {
		if lcsIdx < len(currentLCS) && currentLCS[lcsIdx] == targetVal {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, targetVal.GongGetIdentifier(stage))
		}
	}

	return ops
}
