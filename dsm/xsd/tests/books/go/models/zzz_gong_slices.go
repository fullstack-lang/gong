// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct BookType
	// insertion point per field
	stage.BookType_Credit_reverseMap = make(map[*Credit]*BookType)
	for booktype := range stage.BookTypes {
		_ = booktype
		for _, _credit := range booktype.Credit {
			stage.BookType_Credit_reverseMap[_credit] = booktype
		}
	}

	// Compute reverse map for named struct Books
	// insertion point per field
	stage.Books_Book_reverseMap = make(map[*BookType]*Books)
	for books := range stage.Bookss {
		_ = books
		for _, _booktype := range books.Book {
			stage.Books_Book_reverseMap[_booktype] = books
		}
	}

	// Compute reverse map for named struct Credit
	// insertion point per field
	stage.Credit_Link_reverseMap = make(map[*Link]*Credit)
	for credit := range stage.Credits {
		_ = credit
		for _, _link := range credit.Link {
			stage.Credit_Link_reverseMap[_link] = credit
		}
	}

	// Compute reverse map for named struct Link
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.BookTypes {
		res = append(res, instance)
	}

	for instance := range stage.Bookss {
		res = append(res, instance)
	}

	for instance := range stage.Credits {
		res = append(res, instance)
	}

	for instance := range stage.Links {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (booktype *BookType) GongCopy() GongstructIF {
	newInstance := new(BookType)
	booktype.CopyBasicFields(newInstance)
	return newInstance
}

func (books *Books) GongCopy() GongstructIF {
	newInstance := new(Books)
	books.CopyBasicFields(newInstance)
	return newInstance
}

func (credit *Credit) GongCopy() GongstructIF {
	newInstance := new(Credit)
	credit.CopyBasicFields(newInstance)
	return newInstance
}

func (link *Link) GongCopy() GongstructIF {
	newInstance := new(Link)
	link.CopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (booktype *BookType) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(booktype).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(booktype), uint64(GetOrderPointerGongstruct(stage, booktype)))
	return
}

func (books *Books) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(books).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(books), uint64(GetOrderPointerGongstruct(stage, books)))
	return
}

func (credit *Credit) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(credit).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(credit), uint64(GetOrderPointerGongstruct(stage, credit)))
	return
}

func (link *Link) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(link).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(link), uint64(GetOrderPointerGongstruct(stage, link)))
	return
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	var booktypes_newInstances []*BookType
	var booktypes_deletedInstances []*BookType

	// parse all staged instances and check if they have a reference
	for booktype := range stage.BookTypes {
		if ref, ok := stage.BookTypes_reference[booktype]; !ok {
			booktypes_newInstances = append(booktypes_newInstances, booktype)
			newInstancesSlice = append(newInstancesSlice, booktype.GongMarshallIdentifier(stage))
			if stage.BookTypes_referenceOrder == nil {
				stage.BookTypes_referenceOrder = make(map[*BookType]uint)
			}
			stage.BookTypes_referenceOrder[booktype] = stage.BookType_stagedOrder[booktype]
			newInstancesReverseSlice = append(newInstancesReverseSlice, booktype.GongMarshallUnstaging(stage))
			// delete(stage.BookTypes_referenceOrder, booktype)
			fieldInitializers, pointersInitializations := booktype.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.BookType_stagedOrder[ref] = stage.BookType_stagedOrder[booktype]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := booktype.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, booktype)
			// delete(stage.BookType_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", booktype.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.BookTypes_reference {
		instance := stage.BookTypes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.BookTypes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			booktypes_deletedInstances = append(booktypes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(booktypes_newInstances)
	lenDeletedInstances += len(booktypes_deletedInstances)
	var bookss_newInstances []*Books
	var bookss_deletedInstances []*Books

	// parse all staged instances and check if they have a reference
	for books := range stage.Bookss {
		if ref, ok := stage.Bookss_reference[books]; !ok {
			bookss_newInstances = append(bookss_newInstances, books)
			newInstancesSlice = append(newInstancesSlice, books.GongMarshallIdentifier(stage))
			if stage.Bookss_referenceOrder == nil {
				stage.Bookss_referenceOrder = make(map[*Books]uint)
			}
			stage.Bookss_referenceOrder[books] = stage.Books_stagedOrder[books]
			newInstancesReverseSlice = append(newInstancesReverseSlice, books.GongMarshallUnstaging(stage))
			// delete(stage.Bookss_referenceOrder, books)
			fieldInitializers, pointersInitializations := books.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Books_stagedOrder[ref] = stage.Books_stagedOrder[books]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := books.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, books)
			// delete(stage.Books_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", books.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Bookss_reference {
		instance := stage.Bookss_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Bookss[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			bookss_deletedInstances = append(bookss_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(bookss_newInstances)
	lenDeletedInstances += len(bookss_deletedInstances)
	var credits_newInstances []*Credit
	var credits_deletedInstances []*Credit

	// parse all staged instances and check if they have a reference
	for credit := range stage.Credits {
		if ref, ok := stage.Credits_reference[credit]; !ok {
			credits_newInstances = append(credits_newInstances, credit)
			newInstancesSlice = append(newInstancesSlice, credit.GongMarshallIdentifier(stage))
			if stage.Credits_referenceOrder == nil {
				stage.Credits_referenceOrder = make(map[*Credit]uint)
			}
			stage.Credits_referenceOrder[credit] = stage.Credit_stagedOrder[credit]
			newInstancesReverseSlice = append(newInstancesReverseSlice, credit.GongMarshallUnstaging(stage))
			// delete(stage.Credits_referenceOrder, credit)
			fieldInitializers, pointersInitializations := credit.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Credit_stagedOrder[ref] = stage.Credit_stagedOrder[credit]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := credit.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, credit)
			// delete(stage.Credit_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", credit.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Credits_reference {
		instance := stage.Credits_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Credits[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			credits_deletedInstances = append(credits_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(credits_newInstances)
	lenDeletedInstances += len(credits_deletedInstances)
	var links_newInstances []*Link
	var links_deletedInstances []*Link

	// parse all staged instances and check if they have a reference
	for link := range stage.Links {
		if ref, ok := stage.Links_reference[link]; !ok {
			links_newInstances = append(links_newInstances, link)
			newInstancesSlice = append(newInstancesSlice, link.GongMarshallIdentifier(stage))
			if stage.Links_referenceOrder == nil {
				stage.Links_referenceOrder = make(map[*Link]uint)
			}
			stage.Links_referenceOrder[link] = stage.Link_stagedOrder[link]
			newInstancesReverseSlice = append(newInstancesReverseSlice, link.GongMarshallUnstaging(stage))
			// delete(stage.Links_referenceOrder, link)
			fieldInitializers, pointersInitializations := link.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Link_stagedOrder[ref] = stage.Link_stagedOrder[link]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := link.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, link)
			// delete(stage.Link_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", link.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Links_reference {
		instance := stage.Links_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Links[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			links_deletedInstances = append(links_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(links_newInstances)
	lenDeletedInstances += len(links_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)
		stage.modified = true
	} else {
		stage.modified = false
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	stage.BookTypes_reference = make(map[*BookType]*BookType)
	stage.BookTypes_referenceOrder = make(map[*BookType]uint) // diff Unstage needs the reference order
	stage.BookTypes_instance = make(map[*BookType]*BookType)
	for instance := range stage.BookTypes {
		_copy := instance.GongCopy().(*BookType)
		stage.BookTypes_reference[instance] = _copy
		stage.BookTypes_instance[_copy] = instance
		stage.BookTypes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Bookss_reference = make(map[*Books]*Books)
	stage.Bookss_referenceOrder = make(map[*Books]uint) // diff Unstage needs the reference order
	stage.Bookss_instance = make(map[*Books]*Books)
	for instance := range stage.Bookss {
		_copy := instance.GongCopy().(*Books)
		stage.Bookss_reference[instance] = _copy
		stage.Bookss_instance[_copy] = instance
		stage.Bookss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Credits_reference = make(map[*Credit]*Credit)
	stage.Credits_referenceOrder = make(map[*Credit]uint) // diff Unstage needs the reference order
	stage.Credits_instance = make(map[*Credit]*Credit)
	for instance := range stage.Credits {
		_copy := instance.GongCopy().(*Credit)
		stage.Credits_reference[instance] = _copy
		stage.Credits_instance[_copy] = instance
		stage.Credits_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Links_reference = make(map[*Link]*Link)
	stage.Links_referenceOrder = make(map[*Link]uint) // diff Unstage needs the reference order
	stage.Links_instance = make(map[*Link]*Link)
	for instance := range stage.Links {
		_copy := instance.GongCopy().(*Link)
		stage.Links_reference[instance] = _copy
		stage.Links_instance[_copy] = instance
		stage.Links_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.BookTypes {
		reference := stage.BookTypes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Bookss {
		reference := stage.Bookss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Credits {
		reference := stage.Credits_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Links {
		reference := stage.Links_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (booktype *BookType) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.BookType_stagedOrder[booktype]; ok {
		return order
	}
	if order, ok := stage.BookTypes_referenceOrder[booktype]; ok {
		return order
	} else {
		log.Printf("instance %p of type BookType was not staged and does not have a reference order", booktype)
		return 0
	}
}

func (books *Books) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Books_stagedOrder[books]; ok {
		return order
	}
	if order, ok := stage.Bookss_referenceOrder[books]; ok {
		return order
	} else {
		log.Printf("instance %p of type Books was not staged and does not have a reference order", books)
		return 0
	}
}

func (credit *Credit) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Credit_stagedOrder[credit]; ok {
		return order
	}
	if order, ok := stage.Credits_referenceOrder[credit]; ok {
		return order
	} else {
		log.Printf("instance %p of type Credit was not staged and does not have a reference order", credit)
		return 0
	}
}

func (link *Link) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Link_stagedOrder[link]; ok {
		return order
	}
	if order, ok := stage.Links_referenceOrder[link]; ok {
		return order
	} else {
		log.Printf("instance %p of type Link was not staged and does not have a reference order", link)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (booktype *BookType) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", booktype.GongGetGongstructName(), booktype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (booktype *BookType) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", booktype.GongGetGongstructName(), booktype.GongGetOrder(stage))
}

func (books *Books) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", books.GongGetGongstructName(), books.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (books *Books) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", books.GongGetGongstructName(), books.GongGetOrder(stage))
}

func (credit *Credit) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", credit.GongGetGongstructName(), credit.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (credit *Credit) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", credit.GongGetGongstructName(), credit.GongGetOrder(stage))
}

func (link *Link) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", link.GongGetGongstructName(), link.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (link *Link) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", link.GongGetGongstructName(), link.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (booktype *BookType) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", booktype.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BookType")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(booktype.Name))
	return
}

func (books *Books) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", books.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Books")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(books.Name))
	return
}

func (credit *Credit) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", credit.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Credit")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(credit.Name))
	return
}

func (link *Link) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", link.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Link")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(link.Name))
	return
}

// insertion point for unstaging
func (booktype *BookType) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", booktype.GongGetReferenceIdentifier(stage))
	return
}

func (books *Books) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", books.GongGetReferenceIdentifier(stage))
	return
}

func (credit *Credit) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", credit.GongGetReferenceIdentifier(stage))
	return
}

func (link *Link) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", link.GongGetReferenceIdentifier(stage))
	return
}

// end of template
