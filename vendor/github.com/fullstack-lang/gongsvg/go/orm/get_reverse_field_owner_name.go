// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongsvg/go/models"
)

func GetReverseFieldOwnerName(
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance any,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Animate:
		switch reverseField.GongstructName {
		// insertion point
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
		case "Line":
			switch reverseField.Fieldname {
			case "Animates":
				if _line, ok := stage.Line_Animates_reverseMap[inst]; ok {
					res = _line.Name
				}
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
		case "RectAnchoredText":
			switch reverseField.Fieldname {
			case "Animates":
				if _rectanchoredtext, ok := stage.RectAnchoredText_Animates_reverseMap[inst]; ok {
					res = _rectanchoredtext.Name
				}
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
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Circles":
				if _layer, ok := stage.Layer_Circles_reverseMap[inst]; ok {
					res = _layer.Name
				}
			}
		}

	case *models.Ellipse:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Ellipses":
				if _layer, ok := stage.Layer_Ellipses_reverseMap[inst]; ok {
					res = _layer.Name
				}
			}
		}

	case *models.Layer:
		switch reverseField.GongstructName {
		// insertion point
		case "SVG":
			switch reverseField.Fieldname {
			case "Layers":
				if _svg, ok := stage.SVG_Layers_reverseMap[inst]; ok {
					res = _svg.Name
				}
			}
		}

	case *models.Line:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Lines":
				if _layer, ok := stage.Layer_Lines_reverseMap[inst]; ok {
					res = _layer.Name
				}
			}
		}

	case *models.Link:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Links":
				if _layer, ok := stage.Layer_Links_reverseMap[inst]; ok {
					res = _layer.Name
				}
			}
		}

	case *models.LinkAnchoredText:
		switch reverseField.GongstructName {
		// insertion point
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
		}

	case *models.Path:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Paths":
				if _layer, ok := stage.Layer_Paths_reverseMap[inst]; ok {
					res = _layer.Name
				}
			}
		}

	case *models.Point:
		switch reverseField.GongstructName {
		// insertion point
		case "Link":
			switch reverseField.Fieldname {
			case "ControlPoints":
				if _link, ok := stage.Link_ControlPoints_reverseMap[inst]; ok {
					res = _link.Name
				}
			}
		}

	case *models.Polygone:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Polygones":
				if _layer, ok := stage.Layer_Polygones_reverseMap[inst]; ok {
					res = _layer.Name
				}
			}
		}

	case *models.Polyline:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Polylines":
				if _layer, ok := stage.Layer_Polylines_reverseMap[inst]; ok {
					res = _layer.Name
				}
			}
		}

	case *models.Rect:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Rects":
				if _layer, ok := stage.Layer_Rects_reverseMap[inst]; ok {
					res = _layer.Name
				}
			}
		}

	case *models.RectAnchoredPath:
		switch reverseField.GongstructName {
		// insertion point
		case "Rect":
			switch reverseField.Fieldname {
			case "RectAnchoredPaths":
				if _rect, ok := stage.Rect_RectAnchoredPaths_reverseMap[inst]; ok {
					res = _rect.Name
				}
			}
		}

	case *models.RectAnchoredRect:
		switch reverseField.GongstructName {
		// insertion point
		case "Rect":
			switch reverseField.Fieldname {
			case "RectAnchoredRects":
				if _rect, ok := stage.Rect_RectAnchoredRects_reverseMap[inst]; ok {
					res = _rect.Name
				}
			}
		}

	case *models.RectAnchoredText:
		switch reverseField.GongstructName {
		// insertion point
		case "Rect":
			switch reverseField.Fieldname {
			case "RectAnchoredTexts":
				if _rect, ok := stage.Rect_RectAnchoredTexts_reverseMap[inst]; ok {
					res = _rect.Name
				}
			}
		}

	case *models.RectLinkLink:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "RectLinkLinks":
				if _layer, ok := stage.Layer_RectLinkLinks_reverseMap[inst]; ok {
					res = _layer.Name
				}
			}
		}

	case *models.SVG:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Text:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Texts":
				if _layer, ok := stage.Layer_Texts_reverseMap[inst]; ok {
					res = _layer.Name
				}
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
		switch reverseField.GongstructName {
		// insertion point
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
		case "Line":
			switch reverseField.Fieldname {
			case "Animates":
				res = stage.Line_Animates_reverseMap[inst]
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
		case "RectAnchoredText":
			switch reverseField.Fieldname {
			case "Animates":
				res = stage.RectAnchoredText_Animates_reverseMap[inst]
			}
		case "Text":
			switch reverseField.Fieldname {
			case "Animates":
				res = stage.Text_Animates_reverseMap[inst]
			}
		}

	case *models.Circle:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Circles":
				res = stage.Layer_Circles_reverseMap[inst]
			}
		}

	case *models.Ellipse:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Ellipses":
				res = stage.Layer_Ellipses_reverseMap[inst]
			}
		}

	case *models.Layer:
		switch reverseField.GongstructName {
		// insertion point
		case "SVG":
			switch reverseField.Fieldname {
			case "Layers":
				res = stage.SVG_Layers_reverseMap[inst]
			}
		}

	case *models.Line:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Lines":
				res = stage.Layer_Lines_reverseMap[inst]
			}
		}

	case *models.Link:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Links":
				res = stage.Layer_Links_reverseMap[inst]
			}
		}

	case *models.LinkAnchoredText:
		switch reverseField.GongstructName {
		// insertion point
		case "Link":
			switch reverseField.Fieldname {
			case "TextAtArrowEnd":
				res = stage.Link_TextAtArrowEnd_reverseMap[inst]
			case "TextAtArrowStart":
				res = stage.Link_TextAtArrowStart_reverseMap[inst]
			}
		}

	case *models.Path:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Paths":
				res = stage.Layer_Paths_reverseMap[inst]
			}
		}

	case *models.Point:
		switch reverseField.GongstructName {
		// insertion point
		case "Link":
			switch reverseField.Fieldname {
			case "ControlPoints":
				res = stage.Link_ControlPoints_reverseMap[inst]
			}
		}

	case *models.Polygone:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Polygones":
				res = stage.Layer_Polygones_reverseMap[inst]
			}
		}

	case *models.Polyline:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Polylines":
				res = stage.Layer_Polylines_reverseMap[inst]
			}
		}

	case *models.Rect:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Rects":
				res = stage.Layer_Rects_reverseMap[inst]
			}
		}

	case *models.RectAnchoredPath:
		switch reverseField.GongstructName {
		// insertion point
		case "Rect":
			switch reverseField.Fieldname {
			case "RectAnchoredPaths":
				res = stage.Rect_RectAnchoredPaths_reverseMap[inst]
			}
		}

	case *models.RectAnchoredRect:
		switch reverseField.GongstructName {
		// insertion point
		case "Rect":
			switch reverseField.Fieldname {
			case "RectAnchoredRects":
				res = stage.Rect_RectAnchoredRects_reverseMap[inst]
			}
		}

	case *models.RectAnchoredText:
		switch reverseField.GongstructName {
		// insertion point
		case "Rect":
			switch reverseField.Fieldname {
			case "RectAnchoredTexts":
				res = stage.Rect_RectAnchoredTexts_reverseMap[inst]
			}
		}

	case *models.RectLinkLink:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "RectLinkLinks":
				res = stage.Layer_RectLinkLinks_reverseMap[inst]
			}
		}

	case *models.SVG:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Text:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Texts":
				res = stage.Layer_Texts_reverseMap[inst]
			}
		}

	default:
		_ = inst
	}
	return res
}
