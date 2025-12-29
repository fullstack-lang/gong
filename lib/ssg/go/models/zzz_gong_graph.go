// generated code - do not edit
package models

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Chapter:
		ok = stage.IsStagedChapter(target)

	case *Content:
		ok = stage.IsStagedContent(target)

	case *Page:
		ok = stage.IsStagedPage(target)

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

	case *Page:
		ok = stage.IsStagedPage(target)

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

func (stage *Stage) IsStagedPage(page *Page) (ok bool) {

	_, ok = stage.Pages[page]

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

	case *Page:
		stage.StageBranchPage(target)

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

func (stage *Stage) StageBranchPage(page *Page) {

	// check if instance is already staged
	if IsStaged(stage, page) {
		return
	}

	page.Stage(stage)

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

	case *Page:
		toT := CopyBranchPage(mapOrigCopy, fromT)
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

	case *Page:
		stage.UnstageBranchPage(target)

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

func (stage *Stage) UnstageBranchPage(page *Page) {

	// check if instance is already staged
	if !IsStaged(stage, page) {
		return
	}

	page.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (chapter *Chapter) GongDiff(chapterOther *Chapter) (diffs []string) {
	// insertion point for field diffs
	if chapter.Name != chapterOther.Name {
		diffs = append(diffs, "Name")
	}
	if chapter.MardownContent != chapterOther.MardownContent {
		diffs = append(diffs, "MardownContent")
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
				if len(chapter.Pages[i].GongDiff(chapterOther.Pages[i])) > 0 {
					PagesDifferent = true
					break
				}
			}
		}
	}
	if PagesDifferent {
		diffs = append(diffs, "Pages")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (content *Content) GongDiff(contentOther *Content) (diffs []string) {
	// insertion point for field diffs
	if content.Name != contentOther.Name {
		diffs = append(diffs, "Name")
	}
	if content.MardownContent != contentOther.MardownContent {
		diffs = append(diffs, "MardownContent")
	}
	if content.ContentPath != contentOther.ContentPath {
		diffs = append(diffs, "ContentPath")
	}
	if content.OutputPath != contentOther.OutputPath {
		diffs = append(diffs, "OutputPath")
	}
	if content.LayoutPath != contentOther.LayoutPath {
		diffs = append(diffs, "LayoutPath")
	}
	if content.StaticPath != contentOther.StaticPath {
		diffs = append(diffs, "StaticPath")
	}
	if content.IsBespokeLogoFileName != contentOther.IsBespokeLogoFileName {
		diffs = append(diffs, "IsBespokeLogoFileName")
	}
	if content.BespokeLogoFileName != contentOther.BespokeLogoFileName {
		diffs = append(diffs, "BespokeLogoFileName")
	}
	if content.IsBespokePageTileLogoFileName != contentOther.IsBespokePageTileLogoFileName {
		diffs = append(diffs, "IsBespokePageTileLogoFileName")
	}
	if content.BespokePageTileLogoFileName != contentOther.BespokePageTileLogoFileName {
		diffs = append(diffs, "BespokePageTileLogoFileName")
	}
	if content.Target != contentOther.Target {
		diffs = append(diffs, "Target")
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
				if len(content.Chapters[i].GongDiff(contentOther.Chapters[i])) > 0 {
					ChaptersDifferent = true
					break
				}
			}
		}
	}
	if ChaptersDifferent {
		diffs = append(diffs, "Chapters")
	}
	if content.VersionInfo != contentOther.VersionInfo {
		diffs = append(diffs, "VersionInfo")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (page *Page) GongDiff(pageOther *Page) (diffs []string) {
	// insertion point for field diffs
	if page.Name != pageOther.Name {
		diffs = append(diffs, "Name")
	}
	if page.MardownContent != pageOther.MardownContent {
		diffs = append(diffs, "MardownContent")
	}

	return
}
