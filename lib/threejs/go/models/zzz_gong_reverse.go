// generated code - do not edit
package models

// insertion point
func (inst *AmbiantLight) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *BoxGeometry) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *BufferGeometry) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Camera) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Canvas) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Curve) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CylinderGeometry) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DirectionalLight) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Canvas":
		switch reverseField.Fieldname {
		case "DirectionalLights":
			if _canvas, ok := stage.Canvas_DirectionalLights_reverseMap[inst]; ok {
				res = _canvas.Name
			}
		}
	}
	return
}

func (inst *ExtrudeGeometry) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Mesh) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Canvas":
		switch reverseField.Fieldname {
		case "Meshs":
			if _canvas, ok := stage.Canvas_Meshs_reverseMap[inst]; ok {
				res = _canvas.Name
			}
		}
	}
	return
}

func (inst *MeshMaterialBasic) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MeshPhysicalMaterial) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *PlaneGeometry) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Shape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *SphereGeometry) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TorusGeometry) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Triangle) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "BufferGeometry":
		switch reverseField.Fieldname {
		case "Faces":
			if _buffergeometry, ok := stage.BufferGeometry_Faces_reverseMap[inst]; ok {
				res = _buffergeometry.Name
			}
		}
	}
	return
}

func (inst *TubeGeometry) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Vector2) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Shape":
		switch reverseField.Fieldname {
		case "Points":
			if _shape, ok := stage.Shape_Points_reverseMap[inst]; ok {
				res = _shape.Name
			}
		}
	}
	return
}

func (inst *Vector3) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "BufferGeometry":
		switch reverseField.Fieldname {
		case "Vertices":
			if _buffergeometry, ok := stage.BufferGeometry_Vertices_reverseMap[inst]; ok {
				res = _buffergeometry.Name
			}
		}
	case "Curve":
		switch reverseField.Fieldname {
		case "Points":
			if _curve, ok := stage.Curve_Points_reverseMap[inst]; ok {
				res = _curve.Name
			}
		}
	}
	return
}
