// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongsvg/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Animate:
		animateInstance := any(concreteInstance).(*models.Animate)
		ret2 := backRepo.BackRepoAnimate.GetAnimateDBFromAnimatePtr(animateInstance)
		ret = any(ret2).(*T2)
	case *models.Circle:
		circleInstance := any(concreteInstance).(*models.Circle)
		ret2 := backRepo.BackRepoCircle.GetCircleDBFromCirclePtr(circleInstance)
		ret = any(ret2).(*T2)
	case *models.Ellipse:
		ellipseInstance := any(concreteInstance).(*models.Ellipse)
		ret2 := backRepo.BackRepoEllipse.GetEllipseDBFromEllipsePtr(ellipseInstance)
		ret = any(ret2).(*T2)
	case *models.Layer:
		layerInstance := any(concreteInstance).(*models.Layer)
		ret2 := backRepo.BackRepoLayer.GetLayerDBFromLayerPtr(layerInstance)
		ret = any(ret2).(*T2)
	case *models.Line:
		lineInstance := any(concreteInstance).(*models.Line)
		ret2 := backRepo.BackRepoLine.GetLineDBFromLinePtr(lineInstance)
		ret = any(ret2).(*T2)
	case *models.Link:
		linkInstance := any(concreteInstance).(*models.Link)
		ret2 := backRepo.BackRepoLink.GetLinkDBFromLinkPtr(linkInstance)
		ret = any(ret2).(*T2)
	case *models.LinkAnchoredText:
		linkanchoredtextInstance := any(concreteInstance).(*models.LinkAnchoredText)
		ret2 := backRepo.BackRepoLinkAnchoredText.GetLinkAnchoredTextDBFromLinkAnchoredTextPtr(linkanchoredtextInstance)
		ret = any(ret2).(*T2)
	case *models.Path:
		pathInstance := any(concreteInstance).(*models.Path)
		ret2 := backRepo.BackRepoPath.GetPathDBFromPathPtr(pathInstance)
		ret = any(ret2).(*T2)
	case *models.Point:
		pointInstance := any(concreteInstance).(*models.Point)
		ret2 := backRepo.BackRepoPoint.GetPointDBFromPointPtr(pointInstance)
		ret = any(ret2).(*T2)
	case *models.Polygone:
		polygoneInstance := any(concreteInstance).(*models.Polygone)
		ret2 := backRepo.BackRepoPolygone.GetPolygoneDBFromPolygonePtr(polygoneInstance)
		ret = any(ret2).(*T2)
	case *models.Polyline:
		polylineInstance := any(concreteInstance).(*models.Polyline)
		ret2 := backRepo.BackRepoPolyline.GetPolylineDBFromPolylinePtr(polylineInstance)
		ret = any(ret2).(*T2)
	case *models.Rect:
		rectInstance := any(concreteInstance).(*models.Rect)
		ret2 := backRepo.BackRepoRect.GetRectDBFromRectPtr(rectInstance)
		ret = any(ret2).(*T2)
	case *models.RectAnchoredPath:
		rectanchoredpathInstance := any(concreteInstance).(*models.RectAnchoredPath)
		ret2 := backRepo.BackRepoRectAnchoredPath.GetRectAnchoredPathDBFromRectAnchoredPathPtr(rectanchoredpathInstance)
		ret = any(ret2).(*T2)
	case *models.RectAnchoredRect:
		rectanchoredrectInstance := any(concreteInstance).(*models.RectAnchoredRect)
		ret2 := backRepo.BackRepoRectAnchoredRect.GetRectAnchoredRectDBFromRectAnchoredRectPtr(rectanchoredrectInstance)
		ret = any(ret2).(*T2)
	case *models.RectAnchoredText:
		rectanchoredtextInstance := any(concreteInstance).(*models.RectAnchoredText)
		ret2 := backRepo.BackRepoRectAnchoredText.GetRectAnchoredTextDBFromRectAnchoredTextPtr(rectanchoredtextInstance)
		ret = any(ret2).(*T2)
	case *models.RectLinkLink:
		rectlinklinkInstance := any(concreteInstance).(*models.RectLinkLink)
		ret2 := backRepo.BackRepoRectLinkLink.GetRectLinkLinkDBFromRectLinkLinkPtr(rectlinklinkInstance)
		ret = any(ret2).(*T2)
	case *models.SVG:
		svgInstance := any(concreteInstance).(*models.SVG)
		ret2 := backRepo.BackRepoSVG.GetSVGDBFromSVGPtr(svgInstance)
		ret = any(ret2).(*T2)
	case *models.Text:
		textInstance := any(concreteInstance).(*models.Text)
		ret2 := backRepo.BackRepoText.GetTextDBFromTextPtr(textInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Animate:
		tmp := GetInstanceDBFromInstance[models.Animate, AnimateDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Circle:
		tmp := GetInstanceDBFromInstance[models.Circle, CircleDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Ellipse:
		tmp := GetInstanceDBFromInstance[models.Ellipse, EllipseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Layer:
		tmp := GetInstanceDBFromInstance[models.Layer, LayerDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Line:
		tmp := GetInstanceDBFromInstance[models.Line, LineDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Link:
		tmp := GetInstanceDBFromInstance[models.Link, LinkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.LinkAnchoredText:
		tmp := GetInstanceDBFromInstance[models.LinkAnchoredText, LinkAnchoredTextDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Path:
		tmp := GetInstanceDBFromInstance[models.Path, PathDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Point:
		tmp := GetInstanceDBFromInstance[models.Point, PointDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Polygone:
		tmp := GetInstanceDBFromInstance[models.Polygone, PolygoneDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Polyline:
		tmp := GetInstanceDBFromInstance[models.Polyline, PolylineDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Rect:
		tmp := GetInstanceDBFromInstance[models.Rect, RectDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RectAnchoredPath:
		tmp := GetInstanceDBFromInstance[models.RectAnchoredPath, RectAnchoredPathDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RectAnchoredRect:
		tmp := GetInstanceDBFromInstance[models.RectAnchoredRect, RectAnchoredRectDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RectAnchoredText:
		tmp := GetInstanceDBFromInstance[models.RectAnchoredText, RectAnchoredTextDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RectLinkLink:
		tmp := GetInstanceDBFromInstance[models.RectLinkLink, RectLinkLinkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SVG:
		tmp := GetInstanceDBFromInstance[models.SVG, SVGDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Text:
		tmp := GetInstanceDBFromInstance[models.Text, TextDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Animate:
		tmp := GetInstanceDBFromInstance[models.Animate, AnimateDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Circle:
		tmp := GetInstanceDBFromInstance[models.Circle, CircleDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Ellipse:
		tmp := GetInstanceDBFromInstance[models.Ellipse, EllipseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Layer:
		tmp := GetInstanceDBFromInstance[models.Layer, LayerDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Line:
		tmp := GetInstanceDBFromInstance[models.Line, LineDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Link:
		tmp := GetInstanceDBFromInstance[models.Link, LinkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.LinkAnchoredText:
		tmp := GetInstanceDBFromInstance[models.LinkAnchoredText, LinkAnchoredTextDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Path:
		tmp := GetInstanceDBFromInstance[models.Path, PathDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Point:
		tmp := GetInstanceDBFromInstance[models.Point, PointDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Polygone:
		tmp := GetInstanceDBFromInstance[models.Polygone, PolygoneDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Polyline:
		tmp := GetInstanceDBFromInstance[models.Polyline, PolylineDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Rect:
		tmp := GetInstanceDBFromInstance[models.Rect, RectDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RectAnchoredPath:
		tmp := GetInstanceDBFromInstance[models.RectAnchoredPath, RectAnchoredPathDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RectAnchoredRect:
		tmp := GetInstanceDBFromInstance[models.RectAnchoredRect, RectAnchoredRectDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RectAnchoredText:
		tmp := GetInstanceDBFromInstance[models.RectAnchoredText, RectAnchoredTextDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RectLinkLink:
		tmp := GetInstanceDBFromInstance[models.RectLinkLink, RectLinkLinkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SVG:
		tmp := GetInstanceDBFromInstance[models.SVG, SVGDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Text:
		tmp := GetInstanceDBFromInstance[models.Text, TextDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
