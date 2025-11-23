// generated code - do not edit
package models

// Clean computes the reverse map, for all intances, for all clean to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) Clean() {
	// insertion point per named struct
	// Compute reverse map for named struct Animate
	for animate := range stage.Animates {
		_ = animate
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Circle
	for circle := range stage.Circles {
		_ = circle
		// insertion point per field
		var _Animations []*Animate
		for _, _animate := range circle.Animations {
			if IsStaged(stage, _animate) {
			 	_Animations = append(_Animations, _animate)
			}
		}
		circle.Animations = _Animations
		// insertion point per field
	}

	// Compute reverse map for named struct Condition
	for condition := range stage.Conditions {
		_ = condition
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct ControlPoint
	for controlpoint := range stage.ControlPoints {
		_ = controlpoint
		// insertion point per field
		// insertion point per field
		if !IsStaged(stage, controlpoint.ClosestRect) {
			controlpoint.ClosestRect = nil
		}
	}

	// Compute reverse map for named struct Ellipse
	for ellipse := range stage.Ellipses {
		_ = ellipse
		// insertion point per field
		var _Animates []*Animate
		for _, _animate := range ellipse.Animates {
			if IsStaged(stage, _animate) {
			 	_Animates = append(_Animates, _animate)
			}
		}
		ellipse.Animates = _Animates
		// insertion point per field
	}

	// Compute reverse map for named struct Layer
	for layer := range stage.Layers {
		_ = layer
		// insertion point per field
		var _Rects []*Rect
		for _, _rect := range layer.Rects {
			if IsStaged(stage, _rect) {
			 	_Rects = append(_Rects, _rect)
			}
		}
		layer.Rects = _Rects
		var _Texts []*Text
		for _, _text := range layer.Texts {
			if IsStaged(stage, _text) {
			 	_Texts = append(_Texts, _text)
			}
		}
		layer.Texts = _Texts
		var _Circles []*Circle
		for _, _circle := range layer.Circles {
			if IsStaged(stage, _circle) {
			 	_Circles = append(_Circles, _circle)
			}
		}
		layer.Circles = _Circles
		var _Lines []*Line
		for _, _line := range layer.Lines {
			if IsStaged(stage, _line) {
			 	_Lines = append(_Lines, _line)
			}
		}
		layer.Lines = _Lines
		var _Ellipses []*Ellipse
		for _, _ellipse := range layer.Ellipses {
			if IsStaged(stage, _ellipse) {
			 	_Ellipses = append(_Ellipses, _ellipse)
			}
		}
		layer.Ellipses = _Ellipses
		var _Polylines []*Polyline
		for _, _polyline := range layer.Polylines {
			if IsStaged(stage, _polyline) {
			 	_Polylines = append(_Polylines, _polyline)
			}
		}
		layer.Polylines = _Polylines
		var _Polygones []*Polygone
		for _, _polygone := range layer.Polygones {
			if IsStaged(stage, _polygone) {
			 	_Polygones = append(_Polygones, _polygone)
			}
		}
		layer.Polygones = _Polygones
		var _Paths []*Path
		for _, _path := range layer.Paths {
			if IsStaged(stage, _path) {
			 	_Paths = append(_Paths, _path)
			}
		}
		layer.Paths = _Paths
		var _Links []*Link
		for _, _link := range layer.Links {
			if IsStaged(stage, _link) {
			 	_Links = append(_Links, _link)
			}
		}
		layer.Links = _Links
		var _RectLinkLinks []*RectLinkLink
		for _, _rectlinklink := range layer.RectLinkLinks {
			if IsStaged(stage, _rectlinklink) {
			 	_RectLinkLinks = append(_RectLinkLinks, _rectlinklink)
			}
		}
		layer.RectLinkLinks = _RectLinkLinks
		// insertion point per field
	}

	// Compute reverse map for named struct Line
	for line := range stage.Lines {
		_ = line
		// insertion point per field
		var _Animates []*Animate
		for _, _animate := range line.Animates {
			if IsStaged(stage, _animate) {
			 	_Animates = append(_Animates, _animate)
			}
		}
		line.Animates = _Animates
		// insertion point per field
	}

	// Compute reverse map for named struct Link
	for link := range stage.Links {
		_ = link
		// insertion point per field
		var _TextAtArrowStart []*LinkAnchoredText
		for _, _linkanchoredtext := range link.TextAtArrowStart {
			if IsStaged(stage, _linkanchoredtext) {
			 	_TextAtArrowStart = append(_TextAtArrowStart, _linkanchoredtext)
			}
		}
		link.TextAtArrowStart = _TextAtArrowStart
		var _TextAtArrowEnd []*LinkAnchoredText
		for _, _linkanchoredtext := range link.TextAtArrowEnd {
			if IsStaged(stage, _linkanchoredtext) {
			 	_TextAtArrowEnd = append(_TextAtArrowEnd, _linkanchoredtext)
			}
		}
		link.TextAtArrowEnd = _TextAtArrowEnd
		var _ControlPoints []*ControlPoint
		for _, _controlpoint := range link.ControlPoints {
			if IsStaged(stage, _controlpoint) {
			 	_ControlPoints = append(_ControlPoints, _controlpoint)
			}
		}
		link.ControlPoints = _ControlPoints
		// insertion point per field
		if !IsStaged(stage, link.Start) {
			link.Start = nil
		}
		if !IsStaged(stage, link.End) {
			link.End = nil
		}
	}

	// Compute reverse map for named struct LinkAnchoredText
	for linkanchoredtext := range stage.LinkAnchoredTexts {
		_ = linkanchoredtext
		// insertion point per field
		var _Animates []*Animate
		for _, _animate := range linkanchoredtext.Animates {
			if IsStaged(stage, _animate) {
			 	_Animates = append(_Animates, _animate)
			}
		}
		linkanchoredtext.Animates = _Animates
		// insertion point per field
	}

	// Compute reverse map for named struct Path
	for path := range stage.Paths {
		_ = path
		// insertion point per field
		var _Animates []*Animate
		for _, _animate := range path.Animates {
			if IsStaged(stage, _animate) {
			 	_Animates = append(_Animates, _animate)
			}
		}
		path.Animates = _Animates
		// insertion point per field
	}

	// Compute reverse map for named struct Point
	for point := range stage.Points {
		_ = point
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Polygone
	for polygone := range stage.Polygones {
		_ = polygone
		// insertion point per field
		var _Animates []*Animate
		for _, _animate := range polygone.Animates {
			if IsStaged(stage, _animate) {
			 	_Animates = append(_Animates, _animate)
			}
		}
		polygone.Animates = _Animates
		// insertion point per field
	}

	// Compute reverse map for named struct Polyline
	for polyline := range stage.Polylines {
		_ = polyline
		// insertion point per field
		var _Animates []*Animate
		for _, _animate := range polyline.Animates {
			if IsStaged(stage, _animate) {
			 	_Animates = append(_Animates, _animate)
			}
		}
		polyline.Animates = _Animates
		// insertion point per field
	}

	// Compute reverse map for named struct Rect
	for rect := range stage.Rects {
		_ = rect
		// insertion point per field
		var _HoveringTrigger []*Condition
		for _, _condition := range rect.HoveringTrigger {
			if IsStaged(stage, _condition) {
			 	_HoveringTrigger = append(_HoveringTrigger, _condition)
			}
		}
		rect.HoveringTrigger = _HoveringTrigger
		var _DisplayConditions []*Condition
		for _, _condition := range rect.DisplayConditions {
			if IsStaged(stage, _condition) {
			 	_DisplayConditions = append(_DisplayConditions, _condition)
			}
		}
		rect.DisplayConditions = _DisplayConditions
		var _Animations []*Animate
		for _, _animate := range rect.Animations {
			if IsStaged(stage, _animate) {
			 	_Animations = append(_Animations, _animate)
			}
		}
		rect.Animations = _Animations
		var _RectAnchoredTexts []*RectAnchoredText
		for _, _rectanchoredtext := range rect.RectAnchoredTexts {
			if IsStaged(stage, _rectanchoredtext) {
			 	_RectAnchoredTexts = append(_RectAnchoredTexts, _rectanchoredtext)
			}
		}
		rect.RectAnchoredTexts = _RectAnchoredTexts
		var _RectAnchoredRects []*RectAnchoredRect
		for _, _rectanchoredrect := range rect.RectAnchoredRects {
			if IsStaged(stage, _rectanchoredrect) {
			 	_RectAnchoredRects = append(_RectAnchoredRects, _rectanchoredrect)
			}
		}
		rect.RectAnchoredRects = _RectAnchoredRects
		var _RectAnchoredPaths []*RectAnchoredPath
		for _, _rectanchoredpath := range rect.RectAnchoredPaths {
			if IsStaged(stage, _rectanchoredpath) {
			 	_RectAnchoredPaths = append(_RectAnchoredPaths, _rectanchoredpath)
			}
		}
		rect.RectAnchoredPaths = _RectAnchoredPaths
		// insertion point per field
	}

	// Compute reverse map for named struct RectAnchoredPath
	for rectanchoredpath := range stage.RectAnchoredPaths {
		_ = rectanchoredpath
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct RectAnchoredRect
	for rectanchoredrect := range stage.RectAnchoredRects {
		_ = rectanchoredrect
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct RectAnchoredText
	for rectanchoredtext := range stage.RectAnchoredTexts {
		_ = rectanchoredtext
		// insertion point per field
		var _Animates []*Animate
		for _, _animate := range rectanchoredtext.Animates {
			if IsStaged(stage, _animate) {
			 	_Animates = append(_Animates, _animate)
			}
		}
		rectanchoredtext.Animates = _Animates
		// insertion point per field
	}

	// Compute reverse map for named struct RectLinkLink
	for rectlinklink := range stage.RectLinkLinks {
		_ = rectlinklink
		// insertion point per field
		// insertion point per field
		if !IsStaged(stage, rectlinklink.Start) {
			rectlinklink.Start = nil
		}
		if !IsStaged(stage, rectlinklink.End) {
			rectlinklink.End = nil
		}
	}

	// Compute reverse map for named struct SVG
	for svg := range stage.SVGs {
		_ = svg
		// insertion point per field
		var _Layers []*Layer
		for _, _layer := range svg.Layers {
			if IsStaged(stage, _layer) {
			 	_Layers = append(_Layers, _layer)
			}
		}
		svg.Layers = _Layers
		// insertion point per field
		if !IsStaged(stage, svg.StartRect) {
			svg.StartRect = nil
		}
		if !IsStaged(stage, svg.EndRect) {
			svg.EndRect = nil
		}
	}

	// Compute reverse map for named struct SvgText
	for svgtext := range stage.SvgTexts {
		_ = svgtext
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Text
	for text := range stage.Texts {
		_ = text
		// insertion point per field
		var _Animates []*Animate
		for _, _animate := range text.Animates {
			if IsStaged(stage, _animate) {
			 	_Animates = append(_Animates, _animate)
			}
		}
		text.Animates = _Animates
		// insertion point per field
	}

}
