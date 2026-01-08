// generated code - do not edit
package models

import (
	"fmt"
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

func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesStmt string
	_ = newInstancesStmt
	var fieldsEditStmt string
	_ = fieldsEditStmt
	var deletedInstancesStmt string
	_ = deletedInstancesStmt

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
			newInstancesStmt += chapter.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := chapter.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := chapter.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", chapter.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for chapter := range stage.Chapters_reference {
		if _, ok := stage.Chapters[chapter]; !ok {
			chapters_deletedInstances = append(chapters_deletedInstances, chapter)
			deletedInstancesStmt += chapter.GongMarshallUnstaging(stage)
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
			newInstancesStmt += content.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := content.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := content.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", content.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for content := range stage.Contents_reference {
		if _, ok := stage.Contents[content]; !ok {
			contents_deletedInstances = append(contents_deletedInstances, content)
			deletedInstancesStmt += content.GongMarshallUnstaging(stage)
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
			newInstancesStmt += page.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := page.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := page.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", page.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for page := range stage.Pages_reference {
		if _, ok := stage.Pages[page]; !ok {
			pages_deletedInstances = append(pages_deletedInstances, page)
			deletedInstancesStmt += page.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len(pages_newInstances)
	lenDeletedInstances += len(pages_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		notif := newInstancesStmt+fieldsEditStmt+deletedInstancesStmt
		notif += fmt.Sprintf("\n\t// %s", time.Now().Format(time.RFC3339Nano))
		notif += fmt.Sprintf("\n\tstage.Commit()")
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().AddNotification(
				time.Now(),
				notif,
			)
			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Chapters_reference = make(map[*Chapter]*Chapter)
	for instance := range stage.Chapters {
		stage.Chapters_reference[instance] = instance.GongCopy().(*Chapter)
	}

	stage.Contents_reference = make(map[*Content]*Content)
	for instance := range stage.Contents {
		stage.Contents_reference[instance] = instance.GongCopy().(*Content)
	}

	stage.Pages_reference = make(map[*Page]*Page)
	for instance := range stage.Pages {
		stage.Pages_reference[instance] = instance.GongCopy().(*Page)
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
	return stage.ChapterMap_Staged_Order[chapter]
}

func (content *Content) GongGetOrder(stage *Stage) uint {
	return stage.ContentMap_Staged_Order[content]
}

func (page *Page) GongGetOrder(stage *Stage) uint {
	return stage.PageMap_Staged_Order[page]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (chapter *Chapter) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", chapter.GongGetGongstructName(), chapter.GongGetOrder(stage))
}

func (content *Content) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", content.GongGetGongstructName(), content.GongGetOrder(stage))
}

func (page *Page) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", page.GongGetGongstructName(), page.GongGetOrder(stage))
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
	decl = strings.ReplaceAll(decl, "{{Identifier}}", chapter.GongGetIdentifier(stage))
	return
}
func (content *Content) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", content.GongGetIdentifier(stage))
	return
}
func (page *Page) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", page.GongGetIdentifier(stage))
	return
}
