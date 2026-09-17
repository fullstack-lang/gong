// generated code - do not edit
package models

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct AmbiantLight
	// insertion point per field

	// Compute reverse map for named struct BoxGeometry
	// insertion point per field

	// Compute reverse map for named struct BufferGeometry
	// insertion point per field
	stage.BufferGeometry_Vertices_reverseMap = make(map[*Vector3]*BufferGeometry)
	for buffergeometry := range stage.BufferGeometrys {
		_ = buffergeometry
		for _, _vector3 := range buffergeometry.Vertices {
			stage.BufferGeometry_Vertices_reverseMap[_vector3] = buffergeometry
		}
	}
	stage.BufferGeometry_Faces_reverseMap = make(map[*Triangle]*BufferGeometry)
	for buffergeometry := range stage.BufferGeometrys {
		_ = buffergeometry
		for _, _triangle := range buffergeometry.Faces {
			stage.BufferGeometry_Faces_reverseMap[_triangle] = buffergeometry
		}
	}

	// Compute reverse map for named struct Camera
	// insertion point per field

	// Compute reverse map for named struct Canvas
	// insertion point per field
	stage.Canvas_DirectionalLights_reverseMap = make(map[*DirectionalLight]*Canvas)
	for canvas := range stage.Canvass {
		_ = canvas
		for _, _directionallight := range canvas.DirectionalLights {
			stage.Canvas_DirectionalLights_reverseMap[_directionallight] = canvas
		}
	}
	stage.Canvas_Meshs_reverseMap = make(map[*Mesh]*Canvas)
	for canvas := range stage.Canvass {
		_ = canvas
		for _, _mesh := range canvas.Meshs {
			stage.Canvas_Meshs_reverseMap[_mesh] = canvas
		}
	}

	// Compute reverse map for named struct Curve
	// insertion point per field
	stage.Curve_Points_reverseMap = make(map[*Vector3]*Curve)
	for curve := range stage.Curves {
		_ = curve
		for _, _vector3 := range curve.Points {
			stage.Curve_Points_reverseMap[_vector3] = curve
		}
	}

	// Compute reverse map for named struct CylinderGeometry
	// insertion point per field

	// Compute reverse map for named struct DirectionalLight
	// insertion point per field

	// Compute reverse map for named struct ExtrudeGeometry
	// insertion point per field

	// Compute reverse map for named struct Mesh
	// insertion point per field

	// Compute reverse map for named struct MeshMaterialBasic
	// insertion point per field

	// Compute reverse map for named struct MeshPhysicalMaterial
	// insertion point per field

	// Compute reverse map for named struct PlaneGeometry
	// insertion point per field

	// Compute reverse map for named struct Shape
	// insertion point per field
	stage.Shape_Points_reverseMap = make(map[*Vector2]*Shape)
	for shape := range stage.Shapes {
		_ = shape
		for _, _vector2 := range shape.Points {
			stage.Shape_Points_reverseMap[_vector2] = shape
		}
	}

	// Compute reverse map for named struct SphereGeometry
	// insertion point per field

	// Compute reverse map for named struct TorusGeometry
	// insertion point per field

	// Compute reverse map for named struct Triangle
	// insertion point per field

	// Compute reverse map for named struct TubeGeometry
	// insertion point per field

	// Compute reverse map for named struct Vector2
	// insertion point per field

	// Compute reverse map for named struct Vector3
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.AmbiantLights {
		res = append(res, instance)
	}

	for instance := range stage.BoxGeometrys {
		res = append(res, instance)
	}

	for instance := range stage.BufferGeometrys {
		res = append(res, instance)
	}

	for instance := range stage.Cameras {
		res = append(res, instance)
	}

	for instance := range stage.Canvass {
		res = append(res, instance)
	}

	for instance := range stage.Curves {
		res = append(res, instance)
	}

	for instance := range stage.CylinderGeometrys {
		res = append(res, instance)
	}

	for instance := range stage.DirectionalLights {
		res = append(res, instance)
	}

	for instance := range stage.ExtrudeGeometrys {
		res = append(res, instance)
	}

	for instance := range stage.Meshs {
		res = append(res, instance)
	}

	for instance := range stage.MeshMaterialBasics {
		res = append(res, instance)
	}

	for instance := range stage.MeshPhysicalMaterials {
		res = append(res, instance)
	}

	for instance := range stage.PlaneGeometrys {
		res = append(res, instance)
	}

	for instance := range stage.Shapes {
		res = append(res, instance)
	}

	for instance := range stage.SphereGeometrys {
		res = append(res, instance)
	}

	for instance := range stage.TorusGeometrys {
		res = append(res, instance)
	}

	for instance := range stage.Triangles {
		res = append(res, instance)
	}

	for instance := range stage.TubeGeometrys {
		res = append(res, instance)
	}

	for instance := range stage.Vector2s {
		res = append(res, instance)
	}

	for instance := range stage.Vector3s {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (ambiantlight *AmbiantLight) GongCopy() GongstructIF {
	newInstance := new(AmbiantLight)
	ambiantlight.GongCopyBasicFields(newInstance)
	return newInstance
}

func (boxgeometry *BoxGeometry) GongCopy() GongstructIF {
	newInstance := new(BoxGeometry)
	boxgeometry.GongCopyBasicFields(newInstance)
	return newInstance
}

func (buffergeometry *BufferGeometry) GongCopy() GongstructIF {
	newInstance := new(BufferGeometry)
	buffergeometry.GongCopyBasicFields(newInstance)
	return newInstance
}

func (camera *Camera) GongCopy() GongstructIF {
	newInstance := new(Camera)
	camera.GongCopyBasicFields(newInstance)
	return newInstance
}

func (canvas *Canvas) GongCopy() GongstructIF {
	newInstance := new(Canvas)
	canvas.GongCopyBasicFields(newInstance)
	return newInstance
}

func (curve *Curve) GongCopy() GongstructIF {
	newInstance := new(Curve)
	curve.GongCopyBasicFields(newInstance)
	return newInstance
}

func (cylindergeometry *CylinderGeometry) GongCopy() GongstructIF {
	newInstance := new(CylinderGeometry)
	cylindergeometry.GongCopyBasicFields(newInstance)
	return newInstance
}

func (directionallight *DirectionalLight) GongCopy() GongstructIF {
	newInstance := new(DirectionalLight)
	directionallight.GongCopyBasicFields(newInstance)
	return newInstance
}

func (extrudegeometry *ExtrudeGeometry) GongCopy() GongstructIF {
	newInstance := new(ExtrudeGeometry)
	extrudegeometry.GongCopyBasicFields(newInstance)
	return newInstance
}

func (mesh *Mesh) GongCopy() GongstructIF {
	newInstance := new(Mesh)
	mesh.GongCopyBasicFields(newInstance)
	return newInstance
}

func (meshmaterialbasic *MeshMaterialBasic) GongCopy() GongstructIF {
	newInstance := new(MeshMaterialBasic)
	meshmaterialbasic.GongCopyBasicFields(newInstance)
	return newInstance
}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongCopy() GongstructIF {
	newInstance := new(MeshPhysicalMaterial)
	meshphysicalmaterial.GongCopyBasicFields(newInstance)
	return newInstance
}

func (planegeometry *PlaneGeometry) GongCopy() GongstructIF {
	newInstance := new(PlaneGeometry)
	planegeometry.GongCopyBasicFields(newInstance)
	return newInstance
}

func (shape *Shape) GongCopy() GongstructIF {
	newInstance := new(Shape)
	shape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (spheregeometry *SphereGeometry) GongCopy() GongstructIF {
	newInstance := new(SphereGeometry)
	spheregeometry.GongCopyBasicFields(newInstance)
	return newInstance
}

func (torusgeometry *TorusGeometry) GongCopy() GongstructIF {
	newInstance := new(TorusGeometry)
	torusgeometry.GongCopyBasicFields(newInstance)
	return newInstance
}

func (triangle *Triangle) GongCopy() GongstructIF {
	newInstance := new(Triangle)
	triangle.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tubegeometry *TubeGeometry) GongCopy() GongstructIF {
	newInstance := new(TubeGeometry)
	tubegeometry.GongCopyBasicFields(newInstance)
	return newInstance
}

func (vector2 *Vector2) GongCopy() GongstructIF {
	newInstance := new(Vector2)
	vector2.GongCopyBasicFields(newInstance)
	return newInstance
}

func (vector3 *Vector3) GongCopy() GongstructIF {
	newInstance := new(Vector3)
	vector3.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (ambiantlight *AmbiantLight) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(ambiantlight).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(ambiantlight), uint64(stage.GetOrder(ambiantlight)))
	return
}

func (boxgeometry *BoxGeometry) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(boxgeometry).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(boxgeometry), uint64(stage.GetOrder(boxgeometry)))
	return
}

func (buffergeometry *BufferGeometry) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(buffergeometry).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(buffergeometry), uint64(stage.GetOrder(buffergeometry)))
	return
}

func (camera *Camera) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(camera).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(camera), uint64(stage.GetOrder(camera)))
	return
}

func (canvas *Canvas) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(canvas).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(canvas), uint64(stage.GetOrder(canvas)))
	return
}

func (curve *Curve) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(curve).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(curve), uint64(stage.GetOrder(curve)))
	return
}

func (cylindergeometry *CylinderGeometry) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(cylindergeometry).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(cylindergeometry), uint64(stage.GetOrder(cylindergeometry)))
	return
}

func (directionallight *DirectionalLight) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(directionallight).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(directionallight), uint64(stage.GetOrder(directionallight)))
	return
}

func (extrudegeometry *ExtrudeGeometry) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(extrudegeometry).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(extrudegeometry), uint64(stage.GetOrder(extrudegeometry)))
	return
}

func (mesh *Mesh) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(mesh).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(mesh), uint64(stage.GetOrder(mesh)))
	return
}

func (meshmaterialbasic *MeshMaterialBasic) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(meshmaterialbasic).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(meshmaterialbasic), uint64(stage.GetOrder(meshmaterialbasic)))
	return
}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(meshphysicalmaterial).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(meshphysicalmaterial), uint64(stage.GetOrder(meshphysicalmaterial)))
	return
}

func (planegeometry *PlaneGeometry) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(planegeometry).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(planegeometry), uint64(stage.GetOrder(planegeometry)))
	return
}

func (shape *Shape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(shape), uint64(stage.GetOrder(shape)))
	return
}

func (spheregeometry *SphereGeometry) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(spheregeometry).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(spheregeometry), uint64(stage.GetOrder(spheregeometry)))
	return
}

func (torusgeometry *TorusGeometry) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(torusgeometry).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(torusgeometry), uint64(stage.GetOrder(torusgeometry)))
	return
}

func (triangle *Triangle) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(triangle).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(triangle), uint64(stage.GetOrder(triangle)))
	return
}

func (tubegeometry *TubeGeometry) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(tubegeometry).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(tubegeometry), uint64(stage.GetOrder(tubegeometry)))
	return
}

func (vector2 *Vector2) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(vector2).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(vector2), uint64(stage.GetOrder(vector2)))
	return
}

func (vector3 *Vector3) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(vector3).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(vector3), uint64(stage.GetOrder(vector3)))
	return
}


type GongstructDiffable[T any] interface {
	PointerToGongstruct
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	computeCommitsForType(
		stage,
		stage.AmbiantLights,
		stage.AmbiantLight_stagedOrder,
		stage.AmbiantLights_reference,
		&stage.AmbiantLights_referenceOrder,
		stage.AmbiantLights_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.BoxGeometrys,
		stage.BoxGeometry_stagedOrder,
		stage.BoxGeometrys_reference,
		&stage.BoxGeometrys_referenceOrder,
		stage.BoxGeometrys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.BufferGeometrys,
		stage.BufferGeometry_stagedOrder,
		stage.BufferGeometrys_reference,
		&stage.BufferGeometrys_referenceOrder,
		stage.BufferGeometrys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Cameras,
		stage.Camera_stagedOrder,
		stage.Cameras_reference,
		&stage.Cameras_referenceOrder,
		stage.Cameras_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Canvass,
		stage.Canvas_stagedOrder,
		stage.Canvass_reference,
		&stage.Canvass_referenceOrder,
		stage.Canvass_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Curves,
		stage.Curve_stagedOrder,
		stage.Curves_reference,
		&stage.Curves_referenceOrder,
		stage.Curves_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CylinderGeometrys,
		stage.CylinderGeometry_stagedOrder,
		stage.CylinderGeometrys_reference,
		&stage.CylinderGeometrys_referenceOrder,
		stage.CylinderGeometrys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DirectionalLights,
		stage.DirectionalLight_stagedOrder,
		stage.DirectionalLights_reference,
		&stage.DirectionalLights_referenceOrder,
		stage.DirectionalLights_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ExtrudeGeometrys,
		stage.ExtrudeGeometry_stagedOrder,
		stage.ExtrudeGeometrys_reference,
		&stage.ExtrudeGeometrys_referenceOrder,
		stage.ExtrudeGeometrys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Meshs,
		stage.Mesh_stagedOrder,
		stage.Meshs_reference,
		&stage.Meshs_referenceOrder,
		stage.Meshs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MeshMaterialBasics,
		stage.MeshMaterialBasic_stagedOrder,
		stage.MeshMaterialBasics_reference,
		&stage.MeshMaterialBasics_referenceOrder,
		stage.MeshMaterialBasics_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MeshPhysicalMaterials,
		stage.MeshPhysicalMaterial_stagedOrder,
		stage.MeshPhysicalMaterials_reference,
		&stage.MeshPhysicalMaterials_referenceOrder,
		stage.MeshPhysicalMaterials_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.PlaneGeometrys,
		stage.PlaneGeometry_stagedOrder,
		stage.PlaneGeometrys_reference,
		&stage.PlaneGeometrys_referenceOrder,
		stage.PlaneGeometrys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Shapes,
		stage.Shape_stagedOrder,
		stage.Shapes_reference,
		&stage.Shapes_referenceOrder,
		stage.Shapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SphereGeometrys,
		stage.SphereGeometry_stagedOrder,
		stage.SphereGeometrys_reference,
		&stage.SphereGeometrys_referenceOrder,
		stage.SphereGeometrys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TorusGeometrys,
		stage.TorusGeometry_stagedOrder,
		stage.TorusGeometrys_reference,
		&stage.TorusGeometrys_referenceOrder,
		stage.TorusGeometrys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Triangles,
		stage.Triangle_stagedOrder,
		stage.Triangles_reference,
		&stage.Triangles_referenceOrder,
		stage.Triangles_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TubeGeometrys,
		stage.TubeGeometry_stagedOrder,
		stage.TubeGeometrys_reference,
		&stage.TubeGeometrys_referenceOrder,
		stage.TubeGeometrys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Vector2s,
		stage.Vector2_stagedOrder,
		stage.Vector2s_reference,
		&stage.Vector2s_referenceOrder,
		stage.Vector2s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Vector3s,
		stage.Vector3_stagedOrder,
		stage.Vector3s_reference,
		&stage.Vector3s_referenceOrder,
		stage.Vector3s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)
		stage.modified = true
	} else {
		stage.modified = false
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	stage.AmbiantLights_reference = make(map[*AmbiantLight]*AmbiantLight)
	stage.AmbiantLights_referenceOrder = make(map[*AmbiantLight]uint) // diff Unstage needs the reference order
	stage.AmbiantLights_instance = make(map[*AmbiantLight]*AmbiantLight)
	for instance := range stage.AmbiantLights {
		_copy := instance.GongCopy().(*AmbiantLight)
		stage.AmbiantLights_reference[instance] = _copy
		stage.AmbiantLights_instance[_copy] = instance
		stage.AmbiantLights_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.BoxGeometrys_reference = make(map[*BoxGeometry]*BoxGeometry)
	stage.BoxGeometrys_referenceOrder = make(map[*BoxGeometry]uint) // diff Unstage needs the reference order
	stage.BoxGeometrys_instance = make(map[*BoxGeometry]*BoxGeometry)
	for instance := range stage.BoxGeometrys {
		_copy := instance.GongCopy().(*BoxGeometry)
		stage.BoxGeometrys_reference[instance] = _copy
		stage.BoxGeometrys_instance[_copy] = instance
		stage.BoxGeometrys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.BufferGeometrys_reference = make(map[*BufferGeometry]*BufferGeometry)
	stage.BufferGeometrys_referenceOrder = make(map[*BufferGeometry]uint) // diff Unstage needs the reference order
	stage.BufferGeometrys_instance = make(map[*BufferGeometry]*BufferGeometry)
	for instance := range stage.BufferGeometrys {
		_copy := instance.GongCopy().(*BufferGeometry)
		stage.BufferGeometrys_reference[instance] = _copy
		stage.BufferGeometrys_instance[_copy] = instance
		stage.BufferGeometrys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Cameras_reference = make(map[*Camera]*Camera)
	stage.Cameras_referenceOrder = make(map[*Camera]uint) // diff Unstage needs the reference order
	stage.Cameras_instance = make(map[*Camera]*Camera)
	for instance := range stage.Cameras {
		_copy := instance.GongCopy().(*Camera)
		stage.Cameras_reference[instance] = _copy
		stage.Cameras_instance[_copy] = instance
		stage.Cameras_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Canvass_reference = make(map[*Canvas]*Canvas)
	stage.Canvass_referenceOrder = make(map[*Canvas]uint) // diff Unstage needs the reference order
	stage.Canvass_instance = make(map[*Canvas]*Canvas)
	for instance := range stage.Canvass {
		_copy := instance.GongCopy().(*Canvas)
		stage.Canvass_reference[instance] = _copy
		stage.Canvass_instance[_copy] = instance
		stage.Canvass_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Curves_reference = make(map[*Curve]*Curve)
	stage.Curves_referenceOrder = make(map[*Curve]uint) // diff Unstage needs the reference order
	stage.Curves_instance = make(map[*Curve]*Curve)
	for instance := range stage.Curves {
		_copy := instance.GongCopy().(*Curve)
		stage.Curves_reference[instance] = _copy
		stage.Curves_instance[_copy] = instance
		stage.Curves_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.CylinderGeometrys_reference = make(map[*CylinderGeometry]*CylinderGeometry)
	stage.CylinderGeometrys_referenceOrder = make(map[*CylinderGeometry]uint) // diff Unstage needs the reference order
	stage.CylinderGeometrys_instance = make(map[*CylinderGeometry]*CylinderGeometry)
	for instance := range stage.CylinderGeometrys {
		_copy := instance.GongCopy().(*CylinderGeometry)
		stage.CylinderGeometrys_reference[instance] = _copy
		stage.CylinderGeometrys_instance[_copy] = instance
		stage.CylinderGeometrys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.DirectionalLights_reference = make(map[*DirectionalLight]*DirectionalLight)
	stage.DirectionalLights_referenceOrder = make(map[*DirectionalLight]uint) // diff Unstage needs the reference order
	stage.DirectionalLights_instance = make(map[*DirectionalLight]*DirectionalLight)
	for instance := range stage.DirectionalLights {
		_copy := instance.GongCopy().(*DirectionalLight)
		stage.DirectionalLights_reference[instance] = _copy
		stage.DirectionalLights_instance[_copy] = instance
		stage.DirectionalLights_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ExtrudeGeometrys_reference = make(map[*ExtrudeGeometry]*ExtrudeGeometry)
	stage.ExtrudeGeometrys_referenceOrder = make(map[*ExtrudeGeometry]uint) // diff Unstage needs the reference order
	stage.ExtrudeGeometrys_instance = make(map[*ExtrudeGeometry]*ExtrudeGeometry)
	for instance := range stage.ExtrudeGeometrys {
		_copy := instance.GongCopy().(*ExtrudeGeometry)
		stage.ExtrudeGeometrys_reference[instance] = _copy
		stage.ExtrudeGeometrys_instance[_copy] = instance
		stage.ExtrudeGeometrys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Meshs_reference = make(map[*Mesh]*Mesh)
	stage.Meshs_referenceOrder = make(map[*Mesh]uint) // diff Unstage needs the reference order
	stage.Meshs_instance = make(map[*Mesh]*Mesh)
	for instance := range stage.Meshs {
		_copy := instance.GongCopy().(*Mesh)
		stage.Meshs_reference[instance] = _copy
		stage.Meshs_instance[_copy] = instance
		stage.Meshs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.MeshMaterialBasics_reference = make(map[*MeshMaterialBasic]*MeshMaterialBasic)
	stage.MeshMaterialBasics_referenceOrder = make(map[*MeshMaterialBasic]uint) // diff Unstage needs the reference order
	stage.MeshMaterialBasics_instance = make(map[*MeshMaterialBasic]*MeshMaterialBasic)
	for instance := range stage.MeshMaterialBasics {
		_copy := instance.GongCopy().(*MeshMaterialBasic)
		stage.MeshMaterialBasics_reference[instance] = _copy
		stage.MeshMaterialBasics_instance[_copy] = instance
		stage.MeshMaterialBasics_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.MeshPhysicalMaterials_reference = make(map[*MeshPhysicalMaterial]*MeshPhysicalMaterial)
	stage.MeshPhysicalMaterials_referenceOrder = make(map[*MeshPhysicalMaterial]uint) // diff Unstage needs the reference order
	stage.MeshPhysicalMaterials_instance = make(map[*MeshPhysicalMaterial]*MeshPhysicalMaterial)
	for instance := range stage.MeshPhysicalMaterials {
		_copy := instance.GongCopy().(*MeshPhysicalMaterial)
		stage.MeshPhysicalMaterials_reference[instance] = _copy
		stage.MeshPhysicalMaterials_instance[_copy] = instance
		stage.MeshPhysicalMaterials_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PlaneGeometrys_reference = make(map[*PlaneGeometry]*PlaneGeometry)
	stage.PlaneGeometrys_referenceOrder = make(map[*PlaneGeometry]uint) // diff Unstage needs the reference order
	stage.PlaneGeometrys_instance = make(map[*PlaneGeometry]*PlaneGeometry)
	for instance := range stage.PlaneGeometrys {
		_copy := instance.GongCopy().(*PlaneGeometry)
		stage.PlaneGeometrys_reference[instance] = _copy
		stage.PlaneGeometrys_instance[_copy] = instance
		stage.PlaneGeometrys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Shapes_reference = make(map[*Shape]*Shape)
	stage.Shapes_referenceOrder = make(map[*Shape]uint) // diff Unstage needs the reference order
	stage.Shapes_instance = make(map[*Shape]*Shape)
	for instance := range stage.Shapes {
		_copy := instance.GongCopy().(*Shape)
		stage.Shapes_reference[instance] = _copy
		stage.Shapes_instance[_copy] = instance
		stage.Shapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SphereGeometrys_reference = make(map[*SphereGeometry]*SphereGeometry)
	stage.SphereGeometrys_referenceOrder = make(map[*SphereGeometry]uint) // diff Unstage needs the reference order
	stage.SphereGeometrys_instance = make(map[*SphereGeometry]*SphereGeometry)
	for instance := range stage.SphereGeometrys {
		_copy := instance.GongCopy().(*SphereGeometry)
		stage.SphereGeometrys_reference[instance] = _copy
		stage.SphereGeometrys_instance[_copy] = instance
		stage.SphereGeometrys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TorusGeometrys_reference = make(map[*TorusGeometry]*TorusGeometry)
	stage.TorusGeometrys_referenceOrder = make(map[*TorusGeometry]uint) // diff Unstage needs the reference order
	stage.TorusGeometrys_instance = make(map[*TorusGeometry]*TorusGeometry)
	for instance := range stage.TorusGeometrys {
		_copy := instance.GongCopy().(*TorusGeometry)
		stage.TorusGeometrys_reference[instance] = _copy
		stage.TorusGeometrys_instance[_copy] = instance
		stage.TorusGeometrys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Triangles_reference = make(map[*Triangle]*Triangle)
	stage.Triangles_referenceOrder = make(map[*Triangle]uint) // diff Unstage needs the reference order
	stage.Triangles_instance = make(map[*Triangle]*Triangle)
	for instance := range stage.Triangles {
		_copy := instance.GongCopy().(*Triangle)
		stage.Triangles_reference[instance] = _copy
		stage.Triangles_instance[_copy] = instance
		stage.Triangles_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TubeGeometrys_reference = make(map[*TubeGeometry]*TubeGeometry)
	stage.TubeGeometrys_referenceOrder = make(map[*TubeGeometry]uint) // diff Unstage needs the reference order
	stage.TubeGeometrys_instance = make(map[*TubeGeometry]*TubeGeometry)
	for instance := range stage.TubeGeometrys {
		_copy := instance.GongCopy().(*TubeGeometry)
		stage.TubeGeometrys_reference[instance] = _copy
		stage.TubeGeometrys_instance[_copy] = instance
		stage.TubeGeometrys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Vector2s_reference = make(map[*Vector2]*Vector2)
	stage.Vector2s_referenceOrder = make(map[*Vector2]uint) // diff Unstage needs the reference order
	stage.Vector2s_instance = make(map[*Vector2]*Vector2)
	for instance := range stage.Vector2s {
		_copy := instance.GongCopy().(*Vector2)
		stage.Vector2s_reference[instance] = _copy
		stage.Vector2s_instance[_copy] = instance
		stage.Vector2s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Vector3s_reference = make(map[*Vector3]*Vector3)
	stage.Vector3s_referenceOrder = make(map[*Vector3]uint) // diff Unstage needs the reference order
	stage.Vector3s_instance = make(map[*Vector3]*Vector3)
	for instance := range stage.Vector3s {
		_copy := instance.GongCopy().(*Vector3)
		stage.Vector3s_reference[instance] = _copy
		stage.Vector3s_instance[_copy] = instance
		stage.Vector3s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.AmbiantLights {
		reference := stage.AmbiantLights_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.BoxGeometrys {
		reference := stage.BoxGeometrys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.BufferGeometrys {
		reference := stage.BufferGeometrys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Cameras {
		reference := stage.Cameras_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Canvass {
		reference := stage.Canvass_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Curves {
		reference := stage.Curves_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.CylinderGeometrys {
		reference := stage.CylinderGeometrys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.DirectionalLights {
		reference := stage.DirectionalLights_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ExtrudeGeometrys {
		reference := stage.ExtrudeGeometrys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Meshs {
		reference := stage.Meshs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.MeshMaterialBasics {
		reference := stage.MeshMaterialBasics_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.MeshPhysicalMaterials {
		reference := stage.MeshPhysicalMaterials_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PlaneGeometrys {
		reference := stage.PlaneGeometrys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Shapes {
		reference := stage.Shapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SphereGeometrys {
		reference := stage.SphereGeometrys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TorusGeometrys {
		reference := stage.TorusGeometrys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Triangles {
		reference := stage.Triangles_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TubeGeometrys {
		reference := stage.TubeGeometrys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Vector2s {
		reference := stage.Vector2s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Vector3s {
		reference := stage.Vector3s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (ambiantlight *AmbiantLight) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.AmbiantLight_stagedOrder[ambiantlight]; ok {
		return order
	}
	if order, ok := stage.AmbiantLights_referenceOrder[ambiantlight]; ok {
		return order
	} else {
		log.Printf("instance %p of type AmbiantLight was not staged and does not have a reference order", ambiantlight)
		return 0
	}
}

func (boxgeometry *BoxGeometry) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.BoxGeometry_stagedOrder[boxgeometry]; ok {
		return order
	}
	if order, ok := stage.BoxGeometrys_referenceOrder[boxgeometry]; ok {
		return order
	} else {
		log.Printf("instance %p of type BoxGeometry was not staged and does not have a reference order", boxgeometry)
		return 0
	}
}

func (buffergeometry *BufferGeometry) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.BufferGeometry_stagedOrder[buffergeometry]; ok {
		return order
	}
	if order, ok := stage.BufferGeometrys_referenceOrder[buffergeometry]; ok {
		return order
	} else {
		log.Printf("instance %p of type BufferGeometry was not staged and does not have a reference order", buffergeometry)
		return 0
	}
}

func (camera *Camera) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Camera_stagedOrder[camera]; ok {
		return order
	}
	if order, ok := stage.Cameras_referenceOrder[camera]; ok {
		return order
	} else {
		log.Printf("instance %p of type Camera was not staged and does not have a reference order", camera)
		return 0
	}
}

func (canvas *Canvas) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Canvas_stagedOrder[canvas]; ok {
		return order
	}
	if order, ok := stage.Canvass_referenceOrder[canvas]; ok {
		return order
	} else {
		log.Printf("instance %p of type Canvas was not staged and does not have a reference order", canvas)
		return 0
	}
}

func (curve *Curve) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Curve_stagedOrder[curve]; ok {
		return order
	}
	if order, ok := stage.Curves_referenceOrder[curve]; ok {
		return order
	} else {
		log.Printf("instance %p of type Curve was not staged and does not have a reference order", curve)
		return 0
	}
}

func (cylindergeometry *CylinderGeometry) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.CylinderGeometry_stagedOrder[cylindergeometry]; ok {
		return order
	}
	if order, ok := stage.CylinderGeometrys_referenceOrder[cylindergeometry]; ok {
		return order
	} else {
		log.Printf("instance %p of type CylinderGeometry was not staged and does not have a reference order", cylindergeometry)
		return 0
	}
}

func (directionallight *DirectionalLight) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DirectionalLight_stagedOrder[directionallight]; ok {
		return order
	}
	if order, ok := stage.DirectionalLights_referenceOrder[directionallight]; ok {
		return order
	} else {
		log.Printf("instance %p of type DirectionalLight was not staged and does not have a reference order", directionallight)
		return 0
	}
}

func (extrudegeometry *ExtrudeGeometry) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ExtrudeGeometry_stagedOrder[extrudegeometry]; ok {
		return order
	}
	if order, ok := stage.ExtrudeGeometrys_referenceOrder[extrudegeometry]; ok {
		return order
	} else {
		log.Printf("instance %p of type ExtrudeGeometry was not staged and does not have a reference order", extrudegeometry)
		return 0
	}
}

func (mesh *Mesh) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Mesh_stagedOrder[mesh]; ok {
		return order
	}
	if order, ok := stage.Meshs_referenceOrder[mesh]; ok {
		return order
	} else {
		log.Printf("instance %p of type Mesh was not staged and does not have a reference order", mesh)
		return 0
	}
}

func (meshmaterialbasic *MeshMaterialBasic) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.MeshMaterialBasic_stagedOrder[meshmaterialbasic]; ok {
		return order
	}
	if order, ok := stage.MeshMaterialBasics_referenceOrder[meshmaterialbasic]; ok {
		return order
	} else {
		log.Printf("instance %p of type MeshMaterialBasic was not staged and does not have a reference order", meshmaterialbasic)
		return 0
	}
}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.MeshPhysicalMaterial_stagedOrder[meshphysicalmaterial]; ok {
		return order
	}
	if order, ok := stage.MeshPhysicalMaterials_referenceOrder[meshphysicalmaterial]; ok {
		return order
	} else {
		log.Printf("instance %p of type MeshPhysicalMaterial was not staged and does not have a reference order", meshphysicalmaterial)
		return 0
	}
}

func (planegeometry *PlaneGeometry) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PlaneGeometry_stagedOrder[planegeometry]; ok {
		return order
	}
	if order, ok := stage.PlaneGeometrys_referenceOrder[planegeometry]; ok {
		return order
	} else {
		log.Printf("instance %p of type PlaneGeometry was not staged and does not have a reference order", planegeometry)
		return 0
	}
}

func (shape *Shape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Shape_stagedOrder[shape]; ok {
		return order
	}
	if order, ok := stage.Shapes_referenceOrder[shape]; ok {
		return order
	} else {
		log.Printf("instance %p of type Shape was not staged and does not have a reference order", shape)
		return 0
	}
}

func (spheregeometry *SphereGeometry) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SphereGeometry_stagedOrder[spheregeometry]; ok {
		return order
	}
	if order, ok := stage.SphereGeometrys_referenceOrder[spheregeometry]; ok {
		return order
	} else {
		log.Printf("instance %p of type SphereGeometry was not staged and does not have a reference order", spheregeometry)
		return 0
	}
}

func (torusgeometry *TorusGeometry) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TorusGeometry_stagedOrder[torusgeometry]; ok {
		return order
	}
	if order, ok := stage.TorusGeometrys_referenceOrder[torusgeometry]; ok {
		return order
	} else {
		log.Printf("instance %p of type TorusGeometry was not staged and does not have a reference order", torusgeometry)
		return 0
	}
}

func (triangle *Triangle) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Triangle_stagedOrder[triangle]; ok {
		return order
	}
	if order, ok := stage.Triangles_referenceOrder[triangle]; ok {
		return order
	} else {
		log.Printf("instance %p of type Triangle was not staged and does not have a reference order", triangle)
		return 0
	}
}

func (tubegeometry *TubeGeometry) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TubeGeometry_stagedOrder[tubegeometry]; ok {
		return order
	}
	if order, ok := stage.TubeGeometrys_referenceOrder[tubegeometry]; ok {
		return order
	} else {
		log.Printf("instance %p of type TubeGeometry was not staged and does not have a reference order", tubegeometry)
		return 0
	}
}

func (vector2 *Vector2) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Vector2_stagedOrder[vector2]; ok {
		return order
	}
	if order, ok := stage.Vector2s_referenceOrder[vector2]; ok {
		return order
	} else {
		log.Printf("instance %p of type Vector2 was not staged and does not have a reference order", vector2)
		return 0
	}
}

func (vector3 *Vector3) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Vector3_stagedOrder[vector3]; ok {
		return order
	}
	if order, ok := stage.Vector3s_referenceOrder[vector3]; ok {
		return order
	} else {
		log.Printf("instance %p of type Vector3 was not staged and does not have a reference order", vector3)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (ambiantlight *AmbiantLight) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", ambiantlight.GongGetGongstructName(), ambiantlight.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ambiantlight *AmbiantLight) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", ambiantlight.GongGetGongstructName(), ambiantlight.GongGetOrder(stage))
}

func (boxgeometry *BoxGeometry) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", boxgeometry.GongGetGongstructName(), boxgeometry.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (boxgeometry *BoxGeometry) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", boxgeometry.GongGetGongstructName(), boxgeometry.GongGetOrder(stage))
}

func (buffergeometry *BufferGeometry) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", buffergeometry.GongGetGongstructName(), buffergeometry.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (buffergeometry *BufferGeometry) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", buffergeometry.GongGetGongstructName(), buffergeometry.GongGetOrder(stage))
}

func (camera *Camera) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", camera.GongGetGongstructName(), camera.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (camera *Camera) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", camera.GongGetGongstructName(), camera.GongGetOrder(stage))
}

func (canvas *Canvas) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", canvas.GongGetGongstructName(), canvas.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (canvas *Canvas) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", canvas.GongGetGongstructName(), canvas.GongGetOrder(stage))
}

func (curve *Curve) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", curve.GongGetGongstructName(), curve.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (curve *Curve) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", curve.GongGetGongstructName(), curve.GongGetOrder(stage))
}

func (cylindergeometry *CylinderGeometry) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cylindergeometry.GongGetGongstructName(), cylindergeometry.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cylindergeometry *CylinderGeometry) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cylindergeometry.GongGetGongstructName(), cylindergeometry.GongGetOrder(stage))
}

func (directionallight *DirectionalLight) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", directionallight.GongGetGongstructName(), directionallight.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (directionallight *DirectionalLight) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", directionallight.GongGetGongstructName(), directionallight.GongGetOrder(stage))
}

func (extrudegeometry *ExtrudeGeometry) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", extrudegeometry.GongGetGongstructName(), extrudegeometry.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (extrudegeometry *ExtrudeGeometry) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", extrudegeometry.GongGetGongstructName(), extrudegeometry.GongGetOrder(stage))
}

func (mesh *Mesh) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", mesh.GongGetGongstructName(), mesh.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (mesh *Mesh) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", mesh.GongGetGongstructName(), mesh.GongGetOrder(stage))
}

func (meshmaterialbasic *MeshMaterialBasic) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", meshmaterialbasic.GongGetGongstructName(), meshmaterialbasic.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (meshmaterialbasic *MeshMaterialBasic) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", meshmaterialbasic.GongGetGongstructName(), meshmaterialbasic.GongGetOrder(stage))
}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", meshphysicalmaterial.GongGetGongstructName(), meshphysicalmaterial.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (meshphysicalmaterial *MeshPhysicalMaterial) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", meshphysicalmaterial.GongGetGongstructName(), meshphysicalmaterial.GongGetOrder(stage))
}

func (planegeometry *PlaneGeometry) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", planegeometry.GongGetGongstructName(), planegeometry.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (planegeometry *PlaneGeometry) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", planegeometry.GongGetGongstructName(), planegeometry.GongGetOrder(stage))
}

func (shape *Shape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shape.GongGetGongstructName(), shape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shape *Shape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", shape.GongGetGongstructName(), shape.GongGetOrder(stage))
}

func (spheregeometry *SphereGeometry) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spheregeometry.GongGetGongstructName(), spheregeometry.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spheregeometry *SphereGeometry) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spheregeometry.GongGetGongstructName(), spheregeometry.GongGetOrder(stage))
}

func (torusgeometry *TorusGeometry) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", torusgeometry.GongGetGongstructName(), torusgeometry.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (torusgeometry *TorusGeometry) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", torusgeometry.GongGetGongstructName(), torusgeometry.GongGetOrder(stage))
}

func (triangle *Triangle) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", triangle.GongGetGongstructName(), triangle.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (triangle *Triangle) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", triangle.GongGetGongstructName(), triangle.GongGetOrder(stage))
}

func (tubegeometry *TubeGeometry) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tubegeometry.GongGetGongstructName(), tubegeometry.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tubegeometry *TubeGeometry) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tubegeometry.GongGetGongstructName(), tubegeometry.GongGetOrder(stage))
}

func (vector2 *Vector2) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", vector2.GongGetGongstructName(), vector2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (vector2 *Vector2) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", vector2.GongGetGongstructName(), vector2.GongGetOrder(stage))
}

func (vector3 *Vector3) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", vector3.GongGetGongstructName(), vector3.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (vector3 *Vector3) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", vector3.GongGetGongstructName(), vector3.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (ambiantlight *AmbiantLight) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", ambiantlight.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AmbiantLight")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(ambiantlight.Name))
	return
}

func (boxgeometry *BoxGeometry) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", boxgeometry.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BoxGeometry")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(boxgeometry.Name))
	return
}

func (buffergeometry *BufferGeometry) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", buffergeometry.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BufferGeometry")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(buffergeometry.Name))
	return
}

func (camera *Camera) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", camera.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Camera")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(camera.Name))
	return
}

func (canvas *Canvas) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", canvas.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Canvas")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(canvas.Name))
	return
}

func (curve *Curve) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", curve.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Curve")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(curve.Name))
	return
}

func (cylindergeometry *CylinderGeometry) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cylindergeometry.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CylinderGeometry")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(cylindergeometry.Name))
	return
}

func (directionallight *DirectionalLight) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", directionallight.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DirectionalLight")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(directionallight.Name))
	return
}

func (extrudegeometry *ExtrudeGeometry) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", extrudegeometry.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ExtrudeGeometry")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(extrudegeometry.Name))
	return
}

func (mesh *Mesh) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", mesh.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Mesh")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(mesh.Name))
	return
}

func (meshmaterialbasic *MeshMaterialBasic) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", meshmaterialbasic.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MeshMaterialBasic")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(meshmaterialbasic.Name))
	return
}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", meshphysicalmaterial.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MeshPhysicalMaterial")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(meshphysicalmaterial.Name))
	return
}

func (planegeometry *PlaneGeometry) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", planegeometry.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PlaneGeometry")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(planegeometry.Name))
	return
}

func (shape *Shape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Shape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(shape.Name))
	return
}

func (spheregeometry *SphereGeometry) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spheregeometry.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SphereGeometry")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(spheregeometry.Name))
	return
}

func (torusgeometry *TorusGeometry) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", torusgeometry.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TorusGeometry")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(torusgeometry.Name))
	return
}

func (triangle *Triangle) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", triangle.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Triangle")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(triangle.Name))
	return
}

func (tubegeometry *TubeGeometry) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tubegeometry.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TubeGeometry")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(tubegeometry.Name))
	return
}

func (vector2 *Vector2) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", vector2.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Vector2")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(vector2.Name))
	return
}

func (vector3 *Vector3) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", vector3.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Vector3")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(vector3.Name))
	return
}

// insertion point for unstaging
func (ambiantlight *AmbiantLight) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", ambiantlight.GongGetReferenceIdentifier(stage))
	return
}

func (boxgeometry *BoxGeometry) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", boxgeometry.GongGetReferenceIdentifier(stage))
	return
}

func (buffergeometry *BufferGeometry) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", buffergeometry.GongGetReferenceIdentifier(stage))
	return
}

func (camera *Camera) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", camera.GongGetReferenceIdentifier(stage))
	return
}

func (canvas *Canvas) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", canvas.GongGetReferenceIdentifier(stage))
	return
}

func (curve *Curve) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", curve.GongGetReferenceIdentifier(stage))
	return
}

func (cylindergeometry *CylinderGeometry) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cylindergeometry.GongGetReferenceIdentifier(stage))
	return
}

func (directionallight *DirectionalLight) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", directionallight.GongGetReferenceIdentifier(stage))
	return
}

func (extrudegeometry *ExtrudeGeometry) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", extrudegeometry.GongGetReferenceIdentifier(stage))
	return
}

func (mesh *Mesh) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", mesh.GongGetReferenceIdentifier(stage))
	return
}

func (meshmaterialbasic *MeshMaterialBasic) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", meshmaterialbasic.GongGetReferenceIdentifier(stage))
	return
}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", meshphysicalmaterial.GongGetReferenceIdentifier(stage))
	return
}

func (planegeometry *PlaneGeometry) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", planegeometry.GongGetReferenceIdentifier(stage))
	return
}

func (shape *Shape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shape.GongGetReferenceIdentifier(stage))
	return
}

func (spheregeometry *SphereGeometry) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spheregeometry.GongGetReferenceIdentifier(stage))
	return
}

func (torusgeometry *TorusGeometry) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", torusgeometry.GongGetReferenceIdentifier(stage))
	return
}

func (triangle *Triangle) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", triangle.GongGetReferenceIdentifier(stage))
	return
}

func (tubegeometry *TubeGeometry) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tubegeometry.GongGetReferenceIdentifier(stage))
	return
}

func (vector2 *Vector2) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", vector2.GongGetReferenceIdentifier(stage))
	return
}

func (vector3 *Vector3) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", vector3.GongGetReferenceIdentifier(stage))
	return
}

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

// end of template
