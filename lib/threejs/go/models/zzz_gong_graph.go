// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *AmbiantLight:
		ok = stage.IsStagedAmbiantLight(target)

	case *BoxGeometry:
		ok = stage.IsStagedBoxGeometry(target)

	case *Canvas:
		ok = stage.IsStagedCanvas(target)

	case *Curve:
		ok = stage.IsStagedCurve(target)

	case *CylinderGeometry:
		ok = stage.IsStagedCylinderGeometry(target)

	case *DirectionalLight:
		ok = stage.IsStagedDirectionalLight(target)

	case *ExtrudeGeometry:
		ok = stage.IsStagedExtrudeGeometry(target)

	case *Mesh:
		ok = stage.IsStagedMesh(target)

	case *MeshMaterialBasic:
		ok = stage.IsStagedMeshMaterialBasic(target)

	case *MeshPhysicalMaterial:
		ok = stage.IsStagedMeshPhysicalMaterial(target)

	case *PlaneGeometry:
		ok = stage.IsStagedPlaneGeometry(target)

	case *Shape:
		ok = stage.IsStagedShape(target)

	case *SphereGeometry:
		ok = stage.IsStagedSphereGeometry(target)

	case *TorusGeometry:
		ok = stage.IsStagedTorusGeometry(target)

	case *TubeGeometry:
		ok = stage.IsStagedTubeGeometry(target)

	case *Vector2:
		ok = stage.IsStagedVector2(target)

	case *Vector3:
		ok = stage.IsStagedVector3(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *AmbiantLight:
		ok = stage.IsStagedAmbiantLight(target)

	case *BoxGeometry:
		ok = stage.IsStagedBoxGeometry(target)

	case *Canvas:
		ok = stage.IsStagedCanvas(target)

	case *Curve:
		ok = stage.IsStagedCurve(target)

	case *CylinderGeometry:
		ok = stage.IsStagedCylinderGeometry(target)

	case *DirectionalLight:
		ok = stage.IsStagedDirectionalLight(target)

	case *ExtrudeGeometry:
		ok = stage.IsStagedExtrudeGeometry(target)

	case *Mesh:
		ok = stage.IsStagedMesh(target)

	case *MeshMaterialBasic:
		ok = stage.IsStagedMeshMaterialBasic(target)

	case *MeshPhysicalMaterial:
		ok = stage.IsStagedMeshPhysicalMaterial(target)

	case *PlaneGeometry:
		ok = stage.IsStagedPlaneGeometry(target)

	case *Shape:
		ok = stage.IsStagedShape(target)

	case *SphereGeometry:
		ok = stage.IsStagedSphereGeometry(target)

	case *TorusGeometry:
		ok = stage.IsStagedTorusGeometry(target)

	case *TubeGeometry:
		ok = stage.IsStagedTubeGeometry(target)

	case *Vector2:
		ok = stage.IsStagedVector2(target)

	case *Vector3:
		ok = stage.IsStagedVector3(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedAmbiantLight(ambiantlight *AmbiantLight) (ok bool) {

	_, ok = stage.AmbiantLights[ambiantlight]

	return
}

func (stage *Stage) IsStagedBoxGeometry(boxgeometry *BoxGeometry) (ok bool) {

	_, ok = stage.BoxGeometrys[boxgeometry]

	return
}

func (stage *Stage) IsStagedCanvas(canvas *Canvas) (ok bool) {

	_, ok = stage.Canvass[canvas]

	return
}

func (stage *Stage) IsStagedCurve(curve *Curve) (ok bool) {

	_, ok = stage.Curves[curve]

	return
}

func (stage *Stage) IsStagedCylinderGeometry(cylindergeometry *CylinderGeometry) (ok bool) {

	_, ok = stage.CylinderGeometrys[cylindergeometry]

	return
}

func (stage *Stage) IsStagedDirectionalLight(directionallight *DirectionalLight) (ok bool) {

	_, ok = stage.DirectionalLights[directionallight]

	return
}

func (stage *Stage) IsStagedExtrudeGeometry(extrudegeometry *ExtrudeGeometry) (ok bool) {

	_, ok = stage.ExtrudeGeometrys[extrudegeometry]

	return
}

func (stage *Stage) IsStagedMesh(mesh *Mesh) (ok bool) {

	_, ok = stage.Meshs[mesh]

	return
}

func (stage *Stage) IsStagedMeshMaterialBasic(meshmaterialbasic *MeshMaterialBasic) (ok bool) {

	_, ok = stage.MeshMaterialBasics[meshmaterialbasic]

	return
}

func (stage *Stage) IsStagedMeshPhysicalMaterial(meshphysicalmaterial *MeshPhysicalMaterial) (ok bool) {

	_, ok = stage.MeshPhysicalMaterials[meshphysicalmaterial]

	return
}

func (stage *Stage) IsStagedPlaneGeometry(planegeometry *PlaneGeometry) (ok bool) {

	_, ok = stage.PlaneGeometrys[planegeometry]

	return
}

func (stage *Stage) IsStagedShape(shape *Shape) (ok bool) {

	_, ok = stage.Shapes[shape]

	return
}

func (stage *Stage) IsStagedSphereGeometry(spheregeometry *SphereGeometry) (ok bool) {

	_, ok = stage.SphereGeometrys[spheregeometry]

	return
}

func (stage *Stage) IsStagedTorusGeometry(torusgeometry *TorusGeometry) (ok bool) {

	_, ok = stage.TorusGeometrys[torusgeometry]

	return
}

func (stage *Stage) IsStagedTubeGeometry(tubegeometry *TubeGeometry) (ok bool) {

	_, ok = stage.TubeGeometrys[tubegeometry]

	return
}

func (stage *Stage) IsStagedVector2(vector2 *Vector2) (ok bool) {

	_, ok = stage.Vector2s[vector2]

	return
}

func (stage *Stage) IsStagedVector3(vector3 *Vector3) (ok bool) {

	_, ok = stage.Vector3s[vector3]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *AmbiantLight:
		stage.StageBranchAmbiantLight(target)

	case *BoxGeometry:
		stage.StageBranchBoxGeometry(target)

	case *Canvas:
		stage.StageBranchCanvas(target)

	case *Curve:
		stage.StageBranchCurve(target)

	case *CylinderGeometry:
		stage.StageBranchCylinderGeometry(target)

	case *DirectionalLight:
		stage.StageBranchDirectionalLight(target)

	case *ExtrudeGeometry:
		stage.StageBranchExtrudeGeometry(target)

	case *Mesh:
		stage.StageBranchMesh(target)

	case *MeshMaterialBasic:
		stage.StageBranchMeshMaterialBasic(target)

	case *MeshPhysicalMaterial:
		stage.StageBranchMeshPhysicalMaterial(target)

	case *PlaneGeometry:
		stage.StageBranchPlaneGeometry(target)

	case *Shape:
		stage.StageBranchShape(target)

	case *SphereGeometry:
		stage.StageBranchSphereGeometry(target)

	case *TorusGeometry:
		stage.StageBranchTorusGeometry(target)

	case *TubeGeometry:
		stage.StageBranchTubeGeometry(target)

	case *Vector2:
		stage.StageBranchVector2(target)

	case *Vector3:
		stage.StageBranchVector3(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchAmbiantLight(ambiantlight *AmbiantLight) {

	// check if instance is already staged
	if IsStaged(stage, ambiantlight) {
		return
	}

	ambiantlight.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchBoxGeometry(boxgeometry *BoxGeometry) {

	// check if instance is already staged
	if IsStaged(stage, boxgeometry) {
		return
	}

	boxgeometry.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchCanvas(canvas *Canvas) {

	// check if instance is already staged
	if IsStaged(stage, canvas) {
		return
	}

	canvas.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if canvas.AmbiantLight != nil {
		StageBranch(stage, canvas.AmbiantLight)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _directionallight := range canvas.DirectionalLights {
		StageBranch(stage, _directionallight)
	}
	for _, _mesh := range canvas.Meshs {
		StageBranch(stage, _mesh)
	}

}

func (stage *Stage) StageBranchCurve(curve *Curve) {

	// check if instance is already staged
	if IsStaged(stage, curve) {
		return
	}

	curve.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _vector3 := range curve.Points {
		StageBranch(stage, _vector3)
	}

}

func (stage *Stage) StageBranchCylinderGeometry(cylindergeometry *CylinderGeometry) {

	// check if instance is already staged
	if IsStaged(stage, cylindergeometry) {
		return
	}

	cylindergeometry.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchDirectionalLight(directionallight *DirectionalLight) {

	// check if instance is already staged
	if IsStaged(stage, directionallight) {
		return
	}

	directionallight.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchExtrudeGeometry(extrudegeometry *ExtrudeGeometry) {

	// check if instance is already staged
	if IsStaged(stage, extrudegeometry) {
		return
	}

	extrudegeometry.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if extrudegeometry.Shape != nil {
		StageBranch(stage, extrudegeometry.Shape)
	}
	if extrudegeometry.ExtrudePath != nil {
		StageBranch(stage, extrudegeometry.ExtrudePath)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchMesh(mesh *Mesh) {

	// check if instance is already staged
	if IsStaged(stage, mesh) {
		return
	}

	mesh.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if mesh.MeshMaterialBasic != nil {
		StageBranch(stage, mesh.MeshMaterialBasic)
	}
	if mesh.MeshPhysicalMaterial != nil {
		StageBranch(stage, mesh.MeshPhysicalMaterial)
	}
	if mesh.CylinderGeometry != nil {
		StageBranch(stage, mesh.CylinderGeometry)
	}
	if mesh.BoxGeometry != nil {
		StageBranch(stage, mesh.BoxGeometry)
	}
	if mesh.SphereGeometry != nil {
		StageBranch(stage, mesh.SphereGeometry)
	}
	if mesh.TorusGeometry != nil {
		StageBranch(stage, mesh.TorusGeometry)
	}
	if mesh.PlaneGeometry != nil {
		StageBranch(stage, mesh.PlaneGeometry)
	}
	if mesh.TubeGeometry != nil {
		StageBranch(stage, mesh.TubeGeometry)
	}
	if mesh.ExtrudeGeometry != nil {
		StageBranch(stage, mesh.ExtrudeGeometry)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchMeshMaterialBasic(meshmaterialbasic *MeshMaterialBasic) {

	// check if instance is already staged
	if IsStaged(stage, meshmaterialbasic) {
		return
	}

	meshmaterialbasic.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchMeshPhysicalMaterial(meshphysicalmaterial *MeshPhysicalMaterial) {

	// check if instance is already staged
	if IsStaged(stage, meshphysicalmaterial) {
		return
	}

	meshphysicalmaterial.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPlaneGeometry(planegeometry *PlaneGeometry) {

	// check if instance is already staged
	if IsStaged(stage, planegeometry) {
		return
	}

	planegeometry.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchShape(shape *Shape) {

	// check if instance is already staged
	if IsStaged(stage, shape) {
		return
	}

	shape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _vector2 := range shape.Points {
		StageBranch(stage, _vector2)
	}

}

func (stage *Stage) StageBranchSphereGeometry(spheregeometry *SphereGeometry) {

	// check if instance is already staged
	if IsStaged(stage, spheregeometry) {
		return
	}

	spheregeometry.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTorusGeometry(torusgeometry *TorusGeometry) {

	// check if instance is already staged
	if IsStaged(stage, torusgeometry) {
		return
	}

	torusgeometry.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTubeGeometry(tubegeometry *TubeGeometry) {

	// check if instance is already staged
	if IsStaged(stage, tubegeometry) {
		return
	}

	tubegeometry.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tubegeometry.Path != nil {
		StageBranch(stage, tubegeometry.Path)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchVector2(vector2 *Vector2) {

	// check if instance is already staged
	if IsStaged(stage, vector2) {
		return
	}

	vector2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchVector3(vector3 *Vector3) {

	// check if instance is already staged
	if IsStaged(stage, vector3) {
		return
	}

	vector3.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *AmbiantLight:
		toT := CopyBranchAmbiantLight(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BoxGeometry:
		toT := CopyBranchBoxGeometry(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Canvas:
		toT := CopyBranchCanvas(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Curve:
		toT := CopyBranchCurve(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CylinderGeometry:
		toT := CopyBranchCylinderGeometry(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DirectionalLight:
		toT := CopyBranchDirectionalLight(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ExtrudeGeometry:
		toT := CopyBranchExtrudeGeometry(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Mesh:
		toT := CopyBranchMesh(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MeshMaterialBasic:
		toT := CopyBranchMeshMaterialBasic(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MeshPhysicalMaterial:
		toT := CopyBranchMeshPhysicalMaterial(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PlaneGeometry:
		toT := CopyBranchPlaneGeometry(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Shape:
		toT := CopyBranchShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SphereGeometry:
		toT := CopyBranchSphereGeometry(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TorusGeometry:
		toT := CopyBranchTorusGeometry(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TubeGeometry:
		toT := CopyBranchTubeGeometry(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Vector2:
		toT := CopyBranchVector2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Vector3:
		toT := CopyBranchVector3(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchAmbiantLight(mapOrigCopy map[any]any, ambiantlightFrom *AmbiantLight) (ambiantlightTo *AmbiantLight) {

	// ambiantlightFrom has already been copied
	if _ambiantlightTo, ok := mapOrigCopy[ambiantlightFrom]; ok {
		ambiantlightTo = _ambiantlightTo.(*AmbiantLight)
		return
	}

	ambiantlightTo = new(AmbiantLight)
	mapOrigCopy[ambiantlightFrom] = ambiantlightTo
	ambiantlightFrom.CopyBasicFields(ambiantlightTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchBoxGeometry(mapOrigCopy map[any]any, boxgeometryFrom *BoxGeometry) (boxgeometryTo *BoxGeometry) {

	// boxgeometryFrom has already been copied
	if _boxgeometryTo, ok := mapOrigCopy[boxgeometryFrom]; ok {
		boxgeometryTo = _boxgeometryTo.(*BoxGeometry)
		return
	}

	boxgeometryTo = new(BoxGeometry)
	mapOrigCopy[boxgeometryFrom] = boxgeometryTo
	boxgeometryFrom.CopyBasicFields(boxgeometryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCanvas(mapOrigCopy map[any]any, canvasFrom *Canvas) (canvasTo *Canvas) {

	// canvasFrom has already been copied
	if _canvasTo, ok := mapOrigCopy[canvasFrom]; ok {
		canvasTo = _canvasTo.(*Canvas)
		return
	}

	canvasTo = new(Canvas)
	mapOrigCopy[canvasFrom] = canvasTo
	canvasFrom.CopyBasicFields(canvasTo)

	//insertion point for the staging of instances referenced by pointers
	if canvasFrom.AmbiantLight != nil {
		canvasTo.AmbiantLight = CopyBranchAmbiantLight(mapOrigCopy, canvasFrom.AmbiantLight)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _directionallight := range canvasFrom.DirectionalLights {
		canvasTo.DirectionalLights = append(canvasTo.DirectionalLights, CopyBranchDirectionalLight(mapOrigCopy, _directionallight))
	}
	for _, _mesh := range canvasFrom.Meshs {
		canvasTo.Meshs = append(canvasTo.Meshs, CopyBranchMesh(mapOrigCopy, _mesh))
	}

	return
}

func CopyBranchCurve(mapOrigCopy map[any]any, curveFrom *Curve) (curveTo *Curve) {

	// curveFrom has already been copied
	if _curveTo, ok := mapOrigCopy[curveFrom]; ok {
		curveTo = _curveTo.(*Curve)
		return
	}

	curveTo = new(Curve)
	mapOrigCopy[curveFrom] = curveTo
	curveFrom.CopyBasicFields(curveTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _vector3 := range curveFrom.Points {
		curveTo.Points = append(curveTo.Points, CopyBranchVector3(mapOrigCopy, _vector3))
	}

	return
}

func CopyBranchCylinderGeometry(mapOrigCopy map[any]any, cylindergeometryFrom *CylinderGeometry) (cylindergeometryTo *CylinderGeometry) {

	// cylindergeometryFrom has already been copied
	if _cylindergeometryTo, ok := mapOrigCopy[cylindergeometryFrom]; ok {
		cylindergeometryTo = _cylindergeometryTo.(*CylinderGeometry)
		return
	}

	cylindergeometryTo = new(CylinderGeometry)
	mapOrigCopy[cylindergeometryFrom] = cylindergeometryTo
	cylindergeometryFrom.CopyBasicFields(cylindergeometryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDirectionalLight(mapOrigCopy map[any]any, directionallightFrom *DirectionalLight) (directionallightTo *DirectionalLight) {

	// directionallightFrom has already been copied
	if _directionallightTo, ok := mapOrigCopy[directionallightFrom]; ok {
		directionallightTo = _directionallightTo.(*DirectionalLight)
		return
	}

	directionallightTo = new(DirectionalLight)
	mapOrigCopy[directionallightFrom] = directionallightTo
	directionallightFrom.CopyBasicFields(directionallightTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchExtrudeGeometry(mapOrigCopy map[any]any, extrudegeometryFrom *ExtrudeGeometry) (extrudegeometryTo *ExtrudeGeometry) {

	// extrudegeometryFrom has already been copied
	if _extrudegeometryTo, ok := mapOrigCopy[extrudegeometryFrom]; ok {
		extrudegeometryTo = _extrudegeometryTo.(*ExtrudeGeometry)
		return
	}

	extrudegeometryTo = new(ExtrudeGeometry)
	mapOrigCopy[extrudegeometryFrom] = extrudegeometryTo
	extrudegeometryFrom.CopyBasicFields(extrudegeometryTo)

	//insertion point for the staging of instances referenced by pointers
	if extrudegeometryFrom.Shape != nil {
		extrudegeometryTo.Shape = CopyBranchShape(mapOrigCopy, extrudegeometryFrom.Shape)
	}
	if extrudegeometryFrom.ExtrudePath != nil {
		extrudegeometryTo.ExtrudePath = CopyBranchCurve(mapOrigCopy, extrudegeometryFrom.ExtrudePath)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMesh(mapOrigCopy map[any]any, meshFrom *Mesh) (meshTo *Mesh) {

	// meshFrom has already been copied
	if _meshTo, ok := mapOrigCopy[meshFrom]; ok {
		meshTo = _meshTo.(*Mesh)
		return
	}

	meshTo = new(Mesh)
	mapOrigCopy[meshFrom] = meshTo
	meshFrom.CopyBasicFields(meshTo)

	//insertion point for the staging of instances referenced by pointers
	if meshFrom.MeshMaterialBasic != nil {
		meshTo.MeshMaterialBasic = CopyBranchMeshMaterialBasic(mapOrigCopy, meshFrom.MeshMaterialBasic)
	}
	if meshFrom.MeshPhysicalMaterial != nil {
		meshTo.MeshPhysicalMaterial = CopyBranchMeshPhysicalMaterial(mapOrigCopy, meshFrom.MeshPhysicalMaterial)
	}
	if meshFrom.CylinderGeometry != nil {
		meshTo.CylinderGeometry = CopyBranchCylinderGeometry(mapOrigCopy, meshFrom.CylinderGeometry)
	}
	if meshFrom.BoxGeometry != nil {
		meshTo.BoxGeometry = CopyBranchBoxGeometry(mapOrigCopy, meshFrom.BoxGeometry)
	}
	if meshFrom.SphereGeometry != nil {
		meshTo.SphereGeometry = CopyBranchSphereGeometry(mapOrigCopy, meshFrom.SphereGeometry)
	}
	if meshFrom.TorusGeometry != nil {
		meshTo.TorusGeometry = CopyBranchTorusGeometry(mapOrigCopy, meshFrom.TorusGeometry)
	}
	if meshFrom.PlaneGeometry != nil {
		meshTo.PlaneGeometry = CopyBranchPlaneGeometry(mapOrigCopy, meshFrom.PlaneGeometry)
	}
	if meshFrom.TubeGeometry != nil {
		meshTo.TubeGeometry = CopyBranchTubeGeometry(mapOrigCopy, meshFrom.TubeGeometry)
	}
	if meshFrom.ExtrudeGeometry != nil {
		meshTo.ExtrudeGeometry = CopyBranchExtrudeGeometry(mapOrigCopy, meshFrom.ExtrudeGeometry)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMeshMaterialBasic(mapOrigCopy map[any]any, meshmaterialbasicFrom *MeshMaterialBasic) (meshmaterialbasicTo *MeshMaterialBasic) {

	// meshmaterialbasicFrom has already been copied
	if _meshmaterialbasicTo, ok := mapOrigCopy[meshmaterialbasicFrom]; ok {
		meshmaterialbasicTo = _meshmaterialbasicTo.(*MeshMaterialBasic)
		return
	}

	meshmaterialbasicTo = new(MeshMaterialBasic)
	mapOrigCopy[meshmaterialbasicFrom] = meshmaterialbasicTo
	meshmaterialbasicFrom.CopyBasicFields(meshmaterialbasicTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMeshPhysicalMaterial(mapOrigCopy map[any]any, meshphysicalmaterialFrom *MeshPhysicalMaterial) (meshphysicalmaterialTo *MeshPhysicalMaterial) {

	// meshphysicalmaterialFrom has already been copied
	if _meshphysicalmaterialTo, ok := mapOrigCopy[meshphysicalmaterialFrom]; ok {
		meshphysicalmaterialTo = _meshphysicalmaterialTo.(*MeshPhysicalMaterial)
		return
	}

	meshphysicalmaterialTo = new(MeshPhysicalMaterial)
	mapOrigCopy[meshphysicalmaterialFrom] = meshphysicalmaterialTo
	meshphysicalmaterialFrom.CopyBasicFields(meshphysicalmaterialTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPlaneGeometry(mapOrigCopy map[any]any, planegeometryFrom *PlaneGeometry) (planegeometryTo *PlaneGeometry) {

	// planegeometryFrom has already been copied
	if _planegeometryTo, ok := mapOrigCopy[planegeometryFrom]; ok {
		planegeometryTo = _planegeometryTo.(*PlaneGeometry)
		return
	}

	planegeometryTo = new(PlaneGeometry)
	mapOrigCopy[planegeometryFrom] = planegeometryTo
	planegeometryFrom.CopyBasicFields(planegeometryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchShape(mapOrigCopy map[any]any, shapeFrom *Shape) (shapeTo *Shape) {

	// shapeFrom has already been copied
	if _shapeTo, ok := mapOrigCopy[shapeFrom]; ok {
		shapeTo = _shapeTo.(*Shape)
		return
	}

	shapeTo = new(Shape)
	mapOrigCopy[shapeFrom] = shapeTo
	shapeFrom.CopyBasicFields(shapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _vector2 := range shapeFrom.Points {
		shapeTo.Points = append(shapeTo.Points, CopyBranchVector2(mapOrigCopy, _vector2))
	}

	return
}

func CopyBranchSphereGeometry(mapOrigCopy map[any]any, spheregeometryFrom *SphereGeometry) (spheregeometryTo *SphereGeometry) {

	// spheregeometryFrom has already been copied
	if _spheregeometryTo, ok := mapOrigCopy[spheregeometryFrom]; ok {
		spheregeometryTo = _spheregeometryTo.(*SphereGeometry)
		return
	}

	spheregeometryTo = new(SphereGeometry)
	mapOrigCopy[spheregeometryFrom] = spheregeometryTo
	spheregeometryFrom.CopyBasicFields(spheregeometryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTorusGeometry(mapOrigCopy map[any]any, torusgeometryFrom *TorusGeometry) (torusgeometryTo *TorusGeometry) {

	// torusgeometryFrom has already been copied
	if _torusgeometryTo, ok := mapOrigCopy[torusgeometryFrom]; ok {
		torusgeometryTo = _torusgeometryTo.(*TorusGeometry)
		return
	}

	torusgeometryTo = new(TorusGeometry)
	mapOrigCopy[torusgeometryFrom] = torusgeometryTo
	torusgeometryFrom.CopyBasicFields(torusgeometryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTubeGeometry(mapOrigCopy map[any]any, tubegeometryFrom *TubeGeometry) (tubegeometryTo *TubeGeometry) {

	// tubegeometryFrom has already been copied
	if _tubegeometryTo, ok := mapOrigCopy[tubegeometryFrom]; ok {
		tubegeometryTo = _tubegeometryTo.(*TubeGeometry)
		return
	}

	tubegeometryTo = new(TubeGeometry)
	mapOrigCopy[tubegeometryFrom] = tubegeometryTo
	tubegeometryFrom.CopyBasicFields(tubegeometryTo)

	//insertion point for the staging of instances referenced by pointers
	if tubegeometryFrom.Path != nil {
		tubegeometryTo.Path = CopyBranchCurve(mapOrigCopy, tubegeometryFrom.Path)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchVector2(mapOrigCopy map[any]any, vector2From *Vector2) (vector2To *Vector2) {

	// vector2From has already been copied
	if _vector2To, ok := mapOrigCopy[vector2From]; ok {
		vector2To = _vector2To.(*Vector2)
		return
	}

	vector2To = new(Vector2)
	mapOrigCopy[vector2From] = vector2To
	vector2From.CopyBasicFields(vector2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchVector3(mapOrigCopy map[any]any, vector3From *Vector3) (vector3To *Vector3) {

	// vector3From has already been copied
	if _vector3To, ok := mapOrigCopy[vector3From]; ok {
		vector3To = _vector3To.(*Vector3)
		return
	}

	vector3To = new(Vector3)
	mapOrigCopy[vector3From] = vector3To
	vector3From.CopyBasicFields(vector3To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *AmbiantLight:
		stage.UnstageBranchAmbiantLight(target)

	case *BoxGeometry:
		stage.UnstageBranchBoxGeometry(target)

	case *Canvas:
		stage.UnstageBranchCanvas(target)

	case *Curve:
		stage.UnstageBranchCurve(target)

	case *CylinderGeometry:
		stage.UnstageBranchCylinderGeometry(target)

	case *DirectionalLight:
		stage.UnstageBranchDirectionalLight(target)

	case *ExtrudeGeometry:
		stage.UnstageBranchExtrudeGeometry(target)

	case *Mesh:
		stage.UnstageBranchMesh(target)

	case *MeshMaterialBasic:
		stage.UnstageBranchMeshMaterialBasic(target)

	case *MeshPhysicalMaterial:
		stage.UnstageBranchMeshPhysicalMaterial(target)

	case *PlaneGeometry:
		stage.UnstageBranchPlaneGeometry(target)

	case *Shape:
		stage.UnstageBranchShape(target)

	case *SphereGeometry:
		stage.UnstageBranchSphereGeometry(target)

	case *TorusGeometry:
		stage.UnstageBranchTorusGeometry(target)

	case *TubeGeometry:
		stage.UnstageBranchTubeGeometry(target)

	case *Vector2:
		stage.UnstageBranchVector2(target)

	case *Vector3:
		stage.UnstageBranchVector3(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchAmbiantLight(ambiantlight *AmbiantLight) {

	// check if instance is already staged
	if !IsStaged(stage, ambiantlight) {
		return
	}

	ambiantlight.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchBoxGeometry(boxgeometry *BoxGeometry) {

	// check if instance is already staged
	if !IsStaged(stage, boxgeometry) {
		return
	}

	boxgeometry.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchCanvas(canvas *Canvas) {

	// check if instance is already staged
	if !IsStaged(stage, canvas) {
		return
	}

	canvas.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if canvas.AmbiantLight != nil {
		UnstageBranch(stage, canvas.AmbiantLight)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _directionallight := range canvas.DirectionalLights {
		UnstageBranch(stage, _directionallight)
	}
	for _, _mesh := range canvas.Meshs {
		UnstageBranch(stage, _mesh)
	}

}

func (stage *Stage) UnstageBranchCurve(curve *Curve) {

	// check if instance is already staged
	if !IsStaged(stage, curve) {
		return
	}

	curve.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _vector3 := range curve.Points {
		UnstageBranch(stage, _vector3)
	}

}

func (stage *Stage) UnstageBranchCylinderGeometry(cylindergeometry *CylinderGeometry) {

	// check if instance is already staged
	if !IsStaged(stage, cylindergeometry) {
		return
	}

	cylindergeometry.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchDirectionalLight(directionallight *DirectionalLight) {

	// check if instance is already staged
	if !IsStaged(stage, directionallight) {
		return
	}

	directionallight.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchExtrudeGeometry(extrudegeometry *ExtrudeGeometry) {

	// check if instance is already staged
	if !IsStaged(stage, extrudegeometry) {
		return
	}

	extrudegeometry.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if extrudegeometry.Shape != nil {
		UnstageBranch(stage, extrudegeometry.Shape)
	}
	if extrudegeometry.ExtrudePath != nil {
		UnstageBranch(stage, extrudegeometry.ExtrudePath)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchMesh(mesh *Mesh) {

	// check if instance is already staged
	if !IsStaged(stage, mesh) {
		return
	}

	mesh.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if mesh.MeshMaterialBasic != nil {
		UnstageBranch(stage, mesh.MeshMaterialBasic)
	}
	if mesh.MeshPhysicalMaterial != nil {
		UnstageBranch(stage, mesh.MeshPhysicalMaterial)
	}
	if mesh.CylinderGeometry != nil {
		UnstageBranch(stage, mesh.CylinderGeometry)
	}
	if mesh.BoxGeometry != nil {
		UnstageBranch(stage, mesh.BoxGeometry)
	}
	if mesh.SphereGeometry != nil {
		UnstageBranch(stage, mesh.SphereGeometry)
	}
	if mesh.TorusGeometry != nil {
		UnstageBranch(stage, mesh.TorusGeometry)
	}
	if mesh.PlaneGeometry != nil {
		UnstageBranch(stage, mesh.PlaneGeometry)
	}
	if mesh.TubeGeometry != nil {
		UnstageBranch(stage, mesh.TubeGeometry)
	}
	if mesh.ExtrudeGeometry != nil {
		UnstageBranch(stage, mesh.ExtrudeGeometry)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchMeshMaterialBasic(meshmaterialbasic *MeshMaterialBasic) {

	// check if instance is already staged
	if !IsStaged(stage, meshmaterialbasic) {
		return
	}

	meshmaterialbasic.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchMeshPhysicalMaterial(meshphysicalmaterial *MeshPhysicalMaterial) {

	// check if instance is already staged
	if !IsStaged(stage, meshphysicalmaterial) {
		return
	}

	meshphysicalmaterial.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPlaneGeometry(planegeometry *PlaneGeometry) {

	// check if instance is already staged
	if !IsStaged(stage, planegeometry) {
		return
	}

	planegeometry.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchShape(shape *Shape) {

	// check if instance is already staged
	if !IsStaged(stage, shape) {
		return
	}

	shape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _vector2 := range shape.Points {
		UnstageBranch(stage, _vector2)
	}

}

func (stage *Stage) UnstageBranchSphereGeometry(spheregeometry *SphereGeometry) {

	// check if instance is already staged
	if !IsStaged(stage, spheregeometry) {
		return
	}

	spheregeometry.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTorusGeometry(torusgeometry *TorusGeometry) {

	// check if instance is already staged
	if !IsStaged(stage, torusgeometry) {
		return
	}

	torusgeometry.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTubeGeometry(tubegeometry *TubeGeometry) {

	// check if instance is already staged
	if !IsStaged(stage, tubegeometry) {
		return
	}

	tubegeometry.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tubegeometry.Path != nil {
		UnstageBranch(stage, tubegeometry.Path)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchVector2(vector2 *Vector2) {

	// check if instance is already staged
	if !IsStaged(stage, vector2) {
		return
	}

	vector2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchVector3(vector3 *Vector3) {

	// check if instance is already staged
	if !IsStaged(stage, vector3) {
		return
	}

	vector3.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *AmbiantLight) GongReconstructPointersFromReferences(stage *Stage, instance *AmbiantLight) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *BoxGeometry) GongReconstructPointersFromReferences(stage *Stage, instance *BoxGeometry) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Canvas) GongReconstructPointersFromReferences(stage *Stage, instance *Canvas) {
	// insertion point for pointers field
	if instance.AmbiantLight != nil {
		reference.AmbiantLight = stage.AmbiantLights_reference[instance.AmbiantLight]
	}
	// insertion point for slice of pointers field
	reference.DirectionalLights = reference.DirectionalLights[:0]
	for _, _b := range instance.DirectionalLights {
		reference.DirectionalLights = append(reference.DirectionalLights, stage.DirectionalLights_reference[_b])
	}
	reference.Meshs = reference.Meshs[:0]
	for _, _b := range instance.Meshs {
		reference.Meshs = append(reference.Meshs, stage.Meshs_reference[_b])
	}
}

func (reference *Curve) GongReconstructPointersFromReferences(stage *Stage, instance *Curve) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Points = reference.Points[:0]
	for _, _b := range instance.Points {
		reference.Points = append(reference.Points, stage.Vector3s_reference[_b])
	}
}

func (reference *CylinderGeometry) GongReconstructPointersFromReferences(stage *Stage, instance *CylinderGeometry) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *DirectionalLight) GongReconstructPointersFromReferences(stage *Stage, instance *DirectionalLight) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ExtrudeGeometry) GongReconstructPointersFromReferences(stage *Stage, instance *ExtrudeGeometry) {
	// insertion point for pointers field
	if instance.Shape != nil {
		reference.Shape = stage.Shapes_reference[instance.Shape]
	}
	if instance.ExtrudePath != nil {
		reference.ExtrudePath = stage.Curves_reference[instance.ExtrudePath]
	}
	// insertion point for slice of pointers field
}

func (reference *Mesh) GongReconstructPointersFromReferences(stage *Stage, instance *Mesh) {
	// insertion point for pointers field
	if instance.MeshMaterialBasic != nil {
		reference.MeshMaterialBasic = stage.MeshMaterialBasics_reference[instance.MeshMaterialBasic]
	}
	if instance.MeshPhysicalMaterial != nil {
		reference.MeshPhysicalMaterial = stage.MeshPhysicalMaterials_reference[instance.MeshPhysicalMaterial]
	}
	if instance.CylinderGeometry != nil {
		reference.CylinderGeometry = stage.CylinderGeometrys_reference[instance.CylinderGeometry]
	}
	if instance.BoxGeometry != nil {
		reference.BoxGeometry = stage.BoxGeometrys_reference[instance.BoxGeometry]
	}
	if instance.SphereGeometry != nil {
		reference.SphereGeometry = stage.SphereGeometrys_reference[instance.SphereGeometry]
	}
	if instance.TorusGeometry != nil {
		reference.TorusGeometry = stage.TorusGeometrys_reference[instance.TorusGeometry]
	}
	if instance.PlaneGeometry != nil {
		reference.PlaneGeometry = stage.PlaneGeometrys_reference[instance.PlaneGeometry]
	}
	if instance.TubeGeometry != nil {
		reference.TubeGeometry = stage.TubeGeometrys_reference[instance.TubeGeometry]
	}
	if instance.ExtrudeGeometry != nil {
		reference.ExtrudeGeometry = stage.ExtrudeGeometrys_reference[instance.ExtrudeGeometry]
	}
	// insertion point for slice of pointers field
}

func (reference *MeshMaterialBasic) GongReconstructPointersFromReferences(stage *Stage, instance *MeshMaterialBasic) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MeshPhysicalMaterial) GongReconstructPointersFromReferences(stage *Stage, instance *MeshPhysicalMaterial) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PlaneGeometry) GongReconstructPointersFromReferences(stage *Stage, instance *PlaneGeometry) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Shape) GongReconstructPointersFromReferences(stage *Stage, instance *Shape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Points = reference.Points[:0]
	for _, _b := range instance.Points {
		reference.Points = append(reference.Points, stage.Vector2s_reference[_b])
	}
}

func (reference *SphereGeometry) GongReconstructPointersFromReferences(stage *Stage, instance *SphereGeometry) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TorusGeometry) GongReconstructPointersFromReferences(stage *Stage, instance *TorusGeometry) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TubeGeometry) GongReconstructPointersFromReferences(stage *Stage, instance *TubeGeometry) {
	// insertion point for pointers field
	if instance.Path != nil {
		reference.Path = stage.Curves_reference[instance.Path]
	}
	// insertion point for slice of pointers field
}

func (reference *Vector2) GongReconstructPointersFromReferences(stage *Stage, instance *Vector2) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Vector3) GongReconstructPointersFromReferences(stage *Stage, instance *Vector3) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *AmbiantLight) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *BoxGeometry) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Canvas) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.AmbiantLight; _reference != nil {
		reference.AmbiantLight = nil
		if _instance, ok := stage.AmbiantLights_instance[_reference]; ok {
			reference.AmbiantLight = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _DirectionalLights []*DirectionalLight
	for _, _reference := range reference.DirectionalLights {
		if _instance, ok := stage.DirectionalLights_instance[_reference]; ok {
			_DirectionalLights = append(_DirectionalLights, _instance)
		}
	}
	reference.DirectionalLights = _DirectionalLights
	var _Meshs []*Mesh
	for _, _reference := range reference.Meshs {
		if _instance, ok := stage.Meshs_instance[_reference]; ok {
			_Meshs = append(_Meshs, _instance)
		}
	}
	reference.Meshs = _Meshs
}

func (reference *Curve) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Points []*Vector3
	for _, _reference := range reference.Points {
		if _instance, ok := stage.Vector3s_instance[_reference]; ok {
			_Points = append(_Points, _instance)
		}
	}
	reference.Points = _Points
}

func (reference *CylinderGeometry) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *DirectionalLight) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ExtrudeGeometry) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Shape; _reference != nil {
		reference.Shape = nil
		if _instance, ok := stage.Shapes_instance[_reference]; ok {
			reference.Shape = _instance
		}
	}
	if _reference := reference.ExtrudePath; _reference != nil {
		reference.ExtrudePath = nil
		if _instance, ok := stage.Curves_instance[_reference]; ok {
			reference.ExtrudePath = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Mesh) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.MeshMaterialBasic; _reference != nil {
		reference.MeshMaterialBasic = nil
		if _instance, ok := stage.MeshMaterialBasics_instance[_reference]; ok {
			reference.MeshMaterialBasic = _instance
		}
	}
	if _reference := reference.MeshPhysicalMaterial; _reference != nil {
		reference.MeshPhysicalMaterial = nil
		if _instance, ok := stage.MeshPhysicalMaterials_instance[_reference]; ok {
			reference.MeshPhysicalMaterial = _instance
		}
	}
	if _reference := reference.CylinderGeometry; _reference != nil {
		reference.CylinderGeometry = nil
		if _instance, ok := stage.CylinderGeometrys_instance[_reference]; ok {
			reference.CylinderGeometry = _instance
		}
	}
	if _reference := reference.BoxGeometry; _reference != nil {
		reference.BoxGeometry = nil
		if _instance, ok := stage.BoxGeometrys_instance[_reference]; ok {
			reference.BoxGeometry = _instance
		}
	}
	if _reference := reference.SphereGeometry; _reference != nil {
		reference.SphereGeometry = nil
		if _instance, ok := stage.SphereGeometrys_instance[_reference]; ok {
			reference.SphereGeometry = _instance
		}
	}
	if _reference := reference.TorusGeometry; _reference != nil {
		reference.TorusGeometry = nil
		if _instance, ok := stage.TorusGeometrys_instance[_reference]; ok {
			reference.TorusGeometry = _instance
		}
	}
	if _reference := reference.PlaneGeometry; _reference != nil {
		reference.PlaneGeometry = nil
		if _instance, ok := stage.PlaneGeometrys_instance[_reference]; ok {
			reference.PlaneGeometry = _instance
		}
	}
	if _reference := reference.TubeGeometry; _reference != nil {
		reference.TubeGeometry = nil
		if _instance, ok := stage.TubeGeometrys_instance[_reference]; ok {
			reference.TubeGeometry = _instance
		}
	}
	if _reference := reference.ExtrudeGeometry; _reference != nil {
		reference.ExtrudeGeometry = nil
		if _instance, ok := stage.ExtrudeGeometrys_instance[_reference]; ok {
			reference.ExtrudeGeometry = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *MeshMaterialBasic) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MeshPhysicalMaterial) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PlaneGeometry) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Shape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Points []*Vector2
	for _, _reference := range reference.Points {
		if _instance, ok := stage.Vector2s_instance[_reference]; ok {
			_Points = append(_Points, _instance)
		}
	}
	reference.Points = _Points
}

func (reference *SphereGeometry) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TorusGeometry) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TubeGeometry) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Path; _reference != nil {
		reference.Path = nil
		if _instance, ok := stage.Curves_instance[_reference]; ok {
			reference.Path = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Vector2) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Vector3) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ambiantlight *AmbiantLight) GongDiff(stage *Stage, ambiantlightOther *AmbiantLight) (diffs []string) {
	// insertion point for field diffs
	if ambiantlight.Name != ambiantlightOther.Name {
		diffs = append(diffs, ambiantlight.GongMarshallField(stage, "Name"))
	}
	if ambiantlight.Intensity != ambiantlightOther.Intensity {
		diffs = append(diffs, ambiantlight.GongMarshallField(stage, "Intensity"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (boxgeometry *BoxGeometry) GongDiff(stage *Stage, boxgeometryOther *BoxGeometry) (diffs []string) {
	// insertion point for field diffs
	if boxgeometry.Name != boxgeometryOther.Name {
		diffs = append(diffs, boxgeometry.GongMarshallField(stage, "Name"))
	}
	if boxgeometry.Width != boxgeometryOther.Width {
		diffs = append(diffs, boxgeometry.GongMarshallField(stage, "Width"))
	}
	if boxgeometry.Height != boxgeometryOther.Height {
		diffs = append(diffs, boxgeometry.GongMarshallField(stage, "Height"))
	}
	if boxgeometry.Depth != boxgeometryOther.Depth {
		diffs = append(diffs, boxgeometry.GongMarshallField(stage, "Depth"))
	}
	if boxgeometry.WidthSegments != boxgeometryOther.WidthSegments {
		diffs = append(diffs, boxgeometry.GongMarshallField(stage, "WidthSegments"))
	}
	if boxgeometry.HeightSegments != boxgeometryOther.HeightSegments {
		diffs = append(diffs, boxgeometry.GongMarshallField(stage, "HeightSegments"))
	}
	if boxgeometry.DepthSegments != boxgeometryOther.DepthSegments {
		diffs = append(diffs, boxgeometry.GongMarshallField(stage, "DepthSegments"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (canvas *Canvas) GongDiff(stage *Stage, canvasOther *Canvas) (diffs []string) {
	// insertion point for field diffs
	if canvas.Name != canvasOther.Name {
		diffs = append(diffs, canvas.GongMarshallField(stage, "Name"))
	}
	DirectionalLightsDifferent := false
	if len(canvas.DirectionalLights) != len(canvasOther.DirectionalLights) {
		DirectionalLightsDifferent = true
	} else {
		for i := range canvas.DirectionalLights {
			if (canvas.DirectionalLights[i] == nil) != (canvasOther.DirectionalLights[i] == nil) {
				DirectionalLightsDifferent = true
				break
			} else if canvas.DirectionalLights[i] != nil && canvasOther.DirectionalLights[i] != nil {
				// this is a pointer comparaison
				if canvas.DirectionalLights[i] != canvasOther.DirectionalLights[i] {
					DirectionalLightsDifferent = true
					break
				}
			}
		}
	}
	if DirectionalLightsDifferent {
		ops := Diff(stage, canvas, canvasOther, "DirectionalLights", canvasOther.DirectionalLights, canvas.DirectionalLights)
		diffs = append(diffs, ops)
	}
	if (canvas.AmbiantLight == nil) != (canvasOther.AmbiantLight == nil) {
		diffs = append(diffs, canvas.GongMarshallField(stage, "AmbiantLight"))
	} else if canvas.AmbiantLight != nil && canvasOther.AmbiantLight != nil {
		if canvas.AmbiantLight != canvasOther.AmbiantLight {
			diffs = append(diffs, canvas.GongMarshallField(stage, "AmbiantLight"))
		}
	}
	MeshsDifferent := false
	if len(canvas.Meshs) != len(canvasOther.Meshs) {
		MeshsDifferent = true
	} else {
		for i := range canvas.Meshs {
			if (canvas.Meshs[i] == nil) != (canvasOther.Meshs[i] == nil) {
				MeshsDifferent = true
				break
			} else if canvas.Meshs[i] != nil && canvasOther.Meshs[i] != nil {
				// this is a pointer comparaison
				if canvas.Meshs[i] != canvasOther.Meshs[i] {
					MeshsDifferent = true
					break
				}
			}
		}
	}
	if MeshsDifferent {
		ops := Diff(stage, canvas, canvasOther, "Meshs", canvasOther.Meshs, canvas.Meshs)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (curve *Curve) GongDiff(stage *Stage, curveOther *Curve) (diffs []string) {
	// insertion point for field diffs
	if curve.Name != curveOther.Name {
		diffs = append(diffs, curve.GongMarshallField(stage, "Name"))
	}
	PointsDifferent := false
	if len(curve.Points) != len(curveOther.Points) {
		PointsDifferent = true
	} else {
		for i := range curve.Points {
			if (curve.Points[i] == nil) != (curveOther.Points[i] == nil) {
				PointsDifferent = true
				break
			} else if curve.Points[i] != nil && curveOther.Points[i] != nil {
				// this is a pointer comparaison
				if curve.Points[i] != curveOther.Points[i] {
					PointsDifferent = true
					break
				}
			}
		}
	}
	if PointsDifferent {
		ops := Diff(stage, curve, curveOther, "Points", curveOther.Points, curve.Points)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (cylindergeometry *CylinderGeometry) GongDiff(stage *Stage, cylindergeometryOther *CylinderGeometry) (diffs []string) {
	// insertion point for field diffs
	if cylindergeometry.Name != cylindergeometryOther.Name {
		diffs = append(diffs, cylindergeometry.GongMarshallField(stage, "Name"))
	}
	if cylindergeometry.RadiusTop != cylindergeometryOther.RadiusTop {
		diffs = append(diffs, cylindergeometry.GongMarshallField(stage, "RadiusTop"))
	}
	if cylindergeometry.RadiusBottom != cylindergeometryOther.RadiusBottom {
		diffs = append(diffs, cylindergeometry.GongMarshallField(stage, "RadiusBottom"))
	}
	if cylindergeometry.Height != cylindergeometryOther.Height {
		diffs = append(diffs, cylindergeometry.GongMarshallField(stage, "Height"))
	}
	if cylindergeometry.RadialSegments != cylindergeometryOther.RadialSegments {
		diffs = append(diffs, cylindergeometry.GongMarshallField(stage, "RadialSegments"))
	}
	if cylindergeometry.HeightSegments != cylindergeometryOther.HeightSegments {
		diffs = append(diffs, cylindergeometry.GongMarshallField(stage, "HeightSegments"))
	}
	if cylindergeometry.OpenEnded != cylindergeometryOther.OpenEnded {
		diffs = append(diffs, cylindergeometry.GongMarshallField(stage, "OpenEnded"))
	}
	if cylindergeometry.ThetaStart != cylindergeometryOther.ThetaStart {
		diffs = append(diffs, cylindergeometry.GongMarshallField(stage, "ThetaStart"))
	}
	if cylindergeometry.ThetaLength != cylindergeometryOther.ThetaLength {
		diffs = append(diffs, cylindergeometry.GongMarshallField(stage, "ThetaLength"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (directionallight *DirectionalLight) GongDiff(stage *Stage, directionallightOther *DirectionalLight) (diffs []string) {
	// insertion point for field diffs
	if directionallight.Name != directionallightOther.Name {
		diffs = append(diffs, directionallight.GongMarshallField(stage, "Name"))
	}
	if directionallight.X != directionallightOther.X {
		diffs = append(diffs, directionallight.GongMarshallField(stage, "X"))
	}
	if directionallight.Y != directionallightOther.Y {
		diffs = append(diffs, directionallight.GongMarshallField(stage, "Y"))
	}
	if directionallight.Z != directionallightOther.Z {
		diffs = append(diffs, directionallight.GongMarshallField(stage, "Z"))
	}
	if directionallight.Intensity != directionallightOther.Intensity {
		diffs = append(diffs, directionallight.GongMarshallField(stage, "Intensity"))
	}
	if directionallight.IsWithCastShadow != directionallightOther.IsWithCastShadow {
		diffs = append(diffs, directionallight.GongMarshallField(stage, "IsWithCastShadow"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (extrudegeometry *ExtrudeGeometry) GongDiff(stage *Stage, extrudegeometryOther *ExtrudeGeometry) (diffs []string) {
	// insertion point for field diffs
	if extrudegeometry.Name != extrudegeometryOther.Name {
		diffs = append(diffs, extrudegeometry.GongMarshallField(stage, "Name"))
	}
	if (extrudegeometry.Shape == nil) != (extrudegeometryOther.Shape == nil) {
		diffs = append(diffs, extrudegeometry.GongMarshallField(stage, "Shape"))
	} else if extrudegeometry.Shape != nil && extrudegeometryOther.Shape != nil {
		if extrudegeometry.Shape != extrudegeometryOther.Shape {
			diffs = append(diffs, extrudegeometry.GongMarshallField(stage, "Shape"))
		}
	}
	if (extrudegeometry.ExtrudePath == nil) != (extrudegeometryOther.ExtrudePath == nil) {
		diffs = append(diffs, extrudegeometry.GongMarshallField(stage, "ExtrudePath"))
	} else if extrudegeometry.ExtrudePath != nil && extrudegeometryOther.ExtrudePath != nil {
		if extrudegeometry.ExtrudePath != extrudegeometryOther.ExtrudePath {
			diffs = append(diffs, extrudegeometry.GongMarshallField(stage, "ExtrudePath"))
		}
	}
	if extrudegeometry.Steps != extrudegeometryOther.Steps {
		diffs = append(diffs, extrudegeometry.GongMarshallField(stage, "Steps"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (mesh *Mesh) GongDiff(stage *Stage, meshOther *Mesh) (diffs []string) {
	// insertion point for field diffs
	if mesh.Name != meshOther.Name {
		diffs = append(diffs, mesh.GongMarshallField(stage, "Name"))
	}
	if mesh.X != meshOther.X {
		diffs = append(diffs, mesh.GongMarshallField(stage, "X"))
	}
	if mesh.Y != meshOther.Y {
		diffs = append(diffs, mesh.GongMarshallField(stage, "Y"))
	}
	if mesh.Z != meshOther.Z {
		diffs = append(diffs, mesh.GongMarshallField(stage, "Z"))
	}
	if (mesh.MeshMaterialBasic == nil) != (meshOther.MeshMaterialBasic == nil) {
		diffs = append(diffs, mesh.GongMarshallField(stage, "MeshMaterialBasic"))
	} else if mesh.MeshMaterialBasic != nil && meshOther.MeshMaterialBasic != nil {
		if mesh.MeshMaterialBasic != meshOther.MeshMaterialBasic {
			diffs = append(diffs, mesh.GongMarshallField(stage, "MeshMaterialBasic"))
		}
	}
	if (mesh.MeshPhysicalMaterial == nil) != (meshOther.MeshPhysicalMaterial == nil) {
		diffs = append(diffs, mesh.GongMarshallField(stage, "MeshPhysicalMaterial"))
	} else if mesh.MeshPhysicalMaterial != nil && meshOther.MeshPhysicalMaterial != nil {
		if mesh.MeshPhysicalMaterial != meshOther.MeshPhysicalMaterial {
			diffs = append(diffs, mesh.GongMarshallField(stage, "MeshPhysicalMaterial"))
		}
	}
	if (mesh.CylinderGeometry == nil) != (meshOther.CylinderGeometry == nil) {
		diffs = append(diffs, mesh.GongMarshallField(stage, "CylinderGeometry"))
	} else if mesh.CylinderGeometry != nil && meshOther.CylinderGeometry != nil {
		if mesh.CylinderGeometry != meshOther.CylinderGeometry {
			diffs = append(diffs, mesh.GongMarshallField(stage, "CylinderGeometry"))
		}
	}
	if (mesh.BoxGeometry == nil) != (meshOther.BoxGeometry == nil) {
		diffs = append(diffs, mesh.GongMarshallField(stage, "BoxGeometry"))
	} else if mesh.BoxGeometry != nil && meshOther.BoxGeometry != nil {
		if mesh.BoxGeometry != meshOther.BoxGeometry {
			diffs = append(diffs, mesh.GongMarshallField(stage, "BoxGeometry"))
		}
	}
	if (mesh.SphereGeometry == nil) != (meshOther.SphereGeometry == nil) {
		diffs = append(diffs, mesh.GongMarshallField(stage, "SphereGeometry"))
	} else if mesh.SphereGeometry != nil && meshOther.SphereGeometry != nil {
		if mesh.SphereGeometry != meshOther.SphereGeometry {
			diffs = append(diffs, mesh.GongMarshallField(stage, "SphereGeometry"))
		}
	}
	if (mesh.TorusGeometry == nil) != (meshOther.TorusGeometry == nil) {
		diffs = append(diffs, mesh.GongMarshallField(stage, "TorusGeometry"))
	} else if mesh.TorusGeometry != nil && meshOther.TorusGeometry != nil {
		if mesh.TorusGeometry != meshOther.TorusGeometry {
			diffs = append(diffs, mesh.GongMarshallField(stage, "TorusGeometry"))
		}
	}
	if (mesh.PlaneGeometry == nil) != (meshOther.PlaneGeometry == nil) {
		diffs = append(diffs, mesh.GongMarshallField(stage, "PlaneGeometry"))
	} else if mesh.PlaneGeometry != nil && meshOther.PlaneGeometry != nil {
		if mesh.PlaneGeometry != meshOther.PlaneGeometry {
			diffs = append(diffs, mesh.GongMarshallField(stage, "PlaneGeometry"))
		}
	}
	if (mesh.TubeGeometry == nil) != (meshOther.TubeGeometry == nil) {
		diffs = append(diffs, mesh.GongMarshallField(stage, "TubeGeometry"))
	} else if mesh.TubeGeometry != nil && meshOther.TubeGeometry != nil {
		if mesh.TubeGeometry != meshOther.TubeGeometry {
			diffs = append(diffs, mesh.GongMarshallField(stage, "TubeGeometry"))
		}
	}
	if (mesh.ExtrudeGeometry == nil) != (meshOther.ExtrudeGeometry == nil) {
		diffs = append(diffs, mesh.GongMarshallField(stage, "ExtrudeGeometry"))
	} else if mesh.ExtrudeGeometry != nil && meshOther.ExtrudeGeometry != nil {
		if mesh.ExtrudeGeometry != meshOther.ExtrudeGeometry {
			diffs = append(diffs, mesh.GongMarshallField(stage, "ExtrudeGeometry"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (meshmaterialbasic *MeshMaterialBasic) GongDiff(stage *Stage, meshmaterialbasicOther *MeshMaterialBasic) (diffs []string) {
	// insertion point for field diffs
	if meshmaterialbasic.Name != meshmaterialbasicOther.Name {
		diffs = append(diffs, meshmaterialbasic.GongMarshallField(stage, "Name"))
	}
	if meshmaterialbasic.Color != meshmaterialbasicOther.Color {
		diffs = append(diffs, meshmaterialbasic.GongMarshallField(stage, "Color"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (meshphysicalmaterial *MeshPhysicalMaterial) GongDiff(stage *Stage, meshphysicalmaterialOther *MeshPhysicalMaterial) (diffs []string) {
	// insertion point for field diffs
	if meshphysicalmaterial.Name != meshphysicalmaterialOther.Name {
		diffs = append(diffs, meshphysicalmaterial.GongMarshallField(stage, "Name"))
	}
	if meshphysicalmaterial.Color != meshphysicalmaterialOther.Color {
		diffs = append(diffs, meshphysicalmaterial.GongMarshallField(stage, "Color"))
	}
	if meshphysicalmaterial.Wireframe != meshphysicalmaterialOther.Wireframe {
		diffs = append(diffs, meshphysicalmaterial.GongMarshallField(stage, "Wireframe"))
	}
	if meshphysicalmaterial.Opacity != meshphysicalmaterialOther.Opacity {
		diffs = append(diffs, meshphysicalmaterial.GongMarshallField(stage, "Opacity"))
	}
	if meshphysicalmaterial.Transparent != meshphysicalmaterialOther.Transparent {
		diffs = append(diffs, meshphysicalmaterial.GongMarshallField(stage, "Transparent"))
	}
	if meshphysicalmaterial.Visible != meshphysicalmaterialOther.Visible {
		diffs = append(diffs, meshphysicalmaterial.GongMarshallField(stage, "Visible"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (planegeometry *PlaneGeometry) GongDiff(stage *Stage, planegeometryOther *PlaneGeometry) (diffs []string) {
	// insertion point for field diffs
	if planegeometry.Name != planegeometryOther.Name {
		diffs = append(diffs, planegeometry.GongMarshallField(stage, "Name"))
	}
	if planegeometry.Width != planegeometryOther.Width {
		diffs = append(diffs, planegeometry.GongMarshallField(stage, "Width"))
	}
	if planegeometry.Height != planegeometryOther.Height {
		diffs = append(diffs, planegeometry.GongMarshallField(stage, "Height"))
	}
	if planegeometry.WidthSegments != planegeometryOther.WidthSegments {
		diffs = append(diffs, planegeometry.GongMarshallField(stage, "WidthSegments"))
	}
	if planegeometry.HeightSegments != planegeometryOther.HeightSegments {
		diffs = append(diffs, planegeometry.GongMarshallField(stage, "HeightSegments"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (shape *Shape) GongDiff(stage *Stage, shapeOther *Shape) (diffs []string) {
	// insertion point for field diffs
	if shape.Name != shapeOther.Name {
		diffs = append(diffs, shape.GongMarshallField(stage, "Name"))
	}
	PointsDifferent := false
	if len(shape.Points) != len(shapeOther.Points) {
		PointsDifferent = true
	} else {
		for i := range shape.Points {
			if (shape.Points[i] == nil) != (shapeOther.Points[i] == nil) {
				PointsDifferent = true
				break
			} else if shape.Points[i] != nil && shapeOther.Points[i] != nil {
				// this is a pointer comparaison
				if shape.Points[i] != shapeOther.Points[i] {
					PointsDifferent = true
					break
				}
			}
		}
	}
	if PointsDifferent {
		ops := Diff(stage, shape, shapeOther, "Points", shapeOther.Points, shape.Points)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (spheregeometry *SphereGeometry) GongDiff(stage *Stage, spheregeometryOther *SphereGeometry) (diffs []string) {
	// insertion point for field diffs
	if spheregeometry.Name != spheregeometryOther.Name {
		diffs = append(diffs, spheregeometry.GongMarshallField(stage, "Name"))
	}
	if spheregeometry.Radius != spheregeometryOther.Radius {
		diffs = append(diffs, spheregeometry.GongMarshallField(stage, "Radius"))
	}
	if spheregeometry.WidthSegments != spheregeometryOther.WidthSegments {
		diffs = append(diffs, spheregeometry.GongMarshallField(stage, "WidthSegments"))
	}
	if spheregeometry.HeightSegments != spheregeometryOther.HeightSegments {
		diffs = append(diffs, spheregeometry.GongMarshallField(stage, "HeightSegments"))
	}
	if spheregeometry.PhiStart != spheregeometryOther.PhiStart {
		diffs = append(diffs, spheregeometry.GongMarshallField(stage, "PhiStart"))
	}
	if spheregeometry.PhiLength != spheregeometryOther.PhiLength {
		diffs = append(diffs, spheregeometry.GongMarshallField(stage, "PhiLength"))
	}
	if spheregeometry.ThetaStart != spheregeometryOther.ThetaStart {
		diffs = append(diffs, spheregeometry.GongMarshallField(stage, "ThetaStart"))
	}
	if spheregeometry.ThetaLength != spheregeometryOther.ThetaLength {
		diffs = append(diffs, spheregeometry.GongMarshallField(stage, "ThetaLength"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (torusgeometry *TorusGeometry) GongDiff(stage *Stage, torusgeometryOther *TorusGeometry) (diffs []string) {
	// insertion point for field diffs
	if torusgeometry.Name != torusgeometryOther.Name {
		diffs = append(diffs, torusgeometry.GongMarshallField(stage, "Name"))
	}
	if torusgeometry.Radius != torusgeometryOther.Radius {
		diffs = append(diffs, torusgeometry.GongMarshallField(stage, "Radius"))
	}
	if torusgeometry.Tube != torusgeometryOther.Tube {
		diffs = append(diffs, torusgeometry.GongMarshallField(stage, "Tube"))
	}
	if torusgeometry.RadialSegments != torusgeometryOther.RadialSegments {
		diffs = append(diffs, torusgeometry.GongMarshallField(stage, "RadialSegments"))
	}
	if torusgeometry.TubularSegments != torusgeometryOther.TubularSegments {
		diffs = append(diffs, torusgeometry.GongMarshallField(stage, "TubularSegments"))
	}
	if torusgeometry.Arc != torusgeometryOther.Arc {
		diffs = append(diffs, torusgeometry.GongMarshallField(stage, "Arc"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tubegeometry *TubeGeometry) GongDiff(stage *Stage, tubegeometryOther *TubeGeometry) (diffs []string) {
	// insertion point for field diffs
	if tubegeometry.Name != tubegeometryOther.Name {
		diffs = append(diffs, tubegeometry.GongMarshallField(stage, "Name"))
	}
	if (tubegeometry.Path == nil) != (tubegeometryOther.Path == nil) {
		diffs = append(diffs, tubegeometry.GongMarshallField(stage, "Path"))
	} else if tubegeometry.Path != nil && tubegeometryOther.Path != nil {
		if tubegeometry.Path != tubegeometryOther.Path {
			diffs = append(diffs, tubegeometry.GongMarshallField(stage, "Path"))
		}
	}
	if tubegeometry.TubularSegments != tubegeometryOther.TubularSegments {
		diffs = append(diffs, tubegeometry.GongMarshallField(stage, "TubularSegments"))
	}
	if tubegeometry.Radius != tubegeometryOther.Radius {
		diffs = append(diffs, tubegeometry.GongMarshallField(stage, "Radius"))
	}
	if tubegeometry.RadialSegments != tubegeometryOther.RadialSegments {
		diffs = append(diffs, tubegeometry.GongMarshallField(stage, "RadialSegments"))
	}
	if tubegeometry.Closed != tubegeometryOther.Closed {
		diffs = append(diffs, tubegeometry.GongMarshallField(stage, "Closed"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (vector2 *Vector2) GongDiff(stage *Stage, vector2Other *Vector2) (diffs []string) {
	// insertion point for field diffs
	if vector2.Name != vector2Other.Name {
		diffs = append(diffs, vector2.GongMarshallField(stage, "Name"))
	}
	if vector2.X != vector2Other.X {
		diffs = append(diffs, vector2.GongMarshallField(stage, "X"))
	}
	if vector2.Y != vector2Other.Y {
		diffs = append(diffs, vector2.GongMarshallField(stage, "Y"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (vector3 *Vector3) GongDiff(stage *Stage, vector3Other *Vector3) (diffs []string) {
	// insertion point for field diffs
	if vector3.Name != vector3Other.Name {
		diffs = append(diffs, vector3.GongMarshallField(stage, "Name"))
	}
	if vector3.X != vector3Other.X {
		diffs = append(diffs, vector3.GongMarshallField(stage, "X"))
	}
	if vector3.Y != vector3Other.Y {
		diffs = append(diffs, vector3.GongMarshallField(stage, "Y"))
	}
	if vector3.Z != vector3Other.Z {
		diffs = append(diffs, vector3.GongMarshallField(stage, "Z"))
	}

	return
}

// Diff returns the sequence of operations to transform oldSlice into newSlice.
// It requires type T to be comparable (e.g., pointers, ints, strings).
func Diff[T1, T2 PointerToGongstruct](stage *Stage, a, b T1, fieldName string, oldSlice, newSlice []T2) (ops string) {
	m, n := len(oldSlice), len(newSlice)

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if oldSlice[i] == newSlice[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if oldSlice[i-1] == newSlice[j-1] {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Create a temporary view of what's left after deletions for tracking matches
	var currentLCS []T2
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			currentLCS = append(currentLCS, oldSlice[k])
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k, targetVal := range newSlice {
		if lcsIdx < len(currentLCS) && currentLCS[lcsIdx] == targetVal {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, targetVal.GongGetIdentifier(stage))
		}
	}

	return ops
}
