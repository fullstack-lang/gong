// generated code - do not edit
package models

import "time"

// CleanSlice is the Stage method that removes unstaged elements from a slice of pointers.
func (stage *Stage) CleanSlice[T PointerToGongstruct](slice *[]T) (modified bool) {
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
func (stage *Stage) CleanPointer[T PointerToGongstruct](element *T) (modified bool) {
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
// Clean garbage collect unstaged instances that are referenced by Animate
func (animate *Animate) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Circle
func (circle *Circle) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&circle.Animations) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Condition
func (condition *Condition) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ControlPoint
func (controlpoint *ControlPoint) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&controlpoint.ClosestRect) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Ellipse
func (ellipse *Ellipse) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&ellipse.Animates) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FileToDownload
func (filetodownload *FileToDownload) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Layer
func (layer *Layer) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&layer.Rects) || modified
	modified = stage.CleanSlice(&layer.Texts) || modified
	modified = stage.CleanSlice(&layer.Circles) || modified
	modified = stage.CleanSlice(&layer.Lines) || modified
	modified = stage.CleanSlice(&layer.Ellipses) || modified
	modified = stage.CleanSlice(&layer.Polylines) || modified
	modified = stage.CleanSlice(&layer.Polygones) || modified
	modified = stage.CleanSlice(&layer.Paths) || modified
	modified = stage.CleanSlice(&layer.Links) || modified
	modified = stage.CleanSlice(&layer.RectLinkLinks) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Line
func (line *Line) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&line.Animates) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Link
func (link *Link) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&link.TextAtArrowStart) || modified
	modified = stage.CleanSlice(&link.TextAtArrowEnd) || modified
	modified = stage.CleanSlice(&link.TextAtCorner) || modified
	modified = stage.CleanSlice(&link.PathAtArrowStart) || modified
	modified = stage.CleanSlice(&link.PathAtArrowEnd) || modified
	modified = stage.CleanSlice(&link.PathAtCorner) || modified
	modified = stage.CleanSlice(&link.ControlPoints) || modified
	// insertion point per field
	modified = stage.CleanPointer(&link.Start) || modified
	modified = stage.CleanPointer(&link.End) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LinkAnchoredPath
func (linkanchoredpath *LinkAnchoredPath) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by LinkAnchoredText
func (linkanchoredtext *LinkAnchoredText) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&linkanchoredtext.Animates) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Path
func (path *Path) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&path.Animates) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Point
func (point *Point) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Polygone
func (polygone *Polygone) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&polygone.Animates) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Polyline
func (polyline *Polyline) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&polyline.Animates) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Rect
func (rect *Rect) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&rect.Peers) || modified
	modified = stage.CleanSlice(&rect.Obstacles) || modified
	modified = stage.CleanSlice(&rect.HoveringTrigger) || modified
	modified = stage.CleanSlice(&rect.DisplayConditions) || modified
	modified = stage.CleanSlice(&rect.Animations) || modified
	modified = stage.CleanSlice(&rect.RectAnchoredTexts) || modified
	modified = stage.CleanSlice(&rect.RectAnchoredRects) || modified
	modified = stage.CleanSlice(&rect.RectAnchoredPaths) || modified
	modified = stage.CleanSlice(&rect.RectAnchoredPngImages) || modified
	// insertion point per field
	modified = stage.CleanPointer(&rect.EnclosingRect) || modified
	modified = stage.CleanPointer(&rect.AnchoredTo) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RectAnchoredPath
func (rectanchoredpath *RectAnchoredPath) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RectAnchoredPngImage
func (rectanchoredpngimage *RectAnchoredPngImage) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RectAnchoredRect
func (rectanchoredrect *RectAnchoredRect) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RectAnchoredText
func (rectanchoredtext *RectAnchoredText) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&rectanchoredtext.Animates) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RectLinkLink
func (rectlinklink *RectLinkLink) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&rectlinklink.Start) || modified
	modified = stage.CleanPointer(&rectlinklink.End) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SVG
func (svg *SVG) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&svg.Layers) || modified
	// insertion point per field
	modified = stage.CleanPointer(&svg.StartRect) || modified
	modified = stage.CleanPointer(&svg.EndRect) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SvgText
func (svgtext *SvgText) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Text
func (text *Text) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&text.Animates) || modified
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
