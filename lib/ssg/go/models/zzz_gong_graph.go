// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Chapter:
		ok = stage.IsStagedChapter(target)

	case *Content:
		ok = stage.IsStagedContent(target)

	case *DownloadableFile:
		ok = stage.IsStagedDownloadableFile(target)

	case *JpgImage:
		ok = stage.IsStagedJpgImage(target)

	case *Page:
		ok = stage.IsStagedPage(target)

	case *PngImage:
		ok = stage.IsStagedPngImage(target)

	case *Section:
		ok = stage.IsStagedSection(target)

	case *SvgImage:
		ok = stage.IsStagedSvgImage(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Chapter:
		ok = stage.IsStagedChapter(target)

	case *Content:
		ok = stage.IsStagedContent(target)

	case *DownloadableFile:
		ok = stage.IsStagedDownloadableFile(target)

	case *JpgImage:
		ok = stage.IsStagedJpgImage(target)

	case *Page:
		ok = stage.IsStagedPage(target)

	case *PngImage:
		ok = stage.IsStagedPngImage(target)

	case *Section:
		ok = stage.IsStagedSection(target)

	case *SvgImage:
		ok = stage.IsStagedSvgImage(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedChapter(chapter *Chapter) (ok bool) {

	_, ok = stage.Chapters[chapter]

	return
}

func (stage *Stage) IsStagedContent(content *Content) (ok bool) {

	_, ok = stage.Contents[content]

	return
}

func (stage *Stage) IsStagedDownloadableFile(downloadablefile *DownloadableFile) (ok bool) {

	_, ok = stage.DownloadableFiles[downloadablefile]

	return
}

func (stage *Stage) IsStagedJpgImage(jpgimage *JpgImage) (ok bool) {

	_, ok = stage.JpgImages[jpgimage]

	return
}

func (stage *Stage) IsStagedPage(page *Page) (ok bool) {

	_, ok = stage.Pages[page]

	return
}

func (stage *Stage) IsStagedPngImage(pngimage *PngImage) (ok bool) {

	_, ok = stage.PngImages[pngimage]

	return
}

func (stage *Stage) IsStagedSection(section *Section) (ok bool) {

	_, ok = stage.Sections[section]

	return
}

func (stage *Stage) IsStagedSvgImage(svgimage *SvgImage) (ok bool) {

	_, ok = stage.SvgImages[svgimage]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Chapter:
		stage.StageBranchChapter(target)

	case *Content:
		stage.StageBranchContent(target)

	case *DownloadableFile:
		stage.StageBranchDownloadableFile(target)

	case *JpgImage:
		stage.StageBranchJpgImage(target)

	case *Page:
		stage.StageBranchPage(target)

	case *PngImage:
		stage.StageBranchPngImage(target)

	case *Section:
		stage.StageBranchSection(target)

	case *SvgImage:
		stage.StageBranchSvgImage(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchChapter(chapter *Chapter) {

	// check if instance is already staged
	if IsStaged(stage, chapter) {
		return
	}

	chapter.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _page := range chapter.Pages {
		StageBranch(stage, _page)
	}

}

func (stage *Stage) StageBranchContent(content *Content) {

	// check if instance is already staged
	if IsStaged(stage, content) {
		return
	}

	content.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _chapter := range content.Chapters {
		StageBranch(stage, _chapter)
	}

}

func (stage *Stage) StageBranchDownloadableFile(downloadablefile *DownloadableFile) {

	// check if instance is already staged
	if IsStaged(stage, downloadablefile) {
		return
	}

	downloadablefile.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchJpgImage(jpgimage *JpgImage) {

	// check if instance is already staged
	if IsStaged(stage, jpgimage) {
		return
	}

	jpgimage.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPage(page *Page) {

	// check if instance is already staged
	if IsStaged(stage, page) {
		return
	}

	page.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _section := range page.Sections {
		StageBranch(stage, _section)
	}

}

func (stage *Stage) StageBranchPngImage(pngimage *PngImage) {

	// check if instance is already staged
	if IsStaged(stage, pngimage) {
		return
	}

	pngimage.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSection(section *Section) {

	// check if instance is already staged
	if IsStaged(stage, section) {
		return
	}

	section.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if section.SvgImage != nil {
		StageBranch(stage, section.SvgImage)
	}
	if section.PngImage != nil {
		StageBranch(stage, section.PngImage)
	}
	if section.JpgImage != nil {
		StageBranch(stage, section.JpgImage)
	}
	if section.DownloadableFile != nil {
		StageBranch(stage, section.DownloadableFile)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSvgImage(svgimage *SvgImage) {

	// check if instance is already staged
	if IsStaged(stage, svgimage) {
		return
	}

	svgimage.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Chapter:
		toT := CopyBranchChapter(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Content:
		toT := CopyBranchContent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DownloadableFile:
		toT := CopyBranchDownloadableFile(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *JpgImage:
		toT := CopyBranchJpgImage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Page:
		toT := CopyBranchPage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PngImage:
		toT := CopyBranchPngImage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Section:
		toT := CopyBranchSection(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SvgImage:
		toT := CopyBranchSvgImage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchChapter(mapOrigCopy map[any]any, chapterFrom *Chapter) (chapterTo *Chapter) {

	// chapterFrom has already been copied
	if _chapterTo, ok := mapOrigCopy[chapterFrom]; ok {
		chapterTo = _chapterTo.(*Chapter)
		return
	}

	chapterTo = new(Chapter)
	mapOrigCopy[chapterFrom] = chapterTo
	chapterFrom.CopyBasicFields(chapterTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _page := range chapterFrom.Pages {
		chapterTo.Pages = append(chapterTo.Pages, CopyBranchPage(mapOrigCopy, _page))
	}

	return
}

func CopyBranchContent(mapOrigCopy map[any]any, contentFrom *Content) (contentTo *Content) {

	// contentFrom has already been copied
	if _contentTo, ok := mapOrigCopy[contentFrom]; ok {
		contentTo = _contentTo.(*Content)
		return
	}

	contentTo = new(Content)
	mapOrigCopy[contentFrom] = contentTo
	contentFrom.CopyBasicFields(contentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _chapter := range contentFrom.Chapters {
		contentTo.Chapters = append(contentTo.Chapters, CopyBranchChapter(mapOrigCopy, _chapter))
	}

	return
}

func CopyBranchDownloadableFile(mapOrigCopy map[any]any, downloadablefileFrom *DownloadableFile) (downloadablefileTo *DownloadableFile) {

	// downloadablefileFrom has already been copied
	if _downloadablefileTo, ok := mapOrigCopy[downloadablefileFrom]; ok {
		downloadablefileTo = _downloadablefileTo.(*DownloadableFile)
		return
	}

	downloadablefileTo = new(DownloadableFile)
	mapOrigCopy[downloadablefileFrom] = downloadablefileTo
	downloadablefileFrom.CopyBasicFields(downloadablefileTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchJpgImage(mapOrigCopy map[any]any, jpgimageFrom *JpgImage) (jpgimageTo *JpgImage) {

	// jpgimageFrom has already been copied
	if _jpgimageTo, ok := mapOrigCopy[jpgimageFrom]; ok {
		jpgimageTo = _jpgimageTo.(*JpgImage)
		return
	}

	jpgimageTo = new(JpgImage)
	mapOrigCopy[jpgimageFrom] = jpgimageTo
	jpgimageFrom.CopyBasicFields(jpgimageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPage(mapOrigCopy map[any]any, pageFrom *Page) (pageTo *Page) {

	// pageFrom has already been copied
	if _pageTo, ok := mapOrigCopy[pageFrom]; ok {
		pageTo = _pageTo.(*Page)
		return
	}

	pageTo = new(Page)
	mapOrigCopy[pageFrom] = pageTo
	pageFrom.CopyBasicFields(pageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _section := range pageFrom.Sections {
		pageTo.Sections = append(pageTo.Sections, CopyBranchSection(mapOrigCopy, _section))
	}

	return
}

func CopyBranchPngImage(mapOrigCopy map[any]any, pngimageFrom *PngImage) (pngimageTo *PngImage) {

	// pngimageFrom has already been copied
	if _pngimageTo, ok := mapOrigCopy[pngimageFrom]; ok {
		pngimageTo = _pngimageTo.(*PngImage)
		return
	}

	pngimageTo = new(PngImage)
	mapOrigCopy[pngimageFrom] = pngimageTo
	pngimageFrom.CopyBasicFields(pngimageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSection(mapOrigCopy map[any]any, sectionFrom *Section) (sectionTo *Section) {

	// sectionFrom has already been copied
	if _sectionTo, ok := mapOrigCopy[sectionFrom]; ok {
		sectionTo = _sectionTo.(*Section)
		return
	}

	sectionTo = new(Section)
	mapOrigCopy[sectionFrom] = sectionTo
	sectionFrom.CopyBasicFields(sectionTo)

	//insertion point for the staging of instances referenced by pointers
	if sectionFrom.SvgImage != nil {
		sectionTo.SvgImage = CopyBranchSvgImage(mapOrigCopy, sectionFrom.SvgImage)
	}
	if sectionFrom.PngImage != nil {
		sectionTo.PngImage = CopyBranchPngImage(mapOrigCopy, sectionFrom.PngImage)
	}
	if sectionFrom.JpgImage != nil {
		sectionTo.JpgImage = CopyBranchJpgImage(mapOrigCopy, sectionFrom.JpgImage)
	}
	if sectionFrom.DownloadableFile != nil {
		sectionTo.DownloadableFile = CopyBranchDownloadableFile(mapOrigCopy, sectionFrom.DownloadableFile)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSvgImage(mapOrigCopy map[any]any, svgimageFrom *SvgImage) (svgimageTo *SvgImage) {

	// svgimageFrom has already been copied
	if _svgimageTo, ok := mapOrigCopy[svgimageFrom]; ok {
		svgimageTo = _svgimageTo.(*SvgImage)
		return
	}

	svgimageTo = new(SvgImage)
	mapOrigCopy[svgimageFrom] = svgimageTo
	svgimageFrom.CopyBasicFields(svgimageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Chapter:
		stage.UnstageBranchChapter(target)

	case *Content:
		stage.UnstageBranchContent(target)

	case *DownloadableFile:
		stage.UnstageBranchDownloadableFile(target)

	case *JpgImage:
		stage.UnstageBranchJpgImage(target)

	case *Page:
		stage.UnstageBranchPage(target)

	case *PngImage:
		stage.UnstageBranchPngImage(target)

	case *Section:
		stage.UnstageBranchSection(target)

	case *SvgImage:
		stage.UnstageBranchSvgImage(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchChapter(chapter *Chapter) {

	// check if instance is already staged
	if !IsStaged(stage, chapter) {
		return
	}

	chapter.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _page := range chapter.Pages {
		UnstageBranch(stage, _page)
	}

}

func (stage *Stage) UnstageBranchContent(content *Content) {

	// check if instance is already staged
	if !IsStaged(stage, content) {
		return
	}

	content.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _chapter := range content.Chapters {
		UnstageBranch(stage, _chapter)
	}

}

func (stage *Stage) UnstageBranchDownloadableFile(downloadablefile *DownloadableFile) {

	// check if instance is already staged
	if !IsStaged(stage, downloadablefile) {
		return
	}

	downloadablefile.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchJpgImage(jpgimage *JpgImage) {

	// check if instance is already staged
	if !IsStaged(stage, jpgimage) {
		return
	}

	jpgimage.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPage(page *Page) {

	// check if instance is already staged
	if !IsStaged(stage, page) {
		return
	}

	page.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _section := range page.Sections {
		UnstageBranch(stage, _section)
	}

}

func (stage *Stage) UnstageBranchPngImage(pngimage *PngImage) {

	// check if instance is already staged
	if !IsStaged(stage, pngimage) {
		return
	}

	pngimage.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSection(section *Section) {

	// check if instance is already staged
	if !IsStaged(stage, section) {
		return
	}

	section.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if section.SvgImage != nil {
		UnstageBranch(stage, section.SvgImage)
	}
	if section.PngImage != nil {
		UnstageBranch(stage, section.PngImage)
	}
	if section.JpgImage != nil {
		UnstageBranch(stage, section.JpgImage)
	}
	if section.DownloadableFile != nil {
		UnstageBranch(stage, section.DownloadableFile)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSvgImage(svgimage *SvgImage) {

	// check if instance is already staged
	if !IsStaged(stage, svgimage) {
		return
	}

	svgimage.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *Chapter) GongReconstructPointersFromReferences(stage *Stage, instance *Chapter) () {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Pages = reference.Pages[:0]
	for _, _b := range instance.Pages {
		reference.Pages = append(reference.Pages, stage.Pages_reference[_b])
	}

	return
}

func (reference *Content) GongReconstructPointersFromReferences(stage *Stage, instance *Content) () {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Chapters = reference.Chapters[:0]
	for _, _b := range instance.Chapters {
		reference.Chapters = append(reference.Chapters, stage.Chapters_reference[_b])
	}

	return
}

func (reference *DownloadableFile) GongReconstructPointersFromReferences(stage *Stage, instance *DownloadableFile) () {
	// insertion point for pointers field
	// insertion point for slice of pointers field

	return
}

func (reference *JpgImage) GongReconstructPointersFromReferences(stage *Stage, instance *JpgImage) () {
	// insertion point for pointers field
	// insertion point for slice of pointers field

	return
}

func (reference *Page) GongReconstructPointersFromReferences(stage *Stage, instance *Page) () {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Sections = reference.Sections[:0]
	for _, _b := range instance.Sections {
		reference.Sections = append(reference.Sections, stage.Sections_reference[_b])
	}

	return
}

func (reference *PngImage) GongReconstructPointersFromReferences(stage *Stage, instance *PngImage) () {
	// insertion point for pointers field
	// insertion point for slice of pointers field

	return
}

func (reference *Section) GongReconstructPointersFromReferences(stage *Stage, instance *Section) () {
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

	return
}

func (reference *SvgImage) GongReconstructPointersFromReferences(stage *Stage, instance *SvgImage) () {
	// insertion point for pointers field
	// insertion point for slice of pointers field

	return
}

// insertion point for pointer reconstruction from instances
func (reference *Chapter) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Pages []*Page
	for _, _reference := range reference.Pages {
		if _instance, ok := stage.Pages_instance[_reference]; ok {
			_Pages = append(_Pages, _instance)
		}
	}
	reference.Pages = _Pages

	return
}

func (reference *Content) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Chapters []*Chapter
	for _, _reference := range reference.Chapters {
		if _instance, ok := stage.Chapters_instance[_reference]; ok {
			_Chapters = append(_Chapters, _instance)
		}
	}
	reference.Chapters = _Chapters

	return
}

func (reference *DownloadableFile) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	// insertion point for slice of pointers fields

	return
}

func (reference *JpgImage) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	// insertion point for slice of pointers fields

	return
}

func (reference *Page) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Sections []*Section
	for _, _reference := range reference.Sections {
		if _instance, ok := stage.Sections_instance[_reference]; ok {
			_Sections = append(_Sections, _instance)
		}
	}
	reference.Sections = _Sections

	return
}

func (reference *PngImage) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	// insertion point for slice of pointers fields

	return
}

func (reference *Section) GongReconstructPointersFromInstances(stage *Stage) () {
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

	return
}

func (reference *SvgImage) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	// insertion point for slice of pointers fields

	return
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
		ops := Diff(stage, chapter, chapterOther, "Pages", chapterOther.Pages, chapter.Pages)
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
		ops := Diff(stage, content, contentOther, "Chapters", contentOther.Chapters, content.Chapters)
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
		ops := Diff(stage, page, pageOther, "Sections", pageOther.Sections, page.Sections)
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

// Diff returns the sequence of operations to transform oldSlice into newSlice.
// It requires type T to be comparable (e.g., pointers, ints, strings).
func Diff[T1, T2 PointerToGongstruct](stage *Stage, a, b T1, fieldName string, oldSlice, newSlice []T2) (ops string) {
	m, n := len(oldSlice), len(newSlice)

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if oldSlice[i] == newSlice[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if oldSlice[i-1] == newSlice[j-1] {
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

	// Create a temporary view of what's left after deletions for tracking matches
	var currentLCS []T2
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			currentLCS = append(currentLCS, oldSlice[k])
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k, targetVal := range newSlice {
		if lcsIdx < len(currentLCS) && currentLCS[lcsIdx] == targetVal {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, targetVal.GongGetIdentifier(stage))
		}
	}

	return ops
}
