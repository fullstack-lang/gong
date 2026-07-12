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
	ambiantlight.CopyBasicFields(newInstance)
	return newInstance
}

func (boxgeometry *BoxGeometry) GongCopy() GongstructIF {
	newInstance := new(BoxGeometry)
	boxgeometry.CopyBasicFields(newInstance)
	return newInstance
}

func (camera *Camera) GongCopy() GongstructIF {
	newInstance := new(Camera)
	camera.CopyBasicFields(newInstance)
	return newInstance
}

func (canvas *Canvas) GongCopy() GongstructIF {
	newInstance := new(Canvas)
	canvas.CopyBasicFields(newInstance)
	return newInstance
}

func (curve *Curve) GongCopy() GongstructIF {
	newInstance := new(Curve)
	curve.CopyBasicFields(newInstance)
	return newInstance
}

func (cylindergeometry *CylinderGeometry) GongCopy() GongstructIF {
	newInstance := new(CylinderGeometry)
	cylindergeometry.CopyBasicFields(newInstance)
	return newInstance
}

func (directionallight *DirectionalLight) GongCopy() GongstructIF {
	newInstance := new(DirectionalLight)
	directionallight.CopyBasicFields(newInstance)
	return newInstance
}

func (extrudegeometry *ExtrudeGeometry) GongCopy() GongstructIF {
	newInstance := new(ExtrudeGeometry)
	extrudegeometry.CopyBasicFields(newInstance)
	return newInstance
}

func (mesh *Mesh) GongCopy() GongstructIF {
	newInstance := new(Mesh)
	mesh.CopyBasicFields(newInstance)
	return newInstance
}

func (meshmaterialbasic *MeshMaterialBasic) GongCopy() GongstructIF {
	newInstance := new(MeshMaterialBasic)
	meshmaterialbasic.CopyBasicFields(newInstance)
	return newInstance
}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongCopy() GongstructIF {
	newInstance := new(MeshPhysicalMaterial)
	meshphysicalmaterial.CopyBasicFields(newInstance)
	return newInstance
}

func (planegeometry *PlaneGeometry) GongCopy() GongstructIF {
	newInstance := new(PlaneGeometry)
	planegeometry.CopyBasicFields(newInstance)
	return newInstance
}

func (shape *Shape) GongCopy() GongstructIF {
	newInstance := new(Shape)
	shape.CopyBasicFields(newInstance)
	return newInstance
}

func (spheregeometry *SphereGeometry) GongCopy() GongstructIF {
	newInstance := new(SphereGeometry)
	spheregeometry.CopyBasicFields(newInstance)
	return newInstance
}

func (torusgeometry *TorusGeometry) GongCopy() GongstructIF {
	newInstance := new(TorusGeometry)
	torusgeometry.CopyBasicFields(newInstance)
	return newInstance
}

func (tubegeometry *TubeGeometry) GongCopy() GongstructIF {
	newInstance := new(TubeGeometry)
	tubegeometry.CopyBasicFields(newInstance)
	return newInstance
}

func (vector2 *Vector2) GongCopy() GongstructIF {
	newInstance := new(Vector2)
	vector2.CopyBasicFields(newInstance)
	return newInstance
}

func (vector3 *Vector3) GongCopy() GongstructIF {
	newInstance := new(Vector3)
	vector3.CopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (ambiantlight *AmbiantLight) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(ambiantlight).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(ambiantlight), uint64(GetOrderPointerGongstruct(stage, ambiantlight)))
	return
}

func (boxgeometry *BoxGeometry) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(boxgeometry).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(boxgeometry), uint64(GetOrderPointerGongstruct(stage, boxgeometry)))
	return
}

func (camera *Camera) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(camera).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(camera), uint64(GetOrderPointerGongstruct(stage, camera)))
	return
}

func (canvas *Canvas) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(canvas).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(canvas), uint64(GetOrderPointerGongstruct(stage, canvas)))
	return
}

func (curve *Curve) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(curve).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(curve), uint64(GetOrderPointerGongstruct(stage, curve)))
	return
}

func (cylindergeometry *CylinderGeometry) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(cylindergeometry).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(cylindergeometry), uint64(GetOrderPointerGongstruct(stage, cylindergeometry)))
	return
}

func (directionallight *DirectionalLight) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(directionallight).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(directionallight), uint64(GetOrderPointerGongstruct(stage, directionallight)))
	return
}

func (extrudegeometry *ExtrudeGeometry) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(extrudegeometry).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(extrudegeometry), uint64(GetOrderPointerGongstruct(stage, extrudegeometry)))
	return
}

func (mesh *Mesh) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(mesh).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(mesh), uint64(GetOrderPointerGongstruct(stage, mesh)))
	return
}

func (meshmaterialbasic *MeshMaterialBasic) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(meshmaterialbasic).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(meshmaterialbasic), uint64(GetOrderPointerGongstruct(stage, meshmaterialbasic)))
	return
}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(meshphysicalmaterial).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(meshphysicalmaterial), uint64(GetOrderPointerGongstruct(stage, meshphysicalmaterial)))
	return
}

func (planegeometry *PlaneGeometry) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(planegeometry).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(planegeometry), uint64(GetOrderPointerGongstruct(stage, planegeometry)))
	return
}

func (shape *Shape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(shape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(shape), uint64(GetOrderPointerGongstruct(stage, shape)))
	return
}

func (spheregeometry *SphereGeometry) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(spheregeometry).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(spheregeometry), uint64(GetOrderPointerGongstruct(stage, spheregeometry)))
	return
}

func (torusgeometry *TorusGeometry) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(torusgeometry).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(torusgeometry), uint64(GetOrderPointerGongstruct(stage, torusgeometry)))
	return
}

func (tubegeometry *TubeGeometry) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(tubegeometry).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(tubegeometry), uint64(GetOrderPointerGongstruct(stage, tubegeometry)))
	return
}

func (vector2 *Vector2) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(vector2).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(vector2), uint64(GetOrderPointerGongstruct(stage, vector2)))
	return
}

func (vector3 *Vector3) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(vector3).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(vector3), uint64(GetOrderPointerGongstruct(stage, vector3)))
	return
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
	var ambiantlights_newInstances []*AmbiantLight
	var ambiantlights_deletedInstances []*AmbiantLight

	// parse all staged instances and check if they have a reference
	for ambiantlight := range stage.AmbiantLights {
		if ref, ok := stage.AmbiantLights_reference[ambiantlight]; !ok {
			ambiantlights_newInstances = append(ambiantlights_newInstances, ambiantlight)
			newInstancesSlice = append(newInstancesSlice, ambiantlight.GongMarshallIdentifier(stage))
			if stage.AmbiantLights_referenceOrder == nil {
				stage.AmbiantLights_referenceOrder = make(map[*AmbiantLight]uint)
			}
			stage.AmbiantLights_referenceOrder[ambiantlight] = stage.AmbiantLight_stagedOrder[ambiantlight]
			newInstancesReverseSlice = append(newInstancesReverseSlice, ambiantlight.GongMarshallUnstaging(stage))
			// delete(stage.AmbiantLights_referenceOrder, ambiantlight)
			fieldInitializers, pointersInitializations := ambiantlight.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.AmbiantLight_stagedOrder[ref] = stage.AmbiantLight_stagedOrder[ambiantlight]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := ambiantlight.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, ambiantlight)
			// delete(stage.AmbiantLight_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if ambiantlight.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", ambiantlight.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.AmbiantLights_reference {
		instance := stage.AmbiantLights_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.AmbiantLights[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			ambiantlights_deletedInstances = append(ambiantlights_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(ambiantlights_newInstances)
	lenDeletedInstances += len(ambiantlights_deletedInstances)
	var boxgeometrys_newInstances []*BoxGeometry
	var boxgeometrys_deletedInstances []*BoxGeometry

	// parse all staged instances and check if they have a reference
	for boxgeometry := range stage.BoxGeometrys {
		if ref, ok := stage.BoxGeometrys_reference[boxgeometry]; !ok {
			boxgeometrys_newInstances = append(boxgeometrys_newInstances, boxgeometry)
			newInstancesSlice = append(newInstancesSlice, boxgeometry.GongMarshallIdentifier(stage))
			if stage.BoxGeometrys_referenceOrder == nil {
				stage.BoxGeometrys_referenceOrder = make(map[*BoxGeometry]uint)
			}
			stage.BoxGeometrys_referenceOrder[boxgeometry] = stage.BoxGeometry_stagedOrder[boxgeometry]
			newInstancesReverseSlice = append(newInstancesReverseSlice, boxgeometry.GongMarshallUnstaging(stage))
			// delete(stage.BoxGeometrys_referenceOrder, boxgeometry)
			fieldInitializers, pointersInitializations := boxgeometry.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.BoxGeometry_stagedOrder[ref] = stage.BoxGeometry_stagedOrder[boxgeometry]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := boxgeometry.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, boxgeometry)
			// delete(stage.BoxGeometry_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if boxgeometry.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", boxgeometry.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.BoxGeometrys_reference {
		instance := stage.BoxGeometrys_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.BoxGeometrys[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			boxgeometrys_deletedInstances = append(boxgeometrys_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(boxgeometrys_newInstances)
	lenDeletedInstances += len(boxgeometrys_deletedInstances)
	var cameras_newInstances []*Camera
	var cameras_deletedInstances []*Camera

	// parse all staged instances and check if they have a reference
	for camera := range stage.Cameras {
		if ref, ok := stage.Cameras_reference[camera]; !ok {
			cameras_newInstances = append(cameras_newInstances, camera)
			newInstancesSlice = append(newInstancesSlice, camera.GongMarshallIdentifier(stage))
			if stage.Cameras_referenceOrder == nil {
				stage.Cameras_referenceOrder = make(map[*Camera]uint)
			}
			stage.Cameras_referenceOrder[camera] = stage.Camera_stagedOrder[camera]
			newInstancesReverseSlice = append(newInstancesReverseSlice, camera.GongMarshallUnstaging(stage))
			// delete(stage.Cameras_referenceOrder, camera)
			fieldInitializers, pointersInitializations := camera.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Camera_stagedOrder[ref] = stage.Camera_stagedOrder[camera]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := camera.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, camera)
			// delete(stage.Camera_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if camera.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", camera.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Cameras_reference {
		instance := stage.Cameras_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Cameras[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			cameras_deletedInstances = append(cameras_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(cameras_newInstances)
	lenDeletedInstances += len(cameras_deletedInstances)
	var canvass_newInstances []*Canvas
	var canvass_deletedInstances []*Canvas

	// parse all staged instances and check if they have a reference
	for canvas := range stage.Canvass {
		if ref, ok := stage.Canvass_reference[canvas]; !ok {
			canvass_newInstances = append(canvass_newInstances, canvas)
			newInstancesSlice = append(newInstancesSlice, canvas.GongMarshallIdentifier(stage))
			if stage.Canvass_referenceOrder == nil {
				stage.Canvass_referenceOrder = make(map[*Canvas]uint)
			}
			stage.Canvass_referenceOrder[canvas] = stage.Canvas_stagedOrder[canvas]
			newInstancesReverseSlice = append(newInstancesReverseSlice, canvas.GongMarshallUnstaging(stage))
			// delete(stage.Canvass_referenceOrder, canvas)
			fieldInitializers, pointersInitializations := canvas.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Canvas_stagedOrder[ref] = stage.Canvas_stagedOrder[canvas]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := canvas.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, canvas)
			// delete(stage.Canvas_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if canvas.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", canvas.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Canvass_reference {
		instance := stage.Canvass_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Canvass[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			canvass_deletedInstances = append(canvass_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(canvass_newInstances)
	lenDeletedInstances += len(canvass_deletedInstances)
	var curves_newInstances []*Curve
	var curves_deletedInstances []*Curve

	// parse all staged instances and check if they have a reference
	for curve := range stage.Curves {
		if ref, ok := stage.Curves_reference[curve]; !ok {
			curves_newInstances = append(curves_newInstances, curve)
			newInstancesSlice = append(newInstancesSlice, curve.GongMarshallIdentifier(stage))
			if stage.Curves_referenceOrder == nil {
				stage.Curves_referenceOrder = make(map[*Curve]uint)
			}
			stage.Curves_referenceOrder[curve] = stage.Curve_stagedOrder[curve]
			newInstancesReverseSlice = append(newInstancesReverseSlice, curve.GongMarshallUnstaging(stage))
			// delete(stage.Curves_referenceOrder, curve)
			fieldInitializers, pointersInitializations := curve.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Curve_stagedOrder[ref] = stage.Curve_stagedOrder[curve]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := curve.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, curve)
			// delete(stage.Curve_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if curve.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", curve.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Curves_reference {
		instance := stage.Curves_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Curves[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			curves_deletedInstances = append(curves_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(curves_newInstances)
	lenDeletedInstances += len(curves_deletedInstances)
	var cylindergeometrys_newInstances []*CylinderGeometry
	var cylindergeometrys_deletedInstances []*CylinderGeometry

	// parse all staged instances and check if they have a reference
	for cylindergeometry := range stage.CylinderGeometrys {
		if ref, ok := stage.CylinderGeometrys_reference[cylindergeometry]; !ok {
			cylindergeometrys_newInstances = append(cylindergeometrys_newInstances, cylindergeometry)
			newInstancesSlice = append(newInstancesSlice, cylindergeometry.GongMarshallIdentifier(stage))
			if stage.CylinderGeometrys_referenceOrder == nil {
				stage.CylinderGeometrys_referenceOrder = make(map[*CylinderGeometry]uint)
			}
			stage.CylinderGeometrys_referenceOrder[cylindergeometry] = stage.CylinderGeometry_stagedOrder[cylindergeometry]
			newInstancesReverseSlice = append(newInstancesReverseSlice, cylindergeometry.GongMarshallUnstaging(stage))
			// delete(stage.CylinderGeometrys_referenceOrder, cylindergeometry)
			fieldInitializers, pointersInitializations := cylindergeometry.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.CylinderGeometry_stagedOrder[ref] = stage.CylinderGeometry_stagedOrder[cylindergeometry]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := cylindergeometry.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, cylindergeometry)
			// delete(stage.CylinderGeometry_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if cylindergeometry.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", cylindergeometry.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.CylinderGeometrys_reference {
		instance := stage.CylinderGeometrys_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.CylinderGeometrys[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			cylindergeometrys_deletedInstances = append(cylindergeometrys_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(cylindergeometrys_newInstances)
	lenDeletedInstances += len(cylindergeometrys_deletedInstances)
	var directionallights_newInstances []*DirectionalLight
	var directionallights_deletedInstances []*DirectionalLight

	// parse all staged instances and check if they have a reference
	for directionallight := range stage.DirectionalLights {
		if ref, ok := stage.DirectionalLights_reference[directionallight]; !ok {
			directionallights_newInstances = append(directionallights_newInstances, directionallight)
			newInstancesSlice = append(newInstancesSlice, directionallight.GongMarshallIdentifier(stage))
			if stage.DirectionalLights_referenceOrder == nil {
				stage.DirectionalLights_referenceOrder = make(map[*DirectionalLight]uint)
			}
			stage.DirectionalLights_referenceOrder[directionallight] = stage.DirectionalLight_stagedOrder[directionallight]
			newInstancesReverseSlice = append(newInstancesReverseSlice, directionallight.GongMarshallUnstaging(stage))
			// delete(stage.DirectionalLights_referenceOrder, directionallight)
			fieldInitializers, pointersInitializations := directionallight.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DirectionalLight_stagedOrder[ref] = stage.DirectionalLight_stagedOrder[directionallight]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := directionallight.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, directionallight)
			// delete(stage.DirectionalLight_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if directionallight.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", directionallight.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.DirectionalLights_reference {
		instance := stage.DirectionalLights_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.DirectionalLights[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			directionallights_deletedInstances = append(directionallights_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(directionallights_newInstances)
	lenDeletedInstances += len(directionallights_deletedInstances)
	var extrudegeometrys_newInstances []*ExtrudeGeometry
	var extrudegeometrys_deletedInstances []*ExtrudeGeometry

	// parse all staged instances and check if they have a reference
	for extrudegeometry := range stage.ExtrudeGeometrys {
		if ref, ok := stage.ExtrudeGeometrys_reference[extrudegeometry]; !ok {
			extrudegeometrys_newInstances = append(extrudegeometrys_newInstances, extrudegeometry)
			newInstancesSlice = append(newInstancesSlice, extrudegeometry.GongMarshallIdentifier(stage))
			if stage.ExtrudeGeometrys_referenceOrder == nil {
				stage.ExtrudeGeometrys_referenceOrder = make(map[*ExtrudeGeometry]uint)
			}
			stage.ExtrudeGeometrys_referenceOrder[extrudegeometry] = stage.ExtrudeGeometry_stagedOrder[extrudegeometry]
			newInstancesReverseSlice = append(newInstancesReverseSlice, extrudegeometry.GongMarshallUnstaging(stage))
			// delete(stage.ExtrudeGeometrys_referenceOrder, extrudegeometry)
			fieldInitializers, pointersInitializations := extrudegeometry.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ExtrudeGeometry_stagedOrder[ref] = stage.ExtrudeGeometry_stagedOrder[extrudegeometry]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := extrudegeometry.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, extrudegeometry)
			// delete(stage.ExtrudeGeometry_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if extrudegeometry.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", extrudegeometry.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ExtrudeGeometrys_reference {
		instance := stage.ExtrudeGeometrys_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ExtrudeGeometrys[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			extrudegeometrys_deletedInstances = append(extrudegeometrys_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(extrudegeometrys_newInstances)
	lenDeletedInstances += len(extrudegeometrys_deletedInstances)
	var meshs_newInstances []*Mesh
	var meshs_deletedInstances []*Mesh

	// parse all staged instances and check if they have a reference
	for mesh := range stage.Meshs {
		if ref, ok := stage.Meshs_reference[mesh]; !ok {
			meshs_newInstances = append(meshs_newInstances, mesh)
			newInstancesSlice = append(newInstancesSlice, mesh.GongMarshallIdentifier(stage))
			if stage.Meshs_referenceOrder == nil {
				stage.Meshs_referenceOrder = make(map[*Mesh]uint)
			}
			stage.Meshs_referenceOrder[mesh] = stage.Mesh_stagedOrder[mesh]
			newInstancesReverseSlice = append(newInstancesReverseSlice, mesh.GongMarshallUnstaging(stage))
			// delete(stage.Meshs_referenceOrder, mesh)
			fieldInitializers, pointersInitializations := mesh.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Mesh_stagedOrder[ref] = stage.Mesh_stagedOrder[mesh]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := mesh.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, mesh)
			// delete(stage.Mesh_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if mesh.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", mesh.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Meshs_reference {
		instance := stage.Meshs_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Meshs[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			meshs_deletedInstances = append(meshs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(meshs_newInstances)
	lenDeletedInstances += len(meshs_deletedInstances)
	var meshmaterialbasics_newInstances []*MeshMaterialBasic
	var meshmaterialbasics_deletedInstances []*MeshMaterialBasic

	// parse all staged instances and check if they have a reference
	for meshmaterialbasic := range stage.MeshMaterialBasics {
		if ref, ok := stage.MeshMaterialBasics_reference[meshmaterialbasic]; !ok {
			meshmaterialbasics_newInstances = append(meshmaterialbasics_newInstances, meshmaterialbasic)
			newInstancesSlice = append(newInstancesSlice, meshmaterialbasic.GongMarshallIdentifier(stage))
			if stage.MeshMaterialBasics_referenceOrder == nil {
				stage.MeshMaterialBasics_referenceOrder = make(map[*MeshMaterialBasic]uint)
			}
			stage.MeshMaterialBasics_referenceOrder[meshmaterialbasic] = stage.MeshMaterialBasic_stagedOrder[meshmaterialbasic]
			newInstancesReverseSlice = append(newInstancesReverseSlice, meshmaterialbasic.GongMarshallUnstaging(stage))
			// delete(stage.MeshMaterialBasics_referenceOrder, meshmaterialbasic)
			fieldInitializers, pointersInitializations := meshmaterialbasic.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.MeshMaterialBasic_stagedOrder[ref] = stage.MeshMaterialBasic_stagedOrder[meshmaterialbasic]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := meshmaterialbasic.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, meshmaterialbasic)
			// delete(stage.MeshMaterialBasic_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if meshmaterialbasic.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", meshmaterialbasic.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.MeshMaterialBasics_reference {
		instance := stage.MeshMaterialBasics_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.MeshMaterialBasics[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			meshmaterialbasics_deletedInstances = append(meshmaterialbasics_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(meshmaterialbasics_newInstances)
	lenDeletedInstances += len(meshmaterialbasics_deletedInstances)
	var meshphysicalmaterials_newInstances []*MeshPhysicalMaterial
	var meshphysicalmaterials_deletedInstances []*MeshPhysicalMaterial

	// parse all staged instances and check if they have a reference
	for meshphysicalmaterial := range stage.MeshPhysicalMaterials {
		if ref, ok := stage.MeshPhysicalMaterials_reference[meshphysicalmaterial]; !ok {
			meshphysicalmaterials_newInstances = append(meshphysicalmaterials_newInstances, meshphysicalmaterial)
			newInstancesSlice = append(newInstancesSlice, meshphysicalmaterial.GongMarshallIdentifier(stage))
			if stage.MeshPhysicalMaterials_referenceOrder == nil {
				stage.MeshPhysicalMaterials_referenceOrder = make(map[*MeshPhysicalMaterial]uint)
			}
			stage.MeshPhysicalMaterials_referenceOrder[meshphysicalmaterial] = stage.MeshPhysicalMaterial_stagedOrder[meshphysicalmaterial]
			newInstancesReverseSlice = append(newInstancesReverseSlice, meshphysicalmaterial.GongMarshallUnstaging(stage))
			// delete(stage.MeshPhysicalMaterials_referenceOrder, meshphysicalmaterial)
			fieldInitializers, pointersInitializations := meshphysicalmaterial.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.MeshPhysicalMaterial_stagedOrder[ref] = stage.MeshPhysicalMaterial_stagedOrder[meshphysicalmaterial]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := meshphysicalmaterial.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, meshphysicalmaterial)
			// delete(stage.MeshPhysicalMaterial_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if meshphysicalmaterial.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", meshphysicalmaterial.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.MeshPhysicalMaterials_reference {
		instance := stage.MeshPhysicalMaterials_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.MeshPhysicalMaterials[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			meshphysicalmaterials_deletedInstances = append(meshphysicalmaterials_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(meshphysicalmaterials_newInstances)
	lenDeletedInstances += len(meshphysicalmaterials_deletedInstances)
	var planegeometrys_newInstances []*PlaneGeometry
	var planegeometrys_deletedInstances []*PlaneGeometry

	// parse all staged instances and check if they have a reference
	for planegeometry := range stage.PlaneGeometrys {
		if ref, ok := stage.PlaneGeometrys_reference[planegeometry]; !ok {
			planegeometrys_newInstances = append(planegeometrys_newInstances, planegeometry)
			newInstancesSlice = append(newInstancesSlice, planegeometry.GongMarshallIdentifier(stage))
			if stage.PlaneGeometrys_referenceOrder == nil {
				stage.PlaneGeometrys_referenceOrder = make(map[*PlaneGeometry]uint)
			}
			stage.PlaneGeometrys_referenceOrder[planegeometry] = stage.PlaneGeometry_stagedOrder[planegeometry]
			newInstancesReverseSlice = append(newInstancesReverseSlice, planegeometry.GongMarshallUnstaging(stage))
			// delete(stage.PlaneGeometrys_referenceOrder, planegeometry)
			fieldInitializers, pointersInitializations := planegeometry.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PlaneGeometry_stagedOrder[ref] = stage.PlaneGeometry_stagedOrder[planegeometry]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := planegeometry.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, planegeometry)
			// delete(stage.PlaneGeometry_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if planegeometry.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", planegeometry.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.PlaneGeometrys_reference {
		instance := stage.PlaneGeometrys_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.PlaneGeometrys[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			planegeometrys_deletedInstances = append(planegeometrys_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(planegeometrys_newInstances)
	lenDeletedInstances += len(planegeometrys_deletedInstances)
	var shapes_newInstances []*Shape
	var shapes_deletedInstances []*Shape

	// parse all staged instances and check if they have a reference
	for shape := range stage.Shapes {
		if ref, ok := stage.Shapes_reference[shape]; !ok {
			shapes_newInstances = append(shapes_newInstances, shape)
			newInstancesSlice = append(newInstancesSlice, shape.GongMarshallIdentifier(stage))
			if stage.Shapes_referenceOrder == nil {
				stage.Shapes_referenceOrder = make(map[*Shape]uint)
			}
			stage.Shapes_referenceOrder[shape] = stage.Shape_stagedOrder[shape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, shape.GongMarshallUnstaging(stage))
			// delete(stage.Shapes_referenceOrder, shape)
			fieldInitializers, pointersInitializations := shape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Shape_stagedOrder[ref] = stage.Shape_stagedOrder[shape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := shape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, shape)
			// delete(stage.Shape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if shape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", shape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Shapes_reference {
		instance := stage.Shapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Shapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			shapes_deletedInstances = append(shapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(shapes_newInstances)
	lenDeletedInstances += len(shapes_deletedInstances)
	var spheregeometrys_newInstances []*SphereGeometry
	var spheregeometrys_deletedInstances []*SphereGeometry

	// parse all staged instances and check if they have a reference
	for spheregeometry := range stage.SphereGeometrys {
		if ref, ok := stage.SphereGeometrys_reference[spheregeometry]; !ok {
			spheregeometrys_newInstances = append(spheregeometrys_newInstances, spheregeometry)
			newInstancesSlice = append(newInstancesSlice, spheregeometry.GongMarshallIdentifier(stage))
			if stage.SphereGeometrys_referenceOrder == nil {
				stage.SphereGeometrys_referenceOrder = make(map[*SphereGeometry]uint)
			}
			stage.SphereGeometrys_referenceOrder[spheregeometry] = stage.SphereGeometry_stagedOrder[spheregeometry]
			newInstancesReverseSlice = append(newInstancesReverseSlice, spheregeometry.GongMarshallUnstaging(stage))
			// delete(stage.SphereGeometrys_referenceOrder, spheregeometry)
			fieldInitializers, pointersInitializations := spheregeometry.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SphereGeometry_stagedOrder[ref] = stage.SphereGeometry_stagedOrder[spheregeometry]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := spheregeometry.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, spheregeometry)
			// delete(stage.SphereGeometry_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if spheregeometry.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", spheregeometry.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.SphereGeometrys_reference {
		instance := stage.SphereGeometrys_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.SphereGeometrys[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			spheregeometrys_deletedInstances = append(spheregeometrys_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(spheregeometrys_newInstances)
	lenDeletedInstances += len(spheregeometrys_deletedInstances)
	var torusgeometrys_newInstances []*TorusGeometry
	var torusgeometrys_deletedInstances []*TorusGeometry

	// parse all staged instances and check if they have a reference
	for torusgeometry := range stage.TorusGeometrys {
		if ref, ok := stage.TorusGeometrys_reference[torusgeometry]; !ok {
			torusgeometrys_newInstances = append(torusgeometrys_newInstances, torusgeometry)
			newInstancesSlice = append(newInstancesSlice, torusgeometry.GongMarshallIdentifier(stage))
			if stage.TorusGeometrys_referenceOrder == nil {
				stage.TorusGeometrys_referenceOrder = make(map[*TorusGeometry]uint)
			}
			stage.TorusGeometrys_referenceOrder[torusgeometry] = stage.TorusGeometry_stagedOrder[torusgeometry]
			newInstancesReverseSlice = append(newInstancesReverseSlice, torusgeometry.GongMarshallUnstaging(stage))
			// delete(stage.TorusGeometrys_referenceOrder, torusgeometry)
			fieldInitializers, pointersInitializations := torusgeometry.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TorusGeometry_stagedOrder[ref] = stage.TorusGeometry_stagedOrder[torusgeometry]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := torusgeometry.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, torusgeometry)
			// delete(stage.TorusGeometry_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if torusgeometry.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", torusgeometry.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.TorusGeometrys_reference {
		instance := stage.TorusGeometrys_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TorusGeometrys[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			torusgeometrys_deletedInstances = append(torusgeometrys_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(torusgeometrys_newInstances)
	lenDeletedInstances += len(torusgeometrys_deletedInstances)
	var tubegeometrys_newInstances []*TubeGeometry
	var tubegeometrys_deletedInstances []*TubeGeometry

	// parse all staged instances and check if they have a reference
	for tubegeometry := range stage.TubeGeometrys {
		if ref, ok := stage.TubeGeometrys_reference[tubegeometry]; !ok {
			tubegeometrys_newInstances = append(tubegeometrys_newInstances, tubegeometry)
			newInstancesSlice = append(newInstancesSlice, tubegeometry.GongMarshallIdentifier(stage))
			if stage.TubeGeometrys_referenceOrder == nil {
				stage.TubeGeometrys_referenceOrder = make(map[*TubeGeometry]uint)
			}
			stage.TubeGeometrys_referenceOrder[tubegeometry] = stage.TubeGeometry_stagedOrder[tubegeometry]
			newInstancesReverseSlice = append(newInstancesReverseSlice, tubegeometry.GongMarshallUnstaging(stage))
			// delete(stage.TubeGeometrys_referenceOrder, tubegeometry)
			fieldInitializers, pointersInitializations := tubegeometry.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TubeGeometry_stagedOrder[ref] = stage.TubeGeometry_stagedOrder[tubegeometry]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := tubegeometry.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, tubegeometry)
			// delete(stage.TubeGeometry_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if tubegeometry.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", tubegeometry.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.TubeGeometrys_reference {
		instance := stage.TubeGeometrys_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TubeGeometrys[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			tubegeometrys_deletedInstances = append(tubegeometrys_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(tubegeometrys_newInstances)
	lenDeletedInstances += len(tubegeometrys_deletedInstances)
	var vector2s_newInstances []*Vector2
	var vector2s_deletedInstances []*Vector2

	// parse all staged instances and check if they have a reference
	for vector2 := range stage.Vector2s {
		if ref, ok := stage.Vector2s_reference[vector2]; !ok {
			vector2s_newInstances = append(vector2s_newInstances, vector2)
			newInstancesSlice = append(newInstancesSlice, vector2.GongMarshallIdentifier(stage))
			if stage.Vector2s_referenceOrder == nil {
				stage.Vector2s_referenceOrder = make(map[*Vector2]uint)
			}
			stage.Vector2s_referenceOrder[vector2] = stage.Vector2_stagedOrder[vector2]
			newInstancesReverseSlice = append(newInstancesReverseSlice, vector2.GongMarshallUnstaging(stage))
			// delete(stage.Vector2s_referenceOrder, vector2)
			fieldInitializers, pointersInitializations := vector2.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Vector2_stagedOrder[ref] = stage.Vector2_stagedOrder[vector2]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := vector2.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, vector2)
			// delete(stage.Vector2_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if vector2.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", vector2.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Vector2s_reference {
		instance := stage.Vector2s_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Vector2s[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			vector2s_deletedInstances = append(vector2s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(vector2s_newInstances)
	lenDeletedInstances += len(vector2s_deletedInstances)
	var vector3s_newInstances []*Vector3
	var vector3s_deletedInstances []*Vector3

	// parse all staged instances and check if they have a reference
	for vector3 := range stage.Vector3s {
		if ref, ok := stage.Vector3s_reference[vector3]; !ok {
			vector3s_newInstances = append(vector3s_newInstances, vector3)
			newInstancesSlice = append(newInstancesSlice, vector3.GongMarshallIdentifier(stage))
			if stage.Vector3s_referenceOrder == nil {
				stage.Vector3s_referenceOrder = make(map[*Vector3]uint)
			}
			stage.Vector3s_referenceOrder[vector3] = stage.Vector3_stagedOrder[vector3]
			newInstancesReverseSlice = append(newInstancesReverseSlice, vector3.GongMarshallUnstaging(stage))
			// delete(stage.Vector3s_referenceOrder, vector3)
			fieldInitializers, pointersInitializations := vector3.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Vector3_stagedOrder[ref] = stage.Vector3_stagedOrder[vector3]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := vector3.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, vector3)
			// delete(stage.Vector3_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if vector3.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", vector3.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Vector3s_reference {
		instance := stage.Vector3s_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Vector3s[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			vector3s_deletedInstances = append(vector3s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(vector3s_newInstances)
	lenDeletedInstances += len(vector3s_deletedInstances)

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
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(ambiantlight.Name))
	return
}

func (boxgeometry *BoxGeometry) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", boxgeometry.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BoxGeometry")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(boxgeometry.Name))
	return
}

func (camera *Camera) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", camera.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Camera")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(camera.Name))
	return
}

func (canvas *Canvas) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", canvas.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Canvas")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(canvas.Name))
	return
}

func (curve *Curve) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", curve.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Curve")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(curve.Name))
	return
}

func (cylindergeometry *CylinderGeometry) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cylindergeometry.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CylinderGeometry")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(cylindergeometry.Name))
	return
}

func (directionallight *DirectionalLight) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", directionallight.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DirectionalLight")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(directionallight.Name))
	return
}

func (extrudegeometry *ExtrudeGeometry) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", extrudegeometry.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ExtrudeGeometry")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(extrudegeometry.Name))
	return
}

func (mesh *Mesh) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", mesh.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Mesh")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(mesh.Name))
	return
}

func (meshmaterialbasic *MeshMaterialBasic) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", meshmaterialbasic.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MeshMaterialBasic")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(meshmaterialbasic.Name))
	return
}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", meshphysicalmaterial.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MeshPhysicalMaterial")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(meshphysicalmaterial.Name))
	return
}

func (planegeometry *PlaneGeometry) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", planegeometry.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PlaneGeometry")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(planegeometry.Name))
	return
}

func (shape *Shape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", shape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Shape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(shape.Name))
	return
}

func (spheregeometry *SphereGeometry) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spheregeometry.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SphereGeometry")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(spheregeometry.Name))
	return
}

func (torusgeometry *TorusGeometry) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", torusgeometry.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TorusGeometry")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(torusgeometry.Name))
	return
}

func (tubegeometry *TubeGeometry) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tubegeometry.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TubeGeometry")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(tubegeometry.Name))
	return
}

func (vector2 *Vector2) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", vector2.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Vector2")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(vector2.Name))
	return
}

func (vector3 *Vector3) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", vector3.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Vector3")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(vector3.Name))
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

func IntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += IntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
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
