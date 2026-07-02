// generated code - do not edit
package models

import "time"

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !IsStagedPointerToGongstruct(stage, *element) {
		*element = zero
		modified = true
		return
	}
	return
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by AmbiantLight
func (ambiantlight *AmbiantLight) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by BoxGeometry
func (boxgeometry *BoxGeometry) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Canvas
func (canvas *Canvas) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &canvas.DirectionalLights) || modified
	modified = GongCleanSlice(stage, &canvas.Meshs) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &canvas.AmbiantLight) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Curve
func (curve *Curve) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &curve.Points) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CylinderGeometry
func (cylindergeometry *CylinderGeometry) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DirectionalLight
func (directionallight *DirectionalLight) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ExtrudeGeometry
func (extrudegeometry *ExtrudeGeometry) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &extrudegeometry.Shape) || modified
	modified = GongCleanPointer(stage, &extrudegeometry.ExtrudePath) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Mesh
func (mesh *Mesh) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &mesh.MeshMaterialBasic) || modified
	modified = GongCleanPointer(stage, &mesh.MeshPhysicalMaterial) || modified
	modified = GongCleanPointer(stage, &mesh.CylinderGeometry) || modified
	modified = GongCleanPointer(stage, &mesh.BoxGeometry) || modified
	modified = GongCleanPointer(stage, &mesh.SphereGeometry) || modified
	modified = GongCleanPointer(stage, &mesh.TorusGeometry) || modified
	modified = GongCleanPointer(stage, &mesh.PlaneGeometry) || modified
	modified = GongCleanPointer(stage, &mesh.TubeGeometry) || modified
	modified = GongCleanPointer(stage, &mesh.ExtrudeGeometry) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MeshMaterialBasic
func (meshmaterialbasic *MeshMaterialBasic) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by MeshPhysicalMaterial
func (meshphysicalmaterial *MeshPhysicalMaterial) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PlaneGeometry
func (planegeometry *PlaneGeometry) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Shape
func (shape *Shape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &shape.Points) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SphereGeometry
func (spheregeometry *SphereGeometry) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TorusGeometry
func (torusgeometry *TorusGeometry) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TubeGeometry
func (tubegeometry *TubeGeometry) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &tubegeometry.Path) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Vector2
func (vector2 *Vector2) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Vector3
func (vector3 *Vector3) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		modified = instance.GongClean(stage) || modified
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
