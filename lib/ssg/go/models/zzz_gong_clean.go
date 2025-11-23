// generated code - do not edit
package models

// Clean computes the reverse map, for all intances, for all clean to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) Clean() {
	// insertion point per named struct
	// Compute reverse map for named struct Chapter
	for chapter := range stage.Chapters {
		_ = chapter
		// insertion point per field
		var _Pages []*Page
		for _, _page := range chapter.Pages {
			if IsStaged(stage, _page) {
			 	_Pages = append(_Pages, _page)
			}
		}
		chapter.Pages = _Pages
		// insertion point per field
	}

	// Compute reverse map for named struct Content
	for content := range stage.Contents {
		_ = content
		// insertion point per field
		var _Chapters []*Chapter
		for _, _chapter := range content.Chapters {
			if IsStaged(stage, _chapter) {
			 	_Chapters = append(_Chapters, _chapter)
			}
		}
		content.Chapters = _Chapters
		// insertion point per field
	}

	// Compute reverse map for named struct Page
	for page := range stage.Pages {
		_ = page
		// insertion point per field
		// insertion point per field
	}

}
