// generated code - do not edit
package models

func GetReverseFieldOwnerName(
	stage *Stage,
	instance any,
	reverseField *ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *Animate:
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

	case *Circle:
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

	case *Condition:
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				if _circle, ok := stage.Circle_HoveringTrigger_reverseMap[inst]; ok {
					res = _circle.Name
				}
			case "DisplayConditions":
				if _circle, ok := stage.Circle_DisplayConditions_reverseMap[inst]; ok {
					res = _circle.Name
				}
			}
		case "Ellipse":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				if _ellipse, ok := stage.Ellipse_HoveringTrigger_reverseMap[inst]; ok {
					res = _ellipse.Name
				}
			case "DisplayConditions":
				if _ellipse, ok := stage.Ellipse_DisplayConditions_reverseMap[inst]; ok {
					res = _ellipse.Name
				}
			}
		case "Line":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				if _line, ok := stage.Line_HoveringTrigger_reverseMap[inst]; ok {
					res = _line.Name
				}
			case "DisplayConditions":
				if _line, ok := stage.Line_DisplayConditions_reverseMap[inst]; ok {
					res = _line.Name
				}
			}
		case "Link":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				if _link, ok := stage.Link_HoveringTrigger_reverseMap[inst]; ok {
					res = _link.Name
				}
			case "DisplayConditions":
				if _link, ok := stage.Link_DisplayConditions_reverseMap[inst]; ok {
					res = _link.Name
				}
			}
		case "LinkAnchoredText":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				if _linkanchoredtext, ok := stage.LinkAnchoredText_HoveringTrigger_reverseMap[inst]; ok {
					res = _linkanchoredtext.Name
				}
			case "DisplayConditions":
				if _linkanchoredtext, ok := stage.LinkAnchoredText_DisplayConditions_reverseMap[inst]; ok {
					res = _linkanchoredtext.Name
				}
			}
		case "Path":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				if _path, ok := stage.Path_HoveringTrigger_reverseMap[inst]; ok {
					res = _path.Name
				}
			case "DisplayConditions":
				if _path, ok := stage.Path_DisplayConditions_reverseMap[inst]; ok {
					res = _path.Name
				}
			}
		case "Polygone":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				if _polygone, ok := stage.Polygone_HoveringTrigger_reverseMap[inst]; ok {
					res = _polygone.Name
				}
			case "DisplayConditions":
				if _polygone, ok := stage.Polygone_DisplayConditions_reverseMap[inst]; ok {
					res = _polygone.Name
				}
			}
		case "Polyline":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				if _polyline, ok := stage.Polyline_HoveringTrigger_reverseMap[inst]; ok {
					res = _polyline.Name
				}
			case "DisplayConditions":
				if _polyline, ok := stage.Polyline_DisplayConditions_reverseMap[inst]; ok {
					res = _polyline.Name
				}
			}
		case "Rect":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				if _rect, ok := stage.Rect_HoveringTrigger_reverseMap[inst]; ok {
					res = _rect.Name
				}
			case "DisplayConditions":
				if _rect, ok := stage.Rect_DisplayConditions_reverseMap[inst]; ok {
					res = _rect.Name
				}
			}
		case "RectAnchoredPath":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				if _rectanchoredpath, ok := stage.RectAnchoredPath_HoveringTrigger_reverseMap[inst]; ok {
					res = _rectanchoredpath.Name
				}
			case "DisplayConditions":
				if _rectanchoredpath, ok := stage.RectAnchoredPath_DisplayConditions_reverseMap[inst]; ok {
					res = _rectanchoredpath.Name
				}
			}
		case "RectAnchoredRect":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				if _rectanchoredrect, ok := stage.RectAnchoredRect_HoveringTrigger_reverseMap[inst]; ok {
					res = _rectanchoredrect.Name
				}
			case "DisplayConditions":
				if _rectanchoredrect, ok := stage.RectAnchoredRect_DisplayConditions_reverseMap[inst]; ok {
					res = _rectanchoredrect.Name
				}
			}
		case "RectAnchoredText":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				if _rectanchoredtext, ok := stage.RectAnchoredText_HoveringTrigger_reverseMap[inst]; ok {
					res = _rectanchoredtext.Name
				}
			case "DisplayConditions":
				if _rectanchoredtext, ok := stage.RectAnchoredText_DisplayConditions_reverseMap[inst]; ok {
					res = _rectanchoredtext.Name
				}
			}
		case "RectLinkLink":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				if _rectlinklink, ok := stage.RectLinkLink_HoveringTrigger_reverseMap[inst]; ok {
					res = _rectlinklink.Name
				}
			case "DisplayConditions":
				if _rectlinklink, ok := stage.RectLinkLink_DisplayConditions_reverseMap[inst]; ok {
					res = _rectlinklink.Name
				}
			}
		case "Text":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				if _text, ok := stage.Text_HoveringTrigger_reverseMap[inst]; ok {
					res = _text.Name
				}
			case "DisplayConditions":
				if _text, ok := stage.Text_DisplayConditions_reverseMap[inst]; ok {
					res = _text.Name
				}
			}
		}

	case *Ellipse:
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

	case *Layer:
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

	case *Line:
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

	case *Link:
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

	case *LinkAnchoredText:
		switch reverseField.GongstructName {
		// insertion point
		case "Link":
			switch reverseField.Fieldname {
			case "TextAtArrowStart":
				if _link, ok := stage.Link_TextAtArrowStart_reverseMap[inst]; ok {
					res = _link.Name
				}
			case "TextAtArrowEnd":
				if _link, ok := stage.Link_TextAtArrowEnd_reverseMap[inst]; ok {
					res = _link.Name
				}
			}
		}

	case *Path:
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

	case *Point:
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

	case *Polygone:
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

	case *Polyline:
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

	case *Rect:
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

	case *RectAnchoredPath:
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

	case *RectAnchoredRect:
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

	case *RectAnchoredText:
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

	case *RectLinkLink:
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

	case *SVG:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *SvgText:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Text:
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

func GetReverseFieldOwner[T Gongstruct](
	stage *Stage,
	instance *T,
	reverseField *ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *Animate:
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

	case *Circle:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Circles":
				res = stage.Layer_Circles_reverseMap[inst]
			}
		}

	case *Condition:
		switch reverseField.GongstructName {
		// insertion point
		case "Circle":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				res = stage.Circle_HoveringTrigger_reverseMap[inst]
			case "DisplayConditions":
				res = stage.Circle_DisplayConditions_reverseMap[inst]
			}
		case "Ellipse":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				res = stage.Ellipse_HoveringTrigger_reverseMap[inst]
			case "DisplayConditions":
				res = stage.Ellipse_DisplayConditions_reverseMap[inst]
			}
		case "Line":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				res = stage.Line_HoveringTrigger_reverseMap[inst]
			case "DisplayConditions":
				res = stage.Line_DisplayConditions_reverseMap[inst]
			}
		case "Link":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				res = stage.Link_HoveringTrigger_reverseMap[inst]
			case "DisplayConditions":
				res = stage.Link_DisplayConditions_reverseMap[inst]
			}
		case "LinkAnchoredText":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				res = stage.LinkAnchoredText_HoveringTrigger_reverseMap[inst]
			case "DisplayConditions":
				res = stage.LinkAnchoredText_DisplayConditions_reverseMap[inst]
			}
		case "Path":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				res = stage.Path_HoveringTrigger_reverseMap[inst]
			case "DisplayConditions":
				res = stage.Path_DisplayConditions_reverseMap[inst]
			}
		case "Polygone":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				res = stage.Polygone_HoveringTrigger_reverseMap[inst]
			case "DisplayConditions":
				res = stage.Polygone_DisplayConditions_reverseMap[inst]
			}
		case "Polyline":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				res = stage.Polyline_HoveringTrigger_reverseMap[inst]
			case "DisplayConditions":
				res = stage.Polyline_DisplayConditions_reverseMap[inst]
			}
		case "Rect":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				res = stage.Rect_HoveringTrigger_reverseMap[inst]
			case "DisplayConditions":
				res = stage.Rect_DisplayConditions_reverseMap[inst]
			}
		case "RectAnchoredPath":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				res = stage.RectAnchoredPath_HoveringTrigger_reverseMap[inst]
			case "DisplayConditions":
				res = stage.RectAnchoredPath_DisplayConditions_reverseMap[inst]
			}
		case "RectAnchoredRect":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				res = stage.RectAnchoredRect_HoveringTrigger_reverseMap[inst]
			case "DisplayConditions":
				res = stage.RectAnchoredRect_DisplayConditions_reverseMap[inst]
			}
		case "RectAnchoredText":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				res = stage.RectAnchoredText_HoveringTrigger_reverseMap[inst]
			case "DisplayConditions":
				res = stage.RectAnchoredText_DisplayConditions_reverseMap[inst]
			}
		case "RectLinkLink":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				res = stage.RectLinkLink_HoveringTrigger_reverseMap[inst]
			case "DisplayConditions":
				res = stage.RectLinkLink_DisplayConditions_reverseMap[inst]
			}
		case "Text":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				res = stage.Text_HoveringTrigger_reverseMap[inst]
			case "DisplayConditions":
				res = stage.Text_DisplayConditions_reverseMap[inst]
			}
		}

	case *Ellipse:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Ellipses":
				res = stage.Layer_Ellipses_reverseMap[inst]
			}
		}

	case *Layer:
		switch reverseField.GongstructName {
		// insertion point
		case "SVG":
			switch reverseField.Fieldname {
			case "Layers":
				res = stage.SVG_Layers_reverseMap[inst]
			}
		}

	case *Line:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Lines":
				res = stage.Layer_Lines_reverseMap[inst]
			}
		}

	case *Link:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Links":
				res = stage.Layer_Links_reverseMap[inst]
			}
		}

	case *LinkAnchoredText:
		switch reverseField.GongstructName {
		// insertion point
		case "Link":
			switch reverseField.Fieldname {
			case "TextAtArrowStart":
				res = stage.Link_TextAtArrowStart_reverseMap[inst]
			case "TextAtArrowEnd":
				res = stage.Link_TextAtArrowEnd_reverseMap[inst]
			}
		}

	case *Path:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Paths":
				res = stage.Layer_Paths_reverseMap[inst]
			}
		}

	case *Point:
		switch reverseField.GongstructName {
		// insertion point
		case "Link":
			switch reverseField.Fieldname {
			case "ControlPoints":
				res = stage.Link_ControlPoints_reverseMap[inst]
			}
		}

	case *Polygone:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Polygones":
				res = stage.Layer_Polygones_reverseMap[inst]
			}
		}

	case *Polyline:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Polylines":
				res = stage.Layer_Polylines_reverseMap[inst]
			}
		}

	case *Rect:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Rects":
				res = stage.Layer_Rects_reverseMap[inst]
			}
		}

	case *RectAnchoredPath:
		switch reverseField.GongstructName {
		// insertion point
		case "Rect":
			switch reverseField.Fieldname {
			case "RectAnchoredPaths":
				res = stage.Rect_RectAnchoredPaths_reverseMap[inst]
			}
		}

	case *RectAnchoredRect:
		switch reverseField.GongstructName {
		// insertion point
		case "Rect":
			switch reverseField.Fieldname {
			case "RectAnchoredRects":
				res = stage.Rect_RectAnchoredRects_reverseMap[inst]
			}
		}

	case *RectAnchoredText:
		switch reverseField.GongstructName {
		// insertion point
		case "Rect":
			switch reverseField.Fieldname {
			case "RectAnchoredTexts":
				res = stage.Rect_RectAnchoredTexts_reverseMap[inst]
			}
		}

	case *RectLinkLink:
		switch reverseField.GongstructName {
		// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "RectLinkLinks":
				res = stage.Layer_RectLinkLinks_reverseMap[inst]
			}
		}

	case *SVG:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *SvgText:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Text:
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
