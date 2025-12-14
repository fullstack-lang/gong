// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct ArtefactType
	// insertion point per field

	// Compute reverse map for named struct ArtefactTypeShape
	// insertion point per field

	// Compute reverse map for named struct Artist
	// insertion point per field

	// Compute reverse map for named struct ArtistShape
	// insertion point per field

	// Compute reverse map for named struct ControlPointShape
	// insertion point per field

	// Compute reverse map for named struct Desk
	// insertion point per field

	// Compute reverse map for named struct Diagram
	// insertion point per field
	stage.Diagram_MovementShapes_reverseMap = make(map[*MovementShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _movementshape := range diagram.MovementShapes {
			stage.Diagram_MovementShapes_reverseMap[_movementshape] = diagram
		}
	}
	stage.Diagram_ArtefactTypeShapes_reverseMap = make(map[*ArtefactTypeShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _artefacttypeshape := range diagram.ArtefactTypeShapes {
			stage.Diagram_ArtefactTypeShapes_reverseMap[_artefacttypeshape] = diagram
		}
	}
	stage.Diagram_ArtistShapes_reverseMap = make(map[*ArtistShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _artistshape := range diagram.ArtistShapes {
			stage.Diagram_ArtistShapes_reverseMap[_artistshape] = diagram
		}
	}
	stage.Diagram_InfluenceShapes_reverseMap = make(map[*InfluenceShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _influenceshape := range diagram.InfluenceShapes {
			stage.Diagram_InfluenceShapes_reverseMap[_influenceshape] = diagram
		}
	}

	// Compute reverse map for named struct Influence
	// insertion point per field

	// Compute reverse map for named struct InfluenceShape
	// insertion point per field
	stage.InfluenceShape_ControlPointShapes_reverseMap = make(map[*ControlPointShape]*InfluenceShape)
	for influenceshape := range stage.InfluenceShapes {
		_ = influenceshape
		for _, _controlpointshape := range influenceshape.ControlPointShapes {
			stage.InfluenceShape_ControlPointShapes_reverseMap[_controlpointshape] = influenceshape
		}
	}

	// Compute reverse map for named struct Movement
	// insertion point per field

	// Compute reverse map for named struct MovementShape
	// insertion point per field

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.ArtefactTypes {
		res = append(res, instance)
	}

	for instance := range stage.ArtefactTypeShapes {
		res = append(res, instance)
	}

	for instance := range stage.Artists {
		res = append(res, instance)
	}

	for instance := range stage.ArtistShapes {
		res = append(res, instance)
	}

	for instance := range stage.ControlPointShapes {
		res = append(res, instance)
	}

	for instance := range stage.Desks {
		res = append(res, instance)
	}

	for instance := range stage.Diagrams {
		res = append(res, instance)
	}

	for instance := range stage.Influences {
		res = append(res, instance)
	}

	for instance := range stage.InfluenceShapes {
		res = append(res, instance)
	}

	for instance := range stage.Movements {
		res = append(res, instance)
	}

	for instance := range stage.MovementShapes {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (artefacttype *ArtefactType) GongCopy() GongstructIF {
	newInstance := *artefacttype
	return &newInstance
}

func (artefacttypeshape *ArtefactTypeShape) GongCopy() GongstructIF {
	newInstance := *artefacttypeshape
	return &newInstance
}

func (artist *Artist) GongCopy() GongstructIF {
	newInstance := *artist
	return &newInstance
}

func (artistshape *ArtistShape) GongCopy() GongstructIF {
	newInstance := *artistshape
	return &newInstance
}

func (controlpointshape *ControlPointShape) GongCopy() GongstructIF {
	newInstance := *controlpointshape
	return &newInstance
}

func (desk *Desk) GongCopy() GongstructIF {
	newInstance := *desk
	return &newInstance
}

func (diagram *Diagram) GongCopy() GongstructIF {
	newInstance := *diagram
	return &newInstance
}

func (influence *Influence) GongCopy() GongstructIF {
	newInstance := *influence
	return &newInstance
}

func (influenceshape *InfluenceShape) GongCopy() GongstructIF {
	newInstance := *influenceshape
	return &newInstance
}

func (movement *Movement) GongCopy() GongstructIF {
	newInstance := *movement
	return &newInstance
}

func (movementshape *MovementShape) GongCopy() GongstructIF {
	newInstance := *movementshape
	return &newInstance
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {
	stage.reference = make(map[GongstructIF]GongstructIF)
	for _, instance := range stage.GetInstances() {
		stage.reference[instance] = instance.GongCopy()
	}
	stage.new = make(map[GongstructIF]struct{})
	stage.modified = make(map[GongstructIF]struct{})
	stage.deleted = make(map[GongstructIF]struct{})
}
