// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Animate:
		ok = stage.IsStagedAnimate(target)

	case *Circle:
		ok = stage.IsStagedCircle(target)

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

	case *Text:
		ok = stage.IsStagedText(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedAnimate(animate *Animate) (ok bool) {

	_, ok = stage.Animates[animate]

	return
}

func (stage *StageStruct) IsStagedCircle(circle *Circle) (ok bool) {

	_, ok = stage.Circles[circle]

	return
}

func (stage *StageStruct) IsStagedEllipse(ellipse *Ellipse) (ok bool) {

	_, ok = stage.Ellipses[ellipse]

	return
}

func (stage *StageStruct) IsStagedLayer(layer *Layer) (ok bool) {

	_, ok = stage.Layers[layer]

	return
}

func (stage *StageStruct) IsStagedLine(line *Line) (ok bool) {

	_, ok = stage.Lines[line]

	return
}

func (stage *StageStruct) IsStagedLink(link *Link) (ok bool) {

	_, ok = stage.Links[link]

	return
}

func (stage *StageStruct) IsStagedLinkAnchoredText(linkanchoredtext *LinkAnchoredText) (ok bool) {

	_, ok = stage.LinkAnchoredTexts[linkanchoredtext]

	return
}

func (stage *StageStruct) IsStagedPath(path *Path) (ok bool) {

	_, ok = stage.Paths[path]

	return
}

func (stage *StageStruct) IsStagedPoint(point *Point) (ok bool) {

	_, ok = stage.Points[point]

	return
}

func (stage *StageStruct) IsStagedPolygone(polygone *Polygone) (ok bool) {

	_, ok = stage.Polygones[polygone]

	return
}

func (stage *StageStruct) IsStagedPolyline(polyline *Polyline) (ok bool) {

	_, ok = stage.Polylines[polyline]

	return
}

func (stage *StageStruct) IsStagedRect(rect *Rect) (ok bool) {

	_, ok = stage.Rects[rect]

	return
}

func (stage *StageStruct) IsStagedRectAnchoredPath(rectanchoredpath *RectAnchoredPath) (ok bool) {

	_, ok = stage.RectAnchoredPaths[rectanchoredpath]

	return
}

func (stage *StageStruct) IsStagedRectAnchoredRect(rectanchoredrect *RectAnchoredRect) (ok bool) {

	_, ok = stage.RectAnchoredRects[rectanchoredrect]

	return
}

func (stage *StageStruct) IsStagedRectAnchoredText(rectanchoredtext *RectAnchoredText) (ok bool) {

	_, ok = stage.RectAnchoredTexts[rectanchoredtext]

	return
}

func (stage *StageStruct) IsStagedRectLinkLink(rectlinklink *RectLinkLink) (ok bool) {

	_, ok = stage.RectLinkLinks[rectlinklink]

	return
}

func (stage *StageStruct) IsStagedSVG(svg *SVG) (ok bool) {

	_, ok = stage.SVGs[svg]

	return
}

func (stage *StageStruct) IsStagedText(text *Text) (ok bool) {

	_, ok = stage.Texts[text]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Animate:
		stage.StageBranchAnimate(target)

	case *Circle:
		stage.StageBranchCircle(target)

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

	case *Text:
		stage.StageBranchText(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchAnimate(animate *Animate) {

	// check if instance is already staged
	if IsStaged(stage, animate) {
		return
	}

	animate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchCircle(circle *Circle) {

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

func (stage *StageStruct) StageBranchEllipse(ellipse *Ellipse) {

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

func (stage *StageStruct) StageBranchLayer(layer *Layer) {

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

func (stage *StageStruct) StageBranchLine(line *Line) {

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

func (stage *StageStruct) StageBranchLink(link *Link) {

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
	for _, _linkanchoredtext := range link.TextAtArrowEnd {
		StageBranch(stage, _linkanchoredtext)
	}
	for _, _linkanchoredtext := range link.TextAtArrowStart {
		StageBranch(stage, _linkanchoredtext)
	}
	for _, _point := range link.ControlPoints {
		StageBranch(stage, _point)
	}

}

func (stage *StageStruct) StageBranchLinkAnchoredText(linkanchoredtext *LinkAnchoredText) {

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

func (stage *StageStruct) StageBranchPath(path *Path) {

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

func (stage *StageStruct) StageBranchPoint(point *Point) {

	// check if instance is already staged
	if IsStaged(stage, point) {
		return
	}

	point.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPolygone(polygone *Polygone) {

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

func (stage *StageStruct) StageBranchPolyline(polyline *Polyline) {

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

func (stage *StageStruct) StageBranchRect(rect *Rect) {

	// check if instance is already staged
	if IsStaged(stage, rect) {
		return
	}

	rect.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
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

func (stage *StageStruct) StageBranchRectAnchoredPath(rectanchoredpath *RectAnchoredPath) {

	// check if instance is already staged
	if IsStaged(stage, rectanchoredpath) {
		return
	}

	rectanchoredpath.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchRectAnchoredRect(rectanchoredrect *RectAnchoredRect) {

	// check if instance is already staged
	if IsStaged(stage, rectanchoredrect) {
		return
	}

	rectanchoredrect.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchRectAnchoredText(rectanchoredtext *RectAnchoredText) {

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

func (stage *StageStruct) StageBranchRectLinkLink(rectlinklink *RectLinkLink) {

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

func (stage *StageStruct) StageBranchSVG(svg *SVG) {

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

func (stage *StageStruct) StageBranchText(text *Text) {

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
	for _, _linkanchoredtext := range linkFrom.TextAtArrowEnd {
		linkTo.TextAtArrowEnd = append(linkTo.TextAtArrowEnd, CopyBranchLinkAnchoredText(mapOrigCopy, _linkanchoredtext))
	}
	for _, _linkanchoredtext := range linkFrom.TextAtArrowStart {
		linkTo.TextAtArrowStart = append(linkTo.TextAtArrowStart, CopyBranchLinkAnchoredText(mapOrigCopy, _linkanchoredtext))
	}
	for _, _point := range linkFrom.ControlPoints {
		linkTo.ControlPoints = append(linkTo.ControlPoints, CopyBranchPoint(mapOrigCopy, _point))
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
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Animate:
		stage.UnstageBranchAnimate(target)

	case *Circle:
		stage.UnstageBranchCircle(target)

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

	case *Text:
		stage.UnstageBranchText(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchAnimate(animate *Animate) {

	// check if instance is already staged
	if !IsStaged(stage, animate) {
		return
	}

	animate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchCircle(circle *Circle) {

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

func (stage *StageStruct) UnstageBranchEllipse(ellipse *Ellipse) {

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

func (stage *StageStruct) UnstageBranchLayer(layer *Layer) {

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

func (stage *StageStruct) UnstageBranchLine(line *Line) {

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

func (stage *StageStruct) UnstageBranchLink(link *Link) {

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
	for _, _linkanchoredtext := range link.TextAtArrowEnd {
		UnstageBranch(stage, _linkanchoredtext)
	}
	for _, _linkanchoredtext := range link.TextAtArrowStart {
		UnstageBranch(stage, _linkanchoredtext)
	}
	for _, _point := range link.ControlPoints {
		UnstageBranch(stage, _point)
	}

}

func (stage *StageStruct) UnstageBranchLinkAnchoredText(linkanchoredtext *LinkAnchoredText) {

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

func (stage *StageStruct) UnstageBranchPath(path *Path) {

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

func (stage *StageStruct) UnstageBranchPoint(point *Point) {

	// check if instance is already staged
	if !IsStaged(stage, point) {
		return
	}

	point.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPolygone(polygone *Polygone) {

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

func (stage *StageStruct) UnstageBranchPolyline(polyline *Polyline) {

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

func (stage *StageStruct) UnstageBranchRect(rect *Rect) {

	// check if instance is already staged
	if !IsStaged(stage, rect) {
		return
	}

	rect.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
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

func (stage *StageStruct) UnstageBranchRectAnchoredPath(rectanchoredpath *RectAnchoredPath) {

	// check if instance is already staged
	if !IsStaged(stage, rectanchoredpath) {
		return
	}

	rectanchoredpath.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchRectAnchoredRect(rectanchoredrect *RectAnchoredRect) {

	// check if instance is already staged
	if !IsStaged(stage, rectanchoredrect) {
		return
	}

	rectanchoredrect.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchRectAnchoredText(rectanchoredtext *RectAnchoredText) {

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

func (stage *StageStruct) UnstageBranchRectLinkLink(rectlinklink *RectLinkLink) {

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

func (stage *StageStruct) UnstageBranchSVG(svg *SVG) {

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

func (stage *StageStruct) UnstageBranchText(text *Text) {

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
