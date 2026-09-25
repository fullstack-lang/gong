// generated code - do not edit
package models

// insertion point
func (inst *ArtefactType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ArtefactTypeShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Artist) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ArtistShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *ControlPointShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Desk) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Diagram) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Influence) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *InfluenceShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Library) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "SubLibraries":
			if _library, ok := stage.Library_SubLibraries_reverseMap[inst]; ok {
				res = _library.Name
			}
		case "SubLibrariesWhoseNodeIsExpanded":
			if _library, ok := stage.Library_SubLibrariesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	}
	return
}

func (inst *Movement) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MovementShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Place) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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
