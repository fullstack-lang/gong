// generated code - do not edit
package models

import (
	"fmt"
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
	// Compute reverse map for named struct Content
	// insertion point per field

	// Compute reverse map for named struct JpgImage
	// insertion point per field

	// Compute reverse map for named struct PngImage
	// insertion point per field

	// Compute reverse map for named struct SvgImage
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.Contents {
		res = append(res, instance)
	}

	for instance := range stage.JpgImages {
		res = append(res, instance)
	}

	for instance := range stage.PngImages {
		res = append(res, instance)
	}

	for instance := range stage.SvgImages {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (content *Content) GongCopy() GongstructIF {
	newInstance := *content
	return &newInstance
}

func (jpgimage *JpgImage) GongCopy() GongstructIF {
	newInstance := *jpgimage
	return &newInstance
}

func (pngimage *PngImage) GongCopy() GongstructIF {
	newInstance := *pngimage
	return &newInstance
}

func (svgimage *SvgImage) GongCopy() GongstructIF {
	newInstance := *svgimage
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
	var jpgimages_newInstances []*JpgImage
	var jpgimages_deletedInstances []*JpgImage

	// parse all staged instances and check if they have a reference
	for jpgimage := range stage.JpgImages {
		if ref, ok := stage.JpgImages_reference[jpgimage]; !ok {
			jpgimages_newInstances = append(jpgimages_newInstances, jpgimage)
			newInstancesSlice = append(newInstancesSlice, jpgimage.GongMarshallIdentifier(stage))
			if stage.JpgImages_referenceOrder == nil {
				stage.JpgImages_referenceOrder = make(map[*JpgImage]uint)
			}
			stage.JpgImages_referenceOrder[jpgimage] = stage.JpgImageMap_Staged_Order[jpgimage]
			newInstancesReverseSlice = append(newInstancesReverseSlice, jpgimage.GongMarshallUnstaging(stage))
			delete(stage.JpgImages_referenceOrder, jpgimage)
			fieldInitializers, pointersInitializations := jpgimage.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.JpgImageMap_Staged_Order[ref] = stage.JpgImageMap_Staged_Order[jpgimage]
			diffs := jpgimage.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, jpgimage)
			delete(stage.JpgImageMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", jpgimage.GetName())
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
	for ref := range stage.JpgImages_reference {
		if _, ok := stage.JpgImages[ref]; !ok {
			jpgimages_deletedInstances = append(jpgimages_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(jpgimages_newInstances)
	lenDeletedInstances += len(jpgimages_deletedInstances)
	var pngimages_newInstances []*PngImage
	var pngimages_deletedInstances []*PngImage

	// parse all staged instances and check if they have a reference
	for pngimage := range stage.PngImages {
		if ref, ok := stage.PngImages_reference[pngimage]; !ok {
			pngimages_newInstances = append(pngimages_newInstances, pngimage)
			newInstancesSlice = append(newInstancesSlice, pngimage.GongMarshallIdentifier(stage))
			if stage.PngImages_referenceOrder == nil {
				stage.PngImages_referenceOrder = make(map[*PngImage]uint)
			}
			stage.PngImages_referenceOrder[pngimage] = stage.PngImageMap_Staged_Order[pngimage]
			newInstancesReverseSlice = append(newInstancesReverseSlice, pngimage.GongMarshallUnstaging(stage))
			delete(stage.PngImages_referenceOrder, pngimage)
			fieldInitializers, pointersInitializations := pngimage.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PngImageMap_Staged_Order[ref] = stage.PngImageMap_Staged_Order[pngimage]
			diffs := pngimage.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, pngimage)
			delete(stage.PngImageMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", pngimage.GetName())
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
	for ref := range stage.PngImages_reference {
		if _, ok := stage.PngImages[ref]; !ok {
			pngimages_deletedInstances = append(pngimages_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(pngimages_newInstances)
	lenDeletedInstances += len(pngimages_deletedInstances)
	var svgimages_newInstances []*SvgImage
	var svgimages_deletedInstances []*SvgImage

	// parse all staged instances and check if they have a reference
	for svgimage := range stage.SvgImages {
		if ref, ok := stage.SvgImages_reference[svgimage]; !ok {
			svgimages_newInstances = append(svgimages_newInstances, svgimage)
			newInstancesSlice = append(newInstancesSlice, svgimage.GongMarshallIdentifier(stage))
			if stage.SvgImages_referenceOrder == nil {
				stage.SvgImages_referenceOrder = make(map[*SvgImage]uint)
			}
			stage.SvgImages_referenceOrder[svgimage] = stage.SvgImageMap_Staged_Order[svgimage]
			newInstancesReverseSlice = append(newInstancesReverseSlice, svgimage.GongMarshallUnstaging(stage))
			delete(stage.SvgImages_referenceOrder, svgimage)
			fieldInitializers, pointersInitializations := svgimage.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SvgImageMap_Staged_Order[ref] = stage.SvgImageMap_Staged_Order[svgimage]
			diffs := svgimage.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, svgimage)
			delete(stage.SvgImageMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", svgimage.GetName())
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
	for ref := range stage.SvgImages_reference {
		if _, ok := stage.SvgImages[ref]; !ok {
			svgimages_deletedInstances = append(svgimages_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(svgimages_newInstances)
	lenDeletedInstances += len(svgimages_deletedInstances)

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
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	stage.Contents_reference = make(map[*Content]*Content)
	stage.Contents_referenceOrder = make(map[*Content]uint) // diff Unstage needs the reference order
	for instance := range stage.Contents {
		stage.Contents_reference[instance] = instance.GongCopy().(*Content)
		stage.Contents_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.JpgImages_reference = make(map[*JpgImage]*JpgImage)
	stage.JpgImages_referenceOrder = make(map[*JpgImage]uint) // diff Unstage needs the reference order
	for instance := range stage.JpgImages {
		stage.JpgImages_reference[instance] = instance.GongCopy().(*JpgImage)
		stage.JpgImages_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.PngImages_reference = make(map[*PngImage]*PngImage)
	stage.PngImages_referenceOrder = make(map[*PngImage]uint) // diff Unstage needs the reference order
	for instance := range stage.PngImages {
		stage.PngImages_reference[instance] = instance.GongCopy().(*PngImage)
		stage.PngImages_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.SvgImages_reference = make(map[*SvgImage]*SvgImage)
	stage.SvgImages_referenceOrder = make(map[*SvgImage]uint) // diff Unstage needs the reference order
	for instance := range stage.SvgImages {
		stage.SvgImages_reference[instance] = instance.GongCopy().(*SvgImage)
		stage.SvgImages_referenceOrder[instance] = instance.GongGetOrder(stage)
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
func (content *Content) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ContentMap_Staged_Order[content]; ok {
		return order
	}
	return stage.Contents_referenceOrder[content]
}

func (content *Content) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Contents_referenceOrder[content]
}

func (jpgimage *JpgImage) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.JpgImageMap_Staged_Order[jpgimage]; ok {
		return order
	}
	return stage.JpgImages_referenceOrder[jpgimage]
}

func (jpgimage *JpgImage) GongGetReferenceOrder(stage *Stage) uint {
	return stage.JpgImages_referenceOrder[jpgimage]
}

func (pngimage *PngImage) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PngImageMap_Staged_Order[pngimage]; ok {
		return order
	}
	return stage.PngImages_referenceOrder[pngimage]
}

func (pngimage *PngImage) GongGetReferenceOrder(stage *Stage) uint {
	return stage.PngImages_referenceOrder[pngimage]
}

func (svgimage *SvgImage) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SvgImageMap_Staged_Order[svgimage]; ok {
		return order
	}
	return stage.SvgImages_referenceOrder[svgimage]
}

func (svgimage *SvgImage) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SvgImages_referenceOrder[svgimage]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (content *Content) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", content.GongGetGongstructName(), content.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (content *Content) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", content.GongGetGongstructName(), content.GongGetReferenceOrder(stage))
}

func (jpgimage *JpgImage) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", jpgimage.GongGetGongstructName(), jpgimage.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (jpgimage *JpgImage) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", jpgimage.GongGetGongstructName(), jpgimage.GongGetReferenceOrder(stage))
}

func (pngimage *PngImage) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pngimage.GongGetGongstructName(), pngimage.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pngimage *PngImage) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pngimage.GongGetGongstructName(), pngimage.GongGetReferenceOrder(stage))
}

func (svgimage *SvgImage) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", svgimage.GongGetGongstructName(), svgimage.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (svgimage *SvgImage) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", svgimage.GongGetGongstructName(), svgimage.GongGetReferenceOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (content *Content) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", content.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Content")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", content.Name)
	return
}

func (jpgimage *JpgImage) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", jpgimage.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "JpgImage")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", jpgimage.Name)
	return
}

func (pngimage *PngImage) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pngimage.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PngImage")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", pngimage.Name)
	return
}

func (svgimage *SvgImage) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", svgimage.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SvgImage")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", svgimage.Name)
	return
}

// insertion point for unstaging
func (content *Content) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", content.GongGetReferenceIdentifier(stage))
	return
}

func (jpgimage *JpgImage) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", jpgimage.GongGetReferenceIdentifier(stage))
	return
}

func (pngimage *PngImage) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pngimage.GongGetReferenceIdentifier(stage))
	return
}

func (svgimage *SvgImage) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", svgimage.GongGetReferenceIdentifier(stage))
	return
}

// end of template
