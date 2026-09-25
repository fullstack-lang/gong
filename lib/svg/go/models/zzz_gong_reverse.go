// generated code - do not edit
package models

// insertion point
func (inst *Animate) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Circle) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Condition) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *ControlPoint) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Ellipse) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *FileToDownload) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Layer) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Line) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Link) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *LinkAnchoredPath) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Link":
		switch reverseField.Fieldname {
		case "PathAtArrowStart":
			if _link, ok := stage.Link_PathAtArrowStart_reverseMap[inst]; ok {
				res = _link.Name
			}
		case "PathAtArrowEnd":
			if _link, ok := stage.Link_PathAtArrowEnd_reverseMap[inst]; ok {
				res = _link.Name
			}
		case "PathAtCorner":
			if _link, ok := stage.Link_PathAtCorner_reverseMap[inst]; ok {
				res = _link.Name
			}
		}
	}
	return
}

func (inst *LinkAnchoredText) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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
		case "TextAtCorner":
			if _link, ok := stage.Link_TextAtCorner_reverseMap[inst]; ok {
				res = _link.Name
			}
		}
	}
	return
}

func (inst *Path) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Point) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Polygone) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Polyline) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Rect) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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
	case "Rect":
		switch reverseField.Fieldname {
		case "Peers":
			if _rect, ok := stage.Rect_Peers_reverseMap[inst]; ok {
				res = _rect.Name
			}
		case "Obstacles":
			if _rect, ok := stage.Rect_Obstacles_reverseMap[inst]; ok {
				res = _rect.Name
			}
		}
	}
	return
}

func (inst *RectAnchoredPath) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *RectAnchoredPngImage) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Rect":
		switch reverseField.Fieldname {
		case "RectAnchoredPngImages":
			if _rect, ok := stage.Rect_RectAnchoredPngImages_reverseMap[inst]; ok {
				res = _rect.Name
			}
		}
	}
	return
}

func (inst *RectAnchoredRect) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *RectAnchoredText) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *RectLinkLink) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *SVG) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *SvgText) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Text) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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
