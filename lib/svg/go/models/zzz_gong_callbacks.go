// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Animate:
		if stage.OnAfterAnimateCreateCallback != nil {
			stage.OnAfterAnimateCreateCallback.OnAfterCreate(stage, target)
		}
	case *Circle:
		if stage.OnAfterCircleCreateCallback != nil {
			stage.OnAfterCircleCreateCallback.OnAfterCreate(stage, target)
		}
	case *Condition:
		if stage.OnAfterConditionCreateCallback != nil {
			stage.OnAfterConditionCreateCallback.OnAfterCreate(stage, target)
		}
	case *ControlPoint:
		if stage.OnAfterControlPointCreateCallback != nil {
			stage.OnAfterControlPointCreateCallback.OnAfterCreate(stage, target)
		}
	case *Ellipse:
		if stage.OnAfterEllipseCreateCallback != nil {
			stage.OnAfterEllipseCreateCallback.OnAfterCreate(stage, target)
		}
	case *FileToDownload:
		if stage.OnAfterFileToDownloadCreateCallback != nil {
			stage.OnAfterFileToDownloadCreateCallback.OnAfterCreate(stage, target)
		}
	case *Layer:
		if stage.OnAfterLayerCreateCallback != nil {
			stage.OnAfterLayerCreateCallback.OnAfterCreate(stage, target)
		}
	case *Line:
		if stage.OnAfterLineCreateCallback != nil {
			stage.OnAfterLineCreateCallback.OnAfterCreate(stage, target)
		}
	case *Link:
		if stage.OnAfterLinkCreateCallback != nil {
			stage.OnAfterLinkCreateCallback.OnAfterCreate(stage, target)
		}
	case *LinkAnchoredPath:
		if stage.OnAfterLinkAnchoredPathCreateCallback != nil {
			stage.OnAfterLinkAnchoredPathCreateCallback.OnAfterCreate(stage, target)
		}
	case *LinkAnchoredText:
		if stage.OnAfterLinkAnchoredTextCreateCallback != nil {
			stage.OnAfterLinkAnchoredTextCreateCallback.OnAfterCreate(stage, target)
		}
	case *Path:
		if stage.OnAfterPathCreateCallback != nil {
			stage.OnAfterPathCreateCallback.OnAfterCreate(stage, target)
		}
	case *Point:
		if stage.OnAfterPointCreateCallback != nil {
			stage.OnAfterPointCreateCallback.OnAfterCreate(stage, target)
		}
	case *Polygone:
		if stage.OnAfterPolygoneCreateCallback != nil {
			stage.OnAfterPolygoneCreateCallback.OnAfterCreate(stage, target)
		}
	case *Polyline:
		if stage.OnAfterPolylineCreateCallback != nil {
			stage.OnAfterPolylineCreateCallback.OnAfterCreate(stage, target)
		}
	case *Rect:
		if stage.OnAfterRectCreateCallback != nil {
			stage.OnAfterRectCreateCallback.OnAfterCreate(stage, target)
		}
	case *RectAnchoredPath:
		if stage.OnAfterRectAnchoredPathCreateCallback != nil {
			stage.OnAfterRectAnchoredPathCreateCallback.OnAfterCreate(stage, target)
		}
	case *RectAnchoredPngImage:
		if stage.OnAfterRectAnchoredPngImageCreateCallback != nil {
			stage.OnAfterRectAnchoredPngImageCreateCallback.OnAfterCreate(stage, target)
		}
	case *RectAnchoredRect:
		if stage.OnAfterRectAnchoredRectCreateCallback != nil {
			stage.OnAfterRectAnchoredRectCreateCallback.OnAfterCreate(stage, target)
		}
	case *RectAnchoredText:
		if stage.OnAfterRectAnchoredTextCreateCallback != nil {
			stage.OnAfterRectAnchoredTextCreateCallback.OnAfterCreate(stage, target)
		}
	case *RectLinkLink:
		if stage.OnAfterRectLinkLinkCreateCallback != nil {
			stage.OnAfterRectLinkLinkCreateCallback.OnAfterCreate(stage, target)
		}
	case *SVG:
		if stage.OnAfterSVGCreateCallback != nil {
			stage.OnAfterSVGCreateCallback.OnAfterCreate(stage, target)
		}
	case *SvgText:
		if stage.OnAfterSvgTextCreateCallback != nil {
			stage.OnAfterSvgTextCreateCallback.OnAfterCreate(stage, target)
		}
	case *Text:
		if stage.OnAfterTextCreateCallback != nil {
			stage.OnAfterTextCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterCreateFromFront is a backward-compatible package-level forwarder.
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {
	stage.AfterCreateFromFront(instance)
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront[Type Gongstruct](old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Animate:
		newTarget := any(new).(*Animate)
		if stage.OnAfterAnimateUpdateCallback != nil {
			stage.OnAfterAnimateUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Circle:
		newTarget := any(new).(*Circle)
		if stage.OnAfterCircleUpdateCallback != nil {
			stage.OnAfterCircleUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Condition:
		newTarget := any(new).(*Condition)
		if stage.OnAfterConditionUpdateCallback != nil {
			stage.OnAfterConditionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ControlPoint:
		newTarget := any(new).(*ControlPoint)
		if stage.OnAfterControlPointUpdateCallback != nil {
			stage.OnAfterControlPointUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Ellipse:
		newTarget := any(new).(*Ellipse)
		if stage.OnAfterEllipseUpdateCallback != nil {
			stage.OnAfterEllipseUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FileToDownload:
		newTarget := any(new).(*FileToDownload)
		if stage.OnAfterFileToDownloadUpdateCallback != nil {
			stage.OnAfterFileToDownloadUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Layer:
		newTarget := any(new).(*Layer)
		if stage.OnAfterLayerUpdateCallback != nil {
			stage.OnAfterLayerUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Line:
		newTarget := any(new).(*Line)
		if stage.OnAfterLineUpdateCallback != nil {
			stage.OnAfterLineUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Link:
		newTarget := any(new).(*Link)
		if stage.OnAfterLinkUpdateCallback != nil {
			stage.OnAfterLinkUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *LinkAnchoredPath:
		newTarget := any(new).(*LinkAnchoredPath)
		if stage.OnAfterLinkAnchoredPathUpdateCallback != nil {
			stage.OnAfterLinkAnchoredPathUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *LinkAnchoredText:
		newTarget := any(new).(*LinkAnchoredText)
		if stage.OnAfterLinkAnchoredTextUpdateCallback != nil {
			stage.OnAfterLinkAnchoredTextUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Path:
		newTarget := any(new).(*Path)
		if stage.OnAfterPathUpdateCallback != nil {
			stage.OnAfterPathUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Point:
		newTarget := any(new).(*Point)
		if stage.OnAfterPointUpdateCallback != nil {
			stage.OnAfterPointUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Polygone:
		newTarget := any(new).(*Polygone)
		if stage.OnAfterPolygoneUpdateCallback != nil {
			stage.OnAfterPolygoneUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Polyline:
		newTarget := any(new).(*Polyline)
		if stage.OnAfterPolylineUpdateCallback != nil {
			stage.OnAfterPolylineUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Rect:
		newTarget := any(new).(*Rect)
		if stage.OnAfterRectUpdateCallback != nil {
			stage.OnAfterRectUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RectAnchoredPath:
		newTarget := any(new).(*RectAnchoredPath)
		if stage.OnAfterRectAnchoredPathUpdateCallback != nil {
			stage.OnAfterRectAnchoredPathUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RectAnchoredPngImage:
		newTarget := any(new).(*RectAnchoredPngImage)
		if stage.OnAfterRectAnchoredPngImageUpdateCallback != nil {
			stage.OnAfterRectAnchoredPngImageUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RectAnchoredRect:
		newTarget := any(new).(*RectAnchoredRect)
		if stage.OnAfterRectAnchoredRectUpdateCallback != nil {
			stage.OnAfterRectAnchoredRectUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RectAnchoredText:
		newTarget := any(new).(*RectAnchoredText)
		if stage.OnAfterRectAnchoredTextUpdateCallback != nil {
			stage.OnAfterRectAnchoredTextUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RectLinkLink:
		newTarget := any(new).(*RectLinkLink)
		if stage.OnAfterRectLinkLinkUpdateCallback != nil {
			stage.OnAfterRectLinkLinkUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SVG:
		newTarget := any(new).(*SVG)
		if stage.OnAfterSVGUpdateCallback != nil {
			stage.OnAfterSVGUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SvgText:
		newTarget := any(new).(*SvgText)
		if stage.OnAfterSvgTextUpdateCallback != nil {
			stage.OnAfterSvgTextUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Text:
		newTarget := any(new).(*Text)
		if stage.OnAfterTextUpdateCallback != nil {
			stage.OnAfterTextUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// OnAfterUpdateFromFront is a backward-compatible package-level forwarder.
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {
	stage.OnAfterUpdateFromFront(old, new)
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront[Type Gongstruct](staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Animate:
		if stage.OnAfterAnimateDeleteCallback != nil {
			staged := any(staged).(*Animate)
			stage.OnAfterAnimateDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Circle:
		if stage.OnAfterCircleDeleteCallback != nil {
			staged := any(staged).(*Circle)
			stage.OnAfterCircleDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Condition:
		if stage.OnAfterConditionDeleteCallback != nil {
			staged := any(staged).(*Condition)
			stage.OnAfterConditionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ControlPoint:
		if stage.OnAfterControlPointDeleteCallback != nil {
			staged := any(staged).(*ControlPoint)
			stage.OnAfterControlPointDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Ellipse:
		if stage.OnAfterEllipseDeleteCallback != nil {
			staged := any(staged).(*Ellipse)
			stage.OnAfterEllipseDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FileToDownload:
		if stage.OnAfterFileToDownloadDeleteCallback != nil {
			staged := any(staged).(*FileToDownload)
			stage.OnAfterFileToDownloadDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Layer:
		if stage.OnAfterLayerDeleteCallback != nil {
			staged := any(staged).(*Layer)
			stage.OnAfterLayerDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Line:
		if stage.OnAfterLineDeleteCallback != nil {
			staged := any(staged).(*Line)
			stage.OnAfterLineDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Link:
		if stage.OnAfterLinkDeleteCallback != nil {
			staged := any(staged).(*Link)
			stage.OnAfterLinkDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *LinkAnchoredPath:
		if stage.OnAfterLinkAnchoredPathDeleteCallback != nil {
			staged := any(staged).(*LinkAnchoredPath)
			stage.OnAfterLinkAnchoredPathDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *LinkAnchoredText:
		if stage.OnAfterLinkAnchoredTextDeleteCallback != nil {
			staged := any(staged).(*LinkAnchoredText)
			stage.OnAfterLinkAnchoredTextDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Path:
		if stage.OnAfterPathDeleteCallback != nil {
			staged := any(staged).(*Path)
			stage.OnAfterPathDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Point:
		if stage.OnAfterPointDeleteCallback != nil {
			staged := any(staged).(*Point)
			stage.OnAfterPointDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Polygone:
		if stage.OnAfterPolygoneDeleteCallback != nil {
			staged := any(staged).(*Polygone)
			stage.OnAfterPolygoneDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Polyline:
		if stage.OnAfterPolylineDeleteCallback != nil {
			staged := any(staged).(*Polyline)
			stage.OnAfterPolylineDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Rect:
		if stage.OnAfterRectDeleteCallback != nil {
			staged := any(staged).(*Rect)
			stage.OnAfterRectDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RectAnchoredPath:
		if stage.OnAfterRectAnchoredPathDeleteCallback != nil {
			staged := any(staged).(*RectAnchoredPath)
			stage.OnAfterRectAnchoredPathDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RectAnchoredPngImage:
		if stage.OnAfterRectAnchoredPngImageDeleteCallback != nil {
			staged := any(staged).(*RectAnchoredPngImage)
			stage.OnAfterRectAnchoredPngImageDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RectAnchoredRect:
		if stage.OnAfterRectAnchoredRectDeleteCallback != nil {
			staged := any(staged).(*RectAnchoredRect)
			stage.OnAfterRectAnchoredRectDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RectAnchoredText:
		if stage.OnAfterRectAnchoredTextDeleteCallback != nil {
			staged := any(staged).(*RectAnchoredText)
			stage.OnAfterRectAnchoredTextDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RectLinkLink:
		if stage.OnAfterRectLinkLinkDeleteCallback != nil {
			staged := any(staged).(*RectLinkLink)
			stage.OnAfterRectLinkLinkDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SVG:
		if stage.OnAfterSVGDeleteCallback != nil {
			staged := any(staged).(*SVG)
			stage.OnAfterSVGDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SvgText:
		if stage.OnAfterSvgTextDeleteCallback != nil {
			staged := any(staged).(*SvgText)
			stage.OnAfterSvgTextDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Text:
		if stage.OnAfterTextDeleteCallback != nil {
			staged := any(staged).(*Text)
			stage.OnAfterTextDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
