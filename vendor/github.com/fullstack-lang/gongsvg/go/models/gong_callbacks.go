package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

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
	case *Ellipse:
		if stage.OnAfterEllipseCreateCallback != nil {
			stage.OnAfterEllipseCreateCallback.OnAfterCreate(stage, target)
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
	case *Text:
		if stage.OnAfterTextCreateCallback != nil {
			stage.OnAfterTextCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

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
	case *Ellipse:
		newTarget := any(new).(*Ellipse)
		if stage.OnAfterEllipseUpdateCallback != nil {
			stage.OnAfterEllipseUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *Text:
		newTarget := any(new).(*Text)
		if stage.OnAfterTextUpdateCallback != nil {
			stage.OnAfterTextUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

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
	case *Ellipse:
		if stage.OnAfterEllipseDeleteCallback != nil {
			staged := any(staged).(*Ellipse)
			stage.OnAfterEllipseDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *Text:
		if stage.OnAfterTextDeleteCallback != nil {
			staged := any(staged).(*Text)
			stage.OnAfterTextDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Animate:
		if stage.OnAfterAnimateReadCallback != nil {
			stage.OnAfterAnimateReadCallback.OnAfterRead(stage, target)
		}
	case *Circle:
		if stage.OnAfterCircleReadCallback != nil {
			stage.OnAfterCircleReadCallback.OnAfterRead(stage, target)
		}
	case *Ellipse:
		if stage.OnAfterEllipseReadCallback != nil {
			stage.OnAfterEllipseReadCallback.OnAfterRead(stage, target)
		}
	case *Layer:
		if stage.OnAfterLayerReadCallback != nil {
			stage.OnAfterLayerReadCallback.OnAfterRead(stage, target)
		}
	case *Line:
		if stage.OnAfterLineReadCallback != nil {
			stage.OnAfterLineReadCallback.OnAfterRead(stage, target)
		}
	case *Link:
		if stage.OnAfterLinkReadCallback != nil {
			stage.OnAfterLinkReadCallback.OnAfterRead(stage, target)
		}
	case *LinkAnchoredText:
		if stage.OnAfterLinkAnchoredTextReadCallback != nil {
			stage.OnAfterLinkAnchoredTextReadCallback.OnAfterRead(stage, target)
		}
	case *Path:
		if stage.OnAfterPathReadCallback != nil {
			stage.OnAfterPathReadCallback.OnAfterRead(stage, target)
		}
	case *Point:
		if stage.OnAfterPointReadCallback != nil {
			stage.OnAfterPointReadCallback.OnAfterRead(stage, target)
		}
	case *Polygone:
		if stage.OnAfterPolygoneReadCallback != nil {
			stage.OnAfterPolygoneReadCallback.OnAfterRead(stage, target)
		}
	case *Polyline:
		if stage.OnAfterPolylineReadCallback != nil {
			stage.OnAfterPolylineReadCallback.OnAfterRead(stage, target)
		}
	case *Rect:
		if stage.OnAfterRectReadCallback != nil {
			stage.OnAfterRectReadCallback.OnAfterRead(stage, target)
		}
	case *RectAnchoredRect:
		if stage.OnAfterRectAnchoredRectReadCallback != nil {
			stage.OnAfterRectAnchoredRectReadCallback.OnAfterRead(stage, target)
		}
	case *RectAnchoredText:
		if stage.OnAfterRectAnchoredTextReadCallback != nil {
			stage.OnAfterRectAnchoredTextReadCallback.OnAfterRead(stage, target)
		}
	case *RectLinkLink:
		if stage.OnAfterRectLinkLinkReadCallback != nil {
			stage.OnAfterRectLinkLinkReadCallback.OnAfterRead(stage, target)
		}
	case *SVG:
		if stage.OnAfterSVGReadCallback != nil {
			stage.OnAfterSVGReadCallback.OnAfterRead(stage, target)
		}
	case *Text:
		if stage.OnAfterTextReadCallback != nil {
			stage.OnAfterTextReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Animate:
		stage.OnAfterAnimateUpdateCallback = any(callback).(OnAfterUpdateInterface[Animate])
	
	case *Circle:
		stage.OnAfterCircleUpdateCallback = any(callback).(OnAfterUpdateInterface[Circle])
	
	case *Ellipse:
		stage.OnAfterEllipseUpdateCallback = any(callback).(OnAfterUpdateInterface[Ellipse])
	
	case *Layer:
		stage.OnAfterLayerUpdateCallback = any(callback).(OnAfterUpdateInterface[Layer])
	
	case *Line:
		stage.OnAfterLineUpdateCallback = any(callback).(OnAfterUpdateInterface[Line])
	
	case *Link:
		stage.OnAfterLinkUpdateCallback = any(callback).(OnAfterUpdateInterface[Link])
	
	case *LinkAnchoredText:
		stage.OnAfterLinkAnchoredTextUpdateCallback = any(callback).(OnAfterUpdateInterface[LinkAnchoredText])
	
	case *Path:
		stage.OnAfterPathUpdateCallback = any(callback).(OnAfterUpdateInterface[Path])
	
	case *Point:
		stage.OnAfterPointUpdateCallback = any(callback).(OnAfterUpdateInterface[Point])
	
	case *Polygone:
		stage.OnAfterPolygoneUpdateCallback = any(callback).(OnAfterUpdateInterface[Polygone])
	
	case *Polyline:
		stage.OnAfterPolylineUpdateCallback = any(callback).(OnAfterUpdateInterface[Polyline])
	
	case *Rect:
		stage.OnAfterRectUpdateCallback = any(callback).(OnAfterUpdateInterface[Rect])
	
	case *RectAnchoredRect:
		stage.OnAfterRectAnchoredRectUpdateCallback = any(callback).(OnAfterUpdateInterface[RectAnchoredRect])
	
	case *RectAnchoredText:
		stage.OnAfterRectAnchoredTextUpdateCallback = any(callback).(OnAfterUpdateInterface[RectAnchoredText])
	
	case *RectLinkLink:
		stage.OnAfterRectLinkLinkUpdateCallback = any(callback).(OnAfterUpdateInterface[RectLinkLink])
	
	case *SVG:
		stage.OnAfterSVGUpdateCallback = any(callback).(OnAfterUpdateInterface[SVG])
	
	case *Text:
		stage.OnAfterTextUpdateCallback = any(callback).(OnAfterUpdateInterface[Text])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Animate:
		stage.OnAfterAnimateCreateCallback = any(callback).(OnAfterCreateInterface[Animate])
	
	case *Circle:
		stage.OnAfterCircleCreateCallback = any(callback).(OnAfterCreateInterface[Circle])
	
	case *Ellipse:
		stage.OnAfterEllipseCreateCallback = any(callback).(OnAfterCreateInterface[Ellipse])
	
	case *Layer:
		stage.OnAfterLayerCreateCallback = any(callback).(OnAfterCreateInterface[Layer])
	
	case *Line:
		stage.OnAfterLineCreateCallback = any(callback).(OnAfterCreateInterface[Line])
	
	case *Link:
		stage.OnAfterLinkCreateCallback = any(callback).(OnAfterCreateInterface[Link])
	
	case *LinkAnchoredText:
		stage.OnAfterLinkAnchoredTextCreateCallback = any(callback).(OnAfterCreateInterface[LinkAnchoredText])
	
	case *Path:
		stage.OnAfterPathCreateCallback = any(callback).(OnAfterCreateInterface[Path])
	
	case *Point:
		stage.OnAfterPointCreateCallback = any(callback).(OnAfterCreateInterface[Point])
	
	case *Polygone:
		stage.OnAfterPolygoneCreateCallback = any(callback).(OnAfterCreateInterface[Polygone])
	
	case *Polyline:
		stage.OnAfterPolylineCreateCallback = any(callback).(OnAfterCreateInterface[Polyline])
	
	case *Rect:
		stage.OnAfterRectCreateCallback = any(callback).(OnAfterCreateInterface[Rect])
	
	case *RectAnchoredRect:
		stage.OnAfterRectAnchoredRectCreateCallback = any(callback).(OnAfterCreateInterface[RectAnchoredRect])
	
	case *RectAnchoredText:
		stage.OnAfterRectAnchoredTextCreateCallback = any(callback).(OnAfterCreateInterface[RectAnchoredText])
	
	case *RectLinkLink:
		stage.OnAfterRectLinkLinkCreateCallback = any(callback).(OnAfterCreateInterface[RectLinkLink])
	
	case *SVG:
		stage.OnAfterSVGCreateCallback = any(callback).(OnAfterCreateInterface[SVG])
	
	case *Text:
		stage.OnAfterTextCreateCallback = any(callback).(OnAfterCreateInterface[Text])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Animate:
		stage.OnAfterAnimateDeleteCallback = any(callback).(OnAfterDeleteInterface[Animate])
	
	case *Circle:
		stage.OnAfterCircleDeleteCallback = any(callback).(OnAfterDeleteInterface[Circle])
	
	case *Ellipse:
		stage.OnAfterEllipseDeleteCallback = any(callback).(OnAfterDeleteInterface[Ellipse])
	
	case *Layer:
		stage.OnAfterLayerDeleteCallback = any(callback).(OnAfterDeleteInterface[Layer])
	
	case *Line:
		stage.OnAfterLineDeleteCallback = any(callback).(OnAfterDeleteInterface[Line])
	
	case *Link:
		stage.OnAfterLinkDeleteCallback = any(callback).(OnAfterDeleteInterface[Link])
	
	case *LinkAnchoredText:
		stage.OnAfterLinkAnchoredTextDeleteCallback = any(callback).(OnAfterDeleteInterface[LinkAnchoredText])
	
	case *Path:
		stage.OnAfterPathDeleteCallback = any(callback).(OnAfterDeleteInterface[Path])
	
	case *Point:
		stage.OnAfterPointDeleteCallback = any(callback).(OnAfterDeleteInterface[Point])
	
	case *Polygone:
		stage.OnAfterPolygoneDeleteCallback = any(callback).(OnAfterDeleteInterface[Polygone])
	
	case *Polyline:
		stage.OnAfterPolylineDeleteCallback = any(callback).(OnAfterDeleteInterface[Polyline])
	
	case *Rect:
		stage.OnAfterRectDeleteCallback = any(callback).(OnAfterDeleteInterface[Rect])
	
	case *RectAnchoredRect:
		stage.OnAfterRectAnchoredRectDeleteCallback = any(callback).(OnAfterDeleteInterface[RectAnchoredRect])
	
	case *RectAnchoredText:
		stage.OnAfterRectAnchoredTextDeleteCallback = any(callback).(OnAfterDeleteInterface[RectAnchoredText])
	
	case *RectLinkLink:
		stage.OnAfterRectLinkLinkDeleteCallback = any(callback).(OnAfterDeleteInterface[RectLinkLink])
	
	case *SVG:
		stage.OnAfterSVGDeleteCallback = any(callback).(OnAfterDeleteInterface[SVG])
	
	case *Text:
		stage.OnAfterTextDeleteCallback = any(callback).(OnAfterDeleteInterface[Text])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Animate:
		stage.OnAfterAnimateReadCallback = any(callback).(OnAfterReadInterface[Animate])
	
	case *Circle:
		stage.OnAfterCircleReadCallback = any(callback).(OnAfterReadInterface[Circle])
	
	case *Ellipse:
		stage.OnAfterEllipseReadCallback = any(callback).(OnAfterReadInterface[Ellipse])
	
	case *Layer:
		stage.OnAfterLayerReadCallback = any(callback).(OnAfterReadInterface[Layer])
	
	case *Line:
		stage.OnAfterLineReadCallback = any(callback).(OnAfterReadInterface[Line])
	
	case *Link:
		stage.OnAfterLinkReadCallback = any(callback).(OnAfterReadInterface[Link])
	
	case *LinkAnchoredText:
		stage.OnAfterLinkAnchoredTextReadCallback = any(callback).(OnAfterReadInterface[LinkAnchoredText])
	
	case *Path:
		stage.OnAfterPathReadCallback = any(callback).(OnAfterReadInterface[Path])
	
	case *Point:
		stage.OnAfterPointReadCallback = any(callback).(OnAfterReadInterface[Point])
	
	case *Polygone:
		stage.OnAfterPolygoneReadCallback = any(callback).(OnAfterReadInterface[Polygone])
	
	case *Polyline:
		stage.OnAfterPolylineReadCallback = any(callback).(OnAfterReadInterface[Polyline])
	
	case *Rect:
		stage.OnAfterRectReadCallback = any(callback).(OnAfterReadInterface[Rect])
	
	case *RectAnchoredRect:
		stage.OnAfterRectAnchoredRectReadCallback = any(callback).(OnAfterReadInterface[RectAnchoredRect])
	
	case *RectAnchoredText:
		stage.OnAfterRectAnchoredTextReadCallback = any(callback).(OnAfterReadInterface[RectAnchoredText])
	
	case *RectLinkLink:
		stage.OnAfterRectLinkLinkReadCallback = any(callback).(OnAfterReadInterface[RectLinkLink])
	
	case *SVG:
		stage.OnAfterSVGReadCallback = any(callback).(OnAfterReadInterface[SVG])
	
	case *Text:
		stage.OnAfterTextReadCallback = any(callback).(OnAfterReadInterface[Text])
	
	}
}
