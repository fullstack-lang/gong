// generated code - do not edit
package models

import "time"

// CleanSlice is the Stage method that removes unstaged elements from a slice of pointers.
func (stage *Stage) CleanSlice[T GongstructPtr](slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if stage.IsStaged(element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// CleanPointer is the Stage method that sets the pointer to nil if the referenced element is not staged.
func (stage *Stage) CleanPointer[T GongstructPtr](element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !stage.IsStaged(*element) {
		*element = zero
		modified = true
		return
	}
	return
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by Chapter
func (chapter *Chapter) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&chapter.Sections) || modified
	modified = stage.CleanSlice(&chapter.Pages) || modified
	modified = stage.CleanSlice(&chapter.SubChapters) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Content
func (content *Content) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&content.Chapters) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DownloadableFile
func (downloadablefile *DownloadableFile) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by JpgImage
func (jpgimage *JpgImage) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Page
func (page *Page) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&page.Sections) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PngImage
func (pngimage *PngImage) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Section
func (section *Section) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&section.SvgImage) || modified
	modified = stage.CleanPointer(&section.PngImage) || modified
	modified = stage.CleanPointer(&section.JpgImage) || modified
	modified = stage.CleanPointer(&section.DownloadableFile) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SvgImage
func (svgimage *SvgImage) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		modified = instance.GongClean(stage) || modified
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
