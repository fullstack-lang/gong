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
	chapter.CopyBasicFields(newInstance)
	return newInstance
}

func (content *Content) GongCopy() GongstructIF {
	newInstance := new(Content)
	content.CopyBasicFields(newInstance)
	return newInstance
}

func (downloadablefile *DownloadableFile) GongCopy() GongstructIF {
	newInstance := new(DownloadableFile)
	downloadablefile.CopyBasicFields(newInstance)
	return newInstance
}

func (jpgimage *JpgImage) GongCopy() GongstructIF {
	newInstance := new(JpgImage)
	jpgimage.CopyBasicFields(newInstance)
	return newInstance
}

func (page *Page) GongCopy() GongstructIF {
	newInstance := new(Page)
	page.CopyBasicFields(newInstance)
	return newInstance
}

func (pngimage *PngImage) GongCopy() GongstructIF {
	newInstance := new(PngImage)
	pngimage.CopyBasicFields(newInstance)
	return newInstance
}

func (section *Section) GongCopy() GongstructIF {
	newInstance := new(Section)
	section.CopyBasicFields(newInstance)
	return newInstance
}

func (svgimage *SvgImage) GongCopy() GongstructIF {
	newInstance := new(SvgImage)
	svgimage.CopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (chapter *Chapter) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(chapter).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(chapter), uint64(GetOrderPointerGongstruct(stage, chapter)))
	return
}

func (content *Content) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(content).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(content), uint64(GetOrderPointerGongstruct(stage, content)))
	return
}

func (downloadablefile *DownloadableFile) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(downloadablefile).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(downloadablefile), uint64(GetOrderPointerGongstruct(stage, downloadablefile)))
	return
}

func (jpgimage *JpgImage) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(jpgimage).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(jpgimage), uint64(GetOrderPointerGongstruct(stage, jpgimage)))
	return
}

func (page *Page) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(page).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(page), uint64(GetOrderPointerGongstruct(stage, page)))
	return
}

func (pngimage *PngImage) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(pngimage).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(pngimage), uint64(GetOrderPointerGongstruct(stage, pngimage)))
	return
}

func (section *Section) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(section).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(section), uint64(GetOrderPointerGongstruct(stage, section)))
	return
}

func (svgimage *SvgImage) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(svgimage).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(svgimage), uint64(GetOrderPointerGongstruct(stage, svgimage)))
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
			stage.Chapters_referenceOrder[chapter] = stage.Chapter_stagedOrder[chapter]
			newInstancesReverseSlice = append(newInstancesReverseSlice, chapter.GongMarshallUnstaging(stage))
			// delete(stage.Chapters_referenceOrder, chapter)
			fieldInitializers, pointersInitializations := chapter.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Chapter_stagedOrder[ref] = stage.Chapter_stagedOrder[chapter]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := chapter.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, chapter)
			// delete(stage.Chapter_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if chapter.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", chapter.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
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
	for _, ref := range stage.Chapters_reference {
		instance := stage.Chapters_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Chapters[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
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
			stage.Contents_referenceOrder[content] = stage.Content_stagedOrder[content]
			newInstancesReverseSlice = append(newInstancesReverseSlice, content.GongMarshallUnstaging(stage))
			// delete(stage.Contents_referenceOrder, content)
			fieldInitializers, pointersInitializations := content.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Content_stagedOrder[ref] = stage.Content_stagedOrder[content]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := content.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, content)
			// delete(stage.Content_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if content.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", content.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
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
	for _, ref := range stage.Contents_reference {
		instance := stage.Contents_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Contents[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			contents_deletedInstances = append(contents_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(contents_newInstances)
	lenDeletedInstances += len(contents_deletedInstances)
	var downloadablefiles_newInstances []*DownloadableFile
	var downloadablefiles_deletedInstances []*DownloadableFile

	// parse all staged instances and check if they have a reference
	for downloadablefile := range stage.DownloadableFiles {
		if ref, ok := stage.DownloadableFiles_reference[downloadablefile]; !ok {
			downloadablefiles_newInstances = append(downloadablefiles_newInstances, downloadablefile)
			newInstancesSlice = append(newInstancesSlice, downloadablefile.GongMarshallIdentifier(stage))
			if stage.DownloadableFiles_referenceOrder == nil {
				stage.DownloadableFiles_referenceOrder = make(map[*DownloadableFile]uint)
			}
			stage.DownloadableFiles_referenceOrder[downloadablefile] = stage.DownloadableFile_stagedOrder[downloadablefile]
			newInstancesReverseSlice = append(newInstancesReverseSlice, downloadablefile.GongMarshallUnstaging(stage))
			// delete(stage.DownloadableFiles_referenceOrder, downloadablefile)
			fieldInitializers, pointersInitializations := downloadablefile.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DownloadableFile_stagedOrder[ref] = stage.DownloadableFile_stagedOrder[downloadablefile]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := downloadablefile.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, downloadablefile)
			// delete(stage.DownloadableFile_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if downloadablefile.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", downloadablefile.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
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
	for _, ref := range stage.DownloadableFiles_reference {
		instance := stage.DownloadableFiles_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.DownloadableFiles[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			downloadablefiles_deletedInstances = append(downloadablefiles_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(downloadablefiles_newInstances)
	lenDeletedInstances += len(downloadablefiles_deletedInstances)
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
			stage.JpgImages_referenceOrder[jpgimage] = stage.JpgImage_stagedOrder[jpgimage]
			newInstancesReverseSlice = append(newInstancesReverseSlice, jpgimage.GongMarshallUnstaging(stage))
			// delete(stage.JpgImages_referenceOrder, jpgimage)
			fieldInitializers, pointersInitializations := jpgimage.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.JpgImage_stagedOrder[ref] = stage.JpgImage_stagedOrder[jpgimage]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := jpgimage.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, jpgimage)
			// delete(stage.JpgImage_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if jpgimage.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", jpgimage.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
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
	for _, ref := range stage.JpgImages_reference {
		instance := stage.JpgImages_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.JpgImages[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			jpgimages_deletedInstances = append(jpgimages_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(jpgimages_newInstances)
	lenDeletedInstances += len(jpgimages_deletedInstances)
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
			stage.Pages_referenceOrder[page] = stage.Page_stagedOrder[page]
			newInstancesReverseSlice = append(newInstancesReverseSlice, page.GongMarshallUnstaging(stage))
			// delete(stage.Pages_referenceOrder, page)
			fieldInitializers, pointersInitializations := page.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Page_stagedOrder[ref] = stage.Page_stagedOrder[page]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := page.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, page)
			// delete(stage.Page_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if page.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", page.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
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
	for _, ref := range stage.Pages_reference {
		instance := stage.Pages_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Pages[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			pages_deletedInstances = append(pages_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(pages_newInstances)
	lenDeletedInstances += len(pages_deletedInstances)
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
			stage.PngImages_referenceOrder[pngimage] = stage.PngImage_stagedOrder[pngimage]
			newInstancesReverseSlice = append(newInstancesReverseSlice, pngimage.GongMarshallUnstaging(stage))
			// delete(stage.PngImages_referenceOrder, pngimage)
			fieldInitializers, pointersInitializations := pngimage.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PngImage_stagedOrder[ref] = stage.PngImage_stagedOrder[pngimage]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := pngimage.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, pngimage)
			// delete(stage.PngImage_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if pngimage.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", pngimage.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
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
	for _, ref := range stage.PngImages_reference {
		instance := stage.PngImages_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.PngImages[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			pngimages_deletedInstances = append(pngimages_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(pngimages_newInstances)
	lenDeletedInstances += len(pngimages_deletedInstances)
	var sections_newInstances []*Section
	var sections_deletedInstances []*Section

	// parse all staged instances and check if they have a reference
	for section := range stage.Sections {
		if ref, ok := stage.Sections_reference[section]; !ok {
			sections_newInstances = append(sections_newInstances, section)
			newInstancesSlice = append(newInstancesSlice, section.GongMarshallIdentifier(stage))
			if stage.Sections_referenceOrder == nil {
				stage.Sections_referenceOrder = make(map[*Section]uint)
			}
			stage.Sections_referenceOrder[section] = stage.Section_stagedOrder[section]
			newInstancesReverseSlice = append(newInstancesReverseSlice, section.GongMarshallUnstaging(stage))
			// delete(stage.Sections_referenceOrder, section)
			fieldInitializers, pointersInitializations := section.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Section_stagedOrder[ref] = stage.Section_stagedOrder[section]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := section.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, section)
			// delete(stage.Section_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if section.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", section.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
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
	for _, ref := range stage.Sections_reference {
		instance := stage.Sections_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Sections[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			sections_deletedInstances = append(sections_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(sections_newInstances)
	lenDeletedInstances += len(sections_deletedInstances)
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
			stage.SvgImages_referenceOrder[svgimage] = stage.SvgImage_stagedOrder[svgimage]
			newInstancesReverseSlice = append(newInstancesReverseSlice, svgimage.GongMarshallUnstaging(stage))
			// delete(stage.SvgImages_referenceOrder, svgimage)
			fieldInitializers, pointersInitializations := svgimage.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SvgImage_stagedOrder[ref] = stage.SvgImage_stagedOrder[svgimage]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := svgimage.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, svgimage)
			// delete(stage.SvgImage_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if svgimage.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", svgimage.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
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
	for _, ref := range stage.SvgImages_reference {
		instance := stage.SvgImages_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.SvgImages[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
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
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(chapter.Name))
	return
}

func (content *Content) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", content.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Content")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(content.Name))
	return
}

func (downloadablefile *DownloadableFile) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", downloadablefile.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DownloadableFile")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(downloadablefile.Name))
	return
}

func (jpgimage *JpgImage) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", jpgimage.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "JpgImage")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(jpgimage.Name))
	return
}

func (page *Page) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", page.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Page")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(page.Name))
	return
}

func (pngimage *PngImage) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pngimage.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PngImage")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(pngimage.Name))
	return
}

func (section *Section) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", section.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Section")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(section.Name))
	return
}

func (svgimage *SvgImage) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", svgimage.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SvgImage")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(svgimage.Name))
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

func IntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += IntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
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
