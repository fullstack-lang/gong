// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Category1
	// insertion point per field

	// Compute reverse map for named struct Category1Shape
	// insertion point per field

	// Compute reverse map for named struct Category2
	// insertion point per field

	// Compute reverse map for named struct Category2Shape
	// insertion point per field

	// Compute reverse map for named struct Category3
	// insertion point per field

	// Compute reverse map for named struct Category3Shape
	// insertion point per field

	// Compute reverse map for named struct ControlPointShape
	// insertion point per field

	// Compute reverse map for named struct Desk
	// insertion point per field

	// Compute reverse map for named struct Diagram
	// insertion point per field
	stage.Diagram_Category1Shapes_reverseMap = make(map[*Category1Shape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _category1shape := range diagram.Category1Shapes {
			stage.Diagram_Category1Shapes_reverseMap[_category1shape] = diagram
		}
	}
	stage.Diagram_Category2Shapes_reverseMap = make(map[*Category2Shape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _category2shape := range diagram.Category2Shapes {
			stage.Diagram_Category2Shapes_reverseMap[_category2shape] = diagram
		}
	}
	stage.Diagram_Category3Shapes_reverseMap = make(map[*Category3Shape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _category3shape := range diagram.Category3Shapes {
			stage.Diagram_Category3Shapes_reverseMap[_category3shape] = diagram
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

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.Category1s {
		res = append(res, instance)
	}

	for instance := range stage.Category1Shapes {
		res = append(res, instance)
	}

	for instance := range stage.Category2s {
		res = append(res, instance)
	}

	for instance := range stage.Category2Shapes {
		res = append(res, instance)
	}

	for instance := range stage.Category3s {
		res = append(res, instance)
	}

	for instance := range stage.Category3Shapes {
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

	return
}

// insertion point per named struct
func (category1 *Category1) GongCopy() GongstructIF {
	newInstance := *category1
	return &newInstance
}

func (category1shape *Category1Shape) GongCopy() GongstructIF {
	newInstance := *category1shape
	return &newInstance
}

func (category2 *Category2) GongCopy() GongstructIF {
	newInstance := *category2
	return &newInstance
}

func (category2shape *Category2Shape) GongCopy() GongstructIF {
	newInstance := *category2shape
	return &newInstance
}

func (category3 *Category3) GongCopy() GongstructIF {
	newInstance := *category3
	return &newInstance
}

func (category3shape *Category3Shape) GongCopy() GongstructIF {
	newInstance := *category3shape
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
