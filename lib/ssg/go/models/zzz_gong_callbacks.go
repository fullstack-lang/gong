// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point
func (chapter *Chapter) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterChapterCreateCallback != nil {
		stage.OnAfterChapterCreateCallback.OnAfterCreate(stage, chapter)
	}
}

func (chapter *Chapter) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterChapterUpdateCallback != nil {
		var frontChapter *Chapter
		if front != nil {
			frontChapter, _ = front.(*Chapter)
		}
		stage.OnAfterChapterUpdateCallback.OnAfterUpdate(stage, chapter, frontChapter)
	}
}

func (chapter *Chapter) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterChapterDeleteCallback != nil {
		var frontChapter *Chapter
		if front != nil {
			frontChapter, _ = front.(*Chapter)
		}
		stage.OnAfterChapterDeleteCallback.OnAfterDelete(stage, chapter, frontChapter)
	}
}

func (content *Content) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterContentCreateCallback != nil {
		stage.OnAfterContentCreateCallback.OnAfterCreate(stage, content)
	}
}

func (content *Content) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterContentUpdateCallback != nil {
		var frontContent *Content
		if front != nil {
			frontContent, _ = front.(*Content)
		}
		stage.OnAfterContentUpdateCallback.OnAfterUpdate(stage, content, frontContent)
	}
}

func (content *Content) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterContentDeleteCallback != nil {
		var frontContent *Content
		if front != nil {
			frontContent, _ = front.(*Content)
		}
		stage.OnAfterContentDeleteCallback.OnAfterDelete(stage, content, frontContent)
	}
}

func (downloadablefile *DownloadableFile) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDownloadableFileCreateCallback != nil {
		stage.OnAfterDownloadableFileCreateCallback.OnAfterCreate(stage, downloadablefile)
	}
}

func (downloadablefile *DownloadableFile) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDownloadableFileUpdateCallback != nil {
		var frontDownloadableFile *DownloadableFile
		if front != nil {
			frontDownloadableFile, _ = front.(*DownloadableFile)
		}
		stage.OnAfterDownloadableFileUpdateCallback.OnAfterUpdate(stage, downloadablefile, frontDownloadableFile)
	}
}

func (downloadablefile *DownloadableFile) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDownloadableFileDeleteCallback != nil {
		var frontDownloadableFile *DownloadableFile
		if front != nil {
			frontDownloadableFile, _ = front.(*DownloadableFile)
		}
		stage.OnAfterDownloadableFileDeleteCallback.OnAfterDelete(stage, downloadablefile, frontDownloadableFile)
	}
}

func (jpgimage *JpgImage) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterJpgImageCreateCallback != nil {
		stage.OnAfterJpgImageCreateCallback.OnAfterCreate(stage, jpgimage)
	}
}

func (jpgimage *JpgImage) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterJpgImageUpdateCallback != nil {
		var frontJpgImage *JpgImage
		if front != nil {
			frontJpgImage, _ = front.(*JpgImage)
		}
		stage.OnAfterJpgImageUpdateCallback.OnAfterUpdate(stage, jpgimage, frontJpgImage)
	}
}

func (jpgimage *JpgImage) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterJpgImageDeleteCallback != nil {
		var frontJpgImage *JpgImage
		if front != nil {
			frontJpgImage, _ = front.(*JpgImage)
		}
		stage.OnAfterJpgImageDeleteCallback.OnAfterDelete(stage, jpgimage, frontJpgImage)
	}
}

func (page *Page) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPageCreateCallback != nil {
		stage.OnAfterPageCreateCallback.OnAfterCreate(stage, page)
	}
}

func (page *Page) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPageUpdateCallback != nil {
		var frontPage *Page
		if front != nil {
			frontPage, _ = front.(*Page)
		}
		stage.OnAfterPageUpdateCallback.OnAfterUpdate(stage, page, frontPage)
	}
}

func (page *Page) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPageDeleteCallback != nil {
		var frontPage *Page
		if front != nil {
			frontPage, _ = front.(*Page)
		}
		stage.OnAfterPageDeleteCallback.OnAfterDelete(stage, page, frontPage)
	}
}

func (pngimage *PngImage) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPngImageCreateCallback != nil {
		stage.OnAfterPngImageCreateCallback.OnAfterCreate(stage, pngimage)
	}
}

func (pngimage *PngImage) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPngImageUpdateCallback != nil {
		var frontPngImage *PngImage
		if front != nil {
			frontPngImage, _ = front.(*PngImage)
		}
		stage.OnAfterPngImageUpdateCallback.OnAfterUpdate(stage, pngimage, frontPngImage)
	}
}

func (pngimage *PngImage) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPngImageDeleteCallback != nil {
		var frontPngImage *PngImage
		if front != nil {
			frontPngImage, _ = front.(*PngImage)
		}
		stage.OnAfterPngImageDeleteCallback.OnAfterDelete(stage, pngimage, frontPngImage)
	}
}

func (section *Section) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSectionCreateCallback != nil {
		stage.OnAfterSectionCreateCallback.OnAfterCreate(stage, section)
	}
}

func (section *Section) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSectionUpdateCallback != nil {
		var frontSection *Section
		if front != nil {
			frontSection, _ = front.(*Section)
		}
		stage.OnAfterSectionUpdateCallback.OnAfterUpdate(stage, section, frontSection)
	}
}

func (section *Section) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSectionDeleteCallback != nil {
		var frontSection *Section
		if front != nil {
			frontSection, _ = front.(*Section)
		}
		stage.OnAfterSectionDeleteCallback.OnAfterDelete(stage, section, frontSection)
	}
}

func (svgimage *SvgImage) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSvgImageCreateCallback != nil {
		stage.OnAfterSvgImageCreateCallback.OnAfterCreate(stage, svgimage)
	}
}

func (svgimage *SvgImage) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSvgImageUpdateCallback != nil {
		var frontSvgImage *SvgImage
		if front != nil {
			frontSvgImage, _ = front.(*SvgImage)
		}
		stage.OnAfterSvgImageUpdateCallback.OnAfterUpdate(stage, svgimage, frontSvgImage)
	}
}

func (svgimage *SvgImage) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSvgImageDeleteCallback != nil {
		var frontSvgImage *SvgImage
		if front != nil {
			frontSvgImage, _ = front.(*SvgImage)
		}
		stage.OnAfterSvgImageDeleteCallback.OnAfterDelete(stage, svgimage, frontSvgImage)
	}
}

