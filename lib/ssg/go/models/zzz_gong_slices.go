// generated code - do not edit
package models

import (
	"crypto/sha256"
	"encoding/binary"
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
	// Compute reverse map for named struct Chapter
	// insertion point per field
	stage.Chapter_Sections_reverseMap = make(map[*Section]*Chapter)
	for chapter := range stage.Chapters {
		_ = chapter
		for _, _section := range chapter.Sections {
			stage.Chapter_Sections_reverseMap[_section] = chapter
		}
	}
	stage.Chapter_Pages_reverseMap = make(map[*Page]*Chapter)
	for chapter := range stage.Chapters {
		_ = chapter
		for _, _page := range chapter.Pages {
			stage.Chapter_Pages_reverseMap[_page] = chapter
		}
	}
	stage.Chapter_SubChapters_reverseMap = make(map[*Chapter]*Chapter)
	for chapter := range stage.Chapters {
		_ = chapter
		for _, _chapter := range chapter.SubChapters {
			stage.Chapter_SubChapters_reverseMap[_chapter] = chapter
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

	// Compute reverse map for named struct DownloadableFile
	// insertion point per field

	// Compute reverse map for named struct JpgImage
	// insertion point per field

	// Compute reverse map for named struct Page
	// insertion point per field
	stage.Page_Sections_reverseMap = make(map[*Section]*Page)
	for page := range stage.Pages {
		_ = page
		for _, _section := range page.Sections {
			stage.Page_Sections_reverseMap[_section] = page
		}
	}

	// Compute reverse map for named struct PngImage
	// insertion point per field

	// Compute reverse map for named struct Section
	// insertion point per field

	// Compute reverse map for named struct SvgImage
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.Chapters {
		res = append(res, instance)
	}

	for instance := range stage.Contents {
		res = append(res, instance)
	}

	for instance := range stage.DownloadableFiles {
		res = append(res, instance)
	}

	for instance := range stage.JpgImages {
		res = append(res, instance)
	}

	for instance := range stage.Pages {
		res = append(res, instance)
	}

	for instance := range stage.PngImages {
		res = append(res, instance)
	}

	for instance := range stage.Sections {
		res = append(res, instance)
	}

	for instance := range stage.SvgImages {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (chapter *Chapter) GongCopy() GongstructIF {
	newInstance := new(Chapter)
	chapter.GongCopyBasicFields(newInstance)
	return newInstance
}

func (content *Content) GongCopy() GongstructIF {
	newInstance := new(Content)
	content.GongCopyBasicFields(newInstance)
	return newInstance
}

func (downloadablefile *DownloadableFile) GongCopy() GongstructIF {
	newInstance := new(DownloadableFile)
	downloadablefile.GongCopyBasicFields(newInstance)
	return newInstance
}

func (jpgimage *JpgImage) GongCopy() GongstructIF {
	newInstance := new(JpgImage)
	jpgimage.GongCopyBasicFields(newInstance)
	return newInstance
}

func (page *Page) GongCopy() GongstructIF {
	newInstance := new(Page)
	page.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pngimage *PngImage) GongCopy() GongstructIF {
	newInstance := new(PngImage)
	pngimage.GongCopyBasicFields(newInstance)
	return newInstance
}

func (section *Section) GongCopy() GongstructIF {
	newInstance := new(Section)
	section.GongCopyBasicFields(newInstance)
	return newInstance
}

func (svgimage *SvgImage) GongCopy() GongstructIF {
	newInstance := new(SvgImage)
	svgimage.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (chapter *Chapter) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(chapter).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(chapter), uint64(stage.GetOrder(chapter)))
	return
}

func (content *Content) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(content).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(content), uint64(stage.GetOrder(content)))
	return
}

func (downloadablefile *DownloadableFile) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(downloadablefile).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(downloadablefile), uint64(stage.GetOrder(downloadablefile)))
	return
}

func (jpgimage *JpgImage) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(jpgimage).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(jpgimage), uint64(stage.GetOrder(jpgimage)))
	return
}

func (page *Page) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(page).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(page), uint64(stage.GetOrder(page)))
	return
}

func (pngimage *PngImage) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(pngimage).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(pngimage), uint64(stage.GetOrder(pngimage)))
	return
}

func (section *Section) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(section).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(section), uint64(stage.GetOrder(section)))
	return
}

func (svgimage *SvgImage) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(svgimage).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(svgimage), uint64(stage.GetOrder(svgimage)))
	return
}


type GongstructDiffable[T any] interface {
	PointerToGongstruct
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
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
	computeCommitsForType(
		stage,
		stage.Chapters,
		stage.Chapter_stagedOrder,
		stage.Chapters_reference,
		&stage.Chapters_referenceOrder,
		stage.Chapters_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Contents,
		stage.Content_stagedOrder,
		stage.Contents_reference,
		&stage.Contents_referenceOrder,
		stage.Contents_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DownloadableFiles,
		stage.DownloadableFile_stagedOrder,
		stage.DownloadableFiles_reference,
		&stage.DownloadableFiles_referenceOrder,
		stage.DownloadableFiles_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.JpgImages,
		stage.JpgImage_stagedOrder,
		stage.JpgImages_reference,
		&stage.JpgImages_referenceOrder,
		stage.JpgImages_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Pages,
		stage.Page_stagedOrder,
		stage.Pages_reference,
		&stage.Pages_referenceOrder,
		stage.Pages_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.PngImages,
		stage.PngImage_stagedOrder,
		stage.PngImages_reference,
		&stage.PngImages_referenceOrder,
		stage.PngImages_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Sections,
		stage.Section_stagedOrder,
		stage.Sections_reference,
		&stage.Sections_referenceOrder,
		stage.Sections_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SvgImages,
		stage.SvgImage_stagedOrder,
		stage.SvgImages_reference,
		&stage.SvgImages_referenceOrder,
		stage.SvgImages_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

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
	stage.Chapters_reference = make(map[*Chapter]*Chapter)
	stage.Chapters_referenceOrder = make(map[*Chapter]uint) // diff Unstage needs the reference order
	stage.Chapters_instance = make(map[*Chapter]*Chapter)
	for instance := range stage.Chapters {
		_copy := instance.GongCopy().(*Chapter)
		stage.Chapters_reference[instance] = _copy
		stage.Chapters_instance[_copy] = instance
		stage.Chapters_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Contents_reference = make(map[*Content]*Content)
	stage.Contents_referenceOrder = make(map[*Content]uint) // diff Unstage needs the reference order
	stage.Contents_instance = make(map[*Content]*Content)
	for instance := range stage.Contents {
		_copy := instance.GongCopy().(*Content)
		stage.Contents_reference[instance] = _copy
		stage.Contents_instance[_copy] = instance
		stage.Contents_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.DownloadableFiles_reference = make(map[*DownloadableFile]*DownloadableFile)
	stage.DownloadableFiles_referenceOrder = make(map[*DownloadableFile]uint) // diff Unstage needs the reference order
	stage.DownloadableFiles_instance = make(map[*DownloadableFile]*DownloadableFile)
	for instance := range stage.DownloadableFiles {
		_copy := instance.GongCopy().(*DownloadableFile)
		stage.DownloadableFiles_reference[instance] = _copy
		stage.DownloadableFiles_instance[_copy] = instance
		stage.DownloadableFiles_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.JpgImages_reference = make(map[*JpgImage]*JpgImage)
	stage.JpgImages_referenceOrder = make(map[*JpgImage]uint) // diff Unstage needs the reference order
	stage.JpgImages_instance = make(map[*JpgImage]*JpgImage)
	for instance := range stage.JpgImages {
		_copy := instance.GongCopy().(*JpgImage)
		stage.JpgImages_reference[instance] = _copy
		stage.JpgImages_instance[_copy] = instance
		stage.JpgImages_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Pages_reference = make(map[*Page]*Page)
	stage.Pages_referenceOrder = make(map[*Page]uint) // diff Unstage needs the reference order
	stage.Pages_instance = make(map[*Page]*Page)
	for instance := range stage.Pages {
		_copy := instance.GongCopy().(*Page)
		stage.Pages_reference[instance] = _copy
		stage.Pages_instance[_copy] = instance
		stage.Pages_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PngImages_reference = make(map[*PngImage]*PngImage)
	stage.PngImages_referenceOrder = make(map[*PngImage]uint) // diff Unstage needs the reference order
	stage.PngImages_instance = make(map[*PngImage]*PngImage)
	for instance := range stage.PngImages {
		_copy := instance.GongCopy().(*PngImage)
		stage.PngImages_reference[instance] = _copy
		stage.PngImages_instance[_copy] = instance
		stage.PngImages_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Sections_reference = make(map[*Section]*Section)
	stage.Sections_referenceOrder = make(map[*Section]uint) // diff Unstage needs the reference order
	stage.Sections_instance = make(map[*Section]*Section)
	for instance := range stage.Sections {
		_copy := instance.GongCopy().(*Section)
		stage.Sections_reference[instance] = _copy
		stage.Sections_instance[_copy] = instance
		stage.Sections_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SvgImages_reference = make(map[*SvgImage]*SvgImage)
	stage.SvgImages_referenceOrder = make(map[*SvgImage]uint) // diff Unstage needs the reference order
	stage.SvgImages_instance = make(map[*SvgImage]*SvgImage)
	for instance := range stage.SvgImages {
		_copy := instance.GongCopy().(*SvgImage)
		stage.SvgImages_reference[instance] = _copy
		stage.SvgImages_instance[_copy] = instance
		stage.SvgImages_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.Chapters {
		reference := stage.Chapters_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Contents {
		reference := stage.Contents_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.DownloadableFiles {
		reference := stage.DownloadableFiles_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.JpgImages {
		reference := stage.JpgImages_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Pages {
		reference := stage.Pages_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PngImages {
		reference := stage.PngImages_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Sections {
		reference := stage.Sections_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SvgImages {
		reference := stage.SvgImages_reference[instance]
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
func (chapter *Chapter) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Chapter_stagedOrder[chapter]; ok {
		return order
	}
	if order, ok := stage.Chapters_referenceOrder[chapter]; ok {
		return order
	} else {
		log.Printf("instance %p of type Chapter was not staged and does not have a reference order", chapter)
		return 0
	}
}

func (content *Content) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Content_stagedOrder[content]; ok {
		return order
	}
	if order, ok := stage.Contents_referenceOrder[content]; ok {
		return order
	} else {
		log.Printf("instance %p of type Content was not staged and does not have a reference order", content)
		return 0
	}
}

func (downloadablefile *DownloadableFile) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DownloadableFile_stagedOrder[downloadablefile]; ok {
		return order
	}
	if order, ok := stage.DownloadableFiles_referenceOrder[downloadablefile]; ok {
		return order
	} else {
		log.Printf("instance %p of type DownloadableFile was not staged and does not have a reference order", downloadablefile)
		return 0
	}
}

func (jpgimage *JpgImage) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.JpgImage_stagedOrder[jpgimage]; ok {
		return order
	}
	if order, ok := stage.JpgImages_referenceOrder[jpgimage]; ok {
		return order
	} else {
		log.Printf("instance %p of type JpgImage was not staged and does not have a reference order", jpgimage)
		return 0
	}
}

func (page *Page) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Page_stagedOrder[page]; ok {
		return order
	}
	if order, ok := stage.Pages_referenceOrder[page]; ok {
		return order
	} else {
		log.Printf("instance %p of type Page was not staged and does not have a reference order", page)
		return 0
	}
}

func (pngimage *PngImage) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PngImage_stagedOrder[pngimage]; ok {
		return order
	}
	if order, ok := stage.PngImages_referenceOrder[pngimage]; ok {
		return order
	} else {
		log.Printf("instance %p of type PngImage was not staged and does not have a reference order", pngimage)
		return 0
	}
}

func (section *Section) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Section_stagedOrder[section]; ok {
		return order
	}
	if order, ok := stage.Sections_referenceOrder[section]; ok {
		return order
	} else {
		log.Printf("instance %p of type Section was not staged and does not have a reference order", section)
		return 0
	}
}

func (svgimage *SvgImage) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SvgImage_stagedOrder[svgimage]; ok {
		return order
	}
	if order, ok := stage.SvgImages_referenceOrder[svgimage]; ok {
		return order
	} else {
		log.Printf("instance %p of type SvgImage was not staged and does not have a reference order", svgimage)
		return 0
	}
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
	return fmt.Sprintf("__%s__%08d_", chapter.GongGetGongstructName(), chapter.GongGetOrder(stage))
}

func (content *Content) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", content.GongGetGongstructName(), content.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (content *Content) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", content.GongGetGongstructName(), content.GongGetOrder(stage))
}

func (downloadablefile *DownloadableFile) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", downloadablefile.GongGetGongstructName(), downloadablefile.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (downloadablefile *DownloadableFile) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", downloadablefile.GongGetGongstructName(), downloadablefile.GongGetOrder(stage))
}

func (jpgimage *JpgImage) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", jpgimage.GongGetGongstructName(), jpgimage.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (jpgimage *JpgImage) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", jpgimage.GongGetGongstructName(), jpgimage.GongGetOrder(stage))
}

func (page *Page) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", page.GongGetGongstructName(), page.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (page *Page) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", page.GongGetGongstructName(), page.GongGetOrder(stage))
}

func (pngimage *PngImage) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pngimage.GongGetGongstructName(), pngimage.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pngimage *PngImage) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pngimage.GongGetGongstructName(), pngimage.GongGetOrder(stage))
}

func (section *Section) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", section.GongGetGongstructName(), section.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (section *Section) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", section.GongGetGongstructName(), section.GongGetOrder(stage))
}

func (svgimage *SvgImage) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", svgimage.GongGetGongstructName(), svgimage.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (svgimage *SvgImage) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", svgimage.GongGetGongstructName(), svgimage.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (chapter *Chapter) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", chapter.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Chapter")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(chapter.Name))
	return
}

func (content *Content) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", content.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Content")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(content.Name))
	return
}

func (downloadablefile *DownloadableFile) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", downloadablefile.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DownloadableFile")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(downloadablefile.Name))
	return
}

func (jpgimage *JpgImage) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", jpgimage.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "JpgImage")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(jpgimage.Name))
	return
}

func (page *Page) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", page.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Page")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(page.Name))
	return
}

func (pngimage *PngImage) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pngimage.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PngImage")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(pngimage.Name))
	return
}

func (section *Section) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", section.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Section")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(section.Name))
	return
}

func (svgimage *SvgImage) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", svgimage.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SvgImage")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(svgimage.Name))
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

func (downloadablefile *DownloadableFile) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", downloadablefile.GongGetReferenceIdentifier(stage))
	return
}

func (jpgimage *JpgImage) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", jpgimage.GongGetReferenceIdentifier(stage))
	return
}

func (page *Page) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", page.GongGetReferenceIdentifier(stage))
	return
}

func (pngimage *PngImage) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pngimage.GongGetReferenceIdentifier(stage))
	return
}

func (section *Section) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", section.GongGetReferenceIdentifier(stage))
	return
}

func (svgimage *SvgImage) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", svgimage.GongGetReferenceIdentifier(stage))
	return
}

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

// end of template
