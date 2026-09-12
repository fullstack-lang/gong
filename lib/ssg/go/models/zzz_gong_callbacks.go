// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

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

// AfterCreateFromFront is a backward-compatible package-level forwarder.
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {
	stage.AfterCreateFromFront(instance)
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront[Type Gongstruct](old, new *Type) {

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

// OnAfterUpdateFromFront is a backward-compatible package-level forwarder.
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {
	stage.OnAfterUpdateFromFront(old, new)
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront[Type Gongstruct](staged, front *Type) {

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

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
