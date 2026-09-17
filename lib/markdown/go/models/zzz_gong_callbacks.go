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

