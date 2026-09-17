// generated code - do not edit
package models

import "fmt"

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (animate *Animate) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Animates[animate]

	return
}

func (stage *Stage) IsStagedAnimate(animate *Animate) (ok bool) {

	return animate.GongIsStaged(stage)
}

func (circle *Circle) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Circles[circle]

	return
}

func (stage *Stage) IsStagedCircle(circle *Circle) (ok bool) {

	return circle.GongIsStaged(stage)
}

func (condition *Condition) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Conditions[condition]

	return
}

func (stage *Stage) IsStagedCondition(condition *Condition) (ok bool) {

	return condition.GongIsStaged(stage)
}

func (controlpoint *ControlPoint) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ControlPoints[controlpoint]

	return
}

func (stage *Stage) IsStagedControlPoint(controlpoint *ControlPoint) (ok bool) {

	return controlpoint.GongIsStaged(stage)
}

func (ellipse *Ellipse) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Ellipses[ellipse]

	return
}

func (stage *Stage) IsStagedEllipse(ellipse *Ellipse) (ok bool) {

	return ellipse.GongIsStaged(stage)
}

func (filetodownload *FileToDownload) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.FileToDownloads[filetodownload]

	return
}

func (stage *Stage) IsStagedFileToDownload(filetodownload *FileToDownload) (ok bool) {

	return filetodownload.GongIsStaged(stage)
}

func (layer *Layer) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Layers[layer]

	return
}

func (stage *Stage) IsStagedLayer(layer *Layer) (ok bool) {

	return layer.GongIsStaged(stage)
}

func (line *Line) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Lines[line]

	return
}

func (stage *Stage) IsStagedLine(line *Line) (ok bool) {

	return line.GongIsStaged(stage)
}

func (link *Link) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Links[link]

	return
}

func (stage *Stage) IsStagedLink(link *Link) (ok bool) {

	return link.GongIsStaged(stage)
}

func (linkanchoredpath *LinkAnchoredPath) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.LinkAnchoredPaths[linkanchoredpath]

	return
}

func (stage *Stage) IsStagedLinkAnchoredPath(linkanchoredpath *LinkAnchoredPath) (ok bool) {

	return linkanchoredpath.GongIsStaged(stage)
}

func (linkanchoredtext *LinkAnchoredText) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.LinkAnchoredTexts[linkanchoredtext]

	return
}

func (stage *Stage) IsStagedLinkAnchoredText(linkanchoredtext *LinkAnchoredText) (ok bool) {

	return linkanchoredtext.GongIsStaged(stage)
}

func (path *Path) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Paths[path]

	return
}

func (stage *Stage) IsStagedPath(path *Path) (ok bool) {

	return path.GongIsStaged(stage)
}

func (point *Point) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Points[point]

	return
}

func (stage *Stage) IsStagedPoint(point *Point) (ok bool) {

	return point.GongIsStaged(stage)
}

func (polygone *Polygone) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Polygones[polygone]

	return
}

func (stage *Stage) IsStagedPolygone(polygone *Polygone) (ok bool) {

	return polygone.GongIsStaged(stage)
}

func (polyline *Polyline) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Polylines[polyline]

	return
}

func (stage *Stage) IsStagedPolyline(polyline *Polyline) (ok bool) {

	return polyline.GongIsStaged(stage)
}

func (rect *Rect) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Rects[rect]

	return
}

func (stage *Stage) IsStagedRect(rect *Rect) (ok bool) {

	return rect.GongIsStaged(stage)
}

func (rectanchoredpath *RectAnchoredPath) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.RectAnchoredPaths[rectanchoredpath]

	return
}

func (stage *Stage) IsStagedRectAnchoredPath(rectanchoredpath *RectAnchoredPath) (ok bool) {

	return rectanchoredpath.GongIsStaged(stage)
}

func (rectanchoredpngimage *RectAnchoredPngImage) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.RectAnchoredPngImages[rectanchoredpngimage]

	return
}

func (stage *Stage) IsStagedRectAnchoredPngImage(rectanchoredpngimage *RectAnchoredPngImage) (ok bool) {

	return rectanchoredpngimage.GongIsStaged(stage)
}

func (rectanchoredrect *RectAnchoredRect) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.RectAnchoredRects[rectanchoredrect]

	return
}

func (stage *Stage) IsStagedRectAnchoredRect(rectanchoredrect *RectAnchoredRect) (ok bool) {

	return rectanchoredrect.GongIsStaged(stage)
}

func (rectanchoredtext *RectAnchoredText) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.RectAnchoredTexts[rectanchoredtext]

	return
}

func (stage *Stage) IsStagedRectAnchoredText(rectanchoredtext *RectAnchoredText) (ok bool) {

	return rectanchoredtext.GongIsStaged(stage)
}

func (rectlinklink *RectLinkLink) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.RectLinkLinks[rectlinklink]

	return
}

func (stage *Stage) IsStagedRectLinkLink(rectlinklink *RectLinkLink) (ok bool) {

	return rectlinklink.GongIsStaged(stage)
}

func (svg *SVG) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.SVGs[svg]

	return
}

func (stage *Stage) IsStagedSVG(svg *SVG) (ok bool) {

	return svg.GongIsStaged(stage)
}

func (svgtext *SvgText) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.SvgTexts[svgtext]

	return
}

func (stage *Stage) IsStagedSvgText(svgtext *SvgText) (ok bool) {

	return svgtext.GongIsStaged(stage)
}

func (text *Text) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Texts[text]

	return
}

func (stage *Stage) IsStagedText(text *Text) (ok bool) {

	return text.GongIsStaged(stage)
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// StageBranch is a backward-compatible package-level forwarder.
func StageBranch(stage *Stage, instance GongstructIF) {
	stage.StageBranch(instance)
}

// insertion point for stage branch per struct
func (animate *Animate) GongStageBranch(stage *Stage) {
	stage.StageBranchAnimate(animate)
}

func (stage *Stage) StageBranchAnimate(animate *Animate) {

	// check if instance is already staged
	if stage.IsStaged(animate) {
		return
	}

	animate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (circle *Circle) GongStageBranch(stage *Stage) {
	stage.StageBranchCircle(circle)
}

func (stage *Stage) StageBranchCircle(circle *Circle) {

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
	stage.StageBranchCondition(condition)
}

func (stage *Stage) StageBranchCondition(condition *Condition) {

	// check if instance is already staged
	if stage.IsStaged(condition) {
		return
	}

	condition.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (controlpoint *ControlPoint) GongStageBranch(stage *Stage) {
	stage.StageBranchControlPoint(controlpoint)
}

func (stage *Stage) StageBranchControlPoint(controlpoint *ControlPoint) {

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
	stage.StageBranchEllipse(ellipse)
}

func (stage *Stage) StageBranchEllipse(ellipse *Ellipse) {

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
	stage.StageBranchFileToDownload(filetodownload)
}

func (stage *Stage) StageBranchFileToDownload(filetodownload *FileToDownload) {

	// check if instance is already staged
	if stage.IsStaged(filetodownload) {
		return
	}

	filetodownload.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (layer *Layer) GongStageBranch(stage *Stage) {
	stage.StageBranchLayer(layer)
}

func (stage *Stage) StageBranchLayer(layer *Layer) {

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
	stage.StageBranchLine(line)
}

func (stage *Stage) StageBranchLine(line *Line) {

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
	stage.StageBranchLink(link)
}

func (stage *Stage) StageBranchLink(link *Link) {

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
	stage.StageBranchLinkAnchoredPath(linkanchoredpath)
}

func (stage *Stage) StageBranchLinkAnchoredPath(linkanchoredpath *LinkAnchoredPath) {

	// check if instance is already staged
	if stage.IsStaged(linkanchoredpath) {
		return
	}

	linkanchoredpath.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (linkanchoredtext *LinkAnchoredText) GongStageBranch(stage *Stage) {
	stage.StageBranchLinkAnchoredText(linkanchoredtext)
}

func (stage *Stage) StageBranchLinkAnchoredText(linkanchoredtext *LinkAnchoredText) {

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
	stage.StageBranchPath(path)
}

func (stage *Stage) StageBranchPath(path *Path) {

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
	stage.StageBranchPoint(point)
}

func (stage *Stage) StageBranchPoint(point *Point) {

	// check if instance is already staged
	if stage.IsStaged(point) {
		return
	}

	point.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (polygone *Polygone) GongStageBranch(stage *Stage) {
	stage.StageBranchPolygone(polygone)
}

func (stage *Stage) StageBranchPolygone(polygone *Polygone) {

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
	stage.StageBranchPolyline(polyline)
}

func (stage *Stage) StageBranchPolyline(polyline *Polyline) {

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
	stage.StageBranchRect(rect)
}

func (stage *Stage) StageBranchRect(rect *Rect) {

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
	stage.StageBranchRectAnchoredPath(rectanchoredpath)
}

func (stage *Stage) StageBranchRectAnchoredPath(rectanchoredpath *RectAnchoredPath) {

	// check if instance is already staged
	if stage.IsStaged(rectanchoredpath) {
		return
	}

	rectanchoredpath.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rectanchoredpngimage *RectAnchoredPngImage) GongStageBranch(stage *Stage) {
	stage.StageBranchRectAnchoredPngImage(rectanchoredpngimage)
}

func (stage *Stage) StageBranchRectAnchoredPngImage(rectanchoredpngimage *RectAnchoredPngImage) {

	// check if instance is already staged
	if stage.IsStaged(rectanchoredpngimage) {
		return
	}

	rectanchoredpngimage.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rectanchoredrect *RectAnchoredRect) GongStageBranch(stage *Stage) {
	stage.StageBranchRectAnchoredRect(rectanchoredrect)
}

func (stage *Stage) StageBranchRectAnchoredRect(rectanchoredrect *RectAnchoredRect) {

	// check if instance is already staged
	if stage.IsStaged(rectanchoredrect) {
		return
	}

	rectanchoredrect.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rectanchoredtext *RectAnchoredText) GongStageBranch(stage *Stage) {
	stage.StageBranchRectAnchoredText(rectanchoredtext)
}

func (stage *Stage) StageBranchRectAnchoredText(rectanchoredtext *RectAnchoredText) {

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
	stage.StageBranchRectLinkLink(rectlinklink)
}

func (stage *Stage) StageBranchRectLinkLink(rectlinklink *RectLinkLink) {

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
	stage.StageBranchSVG(svg)
}

func (stage *Stage) StageBranchSVG(svg *SVG) {

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
	stage.StageBranchSvgText(svgtext)
}

func (stage *Stage) StageBranchSvgText(svgtext *SvgText) {

	// check if instance is already staged
	if stage.IsStaged(svgtext) {
		return
	}

	svgtext.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (text *Text) GongStageBranch(stage *Stage) {
	stage.StageBranchText(text)
}

func (stage *Stage) StageBranchText(text *Text) {

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

	// animateFrom has already been copied
	if _animateTo, ok := mapOrigCopy[animateFrom]; ok {
		animateTo = _animateTo.(*Animate)
		return
	}

	animateTo = new(Animate)
	mapOrigCopy[animateFrom] = animateTo
	animateFrom.GongCopyBasicFields(animateTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCircle(mapOrigCopy map[any]any, circleFrom *Circle) (circleTo *Circle) {

	// circleFrom has already been copied
	if _circleTo, ok := mapOrigCopy[circleFrom]; ok {
		circleTo = _circleTo.(*Circle)
		return
	}

	circleTo = new(Circle)
	mapOrigCopy[circleFrom] = circleTo
	circleFrom.GongCopyBasicFields(circleTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range circleFrom.Animations {
		circleTo.Animations = append(circleTo.Animations, GongCopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func GongCopyBranchCondition(mapOrigCopy map[any]any, conditionFrom *Condition) (conditionTo *Condition) {

	// conditionFrom has already been copied
	if _conditionTo, ok := mapOrigCopy[conditionFrom]; ok {
		conditionTo = _conditionTo.(*Condition)
		return
	}

	conditionTo = new(Condition)
	mapOrigCopy[conditionFrom] = conditionTo
	conditionFrom.GongCopyBasicFields(conditionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchControlPoint(mapOrigCopy map[any]any, controlpointFrom *ControlPoint) (controlpointTo *ControlPoint) {

	// controlpointFrom has already been copied
	if _controlpointTo, ok := mapOrigCopy[controlpointFrom]; ok {
		controlpointTo = _controlpointTo.(*ControlPoint)
		return
	}

	controlpointTo = new(ControlPoint)
	mapOrigCopy[controlpointFrom] = controlpointTo
	controlpointFrom.GongCopyBasicFields(controlpointTo)

	//insertion point for the staging of instances referenced by pointers
	if controlpointFrom.ClosestRect != nil {
		controlpointTo.ClosestRect = GongCopyBranchRect(mapOrigCopy, controlpointFrom.ClosestRect)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEllipse(mapOrigCopy map[any]any, ellipseFrom *Ellipse) (ellipseTo *Ellipse) {

	// ellipseFrom has already been copied
	if _ellipseTo, ok := mapOrigCopy[ellipseFrom]; ok {
		ellipseTo = _ellipseTo.(*Ellipse)
		return
	}

	ellipseTo = new(Ellipse)
	mapOrigCopy[ellipseFrom] = ellipseTo
	ellipseFrom.GongCopyBasicFields(ellipseTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range ellipseFrom.Animates {
		ellipseTo.Animates = append(ellipseTo.Animates, GongCopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func GongCopyBranchFileToDownload(mapOrigCopy map[any]any, filetodownloadFrom *FileToDownload) (filetodownloadTo *FileToDownload) {

	// filetodownloadFrom has already been copied
	if _filetodownloadTo, ok := mapOrigCopy[filetodownloadFrom]; ok {
		filetodownloadTo = _filetodownloadTo.(*FileToDownload)
		return
	}

	filetodownloadTo = new(FileToDownload)
	mapOrigCopy[filetodownloadFrom] = filetodownloadTo
	filetodownloadFrom.GongCopyBasicFields(filetodownloadTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLayer(mapOrigCopy map[any]any, layerFrom *Layer) (layerTo *Layer) {

	// layerFrom has already been copied
	if _layerTo, ok := mapOrigCopy[layerFrom]; ok {
		layerTo = _layerTo.(*Layer)
		return
	}

	layerTo = new(Layer)
	mapOrigCopy[layerFrom] = layerTo
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

	// lineFrom has already been copied
	if _lineTo, ok := mapOrigCopy[lineFrom]; ok {
		lineTo = _lineTo.(*Line)
		return
	}

	lineTo = new(Line)
	mapOrigCopy[lineFrom] = lineTo
	lineFrom.GongCopyBasicFields(lineTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range lineFrom.Animates {
		lineTo.Animates = append(lineTo.Animates, GongCopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func GongCopyBranchLink(mapOrigCopy map[any]any, linkFrom *Link) (linkTo *Link) {

	// linkFrom has already been copied
	if _linkTo, ok := mapOrigCopy[linkFrom]; ok {
		linkTo = _linkTo.(*Link)
		return
	}

	linkTo = new(Link)
	mapOrigCopy[linkFrom] = linkTo
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

	// linkanchoredpathFrom has already been copied
	if _linkanchoredpathTo, ok := mapOrigCopy[linkanchoredpathFrom]; ok {
		linkanchoredpathTo = _linkanchoredpathTo.(*LinkAnchoredPath)
		return
	}

	linkanchoredpathTo = new(LinkAnchoredPath)
	mapOrigCopy[linkanchoredpathFrom] = linkanchoredpathTo
	linkanchoredpathFrom.GongCopyBasicFields(linkanchoredpathTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLinkAnchoredText(mapOrigCopy map[any]any, linkanchoredtextFrom *LinkAnchoredText) (linkanchoredtextTo *LinkAnchoredText) {

	// linkanchoredtextFrom has already been copied
	if _linkanchoredtextTo, ok := mapOrigCopy[linkanchoredtextFrom]; ok {
		linkanchoredtextTo = _linkanchoredtextTo.(*LinkAnchoredText)
		return
	}

	linkanchoredtextTo = new(LinkAnchoredText)
	mapOrigCopy[linkanchoredtextFrom] = linkanchoredtextTo
	linkanchoredtextFrom.GongCopyBasicFields(linkanchoredtextTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range linkanchoredtextFrom.Animates {
		linkanchoredtextTo.Animates = append(linkanchoredtextTo.Animates, GongCopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func GongCopyBranchPath(mapOrigCopy map[any]any, pathFrom *Path) (pathTo *Path) {

	// pathFrom has already been copied
	if _pathTo, ok := mapOrigCopy[pathFrom]; ok {
		pathTo = _pathTo.(*Path)
		return
	}

	pathTo = new(Path)
	mapOrigCopy[pathFrom] = pathTo
	pathFrom.GongCopyBasicFields(pathTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range pathFrom.Animates {
		pathTo.Animates = append(pathTo.Animates, GongCopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func GongCopyBranchPoint(mapOrigCopy map[any]any, pointFrom *Point) (pointTo *Point) {

	// pointFrom has already been copied
	if _pointTo, ok := mapOrigCopy[pointFrom]; ok {
		pointTo = _pointTo.(*Point)
		return
	}

	pointTo = new(Point)
	mapOrigCopy[pointFrom] = pointTo
	pointFrom.GongCopyBasicFields(pointTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPolygone(mapOrigCopy map[any]any, polygoneFrom *Polygone) (polygoneTo *Polygone) {

	// polygoneFrom has already been copied
	if _polygoneTo, ok := mapOrigCopy[polygoneFrom]; ok {
		polygoneTo = _polygoneTo.(*Polygone)
		return
	}

	polygoneTo = new(Polygone)
	mapOrigCopy[polygoneFrom] = polygoneTo
	polygoneFrom.GongCopyBasicFields(polygoneTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range polygoneFrom.Animates {
		polygoneTo.Animates = append(polygoneTo.Animates, GongCopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func GongCopyBranchPolyline(mapOrigCopy map[any]any, polylineFrom *Polyline) (polylineTo *Polyline) {

	// polylineFrom has already been copied
	if _polylineTo, ok := mapOrigCopy[polylineFrom]; ok {
		polylineTo = _polylineTo.(*Polyline)
		return
	}

	polylineTo = new(Polyline)
	mapOrigCopy[polylineFrom] = polylineTo
	polylineFrom.GongCopyBasicFields(polylineTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range polylineFrom.Animates {
		polylineTo.Animates = append(polylineTo.Animates, GongCopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func GongCopyBranchRect(mapOrigCopy map[any]any, rectFrom *Rect) (rectTo *Rect) {

	// rectFrom has already been copied
	if _rectTo, ok := mapOrigCopy[rectFrom]; ok {
		rectTo = _rectTo.(*Rect)
		return
	}

	rectTo = new(Rect)
	mapOrigCopy[rectFrom] = rectTo
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

	// rectanchoredpathFrom has already been copied
	if _rectanchoredpathTo, ok := mapOrigCopy[rectanchoredpathFrom]; ok {
		rectanchoredpathTo = _rectanchoredpathTo.(*RectAnchoredPath)
		return
	}

	rectanchoredpathTo = new(RectAnchoredPath)
	mapOrigCopy[rectanchoredpathFrom] = rectanchoredpathTo
	rectanchoredpathFrom.GongCopyBasicFields(rectanchoredpathTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRectAnchoredPngImage(mapOrigCopy map[any]any, rectanchoredpngimageFrom *RectAnchoredPngImage) (rectanchoredpngimageTo *RectAnchoredPngImage) {

	// rectanchoredpngimageFrom has already been copied
	if _rectanchoredpngimageTo, ok := mapOrigCopy[rectanchoredpngimageFrom]; ok {
		rectanchoredpngimageTo = _rectanchoredpngimageTo.(*RectAnchoredPngImage)
		return
	}

	rectanchoredpngimageTo = new(RectAnchoredPngImage)
	mapOrigCopy[rectanchoredpngimageFrom] = rectanchoredpngimageTo
	rectanchoredpngimageFrom.GongCopyBasicFields(rectanchoredpngimageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRectAnchoredRect(mapOrigCopy map[any]any, rectanchoredrectFrom *RectAnchoredRect) (rectanchoredrectTo *RectAnchoredRect) {

	// rectanchoredrectFrom has already been copied
	if _rectanchoredrectTo, ok := mapOrigCopy[rectanchoredrectFrom]; ok {
		rectanchoredrectTo = _rectanchoredrectTo.(*RectAnchoredRect)
		return
	}

	rectanchoredrectTo = new(RectAnchoredRect)
	mapOrigCopy[rectanchoredrectFrom] = rectanchoredrectTo
	rectanchoredrectFrom.GongCopyBasicFields(rectanchoredrectTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRectAnchoredText(mapOrigCopy map[any]any, rectanchoredtextFrom *RectAnchoredText) (rectanchoredtextTo *RectAnchoredText) {

	// rectanchoredtextFrom has already been copied
	if _rectanchoredtextTo, ok := mapOrigCopy[rectanchoredtextFrom]; ok {
		rectanchoredtextTo = _rectanchoredtextTo.(*RectAnchoredText)
		return
	}

	rectanchoredtextTo = new(RectAnchoredText)
	mapOrigCopy[rectanchoredtextFrom] = rectanchoredtextTo
	rectanchoredtextFrom.GongCopyBasicFields(rectanchoredtextTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _animate := range rectanchoredtextFrom.Animates {
		rectanchoredtextTo.Animates = append(rectanchoredtextTo.Animates, GongCopyBranchAnimate(mapOrigCopy, _animate))
	}

	return
}

func GongCopyBranchRectLinkLink(mapOrigCopy map[any]any, rectlinklinkFrom *RectLinkLink) (rectlinklinkTo *RectLinkLink) {

	// rectlinklinkFrom has already been copied
	if _rectlinklinkTo, ok := mapOrigCopy[rectlinklinkFrom]; ok {
		rectlinklinkTo = _rectlinklinkTo.(*RectLinkLink)
		return
	}

	rectlinklinkTo = new(RectLinkLink)
	mapOrigCopy[rectlinklinkFrom] = rectlinklinkTo
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

	// svgFrom has already been copied
	if _svgTo, ok := mapOrigCopy[svgFrom]; ok {
		svgTo = _svgTo.(*SVG)
		return
	}

	svgTo = new(SVG)
	mapOrigCopy[svgFrom] = svgTo
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

	// svgtextFrom has already been copied
	if _svgtextTo, ok := mapOrigCopy[svgtextFrom]; ok {
		svgtextTo = _svgtextTo.(*SvgText)
		return
	}

	svgtextTo = new(SvgText)
	mapOrigCopy[svgtextFrom] = svgtextTo
	svgtextFrom.GongCopyBasicFields(svgtextTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchText(mapOrigCopy map[any]any, textFrom *Text) (textTo *Text) {

	// textFrom has already been copied
	if _textTo, ok := mapOrigCopy[textFrom]; ok {
		textTo = _textTo.(*Text)
		return
	}

	textTo = new(Text)
	mapOrigCopy[textFrom] = textTo
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

// UnstageBranch is a backward-compatible package-level forwarder.
func UnstageBranch(stage *Stage, instance GongstructIF) {
	stage.UnstageBranch(instance)
}

// insertion point for unstage branch per struct
func (animate *Animate) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchAnimate(animate)
}

func (stage *Stage) UnstageBranchAnimate(animate *Animate) {

	// check if instance is already staged
	if !stage.IsStaged(animate) {
		return
	}

	animate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (circle *Circle) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchCircle(circle)
}

func (stage *Stage) UnstageBranchCircle(circle *Circle) {

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
	stage.UnstageBranchCondition(condition)
}

func (stage *Stage) UnstageBranchCondition(condition *Condition) {

	// check if instance is already staged
	if !stage.IsStaged(condition) {
		return
	}

	condition.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (controlpoint *ControlPoint) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchControlPoint(controlpoint)
}

func (stage *Stage) UnstageBranchControlPoint(controlpoint *ControlPoint) {

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
	stage.UnstageBranchEllipse(ellipse)
}

func (stage *Stage) UnstageBranchEllipse(ellipse *Ellipse) {

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
	stage.UnstageBranchFileToDownload(filetodownload)
}

func (stage *Stage) UnstageBranchFileToDownload(filetodownload *FileToDownload) {

	// check if instance is already staged
	if !stage.IsStaged(filetodownload) {
		return
	}

	filetodownload.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (layer *Layer) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchLayer(layer)
}

func (stage *Stage) UnstageBranchLayer(layer *Layer) {

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
	stage.UnstageBranchLine(line)
}

func (stage *Stage) UnstageBranchLine(line *Line) {

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
	stage.UnstageBranchLink(link)
}

func (stage *Stage) UnstageBranchLink(link *Link) {

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
	stage.UnstageBranchLinkAnchoredPath(linkanchoredpath)
}

func (stage *Stage) UnstageBranchLinkAnchoredPath(linkanchoredpath *LinkAnchoredPath) {

	// check if instance is already staged
	if !stage.IsStaged(linkanchoredpath) {
		return
	}

	linkanchoredpath.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (linkanchoredtext *LinkAnchoredText) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchLinkAnchoredText(linkanchoredtext)
}

func (stage *Stage) UnstageBranchLinkAnchoredText(linkanchoredtext *LinkAnchoredText) {

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
	stage.UnstageBranchPath(path)
}

func (stage *Stage) UnstageBranchPath(path *Path) {

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
	stage.UnstageBranchPoint(point)
}

func (stage *Stage) UnstageBranchPoint(point *Point) {

	// check if instance is already staged
	if !stage.IsStaged(point) {
		return
	}

	point.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (polygone *Polygone) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPolygone(polygone)
}

func (stage *Stage) UnstageBranchPolygone(polygone *Polygone) {

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
	stage.UnstageBranchPolyline(polyline)
}

func (stage *Stage) UnstageBranchPolyline(polyline *Polyline) {

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
	stage.UnstageBranchRect(rect)
}

func (stage *Stage) UnstageBranchRect(rect *Rect) {

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
	stage.UnstageBranchRectAnchoredPath(rectanchoredpath)
}

func (stage *Stage) UnstageBranchRectAnchoredPath(rectanchoredpath *RectAnchoredPath) {

	// check if instance is already staged
	if !stage.IsStaged(rectanchoredpath) {
		return
	}

	rectanchoredpath.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rectanchoredpngimage *RectAnchoredPngImage) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchRectAnchoredPngImage(rectanchoredpngimage)
}

func (stage *Stage) UnstageBranchRectAnchoredPngImage(rectanchoredpngimage *RectAnchoredPngImage) {

	// check if instance is already staged
	if !stage.IsStaged(rectanchoredpngimage) {
		return
	}

	rectanchoredpngimage.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rectanchoredrect *RectAnchoredRect) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchRectAnchoredRect(rectanchoredrect)
}

func (stage *Stage) UnstageBranchRectAnchoredRect(rectanchoredrect *RectAnchoredRect) {

	// check if instance is already staged
	if !stage.IsStaged(rectanchoredrect) {
		return
	}

	rectanchoredrect.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rectanchoredtext *RectAnchoredText) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchRectAnchoredText(rectanchoredtext)
}

func (stage *Stage) UnstageBranchRectAnchoredText(rectanchoredtext *RectAnchoredText) {

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
	stage.UnstageBranchRectLinkLink(rectlinklink)
}

func (stage *Stage) UnstageBranchRectLinkLink(rectlinklink *RectLinkLink) {

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
	stage.UnstageBranchSVG(svg)
}

func (stage *Stage) UnstageBranchSVG(svg *SVG) {

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
	stage.UnstageBranchSvgText(svgtext)
}

func (stage *Stage) UnstageBranchSvgText(svgtext *SvgText) {

	// check if instance is already staged
	if !stage.IsStaged(svgtext) {
		return
	}

	svgtext.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (text *Text) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchText(text)
}

func (stage *Stage) UnstageBranchText(text *Text) {

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
	reference.Animations = reference.Animations[:0]
	for _, _b := range instance.Animations {
		reference.Animations = append(reference.Animations, stage.Animates_reference[_b])
	}
}

func (reference *Condition) GongReconstructPointersFromReferences(stage *Stage, instance *Condition) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ControlPoint) GongReconstructPointersFromReferences(stage *Stage, instance *ControlPoint) {
	// insertion point for pointers field
	if instance.ClosestRect != nil {
		reference.ClosestRect = stage.Rects_reference[instance.ClosestRect]
	}
	// insertion point for slice of pointers field
}

func (reference *Ellipse) GongReconstructPointersFromReferences(stage *Stage, instance *Ellipse) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Animates = reference.Animates[:0]
	for _, _b := range instance.Animates {
		reference.Animates = append(reference.Animates, stage.Animates_reference[_b])
	}
}

func (reference *FileToDownload) GongReconstructPointersFromReferences(stage *Stage, instance *FileToDownload) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Layer) GongReconstructPointersFromReferences(stage *Stage, instance *Layer) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Rects = reference.Rects[:0]
	for _, _b := range instance.Rects {
		reference.Rects = append(reference.Rects, stage.Rects_reference[_b])
	}
	reference.Texts = reference.Texts[:0]
	for _, _b := range instance.Texts {
		reference.Texts = append(reference.Texts, stage.Texts_reference[_b])
	}
	reference.Circles = reference.Circles[:0]
	for _, _b := range instance.Circles {
		reference.Circles = append(reference.Circles, stage.Circles_reference[_b])
	}
	reference.Lines = reference.Lines[:0]
	for _, _b := range instance.Lines {
		reference.Lines = append(reference.Lines, stage.Lines_reference[_b])
	}
	reference.Ellipses = reference.Ellipses[:0]
	for _, _b := range instance.Ellipses {
		reference.Ellipses = append(reference.Ellipses, stage.Ellipses_reference[_b])
	}
	reference.Polylines = reference.Polylines[:0]
	for _, _b := range instance.Polylines {
		reference.Polylines = append(reference.Polylines, stage.Polylines_reference[_b])
	}
	reference.Polygones = reference.Polygones[:0]
	for _, _b := range instance.Polygones {
		reference.Polygones = append(reference.Polygones, stage.Polygones_reference[_b])
	}
	reference.Paths = reference.Paths[:0]
	for _, _b := range instance.Paths {
		reference.Paths = append(reference.Paths, stage.Paths_reference[_b])
	}
	reference.Links = reference.Links[:0]
	for _, _b := range instance.Links {
		reference.Links = append(reference.Links, stage.Links_reference[_b])
	}
	reference.RectLinkLinks = reference.RectLinkLinks[:0]
	for _, _b := range instance.RectLinkLinks {
		reference.RectLinkLinks = append(reference.RectLinkLinks, stage.RectLinkLinks_reference[_b])
	}
}

func (reference *Line) GongReconstructPointersFromReferences(stage *Stage, instance *Line) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Animates = reference.Animates[:0]
	for _, _b := range instance.Animates {
		reference.Animates = append(reference.Animates, stage.Animates_reference[_b])
	}
}

func (reference *Link) GongReconstructPointersFromReferences(stage *Stage, instance *Link) {
	// insertion point for pointers field
	if instance.Start != nil {
		reference.Start = stage.Rects_reference[instance.Start]
	}
	if instance.End != nil {
		reference.End = stage.Rects_reference[instance.End]
	}
	// insertion point for slice of pointers field
	reference.TextAtArrowStart = reference.TextAtArrowStart[:0]
	for _, _b := range instance.TextAtArrowStart {
		reference.TextAtArrowStart = append(reference.TextAtArrowStart, stage.LinkAnchoredTexts_reference[_b])
	}
	reference.TextAtArrowEnd = reference.TextAtArrowEnd[:0]
	for _, _b := range instance.TextAtArrowEnd {
		reference.TextAtArrowEnd = append(reference.TextAtArrowEnd, stage.LinkAnchoredTexts_reference[_b])
	}
	reference.TextAtCorner = reference.TextAtCorner[:0]
	for _, _b := range instance.TextAtCorner {
		reference.TextAtCorner = append(reference.TextAtCorner, stage.LinkAnchoredTexts_reference[_b])
	}
	reference.PathAtArrowStart = reference.PathAtArrowStart[:0]
	for _, _b := range instance.PathAtArrowStart {
		reference.PathAtArrowStart = append(reference.PathAtArrowStart, stage.LinkAnchoredPaths_reference[_b])
	}
	reference.PathAtArrowEnd = reference.PathAtArrowEnd[:0]
	for _, _b := range instance.PathAtArrowEnd {
		reference.PathAtArrowEnd = append(reference.PathAtArrowEnd, stage.LinkAnchoredPaths_reference[_b])
	}
	reference.PathAtCorner = reference.PathAtCorner[:0]
	for _, _b := range instance.PathAtCorner {
		reference.PathAtCorner = append(reference.PathAtCorner, stage.LinkAnchoredPaths_reference[_b])
	}
	reference.ControlPoints = reference.ControlPoints[:0]
	for _, _b := range instance.ControlPoints {
		reference.ControlPoints = append(reference.ControlPoints, stage.ControlPoints_reference[_b])
	}
}

func (reference *LinkAnchoredPath) GongReconstructPointersFromReferences(stage *Stage, instance *LinkAnchoredPath) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *LinkAnchoredText) GongReconstructPointersFromReferences(stage *Stage, instance *LinkAnchoredText) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Animates = reference.Animates[:0]
	for _, _b := range instance.Animates {
		reference.Animates = append(reference.Animates, stage.Animates_reference[_b])
	}
}

func (reference *Path) GongReconstructPointersFromReferences(stage *Stage, instance *Path) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Animates = reference.Animates[:0]
	for _, _b := range instance.Animates {
		reference.Animates = append(reference.Animates, stage.Animates_reference[_b])
	}
}

func (reference *Point) GongReconstructPointersFromReferences(stage *Stage, instance *Point) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Polygone) GongReconstructPointersFromReferences(stage *Stage, instance *Polygone) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Animates = reference.Animates[:0]
	for _, _b := range instance.Animates {
		reference.Animates = append(reference.Animates, stage.Animates_reference[_b])
	}
}

func (reference *Polyline) GongReconstructPointersFromReferences(stage *Stage, instance *Polyline) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Animates = reference.Animates[:0]
	for _, _b := range instance.Animates {
		reference.Animates = append(reference.Animates, stage.Animates_reference[_b])
	}
}

func (reference *Rect) GongReconstructPointersFromReferences(stage *Stage, instance *Rect) {
	// insertion point for pointers field
	if instance.EnclosingRect != nil {
		reference.EnclosingRect = stage.Rects_reference[instance.EnclosingRect]
	}
	if instance.AnchoredTo != nil {
		reference.AnchoredTo = stage.Rects_reference[instance.AnchoredTo]
	}
	// insertion point for slice of pointers field
	reference.Peers = reference.Peers[:0]
	for _, _b := range instance.Peers {
		reference.Peers = append(reference.Peers, stage.Rects_reference[_b])
	}
	reference.Obstacles = reference.Obstacles[:0]
	for _, _b := range instance.Obstacles {
		reference.Obstacles = append(reference.Obstacles, stage.Rects_reference[_b])
	}
	reference.HoveringTrigger = reference.HoveringTrigger[:0]
	for _, _b := range instance.HoveringTrigger {
		reference.HoveringTrigger = append(reference.HoveringTrigger, stage.Conditions_reference[_b])
	}
	reference.DisplayConditions = reference.DisplayConditions[:0]
	for _, _b := range instance.DisplayConditions {
		reference.DisplayConditions = append(reference.DisplayConditions, stage.Conditions_reference[_b])
	}
	reference.Animations = reference.Animations[:0]
	for _, _b := range instance.Animations {
		reference.Animations = append(reference.Animations, stage.Animates_reference[_b])
	}
	reference.RectAnchoredTexts = reference.RectAnchoredTexts[:0]
	for _, _b := range instance.RectAnchoredTexts {
		reference.RectAnchoredTexts = append(reference.RectAnchoredTexts, stage.RectAnchoredTexts_reference[_b])
	}
	reference.RectAnchoredRects = reference.RectAnchoredRects[:0]
	for _, _b := range instance.RectAnchoredRects {
		reference.RectAnchoredRects = append(reference.RectAnchoredRects, stage.RectAnchoredRects_reference[_b])
	}
	reference.RectAnchoredPaths = reference.RectAnchoredPaths[:0]
	for _, _b := range instance.RectAnchoredPaths {
		reference.RectAnchoredPaths = append(reference.RectAnchoredPaths, stage.RectAnchoredPaths_reference[_b])
	}
	reference.RectAnchoredPngImages = reference.RectAnchoredPngImages[:0]
	for _, _b := range instance.RectAnchoredPngImages {
		reference.RectAnchoredPngImages = append(reference.RectAnchoredPngImages, stage.RectAnchoredPngImages_reference[_b])
	}
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
	reference.Animates = reference.Animates[:0]
	for _, _b := range instance.Animates {
		reference.Animates = append(reference.Animates, stage.Animates_reference[_b])
	}
}

func (reference *RectLinkLink) GongReconstructPointersFromReferences(stage *Stage, instance *RectLinkLink) {
	// insertion point for pointers field
	if instance.Start != nil {
		reference.Start = stage.Rects_reference[instance.Start]
	}
	if instance.End != nil {
		reference.End = stage.Links_reference[instance.End]
	}
	// insertion point for slice of pointers field
}

func (reference *SVG) GongReconstructPointersFromReferences(stage *Stage, instance *SVG) {
	// insertion point for pointers field
	if instance.StartRect != nil {
		reference.StartRect = stage.Rects_reference[instance.StartRect]
	}
	if instance.EndRect != nil {
		reference.EndRect = stage.Rects_reference[instance.EndRect]
	}
	// insertion point for slice of pointers field
	reference.Layers = reference.Layers[:0]
	for _, _b := range instance.Layers {
		reference.Layers = append(reference.Layers, stage.Layers_reference[_b])
	}
}

func (reference *SvgText) GongReconstructPointersFromReferences(stage *Stage, instance *SvgText) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Text) GongReconstructPointersFromReferences(stage *Stage, instance *Text) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Animates = reference.Animates[:0]
	for _, _b := range instance.Animates {
		reference.Animates = append(reference.Animates, stage.Animates_reference[_b])
	}
}

// insertion point for pointer reconstruction from instances
func (reference *Animate) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Circle) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Animations []*Animate
	for _, _reference := range reference.Animations {
		if _instance, ok := stage.Animates_instance[_reference]; ok {
			_Animations = append(_Animations, _instance)
		}
	}
	reference.Animations = _Animations
}

func (reference *Condition) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ControlPoint) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ClosestRect; _reference != nil {
		reference.ClosestRect = nil
		if _instance, ok := stage.Rects_instance[_reference]; ok {
			reference.ClosestRect = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Ellipse) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Animates []*Animate
	for _, _reference := range reference.Animates {
		if _instance, ok := stage.Animates_instance[_reference]; ok {
			_Animates = append(_Animates, _instance)
		}
	}
	reference.Animates = _Animates
}

func (reference *FileToDownload) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Layer) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Rects []*Rect
	for _, _reference := range reference.Rects {
		if _instance, ok := stage.Rects_instance[_reference]; ok {
			_Rects = append(_Rects, _instance)
		}
	}
	reference.Rects = _Rects
	var _Texts []*Text
	for _, _reference := range reference.Texts {
		if _instance, ok := stage.Texts_instance[_reference]; ok {
			_Texts = append(_Texts, _instance)
		}
	}
	reference.Texts = _Texts
	var _Circles []*Circle
	for _, _reference := range reference.Circles {
		if _instance, ok := stage.Circles_instance[_reference]; ok {
			_Circles = append(_Circles, _instance)
		}
	}
	reference.Circles = _Circles
	var _Lines []*Line
	for _, _reference := range reference.Lines {
		if _instance, ok := stage.Lines_instance[_reference]; ok {
			_Lines = append(_Lines, _instance)
		}
	}
	reference.Lines = _Lines
	var _Ellipses []*Ellipse
	for _, _reference := range reference.Ellipses {
		if _instance, ok := stage.Ellipses_instance[_reference]; ok {
			_Ellipses = append(_Ellipses, _instance)
		}
	}
	reference.Ellipses = _Ellipses
	var _Polylines []*Polyline
	for _, _reference := range reference.Polylines {
		if _instance, ok := stage.Polylines_instance[_reference]; ok {
			_Polylines = append(_Polylines, _instance)
		}
	}
	reference.Polylines = _Polylines
	var _Polygones []*Polygone
	for _, _reference := range reference.Polygones {
		if _instance, ok := stage.Polygones_instance[_reference]; ok {
			_Polygones = append(_Polygones, _instance)
		}
	}
	reference.Polygones = _Polygones
	var _Paths []*Path
	for _, _reference := range reference.Paths {
		if _instance, ok := stage.Paths_instance[_reference]; ok {
			_Paths = append(_Paths, _instance)
		}
	}
	reference.Paths = _Paths
	var _Links []*Link
	for _, _reference := range reference.Links {
		if _instance, ok := stage.Links_instance[_reference]; ok {
			_Links = append(_Links, _instance)
		}
	}
	reference.Links = _Links
	var _RectLinkLinks []*RectLinkLink
	for _, _reference := range reference.RectLinkLinks {
		if _instance, ok := stage.RectLinkLinks_instance[_reference]; ok {
			_RectLinkLinks = append(_RectLinkLinks, _instance)
		}
	}
	reference.RectLinkLinks = _RectLinkLinks
}

func (reference *Line) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Animates []*Animate
	for _, _reference := range reference.Animates {
		if _instance, ok := stage.Animates_instance[_reference]; ok {
			_Animates = append(_Animates, _instance)
		}
	}
	reference.Animates = _Animates
}

func (reference *Link) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Start; _reference != nil {
		reference.Start = nil
		if _instance, ok := stage.Rects_instance[_reference]; ok {
			reference.Start = _instance
		}
	}
	if _reference := reference.End; _reference != nil {
		reference.End = nil
		if _instance, ok := stage.Rects_instance[_reference]; ok {
			reference.End = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _TextAtArrowStart []*LinkAnchoredText
	for _, _reference := range reference.TextAtArrowStart {
		if _instance, ok := stage.LinkAnchoredTexts_instance[_reference]; ok {
			_TextAtArrowStart = append(_TextAtArrowStart, _instance)
		}
	}
	reference.TextAtArrowStart = _TextAtArrowStart
	var _TextAtArrowEnd []*LinkAnchoredText
	for _, _reference := range reference.TextAtArrowEnd {
		if _instance, ok := stage.LinkAnchoredTexts_instance[_reference]; ok {
			_TextAtArrowEnd = append(_TextAtArrowEnd, _instance)
		}
	}
	reference.TextAtArrowEnd = _TextAtArrowEnd
	var _TextAtCorner []*LinkAnchoredText
	for _, _reference := range reference.TextAtCorner {
		if _instance, ok := stage.LinkAnchoredTexts_instance[_reference]; ok {
			_TextAtCorner = append(_TextAtCorner, _instance)
		}
	}
	reference.TextAtCorner = _TextAtCorner
	var _PathAtArrowStart []*LinkAnchoredPath
	for _, _reference := range reference.PathAtArrowStart {
		if _instance, ok := stage.LinkAnchoredPaths_instance[_reference]; ok {
			_PathAtArrowStart = append(_PathAtArrowStart, _instance)
		}
	}
	reference.PathAtArrowStart = _PathAtArrowStart
	var _PathAtArrowEnd []*LinkAnchoredPath
	for _, _reference := range reference.PathAtArrowEnd {
		if _instance, ok := stage.LinkAnchoredPaths_instance[_reference]; ok {
			_PathAtArrowEnd = append(_PathAtArrowEnd, _instance)
		}
	}
	reference.PathAtArrowEnd = _PathAtArrowEnd
	var _PathAtCorner []*LinkAnchoredPath
	for _, _reference := range reference.PathAtCorner {
		if _instance, ok := stage.LinkAnchoredPaths_instance[_reference]; ok {
			_PathAtCorner = append(_PathAtCorner, _instance)
		}
	}
	reference.PathAtCorner = _PathAtCorner
	var _ControlPoints []*ControlPoint
	for _, _reference := range reference.ControlPoints {
		if _instance, ok := stage.ControlPoints_instance[_reference]; ok {
			_ControlPoints = append(_ControlPoints, _instance)
		}
	}
	reference.ControlPoints = _ControlPoints
}

func (reference *LinkAnchoredPath) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *LinkAnchoredText) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Animates []*Animate
	for _, _reference := range reference.Animates {
		if _instance, ok := stage.Animates_instance[_reference]; ok {
			_Animates = append(_Animates, _instance)
		}
	}
	reference.Animates = _Animates
}

func (reference *Path) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Animates []*Animate
	for _, _reference := range reference.Animates {
		if _instance, ok := stage.Animates_instance[_reference]; ok {
			_Animates = append(_Animates, _instance)
		}
	}
	reference.Animates = _Animates
}

func (reference *Point) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Polygone) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Animates []*Animate
	for _, _reference := range reference.Animates {
		if _instance, ok := stage.Animates_instance[_reference]; ok {
			_Animates = append(_Animates, _instance)
		}
	}
	reference.Animates = _Animates
}

func (reference *Polyline) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Animates []*Animate
	for _, _reference := range reference.Animates {
		if _instance, ok := stage.Animates_instance[_reference]; ok {
			_Animates = append(_Animates, _instance)
		}
	}
	reference.Animates = _Animates
}

func (reference *Rect) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.EnclosingRect; _reference != nil {
		reference.EnclosingRect = nil
		if _instance, ok := stage.Rects_instance[_reference]; ok {
			reference.EnclosingRect = _instance
		}
	}
	if _reference := reference.AnchoredTo; _reference != nil {
		reference.AnchoredTo = nil
		if _instance, ok := stage.Rects_instance[_reference]; ok {
			reference.AnchoredTo = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Peers []*Rect
	for _, _reference := range reference.Peers {
		if _instance, ok := stage.Rects_instance[_reference]; ok {
			_Peers = append(_Peers, _instance)
		}
	}
	reference.Peers = _Peers
	var _Obstacles []*Rect
	for _, _reference := range reference.Obstacles {
		if _instance, ok := stage.Rects_instance[_reference]; ok {
			_Obstacles = append(_Obstacles, _instance)
		}
	}
	reference.Obstacles = _Obstacles
	var _HoveringTrigger []*Condition
	for _, _reference := range reference.HoveringTrigger {
		if _instance, ok := stage.Conditions_instance[_reference]; ok {
			_HoveringTrigger = append(_HoveringTrigger, _instance)
		}
	}
	reference.HoveringTrigger = _HoveringTrigger
	var _DisplayConditions []*Condition
	for _, _reference := range reference.DisplayConditions {
		if _instance, ok := stage.Conditions_instance[_reference]; ok {
			_DisplayConditions = append(_DisplayConditions, _instance)
		}
	}
	reference.DisplayConditions = _DisplayConditions
	var _Animations []*Animate
	for _, _reference := range reference.Animations {
		if _instance, ok := stage.Animates_instance[_reference]; ok {
			_Animations = append(_Animations, _instance)
		}
	}
	reference.Animations = _Animations
	var _RectAnchoredTexts []*RectAnchoredText
	for _, _reference := range reference.RectAnchoredTexts {
		if _instance, ok := stage.RectAnchoredTexts_instance[_reference]; ok {
			_RectAnchoredTexts = append(_RectAnchoredTexts, _instance)
		}
	}
	reference.RectAnchoredTexts = _RectAnchoredTexts
	var _RectAnchoredRects []*RectAnchoredRect
	for _, _reference := range reference.RectAnchoredRects {
		if _instance, ok := stage.RectAnchoredRects_instance[_reference]; ok {
			_RectAnchoredRects = append(_RectAnchoredRects, _instance)
		}
	}
	reference.RectAnchoredRects = _RectAnchoredRects
	var _RectAnchoredPaths []*RectAnchoredPath
	for _, _reference := range reference.RectAnchoredPaths {
		if _instance, ok := stage.RectAnchoredPaths_instance[_reference]; ok {
			_RectAnchoredPaths = append(_RectAnchoredPaths, _instance)
		}
	}
	reference.RectAnchoredPaths = _RectAnchoredPaths
	var _RectAnchoredPngImages []*RectAnchoredPngImage
	for _, _reference := range reference.RectAnchoredPngImages {
		if _instance, ok := stage.RectAnchoredPngImages_instance[_reference]; ok {
			_RectAnchoredPngImages = append(_RectAnchoredPngImages, _instance)
		}
	}
	reference.RectAnchoredPngImages = _RectAnchoredPngImages
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
	var _Animates []*Animate
	for _, _reference := range reference.Animates {
		if _instance, ok := stage.Animates_instance[_reference]; ok {
			_Animates = append(_Animates, _instance)
		}
	}
	reference.Animates = _Animates
}

func (reference *RectLinkLink) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Start; _reference != nil {
		reference.Start = nil
		if _instance, ok := stage.Rects_instance[_reference]; ok {
			reference.Start = _instance
		}
	}
	if _reference := reference.End; _reference != nil {
		reference.End = nil
		if _instance, ok := stage.Links_instance[_reference]; ok {
			reference.End = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *SVG) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.StartRect; _reference != nil {
		reference.StartRect = nil
		if _instance, ok := stage.Rects_instance[_reference]; ok {
			reference.StartRect = _instance
		}
	}
	if _reference := reference.EndRect; _reference != nil {
		reference.EndRect = nil
		if _instance, ok := stage.Rects_instance[_reference]; ok {
			reference.EndRect = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Layers []*Layer
	for _, _reference := range reference.Layers {
		if _instance, ok := stage.Layers_instance[_reference]; ok {
			_Layers = append(_Layers, _instance)
		}
	}
	reference.Layers = _Layers
}

func (reference *SvgText) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Text) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Animates []*Animate
	for _, _reference := range reference.Animates {
		if _instance, ok := stage.Animates_instance[_reference]; ok {
			_Animates = append(_Animates, _instance)
		}
	}
	reference.Animates = _Animates
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
		ops := stage.Diff(circle, circleOther, "Animations", circleOther.Animations, circle.Animations)
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
		ops := stage.Diff(ellipse, ellipseOther, "Animates", ellipseOther.Animates, ellipse.Animates)
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
		ops := stage.Diff(layer, layerOther, "Rects", layerOther.Rects, layer.Rects)
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
		ops := stage.Diff(layer, layerOther, "Texts", layerOther.Texts, layer.Texts)
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
		ops := stage.Diff(layer, layerOther, "Circles", layerOther.Circles, layer.Circles)
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
		ops := stage.Diff(layer, layerOther, "Lines", layerOther.Lines, layer.Lines)
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
		ops := stage.Diff(layer, layerOther, "Ellipses", layerOther.Ellipses, layer.Ellipses)
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
		ops := stage.Diff(layer, layerOther, "Polylines", layerOther.Polylines, layer.Polylines)
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
		ops := stage.Diff(layer, layerOther, "Polygones", layerOther.Polygones, layer.Polygones)
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
		ops := stage.Diff(layer, layerOther, "Paths", layerOther.Paths, layer.Paths)
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
		ops := stage.Diff(layer, layerOther, "Links", layerOther.Links, layer.Links)
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
		ops := stage.Diff(layer, layerOther, "RectLinkLinks", layerOther.RectLinkLinks, layer.RectLinkLinks)
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
		ops := stage.Diff(line, lineOther, "Animates", lineOther.Animates, line.Animates)
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
		ops := stage.Diff(link, linkOther, "TextAtArrowStart", linkOther.TextAtArrowStart, link.TextAtArrowStart)
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
		ops := stage.Diff(link, linkOther, "TextAtArrowEnd", linkOther.TextAtArrowEnd, link.TextAtArrowEnd)
		diffs = append(diffs, ops)
	}
	TextAtCornerDifferent := false
	if len(link.TextAtCorner) != len(linkOther.TextAtCorner) {
		TextAtCornerDifferent = true
	} else {
		for i := range link.TextAtCorner {
			if (link.TextAtCorner[i] == nil) != (linkOther.TextAtCorner[i] == nil) {
				TextAtCornerDifferent = true
				break
			} else if link.TextAtCorner[i] != nil && linkOther.TextAtCorner[i] != nil {
				// this is a pointer comparaison
				if link.TextAtCorner[i] != linkOther.TextAtCorner[i] {
					TextAtCornerDifferent = true
					break
				}
			}
		}
	}
	if TextAtCornerDifferent {
		ops := stage.Diff(link, linkOther, "TextAtCorner", linkOther.TextAtCorner, link.TextAtCorner)
		diffs = append(diffs, ops)
	}
	PathAtArrowStartDifferent := false
	if len(link.PathAtArrowStart) != len(linkOther.PathAtArrowStart) {
		PathAtArrowStartDifferent = true
	} else {
		for i := range link.PathAtArrowStart {
			if (link.PathAtArrowStart[i] == nil) != (linkOther.PathAtArrowStart[i] == nil) {
				PathAtArrowStartDifferent = true
				break
			} else if link.PathAtArrowStart[i] != nil && linkOther.PathAtArrowStart[i] != nil {
				// this is a pointer comparaison
				if link.PathAtArrowStart[i] != linkOther.PathAtArrowStart[i] {
					PathAtArrowStartDifferent = true
					break
				}
			}
		}
	}
	if PathAtArrowStartDifferent {
		ops := stage.Diff(link, linkOther, "PathAtArrowStart", linkOther.PathAtArrowStart, link.PathAtArrowStart)
		diffs = append(diffs, ops)
	}
	PathAtArrowEndDifferent := false
	if len(link.PathAtArrowEnd) != len(linkOther.PathAtArrowEnd) {
		PathAtArrowEndDifferent = true
	} else {
		for i := range link.PathAtArrowEnd {
			if (link.PathAtArrowEnd[i] == nil) != (linkOther.PathAtArrowEnd[i] == nil) {
				PathAtArrowEndDifferent = true
				break
			} else if link.PathAtArrowEnd[i] != nil && linkOther.PathAtArrowEnd[i] != nil {
				// this is a pointer comparaison
				if link.PathAtArrowEnd[i] != linkOther.PathAtArrowEnd[i] {
					PathAtArrowEndDifferent = true
					break
				}
			}
		}
	}
	if PathAtArrowEndDifferent {
		ops := stage.Diff(link, linkOther, "PathAtArrowEnd", linkOther.PathAtArrowEnd, link.PathAtArrowEnd)
		diffs = append(diffs, ops)
	}
	PathAtCornerDifferent := false
	if len(link.PathAtCorner) != len(linkOther.PathAtCorner) {
		PathAtCornerDifferent = true
	} else {
		for i := range link.PathAtCorner {
			if (link.PathAtCorner[i] == nil) != (linkOther.PathAtCorner[i] == nil) {
				PathAtCornerDifferent = true
				break
			} else if link.PathAtCorner[i] != nil && linkOther.PathAtCorner[i] != nil {
				// this is a pointer comparaison
				if link.PathAtCorner[i] != linkOther.PathAtCorner[i] {
					PathAtCornerDifferent = true
					break
				}
			}
		}
	}
	if PathAtCornerDifferent {
		ops := stage.Diff(link, linkOther, "PathAtCorner", linkOther.PathAtCorner, link.PathAtCorner)
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
		ops := stage.Diff(link, linkOther, "ControlPoints", linkOther.ControlPoints, link.ControlPoints)
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
		ops := stage.Diff(linkanchoredtext, linkanchoredtextOther, "Animates", linkanchoredtextOther.Animates, linkanchoredtext.Animates)
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
		ops := stage.Diff(path, pathOther, "Animates", pathOther.Animates, path.Animates)
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
		ops := stage.Diff(polygone, polygoneOther, "Animates", polygoneOther.Animates, polygone.Animates)
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
		ops := stage.Diff(polyline, polylineOther, "Animates", polylineOther.Animates, polyline.Animates)
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
	PeersDifferent := false
	if len(rect.Peers) != len(rectOther.Peers) {
		PeersDifferent = true
	} else {
		for i := range rect.Peers {
			if (rect.Peers[i] == nil) != (rectOther.Peers[i] == nil) {
				PeersDifferent = true
				break
			} else if rect.Peers[i] != nil && rectOther.Peers[i] != nil {
				// this is a pointer comparaison
				if rect.Peers[i] != rectOther.Peers[i] {
					PeersDifferent = true
					break
				}
			}
		}
	}
	if PeersDifferent {
		ops := stage.Diff(rect, rectOther, "Peers", rectOther.Peers, rect.Peers)
		diffs = append(diffs, ops)
	}
	if (rect.EnclosingRect == nil) != (rectOther.EnclosingRect == nil) {
		diffs = append(diffs, rect.GongMarshallField(stage, "EnclosingRect"))
	} else if rect.EnclosingRect != nil && rectOther.EnclosingRect != nil {
		if rect.EnclosingRect != rectOther.EnclosingRect {
			diffs = append(diffs, rect.GongMarshallField(stage, "EnclosingRect"))
		}
	}
	ObstaclesDifferent := false
	if len(rect.Obstacles) != len(rectOther.Obstacles) {
		ObstaclesDifferent = true
	} else {
		for i := range rect.Obstacles {
			if (rect.Obstacles[i] == nil) != (rectOther.Obstacles[i] == nil) {
				ObstaclesDifferent = true
				break
			} else if rect.Obstacles[i] != nil && rectOther.Obstacles[i] != nil {
				// this is a pointer comparaison
				if rect.Obstacles[i] != rectOther.Obstacles[i] {
					ObstaclesDifferent = true
					break
				}
			}
		}
	}
	if ObstaclesDifferent {
		ops := stage.Diff(rect, rectOther, "Obstacles", rectOther.Obstacles, rect.Obstacles)
		diffs = append(diffs, ops)
	}
	if (rect.AnchoredTo == nil) != (rectOther.AnchoredTo == nil) {
		diffs = append(diffs, rect.GongMarshallField(stage, "AnchoredTo"))
	} else if rect.AnchoredTo != nil && rectOther.AnchoredTo != nil {
		if rect.AnchoredTo != rectOther.AnchoredTo {
			diffs = append(diffs, rect.GongMarshallField(stage, "AnchoredTo"))
		}
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
		ops := stage.Diff(rect, rectOther, "HoveringTrigger", rectOther.HoveringTrigger, rect.HoveringTrigger)
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
		ops := stage.Diff(rect, rectOther, "DisplayConditions", rectOther.DisplayConditions, rect.DisplayConditions)
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
		ops := stage.Diff(rect, rectOther, "Animations", rectOther.Animations, rect.Animations)
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
		ops := stage.Diff(rect, rectOther, "RectAnchoredTexts", rectOther.RectAnchoredTexts, rect.RectAnchoredTexts)
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
		ops := stage.Diff(rect, rectOther, "RectAnchoredRects", rectOther.RectAnchoredRects, rect.RectAnchoredRects)
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
		ops := stage.Diff(rect, rectOther, "RectAnchoredPaths", rectOther.RectAnchoredPaths, rect.RectAnchoredPaths)
		diffs = append(diffs, ops)
	}
	RectAnchoredPngImagesDifferent := false
	if len(rect.RectAnchoredPngImages) != len(rectOther.RectAnchoredPngImages) {
		RectAnchoredPngImagesDifferent = true
	} else {
		for i := range rect.RectAnchoredPngImages {
			if (rect.RectAnchoredPngImages[i] == nil) != (rectOther.RectAnchoredPngImages[i] == nil) {
				RectAnchoredPngImagesDifferent = true
				break
			} else if rect.RectAnchoredPngImages[i] != nil && rectOther.RectAnchoredPngImages[i] != nil {
				// this is a pointer comparaison
				if rect.RectAnchoredPngImages[i] != rectOther.RectAnchoredPngImages[i] {
					RectAnchoredPngImagesDifferent = true
					break
				}
			}
		}
	}
	if RectAnchoredPngImagesDifferent {
		ops := stage.Diff(rect, rectOther, "RectAnchoredPngImages", rectOther.RectAnchoredPngImages, rect.RectAnchoredPngImages)
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
		ops := stage.Diff(rectanchoredtext, rectanchoredtextOther, "Animates", rectanchoredtextOther.Animates, rectanchoredtext.Animates)
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
		ops := stage.Diff(svg, svgOther, "Layers", svgOther.Layers, svg.Layers)
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
		ops := stage.Diff(text, textOther, "Animates", textOther.Animates, text.Animates)
		diffs = append(diffs, ops)
	}

	return
}

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff[T1, T2 PointerToGongstruct](a, b T1, fieldName string, oldSlice, newSlice []T2) (ops string) {
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
