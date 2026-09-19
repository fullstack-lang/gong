// generated code - do not edit
package models

import "fmt"

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (chapter *Chapter) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Chapters[chapter]

	return
}

func (stage *Stage) IsStagedChapter(chapter *Chapter) (ok bool) {

	return chapter.GongIsStaged(stage)
}

func (content *Content) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Contents[content]

	return
}

func (stage *Stage) IsStagedContent(content *Content) (ok bool) {

	return content.GongIsStaged(stage)
}

func (downloadablefile *DownloadableFile) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.DownloadableFiles[downloadablefile]

	return
}

func (stage *Stage) IsStagedDownloadableFile(downloadablefile *DownloadableFile) (ok bool) {

	return downloadablefile.GongIsStaged(stage)
}

func (jpgimage *JpgImage) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.JpgImages[jpgimage]

	return
}

func (stage *Stage) IsStagedJpgImage(jpgimage *JpgImage) (ok bool) {

	return jpgimage.GongIsStaged(stage)
}

func (page *Page) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Pages[page]

	return
}

func (stage *Stage) IsStagedPage(page *Page) (ok bool) {

	return page.GongIsStaged(stage)
}

func (pngimage *PngImage) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PngImages[pngimage]

	return
}

func (stage *Stage) IsStagedPngImage(pngimage *PngImage) (ok bool) {

	return pngimage.GongIsStaged(stage)
}

func (section *Section) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Sections[section]

	return
}

func (stage *Stage) IsStagedSection(section *Section) (ok bool) {

	return section.GongIsStaged(stage)
}

func (svgimage *SvgImage) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.SvgImages[svgimage]

	return
}

func (stage *Stage) IsStagedSvgImage(svgimage *SvgImage) (ok bool) {

	return svgimage.GongIsStaged(stage)
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// StageBranch is a backward-compatible package-level forwarder.
func StageBranch(stage *Stage, instance GongstructIF) {
	stage.StageBranch(instance)
}

// insertion point for stage branch per struct
func (chapter *Chapter) GongStageBranch(stage *Stage) {
	stage.StageBranchChapter(chapter)
}

func (stage *Stage) StageBranchChapter(chapter *Chapter) {

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
	stage.StageBranchContent(content)
}

func (stage *Stage) StageBranchContent(content *Content) {

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
	stage.StageBranchDownloadableFile(downloadablefile)
}

func (stage *Stage) StageBranchDownloadableFile(downloadablefile *DownloadableFile) {

	// check if instance is already staged
	if stage.IsStaged(downloadablefile) {
		return
	}

	downloadablefile.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (jpgimage *JpgImage) GongStageBranch(stage *Stage) {
	stage.StageBranchJpgImage(jpgimage)
}

func (stage *Stage) StageBranchJpgImage(jpgimage *JpgImage) {

	// check if instance is already staged
	if stage.IsStaged(jpgimage) {
		return
	}

	jpgimage.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (page *Page) GongStageBranch(stage *Stage) {
	stage.StageBranchPage(page)
}

func (stage *Stage) StageBranchPage(page *Page) {

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
	stage.StageBranchPngImage(pngimage)
}

func (stage *Stage) StageBranchPngImage(pngimage *PngImage) {

	// check if instance is already staged
	if stage.IsStaged(pngimage) {
		return
	}

	pngimage.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (section *Section) GongStageBranch(stage *Stage) {
	stage.StageBranchSection(section)
}

func (stage *Stage) StageBranchSection(section *Section) {

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
	stage.StageBranchSvgImage(svgimage)
}

func (stage *Stage) StageBranchSvgImage(svgimage *SvgImage) {

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

	// chapterFrom has already been copied
	if _chapterTo, ok := mapOrigCopy[chapterFrom]; ok {
		chapterTo = _chapterTo.(*Chapter)
		return
	}

	chapterTo = new(Chapter)
	mapOrigCopy[chapterFrom] = chapterTo
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

	// contentFrom has already been copied
	if _contentTo, ok := mapOrigCopy[contentFrom]; ok {
		contentTo = _contentTo.(*Content)
		return
	}

	contentTo = new(Content)
	mapOrigCopy[contentFrom] = contentTo
	contentFrom.GongCopyBasicFields(contentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _chapter := range contentFrom.Chapters {
		contentTo.Chapters = append(contentTo.Chapters, GongCopyBranchChapter(mapOrigCopy, _chapter))
	}

	return
}

func GongCopyBranchDownloadableFile(mapOrigCopy map[any]any, downloadablefileFrom *DownloadableFile) (downloadablefileTo *DownloadableFile) {

	// downloadablefileFrom has already been copied
	if _downloadablefileTo, ok := mapOrigCopy[downloadablefileFrom]; ok {
		downloadablefileTo = _downloadablefileTo.(*DownloadableFile)
		return
	}

	downloadablefileTo = new(DownloadableFile)
	mapOrigCopy[downloadablefileFrom] = downloadablefileTo
	downloadablefileFrom.GongCopyBasicFields(downloadablefileTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchJpgImage(mapOrigCopy map[any]any, jpgimageFrom *JpgImage) (jpgimageTo *JpgImage) {

	// jpgimageFrom has already been copied
	if _jpgimageTo, ok := mapOrigCopy[jpgimageFrom]; ok {
		jpgimageTo = _jpgimageTo.(*JpgImage)
		return
	}

	jpgimageTo = new(JpgImage)
	mapOrigCopy[jpgimageFrom] = jpgimageTo
	jpgimageFrom.GongCopyBasicFields(jpgimageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPage(mapOrigCopy map[any]any, pageFrom *Page) (pageTo *Page) {

	// pageFrom has already been copied
	if _pageTo, ok := mapOrigCopy[pageFrom]; ok {
		pageTo = _pageTo.(*Page)
		return
	}

	pageTo = new(Page)
	mapOrigCopy[pageFrom] = pageTo
	pageFrom.GongCopyBasicFields(pageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _section := range pageFrom.Sections {
		pageTo.Sections = append(pageTo.Sections, GongCopyBranchSection(mapOrigCopy, _section))
	}

	return
}

func GongCopyBranchPngImage(mapOrigCopy map[any]any, pngimageFrom *PngImage) (pngimageTo *PngImage) {

	// pngimageFrom has already been copied
	if _pngimageTo, ok := mapOrigCopy[pngimageFrom]; ok {
		pngimageTo = _pngimageTo.(*PngImage)
		return
	}

	pngimageTo = new(PngImage)
	mapOrigCopy[pngimageFrom] = pngimageTo
	pngimageFrom.GongCopyBasicFields(pngimageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSection(mapOrigCopy map[any]any, sectionFrom *Section) (sectionTo *Section) {

	// sectionFrom has already been copied
	if _sectionTo, ok := mapOrigCopy[sectionFrom]; ok {
		sectionTo = _sectionTo.(*Section)
		return
	}

	sectionTo = new(Section)
	mapOrigCopy[sectionFrom] = sectionTo
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

	// svgimageFrom has already been copied
	if _svgimageTo, ok := mapOrigCopy[svgimageFrom]; ok {
		svgimageTo = _svgimageTo.(*SvgImage)
		return
	}

	svgimageTo = new(SvgImage)
	mapOrigCopy[svgimageFrom] = svgimageTo
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

// UnstageBranch is a backward-compatible package-level forwarder.
func UnstageBranch(stage *Stage, instance GongstructIF) {
	stage.UnstageBranch(instance)
}

// insertion point for unstage branch per struct
func (chapter *Chapter) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchChapter(chapter)
}

func (stage *Stage) UnstageBranchChapter(chapter *Chapter) {

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
	stage.UnstageBranchContent(content)
}

func (stage *Stage) UnstageBranchContent(content *Content) {

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
	stage.UnstageBranchDownloadableFile(downloadablefile)
}

func (stage *Stage) UnstageBranchDownloadableFile(downloadablefile *DownloadableFile) {

	// check if instance is already staged
	if !stage.IsStaged(downloadablefile) {
		return
	}

	downloadablefile.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (jpgimage *JpgImage) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchJpgImage(jpgimage)
}

func (stage *Stage) UnstageBranchJpgImage(jpgimage *JpgImage) {

	// check if instance is already staged
	if !stage.IsStaged(jpgimage) {
		return
	}

	jpgimage.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (page *Page) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPage(page)
}

func (stage *Stage) UnstageBranchPage(page *Page) {

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
	stage.UnstageBranchPngImage(pngimage)
}

func (stage *Stage) UnstageBranchPngImage(pngimage *PngImage) {

	// check if instance is already staged
	if !stage.IsStaged(pngimage) {
		return
	}

	pngimage.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (section *Section) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchSection(section)
}

func (stage *Stage) UnstageBranchSection(section *Section) {

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
	stage.UnstageBranchSvgImage(svgimage)
}

func (stage *Stage) UnstageBranchSvgImage(svgimage *SvgImage) {

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
	reference.Sections = reference.Sections[:0]
	for _, _b := range instance.Sections {
		reference.Sections = append(reference.Sections, stage.Sections_reference[_b])
	}
	reference.Pages = reference.Pages[:0]
	for _, _b := range instance.Pages {
		reference.Pages = append(reference.Pages, stage.Pages_reference[_b])
	}
	reference.SubChapters = reference.SubChapters[:0]
	for _, _b := range instance.SubChapters {
		reference.SubChapters = append(reference.SubChapters, stage.Chapters_reference[_b])
	}
}

func (reference *Content) GongReconstructPointersFromReferences(stage *Stage, instance *Content) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Chapters = reference.Chapters[:0]
	for _, _b := range instance.Chapters {
		reference.Chapters = append(reference.Chapters, stage.Chapters_reference[_b])
	}
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
	reference.Sections = reference.Sections[:0]
	for _, _b := range instance.Sections {
		reference.Sections = append(reference.Sections, stage.Sections_reference[_b])
	}
}

func (reference *PngImage) GongReconstructPointersFromReferences(stage *Stage, instance *PngImage) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Section) GongReconstructPointersFromReferences(stage *Stage, instance *Section) {
	// insertion point for pointers field
	if instance.SvgImage != nil {
		reference.SvgImage = stage.SvgImages_reference[instance.SvgImage]
	}
	if instance.PngImage != nil {
		reference.PngImage = stage.PngImages_reference[instance.PngImage]
	}
	if instance.JpgImage != nil {
		reference.JpgImage = stage.JpgImages_reference[instance.JpgImage]
	}
	if instance.DownloadableFile != nil {
		reference.DownloadableFile = stage.DownloadableFiles_reference[instance.DownloadableFile]
	}
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
	var _Sections []*Section
	for _, _reference := range reference.Sections {
		if _instance, ok := stage.Sections_instance[_reference]; ok {
			_Sections = append(_Sections, _instance)
		}
	}
	reference.Sections = _Sections
	var _Pages []*Page
	for _, _reference := range reference.Pages {
		if _instance, ok := stage.Pages_instance[_reference]; ok {
			_Pages = append(_Pages, _instance)
		}
	}
	reference.Pages = _Pages
	var _SubChapters []*Chapter
	for _, _reference := range reference.SubChapters {
		if _instance, ok := stage.Chapters_instance[_reference]; ok {
			_SubChapters = append(_SubChapters, _instance)
		}
	}
	reference.SubChapters = _SubChapters
}

func (reference *Content) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Chapters []*Chapter
	for _, _reference := range reference.Chapters {
		if _instance, ok := stage.Chapters_instance[_reference]; ok {
			_Chapters = append(_Chapters, _instance)
		}
	}
	reference.Chapters = _Chapters
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
	var _Sections []*Section
	for _, _reference := range reference.Sections {
		if _instance, ok := stage.Sections_instance[_reference]; ok {
			_Sections = append(_Sections, _instance)
		}
	}
	reference.Sections = _Sections
}

func (reference *PngImage) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Section) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.SvgImage; _reference != nil {
		reference.SvgImage = nil
		if _instance, ok := stage.SvgImages_instance[_reference]; ok {
			reference.SvgImage = _instance
		}
	}
	if _reference := reference.PngImage; _reference != nil {
		reference.PngImage = nil
		if _instance, ok := stage.PngImages_instance[_reference]; ok {
			reference.PngImage = _instance
		}
	}
	if _reference := reference.JpgImage; _reference != nil {
		reference.JpgImage = nil
		if _instance, ok := stage.JpgImages_instance[_reference]; ok {
			reference.JpgImage = _instance
		}
	}
	if _reference := reference.DownloadableFile; _reference != nil {
		reference.DownloadableFile = nil
		if _instance, ok := stage.DownloadableFiles_instance[_reference]; ok {
			reference.DownloadableFile = _instance
		}
	}
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
	SectionsDifferent := false
	if len(chapter.Sections) != len(chapterOther.Sections) {
		SectionsDifferent = true
	} else {
		for i := range chapter.Sections {
			if (chapter.Sections[i] == nil) != (chapterOther.Sections[i] == nil) {
				SectionsDifferent = true
				break
			} else if chapter.Sections[i] != nil && chapterOther.Sections[i] != nil {
				// this is a pointer comparaison
				if chapter.Sections[i] != chapterOther.Sections[i] {
					SectionsDifferent = true
					break
				}
			}
		}
	}
	if SectionsDifferent {
		ops := stage.Diff(
			chapter,
			"Sections",
			len(chapterOther.Sections),
			len(chapter.Sections),
			func(i, j int) bool {
				return chapterOther.Sections[i] == chapter.Sections[j]
			},
			func(j int) string {
				return chapter.Sections[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	PagesDifferent := false
	if len(chapter.Pages) != len(chapterOther.Pages) {
		PagesDifferent = true
	} else {
		for i := range chapter.Pages {
			if (chapter.Pages[i] == nil) != (chapterOther.Pages[i] == nil) {
				PagesDifferent = true
				break
			} else if chapter.Pages[i] != nil && chapterOther.Pages[i] != nil {
				// this is a pointer comparaison
				if chapter.Pages[i] != chapterOther.Pages[i] {
					PagesDifferent = true
					break
				}
			}
		}
	}
	if PagesDifferent {
		ops := stage.Diff(
			chapter,
			"Pages",
			len(chapterOther.Pages),
			len(chapter.Pages),
			func(i, j int) bool {
				return chapterOther.Pages[i] == chapter.Pages[j]
			},
			func(j int) string {
				return chapter.Pages[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	SubChaptersDifferent := false
	if len(chapter.SubChapters) != len(chapterOther.SubChapters) {
		SubChaptersDifferent = true
	} else {
		for i := range chapter.SubChapters {
			if (chapter.SubChapters[i] == nil) != (chapterOther.SubChapters[i] == nil) {
				SubChaptersDifferent = true
				break
			} else if chapter.SubChapters[i] != nil && chapterOther.SubChapters[i] != nil {
				// this is a pointer comparaison
				if chapter.SubChapters[i] != chapterOther.SubChapters[i] {
					SubChaptersDifferent = true
					break
				}
			}
		}
	}
	if SubChaptersDifferent {
		ops := stage.Diff(
			chapter,
			"SubChapters",
			len(chapterOther.SubChapters),
			len(chapter.SubChapters),
			func(i, j int) bool {
				return chapterOther.SubChapters[i] == chapter.SubChapters[j]
			},
			func(j int) string {
				return chapter.SubChapters[j].GongGetIdentifier(stage)
			},
		)
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
	ChaptersDifferent := false
	if len(content.Chapters) != len(contentOther.Chapters) {
		ChaptersDifferent = true
	} else {
		for i := range content.Chapters {
			if (content.Chapters[i] == nil) != (contentOther.Chapters[i] == nil) {
				ChaptersDifferent = true
				break
			} else if content.Chapters[i] != nil && contentOther.Chapters[i] != nil {
				// this is a pointer comparaison
				if content.Chapters[i] != contentOther.Chapters[i] {
					ChaptersDifferent = true
					break
				}
			}
		}
	}
	if ChaptersDifferent {
		ops := stage.Diff(
			content,
			"Chapters",
			len(contentOther.Chapters),
			len(content.Chapters),
			func(i, j int) bool {
				return contentOther.Chapters[i] == content.Chapters[j]
			},
			func(j int) string {
				return content.Chapters[j].GongGetIdentifier(stage)
			},
		)
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
	SectionsDifferent := false
	if len(page.Sections) != len(pageOther.Sections) {
		SectionsDifferent = true
	} else {
		for i := range page.Sections {
			if (page.Sections[i] == nil) != (pageOther.Sections[i] == nil) {
				SectionsDifferent = true
				break
			} else if page.Sections[i] != nil && pageOther.Sections[i] != nil {
				// this is a pointer comparaison
				if page.Sections[i] != pageOther.Sections[i] {
					SectionsDifferent = true
					break
				}
			}
		}
	}
	if SectionsDifferent {
		ops := stage.Diff(
			page,
			"Sections",
			len(pageOther.Sections),
			len(page.Sections),
			func(i, j int) bool {
				return pageOther.Sections[i] == page.Sections[j]
			},
			func(j int) string {
				return page.Sections[j].GongGetIdentifier(stage)
			},
		)
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
	if (section.SvgImage == nil) != (sectionOther.SvgImage == nil) {
		diffs = append(diffs, section.GongMarshallField(stage, "SvgImage"))
	} else if section.SvgImage != nil && sectionOther.SvgImage != nil {
		if section.SvgImage != sectionOther.SvgImage {
			diffs = append(diffs, section.GongMarshallField(stage, "SvgImage"))
		}
	}
	if (section.PngImage == nil) != (sectionOther.PngImage == nil) {
		diffs = append(diffs, section.GongMarshallField(stage, "PngImage"))
	} else if section.PngImage != nil && sectionOther.PngImage != nil {
		if section.PngImage != sectionOther.PngImage {
			diffs = append(diffs, section.GongMarshallField(stage, "PngImage"))
		}
	}
	if (section.JpgImage == nil) != (sectionOther.JpgImage == nil) {
		diffs = append(diffs, section.GongMarshallField(stage, "JpgImage"))
	} else if section.JpgImage != nil && sectionOther.JpgImage != nil {
		if section.JpgImage != sectionOther.JpgImage {
			diffs = append(diffs, section.GongMarshallField(stage, "JpgImage"))
		}
	}
	if section.IsDownloadableFile != sectionOther.IsDownloadableFile {
		diffs = append(diffs, section.GongMarshallField(stage, "IsDownloadableFile"))
	}
	if (section.DownloadableFile == nil) != (sectionOther.DownloadableFile == nil) {
		diffs = append(diffs, section.GongMarshallField(stage, "DownloadableFile"))
	} else if section.DownloadableFile != nil && sectionOther.DownloadableFile != nil {
		if section.DownloadableFile != sectionOther.DownloadableFile {
			diffs = append(diffs, section.GongMarshallField(stage, "DownloadableFile"))
		}
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
