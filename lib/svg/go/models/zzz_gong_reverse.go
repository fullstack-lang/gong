// generated code - do not edit
package models

// insertion point
func (inst *Animate) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
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
	return
}

func (inst *Circle) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
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
	return
}

func (inst *Condition) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
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
	}
	return
}

func (inst *ControlPoint) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
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
	return
}

func (inst *Ellipse) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
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
	return
}

func (inst *Layer) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
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
	return
}

func (inst *Line) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
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
	return
}

func (inst *Link) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
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
	return
}

func (inst *LinkAnchoredText) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
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
	return
}

func (inst *Path) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
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
	return
}

func (inst *Point) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Polygone) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
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
	return
}

func (inst *Polyline) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
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
	return
}

func (inst *Rect) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
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
	return
}

func (inst *RectAnchoredPath) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
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
	return
}

func (inst *RectAnchoredRect) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
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
	return
}

func (inst *RectAnchoredText) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
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
	return
}

func (inst *RectLinkLink) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
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
	return
}

func (inst *SVG) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *SvgText) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Text) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
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
	return
}


// insertion point
func (inst *Animate) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
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
	return res
}

func (inst *Circle) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Circles":
				res = stage.Layer_Circles_reverseMap[inst]
			}
	}
	return res
}

func (inst *Condition) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Rect":
			switch reverseField.Fieldname {
			case "HoveringTrigger":
				res = stage.Rect_HoveringTrigger_reverseMap[inst]
			case "DisplayConditions":
				res = stage.Rect_DisplayConditions_reverseMap[inst]
			}
	}
	return res
}

func (inst *ControlPoint) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Link":
			switch reverseField.Fieldname {
			case "ControlPoints":
				res = stage.Link_ControlPoints_reverseMap[inst]
			}
	}
	return res
}

func (inst *Ellipse) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Ellipses":
				res = stage.Layer_Ellipses_reverseMap[inst]
			}
	}
	return res
}

func (inst *Layer) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "SVG":
			switch reverseField.Fieldname {
			case "Layers":
				res = stage.SVG_Layers_reverseMap[inst]
			}
	}
	return res
}

func (inst *Line) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Lines":
				res = stage.Layer_Lines_reverseMap[inst]
			}
	}
	return res
}

func (inst *Link) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Links":
				res = stage.Layer_Links_reverseMap[inst]
			}
	}
	return res
}

func (inst *LinkAnchoredText) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
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
	return res
}

func (inst *Path) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Paths":
				res = stage.Layer_Paths_reverseMap[inst]
			}
	}
	return res
}

func (inst *Point) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Polygone) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Polygones":
				res = stage.Layer_Polygones_reverseMap[inst]
			}
	}
	return res
}

func (inst *Polyline) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Polylines":
				res = stage.Layer_Polylines_reverseMap[inst]
			}
	}
	return res
}

func (inst *Rect) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Rects":
				res = stage.Layer_Rects_reverseMap[inst]
			}
	}
	return res
}

func (inst *RectAnchoredPath) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Rect":
			switch reverseField.Fieldname {
			case "RectAnchoredPaths":
				res = stage.Rect_RectAnchoredPaths_reverseMap[inst]
			}
	}
	return res
}

func (inst *RectAnchoredRect) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Rect":
			switch reverseField.Fieldname {
			case "RectAnchoredRects":
				res = stage.Rect_RectAnchoredRects_reverseMap[inst]
			}
	}
	return res
}

func (inst *RectAnchoredText) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Rect":
			switch reverseField.Fieldname {
			case "RectAnchoredTexts":
				res = stage.Rect_RectAnchoredTexts_reverseMap[inst]
			}
	}
	return res
}

func (inst *RectLinkLink) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "RectLinkLinks":
				res = stage.Layer_RectLinkLinks_reverseMap[inst]
			}
	}
	return res
}

func (inst *SVG) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *SvgText) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Text) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Layer":
			switch reverseField.Fieldname {
			case "Texts":
				res = stage.Layer_Texts_reverseMap[inst]
			}
	}
	return res
}

