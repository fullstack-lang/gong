// generated code - do not edit
package models

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (booktype *BookType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.BookTypes[booktype]
	return ok
}

func (books *Books) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Bookss[books]
	return ok
}

func (credit *Credit) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Credits[credit]
	return ok
}

func (link *Link) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Links[link]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (booktype *BookType) GongStageBranch(stage *Stage) {

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
	var alreadyCopied bool
	booktypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, booktypeFrom)
	if alreadyCopied {
		return
	}
	booktypeFrom.GongCopyBasicFields(booktypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _credit := range booktypeFrom.Credit {
		booktypeTo.Credit = append(booktypeTo.Credit, GongCopyBranchCredit(mapOrigCopy, _credit))
	}

	return
}

func GongCopyBranchBooks(mapOrigCopy map[any]any, booksFrom *Books) (booksTo *Books) {
	var alreadyCopied bool
	booksTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, booksFrom)
	if alreadyCopied {
		return
	}
	booksFrom.GongCopyBasicFields(booksTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _booktype := range booksFrom.Book {
		booksTo.Book = append(booksTo.Book, GongCopyBranchBookType(mapOrigCopy, _booktype))
	}

	return
}

func GongCopyBranchCredit(mapOrigCopy map[any]any, creditFrom *Credit) (creditTo *Credit) {
	var alreadyCopied bool
	creditTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, creditFrom)
	if alreadyCopied {
		return
	}
	creditFrom.GongCopyBasicFields(creditTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _link := range creditFrom.Link {
		creditTo.Link = append(creditTo.Link, GongCopyBranchLink(mapOrigCopy, _link))
	}

	return
}

func GongCopyBranchLink(mapOrigCopy map[any]any, linkFrom *Link) (linkTo *Link) {
	var alreadyCopied bool
	linkTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, linkFrom)
	if alreadyCopied {
		return
	}
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
	__gong__reconstructSliceOfPointersFromReferences(&reference.Credit, stage.Credits_reference, instance.Credit)
}

func (reference *Books) GongReconstructPointersFromReferences(stage *Stage, instance *Books) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Book, stage.BookTypes_reference, instance.Book)
}

func (reference *Credit) GongReconstructPointersFromReferences(stage *Stage, instance *Credit) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Link, stage.Links_reference, instance.Link)
}

func (reference *Link) GongReconstructPointersFromReferences(stage *Stage, instance *Link) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *BookType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Credit, stage.Credits_instance)
}

func (reference *Books) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Book, stage.BookTypes_instance)
}

func (reference *Credit) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Link, stage.Links_instance)
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
	if ops := __gong__diffSliceOfPointers(stage, booktype, "Credit", booktypeOther.Credit, booktype.Credit); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, books, "Book", booksOther.Book, books.Book); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, credit, "Link", creditOther.Link, credit.Link); ops != "" {
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

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
