// generated code - do not edit
package models

// insertion point
func (inst *Coordinate) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Polygon":
		switch reverseField.Fieldname {
		case "Coordinates":
			if _polygon, ok := stage.Polygon_Coordinates_reverseMap[inst]; ok {
				res = _polygon.Name
			}
		}
	case "Polyline":
		switch reverseField.Fieldname {
		case "Coordinates":
			if _polyline, ok := stage.Polyline_Coordinates_reverseMap[inst]; ok {
				res = _polyline.Name
			}
		}
	}
	return
}

func (inst *Layer) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "MapView":
		switch reverseField.Fieldname {
		case "Layers":
			if _mapview, ok := stage.MapView_Layers_reverseMap[inst]; ok {
				res = _mapview.Name
			}
		}
	}
	return
}

func (inst *MapView) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Marker) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Layer":
		switch reverseField.Fieldname {
		case "Markers":
			if _layer, ok := stage.Layer_Markers_reverseMap[inst]; ok {
				res = _layer.Name
			}
		}
	}
	return
}

func (inst *Polygon) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Layer":
		switch reverseField.Fieldname {
		case "Polygons":
			if _layer, ok := stage.Layer_Polygons_reverseMap[inst]; ok {
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
