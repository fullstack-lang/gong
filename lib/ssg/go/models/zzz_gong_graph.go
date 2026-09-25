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
func (chapter *Chapter) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Chapters[chapter]
	return ok
}

func (content *Content) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Contents[content]
	return ok
}

func (downloadablefile *DownloadableFile) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DownloadableFiles[downloadablefile]
	return ok
}

func (jpgimage *JpgImage) GongIsStaged(stage *Stage) bool {
	_, ok := stage.JpgImages[jpgimage]
	return ok
}

func (page *Page) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Pages[page]
	return ok
}

func (pngimage *PngImage) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PngImages[pngimage]
	return ok
}

func (section *Section) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Sections[section]
	return ok
}

func (svgimage *SvgImage) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SvgImages[svgimage]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (chapter *Chapter) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(chapter) {
		return
	}

	chapter.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _section := range chapter.Sections {
		stage.StageBranch(_section)
	}
	for _, _page := range chapter.Pages {
		stage.StageBranch(_page)
	}
	for _, _chapter := range chapter.SubChapters {
		stage.StageBranch(_chapter)
	}

}

func (content *Content) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(content) {
		return
	}

	content.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _chapter := range content.Chapters {
		stage.StageBranch(_chapter)
	}

}

func (downloadablefile *DownloadableFile) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(downloadablefile) {
		return
	}

	downloadablefile.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (jpgimage *JpgImage) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(jpgimage) {
		return
	}

	jpgimage.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (page *Page) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(page) {
		return
	}

	page.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _section := range page.Sections {
		stage.StageBranch(_section)
	}

}

func (pngimage *PngImage) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(pngimage) {
		return
	}

	pngimage.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (section *Section) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(section) {
		return
	}

	section.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if section.SvgImage != nil {
		stage.StageBranch(section.SvgImage)
	}
	if section.PngImage != nil {
		stage.StageBranch(section.PngImage)
	}
	if section.JpgImage != nil {
		stage.StageBranch(section.JpgImage)
	}
	if section.DownloadableFile != nil {
		stage.StageBranch(section.DownloadableFile)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (svgimage *SvgImage) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(svgimage) {
		return
	}

	svgimage.Stage(stage)

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
	case *Chapter:
		toT := GongCopyBranchChapter(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Content:
		toT := GongCopyBranchContent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DownloadableFile:
		toT := GongCopyBranchDownloadableFile(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *JpgImage:
		toT := GongCopyBranchJpgImage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Page:
		toT := GongCopyBranchPage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PngImage:
		toT := GongCopyBranchPngImage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Section:
		toT := GongCopyBranchSection(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SvgImage:
		toT := GongCopyBranchSvgImage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchChapter(mapOrigCopy map[any]any, chapterFrom *Chapter) (chapterTo *Chapter) {
	var alreadyCopied bool
	chapterTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, chapterFrom)
	if alreadyCopied {
		return
	}
	chapterFrom.GongCopyBasicFields(chapterTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _section := range chapterFrom.Sections {
		chapterTo.Sections = append(chapterTo.Sections, GongCopyBranchSection(mapOrigCopy, _section))
	}
	for _, _page := range chapterFrom.Pages {
		chapterTo.Pages = append(chapterTo.Pages, GongCopyBranchPage(mapOrigCopy, _page))
	}
	for _, _chapter := range chapterFrom.SubChapters {
		chapterTo.SubChapters = append(chapterTo.SubChapters, GongCopyBranchChapter(mapOrigCopy, _chapter))
	}

	return
}

func GongCopyBranchContent(mapOrigCopy map[any]any, contentFrom *Content) (contentTo *Content) {
	var alreadyCopied bool
	contentTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, contentFrom)
	if alreadyCopied {
		return
	}
	contentFrom.GongCopyBasicFields(contentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _chapter := range contentFrom.Chapters {
		contentTo.Chapters = append(contentTo.Chapters, GongCopyBranchChapter(mapOrigCopy, _chapter))
	}

	return
}

func GongCopyBranchDownloadableFile(mapOrigCopy map[any]any, downloadablefileFrom *DownloadableFile) (downloadablefileTo *DownloadableFile) {
	var alreadyCopied bool
	downloadablefileTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, downloadablefileFrom)
	if alreadyCopied {
		return
	}
	downloadablefileFrom.GongCopyBasicFields(downloadablefileTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchJpgImage(mapOrigCopy map[any]any, jpgimageFrom *JpgImage) (jpgimageTo *JpgImage) {
	var alreadyCopied bool
	jpgimageTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, jpgimageFrom)
	if alreadyCopied {
		return
	}
	jpgimageFrom.GongCopyBasicFields(jpgimageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPage(mapOrigCopy map[any]any, pageFrom *Page) (pageTo *Page) {
	var alreadyCopied bool
	pageTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, pageFrom)
	if alreadyCopied {
		return
	}
	pageFrom.GongCopyBasicFields(pageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _section := range pageFrom.Sections {
		pageTo.Sections = append(pageTo.Sections, GongCopyBranchSection(mapOrigCopy, _section))
	}

	return
}

func GongCopyBranchPngImage(mapOrigCopy map[any]any, pngimageFrom *PngImage) (pngimageTo *PngImage) {
	var alreadyCopied bool
	pngimageTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, pngimageFrom)
	if alreadyCopied {
		return
	}
	pngimageFrom.GongCopyBasicFields(pngimageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSection(mapOrigCopy map[any]any, sectionFrom *Section) (sectionTo *Section) {
	var alreadyCopied bool
	sectionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, sectionFrom)
	if alreadyCopied {
		return
	}
	sectionFrom.GongCopyBasicFields(sectionTo)

	//insertion point for the staging of instances referenced by pointers
	if sectionFrom.SvgImage != nil {
		sectionTo.SvgImage = GongCopyBranchSvgImage(mapOrigCopy, sectionFrom.SvgImage)
	}
	if sectionFrom.PngImage != nil {
		sectionTo.PngImage = GongCopyBranchPngImage(mapOrigCopy, sectionFrom.PngImage)
	}
	if sectionFrom.JpgImage != nil {
		sectionTo.JpgImage = GongCopyBranchJpgImage(mapOrigCopy, sectionFrom.JpgImage)
	}
	if sectionFrom.DownloadableFile != nil {
		sectionTo.DownloadableFile = GongCopyBranchDownloadableFile(mapOrigCopy, sectionFrom.DownloadableFile)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSvgImage(mapOrigCopy map[any]any, svgimageFrom *SvgImage) (svgimageTo *SvgImage) {
	var alreadyCopied bool
	svgimageTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, svgimageFrom)
	if alreadyCopied {
		return
	}
	svgimageFrom.GongCopyBasicFields(svgimageTo)

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
func (chapter *Chapter) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(chapter) {
		return
	}

	chapter.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _section := range chapter.Sections {
		stage.UnstageBranch(_section)
	}
	for _, _page := range chapter.Pages {
		stage.UnstageBranch(_page)
	}
	for _, _chapter := range chapter.SubChapters {
		stage.UnstageBranch(_chapter)
	}

}

func (content *Content) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(content) {
		return
	}

	content.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _chapter := range content.Chapters {
		stage.UnstageBranch(_chapter)
	}

}

func (downloadablefile *DownloadableFile) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(downloadablefile) {
		return
	}

	downloadablefile.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (jpgimage *JpgImage) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(jpgimage) {
		return
	}

	jpgimage.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (page *Page) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(page) {
		return
	}

	page.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _section := range page.Sections {
		stage.UnstageBranch(_section)
	}

}

func (pngimage *PngImage) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(pngimage) {
		return
	}

	pngimage.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (section *Section) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(section) {
		return
	}

	section.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if section.SvgImage != nil {
		stage.UnstageBranch(section.SvgImage)
	}
	if section.PngImage != nil {
		stage.UnstageBranch(section.PngImage)
	}
	if section.JpgImage != nil {
		stage.UnstageBranch(section.JpgImage)
	}
	if section.DownloadableFile != nil {
		stage.UnstageBranch(section.DownloadableFile)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (svgimage *SvgImage) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(svgimage) {
		return
	}

	svgimage.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *Chapter) GongReconstructPointersFromReferences(stage *Stage, instance *Chapter) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Sections, stage.Sections_reference, instance.Sections)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Pages, stage.Pages_reference, instance.Pages)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubChapters, stage.Chapters_reference, instance.SubChapters)
}

func (reference *Content) GongReconstructPointersFromReferences(stage *Stage, instance *Content) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Chapters, stage.Chapters_reference, instance.Chapters)
}

func (reference *DownloadableFile) GongReconstructPointersFromReferences(stage *Stage, instance *DownloadableFile) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *JpgImage) GongReconstructPointersFromReferences(stage *Stage, instance *JpgImage) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Page) GongReconstructPointersFromReferences(stage *Stage, instance *Page) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Sections, stage.Sections_reference, instance.Sections)
}

func (reference *PngImage) GongReconstructPointersFromReferences(stage *Stage, instance *PngImage) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Section) GongReconstructPointersFromReferences(stage *Stage, instance *Section) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.SvgImage, stage.SvgImages_reference, instance.SvgImage)
	__gong__reconstructPointer(&reference.PngImage, stage.PngImages_reference, instance.PngImage)
	__gong__reconstructPointer(&reference.JpgImage, stage.JpgImages_reference, instance.JpgImage)
	__gong__reconstructPointer(&reference.DownloadableFile, stage.DownloadableFiles_reference, instance.DownloadableFile)
	// insertion point for slice of pointers field
}

func (reference *SvgImage) GongReconstructPointersFromReferences(stage *Stage, instance *SvgImage) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *Chapter) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Sections, stage.Sections_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Pages, stage.Pages_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubChapters, stage.Chapters_instance)
}

func (reference *Content) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Chapters, stage.Chapters_instance)
}

func (reference *DownloadableFile) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *JpgImage) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Page) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Sections, stage.Sections_instance)
}

func (reference *PngImage) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Section) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.SvgImage, stage.SvgImages_instance)
	__gong__reconstructPointerFromInstance(&reference.PngImage, stage.PngImages_instance)
	__gong__reconstructPointerFromInstance(&reference.JpgImage, stage.JpgImages_instance)
	__gong__reconstructPointerFromInstance(&reference.DownloadableFile, stage.DownloadableFiles_instance)
	// insertion point for slice of pointers fields
}

func (reference *SvgImage) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (chapter *Chapter) GongDiff(stage *Stage, chapterOther *Chapter) (diffs []string) {
	// insertion point for field diffs
	if chapter.Name != chapterOther.Name {
		diffs = append(diffs, chapter.GongMarshallField(stage, "Name"))
	}
	if chapter.MardownContent != chapterOther.MardownContent {
		diffs = append(diffs, chapter.GongMarshallField(stage, "MardownContent"))
	}
	if ops := __gong__diffSliceOfPointers(stage, chapter, "Sections", chapterOther.Sections, chapter.Sections); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, chapter, "Pages", chapterOther.Pages, chapter.Pages); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, chapter, "SubChapters", chapterOther.SubChapters, chapter.SubChapters); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (content *Content) GongDiff(stage *Stage, contentOther *Content) (diffs []string) {
	// insertion point for field diffs
	if content.Name != contentOther.Name {
		diffs = append(diffs, content.GongMarshallField(stage, "Name"))
	}
	if content.MardownContent != contentOther.MardownContent {
		diffs = append(diffs, content.GongMarshallField(stage, "MardownContent"))
	}
	if content.ContentPath != contentOther.ContentPath {
		diffs = append(diffs, content.GongMarshallField(stage, "ContentPath"))
	}
	if content.OutputPath != contentOther.OutputPath {
		diffs = append(diffs, content.GongMarshallField(stage, "OutputPath"))
	}
	if content.StaticPath != contentOther.StaticPath {
		diffs = append(diffs, content.GongMarshallField(stage, "StaticPath"))
	}
	if content.LogoSVGFile != contentOther.LogoSVGFile {
		diffs = append(diffs, content.GongMarshallField(stage, "LogoSVGFile"))
	}
	if content.IsBespokeLogoFileName != contentOther.IsBespokeLogoFileName {
		diffs = append(diffs, content.GongMarshallField(stage, "IsBespokeLogoFileName"))
	}
	if content.BespokeLogoFileName != contentOther.BespokeLogoFileName {
		diffs = append(diffs, content.GongMarshallField(stage, "BespokeLogoFileName"))
	}
	if content.IsBespokePageTileLogoFileName != contentOther.IsBespokePageTileLogoFileName {
		diffs = append(diffs, content.GongMarshallField(stage, "IsBespokePageTileLogoFileName"))
	}
	if content.BespokePageTileLogoFileName != contentOther.BespokePageTileLogoFileName {
		diffs = append(diffs, content.GongMarshallField(stage, "BespokePageTileLogoFileName"))
	}
	if content.Target != contentOther.Target {
		diffs = append(diffs, content.GongMarshallField(stage, "Target"))
	}
	if ops := __gong__diffSliceOfPointers(stage, content, "Chapters", contentOther.Chapters, content.Chapters); ops != "" {
		diffs = append(diffs, ops)
	}
	if content.VersionInfo != contentOther.VersionInfo {
		diffs = append(diffs, content.GongMarshallField(stage, "VersionInfo"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (downloadablefile *DownloadableFile) GongDiff(stage *Stage, downloadablefileOther *DownloadableFile) (diffs []string) {
	// insertion point for field diffs
	if downloadablefile.Name != downloadablefileOther.Name {
		diffs = append(diffs, downloadablefile.GongMarshallField(stage, "Name"))
	}
	if downloadablefile.Base64Content != downloadablefileOther.Base64Content {
		diffs = append(diffs, downloadablefile.GongMarshallField(stage, "Base64Content"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (jpgimage *JpgImage) GongDiff(stage *Stage, jpgimageOther *JpgImage) (diffs []string) {
	// insertion point for field diffs
	if jpgimage.Name != jpgimageOther.Name {
		diffs = append(diffs, jpgimage.GongMarshallField(stage, "Name"))
	}
	if jpgimage.Base64Content != jpgimageOther.Base64Content {
		diffs = append(diffs, jpgimage.GongMarshallField(stage, "Base64Content"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (page *Page) GongDiff(stage *Stage, pageOther *Page) (diffs []string) {
	// insertion point for field diffs
	if page.Name != pageOther.Name {
		diffs = append(diffs, page.GongMarshallField(stage, "Name"))
	}
	if page.MardownContent != pageOther.MardownContent {
		diffs = append(diffs, page.GongMarshallField(stage, "MardownContent"))
	}
	if ops := __gong__diffSliceOfPointers(stage, page, "Sections", pageOther.Sections, page.Sections); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (pngimage *PngImage) GongDiff(stage *Stage, pngimageOther *PngImage) (diffs []string) {
	// insertion point for field diffs
	if pngimage.Name != pngimageOther.Name {
		diffs = append(diffs, pngimage.GongMarshallField(stage, "Name"))
	}
	if pngimage.Base64Content != pngimageOther.Base64Content {
		diffs = append(diffs, pngimage.GongMarshallField(stage, "Base64Content"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (section *Section) GongDiff(stage *Stage, sectionOther *Section) (diffs []string) {
	// insertion point for field diffs
	if section.Name != sectionOther.Name {
		diffs = append(diffs, section.GongMarshallField(stage, "Name"))
	}
	if section.MardownContent != sectionOther.MardownContent {
		diffs = append(diffs, section.GongMarshallField(stage, "MardownContent"))
	}
	if section.IsImage != sectionOther.IsImage {
		diffs = append(diffs, section.GongMarshallField(stage, "IsImage"))
	}
	if section.SvgImage != sectionOther.SvgImage {
		diffs = append(diffs, section.GongMarshallField(stage, "SvgImage"))
	}
	if section.PngImage != sectionOther.PngImage {
		diffs = append(diffs, section.GongMarshallField(stage, "PngImage"))
	}
	if section.JpgImage != sectionOther.JpgImage {
		diffs = append(diffs, section.GongMarshallField(stage, "JpgImage"))
	}
	if section.IsDownloadableFile != sectionOther.IsDownloadableFile {
		diffs = append(diffs, section.GongMarshallField(stage, "IsDownloadableFile"))
	}
	if section.DownloadableFile != sectionOther.DownloadableFile {
		diffs = append(diffs, section.GongMarshallField(stage, "DownloadableFile"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (svgimage *SvgImage) GongDiff(stage *Stage, svgimageOther *SvgImage) (diffs []string) {
	// insertion point for field diffs
	if svgimage.Name != svgimageOther.Name {
		diffs = append(diffs, svgimage.GongMarshallField(stage, "Name"))
	}
	if svgimage.Content != svgimageOther.Content {
		diffs = append(diffs, svgimage.GongMarshallField(stage, "Content"))
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
