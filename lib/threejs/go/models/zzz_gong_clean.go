// generated code - do not edit
package models

import "time"

// CleanSlice is the Stage method that removes unstaged elements from a slice of pointers.
func (stage *Stage) CleanSlice[T GongstructPtr](slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if stage.IsStaged(element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// CleanPointer is the Stage method that sets the pointer to nil if the referenced element is not staged.
func (stage *Stage) CleanPointer[T GongstructPtr](element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !stage.IsStaged(*element) {
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

// Clean garbage collect unstaged instances that are referenced by BufferGeometry
func (buffergeometry *BufferGeometry) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&buffergeometry.Vertices) || modified
	modified = stage.CleanSlice(&buffergeometry.Faces) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Camera
func (camera *Camera) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Canvas
func (canvas *Canvas) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&canvas.DirectionalLights) || modified
	modified = stage.CleanSlice(&canvas.Meshs) || modified
	// insertion point per field
	modified = stage.CleanPointer(&canvas.AmbiantLight) || modified
	modified = stage.CleanPointer(&canvas.Camera) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Curve
func (curve *Curve) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&curve.Points) || modified
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
	modified = stage.CleanPointer(&extrudegeometry.Shape) || modified
	modified = stage.CleanPointer(&extrudegeometry.ExtrudePath) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Mesh
func (mesh *Mesh) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&mesh.MeshMaterialBasic) || modified
	modified = stage.CleanPointer(&mesh.MeshPhysicalMaterial) || modified
	modified = stage.CleanPointer(&mesh.CylinderGeometry) || modified
	modified = stage.CleanPointer(&mesh.BoxGeometry) || modified
	modified = stage.CleanPointer(&mesh.SphereGeometry) || modified
	modified = stage.CleanPointer(&mesh.TorusGeometry) || modified
	modified = stage.CleanPointer(&mesh.PlaneGeometry) || modified
	modified = stage.CleanPointer(&mesh.TubeGeometry) || modified
	modified = stage.CleanPointer(&mesh.ExtrudeGeometry) || modified
	modified = stage.CleanPointer(&mesh.BufferGeometry) || modified
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
	modified = stage.CleanSlice(&shape.Points) || modified
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

// Clean garbage collect unstaged instances that are referenced by Triangle
func (triangle *Triangle) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TubeGeometry
func (tubegeometry *TubeGeometry) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&tubegeometry.Path) || modified
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
