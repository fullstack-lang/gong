// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongsvg/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Animate:
		tmp := GetInstanceDBFromInstance[models.Animate, AnimateDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Animate":
			switch reverseField.Fieldname {
			}
		case "Circle":
			switch reverseField.Fieldname {
			case "Animations":
				if tmp.Circle_AnimationsDBID.Int64 != 0 {
					id := uint(tmp.Circle_AnimationsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoCircle.Map_CircleDBID_CirclePtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Ellipse":
			switch reverseField.Fieldname {
			case "Animates":
				if tmp.Ellipse_AnimatesDBID.Int64 != 0 {
					id := uint(tmp.Ellipse_AnimatesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoEllipse.Map_EllipseDBID_EllipsePtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Layer":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			case "Animates":
				if tmp.Line_AnimatesDBID.Int64 != 0 {
					id := uint(tmp.Line_AnimatesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoLine.Map_LineDBID_LinePtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "LinkAnchoredText":
			switch reverseField.Fieldname {
			case "Animates":
				if tmp.LinkAnchoredText_AnimatesDBID.Int64 != 0 {
					id := uint(tmp.LinkAnchoredText_AnimatesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Path":
			switch reverseField.Fieldname {
			case "Animates":
				if tmp.Path_AnimatesDBID.Int64 != 0 {
					id := uint(tmp.Path_AnimatesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoPath.Map_PathDBID_PathPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Point":
			switch reverseField.Fieldname {
			}
		case "Polygone":
			switch reverseField.Fieldname {
			case "Animates":
				if tmp.Polygone_AnimatesDBID.Int64 != 0 {
					id := uint(tmp.Polygone_AnimatesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoPolygone.Map_PolygoneDBID_PolygonePtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Polyline":
			switch reverseField.Fieldname {
			case "Animates":
				if tmp.Polyline_AnimatesDBID.Int64 != 0 {
					id := uint(tmp.Polyline_AnimatesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoPolyline.Map_PolylineDBID_PolylinePtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Rect":
			switch reverseField.Fieldname {
			case "Animations":
				if tmp.Rect_AnimationsDBID.Int64 != 0 {
					id := uint(tmp.Rect_AnimationsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoRect.Map_RectDBID_RectPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "RectAnchoredRect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredText":
			switch reverseField.Fieldname {
			case "Animates":
				if tmp.RectAnchoredText_AnimatesDBID.Int64 != 0 {
					id := uint(tmp.RectAnchoredText_AnimatesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoRectAnchoredText.Map_RectAnchoredTextDBID_RectAnchoredTextPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "RectLinkLink":
			switch reverseField.Fieldname {
			}
		case "SVG":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			case "Animates":
				if tmp.Text_AnimatesDBID.Int64 != 0 {
					id := uint(tmp.Text_AnimatesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoText.Map_TextDBID_TextPtr[id]
					res = reservePointerTarget.Name
				}
			}
		}

	case *models.Circle:
		tmp := GetInstanceDBFromInstance[models.Circle, CircleDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Animate":
			switch reverseField.Fieldname {
			}
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "Ellipse":
			switch reverseField.Fieldname {
			}
		case "Layer":
			switch reverseField.Fieldname {
			case "Circles":
				if tmp.Layer_CirclesDBID.Int64 != 0 {
					id := uint(tmp.Layer_CirclesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "LinkAnchoredText":
			switch reverseField.Fieldname {
			}
		case "Path":
			switch reverseField.Fieldname {
			}
		case "Point":
			switch reverseField.Fieldname {
			}
		case "Polygone":
			switch reverseField.Fieldname {
			}
		case "Polyline":
			switch reverseField.Fieldname {
			}
		case "Rect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredRect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredText":
			switch reverseField.Fieldname {
			}
		case "RectLinkLink":
			switch reverseField.Fieldname {
			}
		case "SVG":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.Ellipse:
		tmp := GetInstanceDBFromInstance[models.Ellipse, EllipseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Animate":
			switch reverseField.Fieldname {
			}
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "Ellipse":
			switch reverseField.Fieldname {
			}
		case "Layer":
			switch reverseField.Fieldname {
			case "Ellipses":
				if tmp.Layer_EllipsesDBID.Int64 != 0 {
					id := uint(tmp.Layer_EllipsesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "LinkAnchoredText":
			switch reverseField.Fieldname {
			}
		case "Path":
			switch reverseField.Fieldname {
			}
		case "Point":
			switch reverseField.Fieldname {
			}
		case "Polygone":
			switch reverseField.Fieldname {
			}
		case "Polyline":
			switch reverseField.Fieldname {
			}
		case "Rect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredRect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredText":
			switch reverseField.Fieldname {
			}
		case "RectLinkLink":
			switch reverseField.Fieldname {
			}
		case "SVG":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.Layer:
		tmp := GetInstanceDBFromInstance[models.Layer, LayerDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Animate":
			switch reverseField.Fieldname {
			}
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "Ellipse":
			switch reverseField.Fieldname {
			}
		case "Layer":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "LinkAnchoredText":
			switch reverseField.Fieldname {
			}
		case "Path":
			switch reverseField.Fieldname {
			}
		case "Point":
			switch reverseField.Fieldname {
			}
		case "Polygone":
			switch reverseField.Fieldname {
			}
		case "Polyline":
			switch reverseField.Fieldname {
			}
		case "Rect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredRect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredText":
			switch reverseField.Fieldname {
			}
		case "RectLinkLink":
			switch reverseField.Fieldname {
			}
		case "SVG":
			switch reverseField.Fieldname {
			case "Layers":
				if tmp.SVG_LayersDBID.Int64 != 0 {
					id := uint(tmp.SVG_LayersDBID.Int64)
					reservePointerTarget := backRepo.BackRepoSVG.Map_SVGDBID_SVGPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.Line:
		tmp := GetInstanceDBFromInstance[models.Line, LineDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Animate":
			switch reverseField.Fieldname {
			}
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "Ellipse":
			switch reverseField.Fieldname {
			}
		case "Layer":
			switch reverseField.Fieldname {
			case "Lines":
				if tmp.Layer_LinesDBID.Int64 != 0 {
					id := uint(tmp.Layer_LinesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "LinkAnchoredText":
			switch reverseField.Fieldname {
			}
		case "Path":
			switch reverseField.Fieldname {
			}
		case "Point":
			switch reverseField.Fieldname {
			}
		case "Polygone":
			switch reverseField.Fieldname {
			}
		case "Polyline":
			switch reverseField.Fieldname {
			}
		case "Rect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredRect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredText":
			switch reverseField.Fieldname {
			}
		case "RectLinkLink":
			switch reverseField.Fieldname {
			}
		case "SVG":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.Link:
		tmp := GetInstanceDBFromInstance[models.Link, LinkDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Animate":
			switch reverseField.Fieldname {
			}
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "Ellipse":
			switch reverseField.Fieldname {
			}
		case "Layer":
			switch reverseField.Fieldname {
			case "Links":
				if tmp.Layer_LinksDBID.Int64 != 0 {
					id := uint(tmp.Layer_LinksDBID.Int64)
					reservePointerTarget := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "LinkAnchoredText":
			switch reverseField.Fieldname {
			}
		case "Path":
			switch reverseField.Fieldname {
			}
		case "Point":
			switch reverseField.Fieldname {
			}
		case "Polygone":
			switch reverseField.Fieldname {
			}
		case "Polyline":
			switch reverseField.Fieldname {
			}
		case "Rect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredRect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredText":
			switch reverseField.Fieldname {
			}
		case "RectLinkLink":
			switch reverseField.Fieldname {
			}
		case "SVG":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.LinkAnchoredText:
		tmp := GetInstanceDBFromInstance[models.LinkAnchoredText, LinkAnchoredTextDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Animate":
			switch reverseField.Fieldname {
			}
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "Ellipse":
			switch reverseField.Fieldname {
			}
		case "Layer":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			case "TextAtArrowEnd":
				if tmp.Link_TextAtArrowEndDBID.Int64 != 0 {
					id := uint(tmp.Link_TextAtArrowEndDBID.Int64)
					reservePointerTarget := backRepo.BackRepoLink.Map_LinkDBID_LinkPtr[id]
					res = reservePointerTarget.Name
				}
			case "TextAtArrowStart":
				if tmp.Link_TextAtArrowStartDBID.Int64 != 0 {
					id := uint(tmp.Link_TextAtArrowStartDBID.Int64)
					reservePointerTarget := backRepo.BackRepoLink.Map_LinkDBID_LinkPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "LinkAnchoredText":
			switch reverseField.Fieldname {
			}
		case "Path":
			switch reverseField.Fieldname {
			}
		case "Point":
			switch reverseField.Fieldname {
			}
		case "Polygone":
			switch reverseField.Fieldname {
			}
		case "Polyline":
			switch reverseField.Fieldname {
			}
		case "Rect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredRect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredText":
			switch reverseField.Fieldname {
			}
		case "RectLinkLink":
			switch reverseField.Fieldname {
			}
		case "SVG":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.Path:
		tmp := GetInstanceDBFromInstance[models.Path, PathDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Animate":
			switch reverseField.Fieldname {
			}
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "Ellipse":
			switch reverseField.Fieldname {
			}
		case "Layer":
			switch reverseField.Fieldname {
			case "Paths":
				if tmp.Layer_PathsDBID.Int64 != 0 {
					id := uint(tmp.Layer_PathsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "LinkAnchoredText":
			switch reverseField.Fieldname {
			}
		case "Path":
			switch reverseField.Fieldname {
			}
		case "Point":
			switch reverseField.Fieldname {
			}
		case "Polygone":
			switch reverseField.Fieldname {
			}
		case "Polyline":
			switch reverseField.Fieldname {
			}
		case "Rect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredRect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredText":
			switch reverseField.Fieldname {
			}
		case "RectLinkLink":
			switch reverseField.Fieldname {
			}
		case "SVG":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.Point:
		tmp := GetInstanceDBFromInstance[models.Point, PointDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Animate":
			switch reverseField.Fieldname {
			}
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "Ellipse":
			switch reverseField.Fieldname {
			}
		case "Layer":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			case "ControlPoints":
				if tmp.Link_ControlPointsDBID.Int64 != 0 {
					id := uint(tmp.Link_ControlPointsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoLink.Map_LinkDBID_LinkPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "LinkAnchoredText":
			switch reverseField.Fieldname {
			}
		case "Path":
			switch reverseField.Fieldname {
			}
		case "Point":
			switch reverseField.Fieldname {
			}
		case "Polygone":
			switch reverseField.Fieldname {
			}
		case "Polyline":
			switch reverseField.Fieldname {
			}
		case "Rect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredRect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredText":
			switch reverseField.Fieldname {
			}
		case "RectLinkLink":
			switch reverseField.Fieldname {
			}
		case "SVG":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.Polygone:
		tmp := GetInstanceDBFromInstance[models.Polygone, PolygoneDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Animate":
			switch reverseField.Fieldname {
			}
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "Ellipse":
			switch reverseField.Fieldname {
			}
		case "Layer":
			switch reverseField.Fieldname {
			case "Polygones":
				if tmp.Layer_PolygonesDBID.Int64 != 0 {
					id := uint(tmp.Layer_PolygonesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "LinkAnchoredText":
			switch reverseField.Fieldname {
			}
		case "Path":
			switch reverseField.Fieldname {
			}
		case "Point":
			switch reverseField.Fieldname {
			}
		case "Polygone":
			switch reverseField.Fieldname {
			}
		case "Polyline":
			switch reverseField.Fieldname {
			}
		case "Rect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredRect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredText":
			switch reverseField.Fieldname {
			}
		case "RectLinkLink":
			switch reverseField.Fieldname {
			}
		case "SVG":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.Polyline:
		tmp := GetInstanceDBFromInstance[models.Polyline, PolylineDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Animate":
			switch reverseField.Fieldname {
			}
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "Ellipse":
			switch reverseField.Fieldname {
			}
		case "Layer":
			switch reverseField.Fieldname {
			case "Polylines":
				if tmp.Layer_PolylinesDBID.Int64 != 0 {
					id := uint(tmp.Layer_PolylinesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "LinkAnchoredText":
			switch reverseField.Fieldname {
			}
		case "Path":
			switch reverseField.Fieldname {
			}
		case "Point":
			switch reverseField.Fieldname {
			}
		case "Polygone":
			switch reverseField.Fieldname {
			}
		case "Polyline":
			switch reverseField.Fieldname {
			}
		case "Rect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredRect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredText":
			switch reverseField.Fieldname {
			}
		case "RectLinkLink":
			switch reverseField.Fieldname {
			}
		case "SVG":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.Rect:
		tmp := GetInstanceDBFromInstance[models.Rect, RectDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Animate":
			switch reverseField.Fieldname {
			}
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "Ellipse":
			switch reverseField.Fieldname {
			}
		case "Layer":
			switch reverseField.Fieldname {
			case "Rects":
				if tmp.Layer_RectsDBID.Int64 != 0 {
					id := uint(tmp.Layer_RectsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "LinkAnchoredText":
			switch reverseField.Fieldname {
			}
		case "Path":
			switch reverseField.Fieldname {
			}
		case "Point":
			switch reverseField.Fieldname {
			}
		case "Polygone":
			switch reverseField.Fieldname {
			}
		case "Polyline":
			switch reverseField.Fieldname {
			}
		case "Rect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredRect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredText":
			switch reverseField.Fieldname {
			}
		case "RectLinkLink":
			switch reverseField.Fieldname {
			}
		case "SVG":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.RectAnchoredRect:
		tmp := GetInstanceDBFromInstance[models.RectAnchoredRect, RectAnchoredRectDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Animate":
			switch reverseField.Fieldname {
			}
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "Ellipse":
			switch reverseField.Fieldname {
			}
		case "Layer":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "LinkAnchoredText":
			switch reverseField.Fieldname {
			}
		case "Path":
			switch reverseField.Fieldname {
			}
		case "Point":
			switch reverseField.Fieldname {
			}
		case "Polygone":
			switch reverseField.Fieldname {
			}
		case "Polyline":
			switch reverseField.Fieldname {
			}
		case "Rect":
			switch reverseField.Fieldname {
			case "RectAnchoredRects":
				if tmp.Rect_RectAnchoredRectsDBID.Int64 != 0 {
					id := uint(tmp.Rect_RectAnchoredRectsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoRect.Map_RectDBID_RectPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "RectAnchoredRect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredText":
			switch reverseField.Fieldname {
			}
		case "RectLinkLink":
			switch reverseField.Fieldname {
			}
		case "SVG":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.RectAnchoredText:
		tmp := GetInstanceDBFromInstance[models.RectAnchoredText, RectAnchoredTextDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Animate":
			switch reverseField.Fieldname {
			}
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "Ellipse":
			switch reverseField.Fieldname {
			}
		case "Layer":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "LinkAnchoredText":
			switch reverseField.Fieldname {
			}
		case "Path":
			switch reverseField.Fieldname {
			}
		case "Point":
			switch reverseField.Fieldname {
			}
		case "Polygone":
			switch reverseField.Fieldname {
			}
		case "Polyline":
			switch reverseField.Fieldname {
			}
		case "Rect":
			switch reverseField.Fieldname {
			case "RectAnchoredTexts":
				if tmp.Rect_RectAnchoredTextsDBID.Int64 != 0 {
					id := uint(tmp.Rect_RectAnchoredTextsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoRect.Map_RectDBID_RectPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "RectAnchoredRect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredText":
			switch reverseField.Fieldname {
			}
		case "RectLinkLink":
			switch reverseField.Fieldname {
			}
		case "SVG":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.RectLinkLink:
		tmp := GetInstanceDBFromInstance[models.RectLinkLink, RectLinkLinkDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Animate":
			switch reverseField.Fieldname {
			}
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "Ellipse":
			switch reverseField.Fieldname {
			}
		case "Layer":
			switch reverseField.Fieldname {
			case "RectLinkLinks":
				if tmp.Layer_RectLinkLinksDBID.Int64 != 0 {
					id := uint(tmp.Layer_RectLinkLinksDBID.Int64)
					reservePointerTarget := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "LinkAnchoredText":
			switch reverseField.Fieldname {
			}
		case "Path":
			switch reverseField.Fieldname {
			}
		case "Point":
			switch reverseField.Fieldname {
			}
		case "Polygone":
			switch reverseField.Fieldname {
			}
		case "Polyline":
			switch reverseField.Fieldname {
			}
		case "Rect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredRect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredText":
			switch reverseField.Fieldname {
			}
		case "RectLinkLink":
			switch reverseField.Fieldname {
			}
		case "SVG":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.SVG:
		tmp := GetInstanceDBFromInstance[models.SVG, SVGDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Animate":
			switch reverseField.Fieldname {
			}
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "Ellipse":
			switch reverseField.Fieldname {
			}
		case "Layer":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "LinkAnchoredText":
			switch reverseField.Fieldname {
			}
		case "Path":
			switch reverseField.Fieldname {
			}
		case "Point":
			switch reverseField.Fieldname {
			}
		case "Polygone":
			switch reverseField.Fieldname {
			}
		case "Polyline":
			switch reverseField.Fieldname {
			}
		case "Rect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredRect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredText":
			switch reverseField.Fieldname {
			}
		case "RectLinkLink":
			switch reverseField.Fieldname {
			}
		case "SVG":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	case *models.Text:
		tmp := GetInstanceDBFromInstance[models.Text, TextDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Animate":
			switch reverseField.Fieldname {
			}
		case "Circle":
			switch reverseField.Fieldname {
			}
		case "Ellipse":
			switch reverseField.Fieldname {
			}
		case "Layer":
			switch reverseField.Fieldname {
			case "Texts":
				if tmp.Layer_TextsDBID.Int64 != 0 {
					id := uint(tmp.Layer_TextsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Line":
			switch reverseField.Fieldname {
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "LinkAnchoredText":
			switch reverseField.Fieldname {
			}
		case "Path":
			switch reverseField.Fieldname {
			}
		case "Point":
			switch reverseField.Fieldname {
			}
		case "Polygone":
			switch reverseField.Fieldname {
			}
		case "Polyline":
			switch reverseField.Fieldname {
			}
		case "Rect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredRect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredText":
			switch reverseField.Fieldname {
			}
		case "RectLinkLink":
			switch reverseField.Fieldname {
			}
		case "SVG":
			switch reverseField.Fieldname {
			}
		case "Text":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return
}
