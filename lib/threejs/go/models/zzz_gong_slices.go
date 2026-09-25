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

	// Compute reverse map for named struct Shape
	// insertion point per field
	stage.Shape_Points_reverseMap = make(map[*Vector2]*Shape)
	for shape := range stage.Shapes {
		_ = shape
		for _, _vector2 := range shape.Points {
			stage.Shape_Points_reverseMap[_vector2] = shape
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.AmbiantLights)

	res = __gong__appendInstances(res, stage.BoxGeometrys)

	res = __gong__appendInstances(res, stage.BufferGeometrys)

	res = __gong__appendInstances(res, stage.Cameras)

	res = __gong__appendInstances(res, stage.Canvass)

	res = __gong__appendInstances(res, stage.Curves)

	res = __gong__appendInstances(res, stage.CylinderGeometrys)

	res = __gong__appendInstances(res, stage.DirectionalLights)

	res = __gong__appendInstances(res, stage.ExtrudeGeometrys)

	res = __gong__appendInstances(res, stage.Meshs)

	res = __gong__appendInstances(res, stage.MeshMaterialBasics)

	res = __gong__appendInstances(res, stage.MeshPhysicalMaterials)

	res = __gong__appendInstances(res, stage.PlaneGeometrys)

	res = __gong__appendInstances(res, stage.Shapes)

	res = __gong__appendInstances(res, stage.SphereGeometrys)

	res = __gong__appendInstances(res, stage.TorusGeometrys)

	res = __gong__appendInstances(res, stage.Triangles)

	res = __gong__appendInstances(res, stage.TubeGeometrys)

	res = __gong__appendInstances(res, stage.Vector2s)

	res = __gong__appendInstances(res, stage.Vector3s)

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
func (ambiantlight *AmbiantLight) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ambiantlight)
}

func (boxgeometry *BoxGeometry) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, boxgeometry)
}

func (buffergeometry *BufferGeometry) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, buffergeometry)
}

func (camera *Camera) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, camera)
}

func (canvas *Canvas) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, canvas)
}

func (curve *Curve) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, curve)
}

func (cylindergeometry *CylinderGeometry) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, cylindergeometry)
}

func (directionallight *DirectionalLight) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, directionallight)
}

func (extrudegeometry *ExtrudeGeometry) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, extrudegeometry)
}

func (mesh *Mesh) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, mesh)
}

func (meshmaterialbasic *MeshMaterialBasic) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, meshmaterialbasic)
}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, meshphysicalmaterial)
}

func (planegeometry *PlaneGeometry) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, planegeometry)
}

func (shape *Shape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, shape)
}

func (spheregeometry *SphereGeometry) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, spheregeometry)
}

func (torusgeometry *TorusGeometry) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, torusgeometry)
}

func (triangle *Triangle) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, triangle)
}

func (tubegeometry *TubeGeometry) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tubegeometry)
}

func (vector2 *Vector2) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, vector2)
}

func (vector3 *Vector3) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, vector3)
}


type GongstructDiffable[T any] interface {
	GongstructPtr
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
	__gong__computeReferencePass1(stage, stage.AmbiantLights, &stage.AmbiantLights_reference, &stage.AmbiantLights_referenceOrder, &stage.AmbiantLights_instance)

	__gong__computeReferencePass1(stage, stage.BoxGeometrys, &stage.BoxGeometrys_reference, &stage.BoxGeometrys_referenceOrder, &stage.BoxGeometrys_instance)

	__gong__computeReferencePass1(stage, stage.BufferGeometrys, &stage.BufferGeometrys_reference, &stage.BufferGeometrys_referenceOrder, &stage.BufferGeometrys_instance)

	__gong__computeReferencePass1(stage, stage.Cameras, &stage.Cameras_reference, &stage.Cameras_referenceOrder, &stage.Cameras_instance)

	__gong__computeReferencePass1(stage, stage.Canvass, &stage.Canvass_reference, &stage.Canvass_referenceOrder, &stage.Canvass_instance)

	__gong__computeReferencePass1(stage, stage.Curves, &stage.Curves_reference, &stage.Curves_referenceOrder, &stage.Curves_instance)

	__gong__computeReferencePass1(stage, stage.CylinderGeometrys, &stage.CylinderGeometrys_reference, &stage.CylinderGeometrys_referenceOrder, &stage.CylinderGeometrys_instance)

	__gong__computeReferencePass1(stage, stage.DirectionalLights, &stage.DirectionalLights_reference, &stage.DirectionalLights_referenceOrder, &stage.DirectionalLights_instance)

	__gong__computeReferencePass1(stage, stage.ExtrudeGeometrys, &stage.ExtrudeGeometrys_reference, &stage.ExtrudeGeometrys_referenceOrder, &stage.ExtrudeGeometrys_instance)

	__gong__computeReferencePass1(stage, stage.Meshs, &stage.Meshs_reference, &stage.Meshs_referenceOrder, &stage.Meshs_instance)

	__gong__computeReferencePass1(stage, stage.MeshMaterialBasics, &stage.MeshMaterialBasics_reference, &stage.MeshMaterialBasics_referenceOrder, &stage.MeshMaterialBasics_instance)

	__gong__computeReferencePass1(stage, stage.MeshPhysicalMaterials, &stage.MeshPhysicalMaterials_reference, &stage.MeshPhysicalMaterials_referenceOrder, &stage.MeshPhysicalMaterials_instance)

	__gong__computeReferencePass1(stage, stage.PlaneGeometrys, &stage.PlaneGeometrys_reference, &stage.PlaneGeometrys_referenceOrder, &stage.PlaneGeometrys_instance)

	__gong__computeReferencePass1(stage, stage.Shapes, &stage.Shapes_reference, &stage.Shapes_referenceOrder, &stage.Shapes_instance)

	__gong__computeReferencePass1(stage, stage.SphereGeometrys, &stage.SphereGeometrys_reference, &stage.SphereGeometrys_referenceOrder, &stage.SphereGeometrys_instance)

	__gong__computeReferencePass1(stage, stage.TorusGeometrys, &stage.TorusGeometrys_reference, &stage.TorusGeometrys_referenceOrder, &stage.TorusGeometrys_instance)

	__gong__computeReferencePass1(stage, stage.Triangles, &stage.Triangles_reference, &stage.Triangles_referenceOrder, &stage.Triangles_instance)

	__gong__computeReferencePass1(stage, stage.TubeGeometrys, &stage.TubeGeometrys_reference, &stage.TubeGeometrys_referenceOrder, &stage.TubeGeometrys_instance)

	__gong__computeReferencePass1(stage, stage.Vector2s, &stage.Vector2s_reference, &stage.Vector2s_referenceOrder, &stage.Vector2s_instance)

	__gong__computeReferencePass1(stage, stage.Vector3s, &stage.Vector3s_reference, &stage.Vector3s_referenceOrder, &stage.Vector3s_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.AmbiantLights, stage.AmbiantLights_reference, stage)

	__gong__computeReferencePass2(stage.BoxGeometrys, stage.BoxGeometrys_reference, stage)

	__gong__computeReferencePass2(stage.BufferGeometrys, stage.BufferGeometrys_reference, stage)

	__gong__computeReferencePass2(stage.Cameras, stage.Cameras_reference, stage)

	__gong__computeReferencePass2(stage.Canvass, stage.Canvass_reference, stage)

	__gong__computeReferencePass2(stage.Curves, stage.Curves_reference, stage)

	__gong__computeReferencePass2(stage.CylinderGeometrys, stage.CylinderGeometrys_reference, stage)

	__gong__computeReferencePass2(stage.DirectionalLights, stage.DirectionalLights_reference, stage)

	__gong__computeReferencePass2(stage.ExtrudeGeometrys, stage.ExtrudeGeometrys_reference, stage)

	__gong__computeReferencePass2(stage.Meshs, stage.Meshs_reference, stage)

	__gong__computeReferencePass2(stage.MeshMaterialBasics, stage.MeshMaterialBasics_reference, stage)

	__gong__computeReferencePass2(stage.MeshPhysicalMaterials, stage.MeshPhysicalMaterials_reference, stage)

	__gong__computeReferencePass2(stage.PlaneGeometrys, stage.PlaneGeometrys_reference, stage)

	__gong__computeReferencePass2(stage.Shapes, stage.Shapes_reference, stage)

	__gong__computeReferencePass2(stage.SphereGeometrys, stage.SphereGeometrys_reference, stage)

	__gong__computeReferencePass2(stage.TorusGeometrys, stage.TorusGeometrys_reference, stage)

	__gong__computeReferencePass2(stage.Triangles, stage.Triangles_reference, stage)

	__gong__computeReferencePass2(stage.TubeGeometrys, stage.TubeGeometrys_reference, stage)

	__gong__computeReferencePass2(stage.Vector2s, stage.Vector2s_reference, stage)

	__gong__computeReferencePass2(stage.Vector3s, stage.Vector3s_reference, stage)

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
	return __gong__getOrder(stage.AmbiantLight_stagedOrder, stage.AmbiantLights_referenceOrder, ambiantlight, "AmbiantLight")
}

func (boxgeometry *BoxGeometry) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.BoxGeometry_stagedOrder, stage.BoxGeometrys_referenceOrder, boxgeometry, "BoxGeometry")
}

func (buffergeometry *BufferGeometry) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.BufferGeometry_stagedOrder, stage.BufferGeometrys_referenceOrder, buffergeometry, "BufferGeometry")
}

func (camera *Camera) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Camera_stagedOrder, stage.Cameras_referenceOrder, camera, "Camera")
}

func (canvas *Canvas) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Canvas_stagedOrder, stage.Canvass_referenceOrder, canvas, "Canvas")
}

func (curve *Curve) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Curve_stagedOrder, stage.Curves_referenceOrder, curve, "Curve")
}

func (cylindergeometry *CylinderGeometry) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CylinderGeometry_stagedOrder, stage.CylinderGeometrys_referenceOrder, cylindergeometry, "CylinderGeometry")
}

func (directionallight *DirectionalLight) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DirectionalLight_stagedOrder, stage.DirectionalLights_referenceOrder, directionallight, "DirectionalLight")
}

func (extrudegeometry *ExtrudeGeometry) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ExtrudeGeometry_stagedOrder, stage.ExtrudeGeometrys_referenceOrder, extrudegeometry, "ExtrudeGeometry")
}

func (mesh *Mesh) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Mesh_stagedOrder, stage.Meshs_referenceOrder, mesh, "Mesh")
}

func (meshmaterialbasic *MeshMaterialBasic) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MeshMaterialBasic_stagedOrder, stage.MeshMaterialBasics_referenceOrder, meshmaterialbasic, "MeshMaterialBasic")
}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MeshPhysicalMaterial_stagedOrder, stage.MeshPhysicalMaterials_referenceOrder, meshphysicalmaterial, "MeshPhysicalMaterial")
}

func (planegeometry *PlaneGeometry) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PlaneGeometry_stagedOrder, stage.PlaneGeometrys_referenceOrder, planegeometry, "PlaneGeometry")
}

func (shape *Shape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Shape_stagedOrder, stage.Shapes_referenceOrder, shape, "Shape")
}

func (spheregeometry *SphereGeometry) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SphereGeometry_stagedOrder, stage.SphereGeometrys_referenceOrder, spheregeometry, "SphereGeometry")
}

func (torusgeometry *TorusGeometry) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TorusGeometry_stagedOrder, stage.TorusGeometrys_referenceOrder, torusgeometry, "TorusGeometry")
}

func (triangle *Triangle) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Triangle_stagedOrder, stage.Triangles_referenceOrder, triangle, "Triangle")
}

func (tubegeometry *TubeGeometry) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TubeGeometry_stagedOrder, stage.TubeGeometrys_referenceOrder, tubegeometry, "TubeGeometry")
}

func (vector2 *Vector2) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Vector2_stagedOrder, stage.Vector2s_referenceOrder, vector2, "Vector2")
}

func (vector3 *Vector3) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Vector3_stagedOrder, stage.Vector3s_referenceOrder, vector3, "Vector3")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (ambiantlight *AmbiantLight) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ambiantlight, ambiantlight.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ambiantlight *AmbiantLight) GongGetReferenceIdentifier(stage *Stage) string {
	return ambiantlight.GongGetIdentifier(stage)
}

func (boxgeometry *BoxGeometry) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(boxgeometry, boxgeometry.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (boxgeometry *BoxGeometry) GongGetReferenceIdentifier(stage *Stage) string {
	return boxgeometry.GongGetIdentifier(stage)
}

func (buffergeometry *BufferGeometry) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(buffergeometry, buffergeometry.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (buffergeometry *BufferGeometry) GongGetReferenceIdentifier(stage *Stage) string {
	return buffergeometry.GongGetIdentifier(stage)
}

func (camera *Camera) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(camera, camera.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (camera *Camera) GongGetReferenceIdentifier(stage *Stage) string {
	return camera.GongGetIdentifier(stage)
}

func (canvas *Canvas) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(canvas, canvas.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (canvas *Canvas) GongGetReferenceIdentifier(stage *Stage) string {
	return canvas.GongGetIdentifier(stage)
}

func (curve *Curve) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(curve, curve.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (curve *Curve) GongGetReferenceIdentifier(stage *Stage) string {
	return curve.GongGetIdentifier(stage)
}

func (cylindergeometry *CylinderGeometry) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(cylindergeometry, cylindergeometry.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cylindergeometry *CylinderGeometry) GongGetReferenceIdentifier(stage *Stage) string {
	return cylindergeometry.GongGetIdentifier(stage)
}

func (directionallight *DirectionalLight) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(directionallight, directionallight.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (directionallight *DirectionalLight) GongGetReferenceIdentifier(stage *Stage) string {
	return directionallight.GongGetIdentifier(stage)
}

func (extrudegeometry *ExtrudeGeometry) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(extrudegeometry, extrudegeometry.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (extrudegeometry *ExtrudeGeometry) GongGetReferenceIdentifier(stage *Stage) string {
	return extrudegeometry.GongGetIdentifier(stage)
}

func (mesh *Mesh) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(mesh, mesh.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (mesh *Mesh) GongGetReferenceIdentifier(stage *Stage) string {
	return mesh.GongGetIdentifier(stage)
}

func (meshmaterialbasic *MeshMaterialBasic) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(meshmaterialbasic, meshmaterialbasic.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (meshmaterialbasic *MeshMaterialBasic) GongGetReferenceIdentifier(stage *Stage) string {
	return meshmaterialbasic.GongGetIdentifier(stage)
}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(meshphysicalmaterial, meshphysicalmaterial.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (meshphysicalmaterial *MeshPhysicalMaterial) GongGetReferenceIdentifier(stage *Stage) string {
	return meshphysicalmaterial.GongGetIdentifier(stage)
}

func (planegeometry *PlaneGeometry) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(planegeometry, planegeometry.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (planegeometry *PlaneGeometry) GongGetReferenceIdentifier(stage *Stage) string {
	return planegeometry.GongGetIdentifier(stage)
}

func (shape *Shape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(shape, shape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (shape *Shape) GongGetReferenceIdentifier(stage *Stage) string {
	return shape.GongGetIdentifier(stage)
}

func (spheregeometry *SphereGeometry) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(spheregeometry, spheregeometry.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spheregeometry *SphereGeometry) GongGetReferenceIdentifier(stage *Stage) string {
	return spheregeometry.GongGetIdentifier(stage)
}

func (torusgeometry *TorusGeometry) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(torusgeometry, torusgeometry.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (torusgeometry *TorusGeometry) GongGetReferenceIdentifier(stage *Stage) string {
	return torusgeometry.GongGetIdentifier(stage)
}

func (triangle *Triangle) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(triangle, triangle.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (triangle *Triangle) GongGetReferenceIdentifier(stage *Stage) string {
	return triangle.GongGetIdentifier(stage)
}

func (tubegeometry *TubeGeometry) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tubegeometry, tubegeometry.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tubegeometry *TubeGeometry) GongGetReferenceIdentifier(stage *Stage) string {
	return tubegeometry.GongGetIdentifier(stage)
}

func (vector2 *Vector2) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(vector2, vector2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (vector2 *Vector2) GongGetReferenceIdentifier(stage *Stage) string {
	return vector2.GongGetIdentifier(stage)
}

func (vector3 *Vector3) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(vector3, vector3.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (vector3 *Vector3) GongGetReferenceIdentifier(stage *Stage) string {
	return vector3.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (ambiantlight *AmbiantLight) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ambiantlight.GongGetIdentifier(stage), "AmbiantLight", ambiantlight.Name)
}

func (boxgeometry *BoxGeometry) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(boxgeometry.GongGetIdentifier(stage), "BoxGeometry", boxgeometry.Name)
}

func (buffergeometry *BufferGeometry) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(buffergeometry.GongGetIdentifier(stage), "BufferGeometry", buffergeometry.Name)
}

func (camera *Camera) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(camera.GongGetIdentifier(stage), "Camera", camera.Name)
}

func (canvas *Canvas) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(canvas.GongGetIdentifier(stage), "Canvas", canvas.Name)
}

func (curve *Curve) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(curve.GongGetIdentifier(stage), "Curve", curve.Name)
}

func (cylindergeometry *CylinderGeometry) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(cylindergeometry.GongGetIdentifier(stage), "CylinderGeometry", cylindergeometry.Name)
}

func (directionallight *DirectionalLight) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(directionallight.GongGetIdentifier(stage), "DirectionalLight", directionallight.Name)
}

func (extrudegeometry *ExtrudeGeometry) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(extrudegeometry.GongGetIdentifier(stage), "ExtrudeGeometry", extrudegeometry.Name)
}

func (mesh *Mesh) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(mesh.GongGetIdentifier(stage), "Mesh", mesh.Name)
}

func (meshmaterialbasic *MeshMaterialBasic) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(meshmaterialbasic.GongGetIdentifier(stage), "MeshMaterialBasic", meshmaterialbasic.Name)
}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(meshphysicalmaterial.GongGetIdentifier(stage), "MeshPhysicalMaterial", meshphysicalmaterial.Name)
}

func (planegeometry *PlaneGeometry) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(planegeometry.GongGetIdentifier(stage), "PlaneGeometry", planegeometry.Name)
}

func (shape *Shape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(shape.GongGetIdentifier(stage), "Shape", shape.Name)
}

func (spheregeometry *SphereGeometry) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(spheregeometry.GongGetIdentifier(stage), "SphereGeometry", spheregeometry.Name)
}

func (torusgeometry *TorusGeometry) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(torusgeometry.GongGetIdentifier(stage), "TorusGeometry", torusgeometry.Name)
}

func (triangle *Triangle) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(triangle.GongGetIdentifier(stage), "Triangle", triangle.Name)
}

func (tubegeometry *TubeGeometry) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tubegeometry.GongGetIdentifier(stage), "TubeGeometry", tubegeometry.Name)
}

func (vector2 *Vector2) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(vector2.GongGetIdentifier(stage), "Vector2", vector2.Name)
}

func (vector3 *Vector3) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(vector3.GongGetIdentifier(stage), "Vector3", vector3.Name)
}

// insertion point for unstaging
func (ambiantlight *AmbiantLight) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ambiantlight.GongGetReferenceIdentifier(stage))
}

func (boxgeometry *BoxGeometry) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(boxgeometry.GongGetReferenceIdentifier(stage))
}

func (buffergeometry *BufferGeometry) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(buffergeometry.GongGetReferenceIdentifier(stage))
}

func (camera *Camera) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(camera.GongGetReferenceIdentifier(stage))
}

func (canvas *Canvas) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(canvas.GongGetReferenceIdentifier(stage))
}

func (curve *Curve) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(curve.GongGetReferenceIdentifier(stage))
}

func (cylindergeometry *CylinderGeometry) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(cylindergeometry.GongGetReferenceIdentifier(stage))
}

func (directionallight *DirectionalLight) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(directionallight.GongGetReferenceIdentifier(stage))
}

func (extrudegeometry *ExtrudeGeometry) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(extrudegeometry.GongGetReferenceIdentifier(stage))
}

func (mesh *Mesh) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(mesh.GongGetReferenceIdentifier(stage))
}

func (meshmaterialbasic *MeshMaterialBasic) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(meshmaterialbasic.GongGetReferenceIdentifier(stage))
}

func (meshphysicalmaterial *MeshPhysicalMaterial) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(meshphysicalmaterial.GongGetReferenceIdentifier(stage))
}

func (planegeometry *PlaneGeometry) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(planegeometry.GongGetReferenceIdentifier(stage))
}

func (shape *Shape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(shape.GongGetReferenceIdentifier(stage))
}

func (spheregeometry *SphereGeometry) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(spheregeometry.GongGetReferenceIdentifier(stage))
}

func (torusgeometry *TorusGeometry) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(torusgeometry.GongGetReferenceIdentifier(stage))
}

func (triangle *Triangle) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(triangle.GongGetReferenceIdentifier(stage))
}

func (tubegeometry *TubeGeometry) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tubegeometry.GongGetReferenceIdentifier(stage))
}

func (vector2 *Vector2) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(vector2.GongGetReferenceIdentifier(stage))
}

func (vector3 *Vector3) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(vector3.GongGetReferenceIdentifier(stage))
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

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
