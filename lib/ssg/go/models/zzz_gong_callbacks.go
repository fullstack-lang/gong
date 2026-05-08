// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Chapter:
		if stage.OnAfterChapterCreateCallback != nil {
			stage.OnAfterChapterCreateCallback.OnAfterCreate(stage, target)
		}
	case *Content:
		if stage.OnAfterContentCreateCallback != nil {
			stage.OnAfterContentCreateCallback.OnAfterCreate(stage, target)
		}
	case *DownloadableFile:
		if stage.OnAfterDownloadableFileCreateCallback != nil {
			stage.OnAfterDownloadableFileCreateCallback.OnAfterCreate(stage, target)
		}
	case *JpgImage:
		if stage.OnAfterJpgImageCreateCallback != nil {
			stage.OnAfterJpgImageCreateCallback.OnAfterCreate(stage, target)
		}
	case *Page:
		if stage.OnAfterPageCreateCallback != nil {
			stage.OnAfterPageCreateCallback.OnAfterCreate(stage, target)
		}
	case *PngImage:
		if stage.OnAfterPngImageCreateCallback != nil {
			stage.OnAfterPngImageCreateCallback.OnAfterCreate(stage, target)
		}
	case *Section:
		if stage.OnAfterSectionCreateCallback != nil {
			stage.OnAfterSectionCreateCallback.OnAfterCreate(stage, target)
		}
	case *SvgImage:
		if stage.OnAfterSvgImageCreateCallback != nil {
			stage.OnAfterSvgImageCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is called after a update from front
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Chapter:
		newTarget := any(new).(*Chapter)
		if stage.OnAfterChapterUpdateCallback != nil {
			stage.OnAfterChapterUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Content:
		newTarget := any(new).(*Content)
		if stage.OnAfterContentUpdateCallback != nil {
			stage.OnAfterContentUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DownloadableFile:
		newTarget := any(new).(*DownloadableFile)
		if stage.OnAfterDownloadableFileUpdateCallback != nil {
			stage.OnAfterDownloadableFileUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *JpgImage:
		newTarget := any(new).(*JpgImage)
		if stage.OnAfterJpgImageUpdateCallback != nil {
			stage.OnAfterJpgImageUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Page:
		newTarget := any(new).(*Page)
		if stage.OnAfterPageUpdateCallback != nil {
			stage.OnAfterPageUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PngImage:
		newTarget := any(new).(*PngImage)
		if stage.OnAfterPngImageUpdateCallback != nil {
			stage.OnAfterPngImageUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Section:
		newTarget := any(new).(*Section)
		if stage.OnAfterSectionUpdateCallback != nil {
			stage.OnAfterSectionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SvgImage:
		newTarget := any(new).(*SvgImage)
		if stage.OnAfterSvgImageUpdateCallback != nil {
			stage.OnAfterSvgImageUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Chapter:
		if stage.OnAfterChapterDeleteCallback != nil {
			staged := any(staged).(*Chapter)
			stage.OnAfterChapterDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Content:
		if stage.OnAfterContentDeleteCallback != nil {
			staged := any(staged).(*Content)
			stage.OnAfterContentDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DownloadableFile:
		if stage.OnAfterDownloadableFileDeleteCallback != nil {
			staged := any(staged).(*DownloadableFile)
			stage.OnAfterDownloadableFileDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *JpgImage:
		if stage.OnAfterJpgImageDeleteCallback != nil {
			staged := any(staged).(*JpgImage)
			stage.OnAfterJpgImageDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Page:
		if stage.OnAfterPageDeleteCallback != nil {
			staged := any(staged).(*Page)
			stage.OnAfterPageDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PngImage:
		if stage.OnAfterPngImageDeleteCallback != nil {
			staged := any(staged).(*PngImage)
			stage.OnAfterPngImageDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Section:
		if stage.OnAfterSectionDeleteCallback != nil {
			staged := any(staged).(*Section)
			stage.OnAfterSectionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SvgImage:
		if stage.OnAfterSvgImageDeleteCallback != nil {
			staged := any(staged).(*SvgImage)
			stage.OnAfterSvgImageDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Chapter:
		if stage.OnAfterChapterReadCallback != nil {
			stage.OnAfterChapterReadCallback.OnAfterRead(stage, target)
		}
	case *Content:
		if stage.OnAfterContentReadCallback != nil {
			stage.OnAfterContentReadCallback.OnAfterRead(stage, target)
		}
	case *DownloadableFile:
		if stage.OnAfterDownloadableFileReadCallback != nil {
			stage.OnAfterDownloadableFileReadCallback.OnAfterRead(stage, target)
		}
	case *JpgImage:
		if stage.OnAfterJpgImageReadCallback != nil {
			stage.OnAfterJpgImageReadCallback.OnAfterRead(stage, target)
		}
	case *Page:
		if stage.OnAfterPageReadCallback != nil {
			stage.OnAfterPageReadCallback.OnAfterRead(stage, target)
		}
	case *PngImage:
		if stage.OnAfterPngImageReadCallback != nil {
			stage.OnAfterPngImageReadCallback.OnAfterRead(stage, target)
		}
	case *Section:
		if stage.OnAfterSectionReadCallback != nil {
			stage.OnAfterSectionReadCallback.OnAfterRead(stage, target)
		}
	case *SvgImage:
		if stage.OnAfterSvgImageReadCallback != nil {
			stage.OnAfterSvgImageReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *Stage, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *Chapter:
		stage.OnAfterChapterUpdateCallback = any(callback).(OnAfterUpdateInterface[Chapter])
	case *Content:
		stage.OnAfterContentUpdateCallback = any(callback).(OnAfterUpdateInterface[Content])
	case *DownloadableFile:
		stage.OnAfterDownloadableFileUpdateCallback = any(callback).(OnAfterUpdateInterface[DownloadableFile])
	case *JpgImage:
		stage.OnAfterJpgImageUpdateCallback = any(callback).(OnAfterUpdateInterface[JpgImage])
	case *Page:
		stage.OnAfterPageUpdateCallback = any(callback).(OnAfterUpdateInterface[Page])
	case *PngImage:
		stage.OnAfterPngImageUpdateCallback = any(callback).(OnAfterUpdateInterface[PngImage])
	case *Section:
		stage.OnAfterSectionUpdateCallback = any(callback).(OnAfterUpdateInterface[Section])
	case *SvgImage:
		stage.OnAfterSvgImageUpdateCallback = any(callback).(OnAfterUpdateInterface[SvgImage])
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *Chapter:
		stage.OnAfterChapterCreateCallback = any(callback).(OnAfterCreateInterface[Chapter])
	case *Content:
		stage.OnAfterContentCreateCallback = any(callback).(OnAfterCreateInterface[Content])
	case *DownloadableFile:
		stage.OnAfterDownloadableFileCreateCallback = any(callback).(OnAfterCreateInterface[DownloadableFile])
	case *JpgImage:
		stage.OnAfterJpgImageCreateCallback = any(callback).(OnAfterCreateInterface[JpgImage])
	case *Page:
		stage.OnAfterPageCreateCallback = any(callback).(OnAfterCreateInterface[Page])
	case *PngImage:
		stage.OnAfterPngImageCreateCallback = any(callback).(OnAfterCreateInterface[PngImage])
	case *Section:
		stage.OnAfterSectionCreateCallback = any(callback).(OnAfterCreateInterface[Section])
	case *SvgImage:
		stage.OnAfterSvgImageCreateCallback = any(callback).(OnAfterCreateInterface[SvgImage])
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *Chapter:
		stage.OnAfterChapterDeleteCallback = any(callback).(OnAfterDeleteInterface[Chapter])
	case *Content:
		stage.OnAfterContentDeleteCallback = any(callback).(OnAfterDeleteInterface[Content])
	case *DownloadableFile:
		stage.OnAfterDownloadableFileDeleteCallback = any(callback).(OnAfterDeleteInterface[DownloadableFile])
	case *JpgImage:
		stage.OnAfterJpgImageDeleteCallback = any(callback).(OnAfterDeleteInterface[JpgImage])
	case *Page:
		stage.OnAfterPageDeleteCallback = any(callback).(OnAfterDeleteInterface[Page])
	case *PngImage:
		stage.OnAfterPngImageDeleteCallback = any(callback).(OnAfterDeleteInterface[PngImage])
	case *Section:
		stage.OnAfterSectionDeleteCallback = any(callback).(OnAfterDeleteInterface[Section])
	case *SvgImage:
		stage.OnAfterSvgImageDeleteCallback = any(callback).(OnAfterDeleteInterface[SvgImage])
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *Chapter:
		stage.OnAfterChapterReadCallback = any(callback).(OnAfterReadInterface[Chapter])
	case *Content:
		stage.OnAfterContentReadCallback = any(callback).(OnAfterReadInterface[Content])
	case *DownloadableFile:
		stage.OnAfterDownloadableFileReadCallback = any(callback).(OnAfterReadInterface[DownloadableFile])
	case *JpgImage:
		stage.OnAfterJpgImageReadCallback = any(callback).(OnAfterReadInterface[JpgImage])
	case *Page:
		stage.OnAfterPageReadCallback = any(callback).(OnAfterReadInterface[Page])
	case *PngImage:
		stage.OnAfterPngImageReadCallback = any(callback).(OnAfterReadInterface[PngImage])
	case *Section:
		stage.OnAfterSectionReadCallback = any(callback).(OnAfterReadInterface[Section])
	case *SvgImage:
		stage.OnAfterSvgImageReadCallback = any(callback).(OnAfterReadInterface[SvgImage])
	}
}
