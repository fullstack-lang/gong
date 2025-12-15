// generated code - do not edit
package models

// insertion point
func (inst *Category1) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Category1Shape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
		case "Diagram":
			switch reverseField.Fieldname {
			case "Category1Shapes":
				if _diagram, ok := stage.Diagram_Category1Shapes_reverseMap[inst]; ok {
					res = _diagram.Name
				}
			}
	}
	return
}

func (inst *Category2) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Category2Shape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
		case "Diagram":
			switch reverseField.Fieldname {
			case "Category2Shapes":
				if _diagram, ok := stage.Diagram_Category2Shapes_reverseMap[inst]; ok {
					res = _diagram.Name
				}
			}
	}
	return
}

func (inst *Category3) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Category3Shape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
		case "Diagram":
			switch reverseField.Fieldname {
			case "Category3Shapes":
				if _diagram, ok := stage.Diagram_Category3Shapes_reverseMap[inst]; ok {
					res = _diagram.Name
				}
			}
	}
	return
}

func (inst *ControlPointShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
		case "InfluenceShape":
			switch reverseField.Fieldname {
			case "ControlPointShapes":
				if _influenceshape, ok := stage.InfluenceShape_ControlPointShapes_reverseMap[inst]; ok {
					res = _influenceshape.Name
				}
			}
	}
	return
}

func (inst *Desk) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Diagram) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Influence) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *InfluenceShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
		case "Diagram":
			switch reverseField.Fieldname {
			case "InfluenceShapes":
				if _diagram, ok := stage.Diagram_InfluenceShapes_reverseMap[inst]; ok {
					res = _diagram.Name
				}
			}
	}
	return
}


// insertion point
func (inst *Category1) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Category1Shape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Diagram":
			switch reverseField.Fieldname {
			case "Category1Shapes":
				res = stage.Diagram_Category1Shapes_reverseMap[inst]
			}
	}
	return res
}

func (inst *Category2) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Category2Shape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Diagram":
			switch reverseField.Fieldname {
			case "Category2Shapes":
				res = stage.Diagram_Category2Shapes_reverseMap[inst]
			}
	}
	return res
}

func (inst *Category3) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Category3Shape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Diagram":
			switch reverseField.Fieldname {
			case "Category3Shapes":
				res = stage.Diagram_Category3Shapes_reverseMap[inst]
			}
	}
	return res
}

func (inst *ControlPointShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "InfluenceShape":
			switch reverseField.Fieldname {
			case "ControlPointShapes":
				res = stage.InfluenceShape_ControlPointShapes_reverseMap[inst]
			}
	}
	return res
}

func (inst *Desk) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Diagram) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Influence) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *InfluenceShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Diagram":
			switch reverseField.Fieldname {
			case "InfluenceShapes":
				res = stage.Diagram_InfluenceShapes_reverseMap[inst]
			}
	}
	return res
}

