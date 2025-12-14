// generated code - do not edit
package models

// insertion point
func (inst *ArtefactType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *ArtefactTypeShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
		case "Diagram":
			switch reverseField.Fieldname {
			case "ArtefactTypeShapes":
				if _diagram, ok := stage.Diagram_ArtefactTypeShapes_reverseMap[inst]; ok {
					res = _diagram.Name
				}
			}
	}
	return
}

func (inst *Artist) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *ArtistShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
		case "Diagram":
			switch reverseField.Fieldname {
			case "ArtistShapes":
				if _diagram, ok := stage.Diagram_ArtistShapes_reverseMap[inst]; ok {
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

func (inst *Movement) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *MovementShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
		case "Diagram":
			switch reverseField.Fieldname {
			case "MovementShapes":
				if _diagram, ok := stage.Diagram_MovementShapes_reverseMap[inst]; ok {
					res = _diagram.Name
				}
			}
	}
	return
}

func (inst *Place) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
		case "Movement":
			switch reverseField.Fieldname {
			case "Places":
				if _movement, ok := stage.Movement_Places_reverseMap[inst]; ok {
					res = _movement.Name
				}
			}
	}
	return
}


// insertion point
func (inst *ArtefactType) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *ArtefactTypeShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Diagram":
			switch reverseField.Fieldname {
			case "ArtefactTypeShapes":
				res = stage.Diagram_ArtefactTypeShapes_reverseMap[inst]
			}
	}
	return res
}

func (inst *Artist) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *ArtistShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Diagram":
			switch reverseField.Fieldname {
			case "ArtistShapes":
				res = stage.Diagram_ArtistShapes_reverseMap[inst]
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

func (inst *Movement) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *MovementShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Diagram":
			switch reverseField.Fieldname {
			case "MovementShapes":
				res = stage.Diagram_MovementShapes_reverseMap[inst]
			}
	}
	return res
}

func (inst *Place) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Movement":
			switch reverseField.Fieldname {
			case "Places":
				res = stage.Movement_Places_reverseMap[inst]
			}
	}
	return res
}

