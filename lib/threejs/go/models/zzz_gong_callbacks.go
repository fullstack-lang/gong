// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *AmbiantLight:
		if stage.OnAfterAmbiantLightCreateCallback != nil {
			stage.OnAfterAmbiantLightCreateCallback.OnAfterCreate(stage, target)
		}
	case *BoxGeometry:
		if stage.OnAfterBoxGeometryCreateCallback != nil {
			stage.OnAfterBoxGeometryCreateCallback.OnAfterCreate(stage, target)
		}
	case *BufferGeometry:
		if stage.OnAfterBufferGeometryCreateCallback != nil {
			stage.OnAfterBufferGeometryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Camera:
		if stage.OnAfterCameraCreateCallback != nil {
			stage.OnAfterCameraCreateCallback.OnAfterCreate(stage, target)
		}
	case *Canvas:
		if stage.OnAfterCanvasCreateCallback != nil {
			stage.OnAfterCanvasCreateCallback.OnAfterCreate(stage, target)
		}
	case *Curve:
		if stage.OnAfterCurveCreateCallback != nil {
			stage.OnAfterCurveCreateCallback.OnAfterCreate(stage, target)
		}
	case *CylinderGeometry:
		if stage.OnAfterCylinderGeometryCreateCallback != nil {
			stage.OnAfterCylinderGeometryCreateCallback.OnAfterCreate(stage, target)
		}
	case *DirectionalLight:
		if stage.OnAfterDirectionalLightCreateCallback != nil {
			stage.OnAfterDirectionalLightCreateCallback.OnAfterCreate(stage, target)
		}
	case *ExtrudeGeometry:
		if stage.OnAfterExtrudeGeometryCreateCallback != nil {
			stage.OnAfterExtrudeGeometryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Mesh:
		if stage.OnAfterMeshCreateCallback != nil {
			stage.OnAfterMeshCreateCallback.OnAfterCreate(stage, target)
		}
	case *MeshMaterialBasic:
		if stage.OnAfterMeshMaterialBasicCreateCallback != nil {
			stage.OnAfterMeshMaterialBasicCreateCallback.OnAfterCreate(stage, target)
		}
	case *MeshPhysicalMaterial:
		if stage.OnAfterMeshPhysicalMaterialCreateCallback != nil {
			stage.OnAfterMeshPhysicalMaterialCreateCallback.OnAfterCreate(stage, target)
		}
	case *PlaneGeometry:
		if stage.OnAfterPlaneGeometryCreateCallback != nil {
			stage.OnAfterPlaneGeometryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Shape:
		if stage.OnAfterShapeCreateCallback != nil {
			stage.OnAfterShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *SphereGeometry:
		if stage.OnAfterSphereGeometryCreateCallback != nil {
			stage.OnAfterSphereGeometryCreateCallback.OnAfterCreate(stage, target)
		}
	case *TorusGeometry:
		if stage.OnAfterTorusGeometryCreateCallback != nil {
			stage.OnAfterTorusGeometryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Triangle:
		if stage.OnAfterTriangleCreateCallback != nil {
			stage.OnAfterTriangleCreateCallback.OnAfterCreate(stage, target)
		}
	case *TubeGeometry:
		if stage.OnAfterTubeGeometryCreateCallback != nil {
			stage.OnAfterTubeGeometryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Vector2:
		if stage.OnAfterVector2CreateCallback != nil {
			stage.OnAfterVector2CreateCallback.OnAfterCreate(stage, target)
		}
	case *Vector3:
		if stage.OnAfterVector3CreateCallback != nil {
			stage.OnAfterVector3CreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterCreateFromFront is a backward-compatible package-level forwarder.
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {
	stage.AfterCreateFromFront(instance)
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront[Type Gongstruct](old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *AmbiantLight:
		newTarget := any(new).(*AmbiantLight)
		if stage.OnAfterAmbiantLightUpdateCallback != nil {
			stage.OnAfterAmbiantLightUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *BoxGeometry:
		newTarget := any(new).(*BoxGeometry)
		if stage.OnAfterBoxGeometryUpdateCallback != nil {
			stage.OnAfterBoxGeometryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *BufferGeometry:
		newTarget := any(new).(*BufferGeometry)
		if stage.OnAfterBufferGeometryUpdateCallback != nil {
			stage.OnAfterBufferGeometryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Camera:
		newTarget := any(new).(*Camera)
		if stage.OnAfterCameraUpdateCallback != nil {
			stage.OnAfterCameraUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Canvas:
		newTarget := any(new).(*Canvas)
		if stage.OnAfterCanvasUpdateCallback != nil {
			stage.OnAfterCanvasUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Curve:
		newTarget := any(new).(*Curve)
		if stage.OnAfterCurveUpdateCallback != nil {
			stage.OnAfterCurveUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *CylinderGeometry:
		newTarget := any(new).(*CylinderGeometry)
		if stage.OnAfterCylinderGeometryUpdateCallback != nil {
			stage.OnAfterCylinderGeometryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DirectionalLight:
		newTarget := any(new).(*DirectionalLight)
		if stage.OnAfterDirectionalLightUpdateCallback != nil {
			stage.OnAfterDirectionalLightUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ExtrudeGeometry:
		newTarget := any(new).(*ExtrudeGeometry)
		if stage.OnAfterExtrudeGeometryUpdateCallback != nil {
			stage.OnAfterExtrudeGeometryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Mesh:
		newTarget := any(new).(*Mesh)
		if stage.OnAfterMeshUpdateCallback != nil {
			stage.OnAfterMeshUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MeshMaterialBasic:
		newTarget := any(new).(*MeshMaterialBasic)
		if stage.OnAfterMeshMaterialBasicUpdateCallback != nil {
			stage.OnAfterMeshMaterialBasicUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MeshPhysicalMaterial:
		newTarget := any(new).(*MeshPhysicalMaterial)
		if stage.OnAfterMeshPhysicalMaterialUpdateCallback != nil {
			stage.OnAfterMeshPhysicalMaterialUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PlaneGeometry:
		newTarget := any(new).(*PlaneGeometry)
		if stage.OnAfterPlaneGeometryUpdateCallback != nil {
			stage.OnAfterPlaneGeometryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Shape:
		newTarget := any(new).(*Shape)
		if stage.OnAfterShapeUpdateCallback != nil {
			stage.OnAfterShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SphereGeometry:
		newTarget := any(new).(*SphereGeometry)
		if stage.OnAfterSphereGeometryUpdateCallback != nil {
			stage.OnAfterSphereGeometryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TorusGeometry:
		newTarget := any(new).(*TorusGeometry)
		if stage.OnAfterTorusGeometryUpdateCallback != nil {
			stage.OnAfterTorusGeometryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Triangle:
		newTarget := any(new).(*Triangle)
		if stage.OnAfterTriangleUpdateCallback != nil {
			stage.OnAfterTriangleUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TubeGeometry:
		newTarget := any(new).(*TubeGeometry)
		if stage.OnAfterTubeGeometryUpdateCallback != nil {
			stage.OnAfterTubeGeometryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Vector2:
		newTarget := any(new).(*Vector2)
		if stage.OnAfterVector2UpdateCallback != nil {
			stage.OnAfterVector2UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Vector3:
		newTarget := any(new).(*Vector3)
		if stage.OnAfterVector3UpdateCallback != nil {
			stage.OnAfterVector3UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// OnAfterUpdateFromFront is a backward-compatible package-level forwarder.
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {
	stage.OnAfterUpdateFromFront(old, new)
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront[Type Gongstruct](staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *AmbiantLight:
		if stage.OnAfterAmbiantLightDeleteCallback != nil {
			staged := any(staged).(*AmbiantLight)
			stage.OnAfterAmbiantLightDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *BoxGeometry:
		if stage.OnAfterBoxGeometryDeleteCallback != nil {
			staged := any(staged).(*BoxGeometry)
			stage.OnAfterBoxGeometryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *BufferGeometry:
		if stage.OnAfterBufferGeometryDeleteCallback != nil {
			staged := any(staged).(*BufferGeometry)
			stage.OnAfterBufferGeometryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Camera:
		if stage.OnAfterCameraDeleteCallback != nil {
			staged := any(staged).(*Camera)
			stage.OnAfterCameraDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Canvas:
		if stage.OnAfterCanvasDeleteCallback != nil {
			staged := any(staged).(*Canvas)
			stage.OnAfterCanvasDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Curve:
		if stage.OnAfterCurveDeleteCallback != nil {
			staged := any(staged).(*Curve)
			stage.OnAfterCurveDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *CylinderGeometry:
		if stage.OnAfterCylinderGeometryDeleteCallback != nil {
			staged := any(staged).(*CylinderGeometry)
			stage.OnAfterCylinderGeometryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DirectionalLight:
		if stage.OnAfterDirectionalLightDeleteCallback != nil {
			staged := any(staged).(*DirectionalLight)
			stage.OnAfterDirectionalLightDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ExtrudeGeometry:
		if stage.OnAfterExtrudeGeometryDeleteCallback != nil {
			staged := any(staged).(*ExtrudeGeometry)
			stage.OnAfterExtrudeGeometryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Mesh:
		if stage.OnAfterMeshDeleteCallback != nil {
			staged := any(staged).(*Mesh)
			stage.OnAfterMeshDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MeshMaterialBasic:
		if stage.OnAfterMeshMaterialBasicDeleteCallback != nil {
			staged := any(staged).(*MeshMaterialBasic)
			stage.OnAfterMeshMaterialBasicDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MeshPhysicalMaterial:
		if stage.OnAfterMeshPhysicalMaterialDeleteCallback != nil {
			staged := any(staged).(*MeshPhysicalMaterial)
			stage.OnAfterMeshPhysicalMaterialDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PlaneGeometry:
		if stage.OnAfterPlaneGeometryDeleteCallback != nil {
			staged := any(staged).(*PlaneGeometry)
			stage.OnAfterPlaneGeometryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Shape:
		if stage.OnAfterShapeDeleteCallback != nil {
			staged := any(staged).(*Shape)
			stage.OnAfterShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SphereGeometry:
		if stage.OnAfterSphereGeometryDeleteCallback != nil {
			staged := any(staged).(*SphereGeometry)
			stage.OnAfterSphereGeometryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TorusGeometry:
		if stage.OnAfterTorusGeometryDeleteCallback != nil {
			staged := any(staged).(*TorusGeometry)
			stage.OnAfterTorusGeometryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Triangle:
		if stage.OnAfterTriangleDeleteCallback != nil {
			staged := any(staged).(*Triangle)
			stage.OnAfterTriangleDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TubeGeometry:
		if stage.OnAfterTubeGeometryDeleteCallback != nil {
			staged := any(staged).(*TubeGeometry)
			stage.OnAfterTubeGeometryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Vector2:
		if stage.OnAfterVector2DeleteCallback != nil {
			staged := any(staged).(*Vector2)
			stage.OnAfterVector2DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Vector3:
		if stage.OnAfterVector3DeleteCallback != nil {
			staged := any(staged).(*Vector3)
			stage.OnAfterVector3DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
