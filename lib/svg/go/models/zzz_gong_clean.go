// generated code - do not edit
package models

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	*slice = cleanedSlice
	return
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element *T) (modified bool) {
	if !IsStagedPointerToGongstruct(stage, *element) {
		var zero T
		*element = zero
		return true
	}
	return false
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
	modified = GongCleanSlice(stage, &circle.Animations) || modified
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
	modified = GongCleanPointer(stage, &controlpoint.ClosestRect) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Ellipse
func (ellipse *Ellipse) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &ellipse.Animates) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Layer
func (layer *Layer) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &layer.Rects) || modified
	modified = GongCleanSlice(stage, &layer.Texts) || modified
	modified = GongCleanSlice(stage, &layer.Circles) || modified
	modified = GongCleanSlice(stage, &layer.Lines) || modified
	modified = GongCleanSlice(stage, &layer.Ellipses) || modified
	modified = GongCleanSlice(stage, &layer.Polylines) || modified
	modified = GongCleanSlice(stage, &layer.Polygones) || modified
	modified = GongCleanSlice(stage, &layer.Paths) || modified
	modified = GongCleanSlice(stage, &layer.Links) || modified
	modified = GongCleanSlice(stage, &layer.RectLinkLinks) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Line
func (line *Line) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &line.Animates) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Link
func (link *Link) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &link.TextAtArrowStart) || modified
	modified = GongCleanSlice(stage, &link.TextAtArrowEnd) || modified
	modified = GongCleanSlice(stage, &link.ControlPoints) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &link.Start) || modified
	modified = GongCleanPointer(stage, &link.End) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LinkAnchoredText
func (linkanchoredtext *LinkAnchoredText) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &linkanchoredtext.Animates) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Path
func (path *Path) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &path.Animates) || modified
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
	modified = GongCleanSlice(stage, &polygone.Animates) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Polyline
func (polyline *Polyline) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &polyline.Animates) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Rect
func (rect *Rect) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &rect.HoveringTrigger) || modified
	modified = GongCleanSlice(stage, &rect.DisplayConditions) || modified
	modified = GongCleanSlice(stage, &rect.Animations) || modified
	modified = GongCleanSlice(stage, &rect.RectAnchoredTexts) || modified
	modified = GongCleanSlice(stage, &rect.RectAnchoredRects) || modified
	modified = GongCleanSlice(stage, &rect.RectAnchoredPaths) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RectAnchoredPath
func (rectanchoredpath *RectAnchoredPath) GongClean(stage *Stage) (modified bool) {
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
	modified = GongCleanSlice(stage, &rectanchoredtext.Animates) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RectLinkLink
func (rectlinklink *RectLinkLink) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &rectlinklink.Start) || modified
	modified = GongCleanPointer(stage, &rectlinklink.End) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SVG
func (svg *SVG) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &svg.Layers) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &svg.StartRect) || modified
	modified = GongCleanPointer(stage, &svg.EndRect) || modified
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
	modified = GongCleanSlice(stage, &text.Animates) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		modified = instance.GongClean(stage) || modified
	}
	return
}
