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
	// Compute reverse map for named struct Content
	// insertion point per field

	// Compute reverse map for named struct JpgImage
	// insertion point per field

	// Compute reverse map for named struct PngImage
	// insertion point per field

	// Compute reverse map for named struct SvgImage
	// insertion point per field

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

	var newInstancesReverseStmt string
	_ = newInstancesReverseStmt
	var fieldsEditReverseStmt string
	_ = fieldsEditReverseStmt
	var deletedInstancesReverseStmt string
	_ = deletedInstancesReverseStmt

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
			newInstancesStmt += content.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += content.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := content.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := content.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, content)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", content.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Contents_reference {
		if _, ok := stage.Contents[ref]; !ok {
			contents_deletedInstances = append(contents_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += jpgimage.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += jpgimage.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := jpgimage.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := jpgimage.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, jpgimage)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", jpgimage.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.JpgImages_reference {
		if _, ok := stage.JpgImages[ref]; !ok {
			jpgimages_deletedInstances = append(jpgimages_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += pngimage.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += pngimage.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := pngimage.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := pngimage.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, pngimage)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", pngimage.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.PngImages_reference {
		if _, ok := stage.PngImages[ref]; !ok {
			pngimages_deletedInstances = append(pngimages_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += svgimage.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += svgimage.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := svgimage.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := svgimage.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, svgimage)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", svgimage.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.SvgImages_reference {
		if _, ok := stage.SvgImages[ref]; !ok {
			svgimages_deletedInstances = append(svgimages_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
		}
	}

	lenNewInstances += len(svgimages_newInstances)
	lenDeletedInstances += len(svgimages_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += fmt.Sprintf("\n\t// %s", time.Now().Format(time.RFC3339Nano))
		forwardCommit += "\n\tstage.Commit()\n"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += fmt.Sprintf("\n\t// %s", time.Now().Format(time.RFC3339Nano))
		backwardCommit += "\n\tstage.Commit()\n"
		// append to the start of the backward commits slice
		stage.backwardCommits = append([]string{backwardCommit}, stage.backwardCommits...)

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

}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (content *Content) GongGetOrder(stage *Stage) uint {
	return stage.ContentMap_Staged_Order[content]
}

func (content *Content) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Contents_referenceOrder[content]
}

func (jpgimage *JpgImage) GongGetOrder(stage *Stage) uint {
	return stage.JpgImageMap_Staged_Order[jpgimage]
}

func (jpgimage *JpgImage) GongGetReferenceOrder(stage *Stage) uint {
	return stage.JpgImages_referenceOrder[jpgimage]
}

func (pngimage *PngImage) GongGetOrder(stage *Stage) uint {
	return stage.PngImageMap_Staged_Order[pngimage]
}

func (pngimage *PngImage) GongGetReferenceOrder(stage *Stage) uint {
	return stage.PngImages_referenceOrder[pngimage]
}

func (svgimage *SvgImage) GongGetOrder(stage *Stage) uint {
	return stage.SvgImageMap_Staged_Order[svgimage]
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
