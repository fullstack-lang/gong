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

	// Compute reverse map for named struct Page
	// insertion point per field
	stage.Page_Sections_reverseMap = make(map[*Section]*Page)
	for page := range stage.Pages {
		_ = page
		for _, _section := range page.Sections {
			stage.Page_Sections_reverseMap[_section] = page
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.Chapters)

	res = __gong__appendInstances(res, stage.Contents)

	res = __gong__appendInstances(res, stage.DownloadableFiles)

	res = __gong__appendInstances(res, stage.JpgImages)

	res = __gong__appendInstances(res, stage.Pages)

	res = __gong__appendInstances(res, stage.PngImages)

	res = __gong__appendInstances(res, stage.Sections)

	res = __gong__appendInstances(res, stage.SvgImages)

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
func (chapter *Chapter) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, chapter)
}

func (content *Content) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, content)
}

func (downloadablefile *DownloadableFile) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, downloadablefile)
}

func (jpgimage *JpgImage) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, jpgimage)
}

func (page *Page) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, page)
}

func (pngimage *PngImage) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pngimage)
}

func (section *Section) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, section)
}

func (svgimage *SvgImage) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, svgimage)
}


type GongstructDiffable[T any] interface {
	GongstructPtr
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
	__gong__computeReferencePass1(stage, stage.Chapters, &stage.Chapters_reference, &stage.Chapters_referenceOrder, &stage.Chapters_instance)

	__gong__computeReferencePass1(stage, stage.Contents, &stage.Contents_reference, &stage.Contents_referenceOrder, &stage.Contents_instance)

	__gong__computeReferencePass1(stage, stage.DownloadableFiles, &stage.DownloadableFiles_reference, &stage.DownloadableFiles_referenceOrder, &stage.DownloadableFiles_instance)

	__gong__computeReferencePass1(stage, stage.JpgImages, &stage.JpgImages_reference, &stage.JpgImages_referenceOrder, &stage.JpgImages_instance)

	__gong__computeReferencePass1(stage, stage.Pages, &stage.Pages_reference, &stage.Pages_referenceOrder, &stage.Pages_instance)

	__gong__computeReferencePass1(stage, stage.PngImages, &stage.PngImages_reference, &stage.PngImages_referenceOrder, &stage.PngImages_instance)

	__gong__computeReferencePass1(stage, stage.Sections, &stage.Sections_reference, &stage.Sections_referenceOrder, &stage.Sections_instance)

	__gong__computeReferencePass1(stage, stage.SvgImages, &stage.SvgImages_reference, &stage.SvgImages_referenceOrder, &stage.SvgImages_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.Chapters, stage.Chapters_reference, stage)

	__gong__computeReferencePass2(stage.Contents, stage.Contents_reference, stage)

	__gong__computeReferencePass2(stage.DownloadableFiles, stage.DownloadableFiles_reference, stage)

	__gong__computeReferencePass2(stage.JpgImages, stage.JpgImages_reference, stage)

	__gong__computeReferencePass2(stage.Pages, stage.Pages_reference, stage)

	__gong__computeReferencePass2(stage.PngImages, stage.PngImages_reference, stage)

	__gong__computeReferencePass2(stage.Sections, stage.Sections_reference, stage)

	__gong__computeReferencePass2(stage.SvgImages, stage.SvgImages_reference, stage)

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
	return __gong__getOrder(stage.Chapter_stagedOrder, stage.Chapters_referenceOrder, chapter, "Chapter")
}

func (content *Content) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Content_stagedOrder, stage.Contents_referenceOrder, content, "Content")
}

func (downloadablefile *DownloadableFile) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DownloadableFile_stagedOrder, stage.DownloadableFiles_referenceOrder, downloadablefile, "DownloadableFile")
}

func (jpgimage *JpgImage) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.JpgImage_stagedOrder, stage.JpgImages_referenceOrder, jpgimage, "JpgImage")
}

func (page *Page) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Page_stagedOrder, stage.Pages_referenceOrder, page, "Page")
}

func (pngimage *PngImage) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PngImage_stagedOrder, stage.PngImages_referenceOrder, pngimage, "PngImage")
}

func (section *Section) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Section_stagedOrder, stage.Sections_referenceOrder, section, "Section")
}

func (svgimage *SvgImage) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SvgImage_stagedOrder, stage.SvgImages_referenceOrder, svgimage, "SvgImage")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (chapter *Chapter) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(chapter, chapter.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (chapter *Chapter) GongGetReferenceIdentifier(stage *Stage) string {
	return chapter.GongGetIdentifier(stage)
}

func (content *Content) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(content, content.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (content *Content) GongGetReferenceIdentifier(stage *Stage) string {
	return content.GongGetIdentifier(stage)
}

func (downloadablefile *DownloadableFile) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(downloadablefile, downloadablefile.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (downloadablefile *DownloadableFile) GongGetReferenceIdentifier(stage *Stage) string {
	return downloadablefile.GongGetIdentifier(stage)
}

func (jpgimage *JpgImage) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(jpgimage, jpgimage.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (jpgimage *JpgImage) GongGetReferenceIdentifier(stage *Stage) string {
	return jpgimage.GongGetIdentifier(stage)
}

func (page *Page) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(page, page.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (page *Page) GongGetReferenceIdentifier(stage *Stage) string {
	return page.GongGetIdentifier(stage)
}

func (pngimage *PngImage) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pngimage, pngimage.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pngimage *PngImage) GongGetReferenceIdentifier(stage *Stage) string {
	return pngimage.GongGetIdentifier(stage)
}

func (section *Section) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(section, section.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (section *Section) GongGetReferenceIdentifier(stage *Stage) string {
	return section.GongGetIdentifier(stage)
}

func (svgimage *SvgImage) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(svgimage, svgimage.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (svgimage *SvgImage) GongGetReferenceIdentifier(stage *Stage) string {
	return svgimage.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (chapter *Chapter) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(chapter.GongGetIdentifier(stage), "Chapter", chapter.Name)
}

func (content *Content) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(content.GongGetIdentifier(stage), "Content", content.Name)
}

func (downloadablefile *DownloadableFile) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(downloadablefile.GongGetIdentifier(stage), "DownloadableFile", downloadablefile.Name)
}

func (jpgimage *JpgImage) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(jpgimage.GongGetIdentifier(stage), "JpgImage", jpgimage.Name)
}

func (page *Page) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(page.GongGetIdentifier(stage), "Page", page.Name)
}

func (pngimage *PngImage) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pngimage.GongGetIdentifier(stage), "PngImage", pngimage.Name)
}

func (section *Section) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(section.GongGetIdentifier(stage), "Section", section.Name)
}

func (svgimage *SvgImage) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(svgimage.GongGetIdentifier(stage), "SvgImage", svgimage.Name)
}

// insertion point for unstaging
func (chapter *Chapter) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(chapter.GongGetReferenceIdentifier(stage))
}

func (content *Content) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(content.GongGetReferenceIdentifier(stage))
}

func (downloadablefile *DownloadableFile) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(downloadablefile.GongGetReferenceIdentifier(stage))
}

func (jpgimage *JpgImage) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(jpgimage.GongGetReferenceIdentifier(stage))
}

func (page *Page) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(page.GongGetReferenceIdentifier(stage))
}

func (pngimage *PngImage) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pngimage.GongGetReferenceIdentifier(stage))
}

func (section *Section) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(section.GongGetReferenceIdentifier(stage))
}

func (svgimage *SvgImage) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(svgimage.GongGetReferenceIdentifier(stage))
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

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
