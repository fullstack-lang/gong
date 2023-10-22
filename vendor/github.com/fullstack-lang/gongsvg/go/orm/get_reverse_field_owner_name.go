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
				if _circle, ok := stage.Circle_Animations_reverseMap[inst]; ok {
					res = _circle.Name
				}
			}
		case "Ellipse":
			switch reverseField.Fieldname {
			case "Animates":
				if _ellipse, ok := stage.Ellipse_Animates_reverseMap[inst]; ok {
					res = _ellipse.Name
				}
			}
		case "Layer":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			case "Animates":
				if _line, ok := stage.Line_Animates_reverseMap[inst]; ok {
					res = _line.Name
				}
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "LinkAnchoredText":
			switch reverseField.Fieldname {
			case "Animates":
				if _linkanchoredtext, ok := stage.LinkAnchoredText_Animates_reverseMap[inst]; ok {
					res = _linkanchoredtext.Name
				}
			}
		case "Path":
			switch reverseField.Fieldname {
			case "Animates":
				if _path, ok := stage.Path_Animates_reverseMap[inst]; ok {
					res = _path.Name
				}
			}
		case "Point":
			switch reverseField.Fieldname {
			}
		case "Polygone":
			switch reverseField.Fieldname {
			case "Animates":
				if _polygone, ok := stage.Polygone_Animates_reverseMap[inst]; ok {
					res = _polygone.Name
				}
			}
		case "Polyline":
			switch reverseField.Fieldname {
			case "Animates":
				if _polyline, ok := stage.Polyline_Animates_reverseMap[inst]; ok {
					res = _polyline.Name
				}
			}
		case "Rect":
			switch reverseField.Fieldname {
			case "Animations":
				if _rect, ok := stage.Rect_Animations_reverseMap[inst]; ok {
					res = _rect.Name
				}
			}
		case "RectAnchoredRect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredText":
			switch reverseField.Fieldname {
			case "Animates":
				if _rectanchoredtext, ok := stage.RectAnchoredText_Animates_reverseMap[inst]; ok {
					res = _rectanchoredtext.Name
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
				if _text, ok := stage.Text_Animates_reverseMap[inst]; ok {
					res = _text.Name
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
				if _layer, ok := stage.Layer_Circles_reverseMap[inst]; ok {
					res = _layer.Name
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
				if _layer, ok := stage.Layer_Ellipses_reverseMap[inst]; ok {
					res = _layer.Name
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
				if _svg, ok := stage.SVG_Layers_reverseMap[inst]; ok {
					res = _svg.Name
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
				if _layer, ok := stage.Layer_Lines_reverseMap[inst]; ok {
					res = _layer.Name
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
				if _layer, ok := stage.Layer_Links_reverseMap[inst]; ok {
					res = _layer.Name
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
				if _link, ok := stage.Link_TextAtArrowEnd_reverseMap[inst]; ok {
					res = _link.Name
				}
			case "TextAtArrowStart":
				if _link, ok := stage.Link_TextAtArrowStart_reverseMap[inst]; ok {
					res = _link.Name
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
				if _layer, ok := stage.Layer_Paths_reverseMap[inst]; ok {
					res = _layer.Name
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
				if _link, ok := stage.Link_ControlPoints_reverseMap[inst]; ok {
					res = _link.Name
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
				if _layer, ok := stage.Layer_Polygones_reverseMap[inst]; ok {
					res = _layer.Name
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
				if _layer, ok := stage.Layer_Polylines_reverseMap[inst]; ok {
					res = _layer.Name
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
				if _layer, ok := stage.Layer_Rects_reverseMap[inst]; ok {
					res = _layer.Name
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
				if _rect, ok := stage.Rect_RectAnchoredRects_reverseMap[inst]; ok {
					res = _rect.Name
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
				if _rect, ok := stage.Rect_RectAnchoredTexts_reverseMap[inst]; ok {
					res = _rect.Name
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
				if _layer, ok := stage.Layer_RectLinkLinks_reverseMap[inst]; ok {
					res = _layer.Name
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
				if _layer, ok := stage.Layer_Texts_reverseMap[inst]; ok {
					res = _layer.Name
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

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
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
				res = stage.Circle_Animations_reverseMap[inst]
			}
		case "Ellipse":
			switch reverseField.Fieldname {
			case "Animates":
				res = stage.Ellipse_Animates_reverseMap[inst]
			}
		case "Layer":
			switch reverseField.Fieldname {
			}
		case "Line":
			switch reverseField.Fieldname {
			case "Animates":
				res = stage.Line_Animates_reverseMap[inst]
			}
		case "Link":
			switch reverseField.Fieldname {
			}
		case "LinkAnchoredText":
			switch reverseField.Fieldname {
			case "Animates":
				res = stage.LinkAnchoredText_Animates_reverseMap[inst]
			}
		case "Path":
			switch reverseField.Fieldname {
			case "Animates":
				res = stage.Path_Animates_reverseMap[inst]
			}
		case "Point":
			switch reverseField.Fieldname {
			}
		case "Polygone":
			switch reverseField.Fieldname {
			case "Animates":
				res = stage.Polygone_Animates_reverseMap[inst]
			}
		case "Polyline":
			switch reverseField.Fieldname {
			case "Animates":
				res = stage.Polyline_Animates_reverseMap[inst]
			}
		case "Rect":
			switch reverseField.Fieldname {
			case "Animations":
				res = stage.Rect_Animations_reverseMap[inst]
			}
		case "RectAnchoredRect":
			switch reverseField.Fieldname {
			}
		case "RectAnchoredText":
			switch reverseField.Fieldname {
			case "Animates":
				res = stage.RectAnchoredText_Animates_reverseMap[inst]
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
				res = stage.Text_Animates_reverseMap[inst]
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
				res = stage.Layer_Circles_reverseMap[inst]
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
				res = stage.Layer_Ellipses_reverseMap[inst]
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
				res = stage.SVG_Layers_reverseMap[inst]
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
				res = stage.Layer_Lines_reverseMap[inst]
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
				res = stage.Layer_Links_reverseMap[inst]
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
				res = stage.Link_TextAtArrowEnd_reverseMap[inst]
			case "TextAtArrowStart":
				res = stage.Link_TextAtArrowStart_reverseMap[inst]
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
				res = stage.Layer_Paths_reverseMap[inst]
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
				res = stage.Link_ControlPoints_reverseMap[inst]
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
				res = stage.Layer_Polygones_reverseMap[inst]
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
				res = stage.Layer_Polylines_reverseMap[inst]
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
				res = stage.Layer_Rects_reverseMap[inst]
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
				res = stage.Rect_RectAnchoredRects_reverseMap[inst]
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
				res = stage.Rect_RectAnchoredTexts_reverseMap[inst]
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
				res = stage.Layer_RectLinkLinks_reverseMap[inst]
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
				res = stage.Layer_Texts_reverseMap[inst]
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
	return res
}
