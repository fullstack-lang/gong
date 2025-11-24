// generated code - do not edit
package models

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice []T) []T {
	if slice == nil {
		return nil
	}

	var cleanedSlice []T
	for _, element := range slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	return cleanedSlice
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element T) T {
	if !IsStagedPointerToGongstruct(stage, element) {
		var zero T
		return zero
	}
	return element
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by Animate
func (animate *Animate) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Circle
func (circle *Circle) GongClean(stage *Stage) {
	// insertion point per field
	circle.Animations = GongCleanSlice(stage, circle.Animations)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Condition
func (condition *Condition) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by ControlPoint
func (controlpoint *ControlPoint) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	controlpoint.ClosestRect = GongCleanPointer(stage, controlpoint.ClosestRect)
}

// Clean garbage collect unstaged instances that are referenced by Ellipse
func (ellipse *Ellipse) GongClean(stage *Stage) {
	// insertion point per field
	ellipse.Animates = GongCleanSlice(stage, ellipse.Animates)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Layer
func (layer *Layer) GongClean(stage *Stage) {
	// insertion point per field
	layer.Rects = GongCleanSlice(stage, layer.Rects)
	layer.Texts = GongCleanSlice(stage, layer.Texts)
	layer.Circles = GongCleanSlice(stage, layer.Circles)
	layer.Lines = GongCleanSlice(stage, layer.Lines)
	layer.Ellipses = GongCleanSlice(stage, layer.Ellipses)
	layer.Polylines = GongCleanSlice(stage, layer.Polylines)
	layer.Polygones = GongCleanSlice(stage, layer.Polygones)
	layer.Paths = GongCleanSlice(stage, layer.Paths)
	layer.Links = GongCleanSlice(stage, layer.Links)
	layer.RectLinkLinks = GongCleanSlice(stage, layer.RectLinkLinks)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Line
func (line *Line) GongClean(stage *Stage) {
	// insertion point per field
	line.Animates = GongCleanSlice(stage, line.Animates)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Link
func (link *Link) GongClean(stage *Stage) {
	// insertion point per field
	link.TextAtArrowStart = GongCleanSlice(stage, link.TextAtArrowStart)
	link.TextAtArrowEnd = GongCleanSlice(stage, link.TextAtArrowEnd)
	link.ControlPoints = GongCleanSlice(stage, link.ControlPoints)
	// insertion point per field
	link.Start = GongCleanPointer(stage, link.Start)
	link.End = GongCleanPointer(stage, link.End)
}

// Clean garbage collect unstaged instances that are referenced by LinkAnchoredText
func (linkanchoredtext *LinkAnchoredText) GongClean(stage *Stage) {
	// insertion point per field
	linkanchoredtext.Animates = GongCleanSlice(stage, linkanchoredtext.Animates)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Path
func (path *Path) GongClean(stage *Stage) {
	// insertion point per field
	path.Animates = GongCleanSlice(stage, path.Animates)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Point
func (point *Point) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Polygone
func (polygone *Polygone) GongClean(stage *Stage) {
	// insertion point per field
	polygone.Animates = GongCleanSlice(stage, polygone.Animates)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Polyline
func (polyline *Polyline) GongClean(stage *Stage) {
	// insertion point per field
	polyline.Animates = GongCleanSlice(stage, polyline.Animates)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Rect
func (rect *Rect) GongClean(stage *Stage) {
	// insertion point per field
	rect.HoveringTrigger = GongCleanSlice(stage, rect.HoveringTrigger)
	rect.DisplayConditions = GongCleanSlice(stage, rect.DisplayConditions)
	rect.Animations = GongCleanSlice(stage, rect.Animations)
	rect.RectAnchoredTexts = GongCleanSlice(stage, rect.RectAnchoredTexts)
	rect.RectAnchoredRects = GongCleanSlice(stage, rect.RectAnchoredRects)
	rect.RectAnchoredPaths = GongCleanSlice(stage, rect.RectAnchoredPaths)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by RectAnchoredPath
func (rectanchoredpath *RectAnchoredPath) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by RectAnchoredRect
func (rectanchoredrect *RectAnchoredRect) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by RectAnchoredText
func (rectanchoredtext *RectAnchoredText) GongClean(stage *Stage) {
	// insertion point per field
	rectanchoredtext.Animates = GongCleanSlice(stage, rectanchoredtext.Animates)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by RectLinkLink
func (rectlinklink *RectLinkLink) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	rectlinklink.Start = GongCleanPointer(stage, rectlinklink.Start)
	rectlinklink.End = GongCleanPointer(stage, rectlinklink.End)
}

// Clean garbage collect unstaged instances that are referenced by SVG
func (svg *SVG) GongClean(stage *Stage) {
	// insertion point per field
	svg.Layers = GongCleanSlice(stage, svg.Layers)
	// insertion point per field
	svg.StartRect = GongCleanPointer(stage, svg.StartRect)
	svg.EndRect = GongCleanPointer(stage, svg.EndRect)
}

// Clean garbage collect unstaged instances that are referenced by SvgText
func (svgtext *SvgText) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Text
func (text *Text) GongClean(stage *Stage) {
	// insertion point per field
	text.Animates = GongCleanSlice(stage, text.Animates)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() {
	for _, instance := range stage.GetInstances() {
		instance.GongClean(stage)
	}
}
