// generated code - do not edit
package models

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (animate *Animate) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Animates[animate]
	return ok
}

func (circle *Circle) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Circles[circle]
	return ok
}

func (condition *Condition) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Conditions[condition]
	return ok
}

func (controlpoint *ControlPoint) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ControlPoints[controlpoint]
	return ok
}

func (ellipse *Ellipse) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Ellipses[ellipse]
	return ok
}

func (filetodownload *FileToDownload) GongIsStaged(stage *Stage) bool {
	_, ok := stage.FileToDownloads[filetodownload]
	return ok
}

func (layer *Layer) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Layers[layer]
	return ok
}

func (line *Line) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Lines[line]
	return ok
}

func (link *Link) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Links[link]
	return ok
}

func (linkanchoredpath *LinkAnchoredPath) GongIsStaged(stage *Stage) bool {
	_, ok := stage.LinkAnchoredPaths[linkanchoredpath]
	return ok
}

func (linkanchoredtext *LinkAnchoredText) GongIsStaged(stage *Stage) bool {
	_, ok := stage.LinkAnchoredTexts[linkanchoredtext]
	return ok
}

func (path *Path) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Paths[path]
	return ok
}

func (point *Point) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Points[point]
	return ok
}

func (polygone *Polygone) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Polygones[polygone]
	return ok
}

func (polyline *Polyline) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Polylines[polyline]
	return ok
}

func (rect *Rect) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Rects[rect]
	return ok
}

func (rectanchoredpath *RectAnchoredPath) GongIsStaged(stage *Stage) bool {
	_, ok := stage.RectAnchoredPaths[rectanchoredpath]
	return ok
}

func (rectanchoredpngimage *RectAnchoredPngImage) GongIsStaged(stage *Stage) bool {
	_, ok := stage.RectAnchoredPngImages[rectanchoredpngimage]
	return ok
}

func (rectanchoredrect *RectAnchoredRect) GongIsStaged(stage *Stage) bool {
	_, ok := stage.RectAnchoredRects[rectanchoredrect]
	return ok
}

func (rectanchoredtext *RectAnchoredText) GongIsStaged(stage *Stage) bool {
	_, ok := stage.RectAnchoredTexts[rectanchoredtext]
	return ok
}

func (rectlinklink *RectLinkLink) GongIsStaged(stage *Stage) bool {
	_, ok := stage.RectLinkLinks[rectlinklink]
	return ok
}

func (svg *SVG) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SVGs[svg]
	return ok
}

func (svgtext *SvgText) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SvgTexts[svgtext]
	return ok
}

func (text *Text) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Texts[text]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (animate *Animate) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(animate) {
		return
	}

	animate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (circle *Circle) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(circle) {
		return
	}

	circle.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range circle.Animations {
		stage.StageBranch(_animate)
	}

}

func (condition *Condition) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(condition) {
		return
	}

	condition.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (controlpoint *ControlPoint) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(controlpoint) {
		return
	}

	controlpoint.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if controlpoint.ClosestRect != nil {
		stage.StageBranch(controlpoint.ClosestRect)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ellipse *Ellipse) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ellipse) {
		return
	}

	ellipse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range ellipse.Animates {
		stage.StageBranch(_animate)
	}

}

func (filetodownload *FileToDownload) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(filetodownload) {
		return
	}

	filetodownload.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (layer *Layer) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(layer) {
		return
	}

	layer.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rect := range layer.Rects {
		stage.StageBranch(_rect)
	}
	for _, _text := range layer.Texts {
		stage.StageBranch(_text)
	}
	for _, _circle := range layer.Circles {
		stage.StageBranch(_circle)
	}
	for _, _line := range layer.Lines {
		stage.StageBranch(_line)
	}
	for _, _ellipse := range layer.Ellipses {
		stage.StageBranch(_ellipse)
	}
	for _, _polyline := range layer.Polylines {
		stage.StageBranch(_polyline)
	}
	for _, _polygone := range layer.Polygones {
		stage.StageBranch(_polygone)
	}
	for _, _path := range layer.Paths {
		stage.StageBranch(_path)
	}
	for _, _link := range layer.Links {
		stage.StageBranch(_link)
	}
	for _, _rectlinklink := range layer.RectLinkLinks {
		stage.StageBranch(_rectlinklink)
	}

}

func (line *Line) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(line) {
		return
	}

	line.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range line.Animates {
		stage.StageBranch(_animate)
	}

}

func (link *Link) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(link) {
		return
	}

	link.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if link.Start != nil {
		stage.StageBranch(link.Start)
	}
	if link.End != nil {
		stage.StageBranch(link.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _linkanchoredtext := range link.TextAtArrowStart {
		stage.StageBranch(_linkanchoredtext)
	}
	for _, _linkanchoredtext := range link.TextAtArrowEnd {
		stage.StageBranch(_linkanchoredtext)
	}
	for _, _linkanchoredtext := range link.TextAtCorner {
		stage.StageBranch(_linkanchoredtext)
	}
	for _, _linkanchoredpath := range link.PathAtArrowStart {
		stage.StageBranch(_linkanchoredpath)
	}
	for _, _linkanchoredpath := range link.PathAtArrowEnd {
		stage.StageBranch(_linkanchoredpath)
	}
	for _, _linkanchoredpath := range link.PathAtCorner {
		stage.StageBranch(_linkanchoredpath)
	}
	for _, _controlpoint := range link.ControlPoints {
		stage.StageBranch(_controlpoint)
	}

}

func (linkanchoredpath *LinkAnchoredPath) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(linkanchoredpath) {
		return
	}

	linkanchoredpath.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (linkanchoredtext *LinkAnchoredText) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(linkanchoredtext) {
		return
	}

	linkanchoredtext.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range linkanchoredtext.Animates {
		stage.StageBranch(_animate)
	}

}

func (path *Path) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(path) {
		return
	}

	path.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range path.Animates {
		stage.StageBranch(_animate)
	}

}

func (point *Point) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(point) {
		return
	}

	point.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (polygone *Polygone) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(polygone) {
		return
	}

	polygone.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range polygone.Animates {
		stage.StageBranch(_animate)
	}

}

func (polyline *Polyline) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(polyline) {
		return
	}

	polyline.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range polyline.Animates {
		stage.StageBranch(_animate)
	}

}

func (rect *Rect) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(rect) {
		return
	}

	rect.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if rect.EnclosingRect != nil {
		stage.StageBranch(rect.EnclosingRect)
	}
	if rect.AnchoredTo != nil {
		stage.StageBranch(rect.AnchoredTo)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rect := range rect.Peers {
		stage.StageBranch(_rect)
	}
	for _, _rect := range rect.Obstacles {
		stage.StageBranch(_rect)
	}
	for _, _condition := range rect.HoveringTrigger {
		stage.StageBranch(_condition)
	}
	for _, _condition := range rect.DisplayConditions {
		stage.StageBranch(_condition)
	}
	for _, _animate := range rect.Animations {
		stage.StageBranch(_animate)
	}
	for _, _rectanchoredtext := range rect.RectAnchoredTexts {
		stage.StageBranch(_rectanchoredtext)
	}
	for _, _rectanchoredrect := range rect.RectAnchoredRects {
		stage.StageBranch(_rectanchoredrect)
	}
	for _, _rectanchoredpath := range rect.RectAnchoredPaths {
		stage.StageBranch(_rectanchoredpath)
	}
	for _, _rectanchoredpngimage := range rect.RectAnchoredPngImages {
		stage.StageBranch(_rectanchoredpngimage)
	}

}

func (rectanchoredpath *RectAnchoredPath) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(rectanchoredpath) {
		return
	}

	rectanchoredpath.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rectanchoredpngimage *RectAnchoredPngImage) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(rectanchoredpngimage) {
		return
	}

	rectanchoredpngimage.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rectanchoredrect *RectAnchoredRect) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(rectanchoredrect) {
		return
	}

	rectanchoredrect.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rectanchoredtext *RectAnchoredText) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(rectanchoredtext) {
		return
	}

	rectanchoredtext.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range rectanchoredtext.Animates {
		stage.StageBranch(_animate)
	}

}

func (rectlinklink *RectLinkLink) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(rectlinklink) {
		return
	}

	rectlinklink.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if rectlinklink.Start != nil {
		stage.StageBranch(rectlinklink.Start)
	}
	if rectlinklink.End != nil {
		stage.StageBranch(rectlinklink.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (svg *SVG) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(svg) {
		return
	}

	svg.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if svg.StartRect != nil {
		stage.StageBranch(svg.StartRect)
	}
	if svg.EndRect != nil {
		stage.StageBranch(svg.EndRect)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _layer := range svg.Layers {
		stage.StageBranch(_layer)
	}

}

func (svgtext *SvgText) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(svgtext) {
		return
	}

	svgtext.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (text *Text) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(text) {
		return
	}

	text.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range text.Animates {
		stage.StageBranch(_animate)
	}

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Animate:
		toT := GongCopyBranchAnimate(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Circle:
		toT := GongCopyBranchCircle(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Condition:
		toT := GongCopyBranchCondition(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ControlPoint:
		toT := GongCopyBranchControlPoint(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Ellipse:
		toT := GongCopyBranchEllipse(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FileToDownload:
		toT := GongCopyBranchFileToDownload(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Layer:
		toT := GongCopyBranchLayer(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Line:
		toT := GongCopyBranchLine(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Link:
		toT := GongCopyBranchLink(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LinkAnchoredPath:
		toT := GongCopyBranchLinkAnchoredPath(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LinkAnchoredText:
		toT := GongCopyBranchLinkAnchoredText(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Path:
		toT := GongCopyBranchPath(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Point:
		toT := GongCopyBranchPoint(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Polygone:
		toT := GongCopyBranchPolygone(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Polyline:
		toT := GongCopyBranchPolyline(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Rect:
		toT := GongCopyBranchRect(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RectAnchoredPath:
		toT := GongCopyBranchRectAnchoredPath(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RectAnchoredPngImage:
		toT := GongCopyBranchRectAnchoredPngImage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RectAnchoredRect:
		toT := GongCopyBranchRectAnchoredRect(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RectAnchoredText:
		toT := GongCopyBranchRectAnchoredText(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RectLinkLink:
		toT := GongCopyBranchRectLinkLink(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SVG:
		toT := GongCopyBranchSVG(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SvgText:
		toT := GongCopyBranchSvgText(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Text:
		toT := GongCopyBranchText(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAnimate(mapOrigCopy map[any]any, animateFrom *Animate) (animateTo *Animate) {
	var alreadyCopied bool
	animateTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, animateFrom)
	if alreadyCopied {
		return
	}
	animateFrom.GongCopyBasicFields(animateTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCircle(mapOrigCopy map[any]any, circleFrom *Circle) (circleTo *Circle) {
	var alreadyCopied bool
	circleTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, circleFrom)
	if alreadyCopied {
		return
	}
	circleFrom.GongCopyBasicFields(circleTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range circleFrom.Animations {
		circleTo.Animations = append(circleTo.Animations, GongCopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func GongCopyBranchCondition(mapOrigCopy map[any]any, conditionFrom *Condition) (conditionTo *Condition) {
	var alreadyCopied bool
	conditionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, conditionFrom)
	if alreadyCopied {
		return
	}
	conditionFrom.GongCopyBasicFields(conditionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchControlPoint(mapOrigCopy map[any]any, controlpointFrom *ControlPoint) (controlpointTo *ControlPoint) {
	var alreadyCopied bool
	controlpointTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, controlpointFrom)
	if alreadyCopied {
		return
	}
	controlpointFrom.GongCopyBasicFields(controlpointTo)

	//insertion point for the staging of instances referenced by pointers
	if controlpointFrom.ClosestRect != nil {
		controlpointTo.ClosestRect = GongCopyBranchRect(mapOrigCopy, controlpointFrom.ClosestRect)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEllipse(mapOrigCopy map[any]any, ellipseFrom *Ellipse) (ellipseTo *Ellipse) {
	var alreadyCopied bool
	ellipseTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, ellipseFrom)
	if alreadyCopied {
		return
	}
	ellipseFrom.GongCopyBasicFields(ellipseTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range ellipseFrom.Animates {
		ellipseTo.Animates = append(ellipseTo.Animates, GongCopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func GongCopyBranchFileToDownload(mapOrigCopy map[any]any, filetodownloadFrom *FileToDownload) (filetodownloadTo *FileToDownload) {
	var alreadyCopied bool
	filetodownloadTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, filetodownloadFrom)
	if alreadyCopied {
		return
	}
	filetodownloadFrom.GongCopyBasicFields(filetodownloadTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLayer(mapOrigCopy map[any]any, layerFrom *Layer) (layerTo *Layer) {
	var alreadyCopied bool
	layerTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, layerFrom)
	if alreadyCopied {
		return
	}
	layerFrom.GongCopyBasicFields(layerTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rect := range layerFrom.Rects {
		layerTo.Rects = append(layerTo.Rects, GongCopyBranchRect(mapOrigCopy, _rect))
	}
	for _, _text := range layerFrom.Texts {
		layerTo.Texts = append(layerTo.Texts, GongCopyBranchText(mapOrigCopy, _text))
	}
	for _, _circle := range layerFrom.Circles {
		layerTo.Circles = append(layerTo.Circles, GongCopyBranchCircle(mapOrigCopy, _circle))
	}
	for _, _line := range layerFrom.Lines {
		layerTo.Lines = append(layerTo.Lines, GongCopyBranchLine(mapOrigCopy, _line))
	}
	for _, _ellipse := range layerFrom.Ellipses {
		layerTo.Ellipses = append(layerTo.Ellipses, GongCopyBranchEllipse(mapOrigCopy, _ellipse))
	}
	for _, _polyline := range layerFrom.Polylines {
		layerTo.Polylines = append(layerTo.Polylines, GongCopyBranchPolyline(mapOrigCopy, _polyline))
	}
	for _, _polygone := range layerFrom.Polygones {
		layerTo.Polygones = append(layerTo.Polygones, GongCopyBranchPolygone(mapOrigCopy, _polygone))
	}
	for _, _path := range layerFrom.Paths {
		layerTo.Paths = append(layerTo.Paths, GongCopyBranchPath(mapOrigCopy, _path))
	}
	for _, _link := range layerFrom.Links {
		layerTo.Links = append(layerTo.Links, GongCopyBranchLink(mapOrigCopy, _link))
	}
	for _, _rectlinklink := range layerFrom.RectLinkLinks {
		layerTo.RectLinkLinks = append(layerTo.RectLinkLinks, GongCopyBranchRectLinkLink(mapOrigCopy, _rectlinklink))
	}

	return
}

func GongCopyBranchLine(mapOrigCopy map[any]any, lineFrom *Line) (lineTo *Line) {
	var alreadyCopied bool
	lineTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, lineFrom)
	if alreadyCopied {
		return
	}
	lineFrom.GongCopyBasicFields(lineTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range lineFrom.Animates {
		lineTo.Animates = append(lineTo.Animates, GongCopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func GongCopyBranchLink(mapOrigCopy map[any]any, linkFrom *Link) (linkTo *Link) {
	var alreadyCopied bool
	linkTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, linkFrom)
	if alreadyCopied {
		return
	}
	linkFrom.GongCopyBasicFields(linkTo)

	//insertion point for the staging of instances referenced by pointers
	if linkFrom.Start != nil {
		linkTo.Start = GongCopyBranchRect(mapOrigCopy, linkFrom.Start)
	}
	if linkFrom.End != nil {
		linkTo.End = GongCopyBranchRect(mapOrigCopy, linkFrom.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _linkanchoredtext := range linkFrom.TextAtArrowStart {
		linkTo.TextAtArrowStart = append(linkTo.TextAtArrowStart, GongCopyBranchLinkAnchoredText(mapOrigCopy, _linkanchoredtext))
	}
	for _, _linkanchoredtext := range linkFrom.TextAtArrowEnd {
		linkTo.TextAtArrowEnd = append(linkTo.TextAtArrowEnd, GongCopyBranchLinkAnchoredText(mapOrigCopy, _linkanchoredtext))
	}
	for _, _linkanchoredtext := range linkFrom.TextAtCorner {
		linkTo.TextAtCorner = append(linkTo.TextAtCorner, GongCopyBranchLinkAnchoredText(mapOrigCopy, _linkanchoredtext))
	}
	for _, _linkanchoredpath := range linkFrom.PathAtArrowStart {
		linkTo.PathAtArrowStart = append(linkTo.PathAtArrowStart, GongCopyBranchLinkAnchoredPath(mapOrigCopy, _linkanchoredpath))
	}
	for _, _linkanchoredpath := range linkFrom.PathAtArrowEnd {
		linkTo.PathAtArrowEnd = append(linkTo.PathAtArrowEnd, GongCopyBranchLinkAnchoredPath(mapOrigCopy, _linkanchoredpath))
	}
	for _, _linkanchoredpath := range linkFrom.PathAtCorner {
		linkTo.PathAtCorner = append(linkTo.PathAtCorner, GongCopyBranchLinkAnchoredPath(mapOrigCopy, _linkanchoredpath))
	}
	for _, _controlpoint := range linkFrom.ControlPoints {
		linkTo.ControlPoints = append(linkTo.ControlPoints, GongCopyBranchControlPoint(mapOrigCopy, _controlpoint))
	}

	return
}

func GongCopyBranchLinkAnchoredPath(mapOrigCopy map[any]any, linkanchoredpathFrom *LinkAnchoredPath) (linkanchoredpathTo *LinkAnchoredPath) {
	var alreadyCopied bool
	linkanchoredpathTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, linkanchoredpathFrom)
	if alreadyCopied {
		return
	}
	linkanchoredpathFrom.GongCopyBasicFields(linkanchoredpathTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLinkAnchoredText(mapOrigCopy map[any]any, linkanchoredtextFrom *LinkAnchoredText) (linkanchoredtextTo *LinkAnchoredText) {
	var alreadyCopied bool
	linkanchoredtextTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, linkanchoredtextFrom)
	if alreadyCopied {
		return
	}
	linkanchoredtextFrom.GongCopyBasicFields(linkanchoredtextTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range linkanchoredtextFrom.Animates {
		linkanchoredtextTo.Animates = append(linkanchoredtextTo.Animates, GongCopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func GongCopyBranchPath(mapOrigCopy map[any]any, pathFrom *Path) (pathTo *Path) {
	var alreadyCopied bool
	pathTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, pathFrom)
	if alreadyCopied {
		return
	}
	pathFrom.GongCopyBasicFields(pathTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range pathFrom.Animates {
		pathTo.Animates = append(pathTo.Animates, GongCopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func GongCopyBranchPoint(mapOrigCopy map[any]any, pointFrom *Point) (pointTo *Point) {
	var alreadyCopied bool
	pointTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, pointFrom)
	if alreadyCopied {
		return
	}
	pointFrom.GongCopyBasicFields(pointTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPolygone(mapOrigCopy map[any]any, polygoneFrom *Polygone) (polygoneTo *Polygone) {
	var alreadyCopied bool
	polygoneTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, polygoneFrom)
	if alreadyCopied {
		return
	}
	polygoneFrom.GongCopyBasicFields(polygoneTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range polygoneFrom.Animates {
		polygoneTo.Animates = append(polygoneTo.Animates, GongCopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func GongCopyBranchPolyline(mapOrigCopy map[any]any, polylineFrom *Polyline) (polylineTo *Polyline) {
	var alreadyCopied bool
	polylineTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, polylineFrom)
	if alreadyCopied {
		return
	}
	polylineFrom.GongCopyBasicFields(polylineTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range polylineFrom.Animates {
		polylineTo.Animates = append(polylineTo.Animates, GongCopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func GongCopyBranchRect(mapOrigCopy map[any]any, rectFrom *Rect) (rectTo *Rect) {
	var alreadyCopied bool
	rectTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, rectFrom)
	if alreadyCopied {
		return
	}
	rectFrom.GongCopyBasicFields(rectTo)

	//insertion point for the staging of instances referenced by pointers
	if rectFrom.EnclosingRect != nil {
		rectTo.EnclosingRect = GongCopyBranchRect(mapOrigCopy, rectFrom.EnclosingRect)
	}
	if rectFrom.AnchoredTo != nil {
		rectTo.AnchoredTo = GongCopyBranchRect(mapOrigCopy, rectFrom.AnchoredTo)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rect := range rectFrom.Peers {
		rectTo.Peers = append(rectTo.Peers, GongCopyBranchRect(mapOrigCopy, _rect))
	}
	for _, _rect := range rectFrom.Obstacles {
		rectTo.Obstacles = append(rectTo.Obstacles, GongCopyBranchRect(mapOrigCopy, _rect))
	}
	for _, _condition := range rectFrom.HoveringTrigger {
		rectTo.HoveringTrigger = append(rectTo.HoveringTrigger, GongCopyBranchCondition(mapOrigCopy, _condition))
	}
	for _, _condition := range rectFrom.DisplayConditions {
		rectTo.DisplayConditions = append(rectTo.DisplayConditions, GongCopyBranchCondition(mapOrigCopy, _condition))
	}
	for _, _animate := range rectFrom.Animations {
		rectTo.Animations = append(rectTo.Animations, GongCopyBranchAnimate(mapOrigCopy, _animate))
	}
	for _, _rectanchoredtext := range rectFrom.RectAnchoredTexts {
		rectTo.RectAnchoredTexts = append(rectTo.RectAnchoredTexts, GongCopyBranchRectAnchoredText(mapOrigCopy, _rectanchoredtext))
	}
	for _, _rectanchoredrect := range rectFrom.RectAnchoredRects {
		rectTo.RectAnchoredRects = append(rectTo.RectAnchoredRects, GongCopyBranchRectAnchoredRect(mapOrigCopy, _rectanchoredrect))
	}
	for _, _rectanchoredpath := range rectFrom.RectAnchoredPaths {
		rectTo.RectAnchoredPaths = append(rectTo.RectAnchoredPaths, GongCopyBranchRectAnchoredPath(mapOrigCopy, _rectanchoredpath))
	}
	for _, _rectanchoredpngimage := range rectFrom.RectAnchoredPngImages {
		rectTo.RectAnchoredPngImages = append(rectTo.RectAnchoredPngImages, GongCopyBranchRectAnchoredPngImage(mapOrigCopy, _rectanchoredpngimage))
	}

	return
}

func GongCopyBranchRectAnchoredPath(mapOrigCopy map[any]any, rectanchoredpathFrom *RectAnchoredPath) (rectanchoredpathTo *RectAnchoredPath) {
	var alreadyCopied bool
	rectanchoredpathTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, rectanchoredpathFrom)
	if alreadyCopied {
		return
	}
	rectanchoredpathFrom.GongCopyBasicFields(rectanchoredpathTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRectAnchoredPngImage(mapOrigCopy map[any]any, rectanchoredpngimageFrom *RectAnchoredPngImage) (rectanchoredpngimageTo *RectAnchoredPngImage) {
	var alreadyCopied bool
	rectanchoredpngimageTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, rectanchoredpngimageFrom)
	if alreadyCopied {
		return
	}
	rectanchoredpngimageFrom.GongCopyBasicFields(rectanchoredpngimageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRectAnchoredRect(mapOrigCopy map[any]any, rectanchoredrectFrom *RectAnchoredRect) (rectanchoredrectTo *RectAnchoredRect) {
	var alreadyCopied bool
	rectanchoredrectTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, rectanchoredrectFrom)
	if alreadyCopied {
		return
	}
	rectanchoredrectFrom.GongCopyBasicFields(rectanchoredrectTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRectAnchoredText(mapOrigCopy map[any]any, rectanchoredtextFrom *RectAnchoredText) (rectanchoredtextTo *RectAnchoredText) {
	var alreadyCopied bool
	rectanchoredtextTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, rectanchoredtextFrom)
	if alreadyCopied {
		return
	}
	rectanchoredtextFrom.GongCopyBasicFields(rectanchoredtextTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range rectanchoredtextFrom.Animates {
		rectanchoredtextTo.Animates = append(rectanchoredtextTo.Animates, GongCopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func GongCopyBranchRectLinkLink(mapOrigCopy map[any]any, rectlinklinkFrom *RectLinkLink) (rectlinklinkTo *RectLinkLink) {
	var alreadyCopied bool
	rectlinklinkTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, rectlinklinkFrom)
	if alreadyCopied {
		return
	}
	rectlinklinkFrom.GongCopyBasicFields(rectlinklinkTo)

	//insertion point for the staging of instances referenced by pointers
	if rectlinklinkFrom.Start != nil {
		rectlinklinkTo.Start = GongCopyBranchRect(mapOrigCopy, rectlinklinkFrom.Start)
	}
	if rectlinklinkFrom.End != nil {
		rectlinklinkTo.End = GongCopyBranchLink(mapOrigCopy, rectlinklinkFrom.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSVG(mapOrigCopy map[any]any, svgFrom *SVG) (svgTo *SVG) {
	var alreadyCopied bool
	svgTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, svgFrom)
	if alreadyCopied {
		return
	}
	svgFrom.GongCopyBasicFields(svgTo)

	//insertion point for the staging of instances referenced by pointers
	if svgFrom.StartRect != nil {
		svgTo.StartRect = GongCopyBranchRect(mapOrigCopy, svgFrom.StartRect)
	}
	if svgFrom.EndRect != nil {
		svgTo.EndRect = GongCopyBranchRect(mapOrigCopy, svgFrom.EndRect)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _layer := range svgFrom.Layers {
		svgTo.Layers = append(svgTo.Layers, GongCopyBranchLayer(mapOrigCopy, _layer))
	}

	return
}

func GongCopyBranchSvgText(mapOrigCopy map[any]any, svgtextFrom *SvgText) (svgtextTo *SvgText) {
	var alreadyCopied bool
	svgtextTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, svgtextFrom)
	if alreadyCopied {
		return
	}
	svgtextFrom.GongCopyBasicFields(svgtextTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchText(mapOrigCopy map[any]any, textFrom *Text) (textTo *Text) {
	var alreadyCopied bool
	textTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, textFrom)
	if alreadyCopied {
		return
	}
	textFrom.GongCopyBasicFields(textTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range textFrom.Animates {
		textTo.Animates = append(textTo.Animates, GongCopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// insertion point for unstage branch per struct
func (animate *Animate) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(animate) {
		return
	}

	animate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (circle *Circle) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(circle) {
		return
	}

	circle.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range circle.Animations {
		stage.UnstageBranch(_animate)
	}

}

func (condition *Condition) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(condition) {
		return
	}

	condition.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (controlpoint *ControlPoint) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(controlpoint) {
		return
	}

	controlpoint.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if controlpoint.ClosestRect != nil {
		stage.UnstageBranch(controlpoint.ClosestRect)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ellipse *Ellipse) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ellipse) {
		return
	}

	ellipse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range ellipse.Animates {
		stage.UnstageBranch(_animate)
	}

}

func (filetodownload *FileToDownload) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(filetodownload) {
		return
	}

	filetodownload.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (layer *Layer) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(layer) {
		return
	}

	layer.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rect := range layer.Rects {
		stage.UnstageBranch(_rect)
	}
	for _, _text := range layer.Texts {
		stage.UnstageBranch(_text)
	}
	for _, _circle := range layer.Circles {
		stage.UnstageBranch(_circle)
	}
	for _, _line := range layer.Lines {
		stage.UnstageBranch(_line)
	}
	for _, _ellipse := range layer.Ellipses {
		stage.UnstageBranch(_ellipse)
	}
	for _, _polyline := range layer.Polylines {
		stage.UnstageBranch(_polyline)
	}
	for _, _polygone := range layer.Polygones {
		stage.UnstageBranch(_polygone)
	}
	for _, _path := range layer.Paths {
		stage.UnstageBranch(_path)
	}
	for _, _link := range layer.Links {
		stage.UnstageBranch(_link)
	}
	for _, _rectlinklink := range layer.RectLinkLinks {
		stage.UnstageBranch(_rectlinklink)
	}

}

func (line *Line) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(line) {
		return
	}

	line.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range line.Animates {
		stage.UnstageBranch(_animate)
	}

}

func (link *Link) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(link) {
		return
	}

	link.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if link.Start != nil {
		stage.UnstageBranch(link.Start)
	}
	if link.End != nil {
		stage.UnstageBranch(link.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _linkanchoredtext := range link.TextAtArrowStart {
		stage.UnstageBranch(_linkanchoredtext)
	}
	for _, _linkanchoredtext := range link.TextAtArrowEnd {
		stage.UnstageBranch(_linkanchoredtext)
	}
	for _, _linkanchoredtext := range link.TextAtCorner {
		stage.UnstageBranch(_linkanchoredtext)
	}
	for _, _linkanchoredpath := range link.PathAtArrowStart {
		stage.UnstageBranch(_linkanchoredpath)
	}
	for _, _linkanchoredpath := range link.PathAtArrowEnd {
		stage.UnstageBranch(_linkanchoredpath)
	}
	for _, _linkanchoredpath := range link.PathAtCorner {
		stage.UnstageBranch(_linkanchoredpath)
	}
	for _, _controlpoint := range link.ControlPoints {
		stage.UnstageBranch(_controlpoint)
	}

}

func (linkanchoredpath *LinkAnchoredPath) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(linkanchoredpath) {
		return
	}

	linkanchoredpath.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (linkanchoredtext *LinkAnchoredText) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(linkanchoredtext) {
		return
	}

	linkanchoredtext.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range linkanchoredtext.Animates {
		stage.UnstageBranch(_animate)
	}

}

func (path *Path) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(path) {
		return
	}

	path.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range path.Animates {
		stage.UnstageBranch(_animate)
	}

}

func (point *Point) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(point) {
		return
	}

	point.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (polygone *Polygone) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(polygone) {
		return
	}

	polygone.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range polygone.Animates {
		stage.UnstageBranch(_animate)
	}

}

func (polyline *Polyline) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(polyline) {
		return
	}

	polyline.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range polyline.Animates {
		stage.UnstageBranch(_animate)
	}

}

func (rect *Rect) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(rect) {
		return
	}

	rect.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if rect.EnclosingRect != nil {
		stage.UnstageBranch(rect.EnclosingRect)
	}
	if rect.AnchoredTo != nil {
		stage.UnstageBranch(rect.AnchoredTo)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rect := range rect.Peers {
		stage.UnstageBranch(_rect)
	}
	for _, _rect := range rect.Obstacles {
		stage.UnstageBranch(_rect)
	}
	for _, _condition := range rect.HoveringTrigger {
		stage.UnstageBranch(_condition)
	}
	for _, _condition := range rect.DisplayConditions {
		stage.UnstageBranch(_condition)
	}
	for _, _animate := range rect.Animations {
		stage.UnstageBranch(_animate)
	}
	for _, _rectanchoredtext := range rect.RectAnchoredTexts {
		stage.UnstageBranch(_rectanchoredtext)
	}
	for _, _rectanchoredrect := range rect.RectAnchoredRects {
		stage.UnstageBranch(_rectanchoredrect)
	}
	for _, _rectanchoredpath := range rect.RectAnchoredPaths {
		stage.UnstageBranch(_rectanchoredpath)
	}
	for _, _rectanchoredpngimage := range rect.RectAnchoredPngImages {
		stage.UnstageBranch(_rectanchoredpngimage)
	}

}

func (rectanchoredpath *RectAnchoredPath) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(rectanchoredpath) {
		return
	}

	rectanchoredpath.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rectanchoredpngimage *RectAnchoredPngImage) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(rectanchoredpngimage) {
		return
	}

	rectanchoredpngimage.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rectanchoredrect *RectAnchoredRect) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(rectanchoredrect) {
		return
	}

	rectanchoredrect.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rectanchoredtext *RectAnchoredText) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(rectanchoredtext) {
		return
	}

	rectanchoredtext.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range rectanchoredtext.Animates {
		stage.UnstageBranch(_animate)
	}

}

func (rectlinklink *RectLinkLink) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(rectlinklink) {
		return
	}

	rectlinklink.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if rectlinklink.Start != nil {
		stage.UnstageBranch(rectlinklink.Start)
	}
	if rectlinklink.End != nil {
		stage.UnstageBranch(rectlinklink.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (svg *SVG) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(svg) {
		return
	}

	svg.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if svg.StartRect != nil {
		stage.UnstageBranch(svg.StartRect)
	}
	if svg.EndRect != nil {
		stage.UnstageBranch(svg.EndRect)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _layer := range svg.Layers {
		stage.UnstageBranch(_layer)
	}

}

func (svgtext *SvgText) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(svgtext) {
		return
	}

	svgtext.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (text *Text) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(text) {
		return
	}

	text.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range text.Animates {
		stage.UnstageBranch(_animate)
	}

}

// insertion point for pointer reconstruction from references
func (reference *Animate) GongReconstructPointersFromReferences(stage *Stage, instance *Animate) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Circle) GongReconstructPointersFromReferences(stage *Stage, instance *Circle) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Animations, stage.Animates_reference, instance.Animations)
}

func (reference *Condition) GongReconstructPointersFromReferences(stage *Stage, instance *Condition) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ControlPoint) GongReconstructPointersFromReferences(stage *Stage, instance *ControlPoint) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.ClosestRect, stage.Rects_reference, instance.ClosestRect)
	// insertion point for slice of pointers field
}

func (reference *Ellipse) GongReconstructPointersFromReferences(stage *Stage, instance *Ellipse) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Animates, stage.Animates_reference, instance.Animates)
}

func (reference *FileToDownload) GongReconstructPointersFromReferences(stage *Stage, instance *FileToDownload) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Layer) GongReconstructPointersFromReferences(stage *Stage, instance *Layer) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Rects, stage.Rects_reference, instance.Rects)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Texts, stage.Texts_reference, instance.Texts)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Circles, stage.Circles_reference, instance.Circles)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Lines, stage.Lines_reference, instance.Lines)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Ellipses, stage.Ellipses_reference, instance.Ellipses)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Polylines, stage.Polylines_reference, instance.Polylines)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Polygones, stage.Polygones_reference, instance.Polygones)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Paths, stage.Paths_reference, instance.Paths)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Links, stage.Links_reference, instance.Links)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RectLinkLinks, stage.RectLinkLinks_reference, instance.RectLinkLinks)
}

func (reference *Line) GongReconstructPointersFromReferences(stage *Stage, instance *Line) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Animates, stage.Animates_reference, instance.Animates)
}

func (reference *Link) GongReconstructPointersFromReferences(stage *Stage, instance *Link) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Start, stage.Rects_reference, instance.Start)
	__gong__reconstructPointer(&reference.End, stage.Rects_reference, instance.End)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.TextAtArrowStart, stage.LinkAnchoredTexts_reference, instance.TextAtArrowStart)
	__gong__reconstructSliceOfPointersFromReferences(&reference.TextAtArrowEnd, stage.LinkAnchoredTexts_reference, instance.TextAtArrowEnd)
	__gong__reconstructSliceOfPointersFromReferences(&reference.TextAtCorner, stage.LinkAnchoredTexts_reference, instance.TextAtCorner)
	__gong__reconstructSliceOfPointersFromReferences(&reference.PathAtArrowStart, stage.LinkAnchoredPaths_reference, instance.PathAtArrowStart)
	__gong__reconstructSliceOfPointersFromReferences(&reference.PathAtArrowEnd, stage.LinkAnchoredPaths_reference, instance.PathAtArrowEnd)
	__gong__reconstructSliceOfPointersFromReferences(&reference.PathAtCorner, stage.LinkAnchoredPaths_reference, instance.PathAtCorner)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ControlPoints, stage.ControlPoints_reference, instance.ControlPoints)
}

func (reference *LinkAnchoredPath) GongReconstructPointersFromReferences(stage *Stage, instance *LinkAnchoredPath) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *LinkAnchoredText) GongReconstructPointersFromReferences(stage *Stage, instance *LinkAnchoredText) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Animates, stage.Animates_reference, instance.Animates)
}

func (reference *Path) GongReconstructPointersFromReferences(stage *Stage, instance *Path) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Animates, stage.Animates_reference, instance.Animates)
}

func (reference *Point) GongReconstructPointersFromReferences(stage *Stage, instance *Point) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Polygone) GongReconstructPointersFromReferences(stage *Stage, instance *Polygone) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Animates, stage.Animates_reference, instance.Animates)
}

func (reference *Polyline) GongReconstructPointersFromReferences(stage *Stage, instance *Polyline) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Animates, stage.Animates_reference, instance.Animates)
}

func (reference *Rect) GongReconstructPointersFromReferences(stage *Stage, instance *Rect) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.EnclosingRect, stage.Rects_reference, instance.EnclosingRect)
	__gong__reconstructPointer(&reference.AnchoredTo, stage.Rects_reference, instance.AnchoredTo)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Peers, stage.Rects_reference, instance.Peers)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Obstacles, stage.Rects_reference, instance.Obstacles)
	__gong__reconstructSliceOfPointersFromReferences(&reference.HoveringTrigger, stage.Conditions_reference, instance.HoveringTrigger)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DisplayConditions, stage.Conditions_reference, instance.DisplayConditions)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Animations, stage.Animates_reference, instance.Animations)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RectAnchoredTexts, stage.RectAnchoredTexts_reference, instance.RectAnchoredTexts)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RectAnchoredRects, stage.RectAnchoredRects_reference, instance.RectAnchoredRects)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RectAnchoredPaths, stage.RectAnchoredPaths_reference, instance.RectAnchoredPaths)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RectAnchoredPngImages, stage.RectAnchoredPngImages_reference, instance.RectAnchoredPngImages)
}

func (reference *RectAnchoredPath) GongReconstructPointersFromReferences(stage *Stage, instance *RectAnchoredPath) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *RectAnchoredPngImage) GongReconstructPointersFromReferences(stage *Stage, instance *RectAnchoredPngImage) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *RectAnchoredRect) GongReconstructPointersFromReferences(stage *Stage, instance *RectAnchoredRect) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *RectAnchoredText) GongReconstructPointersFromReferences(stage *Stage, instance *RectAnchoredText) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Animates, stage.Animates_reference, instance.Animates)
}

func (reference *RectLinkLink) GongReconstructPointersFromReferences(stage *Stage, instance *RectLinkLink) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Start, stage.Rects_reference, instance.Start)
	__gong__reconstructPointer(&reference.End, stage.Links_reference, instance.End)
	// insertion point for slice of pointers field
}

func (reference *SVG) GongReconstructPointersFromReferences(stage *Stage, instance *SVG) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.StartRect, stage.Rects_reference, instance.StartRect)
	__gong__reconstructPointer(&reference.EndRect, stage.Rects_reference, instance.EndRect)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Layers, stage.Layers_reference, instance.Layers)
}

func (reference *SvgText) GongReconstructPointersFromReferences(stage *Stage, instance *SvgText) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Text) GongReconstructPointersFromReferences(stage *Stage, instance *Text) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Animates, stage.Animates_reference, instance.Animates)
}

// insertion point for pointer reconstruction from instances
func (reference *Animate) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Circle) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Animations, stage.Animates_instance)
}

func (reference *Condition) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ControlPoint) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.ClosestRect, stage.Rects_instance)
	// insertion point for slice of pointers fields
}

func (reference *Ellipse) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Animates, stage.Animates_instance)
}

func (reference *FileToDownload) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Layer) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Rects, stage.Rects_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Texts, stage.Texts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Circles, stage.Circles_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Lines, stage.Lines_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Ellipses, stage.Ellipses_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Polylines, stage.Polylines_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Polygones, stage.Polygones_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Paths, stage.Paths_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Links, stage.Links_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RectLinkLinks, stage.RectLinkLinks_instance)
}

func (reference *Line) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Animates, stage.Animates_instance)
}

func (reference *Link) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Start, stage.Rects_instance)
	__gong__reconstructPointerFromInstance(&reference.End, stage.Rects_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.TextAtArrowStart, stage.LinkAnchoredTexts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.TextAtArrowEnd, stage.LinkAnchoredTexts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.TextAtCorner, stage.LinkAnchoredTexts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.PathAtArrowStart, stage.LinkAnchoredPaths_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.PathAtArrowEnd, stage.LinkAnchoredPaths_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.PathAtCorner, stage.LinkAnchoredPaths_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ControlPoints, stage.ControlPoints_instance)
}

func (reference *LinkAnchoredPath) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *LinkAnchoredText) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Animates, stage.Animates_instance)
}

func (reference *Path) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Animates, stage.Animates_instance)
}

func (reference *Point) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Polygone) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Animates, stage.Animates_instance)
}

func (reference *Polyline) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Animates, stage.Animates_instance)
}

func (reference *Rect) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.EnclosingRect, stage.Rects_instance)
	__gong__reconstructPointerFromInstance(&reference.AnchoredTo, stage.Rects_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Peers, stage.Rects_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Obstacles, stage.Rects_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.HoveringTrigger, stage.Conditions_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DisplayConditions, stage.Conditions_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Animations, stage.Animates_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RectAnchoredTexts, stage.RectAnchoredTexts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RectAnchoredRects, stage.RectAnchoredRects_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RectAnchoredPaths, stage.RectAnchoredPaths_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RectAnchoredPngImages, stage.RectAnchoredPngImages_instance)
}

func (reference *RectAnchoredPath) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *RectAnchoredPngImage) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *RectAnchoredRect) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *RectAnchoredText) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Animates, stage.Animates_instance)
}

func (reference *RectLinkLink) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Start, stage.Rects_instance)
	__gong__reconstructPointerFromInstance(&reference.End, stage.Links_instance)
	// insertion point for slice of pointers fields
}

func (reference *SVG) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.StartRect, stage.Rects_instance)
	__gong__reconstructPointerFromInstance(&reference.EndRect, stage.Rects_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Layers, stage.Layers_instance)
}

func (reference *SvgText) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Text) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Animates, stage.Animates_instance)
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
	if ops := __gong__diffSliceOfPointers(stage, circle, "Animations", circleOther.Animations, circle.Animations); ops != "" {
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
	if controlpoint.ClosestRect != controlpointOther.ClosestRect {
		diffs = append(diffs, controlpoint.GongMarshallField(stage, "ClosestRect"))
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
	if ops := __gong__diffSliceOfPointers(stage, ellipse, "Animates", ellipseOther.Animates, ellipse.Animates); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (filetodownload *FileToDownload) GongDiff(stage *Stage, filetodownloadOther *FileToDownload) (diffs []string) {
	// insertion point for field diffs
	if filetodownload.Name != filetodownloadOther.Name {
		diffs = append(diffs, filetodownload.GongMarshallField(stage, "Name"))
	}
	if filetodownload.Base64EncodedContent != filetodownloadOther.Base64EncodedContent {
		diffs = append(diffs, filetodownload.GongMarshallField(stage, "Base64EncodedContent"))
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
	if ops := __gong__diffSliceOfPointers(stage, layer, "Rects", layerOther.Rects, layer.Rects); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, layer, "Texts", layerOther.Texts, layer.Texts); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, layer, "Circles", layerOther.Circles, layer.Circles); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, layer, "Lines", layerOther.Lines, layer.Lines); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, layer, "Ellipses", layerOther.Ellipses, layer.Ellipses); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, layer, "Polylines", layerOther.Polylines, layer.Polylines); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, layer, "Polygones", layerOther.Polygones, layer.Polygones); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, layer, "Paths", layerOther.Paths, layer.Paths); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, layer, "Links", layerOther.Links, layer.Links); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, layer, "RectLinkLinks", layerOther.RectLinkLinks, layer.RectLinkLinks); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, line, "Animates", lineOther.Animates, line.Animates); ops != "" {
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
	if link.Start != linkOther.Start {
		diffs = append(diffs, link.GongMarshallField(stage, "Start"))
	}
	if link.StartAnchorType != linkOther.StartAnchorType {
		diffs = append(diffs, link.GongMarshallField(stage, "StartAnchorType"))
	}
	if link.End != linkOther.End {
		diffs = append(diffs, link.GongMarshallField(stage, "End"))
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
	if ops := __gong__diffSliceOfPointers(stage, link, "TextAtArrowStart", linkOther.TextAtArrowStart, link.TextAtArrowStart); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, link, "TextAtArrowEnd", linkOther.TextAtArrowEnd, link.TextAtArrowEnd); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, link, "TextAtCorner", linkOther.TextAtCorner, link.TextAtCorner); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, link, "PathAtArrowStart", linkOther.PathAtArrowStart, link.PathAtArrowStart); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, link, "PathAtArrowEnd", linkOther.PathAtArrowEnd, link.PathAtArrowEnd); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, link, "PathAtCorner", linkOther.PathAtCorner, link.PathAtCorner); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, link, "ControlPoints", linkOther.ControlPoints, link.ControlPoints); ops != "" {
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
func (linkanchoredpath *LinkAnchoredPath) GongDiff(stage *Stage, linkanchoredpathOther *LinkAnchoredPath) (diffs []string) {
	// insertion point for field diffs
	if linkanchoredpath.Name != linkanchoredpathOther.Name {
		diffs = append(diffs, linkanchoredpath.GongMarshallField(stage, "Name"))
	}
	if linkanchoredpath.Definition != linkanchoredpathOther.Definition {
		diffs = append(diffs, linkanchoredpath.GongMarshallField(stage, "Definition"))
	}
	if linkanchoredpath.X_Offset != linkanchoredpathOther.X_Offset {
		diffs = append(diffs, linkanchoredpath.GongMarshallField(stage, "X_Offset"))
	}
	if linkanchoredpath.Y_Offset != linkanchoredpathOther.Y_Offset {
		diffs = append(diffs, linkanchoredpath.GongMarshallField(stage, "Y_Offset"))
	}
	if linkanchoredpath.ScalePropotionnally != linkanchoredpathOther.ScalePropotionnally {
		diffs = append(diffs, linkanchoredpath.GongMarshallField(stage, "ScalePropotionnally"))
	}
	if linkanchoredpath.AppliedScaling != linkanchoredpathOther.AppliedScaling {
		diffs = append(diffs, linkanchoredpath.GongMarshallField(stage, "AppliedScaling"))
	}
	if linkanchoredpath.Color != linkanchoredpathOther.Color {
		diffs = append(diffs, linkanchoredpath.GongMarshallField(stage, "Color"))
	}
	if linkanchoredpath.FillOpacity != linkanchoredpathOther.FillOpacity {
		diffs = append(diffs, linkanchoredpath.GongMarshallField(stage, "FillOpacity"))
	}
	if linkanchoredpath.Stroke != linkanchoredpathOther.Stroke {
		diffs = append(diffs, linkanchoredpath.GongMarshallField(stage, "Stroke"))
	}
	if linkanchoredpath.StrokeOpacity != linkanchoredpathOther.StrokeOpacity {
		diffs = append(diffs, linkanchoredpath.GongMarshallField(stage, "StrokeOpacity"))
	}
	if linkanchoredpath.StrokeWidth != linkanchoredpathOther.StrokeWidth {
		diffs = append(diffs, linkanchoredpath.GongMarshallField(stage, "StrokeWidth"))
	}
	if linkanchoredpath.StrokeDashArray != linkanchoredpathOther.StrokeDashArray {
		diffs = append(diffs, linkanchoredpath.GongMarshallField(stage, "StrokeDashArray"))
	}
	if linkanchoredpath.StrokeDashArrayWhenSelected != linkanchoredpathOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, linkanchoredpath.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if linkanchoredpath.Transform != linkanchoredpathOther.Transform {
		diffs = append(diffs, linkanchoredpath.GongMarshallField(stage, "Transform"))
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
	if ops := __gong__diffSliceOfPointers(stage, linkanchoredtext, "Animates", linkanchoredtextOther.Animates, linkanchoredtext.Animates); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, path, "Animates", pathOther.Animates, path.Animates); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, polygone, "Animates", polygoneOther.Animates, polygone.Animates); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, polyline, "Animates", polylineOther.Animates, polyline.Animates); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, rect, "Peers", rectOther.Peers, rect.Peers); ops != "" {
		diffs = append(diffs, ops)
	}
	if rect.EnclosingRect != rectOther.EnclosingRect {
		diffs = append(diffs, rect.GongMarshallField(stage, "EnclosingRect"))
	}
	if ops := __gong__diffSliceOfPointers(stage, rect, "Obstacles", rectOther.Obstacles, rect.Obstacles); ops != "" {
		diffs = append(diffs, ops)
	}
	if rect.AnchoredTo != rectOther.AnchoredTo {
		diffs = append(diffs, rect.GongMarshallField(stage, "AnchoredTo"))
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
	if ops := __gong__diffSliceOfPointers(stage, rect, "HoveringTrigger", rectOther.HoveringTrigger, rect.HoveringTrigger); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, rect, "DisplayConditions", rectOther.DisplayConditions, rect.DisplayConditions); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, rect, "Animations", rectOther.Animations, rect.Animations); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, rect, "RectAnchoredTexts", rectOther.RectAnchoredTexts, rect.RectAnchoredTexts); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, rect, "RectAnchoredRects", rectOther.RectAnchoredRects, rect.RectAnchoredRects); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, rect, "RectAnchoredPaths", rectOther.RectAnchoredPaths, rect.RectAnchoredPaths); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, rect, "RectAnchoredPngImages", rectOther.RectAnchoredPngImages, rect.RectAnchoredPngImages); ops != "" {
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
	if rect.URLPath != rectOther.URLPath {
		diffs = append(diffs, rect.GongMarshallField(stage, "URLPath"))
	}
	if rect.URLTarget != rectOther.URLTarget {
		diffs = append(diffs, rect.GongMarshallField(stage, "URLTarget"))
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
func (rectanchoredpngimage *RectAnchoredPngImage) GongDiff(stage *Stage, rectanchoredpngimageOther *RectAnchoredPngImage) (diffs []string) {
	// insertion point for field diffs
	if rectanchoredpngimage.Name != rectanchoredpngimageOther.Name {
		diffs = append(diffs, rectanchoredpngimage.GongMarshallField(stage, "Name"))
	}
	if rectanchoredpngimage.X != rectanchoredpngimageOther.X {
		diffs = append(diffs, rectanchoredpngimage.GongMarshallField(stage, "X"))
	}
	if rectanchoredpngimage.Y != rectanchoredpngimageOther.Y {
		diffs = append(diffs, rectanchoredpngimage.GongMarshallField(stage, "Y"))
	}
	if rectanchoredpngimage.Width != rectanchoredpngimageOther.Width {
		diffs = append(diffs, rectanchoredpngimage.GongMarshallField(stage, "Width"))
	}
	if rectanchoredpngimage.Height != rectanchoredpngimageOther.Height {
		diffs = append(diffs, rectanchoredpngimage.GongMarshallField(stage, "Height"))
	}
	if rectanchoredpngimage.RX != rectanchoredpngimageOther.RX {
		diffs = append(diffs, rectanchoredpngimage.GongMarshallField(stage, "RX"))
	}
	if rectanchoredpngimage.X_Offset != rectanchoredpngimageOther.X_Offset {
		diffs = append(diffs, rectanchoredpngimage.GongMarshallField(stage, "X_Offset"))
	}
	if rectanchoredpngimage.Y_Offset != rectanchoredpngimageOther.Y_Offset {
		diffs = append(diffs, rectanchoredpngimage.GongMarshallField(stage, "Y_Offset"))
	}
	if rectanchoredpngimage.RectAnchorType != rectanchoredpngimageOther.RectAnchorType {
		diffs = append(diffs, rectanchoredpngimage.GongMarshallField(stage, "RectAnchorType"))
	}
	if rectanchoredpngimage.Base64Content != rectanchoredpngimageOther.Base64Content {
		diffs = append(diffs, rectanchoredpngimage.GongMarshallField(stage, "Base64Content"))
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
	if ops := __gong__diffSliceOfPointers(stage, rectanchoredtext, "Animates", rectanchoredtextOther.Animates, rectanchoredtext.Animates); ops != "" {
		diffs = append(diffs, ops)
	}
	if rectanchoredtext.URLPath != rectanchoredtextOther.URLPath {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "URLPath"))
	}
	if rectanchoredtext.URLTarget != rectanchoredtextOther.URLTarget {
		diffs = append(diffs, rectanchoredtext.GongMarshallField(stage, "URLTarget"))
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
	if rectlinklink.Start != rectlinklinkOther.Start {
		diffs = append(diffs, rectlinklink.GongMarshallField(stage, "Start"))
	}
	if rectlinklink.End != rectlinklinkOther.End {
		diffs = append(diffs, rectlinklink.GongMarshallField(stage, "End"))
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
	if ops := __gong__diffSliceOfPointers(stage, svg, "Layers", svgOther.Layers, svg.Layers); ops != "" {
		diffs = append(diffs, ops)
	}
	if svg.DrawingState != svgOther.DrawingState {
		diffs = append(diffs, svg.GongMarshallField(stage, "DrawingState"))
	}
	if svg.StartRect != svgOther.StartRect {
		diffs = append(diffs, svg.GongMarshallField(stage, "StartRect"))
	}
	if svg.EndRect != svgOther.EndRect {
		diffs = append(diffs, svg.GongMarshallField(stage, "EndRect"))
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
	if svg.PanX != svgOther.PanX {
		diffs = append(diffs, svg.GongMarshallField(stage, "PanX"))
	}
	if svg.PanY != svgOther.PanY {
		diffs = append(diffs, svg.GongMarshallField(stage, "PanY"))
	}
	if svg.Zoom != svgOther.Zoom {
		diffs = append(diffs, svg.GongMarshallField(stage, "Zoom"))
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
	if ops := __gong__diffSliceOfPointers(stage, text, "Animates", textOther.Animates, text.Animates); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
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

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
