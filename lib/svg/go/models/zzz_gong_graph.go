// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Animate:
		ok = stage.IsStagedAnimate(target)

	case *Circle:
		ok = stage.IsStagedCircle(target)

	case *Condition:
		ok = stage.IsStagedCondition(target)

	case *ControlPoint:
		ok = stage.IsStagedControlPoint(target)

	case *Ellipse:
		ok = stage.IsStagedEllipse(target)

	case *Layer:
		ok = stage.IsStagedLayer(target)

	case *Line:
		ok = stage.IsStagedLine(target)

	case *Link:
		ok = stage.IsStagedLink(target)

	case *LinkAnchoredText:
		ok = stage.IsStagedLinkAnchoredText(target)

	case *Path:
		ok = stage.IsStagedPath(target)

	case *Point:
		ok = stage.IsStagedPoint(target)

	case *Polygone:
		ok = stage.IsStagedPolygone(target)

	case *Polyline:
		ok = stage.IsStagedPolyline(target)

	case *Rect:
		ok = stage.IsStagedRect(target)

	case *RectAnchoredPath:
		ok = stage.IsStagedRectAnchoredPath(target)

	case *RectAnchoredRect:
		ok = stage.IsStagedRectAnchoredRect(target)

	case *RectAnchoredText:
		ok = stage.IsStagedRectAnchoredText(target)

	case *RectLinkLink:
		ok = stage.IsStagedRectLinkLink(target)

	case *SVG:
		ok = stage.IsStagedSVG(target)

	case *SvgText:
		ok = stage.IsStagedSvgText(target)

	case *Text:
		ok = stage.IsStagedText(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Animate:
		ok = stage.IsStagedAnimate(target)

	case *Circle:
		ok = stage.IsStagedCircle(target)

	case *Condition:
		ok = stage.IsStagedCondition(target)

	case *ControlPoint:
		ok = stage.IsStagedControlPoint(target)

	case *Ellipse:
		ok = stage.IsStagedEllipse(target)

	case *Layer:
		ok = stage.IsStagedLayer(target)

	case *Line:
		ok = stage.IsStagedLine(target)

	case *Link:
		ok = stage.IsStagedLink(target)

	case *LinkAnchoredText:
		ok = stage.IsStagedLinkAnchoredText(target)

	case *Path:
		ok = stage.IsStagedPath(target)

	case *Point:
		ok = stage.IsStagedPoint(target)

	case *Polygone:
		ok = stage.IsStagedPolygone(target)

	case *Polyline:
		ok = stage.IsStagedPolyline(target)

	case *Rect:
		ok = stage.IsStagedRect(target)

	case *RectAnchoredPath:
		ok = stage.IsStagedRectAnchoredPath(target)

	case *RectAnchoredRect:
		ok = stage.IsStagedRectAnchoredRect(target)

	case *RectAnchoredText:
		ok = stage.IsStagedRectAnchoredText(target)

	case *RectLinkLink:
		ok = stage.IsStagedRectLinkLink(target)

	case *SVG:
		ok = stage.IsStagedSVG(target)

	case *SvgText:
		ok = stage.IsStagedSvgText(target)

	case *Text:
		ok = stage.IsStagedText(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedAnimate(animate *Animate) (ok bool) {

	_, ok = stage.Animates[animate]

	return
}

func (stage *Stage) IsStagedCircle(circle *Circle) (ok bool) {

	_, ok = stage.Circles[circle]

	return
}

func (stage *Stage) IsStagedCondition(condition *Condition) (ok bool) {

	_, ok = stage.Conditions[condition]

	return
}

func (stage *Stage) IsStagedControlPoint(controlpoint *ControlPoint) (ok bool) {

	_, ok = stage.ControlPoints[controlpoint]

	return
}

func (stage *Stage) IsStagedEllipse(ellipse *Ellipse) (ok bool) {

	_, ok = stage.Ellipses[ellipse]

	return
}

func (stage *Stage) IsStagedLayer(layer *Layer) (ok bool) {

	_, ok = stage.Layers[layer]

	return
}

func (stage *Stage) IsStagedLine(line *Line) (ok bool) {

	_, ok = stage.Lines[line]

	return
}

func (stage *Stage) IsStagedLink(link *Link) (ok bool) {

	_, ok = stage.Links[link]

	return
}

func (stage *Stage) IsStagedLinkAnchoredText(linkanchoredtext *LinkAnchoredText) (ok bool) {

	_, ok = stage.LinkAnchoredTexts[linkanchoredtext]

	return
}

func (stage *Stage) IsStagedPath(path *Path) (ok bool) {

	_, ok = stage.Paths[path]

	return
}

func (stage *Stage) IsStagedPoint(point *Point) (ok bool) {

	_, ok = stage.Points[point]

	return
}

func (stage *Stage) IsStagedPolygone(polygone *Polygone) (ok bool) {

	_, ok = stage.Polygones[polygone]

	return
}

func (stage *Stage) IsStagedPolyline(polyline *Polyline) (ok bool) {

	_, ok = stage.Polylines[polyline]

	return
}

func (stage *Stage) IsStagedRect(rect *Rect) (ok bool) {

	_, ok = stage.Rects[rect]

	return
}

func (stage *Stage) IsStagedRectAnchoredPath(rectanchoredpath *RectAnchoredPath) (ok bool) {

	_, ok = stage.RectAnchoredPaths[rectanchoredpath]

	return
}

func (stage *Stage) IsStagedRectAnchoredRect(rectanchoredrect *RectAnchoredRect) (ok bool) {

	_, ok = stage.RectAnchoredRects[rectanchoredrect]

	return
}

func (stage *Stage) IsStagedRectAnchoredText(rectanchoredtext *RectAnchoredText) (ok bool) {

	_, ok = stage.RectAnchoredTexts[rectanchoredtext]

	return
}

func (stage *Stage) IsStagedRectLinkLink(rectlinklink *RectLinkLink) (ok bool) {

	_, ok = stage.RectLinkLinks[rectlinklink]

	return
}

func (stage *Stage) IsStagedSVG(svg *SVG) (ok bool) {

	_, ok = stage.SVGs[svg]

	return
}

func (stage *Stage) IsStagedSvgText(svgtext *SvgText) (ok bool) {

	_, ok = stage.SvgTexts[svgtext]

	return
}

func (stage *Stage) IsStagedText(text *Text) (ok bool) {

	_, ok = stage.Texts[text]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Animate:
		stage.StageBranchAnimate(target)

	case *Circle:
		stage.StageBranchCircle(target)

	case *Condition:
		stage.StageBranchCondition(target)

	case *ControlPoint:
		stage.StageBranchControlPoint(target)

	case *Ellipse:
		stage.StageBranchEllipse(target)

	case *Layer:
		stage.StageBranchLayer(target)

	case *Line:
		stage.StageBranchLine(target)

	case *Link:
		stage.StageBranchLink(target)

	case *LinkAnchoredText:
		stage.StageBranchLinkAnchoredText(target)

	case *Path:
		stage.StageBranchPath(target)

	case *Point:
		stage.StageBranchPoint(target)

	case *Polygone:
		stage.StageBranchPolygone(target)

	case *Polyline:
		stage.StageBranchPolyline(target)

	case *Rect:
		stage.StageBranchRect(target)

	case *RectAnchoredPath:
		stage.StageBranchRectAnchoredPath(target)

	case *RectAnchoredRect:
		stage.StageBranchRectAnchoredRect(target)

	case *RectAnchoredText:
		stage.StageBranchRectAnchoredText(target)

	case *RectLinkLink:
		stage.StageBranchRectLinkLink(target)

	case *SVG:
		stage.StageBranchSVG(target)

	case *SvgText:
		stage.StageBranchSvgText(target)

	case *Text:
		stage.StageBranchText(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchAnimate(animate *Animate) {

	// check if instance is already staged
	if IsStaged(stage, animate) {
		return
	}

	animate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchCircle(circle *Circle) {

	// check if instance is already staged
	if IsStaged(stage, circle) {
		return
	}

	circle.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range circle.Animations {
		StageBranch(stage, _animate)
	}

}

func (stage *Stage) StageBranchCondition(condition *Condition) {

	// check if instance is already staged
	if IsStaged(stage, condition) {
		return
	}

	condition.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchControlPoint(controlpoint *ControlPoint) {

	// check if instance is already staged
	if IsStaged(stage, controlpoint) {
		return
	}

	controlpoint.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if controlpoint.ClosestRect != nil {
		StageBranch(stage, controlpoint.ClosestRect)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchEllipse(ellipse *Ellipse) {

	// check if instance is already staged
	if IsStaged(stage, ellipse) {
		return
	}

	ellipse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range ellipse.Animates {
		StageBranch(stage, _animate)
	}

}

func (stage *Stage) StageBranchLayer(layer *Layer) {

	// check if instance is already staged
	if IsStaged(stage, layer) {
		return
	}

	layer.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rect := range layer.Rects {
		StageBranch(stage, _rect)
	}
	for _, _text := range layer.Texts {
		StageBranch(stage, _text)
	}
	for _, _circle := range layer.Circles {
		StageBranch(stage, _circle)
	}
	for _, _line := range layer.Lines {
		StageBranch(stage, _line)
	}
	for _, _ellipse := range layer.Ellipses {
		StageBranch(stage, _ellipse)
	}
	for _, _polyline := range layer.Polylines {
		StageBranch(stage, _polyline)
	}
	for _, _polygone := range layer.Polygones {
		StageBranch(stage, _polygone)
	}
	for _, _path := range layer.Paths {
		StageBranch(stage, _path)
	}
	for _, _link := range layer.Links {
		StageBranch(stage, _link)
	}
	for _, _rectlinklink := range layer.RectLinkLinks {
		StageBranch(stage, _rectlinklink)
	}

}

func (stage *Stage) StageBranchLine(line *Line) {

	// check if instance is already staged
	if IsStaged(stage, line) {
		return
	}

	line.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range line.Animates {
		StageBranch(stage, _animate)
	}

}

func (stage *Stage) StageBranchLink(link *Link) {

	// check if instance is already staged
	if IsStaged(stage, link) {
		return
	}

	link.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if link.Start != nil {
		StageBranch(stage, link.Start)
	}
	if link.End != nil {
		StageBranch(stage, link.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _linkanchoredtext := range link.TextAtArrowStart {
		StageBranch(stage, _linkanchoredtext)
	}
	for _, _linkanchoredtext := range link.TextAtArrowEnd {
		StageBranch(stage, _linkanchoredtext)
	}
	for _, _controlpoint := range link.ControlPoints {
		StageBranch(stage, _controlpoint)
	}

}

func (stage *Stage) StageBranchLinkAnchoredText(linkanchoredtext *LinkAnchoredText) {

	// check if instance is already staged
	if IsStaged(stage, linkanchoredtext) {
		return
	}

	linkanchoredtext.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range linkanchoredtext.Animates {
		StageBranch(stage, _animate)
	}

}

func (stage *Stage) StageBranchPath(path *Path) {

	// check if instance is already staged
	if IsStaged(stage, path) {
		return
	}

	path.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range path.Animates {
		StageBranch(stage, _animate)
	}

}

func (stage *Stage) StageBranchPoint(point *Point) {

	// check if instance is already staged
	if IsStaged(stage, point) {
		return
	}

	point.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPolygone(polygone *Polygone) {

	// check if instance is already staged
	if IsStaged(stage, polygone) {
		return
	}

	polygone.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range polygone.Animates {
		StageBranch(stage, _animate)
	}

}

func (stage *Stage) StageBranchPolyline(polyline *Polyline) {

	// check if instance is already staged
	if IsStaged(stage, polyline) {
		return
	}

	polyline.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range polyline.Animates {
		StageBranch(stage, _animate)
	}

}

func (stage *Stage) StageBranchRect(rect *Rect) {

	// check if instance is already staged
	if IsStaged(stage, rect) {
		return
	}

	rect.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _condition := range rect.HoveringTrigger {
		StageBranch(stage, _condition)
	}
	for _, _condition := range rect.DisplayConditions {
		StageBranch(stage, _condition)
	}
	for _, _animate := range rect.Animations {
		StageBranch(stage, _animate)
	}
	for _, _rectanchoredtext := range rect.RectAnchoredTexts {
		StageBranch(stage, _rectanchoredtext)
	}
	for _, _rectanchoredrect := range rect.RectAnchoredRects {
		StageBranch(stage, _rectanchoredrect)
	}
	for _, _rectanchoredpath := range rect.RectAnchoredPaths {
		StageBranch(stage, _rectanchoredpath)
	}

}

func (stage *Stage) StageBranchRectAnchoredPath(rectanchoredpath *RectAnchoredPath) {

	// check if instance is already staged
	if IsStaged(stage, rectanchoredpath) {
		return
	}

	rectanchoredpath.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchRectAnchoredRect(rectanchoredrect *RectAnchoredRect) {

	// check if instance is already staged
	if IsStaged(stage, rectanchoredrect) {
		return
	}

	rectanchoredrect.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchRectAnchoredText(rectanchoredtext *RectAnchoredText) {

	// check if instance is already staged
	if IsStaged(stage, rectanchoredtext) {
		return
	}

	rectanchoredtext.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range rectanchoredtext.Animates {
		StageBranch(stage, _animate)
	}

}

func (stage *Stage) StageBranchRectLinkLink(rectlinklink *RectLinkLink) {

	// check if instance is already staged
	if IsStaged(stage, rectlinklink) {
		return
	}

	rectlinklink.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if rectlinklink.Start != nil {
		StageBranch(stage, rectlinklink.Start)
	}
	if rectlinklink.End != nil {
		StageBranch(stage, rectlinklink.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSVG(svg *SVG) {

	// check if instance is already staged
	if IsStaged(stage, svg) {
		return
	}

	svg.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if svg.StartRect != nil {
		StageBranch(stage, svg.StartRect)
	}
	if svg.EndRect != nil {
		StageBranch(stage, svg.EndRect)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _layer := range svg.Layers {
		StageBranch(stage, _layer)
	}

}

func (stage *Stage) StageBranchSvgText(svgtext *SvgText) {

	// check if instance is already staged
	if IsStaged(stage, svgtext) {
		return
	}

	svgtext.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchText(text *Text) {

	// check if instance is already staged
	if IsStaged(stage, text) {
		return
	}

	text.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range text.Animates {
		StageBranch(stage, _animate)
	}

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Animate:
		toT := CopyBranchAnimate(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Circle:
		toT := CopyBranchCircle(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Condition:
		toT := CopyBranchCondition(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ControlPoint:
		toT := CopyBranchControlPoint(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Ellipse:
		toT := CopyBranchEllipse(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Layer:
		toT := CopyBranchLayer(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Line:
		toT := CopyBranchLine(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Link:
		toT := CopyBranchLink(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LinkAnchoredText:
		toT := CopyBranchLinkAnchoredText(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Path:
		toT := CopyBranchPath(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Point:
		toT := CopyBranchPoint(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Polygone:
		toT := CopyBranchPolygone(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Polyline:
		toT := CopyBranchPolyline(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Rect:
		toT := CopyBranchRect(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RectAnchoredPath:
		toT := CopyBranchRectAnchoredPath(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RectAnchoredRect:
		toT := CopyBranchRectAnchoredRect(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RectAnchoredText:
		toT := CopyBranchRectAnchoredText(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RectLinkLink:
		toT := CopyBranchRectLinkLink(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SVG:
		toT := CopyBranchSVG(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SvgText:
		toT := CopyBranchSvgText(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Text:
		toT := CopyBranchText(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchAnimate(mapOrigCopy map[any]any, animateFrom *Animate) (animateTo *Animate) {

	// animateFrom has already been copied
	if _animateTo, ok := mapOrigCopy[animateFrom]; ok {
		animateTo = _animateTo.(*Animate)
		return
	}

	animateTo = new(Animate)
	mapOrigCopy[animateFrom] = animateTo
	animateFrom.CopyBasicFields(animateTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCircle(mapOrigCopy map[any]any, circleFrom *Circle) (circleTo *Circle) {

	// circleFrom has already been copied
	if _circleTo, ok := mapOrigCopy[circleFrom]; ok {
		circleTo = _circleTo.(*Circle)
		return
	}

	circleTo = new(Circle)
	mapOrigCopy[circleFrom] = circleTo
	circleFrom.CopyBasicFields(circleTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range circleFrom.Animations {
		circleTo.Animations = append(circleTo.Animations, CopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func CopyBranchCondition(mapOrigCopy map[any]any, conditionFrom *Condition) (conditionTo *Condition) {

	// conditionFrom has already been copied
	if _conditionTo, ok := mapOrigCopy[conditionFrom]; ok {
		conditionTo = _conditionTo.(*Condition)
		return
	}

	conditionTo = new(Condition)
	mapOrigCopy[conditionFrom] = conditionTo
	conditionFrom.CopyBasicFields(conditionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchControlPoint(mapOrigCopy map[any]any, controlpointFrom *ControlPoint) (controlpointTo *ControlPoint) {

	// controlpointFrom has already been copied
	if _controlpointTo, ok := mapOrigCopy[controlpointFrom]; ok {
		controlpointTo = _controlpointTo.(*ControlPoint)
		return
	}

	controlpointTo = new(ControlPoint)
	mapOrigCopy[controlpointFrom] = controlpointTo
	controlpointFrom.CopyBasicFields(controlpointTo)

	//insertion point for the staging of instances referenced by pointers
	if controlpointFrom.ClosestRect != nil {
		controlpointTo.ClosestRect = CopyBranchRect(mapOrigCopy, controlpointFrom.ClosestRect)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEllipse(mapOrigCopy map[any]any, ellipseFrom *Ellipse) (ellipseTo *Ellipse) {

	// ellipseFrom has already been copied
	if _ellipseTo, ok := mapOrigCopy[ellipseFrom]; ok {
		ellipseTo = _ellipseTo.(*Ellipse)
		return
	}

	ellipseTo = new(Ellipse)
	mapOrigCopy[ellipseFrom] = ellipseTo
	ellipseFrom.CopyBasicFields(ellipseTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range ellipseFrom.Animates {
		ellipseTo.Animates = append(ellipseTo.Animates, CopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func CopyBranchLayer(mapOrigCopy map[any]any, layerFrom *Layer) (layerTo *Layer) {

	// layerFrom has already been copied
	if _layerTo, ok := mapOrigCopy[layerFrom]; ok {
		layerTo = _layerTo.(*Layer)
		return
	}

	layerTo = new(Layer)
	mapOrigCopy[layerFrom] = layerTo
	layerFrom.CopyBasicFields(layerTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rect := range layerFrom.Rects {
		layerTo.Rects = append(layerTo.Rects, CopyBranchRect(mapOrigCopy, _rect))
	}
	for _, _text := range layerFrom.Texts {
		layerTo.Texts = append(layerTo.Texts, CopyBranchText(mapOrigCopy, _text))
	}
	for _, _circle := range layerFrom.Circles {
		layerTo.Circles = append(layerTo.Circles, CopyBranchCircle(mapOrigCopy, _circle))
	}
	for _, _line := range layerFrom.Lines {
		layerTo.Lines = append(layerTo.Lines, CopyBranchLine(mapOrigCopy, _line))
	}
	for _, _ellipse := range layerFrom.Ellipses {
		layerTo.Ellipses = append(layerTo.Ellipses, CopyBranchEllipse(mapOrigCopy, _ellipse))
	}
	for _, _polyline := range layerFrom.Polylines {
		layerTo.Polylines = append(layerTo.Polylines, CopyBranchPolyline(mapOrigCopy, _polyline))
	}
	for _, _polygone := range layerFrom.Polygones {
		layerTo.Polygones = append(layerTo.Polygones, CopyBranchPolygone(mapOrigCopy, _polygone))
	}
	for _, _path := range layerFrom.Paths {
		layerTo.Paths = append(layerTo.Paths, CopyBranchPath(mapOrigCopy, _path))
	}
	for _, _link := range layerFrom.Links {
		layerTo.Links = append(layerTo.Links, CopyBranchLink(mapOrigCopy, _link))
	}
	for _, _rectlinklink := range layerFrom.RectLinkLinks {
		layerTo.RectLinkLinks = append(layerTo.RectLinkLinks, CopyBranchRectLinkLink(mapOrigCopy, _rectlinklink))
	}

	return
}

func CopyBranchLine(mapOrigCopy map[any]any, lineFrom *Line) (lineTo *Line) {

	// lineFrom has already been copied
	if _lineTo, ok := mapOrigCopy[lineFrom]; ok {
		lineTo = _lineTo.(*Line)
		return
	}

	lineTo = new(Line)
	mapOrigCopy[lineFrom] = lineTo
	lineFrom.CopyBasicFields(lineTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range lineFrom.Animates {
		lineTo.Animates = append(lineTo.Animates, CopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func CopyBranchLink(mapOrigCopy map[any]any, linkFrom *Link) (linkTo *Link) {

	// linkFrom has already been copied
	if _linkTo, ok := mapOrigCopy[linkFrom]; ok {
		linkTo = _linkTo.(*Link)
		return
	}

	linkTo = new(Link)
	mapOrigCopy[linkFrom] = linkTo
	linkFrom.CopyBasicFields(linkTo)

	//insertion point for the staging of instances referenced by pointers
	if linkFrom.Start != nil {
		linkTo.Start = CopyBranchRect(mapOrigCopy, linkFrom.Start)
	}
	if linkFrom.End != nil {
		linkTo.End = CopyBranchRect(mapOrigCopy, linkFrom.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _linkanchoredtext := range linkFrom.TextAtArrowStart {
		linkTo.TextAtArrowStart = append(linkTo.TextAtArrowStart, CopyBranchLinkAnchoredText(mapOrigCopy, _linkanchoredtext))
	}
	for _, _linkanchoredtext := range linkFrom.TextAtArrowEnd {
		linkTo.TextAtArrowEnd = append(linkTo.TextAtArrowEnd, CopyBranchLinkAnchoredText(mapOrigCopy, _linkanchoredtext))
	}
	for _, _controlpoint := range linkFrom.ControlPoints {
		linkTo.ControlPoints = append(linkTo.ControlPoints, CopyBranchControlPoint(mapOrigCopy, _controlpoint))
	}

	return
}

func CopyBranchLinkAnchoredText(mapOrigCopy map[any]any, linkanchoredtextFrom *LinkAnchoredText) (linkanchoredtextTo *LinkAnchoredText) {

	// linkanchoredtextFrom has already been copied
	if _linkanchoredtextTo, ok := mapOrigCopy[linkanchoredtextFrom]; ok {
		linkanchoredtextTo = _linkanchoredtextTo.(*LinkAnchoredText)
		return
	}

	linkanchoredtextTo = new(LinkAnchoredText)
	mapOrigCopy[linkanchoredtextFrom] = linkanchoredtextTo
	linkanchoredtextFrom.CopyBasicFields(linkanchoredtextTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range linkanchoredtextFrom.Animates {
		linkanchoredtextTo.Animates = append(linkanchoredtextTo.Animates, CopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func CopyBranchPath(mapOrigCopy map[any]any, pathFrom *Path) (pathTo *Path) {

	// pathFrom has already been copied
	if _pathTo, ok := mapOrigCopy[pathFrom]; ok {
		pathTo = _pathTo.(*Path)
		return
	}

	pathTo = new(Path)
	mapOrigCopy[pathFrom] = pathTo
	pathFrom.CopyBasicFields(pathTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range pathFrom.Animates {
		pathTo.Animates = append(pathTo.Animates, CopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func CopyBranchPoint(mapOrigCopy map[any]any, pointFrom *Point) (pointTo *Point) {

	// pointFrom has already been copied
	if _pointTo, ok := mapOrigCopy[pointFrom]; ok {
		pointTo = _pointTo.(*Point)
		return
	}

	pointTo = new(Point)
	mapOrigCopy[pointFrom] = pointTo
	pointFrom.CopyBasicFields(pointTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPolygone(mapOrigCopy map[any]any, polygoneFrom *Polygone) (polygoneTo *Polygone) {

	// polygoneFrom has already been copied
	if _polygoneTo, ok := mapOrigCopy[polygoneFrom]; ok {
		polygoneTo = _polygoneTo.(*Polygone)
		return
	}

	polygoneTo = new(Polygone)
	mapOrigCopy[polygoneFrom] = polygoneTo
	polygoneFrom.CopyBasicFields(polygoneTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range polygoneFrom.Animates {
		polygoneTo.Animates = append(polygoneTo.Animates, CopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func CopyBranchPolyline(mapOrigCopy map[any]any, polylineFrom *Polyline) (polylineTo *Polyline) {

	// polylineFrom has already been copied
	if _polylineTo, ok := mapOrigCopy[polylineFrom]; ok {
		polylineTo = _polylineTo.(*Polyline)
		return
	}

	polylineTo = new(Polyline)
	mapOrigCopy[polylineFrom] = polylineTo
	polylineFrom.CopyBasicFields(polylineTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range polylineFrom.Animates {
		polylineTo.Animates = append(polylineTo.Animates, CopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func CopyBranchRect(mapOrigCopy map[any]any, rectFrom *Rect) (rectTo *Rect) {

	// rectFrom has already been copied
	if _rectTo, ok := mapOrigCopy[rectFrom]; ok {
		rectTo = _rectTo.(*Rect)
		return
	}

	rectTo = new(Rect)
	mapOrigCopy[rectFrom] = rectTo
	rectFrom.CopyBasicFields(rectTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _condition := range rectFrom.HoveringTrigger {
		rectTo.HoveringTrigger = append(rectTo.HoveringTrigger, CopyBranchCondition(mapOrigCopy, _condition))
	}
	for _, _condition := range rectFrom.DisplayConditions {
		rectTo.DisplayConditions = append(rectTo.DisplayConditions, CopyBranchCondition(mapOrigCopy, _condition))
	}
	for _, _animate := range rectFrom.Animations {
		rectTo.Animations = append(rectTo.Animations, CopyBranchAnimate(mapOrigCopy, _animate))
	}
	for _, _rectanchoredtext := range rectFrom.RectAnchoredTexts {
		rectTo.RectAnchoredTexts = append(rectTo.RectAnchoredTexts, CopyBranchRectAnchoredText(mapOrigCopy, _rectanchoredtext))
	}
	for _, _rectanchoredrect := range rectFrom.RectAnchoredRects {
		rectTo.RectAnchoredRects = append(rectTo.RectAnchoredRects, CopyBranchRectAnchoredRect(mapOrigCopy, _rectanchoredrect))
	}
	for _, _rectanchoredpath := range rectFrom.RectAnchoredPaths {
		rectTo.RectAnchoredPaths = append(rectTo.RectAnchoredPaths, CopyBranchRectAnchoredPath(mapOrigCopy, _rectanchoredpath))
	}

	return
}

func CopyBranchRectAnchoredPath(mapOrigCopy map[any]any, rectanchoredpathFrom *RectAnchoredPath) (rectanchoredpathTo *RectAnchoredPath) {

	// rectanchoredpathFrom has already been copied
	if _rectanchoredpathTo, ok := mapOrigCopy[rectanchoredpathFrom]; ok {
		rectanchoredpathTo = _rectanchoredpathTo.(*RectAnchoredPath)
		return
	}

	rectanchoredpathTo = new(RectAnchoredPath)
	mapOrigCopy[rectanchoredpathFrom] = rectanchoredpathTo
	rectanchoredpathFrom.CopyBasicFields(rectanchoredpathTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRectAnchoredRect(mapOrigCopy map[any]any, rectanchoredrectFrom *RectAnchoredRect) (rectanchoredrectTo *RectAnchoredRect) {

	// rectanchoredrectFrom has already been copied
	if _rectanchoredrectTo, ok := mapOrigCopy[rectanchoredrectFrom]; ok {
		rectanchoredrectTo = _rectanchoredrectTo.(*RectAnchoredRect)
		return
	}

	rectanchoredrectTo = new(RectAnchoredRect)
	mapOrigCopy[rectanchoredrectFrom] = rectanchoredrectTo
	rectanchoredrectFrom.CopyBasicFields(rectanchoredrectTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRectAnchoredText(mapOrigCopy map[any]any, rectanchoredtextFrom *RectAnchoredText) (rectanchoredtextTo *RectAnchoredText) {

	// rectanchoredtextFrom has already been copied
	if _rectanchoredtextTo, ok := mapOrigCopy[rectanchoredtextFrom]; ok {
		rectanchoredtextTo = _rectanchoredtextTo.(*RectAnchoredText)
		return
	}

	rectanchoredtextTo = new(RectAnchoredText)
	mapOrigCopy[rectanchoredtextFrom] = rectanchoredtextTo
	rectanchoredtextFrom.CopyBasicFields(rectanchoredtextTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range rectanchoredtextFrom.Animates {
		rectanchoredtextTo.Animates = append(rectanchoredtextTo.Animates, CopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func CopyBranchRectLinkLink(mapOrigCopy map[any]any, rectlinklinkFrom *RectLinkLink) (rectlinklinkTo *RectLinkLink) {

	// rectlinklinkFrom has already been copied
	if _rectlinklinkTo, ok := mapOrigCopy[rectlinklinkFrom]; ok {
		rectlinklinkTo = _rectlinklinkTo.(*RectLinkLink)
		return
	}

	rectlinklinkTo = new(RectLinkLink)
	mapOrigCopy[rectlinklinkFrom] = rectlinklinkTo
	rectlinklinkFrom.CopyBasicFields(rectlinklinkTo)

	//insertion point for the staging of instances referenced by pointers
	if rectlinklinkFrom.Start != nil {
		rectlinklinkTo.Start = CopyBranchRect(mapOrigCopy, rectlinklinkFrom.Start)
	}
	if rectlinklinkFrom.End != nil {
		rectlinklinkTo.End = CopyBranchLink(mapOrigCopy, rectlinklinkFrom.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSVG(mapOrigCopy map[any]any, svgFrom *SVG) (svgTo *SVG) {

	// svgFrom has already been copied
	if _svgTo, ok := mapOrigCopy[svgFrom]; ok {
		svgTo = _svgTo.(*SVG)
		return
	}

	svgTo = new(SVG)
	mapOrigCopy[svgFrom] = svgTo
	svgFrom.CopyBasicFields(svgTo)

	//insertion point for the staging of instances referenced by pointers
	if svgFrom.StartRect != nil {
		svgTo.StartRect = CopyBranchRect(mapOrigCopy, svgFrom.StartRect)
	}
	if svgFrom.EndRect != nil {
		svgTo.EndRect = CopyBranchRect(mapOrigCopy, svgFrom.EndRect)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _layer := range svgFrom.Layers {
		svgTo.Layers = append(svgTo.Layers, CopyBranchLayer(mapOrigCopy, _layer))
	}

	return
}

func CopyBranchSvgText(mapOrigCopy map[any]any, svgtextFrom *SvgText) (svgtextTo *SvgText) {

	// svgtextFrom has already been copied
	if _svgtextTo, ok := mapOrigCopy[svgtextFrom]; ok {
		svgtextTo = _svgtextTo.(*SvgText)
		return
	}

	svgtextTo = new(SvgText)
	mapOrigCopy[svgtextFrom] = svgtextTo
	svgtextFrom.CopyBasicFields(svgtextTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchText(mapOrigCopy map[any]any, textFrom *Text) (textTo *Text) {

	// textFrom has already been copied
	if _textTo, ok := mapOrigCopy[textFrom]; ok {
		textTo = _textTo.(*Text)
		return
	}

	textTo = new(Text)
	mapOrigCopy[textFrom] = textTo
	textFrom.CopyBasicFields(textTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range textFrom.Animates {
		textTo.Animates = append(textTo.Animates, CopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Animate:
		stage.UnstageBranchAnimate(target)

	case *Circle:
		stage.UnstageBranchCircle(target)

	case *Condition:
		stage.UnstageBranchCondition(target)

	case *ControlPoint:
		stage.UnstageBranchControlPoint(target)

	case *Ellipse:
		stage.UnstageBranchEllipse(target)

	case *Layer:
		stage.UnstageBranchLayer(target)

	case *Line:
		stage.UnstageBranchLine(target)

	case *Link:
		stage.UnstageBranchLink(target)

	case *LinkAnchoredText:
		stage.UnstageBranchLinkAnchoredText(target)

	case *Path:
		stage.UnstageBranchPath(target)

	case *Point:
		stage.UnstageBranchPoint(target)

	case *Polygone:
		stage.UnstageBranchPolygone(target)

	case *Polyline:
		stage.UnstageBranchPolyline(target)

	case *Rect:
		stage.UnstageBranchRect(target)

	case *RectAnchoredPath:
		stage.UnstageBranchRectAnchoredPath(target)

	case *RectAnchoredRect:
		stage.UnstageBranchRectAnchoredRect(target)

	case *RectAnchoredText:
		stage.UnstageBranchRectAnchoredText(target)

	case *RectLinkLink:
		stage.UnstageBranchRectLinkLink(target)

	case *SVG:
		stage.UnstageBranchSVG(target)

	case *SvgText:
		stage.UnstageBranchSvgText(target)

	case *Text:
		stage.UnstageBranchText(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchAnimate(animate *Animate) {

	// check if instance is already staged
	if !IsStaged(stage, animate) {
		return
	}

	animate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchCircle(circle *Circle) {

	// check if instance is already staged
	if !IsStaged(stage, circle) {
		return
	}

	circle.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range circle.Animations {
		UnstageBranch(stage, _animate)
	}

}

func (stage *Stage) UnstageBranchCondition(condition *Condition) {

	// check if instance is already staged
	if !IsStaged(stage, condition) {
		return
	}

	condition.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchControlPoint(controlpoint *ControlPoint) {

	// check if instance is already staged
	if !IsStaged(stage, controlpoint) {
		return
	}

	controlpoint.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if controlpoint.ClosestRect != nil {
		UnstageBranch(stage, controlpoint.ClosestRect)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchEllipse(ellipse *Ellipse) {

	// check if instance is already staged
	if !IsStaged(stage, ellipse) {
		return
	}

	ellipse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range ellipse.Animates {
		UnstageBranch(stage, _animate)
	}

}

func (stage *Stage) UnstageBranchLayer(layer *Layer) {

	// check if instance is already staged
	if !IsStaged(stage, layer) {
		return
	}

	layer.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rect := range layer.Rects {
		UnstageBranch(stage, _rect)
	}
	for _, _text := range layer.Texts {
		UnstageBranch(stage, _text)
	}
	for _, _circle := range layer.Circles {
		UnstageBranch(stage, _circle)
	}
	for _, _line := range layer.Lines {
		UnstageBranch(stage, _line)
	}
	for _, _ellipse := range layer.Ellipses {
		UnstageBranch(stage, _ellipse)
	}
	for _, _polyline := range layer.Polylines {
		UnstageBranch(stage, _polyline)
	}
	for _, _polygone := range layer.Polygones {
		UnstageBranch(stage, _polygone)
	}
	for _, _path := range layer.Paths {
		UnstageBranch(stage, _path)
	}
	for _, _link := range layer.Links {
		UnstageBranch(stage, _link)
	}
	for _, _rectlinklink := range layer.RectLinkLinks {
		UnstageBranch(stage, _rectlinklink)
	}

}

func (stage *Stage) UnstageBranchLine(line *Line) {

	// check if instance is already staged
	if !IsStaged(stage, line) {
		return
	}

	line.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range line.Animates {
		UnstageBranch(stage, _animate)
	}

}

func (stage *Stage) UnstageBranchLink(link *Link) {

	// check if instance is already staged
	if !IsStaged(stage, link) {
		return
	}

	link.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if link.Start != nil {
		UnstageBranch(stage, link.Start)
	}
	if link.End != nil {
		UnstageBranch(stage, link.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _linkanchoredtext := range link.TextAtArrowStart {
		UnstageBranch(stage, _linkanchoredtext)
	}
	for _, _linkanchoredtext := range link.TextAtArrowEnd {
		UnstageBranch(stage, _linkanchoredtext)
	}
	for _, _controlpoint := range link.ControlPoints {
		UnstageBranch(stage, _controlpoint)
	}

}

func (stage *Stage) UnstageBranchLinkAnchoredText(linkanchoredtext *LinkAnchoredText) {

	// check if instance is already staged
	if !IsStaged(stage, linkanchoredtext) {
		return
	}

	linkanchoredtext.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range linkanchoredtext.Animates {
		UnstageBranch(stage, _animate)
	}

}

func (stage *Stage) UnstageBranchPath(path *Path) {

	// check if instance is already staged
	if !IsStaged(stage, path) {
		return
	}

	path.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range path.Animates {
		UnstageBranch(stage, _animate)
	}

}

func (stage *Stage) UnstageBranchPoint(point *Point) {

	// check if instance is already staged
	if !IsStaged(stage, point) {
		return
	}

	point.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPolygone(polygone *Polygone) {

	// check if instance is already staged
	if !IsStaged(stage, polygone) {
		return
	}

	polygone.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range polygone.Animates {
		UnstageBranch(stage, _animate)
	}

}

func (stage *Stage) UnstageBranchPolyline(polyline *Polyline) {

	// check if instance is already staged
	if !IsStaged(stage, polyline) {
		return
	}

	polyline.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range polyline.Animates {
		UnstageBranch(stage, _animate)
	}

}

func (stage *Stage) UnstageBranchRect(rect *Rect) {

	// check if instance is already staged
	if !IsStaged(stage, rect) {
		return
	}

	rect.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _condition := range rect.HoveringTrigger {
		UnstageBranch(stage, _condition)
	}
	for _, _condition := range rect.DisplayConditions {
		UnstageBranch(stage, _condition)
	}
	for _, _animate := range rect.Animations {
		UnstageBranch(stage, _animate)
	}
	for _, _rectanchoredtext := range rect.RectAnchoredTexts {
		UnstageBranch(stage, _rectanchoredtext)
	}
	for _, _rectanchoredrect := range rect.RectAnchoredRects {
		UnstageBranch(stage, _rectanchoredrect)
	}
	for _, _rectanchoredpath := range rect.RectAnchoredPaths {
		UnstageBranch(stage, _rectanchoredpath)
	}

}

func (stage *Stage) UnstageBranchRectAnchoredPath(rectanchoredpath *RectAnchoredPath) {

	// check if instance is already staged
	if !IsStaged(stage, rectanchoredpath) {
		return
	}

	rectanchoredpath.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchRectAnchoredRect(rectanchoredrect *RectAnchoredRect) {

	// check if instance is already staged
	if !IsStaged(stage, rectanchoredrect) {
		return
	}

	rectanchoredrect.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchRectAnchoredText(rectanchoredtext *RectAnchoredText) {

	// check if instance is already staged
	if !IsStaged(stage, rectanchoredtext) {
		return
	}

	rectanchoredtext.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range rectanchoredtext.Animates {
		UnstageBranch(stage, _animate)
	}

}

func (stage *Stage) UnstageBranchRectLinkLink(rectlinklink *RectLinkLink) {

	// check if instance is already staged
	if !IsStaged(stage, rectlinklink) {
		return
	}

	rectlinklink.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if rectlinklink.Start != nil {
		UnstageBranch(stage, rectlinklink.Start)
	}
	if rectlinklink.End != nil {
		UnstageBranch(stage, rectlinklink.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSVG(svg *SVG) {

	// check if instance is already staged
	if !IsStaged(stage, svg) {
		return
	}

	svg.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if svg.StartRect != nil {
		UnstageBranch(stage, svg.StartRect)
	}
	if svg.EndRect != nil {
		UnstageBranch(stage, svg.EndRect)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _layer := range svg.Layers {
		UnstageBranch(stage, _layer)
	}

}

func (stage *Stage) UnstageBranchSvgText(svgtext *SvgText) {

	// check if instance is already staged
	if !IsStaged(stage, svgtext) {
		return
	}

	svgtext.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchText(text *Text) {

	// check if instance is already staged
	if !IsStaged(stage, text) {
		return
	}

	text.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range text.Animates {
		UnstageBranch(stage, _animate)
	}

}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (animate *Animate) GongDiff(stage *Stage, animateOther *Animate) (diffs []string) {
	// insertion point for field diffs
	if animate.Name != animateOther.Name {
		diffs = append(diffs, animate.GongMarshallField(stage, "Name"))
	}
	if animate.AttributeName != animateOther.AttributeName {
		diffs = append(diffs, animate.GongMarshallField(stage, "AttributeName"))
	}
	if animate.Values != animateOther.Values {
		diffs = append(diffs, animate.GongMarshallField(stage, "Values"))
	}
	if animate.From != animateOther.From {
		diffs = append(diffs, animate.GongMarshallField(stage, "From"))
	}
	if animate.To != animateOther.To {
		diffs = append(diffs, animate.GongMarshallField(stage, "To"))
	}
	if animate.Dur != animateOther.Dur {
		diffs = append(diffs, animate.GongMarshallField(stage, "Dur"))
	}
	if animate.RepeatCount != animateOther.RepeatCount {
		diffs = append(diffs, animate.GongMarshallField(stage, "RepeatCount"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (circle *Circle) GongDiff(stage *Stage, circleOther *Circle) (diffs []string) {
	// insertion point for field diffs
	if circle.Name != circleOther.Name {
		diffs = append(diffs, circle.GongMarshallField(stage, "Name"))
	}
	if circle.CX != circleOther.CX {
		diffs = append(diffs, circle.GongMarshallField(stage, "CX"))
	}
	if circle.CY != circleOther.CY {
		diffs = append(diffs, circle.GongMarshallField(stage, "CY"))
	}
	if circle.Radius != circleOther.Radius {
		diffs = append(diffs, circle.GongMarshallField(stage, "Radius"))
	}
	if circle.Color != circleOther.Color {
		diffs = append(diffs, circle.GongMarshallField(stage, "Color"))
	}
	if circle.FillOpacity != circleOther.FillOpacity {
		diffs = append(diffs, circle.GongMarshallField(stage, "FillOpacity"))
	}
	if circle.Stroke != circleOther.Stroke {
		diffs = append(diffs, circle.GongMarshallField(stage, "Stroke"))
	}
	if circle.StrokeOpacity != circleOther.StrokeOpacity {
		diffs = append(diffs, circle.GongMarshallField(stage, "StrokeOpacity"))
	}
	if circle.StrokeWidth != circleOther.StrokeWidth {
		diffs = append(diffs, circle.GongMarshallField(stage, "StrokeWidth"))
	}
	if circle.StrokeDashArray != circleOther.StrokeDashArray {
		diffs = append(diffs, circle.GongMarshallField(stage, "StrokeDashArray"))
	}
	if circle.StrokeDashArrayWhenSelected != circleOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, circle.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if circle.Transform != circleOther.Transform {
		diffs = append(diffs, circle.GongMarshallField(stage, "Transform"))
	}
	AnimationsDifferent := false
	if len(circle.Animations) != len(circleOther.Animations) {
		AnimationsDifferent = true
	} else {
		for i := range circle.Animations {
			if (circle.Animations[i] == nil) != (circleOther.Animations[i] == nil) {
				AnimationsDifferent = true
				break
			} else if circle.Animations[i] != nil && circleOther.Animations[i] != nil {
				// this is a pointer comparaison
				if circle.Animations[i] != circleOther.Animations[i] {
					AnimationsDifferent = true
					break
				}
			}
		}
	}
	if AnimationsDifferent {
		ops := Diff(stage, circle, circleOther, "Animations", circleOther.Animations, circle.Animations)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (condition *Condition) GongDiff(stage *Stage, conditionOther *Condition) (diffs []string) {
	// insertion point for field diffs
	if condition.Name != conditionOther.Name {
		diffs = append(diffs, condition.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (controlpoint *ControlPoint) GongDiff(stage *Stage, controlpointOther *ControlPoint) (diffs []string) {
	// insertion point for field diffs
	if controlpoint.Name != controlpointOther.Name {
		diffs = append(diffs, controlpoint.GongMarshallField(stage, "Name"))
	}
	if controlpoint.X_Relative != controlpointOther.X_Relative {
		diffs = append(diffs, controlpoint.GongMarshallField(stage, "X_Relative"))
	}
	if controlpoint.Y_Relative != controlpointOther.Y_Relative {
		diffs = append(diffs, controlpoint.GongMarshallField(stage, "Y_Relative"))
	}
	if (controlpoint.ClosestRect == nil) != (controlpointOther.ClosestRect == nil) {
		diffs = append(diffs, controlpoint.GongMarshallField(stage, "ClosestRect"))
	} else if controlpoint.ClosestRect != nil && controlpointOther.ClosestRect != nil {
		if controlpoint.ClosestRect != controlpointOther.ClosestRect {
			diffs = append(diffs, controlpoint.GongMarshallField(stage, "ClosestRect"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ellipse *Ellipse) GongDiff(stage *Stage, ellipseOther *Ellipse) (diffs []string) {
	// insertion point for field diffs
	if ellipse.Name != ellipseOther.Name {
		diffs = append(diffs, ellipse.GongMarshallField(stage, "Name"))
	}
	if ellipse.CX != ellipseOther.CX {
		diffs = append(diffs, ellipse.GongMarshallField(stage, "CX"))
	}
	if ellipse.CY != ellipseOther.CY {
		diffs = append(diffs, ellipse.GongMarshallField(stage, "CY"))
	}
	if ellipse.RX != ellipseOther.RX {
		diffs = append(diffs, ellipse.GongMarshallField(stage, "RX"))
	}
	if ellipse.RY != ellipseOther.RY {
		diffs = append(diffs, ellipse.GongMarshallField(stage, "RY"))
	}
	if ellipse.Color != ellipseOther.Color {
		diffs = append(diffs, ellipse.GongMarshallField(stage, "Color"))
	}
	if ellipse.FillOpacity != ellipseOther.FillOpacity {
		diffs = append(diffs, ellipse.GongMarshallField(stage, "FillOpacity"))
	}
	if ellipse.Stroke != ellipseOther.Stroke {
		diffs = append(diffs, ellipse.GongMarshallField(stage, "Stroke"))
	}
	if ellipse.StrokeOpacity != ellipseOther.StrokeOpacity {
		diffs = append(diffs, ellipse.GongMarshallField(stage, "StrokeOpacity"))
	}
	if ellipse.StrokeWidth != ellipseOther.StrokeWidth {
		diffs = append(diffs, ellipse.GongMarshallField(stage, "StrokeWidth"))
	}
	if ellipse.StrokeDashArray != ellipseOther.StrokeDashArray {
		diffs = append(diffs, ellipse.GongMarshallField(stage, "StrokeDashArray"))
	}
	if ellipse.StrokeDashArrayWhenSelected != ellipseOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, ellipse.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if ellipse.Transform != ellipseOther.Transform {
		diffs = append(diffs, ellipse.GongMarshallField(stage, "Transform"))
	}
	AnimatesDifferent := false
	if len(ellipse.Animates) != len(ellipseOther.Animates) {
		AnimatesDifferent = true
	} else {
		for i := range ellipse.Animates {
			if (ellipse.Animates[i] == nil) != (ellipseOther.Animates[i] == nil) {
				AnimatesDifferent = true
				break
			} else if ellipse.Animates[i] != nil && ellipseOther.Animates[i] != nil {
				// this is a pointer comparaison
				if ellipse.Animates[i] != ellipseOther.Animates[i] {
					AnimatesDifferent = true
					break
				}
			}
		}
	}
	if AnimatesDifferent {
		ops := Diff(stage, ellipse, ellipseOther, "Animates", ellipseOther.Animates, ellipse.Animates)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (layer *Layer) GongDiff(stage *Stage, layerOther *Layer) (diffs []string) {
	// insertion point for field diffs
	if layer.Name != layerOther.Name {
		diffs = append(diffs, layer.GongMarshallField(stage, "Name"))
	}
	RectsDifferent := false
	if len(layer.Rects) != len(layerOther.Rects) {
		RectsDifferent = true
	} else {
		for i := range layer.Rects {
			if (layer.Rects[i] == nil) != (layerOther.Rects[i] == nil) {
				RectsDifferent = true
				break
			} else if layer.Rects[i] != nil && layerOther.Rects[i] != nil {
				// this is a pointer comparaison
				if layer.Rects[i] != layerOther.Rects[i] {
					RectsDifferent = true
					break
				}
			}
		}
	}
	if RectsDifferent {
		ops := Diff(stage, layer, layerOther, "Rects", layerOther.Rects, layer.Rects)
		diffs = append(diffs, ops)
	}
	TextsDifferent := false
	if len(layer.Texts) != len(layerOther.Texts) {
		TextsDifferent = true
	} else {
		for i := range layer.Texts {
			if (layer.Texts[i] == nil) != (layerOther.Texts[i] == nil) {
				TextsDifferent = true
				break
			} else if layer.Texts[i] != nil && layerOther.Texts[i] != nil {
				// this is a pointer comparaison
				if layer.Texts[i] != layerOther.Texts[i] {
					TextsDifferent = true
					break
				}
			}
		}
	}
	if TextsDifferent {
		ops := Diff(stage, layer, layerOther, "Texts", layerOther.Texts, layer.Texts)
		diffs = append(diffs, ops)
	}
	CirclesDifferent := false
	if len(layer.Circles) != len(layerOther.Circles) {
		CirclesDifferent = true
	} else {
		for i := range layer.Circles {
			if (layer.Circles[i] == nil) != (layerOther.Circles[i] == nil) {
				CirclesDifferent = true
				break
			} else if layer.Circles[i] != nil && layerOther.Circles[i] != nil {
				// this is a pointer comparaison
				if layer.Circles[i] != layerOther.Circles[i] {
					CirclesDifferent = true
					break
				}
			}
		}
	}
	if CirclesDifferent {
		ops := Diff(stage, layer, layerOther, "Circles", layerOther.Circles, layer.Circles)
		diffs = append(diffs, ops)
	}
	LinesDifferent := false
	if len(layer.Lines) != len(layerOther.Lines) {
		LinesDifferent = true
	} else {
		for i := range layer.Lines {
			if (layer.Lines[i] == nil) != (layerOther.Lines[i] == nil) {
				LinesDifferent = true
				break
			} else if layer.Lines[i] != nil && layerOther.Lines[i] != nil {
				// this is a pointer comparaison
				if layer.Lines[i] != layerOther.Lines[i] {
					LinesDifferent = true
					break
				}
			}
		}
	}
	if LinesDifferent {
		ops := Diff(stage, layer, layerOther, "Lines", layerOther.Lines, layer.Lines)
		diffs = append(diffs, ops)
	}
	EllipsesDifferent := false
	if len(layer.Ellipses) != len(layerOther.Ellipses) {
		EllipsesDifferent = true
	} else {
		for i := range layer.Ellipses {
			if (layer.Ellipses[i] == nil) != (layerOther.Ellipses[i] == nil) {
				EllipsesDifferent = true
				break
			} else if layer.Ellipses[i] != nil && layerOther.Ellipses[i] != nil {
				// this is a pointer comparaison
				if layer.Ellipses[i] != layerOther.Ellipses[i] {
					EllipsesDifferent = true
					break
				}
			}
		}
	}
	if EllipsesDifferent {
		ops := Diff(stage, layer, layerOther, "Ellipses", layerOther.Ellipses, layer.Ellipses)
		diffs = append(diffs, ops)
	}
	PolylinesDifferent := false
	if len(layer.Polylines) != len(layerOther.Polylines) {
		PolylinesDifferent = true
	} else {
		for i := range layer.Polylines {
			if (layer.Polylines[i] == nil) != (layerOther.Polylines[i] == nil) {
				PolylinesDifferent = true
				break
			} else if layer.Polylines[i] != nil && layerOther.Polylines[i] != nil {
				// this is a pointer comparaison
				if layer.Polylines[i] != layerOther.Polylines[i] {
					PolylinesDifferent = true
					break
				}
			}
		}
	}
	if PolylinesDifferent {
		ops := Diff(stage, layer, layerOther, "Polylines", layerOther.Polylines, layer.Polylines)
		diffs = append(diffs, ops)
	}
	PolygonesDifferent := false
	if len(layer.Polygones) != len(layerOther.Polygones) {
		PolygonesDifferent = true
	} else {
		for i := range layer.Polygones {
			if (layer.Polygones[i] == nil) != (layerOther.Polygones[i] == nil) {
				PolygonesDifferent = true
				break
			} else if layer.Polygones[i] != nil && layerOther.Polygones[i] != nil {
				// this is a pointer comparaison
				if layer.Polygones[i] != layerOther.Polygones[i] {
					PolygonesDifferent = true
					break
				}
			}
		}
	}
	if PolygonesDifferent {
		ops := Diff(stage, layer, layerOther, "Polygones", layerOther.Polygones, layer.Polygones)
		diffs = append(diffs, ops)
	}
	PathsDifferent := false
	if len(layer.Paths) != len(layerOther.Paths) {
		PathsDifferent = true
	} else {
		for i := range layer.Paths {
			if (layer.Paths[i] == nil) != (layerOther.Paths[i] == nil) {
				PathsDifferent = true
				break
			} else if layer.Paths[i] != nil && layerOther.Paths[i] != nil {
				// this is a pointer comparaison
				if layer.Paths[i] != layerOther.Paths[i] {
					PathsDifferent = true
					break
				}
			}
		}
	}
	if PathsDifferent {
		ops := Diff(stage, layer, layerOther, "Paths", layerOther.Paths, layer.Paths)
		diffs = append(diffs, ops)
	}
	LinksDifferent := false
	if len(layer.Links) != len(layerOther.Links) {
		LinksDifferent = true
	} else {
		for i := range layer.Links {
			if (layer.Links[i] == nil) != (layerOther.Links[i] == nil) {
				LinksDifferent = true
				break
			} else if layer.Links[i] != nil && layerOther.Links[i] != nil {
				// this is a pointer comparaison
				if layer.Links[i] != layerOther.Links[i] {
					LinksDifferent = true
					break
				}
			}
		}
	}
	if LinksDifferent {
		ops := Diff(stage, layer, layerOther, "Links", layerOther.Links, layer.Links)
		diffs = append(diffs, ops)
	}
	RectLinkLinksDifferent := false
	if len(layer.RectLinkLinks) != len(layerOther.RectLinkLinks) {
		RectLinkLinksDifferent = true
	} else {
		for i := range layer.RectLinkLinks {
			if (layer.RectLinkLinks[i] == nil) != (layerOther.RectLinkLinks[i] == nil) {
				RectLinkLinksDifferent = true
				break
			} else if layer.RectLinkLinks[i] != nil && layerOther.RectLinkLinks[i] != nil {
				// this is a pointer comparaison
				if layer.RectLinkLinks[i] != layerOther.RectLinkLinks[i] {
					RectLinkLinksDifferent = true
					break
				}
			}
		}
	}
	if RectLinkLinksDifferent {
		ops := Diff(stage, layer, layerOther, "RectLinkLinks", layerOther.RectLinkLinks, layer.RectLinkLinks)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (line *Line) GongDiff(stage *Stage, lineOther *Line) (diffs []string) {
	// insertion point for field diffs
	if line.Name != lineOther.Name {
		diffs = append(diffs, line.GongMarshallField(stage, "Name"))
	}
	if line.X1 != lineOther.X1 {
		diffs = append(diffs, line.GongMarshallField(stage, "X1"))
	}
	if line.Y1 != lineOther.Y1 {
		diffs = append(diffs, line.GongMarshallField(stage, "Y1"))
	}
	if line.X2 != lineOther.X2 {
		diffs = append(diffs, line.GongMarshallField(stage, "X2"))
	}
	if line.Y2 != lineOther.Y2 {
		diffs = append(diffs, line.GongMarshallField(stage, "Y2"))
	}
	if line.Color != lineOther.Color {
		diffs = append(diffs, line.GongMarshallField(stage, "Color"))
	}
	if line.FillOpacity != lineOther.FillOpacity {
		diffs = append(diffs, line.GongMarshallField(stage, "FillOpacity"))
	}
	if line.Stroke != lineOther.Stroke {
		diffs = append(diffs, line.GongMarshallField(stage, "Stroke"))
	}
	if line.StrokeOpacity != lineOther.StrokeOpacity {
		diffs = append(diffs, line.GongMarshallField(stage, "StrokeOpacity"))
	}
	if line.StrokeWidth != lineOther.StrokeWidth {
		diffs = append(diffs, line.GongMarshallField(stage, "StrokeWidth"))
	}
	if line.StrokeDashArray != lineOther.StrokeDashArray {
		diffs = append(diffs, line.GongMarshallField(stage, "StrokeDashArray"))
	}
	if line.StrokeDashArrayWhenSelected != lineOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, line.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if line.Transform != lineOther.Transform {
		diffs = append(diffs, line.GongMarshallField(stage, "Transform"))
	}
	AnimatesDifferent := false
	if len(line.Animates) != len(lineOther.Animates) {
		AnimatesDifferent = true
	} else {
		for i := range line.Animates {
			if (line.Animates[i] == nil) != (lineOther.Animates[i] == nil) {
				AnimatesDifferent = true
				break
			} else if line.Animates[i] != nil && lineOther.Animates[i] != nil {
				// this is a pointer comparaison
				if line.Animates[i] != lineOther.Animates[i] {
					AnimatesDifferent = true
					break
				}
			}
		}
	}
	if AnimatesDifferent {
		ops := Diff(stage, line, lineOther, "Animates", lineOther.Animates, line.Animates)
		diffs = append(diffs, ops)
	}
	if line.MouseClickX != lineOther.MouseClickX {
		diffs = append(diffs, line.GongMarshallField(stage, "MouseClickX"))
	}
	if line.MouseClickY != lineOther.MouseClickY {
		diffs = append(diffs, line.GongMarshallField(stage, "MouseClickY"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (link *Link) GongDiff(stage *Stage, linkOther *Link) (diffs []string) {
	// insertion point for field diffs
	if link.Name != linkOther.Name {
		diffs = append(diffs, link.GongMarshallField(stage, "Name"))
	}
	if link.Type != linkOther.Type {
		diffs = append(diffs, link.GongMarshallField(stage, "Type"))
	}
	if link.IsBezierCurve != linkOther.IsBezierCurve {
		diffs = append(diffs, link.GongMarshallField(stage, "IsBezierCurve"))
	}
	if (link.Start == nil) != (linkOther.Start == nil) {
		diffs = append(diffs, link.GongMarshallField(stage, "Start"))
	} else if link.Start != nil && linkOther.Start != nil {
		if link.Start != linkOther.Start {
			diffs = append(diffs, link.GongMarshallField(stage, "Start"))
		}
	}
	if link.StartAnchorType != linkOther.StartAnchorType {
		diffs = append(diffs, link.GongMarshallField(stage, "StartAnchorType"))
	}
	if (link.End == nil) != (linkOther.End == nil) {
		diffs = append(diffs, link.GongMarshallField(stage, "End"))
	} else if link.End != nil && linkOther.End != nil {
		if link.End != linkOther.End {
			diffs = append(diffs, link.GongMarshallField(stage, "End"))
		}
	}
	if link.EndAnchorType != linkOther.EndAnchorType {
		diffs = append(diffs, link.GongMarshallField(stage, "EndAnchorType"))
	}
	if link.StartOrientation != linkOther.StartOrientation {
		diffs = append(diffs, link.GongMarshallField(stage, "StartOrientation"))
	}
	if link.StartRatio != linkOther.StartRatio {
		diffs = append(diffs, link.GongMarshallField(stage, "StartRatio"))
	}
	if link.EndOrientation != linkOther.EndOrientation {
		diffs = append(diffs, link.GongMarshallField(stage, "EndOrientation"))
	}
	if link.EndRatio != linkOther.EndRatio {
		diffs = append(diffs, link.GongMarshallField(stage, "EndRatio"))
	}
	if link.CornerOffsetRatio != linkOther.CornerOffsetRatio {
		diffs = append(diffs, link.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if link.CornerRadius != linkOther.CornerRadius {
		diffs = append(diffs, link.GongMarshallField(stage, "CornerRadius"))
	}
	if link.HasEndArrow != linkOther.HasEndArrow {
		diffs = append(diffs, link.GongMarshallField(stage, "HasEndArrow"))
	}
	if link.EndArrowSize != linkOther.EndArrowSize {
		diffs = append(diffs, link.GongMarshallField(stage, "EndArrowSize"))
	}
	if link.EndArrowOffset != linkOther.EndArrowOffset {
		diffs = append(diffs, link.GongMarshallField(stage, "EndArrowOffset"))
	}
	if link.HasStartArrow != linkOther.HasStartArrow {
		diffs = append(diffs, link.GongMarshallField(stage, "HasStartArrow"))
	}
	if link.StartArrowSize != linkOther.StartArrowSize {
		diffs = append(diffs, link.GongMarshallField(stage, "StartArrowSize"))
	}
	if link.StartArrowOffset != linkOther.StartArrowOffset {
		diffs = append(diffs, link.GongMarshallField(stage, "StartArrowOffset"))
	}
	TextAtArrowStartDifferent := false
	if len(link.TextAtArrowStart) != len(linkOther.TextAtArrowStart) {
		TextAtArrowStartDifferent = true
	} else {
		for i := range link.TextAtArrowStart {
			if (link.TextAtArrowStart[i] == nil) != (linkOther.TextAtArrowStart[i] == nil) {
				TextAtArrowStartDifferent = true
				break
			} else if link.TextAtArrowStart[i] != nil && linkOther.TextAtArrowStart[i] != nil {
				// this is a pointer comparaison
				if link.TextAtArrowStart[i] != linkOther.TextAtArrowStart[i] {
					TextAtArrowStartDifferent = true
					break
				}
			}
		}
	}
	if TextAtArrowStartDifferent {
		ops := Diff(stage, link, linkOther, "TextAtArrowStart", linkOther.TextAtArrowStart, link.TextAtArrowStart)
		diffs = append(diffs, ops)
	}
	TextAtArrowEndDifferent := false
	if len(link.TextAtArrowEnd) != len(linkOther.TextAtArrowEnd) {
		TextAtArrowEndDifferent = true
	} else {
		for i := range link.TextAtArrowEnd {
			if (link.TextAtArrowEnd[i] == nil) != (linkOther.TextAtArrowEnd[i] == nil) {
				TextAtArrowEndDifferent = true
				break
			} else if link.TextAtArrowEnd[i] != nil && linkOther.TextAtArrowEnd[i] != nil {
				// this is a pointer comparaison
				if link.TextAtArrowEnd[i] != linkOther.TextAtArrowEnd[i] {
					TextAtArrowEndDifferent = true
					break
				}
			}
		}
	}
	if TextAtArrowEndDifferent {
		ops := Diff(stage, link, linkOther, "TextAtArrowEnd", linkOther.TextAtArrowEnd, link.TextAtArrowEnd)
		diffs = append(diffs, ops)
	}
	ControlPointsDifferent := false
	if len(link.ControlPoints) != len(linkOther.ControlPoints) {
		ControlPointsDifferent = true
	} else {
		for i := range link.ControlPoints {
			if (link.ControlPoints[i] == nil) != (linkOther.ControlPoints[i] == nil) {
				ControlPointsDifferent = true
				break
			} else if link.ControlPoints[i] != nil && linkOther.ControlPoints[i] != nil {
				// this is a pointer comparaison
				if link.ControlPoints[i] != linkOther.ControlPoints[i] {
					ControlPointsDifferent = true
					break
				}
			}
		}
	}
	if ControlPointsDifferent {
		ops := Diff(stage, link, linkOther, "ControlPoints", linkOther.ControlPoints, link.ControlPoints)
		diffs = append(diffs, ops)
	}
	if link.Color != linkOther.Color {
		diffs = append(diffs, link.GongMarshallField(stage, "Color"))
	}
	if link.FillOpacity != linkOther.FillOpacity {
		diffs = append(diffs, link.GongMarshallField(stage, "FillOpacity"))
	}
	if link.Stroke != linkOther.Stroke {
		diffs = append(diffs, link.GongMarshallField(stage, "Stroke"))
	}
	if link.StrokeOpacity != linkOther.StrokeOpacity {
		diffs = append(diffs, link.GongMarshallField(stage, "StrokeOpacity"))
	}
	if link.StrokeWidth != linkOther.StrokeWidth {
		diffs = append(diffs, link.GongMarshallField(stage, "StrokeWidth"))
	}
	if link.StrokeDashArray != linkOther.StrokeDashArray {
		diffs = append(diffs, link.GongMarshallField(stage, "StrokeDashArray"))
	}
	if link.StrokeDashArrayWhenSelected != linkOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, link.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if link.Transform != linkOther.Transform {
		diffs = append(diffs, link.GongMarshallField(stage, "Transform"))
	}
	if link.MouseX != linkOther.MouseX {
		diffs = append(diffs, link.GongMarshallField(stage, "MouseX"))
	}
	if link.MouseY != linkOther.MouseY {
		diffs = append(diffs, link.GongMarshallField(stage, "MouseY"))
	}
	if link.MouseEventKey != linkOther.MouseEventKey {
		diffs = append(diffs, link.GongMarshallField(stage, "MouseEventKey"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (linkanchoredtext *LinkAnchoredText) GongDiff(stage *Stage, linkanchoredtextOther *LinkAnchoredText) (diffs []string) {
	// insertion point for field diffs
	if linkanchoredtext.Name != linkanchoredtextOther.Name {
		diffs = append(diffs, linkanchoredtext.GongMarshallField(stage, "Name"))
	}
	if linkanchoredtext.Content != linkanchoredtextOther.Content {
		diffs = append(diffs, linkanchoredtext.GongMarshallField(stage, "Content"))
	}
	if linkanchoredtext.AutomaticLayout != linkanchoredtextOther.AutomaticLayout {
		diffs = append(diffs, linkanchoredtext.GongMarshallField(stage, "AutomaticLayout"))
	}
	if linkanchoredtext.LinkAnchorType != linkanchoredtextOther.LinkAnchorType {
		diffs = append(diffs, linkanchoredtext.GongMarshallField(stage, "LinkAnchorType"))
	}
	if linkanchoredtext.X_Offset != linkanchoredtextOther.X_Offset {
		diffs = append(diffs, linkanchoredtext.GongMarshallField(stage, "X_Offset"))
	}
	if linkanchoredtext.Y_Offset != linkanchoredtextOther.Y_Offset {
		diffs = append(diffs, linkanchoredtext.GongMarshallField(stage, "Y_Offset"))
	}
	if linkanchoredtext.FontWeight != linkanchoredtextOther.FontWeight {
		diffs = append(diffs, linkanchoredtext.GongMarshallField(stage, "FontWeight"))
	}
	if linkanchoredtext.FontSize != linkanchoredtextOther.FontSize {
		diffs = append(diffs, linkanchoredtext.GongMarshallField(stage, "FontSize"))
	}
	if linkanchoredtext.FontStyle != linkanchoredtextOther.FontStyle {
		diffs = append(diffs, linkanchoredtext.GongMarshallField(stage, "FontStyle"))
	}
	if linkanchoredtext.LetterSpacing != linkanchoredtextOther.LetterSpacing {
		diffs = append(diffs, linkanchoredtext.GongMarshallField(stage, "LetterSpacing"))
	}
	if linkanchoredtext.FontFamily != linkanchoredtextOther.FontFamily {
		diffs = append(diffs, linkanchoredtext.GongMarshallField(stage, "FontFamily"))
	}
	if linkanchoredtext.WhiteSpace != linkanchoredtextOther.WhiteSpace {
		diffs = append(diffs, linkanchoredtext.GongMarshallField(stage, "WhiteSpace"))
	}
	if linkanchoredtext.Color != linkanchoredtextOther.Color {
		diffs = append(diffs, linkanchoredtext.GongMarshallField(stage, "Color"))
	}
	if linkanchoredtext.FillOpacity != linkanchoredtextOther.FillOpacity {
		diffs = append(diffs, linkanchoredtext.GongMarshallField(stage, "FillOpacity"))
	}
	if linkanchoredtext.Stroke != linkanchoredtextOther.Stroke {
		diffs = append(diffs, linkanchoredtext.GongMarshallField(stage, "Stroke"))
	}
	if linkanchoredtext.StrokeOpacity != linkanchoredtextOther.StrokeOpacity {
		diffs = append(diffs, linkanchoredtext.GongMarshallField(stage, "StrokeOpacity"))
	}
	if linkanchoredtext.StrokeWidth != linkanchoredtextOther.StrokeWidth {
		diffs = append(diffs, linkanchoredtext.GongMarshallField(stage, "StrokeWidth"))
	}
	if linkanchoredtext.StrokeDashArray != linkanchoredtextOther.StrokeDashArray {
		diffs = append(diffs, linkanchoredtext.GongMarshallField(stage, "StrokeDashArray"))
	}
	if linkanchoredtext.StrokeDashArrayWhenSelected != linkanchoredtextOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, linkanchoredtext.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if linkanchoredtext.Transform != linkanchoredtextOther.Transform {
		diffs = append(diffs, linkanchoredtext.GongMarshallField(stage, "Transform"))
	}
	AnimatesDifferent := false
	if len(linkanchoredtext.Animates) != len(linkanchoredtextOther.Animates) {
		AnimatesDifferent = true
	} else {
		for i := range linkanchoredtext.Animates {
			if (linkanchoredtext.Animates[i] == nil) != (linkanchoredtextOther.Animates[i] == nil) {
				AnimatesDifferent = true
				break
			} else if linkanchoredtext.Animates[i] != nil && linkanchoredtextOther.Animates[i] != nil {
				// this is a pointer comparaison
				if linkanchoredtext.Animates[i] != linkanchoredtextOther.Animates[i] {
					AnimatesDifferent = true
					break
				}
			}
		}
	}
	if AnimatesDifferent {
		ops := Diff(stage, linkanchoredtext, linkanchoredtextOther, "Animates", linkanchoredtextOther.Animates, linkanchoredtext.Animates)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (path *Path) GongDiff(stage *Stage, pathOther *Path) (diffs []string) {
	// insertion point for field diffs
	if path.Name != pathOther.Name {
		diffs = append(diffs, path.GongMarshallField(stage, "Name"))
	}
	if path.Definition != pathOther.Definition {
		diffs = append(diffs, path.GongMarshallField(stage, "Definition"))
	}
	if path.Color != pathOther.Color {
		diffs = append(diffs, path.GongMarshallField(stage, "Color"))
	}
	if path.FillOpacity != pathOther.FillOpacity {
		diffs = append(diffs, path.GongMarshallField(stage, "FillOpacity"))
	}
	if path.Stroke != pathOther.Stroke {
		diffs = append(diffs, path.GongMarshallField(stage, "Stroke"))
	}
	if path.StrokeOpacity != pathOther.StrokeOpacity {
		diffs = append(diffs, path.GongMarshallField(stage, "StrokeOpacity"))
	}
	if path.StrokeWidth != pathOther.StrokeWidth {
		diffs = append(diffs, path.GongMarshallField(stage, "StrokeWidth"))
	}
	if path.StrokeDashArray != pathOther.StrokeDashArray {
		diffs = append(diffs, path.GongMarshallField(stage, "StrokeDashArray"))
	}
	if path.StrokeDashArrayWhenSelected != pathOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, path.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if path.Transform != pathOther.Transform {
		diffs = append(diffs, path.GongMarshallField(stage, "Transform"))
	}
	AnimatesDifferent := false
	if len(path.Animates) != len(pathOther.Animates) {
		AnimatesDifferent = true
	} else {
		for i := range path.Animates {
			if (path.Animates[i] == nil) != (pathOther.Animates[i] == nil) {
				AnimatesDifferent = true
				break
			} else if path.Animates[i] != nil && pathOther.Animates[i] != nil {
				// this is a pointer comparaison
				if path.Animates[i] != pathOther.Animates[i] {
					AnimatesDifferent = true
					break
				}
			}
		}
	}
	if AnimatesDifferent {
		ops := Diff(stage, path, pathOther, "Animates", pathOther.Animates, path.Animates)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (point *Point) GongDiff(stage *Stage, pointOther *Point) (diffs []string) {
	// insertion point for field diffs
	if point.Name != pointOther.Name {
		diffs = append(diffs, point.GongMarshallField(stage, "Name"))
	}
	if point.X != pointOther.X {
		diffs = append(diffs, point.GongMarshallField(stage, "X"))
	}
	if point.Y != pointOther.Y {
		diffs = append(diffs, point.GongMarshallField(stage, "Y"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (polygone *Polygone) GongDiff(stage *Stage, polygoneOther *Polygone) (diffs []string) {
	// insertion point for field diffs
	if polygone.Name != polygoneOther.Name {
		diffs = append(diffs, polygone.GongMarshallField(stage, "Name"))
	}
	if polygone.Points != polygoneOther.Points {
		diffs = append(diffs, polygone.GongMarshallField(stage, "Points"))
	}
	if polygone.Color != polygoneOther.Color {
		diffs = append(diffs, polygone.GongMarshallField(stage, "Color"))
	}
	if polygone.FillOpacity != polygoneOther.FillOpacity {
		diffs = append(diffs, polygone.GongMarshallField(stage, "FillOpacity"))
	}
	if polygone.Stroke != polygoneOther.Stroke {
		diffs = append(diffs, polygone.GongMarshallField(stage, "Stroke"))
	}
	if polygone.StrokeOpacity != polygoneOther.StrokeOpacity {
		diffs = append(diffs, polygone.GongMarshallField(stage, "StrokeOpacity"))
	}
	if polygone.StrokeWidth != polygoneOther.StrokeWidth {
		diffs = append(diffs, polygone.GongMarshallField(stage, "StrokeWidth"))
	}
	if polygone.StrokeDashArray != polygoneOther.StrokeDashArray {
		diffs = append(diffs, polygone.GongMarshallField(stage, "StrokeDashArray"))
	}
	if polygone.StrokeDashArrayWhenSelected != polygoneOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, polygone.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if polygone.Transform != polygoneOther.Transform {
		diffs = append(diffs, polygone.GongMarshallField(stage, "Transform"))
	}
	AnimatesDifferent := false
	if len(polygone.Animates) != len(polygoneOther.Animates) {
		AnimatesDifferent = true
	} else {
		for i := range polygone.Animates {
			if (polygone.Animates[i] == nil) != (polygoneOther.Animates[i] == nil) {
				AnimatesDifferent = true
				break
			} else if polygone.Animates[i] != nil && polygoneOther.Animates[i] != nil {
				// this is a pointer comparaison
				if polygone.Animates[i] != polygoneOther.Animates[i] {
					AnimatesDifferent = true
					break
				}
			}
		}
	}
	if AnimatesDifferent {
		ops := Diff(stage, polygone, polygoneOther, "Animates", polygoneOther.Animates, polygone.Animates)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (polyline *Polyline) GongDiff(stage *Stage, polylineOther *Polyline) (diffs []string) {
	// insertion point for field diffs
	if polyline.Name != polylineOther.Name {
		diffs = append(diffs, polyline.GongMarshallField(stage, "Name"))
	}
	if polyline.Points != polylineOther.Points {
		diffs = append(diffs, polyline.GongMarshallField(stage, "Points"))
	}
	if polyline.Color != polylineOther.Color {
		diffs = append(diffs, polyline.GongMarshallField(stage, "Color"))
	}
	if polyline.FillOpacity != polylineOther.FillOpacity {
		diffs = append(diffs, polyline.GongMarshallField(stage, "FillOpacity"))
	}
	if polyline.Stroke != polylineOther.Stroke {
		diffs = append(diffs, polyline.GongMarshallField(stage, "Stroke"))
	}
	if polyline.StrokeOpacity != polylineOther.StrokeOpacity {
		diffs = append(diffs, polyline.GongMarshallField(stage, "StrokeOpacity"))
	}
	if polyline.StrokeWidth != polylineOther.StrokeWidth {
		diffs = append(diffs, polyline.GongMarshallField(stage, "StrokeWidth"))
	}
	if polyline.StrokeDashArray != polylineOther.StrokeDashArray {
		diffs = append(diffs, polyline.GongMarshallField(stage, "StrokeDashArray"))
	}
	if polyline.StrokeDashArrayWhenSelected != polylineOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, polyline.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if polyline.Transform != polylineOther.Transform {
		diffs = append(diffs, polyline.GongMarshallField(stage, "Transform"))
	}
	AnimatesDifferent := false
	if len(polyline.Animates) != len(polylineOther.Animates) {
		AnimatesDifferent = true
	} else {
		for i := range polyline.Animates {
			if (polyline.Animates[i] == nil) != (polylineOther.Animates[i] == nil) {
				AnimatesDifferent = true
				break
			} else if polyline.Animates[i] != nil && polylineOther.Animates[i] != nil {
				// this is a pointer comparaison
				if polyline.Animates[i] != polylineOther.Animates[i] {
					AnimatesDifferent = true
					break
				}
			}
		}
	}
	if AnimatesDifferent {
		ops := Diff(stage, polyline, polylineOther, "Animates", polylineOther.Animates, polyline.Animates)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (rect *Rect) GongDiff(stage *Stage, rectOther *Rect) (diffs []string) {
	// insertion point for field diffs
	if rect.Name != rectOther.Name {
		diffs = append(diffs, rect.GongMarshallField(stage, "Name"))
	}
	if rect.X != rectOther.X {
		diffs = append(diffs, rect.GongMarshallField(stage, "X"))
	}
	if rect.Y != rectOther.Y {
		diffs = append(diffs, rect.GongMarshallField(stage, "Y"))
	}
	if rect.Width != rectOther.Width {
		diffs = append(diffs, rect.GongMarshallField(stage, "Width"))
	}
	if rect.Height != rectOther.Height {
		diffs = append(diffs, rect.GongMarshallField(stage, "Height"))
	}
	if rect.RX != rectOther.RX {
		diffs = append(diffs, rect.GongMarshallField(stage, "RX"))
	}
	if rect.Color != rectOther.Color {
		diffs = append(diffs, rect.GongMarshallField(stage, "Color"))
	}
	if rect.FillOpacity != rectOther.FillOpacity {
		diffs = append(diffs, rect.GongMarshallField(stage, "FillOpacity"))
	}
	if rect.Stroke != rectOther.Stroke {
		diffs = append(diffs, rect.GongMarshallField(stage, "Stroke"))
	}
	if rect.StrokeOpacity != rectOther.StrokeOpacity {
		diffs = append(diffs, rect.GongMarshallField(stage, "StrokeOpacity"))
	}
	if rect.StrokeWidth != rectOther.StrokeWidth {
		diffs = append(diffs, rect.GongMarshallField(stage, "StrokeWidth"))
	}
	if rect.StrokeDashArray != rectOther.StrokeDashArray {
		diffs = append(diffs, rect.GongMarshallField(stage, "StrokeDashArray"))
	}
	if rect.StrokeDashArrayWhenSelected != rectOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, rect.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if rect.Transform != rectOther.Transform {
		diffs = append(diffs, rect.GongMarshallField(stage, "Transform"))
	}
	HoveringTriggerDifferent := false
	if len(rect.HoveringTrigger) != len(rectOther.HoveringTrigger) {
		HoveringTriggerDifferent = true
	} else {
		for i := range rect.HoveringTrigger {
			if (rect.HoveringTrigger[i] == nil) != (rectOther.HoveringTrigger[i] == nil) {
				HoveringTriggerDifferent = true
				break
			} else if rect.HoveringTrigger[i] != nil && rectOther.HoveringTrigger[i] != nil {
				// this is a pointer comparaison
				if rect.HoveringTrigger[i] != rectOther.HoveringTrigger[i] {
					HoveringTriggerDifferent = true
					break
				}
			}
		}
	}
	if HoveringTriggerDifferent {
		ops := Diff(stage, rect, rectOther, "HoveringTrigger", rectOther.HoveringTrigger, rect.HoveringTrigger)
		diffs = append(diffs, ops)
	}
	DisplayConditionsDifferent := false
	if len(rect.DisplayConditions) != len(rectOther.DisplayConditions) {
		DisplayConditionsDifferent = true
	} else {
		for i := range rect.DisplayConditions {
			if (rect.DisplayConditions[i] == nil) != (rectOther.DisplayConditions[i] == nil) {
				DisplayConditionsDifferent = true
				break
			} else if rect.DisplayConditions[i] != nil && rectOther.DisplayConditions[i] != nil {
				// this is a pointer comparaison
				if rect.DisplayConditions[i] != rectOther.DisplayConditions[i] {
					DisplayConditionsDifferent = true
					break
				}
			}
		}
	}
	if DisplayConditionsDifferent {
		ops := Diff(stage, rect, rectOther, "DisplayConditions", rectOther.DisplayConditions, rect.DisplayConditions)
		diffs = append(diffs, ops)
	}
	AnimationsDifferent := false
	if len(rect.Animations) != len(rectOther.Animations) {
		AnimationsDifferent = true
	} else {
		for i := range rect.Animations {
			if (rect.Animations[i] == nil) != (rectOther.Animations[i] == nil) {
				AnimationsDifferent = true
				break
			} else if rect.Animations[i] != nil && rectOther.Animations[i] != nil {
				// this is a pointer comparaison
				if rect.Animations[i] != rectOther.Animations[i] {
					AnimationsDifferent = true
					break
				}
			}
		}
	}
	if AnimationsDifferent {
		ops := Diff(stage, rect, rectOther, "Animations", rectOther.Animations, rect.Animations)
		diffs = append(diffs, ops)
	}
	if rect.IsSelectable != rectOther.IsSelectable {
		diffs = append(diffs, rect.GongMarshallField(stage, "IsSelectable"))
	}
	if rect.IsSelected != rectOther.IsSelected {
		diffs = append(diffs, rect.GongMarshallField(stage, "IsSelected"))
	}
	if rect.CanHaveLeftHandle != rectOther.CanHaveLeftHandle {
		diffs = append(diffs, rect.GongMarshallField(stage, "CanHaveLeftHandle"))
	}
	if rect.HasLeftHandle != rectOther.HasLeftHandle {
		diffs = append(diffs, rect.GongMarshallField(stage, "HasLeftHandle"))
	}
	if rect.CanHaveRightHandle != rectOther.CanHaveRightHandle {
		diffs = append(diffs, rect.GongMarshallField(stage, "CanHaveRightHandle"))
	}
	if rect.HasRightHandle != rectOther.HasRightHandle {
		diffs = append(diffs, rect.GongMarshallField(stage, "HasRightHandle"))
	}
	if rect.CanHaveTopHandle != rectOther.CanHaveTopHandle {
		diffs = append(diffs, rect.GongMarshallField(stage, "CanHaveTopHandle"))
	}
	if rect.HasTopHandle != rectOther.HasTopHandle {
		diffs = append(diffs, rect.GongMarshallField(stage, "HasTopHandle"))
	}
	if rect.IsScalingProportionally != rectOther.IsScalingProportionally {
		diffs = append(diffs, rect.GongMarshallField(stage, "IsScalingProportionally"))
	}
	if rect.CanHaveBottomHandle != rectOther.CanHaveBottomHandle {
		diffs = append(diffs, rect.GongMarshallField(stage, "CanHaveBottomHandle"))
	}
	if rect.HasBottomHandle != rectOther.HasBottomHandle {
		diffs = append(diffs, rect.GongMarshallField(stage, "HasBottomHandle"))
	}
	if rect.CanMoveHorizontaly != rectOther.CanMoveHorizontaly {
		diffs = append(diffs, rect.GongMarshallField(stage, "CanMoveHorizontaly"))
	}
	if rect.CanMoveVerticaly != rectOther.CanMoveVerticaly {
		diffs = append(diffs, rect.GongMarshallField(stage, "CanMoveVerticaly"))
	}
	RectAnchoredTextsDifferent := false
	if len(rect.RectAnchoredTexts) != len(rectOther.RectAnchoredTexts) {
		RectAnchoredTextsDifferent = true
	} else {
		for i := range rect.RectAnchoredTexts {
			if (rect.RectAnchoredTexts[i] == nil) != (rectOther.RectAnchoredTexts[i] == nil) {
				RectAnchoredTextsDifferent = true
				break
			} else if rect.RectAnchoredTexts[i] != nil && rectOther.RectAnchoredTexts[i] != nil {
				// this is a pointer comparaison
				if rect.RectAnchoredTexts[i] != rectOther.RectAnchoredTexts[i] {
					RectAnchoredTextsDifferent = true
					break
				}
			}
		}
	}
	if RectAnchoredTextsDifferent {
		ops := Diff(stage, rect, rectOther, "RectAnchoredTexts", rectOther.RectAnchoredTexts, rect.RectAnchoredTexts)
		diffs = append(diffs, ops)
	}
	RectAnchoredRectsDifferent := false
	if len(rect.RectAnchoredRects) != len(rectOther.RectAnchoredRects) {
		RectAnchoredRectsDifferent = true
	} else {
		for i := range rect.RectAnchoredRects {
			if (rect.RectAnchoredRects[i] == nil) != (rectOther.RectAnchoredRects[i] == nil) {
				RectAnchoredRectsDifferent = true
				break
			} else if rect.RectAnchoredRects[i] != nil && rectOther.RectAnchoredRects[i] != nil {
				// this is a pointer comparaison
				if rect.RectAnchoredRects[i] != rectOther.RectAnchoredRects[i] {
					RectAnchoredRectsDifferent = true
					break
				}
			}
		}
	}
	if RectAnchoredRectsDifferent {
		ops := Diff(stage, rect, rectOther, "RectAnchoredRects", rectOther.RectAnchoredRects, rect.RectAnchoredRects)
		diffs = append(diffs, ops)
	}
	RectAnchoredPathsDifferent := false
	if len(rect.RectAnchoredPaths) != len(rectOther.RectAnchoredPaths) {
		RectAnchoredPathsDifferent = true
	} else {
		for i := range rect.RectAnchoredPaths {
			if (rect.RectAnchoredPaths[i] == nil) != (rectOther.RectAnchoredPaths[i] == nil) {
				RectAnchoredPathsDifferent = true
				break
			} else if rect.RectAnchoredPaths[i] != nil && rectOther.RectAnchoredPaths[i] != nil {
				// this is a pointer comparaison
				if rect.RectAnchoredPaths[i] != rectOther.RectAnchoredPaths[i] {
					RectAnchoredPathsDifferent = true
					break
				}
			}
		}
	}
	if RectAnchoredPathsDifferent {
		ops := Diff(stage, rect, rectOther, "RectAnchoredPaths", rectOther.RectAnchoredPaths, rect.RectAnchoredPaths)
		diffs = append(diffs, ops)
	}
	if rect.ChangeColorWhenHovered != rectOther.ChangeColorWhenHovered {
		diffs = append(diffs, rect.GongMarshallField(stage, "ChangeColorWhenHovered"))
	}
	if rect.ColorWhenHovered != rectOther.ColorWhenHovered {
		diffs = append(diffs, rect.GongMarshallField(stage, "ColorWhenHovered"))
	}
	if rect.OriginalColor != rectOther.OriginalColor {
		diffs = append(diffs, rect.GongMarshallField(stage, "OriginalColor"))
	}
	if rect.FillOpacityWhenHovered != rectOther.FillOpacityWhenHovered {
		diffs = append(diffs, rect.GongMarshallField(stage, "FillOpacityWhenHovered"))
	}
	if rect.OriginalFillOpacity != rectOther.OriginalFillOpacity {
		diffs = append(diffs, rect.GongMarshallField(stage, "OriginalFillOpacity"))
	}
	if rect.HasToolTip != rectOther.HasToolTip {
		diffs = append(diffs, rect.GongMarshallField(stage, "HasToolTip"))
	}
	if rect.ToolTipText != rectOther.ToolTipText {
		diffs = append(diffs, rect.GongMarshallField(stage, "ToolTipText"))
	}
	if rect.ToolTipPosition != rectOther.ToolTipPosition {
		diffs = append(diffs, rect.GongMarshallField(stage, "ToolTipPosition"))
	}
	if rect.MouseX != rectOther.MouseX {
		diffs = append(diffs, rect.GongMarshallField(stage, "MouseX"))
	}
	if rect.MouseY != rectOther.MouseY {
		diffs = append(diffs, rect.GongMarshallField(stage, "MouseY"))
	}
	if rect.MouseEventKey != rectOther.MouseEventKey {
		diffs = append(diffs, rect.GongMarshallField(stage, "MouseEventKey"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (rectanchoredpath *RectAnchoredPath) GongDiff(stage *Stage, rectanchoredpathOther *RectAnchoredPath) (diffs []string) {
	// insertion point for field diffs
	if rectanchoredpath.Name != rectanchoredpathOther.Name {
		diffs = append(diffs, rectanchoredpath.GongMarshallField(stage, "Name"))
	}
	if rectanchoredpath.Definition != rectanchoredpathOther.Definition {
		diffs = append(diffs, rectanchoredpath.GongMarshallField(stage, "Definition"))
	}
	if rectanchoredpath.X_Offset != rectanchoredpathOther.X_Offset {
		diffs = append(diffs, rectanchoredpath.GongMarshallField(stage, "X_Offset"))
	}
	if rectanchoredpath.Y_Offset != rectanchoredpathOther.Y_Offset {
		diffs = append(diffs, rectanchoredpath.GongMarshallField(stage, "Y_Offset"))
	}
	if rectanchoredpath.RectAnchorType != rectanchoredpathOther.RectAnchorType {
		diffs = append(diffs, rectanchoredpath.GongMarshallField(stage, "RectAnchorType"))
	}
	if rectanchoredpath.ScalePropotionnally != rectanchoredpathOther.ScalePropotionnally {
		diffs = append(diffs, rectanchoredpath.GongMarshallField(stage, "ScalePropotionnally"))
	}
	if rectanchoredpath.AppliedScaling != rectanchoredpathOther.AppliedScaling {
		diffs = append(diffs, rectanchoredpath.GongMarshallField(stage, "AppliedScaling"))
	}
	if rectanchoredpath.Color != rectanchoredpathOther.Color {
		diffs = append(diffs, rectanchoredpath.GongMarshallField(stage, "Color"))
	}
	if rectanchoredpath.FillOpacity != rectanchoredpathOther.FillOpacity {
		diffs = append(diffs, rectanchoredpath.GongMarshallField(stage, "FillOpacity"))
	}
	if rectanchoredpath.Stroke != rectanchoredpathOther.Stroke {
		diffs = append(diffs, rectanchoredpath.GongMarshallField(stage, "Stroke"))
	}
	if rectanchoredpath.StrokeOpacity != rectanchoredpathOther.StrokeOpacity {
		diffs = append(diffs, rectanchoredpath.GongMarshallField(stage, "StrokeOpacity"))
	}
	if rectanchoredpath.StrokeWidth != rectanchoredpathOther.StrokeWidth {
		diffs = append(diffs, rectanchoredpath.GongMarshallField(stage, "StrokeWidth"))
	}
	if rectanchoredpath.StrokeDashArray != rectanchoredpathOther.StrokeDashArray {
		diffs = append(diffs, rectanchoredpath.GongMarshallField(stage, "StrokeDashArray"))
	}
	if rectanchoredpath.StrokeDashArrayWhenSelected != rectanchoredpathOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, rectanchoredpath.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if rectanchoredpath.Transform != rectanchoredpathOther.Transform {
		diffs = append(diffs, rectanchoredpath.GongMarshallField(stage, "Transform"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (rectanchoredrect *RectAnchoredRect) GongDiff(stage *Stage, rectanchoredrectOther *RectAnchoredRect) (diffs []string) {
	// insertion point for field diffs
	if rectanchoredrect.Name != rectanchoredrectOther.Name {
		diffs = append(diffs, rectanchoredrect.GongMarshallField(stage, "Name"))
	}
	if rectanchoredrect.X != rectanchoredrectOther.X {
		diffs = append(diffs, rectanchoredrect.GongMarshallField(stage, "X"))
	}
	if rectanchoredrect.Y != rectanchoredrectOther.Y {
		diffs = append(diffs, rectanchoredrect.GongMarshallField(stage, "Y"))
	}
	if rectanchoredrect.Width != rectanchoredrectOther.Width {
		diffs = append(diffs, rectanchoredrect.GongMarshallField(stage, "Width"))
	}
	if rectanchoredrect.Height != rectanchoredrectOther.Height {
		diffs = append(diffs, rectanchoredrect.GongMarshallField(stage, "Height"))
	}
	if rectanchoredrect.RX != rectanchoredrectOther.RX {
		diffs = append(diffs, rectanchoredrect.GongMarshallField(stage, "RX"))
	}
	if rectanchoredrect.X_Offset != rectanchoredrectOther.X_Offset {
		diffs = append(diffs, rectanchoredrect.GongMarshallField(stage, "X_Offset"))
	}
	if rectanchoredrect.Y_Offset != rectanchoredrectOther.Y_Offset {
		diffs = append(diffs, rectanchoredrect.GongMarshallField(stage, "Y_Offset"))
	}
	if rectanchoredrect.RectAnchorType != rectanchoredrectOther.RectAnchorType {
		diffs = append(diffs, rectanchoredrect.GongMarshallField(stage, "RectAnchorType"))
	}
	if rectanchoredrect.WidthFollowRect != rectanchoredrectOther.WidthFollowRect {
		diffs = append(diffs, rectanchoredrect.GongMarshallField(stage, "WidthFollowRect"))
	}
	if rectanchoredrect.HeightFollowRect != rectanchoredrectOther.HeightFollowRect {
		diffs = append(diffs, rectanchoredrect.GongMarshallField(stage, "HeightFollowRect"))
	}
	if rectanchoredrect.HasToolTip != rectanchoredrectOther.HasToolTip {
		diffs = append(diffs, rectanchoredrect.GongMarshallField(stage, "HasToolTip"))
	}
	if rectanchoredrect.ToolTipText != rectanchoredrectOther.ToolTipText {
		diffs = append(diffs, rectanchoredrect.GongMarshallField(stage, "ToolTipText"))
	}
	if rectanchoredrect.Color != rectanchoredrectOther.Color {
		diffs = append(diffs, rectanchoredrect.GongMarshallField(stage, "Color"))
	}
	if rectanchoredrect.FillOpacity != rectanchoredrectOther.FillOpacity {
		diffs = append(diffs, rectanchoredrect.GongMarshallField(stage, "FillOpacity"))
	}
	if rectanchoredrect.Stroke != rectanchoredrectOther.Stroke {
		diffs = append(diffs, rectanchoredrect.GongMarshallField(stage, "Stroke"))
	}
	if rectanchoredrect.StrokeOpacity != rectanchoredrectOther.StrokeOpacity {
		diffs = append(diffs, rectanchoredrect.GongMarshallField(stage, "StrokeOpacity"))
	}
	if rectanchoredrect.StrokeWidth != rectanchoredrectOther.StrokeWidth {
		diffs = append(diffs, rectanchoredrect.GongMarshallField(stage, "StrokeWidth"))
	}
	if rectanchoredrect.StrokeDashArray != rectanchoredrectOther.StrokeDashArray {
		diffs = append(diffs, rectanchoredrect.GongMarshallField(stage, "StrokeDashArray"))
	}
	if rectanchoredrect.StrokeDashArrayWhenSelected != rectanchoredrectOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, rectanchoredrect.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if rectanchoredrect.Transform != rectanchoredrectOther.Transform {
		diffs = append(diffs, rectanchoredrect.GongMarshallField(stage, "Transform"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (rectanchoredtext *RectAnchoredText) GongDiff(stage *Stage, rectanchoredtextOther *RectAnchoredText) (diffs []string) {
	// insertion point for field diffs
	if rectanchoredtext.Name != rectanchoredtextOther.Name {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "Name"))
	}
	if rectanchoredtext.Content != rectanchoredtextOther.Content {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "Content"))
	}
	if rectanchoredtext.FontWeight != rectanchoredtextOther.FontWeight {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "FontWeight"))
	}
	if rectanchoredtext.FontSize != rectanchoredtextOther.FontSize {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "FontSize"))
	}
	if rectanchoredtext.FontStyle != rectanchoredtextOther.FontStyle {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "FontStyle"))
	}
	if rectanchoredtext.LetterSpacing != rectanchoredtextOther.LetterSpacing {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "LetterSpacing"))
	}
	if rectanchoredtext.FontFamily != rectanchoredtextOther.FontFamily {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "FontFamily"))
	}
	if rectanchoredtext.WhiteSpace != rectanchoredtextOther.WhiteSpace {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "WhiteSpace"))
	}
	if rectanchoredtext.X_Offset != rectanchoredtextOther.X_Offset {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "X_Offset"))
	}
	if rectanchoredtext.Y_Offset != rectanchoredtextOther.Y_Offset {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "Y_Offset"))
	}
	if rectanchoredtext.RectAnchorType != rectanchoredtextOther.RectAnchorType {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "RectAnchorType"))
	}
	if rectanchoredtext.TextAnchorType != rectanchoredtextOther.TextAnchorType {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "TextAnchorType"))
	}
	if rectanchoredtext.DominantBaseline != rectanchoredtextOther.DominantBaseline {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "DominantBaseline"))
	}
	if rectanchoredtext.WritingMode != rectanchoredtextOther.WritingMode {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "WritingMode"))
	}
	if rectanchoredtext.Color != rectanchoredtextOther.Color {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "Color"))
	}
	if rectanchoredtext.FillOpacity != rectanchoredtextOther.FillOpacity {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "FillOpacity"))
	}
	if rectanchoredtext.Stroke != rectanchoredtextOther.Stroke {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "Stroke"))
	}
	if rectanchoredtext.StrokeOpacity != rectanchoredtextOther.StrokeOpacity {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "StrokeOpacity"))
	}
	if rectanchoredtext.StrokeWidth != rectanchoredtextOther.StrokeWidth {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "StrokeWidth"))
	}
	if rectanchoredtext.StrokeDashArray != rectanchoredtextOther.StrokeDashArray {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "StrokeDashArray"))
	}
	if rectanchoredtext.StrokeDashArrayWhenSelected != rectanchoredtextOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if rectanchoredtext.Transform != rectanchoredtextOther.Transform {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "Transform"))
	}
	AnimatesDifferent := false
	if len(rectanchoredtext.Animates) != len(rectanchoredtextOther.Animates) {
		AnimatesDifferent = true
	} else {
		for i := range rectanchoredtext.Animates {
			if (rectanchoredtext.Animates[i] == nil) != (rectanchoredtextOther.Animates[i] == nil) {
				AnimatesDifferent = true
				break
			} else if rectanchoredtext.Animates[i] != nil && rectanchoredtextOther.Animates[i] != nil {
				// this is a pointer comparaison
				if rectanchoredtext.Animates[i] != rectanchoredtextOther.Animates[i] {
					AnimatesDifferent = true
					break
				}
			}
		}
	}
	if AnimatesDifferent {
		ops := Diff(stage, rectanchoredtext, rectanchoredtextOther, "Animates", rectanchoredtextOther.Animates, rectanchoredtext.Animates)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (rectlinklink *RectLinkLink) GongDiff(stage *Stage, rectlinklinkOther *RectLinkLink) (diffs []string) {
	// insertion point for field diffs
	if rectlinklink.Name != rectlinklinkOther.Name {
		diffs = append(diffs, rectlinklink.GongMarshallField(stage, "Name"))
	}
	if (rectlinklink.Start == nil) != (rectlinklinkOther.Start == nil) {
		diffs = append(diffs, rectlinklink.GongMarshallField(stage, "Start"))
	} else if rectlinklink.Start != nil && rectlinklinkOther.Start != nil {
		if rectlinklink.Start != rectlinklinkOther.Start {
			diffs = append(diffs, rectlinklink.GongMarshallField(stage, "Start"))
		}
	}
	if (rectlinklink.End == nil) != (rectlinklinkOther.End == nil) {
		diffs = append(diffs, rectlinklink.GongMarshallField(stage, "End"))
	} else if rectlinklink.End != nil && rectlinklinkOther.End != nil {
		if rectlinklink.End != rectlinklinkOther.End {
			diffs = append(diffs, rectlinklink.GongMarshallField(stage, "End"))
		}
	}
	if rectlinklink.TargetAnchorPosition != rectlinklinkOther.TargetAnchorPosition {
		diffs = append(diffs, rectlinklink.GongMarshallField(stage, "TargetAnchorPosition"))
	}
	if rectlinklink.Color != rectlinklinkOther.Color {
		diffs = append(diffs, rectlinklink.GongMarshallField(stage, "Color"))
	}
	if rectlinklink.FillOpacity != rectlinklinkOther.FillOpacity {
		diffs = append(diffs, rectlinklink.GongMarshallField(stage, "FillOpacity"))
	}
	if rectlinklink.Stroke != rectlinklinkOther.Stroke {
		diffs = append(diffs, rectlinklink.GongMarshallField(stage, "Stroke"))
	}
	if rectlinklink.StrokeOpacity != rectlinklinkOther.StrokeOpacity {
		diffs = append(diffs, rectlinklink.GongMarshallField(stage, "StrokeOpacity"))
	}
	if rectlinklink.StrokeWidth != rectlinklinkOther.StrokeWidth {
		diffs = append(diffs, rectlinklink.GongMarshallField(stage, "StrokeWidth"))
	}
	if rectlinklink.StrokeDashArray != rectlinklinkOther.StrokeDashArray {
		diffs = append(diffs, rectlinklink.GongMarshallField(stage, "StrokeDashArray"))
	}
	if rectlinklink.StrokeDashArrayWhenSelected != rectlinklinkOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, rectlinklink.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if rectlinklink.Transform != rectlinklinkOther.Transform {
		diffs = append(diffs, rectlinklink.GongMarshallField(stage, "Transform"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (svg *SVG) GongDiff(stage *Stage, svgOther *SVG) (diffs []string) {
	// insertion point for field diffs
	if svg.Name != svgOther.Name {
		diffs = append(diffs, svg.GongMarshallField(stage, "Name"))
	}
	LayersDifferent := false
	if len(svg.Layers) != len(svgOther.Layers) {
		LayersDifferent = true
	} else {
		for i := range svg.Layers {
			if (svg.Layers[i] == nil) != (svgOther.Layers[i] == nil) {
				LayersDifferent = true
				break
			} else if svg.Layers[i] != nil && svgOther.Layers[i] != nil {
				// this is a pointer comparaison
				if svg.Layers[i] != svgOther.Layers[i] {
					LayersDifferent = true
					break
				}
			}
		}
	}
	if LayersDifferent {
		ops := Diff(stage, svg, svgOther, "Layers", svgOther.Layers, svg.Layers)
		diffs = append(diffs, ops)
	}
	if svg.DrawingState != svgOther.DrawingState {
		diffs = append(diffs, svg.GongMarshallField(stage, "DrawingState"))
	}
	if (svg.StartRect == nil) != (svgOther.StartRect == nil) {
		diffs = append(diffs, svg.GongMarshallField(stage, "StartRect"))
	} else if svg.StartRect != nil && svgOther.StartRect != nil {
		if svg.StartRect != svgOther.StartRect {
			diffs = append(diffs, svg.GongMarshallField(stage, "StartRect"))
		}
	}
	if (svg.EndRect == nil) != (svgOther.EndRect == nil) {
		diffs = append(diffs, svg.GongMarshallField(stage, "EndRect"))
	} else if svg.EndRect != nil && svgOther.EndRect != nil {
		if svg.EndRect != svgOther.EndRect {
			diffs = append(diffs, svg.GongMarshallField(stage, "EndRect"))
		}
	}
	if svg.IsEditable != svgOther.IsEditable {
		diffs = append(diffs, svg.GongMarshallField(stage, "IsEditable"))
	}
	if svg.IsSVGFrontEndFileGenerated != svgOther.IsSVGFrontEndFileGenerated {
		diffs = append(diffs, svg.GongMarshallField(stage, "IsSVGFrontEndFileGenerated"))
	}
	if svg.IsSVGBackEndFileGenerated != svgOther.IsSVGBackEndFileGenerated {
		diffs = append(diffs, svg.GongMarshallField(stage, "IsSVGBackEndFileGenerated"))
	}
	if svg.DefaultDirectoryForGeneratedImages != svgOther.DefaultDirectoryForGeneratedImages {
		diffs = append(diffs, svg.GongMarshallField(stage, "DefaultDirectoryForGeneratedImages"))
	}
	if svg.IsControlBannerHidden != svgOther.IsControlBannerHidden {
		diffs = append(diffs, svg.GongMarshallField(stage, "IsControlBannerHidden"))
	}
	if svg.OverrideWidth != svgOther.OverrideWidth {
		diffs = append(diffs, svg.GongMarshallField(stage, "OverrideWidth"))
	}
	if svg.OverriddenWidth != svgOther.OverriddenWidth {
		diffs = append(diffs, svg.GongMarshallField(stage, "OverriddenWidth"))
	}
	if svg.OverrideHeight != svgOther.OverrideHeight {
		diffs = append(diffs, svg.GongMarshallField(stage, "OverrideHeight"))
	}
	if svg.OverriddenHeight != svgOther.OverriddenHeight {
		diffs = append(diffs, svg.GongMarshallField(stage, "OverriddenHeight"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (svgtext *SvgText) GongDiff(stage *Stage, svgtextOther *SvgText) (diffs []string) {
	// insertion point for field diffs
	if svgtext.Name != svgtextOther.Name {
		diffs = append(diffs, svgtext.GongMarshallField(stage, "Name"))
	}
	if svgtext.Text != svgtextOther.Text {
		diffs = append(diffs, svgtext.GongMarshallField(stage, "Text"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (text *Text) GongDiff(stage *Stage, textOther *Text) (diffs []string) {
	// insertion point for field diffs
	if text.Name != textOther.Name {
		diffs = append(diffs, text.GongMarshallField(stage, "Name"))
	}
	if text.X != textOther.X {
		diffs = append(diffs, text.GongMarshallField(stage, "X"))
	}
	if text.Y != textOther.Y {
		diffs = append(diffs, text.GongMarshallField(stage, "Y"))
	}
	if text.Content != textOther.Content {
		diffs = append(diffs, text.GongMarshallField(stage, "Content"))
	}
	if text.Color != textOther.Color {
		diffs = append(diffs, text.GongMarshallField(stage, "Color"))
	}
	if text.FillOpacity != textOther.FillOpacity {
		diffs = append(diffs, text.GongMarshallField(stage, "FillOpacity"))
	}
	if text.Stroke != textOther.Stroke {
		diffs = append(diffs, text.GongMarshallField(stage, "Stroke"))
	}
	if text.StrokeOpacity != textOther.StrokeOpacity {
		diffs = append(diffs, text.GongMarshallField(stage, "StrokeOpacity"))
	}
	if text.StrokeWidth != textOther.StrokeWidth {
		diffs = append(diffs, text.GongMarshallField(stage, "StrokeWidth"))
	}
	if text.StrokeDashArray != textOther.StrokeDashArray {
		diffs = append(diffs, text.GongMarshallField(stage, "StrokeDashArray"))
	}
	if text.StrokeDashArrayWhenSelected != textOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, text.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if text.Transform != textOther.Transform {
		diffs = append(diffs, text.GongMarshallField(stage, "Transform"))
	}
	if text.FontWeight != textOther.FontWeight {
		diffs = append(diffs, text.GongMarshallField(stage, "FontWeight"))
	}
	if text.FontSize != textOther.FontSize {
		diffs = append(diffs, text.GongMarshallField(stage, "FontSize"))
	}
	if text.FontStyle != textOther.FontStyle {
		diffs = append(diffs, text.GongMarshallField(stage, "FontStyle"))
	}
	if text.LetterSpacing != textOther.LetterSpacing {
		diffs = append(diffs, text.GongMarshallField(stage, "LetterSpacing"))
	}
	if text.FontFamily != textOther.FontFamily {
		diffs = append(diffs, text.GongMarshallField(stage, "FontFamily"))
	}
	if text.WhiteSpace != textOther.WhiteSpace {
		diffs = append(diffs, text.GongMarshallField(stage, "WhiteSpace"))
	}
	AnimatesDifferent := false
	if len(text.Animates) != len(textOther.Animates) {
		AnimatesDifferent = true
	} else {
		for i := range text.Animates {
			if (text.Animates[i] == nil) != (textOther.Animates[i] == nil) {
				AnimatesDifferent = true
				break
			} else if text.Animates[i] != nil && textOther.Animates[i] != nil {
				// this is a pointer comparaison
				if text.Animates[i] != textOther.Animates[i] {
					AnimatesDifferent = true
					break
				}
			}
		}
	}
	if AnimatesDifferent {
		ops := Diff(stage, text, textOther, "Animates", textOther.Animates, text.Animates)
		diffs = append(diffs, ops)
	}

	return
}

// Diff returns the sequence of operations to transform oldSlice into newSlice.
// It requires type T to be comparable (e.g., pointers, ints, strings).
func Diff[T1, T2 PointerToGongstruct](stage *Stage, a, b T1, fieldName string, oldSlice, newSlice []T2) (ops string) {
	m, n := len(oldSlice), len(newSlice)

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if oldSlice[i] == newSlice[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if oldSlice[i-1] == newSlice[j-1] {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Create a temporary view of what's left after deletions for tracking matches
	var currentLCS []T2
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			currentLCS = append(currentLCS, oldSlice[k])
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k, targetVal := range newSlice {
		if lcsIdx < len(currentLCS) && currentLCS[lcsIdx] == targetVal {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, targetVal.GongGetIdentifier(stage))
		}
	}

	return ops
}
