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

// Clean computes the reverse map, for all intances, for all clean to pointers field
func (stage *Stage) Clean() {
	// insertion point per named struct
	// clean up Animate
	for animate := range stage.Animates {
		_ = animate
		// insertion point per field
		// insertion point per field
	}

	// clean up Circle
	for circle := range stage.Circles {
		_ = circle
		// insertion point per field
		circle.Animations = GongCleanSlice(stage, circle.Animations)
		// insertion point per field
	}

	// clean up Condition
	for condition := range stage.Conditions {
		_ = condition
		// insertion point per field
		// insertion point per field
	}

	// clean up ControlPoint
	for controlpoint := range stage.ControlPoints {
		_ = controlpoint
		// insertion point per field
		// insertion point per field
		controlpoint.ClosestRect = GongCleanPointer(stage, controlpoint.ClosestRect)
	}

	// clean up Ellipse
	for ellipse := range stage.Ellipses {
		_ = ellipse
		// insertion point per field
		ellipse.Animates = GongCleanSlice(stage, ellipse.Animates)
		// insertion point per field
	}

	// clean up Layer
	for layer := range stage.Layers {
		_ = layer
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

	// clean up Line
	for line := range stage.Lines {
		_ = line
		// insertion point per field
		line.Animates = GongCleanSlice(stage, line.Animates)
		// insertion point per field
	}

	// clean up Link
	for link := range stage.Links {
		_ = link
		// insertion point per field
		link.TextAtArrowStart = GongCleanSlice(stage, link.TextAtArrowStart)
		link.TextAtArrowEnd = GongCleanSlice(stage, link.TextAtArrowEnd)
		link.ControlPoints = GongCleanSlice(stage, link.ControlPoints)
		// insertion point per field
		link.Start = GongCleanPointer(stage, link.Start)
		link.End = GongCleanPointer(stage, link.End)
	}

	// clean up LinkAnchoredText
	for linkanchoredtext := range stage.LinkAnchoredTexts {
		_ = linkanchoredtext
		// insertion point per field
		linkanchoredtext.Animates = GongCleanSlice(stage, linkanchoredtext.Animates)
		// insertion point per field
	}

	// clean up Path
	for path := range stage.Paths {
		_ = path
		// insertion point per field
		path.Animates = GongCleanSlice(stage, path.Animates)
		// insertion point per field
	}

	// clean up Point
	for point := range stage.Points {
		_ = point
		// insertion point per field
		// insertion point per field
	}

	// clean up Polygone
	for polygone := range stage.Polygones {
		_ = polygone
		// insertion point per field
		polygone.Animates = GongCleanSlice(stage, polygone.Animates)
		// insertion point per field
	}

	// clean up Polyline
	for polyline := range stage.Polylines {
		_ = polyline
		// insertion point per field
		polyline.Animates = GongCleanSlice(stage, polyline.Animates)
		// insertion point per field
	}

	// clean up Rect
	for rect := range stage.Rects {
		_ = rect
		// insertion point per field
		rect.HoveringTrigger = GongCleanSlice(stage, rect.HoveringTrigger)
		rect.DisplayConditions = GongCleanSlice(stage, rect.DisplayConditions)
		rect.Animations = GongCleanSlice(stage, rect.Animations)
		rect.RectAnchoredTexts = GongCleanSlice(stage, rect.RectAnchoredTexts)
		rect.RectAnchoredRects = GongCleanSlice(stage, rect.RectAnchoredRects)
		rect.RectAnchoredPaths = GongCleanSlice(stage, rect.RectAnchoredPaths)
		// insertion point per field
	}

	// clean up RectAnchoredPath
	for rectanchoredpath := range stage.RectAnchoredPaths {
		_ = rectanchoredpath
		// insertion point per field
		// insertion point per field
	}

	// clean up RectAnchoredRect
	for rectanchoredrect := range stage.RectAnchoredRects {
		_ = rectanchoredrect
		// insertion point per field
		// insertion point per field
	}

	// clean up RectAnchoredText
	for rectanchoredtext := range stage.RectAnchoredTexts {
		_ = rectanchoredtext
		// insertion point per field
		rectanchoredtext.Animates = GongCleanSlice(stage, rectanchoredtext.Animates)
		// insertion point per field
	}

	// clean up RectLinkLink
	for rectlinklink := range stage.RectLinkLinks {
		_ = rectlinklink
		// insertion point per field
		// insertion point per field
		rectlinklink.Start = GongCleanPointer(stage, rectlinklink.Start)
		rectlinklink.End = GongCleanPointer(stage, rectlinklink.End)
	}

	// clean up SVG
	for svg := range stage.SVGs {
		_ = svg
		// insertion point per field
		svg.Layers = GongCleanSlice(stage, svg.Layers)
		// insertion point per field
		svg.StartRect = GongCleanPointer(stage, svg.StartRect)
		svg.EndRect = GongCleanPointer(stage, svg.EndRect)
	}

	// clean up SvgText
	for svgtext := range stage.SvgTexts {
		_ = svgtext
		// insertion point per field
		// insertion point per field
	}

	// clean up Text
	for text := range stage.Texts {
		_ = text
		// insertion point per field
		text.Animates = GongCleanSlice(stage, text.Animates)
		// insertion point per field
	}

}
