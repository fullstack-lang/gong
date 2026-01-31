// generated code - do not edit
package models

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

var __GongSliceTemplate_time__dummyDeclaration time.Duration
var _ = __GongSliceTemplate_time__dummyDeclaration

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Chapter
	// insertion point per field
	stage.Chapter_Pages_reverseMap = make(map[*Page]*Chapter)
	for chapter := range stage.Chapters {
		_ = chapter
		for _, _page := range chapter.Pages {
			stage.Chapter_Pages_reverseMap[_page] = chapter
		}
	}

	// Compute reverse map for named struct Content
	// insertion point per field
	stage.Content_Chapters_reverseMap = make(map[*Chapter]*Content)
	for content := range stage.Contents {
		_ = content
		for _, _chapter := range content.Chapters {
			stage.Content_Chapters_reverseMap[_chapter] = content
		}
	}

	// Compute reverse map for named struct Page
	// insertion point per field

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.Chapters {
		res = append(res, instance)
	}

	for instance := range stage.Contents {
		res = append(res, instance)
	}

	for instance := range stage.Pages {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (chapter *Chapter) GongCopy() GongstructIF {
	newInstance := *chapter
	return &newInstance
}

func (content *Content) GongCopy() GongstructIF {
	newInstance := *content
	return &newInstance
}

func (page *Page) GongCopy() GongstructIF {
	newInstance := *page
	return &newInstance
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
	var chapters_newInstances []*Chapter
	var chapters_deletedInstances []*Chapter

	// parse all staged instances and check if they have a reference
	for chapter := range stage.Chapters {
		if ref, ok := stage.Chapters_reference[chapter]; !ok {
			chapters_newInstances = append(chapters_newInstances, chapter)
			newInstancesSlice = append(newInstancesSlice, chapter.GongMarshallIdentifier(stage))
			if stage.Chapters_referenceOrder == nil {
				stage.Chapters_referenceOrder = make(map[*Chapter]uint)
			}
			stage.Chapters_referenceOrder[chapter] = stage.ChapterMap_Staged_Order[chapter]
			newInstancesReverseSlice = append(newInstancesReverseSlice, chapter.GongMarshallUnstaging(stage))
			delete(stage.Chapters_referenceOrder, chapter)
			fieldInitializers, pointersInitializations := chapter.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ChapterMap_Staged_Order[ref] = stage.ChapterMap_Staged_Order[chapter]
			diffs := chapter.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, chapter)
			delete(stage.ChapterMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", chapter.GetName())
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
	for ref := range stage.Chapters_reference {
		if _, ok := stage.Chapters[ref]; !ok {
			chapters_deletedInstances = append(chapters_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(chapters_newInstances)
	lenDeletedInstances += len(chapters_deletedInstances)
	var contents_newInstances []*Content
	var contents_deletedInstances []*Content

	// parse all staged instances and check if they have a reference
	for content := range stage.Contents {
		if ref, ok := stage.Contents_reference[content]; !ok {
			contents_newInstances = append(contents_newInstances, content)
			newInstancesSlice = append(newInstancesSlice, content.GongMarshallIdentifier(stage))
			if stage.Contents_referenceOrder == nil {
				stage.Contents_referenceOrder = make(map[*Content]uint)
			}
			stage.Contents_referenceOrder[content] = stage.ContentMap_Staged_Order[content]
			newInstancesReverseSlice = append(newInstancesReverseSlice, content.GongMarshallUnstaging(stage))
			delete(stage.Contents_referenceOrder, content)
			fieldInitializers, pointersInitializations := content.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ContentMap_Staged_Order[ref] = stage.ContentMap_Staged_Order[content]
			diffs := content.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, content)
			delete(stage.ContentMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", content.GetName())
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
	for ref := range stage.Contents_reference {
		if _, ok := stage.Contents[ref]; !ok {
			contents_deletedInstances = append(contents_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(contents_newInstances)
	lenDeletedInstances += len(contents_deletedInstances)
	var pages_newInstances []*Page
	var pages_deletedInstances []*Page

	// parse all staged instances and check if they have a reference
	for page := range stage.Pages {
		if ref, ok := stage.Pages_reference[page]; !ok {
			pages_newInstances = append(pages_newInstances, page)
			newInstancesSlice = append(newInstancesSlice, page.GongMarshallIdentifier(stage))
			if stage.Pages_referenceOrder == nil {
				stage.Pages_referenceOrder = make(map[*Page]uint)
			}
			stage.Pages_referenceOrder[page] = stage.PageMap_Staged_Order[page]
			newInstancesReverseSlice = append(newInstancesReverseSlice, page.GongMarshallUnstaging(stage))
			delete(stage.Pages_referenceOrder, page)
			fieldInitializers, pointersInitializations := page.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PageMap_Staged_Order[ref] = stage.PageMap_Staged_Order[page]
			diffs := page.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, page)
			delete(stage.PageMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", page.GetName())
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
	for ref := range stage.Pages_reference {
		if _, ok := stage.Pages[ref]; !ok {
			pages_deletedInstances = append(pages_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(pages_newInstances)
	lenDeletedInstances += len(pages_deletedInstances)

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

		if stage.GetProbeIF() != nil {
			var mergedCommits string
			for _, commit := range stage.forwardCommits {
				mergedCommits += commit
			}
			stage.GetProbeIF().AddNotification(
				time.Now(),
				"	// Forward commits:\n"+
					mergedCommits,
			)

			var reverseMergedCommits string
			for _, reverserCommit := range stage.backwardCommits {
				reverseMergedCommits += reverserCommit
			}
			stage.GetProbeIF().AddNotification(
				time.Now(),
				"	// Backward commits:\n"+
					reverseMergedCommits,
			)

			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Chapters_reference = make(map[*Chapter]*Chapter)
	stage.Chapters_referenceOrder = make(map[*Chapter]uint) // diff Unstage needs the reference order
	for instance := range stage.Chapters {
		stage.Chapters_reference[instance] = instance.GongCopy().(*Chapter)
		stage.Chapters_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Contents_reference = make(map[*Content]*Content)
	stage.Contents_referenceOrder = make(map[*Content]uint) // diff Unstage needs the reference order
	for instance := range stage.Contents {
		stage.Contents_reference[instance] = instance.GongCopy().(*Content)
		stage.Contents_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Pages_reference = make(map[*Page]*Page)
	stage.Pages_referenceOrder = make(map[*Page]uint) // diff Unstage needs the reference order
	for instance := range stage.Pages {
		stage.Pages_reference[instance] = instance.GongCopy().(*Page)
		stage.Pages_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (chapter *Chapter) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ChapterMap_Staged_Order[chapter]; ok {
		return order
	}
	return stage.Chapters_referenceOrder[chapter]
}

func (chapter *Chapter) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Chapters_referenceOrder[chapter]
}

func (content *Content) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ContentMap_Staged_Order[content]; ok {
		return order
	}
	return stage.Contents_referenceOrder[content]
}

func (content *Content) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Contents_referenceOrder[content]
}

func (page *Page) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PageMap_Staged_Order[page]; ok {
		return order
	}
	return stage.Pages_referenceOrder[page]
}

func (page *Page) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Pages_referenceOrder[page]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (chapter *Chapter) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", chapter.GongGetGongstructName(), chapter.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (chapter *Chapter) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", chapter.GongGetGongstructName(), chapter.GongGetReferenceOrder(stage))
}

func (content *Content) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", content.GongGetGongstructName(), content.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (content *Content) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", content.GongGetGongstructName(), content.GongGetReferenceOrder(stage))
}

func (page *Page) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", page.GongGetGongstructName(), page.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (page *Page) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", page.GongGetGongstructName(), page.GongGetReferenceOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (chapter *Chapter) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", chapter.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Chapter")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", chapter.Name)
	return
}
func (content *Content) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", content.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Content")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", content.Name)
	return
}
func (page *Page) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", page.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Page")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", page.Name)
	return
}

// insertion point for unstaging
func (chapter *Chapter) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", chapter.GongGetReferenceIdentifier(stage))
	return
}
func (content *Content) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", content.GongGetReferenceIdentifier(stage))
	return
}
func (page *Page) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", page.GongGetReferenceIdentifier(stage))
	return
}
