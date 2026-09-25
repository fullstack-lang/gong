// generated code - do not edit
package models

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (ambiantlight *AmbiantLight) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AmbiantLights[ambiantlight]
	return ok
}

func (boxgeometry *BoxGeometry) GongIsStaged(stage *Stage) bool {
	_, ok := stage.BoxGeometrys[boxgeometry]
	return ok
}

func (buffergeometry *BufferGeometry) GongIsStaged(stage *Stage) bool {
	_, ok := stage.BufferGeometrys[buffergeometry]
	return ok
}

func (camera *Camera) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Cameras[camera]
	return ok
}

func (canvas *Canvas) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Canvass[canvas]
	return ok
}

func (curve *Curve) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Curves[curve]
	return ok
}

func (cylindergeometry *CylinderGeometry) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CylinderGeometrys[cylindergeometry]
	return ok
}

func (directionallight *DirectionalLight) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DirectionalLights[directionallight]
	return ok
}

func (extrudegeometry *ExtrudeGeometry) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ExtrudeGeometrys[extrudegeometry]
	return ok
}

func (mesh *Mesh) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Meshs[mesh]
	return ok
}

func (meshmaterialbasic *MeshMaterialBasic) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MeshMaterialBasics[meshmaterialbasic]
	return ok
}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MeshPhysicalMaterials[meshphysicalmaterial]
	return ok
}

func (planegeometry *PlaneGeometry) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PlaneGeometrys[planegeometry]
	return ok
}

func (shape *Shape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Shapes[shape]
	return ok
}

func (spheregeometry *SphereGeometry) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SphereGeometrys[spheregeometry]
	return ok
}

func (torusgeometry *TorusGeometry) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TorusGeometrys[torusgeometry]
	return ok
}

func (triangle *Triangle) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Triangles[triangle]
	return ok
}

func (tubegeometry *TubeGeometry) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TubeGeometrys[tubegeometry]
	return ok
}

func (vector2 *Vector2) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Vector2s[vector2]
	return ok
}

func (vector3 *Vector3) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Vector3s[vector3]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (ambiantlight *AmbiantLight) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ambiantlight) {
		return
	}

	ambiantlight.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (boxgeometry *BoxGeometry) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(boxgeometry) {
		return
	}

	boxgeometry.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (buffergeometry *BufferGeometry) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(buffergeometry) {
		return
	}

	buffergeometry.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _vector3 := range buffergeometry.Vertices {
		stage.StageBranch(_vector3)
	}
	for _, _triangle := range buffergeometry.Faces {
		stage.StageBranch(_triangle)
	}

}

func (camera *Camera) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(camera) {
		return
	}

	camera.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (canvas *Canvas) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(canvas) {
		return
	}

	canvas.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if canvas.AmbiantLight != nil {
		stage.StageBranch(canvas.AmbiantLight)
	}
	if canvas.Camera != nil {
		stage.StageBranch(canvas.Camera)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _directionallight := range canvas.DirectionalLights {
		stage.StageBranch(_directionallight)
	}
	for _, _mesh := range canvas.Meshs {
		stage.StageBranch(_mesh)
	}

}

func (curve *Curve) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(curve) {
		return
	}

	curve.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _vector3 := range curve.Points {
		stage.StageBranch(_vector3)
	}

}

func (cylindergeometry *CylinderGeometry) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(cylindergeometry) {
		return
	}

	cylindergeometry.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (directionallight *DirectionalLight) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(directionallight) {
		return
	}

	directionallight.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (extrudegeometry *ExtrudeGeometry) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(extrudegeometry) {
		return
	}

	extrudegeometry.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if extrudegeometry.Shape != nil {
		stage.StageBranch(extrudegeometry.Shape)
	}
	if extrudegeometry.ExtrudePath != nil {
		stage.StageBranch(extrudegeometry.ExtrudePath)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (mesh *Mesh) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(mesh) {
		return
	}

	mesh.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if mesh.MeshMaterialBasic != nil {
		stage.StageBranch(mesh.MeshMaterialBasic)
	}
	if mesh.MeshPhysicalMaterial != nil {
		stage.StageBranch(mesh.MeshPhysicalMaterial)
	}
	if mesh.CylinderGeometry != nil {
		stage.StageBranch(mesh.CylinderGeometry)
	}
	if mesh.BoxGeometry != nil {
		stage.StageBranch(mesh.BoxGeometry)
	}
	if mesh.SphereGeometry != nil {
		stage.StageBranch(mesh.SphereGeometry)
	}
	if mesh.TorusGeometry != nil {
		stage.StageBranch(mesh.TorusGeometry)
	}
	if mesh.PlaneGeometry != nil {
		stage.StageBranch(mesh.PlaneGeometry)
	}
	if mesh.TubeGeometry != nil {
		stage.StageBranch(mesh.TubeGeometry)
	}
	if mesh.ExtrudeGeometry != nil {
		stage.StageBranch(mesh.ExtrudeGeometry)
	}
	if mesh.BufferGeometry != nil {
		stage.StageBranch(mesh.BufferGeometry)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (meshmaterialbasic *MeshMaterialBasic) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(meshmaterialbasic) {
		return
	}

	meshmaterialbasic.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(meshphysicalmaterial) {
		return
	}

	meshphysicalmaterial.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (planegeometry *PlaneGeometry) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(planegeometry) {
		return
	}

	planegeometry.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shape *Shape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(shape) {
		return
	}

	shape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _vector2 := range shape.Points {
		stage.StageBranch(_vector2)
	}

}

func (spheregeometry *SphereGeometry) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(spheregeometry) {
		return
	}

	spheregeometry.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (torusgeometry *TorusGeometry) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(torusgeometry) {
		return
	}

	torusgeometry.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (triangle *Triangle) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(triangle) {
		return
	}

	triangle.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tubegeometry *TubeGeometry) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tubegeometry) {
		return
	}

	tubegeometry.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tubegeometry.Path != nil {
		stage.StageBranch(tubegeometry.Path)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (vector2 *Vector2) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(vector2) {
		return
	}

	vector2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (vector3 *Vector3) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(vector3) {
		return
	}

	vector3.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *AmbiantLight:
		toT := GongCopyBranchAmbiantLight(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BoxGeometry:
		toT := GongCopyBranchBoxGeometry(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BufferGeometry:
		toT := GongCopyBranchBufferGeometry(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Camera:
		toT := GongCopyBranchCamera(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Canvas:
		toT := GongCopyBranchCanvas(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Curve:
		toT := GongCopyBranchCurve(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CylinderGeometry:
		toT := GongCopyBranchCylinderGeometry(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DirectionalLight:
		toT := GongCopyBranchDirectionalLight(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ExtrudeGeometry:
		toT := GongCopyBranchExtrudeGeometry(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Mesh:
		toT := GongCopyBranchMesh(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MeshMaterialBasic:
		toT := GongCopyBranchMeshMaterialBasic(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MeshPhysicalMaterial:
		toT := GongCopyBranchMeshPhysicalMaterial(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PlaneGeometry:
		toT := GongCopyBranchPlaneGeometry(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Shape:
		toT := GongCopyBranchShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SphereGeometry:
		toT := GongCopyBranchSphereGeometry(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TorusGeometry:
		toT := GongCopyBranchTorusGeometry(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Triangle:
		toT := GongCopyBranchTriangle(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TubeGeometry:
		toT := GongCopyBranchTubeGeometry(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Vector2:
		toT := GongCopyBranchVector2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Vector3:
		toT := GongCopyBranchVector3(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAmbiantLight(mapOrigCopy map[any]any, ambiantlightFrom *AmbiantLight) (ambiantlightTo *AmbiantLight) {
	var alreadyCopied bool
	ambiantlightTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, ambiantlightFrom)
	if alreadyCopied {
		return
	}
	ambiantlightFrom.GongCopyBasicFields(ambiantlightTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBoxGeometry(mapOrigCopy map[any]any, boxgeometryFrom *BoxGeometry) (boxgeometryTo *BoxGeometry) {
	var alreadyCopied bool
	boxgeometryTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, boxgeometryFrom)
	if alreadyCopied {
		return
	}
	boxgeometryFrom.GongCopyBasicFields(boxgeometryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBufferGeometry(mapOrigCopy map[any]any, buffergeometryFrom *BufferGeometry) (buffergeometryTo *BufferGeometry) {
	var alreadyCopied bool
	buffergeometryTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, buffergeometryFrom)
	if alreadyCopied {
		return
	}
	buffergeometryFrom.GongCopyBasicFields(buffergeometryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _vector3 := range buffergeometryFrom.Vertices {
		buffergeometryTo.Vertices = append(buffergeometryTo.Vertices, GongCopyBranchVector3(mapOrigCopy, _vector3))
	}
	for _, _triangle := range buffergeometryFrom.Faces {
		buffergeometryTo.Faces = append(buffergeometryTo.Faces, GongCopyBranchTriangle(mapOrigCopy, _triangle))
	}

	return
}

func GongCopyBranchCamera(mapOrigCopy map[any]any, cameraFrom *Camera) (cameraTo *Camera) {
	var alreadyCopied bool
	cameraTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, cameraFrom)
	if alreadyCopied {
		return
	}
	cameraFrom.GongCopyBasicFields(cameraTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCanvas(mapOrigCopy map[any]any, canvasFrom *Canvas) (canvasTo *Canvas) {
	var alreadyCopied bool
	canvasTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, canvasFrom)
	if alreadyCopied {
		return
	}
	canvasFrom.GongCopyBasicFields(canvasTo)

	//insertion point for the staging of instances referenced by pointers
	if canvasFrom.AmbiantLight != nil {
		canvasTo.AmbiantLight = GongCopyBranchAmbiantLight(mapOrigCopy, canvasFrom.AmbiantLight)
	}
	if canvasFrom.Camera != nil {
		canvasTo.Camera = GongCopyBranchCamera(mapOrigCopy, canvasFrom.Camera)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _directionallight := range canvasFrom.DirectionalLights {
		canvasTo.DirectionalLights = append(canvasTo.DirectionalLights, GongCopyBranchDirectionalLight(mapOrigCopy, _directionallight))
	}
	for _, _mesh := range canvasFrom.Meshs {
		canvasTo.Meshs = append(canvasTo.Meshs, GongCopyBranchMesh(mapOrigCopy, _mesh))
	}

	return
}

func GongCopyBranchCurve(mapOrigCopy map[any]any, curveFrom *Curve) (curveTo *Curve) {
	var alreadyCopied bool
	curveTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, curveFrom)
	if alreadyCopied {
		return
	}
	curveFrom.GongCopyBasicFields(curveTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _vector3 := range curveFrom.Points {
		curveTo.Points = append(curveTo.Points, GongCopyBranchVector3(mapOrigCopy, _vector3))
	}

	return
}

func GongCopyBranchCylinderGeometry(mapOrigCopy map[any]any, cylindergeometryFrom *CylinderGeometry) (cylindergeometryTo *CylinderGeometry) {
	var alreadyCopied bool
	cylindergeometryTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, cylindergeometryFrom)
	if alreadyCopied {
		return
	}
	cylindergeometryFrom.GongCopyBasicFields(cylindergeometryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDirectionalLight(mapOrigCopy map[any]any, directionallightFrom *DirectionalLight) (directionallightTo *DirectionalLight) {
	var alreadyCopied bool
	directionallightTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, directionallightFrom)
	if alreadyCopied {
		return
	}
	directionallightFrom.GongCopyBasicFields(directionallightTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchExtrudeGeometry(mapOrigCopy map[any]any, extrudegeometryFrom *ExtrudeGeometry) (extrudegeometryTo *ExtrudeGeometry) {
	var alreadyCopied bool
	extrudegeometryTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, extrudegeometryFrom)
	if alreadyCopied {
		return
	}
	extrudegeometryFrom.GongCopyBasicFields(extrudegeometryTo)

	//insertion point for the staging of instances referenced by pointers
	if extrudegeometryFrom.Shape != nil {
		extrudegeometryTo.Shape = GongCopyBranchShape(mapOrigCopy, extrudegeometryFrom.Shape)
	}
	if extrudegeometryFrom.ExtrudePath != nil {
		extrudegeometryTo.ExtrudePath = GongCopyBranchCurve(mapOrigCopy, extrudegeometryFrom.ExtrudePath)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMesh(mapOrigCopy map[any]any, meshFrom *Mesh) (meshTo *Mesh) {
	var alreadyCopied bool
	meshTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, meshFrom)
	if alreadyCopied {
		return
	}
	meshFrom.GongCopyBasicFields(meshTo)

	//insertion point for the staging of instances referenced by pointers
	if meshFrom.MeshMaterialBasic != nil {
		meshTo.MeshMaterialBasic = GongCopyBranchMeshMaterialBasic(mapOrigCopy, meshFrom.MeshMaterialBasic)
	}
	if meshFrom.MeshPhysicalMaterial != nil {
		meshTo.MeshPhysicalMaterial = GongCopyBranchMeshPhysicalMaterial(mapOrigCopy, meshFrom.MeshPhysicalMaterial)
	}
	if meshFrom.CylinderGeometry != nil {
		meshTo.CylinderGeometry = GongCopyBranchCylinderGeometry(mapOrigCopy, meshFrom.CylinderGeometry)
	}
	if meshFrom.BoxGeometry != nil {
		meshTo.BoxGeometry = GongCopyBranchBoxGeometry(mapOrigCopy, meshFrom.BoxGeometry)
	}
	if meshFrom.SphereGeometry != nil {
		meshTo.SphereGeometry = GongCopyBranchSphereGeometry(mapOrigCopy, meshFrom.SphereGeometry)
	}
	if meshFrom.TorusGeometry != nil {
		meshTo.TorusGeometry = GongCopyBranchTorusGeometry(mapOrigCopy, meshFrom.TorusGeometry)
	}
	if meshFrom.PlaneGeometry != nil {
		meshTo.PlaneGeometry = GongCopyBranchPlaneGeometry(mapOrigCopy, meshFrom.PlaneGeometry)
	}
	if meshFrom.TubeGeometry != nil {
		meshTo.TubeGeometry = GongCopyBranchTubeGeometry(mapOrigCopy, meshFrom.TubeGeometry)
	}
	if meshFrom.ExtrudeGeometry != nil {
		meshTo.ExtrudeGeometry = GongCopyBranchExtrudeGeometry(mapOrigCopy, meshFrom.ExtrudeGeometry)
	}
	if meshFrom.BufferGeometry != nil {
		meshTo.BufferGeometry = GongCopyBranchBufferGeometry(mapOrigCopy, meshFrom.BufferGeometry)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMeshMaterialBasic(mapOrigCopy map[any]any, meshmaterialbasicFrom *MeshMaterialBasic) (meshmaterialbasicTo *MeshMaterialBasic) {
	var alreadyCopied bool
	meshmaterialbasicTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, meshmaterialbasicFrom)
	if alreadyCopied {
		return
	}
	meshmaterialbasicFrom.GongCopyBasicFields(meshmaterialbasicTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMeshPhysicalMaterial(mapOrigCopy map[any]any, meshphysicalmaterialFrom *MeshPhysicalMaterial) (meshphysicalmaterialTo *MeshPhysicalMaterial) {
	var alreadyCopied bool
	meshphysicalmaterialTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, meshphysicalmaterialFrom)
	if alreadyCopied {
		return
	}
	meshphysicalmaterialFrom.GongCopyBasicFields(meshphysicalmaterialTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPlaneGeometry(mapOrigCopy map[any]any, planegeometryFrom *PlaneGeometry) (planegeometryTo *PlaneGeometry) {
	var alreadyCopied bool
	planegeometryTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, planegeometryFrom)
	if alreadyCopied {
		return
	}
	planegeometryFrom.GongCopyBasicFields(planegeometryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchShape(mapOrigCopy map[any]any, shapeFrom *Shape) (shapeTo *Shape) {
	var alreadyCopied bool
	shapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, shapeFrom)
	if alreadyCopied {
		return
	}
	shapeFrom.GongCopyBasicFields(shapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _vector2 := range shapeFrom.Points {
		shapeTo.Points = append(shapeTo.Points, GongCopyBranchVector2(mapOrigCopy, _vector2))
	}

	return
}

func GongCopyBranchSphereGeometry(mapOrigCopy map[any]any, spheregeometryFrom *SphereGeometry) (spheregeometryTo *SphereGeometry) {
	var alreadyCopied bool
	spheregeometryTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, spheregeometryFrom)
	if alreadyCopied {
		return
	}
	spheregeometryFrom.GongCopyBasicFields(spheregeometryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTorusGeometry(mapOrigCopy map[any]any, torusgeometryFrom *TorusGeometry) (torusgeometryTo *TorusGeometry) {
	var alreadyCopied bool
	torusgeometryTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, torusgeometryFrom)
	if alreadyCopied {
		return
	}
	torusgeometryFrom.GongCopyBasicFields(torusgeometryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTriangle(mapOrigCopy map[any]any, triangleFrom *Triangle) (triangleTo *Triangle) {
	var alreadyCopied bool
	triangleTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, triangleFrom)
	if alreadyCopied {
		return
	}
	triangleFrom.GongCopyBasicFields(triangleTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTubeGeometry(mapOrigCopy map[any]any, tubegeometryFrom *TubeGeometry) (tubegeometryTo *TubeGeometry) {
	var alreadyCopied bool
	tubegeometryTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tubegeometryFrom)
	if alreadyCopied {
		return
	}
	tubegeometryFrom.GongCopyBasicFields(tubegeometryTo)

	//insertion point for the staging of instances referenced by pointers
	if tubegeometryFrom.Path != nil {
		tubegeometryTo.Path = GongCopyBranchCurve(mapOrigCopy, tubegeometryFrom.Path)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchVector2(mapOrigCopy map[any]any, vector2From *Vector2) (vector2To *Vector2) {
	var alreadyCopied bool
	vector2To, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, vector2From)
	if alreadyCopied {
		return
	}
	vector2From.GongCopyBasicFields(vector2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchVector3(mapOrigCopy map[any]any, vector3From *Vector3) (vector3To *Vector3) {
	var alreadyCopied bool
	vector3To, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, vector3From)
	if alreadyCopied {
		return
	}
	vector3From.GongCopyBasicFields(vector3To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// insertion point for unstage branch per struct
func (ambiantlight *AmbiantLight) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ambiantlight) {
		return
	}

	ambiantlight.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (boxgeometry *BoxGeometry) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(boxgeometry) {
		return
	}

	boxgeometry.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (buffergeometry *BufferGeometry) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(buffergeometry) {
		return
	}

	buffergeometry.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _vector3 := range buffergeometry.Vertices {
		stage.UnstageBranch(_vector3)
	}
	for _, _triangle := range buffergeometry.Faces {
		stage.UnstageBranch(_triangle)
	}

}

func (camera *Camera) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(camera) {
		return
	}

	camera.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (canvas *Canvas) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(canvas) {
		return
	}

	canvas.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if canvas.AmbiantLight != nil {
		stage.UnstageBranch(canvas.AmbiantLight)
	}
	if canvas.Camera != nil {
		stage.UnstageBranch(canvas.Camera)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _directionallight := range canvas.DirectionalLights {
		stage.UnstageBranch(_directionallight)
	}
	for _, _mesh := range canvas.Meshs {
		stage.UnstageBranch(_mesh)
	}

}

func (curve *Curve) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(curve) {
		return
	}

	curve.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _vector3 := range curve.Points {
		stage.UnstageBranch(_vector3)
	}

}

func (cylindergeometry *CylinderGeometry) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(cylindergeometry) {
		return
	}

	cylindergeometry.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (directionallight *DirectionalLight) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(directionallight) {
		return
	}

	directionallight.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (extrudegeometry *ExtrudeGeometry) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(extrudegeometry) {
		return
	}

	extrudegeometry.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if extrudegeometry.Shape != nil {
		stage.UnstageBranch(extrudegeometry.Shape)
	}
	if extrudegeometry.ExtrudePath != nil {
		stage.UnstageBranch(extrudegeometry.ExtrudePath)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (mesh *Mesh) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(mesh) {
		return
	}

	mesh.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if mesh.MeshMaterialBasic != nil {
		stage.UnstageBranch(mesh.MeshMaterialBasic)
	}
	if mesh.MeshPhysicalMaterial != nil {
		stage.UnstageBranch(mesh.MeshPhysicalMaterial)
	}
	if mesh.CylinderGeometry != nil {
		stage.UnstageBranch(mesh.CylinderGeometry)
	}
	if mesh.BoxGeometry != nil {
		stage.UnstageBranch(mesh.BoxGeometry)
	}
	if mesh.SphereGeometry != nil {
		stage.UnstageBranch(mesh.SphereGeometry)
	}
	if mesh.TorusGeometry != nil {
		stage.UnstageBranch(mesh.TorusGeometry)
	}
	if mesh.PlaneGeometry != nil {
		stage.UnstageBranch(mesh.PlaneGeometry)
	}
	if mesh.TubeGeometry != nil {
		stage.UnstageBranch(mesh.TubeGeometry)
	}
	if mesh.ExtrudeGeometry != nil {
		stage.UnstageBranch(mesh.ExtrudeGeometry)
	}
	if mesh.BufferGeometry != nil {
		stage.UnstageBranch(mesh.BufferGeometry)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (meshmaterialbasic *MeshMaterialBasic) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(meshmaterialbasic) {
		return
	}

	meshmaterialbasic.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(meshphysicalmaterial) {
		return
	}

	meshphysicalmaterial.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (planegeometry *PlaneGeometry) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(planegeometry) {
		return
	}

	planegeometry.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (shape *Shape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(shape) {
		return
	}

	shape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _vector2 := range shape.Points {
		stage.UnstageBranch(_vector2)
	}

}

func (spheregeometry *SphereGeometry) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(spheregeometry) {
		return
	}

	spheregeometry.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (torusgeometry *TorusGeometry) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(torusgeometry) {
		return
	}

	torusgeometry.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (triangle *Triangle) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(triangle) {
		return
	}

	triangle.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tubegeometry *TubeGeometry) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tubegeometry) {
		return
	}

	tubegeometry.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tubegeometry.Path != nil {
		stage.UnstageBranch(tubegeometry.Path)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (vector2 *Vector2) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(vector2) {
		return
	}

	vector2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (vector3 *Vector3) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(vector3) {
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

func (reference *BufferGeometry) GongReconstructPointersFromReferences(stage *Stage, instance *BufferGeometry) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Vertices, stage.Vector3s_reference, instance.Vertices)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Faces, stage.Triangles_reference, instance.Faces)
}

func (reference *Camera) GongReconstructPointersFromReferences(stage *Stage, instance *Camera) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Canvas) GongReconstructPointersFromReferences(stage *Stage, instance *Canvas) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.AmbiantLight, stage.AmbiantLights_reference, instance.AmbiantLight)
	__gong__reconstructPointer(&reference.Camera, stage.Cameras_reference, instance.Camera)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.DirectionalLights, stage.DirectionalLights_reference, instance.DirectionalLights)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Meshs, stage.Meshs_reference, instance.Meshs)
}

func (reference *Curve) GongReconstructPointersFromReferences(stage *Stage, instance *Curve) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Points, stage.Vector3s_reference, instance.Points)
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
	__gong__reconstructPointer(&reference.Shape, stage.Shapes_reference, instance.Shape)
	__gong__reconstructPointer(&reference.ExtrudePath, stage.Curves_reference, instance.ExtrudePath)
	// insertion point for slice of pointers field
}

func (reference *Mesh) GongReconstructPointersFromReferences(stage *Stage, instance *Mesh) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MeshMaterialBasic, stage.MeshMaterialBasics_reference, instance.MeshMaterialBasic)
	__gong__reconstructPointer(&reference.MeshPhysicalMaterial, stage.MeshPhysicalMaterials_reference, instance.MeshPhysicalMaterial)
	__gong__reconstructPointer(&reference.CylinderGeometry, stage.CylinderGeometrys_reference, instance.CylinderGeometry)
	__gong__reconstructPointer(&reference.BoxGeometry, stage.BoxGeometrys_reference, instance.BoxGeometry)
	__gong__reconstructPointer(&reference.SphereGeometry, stage.SphereGeometrys_reference, instance.SphereGeometry)
	__gong__reconstructPointer(&reference.TorusGeometry, stage.TorusGeometrys_reference, instance.TorusGeometry)
	__gong__reconstructPointer(&reference.PlaneGeometry, stage.PlaneGeometrys_reference, instance.PlaneGeometry)
	__gong__reconstructPointer(&reference.TubeGeometry, stage.TubeGeometrys_reference, instance.TubeGeometry)
	__gong__reconstructPointer(&reference.ExtrudeGeometry, stage.ExtrudeGeometrys_reference, instance.ExtrudeGeometry)
	__gong__reconstructPointer(&reference.BufferGeometry, stage.BufferGeometrys_reference, instance.BufferGeometry)
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
	__gong__reconstructSliceOfPointersFromReferences(&reference.Points, stage.Vector2s_reference, instance.Points)
}

func (reference *SphereGeometry) GongReconstructPointersFromReferences(stage *Stage, instance *SphereGeometry) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TorusGeometry) GongReconstructPointersFromReferences(stage *Stage, instance *TorusGeometry) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Triangle) GongReconstructPointersFromReferences(stage *Stage, instance *Triangle) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TubeGeometry) GongReconstructPointersFromReferences(stage *Stage, instance *TubeGeometry) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Path, stage.Curves_reference, instance.Path)
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

func (reference *BufferGeometry) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Vertices, stage.Vector3s_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Faces, stage.Triangles_instance)
}

func (reference *Camera) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Canvas) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.AmbiantLight, stage.AmbiantLights_instance)
	__gong__reconstructPointerFromInstance(&reference.Camera, stage.Cameras_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.DirectionalLights, stage.DirectionalLights_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Meshs, stage.Meshs_instance)
}

func (reference *Curve) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Points, stage.Vector3s_instance)
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
	__gong__reconstructPointerFromInstance(&reference.Shape, stage.Shapes_instance)
	__gong__reconstructPointerFromInstance(&reference.ExtrudePath, stage.Curves_instance)
	// insertion point for slice of pointers fields
}

func (reference *Mesh) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MeshMaterialBasic, stage.MeshMaterialBasics_instance)
	__gong__reconstructPointerFromInstance(&reference.MeshPhysicalMaterial, stage.MeshPhysicalMaterials_instance)
	__gong__reconstructPointerFromInstance(&reference.CylinderGeometry, stage.CylinderGeometrys_instance)
	__gong__reconstructPointerFromInstance(&reference.BoxGeometry, stage.BoxGeometrys_instance)
	__gong__reconstructPointerFromInstance(&reference.SphereGeometry, stage.SphereGeometrys_instance)
	__gong__reconstructPointerFromInstance(&reference.TorusGeometry, stage.TorusGeometrys_instance)
	__gong__reconstructPointerFromInstance(&reference.PlaneGeometry, stage.PlaneGeometrys_instance)
	__gong__reconstructPointerFromInstance(&reference.TubeGeometry, stage.TubeGeometrys_instance)
	__gong__reconstructPointerFromInstance(&reference.ExtrudeGeometry, stage.ExtrudeGeometrys_instance)
	__gong__reconstructPointerFromInstance(&reference.BufferGeometry, stage.BufferGeometrys_instance)
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
	__gong__reconstructSliceOfPointersFromInstances(&reference.Points, stage.Vector2s_instance)
}

func (reference *SphereGeometry) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TorusGeometry) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Triangle) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TubeGeometry) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Path, stage.Curves_instance)
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
func (buffergeometry *BufferGeometry) GongDiff(stage *Stage, buffergeometryOther *BufferGeometry) (diffs []string) {
	// insertion point for field diffs
	if buffergeometry.Name != buffergeometryOther.Name {
		diffs = append(diffs, buffergeometry.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, buffergeometry, "Vertices", buffergeometryOther.Vertices, buffergeometry.Vertices); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, buffergeometry, "Faces", buffergeometryOther.Faces, buffergeometry.Faces); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (camera *Camera) GongDiff(stage *Stage, cameraOther *Camera) (diffs []string) {
	// insertion point for field diffs
	if camera.Name != cameraOther.Name {
		diffs = append(diffs, camera.GongMarshallField(stage, "Name"))
	}
	if camera.X != cameraOther.X {
		diffs = append(diffs, camera.GongMarshallField(stage, "X"))
	}
	if camera.Y != cameraOther.Y {
		diffs = append(diffs, camera.GongMarshallField(stage, "Y"))
	}
	if camera.Z != cameraOther.Z {
		diffs = append(diffs, camera.GongMarshallField(stage, "Z"))
	}
	if camera.TargetX != cameraOther.TargetX {
		diffs = append(diffs, camera.GongMarshallField(stage, "TargetX"))
	}
	if camera.TargetY != cameraOther.TargetY {
		diffs = append(diffs, camera.GongMarshallField(stage, "TargetY"))
	}
	if camera.TargetZ != cameraOther.TargetZ {
		diffs = append(diffs, camera.GongMarshallField(stage, "TargetZ"))
	}
	if camera.Fov != cameraOther.Fov {
		diffs = append(diffs, camera.GongMarshallField(stage, "Fov"))
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
	if ops := __gong__diffSliceOfPointers(stage, canvas, "DirectionalLights", canvasOther.DirectionalLights, canvas.DirectionalLights); ops != "" {
		diffs = append(diffs, ops)
	}
	if canvas.AmbiantLight != canvasOther.AmbiantLight {
		diffs = append(diffs, canvas.GongMarshallField(stage, "AmbiantLight"))
	}
	if ops := __gong__diffSliceOfPointers(stage, canvas, "Meshs", canvasOther.Meshs, canvas.Meshs); ops != "" {
		diffs = append(diffs, ops)
	}
	if canvas.Camera != canvasOther.Camera {
		diffs = append(diffs, canvas.GongMarshallField(stage, "Camera"))
	}
	if canvas.IsWithLastRenderingUpdate != canvasOther.IsWithLastRenderingUpdate {
		diffs = append(diffs, canvas.GongMarshallField(stage, "IsWithLastRenderingUpdate"))
	}
	if canvas.LastRendering != canvasOther.LastRendering {
		diffs = append(diffs, canvas.GongMarshallField(stage, "LastRendering"))
	}
	if canvas.Frame64BitsEncoded != canvasOther.Frame64BitsEncoded {
		diffs = append(diffs, canvas.GongMarshallField(stage, "Frame64BitsEncoded"))
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
	if ops := __gong__diffSliceOfPointers(stage, curve, "Points", curveOther.Points, curve.Points); ops != "" {
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
	if extrudegeometry.Shape != extrudegeometryOther.Shape {
		diffs = append(diffs, extrudegeometry.GongMarshallField(stage, "Shape"))
	}
	if extrudegeometry.ExtrudePath != extrudegeometryOther.ExtrudePath {
		diffs = append(diffs, extrudegeometry.GongMarshallField(stage, "ExtrudePath"))
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
	if mesh.MeshMaterialBasic != meshOther.MeshMaterialBasic {
		diffs = append(diffs, mesh.GongMarshallField(stage, "MeshMaterialBasic"))
	}
	if mesh.MeshPhysicalMaterial != meshOther.MeshPhysicalMaterial {
		diffs = append(diffs, mesh.GongMarshallField(stage, "MeshPhysicalMaterial"))
	}
	if mesh.CylinderGeometry != meshOther.CylinderGeometry {
		diffs = append(diffs, mesh.GongMarshallField(stage, "CylinderGeometry"))
	}
	if mesh.BoxGeometry != meshOther.BoxGeometry {
		diffs = append(diffs, mesh.GongMarshallField(stage, "BoxGeometry"))
	}
	if mesh.SphereGeometry != meshOther.SphereGeometry {
		diffs = append(diffs, mesh.GongMarshallField(stage, "SphereGeometry"))
	}
	if mesh.TorusGeometry != meshOther.TorusGeometry {
		diffs = append(diffs, mesh.GongMarshallField(stage, "TorusGeometry"))
	}
	if mesh.PlaneGeometry != meshOther.PlaneGeometry {
		diffs = append(diffs, mesh.GongMarshallField(stage, "PlaneGeometry"))
	}
	if mesh.TubeGeometry != meshOther.TubeGeometry {
		diffs = append(diffs, mesh.GongMarshallField(stage, "TubeGeometry"))
	}
	if mesh.ExtrudeGeometry != meshOther.ExtrudeGeometry {
		diffs = append(diffs, mesh.GongMarshallField(stage, "ExtrudeGeometry"))
	}
	if mesh.BufferGeometry != meshOther.BufferGeometry {
		diffs = append(diffs, mesh.GongMarshallField(stage, "BufferGeometry"))
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
	if ops := __gong__diffSliceOfPointers(stage, shape, "Points", shapeOther.Points, shape.Points); ops != "" {
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
func (triangle *Triangle) GongDiff(stage *Stage, triangleOther *Triangle) (diffs []string) {
	// insertion point for field diffs
	if triangle.Name != triangleOther.Name {
		diffs = append(diffs, triangle.GongMarshallField(stage, "Name"))
	}
	if triangle.V1 != triangleOther.V1 {
		diffs = append(diffs, triangle.GongMarshallField(stage, "V1"))
	}
	if triangle.V2 != triangleOther.V2 {
		diffs = append(diffs, triangle.GongMarshallField(stage, "V2"))
	}
	if triangle.V3 != triangleOther.V3 {
		diffs = append(diffs, triangle.GongMarshallField(stage, "V3"))
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
	if tubegeometry.Path != tubegeometryOther.Path {
		diffs = append(diffs, tubegeometry.GongMarshallField(stage, "Path"))
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

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
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

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
