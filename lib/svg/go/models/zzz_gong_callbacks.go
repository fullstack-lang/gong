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
func (animate *Animate) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAnimateCreateCallback != nil {
		stage.OnAfterAnimateCreateCallback.OnAfterCreate(stage, animate)
	}
}

func (animate *Animate) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAnimateUpdateCallback != nil {
		var frontAnimate *Animate
		if front != nil {
			frontAnimate, _ = front.(*Animate)
		}
		stage.OnAfterAnimateUpdateCallback.OnAfterUpdate(stage, animate, frontAnimate)
	}
}

func (animate *Animate) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAnimateDeleteCallback != nil {
		var frontAnimate *Animate
		if front != nil {
			frontAnimate, _ = front.(*Animate)
		}
		stage.OnAfterAnimateDeleteCallback.OnAfterDelete(stage, animate, frontAnimate)
	}
}

func (circle *Circle) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCircleCreateCallback != nil {
		stage.OnAfterCircleCreateCallback.OnAfterCreate(stage, circle)
	}
}

func (circle *Circle) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCircleUpdateCallback != nil {
		var frontCircle *Circle
		if front != nil {
			frontCircle, _ = front.(*Circle)
		}
		stage.OnAfterCircleUpdateCallback.OnAfterUpdate(stage, circle, frontCircle)
	}
}

func (circle *Circle) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCircleDeleteCallback != nil {
		var frontCircle *Circle
		if front != nil {
			frontCircle, _ = front.(*Circle)
		}
		stage.OnAfterCircleDeleteCallback.OnAfterDelete(stage, circle, frontCircle)
	}
}

func (condition *Condition) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterConditionCreateCallback != nil {
		stage.OnAfterConditionCreateCallback.OnAfterCreate(stage, condition)
	}
}

func (condition *Condition) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConditionUpdateCallback != nil {
		var frontCondition *Condition
		if front != nil {
			frontCondition, _ = front.(*Condition)
		}
		stage.OnAfterConditionUpdateCallback.OnAfterUpdate(stage, condition, frontCondition)
	}
}

func (condition *Condition) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConditionDeleteCallback != nil {
		var frontCondition *Condition
		if front != nil {
			frontCondition, _ = front.(*Condition)
		}
		stage.OnAfterConditionDeleteCallback.OnAfterDelete(stage, condition, frontCondition)
	}
}

func (controlpoint *ControlPoint) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterControlPointCreateCallback != nil {
		stage.OnAfterControlPointCreateCallback.OnAfterCreate(stage, controlpoint)
	}
}

func (controlpoint *ControlPoint) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterControlPointUpdateCallback != nil {
		var frontControlPoint *ControlPoint
		if front != nil {
			frontControlPoint, _ = front.(*ControlPoint)
		}
		stage.OnAfterControlPointUpdateCallback.OnAfterUpdate(stage, controlpoint, frontControlPoint)
	}
}

func (controlpoint *ControlPoint) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterControlPointDeleteCallback != nil {
		var frontControlPoint *ControlPoint
		if front != nil {
			frontControlPoint, _ = front.(*ControlPoint)
		}
		stage.OnAfterControlPointDeleteCallback.OnAfterDelete(stage, controlpoint, frontControlPoint)
	}
}

func (ellipse *Ellipse) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEllipseCreateCallback != nil {
		stage.OnAfterEllipseCreateCallback.OnAfterCreate(stage, ellipse)
	}
}

func (ellipse *Ellipse) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEllipseUpdateCallback != nil {
		var frontEllipse *Ellipse
		if front != nil {
			frontEllipse, _ = front.(*Ellipse)
		}
		stage.OnAfterEllipseUpdateCallback.OnAfterUpdate(stage, ellipse, frontEllipse)
	}
}

func (ellipse *Ellipse) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEllipseDeleteCallback != nil {
		var frontEllipse *Ellipse
		if front != nil {
			frontEllipse, _ = front.(*Ellipse)
		}
		stage.OnAfterEllipseDeleteCallback.OnAfterDelete(stage, ellipse, frontEllipse)
	}
}

func (filetodownload *FileToDownload) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFileToDownloadCreateCallback != nil {
		stage.OnAfterFileToDownloadCreateCallback.OnAfterCreate(stage, filetodownload)
	}
}

func (filetodownload *FileToDownload) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFileToDownloadUpdateCallback != nil {
		var frontFileToDownload *FileToDownload
		if front != nil {
			frontFileToDownload, _ = front.(*FileToDownload)
		}
		stage.OnAfterFileToDownloadUpdateCallback.OnAfterUpdate(stage, filetodownload, frontFileToDownload)
	}
}

func (filetodownload *FileToDownload) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFileToDownloadDeleteCallback != nil {
		var frontFileToDownload *FileToDownload
		if front != nil {
			frontFileToDownload, _ = front.(*FileToDownload)
		}
		stage.OnAfterFileToDownloadDeleteCallback.OnAfterDelete(stage, filetodownload, frontFileToDownload)
	}
}

func (layer *Layer) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLayerCreateCallback != nil {
		stage.OnAfterLayerCreateCallback.OnAfterCreate(stage, layer)
	}
}

func (layer *Layer) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLayerUpdateCallback != nil {
		var frontLayer *Layer
		if front != nil {
			frontLayer, _ = front.(*Layer)
		}
		stage.OnAfterLayerUpdateCallback.OnAfterUpdate(stage, layer, frontLayer)
	}
}

func (layer *Layer) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLayerDeleteCallback != nil {
		var frontLayer *Layer
		if front != nil {
			frontLayer, _ = front.(*Layer)
		}
		stage.OnAfterLayerDeleteCallback.OnAfterDelete(stage, layer, frontLayer)
	}
}

func (line *Line) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLineCreateCallback != nil {
		stage.OnAfterLineCreateCallback.OnAfterCreate(stage, line)
	}
}

func (line *Line) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLineUpdateCallback != nil {
		var frontLine *Line
		if front != nil {
			frontLine, _ = front.(*Line)
		}
		stage.OnAfterLineUpdateCallback.OnAfterUpdate(stage, line, frontLine)
	}
}

func (line *Line) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLineDeleteCallback != nil {
		var frontLine *Line
		if front != nil {
			frontLine, _ = front.(*Line)
		}
		stage.OnAfterLineDeleteCallback.OnAfterDelete(stage, line, frontLine)
	}
}

func (link *Link) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLinkCreateCallback != nil {
		stage.OnAfterLinkCreateCallback.OnAfterCreate(stage, link)
	}
}

func (link *Link) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinkUpdateCallback != nil {
		var frontLink *Link
		if front != nil {
			frontLink, _ = front.(*Link)
		}
		stage.OnAfterLinkUpdateCallback.OnAfterUpdate(stage, link, frontLink)
	}
}

func (link *Link) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinkDeleteCallback != nil {
		var frontLink *Link
		if front != nil {
			frontLink, _ = front.(*Link)
		}
		stage.OnAfterLinkDeleteCallback.OnAfterDelete(stage, link, frontLink)
	}
}

func (linkanchoredpath *LinkAnchoredPath) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLinkAnchoredPathCreateCallback != nil {
		stage.OnAfterLinkAnchoredPathCreateCallback.OnAfterCreate(stage, linkanchoredpath)
	}
}

func (linkanchoredpath *LinkAnchoredPath) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinkAnchoredPathUpdateCallback != nil {
		var frontLinkAnchoredPath *LinkAnchoredPath
		if front != nil {
			frontLinkAnchoredPath, _ = front.(*LinkAnchoredPath)
		}
		stage.OnAfterLinkAnchoredPathUpdateCallback.OnAfterUpdate(stage, linkanchoredpath, frontLinkAnchoredPath)
	}
}

func (linkanchoredpath *LinkAnchoredPath) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinkAnchoredPathDeleteCallback != nil {
		var frontLinkAnchoredPath *LinkAnchoredPath
		if front != nil {
			frontLinkAnchoredPath, _ = front.(*LinkAnchoredPath)
		}
		stage.OnAfterLinkAnchoredPathDeleteCallback.OnAfterDelete(stage, linkanchoredpath, frontLinkAnchoredPath)
	}
}

func (linkanchoredtext *LinkAnchoredText) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLinkAnchoredTextCreateCallback != nil {
		stage.OnAfterLinkAnchoredTextCreateCallback.OnAfterCreate(stage, linkanchoredtext)
	}
}

func (linkanchoredtext *LinkAnchoredText) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinkAnchoredTextUpdateCallback != nil {
		var frontLinkAnchoredText *LinkAnchoredText
		if front != nil {
			frontLinkAnchoredText, _ = front.(*LinkAnchoredText)
		}
		stage.OnAfterLinkAnchoredTextUpdateCallback.OnAfterUpdate(stage, linkanchoredtext, frontLinkAnchoredText)
	}
}

func (linkanchoredtext *LinkAnchoredText) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinkAnchoredTextDeleteCallback != nil {
		var frontLinkAnchoredText *LinkAnchoredText
		if front != nil {
			frontLinkAnchoredText, _ = front.(*LinkAnchoredText)
		}
		stage.OnAfterLinkAnchoredTextDeleteCallback.OnAfterDelete(stage, linkanchoredtext, frontLinkAnchoredText)
	}
}

func (path *Path) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPathCreateCallback != nil {
		stage.OnAfterPathCreateCallback.OnAfterCreate(stage, path)
	}
}

func (path *Path) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPathUpdateCallback != nil {
		var frontPath *Path
		if front != nil {
			frontPath, _ = front.(*Path)
		}
		stage.OnAfterPathUpdateCallback.OnAfterUpdate(stage, path, frontPath)
	}
}

func (path *Path) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPathDeleteCallback != nil {
		var frontPath *Path
		if front != nil {
			frontPath, _ = front.(*Path)
		}
		stage.OnAfterPathDeleteCallback.OnAfterDelete(stage, path, frontPath)
	}
}

func (point *Point) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPointCreateCallback != nil {
		stage.OnAfterPointCreateCallback.OnAfterCreate(stage, point)
	}
}

func (point *Point) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointUpdateCallback != nil {
		var frontPoint *Point
		if front != nil {
			frontPoint, _ = front.(*Point)
		}
		stage.OnAfterPointUpdateCallback.OnAfterUpdate(stage, point, frontPoint)
	}
}

func (point *Point) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointDeleteCallback != nil {
		var frontPoint *Point
		if front != nil {
			frontPoint, _ = front.(*Point)
		}
		stage.OnAfterPointDeleteCallback.OnAfterDelete(stage, point, frontPoint)
	}
}

func (polygone *Polygone) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPolygoneCreateCallback != nil {
		stage.OnAfterPolygoneCreateCallback.OnAfterCreate(stage, polygone)
	}
}

func (polygone *Polygone) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPolygoneUpdateCallback != nil {
		var frontPolygone *Polygone
		if front != nil {
			frontPolygone, _ = front.(*Polygone)
		}
		stage.OnAfterPolygoneUpdateCallback.OnAfterUpdate(stage, polygone, frontPolygone)
	}
}

func (polygone *Polygone) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPolygoneDeleteCallback != nil {
		var frontPolygone *Polygone
		if front != nil {
			frontPolygone, _ = front.(*Polygone)
		}
		stage.OnAfterPolygoneDeleteCallback.OnAfterDelete(stage, polygone, frontPolygone)
	}
}

func (polyline *Polyline) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPolylineCreateCallback != nil {
		stage.OnAfterPolylineCreateCallback.OnAfterCreate(stage, polyline)
	}
}

func (polyline *Polyline) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPolylineUpdateCallback != nil {
		var frontPolyline *Polyline
		if front != nil {
			frontPolyline, _ = front.(*Polyline)
		}
		stage.OnAfterPolylineUpdateCallback.OnAfterUpdate(stage, polyline, frontPolyline)
	}
}

func (polyline *Polyline) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPolylineDeleteCallback != nil {
		var frontPolyline *Polyline
		if front != nil {
			frontPolyline, _ = front.(*Polyline)
		}
		stage.OnAfterPolylineDeleteCallback.OnAfterDelete(stage, polyline, frontPolyline)
	}
}

func (rect *Rect) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRectCreateCallback != nil {
		stage.OnAfterRectCreateCallback.OnAfterCreate(stage, rect)
	}
}

func (rect *Rect) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRectUpdateCallback != nil {
		var frontRect *Rect
		if front != nil {
			frontRect, _ = front.(*Rect)
		}
		stage.OnAfterRectUpdateCallback.OnAfterUpdate(stage, rect, frontRect)
	}
}

func (rect *Rect) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRectDeleteCallback != nil {
		var frontRect *Rect
		if front != nil {
			frontRect, _ = front.(*Rect)
		}
		stage.OnAfterRectDeleteCallback.OnAfterDelete(stage, rect, frontRect)
	}
}

func (rectanchoredpath *RectAnchoredPath) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRectAnchoredPathCreateCallback != nil {
		stage.OnAfterRectAnchoredPathCreateCallback.OnAfterCreate(stage, rectanchoredpath)
	}
}

func (rectanchoredpath *RectAnchoredPath) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRectAnchoredPathUpdateCallback != nil {
		var frontRectAnchoredPath *RectAnchoredPath
		if front != nil {
			frontRectAnchoredPath, _ = front.(*RectAnchoredPath)
		}
		stage.OnAfterRectAnchoredPathUpdateCallback.OnAfterUpdate(stage, rectanchoredpath, frontRectAnchoredPath)
	}
}

func (rectanchoredpath *RectAnchoredPath) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRectAnchoredPathDeleteCallback != nil {
		var frontRectAnchoredPath *RectAnchoredPath
		if front != nil {
			frontRectAnchoredPath, _ = front.(*RectAnchoredPath)
		}
		stage.OnAfterRectAnchoredPathDeleteCallback.OnAfterDelete(stage, rectanchoredpath, frontRectAnchoredPath)
	}
}

func (rectanchoredpngimage *RectAnchoredPngImage) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRectAnchoredPngImageCreateCallback != nil {
		stage.OnAfterRectAnchoredPngImageCreateCallback.OnAfterCreate(stage, rectanchoredpngimage)
	}
}

func (rectanchoredpngimage *RectAnchoredPngImage) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRectAnchoredPngImageUpdateCallback != nil {
		var frontRectAnchoredPngImage *RectAnchoredPngImage
		if front != nil {
			frontRectAnchoredPngImage, _ = front.(*RectAnchoredPngImage)
		}
		stage.OnAfterRectAnchoredPngImageUpdateCallback.OnAfterUpdate(stage, rectanchoredpngimage, frontRectAnchoredPngImage)
	}
}

func (rectanchoredpngimage *RectAnchoredPngImage) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRectAnchoredPngImageDeleteCallback != nil {
		var frontRectAnchoredPngImage *RectAnchoredPngImage
		if front != nil {
			frontRectAnchoredPngImage, _ = front.(*RectAnchoredPngImage)
		}
		stage.OnAfterRectAnchoredPngImageDeleteCallback.OnAfterDelete(stage, rectanchoredpngimage, frontRectAnchoredPngImage)
	}
}

func (rectanchoredrect *RectAnchoredRect) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRectAnchoredRectCreateCallback != nil {
		stage.OnAfterRectAnchoredRectCreateCallback.OnAfterCreate(stage, rectanchoredrect)
	}
}

func (rectanchoredrect *RectAnchoredRect) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRectAnchoredRectUpdateCallback != nil {
		var frontRectAnchoredRect *RectAnchoredRect
		if front != nil {
			frontRectAnchoredRect, _ = front.(*RectAnchoredRect)
		}
		stage.OnAfterRectAnchoredRectUpdateCallback.OnAfterUpdate(stage, rectanchoredrect, frontRectAnchoredRect)
	}
}

func (rectanchoredrect *RectAnchoredRect) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRectAnchoredRectDeleteCallback != nil {
		var frontRectAnchoredRect *RectAnchoredRect
		if front != nil {
			frontRectAnchoredRect, _ = front.(*RectAnchoredRect)
		}
		stage.OnAfterRectAnchoredRectDeleteCallback.OnAfterDelete(stage, rectanchoredrect, frontRectAnchoredRect)
	}
}

func (rectanchoredtext *RectAnchoredText) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRectAnchoredTextCreateCallback != nil {
		stage.OnAfterRectAnchoredTextCreateCallback.OnAfterCreate(stage, rectanchoredtext)
	}
}

func (rectanchoredtext *RectAnchoredText) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRectAnchoredTextUpdateCallback != nil {
		var frontRectAnchoredText *RectAnchoredText
		if front != nil {
			frontRectAnchoredText, _ = front.(*RectAnchoredText)
		}
		stage.OnAfterRectAnchoredTextUpdateCallback.OnAfterUpdate(stage, rectanchoredtext, frontRectAnchoredText)
	}
}

func (rectanchoredtext *RectAnchoredText) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRectAnchoredTextDeleteCallback != nil {
		var frontRectAnchoredText *RectAnchoredText
		if front != nil {
			frontRectAnchoredText, _ = front.(*RectAnchoredText)
		}
		stage.OnAfterRectAnchoredTextDeleteCallback.OnAfterDelete(stage, rectanchoredtext, frontRectAnchoredText)
	}
}

func (rectlinklink *RectLinkLink) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRectLinkLinkCreateCallback != nil {
		stage.OnAfterRectLinkLinkCreateCallback.OnAfterCreate(stage, rectlinklink)
	}
}

func (rectlinklink *RectLinkLink) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRectLinkLinkUpdateCallback != nil {
		var frontRectLinkLink *RectLinkLink
		if front != nil {
			frontRectLinkLink, _ = front.(*RectLinkLink)
		}
		stage.OnAfterRectLinkLinkUpdateCallback.OnAfterUpdate(stage, rectlinklink, frontRectLinkLink)
	}
}

func (rectlinklink *RectLinkLink) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRectLinkLinkDeleteCallback != nil {
		var frontRectLinkLink *RectLinkLink
		if front != nil {
			frontRectLinkLink, _ = front.(*RectLinkLink)
		}
		stage.OnAfterRectLinkLinkDeleteCallback.OnAfterDelete(stage, rectlinklink, frontRectLinkLink)
	}
}

func (svg *SVG) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSVGCreateCallback != nil {
		stage.OnAfterSVGCreateCallback.OnAfterCreate(stage, svg)
	}
}

func (svg *SVG) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSVGUpdateCallback != nil {
		var frontSVG *SVG
		if front != nil {
			frontSVG, _ = front.(*SVG)
		}
		stage.OnAfterSVGUpdateCallback.OnAfterUpdate(stage, svg, frontSVG)
	}
}

func (svg *SVG) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSVGDeleteCallback != nil {
		var frontSVG *SVG
		if front != nil {
			frontSVG, _ = front.(*SVG)
		}
		stage.OnAfterSVGDeleteCallback.OnAfterDelete(stage, svg, frontSVG)
	}
}

func (svgtext *SvgText) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSvgTextCreateCallback != nil {
		stage.OnAfterSvgTextCreateCallback.OnAfterCreate(stage, svgtext)
	}
}

func (svgtext *SvgText) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSvgTextUpdateCallback != nil {
		var frontSvgText *SvgText
		if front != nil {
			frontSvgText, _ = front.(*SvgText)
		}
		stage.OnAfterSvgTextUpdateCallback.OnAfterUpdate(stage, svgtext, frontSvgText)
	}
}

func (svgtext *SvgText) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSvgTextDeleteCallback != nil {
		var frontSvgText *SvgText
		if front != nil {
			frontSvgText, _ = front.(*SvgText)
		}
		stage.OnAfterSvgTextDeleteCallback.OnAfterDelete(stage, svgtext, frontSvgText)
	}
}

func (text *Text) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTextCreateCallback != nil {
		stage.OnAfterTextCreateCallback.OnAfterCreate(stage, text)
	}
}

func (text *Text) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextUpdateCallback != nil {
		var frontText *Text
		if front != nil {
			frontText, _ = front.(*Text)
		}
		stage.OnAfterTextUpdateCallback.OnAfterUpdate(stage, text, frontText)
	}
}

func (text *Text) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextDeleteCallback != nil {
		var frontText *Text
		if front != nil {
			frontText, _ = front.(*Text)
		}
		stage.OnAfterTextDeleteCallback.OnAfterDelete(stage, text, frontText)
	}
}

